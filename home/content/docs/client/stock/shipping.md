---
title: Shipping
type: docs
weight: 10
bookToC: true
---

## Overview

In Nervatura database a product can be moved into or moved out from a warehouse only based on a certain document or action. This ensures that the origin of all the delivered products can be reliably traced if necessary.

Warehouse traffic can be generated from items of ORDER, WORKSHEET and RENTAL type documents, as well as from corrective items of INVENTORY and TRANSFER documents. The SHIPPING is the stock control form of ORDER, WORKSHEET and RENTAL documents.

## Input fields

### CREATE DELIVERY

### Shipping Date
Stock movements (stock release or receipt) date. The value is valid for all line items.

### Warehouse
A warehouse can be chosen from the search field from *warehouse* type items of [**PLACE**](/docs/client/settings/place#type). The value is valid for all line items.

### Product No., Product Name, Unit
Information about the product. Data serves just information purposes, cannot be modified in this section.

### Batch No.
The quantity is put into this group, or taken off from here. Usage is optional.

### Quantity
Quantity to be released or received. By default, the missing quantity compared to the whole document, but it can be changed freely. Directional sign should not be provided, based on the document type it is handled automatically by the program.

{{< hint info >}}

With **Batch No.** several virtual groups of the same product can be created within a certain warehouse. The program breaks down the actual stock into these groups. The warehouse turnover related to the documents (order, worksheet, rental) shows which items were released from (or were added to) which group.

Its usage is highly efficient and easy way of basic monitoring of goods. The content and structure of the Batch No. allows defining information related to a specific product group.

The two most common cases are defining of supplier/origin of the product and possibly the information related to expiry date. The first one is useful if a particular product is purchased from several sources. The group identifier helps to track even that from exactly which supplier was the product delivered to our customer received originally.

The expiry information can be useful in manging the actual stock. By using it we can have a more accurate picture of the content of the stock,  at delivery we can prioritize certain groups (eg. products soon to expire), or can make pricing decisions (campaign products).

{{< /hint >}}

## Related data

### DOCUMENT NO., DELIVERY TYPE, CUSTOMER NAME
Related fields of ORDER, WORKSHEET and RENTAL documents, their value cannot be changed. The value of DELIVERY TYPE is based on the type of the document. CUSTOMER type documents will receive SHIPPING OUT (stock release), SUPPLIER types will get SHIPPING IN (goods receipt) value.

### DOCUMENT ITEMS
- **Shipped Product**: Product No.|Product Name. The list includes the item type products of the document. If there are [Element Product](/docs/client/resources/product) items linked, then the components of those can be seen here.
- ***Doc. Qty***: The document quantity. If there are [Element Product](/docs/client/resources/product) items linked, then proportional quantity is shown.
- **Turnover**: The aggregate turnover of products related to the document. If movements of the opposite direction have been made as well (eg. correction), then the net result.
- **Difference**: The difference between the two previous columns. Amount other than zero is indicated in red by the program.

### [DELIVERY ITEMS](/docs/client/stock/delivery)
The items of former warehouse movements related to the document (stock release or receipt).

## Operations

### ADD, ALL DIFFERENCE
The Difference column value will be used for generating delivery note line item.

The **ADD** applies only to a specific row, the **ALL DIFFERENCE** will perform the operation for all non-zero quantities.

### STOCK
Current stock of a given product.

The total inventory of the product in Warehouse/Unit/Batch No.

### CREATE
Warehouse movements execution, delivery note generation.
