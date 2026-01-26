# Gemini Integration for LangCare MCP FHIR Server

**Connect Google Gemini to LangCare MCP FHIR Server using function calling and HTTP/SSE transport.**

## Overview

Google Gemini supports function calling (tool use) through the Gemini API. This integration runs the MCP server in HTTP mode and maps MCP tools to Gemini function declarations.

**Best for:**
- Google Cloud deployments
- Vertex AI integration
- Custom applications with Gemini API
- Multi-modal healthcare AI (text + images)

## Prerequisites

- **Google Cloud Project** with Gemini API enabled
- **Gemini API Key** or Application Default Credentials (ADC)
- **LangCare MCP FHIR Server** built (`make build`)
- **FHIR Server Credentials** (EPIC, Cerner, GCP Healthcare API)
- **Public endpoint** (Cloud Run recommended)

## Architecture

```
┌─────────────┐         ┌──────────────┐         ┌─────────────┐
│   Gemini    │  HTTPS  │  MCP Server  │  OAuth  │  FHIR API   │
│   API/SDK   │────────▶│   (HTTP/SSE) │────────▶│   (EMR)     │
│  + Functions│         │   + Auth     │         │             │
└─────────────┘         └──────────────┘         └─────────────┘
```

**Key components:**
- **Gemini Function Calling** - Maps MCP tools to Gemini functions
- **HTTP/SSE Transport** - MCP server in HTTP mode
- **Cloud Run Deployment** - Serverless, auto-scaling
- **System Instructions** - Clinical skills in Gemini system prompt

## Setup Instructions

### Step 1: Deploy MCP Server to Cloud Run

**Create HTTP-mode config:**

```yaml
# configs/config.cloudrun.yaml
fhir_server:
  provider: "gcp"  # Or "epic", "cerner"

  # For GCP Healthcare API
  gcp:
    project_id: "your-gcp-project"
    location: "us-central1"
    dataset_id: "your-dataset"
    fhir_store_id: "your-fhir-store"

  # Or for EPIC/Cerner, use their configs

transport:
  stdio: false
  http:
    enabled: true
    port: 8080

security:
  mcp_auth_tokens: "${MCP_AUTH_TOKENS}"  # Set via env var
  rate_limit:
    enabled: true
    rate: 100
    burst: 200

logging:
  scrub_phi: true
```

**Deploy to Cloud Run:**

```bash
# Build Docker image
docker build -t gcr.io/your-project/langcare-mcp-fhir .

# Push to Google Container Registry
docker push gcr.io/your-project/langcare-mcp-fhir

# Deploy to Cloud Run
gcloud run deploy langcare-mcp-fhir \
  --image gcr.io/your-project/langcare-mcp-fhir \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --port 8080 \
  --set-env-vars MCP_AUTH_TOKENS=your-secret-token \
  --min-instances 1 \
  --max-instances 10 \
  --memory 512Mi \
  --cpu 1
```

**Get Cloud Run URL:**
```bash
gcloud run services describe langcare-mcp-fhir --region us-central1 --format 'value(status.url)'
# Example: https://langcare-mcp-fhir-abc123-uc.a.run.app
```

### Step 2: Setup Gemini with Function Calling

**Install Gemini SDK:**

```bash
pip install google-generativeai
```

**Create Python integration:**

