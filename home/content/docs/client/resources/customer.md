---
title: Customer
type: docs
weight: 10
bookToC: true
---

## Overview

Nervatura handles all the partners, suppliers and customers of the company consistantly, at one place.

## Input fields

### Customer No.
Unique ID, generated at the first data save. The format and value of the next data in row is taken from the [**DOCUMENT NUMBERING**](/docs/client/settings/numberdef) (code = custnumber) data series.

### Customer Name
Full name of the customer.

### Customer Type
company, private, other

### Taxnumber, Account No., Tax-free
Main general data.

### Credit line
Customer's credit limit. Data is used by financial reports.

### Payment period
Value given here overrides the [**default deadline**](/docs/client/settings/setting#default-deadline) field.

### Discount
If new product line is added (offer, order, invoice etc.) all products will receive the discount percentage specified in this field. If the product has a separate customer price, the value specified here will not be considered by the program.

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


