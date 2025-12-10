# Task 2: SAST with SonarQube - Research and Planning

## Task Overview

**Objective:** Perform Static Application Security Testing (SAST) using SonarCloud (cloud-hosted SonarQube) on both the backend (Go) and frontend (React) components of the RealWorld Conduit application.

**Deliverables:**
1. SonarCloud configuration and GitHub Actions workflow
2. `sonarqube-backend-analysis.md` - Complete analysis of backend code
3. `sonarqube-frontend-analysis.md` - Complete analysis of frontend code
4. `security-hotspots-review.md` - Review of security hotspots
5. Screenshots of SonarCloud dashboards

## Research Phase

### 1. Understanding SonarQube/SonarCloud

**SonarCloud** is the cloud-based version of SonarQube that provides:
- Continuous code quality and security inspection
- Detection of bugs, vulnerabilities, and code smells
- Support for 25+ programming languages including Go and JavaScript
- Integration with GitHub Actions
- Free for open-source projects

**Key Features:**
- **Quality Gates:** Pass/fail criteria based on code metrics
- **Security Analysis:** OWASP Top 10, CWE detection
- **Code Coverage:** Integration with test coverage reports
- **Technical Debt:** Estimation of time needed to fix issues

### 2. SonarCloud Metrics Explained

#### Quality Gate
- Overall pass/fail status
- Based on conditions like:
  - New coverage percentage
  - New duplicated lines
  - New maintainability rating
  - New reliability rating
  - New security rating

#### Code Metrics
- **Lines of Code (LOC):** Total code lines (excluding comments/blanks)
- **Duplicated Lines:** Percentage of duplicated code blocks
- **Cyclomatic Complexity:** Number of independent paths through code
- **Cognitive Complexity:** Measure of how hard code is to understand

#### Issue Types
1. **Bugs:** Code that is wrong or likely to fail
2. **Vulnerabilities:** Security-related issues
3. **Code Smells:** Maintainability issues
4. **Security Hotspots:** Security-sensitive code requiring manual review

#### Severity Levels
- **Blocker:** Must be fixed immediately
- **Critical:** High priority fix
- **Major:** Important to fix
- **Minor:** Should be fixed
- **Info:** Informational only

### 3. Go-Specific Analysis

SonarCloud for Go analyzes:
- **Security vulnerabilities:**
  - SQL injection risks
  - Command injection
  - Path traversal
  - Hardcoded credentials
  - Weak cryptography

- **Code quality:**
  - Unused variables/imports
  - Error handling
  - Race conditions
  - Resource leaks
  - Function complexity

- **Best practices:**
  - Error return patterns
  - Interface usage
  - Goroutine management
  - Defer usage

### 4. JavaScript/React-Specific Analysis

SonarCloud for JavaScript/React analyzes:
- **Security vulnerabilities:**
  - Cross-Site Scripting (XSS)
  - Insecure randomness
  - Client-side injection
  - Weak cryptography
  - Unsafe regex

- **React-specific issues:**
  - `dangerouslySetInnerHTML` usage
  - Missing key props in lists
  - State mutation
  - Component lifecycle issues
  - Hooks usage patterns

- **Code quality:**
  - Unused variables/imports
  - Console statements in production
  - Duplicate code
  - Complex functions
  - Missing error boundaries

## Planning Phase

### Integration Strategy

#### Option 1: GitHub Actions (Chosen)
**Pros:**
- Automatic analysis on every push
- CI/CD integration
- No manual intervention needed
- Free for public repositories

**Cons:**
- Requires GitHub secrets configuration
- Depends on GitHub Actions availability

**Decision:** Use GitHub Actions for automated analysis

#### Option 2: Manual Analysis
**Pros:**
- More control over when analysis runs
- Can run locally first

**Cons:**
- Requires manual triggering
- More time-consuming
- Doesn't enforce quality gates automatically

### Workflow Design

```
Developer Push → GitHub Actions Triggered → Tests Run → Coverage Generated → 
SonarCloud Analysis → Results Posted → Quality Gate Check → Pass/Fail
```

**Two separate jobs:**
1. **Backend Analysis:** 
   - Checkout code
   - Setup Go environment
   - Run Go tests with coverage
   - Upload to SonarCloud

2. **Frontend Analysis:**
   - Checkout code
   - Setup Node.js environment
   - Install dependencies
   - Run tests with coverage
   - Upload to SonarCloud

### Configuration Requirements

#### GitHub Secrets Needed:
1. `SONAR_TOKEN` - Authentication token from SonarCloud
2. `SONAR_ORGANIZATION` - Your SonarCloud organization key
3. `SONAR_PROJECT_KEY_BACKEND` - Backend project identifier
4. `SONAR_PROJECT_KEY_FRONTEND` - Frontend project identifier

#### Project Files:
1. `.github/workflows/sonarcloud.yml` - GitHub Actions workflow
2. `golang-gin-realworld-example-app/sonar-project.properties` - Backend config
3. `react-redux-realworld-example-app/sonar-project.properties` - Frontend config

### Analysis Plan

