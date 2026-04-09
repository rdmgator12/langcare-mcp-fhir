#!/bin/bash
# This script creates a JWKS file for the OpenEMR FHIR API.
#
# OpenEMR uses SMART on FHIR Backend Services authentication with
# private_key_jwt and RS384. The JWKS produced by this script must be
# pasted into the "JSON Web Key Set" field when registering a
# Confidential API Client in OpenEMR (Administration -> System ->
# API Clients -> Register New API Client).
#
# Generate private key (4096-bit recommended by OpenEMR docs)
# cd keys/openemr
# openssl genrsa -out private.key 4096
# Extract public key
# openssl rsa -in private.key -pubout -out public.pem
#
# Then run:
# chmod +x create_jwks_openemr.sh && ./create_jwks_openemr.sh

# Extract modulus and exponent from public key
PUBLIC_KEY="../keys/openemr/public.pem"
KEY_ID="openemr-key-001"

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
cat <<EOF > ../keys/openemr/jwks.json
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

echo "JWKS created in ../keys/openemr/jwks.json"
echo "Key ID (kid): $KEY_ID"
echo
echo "Next steps:"
echo "  1. Copy the contents of ../keys/openemr/jwks.json"
echo "  2. In OpenEMR: Administration -> System -> API Clients -> Register New API Client"
echo "       Client Type:           Confidential"
echo "       Grant Types:           Client Credentials"
echo "       Authentication Method: private_key_jwt"
echo "       JSON Web Key Set:      paste the JWKS"
echo "  3. Save and Enable the client, then copy the generated Client ID"
echo "     into your config.openemr.yaml as fhir_server.openemr.client_id"
echo "  4. Ensure fhir_server.openemr.key_id matches: $KEY_ID"
