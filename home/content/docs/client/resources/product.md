---
title: Product
type: docs
weight: 20
bookToC: true
---

## Input fields

### Product No.
The unique ID of the product. Can be set when a new product is added, before the first save. Later cannot be modified any more.<br />
If the field is left empty, the program automatically will generate one when the save happens. The format and value of the next data in row is taken from the [**DOCUMENT NUMBERING**](/docs/client/settings/numberdef) (code = partnumber) data series.

### Product Name
The full name of the product or short description.

### Product Type
item, service

### Unit, Tax, Web Item
Other product attributes. When new product line is added (eg. offer, order, invoice etc.) the default value for Unit and Tax fields will be the data set here.

### Comment
Remarks field.

## Related data

### [**METADATA**](/docs/client/settings/metadata)
Unlimited number of supplementary data can be added.

{{< hint info >}}

The **Element Product** is an automatically generated, special additional data. With its help we can create special, so called virtual products. In practice this means that we create groups of some products which also exist independently and then treat these groups as separate products. We can add them to the product list, name, number and price them, use them in an offer, order or issue an invoice for them.<br />
If we produce or assemble a product, then it will usually be created physically as well. The components it is made of will be withdrawn from the warehouse and a new product is added to our stock.<br /> 
The virtual product may contain existing products and services, but these products will still be tracked separately (although they are not necessarily distributed separately). If the virtual product is sold (or rented out) the order and the invoice will be issued for the virtual product, but the delivery note will list all the content of the virtual product (group) and our stocks will also be reduced by these individual products.<br />
The virtual products are **item** types, and when created the content should be added at additional data by using **Element Product**. Here also the quantity of a certain product included in the pack can be specified.

{{< /hint >}}

### BARCODE (QR CODE)
Each product can be connected to any number of bar codes, but the code must remain unique to ensure that the product is clearly identifiable.

{{< hint danger >}}

The reader and the reports of the program do not support all possible code types!

{{< /hint >}}
- **Barcode**: The barcode (QR code) in alphanumeric form. The program will only check the uniqueness, not the compliance with the type of the barcode.
- **Type**: AZTEC, CODABAR, CODE_128, CODE_39, CODE_93, DATA_MATRIX, EAN_13, EAN_8, ITF, MSI, PDF417, QR_CODE, RSS_EXPANDED, RSS14, UPC_A, UPC_E
- **Quantity**: The actual amount of the products identified by the barcode. For example can be used for packaged goods, tray packaging.
- **Default**: If more than one bar code is assigned, this will be the default. Because of the uniqueness of the barcode the product is always clearly identifiable, but in reverse case (eg. in case the barcode should be printed on a document) we must assign one being the default for that product.
- **Description**: Comment related to the barcode. Informal, has no role in identification.

### [**EVENT**](/docs/client/resources/event)
Unlimited number of events can be added.

### PRICE
The product can be connected to several prices based on different schemes, which are stored by the program historically, for previous periods as well. When a new product line is added (ie. offer, order, invoice, etc.) the program will select these as defaults, also they can be used in reports, for example when making Pricelists.

{{< hint danger >}}

The Nervatura Client only supports the usage of list prices and customer specific prices! The Nervatura Framework also uses prices linked to a certain customer group.

{{< /hint >}}

- **Start Date**: Start of validity, mandatory data.
- **End Date**: End of validity, can be left empty.
- **Supplier**: Supplier (if marked) or customer price. By default the customer price.
- **Currency**: By default the value given in [*default currency*](/docs/client/settings/setting#default-currency).
- **Quantity**: Price ranges can also be specified, thus different price can be set for a smaller and bigger quantity of the same product. The quantity should be used as the lower threshold, ie. this should be the minimum quantity for the price set. Default value: 0.
- **Price**
- **Customer Name**: The price is valid only for this specific customer. If the field is empty the general list price is used.

{{< hint info >}}

The program for pricing considers the following conditions: (product and transaction side), date, currency, quantity, customer. Also follows the below pricing validity rules:<br />
The valid price is the one set in the given currency, having a valid from date which is earlier or equal to the date when the price is used and valid till a date which is bigger or equal to the date when the price is used, or it is missing. The quantity set for the price should be lower or equal to the given quantity. These criteria should be met to have a general list price for the product. Additionally, if these criteria are also met by a specific customer price then this will be applied, otherwise the valid price will be the general list price.

{{< /hint >}}

{{< hint info >}}

It is possible to set a price valid for a certain period. The product should have a price with an earlier valid from date and undefined valid till date. If then a later valid from date and also a valid till date is given, than the product will have the set price only during this period. In periods before and after the original price with undefined ending will be valid.

{{< /hint >}}

## Operations

### REPORT
[**DATA EXPORT**](/docs/client/program/export)

### BOOKMARK
Set a bookmark for the record. Later can be loaded from bookmarks at any time.