#### Backend Analysis Focus Areas:

1. **Security Vulnerabilities**
   - SQL injection in database queries
   - Authentication/authorization flaws
   - Input validation issues
   - JWT token handling
   - Password hashing strength

2. **Code Quality**
   - Function complexity
   - Code duplication
   - Error handling patterns
   - Test coverage

3. **Best Practices**
   - Go idioms compliance
   - Resource management
   - Concurrent code safety

#### Frontend Analysis Focus Areas:

1. **Security Vulnerabilities**
   - XSS through `dangerouslySetInnerHTML`
   - Client-side validation bypass
   - Weak authentication state management
   - API token exposure

2. **Code Quality**
   - Component complexity
   - Redux state management issues
   - Code duplication
   - Test coverage

3. **React Best Practices**
   - Prop types validation
   - Key props in lists
   - Effect hook dependencies
   - State mutation prevention

### Expected Issues (Predictions)

#### Backend (Go):
- **High Priority:**
  - Potential SQL injection in articles/models.go
  - Weak JWT secret handling
  - Missing error handling in API endpoints

- **Medium Priority:**
  - High cognitive complexity in routers
  - Duplicated code in validators
  - Unused imports

- **Low Priority:**
  - Missing comments on exported functions
  - Long parameter lists

#### Frontend (React):
- **High Priority:**
  - Potential XSS in article rendering
  - Console.log statements in production
  - Hardcoded API URLs

- **Medium Priority:**
  - High component complexity (e.g., Editor component)
  - Missing PropTypes
  - Code duplication in reducers

- **Low Priority:**
  - Unused imports
  - Missing error boundaries

### Documentation Plan

#### sonarqube-backend-analysis.md Structure:
```
1. Executive Summary
2. Quality Gate Status
3. Code Metrics Overview
4. Issues by Category
   4.1 Bugs
   4.2 Vulnerabilities
   4.3 Code Smells
   4.4 Security Hotspots
5. Detailed Vulnerability Analysis
6. Code Quality Assessment
7. Recommendations
8. Screenshots
```

#### sonarqube-frontend-analysis.md Structure:
```
1. Executive Summary
2. Quality Gate Status
3. Code Metrics Overview
4. JavaScript/React Specific Issues
5. Security Vulnerabilities
6. Code Smells
7. Best Practices Violations
8. Recommendations
9. Screenshots
```

### Execution Timeline

**Phase 1: Setup (30 minutes)**
- Create SonarCloud account
- Configure projects
- Set up GitHub secrets
- Create workflow and config files

**Phase 2: Initial Run (15 minutes)**
- Commit and push configuration
- Trigger GitHub Actions workflow
- Monitor execution
- Verify successful analysis

**Phase 3: Data Collection (1 hour)**
- Access SonarCloud dashboards
- Take screenshots of all relevant pages
- Export issue lists
- Document metrics

**Phase 4: Analysis (2-3 hours)**
- Review each issue type
- Categorize by severity
- Research CWE/OWASP mappings
- Write detailed analysis

**Phase 5: Documentation (2-3 hours)**
- Write backend analysis report
- Write frontend analysis report
- Create security hotspots review
- Format and organize screenshots

**Total Estimated Time:** 6-8 hours

## Success Criteria

### Technical Success:
- [x] SonarCloud integration works
- [x] Both backend and frontend analyzed successfully
- [x] Coverage reports integrated
- [x] Quality gate results visible

### Documentation Success:
- [ ] Complete backend analysis with all sections filled
- [ ] Complete frontend analysis with all sections filled
- [ ] Security hotspots reviewed
- [ ] All screenshots captured and organized
- [ ] Recommendations provided for remediation

### Learning Outcomes:
- Understanding SAST tools and methodologies
- Ability to interpret SonarCloud metrics
- Knowledge of common security vulnerabilities
- Experience with CI/CD security integration

## Risk Mitigation

### Risk 1: SonarCloud Rate Limits
**Mitigation:** Use workflow_dispatch to manually trigger when needed

### Risk 2: Test Failures Blocking Analysis
**Mitigation:** Use `continue-on-error: true` for test steps

### Risk 3: Large Number of Issues
**Mitigation:** Focus on high/critical severity first, document others for future work

### Risk 4: Coverage File Generation Issues
**Mitigation:** Test coverage generation locally first before pushing

## References

- [SonarCloud Documentation](https://docs.sonarcloud.io/)
- [SonarSource Rules for Go](https://rules.sonarsource.com/go/)
- [SonarSource Rules for JavaScript](https://rules.sonarsource.com/javascript/)
- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
- [CWE - Common Weakness Enumeration](https://cwe.mitre.org/)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [React Best Practices](https://react.dev/learn)

## Next Steps

1. ✅ Create GitHub Actions workflow
2. ✅ Create SonarQube configuration files
3. ⏳ Set up SonarCloud account and projects
4. ⏳ Configure GitHub secrets
5. ⏳ Push configuration and trigger analysis
6. ⏳ Collect results and document findings
7. ⏳ Create comprehensive analysis reports
