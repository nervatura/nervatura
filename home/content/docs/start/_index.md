---
title: Getting Started
type: docs
weight: 40
bookFlatSection: true
bookCollapseSection: false
bookToC: true
---

## 1. Create a new demo database

### Docker

Create a Docker container and set some options
```
$ mkdir data
$ docker run -i -t --rm --name nervatura \
  -e NT_API_KEY=DEMO_API_KEY \
  -p 5000:5000 -v "$(pwd)"/data:/data nervatura/nervatura:latest
```
In a new command window:
```
$ docker exec -i nervatura /nervatura \
  -c DatabaseCreate -k DEMO_API_KEY \
  -o "{\"database\":\"demo\",\"demo\":true}"
```
### Snap
```
$ sudo NT_API_KEY=DEMO_API_KEY \
  NT_ALIAS_DEMO="sqlite://file:/var/snap/nervatura/common/demo.db?cache=shared&mode=rwc" \
  /snap/nervatura/current/nervatura -c DatabaseCreate \
  -k DEMO_API_KEY -o "{\"database\":\"demo\",\"demo\":true}"
```

### Windows

Open a PowerShell and set the (temporary) variables:
```
$env:NT_API_KEY="DEMO_API_KEY"
$env:NT_ALIAS_DEMO="sqlite://file:///C:/ProgramData/Nervatura/data/demo.db?cache=shared&mode=rwc"
```
Launch nervatura CLI (in the same command window!), and create a database:
```
& "C:\Program Files\Nervatura\nervatura.exe" -c DatabaseCreate -k DEMO_API_KEY -o '{\"database\":\"demo\",\"demo\":true}'
```
Default Nervatura data directory: `C:/ProgramData/Nervatura`

Launch nervatura server (in the same command window or Start menu):
```
& "C:\Program Files\Nervatura\nervatura.exe"
```
Of course, the environment variables can also be set permanently (see more SystemPropertiesAdvanced.exe). In this case, the server can be started from anywhere.

Alternatively, the settings can be specified in the parameter:
```
& "C:\Program Files\Nervatura\nervatura.exe" -env C:\ProgramData\Nervatura\.env.example
```

### NPM

See [Node.js Examples](/docs/start/examples)

### Admin GUI

You can use the [**ADMIN GUI**](/docs/start/screenshot#service-admin-gui) Database section:

API-KEY: **DEMO_API_KEY**<br />
Alias name: **demo**<br />
Demo database: **true**

## 2. Login to the database: 

[**Nervatura Client**](/docs/start/screenshot#web-client)

Username: **admin**<br />
Password: **Empty password: Please change after the first login!**<br />
Database: **demo**
