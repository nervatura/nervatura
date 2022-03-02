---
title: Tool
type: docs
weight: 40
bookToC: true
---

## Overview

The stock inventory list shows the availability of goods by quantities. However in some cases it might be needed to have the possibility to follow up a certain item individually. In this case it gets a unique identifier, specific data can be connected to it through additional data, events can be assigned, and [**also its move can be tracked**](/docs/client/stock/waybill).

{{< hint info >}}

A typical example could be the management of company cars, which requires recording of many different data types and events. Similarly to laptops and mobile phones, in which case it is possible to track also which user owns the phone at a certain time, and not only have the static recording of the characteristics of the device.

{{< /hint >}}

## Input fields

### Serial
The unique identifier of the TOOL. Can be set when the new TOOL is created, before first save. Cannot be modified later. If left empty, the program will automatically generate one before saving. The format and value of the next data in row is taken from the [**DOCUMENT NUMBERING**](/docs/client/settings/numberdef) (code = serial) data series.

### Description
The name of the TOOL or a brief description.

### Group
In this field a valid element of [**toolgroup**](/docs/client/settings/groups) group should be given.

### Product No.
The [**product type**](/docs/client/resources/product#product-type) of the TOOL. Mandatory.

### Comment
Remarks field.

## Related data

### [**METADATA**](/docs/client/settings/metadata)
Unlimited number of supplementary data can be added.

### [**EVENT**](/docs/client/resources/event)
Unlimited number of events can be added.

## Operations

### REPORT
[**DATA EXPORT**](/docs/client/program/export)

### BOOKMARK
Set a bookmark for the record. Later can be loaded from bookmarks at any time.
