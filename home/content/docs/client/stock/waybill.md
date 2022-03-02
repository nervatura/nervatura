---
title: Tool movement
type: docs
weight: 40
bookToC: true
---

## Overview

The movement of [**products**](/docs/client/resources/product) in warehouses can be tracked with INVENTORY and DELIVERY document types. To track the moves of [**tools**](/docs/client/resources/tool) the WAYBILL type should be used. It helps to connect  the tool for a certain period to a [**customer**](/docs/client/resources/customer), [**employee**](/docs/client/resources/employee) or [**document**](/docs/client/document/document)  (orders, worksheets, rental, invoice).

The forms to provide easy handling are designed to enable connecting multiple tools to a customer, employee or document with one data sheet. This allows to easily manage cases like following up on  equipment being handed out to or taken back from your employees, as well as tracking the tools being used for a given project.

## Input fields

### Document No.
Unique ID, generated at the first data save. The format and value of the next data in row is taken from the [**DOCUMENT NUMBERING**](/docs/client/settings/numberdef) (code = waybill) data series.

### Direction
OUT, IN. Value cannot be changed after first save!

### Creation
Date of creation. Automatic value, cannot be changed.

### State
Its value and editing possibility is linked to [**ACCESS RIGHTS**](/docs/client/settings/usergroup#supervisor) setting. Not used in current version.

### Reference Type
DOCUMENT, CUSTOMER, EMPLOYEE

### Reference
Depending on the Reference Type, an ID for DOCUMENT, CUSTOMER, EMPLOYEE.

### Comment
Remarks field.

### Internal notes
Internal comments. Text defined in this field will not appear on the document.

## Related data

### [**METADATA**](/docs/client/settings/metadata)
Unlimited number of supplementary data can be added.

### [**REPORT NOTES**](/docs/client/program/notes)
Editable remarks, data for reports.

### DOCUMENT ITEM 
- **Shipping Date**: Stock release or receipt date.
- **Serial**: One of the items of [**TOOL**](/docs/client/resources/tool).
- **Comment**: Other remarks, data.

## Operations

### COPY FROM
Create a new, *same transaction type* document on the basis of current document's data. 

The dates and information related to creation gets updated and the references from the original document will not be transferred either.

### REPORT
[**DATA EXPORT**](/docs/client/program/export)

### BOOKMARK
Set a bookmark for the record. Later can be loaded from bookmarks at any time.