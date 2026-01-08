/*
# Nervatura v6

- Open Source Business Management Framework

# Features

Nervatura is a business management framework. It can handle any type of business related information,
starting from customer details, up to shipping, stock or payment information. Developed as open-source
project and can be used freely under the scope of LGPLv3 License.

The framework is based on Nervatura Object Model (https://nervatura.github.io/nervatura/docs/model/) specification.
It is a general open-data model, which can store all information generated in the operation of a usual corporation.

The Nervatura service is small and fast. A single ~9 MB file contains all the necessary dependencies.
The framework includes:

• CLI API (https://nervatura.github.io/nervatura/docs/cli/) (command line interface)

• CGO API (https://nervatura.github.io/nervatura/docs/cli/) (C shared library)

• standard HTTP OPEN API (https://nervatura.github.io/nervatura/docs/open/) for client communication

• HTTP/2-based gRPC API (https://nervatura.github.io/nervatura/docs/grpc/) for server-side communication

• MCP (Model Context Protocol) for LLM integration (https://nervatura.github.io/nervatura/docs/mcp/)

• JWT generation, external token validation, SSL/TLS support and other HTTP security settings (https://github.com/nervatura/nervatura-service/blob/master/.env.example)

• built-in database drivers for postgres, mysql, mssql, sqlite databases

• a basic report generation library for creating simple PDF documents (eg. order, invoice, etc.)
or CSV data files

• sample report templates and Report Editor extension (https://nervatura.github.io/nervatura/docs/editor/)

• Nervatura Client (https://nervatura.github.io/nervatura/docs/client/) graphical user interface

The framework can be easily extended with additional interfaces and functions in the any languages.

# Installation

https://nervatura.github.io/nervatura/docs/install/

# Quick Start

https://nervatura.github.io/nervatura/docs/create/

More info see http://www.nervatura.com.
*/
package main

import (
	"log"
	"os"

	"github.com/nervatura/nervatura/v6/pkg/app"
)

var (
	version = "dev"
)

func main() {
	log.Printf("Version: %s\n", version)
	if _, err := app.New(version, nil); err != nil {
		log.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