```python
# gemini_fhir_client.py
import google.generativeai as genai
import requests
import json

# Configure Gemini
genai.configure(api_key='YOUR_GEMINI_API_KEY')

# MCP Server config
MCP_SERVER_URL = "https://your-cloud-run-url.run.app/mcp"
MCP_AUTH_TOKEN = "your-secret-token"

# Define FHIR tool functions for Gemini
fhir_tools = [
    {
        "name": "fhir_read",
        "description": "Read a FHIR resource by type and ID",
        "parameters": {
            "type": "object",
            "properties": {
                "resourceType": {
                    "type": "string",
                    "description": "FHIR resource type (Patient, Observation, etc.)"
                },
                "id": {
                    "type": "string",
                    "description": "Resource ID"
                }
            },
            "required": ["resourceType", "id"]
        }
    },
    {
        "name": "fhir_search",
        "description": "Search FHIR resources with query parameters",
        "parameters": {
            "type": "object",
            "properties": {
                "resourceType": {
                    "type": "string",
                    "description": "FHIR resource type to search"
                },
                "queryParams": {
                    "type": "string",
                    "description": "Query parameters (e.g., 'name=Smith&birthdate=1980-01-01')"
                }
            },
            "required": ["resourceType"]
        }
    },
    {
        "name": "fhir_create",
        "description": "Create a new FHIR resource",
        "parameters": {
            "type": "object",
            "properties": {
                "resourceType": {
                    "type": "string",
                    "description": "FHIR resource type"
                },
                "resource": {
                    "type": "object",
                    "description": "Complete FHIR resource JSON"
                }
            },
            "required": ["resourceType", "resource"]
        }
    },
    {
        "name": "fhir_update",
        "description": "Update an existing FHIR resource",
        "parameters": {
            "type": "object",
            "properties": {
                "resourceType": {
                    "type": "string",
                    "description": "FHIR resource type"
                },
                "id": {
                    "type": "string",
                    "description": "Resource ID"
                },
                "resource": {
                    "type": "object",
                    "description": "Updated FHIR resource JSON"
                }
            },
            "required": ["resourceType", "id", "resource"]
        }
    }
]

def call_mcp_tool(tool_name, arguments):
    """Call MCP server tool"""
    headers = {
        "Authorization": f"Bearer {MCP_AUTH_TOKEN}",
        "Content-Type": "application/json"
    }
    payload = {
        "jsonrpc": "2.0",
        "id": 1,
        "method": "tools/call",
        "params": {
            "name": tool_name,
            "arguments": arguments
        }
    }
    response = requests.post(MCP_SERVER_URL, json=payload, headers=headers)
    return response.json()

def chat_with_fhir(user_message, system_instructions=None):
    """Chat with Gemini using FHIR tools"""
    model = genai.GenerativeModel(
        model_name='gemini-1.5-pro',
        tools=fhir_tools,
        system_instruction=system_instructions
    )

    chat = model.start_chat(enable_automatic_function_calling=False)
    response = chat.send_message(user_message)

    # Handle function calls
    while response.candidates[0].content.parts[0].function_call:
        function_call = response.candidates[0].content.parts[0].function_call
        tool_name = function_call.name
        arguments = dict(function_call.args)

        print(f"🔧 Calling {tool_name} with args: {arguments}")

        # Call MCP server
        result = call_mcp_tool(tool_name, arguments)

        # Send result back to Gemini
        response = chat.send_message(
            genai.protos.Content(
                parts=[genai.protos.Part(
                    function_response=genai.protos.FunctionResponse(
                        name=tool_name,
                        response={"result": result}
                    )
                )]
            )
        )

    return response.text

# System instructions with optional clinical skills
SYSTEM_INSTRUCTIONS = """
You are a FHIR-enabled medical AI assistant with access to electronic health records.

You have 4 FHIR tools:
- fhir_read: Read specific resource by type and ID
- fhir_search: Search resources with query parameters
- fhir_create: Create new resource
- fhir_update: Update existing resource

Always verify patient identity before operations. Present clinical data clearly with reference ranges for labs. Follow HIPAA best practices.
"""

# Example usage
if __name__ == "__main__":
    # Add clinical skills (optional)
    with open('../../skills/core/fhir-clinical/SKILL.md', 'r') as f:
        clinical_skills = f.read()

    full_instructions = SYSTEM_INSTRUCTIONS + "\n\n" + clinical_skills

    # Chat with FHIR access
    response = chat_with_fhir(
        "Search for patients with last name Lopez",
        system_instructions=full_instructions
    )
    print(response)
```

### Step 3: Test the Integration

```python
# test_gemini_fhir.py
from gemini_fhir_client import chat_with_fhir, SYSTEM_INSTRUCTIONS

# Test 1: Search patients
print("=== Test 1: Search Patients ===")
response = chat_with_fhir("Search for patients named Smith", SYSTEM_INSTRUCTIONS)
print(response)

# Test 2: Read patient
print("\n=== Test 2: Read Patient ===")
response = chat_with_fhir("Read patient with ID 12345", SYSTEM_INSTRUCTIONS)
print(response)

# Test 3: Chart review (with clinical skills)
print("\n=== Test 3: Chart Review ===")
with open('../../skills/core/fhir-clinical/SKILL.md', 'r') as f:
    skills = SYSTEM_INSTRUCTIONS + "\n\n" + f.read()

response = chat_with_fhir("Review chart for patient John Doe, DOB 1965-03-15", skills)
print(response)
```

