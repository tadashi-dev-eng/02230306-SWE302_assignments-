# Task 2: SAST with SonarQube - Execution Guide

## Overview
This guide provides step-by-step instructions for executing Task 2 of Assignment 2: Static Application Security Testing using SonarCloud (cloud-hosted SonarQube).

## Prerequisites Setup

### 1. SonarCloud Account Setup
1. Visit [SonarCloud](https://sonarcloud.io/)
2. Sign in with your GitHub account
3. Import your GitHub repository: `02230306-SWE302_assignments-`

### 2. Create SonarCloud Projects

#### Backend Project (Go)
1. In SonarCloud dashboard, click "Analyze new project"
2. Select your repository
3. Create project with key: `tadashi-dev-eng_02230306-SWE302_assignments-backend`
4. Choose "GitHub Actions" as analysis method
5. Copy the generated `SONAR_TOKEN`

#### Frontend Project (React)
1. Create another project in SonarCloud
2. Project key: `tadashi-dev-eng_02230306-SWE302_assignments-frontend`
3. Choose "GitHub Actions" as analysis method
4. Use the same `SONAR_TOKEN` or generate a new one

### 3. Configure GitHub Secrets

Go to your GitHub repository settings → Secrets and variables → Actions → New repository secret

Add the following secrets:

1. **SONAR_TOKEN**
   - Value: Token from SonarCloud (starts with `sqp_`)
   
2. **SONAR_ORGANIZATION**
   - Value: Your SonarCloud organization key (usually your GitHub username)
   
3. **SONAR_PROJECT_KEY_BACKEND**
   - Value: `tadashi-dev-eng_02230306-SWE302_assignments-backend`
   
4. **SONAR_PROJECT_KEY_FRONTEND**
   - Value: `tadashi-dev-eng_02230306-SWE302_assignments-frontend`

## Execution Steps

### Step 1: Verify Configuration Files

Ensure the following files exist in your repository:

```bash
.github/workflows/sonarcloud.yml
golang-gin-realworld-example-app/sonar-project.properties
react-redux-realworld-example-app/sonar-project.properties
```

### Step 2: Generate Test Coverage (Local Testing)

#### Backend Coverage:
```bash
cd golang-gin-realworld-example-app
go test ./... -coverprofile=coverage.out -covermode=atomic
go test ./... -json > test-report.json
```

#### Frontend Coverage:
```bash
cd react-redux-realworld-example-app
npm install
npm test -- --coverage --watchAll=false
```

### Step 3: Commit and Push Configuration

```bash
git add .github/workflows/sonarcloud.yml
git add golang-gin-realworld-example-app/sonar-project.properties
git add react-redux-realworld-example-app/sonar-project.properties
git commit -m "feat: Add SonarCloud SAST configuration for Task 2"
git push origin master
```

### Step 4: Trigger GitHub Actions Workflow

The workflow will automatically run on push. You can also trigger it manually:

1. Go to GitHub repository → Actions tab
2. Select "SonarCloud Analysis" workflow
3. Click "Run workflow" button
4. Select branch (master)
5. Click "Run workflow"

### Step 5: Monitor Workflow Execution

1. Watch the workflow run in GitHub Actions
2. Check for any errors in the logs
3. Wait for both jobs to complete:
   - `sonarcloud-backend`
   - `sonarcloud-frontend`

### Step 6: Access SonarCloud Results

1. Go to [SonarCloud Dashboard](https://sonarcloud.io/)
2. Navigate to your organization
3. View backend project results
4. View frontend project results

## Expected Outcomes

### Backend Analysis Results
- Quality Gate status (Pass/Fail)
- Code metrics (LOC, complexity, duplication)
- Issues breakdown (Bugs, Vulnerabilities, Code Smells)
- Security hotspots
- Code coverage percentage

### Frontend Analysis Results
- Quality Gate status
- JavaScript/React-specific issues
- Security vulnerabilities
- Code smells and technical debt
- Test coverage

## Data Collection for Reports

### For Backend Analysis (`sonarqube-backend-analysis.md`)

Collect from SonarCloud dashboard:

1. **Quality Gate**
   - Status (Pass/Fail)
   - Conditions not met
   - Screenshot: Main dashboard

2. **Metrics**
   - Lines of Code
   - Duplicated Lines (%)
   - Cyclomatic Complexity
   - Cognitive Complexity
   - Screenshot: Measures tab

3. **Issues**
   - Total bugs (count by severity)
   - Total vulnerabilities (count by severity)
   - Total code smells (count by severity)
   - Screenshot: Issues tab

4. **Security**
   - Security hotspots count
   - Security rating (A-E)
   - Vulnerability details
   - Screenshot: Security tab

5. **Maintainability**
   - Technical debt
   - Maintainability rating
   - Code smells by type
   - Screenshot: Measures tab

### For Frontend Analysis (`sonarqube-frontend-analysis.md`)

Collect similar data as backend, plus:

1. **React-Specific Issues**
   - Component complexity
   - JSX security issues
   - React anti-patterns

2. **JavaScript Issues**
   - Unused variables/imports
   - Console statements
   - Weak cryptography

## Troubleshooting

### Issue: Workflow Fails - Missing Secrets
**Solution:** Verify all GitHub secrets are correctly set

### Issue: SonarCloud Analysis Fails
**Solution:** Check `sonar-project.properties` syntax and project keys

### Issue: Coverage Not Showing
**Solution:** Ensure test commands run successfully and generate coverage files

### Issue: Quality Gate Fails
**Solution:** This is expected - document the failures in your analysis report

## Next Steps After Execution

1. Take screenshots of all relevant SonarCloud pages
2. Document findings in analysis markdown files
3. Complete security hotspots review
4. Create remediation plan for identified issues
5. Optionally: Fix some issues and re-run analysis to show improvement

## Timeline

- **Setup (30 min):** Create SonarCloud account, configure secrets
- **Execution (15 min):** Push changes, trigger workflow
- **Analysis (10 min):** Wait for workflow to complete
- **Documentation (2-3 hours):** Collect data, write analysis reports

## References

- [SonarCloud Documentation](https://docs.sonarcloud.io/)
- [SonarCloud GitHub Action](https://github.com/SonarSource/sonarcloud-github-action)
- [Go Test Coverage](https://go.dev/blog/cover)
- [React Test Coverage](https://create-react-app.dev/docs/running-tests/#coverage-reporting)
