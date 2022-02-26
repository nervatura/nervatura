---
title: Object Model
type: docs
weight: 20
bookFlatSection: true
bookCollapseSection: false
bookToC: true
---

## Overview
It is a general **open-data model**, which can store all information generated in the operation of a usual corporation. This covers all manufacturer, retailer and service companies (or governmental units) where the business operation can be defined and described within a **GOODS** (items, services to be sold, provided) – **CLIENT** (the recipient of goods) - **RESOURCE** (assets used to produce the goods) triangle.

{{< hint info >}}

The open-data principle regulates the access to our business data. The point is the logic of the data storage. It means that the data are defined for storage so, that those are compliant with an open data-model which could be accessed and interpreted by anyone. It doesn't concern the physical storage of the data, that can be implemented according to one's needs. However it should ensure that data can be managed safely according to published description. Retrieving, new data creation, possibility to export the entire data structure should be provided.

*What are the main advantages of open-data applications?*
- **safety**: provided one's have proper usage rights and physical access to a database, then will be able to interpret and process the data correctly without any help or permission from a third party. Information in the data becomes independent of any management system, its treatment is not tied to specific programs or technologies.
- **efficiency and cost reduction**: the business management system can be developed and diversified in accordance with one's needs. There is no need and pressure to be tied to a solution of any vendor, the most appropriate tools and programs can be selected for all tasks. The only criteria is that selected applications should be able to communicate and exchange data with each other or with a central database according to the open-data description. The elements of the system can at any time be flexibly developed or new ones added by choosing the best offers available on the market.

{{< /hint >}}

It is located between the application surfaces that are using and creating the data and the real data storage layer. It defines logical objects; data is stored in these freely defined attributes and in relations between them. Its flexible structure allows defining new properties or assigning events to our objects.

The number of objects is minimal, their structure is simple. It has an easy to learn, clear and straightforward logic. However it is capable to store the required data in structures. It ensures the possibility to attach defined type metadata of any kind to each object and also makes the objects linkable to each other arbitrarily.

The data model is independent from data storage layer. The data storage can be implemented in any way or with any device but as a main requirement the user of the data model must not sense this at all.

## Objects

Such ***pre-defined functional roles*** which can have any type of attributs, events can be attached to them as well as their elements can be attached to elements of other objects.

<img alt="Nervatura Object Model" src="/images/nom.svg"
  style="width:325px;display:block;margin-left:auto;margin-right:auto;" />

## Base objects

**CUSTOMER** - all external partners of the company, including the buyer, consumer and supplier side

**PRODUCT** - all raw materials, semi-finished and end-products that are related to our activity (as customer or vendor), produced by us as a manufacturer or offered as service

**TOOL, EMPLOYEE, PLACE** resources which are available for executing the activity and they contribute to it. These can be human resources (EMPLOYEE), material devices, tools, machines (TOOL) or financial, potentially infrastructural conditions such as warehouses, bank account, petty cash (PLACE)

## Metadata

All data that **describe a given object**, we want to attach to it as information. Some of them are pre-defined but further ones can freely be defined for any of the objects.

By using the **DEFFIELD** object we can define data storage metadata for other objects. Besides the classical data types (bool, integer, float, date, time, string, notes) these can contain list of values (valuelist), url links (urlink) or references to concrete items of other objects (customer, product, tool, employee, etc.).

Through the **FIELDVALUE** object every defined feature of the elements of all objects can be queried.

## Events

**Extended object metadata, usually connected to a time or an interval.** With the help of events we can make the static metadata of an object into dynamic so the feature of a given component is able to take various values at different times. An event can also be valid for a period of time, so having a start and an end date. Optional number of supplementary data of a given object can be attached to it, and it can be grouped as well.

We can manage the events through the **EVENT** object. Beside the base object we can also assign events to projects.

## Transaction

Transaction **is such a sort of event to which at least two base objects are joined**.

An event is always attached to a given object. As a further event feature another base object can be specified but it's just an optional additional data in this case.<br />
In the transactions the relation between the base objects is an indispensable and essential component of the given event. The transaction doesn’t belong to any of the base objects but the base objects are joined to a transaction. From these some base objects might be optional components but at least two should be indispensable part of it.

The most common object pair is the customer and product relationship (e.g.: offer, order, invoice) but any other combination is also possible, for example product-place (stock management), customer-tool (rental), employee-customer (worksheet) etc.

We can link additional data to transactions just as to events, but in contrary to events, here we don’t use the features of the linked base object but we can declare own metadata. Transactions can also be linked to each other or can "originate" from each other, for instance offer -> order -> inventory move -> invoice.

