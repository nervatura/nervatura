---
title: Place
type: docs
weight: 70
bookToC: true
---

## Overview

In Nervatura those logistic points where our own or someone else’s assets and resources are recorded or physically stored, belong to a special group called PLACE. These are the warehouses, but also the petty cash, where cash is stored and even a virtual place like the bank account belong to this group.

## Input fields

### Place No.
Unique ID, generated at the first data save. The format and value of the next data in row is taken from the [**DOCUMENT NUMBERING**](/docs/client/settings/numberdef) (code = planumber) data series.

### Type
bank, cash, warehouse

### Description
Short description.

### Currency
The elements of [**CURRENCY**](/docs/client/settings/currency). Only for *bank* and *cash* types, appears after first save. Its default value is taken from [**DEFAULT SETTINGS**](/docs/client/settings/setting) *default currency*.

### Address data
**Zipcode**, **City**, **Street**

### Comment
Remarks field.

## Related data

### [**METADATA**](/docs/client/settings/metadata)
Unlimited number of supplementary data can be added.

### CONTACT INFO
***Firtsname, Surname, Status, Phone, Mobile, Other, Email, Comment***<br />
Unlimited number of contact data can be added.

## Operations

### REPORT
[**DATA EXPORT**](/docs/client/program/export)

### BOOKMARK
Set a bookmark for the record. Later can be loaded from bookmarks at any time.
