#!/usr/bin/env node

const fs = require('fs');
const path = require('path');
const os = require('os');

const binDir = path.join(__dirname, '..', 'bin');
const platform = os.platform();
const arch = os.arch();

let binaryName;

if (platform === 'darwin' && arch === 'arm64') {
  binaryName = 'langcare-mcp-fhir-darwin-arm64';
} else if (platform === 'darwin' && arch === 'x64') {
  binaryName = 'langcare-mcp-fhir-darwin-amd64';
} else if (platform === 'linux' && arch === 'x64') {
  binaryName = 'langcare-mcp-fhir-linux-amd64';
} else if (platform === 'win32' && arch === 'x64') {
  binaryName = 'langcare-mcp-fhir-windows-amd64.exe';
} else {
  console.error(`❌ Unsupported platform: ${platform} ${arch}`);
  console.error('   Supported platforms: macOS (Intel/ARM), Linux (x64), Windows (x64)');
  process.exit(1);
}

const source = path.join(binDir, binaryName);
const target = path.join(binDir, 'langcare-mcp-fhir');

if (!fs.existsSync(source)) {
  console.error(`❌ Binary not found: ${source}`);
  process.exit(1);
}

try {
  // Remove existing symlink/file
  if (fs.existsSync(target)) {
    fs.unlinkSync(target);
  }

  // Create symlink on Unix, copy on Windows
  if (platform === 'win32') {
    fs.copyFileSync(source, target);
  } else {
    fs.symlinkSync(binaryName, target, 'file');
  }

  console.log(`✓ LangCare MCP FHIR Server installed for ${platform} ${arch}`);
  console.log(`✓ Binary available at: ${target}`);
  console.log(`✓ Run with: langcare-mcp-fhir -config /path/to/config.yaml`);
} catch (err) {
  console.error(`❌ Failed to set up binary: ${err.message}`);
  process.exit(1);
}