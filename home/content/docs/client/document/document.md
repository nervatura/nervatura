---
title: Document
type: docs
weight: 10
bookToC: true
---

## Overview

Nervatura was made for being able to store all data generated during ones working process. **OFFER** type forms are used to record the given (CUSTOMER) and also the received (SUPPLIER) offers/quotations. If any of the offers finally results in business deal, then its data are easy to use for further transaction types.

The **ORDER**, **WORKSHEET** and **RENTAL** transaction types play a significant role in the workflow. They gather all those goods/services that we either committed to deliver or we expect to be delivered by our suppliers. The data sheet screens therefore were designed to enable easy review and follow up of all physical movements of goods (inventory, asset movement), as well as the financial aspects related to a transaction.

The forms were optimized to cover three different transaction types.
- The **ORDER** is for the classic working method used in the majority of firms. We sell goods or services to our business partners or we are the ones provided with these by our suppliers.
- The **WORKSHEET** is useful when our activity is rather service related. This can also be linked to a product (stock movement), however mostly product is concerned not as the subject of a sales action but used as a material and thus rather it is part of the service itself.
- In case of **RENTAL** the focus moves from labor to time. We can follow that our own (or borrowed) products/tools under what conditions and at what time frame were used by our partners.

The **INVOICE** is used for recording and follow up of our financial obligations.

The **RECEIPT** is the simplified version of the above one, with following differences:
- Only CUSTOMER type (not available for SUPPLIER direction)
- Customer cannot be defined
- The fulfillment date (Receipt Date) and financial due date (Due Date) is always the same.
- By default it is counted as paid (Paid).
- Cannot be linked to a bank transfer or cash receipts.
- Cannot be linked to TOOL movements.

## Input fields

### Document No.
Unique ID, generated at the first data save. The format and value of the next data in row is taken from the [**DOCUMENT NUMBERING**](/docs/client/settings/numberdef) data series depending on the data type (CUSTOMER / SUPPLIER):
- Offer No. (code = offer_out/offer_in)
- Order No. (code = order_out/order_in)
- Worksheet No. (code = worksheet_out)
- Rental No. (code = rent_out/rent_in)
- Invoice No. (code = invoice_out/invoice_in)
- Receipt No. (code = receipt_out)

### Document Type
CUSTOMER or SUPPLIER direction. Its value cannot be changed after first save. For WORKSHEET and RECEIPT types currently only CUSTOMER type is allowed.

### Reference No.
Other reference number. Optional, its value can be freely defined.

{{< hint info >}}

In case the invoce is created by using **CORRECTIVE**, **CANCELLATION** actions, then the field value will be the number of that document which the data originate from. The two documents become connected, the value of the field cannot be changed.

In case of  **CREATE FROM** action it can be defined if the program should connect the two documents or not.

If **COPY FROM** action is used, this connection is not established, the field remains freely editable.

{{< /hint >}}

