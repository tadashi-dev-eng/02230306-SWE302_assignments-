# Testing analysis for golang-gin-realworld-example-app

- `common` — has tests: `common/unit_test.go`
- `users` — has tests: `users/unit_test.go`
- `articles` — no tests (no `*_test.go` found).
- Other directories (front-end `react-redux-realworld-example-app`) are not Go packages, so they do not have Go unit tests in this run.

Run used: `go test ./... -v`

## Test run results

- `realworld-backend` — no test files
- `realworld-backend/articles` — no test files
- `realworld-backend/common` — tests executed; `TestNewValidatorError` failed (panic due to validator tag), others passed or logged expected behaviors.
- `realworld-backend/users` — tests executed; `TestWithoutAuth` failed (assertion mismatches) along with some logged DB issues.

## Observed failures and root causes

1) `common.TestNewValidatorError`
   - Symptom: Run-time panic: "Undefined validation function 'exists' on field 'Username'", and HTTP 500 responses instead of expected codes.
   - Why: The test struct uses `binding:"exists,..."`. The project does not register a custom validator function named `exists` with the go-playground validator used by Gin. The `exists` tag is not a built-in validator tag and therefore unknown.
   - Suggested fixes:
     - Replace `exists` with `required` when the test is only asserting field presence.
     - Or add a custom validator registration (e.g., `validate.RegisterValidation("exists", <fn>)`) at application initialization so the tag is recognized by Gin's validator (v10).
     - Verify and rationalize validator versions in `go.mod` — avoid mixing `gopkg.in/go-playground/validator.v8` and `github.com/go-playground/validator/v10` unless intended. Prefer one and ensure registration/usage is compatible with the version.

2) `users.TestWithoutAuth` failing case ("test database pk error for user update")
   - Symptom: Test expected errors indicating invalid email/username format but got `required` errors for both fields. The test assertion fails due to mismatch.
   - Why: The test uses a body that does not include `user` fields or contains malformed JSON (`{"password": "password321"}}`, extra `}`), so the request gets parsed into a structure missing `username`/`email`, yielding `required` errors rather than format (email/alphanum) errors. Also one of the test intentions is to trigger a DB pk error, but the parser-level errors occur before it reaches the DB layer.
   - Suggested fixes:
     - Fix JSON body (remove trailing `}`) and include the fields required for the case to validate format (not presence) if testing specific validations.
     - If the goal is to test DB errors, ensure the request is well-formed and database is in the desired state before running the test so the database layer is reached and the appropriate error occurs.
     - Adjust expected assertions to match actual behavior in the presence of missing fields.

3) `common.TestNewError` — AutoMigrate log for `NotExist` model
   - Symptom: Logs show SQLite error `near ")": syntax error` after `CREATE TABLE "not_exists" ( )` (empty columns). The test subsequently reports `no such table: not_exists` during the query but the test passes.
   - Why: The `NotExist` struct in this test has an unexported field (`heheda` lowercased), which GORM ignores when creating table columns. This leads GORM to attempt to create a table with no columns, and SQLite rejects that SQL.
   - Suggested fixes:
     - Export fields used for AutoMigrate (e.g., `Heheda`), or use an entirely empty struct (and skip AutoMigrate), or avoid AutoMigrate on empty models.
     - If the intent is to assert `no such table`, then the test should not call `AutoMigrate` on the struct.

4) General environment and DB related observations
   - Some test cases intentionally chmode database file to `0000` or call `TestDBFree` to assert behavior; the run logs show `unable to open database file: permission denied`, which is expected and part of the test.
   - The repository includes both `gopkg.in/go-playground/validator.v8` in `go.mod` and `github.com/go-playground/validator/v10` (indirect). Ensure compatibility and avoid duplicate versions causing inconsistent expectations.

## Actionable fixes & next steps
1. Replace `exists` tag with `required` in `common/unit_test.go` struct to avoid panic and match the current validator set up — OR register an `exists` custom validator in initialization (prefer `required` unless there are other parts of the app relying on `exists`).
2. In `users/unit_test.go`:
   - Fix malformed JSON in test case (`{"password": "password321"}}` -> `{"password": "password321"}`)
   - Ensure the request includes proper `user` wrapper and fields when testing format validations to trigger `email`/`alphanum` errors — or adjust expected error messages to `required` when not present.
   - If intention is to test DB PK errors, use a valid request body and carefully pre-set the DB to the state that causes the pk error.
3. Modify `common.TestNewError` to use exported fields in `NotExist` model or avoid `AutoMigrate` call if the test expects 'no such table'.
4. Clean up validator versioning and registration:
   - Confirm which validator version the project should use (v8 or v10). Update `go.mod` and code references accordingly.
   - If v10 is required (Gin uses it), register custom validators used in tests (like `exists`) using `validate.RegisterValidation(...)` with proper function signatures.
5. Add tests for `articles` package (no tests present) to satisfy assignment goals (per `ASSIGNMENT_1.md` there should be tests added for `articles/`).

## Quick checklist for test repairs
- [ ] Replace `exists` with `required` in `common/unit_test.go` (or register custom validator).
- [ ] Fix malformed JSON bodies and wrap fields appropriately in `users/unit_test.go` test cases.
- [ ] Export fields in test-only model for AutoMigrate or drop AutoMigrate call for empty struct.
- [ ] Decide and enforce a validator version in `go.mod` and code (v10 recommended).
- [ ] Add tests for `articles`.
