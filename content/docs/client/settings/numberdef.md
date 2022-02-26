---
title: Document numbering
type: docs
weight: 20
bookToC: true
---

## Overview

Format settings of unique identifiers of documents (eg. orders, invoices, cash receipts, delivery notes etc.) and other resources (eg. customers, products, employees, etc.). If a new item is created the identifiers will follow the rule which was set here.

## Input fields

### Code
A unique identifier for a certain set of rules. Its value can not be changed.

### Prefix
The text prefix of the identifier. It can be any length, but usage of special characters,  spaces in the text is not recommended.

### Year
If selected, the created identifier will contain the year. The number is not formed automatically from the current date, but can be set in the [**DEFAULT SETTINGS**](/docs/client/settings/setting) section, 
in the ***business year*** field. This is due to avoid any technical issues resulting 
from changes in different fiscal or calendar years.

### Separator
The separator character in the identifier. Default value: /

### Lenght
The value field is arranged in such length to the right and filled with zeros.

### Value
The current status of the counter, the next sequence number will be one value higher than this one. It is possible to re-set the counter, but the uniqueness must be ensured in all cases!
