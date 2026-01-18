Nervatura
=========
![Coverage](https://img.shields.io/badge/Coverage-100%25-brightgreen)
[![Go Report Card](https://goreportcard.com/badge/github.com/nervatura/nervatura/v6)](https://goreportcard.com/report/github.com/nervatura/nervatura/v6)
[![GoDoc](https://godoc.org/github.com/nervatura/nervatura?status.svg)](https://pkg.go.dev/github.com/nervatura/nervatura/v6)
[![Release](https://img.shields.io/github/v/release/nervatura/nervatura)](https://github.com/nervatura/nervatura/releases)

Open Source Business Management Framework

The Nervatura **v6 is currently beta** version. Available for download from the [GitHub releases](https://github.com/nervatura/nervatura/releases) page. Docker image is also available in beta version:

    docker pull nervatura/nervatura:beta

## Features

Nervatura is a business management framework based on **open-data principle**. It can handle any type of business related information, starting from customer details, up to shipping, stock or payment information.

The main aspects of its design were:

* simple and transparent structure
* capability of storing different data types of an average company
* effective, easily expandable and secure data storage
* support of several database types
* well documented, easy data management
* ready for LLM integration and AI-assisted data processing

The framework is based on Nervatura Object [**MODEL**](https://nervatura.github.io/nervatura/model) specification. It is a general **open-data model**, which can store all information generated in the operation of a usual corporation.

The Nervatura service is small and fast. A single ~9 MB file contains all the necessary dependencies.
The framework includes:
- [**CLI API**](https://nervatura.github.io/nervatura/cli/#cli) (command line)
- [**CGO API**](https://nervatura.github.io/nervatura/cli/#cgo) (C shared library)
- standard HTTP [**OpenAPI**](https://nervatura.github.io/nervatura/open/) for client communication
- HTTP/2-based [**gRPC API**](https://nervatura.github.io/nervatura/grpc/) for server-side communication
- [**MCP**](https://nervatura.github.io/nervatura/mcp/) (Model Context Protocol) for LLM integration
- built-in database drivers for postgres, mysql, mssql, sqlite databases
- a basic report generation library for creating simple PDF documents (eg. order, invoice, etc.) 
or CSV data files
- sample report templates and [**Report Editor**](https://nervatura.github.io/nervatura/editor/) extension
- [**Nervatura Client**](https://nervatura.github.io/nervatura/client/) responsive graphical user interface

The Nervatura Backend is a simple interface layer that provides multiple, well-documented data access protocols for handling data. With their help, we can use the best data access for every development language and environment. Using the functions of the interfaces, we can be sure that the data is always read or written from the databases correctly and simply. The data can be stored in several types of databases, but they can be handled in the same format, and the database types can be easily changed.

The Nervatura Service has a modular structure, where most modules are optional. The default service includes all modules, but you [can build a customized service from them](https://nervatura.github.io/nervatura/install/#custom).

The Nervatura Client is a [server-side responsive component](https://github.com/nervatura/component) application. It was created so that all the business data of the framework can be managed immediately after installation through a graphical interface. The client and report interface supports [multilingual interface](https://nervatura.github.io/nervatura/start/#customization).

The Nervatura Framework **can be used independently**, but it is basically designed to provide a stable and secure foundation for self-developed, customized enterprise business systems. The framework **can be easily extended** with additional user interfaces or data management functions in **any programming language** or technology. Using the data from the framework, you can easily create your own web stores, user input interfaces or data interfaces for other systems. The framework is **ready for LLM integration and AI-powered data processing** with a purpose-built, secure and fast built-in [**MCP**](https://nervatura.github.io/nervatura/mcp/) server.

Nervatura Client supports the business processes that most companies may need. During your own developments, you can only focus on those that really require unique solutions, and you can use the technology that best suits the purpose. This type of development means greater **technological independence and security**, since your self-developed applications are only connected to other systems through well-documented interfaces, so unnecessary external technological dependencies cannot develop.

You can find more information about the use of different programming languages and development environments in the [**Examples**](https://nervatura.github.io/nervatura/examples/) section.

[**Installation**](https://nervatura.github.io/nervatura/install/) and [**Quick Start**](https://nervatura.github.io/nervatura/start/)

[**Upgrade**](https://nervatura.github.io/nervatura/upgrade/) from v5 version to v6.*

More info see 

http://www.nervatura.com
