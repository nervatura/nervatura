---
title: Report queue
type: docs
weight: 50
bookToC: true
---

## Overview

Items to the list can be added by **REPORT** command of the forms specified in **Report Type**, then by the **Report Queue** command.

## Input fields

### DATA FILTER
- **Report Type**: CUSTOMER, PRODUCT, EMPLOYEE, TOOL, PROJECT, ORDER, OFFER, INVOICE, RECEIPT, RENTAL,
WORKSHEET, DELIVERY, CORRECTION, TOOL MOVEMENT, PRODUCTION, FORMULA, BANK STATEMENT, PETTY CASH
- **Start Date / End Date**: The date when added to the REPORT QUEUE list
- **Document No.**: Number of the transaction, event or resource (eg. Invoice No., Customer No., Serial etc.).
- **Username**: The user who has added it to the REPORT QUEUE list

### OPTIONS
- **Export Mode**: PRINT, PDF, XML. When command **EXPORT ALL** is used it is mandatory to fill in, in case of printing it is not taken into account by the program.
- **Orientation**: Portrait, Landscape
- **Size**: A3, A4, A5, Letter, Legal

{{< hint info >}}

The PRINT mode additional parameters are the **Size** and **Orientation**. The number of **Copies** is specified per each item (list row).

{{< /hint >}}

## Operations

### SEARCH
Lists the items filtered by the REPORT QUEUE **DATA FILTER** or all.

{{< hint info >}}

When **SEARCH** is used the **OPTIONS** and **DATA FILTER** part closes automatically, can be opened any time again by selecting the header.

{{< /hint >}}

### EXPORT ALL
Exports all the items listed in REPORT QUEUE in the format specified in **Export Mode** option. More information: [**DATA EXPORT**](/docs/client/program/export)

{{< hint info >}}

Only the PDF and XML options are available in **Export Mode**. Additional parameters for PDF the **Orientation** and **Size**. The PRINT option only works for items selected line by line. Thus it is possible for each row to export it to either PDF or XML formats.

{{< /hint >}}
