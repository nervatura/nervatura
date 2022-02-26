---
title: Formula
type: docs
weight: 60
bookToC: true
---

## Overview

Production of a product in some cases might require significant amount of raw materials. It can occur as well that the same product can also be produced from different components. In these cases, the work with [**PRODUCTION**](/docs/client/stock/production) forms can be sped up by FORMULA data sheets prepared for the products to be manufactured.

These can also be imagined as production recipes for a product. All the raw material requirements of the product can be specified for a given quantity. Then on the new production form it is enough to load one of the FORMULA data lines related to the product and the program will fill in the list of raw materials in proportion to the production quantity.

## Input fields

### Document No.
Unique ID, generated at the first data save. The format and value of the next data in row is taken from the [**DOCUMENT NUMBERING**](/docs/client/settings/numberdef) (code = formula_transfer) data series.

### Creation
Date of creation. Automatic value, cannot be changed.

### Closed
*Technical* closing of the document.

{{< hint danger >}}

If set, document data become read only. **Marking is not revocable on the user interface!**

{{< /hint >}}

### State
Its value and editing possibility is linked to [**ACCESS RIGHTS**](/docs/client/settings/usergroup#supervisor) setting. Not used in current version.

### Product No.
One of the items of [**PRODUCT**](/docs/client/resources/product), the subject of the production, the form is related to this product. Mandatory. Value can be defined with a search field or with the barcode of the product.

### Reference No.
Other reference number. Optional, its value can be freely defined.

### Quantity
The listed raw material quantity relates to the product quantity shown in this field.
                
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
- **Product No.**: Raw material used for production. Mandatory.
- **Quantity**: The given raw material's quantity that is required for production. Should not contain any directional sign, will automatically interpreted as stock decrease by the program. If negative value is given, then it will increase the stock level (eg. by-product formation).
- **Not Split**: If selected, raw material quantity in proportion to production requirement will be rounded up to whole numbers (integer).

{{< hint info >}}

Example: The FORMULA row data list was prepared for production of 8 products. There is one raw material on the list which was marked as "not split". 12 products are going to be produced based on the raw material requirement specified in the FORMULA, then the program will offer 2 pieces of the marked material (vs. 1.5 pieces).

{{< /hint >}}

- **Warehouse**: The raw material usage will be deducted from the stock of this warehouse. Optional. If the field is not filled, then the default warehouse will be selected for production.
- **Comment**: Other remarks.

## Operations

### COPY FROM
Create a new, *same transaction type* document on the basis of current document's data. 

The dates and information related to creation gets updated and the references from the original document will not be transferred either.

### REPORT
[**DATA EXPORT**](/docs/client/program/export)

### BOOKMARK
Set a bookmark for the record. Later can be loaded from bookmarks at any time.