The object of transactions is the **TRANS** which contains the main data of transactions as well as the single object relations. **ITEM** object contains PRODUCT lines linked to transactions, **PAYMENT** object contains financial settlements, **MOVEMENT** object contains warehouse and tool movements.

## Relations

There are several possibilities to link single objects. Usually the object has the possibility of applying the one to one relation by default, if it is required so by its type. In case of need the additional data pointing to the proper type of object can also be generated at any time.<br />
For example we can set a customer type feature for CUSTOMER object wherewith we can link a given customer to another customer. With the same method one to many relations can also be set, so in this case we can also link our customer to some other customers.

If any linked customer is also linked to an other customer it results in a many to many relation.<br />
**LINK** object can also be used to set relations to objects. This way two objects can be linked without setting further object features.

## Group settings

Several options are available for grouping the objects. Using supplementary data, further to data storage opportunities allows also grouping to a certain degree.

In the **GROUPS** object we can create groups by object types. If needed, further features can be defined for these groups. These can then be used for assignments of pre-defined values on a given object (for example type options), but through LINK object can also be used for creating classical one to many groups (for example customer or product groups).

Actually the **PROJECT** object can be interpreted as the extension of GROUPS object. Surely it is possible to set metadata here as well but at PROJECT object time related extension is also possible, just like it is in case of events vs. metadata. Optionally it can have start or end date, we can also link it to customers or places. Projects can also have their own events as well as any transaction can be linked to them.

## Complex data types

When adding features to objects in some cases complex data feature setting is needed. Essentially **these are such sub objects which possess own features**. For example if we want to add address data to a customer then by setting the address we can give the city, the zip code or the name of the street as well.<br />
Some complex data types can be linked not only to a single object but the same element can also be attached to many others. In some of them it is possible to define further metadata.

One to many linked sub objects:<br />
**ADDRESS, BARCODE, CONTACT, PRICE**

One to one linked sub objects:<br />
**CURRENCY, PATTERN, RATE, TAX**

## Other objects

The objects of rights management, logging and other options:<br />
**LOG, NUMBERDEF**

## User interface objects

These objects **are not part of the object model**, they are not needed for recording the workflow data. However certain applications to ensure their own operation might require data storage possibilities.

Storage of data of Reports:<br />
**UI_REPORT**

Settings of user interfaces:<br />
**UI_MENU, UI_MENUFIELDS**

User rights, settings:<br />
**UI_AUDIT, UI_USERCONFIG**

Regional settings:<br />
**UI_MESSAGE**

Printing:<br />
**UI_PRINTQUEUE**

## Relations pyramid

*For safe data export and import go from top to the bottom.*

| Level | | Metadata	| Objects |
|---|---|---|---|
| level 1a	| no external link	| no	| GROUPS, NUMBERDEF |
| level 1b	| no external link	| yes* |	CURRENCY, TAX |
| level 2a	| link to <= level 1	| no |	DEFFIELD, PATTERN |
| level 2b	| link to <= level 1	| yes* |	CUSTOMER, EMPLOYEE, PLACE, PRODUCT |
| level 3	| link to <= level 2	| yes* |	BARCODE, PRICE, PROJECT, RATE, TOOL |
| level 4	| link to <= level 3	| yes* |	TRANS |
| level 5	| link to <= level 4	| yes* |	EVENT, ITEM, MOVEMENT, PAYMENT |
| level 6	| link to <= level 5	| yes* |	ADDRESS, CONTACT |
| level 7	| link to <= level 6	| yes* |	LINK, LOG |
| level 8	| link to <= level 7	| no |	FIELDVALUE |
**Export with the FIELDVALUE (cross-references fields)*

| Level | |	Objects |
|---|---|---|
| level 1 |	no external link |	UI_MENU |
| level 2a |	link to <= level 1 |	UI_MESSAGE |
| level 2b |	link to <= NOM level 1 |	UI_REPORT, UI_AUDIT |
| level 3 |	link to <= level 2b |	UI_MENUFIELDS |
| level 4 |	link to <= NOM level 2 |	UI_USERCONFIG, UI_PRINTQUEUE |

## Objects relations

*1. picture: Document type (transtype) relations.*

<img alt="Document type relations" style="max-width:100%;" src="/images/trans.svg" />

*2. picture: A possible relational database plan of NOM objects.*

<img alt="Database plan of NOM objects" style="max-width:100%" src="/images/nom_rel.svg" />

*3. picture: A possible relational database plan of user interface objects.*

<img alt="Database plan of user interface objects" style="max-width:100%" src="/images/nom_uio.svg" />


