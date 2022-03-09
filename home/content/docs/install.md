---
title: Installation
type: docs
weight: 30
bookFlatSection: true
bookCollapseSection: false
bookToC: true
---
## Installation

### **Docker** image
```
$ docker pull nervatura/nervatura:latest
```

### **Snap** package (Linux daemon)

```
$ sudo snap install --beta nervatura
```

Checking service status and last logs:

```
$ systemctl status -l snap.nervatura.nervatura.service
```

Default snap data and http.log path:  `/var/snap/nervatura/common`

### Node.js **NPM**
```
$ npm install --save nervatura
```
Add a run script to your `package.json` file like this:

`"scripts": {
  "nervatura": "./node_modules/nervatura/bin/nervatura"
}`
```
$ npm run nervatura
```

### Prebuild binaries

[Linux and Windows x64](https://github.com/nervatura/nervatura/releases)

### Other platforms
```
$ git clone https://github.com/nervatura/nervatura.git
$ cd nervatura
$ CGO_ENABLED=0 GOOS=$(OS_NAME) GOARCH=$(ARCH_NAME) \
  go build -tags "$(TAGS)" -ldflags="-w -s -X main.Version=$(VERSION)" \
  -o $(APP_NAME) main.go
```
See more: [Building Applications in GoLang](https://golangdocs.com/building-applications-in-golang)

## Configuration Options

The application uses environment variables to set configuration options. It will be read from the [.env.example](https://raw.githubusercontent.com/nervatura/nervatura/master/service/.env.example) file. Set the environment variables as needed!