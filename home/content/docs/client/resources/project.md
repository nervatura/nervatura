---
title: Project
type: docs
weight: 50
bookToC: true
---

## Overview

[**OFFER, ORDER, WORKSHEET, RENTAL, INVOICE, RECEIPT**](/docs/client/document/document) documents could be linked to projects by setting so their Project No. field.

## Input fields

### Project No.
Unique ID, generated at the first data save. The format and value of the next data in row is taken from the [**DOCUMENT NUMBERING**](/docs/client/settings/numberdef) (code = pronumber) data series.

### Description
The name of the project.

### Start Date, End Date
Optional. The project's start and end dates.

### Customer
Optional. The name of the [**customer**](/docs/client/resources/customer) of the project.

### Comment
Remarks field.

## Related data

### [**METADATA**](/docs/client/settings/metadata)
Unlimited number of supplementary data can be added.

### ADDRESS DATA
***Country, State, Zipcode, City, Street, Comment***<br />
Unlimited number of address data can be added. The default report templates will consider the 
first address row being the address.

### CONTACT INFO
***Firtsname, Surname, Status, Phone, Mobile, Other, Email, Comment***<br />
Unlimited number of contact data can be added.

### [**EVENT**](/docs/client/resources/event)
Unlimited number of events can be added.

## Operations

### REPORT
[**DATA EXPORT**](/docs/client/program/export)

### BOOKMARK
Set a bookmark for the record. Later can be loaded from bookmarks at any time.
