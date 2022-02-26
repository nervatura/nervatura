---
title: Default settings
type: docs
weight: 10
bookToC: true
---

## Overview

The database settings are not necessarily related to Nervatura Touch program. **Could affect the server's settings, the method of data storage or even operation of other programs!**<br />
For example, the default values set here will be used by the [**Nervatura HTTP API**](/docs/service/api) functions as well.

## Some important settings

### business year
Set the current fiscal year. This as used by [**DOCUMENT NUMBERING**](/docs/client/settings/numberdef) as the value of the Year

### default bank, default petty cash, default warehouse
These will be the default values for new input forms. The value for [**Place No.**](/docs/client/settings/place) should be set in the field.

### default currency
This will be the default value for new input forms. The field should include the [**currency**](/docs/client/settings/currency) code.

### default deadline
Number of days. This will be the basis for calculating the default [**duedate**](/docs/client/document/document#duedate) field in case of a new invoice. The value can be effected by the [**CUSTOMER**](/docs/client/resources/customer) settings of the invoice.

### default paidtype
This will be the default value for new input forms. A valid element of the [**paidtype**](/docs/client/settings/groups) group should be set in the field.

### default taxcode
This will be the default value for new line items. The field should include the *code* field of the **TAX**.

### default unit
Setting applies to a default unit field of a new [**PRODUCT**](/docs/client/resources/product)
