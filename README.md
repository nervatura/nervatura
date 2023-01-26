Nervatura
=========

Open Source Business Management Framework

## Features

Nervatura is a business management framework based on **open-data principle**. It can handle any type of business related information, starting from customer details, up to shipping, stock or payment information.

The main aspects of its design were:

* simple and transparent structure
* capability of storing different data types of an average company
* effective, easily expandable and secure data storage
* support of several database types
* well documented, easy data management

The framework is based on Nervatura Object [**MODEL**](https://nervatura.github.io/nervatura/docs/model) specification. It is a general **open-data model**, which can store all information generated in the operation of a usual corporation.

The Nervatura service is small and fast. A single ~6 MB file contains all the necessary dependencies.
The framework includes:
- [**CLI API**](https://nervatura.github.io/nervatura/docs/service/cli#cli-api) (command line)
- [**CGO API**](https://nervatura.github.io/nervatura/docs/service/cli#cgo-api) (C shared library)
- standard HTTP [**RESTful API**](https://nervatura.github.io/nervatura/docs/service/api) for client communication
- HTTP/2-based [**gRPC API**](https://nervatura.github.io/nervatura/docs/service/grpc) for server-side communication
- JWT generation, external token validation, SSL/TLS support and other HTTP security [settings](https://github.com/nervatura/nervatura-service/blob/master/.env.example)
- built-in database drivers for postgres, mysql, sqlite databases
- a basic report generation library for creating simple PDF documents (eg. order, invoice, etc.) 
or CSV data files
- sample report templates and [**REPORT EDITOR**](https://nervatura.github.io/nervatura/docs/client/program/editor) GUI
- [**CLIENT**](https://nervatura.github.io/nervatura/docs/client) Web Component application and a basic **ADMIN** interface

The Nervatura [**Service**](https://nervatura.github.io/nervatura/docs/service) is a simple interface layer that provides multiple, well-documented data access protocols for handling data. With their help, we can use the best data access for every development language and environment. Using the functions of the interfaces, we can be sure that the data is always read or written from the databases correctly and simply. The data can be stored in several types of databases, but they can be handled in the same format, and the database types can be easily changed.

The Nervatura Service has a modular structure, where most modules are optional. The default service includes all modules, but you [can build a customized service from them](https://nervatura.github.io/nervatura/docs/install/#other-platforms-and-custom-build).

The Nervatura [**Client**](https://nervatura.github.io/nervatura/docs/client) is a standard HTML5/ES6 [Web Component](https://developer.mozilla.org/en-US/docs/Web/Web_Components) application that contains no other external dependencies apart from the [lit helper functions](https://lit.dev). The standard HTML5 Web Components can be easily integrated or called from other javascript frameworks. It was created so that all the business data of the framework can be managed immediately after installation through a graphical interface. The client and report interface supports [multilingualism](https://nervatura.github.io/nervatura/docs/start/customization#customize-the-appearance).

The Nervatura Framework **can be used independently**, but it is basically designed to provide a stable and secure foundation for self-developed, customized enterprise business systems. The framework **can be easily extended** with additional user interfaces or data management functions in **any programming language** or technology. Using the data from the framework, you can easily create your own web stores, user input interfaces or data interfaces for other systems. 

Nervatura Client supports the business processes that most companies may need. During your own developments, you can only focus on those that really require unique solutions, and you can use the technology that best suits the purpose. This type of development means greater **technological independence and security**, since your self-developed applications are only connected to other systems through well-documented interfaces, so unnecessary external technological dependencies cannot develop.

You can find more information about the use of different programming languages and development environments in the [**Examples**](https://nervatura.github.io/nervatura/docs/start/examples) section:
- [Node.js and NPM](https://nervatura.github.io/nervatura/docs/start/examples/#nodejs-and-npm)
- [Python and Snap or Windows Package Manager](https://nervatura.github.io/nervatura/docs/start/examples/#python-and-snap-or-windows-package-manager)
- [Go and Docker](https://nervatura.github.io/nervatura/docs/start/examples/#go-and-docker)

[**Installation**](https://nervatura.github.io/nervatura/docs/install) and [**Quick Start**](https://nervatura.github.io/nervatura/docs/start)

More info see 

http://www.nervatura.com
