#!/bin/bash
# This script creates a JWKS file for the EPIC API.

# Generate private key (PKCS8 format recommended)
# cd keys
# openssl genrsa -out private-key.pem 2048
# Extract public key
# openssl rsa -in private-key.pem -pubout -out public-key.pem

# Convert to JWK format (EPIC accepts this)
# chmod +x create_jwks.sh && ./create_jwks.sh

# Extract modulus and exponent from public key
PUBLIC_KEY="../keys/public-key.pem"
KEY_ID="langcare-non-prod-epic-key"

# Get modulus (n)
MODULUS=$(openssl rsa -pubin -in "$PUBLIC_KEY" -noout -modulus | \
  sed 's/Modulus=//' | \
  xxd -r -p | \
  base64 -w 0 | \
  tr '+/' '-_' | \
  tr -d '=')

# Get exponent (e) - usually 65537 (AQAB in base64url)
EXPONENT="AQAB"

# Create JWKS JSON
cat <<EOF > ../keys/jwks.json 
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

echo "JWKS created in ../keys/jwks.json"

## 📋 Register with EPIC
## Once deployed, provide EPIC with:
## JWKS URL: https://langcareai.com/.well-known/jwks.json