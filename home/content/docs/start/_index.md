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
$env:NT_ALIAS_DEMO="sqlite://file:///C:/YOUR_DATA_PATH/demo.db"
```
The "C:/YOUR_DATA_PATH" is an existing and writable directory, e.g. `$env:NT_ALIAS_DEMO="sqlite://file:///C:/mydata/demo.db"` The SQLite database are created automatically.

Launch nervatura CLI (in the same command window!), and create a database:
```
nervatura -c DatabaseCreate -k DEMO_API_KEY -o '{\"database\":\"demo\",\"demo\":true}'
```

Launch nervatura server (in the same command window!):
```
nervatura
```
Of course, the environment variables can also be set permanently (see more SystemPropertiesAdvanced.exe). In this case, the server can be started from anywhere, not just in the session.

### Node.js

See [Examples](/docs/start/examples)

### Python

[Python gRPC packages](https://pypi.org/project/nervatura/)

See [Python sample application](https://github.com/nervatura/nervatura-fastapi)

### Admin GUI

You can use the [**ADMIN GUI**](http://localhost:5000/admin/) Database section:

API-KEY: **DEMO_API_KEY**<br />
Alias name: **demo**<br />
Demo database: **true**

## 2. Login to the database: 

[**Nervatura Client**](http://localhost:5000/client/)

Username: **admin**<br />
Password: **Empty password: Please change after the first login!**<br />
Database: **demo**
