---
title: Examples
type: docs
weight: 30
bookToC: true
---

## **Quick Start**

### Node.js

```
git clone https://github.com/nervatura/nervatura-examples.git
cd nervatura-examples/node
npm install
npm start
```
Open your browser to http://localhost:8080

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