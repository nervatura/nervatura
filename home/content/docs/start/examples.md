---
title: Examples
type: docs
weight: 30
bookToC: true
---

## **Quick Start**

### **Node.js and NPM**

Recommended API: 
- [gRPC](/docs/service/grpc) 
- [CGO](/docs/service/cli#cgo-api) (Linux and Windows x64)

```
$ git clone https://github.com/nervatura/nervatura-examples.git
$ cd nervatura-examples/node
$ npm install
$ npm start
```
Open your browser to http://localhost:8080

You do not need a Nervatura backend server to use the [CLI](/docs/service/cli#cli-api) and [CGO](/docs/service/cli#cgo-api). Automatic server start can be turned off with `NT_EXAMPLE_SERVICE_DISABLED=true`  (see in the `nervatura-examples/node/.env` file).

### **Python and Snap or prebuild binaries**

Recommended API:  
- [CGO](/docs/service/cli#cgo-api) (Linux and Windows x64)
- [gRPC](/docs/service/grpc)

1. Download the Python examples

    ```
    $ git clone https://github.com/nervatura/nervatura-examples.git
    $ cd nervatura-examples/python
    $ pip install -r requirements.txt
    $ python main.py
    ```
    Open your browser to http://localhost:8000

2. Nervatura backend

  - Linux
    ```
    $ sudo snap install nervatura
    ```
    The [CLI](/docs/service/cli#cli-api) and [CGO](/docs/service/cli#cgo-api) is ready to use. To use the [gRPC](/docs/service/grpc) and [HTTP](/docs/service/api), start Nervatura service with the .env.example settings (`nervatura-examples/python` directory):
    ```
    $ /snap/nervatura/current/nervatura -env $(pwd)/.env.example
    ```

  - Windows users:
    - download the [latest version](https://github.com/nervatura/nervatura/releases/latest) to the `nervatura-examples/python/bin` directory
    - change the value of the `NT_EXAMPLE_SERVICE_PATH` (`nervatura-examples/python/.env.example` file): "/snap/nervatura/current/nervatura" -> "bin/nervatura.exe"
    - change the value of the `NT_EXAMPLE_SERVICE_LIB` (`nervatura-examples/python/.env.example` file): "/snap/nervatura/current/nervatura.so" -> "bin/nervatura.dll"

    Start the Nervatura backend server ([gRPC](/docs/service/grpc) and [HTTP](/docs/service/api) examples)

    ```
      $ bin/nervatura -env .env.example
    ```
### **Go and Docker**

Recommended API:
- [gRPC](/docs/service/grpc)

1. Download the Go examples

    ```
    $ git clone https://github.com/nervatura/nervatura-examples.git
    $ cd nervatura-examples/go
    $ go mod vendor
    $ go run ./main.go
    ```
    Open your browser to http://localhost:7000

2. Nervatura backend (`nervatura-examples/go` directory)
    ```
    $ docker run -i -t --rm --name nervatura --env-file .env.example -p 5000:5000 -p 9200:9200 -v $(pwd)/data:/data nervatura/nervatura:latest
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
  - with the HMAC algorithm example
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