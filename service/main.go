/*
Open Source Business Management Framework

# Features

Nervatura is a business management framework. It can handle any type of business related information,
starting from customer details, up to shipping, stock or payment information. Developed as open-source
project and can be used freely under the scope of LGPLv3 License.

The framework is based on Nervatura Object Model (https://nervatura.github.io/nervatura/docs/model) specification.
It is a general open-data model, which can store all information generated in the operation of a usual corporation.

The Nervatura service is small and fast. A single ~6 MB file contains all the necessary dependencies.
The framework includes:

• CLI API (https://nervatura.github.io/nervatura/docs/service/cli#cli-api) (command line)

• CGO API (https://nervatura.github.io/nervatura/docs/service/cli#cgo-api) (C shared library)

• standard HTTP OPEN API (https://nervatura.github.io/nervatura/docs/service/api) for client communication

• HTTP/2-based gRPC API (https://nervatura.github.io/nervatura/docs/service/grpc) for server-side communication

• JWT generation, external token validation, SSL/TLS support and other HTTP security settings (https://github.com/nervatura/nervatura-service/blob/master/.env.example)

• built-in database drivers for postgres, mysql, mssql, sqlite databases

• a basic report generation library for creating simple PDF documents (eg. order, invoice, etc.)
or CSV data files

• sample report templates and REPORT EDITOR (https://nervatura.github.io/nervatura/docs/client/program/editor) GUI

• CLIENT (https://nervatura.github.io/nervatura/docs/client) Web Component application and a basic **ADMIN** interface

The client and report interface supports multilingualism (https://nervatura.github.io/nervatura/docs/start/customization#customize-the-appearance).

The framework can be easily extended with additional interfaces and functions in the supported languages:
C#, C++, Dart, Go, Java, Kotlin, Node, Objective-C, PHP, Python, Ruby

# Installation

https://nervatura.github.io/nervatura/docs/install

# Quick Start

https://nervatura.github.io/nervatura/docs/start

More info see http://www.nervatura.com.
*/
package main

import (
	"fmt"

	"github.com/nervatura/nervatura/service/app"
)

var (
	version = "dev"
)

func main() {
	fmt.Printf("Version: %s\n", version)
	_, err := app.New(version, nil)
	if err != nil {
		return
	}
}
