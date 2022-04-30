

---
title: gRPC API
type: docs
weight: 30
bookToC: true
---

# Overview

Nervatura [gRPC](https://grpc.io/) specification. For more examples, see 
[Nervatura example application](https://github.com/nervatura/nervatura-example)

| Method Name | Request Type | Response Type | Description |
| --- | --- | --- | --- |
| UserLogin | [RequestUserLogin](/docs/service/grpc#requestuserlogin) | [ResponseUserLogin](/docs/service/grpc#responseuserlogin) | Logs in user by username and password |
| UserPassword | [RequestUserPassword](/docs/service/grpc#requestuserpassword) | [ResponseEmpty](/docs/service/grpc#responseempty) | User (employee or customer) password change. |
| TokenLogin | [RequestEmpty](/docs/service/grpc#requestempty) | [ResponseTokenLogin](/docs/service/grpc#responsetokenlogin) | JWT token auth. |
| TokenRefresh | [RequestEmpty](/docs/service/grpc#requestempty) | [ResponseTokenRefresh](/docs/service/grpc#responsetokenrefresh) | Refreshes JWT token by checking at database whether refresh token exists. |
| TokenDecode | [RequestTokenDecode](/docs/service/grpc#requesttokendecode) | [ResponseTokenDecode](/docs/service/grpc#responsetokendecode) | Decoded JWT token but doesn't validate the signature. |
| Get | [RequestGet](/docs/service/grpc#requestget) | [ResponseGet](/docs/service/grpc#responseget) | Get returns one or more records |
| Update | [RequestUpdate](/docs/service/grpc#requestupdate) | [ResponseUpdate](/docs/service/grpc#responseupdate) | Add/update one or more items |
| Delete | [RequestDelete](/docs/service/grpc#requestdelete) | [ResponseEmpty](/docs/service/grpc#responseempty) | Delete an item |
| View | [RequestView](/docs/service/grpc#requestview) | [ResponseView](/docs/service/grpc#responseview) | Run raw SQL queries in safe mode |
| Function | [RequestFunction](/docs/service/grpc#requestfunction) | [ResponseFunction](/docs/service/grpc#responsefunction) | Call a server-side function |
| Report | [RequestReport](/docs/service/grpc#requestreport) | [ResponseReport](/docs/service/grpc#responsereport) | Create and download a Nervatura Report |
| ReportList | [RequestReportList](/docs/service/grpc#requestreportlist) | [ResponseReportList](/docs/service/grpc#responsereportlist) | List all available Nervatura Report. Admin user group membership required. |
| ReportInstall | [RequestReportInstall](/docs/service/grpc#requestreportinstall) | [ResponseReportInstall](/docs/service/grpc#responsereportinstall) | Install a report to the database. Admin user group membership required. |
| ReportDelete | [RequestReportDelete](/docs/service/grpc#requestreportdelete) | [ResponseEmpty](/docs/service/grpc#responseempty) | Delete a report from the database. Admin user group membership required. |
| DatabaseCreate | [RequestDatabaseCreate](/docs/service/grpc#requestdatabasecreate) | [ResponseDatabaseCreate](/docs/service/grpc#responsedatabasecreate) | Create a new Nervatura database |
 
 
<br />

---
# Table of Contents


- Messages
    - [Address](/docs/service/grpc#address)
    - [Barcode](/docs/service/grpc#barcode)
    - [Contact](/docs/service/grpc#contact)
    - [Currency](/docs/service/grpc#currency)
    - [Customer](/docs/service/grpc#customer)
    - [Deffield](/docs/service/grpc#deffield)
    - [Employee](/docs/service/grpc#employee)
    - [Event](/docs/service/grpc#event)
    - [Fieldvalue](/docs/service/grpc#fieldvalue)
    - [Groups](/docs/service/grpc#groups)
    - [Item](/docs/service/grpc#item)
    - [Link](/docs/service/grpc#link)
    - [Log](/docs/service/grpc#log)
    - [MetaData](/docs/service/grpc#metadata)
    - [Movement](/docs/service/grpc#movement)
    - [Numberdef](/docs/service/grpc#numberdef)
    - [Pattern](/docs/service/grpc#pattern)
    - [Payment](/docs/service/grpc#payment)
    - [Place](/docs/service/grpc#place)
    - [Price](/docs/service/grpc#price)
    - [Product](/docs/service/grpc#product)
    - [Project](/docs/service/grpc#project)
    - [Rate](/docs/service/grpc#rate)
    - [RequestDatabaseCreate](/docs/service/grpc#requestdatabasecreate)
    - [RequestDelete](/docs/service/grpc#requestdelete)
    - [RequestEmpty](/docs/service/grpc#requestempty)
    - [RequestFunction](/docs/service/grpc#requestfunction)
    - [RequestFunction.ValuesEntry](/docs/service/grpc#requestfunctionvaluesentry)
    - [RequestGet](/docs/service/grpc#requestget)
    - [RequestReport](/docs/service/grpc#requestreport)
    - [RequestReport.FiltersEntry](/docs/service/grpc#requestreportfiltersentry)
    - [RequestReportDelete](/docs/service/grpc#requestreportdelete)
    - [RequestReportInstall](/docs/service/grpc#requestreportinstall)
    - [RequestReportList](/docs/service/grpc#requestreportlist)
    - [RequestTokenDecode](/docs/service/grpc#requesttokendecode)
    - [RequestUpdate](/docs/service/grpc#requestupdate)
    - [RequestUpdate.Item](/docs/service/grpc#requestupdateitem)
    - [RequestUpdate.Item.KeysEntry](/docs/service/grpc#requestupdateitemkeysentry)
    - [RequestUpdate.Item.ValuesEntry](/docs/service/grpc#requestupdateitemvaluesentry)
    - [RequestUserLogin](/docs/service/grpc#requestuserlogin)
    - [RequestUserPassword](/docs/service/grpc#requestuserpassword)
    - [RequestView](/docs/service/grpc#requestview)
    - [RequestView.Query](/docs/service/grpc#requestviewquery)
    - [ResponseDatabaseCreate](/docs/service/grpc#responsedatabasecreate)
    - [ResponseEmpty](/docs/service/grpc#responseempty)
    - [ResponseFunction](/docs/service/grpc#responsefunction)
    - [ResponseGet](/docs/service/grpc#responseget)
    - [ResponseGet.Value](/docs/service/grpc#responsegetvalue)
    - [ResponseReport](/docs/service/grpc#responsereport)
    - [ResponseReportInstall](/docs/service/grpc#responsereportinstall)
    - [ResponseReportList](/docs/service/grpc#responsereportlist)
    - [ResponseReportList.Info](/docs/service/grpc#responsereportlistinfo)
    - [ResponseRows](/docs/service/grpc#responserows)
    - [ResponseRows.Item](/docs/service/grpc#responserowsitem)
    - [ResponseRows.Item.ValuesEntry](/docs/service/grpc#responserowsitemvaluesentry)
    - [ResponseTokenDecode](/docs/service/grpc#responsetokendecode)
    - [ResponseTokenLogin](/docs/service/grpc#responsetokenlogin)
    - [ResponseTokenRefresh](/docs/service/grpc#responsetokenrefresh)
    - [ResponseUpdate](/docs/service/grpc#responseupdate)
    - [ResponseUserLogin](/docs/service/grpc#responseuserlogin)
    - [ResponseView](/docs/service/grpc#responseview)
    - [ResponseView.ValuesEntry](/docs/service/grpc#responseviewvaluesentry)
    - [Tax](/docs/service/grpc#tax)
    - [Tool](/docs/service/grpc#tool)
    - [Trans](/docs/service/grpc#trans)
    - [UiAudit](/docs/service/grpc#uiaudit)
    - [UiMenu](/docs/service/grpc#uimenu)
    - [UiMenufields](/docs/service/grpc#uimenufields)
    - [UiMessage](/docs/service/grpc#uimessage)
    - [UiPrintqueue](/docs/service/grpc#uiprintqueue)
    - [UiReport](/docs/service/grpc#uireport)
    - [UiUserconfig](/docs/service/grpc#uiuserconfig)
    - [Value](/docs/service/grpc#value)
  


- Enums
    - [DataType](/docs/service/grpc#datatype)
    - [ReportOrientation](/docs/service/grpc#reportorientation)
    - [ReportOutput](/docs/service/grpc#reportoutput)
    - [ReportSize](/docs/service/grpc#reportsize)
    - [ReportType](/docs/service/grpc#reporttype)
  


- [Scalar Value Types](#scalar-value-types)

<br />

---
# Messages



## Address
RequestUpdate Key->ID keys:

- ```id```: Value is a generated unique key identifier: *{nervatype}/{refnumber}~{rownumber}*. The *rownumber* is the order of multiple *{nervatype}/{refnumber}* keys. For example: ```customer/DMCUST/00001~1```

- ```nervatype```: Valid values: *customer, employee, event, place, product, project, tool, trans*

- ```ref_id```: Valid values: *customer/{custnumber}, employee/{empnumber}, event/{calnumber}, place/{planumber}, product/{partnumber}, 
project/{pronumber}, tool/{serial}, trans/{transnumber}*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| nervatype | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'nervatype') |
| ref_id | [ int64](/docs/service/grpc#int64) | Reference to [Customer](#customer).id, [Employee](#employee).id, [Event](#event).id, [Place](#place).id, [Product](#product).id, [Project](#project).id, [Tool](#tool).id, [Trans](#trans).id |
| country | [ string](/docs/service/grpc#string) |  |
| state | [ string](/docs/service/grpc#string) |  |
| zipcode | [ string](/docs/service/grpc#string) |  |
| city | [ string](/docs/service/grpc#string) |  |
| street | [ string](/docs/service/grpc#string) |  |
| notes | [ string](/docs/service/grpc#string) |  |
| metadata | [repeated MetaData](/docs/service/grpc#metadata) | Address meta data |

<br />



## Barcode
RequestUpdate Key->ID keys:

- ```id```: Barcode *code*

- ```barcodetype```: Valid values: *CODE_128, CODE_39, EAN_13, EAN_8, QR*

- ```product_id```: Product *partnumber*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| code | [ string](/docs/service/grpc#string) | Each product can be connected to any number of bar codes, but the code must remain unique to ensure that the product is clearly identifiable. |
| product_id | [ int64](/docs/service/grpc#int64) | Reference to [Product](#product).id |
| description | [ string](/docs/service/grpc#string) | Comment related to the barcode. Informal, has no role in identification. |
| barcodetype | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id  (only where groupname = 'barcodetype'). |
| qty | [ double](/docs/service/grpc#double) | The actual amount of the products identified by the barcode. For example can be used for packaged goods, tray packaging. |
| defcode | [ bool](/docs/service/grpc#bool) | If more than one bar code is assigned, this will be the default. Because of the uniqueness of the barcode the product is always clearly identifiable, but in reverse case (eg. in case the barcode should be printed on a document) we must assign one being the default for that product. |

<br />



## Contact
RequestUpdate Key->ID keys:

- ```id```: The value is a generated constant key identifier: *{nervatype}/{refnumber}~{rownumber}*. The rownumber is the order of multiple *{nervatype}/{refnumber}* keys.

- ```nervatype```: Valid values: *customer, employee, event, place, product, project, tool, trans*

- ```ref_id```: Valid values: *customer/{custnumber}, employee/{empnumber}, event/{calnumber}, place/{planumber}, product/{partnumber}, project/{pronumber}, tool/{serial}, trans/{transnumber}*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| nervatype | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'nervatype') |
| ref_id | [ int64](/docs/service/grpc#int64) | Reference to [Customer](#customer).id, [Employee](#employee).id, [Event](#event).id, [Place](#place).id, [Product](#product).id, [Project](#project).id, [Tool](#tool).id, [Trans](#trans).id |
| firstname | [ string](/docs/service/grpc#string) |  |
| surname | [ string](/docs/service/grpc#string) |  |
| status | [ string](/docs/service/grpc#string) |  |
| phone | [ string](/docs/service/grpc#string) |  |
| fax | [ string](/docs/service/grpc#string) |  |
| mobil | [ string](/docs/service/grpc#string) |  |
| email | [ string](/docs/service/grpc#string) |  |
| notes | [ string](/docs/service/grpc#string) |  |
| metadata | [repeated MetaData](/docs/service/grpc#metadata) | Contact meta data |

<br />



## Currency
RequestUpdate Key->ID keys:

- ```id```: Currency *curr*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| curr | [ string](/docs/service/grpc#string) | The ISO 4217 code of the currency. |
| description | [ string](/docs/service/grpc#string) | The name of the currency. |
| digit | [ int64](/docs/service/grpc#int64) | The number of decimal places used for recording and rounding by the program. Default: 2 |
| defrate | [ double](/docs/service/grpc#double) | Default Rate. You can specify an exchange rate vs. the default currency, which will be used by the reports. |
| cround | [ int64](/docs/service/grpc#int64) | Rounding value for cash. Could be used in case the smallest banknote in circulation for that certain currency is not 1. |
| metadata | [repeated MetaData](/docs/service/grpc#metadata) | Currency meta data |

<br />



## Customer
RequestUpdate Key->ID keys:

- ```id```: Customer *custnumber*

- ```custtype```: Valid values: *own, company, private, other*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| custtype | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'custtype') |
| custnumber | [ string](/docs/service/grpc#string) | Unique ID. If you set it to numberdef, it will be generated at the first data save. The format and value of the next data in row is taken from the numberdef (numberkey = custnumber) data series. |
| custname | [ string](/docs/service/grpc#string) | Full name of the customer |
| taxnumber | [ string](/docs/service/grpc#string) |  |
| account | [ string](/docs/service/grpc#string) |  |
| notax | [ bool](/docs/service/grpc#bool) | Tax-free |
| terms | [ int64](/docs/service/grpc#int64) | Payment per. |
| creditlimit | [ double](/docs/service/grpc#double) | Customer's credit limit. Data is used by financial reports. |
| discount | [ double](/docs/service/grpc#double) | If new product line is added (offer, order, invoice etc.) all products will receive the discount percentage specified in this field. If the product has a separate customer price, the value specified here will not be considered by the program. |
| notes | [ string](/docs/service/grpc#string) |  |
| inactive | [ bool](/docs/service/grpc#bool) |  |
| metadata | [repeated MetaData](/docs/service/grpc#metadata) | Customer meta data |

<br />



## Deffield
RequestUpdate Key->ID keys:

- ```id```: Deffield *fieldname*

- ```nervatype```: Valid values: *address, barcode, contact, currency, customer, employee, event, item, link, log, movement, payment, price, place, product, project, rate, tax, tool, trans, setting*

- ```subtype```: All groupvalue from Groups, where groupname equal *custtype, placetype, protype, toolgroup, transtype*

- ```fieldtype```: Valid values: *bool, date, time, float, integer, string, valuelist, notes, urlink, password, customer, tool, transitem, transmovement, transpayment, product, project, employee, place*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| fieldname | [ string](/docs/service/grpc#string) |  |
| nervatype | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'nervatype') |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _subtype.subtype | [optional int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (where groupname in ('custtype','placetype','  protype','toolgroup','transtype')) |
| fieldtype | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id  (only where groupname = 'fieldtype') |
| description | [ string](/docs/service/grpc#string) |  |
| valuelist | [ string](/docs/service/grpc#string) | If fieldtype=valuelist: valid values are listed, separated by ~ |
| addnew | [ bool](/docs/service/grpc#bool) | When selected, the attribute in case of adding a new element (eg a new customer or employee is created) will automatically be created with the default value according to its type and also will be attached to the new element. |
| visible | [ bool](/docs/service/grpc#bool) | Can appear or not (hidden value) on the entry forms |
| readonly | [ bool](/docs/service/grpc#bool) | The value of the attribute can not be changed in the program interface |

<br />



## Employee
RequestUpdate Key->ID keys:

- ```id```: Employee *empnumber*

- ```usergroup```: All groupvalue from Groups, where groupname equal usergroup

- ```department```: All groupvalue from Groups, where groupname equal department


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| empnumber | [ string](/docs/service/grpc#string) | Unique ID. If you set it to numberdef, it will be generated at the first data save. The format and value of the next data in row is taken from the numberdef (numberkey = empnumber) data series. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _username.username | [optional string](/docs/service/grpc#string) | Database login name. Should be unique on database level. |
| usergroup | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'usergroup') |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _startdate.startdate | [optional string](/docs/service/grpc#string) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _enddate.enddate | [optional string](/docs/service/grpc#string) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _department.department | [optional int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'department') |
| registration_key | [ string](/docs/service/grpc#string) |  |
| inactive | [ bool](/docs/service/grpc#bool) |  |
| metadata | [repeated MetaData](/docs/service/grpc#metadata) | Employee meta data |

<br />



## Event
RequestUpdate Key->ID keys:

- ```id```: Event *calnumber*

- ```nervatype```: Valid values: *customer, employee, place, product, project, tool, trans*

- ```ref_id```: Valid values: *customer/{custnumber}, employee/{empnumber}, place/{planumber}, product/{partnumber}, project/{pronumber}, tool/{serial}, trans/{transnumber}*

- ```eventgroup```: All groupvalue from Groups, where groupname equal eventgroup


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| calnumber | [ string](/docs/service/grpc#string) | Unique ID. If you set it to numberdef, it will be generated at the first data save. The format and value of the next data in row is taken from the numberdef (numberkey = calnumber) data series. |
| nervatype | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'nervatype') |
| ref_id | [ int64](/docs/service/grpc#int64) | Reference to [Customer](#customer).id, [Employee](#employee).id, [Place](#place).id, [Product](#product).id, [Project](#project).id, [Tool](#tool).id, [Trans](#trans).id |
| uid | [ string](/docs/service/grpc#string) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _eventgroup.eventgroup | [optional int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'eventgroup') |
| fromdate | [ string](/docs/service/grpc#string) | Datetime |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _todate.todate | [optional string](/docs/service/grpc#string) | Datetime |
| subject | [ string](/docs/service/grpc#string) |  |
| place | [ string](/docs/service/grpc#string) |  |
| description | [ string](/docs/service/grpc#string) |  |
| metadata | [repeated MetaData](/docs/service/grpc#metadata) | Event meta data |

<br />



## Fieldvalue
RequestUpdate Key->ID keys:

- ```id```: The value is a generated constant key identifier: *{refnumber}~~{fieldname}~{rownumber}*. The rownumber is the order of multiple *{refnumber}~~{fieldname}* keys.

- ```ref_id```: Valid values: *{nervatype}/{refnumber}*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| fieldname | [ string](/docs/service/grpc#string) | Reference to [Deffield](#deffield).fieldname. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _ref_id.ref_id | [optional int64](/docs/service/grpc#int64) | Reference to any type.id where type = [Deffield](#deffield).nervatype. If it is null then nervatype = setting. |
| value | [ string](/docs/service/grpc#string) |  |
| notes | [ string](/docs/service/grpc#string) |  |

<br />



## Groups
RequestUpdate Key->ID keys:

- ```id```: Group *groupname~groupvalue*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| groupname | [ string](/docs/service/grpc#string) |  |
| groupvalue | [ string](/docs/service/grpc#string) |  |
| description | [ string](/docs/service/grpc#string) |  |
| inactive | [ bool](/docs/service/grpc#bool) |  |

<br />



## Item
RequestUpdate Key->ID keys:

- ```id```: Trans and Item *transnumber~rownumber*

- ```trans_id```: Trans *transnumber*

- ```product_id```: Product *partnumber*

- ```tax_id```: Tax *taxcode*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| trans_id | [ int64](/docs/service/grpc#int64) | Reference to [trans](#trans).id |
| product_id | [ int64](/docs/service/grpc#int64) | Reference to [product](#product).id |
| unit | [ string](/docs/service/grpc#string) |  |
| qty | [ double](/docs/service/grpc#double) |  |
| fxprice | [ double](/docs/service/grpc#double) |  |
| netamount | [ double](/docs/service/grpc#double) |  |
| discount | [ double](/docs/service/grpc#double) |  |
| tax_id | [ int64](/docs/service/grpc#int64) | Reference to [Tax](#tax).id |
| vatamount | [ double](/docs/service/grpc#double) |  |
| amount | [ double](/docs/service/grpc#double) |  |
| description | [ string](/docs/service/grpc#string) |  |
| deposit | [ bool](/docs/service/grpc#bool) |  |
| ownstock | [ double](/docs/service/grpc#double) |  |
| actionprice | [ bool](/docs/service/grpc#bool) |  |
| metadata | [repeated MetaData](/docs/service/grpc#metadata) | Item meta data |

<br />



## Link
RequestUpdate Key->ID keys:

- ```id```: *{nervatype_1}~{refnumber_1}~~{nervatype_2}~{refnumber_2}*

- ```nervatype_1```: All groupvalue from Groups, where groupname equal nervatype

- ```ref_id_1```: *{nervatype_1}/{refnumber_1}*

- ```nervatype_2```: All groupvalue from Groups, where groupname equal nervatype


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| nervatype_1 | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'nervatype') |
| ref_id_1 | [ int64](/docs/service/grpc#int64) | Reference to {nervatype}.id |
| nervatype_2 | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'nervatype') |
| ref_id_2 | [ int64](/docs/service/grpc#int64) | Reference to {nervatype}.id |
| metadata | [repeated MetaData](/docs/service/grpc#metadata) | Link meta data |

<br />



## Log
RequestUpdate Key->ID keys:

- ```id```: *{empnumber}~{crdate}'*

- ```employee_id```: Employee *empnumber*

- ```ref_id```: *{nervatype}/{refnumber}*

- ```nervatype```: All groupvalue from Groups, where groupname equal nervatype

- ```logstate```: Valid values: update, closed, deleted, print, login, logout


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| employee_id | [ int64](/docs/service/grpc#int64) | Reference to [Employee](#employee).id |
| crdate | [ string](/docs/service/grpc#string) | Date-time |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _nervatype.nervatype | [optional int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'nervatype') |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _ref_id.ref_id | [optional int64](/docs/service/grpc#int64) | Reference to {nervatype}.id |
| logstate | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'logstate') |
| metadata | [repeated MetaData](/docs/service/grpc#metadata) | Log meta data |

<br />



## MetaData



| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| fieldname | [ string](/docs/service/grpc#string) | Reference to [Deffield](#deffield).fieldname. |
| fieldtype | [ string](/docs/service/grpc#string) | Reference to [Deffield](#deffield).fieldtype. |
| value | [ string](/docs/service/grpc#string) |  |
| notes | [ string](/docs/service/grpc#string) |  |

<br />



## Movement
RequestUpdate Key->ID keys:

- ```id```: Trans and Item *transnumber~rownumber*

- ```trans_id```: Trans *transnumber*

- ```product_id```: Product *partnumber*

- ```movetype```: Valid values: *inventory, tool, plan, head*

- ```tool_id```: Tool *serial*

- ```place_id```: Place *planumber*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| trans_id | [ int64](/docs/service/grpc#int64) | Reference to [Trans](#trans).id |
| shippingdate | [ string](/docs/service/grpc#string) | Date-time |
| movetype | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'movetype') |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _product_id.product_id | [optional int64](/docs/service/grpc#int64) | Reference to [Product](#product).id |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _tool_id.tool_id | [optional int64](/docs/service/grpc#int64) | Reference to [Tool](#tool).id |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _place_id.place_id | [optional int64](/docs/service/grpc#int64) | Reference to [Place](#place).id |
| qty | [ double](/docs/service/grpc#double) |  |
| description | [ string](/docs/service/grpc#string) |  |
| shared | [ bool](/docs/service/grpc#bool) |  |
| metadata | [repeated MetaData](/docs/service/grpc#metadata) | Movement meta data |

<br />



## Numberdef
RequestUpdate Key->ID keys:

- ```id```: Numberdef *numberkey*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| numberkey | [ string](/docs/service/grpc#string) | Unique key |
| prefix | [ string](/docs/service/grpc#string) | The text prefix of the identifier. It can be any length, but usage of special characters, spaces in the text is not recommended. |
| curvalue | [ int64](/docs/service/grpc#int64) | The current status of the counter, the next sequence number will be one value higher than this one. It is possible to re-set the counter, but the uniqueness must be ensured in all cases! |
| isyear | [ bool](/docs/service/grpc#bool) | If selected, the created identifier will contain the year. |
| sep | [ string](/docs/service/grpc#string) | The separator character in the identifier. Default: "/" |
| len | [ int64](/docs/service/grpc#int64) | The value field is arranged in such length to the right and filled with zeros. |
| description | [ string](/docs/service/grpc#string) |  |
| visible | [ bool](/docs/service/grpc#bool) |  |
| readonly | [ bool](/docs/service/grpc#bool) |  |
| orderby | [ int64](/docs/service/grpc#int64) |  |

<br />



## Pattern
RequestUpdate Key->ID keys:

- ```id```: Pattern *description*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| description | [ string](/docs/service/grpc#string) |  |
| transtype | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'transtype') |
| notes | [ string](/docs/service/grpc#string) |  |
| defpattern | [ bool](/docs/service/grpc#bool) |  |

<br />



## Payment
RequestUpdate Key->ID keys:

- ```id```: Trans and Item *transnumber~rownumber*

- ```trans_id```: Trans *transnumber*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| trans_id | [ int64](/docs/service/grpc#int64) | Reference to [Trans](#trans).id |
| paiddate | [ string](/docs/service/grpc#string) |  |
| amount | [ double](/docs/service/grpc#double) |  |
| notes | [ string](/docs/service/grpc#string) |  |
| metadata | [repeated MetaData](/docs/service/grpc#metadata) | Payment meta data |

<br />



## Place
RequestUpdate Key->ID keys:

- ```id```: Place *planumber*

- ```placetype```: Valid values: *bank, cash, warehouse, other*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| planumber | [ string](/docs/service/grpc#string) | Unique ID. If you set it to numberdef, it will be generated at the first data save. The format and value of the next data in row is taken from the numberdef (numberkey = planumber) data series. |
| placetype | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'placetype') |
| description | [ string](/docs/service/grpc#string) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _curr.curr | [optional string](/docs/service/grpc#string) |  |
| defplace | [ bool](/docs/service/grpc#bool) |  |
| notes | [ string](/docs/service/grpc#string) |  |
| inactive | [ bool](/docs/service/grpc#bool) |  |
| metadata | [repeated MetaData](/docs/service/grpc#metadata) | Place meta data |

<br />



## Price
RequestUpdate Key->ID keys:

- ```id```: Price *partnumber~validfrom~curr~qty*

- ```product_id```: Product *partnumber*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| product_id | [ int64](/docs/service/grpc#int64) | Reference to [Product](#product).id |
| validfrom | [ string](/docs/service/grpc#string) | Start of validity, mandatory data. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _validto.validto | [optional string](/docs/service/grpc#string) | End of validity, can be left empty. |
| curr | [ string](/docs/service/grpc#string) |  |
| qty | [ double](/docs/service/grpc#double) | Price ranges can also be specified, thus different price can be set for a smaller and bigger quantity of the same product. The quantity should be used as the lower threshold, ie. this should be the minimum quantity for the price set. |
| pricevalue | [ double](/docs/service/grpc#double) | Price value |
| vendorprice | [ bool](/docs/service/grpc#bool) | Supplier (if marked) or customer price. By default the customer price. |
| metadata | [repeated MetaData](/docs/service/grpc#metadata) | Price meta data |

<br />



## Product
RequestUpdate Key->ID keys:

- ```id```: Product *partnumber*

- ```protype```: Valid values: *item, service*

- ```tax_id```: Tax *taxcode*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| partnumber | [ string](/docs/service/grpc#string) | Unique ID. If you set it to numberdef, it will be generated at the first data save. The format and value of the next data in row is taken from the numberdef (numberkey = partnumber) data series. |
| protype | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'protype') |
| description | [ string](/docs/service/grpc#string) | The full name of the product or short description. |
| unit | [ string](/docs/service/grpc#string) |  |
| tax_id | [ int64](/docs/service/grpc#int64) | Reference to [Tax](#tax).id |
| notes | [ string](/docs/service/grpc#string) |  |
| webitem | [ bool](/docs/service/grpc#bool) |  |
| inactive | [ bool](/docs/service/grpc#bool) |  |
| metadata | [repeated MetaData](/docs/service/grpc#metadata) | Product meta data |

<br />



## Project
RequestUpdate Key->ID keys:

- ```id```: Project *pronumber*

- ```customer_id```: Tax *custnumber*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| pronumber | [ string](/docs/service/grpc#string) | Unique ID. If you set it to numberdef, it will be generated at the first data save. The format and value of the next data in row is taken from the numberdef (numberkey = pronumber) data series. |
| description | [ string](/docs/service/grpc#string) | The name of the project. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _customer_id.customer_id | [optional int64](/docs/service/grpc#int64) | Reference to [Customer](#customer).id |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _startdate.startdate | [optional string](/docs/service/grpc#string) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _enddate.enddate | [optional string](/docs/service/grpc#string) |  |
| notes | [ string](/docs/service/grpc#string) |  |
| inactive | [ bool](/docs/service/grpc#bool) |  |
| metadata | [repeated MetaData](/docs/service/grpc#metadata) | Project meta data |

<br />



## Rate
RequestUpdate Key->ID keys:

- ```id```: Rate *ratetype~ratedate~curr~planumber*

- ```place_id```: Place *planumber*

- ```ratetype```: Valid values: *rate, buy, sell,average*

- ```rategroup```: all groupvalue from Groups, where groupname equal rategroup


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| ratetype | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'ratetype') |
| ratedate | [ string](/docs/service/grpc#string) |  |
| curr | [ string](/docs/service/grpc#string) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _place_id.place_id | [optional int64](/docs/service/grpc#int64) | Reference to [Place](#place).id |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _rategroup.rategroup | [optional int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'rategroup') |
| ratevalue | [ double](/docs/service/grpc#double) | Rate or interest value |
| metadata | [repeated MetaData](/docs/service/grpc#metadata) | Rate meta data |

<br />



## RequestDatabaseCreate
New database props.


| Field | Type | Description |
| ----- | ---- | ----------- |
| alias | [ string](/docs/service/grpc#string) | Alias name of the database |
| demo | [ bool](/docs/service/grpc#bool) | Create a DEMO database |

<br />



## RequestDelete
Delete parameters


| Field | Type | Description |
| ----- | ---- | ----------- |
| nervatype | [ DataType](/docs/service/grpc#datatype) |  |
| id | [ int64](/docs/service/grpc#int64) | The object ID |
| key | [ string](/docs/service/grpc#string) | Use Key instead of ID |

<br />



## RequestEmpty
No parameters




## RequestFunction



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](/docs/service/grpc#string) | Server function name |
| values | [map RequestFunction.ValuesEntry](/docs/service/grpc#requestfunctionvaluesentry) | The array of parameter values |

<br />



## RequestFunction.ValuesEntry



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](/docs/service/grpc#string) |  |
| value | [ Value](/docs/service/grpc#value) |  |

<br />



## RequestGet



| Field | Type | Description |
| ----- | ---- | ----------- |
| nervatype | [ DataType](/docs/service/grpc#datatype) |  |
| metadata | [ bool](/docs/service/grpc#bool) |  |
| ids | [repeated int64](/docs/service/grpc#int64) |  |
| filter | [repeated string](/docs/service/grpc#string) |  |

<br />



## RequestReport



| Field | Type | Description |
| ----- | ---- | ----------- |
| reportkey | [ string](/docs/service/grpc#string) | Example : ntr_invoice_en |
| orientation | [ ReportOrientation](/docs/service/grpc#reportorientation) |  |
| size | [ ReportSize](/docs/service/grpc#reportsize) |  |
| output | [ ReportOutput](/docs/service/grpc#reportoutput) |  |
| type | [ ReportType](/docs/service/grpc#reporttype) |  |
| refnumber | [ string](/docs/service/grpc#string) | Example : DMINV/00001 |
| template | [ string](/docs/service/grpc#string) | Custom report JSON template |
| filters | [map RequestReport.FiltersEntry](/docs/service/grpc#requestreportfiltersentry) |  |

<br />



## RequestReport.FiltersEntry



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](/docs/service/grpc#string) |  |
| value | [ Value](/docs/service/grpc#value) |  |

<br />



## RequestReportDelete



| Field | Type | Description |
| ----- | ---- | ----------- |
| reportkey | [ string](/docs/service/grpc#string) | Example : ntr_invoice_en |

<br />



## RequestReportInstall
Admin user group membership required.


| Field | Type | Description |
| ----- | ---- | ----------- |
| reportkey | [ string](/docs/service/grpc#string) | Example : ntr_invoice_en |

<br />



## RequestReportList



| Field | Type | Description |
| ----- | ---- | ----------- |
| label | [ string](/docs/service/grpc#string) |  |

<br />



## RequestTokenDecode



| Field | Type | Description |
| ----- | ---- | ----------- |
| value | [ string](/docs/service/grpc#string) | Access token code. |

<br />



## RequestUpdate



| Field | Type | Description |
| ----- | ---- | ----------- |
| nervatype | [ DataType](/docs/service/grpc#datatype) |  |
| items | [repeated RequestUpdate.Item](/docs/service/grpc#requestupdateitem) |  |

<br />



## RequestUpdate.Item



| Field | Type | Description |
| ----- | ---- | ----------- |
| values | [map RequestUpdate.Item.ValuesEntry](/docs/service/grpc#requestupdateitemvaluesentry) |  |
| keys | [map RequestUpdate.Item.KeysEntry](/docs/service/grpc#requestupdateitemkeysentry) |  |

<br />



## RequestUpdate.Item.KeysEntry



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](/docs/service/grpc#string) |  |
| value | [ Value](/docs/service/grpc#value) |  |

<br />



## RequestUpdate.Item.ValuesEntry



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](/docs/service/grpc#string) |  |
| value | [ Value](/docs/service/grpc#value) |  |

<br />



## RequestUserLogin



| Field | Type | Description |
| ----- | ---- | ----------- |
| username | [ string](/docs/service/grpc#string) | Employee username or Customer custnumber (email or phone number) |
| password | [ string](/docs/service/grpc#string) |  |
| database | [ string](/docs/service/grpc#string) | Optional. Default value: NT_DEFAULT_ALIAS |

<br />



## RequestUserPassword



| Field | Type | Description |
| ----- | ---- | ----------- |
| password | [ string](/docs/service/grpc#string) | New password |
| confirm | [ string](/docs/service/grpc#string) | New password confirmation |
| username | [ string](/docs/service/grpc#string) | Optional. Only if different from the logged in user. Admin user group membership required. |
| custnumber | [ string](/docs/service/grpc#string) | Optional. Only if different from the logged in user. Admin user group membership required. |

<br />



## RequestView
Only "select" queries and functions can be executed. Changes to the data are not saved in the database.


| Field | Type | Description |
| ----- | ---- | ----------- |
| options | [repeated RequestView.Query](/docs/service/grpc#requestviewquery) | The array of Query object |

<br />



## RequestView.Query



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](/docs/service/grpc#string) | Give the query a unique name |
| text | [ string](/docs/service/grpc#string) | The SQL query as a string |
| values | [repeated Value](/docs/service/grpc#value) | The array of parameter values |

<br />



## ResponseDatabaseCreate
Result log data


| Field | Type | Description |
| ----- | ---- | ----------- |
| details | [ ResponseRows](/docs/service/grpc#responserows) |  |

<br />



## ResponseEmpty
Does not return content.




## ResponseFunction



| Field | Type | Description |
| ----- | ---- | ----------- |
| value | [ bytes](/docs/service/grpc#bytes) |  |

<br />



## ResponseGet



| Field | Type | Description |
| ----- | ---- | ----------- |
| values | [repeated ResponseGet.Value](/docs/service/grpc#responsegetvalue) |  |

<br />



## ResponseGet.Value



| Field | Type | Description |
| ----- | ---- | ----------- |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.address | [ Address](/docs/service/grpc#address) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.barcode | [ Barcode](/docs/service/grpc#barcode) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.contact | [ Contact](/docs/service/grpc#contact) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.currency | [ Currency](/docs/service/grpc#currency) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.customer | [ Customer](/docs/service/grpc#customer) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.deffield | [ Deffield](/docs/service/grpc#deffield) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.employee | [ Employee](/docs/service/grpc#employee) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.event | [ Event](/docs/service/grpc#event) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.fieldvalue | [ Fieldvalue](/docs/service/grpc#fieldvalue) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.groups | [ Groups](/docs/service/grpc#groups) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.item | [ Item](/docs/service/grpc#item) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.link | [ Link](/docs/service/grpc#link) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.log | [ Log](/docs/service/grpc#log) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.movement | [ Movement](/docs/service/grpc#movement) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.numberdef | [ Numberdef](/docs/service/grpc#numberdef) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.pattern | [ Pattern](/docs/service/grpc#pattern) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.payment | [ Payment](/docs/service/grpc#payment) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.place | [ Place](/docs/service/grpc#place) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.price | [ Price](/docs/service/grpc#price) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.product | [ Product](/docs/service/grpc#product) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.project | [ Project](/docs/service/grpc#project) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.rate | [ Rate](/docs/service/grpc#rate) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.tax | [ Tax](/docs/service/grpc#tax) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.tool | [ Tool](/docs/service/grpc#tool) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.trans | [ Trans](/docs/service/grpc#trans) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.ui_audit | [ UiAudit](/docs/service/grpc#uiaudit) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.ui_menu | [ UiMenu](/docs/service/grpc#uimenu) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.ui_menufields | [ UiMenufields](/docs/service/grpc#uimenufields) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.ui_message | [ UiMessage](/docs/service/grpc#uimessage) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.ui_printqueue | [ UiPrintqueue](/docs/service/grpc#uiprintqueue) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.ui_report | [ UiReport](/docs/service/grpc#uireport) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.ui_userconfig | [ UiUserconfig](/docs/service/grpc#uiuserconfig) |  |

<br />



## ResponseReport



| Field | Type | Description |
| ----- | ---- | ----------- |
| value | [ bytes](/docs/service/grpc#bytes) |  |

<br />



## ResponseReportInstall



| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) | Returns a new report ID. |

<br />



## ResponseReportList
Returns all installable files from the NT_REPORT_DIR directory (empty value: all available built-in Nervatura Reports)


| Field | Type | Description |
| ----- | ---- | ----------- |
| items | [repeated ResponseReportList.Info](/docs/service/grpc#responsereportlistinfo) |  |

<br />



## ResponseReportList.Info



| Field | Type | Description |
| ----- | ---- | ----------- |
| reportkey | [ string](/docs/service/grpc#string) |  |
| repname | [ string](/docs/service/grpc#string) |  |
| description | [ string](/docs/service/grpc#string) |  |
| label | [ string](/docs/service/grpc#string) |  |
| reptype | [ string](/docs/service/grpc#string) |  |
| filename | [ string](/docs/service/grpc#string) |  |
| installed | [ bool](/docs/service/grpc#bool) |  |

<br />



## ResponseRows



| Field | Type | Description |
| ----- | ---- | ----------- |
| items | [repeated ResponseRows.Item](/docs/service/grpc#responserowsitem) |  |

<br />



## ResponseRows.Item



| Field | Type | Description |
| ----- | ---- | ----------- |
| values | [map ResponseRows.Item.ValuesEntry](/docs/service/grpc#responserowsitemvaluesentry) |  |

<br />



## ResponseRows.Item.ValuesEntry



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](/docs/service/grpc#string) |  |
| value | [ Value](/docs/service/grpc#value) |  |

<br />



## ResponseTokenDecode
Access token claims.


| Field | Type | Description |
| ----- | ---- | ----------- |
| username | [ string](/docs/service/grpc#string) |  |
| database | [ string](/docs/service/grpc#string) |  |
| exp | [ double](/docs/service/grpc#double) | JWT expiration time |
| iss | [ string](/docs/service/grpc#string) |  |

<br />



## ResponseTokenLogin
Token user properties


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| username | [ string](/docs/service/grpc#string) |  |
| empnumber | [ string](/docs/service/grpc#string) |  |
| usergroup | [ int64](/docs/service/grpc#int64) |  |
| scope | [ string](/docs/service/grpc#string) |  |
| department | [ string](/docs/service/grpc#string) |  |

<br />



## ResponseTokenRefresh



| Field | Type | Description |
| ----- | ---- | ----------- |
| value | [ string](/docs/service/grpc#string) | Access token code. |

<br />



## ResponseUpdate
If the ID (or Key) value is missing, it creates a new item.


| Field | Type | Description |
| ----- | ---- | ----------- |
| values | [repeated int64](/docs/service/grpc#int64) | Returns the all new/updated IDs values. |

<br />



## ResponseUserLogin



| Field | Type | Description |
| ----- | ---- | ----------- |
| token | [ string](/docs/service/grpc#string) | Access JWT token |
| engine | [ string](/docs/service/grpc#string) | Type of database |
| version | [ string](/docs/service/grpc#string) | Service version |

<br />



## ResponseView



| Field | Type | Description |
| ----- | ---- | ----------- |
| values | [map ResponseView.ValuesEntry](/docs/service/grpc#responseviewvaluesentry) | key - results map |

<br />



## ResponseView.ValuesEntry



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](/docs/service/grpc#string) |  |
| value | [ ResponseRows](/docs/service/grpc#responserows) |  |

<br />



## Tax
RequestUpdate Key->ID keys:

- ```id```: Tax *taxcode*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| taxcode | [ string](/docs/service/grpc#string) | Unique ID. |
| description | [ string](/docs/service/grpc#string) |  |
| rate | [ double](/docs/service/grpc#double) | Rate or interest value |
| inactive | [ bool](/docs/service/grpc#bool) |  |
| metadata | [repeated MetaData](/docs/service/grpc#metadata) | Tax meta data |

<br />



## Tool
RequestUpdate Key->ID keys:

- ```id```: Tool *serial*

- ```toolgroup```: all groupvalue from Groups, where groupname equal toolgroup

- ```product_id```: Product *partnumber*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| serial | [ string](/docs/service/grpc#string) | Unique ID. If you set it to numberdef, it will be generated at the first data save. The format and value of the next data in row is taken from the numberdef (numberkey = serial) data series. |
| description | [ string](/docs/service/grpc#string) |  |
| product_id | [ int64](/docs/service/grpc#int64) | Reference to [Product](#product).id |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _toolgroup.toolgroup | [optional int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'toolgroup') |
| notes | [ string](/docs/service/grpc#string) |  |
| inactive | [ bool](/docs/service/grpc#bool) |  |
| metadata | [repeated MetaData](/docs/service/grpc#metadata) | Tool meta data |

<br />



## Trans
RequestUpdate Key->ID keys:

- ```id```: Trans *transnumber*

- ```transtype```: all groupvalue from Groups, where groupname equal transtype

- ```direction```: Valid values *in, out, transfer*

- ```customer_id```: Customer *custnumber*

- ```employee_id```: Employee *empnumber*

- ```department```: all groupvalue from Groups, where groupname equal department

- ```project_id```: Project *pronumber*

- ```place_id```: Place *planumber*

- ```paidtype```: all groupvalue from Groups, where groupname equal paidtype

- ```transtate```: all groupvalue from Groups, where groupname equal transtate


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| transnumber | [ string](/docs/service/grpc#string) | Unique ID. If you set it to numberdef, it will be generated at the first data save. The format and value of the next data in row is taken from the numberdef (numberkey = transnumber) data series. |
| transtype | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'transtype')[Groups](#groups).id |
| direction | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'direction') |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _ref_transnumber.ref_transnumber | [optional string](/docs/service/grpc#string) |  |
| crdate | [ string](/docs/service/grpc#string) |  |
| transdate | [ string](/docs/service/grpc#string) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _duedate.duedate | [optional string](/docs/service/grpc#string) | Date-time |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _customer_id.customer_id | [optional int64](/docs/service/grpc#int64) | Reference to [Customer](#customer).id |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _employee_id.employee_id | [optional int64](/docs/service/grpc#int64) | Reference to [Employee](#employee).id |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _department.department | [optional int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'department') |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _project_id.project_id | [optional int64](/docs/service/grpc#int64) | Reference to [Project](#project).id |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _place_id.place_id | [optional int64](/docs/service/grpc#int64) | Reference to [Place](#place).id |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _paidtype.paidtype | [optional int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'paidtype') |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _curr.curr | [optional string](/docs/service/grpc#string) |  |
| notax | [ bool](/docs/service/grpc#bool) |  |
| paid | [ bool](/docs/service/grpc#bool) |  |
| acrate | [ double](/docs/service/grpc#double) |  |
| notes | [ string](/docs/service/grpc#string) |  |
| intnotes | [ string](/docs/service/grpc#string) |  |
| fnote | [ string](/docs/service/grpc#string) |  |
| transtate | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'transtate') |
| closed | [ bool](/docs/service/grpc#bool) |  |
| metadata | [repeated MetaData](/docs/service/grpc#metadata) | Trans meta data |

<br />



## UiAudit
RequestUpdate Key->ID keys:

- ```id```: UiAudit *{usergroup}~{nervatype}~{transtype}*

- ```usergroup```: all groupvalue from Groups, where groupname equal usergroup

- ```nervatype```: all groupvalue from Groups, where groupname equal nervatype

- ```subtype```: all groupvalue from Groups, where groupname equal transtype, movetype, protype, custtype, placetype

- ```inputfilter```: Valid values *disabled, readonly, update, all*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| usergroup | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'usergroup') |
| nervatype | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'nervatype') |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _subtype.subtype | [optional int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'movetype') |
| inputfilter | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'inputfilter') |
| supervisor | [ bool](/docs/service/grpc#bool) |  |

<br />



## UiMenu
RequestUpdate Key->ID keys:

- ```id```: UiMenu *menukey*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| menukey | [ string](/docs/service/grpc#string) |  |
| description | [ string](/docs/service/grpc#string) |  |
| modul | [ string](/docs/service/grpc#string) |  |
| icon | [ string](/docs/service/grpc#string) |  |
| method | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'method') |
| funcname | [ string](/docs/service/grpc#string) |  |
| address | [ string](/docs/service/grpc#string) |  |

<br />



## UiMenufields
RequestUpdate Key->ID keys:

- ```id```: UiMenufields *{menukey}~{fieldname}*

- ```menu_id```: UiMenu *menukey*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| menu_id | [ int64](/docs/service/grpc#int64) | Reference to [UiMenu](#UiMenu).id |
| fieldname | [ string](/docs/service/grpc#string) |  |
| description | [ string](/docs/service/grpc#string) |  |
| fieldtype | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'fieldtype') |
| orderby | [ int64](/docs/service/grpc#int64) |  |

<br />



## UiMessage
RequestUpdate Key->ID keys:

- ```id```: UiMessage *{secname}~{fieldname}~{lang}*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| secname | [ string](/docs/service/grpc#string) |  |
| fieldname | [ string](/docs/service/grpc#string) |  |
| lang | [ string](/docs/service/grpc#string) |  |
| msg | [ string](/docs/service/grpc#string) |  |

<br />



## UiPrintqueue



| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _nervatype.nervatype | [optional int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'nervatype') |
| ref_id | [ int64](/docs/service/grpc#int64) | Reference to {nervatype}.id |
| qty | [ double](/docs/service/grpc#double) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _employee_id.employee_id | [optional int64](/docs/service/grpc#int64) | Reference to [Employee](#employee).id |
| report_id | [ int64](/docs/service/grpc#int64) | Reference to [UiReport](#UiReport).id |
| crdate | [ string](/docs/service/grpc#string) | Date-time |

<br />



## UiReport
RequestUpdate Key->ID keys:

- ```id```: UiReport *reportkey*

- ```nervatype```: all groupvalue from Groups, where groupname equal nervatype

- ```transtype```: all groupvalue from Groups, where groupname equal transtype

- ```direction```: all groupvalue from Groups, where groupname equal direction

- ```filetype```: all groupvalue from Groups, where groupname equal filetype


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| reportkey | [ string](/docs/service/grpc#string) |  |
| nervatype | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'nervatype') |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _transtype.transtype | [optional int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'transtype') |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _direction.direction | [optional int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'direction') |
| repname | [ string](/docs/service/grpc#string) |  |
| description | [ string](/docs/service/grpc#string) |  |
| label | [ string](/docs/service/grpc#string) |  |
| filetype | [ int64](/docs/service/grpc#int64) | Reference to [Groups](#groups).id (only where groupname = 'filetype') |
| report | [ string](/docs/service/grpc#string) |  |

<br />



## UiUserconfig
RequestUpdate Key->ID keys:

- ```id```: *{empnumber}~{section}~{cfgroup}~{cfname}*

- ```employee_id```: Employee *{empnumber}*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](/docs/service/grpc#int64) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _employee_id.employee_id | [optional int64](/docs/service/grpc#int64) | Reference to [Employee](#employee).id |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _section.section | [optional string](/docs/service/grpc#string) |  |
| cfgroup | [ string](/docs/service/grpc#string) |  |
| cfname | [ string](/docs/service/grpc#string) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _cfvalue.cfvalue | [optional string](/docs/service/grpc#string) |  |
| orderby | [ int64](/docs/service/grpc#int64) |  |

<br />



## Value



| Field | Type | Description |
| ----- | ---- | ----------- |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.boolean | [ bool](/docs/service/grpc#bool) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.number | [ double](/docs/service/grpc#double) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.text | [ string](/docs/service/grpc#string) | google.protobuf.NullValue null = 4; |

<br />



---
# Enums



## DataType


| Name | Number | Description |
| ---- | ------ | ----------- |
| address | 0 | [Address](/docs/service/grpc#address) |
| barcode | 1 | [Barcode](/docs/service/grpc#barcode) |
| contact | 2 | [Contact](/docs/service/grpc#contact) |
| currency | 3 | [Currency](/docs/service/grpc#currency) |
| customer | 4 | [Customer](/docs/service/grpc#customer) |
| deffield | 5 | [Deffield](/docs/service/grpc#deffield) |
| employee | 6 | [Employee](/docs/service/grpc#employee) |
| event | 7 | [Event](/docs/service/grpc#event) |
| fieldvalue | 8 | [Fieldvalue](/docs/service/grpc#fieldvalue) |
| groups | 9 | [Groups](/docs/service/grpc#groups) |
| item | 10 | [Item](/docs/service/grpc#item) |
| link | 11 | [Link](/docs/service/grpc#link) |
| log | 12 | [Log](/docs/service/grpc#log) |
| movement | 13 | [Movement](/docs/service/grpc#movement) |
| numberdef | 14 | [Numberdef](/docs/service/grpc#numberdef) |
| pattern | 15 | [Pattern](/docs/service/grpc#pattern) |
| payment | 16 | [Payment](/docs/service/grpc#payment) |
| place | 17 | [Place](/docs/service/grpc#place) |
| price | 18 | [Price](/docs/service/grpc#price) |
| product | 19 | [Product](/docs/service/grpc#product) |
| project | 20 | [Project](/docs/service/grpc#project) |
| rate | 21 | [Rate](/docs/service/grpc#rate) |
| tax | 22 | [Tax](/docs/service/grpc#tax) |
| tool | 23 | [Tool](/docs/service/grpc#tool) |
| trans | 24 | [Trans](/docs/service/grpc#trans) |
| ui_audit | 25 | [UiAudit](/docs/service/grpc#uiaudit) |
| ui_menu | 26 | [UiMenu](/docs/service/grpc#uimenu) |
| ui_menufields | 27 | [UiMenufields](/docs/service/grpc#uimenufields) |
| ui_message | 28 | [UiMessage](/docs/service/grpc#uimessage) |
| ui_printqueue | 29 | [UiPrintqueue](/docs/service/grpc#uiprintqueue) |
| ui_report | 30 | [UiReport](/docs/service/grpc#uireport) |
| ui_userconfig | 31 | [UiUserconfig](/docs/service/grpc#uiuserconfig) |

<br />


## ReportOrientation


| Name | Number | Description |
| ---- | ------ | ----------- |
| portrait | 0 |  |
| landscape | 1 |  |

<br />


## ReportOutput


| Name | Number | Description |
| ---- | ------ | ----------- |
| auto | 0 |  |
| xml | 1 |  |
| data | 2 |  |
| base64 | 3 |  |

<br />


## ReportSize


| Name | Number | Description |
| ---- | ------ | ----------- |
| a3 | 0 |  |
| a4 | 1 |  |
| a5 | 2 |  |
| letter | 3 |  |
| legal | 4 |  |

<br />


## ReportType


| Name | Number | Description |
| ---- | ------ | ----------- |
| report_none | 0 |  |
| report_customer | 1 |  |
| report_employee | 2 |  |
| report_event | 3 |  |
| report_place | 4 |  |
| report_product | 5 |  |
| report_project | 6 |  |
| report_tool | 7 |  |
| report_trans | 8 |  |

<br />



---
# Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <div><h4 id="double" /></div><a name="double" /> double |  | double | double | float |
| <div><h4 id="int64" /></div><a name="int64" /> int64 | Uses variable-length encoding | int64 | long | int/long |
| <div><h4 id="bool" /></div><a name="bool" /> bool |  | bool | boolean | boolean |
| <div><h4 id="string" /></div><a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <div><h4 id="bytes" /></div><a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

