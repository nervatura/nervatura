---
title: Metadata
type: docs
weight: 50
bookToC: true
---

## Overview

With Nervatura it is easy to store a variety of data. If some new information is needed for which there have not been any data collected yet, the case is simple. Just create a new attribute, specify its type and connect it to the data you would like to use it with.

## Input fields

### Data GUID
Automatically generated internal ID. A unique value, can not be changed.

### Description
The name set here will be displayed on the user interface.

### Data Type
Valid values: customer, employee, event, place, product, project, tool, trans. Those data types which we would like the new feature to be connected to. If the trans value is selected then it can be used on all forms under DOCUMENT, PAYMENT and STOCK CONTROL menus (eg. offer, order, invoice, etc.).

### Value Type
- **bool**: Two positions, YES/NO or TRUE/FALSE values
- **integer, float**: Only numbers can be entered, in case of integer only numbers without decimals
- **date, time**: Only valid date or time can be set
- **string, notes**: Any text type value. In case of string shorter, for notes longer editable text.
- **password**: For data entry a "password" type field will be displayed, where the typed in characters cannot be identified. Note that the value in the database is not going to be encrypted, stored only as plain text!
- **valuelist**: After saving a **Value list** field will be displayed, in which a list of items can be entered. The elements can be separated from each other by | sign. When used, only those values will be allowed to enter which have been set in this list, other value is not accepted.
- **urlink**: Any URL link.
- **customer, employee, place, product, project, tool**: only valid data stored in the database and chosen through the search engine can be used.
- **transitem**: we can choose from data in DOCUMENT menu (eg offer, order, invoice, etc.) by using the search tool.
- **transpayment**: we can choose from data in PAYMENT menu (bank, petty cash) by using the search tool.
- **transmovement**: we can choose from data in STOCK CONTROL menu (delivery, inventory, etc.) by using the search tool.

### Auto create
When selected, the attribute in case of adding a new element (eg a new customer or employee is created) will automatically be created with the default value according to its type and also will be attached to the new element.

### Visible
Can appear or not (hidden value) on the entry forms.

### Readonly
The value of the attribute can not be changed in the program interface.

{{< hint info >}}

The additional data can be useful also to group our customers, products (value list data type).

{{< /hint >}}
