Nervatura
=========

Open Source Business Management Framework

## Features

Nervatura is a business management framework based on **open-data principle**. It can handle any type of business related information, starting from customer details, up to shipping, stock or payment information.

The framework is based on Nervatura [Object Model](https://nervatura.github.io/nervatura/model) specification. The main aspects of its design were:

* simple and transparent structure
* capability of storing different data types of an average company
* effective, easily expandable and secure data storage
* support of several database types
* well documented, easy data management

The framework is based on Nervatura Object [**MODEL**](https://nervatura.github.io/nervatura/model) specification. It is a general **open-data model**, which can store all information generated in the operation of a usual corporation.

The Nervatura service is small and fast. A single ~6 MB file contains all the necessary dependencies.
The framework includes:
- [**CLI API**](https://nervatura.github.io/nervatura/#cli-api) (command line)
- [**CGO API**](https://nervatura.github.io/nervatura/#cgo-api) (C shared library)
- standard HTTP [**RESTful API**](https://nervatura.github.io/nervatura/api) for client communication
- HTTP/2-based [**gRPC API**](https://nervatura.github.io/nervatura/grpc) for server-side communication
- JWT generation, external token validation, SSL/TLS support and other HTTP security [settings](https://github.com/nervatura/nervatura-service/blob/master/.env.example)
- built-in database drivers for postgres, mysql, sqlite databases
- a basic report generation library for creating simple PDF documents (eg. order, invoice, etc.) 
or CSV data files
- sample report templates and [**REPORT EDITOR**](https://nervatura.github.io/nervatura/docs/editor) GUI
- PWA [**CLIENT**](https://nervatura.github.io/nervatura/docs) application and a basic **ADMIN** interface

The client and report interface supports [multilingualism](https://nervatura.github.io/nervatura/#customize-the-appearance). The framework can be easily extended with additional interfaces and functions in the supported languages: 
C, C#, C++, Dart, Go, Java, Kotlin, Node, Objective-C, PHP, Python, Ruby (CGO and gRPC API)

[**Installation**](https://nervatura.github.io/nervatura/#installation) and [**Quick Start**](https://nervatura.github.io/nervatura/#quick-start)

More info see 

http://www.nervatura.com
