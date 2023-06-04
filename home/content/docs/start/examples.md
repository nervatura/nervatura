---
title: Examples
type: docs
weight: 30
bookToC: true
---

## **Quick Start**

### **Node.js**

Recommended API: 
- [gRPC](/docs/service/grpc) 
- [CGO](/docs/service/cli#cgo-api) (Linux and Windows x64, without Admin and Client GUI)

```
git clone https://github.com/nervatura/nervatura-examples.git
```
```
cd nervatura-examples/node
```
```
npm install
```
```
npm start
```
Open your browser to http://localhost:8080

You do not need a Nervatura backend server to use the [CLI](/docs/service/cli#cli-api) and [CGO](/docs/service/cli#cgo-api). Automatic server start can be turned off with `NT_EXAMPLE_SERVICE_DISABLED=true`  (see in the `nervatura-examples/node/.env` file).

### **Python**

Recommended API:
- [gRPC](/docs/service/grpc)
- [CGO](/docs/service/cli#cgo-api) (Linux and Windows x64, without Admin and Client GUI)

1. ***Download the Python examples***

    ```
    git clone https://github.com/nervatura/nervatura-examples.git
    ```
    ```
    cd nervatura-examples/python
    ```
    ```
    pip install -r requirements.txt
    ```
    ```
    python main.py
    ```
    Open your browser to http://localhost:8000

2. ***Nervatura backend***

  - Linux
    ```
    sudo snap install nervatura
    ```
    ```
    sudo systemctl stop snap.nervatura.nervatura.service
    ```
    The [CLI](/docs/service/cli#cli-api) and [CGO](/docs/service/cli#cgo-api) is ready to use. To use the [gRPC](/docs/service/grpc) and [HTTP](/docs/service/api), start Nervatura service with the .env.example settings (`nervatura-examples/python` directory):
    ```
    /snap/nervatura/current/nervatura -tray -env $(pwd)/.env.example
    ```

  - Windows:
    ```
    winget install --id Nervatura.Nervatura --source winget
    ```
    - change the value of the `NT_EXAMPLE_SERVICE_PATH` (`nervatura-examples/python/.env.example` file): "/snap/nervatura/current/nervatura" -> *"C:/Program Files/Nervatura/nervatura.exe"*
    - change the value of the `NT_EXAMPLE_SERVICE_LIB` (`nervatura-examples/python/.env.example` file): "/snap/nervatura/current/nervatura.so" -> *"C:/ProgramData/Nervatura/nervatura.dll"*

    Start the Nervatura backend server ([gRPC](/docs/service/grpc) and [HTTP](/docs/service/api) examples) with the .env.example settings (`nervatura-examples/python` directory):

    ```
      & "C:\Program Files\Nervatura\nervatura.exe" -tray -env .env.example
    ```
  
  - Docker (without [CGO](/docs/service/cli#cgo-api))
    
    Change the value of the `NT_EXAMPLE_SERVICE_PATH` (`nervatura-examples/python/.env.example` file): "/snap/nervatura/current/nervatura" -> *"docker"*

    Start the Nervatura backend server with the .env.example settings (`nervatura-examples/python` directory):

    ```
    docker run -i -t --rm --name nervatura --env-file .env.example -p 5000:5000 -p 9200:9200 -v $(pwd)/data:/data nervatura/nervatura:latest
    ```

### **Go**

Recommended API:
- [gRPC](/docs/service/grpc)

1. ***Download the Go examples***

    ```
    git clone https://github.com/nervatura/nervatura-examples.git
    ```
    ```
    cd nervatura-examples/go
    ```
    ```
    go mod vendor
    ```
    ```
    go run ./main.go
    ```
    Open your browser to http://localhost:7000

2. ***Nervatura backend*** (`nervatura-examples/go` directory)

- Docker
    ```
    docker run -i -t --rm --name nervatura --env-file .env.example -p 5000:5000 -p 9200:9200 -v $(pwd)/data:/data nervatura/nervatura:latest
    ```
- Linux and Windows

  Follow the instructions in the [python example](/docs/start/examples/#python)

### **PHP**

Recommended API:
- [gRPC](/docs/service/grpc)
- [HTTP](/docs/service/api)

1. ***Download and install the PHP examples*** (linux)

    ```
    git clone https://github.com/nervatura/nervatura-examples.git
    ```
    ```
    cd nervatura-examples/php
    ```
    ```
    composer install
    ```
    ```
    cd public
    ```
    ```
    php -S localhost:8000
    ```
    Open your browser to http://localhost:8000

2. ***gRPC install*** (linux, optional)

    ```
    sudo pecl install grpc
    ```
    ```
    sudo pecl install protobuf
    ```
    ```
    php -d extension=grpc.so -d extension=protobuf.so -S localhost:8000
    ```
    More details or Windows installation: [Install gRPC for PHP](https://cloud.google.com/php/grpc)

2. ***Nervatura backend*** (`nervatura-examples/php` directory)
  
  - Docker
    ```
    docker run -i -t --rm --name nervatura --env-file .env.example -p 5000:5000 -p 9200:9200 -v $(pwd)/data:/data nervatura/nervatura:latest
    ```

  - Linux and Windows

    Follow the instructions in the [python example](/docs/start/examples/#python)

### **Flutter** and **Dart**

Recommended API:
- [gRPC](/docs/service/grpc)
- [CGO](/docs/service/cli#cgo-api) (Linux and Windows x64, without Admin and Client GUI)

1. ***Download the Flutter/Dart examples***

    ```
    git clone https://github.com/nervatura/nervatura-examples.git
    ```
    ```
    cd nervatura-examples/flutter
    ```
2. ***Nervatura backend***

  - Linux
    ```
    sudo snap install nervatura
    ```
    ```
    sudo systemctl stop snap.nervatura.nervatura.service
    ```
    The [CLI](/docs/service/cli#cli-api) and [CGO](/docs/service/cli#cgo-api) is ready to use. To use the [gRPC](/docs/service/grpc) and [HTTP](/docs/service/api), start Nervatura service with the .env.example settings (`nervatura-examples/flutter` directory):
    ```
    /snap/nervatura/current/nervatura -env $(pwd)/.env.example
    ```
  
  - Windows:
    ```
    winget install --id Nervatura.Nervatura --source winget
    ```
    - change the value of the `NT_EXAMPLE_SERVICE_PATH` (`nervatura-examples/flutter/.env.example` file): "/snap/nervatura/current/nervatura" -> *"C:/Program Files/Nervatura/nervatura.exe"*
    - change the value of the `NT_EXAMPLE_SERVICE_LIB` (`nervatura-examples/flutter/.env.example` file): "/snap/nervatura/current/nervatura.so" -> *"C:/ProgramData/Nervatura/nervatura.dll"*

    Start the Nervatura backend server ([gRPC](/docs/service/grpc) and [HTTP](/docs/service/api) examples) with the .env.example settings (`nervatura-examples/flutter` directory):

    ```
      & "C:\Program Files\Nervatura\nervatura.exe" -env .env.example
    ```
  
  - Docker (without [CGO](/docs/service/cli#cgo-api))
    
    Change the value of the `NT_EXAMPLE_SERVICE_PATH` (`nervatura-examples/flutter/.env.example` file): "/snap/nervatura/current/nervatura" -> *"docker"*

    Start the Nervatura backend server with the .env.example settings (`nervatura-examples/flutter` directory):

    ```
    docker run -i -t --rm --name nervatura --env-file .env.example -p 5000:5000 -p 9200:9200 -v $(pwd)/data:/data nervatura/nervatura:latest
    ```

3. ***Dart backend***

    Install the dependencies and start the Dart backend server with the .env.example settings (`nervatura-examples/flutter` directory):

    ```
    dart pub get -C server
    ```

  - Run without [CGO](/docs/service/cli#cgo-api) settings:

    ```
    dart run server/lib/server.dart
    ```

  - Run with [CGO](/docs/service/cli#cgo-api) (linux example):

    ```
    NT_API_KEY=EXAMPLE_API_KEY NT_TOKEN_PUBLIC_KID=PUBLIC_KID NT_TOKEN_PUBLIC_KEY=data/public.key dart run server/lib/server.dart
    ```

4. ***Flutter client***

  - Install the all dependencies (`nervatura-examples/flutter/client` directory):

    ```
    flutter pub get
    ```

  - Windows or Linux desktop:

    Run without [CGO](/docs/service/cli#cgo-api) settings (linux example):

    ```
    flutter run -d linux --dart-define-from-file=env-example.json
    ```

    Run with [CGO](/docs/service/cli#cgo-api) (linux example):

    ```
    NT_API_KEY=EXAMPLE_API_KEY NT_TOKEN_PUBLIC_KID=PUBLIC_KID NT_TOKEN_PUBLIC_KEY="../data/public.key" NT_ALIAS_DEMO="sqlite://file:../data/demo.db?cache=shared&mode=rwc" flutter run -d linux --dart-define-from-file=env-example.json
    ```

  - Web or other platform:

    ```
    flutter run -d chrome --dart-define-from-file=env-example.json
    ```



## **Examples**

### Create a demo database

  - [HTTP](/docs/service/api), [CLI](/docs/service/cli#cli-api), [CGO](/docs/service/cli#cgo-api)
, [gRPC](/docs/service/grpc)
  - *All examples require a demo database. Please run this first!*

### Basic password login

  - [HTTP](/docs/service/api), [CLI](/docs/service/cli#cli-api), [CGO](/docs/service/cli#cgo-api)
  , [gRPC](/docs/service/grpc)

### Custom token (passwordless) login

  - [HTTP](/docs/service/api), [CLI](/docs/service/cli#cli-api), [CGO](/docs/service/cli#cgo-api)
  , [gRPC](/docs/service/grpc)
  - public/private key pair example

### Nervatura Client custom token login
  
  - [HTTP](/docs/service/api), [CLI](/docs/service/cli#cli-api), [CGO](/docs/service/cli#cgo-api)
  , [gRPC](/docs/service/grpc)
  - Authorization Code
  - Implicit (token) Grant
  - error or logout callback

### Using external API for token based authentication

  - [HTTP](/docs/service/api), [gRPC](/docs/service/grpc)
  - use of an external identification service provider (Facebook, GitHub, Google, Firebase and others)

### Nervatura Client language translation
  
  - [client_config.json](https://github.com/nervatura/nervatura/tree/master/dist) file example

### Create an invoice
  
  - [HTTP](/docs/service/api), [CLI](/docs/service/cli#cli-api), [CGO](/docs/service/cli#cgo-api)
  , [gRPC](/docs/service/grpc)
  - creating or updating a new or existing customer, customer address and contact data
  - create a new invoice
  - create and download a PDF invoice

### Nervatura Client menu shortcuts

  - [HTTP](/docs/service/api), [CLI](/docs/service/cli#cli-api), [CGO](/docs/service/cli#cgo-api)
  , [gRPC](/docs/service/grpc)
  - External page - GET type example
  - Email sending - POST type example

### Nervatura CSV Report Example

  - [HTTP](/docs/service/api), [CLI](/docs/service/cli#cli-api), [CGO](/docs/service/cli#cgo-api)
  , [gRPC](/docs/service/grpc)
  - simple customer contact list example
  - input parameters

### An example of all available functions in all APIs

  - [HTTP](/docs/service/api), [CLI](/docs/service/cli#cli-api), [CGO](/docs/service/cli#cgo-api)
  , [gRPC](/docs/service/grpc)