## Alternative: Vertex AI Integration

For enterprise deployments, use Vertex AI instead of Gemini API:

```python
# vertex_fhir_client.py
from vertexai.preview.generative_models import GenerativeModel, Tool, FunctionDeclaration
import vertexai

# Initialize Vertex AI
vertexai.init(project="your-project-id", location="us-central1")

# Define FHIR tools as Vertex AI functions
fhir_read = FunctionDeclaration(
    name="fhir_read",
    description="Read a FHIR resource by type and ID",
    parameters={
        "type": "object",
        "properties": {
            "resourceType": {"type": "string"},
            "id": {"type": "string"}
        },
        "required": ["resourceType", "id"]
    }
)

# Define other tools (fhir_search, fhir_create, fhir_update)...

fhir_tool = Tool(function_declarations=[fhir_read, fhir_search, fhir_create, fhir_update])

# Create model with tools
model = GenerativeModel(
    "gemini-1.5-pro",
    tools=[fhir_tool],
    system_instruction=SYSTEM_INSTRUCTIONS
)

# Chat with automatic function calling
chat = model.start_chat()
response = chat.send_message("Search for patients named Smith")

# Handle function calls (similar to Gemini API example)
```

**Advantages of Vertex AI:**
- Enterprise SLAs and support
- Private network connectivity (VPC)
- Fine-tuning and model customization
- Integrated with GCP services (BigQuery, Cloud Healthcare API)
- Audit logging with Cloud Audit Logs

## Web Application Example

Create a Streamlit app with FHIR access:

```python
# app.py
import streamlit as st
from gemini_fhir_client import chat_with_fhir, SYSTEM_INSTRUCTIONS

st.title("FHIR Medical Assistant (Gemini)")

# Load clinical skills
with open('skills/core/fhir-clinical/SKILL.md', 'r') as f:
    skills = SYSTEM_INSTRUCTIONS + "\n\n" + f.read()

# Chat interface
if "messages" not in st.session_state:
    st.session_state.messages = []

for message in st.session_state.messages:
    with st.chat_message(message["role"]):
        st.markdown(message["content"])

if prompt := st.chat_input("Ask about patient records..."):
    st.session_state.messages.append({"role": "user", "content": prompt})
    with st.chat_message("user"):
        st.markdown(prompt)

    with st.chat_message("assistant"):
        with st.spinner("Accessing FHIR records..."):
            response = chat_with_fhir(prompt, skills)
        st.markdown(response)
        st.session_state.messages.append({"role": "assistant", "content": response})
```

**Run:**
```bash
streamlit run app.py
```

## Adding Clinical Skills

Clinical skills are optional enhancements that provide structured workflows.

**Method 1: Append to System Instructions**

```python
with open('skills/core/fhir-clinical/SKILL.md', 'r') as f:
    clinical_skills = f.read()

system_instructions = BASE_INSTRUCTIONS + "\n\n" + clinical_skills

model = genai.GenerativeModel(
    model_name='gemini-1.5-pro',
    tools=fhir_tools,
    system_instruction=system_instructions
)
```

**Method 2: RAG with Embeddings**

For large skills, use retrieval-augmented generation:

```python
from langchain.vectorstores import FAISS
from langchain.embeddings import VertexAIEmbeddings

# Index clinical skills
embeddings = VertexAIEmbeddings()
skill_docs = load_skill_documents('skills/core/fhir-clinical/SKILL.md')
vectorstore = FAISS.from_documents(skill_docs, embeddings)

# Retrieve relevant sections
def chat_with_rag(query):
    relevant_skills = vectorstore.similarity_search(query, k=3)
    context = "\n\n".join([doc.page_content for doc in relevant_skills])

    prompt = f"""
    Context from clinical guidelines:
    {context}

    User query: {query}
    """

    return chat_with_fhir(prompt, SYSTEM_INSTRUCTIONS)
```

## Troubleshooting

### Function Calls Not Working

**Symptom:** Gemini doesn't call FHIR tools

**Solutions:**
1. Verify tool declarations match MCP tool schemas
2. Check that `enable_automatic_function_calling=False` (for manual handling)
3. Ensure system instructions mention the tools explicitly
4. Test with explicit prompt: "Use the fhir_search tool to find patients"

