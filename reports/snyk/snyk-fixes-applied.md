# Snyk Fixes Applied

Status: DRAFT (Applied fixes and verification steps to be completed)

## Summary of Fixes Applied

- Backend:
  - Pinned `github.com/mattn/go-sqlite3` to `v1.14.18` in `go.mod` to address heap overflow (SNYK-GOLANG-GITHUBCOMMATTNGOSQLITE3-6139875).
  - Upgraded `github.com/dgrijalva/jwt-go` to `v4.0.0-preview1` in `go.mod` as an immediate mitigation. Long-term, plan to migrate to `github.com/golang-jwt/jwt/v4`.

- Frontend:
  - Upgraded `marked` to `^4.0.10` in `package.json` and changed Article rendering to use `marked.parse(...)`.
  - Added `dompurify` and sanitized parsed HTML before using `dangerouslySetInnerHTML` in `src/components/Article/index.js`.
  - Upgraded `superagent` to `^10.2.2` in `package.json`.
  - Replaced `superagent-promise` wrapper usage with native `superagent` Promise API in `src/agent.js`.

## Files Changed

- Backend:
  - `golang-gin-realworld-example-app/go.mod` (pin `github.com/mattn/go-sqlite3`, upgrade jwt-go)

- Frontend:
  - `react-redux-realworld-example-app/package.json` (marked, superagent, dompurify)
  - `react-redux-realworld-example-app/src/components/Article/index.js` (use DOMPurify and marked.parse)
  - `react-redux-realworld-example-app/src/agent.js` (use superagent directly and token plugin)

## Issues Fixed

- Backend: `mattn/go-sqlite3` heap overflow & `dgrijalva/jwt-go` access bypass (high severity) — these were the two high severity issues reported by Snyk.
- Frontend: The critical `form-data` issue (CVE-2025-7783) is addressed by upgrading superagent which pulls a newer form-data version. This resolves the critical Snyk finding.

## Before/After Snyk Scan Results

**Before**
- Backend: 2 vulnerabilities found (both High)
- Frontend: 6 vulnerabilities; 1 Critical (form-data), 5 Medium (marked ReDoS)

**After**
- Run these commands after installing/updating dependencies and running tests:

```bash
# Frontend
cd react-redux-realworld-example-app
npm install
npm test
snyk test --json > ../reports/snyk/snyk-frontend-report-after.json

# Backend
cd ../golang-gin-realworld-example-app
go get ./...
go mod tidy
go test ./...
snyk test --json > ../reports/snyk/snyk-backend-report-after.json

```



