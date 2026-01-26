# LangCare MCP FHIR Scripts

Utility scripts for development and testing.

---

## test_epic_token.go

Standalone test script to verify EPIC credentials and obtain an access token with **comprehensive FHIR scopes (60+)**.

### Purpose

Tests your EPIC authentication setup by:
1. Loading your private key
2. Creating a JWT assertion with **60+ comprehensive FHIR scopes**
3. Exchanging it for an access token
4. **Analyzing which scopes were granted vs. requested**
5. Testing the token against EPIC's FHIR API

**Key Features:**
- ✅ Requests all scopes from `config.local.epic.yaml` (60+ scopes)
- ✅ Shows detailed scope analysis (requested vs. granted)
- ✅ Identifies missing scopes that need EPIC App Orchard registration
- ✅ Tests token with live FHIR API call

**Scopes Tested:**
- Patient, Observation, Condition, MedicationRequest, Encounter
- Procedure, DiagnosticReport, AllergyIntolerance, Immunization
- DocumentReference, Binary, Composition, Media (clinical notes & documents)
- CarePlan, Goal, Appointment, Coverage, Practitioner
- And 40+ more FHIR resources!

### Usage

**Option 1: Environment Variables**

```bash
export EPIC_CLIENT_ID="your-client-id"
export EPIC_PRIVATE_KEY_PATH="/absolute/path/to/keys/private-key.pem"

go run test/test_epic_token.go
```

**Option 2: Command Line Arguments**

```bash
go run test/test_epic_token.go \
  "your-client-id" \
  "/absolute/path/to/keys/private-key.pem"
```

**Quick Test (with your actual values):**

```bash
go run test/test_epic_token.go \
  "085e800e-401a-4303-9613-0dabec0f84c5" \
  "/Users/{your-name}}/langcare/code/langcare-mcp-fhir/keys/private-key.pem"
```

### Expected Output

```
=== EPIC Token Test ===
Client ID: 085e800e-401a-4303-9613-0dabec0f84c5
Private Key: /Users/you/langcare-mcp-fhir/keys/private-key.pem
Token URL: https://fhir.epic.com/interconnect-fhir-oauth/oauth2/token
Requesting 62 scopes (comprehensive FHIR access)
Scopes: system/Patient.read, system/Patient.write, system/Observation.read, ...

Step 1: Loading private key...
✅ Private key loaded successfully

Step 2: Creating JWT assertion...
✅ JWT assertion created and signed
   JWT (first 50 chars): eyJhbGciOiJSUzM4NCIsInR5cCI6IkpXVCJ9.eyJhdWQiOi...

Step 3: Exchanging JWT for access token...

Step 4: Parsing token response...
✅ Access token obtained successfully!

=== TOKEN DETAILS ===
Token Type: Bearer
Expires In: 3600 seconds (60 minutes)

Requested Scopes: 62
Granted Scopes: 62
✅ All requested scopes were granted!

Granted Scopes:
system/Patient.read system/Patient.write system/Observation.read ...

Access Token (first 50 chars): eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6...

=== FULL ACCESS TOKEN ===
eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjEifQ.eyJpc3MiOiJ...

=== Testing Token with FHIR API ===
Calling: https://fhir.epic.com/interconnect-fhir-oauth/api/FHIR/R4/metadata
Response Status: 200 OK
✅ Token works! Successfully called FHIR API
   Resource Type: CapabilityStatement
   FHIR Version: 4.0.1

=== TEST COMPLETE ===
```

### If Some Scopes Are Missing

If you haven't registered all scopes in EPIC App Orchard, you'll see:

```
=== TOKEN DETAILS ===
Token Type: Bearer
Expires In: 3600 seconds (60 minutes)

Requested Scopes: 62
Granted Scopes: 45

⚠️  WARNING: Not all scopes were granted!
   You may need to register additional scopes in EPIC App Orchard.

Missing Scopes (17):
  - system/Binary.read
  - system/Composition.read
  - system/Composition.write
  - system/Media.read
  - system/Media.write
  ...
```