### MCP Server Connection Errors

**Symptom:** "Connection refused" or timeout

**Solutions:**
1. Verify Cloud Run URL is correct: `curl https://your-url.run.app/health`
2. Check Bearer token matches config
3. Ensure Cloud Run allows unauthenticated requests (or configure IAM)
4. Review Cloud Run logs: `gcloud run logs read --service langcare-mcp-fhir`

### FHIR Authentication Errors

**Symptom:** "401 Unauthorized" from FHIR server

**Solutions:**
1. Verify FHIR credentials in config
2. Check OAuth scopes for operation
3. Test FHIR access manually with curl
4. For GCP Healthcare API, verify service account has `healthcare.fhirResources.*` permissions

### Rate Limiting

**Symptom:** "Too many requests"

**Solutions:**
1. Increase rate limit in MCP server config
2. Implement exponential backoff in client code
3. Scale Cloud Run instances: `--max-instances 20`
4. Use multiple MCP auth tokens (rate limit per token)

## Advanced Configuration

### Multi-Modal FHIR AI

Gemini supports images - use for clinical photos, radiology, pathology:

```python
import PIL.Image

model = genai.GenerativeModel('gemini-1.5-pro', tools=fhir_tools)

# Upload clinical image
image = PIL.Image.open('chest_xray.jpg')

response = model.generate_content([
    "Analyze this chest X-ray and document findings in FHIR DiagnosticReport",
    image
])

# Extract structured data and create FHIR resource
if response.candidates[0].content.parts[0].function_call:
    # Use fhir_create to save DiagnosticReport
    ...
```

### Grounding with Google Search

Combine FHIR data with medical knowledge:

```python
from vertexai.preview.generative_models import grounding

model = GenerativeModel(
    "gemini-1.5-pro",
    tools=[fhir_tool],
    tool_config=grounding.ToolConfig(
        function_calling_config=grounding.FunctionCallingConfig(
            mode=grounding.FunctionCallingConfig.Mode.AUTO
        ),
        google_search_retrieval=grounding.GoogleSearchRetrieval()
    )
)

# Gemini can now ground clinical decisions in medical literature
response = chat.send_message(
    "Patient has HbA1c of 8.5%. What are current ADA guidelines for treatment?"
)
```

### Streaming Responses

For real-time UX:

```python
model = genai.GenerativeModel('gemini-1.5-pro', tools=fhir_tools)

for chunk in model.generate_content("Review patient chart...", stream=True):
    print(chunk.text, end='')
```

## Production Deployment

### Cloud Run Configuration

**Recommended settings:**

```bash
gcloud run deploy langcare-mcp-fhir \
  --image gcr.io/project/langcare-mcp-fhir \
  --platform managed \
  --region us-central1 \
  --service-account fhir-mcp-sa@project.iam.gserviceaccount.com \
  --vpc-connector vpc-connector \
  --ingress internal-and-cloud-load-balancing \
  --min-instances 2 \
  --max-instances 50 \
  --cpu 2 \
  --memory 1Gi \
  --timeout 300 \
  --concurrency 80 \
  --set-env-vars MCP_AUTH_TOKENS=${MCP_TOKENS} \
  --set-secrets EPIC_PRIVATE_KEY=epic-key:latest
```

**Security:**
- Use service account with minimal permissions
- VPC connector for private FHIR servers
- Internal ingress (no public access, use load balancer)
- Secrets in Secret Manager (not env vars)

### Monitoring

**Enable Cloud Monitoring:**

```python
from google.cloud import monitoring_v3

client = monitoring_v3.MetricServiceClient()
project_name = f"projects/{project_id}"

# Track FHIR operations
metric = monitoring_v3.TimeSeries()
metric.metric.type = "custom.googleapis.com/fhir/operations"
metric.resource.type = "cloud_run_revision"
# ... send metrics
```

**Dashboards:**
- FHIR operation latency (p50, p95, p99)
- Error rate by resource type
- Authentication failures
- Rate limit hits
- Gemini API usage and costs

### Cost Optimization

**Gemini API costs:**
- Input tokens: ~$3.50 per 1M tokens (Gemini 1.5 Pro)
- Output tokens: ~$10.50 per 1M tokens
- Function calls count as output tokens

