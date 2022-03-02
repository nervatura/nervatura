

# Nervatura gRPC API

Nervatura [gRPC](https://grpc.io/) specification. For more examples, see [Node.js sample application](https://github.com/nervatura/nervatura-express)

| Method Name | Request Type | Response Type | Description |
| --- | --- | --- | --- |
| UserLogin | [RequestUserLogin](#requestuserlogin) | [ResponseUserLogin](#responseuserlogin) | Logs in user by username and password |
| UserPassword | [RequestUserPassword](#requestuserpassword) | [ResponseEmpty](#responseempty) | User (employee or customer) password change. |
| TokenLogin | [RequestEmpty](#requestempty) | [ResponseTokenLogin](#responsetokenlogin) | JWT token auth. |
| TokenRefresh | [RequestEmpty](#requestempty) | [ResponseTokenRefresh](#responsetokenrefresh) | Refreshes JWT token by checking at database whether refresh token exists. |
| TokenDecode | [RequestTokenDecode](#requesttokendecode) | [ResponseTokenDecode](#responsetokendecode) | Decoded JWT token but doesn't validate the signature. |
| Get | [RequestGet](#requestget) | [ResponseGet](#responseget) | Get returns one or more records |
| Update | [RequestUpdate](#requestupdate) | [ResponseUpdate](#responseupdate) | Add/update one or more items |
| Delete | [RequestDelete](#requestdelete) | [ResponseEmpty](#responseempty) | Delete an item |
| View | [RequestView](#requestview) | [ResponseView](#responseview) | Run raw SQL queries in safe mode |
| Function | [RequestFunction](#requestfunction) | [ResponseFunction](#responsefunction) | Call a server-side function |
| Report | [RequestReport](#requestreport) | [ResponseReport](#responsereport) | Create and download a Nervatura Report |
| ReportList | [RequestReportList](#requestreportlist) | [ResponseReportList](#responsereportlist) | List all available Nervatura Report. Admin user group membership required. |
| ReportInstall | [RequestReportInstall](#requestreportinstall) | [ResponseReportInstall](#responsereportinstall) | Install a report to the database. Admin user group membership required. |
| ReportDelete | [RequestReportDelete](#requestreportdelete) | [ResponseEmpty](#responseempty) | Delete a report from the database. Admin user group membership required. |
| DatabaseCreate | [RequestDatabaseCreate](#requestdatabasecreate) | [ResponseDatabaseCreate](#responsedatabasecreate) | Create a new Nervatura database |
 
 
<br />

---
# Table of Contents


- Messages
    - [Address](#address)
    - [Barcode](#barcode)
    - [Contact](#contact)
    - [Currency](#currency)
    - [Customer](#customer)
    - [Deffield](#deffield)
    - [Employee](#employee)
    - [Event](#event)
    - [Fieldvalue](#fieldvalue)
    - [Groups](#groups)
    - [Item](#item)
    - [Link](#link)
    - [Log](#log)
    - [MetaData](#metadata)
    - [Movement](#movement)
    - [Numberdef](#numberdef)
    - [Pattern](#pattern)
    - [Payment](#payment)
    - [Place](#place)
    - [Price](#price)
    - [Product](#product)
    - [Project](#project)
    - [Rate](#rate)
    - [RequestDatabaseCreate](#requestdatabasecreate)
    - [RequestDelete](#requestdelete)
    - [RequestEmpty](#requestempty)
    - [RequestFunction](#requestfunction)
    - [RequestFunction.ValuesEntry](#requestfunctionvaluesentry)
    - [RequestGet](#requestget)
    - [RequestReport](#requestreport)
    - [RequestReport.FiltersEntry](#requestreportfiltersentry)
    - [RequestReportDelete](#requestreportdelete)
    - [RequestReportInstall](#requestreportinstall)
    - [RequestReportList](#requestreportlist)
    - [RequestTokenDecode](#requesttokendecode)
    - [RequestUpdate](#requestupdate)
    - [RequestUpdate.Item](#requestupdateitem)
    - [RequestUpdate.Item.KeysEntry](#requestupdateitemkeysentry)
    - [RequestUpdate.Item.ValuesEntry](#requestupdateitemvaluesentry)
    - [RequestUserLogin](#requestuserlogin)
    - [RequestUserPassword](#requestuserpassword)
    - [RequestView](#requestview)
    - [RequestView.Query](#requestviewquery)
    - [ResponseDatabaseCreate](#responsedatabasecreate)
    - [ResponseEmpty](#responseempty)
    - [ResponseFunction](#responsefunction)
    - [ResponseGet](#responseget)
    - [ResponseGet.Value](#responsegetvalue)
    - [ResponseReport](#responsereport)
    - [ResponseReportInstall](#responsereportinstall)
    - [ResponseReportList](#responsereportlist)
    - [ResponseReportList.Info](#responsereportlistinfo)
    - [ResponseRows](#responserows)
    - [ResponseRows.Item](#responserowsitem)
    - [ResponseRows.Item.ValuesEntry](#responserowsitemvaluesentry)
    - [ResponseTokenDecode](#responsetokendecode)
    - [ResponseTokenLogin](#responsetokenlogin)
    - [ResponseTokenRefresh](#responsetokenrefresh)
    - [ResponseUpdate](#responseupdate)
    - [ResponseUserLogin](#responseuserlogin)
    - [ResponseView](#responseview)
    - [ResponseView.ValuesEntry](#responseviewvaluesentry)
    - [Tax](#tax)
    - [Tool](#tool)
    - [Trans](#trans)
    - [UiAudit](#uiaudit)
    - [UiMenu](#uimenu)
    - [UiMenufields](#uimenufields)
    - [UiMessage](#uimessage)
    - [UiPrintqueue](#uiprintqueue)
    - [UiReport](#uireport)
    - [UiUserconfig](#uiuserconfig)
    - [Value](#value)
  


- Enums
    - [DataType](#datatype)
    - [ReportOrientation](#reportorientation)
    - [ReportOutput](#reportoutput)
    - [ReportSize](#reportsize)
    - [ReportType](#reporttype)
  


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
| id | [ int64](#int64) |  |
| nervatype | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'nervatype') |
| ref_id | [ int64](#int64) | Reference to [Customer](#customer).id, [Employee](#employee).id, [Event](#event).id, [Place](#place).id, [Product](#product).id, [Project](#project).id, [Tool](#tool).id, [Trans](#trans).id |
| country | [ string](#string) |  |
| state | [ string](#string) |  |
| zipcode | [ string](#string) |  |
| city | [ string](#string) |  |
| street | [ string](#string) |  |
| notes | [ string](#string) |  |
| metadata | [repeated MetaData](#metadata) | Address meta data |

<br />



## Barcode
RequestUpdate Key->ID keys:

- ```id```: Barcode *code*

- ```barcodetype```: Valid values: *CODE_128, CODE_39, EAN_13, EAN_8, QR*

- ```product_id```: Product *partnumber*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| code | [ string](#string) | Each product can be connected to any number of bar codes, but the code must remain unique to ensure that the product is clearly identifiable. |
| product_id | [ int64](#int64) | Reference to [Product](#product).id |
| description | [ string](#string) | Comment related to the barcode. Informal, has no role in identification. |
| barcodetype | [ int64](#int64) | Reference to [Groups](#groups).id  (only where groupname = 'barcodetype'). |
| qty | [ double](#double) | The actual amount of the products identified by the barcode. For example can be used for packaged goods, tray packaging. |
| defcode | [ bool](#bool) | If more than one bar code is assigned, this will be the default. Because of the uniqueness of the barcode the product is always clearly identifiable, but in reverse case (eg. in case the barcode should be printed on a document) we must assign one being the default for that product. |

<br />



## Contact
RequestUpdate Key->ID keys:

- ```id```: The value is a generated constant key identifier: *{nervatype}/{refnumber}~{rownumber}*. The rownumber is the order of multiple *{nervatype}/{refnumber}* keys.

- ```nervatype```: Valid values: *customer, employee, event, place, product, project, tool, trans*

- ```ref_id```: Valid values: *customer/{custnumber}, employee/{empnumber}, event/{calnumber}, place/{planumber}, product/{partnumber}, project/{pronumber}, tool/{serial}, trans/{transnumber}*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| nervatype | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'nervatype') |
| ref_id | [ int64](#int64) | Reference to [Customer](#customer).id, [Employee](#employee).id, [Event](#event).id, [Place](#place).id, [Product](#product).id, [Project](#project).id, [Tool](#tool).id, [Trans](#trans).id |
| firstname | [ string](#string) |  |
| surname | [ string](#string) |  |
| status | [ string](#string) |  |
| phone | [ string](#string) |  |
| fax | [ string](#string) |  |
| mobil | [ string](#string) |  |
| email | [ string](#string) |  |
| notes | [ string](#string) |  |
| metadata | [repeated MetaData](#metadata) | Contact meta data |

<br />



## Currency
RequestUpdate Key->ID keys:

- ```id```: Currency *curr*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| curr | [ string](#string) | The ISO 4217 code of the currency. |
| description | [ string](#string) | The name of the currency. |
| digit | [ int64](#int64) | The number of decimal places used for recording and rounding by the program. Default: 2 |
| defrate | [ double](#double) | Default Rate. You can specify an exchange rate vs. the default currency, which will be used by the reports. |
| cround | [ int64](#int64) | Rounding value for cash. Could be used in case the smallest banknote in circulation for that certain currency is not 1. |
| metadata | [repeated MetaData](#metadata) | Currency meta data |

<br />



## Customer
RequestUpdate Key->ID keys:

- ```id```: Customer *custnumber*

- ```custtype```: Valid values: *own, company, private, other*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| custtype | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'custtype') |
| custnumber | [ string](#string) | Unique ID. If you set it to numberdef, it will be generated at the first data save. The format and value of the next data in row is taken from the numberdef (numberkey = custnumber) data series. |
| custname | [ string](#string) | Full name of the customer |
| taxnumber | [ string](#string) |  |
| account | [ string](#string) |  |
| notax | [ bool](#bool) | Tax-free |
| terms | [ int64](#int64) | Payment per. |
| creditlimit | [ double](#double) | Customer's credit limit. Data is used by financial reports. |
| discount | [ double](#double) | If new product line is added (offer, order, invoice etc.) all products will receive the discount percentage specified in this field. If the product has a separate customer price, the value specified here will not be considered by the program. |
| notes | [ string](#string) |  |
| inactive | [ bool](#bool) |  |
| metadata | [repeated MetaData](#metadata) | Customer meta data |

<br />



## Deffield
RequestUpdate Key->ID keys:

- ```id```: Deffield *fieldname*

- ```nervatype```: Valid values: *address, barcode, contact, currency, customer, employee, event, item, link, log, movement, payment, price, place, product, project, rate, tax, tool, trans, setting*

- ```subtype```: All groupvalue from Groups, where groupname equal *custtype, placetype, protype, toolgroup, transtype*

- ```fieldtype```: Valid values: *bool, date, time, float, integer, string, valuelist, notes, urlink, password, customer, tool, transitem, transmovement, transpayment, product, project, employee, place*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| fieldname | [ string](#string) |  |
| nervatype | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'nervatype') |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _subtype.subtype | [optional int64](#int64) | Reference to [Groups](#groups).id (where groupname in ('custtype','placetype','  protype','toolgroup','transtype')) |
| fieldtype | [ int64](#int64) | Reference to [Groups](#groups).id  (only where groupname = 'fieldtype') |
| description | [ string](#string) |  |
| valuelist | [ string](#string) | If fieldtype=valuelist: valid values are listed, separated by ~ |
| addnew | [ bool](#bool) | When selected, the attribute in case of adding a new element (eg a new customer or employee is created) will automatically be created with the default value according to its type and also will be attached to the new element. |
| visible | [ bool](#bool) | Can appear or not (hidden value) on the entry forms |
| readonly | [ bool](#bool) | The value of the attribute can not be changed in the program interface |

<br />



## Employee
RequestUpdate Key->ID keys:

- ```id```: Employee *empnumber*

- ```usergroup```: All groupvalue from Groups, where groupname equal usergroup

- ```department```: All groupvalue from Groups, where groupname equal department


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| empnumber | [ string](#string) | Unique ID. If you set it to numberdef, it will be generated at the first data save. The format and value of the next data in row is taken from the numberdef (numberkey = empnumber) data series. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _username.username | [optional string](#string) | Database login name. Should be unique on database level. |
| usergroup | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'usergroup') |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _startdate.startdate | [optional string](#string) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _enddate.enddate | [optional string](#string) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _department.department | [optional int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'department') |
| registration_key | [ string](#string) |  |
| inactive | [ bool](#bool) |  |
| metadata | [repeated MetaData](#metadata) | Employee meta data |

<br />



## Event
RequestUpdate Key->ID keys:

- ```id```: Event *calnumber*

- ```nervatype```: Valid values: *customer, employee, place, product, project, tool, trans*

- ```ref_id```: Valid values: *customer/{custnumber}, employee/{empnumber}, place/{planumber}, product/{partnumber}, project/{pronumber}, tool/{serial}, trans/{transnumber}*

- ```eventgroup```: All groupvalue from Groups, where groupname equal eventgroup


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| calnumber | [ string](#string) | Unique ID. If you set it to numberdef, it will be generated at the first data save. The format and value of the next data in row is taken from the numberdef (numberkey = calnumber) data series. |
| nervatype | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'nervatype') |
| ref_id | [ int64](#int64) | Reference to [Customer](#customer).id, [Employee](#employee).id, [Place](#place).id, [Product](#product).id, [Project](#project).id, [Tool](#tool).id, [Trans](#trans).id |
| uid | [ string](#string) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _eventgroup.eventgroup | [optional int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'eventgroup') |
| fromdate | [ string](#string) | Datetime |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _todate.todate | [optional string](#string) | Datetime |
| subject | [ string](#string) |  |
| place | [ string](#string) |  |
| description | [ string](#string) |  |
| metadata | [repeated MetaData](#metadata) | Event meta data |

<br />



## Fieldvalue
RequestUpdate Key->ID keys:

- ```id```: The value is a generated constant key identifier: *{refnumber}~~{fieldname}~{rownumber}*. The rownumber is the order of multiple *{refnumber}~~{fieldname}* keys.

- ```ref_id```: Valid values: *{nervatype}/{refnumber}*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| fieldname | [ string](#string) | Reference to [Deffield](#deffield).fieldname. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _ref_id.ref_id | [optional int64](#int64) | Reference to any type.id where type = [Deffield](#deffield).nervatype. If it is null then nervatype = setting. |
| value | [ string](#string) |  |
| notes | [ string](#string) |  |

<br />



## Groups
RequestUpdate Key->ID keys:

- ```id```: Group *groupname~groupvalue*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| groupname | [ string](#string) |  |
| groupvalue | [ string](#string) |  |
| description | [ string](#string) |  |
| inactive | [ bool](#bool) |  |

<br />



## Item
RequestUpdate Key->ID keys:

- ```id```: Trans and Item *transnumber~rownumber*

- ```trans_id```: Trans *transnumber*

- ```product_id```: Product *partnumber*

- ```tax_id```: Tax *taxcode*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| trans_id | [ int64](#int64) | Reference to [trans](#trans).id |
| product_id | [ int64](#int64) | Reference to [product](#product).id |
| unit | [ string](#string) |  |
| qty | [ double](#double) |  |
| fxprice | [ double](#double) |  |
| netamount | [ double](#double) |  |
| discount | [ double](#double) |  |
| tax_id | [ int64](#int64) | Reference to [Tax](#tax).id |
| vatamount | [ double](#double) |  |
| amount | [ double](#double) |  |
| description | [ string](#string) |  |
| deposit | [ bool](#bool) |  |
| ownstock | [ double](#double) |  |
| actionprice | [ bool](#bool) |  |
| metadata | [repeated MetaData](#metadata) | Item meta data |

<br />



## Link
RequestUpdate Key->ID keys:

- ```id```: *{nervatype_1}~{refnumber_1}~~{nervatype_2}~{refnumber_2}*

- ```nervatype_1```: All groupvalue from Groups, where groupname equal nervatype

- ```ref_id_1```: *{nervatype_1}/{refnumber_1}*

- ```nervatype_2```: All groupvalue from Groups, where groupname equal nervatype


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| nervatype_1 | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'nervatype') |
| ref_id_1 | [ int64](#int64) | Reference to {nervatype}.id |
| nervatype_2 | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'nervatype') |
| ref_id_2 | [ int64](#int64) | Reference to {nervatype}.id |
| metadata | [repeated MetaData](#metadata) | Link meta data |

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
| id | [ int64](#int64) |  |
| employee_id | [ int64](#int64) | Reference to [Employee](#employee).id |
| crdate | [ string](#string) | Date-time |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _nervatype.nervatype | [optional int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'nervatype') |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _ref_id.ref_id | [optional int64](#int64) | Reference to {nervatype}.id |
| logstate | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'logstate') |
| metadata | [repeated MetaData](#metadata) | Log meta data |

<br />



## MetaData



| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| fieldname | [ string](#string) | Reference to [Deffield](#deffield).fieldname. |
| fieldtype | [ string](#string) | Reference to [Deffield](#deffield).fieldtype. |
| value | [ string](#string) |  |
| notes | [ string](#string) |  |

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
| id | [ int64](#int64) |  |
| trans_id | [ int64](#int64) | Reference to [Trans](#trans).id |
| shippingdate | [ string](#string) | Date-time |
| movetype | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'movetype') |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _product_id.product_id | [optional int64](#int64) | Reference to [Product](#product).id |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _tool_id.tool_id | [optional int64](#int64) | Reference to [Tool](#tool).id |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _place_id.place_id | [optional int64](#int64) | Reference to [Place](#place).id |
| qty | [ double](#double) |  |
| description | [ string](#string) |  |
| shared | [ bool](#bool) |  |
| metadata | [repeated MetaData](#metadata) | Movement meta data |

<br />



## Numberdef
RequestUpdate Key->ID keys:

- ```id```: Numberdef *numberkey*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| numberkey | [ string](#string) | Unique key |
| prefix | [ string](#string) | The text prefix of the identifier. It can be any length, but usage of special characters, spaces in the text is not recommended. |
| curvalue | [ int64](#int64) | The current status of the counter, the next sequence number will be one value higher than this one. It is possible to re-set the counter, but the uniqueness must be ensured in all cases! |
| isyear | [ bool](#bool) | If selected, the created identifier will contain the year. |
| sep | [ string](#string) | The separator character in the identifier. Default: "/" |
| len | [ int64](#int64) | The value field is arranged in such length to the right and filled with zeros. |
| description | [ string](#string) |  |
| visible | [ bool](#bool) |  |
| readonly | [ bool](#bool) |  |
| orderby | [ int64](#int64) |  |

<br />



## Pattern
RequestUpdate Key->ID keys:

- ```id```: Pattern *description*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| description | [ string](#string) |  |
| transtype | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'transtype') |
| notes | [ string](#string) |  |
| defpattern | [ bool](#bool) |  |

<br />



## Payment
RequestUpdate Key->ID keys:

- ```id```: Trans and Item *transnumber~rownumber*

- ```trans_id```: Trans *transnumber*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| trans_id | [ int64](#int64) | Reference to [Trans](#trans).id |
| paiddate | [ string](#string) |  |
| amount | [ double](#double) |  |
| notes | [ string](#string) |  |
| metadata | [repeated MetaData](#metadata) | Payment meta data |

<br />



## Place
RequestUpdate Key->ID keys:

- ```id```: Place *planumber*

- ```placetype```: Valid values: *bank, cash, warehouse, other*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| planumber | [ string](#string) | Unique ID. If you set it to numberdef, it will be generated at the first data save. The format and value of the next data in row is taken from the numberdef (numberkey = planumber) data series. |
| placetype | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'placetype') |
| description | [ string](#string) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _curr.curr | [optional string](#string) |  |
| defplace | [ bool](#bool) |  |
| notes | [ string](#string) |  |
| inactive | [ bool](#bool) |  |
| metadata | [repeated MetaData](#metadata) | Place meta data |

<br />



## Price
RequestUpdate Key->ID keys:

- ```id```: Price *partnumber~validfrom~curr~qty*

- ```product_id```: Product *partnumber*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| product_id | [ int64](#int64) | Reference to [Product](#product).id |
| validfrom | [ string](#string) | Start of validity, mandatory data. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _validto.validto | [optional string](#string) | End of validity, can be left empty. |
| curr | [ string](#string) |  |
| qty | [ double](#double) | Price ranges can also be specified, thus different price can be set for a smaller and bigger quantity of the same product. The quantity should be used as the lower threshold, ie. this should be the minimum quantity for the price set. |
| pricevalue | [ double](#double) | Price value |
| vendorprice | [ bool](#bool) | Supplier (if marked) or customer price. By default the customer price. |
| metadata | [repeated MetaData](#metadata) | Price meta data |

<br />



## Product
RequestUpdate Key->ID keys:

- ```id```: Product *partnumber*

- ```protype```: Valid values: *item, service*

- ```tax_id```: Tax *taxcode*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| partnumber | [ string](#string) | Unique ID. If you set it to numberdef, it will be generated at the first data save. The format and value of the next data in row is taken from the numberdef (numberkey = partnumber) data series. |
| protype | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'protype') |
| description | [ string](#string) | The full name of the product or short description. |
| unit | [ string](#string) |  |
| tax_id | [ int64](#int64) | Reference to [Tax](#tax).id |
| notes | [ string](#string) |  |
| webitem | [ bool](#bool) |  |
| inactive | [ bool](#bool) |  |
| metadata | [repeated MetaData](#metadata) | Product meta data |

<br />



## Project
RequestUpdate Key->ID keys:

- ```id```: Project *pronumber*

- ```customer_id```: Tax *custnumber*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| pronumber | [ string](#string) | Unique ID. If you set it to numberdef, it will be generated at the first data save. The format and value of the next data in row is taken from the numberdef (numberkey = pronumber) data series. |
| description | [ string](#string) | The name of the project. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _customer_id.customer_id | [optional int64](#int64) | Reference to [Customer](#customer).id |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _startdate.startdate | [optional string](#string) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _enddate.enddate | [optional string](#string) |  |
| notes | [ string](#string) |  |
| inactive | [ bool](#bool) |  |
| metadata | [repeated MetaData](#metadata) | Project meta data |

<br />



## Rate
RequestUpdate Key->ID keys:

- ```id```: Rate *ratetype~ratedate~curr~planumber*

- ```place_id```: Place *planumber*

- ```ratetype```: Valid values: *rate, buy, sell,average*

- ```rategroup```: all groupvalue from Groups, where groupname equal rategroup


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| ratetype | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'ratetype') |
| ratedate | [ string](#string) |  |
| curr | [ string](#string) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _place_id.place_id | [optional int64](#int64) | Reference to [Place](#place).id |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _rategroup.rategroup | [optional int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'rategroup') |
| ratevalue | [ double](#double) | Rate or interest value |
| metadata | [repeated MetaData](#metadata) | Rate meta data |

<br />



## RequestDatabaseCreate
New database props.


| Field | Type | Description |
| ----- | ---- | ----------- |
| alias | [ string](#string) | Alias name of the database |
| demo | [ bool](#bool) | Create a DEMO database |

<br />



## RequestDelete
Delete parameters


| Field | Type | Description |
| ----- | ---- | ----------- |
| nervatype | [ DataType](#datatype) |  |
| id | [ int64](#int64) | The object ID |
| key | [ string](#string) | Use Key instead of ID |

<br />



## RequestEmpty
No parameters




## RequestFunction



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | Server function name |
| values | [map RequestFunction.ValuesEntry](#requestfunctionvaluesentry) | The array of parameter values |

<br />



## RequestFunction.ValuesEntry



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) |  |
| value | [ Value](#value) |  |

<br />



## RequestGet



| Field | Type | Description |
| ----- | ---- | ----------- |
| nervatype | [ DataType](#datatype) |  |
| metadata | [ bool](#bool) |  |
| ids | [repeated int64](#int64) |  |
| filter | [repeated string](#string) |  |

<br />



## RequestReport



| Field | Type | Description |
| ----- | ---- | ----------- |
| reportkey | [ string](#string) | Example : ntr_invoice_en |
| orientation | [ ReportOrientation](#reportorientation) |  |
| size | [ ReportSize](#reportsize) |  |
| output | [ ReportOutput](#reportoutput) |  |
| type | [ ReportType](#reporttype) |  |
| refnumber | [ string](#string) | Example : DMINV/00001 |
| template | [ string](#string) | Custom report JSON template |
| filters | [map RequestReport.FiltersEntry](#requestreportfiltersentry) |  |

<br />



## RequestReport.FiltersEntry



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) |  |
| value | [ Value](#value) |  |

<br />



## RequestReportDelete



| Field | Type | Description |
| ----- | ---- | ----------- |
| reportkey | [ string](#string) | Example : ntr_invoice_en |

<br />



## RequestReportInstall
Admin user group membership required.


| Field | Type | Description |
| ----- | ---- | ----------- |
| reportkey | [ string](#string) | Example : ntr_invoice_en |

<br />



## RequestReportList



| Field | Type | Description |
| ----- | ---- | ----------- |
| label | [ string](#string) |  |

<br />



## RequestTokenDecode



| Field | Type | Description |
| ----- | ---- | ----------- |
| value | [ string](#string) | Access token code. |

<br />



## RequestUpdate



| Field | Type | Description |
| ----- | ---- | ----------- |
| nervatype | [ DataType](#datatype) |  |
| items | [repeated RequestUpdate.Item](#requestupdateitem) |  |

<br />



## RequestUpdate.Item



| Field | Type | Description |
| ----- | ---- | ----------- |
| values | [map RequestUpdate.Item.ValuesEntry](#requestupdateitemvaluesentry) |  |
| keys | [map RequestUpdate.Item.KeysEntry](#requestupdateitemkeysentry) |  |

<br />



## RequestUpdate.Item.KeysEntry



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) |  |
| value | [ Value](#value) |  |

<br />



## RequestUpdate.Item.ValuesEntry



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) |  |
| value | [ Value](#value) |  |

<br />



## RequestUserLogin



| Field | Type | Description |
| ----- | ---- | ----------- |
| username | [ string](#string) | Employee username or Customer custnumber (email or phone number) |
| password | [ string](#string) |  |
| database | [ string](#string) | Optional. Default value: NT_DEFAULT_ALIAS |

<br />



## RequestUserPassword



| Field | Type | Description |
| ----- | ---- | ----------- |
| password | [ string](#string) | New password |
| confirm | [ string](#string) | New password confirmation |
| username | [ string](#string) | Optional. Only if different from the logged in user. Admin user group membership required. |
| custnumber | [ string](#string) | Optional. Only if different from the logged in user. Admin user group membership required. |

<br />



## RequestView
Only "select" queries and functions can be executed. Changes to the data are not saved in the database.


| Field | Type | Description |
| ----- | ---- | ----------- |
| options | [repeated RequestView.Query](#requestviewquery) | The array of Query object |

<br />



## RequestView.Query



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | Give the query a unique name |
| text | [ string](#string) | The SQL query as a string |
| values | [repeated Value](#value) | The array of parameter values |

<br />



## ResponseDatabaseCreate
Result log data


| Field | Type | Description |
| ----- | ---- | ----------- |
| details | [ ResponseRows](#responserows) |  |

<br />



## ResponseEmpty
Does not return content.




## ResponseFunction



| Field | Type | Description |
| ----- | ---- | ----------- |
| value | [ bytes](#bytes) |  |

<br />



## ResponseGet



| Field | Type | Description |
| ----- | ---- | ----------- |
| values | [repeated ResponseGet.Value](#responsegetvalue) |  |

<br />



## ResponseGet.Value



| Field | Type | Description |
| ----- | ---- | ----------- |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.address | [ Address](#address) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.barcode | [ Barcode](#barcode) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.contact | [ Contact](#contact) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.currency | [ Currency](#currency) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.customer | [ Customer](#customer) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.deffield | [ Deffield](#deffield) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.employee | [ Employee](#employee) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.event | [ Event](#event) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.fieldvalue | [ Fieldvalue](#fieldvalue) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.groups | [ Groups](#groups) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.item | [ Item](#item) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.link | [ Link](#link) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.log | [ Log](#log) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.movement | [ Movement](#movement) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.numberdef | [ Numberdef](#numberdef) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.pattern | [ Pattern](#pattern) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.payment | [ Payment](#payment) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.place | [ Place](#place) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.price | [ Price](#price) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.product | [ Product](#product) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.project | [ Project](#project) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.rate | [ Rate](#rate) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.tax | [ Tax](#tax) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.tool | [ Tool](#tool) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.trans | [ Trans](#trans) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.ui_audit | [ UiAudit](#uiaudit) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.ui_menu | [ UiMenu](#uimenu) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.ui_menufields | [ UiMenufields](#uimenufields) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.ui_message | [ UiMessage](#uimessage) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.ui_printqueue | [ UiPrintqueue](#uiprintqueue) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.ui_report | [ UiReport](#uireport) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.ui_userconfig | [ UiUserconfig](#uiuserconfig) |  |

<br />



## ResponseReport



| Field | Type | Description |
| ----- | ---- | ----------- |
| value | [ bytes](#bytes) |  |

<br />



## ResponseReportInstall



| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) | Returns a new report ID. |

<br />



## ResponseReportList
Returns all installable files from the NT_REPORT_DIR directory (empty value: all available built-in Nervatura Reports)


| Field | Type | Description |
| ----- | ---- | ----------- |
| items | [repeated ResponseReportList.Info](#responsereportlistinfo) |  |

<br />



## ResponseReportList.Info



| Field | Type | Description |
| ----- | ---- | ----------- |
| reportkey | [ string](#string) |  |
| repname | [ string](#string) |  |
| description | [ string](#string) |  |
| label | [ string](#string) |  |
| reptype | [ string](#string) |  |
| filename | [ string](#string) |  |
| installed | [ bool](#bool) |  |

<br />



## ResponseRows



| Field | Type | Description |
| ----- | ---- | ----------- |
| items | [repeated ResponseRows.Item](#responserowsitem) |  |

<br />



## ResponseRows.Item



| Field | Type | Description |
| ----- | ---- | ----------- |
| values | [map ResponseRows.Item.ValuesEntry](#responserowsitemvaluesentry) |  |

<br />



## ResponseRows.Item.ValuesEntry



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) |  |
| value | [ Value](#value) |  |

<br />



## ResponseTokenDecode
Access token claims.


| Field | Type | Description |
| ----- | ---- | ----------- |
| username | [ string](#string) |  |
| database | [ string](#string) |  |
| exp | [ double](#double) | JWT expiration time |
| iss | [ string](#string) |  |

<br />



## ResponseTokenLogin
Token user properties


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| username | [ string](#string) |  |
| empnumber | [ string](#string) |  |
| usergroup | [ int64](#int64) |  |
| scope | [ string](#string) |  |
| department | [ string](#string) |  |

<br />



## ResponseTokenRefresh



| Field | Type | Description |
| ----- | ---- | ----------- |
| value | [ string](#string) | Access token code. |

<br />



## ResponseUpdate
If the ID (or Key) value is missing, it creates a new item.


| Field | Type | Description |
| ----- | ---- | ----------- |
| values | [repeated int64](#int64) | Returns the all new/updated IDs values. |

<br />



## ResponseUserLogin



| Field | Type | Description |
| ----- | ---- | ----------- |
| token | [ string](#string) | Access JWT token |
| engine | [ string](#string) | Type of database |
| version | [ string](#string) | Service version |

<br />



## ResponseView



| Field | Type | Description |
| ----- | ---- | ----------- |
| values | [map ResponseView.ValuesEntry](#responseviewvaluesentry) | key - results map |

<br />



## ResponseView.ValuesEntry



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) |  |
| value | [ ResponseRows](#responserows) |  |

<br />



## Tax
RequestUpdate Key->ID keys:

- ```id```: Tax *taxcode*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| taxcode | [ string](#string) | Unique ID. |
| description | [ string](#string) |  |
| rate | [ double](#double) | Rate or interest value |
| inactive | [ bool](#bool) |  |
| metadata | [repeated MetaData](#metadata) | Tax meta data |

<br />



## Tool
RequestUpdate Key->ID keys:

- ```id```: Tool *serial*

- ```toolgroup```: all groupvalue from Groups, where groupname equal toolgroup

- ```product_id```: Product *partnumber*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| serial | [ string](#string) | Unique ID. If you set it to numberdef, it will be generated at the first data save. The format and value of the next data in row is taken from the numberdef (numberkey = serial) data series. |
| description | [ string](#string) |  |
| product_id | [ int64](#int64) | Reference to [Product](#product).id |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _toolgroup.toolgroup | [optional int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'toolgroup') |
| notes | [ string](#string) |  |
| inactive | [ bool](#bool) |  |
| metadata | [repeated MetaData](#metadata) | Tool meta data |

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
| id | [ int64](#int64) |  |
| transnumber | [ string](#string) | Unique ID. If you set it to numberdef, it will be generated at the first data save. The format and value of the next data in row is taken from the numberdef (numberkey = transnumber) data series. |
| transtype | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'transtype')[Groups](#groups).id |
| direction | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'direction') |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _ref_transnumber.ref_transnumber | [optional string](#string) |  |
| crdate | [ string](#string) |  |
| transdate | [ string](#string) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _duedate.duedate | [optional string](#string) | Date-time |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _customer_id.customer_id | [optional int64](#int64) | Reference to [Customer](#customer).id |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _employee_id.employee_id | [optional int64](#int64) | Reference to [Employee](#employee).id |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _department.department | [optional int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'department') |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _project_id.project_id | [optional int64](#int64) | Reference to [Project](#project).id |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _place_id.place_id | [optional int64](#int64) | Reference to [Place](#place).id |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _paidtype.paidtype | [optional int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'paidtype') |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _curr.curr | [optional string](#string) |  |
| notax | [ bool](#bool) |  |
| paid | [ bool](#bool) |  |
| acrate | [ double](#double) |  |
| notes | [ string](#string) |  |
| intnotes | [ string](#string) |  |
| fnote | [ string](#string) |  |
| transtate | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'transtate') |
| closed | [ bool](#bool) |  |
| metadata | [repeated MetaData](#metadata) | Trans meta data |

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
| id | [ int64](#int64) |  |
| usergroup | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'usergroup') |
| nervatype | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'nervatype') |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _subtype.subtype | [optional int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'movetype') |
| inputfilter | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'inputfilter') |
| supervisor | [ bool](#bool) |  |

<br />



## UiMenu
RequestUpdate Key->ID keys:

- ```id```: UiMenu *menukey*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| menukey | [ string](#string) |  |
| description | [ string](#string) |  |
| modul | [ string](#string) |  |
| icon | [ string](#string) |  |
| method | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'method') |
| funcname | [ string](#string) |  |
| address | [ string](#string) |  |

<br />



## UiMenufields
RequestUpdate Key->ID keys:

- ```id```: UiMenufields *{menukey}~{fieldname}*

- ```menu_id```: UiMenu *menukey*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| menu_id | [ int64](#int64) | Reference to [UiMenu](#UiMenu).id |
| fieldname | [ string](#string) |  |
| description | [ string](#string) |  |
| fieldtype | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'fieldtype') |
| orderby | [ int64](#int64) |  |

<br />



## UiMessage
RequestUpdate Key->ID keys:

- ```id```: UiMessage *{secname}~{fieldname}~{lang}*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| secname | [ string](#string) |  |
| fieldname | [ string](#string) |  |
| lang | [ string](#string) |  |
| msg | [ string](#string) |  |

<br />



## UiPrintqueue



| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _nervatype.nervatype | [optional int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'nervatype') |
| ref_id | [ int64](#int64) | Reference to {nervatype}.id |
| qty | [ double](#double) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _employee_id.employee_id | [optional int64](#int64) | Reference to [Employee](#employee).id |
| report_id | [ int64](#int64) | Reference to [UiReport](#UiReport).id |
| crdate | [ string](#string) | Date-time |

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
| id | [ int64](#int64) |  |
| reportkey | [ string](#string) |  |
| nervatype | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'nervatype') |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _transtype.transtype | [optional int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'transtype') |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _direction.direction | [optional int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'direction') |
| repname | [ string](#string) |  |
| description | [ string](#string) |  |
| label | [ string](#string) |  |
| filetype | [ int64](#int64) | Reference to [Groups](#groups).id (only where groupname = 'filetype') |
| report | [ string](#string) |  |

<br />



## UiUserconfig
RequestUpdate Key->ID keys:

- ```id```: *{empnumber}~{section}~{cfgroup}~{cfname}*

- ```employee_id```: Employee *{empnumber}*


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _employee_id.employee_id | [optional int64](#int64) | Reference to [Employee](#employee).id |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _section.section | [optional string](#string) |  |
| cfgroup | [ string](#string) |  |
| cfname | [ string](#string) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) _cfvalue.cfvalue | [optional string](#string) |  |
| orderby | [ int64](#int64) |  |

<br />



## Value



| Field | Type | Description |
| ----- | ---- | ----------- |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.boolean | [ bool](#bool) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.number | [ double](#double) |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) value.text | [ string](#string) | google.protobuf.NullValue null = 4; |

<br />



---
# Enums



## DataType


| Name | Number | Description |
| ---- | ------ | ----------- |
| address | 0 | [Address](#address) |
| barcode | 1 | [Barcode](#barcode) |
| contact | 2 | [Contact](#contact) |
| currency | 3 | [Currency](#currency) |
| customer | 4 | [Customer](#customer) |
| deffield | 5 | [Deffield](#deffield) |
| employee | 6 | [Employee](#employee) |
| event | 7 | [Event](#event) |
| fieldvalue | 8 | [Fieldvalue](#fieldvalue) |
| groups | 9 | [Groups](#groups) |
| item | 10 | [Item](#item) |
| link | 11 | [Link](#link) |
| log | 12 | [Log](#log) |
| movement | 13 | [Movement](#movement) |
| numberdef | 14 | [Numberdef](#numberdef) |
| pattern | 15 | [Pattern](#pattern) |
| payment | 16 | [Payment](#payment) |
| place | 17 | [Place](#place) |
| price | 18 | [Price](#price) |
| product | 19 | [Product](#product) |
| project | 20 | [Project](#project) |
| rate | 21 | [Rate](#rate) |
| tax | 22 | [Tax](#tax) |
| tool | 23 | [Tool](#tool) |
| trans | 24 | [Trans](#trans) |
| ui_audit | 25 | [UiAudit](#uiaudit) |
| ui_menu | 26 | [UiMenu](#uimenu) |
| ui_menufields | 27 | [UiMenufields](#uimenufields) |
| ui_message | 28 | [UiMessage](#uimessage) |
| ui_printqueue | 29 | [UiPrintqueue](#uiprintqueue) |
| ui_report | 30 | [UiReport](#uireport) |
| ui_userconfig | 31 | [UiUserconfig](#uiuserconfig) |

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

