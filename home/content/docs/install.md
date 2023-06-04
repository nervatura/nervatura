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
docker pull nervatura/nervatura:latest
```

Example:
```
docker run -i -t --rm --name nervatura --env-file YOUR_ENV_FILE_PATH/.env -p 5000:5000 -p 9200:9200 -v YOUR_DATA_PATH:/data nervatura/nervatura:latest
```

### **Snap** package (Linux daemon)

```
sudo snap install nervatura
```

Checking service status and last logs:

```
systemctl status -l snap.nervatura.nervatura.service
```

Default snap data and http log path:  `/var/snap/nervatura/common`

The nervatura service starts by default. More examples:

- `sudo systemctl` `stop` or `start` or `restart` `snap.nervatura.nervatura.service`</br>
  `sudo systemctl` `disable` or `enable` `snap.nervatura.nervatura.service`

- Stop the service and start the program with the system tray icon/menu:
  ```
  sudo systemctl stop snap.nervatura.nervatura.service
  ```
  ```
  /snap/nervatura/current/nervatura -tray
  ```

- Stop the service and start with the tray icon/menu as a background process. The (temporary) log file of the application must also be set, otherwise the info and error messages that have been continuously written to the screen until now will be lost. The new HTTP log file is required for file access reasons.
  ```
  sudo systemctl stop snap.nervatura.nervatura.service
  ```
  ```
  NT_APP_LOG_FILE=./nervatura.log NT_HTTP_LOG_FILE=./nervatura-http.log /snap/nervatura/current/nervatura -tray &> /dev/null &
  ```

### Node.js **NPM**
```
npm install --save nervatura
```
Add a run script to your `package.json` file like this:

`"scripts": {
  "nervatura": "./node_modules/nervatura/bin/nervatura"
}`
```
npm run nervatura
```

### **Winget** (Windows Package Manager)

```
winget install --id Nervatura.Nervatura --source winget
```

- Start with the tray icon/menu and an .env file
  ```
  & "C:\Program Files\Nervatura\nervatura.exe" -tray -env C:\ProgramData\Nervatura\.env.example
  ```

- Start with the tray icon/menu as a background process (PS terminal). The (temporary) app and HTTP log file must also be set, otherwise the info and error messages that have been continuously written to the screen until now will be lost.
  ```
  $env:NT_APP_LOG_FILE="C:/ProgramData/Nervatura/data/nervatura.log"; $env:NT_HTTP_LOG_FILE="C:/ProgramData/Nervatura/data/http.log"; Start-Process -FilePath "nervatura.exe" -ArgumentList "-tray" -WindowStyle Hidden
  ```
  or just simply run this
  ```
  & "C:\Program Files\Nervatura\nervatura.bat"
  ```
Default Nervatura data directory: `C:\ProgramData\Nervatura\data`

### Prebuild binaries

[Linux and Windows](https://github.com/nervatura/nervatura/releases)

### Other platforms and custom build
```
git clone https://github.com/nervatura/nervatura.git
```
```
cd nervatura/service
```

Build command:

`
CGO_ENABLED=0 GOOS=$(OS_NAME) GOARCH=$(ARCH_NAME) go build -tags "$(TAGS)" -ldflags="-w -s -X main.Version=$(VERSION)" -o nervatura main.go
`

- `$(OS_NAME)` and `$(ARCH_NAME)`

You can see the list of supported platform by running:
```
go tool dist list
```
- `$(TAGS)` optional modul list: ***all, http, grpc, postgres, mysql, mssql, sqlite***
- `$(VERSION)` application version

Example:

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -tags "http postgres" -ldflags="-w -s -X main.Version=5.1.0" \
  -o nervatura main.go
```
See more: [Building Applications in GoLang](https://golangdocs.com/building-applications-in-golang)

## Configuration Options

The application uses environment variables to set configuration options. It will be read from the [.env.example](https://raw.githubusercontent.com/nervatura/nervatura/master/service/.env.example) file. Set the environment variables as needed!

The `.env` file can be created in the current working directory, where the command is executed (in development mode). The name and location of the configuration file can also be specified in the command line parameters:
```
nervatura -tray -env /path/.env.example
```

System tray icon / menu -> Configuration values