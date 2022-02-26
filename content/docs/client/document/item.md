---
title: Document Item
type: docs
weight: 20
bookToC: true
---

## Input fields

### Product No.
One of the items of the [**PRODUCT**](/docs/client/resources/product#product-no.). Field filling is compulsory. Its value can be defined using a search field but can be set with the help of a barcode as well.

{{< hint info >}}

When the product has been chosen the program will display on the form all available (already defined) data about the product: Description, Unit, Tax Rate, Discount, Unit Price.

The price of the product is searched by taking into account the following criteria: (product and transaction side), current day, currency, volume, customer. With specified amount (if zero, then adjusted to one) calculates the Net Amount and Amount prices based on the Unit Price.

{{< /hint >}}

### Description, Unit
By default, the Product Name and Unit, but ones are free to specify any data.

### Tax Rate
A valid item of the **TAX**. By default, the Product Tax. If it is changed the program will recalculate the values.

### Own Stock
Technical field for creating specific reports. Usage is optional. 

### Option
OFFER. To indicate special/optional items of the offer. Usage is optional. 

### Deposit
INVOICE. To handle the item lines of the advance payment. 

### Discount
By default the value of the [**CUSTOMER**](/docs/client/resources/customer#discount) discount field. If the product has a valid customer price, the value will always be zero! 

### Quantity, Discount, Unit Price, Net Amount, Amount
If any of the fields is changed, based on the indicated value the program will recalculate the others. 