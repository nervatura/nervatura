---
title: Payment
type: docs
weight: 30
bookToC: true
---

## Input fields

### Document No.
Unique ID, generated at the first data save. The format and value of the next data in row is taken from the [**DOCUMENT NUMBERING**](/docs/client/settings/numberdef) (code = bank_transfer/cash) data series. 

### Reference No.
Other reference or bank statement number. Optional, its value can be freely defined.

### Creation
Date of creation. Automatic value, cannot be changed.

### Closed
*Technical* closing of the document.

{{< hint danger >}}

If set, document data become read only. **Marking is not revocable on the user interface!**

{{< /hint >}}

### State
Its value and editing possibility is linked to [**ACCESS RIGHTS**](/docs/client/settings/usergroup#supervisor) setting. Not used in current version.

### Comment
Remarks field

### Internal notes
Internal comments. Text defined in this field will not appear on the document

---
### BANK

### Account Date
The date of the statement. Field filling is compulsory.

### Bank Account
A bank account can be chosen from the search field from *bank* type items of [**PLACE**](/docs/client/settings/place#type).

---
### PETTY CASH

### Direction
OUT, IN. Its value cannot be changed after first save!

### Payment Date
The date of the transaction. Field filling is compulsory.

### Amount
The amount of inpayment or outpayment. Regardless of direction always with positive sign!

### Petty cash
A petty cash can be chosen from the search field from *cash* type items of [**PLACE**](/docs/client/settings/place#type).

### Employee No.
One of the items of [**EMPLOYEE**](/docs/client/resources/employee#employee-no.). Optional.

## Related data

### [**METADATA**](/docs/client/settings/metadata)
Unlimited number of supplementary data can be added.

### [**REPORT NOTES**](/docs/client/program/notes)
Editable remarks, data for reports.

### INVOICE
Invoices which are linked to bank statement items and to cash receipts.
- *Invoice No.*: The ID of the [**INVOICE**](/docs/client/document/document).
- *Currency*: The currency of related amount. Cannot be changed, set by the program based on currency of Invoice No.
- *Amount*: The payment amount. Can be lower than the total sum of the invoice (in case of payment in installments), and can be higher as well, and also can be an amount of opposite sign (net transaction).
- *Rate*: Exchange rate in case the payment currency differs from the invoice currency. Default value: 1.

### DOCUMENT ITEM
BANK. The line items of the statement.
- *Row ID*: Unique ID of the row item. Automatic value, cannot be changed.
- *Payment Date*: Accounting date of the amount. Field filling is compulsory.
- *Amount*: Amount with sign. Amounts with positive sign indicate the incoming, with negative signs the outgoing items.
- *Description*: Other data of the item.

## Operations

### COPY FROM
Create a new, *same transaction type* document on the basis of current document's data. 

The dates and information related to creation gets updated and the references from the original document will not be transferred either.

### REPORT
[**DATA EXPORT**](/docs/client/program/export)

### BOOKMARK
Set a bookmark for the record. Later can be loaded from bookmarks at any time.
