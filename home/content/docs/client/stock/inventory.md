---
title: Inventory control
type: docs
weight: 30
bookToC: true
---

## Overview

Warehouse stock is influenced by a wide variety of things: merchandise arrives from our supplier, we do deliver to a customer, new products are prepared from raw material, or goods are simply delivered from one warehouse to another. The program offers a user interface for all cases listed before.

However in daily operation, even more situations can happen, which effect the actual quantity: scrapping, shortage of goods (ie. theft), physical damage, etc. These cases can be handled by Nervatura INVENTORY document type, and Inventory Control is used as input form.

## Input fields

### Document No.
Unique ID, generated at the first data save. The format and value of the next data in row is taken from the [**DOCUMENT NUMBERING**](/docs/client/settings/numberdef) (code = inventory_transfer) data series.

### Warehouse
A warehouse can be chosen from the search field from *warehouse* type items of [**PLACE**](/docs/client/settings/place#type).

### Reference No.
Optional, its value can be freely defined.

### State
Its value and editing possibility is linked to [**ACCESS RIGHTS**](/docs/client/settings/usergroup#supervisor) setting. Not used in current version.

### Creation
Date of creation. Automatic value, cannot be changed.

### Inventory Date
Date of warehouse movement, mandatory.

### Closed
*Technical* closing of the document.

{{< hint danger >}}

If set, document data become read only. **Marking is not revocable on the user interface!**

{{< /hint >}}

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
- **Shipping Date**: Stock movements (stock release or receipt) date. The value can not be modified at line item level.
- **Product No.**: One of the *items* of [**PRODUCT**](/docs/client/resources/product#product-type). Mandatory. Value can be defined with a search field or with the barcode of the product.
- **Batch No.**: The quantity is put into this group, or taken off from here. Usage is optional.
- **Quantity**: The negative quantity decreases, the positive increases the warehouse stock level.

## Operations

### COPY FROM
Create a new, *same transaction type* document on the basis of current document's data. 

The dates and information related to creation gets updated and the references from the original document will not be transferred either.

### CANCELLATION
Cancellation document creation.

New statement will be created based on the document, the two will be connected. The new document will be the exact copy of the previous document, but the line items will be shown with opposite sign.

### REPORT
[**DATA EXPORT**](/docs/client/program/export)

### BOOKMARK
Set a bookmark for the record. Later can be loaded from bookmarks at any time.