This helps you identify exactly which scopes need to be added in EPIC App Orchard!

### Understanding Scope Mismatches

**Why Some Scopes Might Not Be Granted:**

1. **Not Registered in EPIC App Orchard**
   - You must explicitly register each scope in your EPIC app configuration
   - Go to https://apporchard.epic.com/ → Your App → Authentication → Backend Services
   - Check the boxes for each resource type you need

2. **Scope Requires Special Approval**
   - Some scopes (like `*.write`) may require additional approval
   - Contact EPIC support if a scope is rejected

3. **Typo in Scope Name**
   - Scope names are case-sensitive
   - Must exactly match: `system/ResourceType.read` or `system/ResourceType.write`

4. **EPIC Doesn't Support That Resource**
   - EPIC may not support all FHIR R4 resources
   - Check EPIC's CapabilityStatement for supported resources

**How to Fix:**

1. Note which scopes are missing from the test output
2. Go to EPIC App Orchard
3. Add those scopes to your app configuration
4. Wait 5-10 minutes for changes to propagate
5. Run the test again

See [docs/EPIC-SCOPES.md](../docs/EPIC-SCOPES.md) for complete scope reference.

---

### Troubleshooting

#### Error: "failed to load private key"

```bash
# Check file exists and is readable
ls -la /path/to/keys/private-key.pem

# Verify it's a valid RSA key
openssl rsa -in /path/to/keys/private-key.pem -text -noout

# Check permissions
chmod 600 /path/to/keys/private-key.pem
```

#### Error: "Invalid client assertion signature"

**Causes**:
- Public key not registered with EPIC
- Wrong client ID
- Private/public key mismatch

**Fix**:
```bash
# Verify keys match
openssl rsa -in keys/private-key.pem -noout -modulus | openssl md5
openssl rsa -pubin -in keys/public-key.pem -noout -modulus | openssl md5
# These MD5 hashes should be identical
```

#### Error: "Invalid audience"

Token URL must exactly match EPIC's endpoint:
```
https://fhir.epic.com/interconnect-fhir-oauth/oauth2/token
```

#### Error: "JWT expired"

System clock may be out of sync:
```bash
# macOS
sudo ntpdate -u time.apple.com

# Linux
sudo ntpdate -u pool.ntp.org
```

---

## create_jwks.sh

Creates a JWKS (JSON Web Key Set) from your public key for EPIC registration.

### Usage

```bash
cd scripts
./create_jwks.sh
```

This creates `keys/jwks.json` which you upload to EPIC App Orchard.

See [EPIC-APP-SECURITY.md](../docs/EPIC-APP-SECURITY.md) for details.

---

## Building a Test Binary

To create a standalone test binary:

```bash
# Build the test program
go build -o bin/test-epic-token scripts/test_epic_token.go

# Run it
./bin/test-epic-token "your-client-id" "/path/to/private-key.pem"
```

---

## Using in CI/CD

```bash
#!/bin/bash
# ci-test-epic.sh

set -e

# Load credentials from secrets
export EPIC_CLIENT_ID="${CI_EPIC_CLIENT_ID}"
export EPIC_PRIVATE_KEY_PATH="/tmp/private-key.pem"

# Write private key from secret
echo "${CI_EPIC_PRIVATE_KEY}" > /tmp/private-key.pem
chmod 600 /tmp/private-key.pem

# Test connection
go run scripts/test_epic_token.go

# Cleanup
rm -f /tmp/private-key.pem
```

---

## Additional Resources

- [EPIC Authentication Guide](../docs/EPIC-APP-SECURITY.md)
- [Local Testing Guide](../docs/LOCAL-TESTING.md)
- [EPIC OAuth Documentation](https://fhir.epic.com/Documentation?docId=oauth2)
