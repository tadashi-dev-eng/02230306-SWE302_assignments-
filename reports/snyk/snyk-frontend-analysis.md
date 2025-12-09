# Snyk Frontend Analysis Report


## Key Findings
- Critical: `form-data` (CVE-2025-7783) introduced via `superagent@3.8.3`. Snyk recommends upgrading `superagent` to `10.2.2` to use a safe `form-data` version.
- Medium: `marked@0.3.19` (multiple ReDoS advisories). Upgrade to `marked@0.6.2` or `marked@4.0.10`.
- Code (Snyk Code): No issues found by Snyk Code.
- Manual review findings:
  - `dangerouslySetInnerHTML` in `src/components/Article/index.js` rendering parsed markdown; `marked` `sanitize: true` is insufficient as a full defense.
  - Tokens stored in `localStorage` are vulnerable to XSS-based token theft.

## Evidence
- Snyk scan output exported to `reports/snyk/snyk-frontend-report.json`.
- Snyk Code output exported to `reports/snyk/snyk-code-report.json`.
- Screenshot of Snyk findings attached in the repo and also referenced below:

![Snyk Frontend Findings](/Assignment_2/assets/image%20copy%202.png)

see full output at [snyk-code-report.json](/swe302_assignments/reports/snyk/snyk-frontend-analysis.md).