**Optimization strategies:**
1. Use Gemini 1.5 Flash for simple queries (cheaper)
2. Cache system instructions (reduce input tokens)
3. Batch operations where possible
4. Implement smart routing (simple queries → Flash, complex → Pro)
5. Set max_output_tokens to prevent runaway generation

**Cloud Run costs:**
- vCPU: ~$0.024/hour
- Memory: ~$0.0025/GB-hour
- Requests: ~$0.40 per million

**Cost-effective config:**
```yaml
# For low-traffic deployments
min-instances: 0  # Scale to zero when idle
cpu: 1
memory: 512Mi
```

## Security Considerations

### Gemini API Keys

**Store securely:**
```bash
# Use Secret Manager
gcloud secrets create gemini-api-key --data-file=- <<< "your-key"

# Mount in Cloud Run
gcloud run deploy ... \
  --set-secrets GEMINI_API_KEY=gemini-api-key:latest
```

**In code:**
```python
import os
genai.configure(api_key=os.environ.get('GEMINI_API_KEY'))
```

### HIPAA Compliance

For PHI, use **Vertex AI** (HIPAA-compliant) instead of Gemini API:

```python
# Vertex AI is HIPAA compliant with BAA
import vertexai
vertexai.init(project="your-project", location="us-central1")

# Gemini API is NOT HIPAA compliant
# Do not use for PHI without de-identification
```

**Compliance checklist:**
- ✅ Use Vertex AI (not Gemini API) for PHI
- ✅ Business Associate Agreement (BAA) with Google Cloud
- ✅ VPC Service Controls for data isolation
- ✅ Cloud Audit Logs enabled
- ✅ Data residency controls (region selection)
- ✅ Encryption in transit and at rest (default)

## Examples

### Example 1: Clinical Decision Support

```python
from gemini_fhir_client import chat_with_fhir

# Load clinical skills
with open('skills/core/fhir-clinical/SKILL.md', 'r') as f:
    skills = f.read()

system_prompt = f"""
You are a clinical decision support system using FHIR EHR data.

Provide evidence-based recommendations following current guidelines (USPSTF, ADA, ACC/AHA).

{skills}
"""

# Use case: Diabetes management
response = chat_with_fhir(
    "Review diabetes management for patient ID 12345. Check HbA1c, medications, and complications screening.",
    system_instructions=system_prompt
)

print(response)
```

### Example 2: Preventive Care Outreach

```python
# Batch identify patients overdue for screenings
patients = ["patient-1", "patient-2", "patient-3"]

for patient_id in patients:
    response = chat_with_fhir(
        f"What preventive care screenings is patient {patient_id} overdue for?",
        system_instructions=skills
    )

    if "overdue" in response.lower():
        print(f"📋 {patient_id}: {response}")
        # Send to outreach system
```

### Example 3: Multimodal Clinical Documentation

```python
import PIL.Image

# Document wound with photo
wound_image = PIL.Image.open('wound.jpg')

model = genai.GenerativeModel('gemini-1.5-pro', tools=fhir_tools)

response = model.generate_content([
    "Describe this wound and create a FHIR Observation resource documenting size, depth, and appearance",
    wound_image
])

# Create structured FHIR Observation
# Use fhir_create tool to save
```

## Getting Help

**Gemini/Vertex AI Issues:**
- Google AI docs: https://ai.google.dev/docs
- Vertex AI docs: https://cloud.google.com/vertex-ai/docs
- Support: https://cloud.google.com/support

**MCP Server Issues:**
- GitHub Issues: https://github.com/langcare/langcare-mcp-fhir/issues
- Security guide: [SECURITY.md](../../docs/SECURITY.md)
- Local testing: [LOCAL-TESTING.md](../../docs/LOCAL-TESTING.md)

**Skills:**
- Skills README: [skills/README.md](../../skills/README.md)
- Contributing: [CONTRIBUTING.md](../../CONTRIBUTING.md)

## Next Steps

1. **Deploy MCP server** to Cloud Run
2. **Setup Gemini/Vertex AI** with function calling
3. **Add clinical skills** to system instructions
4. **Build application** (Streamlit, web app, or API)
5. **Test thoroughly** with synthetic data
6. **Monitor and optimize** costs and performance
7. **Scale** to production with proper security

**Questions?** Open an issue or discussion on GitHub!