### State
Its value and editing possibility is linked to [**ACCESS RIGHTS**](/docs/client/settings/usergroup#supervisor) setting. Not used in current version.

---

### Creation 
Date of creation. Automatic value, cannot be changed.

### Offer Date, Valid Date 
OFFER. Freely definable transaction dates. Field filling is compulsory.

### Order Date, Delivery Date
ORDER. Freely definable transaction dates. Field filling is compulsory.

### Start Date, End Date
WORKSHEET, RENTAL. Freely definable transaction dates. Field filling is compulsory.

### Invoice Date, Due Date
INVOICE. Freely definable transaction dates. Field filling is compulsory.

### Receipt Date, Due Date
RECEIPT. Freely definable transaction dates. Field filling is compulsory.

---

### Customer Name
One of the items of the [**CUSTOMER**](/docs/client/resources/customer#customer-name). Mandatory to fill in, except in case of RECEIPT.

---

### Distance (km), Repair time (h), Total time (h), Justification
WORKSHEET. Usage is optional.

### Holidays, Bad machine, Other non-eligible, Justification
RENTAL. Usage is optional.

---

### Currency
One of the items of the [**CURRENCY**](/docs/client/settings/currency). Field filling is compulsory.

### Payment Days: 
OFFER, ORDER, WORKSHEET, RENTAL. Can be used for setting an individual payment due date related to invoicing data of a transaction. Optional, overrides the settings of [**Payment Period**](/docs/client/resources/customer#payment-period) as well as [**default deadline**](/docs/client/settings/setting#default-deadline).

### Acc. Rate
INVOICE, RECEIPT. To define accounting exchange rate. Optional.

### Released 
*Logical* closing of the document. Just indicates the action, has no technical significance.

### Closed
*Technical* closing of the document.

{{< hint danger >}}

If set, document data become read only. **Marking is not revocable on the user interface!**

{{< /hint >}}

---

### Payment
A valid item of the [**paidtype**](/docs/client/settings/groups) group should be defined. Field filling is compulsory.

### Department
A valid item of the [**department**](/docs/client/settings/groups) group should be defined. Optional.

### Employee No.
An item of the [**EMPLOYEE**](/docs/client/resources/employee#employee-no.). Optional.

### Project No.
An item of the [**PROJECT**](/docs/client/resources/project#project-no.). Optional.

### Comment
Remarks field.

### Internal notes
Internal comments. Text defined in this field will not appear on the document.

## Related data

### [**METADATA**](/docs/client/settings/metadata)
Unlimited number of supplementary data can be added.

### [**REPORT NOTES**](/docs/client/program/notes)
Editable remarks, data for reports.

### [**DOCUMENT ITEM**](/docs/client/document/item)
Unlimited number of item lines can be added.

---

### INVOICE
ORDER, WORKSHEET, RENTAL. Item lines of invoices related to the document.

{{< hint warning >}}

Invoices can be linked to the document with [**CREATE FROM**](/docs/client/document/document#create-from) action. For new invoices those invoice items of the document which have not been invoiced yet are offered by the program, these of course can be freely modified. The invoicing can be based on the line items of the document or related delivery notes (invoicing of completed deliveries) as well, however the two methods cannot be mixed within a single transaction.

{{< /hint >}}

### [**SHIPPING**](/docs/client/stock/shipping)
ORDER, WORKSHEET, RENTAL. The aggregate stock move of ***item type*** products of the document.
- *Item Product*: Product description of a line item.
- *Shipped Product*: Description of the product moved out from the warehouse actually. The two description fields can differ from each other when [Element Product](/docs/client/resources/product) products are linked to the products of the document.
- *Shipping Quantity*: The total amount of warehouse movements. The stock release is shown with negative, the entries with positive signs. Either way stock correction will also be summarized!

### [**TOOL MOVEMENT**](/docs/client/stock/waybill)
ORDER, WORKSHEET, RENTAL, INVOICE. Tool movements related to the document.

### PAYMENTS
INVOICE. [**Bank transfers**](/docs/client/document/payment#bank) and [**cash receipts**](/docs/client/document/payment#petty-cash) related to the document.
- *Payment No.*: The identity code of the [**bank transfers**](/docs/client/document/payment#bank) or [**cash receipts**](/docs/client/document/payment#petty-cash).
- *Currency*: The currency of related amount. Cannot be changed, set by the program based on currency of Payment No.
- *Amount*: The payment amount. Can be lower than the total sum of the invoice (in case of payment in installments), and can be higher as well, and also can be an amount of opposite sign (net transaction).
- *Rate*: Exchange rate in case the payment currency differs from the invoice currency. Default value: 1.

## Operations

### COPY FROM
Create a new, *same transaction type* document on the basis of current document's data. 

The dates and information related to creation gets updated and the references from the original document will not be transferred either.

### CREATE FROM
Create a new , different transaction type document on the basis of current document's data.

The following documents can be created:
- **OFFER**: OFFER, ORDER, WORKSHEET, RENTAL. Default: ORDER.
- **ORDER**: OFFER, ORDER, WORKSHEET, RENTAL, INVOICE, RECEIPT. Default: INVOICE.
- **WORKSHEET**: OFFER, ORDER, WORKSHEET, RENTAL, INVOICE, RECEIPT. Default: INVOICE.
- **RENTAL**: OFFER, ORDER, WORKSHEET, RENTAL, INVOICE, RECEIPT. Default: INVOICE.
- **INVOICE**: ORDER, WORKSHEET, RENTAL, INVOICE, RECEIPT. Default: ORDER.

The dates and information related to creation gets updated and the references from the original document will not be transferred either. You can specify the new document type (CUSTOMER - out, SUPPLIER - in), as well as the following parameters:
- **Set Reference No.**: connecting the two documents, setting the Reference No. field of the new document.
- **Invoiced amount deduction**: For ORDER, WORKSHEET, RENTAL types at INVOICE and RECEIPT creation. Not all item lines are copied, if selected, the already invoiced items will be deducted from the list.
- **Create based on delivery**: For ORDER, WORKSHEET, RENTAL types at INVOICE and RECEIPT creation. The item lines can be based on document items or delivery note items. The two methods cannot be mixed for a given document.

### CORRECTIVE
CUSTOMER INVOICE, RECEIPT. Corrective invoice creation.

New invoice will be created based on the document, the two will be connected. All line items are double-taken, with the pair of opposite sign, so the invoice total will be zero after it is created. If invoice correction is needed any of the lines can be simply deleted.

### CANCELLATION
CUSTOMER INVOICE, RECEIPT. Cancellation invoice creation.

New invoice will be created based on the document, the two will be connected. The new invoice will be the exact copy of the previous invoice, but the line items will be shown with opposite sign.  The initial invoice shall have CANCELED status.

### REPORT
[**DATA EXPORT**](/docs/client/program/export)

### BOOKMARK
Set a bookmark for the record. Later can be loaded from bookmarks at any time.
