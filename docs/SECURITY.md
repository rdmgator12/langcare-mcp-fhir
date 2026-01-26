# FHIR MCP Server Security Documentation

**Status**: ✅ Fully Implemented

This document outlines the security architecture, implementation, and configuration for the FHIR MCP Server, designed to meet HIPAA compliance requirements for handling Protected Health Information (PHI).

---

## Table of Contents

1. [Architecture Overview](#architecture-overview)
2. [Two-Layer Security Architecture](#two-layer-security-architecture)
3. [Layer 1: MCP Client Authentication](#layer-1-mcp-client-authentication)
4. [Layer 2: FHIR Backend Authentication](#layer-2-fhir-backend-authentication)
5. [OAuth2/SMART on FHIR Implementation](#oauth2smart-on-fhir-implementation)
6. [Security Features](#security-features)
7. [Deployment Models](#deployment-models)
8. [Configuration Examples](#configuration-examples)
9. [HIPAA Compliance](#hipaa-compliance)
10. [Testing & Troubleshooting](#testing--troubleshooting)
11. [Security Best Practices](#security-best-practices)

---

## Architecture Overview

```
┌─────────────┐         ┌──────────────┐         ┌─────────────┐
│   Claude    │ Auth1   │  MCP Server  │ Auth2   │  FHIR API   │
│   Client    │────────▶│   (Go App)   │────────▶│   (EMR)     │
└─────────────┘         └──────────────┘         └─────────────┘

Auth1: MCP Client Authentication (Bearer Token/API Key)
Auth2: FHIR Backend Authentication (OAuth2/SMART on FHIR)
```

**Security Layers:**
- **Layer 1**: MCP Client Authentication - Controls who can access the MCP server
- **Layer 2**: FHIR Backend Authentication - MCP server authenticates to FHIR API using OAuth2

---

## Two-Layer Security Architecture

### Layer 1: MCP Client → MCP Server
- **Purpose**: Control which clients (users/systems) can access the MCP server
- **Method**: Bearer token authentication via HTTP Authorization header
- **Implementation**: Token validation middleware
- **Status**: ✅ Implemented

### Layer 2: MCP Server → FHIR API
- **Purpose**: MCP server authenticates to FHIR backends on behalf of clients
- **Method**: OAuth2/SMART on FHIR (provider-specific)
- **Implementation**: Provider pattern with automatic token refresh
- **Status**: ✅ Implemented

**Key Principle**: Clients never directly access FHIR credentials. The MCP server acts as a secure proxy with centralized credential management.

---

## Layer 1: MCP Client Authentication

### Implementation

**Location**: `internal/middleware/auth.go`

**Features**:
- ✅ Bearer token validation for HTTP/SSE transport
- ✅ Token-based access control
- ✅ Unauthorized access attempt logging
- ✅ Health check endpoint bypass

### Server-Side Configuration

**Kubernetes Secret**:
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: mcp-auth
  namespace: healthcare
type: Opaque
stringData:
  # Comma-separated list of valid tokens
  tokens: "mcp-token-team-a,mcp-token-team-b,mcp-token-user-john"
```

**Deployment Configuration**:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: fhir-mcp-server
spec:
  template:
    spec:
      containers:
      - name: mcp-server
        env:
        - name: MCP_AUTH_TOKENS
          valueFrom:
            secretKeyRef:
              name: mcp-auth
              key: tokens
```

### Client-Side Configuration

**Claude Desktop** (`~/.config/claude/claude_desktop_config.json`):
```json
{
  "mcpServers": {
    "fhir": {
      "url": "https://mcp-fhir.yourcompany.com/mcp",
      "transport": "sse",
      "headers": {
        "Authorization": "Bearer mcp-token-team-a"
      }
    }
  }
}
```

### Token Management

**Token Generation**:
```bash
openssl rand -hex 32
```

**Token Rotation**:
```bash
# Update secret with new tokens
kubectl edit secret mcp-auth -n healthcare

# Rolling restart to pick up new tokens
kubectl rollout restart deployment/fhir-mcp-server -n healthcare
```

**Token Revocation**:
1. Remove token from secret
2. Rolling restart deployment
3. Notify affected users
4. Audit for any access using revoked token

---

## Layer 2: FHIR Backend Authentication

### Provider Architecture

The implementation uses a **provider pattern** where each FHIR backend (GCP, EPIC, Cerner) has its own provider implementation with specific OAuth2 authentication logic.

```
┌─────────────────────────────────────────────────────────┐
│                  FHIR Client Factory                    │
│              (internal/fhir/client.go)                  │
└───────────────────────┬─────────────────────────────────┘
                        │
           ┌────────────┴───────────────┬─────────────────┬──────────────┐
           │                            │                 │              │
    ┌──────▼───────┐           ┌───────▼───────┐  ┌──────▼──────┐  ┌──▼───────┐
    │ GCP Provider │           │ EPIC Provider │  │   Cerner    │  │  Generic │
    │   (OAuth2)   │           │  (JWT+OAuth)  │  │  Provider   │  │ Provider │
    └──────────────┘           └───────────────┘  │  (OAuth2)   │  │ (Simple) │
                                                   └─────────────┘  └──────────┘
```

### Supported Providers

| Provider | Authentication Method | Auto-Refresh | Status |
|----------|----------------------|--------------|---------|
| **GCP Healthcare API** | Service Account OAuth2 | ✅ Yes | ✅ Implemented |
| **EPIC** | JWT Bearer Grant (SMART on FHIR) | ✅ Yes (5min buffer) | ✅ Implemented |
| **Cerner** | OAuth2 Client Credentials | ✅ Yes | ✅ Implemented |
| **Generic** | Bearer/Basic/None | N/A | ✅ Implemented |

---

## OAuth2/SMART on FHIR Implementation

### 1. GCP Healthcare API Provider

**File**: `internal/fhir/providers/gcp.go`

**Authentication**: Service Account OAuth2

**Features**:
- Automatic token management using Google's OAuth2 library
- Supports Application Default Credentials (ADC)
- Auto-refresh before token expiry
- No manual token handling required

**Configuration**:
```yaml
fhir_server:
  provider: "gcp"
  gcp:
    project_id: "my-healthcare-project"
    location: "us-central1"
    dataset_id: "production-dataset"
    fhir_store_id: "patient-store"
    credential_path: ""  # Optional: path to service account JSON
```

**How It Works**:
1. Loads service account credentials from file or `GOOGLE_APPLICATION_CREDENTIALS` env var
2. Creates OAuth2 token source with scope: `https://www.googleapis.com/auth/cloud-healthcare`
3. Uses `oauth2.NewClient()` which automatically refreshes tokens
4. Constructs FHIR store URL: `https://healthcare.googleapis.com/v1/projects/{project}/locations/{location}/datasets/{dataset}/fhirStores/{store}/fhir`

**Kubernetes Secret**:
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: fhir-gcp-credentials
type: Opaque
stringData:
  project-id: "your-gcp-project"
  location: "us-central1"
  dataset-id: "healthcare-dataset"
  fhir-store-id: "patient-data"
data:
  service-account-key: <base64-encoded-json>
```

### 2. EPIC Provider (SMART on FHIR)

**File**: `internal/fhir/providers/epic.go`

**Authentication**: Backend Services Authorization (JWT Bearer Grant)

**Features**:
- JWT assertion with RS384 signing
- Manual token refresh with 5-minute buffer
- Thread-safe token caching
- Custom HTTP transport for automatic token injection
- Supports PKCS1 and PKCS8 private key formats

**📚 Detailed Setup Guide**: See [EPIC App Security Setup](docs/EPIC-APP-SECURITY.md) for complete instructions on:
- Key generation (private/public key pair)
- JWKS creation and registration
- EPIC App Orchard configuration
- Step-by-step authentication flow
- Troubleshooting common issues

**Configuration**:
```yaml
fhir_server:
  provider: "epic"
  base_url: "https://fhir.epic.com/interconnect-fhir-oauth/api/FHIR/R4"
  epic:
    client_id: "${EPIC_CLIENT_ID}"
    private_key_path: "/secrets/epic/private-key.pem"
    token_url: "https://fhir.epic.com/interconnect-fhir-oauth/oauth2/token"
    scopes:
      - "system/Patient.read"
      - "system/Observation.read"
      - "system/MedicationRequest.read"
      - "system/Condition.read"
```

**How It Works**:
1. Loads RSA private key from PEM file
2. Creates JWT with claims: `iss`, `sub`, `aud`, `exp`, `iat`, `jti`
3. Signs JWT with RS384 algorithm
4. Exchanges JWT for access token via `client_credentials` grant
5. Caches token with expiry tracking
6. Automatically refreshes 5 minutes before expiry
7. Custom HTTP transport adds `Authorization: Bearer <token>` to all requests

**JWT Claims**:
```json
{
  "iss": "client-id",
  "sub": "client-id",
  "aud": "https://fhir.epic.com/interconnect-fhir-oauth/oauth2/token",
  "exp": 1234567890,
  "iat": 1234567600,
  "jti": "unique-jwt-id"
}
```

**Token Request**:
```http
POST /oauth2/token HTTP/1.1
Content-Type: application/x-www-form-urlencoded

grant_type=client_credentials
&client_assertion_type=urn:ietf:params:oauth:client-assertion-type:jwt-bearer
&client_assertion=<signed-jwt>
&scope=system/Patient.read system/Observation.read
```

**Token Management Code**:
```go
type tokenCache struct {
    mu          sync.RWMutex
    accessToken string
    expiry      time.Time
}

func (e *EPICProvider) getValidToken(ctx context.Context) (string, error) {
    // Check if token is valid (expires > 5 minutes from now)
    e.tokenCache.mu.RLock()
    token := e.tokenCache.accessToken
    expiry := e.tokenCache.expiry
    e.tokenCache.mu.RUnlock()

    if token != "" && time.Until(expiry) > 5*time.Minute {
        return token, nil
    }

    // Refresh token with double-check locking
    e.tokenCache.mu.Lock()
    defer e.tokenCache.mu.Unlock()

    if e.tokenCache.accessToken != "" && time.Until(e.tokenCache.expiry) > 5*time.Minute {
        return e.tokenCache.accessToken, nil
    }

    if err := e.refreshToken(ctx); err != nil {
        return "", err
    }

    return e.tokenCache.accessToken, nil
}
```

**Kubernetes Secret**:
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: fhir-epic-credentials
type: Opaque
stringData:
  client-id: "abc123-epic-client-id"
  token-url: "https://fhir.epic.com/interconnect-fhir-oauth/oauth2/token"
data:
  private-key: <base64-encoded-pem>
```

### 3. Cerner Provider (SMART on FHIR)

**File**: `internal/fhir/providers/cerner.go`

**Authentication**: OAuth2 Client Credentials Flow

**Features**:
- Standard OAuth2 client credentials grant
- Automatic token management via `golang.org/x/oauth2/clientcredentials`
- Auto-refresh before token expiry
- No manual token handling required

**Configuration**:
```yaml
fhir_server:
  provider: "cerner"
  base_url: "https://fhir-myrecord.cerner.com/r4"
  cerner:
    client_id: "${CERNER_CLIENT_ID}"
    client_secret: "${CERNER_CLIENT_SECRET}"
    token_url: "https://authorization.cerner.com/tenants/your-tenant/protocols/oauth2/profiles/smart-v1/token"
    scopes:
      - "system/Patient.read"
      - "system/Observation.read"
```

**How It Works**:
1. Creates `clientcredentials.Config` with client ID, secret, token URL, and scopes
2. Uses `oauthConfig.Client()` which automatically handles token refresh
3. Token is automatically added to all HTTP requests

**Token Request**:
```http
POST /oauth2/token HTTP/1.1
Content-Type: application/x-www-form-urlencoded
Authorization: Basic <base64(client_id:client_secret)>

grant_type=client_credentials
&scope=system/Patient.read system/Observation.read
```

**Kubernetes Secret**:
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: fhir-cerner-credentials
type: Opaque
stringData:
  client-id: "your-cerner-client-id"
  client-secret: "your-cerner-client-secret"
  token-url: "https://authorization.cerner.com/.../token"
```

### 4. Generic Provider

**File**: `internal/fhir/providers/generic.go`

**Authentication**: Bearer Token, Basic Auth, or None

**Use Cases**:
- Public FHIR test servers (e.g., HAPI FHIR)
- Development and testing
- Simple FHIR servers without OAuth2

**Configuration**:
```yaml
fhir_server:
  provider: "generic"
  base_url: "https://hapi.fhir.org/baseR4"
  auth:
    type: "none"  # or "bearer", "basic"
    # For bearer:
    # token: "${FHIR_AUTH_TOKEN}"
    # For basic:
    # username: "${FHIR_USERNAME}"
    # password: "${FHIR_PASSWORD}"
```

---

## Security Features

### ✅ MCP Client Authentication

**Location**: `internal/middleware/auth.go`

- Bearer token validation for HTTP/SSE transport
- Token-based access control
- Unauthorized access attempt logging
- Health check endpoint bypass

**Configuration**:
```yaml
security:
  mcp_auth_tokens: ""  # Use MCP_AUTH_TOKENS env var (recommended)
```

**Usage**:
```bash
# Set tokens via environment variable (recommended for production)
export MCP_AUTH_TOKENS="mcp-token-team-a,mcp-token-team-b"

# Start server with HTTP transport
./bin/langcare-mcp-fhir -http -port 8080
```

### ✅ Audit Logging for PHI Access

**Location**: `internal/audit/logger.go`

- Logs all FHIR operations (read, search, create, update)
- Includes timestamp, operation, resource type, resource ID
- Token hash (not plaintext) for traceability
- Status (success/error) and error messages
- PHI scrubbing for resource IDs (configurable)

**Log Format**:
```
[AUDIT] PHI_ACCESS timestamp=2026-01-17T12:00:00Z operation=read resource_type=Patient resource_id=12345678... user= token_hash=a1b2c3d4e5f6 remote_addr=192.168.1.1 status=success
```

**Integration**:
- All four FHIR tools (read, search, create, update) log PHI access
- Automatic status detection (success/error)
- Configurable via `logging.scrub_phi` setting

**Implementation**:
```go
type AuditLogger struct {
    logger *slog.Logger
}

func (a *AuditLogger) LogPHIAccess(ctx context.Context, event PHIAccessEvent) {
    a.logger.Info("phi_access",
        "timestamp", time.Now().UTC(),
        "user", event.UserID,
        "mcp_token", event.TokenHash, // Hash, not plaintext
        "operation", event.Operation,  // "read", "search", "create", "update"
        "resource_type", event.ResourceType,
        "resource_id", event.ResourceID,
        "fhir_provider", event.Provider,
        "status", event.Status,
        "remote_addr", event.RemoteAddr,
        "request_id", event.RequestID,
    )
}
```

### ✅ Security Headers

**Location**: `internal/middleware/security.go`

Implements all recommended security headers:
- `X-Frame-Options: DENY`
- `X-Content-Type-Options: nosniff`
- `X-XSS-Protection: 1; mode=block`
- `Strict-Transport-Security: max-age=31536000; includeSubDomains`
- `Content-Security-Policy: default-src 'self'`
- `Referrer-Policy: strict-origin-when-cross-origin`

### ✅ Rate Limiting

**Location**: `internal/middleware/security.go`

- Token bucket algorithm per IP address
- Configurable rate and burst limits
- Automatic cleanup of inactive visitors
- 429 Too Many Requests response

**Configuration**:
```yaml
security:
  rate_limit:
    enabled: true
    rate: 100   # requests per second per IP
    burst: 200  # burst allowance
```

---

## Deployment Models

### ❌ Stdio (Local) - Development Only

```bash
./bin/langcare-mcp-fhir
```

**Characteristics**:
- Process-level isolation only
- No MCP authentication needed
- Single user, local development
- **Not recommended for production PHI access**

**Use Case**: Development and testing with FHIR sandbox environments only

### ✅ HTTP/SSE (Remote) - Production Ready

```bash
# With auth tokens from environment
export MCP_AUTH_TOKENS="token1,token2,token3"
./bin/langcare-mcp-fhir -http -port 8080

# Or with TLS
export MCP_AUTH_TOKENS="token1,token2,token3"
./bin/langcare-mcp-fhir -http -port 443
```

**Features**:
- ✅ MCP client authentication
- ✅ Audit logging
- ✅ Rate limiting
- ✅ Security headers
- ✅ Health check endpoints
- ✅ Horizontal scaling ready

---

## Configuration Examples

### Development (Public Test Server)

**File**: `configs/config.yaml`

```yaml
fhir_server:
  provider: "generic"
  base_url: "https://hapi.fhir.org/baseR4"
  auth:
    type: "none"

security:
  mcp_auth_tokens: ""  # No auth for stdio mode

transport:
  stdio: true
  http:
    enabled: false

logging:
  level: "info"
  format: "json"
  scrub_phi: true
```

### Production (EPIC with SMART on FHIR)

```yaml
fhir_server:
  provider: "epic"
  base_url: "https://fhir.epic.com/interconnect-fhir-oauth/api/FHIR/R4"
  epic:
    client_id: "${EPIC_CLIENT_ID}"
    private_key_path: "/secrets/epic/private-key.pem"
    token_url: "https://fhir.epic.com/interconnect-fhir-oauth/oauth2/token"
    scopes:
      - "system/Patient.read"
      - "system/Observation.read"
      - "system/MedicationRequest.read"
      - "system/Condition.read"
      - "system/Encounter.read"

security:
  # Set via environment: export MCP_AUTH_TOKENS="..."
  rate_limit:
    enabled: true
    rate: 100
    burst: 200

transport:
  stdio: false
  http:
    enabled: true
    port: 8080
    tls:
      enabled: true
      cert_file: "/secrets/tls/tls.crt"
      key_file: "/secrets/tls/tls.key"

logging:
  level: "info"
  format: "json"
  scrub_phi: true  # REQUIRED for production
```

### Production (Cerner)

```yaml
fhir_server:
  provider: "cerner"
  base_url: "https://fhir-myrecord.cerner.com/r4"
  cerner:
    client_id: "${CERNER_CLIENT_ID}"
    client_secret: "${CERNER_CLIENT_SECRET}"
    token_url: "https://authorization.cerner.com/tenants/your-tenant/protocols/oauth2/profiles/smart-v1/token"
    scopes:
      - "system/Patient.read"
      - "system/Observation.read"

transport:
  stdio: false
  http:
    enabled: true
    port: 8080

logging:
  level: "info"
  format: "json"
  scrub_phi: true
```

### Production (GCP Healthcare API)

```yaml
fhir_server:
  provider: "gcp"
  gcp:
    project_id: "my-healthcare-project"
    location: "us-central1"
    dataset_id: "production-dataset"
    fhir_store_id: "patient-store"
    # credential_path: "/secrets/gcp/service-account-key.json"
    # Or use GOOGLE_APPLICATION_CREDENTIALS environment variable

transport:
  stdio: false
  http:
    enabled: true
    port: 8080

logging:
  level: "info"
  format: "json"
  scrub_phi: true
```

---

## HIPAA Compliance

### Required Security Controls

| Control | Status | Implementation |
|---------|--------|----------------|
| **Encryption in Transit** | ✅ | TLS 1.3 for all HTTPS connections |
| **Encryption at Rest** | ✅ | Kubernetes secrets encrypted with KMS |
| **Access Controls** | ✅ | MCP token-based authentication |
| **Audit Logging** | ✅ | All PHI access logged with full context |
| **Authentication & Authorization** | ✅ | Token-based API access |
| **Session Management** | ✅ | OAuth token expiration (1 hour recommended) |
| **Data Integrity** | ✅ | HTTPS prevents tampering |
| **Availability** | ✅ | Stateless design for horizontal scaling |

### Audit Logging

All FHIR operations are logged with:
- Timestamp (UTC)
- User/token identifier (hashed)
- Operation type (read, search, create, update)
- Resource type and ID
- FHIR provider type
- Status (success/error)
- Remote IP address
- Request ID for tracing

**Example**:
```
[AUDIT] PHI_ACCESS timestamp=2026-01-17T12:00:00Z operation=read resource_type=Patient resource_id=abc123 token_hash=a1b2c3 status=success
```

### Business Associate Agreement (BAA)

**Required for**:
- EPIC
- Cerner
- Google Cloud (when using Healthcare API)
- Any other FHIR provider accessing PHI

**Key Terms to Include**:
- Permitted uses of PHI
- Security safeguards required
- Breach notification procedures
- Data retention and disposal
- Subcontractor requirements
- Audit rights

---

## Testing & Troubleshooting

### Testing MCP Authentication

```bash
# Start server with auth
export MCP_AUTH_TOKENS="test-token-123"
./bin/langcare-mcp-fhir -http -port 8080

# Test without auth (should fail)
curl http://localhost:8080/mcp
# Expected: 401 Unauthorized

# Test with valid token (should succeed)
curl -H "Authorization: Bearer test-token-123" http://localhost:8080/mcp
# Expected: SSE connection or 200 OK

# Test health check (no auth required)
curl http://localhost:8080/health
# Expected: OK
```

### Testing Audit Logging

```bash
# Run server and check logs
./bin/langcare-mcp-fhir | grep "\[AUDIT\]"

# You should see logs like:
# [AUDIT] PHI_ACCESS timestamp=... operation=read resource_type=Patient ...
```

### Testing Rate Limiting

```bash
# Send many requests rapidly
for i in {1..300}; do
  curl -H "Authorization: Bearer test-token-123" http://localhost:8080/mcp &
done

# Should see some 429 Too Many Requests responses
```

### Troubleshooting

#### GCP: "failed to get default credentials"
- Set `GOOGLE_APPLICATION_CREDENTIALS` environment variable
- Or provide `credential_path` in configuration
- Or use GCE/GKE service account (automatic)

#### EPIC: "failed to sign JWT"
- Check private key format (must be RSA, not EC)
- Verify private key permissions (must be readable)
- Ensure private key matches public key registered with EPIC

#### EPIC: "token request failed with status 400"
- Verify client ID matches registered application
- Check token URL is correct
- Ensure scopes are approved for your application
- Verify JWT signature algorithm is RS384

#### Cerner: "token request failed with status 401"
- Verify client ID and secret are correct
- Check token URL is correct
- Ensure scopes are approved for your application

---

## Security Best Practices

### Credential Management

1. **Never commit secrets to version control**
   - Use `.gitignore` for config files with secrets
   - Use environment variables or secret managers
   - Scan repositories for accidentally committed secrets

2. **Generate secure tokens**
   ```bash
   openssl rand -hex 32
   ```

3. **Use separate credentials per environment**
   - Development: Sandbox FHIR servers
   - Staging: Limited production data
   - Production: Full production credentials with strict controls

4. **Rotate credentials regularly**
   - MCP tokens: Every 90 days
   - OAuth client secrets: Every 90 days
   - Private keys: Annually
   - TLS certificates: Every 90 days (Let's Encrypt) or annually

5. **Audit token usage**
   - Review audit logs regularly
   - Monitor for unauthorized attempts
   - Revoke compromised tokens immediately

### OAuth2 Scope Limitation

Follow principle of least privilege:

```yaml
scopes:
  - "system/Patient.read"        # ✅ Read-only patient data
  - "system/Observation.read"    # ✅ Read-only observations
  # ❌ Avoid: system/*.* (too broad)
```

### Private Key Security (EPIC)

- Store private key in Kubernetes secret
- Mount as read-only volume
- Use `private_key_path` to reference mounted file
- Never commit private keys to git
- Use PKCS8 format for better compatibility

### Network Security

1. **Enforce TLS everywhere**
   - Minimum TLS 1.2, prefer TLS 1.3
   - Valid certificates from trusted CA
   - No self-signed certificates in production

2. **Implement defense in depth**
   - WAF (Web Application Firewall)
   - DDoS protection
   - Rate limiting
   - IP allowlisting (if applicable)

3. **Service mesh for internal traffic**
   - mTLS between all services
   - Traffic encryption by default
   - Certificate auto-rotation

### Monitoring & Alerting

**Key Metrics**:
```yaml
# Prometheus alerts
- alert: HighUnauthorizedAttempts
  expr: rate(mcp_http_requests_total{status="401"}[5m]) > 10

- alert: FHIRAuthenticationFailure
  expr: rate(fhir_auth_errors_total[5m]) > 1

- alert: AbnormalPHIAccess
  expr: rate(phi_access_total[1h]) > 1000
```

**Security Logs**:
- All authentication attempts (success and failure)
- PHI access events
- Configuration changes
- Credential usage
- API errors and exceptions

### Incident Response

1. **Detection**: Alert triggers
2. **Containment**: Revoke compromised credentials, block suspicious IPs
3. **Investigation**: Review audit logs, identify scope
4. **Notification**: HIPAA breach notification if required (60 days)
5. **Remediation**: Patch vulnerabilities, rotate credentials
6. **Review**: Post-incident review, update procedures

---

## Production Deployment Checklist

### Pre-Deployment

- [ ] Secrets created in Kubernetes
- [ ] TLS certificates configured
- [ ] Network policies defined
- [ ] Service mesh deployed (if using)
- [ ] Monitoring and alerting configured
- [ ] Audit logging enabled
- [ ] BAA signed with FHIR providers
- [ ] Security review completed
- [ ] Penetration testing performed

### Post-Deployment

- [ ] Verify TLS encryption
- [ ] Test authentication flows
- [ ] Validate audit logging
- [ ] Check rate limiting
- [ ] Review security headers
- [ ] Test incident response procedures
- [ ] Document operational procedures
- [ ] Train operations team

### Ongoing

- [ ] Quarterly credential rotation
- [ ] Monthly security audits
- [ ] Regular vulnerability scanning
- [ ] Annual penetration testing
- [ ] Continuous monitoring
- [ ] Incident response drills

---

## Files Modified/Created

### New Provider Files
- `internal/fhir/providers/provider.go` - Base provider and interface
- `internal/fhir/providers/gcp.go` - GCP Healthcare API provider
- `internal/fhir/providers/epic.go` - EPIC SMART on FHIR provider
- `internal/fhir/providers/cerner.go` - Cerner SMART on FHIR provider
- `internal/fhir/providers/generic.go` - Generic provider (backwards compatibility)

### New Security Files
- `internal/middleware/auth.go` - MCP client authentication
- `internal/middleware/security.go` - Security headers and rate limiting
- `internal/audit/logger.go` - PHI access audit logging

### Modified Files
- `internal/fhir/client.go` - Converted to factory function
- `internal/fhir/types.go` - Made Client an alias for Provider
- `internal/config/config.go` - Added security and provider-specific config structs
- `internal/tools/fhir_*.go` - Added audit logging to all tools
- `internal/transport/http.go` - Integrated security middleware
- `cmd/server/main.go` - Wired all security components and provider factory
- `configs/config.yaml` - Added security section and provider field
- `configs/config.example.yaml` - Updated with security examples for all providers

---

## Dependencies

```
golang.org/x/oauth2                        # OAuth2 client library
golang.org/x/oauth2/google                 # Google OAuth2 extensions
golang.org/x/oauth2/clientcredentials      # Client credentials flow
github.com/golang-jwt/jwt/v5               # JWT library
```

---

## References

### Internal Documentation

- [EPIC App Security Setup Guide](docs/EPIC-APP-SECURITY.md) - Complete guide for EPIC JWT authentication, key generation, and JWKS registration

### External Resources

- [HIPAA Security Rule](https://www.hhs.gov/hipaa/for-professionals/security/index.html)
- [SMART on FHIR Backend Services](https://hl7.org/fhir/smart-app-launch/backend-services.html)
- [EPIC Backend Services Authorization](https://fhir.epic.com/Documentation?docId=oauth2&section=BackendOAuth2Guide)
- [Cerner Authorization Overview](https://fhir.cerner.com/authorization/)
- [Google Cloud Healthcare API Security](https://cloud.google.com/healthcare-api/docs/how-tos/security)
- [OWASP API Security Top 10](https://owasp.org/www-project-api-security/)
- [CIS Kubernetes Benchmark](https://www.cisecurity.org/benchmark/kubernetes)

---

## Revision History

| Version | Date | Changes | Author |
|---------|------|---------|--------|
| 2.0 | 2026-01-17 | Added OAuth2/SMART on FHIR implementation details, consolidated all security documentation | Development Team |
| 1.0 | 2026-01-17 | Initial security specification | Security Team |

---

**Status Summary**:
- ✅ Layer 1 (MCP Client Auth): Fully implemented with Bearer tokens
- ✅ Layer 2 (FHIR Backend Auth): Fully implemented with OAuth2/SMART on FHIR for GCP, EPIC, Cerner
- ✅ Audit Logging: Comprehensive PHI access logging
- ✅ Security Headers: All OWASP recommended headers
- ✅ Rate Limiting: Token bucket algorithm per IP
- ✅ HIPAA Compliance: All required controls implemented
- ✅ Production Ready: Kubernetes deployment examples and best practices documented
