---
title: Examples
type: docs
weight: 30
bookToC: true
---

## **Quick Start**

### Node.js

```
$ git clone https://github.com/nervatura/nervatura-examples.git
$ cd nervatura-examples/node
$ npm install
$ npm start
```
Open your browser to http://localhost:8080

### Python

1. Download the Python examples

    ```
    $ git clone https://github.com/nervatura/nervatura-examples.git
    $ cd nervatura-examples/python
    $ pip install -r requirements.txt
    $ python main.py
    ```
    Open your browser to http://localhost:8000

2. Download the Nervatura

  - Linux x64
    ```
    $ mkdir bin
    $ cd bin
    $ curl -L -s https://api.github.com/repos/nervatura/nervatura/releases/latest | grep -o -E "https://(.*)nervatura_(.*)_linux_amd64.tar.gz" | wget -qi -
    $ tar -zxf *.gz
    $ rm *.gz
    ```

  - Windows users:
    - download the [latest version](https://github.com/nervatura/nervatura/releases/latest) to the `/bin` directory
    - change the value of the `NT_EXAMPLE_SERVICE_PATH` (`.env.example` file): "bin/nervatura" -> "bin/nervatura.exe"
    - change the value of the `NT_EXAMPLE_SERVICE_LIB` (`.env.example` file): "bin/nervatura.so" -> "bin/nervatura.dll"

3. Start the Nervatura backend server (gRPC and HTTP examples)

    ```
      $ bin/nervatura -env .env.example
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