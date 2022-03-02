---
title: Menu shortcuts
type: docs
weight: 40
bookToC: true
---

## Overview

It is possible to create shortcuts in the program, which can call remote functions. Input parameters can also be set and the user will be informed about the result by a message box.

## Input fields

### Menukey
The shortcut ID. A unique value, can not be repeated.

### Description
The name set here will be displayed on the user interface.

### Method
When set ***GET***, the program opens your default browser and hands over to it the given **Address** and also the input parameters if there were any. In case of a ***POST*** value, the program will contact the server (default host URL or **Address**), run the specified program (**Funcname**) and wait for the response.

### Modul
The shortcuts can also be displayed in groups. This feature is not supported by the current version of the program.

### Icon
Individual images can be assigned to shortcuts. This feature is not supported by the current version of the program.

### Funcname
The name of the server-side function which is called when the **Method** is set to ***POST***. The server’s address and protocol will automatically be taken by the program from login data (eg http://localhost:8000/) 
or **Address** field value.

### Address
The full URL address if the the address differs from the server's one. (eg https://www.google.com.)

### Parameters
Undefined number of parameters can also be assigned to shortcuts. These will be requested when the shortcut is used. When a new parameter is created the following can be specified:
- **Name**: Parameter ID. A unique value, can not be repeated.
- **Description**: The name set here will be displayed on the user interface.
- **Type**: bool, date, float, integer, string. When keyed in only the data type specified here will be accepted by the program.
- **Order**: The order of the parameters can be specified
