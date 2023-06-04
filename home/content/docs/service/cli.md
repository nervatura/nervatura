---
title: CLI AND CGO API
type: docs
weight: 10
bookToC: true
---

## CLI API

Full command line API. Program usage:
```
docker exec -i nervatura /nervatura -h
```
Windows (PowerShell):
```
& "C:\Program Files\Nervatura\nervatura.exe" -h
```
Linux Snap:
```
/snap/nervatura/current/nervatura -h
```
NPM:
```
./node_modules/nervatura/bin/nervatura -h
```
Example:
```
docker exec -i nervatura /nervatura \
  -c UserLogin -o "{\"username\":\"admin\",\"database\":\"demo\"}"
```

Golang docs:

[![GoDoc](https://godoc.org/github.com/nervatura/nervatura?status.svg)](https://pkg.go.dev/github.com/nervatura/nervatura/service/pkg/service#CLIService)

## CGO API

The CGO API a standard shared object binary file (.so or .dll) exposing Nervatura functions as a C-style APIs. It can be called from C, Python, Ruby, Node, Java, Lua, Dart, C#. Supported operating systems: Linux and Windows x64.

## For more examples, 
see 
- [Nervatura examples](https://github.com/nervatura/nervatura-examples)
