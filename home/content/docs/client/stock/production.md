---
title: Production
type: docs
weight: 50
bookToC: true
---

## Overview

PRODUCTION helps to produce new products from raw material stored in the warehouses. The inventory level will decrease by the amount of raw materials used, and the manufactured new product quantity will appear in stock. The data sheet tracks the material usage. If other costs, resource usage (eg. used energy, time spent, tool used etc.) is needed to be tracked, then through additional data it can be linked to [**WORKSHEET**](/docs/client/document/document) forms as well.

## Input fields

### Document No.
Unique ID, generated at the first data save. The format and value of the next data in row is taken from the [**DOCUMENT NUMBERING**](/docs/client/settings/numberdef) (code = production_transfer) data series.

### Creation
Date of creation. Automatic value, cannot be changed.

### Closed
*Technical* closing of the document.

{{< hint danger >}}

If set, document data become read only. **Marking is not revocable on the user interface!**

{{< /hint >}}

### State
Its value and editing possibility is linked to [**ACCESS RIGHTS**](/docs/client/settings/usergroup#supervisor) setting. Not used in current version.

### Start Date
Production starting date. Mandatory.

### End Date
Production ending date. The day, when the quantity will appear in the stock of the given warehouse.

### Product No.
One of the items of [**PRODUCT**](/docs/client/resources/product), the subject of the production, the form is related to this product. Mandatory. Value can be defined with a search field or with the barcode of the product.

### Reference No.
Other reference number. Optional, its value can be freely defined.

### Warehouse
The produced quantity will be shown in this warehouse. A warehouse can be chosen from the search field from *warehouse* type items of [**PLACE**](/docs/client/settings/place#type). Mandatory.

### Batch No.
The quantity is put into this group. Usage is optional.

### Quantity
Quantity of produced goods.
  
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
- **Shipping Date**: The usage date of raw materials. The program will deduct the 
quantity from the given warehouse stock on this day.
- **Warehouse**: The raw material quantity will be deducted from this warehouse by the program. A warehouse can be chosen from the search field from *warehouse* type items of [**PLACE**](/docs/client/settings/place#type). Mandatory.
- **Product No.**: Raw material used for production. Mandatory.
- **Batch No.**: The quantity is taken off from this group. Usage is optional.
- **Quantity**: The used quantity. Should not contain any directional sign, will 
automatically interpreted as stock decrease by the program. If negative value is given, 
then it will increase the stock level (eg. by-product formation).

## Operations

### COPY FROM
Create a new, *same transaction type* document on the basis of current document's data. 

The dates and information related to creation gets updated and the references from the original document will not be transferred either.

### LOAD FORMULA
The program will display the [**FORMULA**](/docs/client/stock/formula) templates available for the produced product in a drop down list. The list of raw materials will be loaded according to the chosen template. **If there were lines already, those will be deleted!** Raw material quantity is taken proportionally to the product quantity to be produced.

### REPORT
[**DATA EXPORT**](/docs/client/program/export)

### BOOKMARK
Set a bookmark for the record. Later can be loaded from bookmarks at any time.
