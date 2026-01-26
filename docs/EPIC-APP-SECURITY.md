# EPIC Backend Services (JWT Bearer) Authentication Guide

This document provides a comprehensive guide to setting up and configuring JWT-based authentication for EPIC's FHIR API using Backend Services Authorization.

---

## Table of Contents

1. [Overview](#overview)
2. [Key Concepts](#key-concepts)
3. [Setup Process](#setup-process)
4. [Key Generation](#key-generation)
5. [JWKS Creation](#jwks-creation)
6. [EPIC App Registration](#epic-app-registration)
7. [Authentication Flow](#authentication-flow)
8. [Implementation Details](#implementation-details)
9. [Security Best Practices](#security-best-practices)
10. [Troubleshooting](#troubleshooting)

---

## Overview

EPIC's Backend Services Authorization uses **JWT Bearer Token** authentication, which is based on the [SMART on FHIR Backend Services specification](https://hl7.org/fhir/smart-app-launch/backend-services.html).

**Key Principle**: Your application signs a JWT with its **private key**, and EPIC verifies the signature using your **public key** that you registered with them.

### Architecture

```
┌─────────────┐                                    ┌──────────────┐
│  Your App   │                                    │  EPIC FHIR   │
│             │                                    │              │
│ 1. Load     │                                    │              │
│ private key │                                    │              │
│ from file   │                                    │              │
│             │                                    │              │
│ 2. Create   │                                    │              │
│ JWT claims  │                                    │              │
│             │                                    │              │
│ 3. SIGN JWT │                                    │              │
│ with your   │                                    │              │
│ private key │                                    │              │
│ (RS384)     │                                    │              │
│             │                                    │              │
│ 4. Send     │──────── Signed JWT ──────────────>│              │
│ signed JWT  │                                    │ 5. VERIFY    │
│ to token    │                                    │ signature    │
│ endpoint    │                                    │ using YOUR   │
│             │                                    │ public key   │
│             │                                    │ (registered) │
│             │                                    │              │
│             │<──────── Access Token ────────────│ 6. Issue     │
│ 7. Use      │                                    │ access token │
│ access token│                                    │              │
└─────────────┘                                    └──────────────┘
```

---

## Key Concepts

### Asymmetric Cryptography (RSA)

The authentication process uses RSA public-key cryptography:

```
Private Key (YOU) ──> Sign JWT ──> Signed JWT ──> Send to EPIC
                                                         │
                                                         ▼
Public Key (EPIC) ──> Verify Signature ──> ✅ Valid ──> Issue Token
```

**Important Rules**:
- Only the holder of the **private key** can create valid signatures
- Anyone with the **public key** can verify those signatures
- The **private key** NEVER leaves your infrastructure
- The **public key** is registered with EPIC

### Components

| Component | Location | Purpose |
|-----------|----------|---------|
| **Private Key** | Your server only | Sign JWTs |
| **Public Key** | EPIC's servers | Verify JWT signatures |
| **Signed JWT** | Sent to EPIC | Prove your identity |
| **Access Token** | Returned by EPIC | Access FHIR resources |

---

## Setup Process

### High-Level Steps

1. **Generate RSA key pair** (private + public key)
2. **Create JWKS** (JSON Web Key Set) from public key
3. **Register public key with EPIC** (via App Orchard or JWKS URL)
4. **Configure your application** with private key path
5. **Test authentication** by requesting access tokens

---

## Key Generation

### Step 1: Generate Private Key

```bash
# Generate 2048-bit RSA private key in PKCS8 format (recommended)
openssl genrsa -out private-key.pem 2048
```

**Output**: `private-key.pem` (KEEP THIS SECRET!)

### Step 2: Extract Public Key

```bash
# Extract public key from private key
openssl rsa -in private-key.pem -pubout -out public-key.pem
```

**Output**: `public-key.pem` (this will be shared with EPIC)

### Verify Your Keys

```bash
# View private key details
openssl rsa -in private-key.pem -text -noout

# View public key details
openssl rsa -pubin -in public-key.pem -text -noout
```

### Key Format Support

The EPIC provider implementation (`internal/fhir/providers/epic.go`) supports both:
- **PKCS1 format** (traditional RSA format)
- **PKCS8 format** (modern, recommended)

If you need to convert:
```bash
# Convert PKCS1 to PKCS8
openssl pkcs8 -topk8 -inform PEM -outform PEM -nocrypt \
  -in private-key.pem -out private-key-pkcs8.pem
```

---

## JWKS Creation

### What is JWKS?

**JWKS** (JSON Web Key Set) is a JSON structure containing one or more public keys in JWK (JSON Web Key) format. EPIC uses this to verify your JWT signatures.

### JWKS Structure

```json
{
  "keys": [
    {
      "kty": "RSA",              // Key type: RSA
      "kid": "my-key-1",         // Key ID: unique identifier
      "use": "sig",              // Use: signature verification
      "alg": "RS384",            // Algorithm: RS384
      "n": "<modulus>",          // RSA modulus (base64url-encoded)
      "e": "AQAB"                // RSA exponent (65537 = AQAB)
    }
  ]
}
```

### Automated JWKS Generation Script

Save this script as `create_jwks.sh`:

```bash
#!/bin/bash
# Create JWKS file for EPIC API registration
# Usage: chmod +x create_jwks.sh && ./create_jwks.sh

set -e

PUBLIC_KEY="public-key.pem"
KEY_ID="langcare-non-prod-epic-key"  # Change this to your key ID

# Check if public key exists
if [ ! -f "$PUBLIC_KEY" ]; then
    echo "Error: $PUBLIC_KEY not found!"
    echo "Generate keys first:"
    echo "  openssl genrsa -out private-key.pem 2048"
    echo "  openssl rsa -in private-key.pem -pubout -out public-key.pem"
    exit 1
fi

echo "Generating JWKS from $PUBLIC_KEY..."

# Extract modulus (n) and convert to base64url
MODULUS=$(openssl rsa -pubin -in "$PUBLIC_KEY" -noout -modulus | \
  sed 's/Modulus=//' | \
  xxd -r -p | \
  base64 -w 0 | \
  tr '+/' '-_' | \
  tr -d '=')

# Exponent is usually 65537, which is "AQAB" in base64url
EXPONENT="AQAB"

# Create JWKS JSON
cat <<EOF > jwks.json
{
  "keys": [
    {
      "kty": "RSA",
      "kid": "$KEY_ID",
      "use": "sig",
      "alg": "RS384",
      "n": "$MODULUS",
      "e": "$EXPONENT"
    }
  ]
}
EOF

echo "✅ JWKS created: jwks.json"
echo ""
echo "📋 Next Steps:"
echo "1. Upload jwks.json to EPIC App Orchard, OR"
echo "2. Host at https://yourdomain.com/.well-known/jwks.json"
echo "3. Register the JWKS URL with EPIC"
```

### Run the Script

```bash
# Make executable
chmod +x create_jwks.sh

# Generate JWKS
./create_jwks.sh
```

**Output**: `jwks.json` containing your public key in JWK format

---

## EPIC App Registration

### Option A: Upload JWKS Directly

1. **Log in to EPIC App Orchard**
   - For vendors: https://apporchard.epic.com/
   - For EPIC customers: MyApps portal

2. **Navigate to Your App**
   - Go to your app's configuration page
   - Find the "Backend Services" or "JWT Bearer" authentication section

3. **Upload JWKS**
   - Upload your `jwks.json` file
   - Or paste the JSON content directly

4. **Configure Scopes**
   ```
   system/Patient.read
   system/Observation.read
   system/MedicationRequest.read
   system/Condition.read
   system/Encounter.read
   ```

### Option B: Provide JWKS URL (Recommended for Production)

1. **Host JWKS Publicly**

   Deploy your `jwks.json` to a public URL:
   ```
   https://yourdomain.com/.well-known/jwks.json
   ```

   Example nginx configuration:
   ```nginx
   location /.well-known/jwks.json {
       alias /var/www/html/.well-known/jwks.json;
       add_header Content-Type application/json;
       add_header Cache-Control "public, max-age=3600";
   }
   ```

2. **Register URL with EPIC**
   - Enter the JWKS URL in your app configuration
   - EPIC will periodically fetch your keys from this endpoint

3. **Benefits of JWKS URL**
   - Zero-downtime key rotation
   - Register multiple keys simultaneously
   - Automatic key distribution

### Key Registration Fields

When configuring your app in EPIC, you'll need:

| Field | Value |
|-------|-------|
| **Authentication Method** | Backend Services / JWT Bearer Token |
| **Client ID** | Provided by EPIC (e.g., `abc123-client-id`) |
| **Public Key / JWKS** | Upload JWKS or provide JWKS URL |
| **Token Endpoint** | `https://fhir.epic.com/interconnect-fhir-oauth/oauth2/token` |
| **Scopes** | List of FHIR resource scopes you need |

---

## Authentication Flow

### Detailed Step-by-Step Process

#### Step 1: Load Private Key

**Your Application** (happens once at startup):
```go
// Load RSA private key from PEM file
privateKey, err := loadPrivateKey("/secrets/epic/private-key.pem")
if err != nil {
    log.Fatal("Failed to load private key:", err)
}
```

**Implementation**: See `internal/fhir/providers/epic.go`, `loadPrivateKey()` function

#### Step 2: Create JWT Claims

**Your Application** (for each token request):
```go
now := time.Now()
claims := jwt.MapClaims{
    "iss": "your-client-id",        // Issuer: Your EPIC client ID
    "sub": "your-client-id",        // Subject: Your EPIC client ID
    "aud": "https://fhir.epic.com/interconnect-fhir-oauth/oauth2/token",
    "exp": now.Add(5 * time.Minute).Unix(),  // Expiry: 5 minutes from now
    "iat": now.Unix(),                       // Issued at: now
    "jti": generateJTI(),                    // JWT ID: unique identifier
}
```

**JWT Claims Explained**:
- `iss` (issuer): Your EPIC client ID
- `sub` (subject): Your EPIC client ID (same as issuer)
- `aud` (audience): EPIC's token endpoint URL
- `exp` (expiration): When this JWT expires (max 5 minutes)
- `iat` (issued at): When this JWT was created
- `jti` (JWT ID): Unique identifier for this JWT (prevents replay attacks)

#### Step 3: Sign JWT with Private Key

**Your Application**:
```go
// Create JWT with RS384 algorithm
token := jwt.NewWithClaims(jwt.SigningMethodRS384, claims)

// Sign JWT with your private key
signedJWT, err := token.SignedString(privateKey)
if err != nil {
    return fmt.Errorf("failed to sign JWT: %w", err)
}
```

**Important**:
- Algorithm MUST be **RS384** (RSA signature with SHA-384)
- The signed JWT is a cryptographic proof of your identity
- Only you can create this signature (because only you have the private key)

#### Step 4: Request Access Token from EPIC

**Your Application**:
```go
// Prepare token request
data := url.Values{}
data.Set("grant_type", "client_credentials")
data.Set("client_assertion_type", "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
data.Set("client_assertion", signedJWT)  // <-- The signed JWT
data.Set("scope", "system/Patient.read system/Observation.read")

// Send POST request to EPIC
resp, err := http.PostForm(
    "https://fhir.epic.com/interconnect-fhir-oauth/oauth2/token",
    data,
)
```

**HTTP Request Example**:
```http
POST /oauth2/token HTTP/1.1
Host: fhir.epic.com
Content-Type: application/x-www-form-urlencoded

grant_type=client_credentials
&client_assertion_type=urn:ietf:params:oauth:client-assertion-type:jwt-bearer
&client_assertion=eyJhbGciOiJSUzM4NCIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ5b3VyLWNsaWVudC1pZCIsInN1YiI6InlvdXItY2xpZW50LWlkIiwiYXVkIjoiaHR0cHM6Ly9maGlyLmVwaWMuY29tL2ludGVyY29ubmVjdC1maGlyLW9hdXRoL29hdXRoMi90b2tlbiIsImV4cCI6MTczNzE0MTYwMCwiaWF0IjoxNzM3MTQxMzAwLCJqdGkiOiJhYmMxMjMifQ.signature
&scope=system/Patient.read%20system/Observation.read
```

#### Step 5: EPIC Verifies JWT Signature

**EPIC's Token Server**:
1. Extracts the signed JWT from `client_assertion`
2. Decodes the JWT header to get `alg` (should be RS384) and `kid` (key ID)
3. Fetches your public key from registered JWKS (using `kid` if provided)
4. Verifies the JWT signature using your public key
5. Validates JWT claims:
   - `iss` matches your registered client ID
   - `sub` matches your registered client ID
   - `aud` is EPIC's token endpoint
   - `exp` is in the future (not expired)
   - `iat` is reasonable (not too far in past/future)
   - `jti` hasn't been used before (prevents replay)

**If verification succeeds**: Continue to step 6
**If verification fails**: Return 400 Bad Request with error details

#### Step 6: EPIC Issues Access Token

**EPIC's Response** (if verification succeeds):
```json
{
  "access_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9...",
  "token_type": "Bearer",
  "expires_in": 3600,
  "scope": "system/Patient.read system/Observation.read"
}
```

**Your Application**:
```go
// Parse response
var tokenResp struct {
    AccessToken string `json:"access_token"`
    ExpiresIn   int    `json:"expires_in"`
    TokenType   string `json:"token_type"`
}
json.NewDecoder(resp.Body).Decode(&tokenResp)

// Cache token with expiry
e.tokenCache.accessToken = tokenResp.AccessToken
e.tokenCache.expiry = time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)
```

#### Step 7: Use Access Token for FHIR API Calls

**Your Application** (for all FHIR requests):
```go
req, _ := http.NewRequest("GET", "https://fhir.epic.com/interconnect-fhir-oauth/api/FHIR/R4/Patient/123", nil)
req.Header.Set("Authorization", "Bearer " + accessToken)
req.Header.Set("Accept", "application/fhir+json")

resp, err := httpClient.Do(req)
```

---

## Implementation Details

### Current Implementation

The LangCare MCP FHIR Server implements EPIC authentication in:

**File**: `internal/fhir/providers/epic.go`

**Key Features**:
- ✅ JWT creation with RS384 signing
- ✅ Automatic token refresh (5-minute buffer before expiry)
- ✅ Thread-safe token caching
- ✅ Custom HTTP transport for automatic token injection
- ✅ Support for PKCS1 and PKCS8 private key formats
- ✅ Unique JWT ID generation (prevents replay attacks)

### Configuration

**File**: `configs/config.yaml`

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
```

### Kubernetes Deployment

**Private Key Secret**:
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: fhir-epic-credentials
  namespace: healthcare
type: Opaque
stringData:
  client-id: "your-epic-client-id"
  token-url: "https://fhir.epic.com/interconnect-fhir-oauth/oauth2/token"
data:
  # Base64-encoded private key PEM
  private-key: <base64-encoded-pem>
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
        - name: EPIC_CLIENT_ID
          valueFrom:
            secretKeyRef:
              name: fhir-epic-credentials
              key: client-id
        volumeMounts:
        - name: epic-private-key
          mountPath: /secrets/epic
          readOnly: true
      volumes:
      - name: epic-private-key
        secret:
          secretName: fhir-epic-credentials
          items:
          - key: private-key
            path: private-key.pem
          defaultMode: 0400  # Read-only for owner
```

### Token Management

The implementation includes automatic token refresh:

```go
func (e *EPICProvider) getValidToken(ctx context.Context) (string, error) {
    // Check if cached token is still valid (>5 minutes remaining)
    e.tokenCache.mu.RLock()
    token := e.tokenCache.accessToken
    expiry := e.tokenCache.expiry
    e.tokenCache.mu.RUnlock()

    if token != "" && time.Until(expiry) > 5*time.Minute {
        return token, nil  // Use cached token
    }

    // Token expired or about to expire, refresh it
    e.tokenCache.mu.Lock()
    defer e.tokenCache.mu.Unlock()

    // Double-check after acquiring write lock (prevents race condition)
    if e.tokenCache.accessToken != "" && time.Until(e.tokenCache.expiry) > 5*time.Minute {
        return e.tokenCache.accessToken, nil
    }

    // Refresh token
    if err := e.refreshToken(ctx); err != nil {
        return "", err
    }

    return e.tokenCache.accessToken, nil
}
```

**Features**:
- Thread-safe using `sync.RWMutex`
- Proactive refresh (5-minute buffer)
- Double-check locking pattern (prevents race conditions)
- Automatic retry on failure

---

## Security Best Practices

### Private Key Security

#### ✅ DO:
- Store private key in Kubernetes Secret
- Mount as read-only volume (`defaultMode: 0400`)
- Use separate keys per environment (dev, staging, prod)
- Rotate keys annually or when compromised
- Use hardware security modules (HSM) for production keys
- Monitor private key access in audit logs

#### ❌ DON'T:
- Never commit private key to git
- Never log private key contents
- Never send private key over network
- Never share private key between environments
- Never store private key unencrypted

### Key Rotation

To rotate keys without downtime:

1. **Generate new key pair**:
   ```bash
   openssl genrsa -out private-key-new.pem 2048
   openssl rsa -in private-key-new.pem -pubout -out public-key-new.pem
   ```

2. **Create JWKS with BOTH keys**:
   ```json
   {
     "keys": [
       {
         "kty": "RSA",
         "kid": "key-2024-01",  // OLD key
         "use": "sig",
         "alg": "RS384",
         "n": "<old-modulus>",
         "e": "AQAB"
       },
       {
         "kty": "RSA",
         "kid": "key-2024-02",  // NEW key
         "use": "sig",
         "alg": "RS384",
         "n": "<new-modulus>",
         "e": "AQAB"
       }
     ]
   }
   ```

3. **Update JWKS in EPIC** (upload or update JWKS URL)

4. **Deploy new private key** to your application

5. **Wait 24 hours** (ensure all tokens using old key expire)

6. **Remove old key** from JWKS

### JWT Best Practices

#### ✅ Proper JWT Claims:
```go
claims := jwt.MapClaims{
    "iss": clientID,              // ✅ Always use your registered client ID
    "sub": clientID,              // ✅ Must match iss
    "aud": tokenURL,              // ✅ Must be EPIC's token endpoint
    "exp": now.Add(5*time.Minute).Unix(),  // ✅ Short expiry (max 5 min)
    "iat": now.Unix(),            // ✅ Current timestamp
    "jti": uuid.New().String(),   // ✅ Unique ID for each JWT
}
```

#### ❌ Common Mistakes:
- Using wrong `aud` (must be EPIC's token endpoint, not FHIR base URL)
- JWT expiry > 5 minutes (EPIC rejects)
- Missing or duplicate `jti` (enables replay attacks)
- Wrong algorithm (must be RS384)

### Scope Management

Follow principle of least privilege:

```yaml
scopes:
  # ✅ Good: Specific, read-only scopes
  - "system/Patient.read"
  - "system/Observation.read"

  # ⚠️ Use with caution: Write access
  - "system/MedicationRequest.write"

  # ❌ Bad: Too broad
  - "system/*.*"
```

---

## Troubleshooting

### Common Errors

#### Error: "Invalid client assertion signature"

**Cause**: EPIC cannot verify your JWT signature

**Solutions**:
1. Verify public key is registered with EPIC:
   ```bash
   # Check your public key
   openssl rsa -pubin -in public-key.pem -text -noout
   ```

2. Verify private key matches public key:
   ```bash
   # Private key modulus
   openssl rsa -in private-key.pem -noout -modulus

   # Public key modulus (should match)
   openssl rsa -pubin -in public-key.pem -noout -modulus
   ```

3. Verify JWT algorithm is RS384:
   ```bash
   # Decode JWT header (first part before first dot)
   echo "eyJhbGc..." | base64 -d
   # Should show: {"alg":"RS384","typ":"JWT"}
   ```

#### Error: "Invalid client"

**Cause**: Client ID in JWT doesn't match registered app

**Solutions**:
1. Verify `client_id` in config matches EPIC registration
2. Check both `iss` and `sub` claims match your client ID
3. Ensure client ID hasn't been revoked

#### Error: "Invalid audience"

**Cause**: `aud` claim doesn't match EPIC's token endpoint

**Solution**:
```yaml
epic:
  token_url: "https://fhir.epic.com/interconnect-fhir-oauth/oauth2/token"
  # NOT the FHIR base URL!
```

#### Error: "JWT expired"

**Cause**: JWT `exp` claim is in the past

**Solutions**:
1. Ensure system clock is synchronized (use NTP)
2. Verify JWT expiry is <= 5 minutes from `iat`
3. Check for clock skew between your server and EPIC

#### Error: "Failed to load private key"

**Cause**: Private key file not found or wrong format

**Solutions**:
1. Verify file path is correct
2. Check file permissions (must be readable)
3. Verify key format (PKCS1 or PKCS8):
   ```bash
   # Check key format
   openssl rsa -in private-key.pem -text -noout
   ```

### Debugging Tips

#### Enable Debug Logging

```yaml
logging:
  level: "debug"  # Shows JWT creation and token requests
```

#### Decode JWT Manually

```bash
# Install jq
# Decode JWT (3 parts: header.payload.signature)
JWT="eyJhbGc..."

# Decode header
echo "$JWT" | cut -d. -f1 | base64 -d | jq

# Decode payload
echo "$JWT" | cut -d. -f2 | base64 -d | jq
```

#### Test JWT Signing Locally

```bash
# Create test JWT with openssl
echo -n '{"alg":"RS384","typ":"JWT"}' | base64 | tr -d '=' | tr '+/' '-_'
echo -n '{"iss":"test","sub":"test","aud":"test","exp":9999999999,"iat":1234567890,"jti":"test"}' | base64 | tr -d '=' | tr '+/' '-_'

# Sign with private key
echo -n "header.payload" | openssl dgst -sha384 -sign private-key.pem | base64 | tr -d '=' | tr '+/' '-_'
```

### Testing Checklist

Before deploying to production:

- [ ] Private key and public key pair generated
- [ ] JWKS created from public key
- [ ] Public key registered with EPIC
- [ ] Client ID configured correctly
- [ ] Token URL configured correctly
- [ ] Scopes approved by EPIC
- [ ] Private key accessible from application
- [ ] JWT signature verification passes
- [ ] Access token successfully retrieved
- [ ] FHIR API calls work with access token
- [ ] Token refresh works automatically
- [ ] Audit logging captures all requests

---

## References

### Official Documentation

- [EPIC Backend Services Authorization](https://fhir.epic.com/Documentation?docId=oauth2&section=BackendOAuth2Guide)
- [SMART on FHIR Backend Services](https://hl7.org/fhir/smart-app-launch/backend-services.html)
- [RFC 7515: JSON Web Signature (JWS)](https://tools.ietf.org/html/rfc7515)
- [RFC 7517: JSON Web Key (JWK)](https://tools.ietf.org/html/rfc7517)
- [RFC 7519: JSON Web Token (JWT)](https://tools.ietf.org/html/rfc7519)

### Tools

- [JWT.io Debugger](https://jwt.io/) - Decode and verify JWTs
- [JWK Creator](https://mkjwk.org/) - Generate JWK from keys
- [OpenSSL](https://www.openssl.org/) - Key generation and manipulation

### Related Documentation

- [Main Security Documentation](../SECURITY.md)
- [Configuration Examples](../configs/config.example.yaml)
- [EPIC Provider Implementation](../internal/fhir/providers/epic.go)

---

## Summary

### What You Need to Know

| Component | What It Is | Where It Goes |
|-----------|------------|---------------|
| **Private Key** | Your secret key for signing JWTs | Your server only (`/secrets/epic/private-key.pem`) |
| **Public Key** | Verification key derived from private key | EPIC's servers (via JWKS) |
| **JWKS** | JSON containing your public key | EPIC App Orchard or public URL |
| **Client ID** | Your app's identifier in EPIC | Configuration file |
| **Signed JWT** | Cryptographic proof of identity | Sent to EPIC for token exchange |
| **Access Token** | Bearer token for FHIR API calls | Cached and auto-refreshed |

### Security Guarantees

✅ **Private key never leaves your infrastructure**
✅ **Only you can create valid JWTs** (cryptographic proof)
✅ **EPIC verifies your identity** using your public key
✅ **Tokens auto-refresh** before expiry
✅ **Thread-safe** token caching
✅ **Audit logging** of all PHI access

---

**Version**: 1.0
**Last Updated**: 2026-01-17
**Maintained By**: Development Team
