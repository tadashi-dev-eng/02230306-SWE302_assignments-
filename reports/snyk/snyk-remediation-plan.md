# Snyk Remediation Plan

Status: DRAFT (Actionable plan based on Snyk backend/frontend scans and manual review)

## Summary / Goal
Fix critical and high-severity vulnerabilities in the repo so that Snyk scan shows zero critical issues and a reduced number of high issues; implement short-term mitigations for XSS and token-safety; and prepare a long-term dependency hygiene plan.

## Critical Issues (Must Fix Immediately)

### 1) `form-data` issue (CVE-2025-7783) via `superagent` (Critical)
- Severity: Critical (CVSS 9.4)
- Affected component: `form-data` v2.3.3 introduced by `superagent@3.8.3`.
- Remediation steps:
  - Upgrade `superagent` to `10.2.2` which pulls `form-data@4.0.5` to mitigate the vulnerability.
  - Replace any deprecated wrappers (superagent-promise) or migrate to native `superagent` Promise API.
- Estimated time to fix: 1-2 hours (upgrade + manual tests + run Snyk)
- Risk/Regression: Potential changes in superagent API; must test all requests and auth flows.

## High Priority Issues

### 2) `mattn/go-sqlite3` heap overflow (High)
- Severity: High
- Affected component: transitive `github.com/mattn/go-sqlite3@v1.14.15` introduced via `github.com/jinzhu/gorm`.
- Remediation steps:
  - Add explicit requirement in `go.mod`: `github.com/mattn/go-sqlite3 v1.14.18` and run `go get`/`go mod tidy`.
  - Re-run `go test ./...` and `snyk test` to validate.
- Estimated time: 30-90 minutes.

### 3) `dgrijalva/jwt-go` access bypass (High)
- Severity: High
- Affected component: `github.com/dgrijalva/jwt-go` v3.2.0
- Remediation steps:
  - Short-term: upgrade to `github.com/dgrijalva/jwt-go v4.0.0-preview1` or patch release to reduce immediate risk.
  - Preferred long-term: Migrate to the maintained fork `github.com/golang-jwt/jwt/v4`. Update imports and API usage in `users/middlewares.go` and related code.
  - Add tests around authentication to ensure no regression.
- Estimated time: 2-4 hours (migration + tests); shorter if using the v4 preview.

## Medium/Low Priority Issues

### 4) `marked` ReDoS (Medium)
- Severity: Medium
- Affected component: `marked` v0.3.19
- Remediation steps:
  - Upgrade to `marked@0.6.2` for a low-risk patch or `marked@4.0.10` and update code (breaking APIs) as required.
  - Use an HTML sanitizer (DOMPurify) before `dangerouslySetInnerHTML` to reduce XSS risk.
- Estimated time: 1-3 hours depending on chosen version and migration effort.

Other low priority items (e.g., minor package updates) will be scheduled in a dependency hygiene plan.

## Dependency Update Strategy

1. Apply high-impact security upgrades first (superagent, go-sqlite3, jwt).
2. Use explicit pinning for vulnerable transitive dependencies where necessary.
3. Upgrade packages in batches to limit risk; run unit tests and a small regression suite for each batch.
4. Consider larger migrations (GORM v1 to v2, marked v4) as medium-term projects with appropriate branching and testing.

Testing Plan After Upgrades:
- Run unit tests (`npm test` and `go test ./...`) and run frontend smoke tests.
- Re-run Snyk CLI (snyk test --json) and Snyk Code (`snyk code test --json`) and save reports.
- Run a quick ZAP baseline for the frontend to capture immediate regressions.

## Rollback and Mitigation
- If a dependency upgrade causes regressions or breaks functionality, revert the change and apply mitigation, such as:
  - Patch or explicit pin for the insecure transitive package.
  - Introduce runtime checks to minimize the vulnerable code path exposure.

  ## Implementation and Verification Steps

  1. Implement fixes and pin dependencies in code:
    - Frontend: Update `package.json` (marked, superagent) and integrate DOMPurify; update `src/agent.js` usage.
    - Backend: Update `go.mod` with pinned `github.com/mattn/go-sqlite3` and a patched `jwt-go` release (preview or migrate to `golang-jwt`).

  2. Run unit and integration tests:
    - Frontend: `npm test`
    - Backend: `go test ./...`

  3. Re-run Snyk scans and Snyk Code analysis and compare results to the baseline:
    - `snyk test --json > reports/snyk/snyk-backend-report-after.json`
    - `snyk test --json > reports/snyk/snyk-frontend-report-after.json`
    - `snyk code test --json > reports/snyk/snyk-code-report-after.json`

  4. Evidence & Documentation:
    - Include before/after Snyk reports and screenshot evidence of the Snyk dashboard improvements in `reports/snyk`.
    - Update `snyk-fixes-applied.md` with the list of changes, test results and links to PRs or commit IDs.

  5. Deployment & Monitoring:
    - Monitor the running application for regressions and add Snyk checks to CI for future PRs.

---

This remediation plan targets the immediate resolution of critical and high-risk items discovered by Snyk and manual code review. After completing these actions we will verify scans and create `snyk-fixes-applied.md` documenting changes and before/after Snyk outputs.
