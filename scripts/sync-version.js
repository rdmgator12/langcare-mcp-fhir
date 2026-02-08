#!/usr/bin/env node

const fs = require('fs');
const path = require('path');

const rootDir = path.join(__dirname, '..');
const version = require(path.join(rootDir, 'package.json')).version;

// Update server.json
const serverJsonPath = path.join(rootDir, 'server.json');
const serverJson = JSON.parse(fs.readFileSync(serverJsonPath, 'utf8'));
serverJson.version = version;
if (serverJson.packages && serverJson.packages[0]) {
  serverJson.packages[0].version = version;
}
fs.writeFileSync(serverJsonPath, JSON.stringify(serverJson, null, 2) + '\n');
console.log(`✓ server.json updated to ${version}`);

// Update internal/config/config.go
const configGoPath = path.join(rootDir, 'internal/config/config.go');
let configGo = fs.readFileSync(configGoPath, 'utf8');
configGo = configGo.replace(
  /c\.MCP\.Version = "[^"]+"/,
  `c.MCP.Version = "${version}"`
);
fs.writeFileSync(configGoPath, configGo);
console.log(`✓ internal/config/config.go updated to ${version}`);

console.log(`\nVersion synced to ${version}`);
