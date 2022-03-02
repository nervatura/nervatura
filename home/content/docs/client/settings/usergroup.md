---
title: Access rights
type: docs
weight: 30
bookToC: true
---

## Overview

The user access rights in Nervatura are managed through user groups. The access rules are assigned to these groups and are applicable to group members. When the rule is changed, the new settings will automatically be valid for all members.<br />
Each user must be a member of an access rights group, but can only be part of one of these groups.  This can be set in [**EMPLOYEE**](/docs/client/resources/employee) usergroup field.

{{< hint danger >}}

**The rights set here will apply also to other user interfaces of Nervatura**, such as the functions of [**Nervatura HTTP API**](/docs/service/api)!

{{< /hint >}}

## Input fields

### Group
Group ID. A unique value, can not be repeated.

An user access rights group by default is created with full access, though the scope can be limited by rules. When a new rule is created the following can be set:
- **Type**: audit, customer, employee, event, menu, price, product, project, report, setting, tool, trans
- **Subtype**: just in case of *report*, *menu* or *trans* types. New rule appears after first save.
  - ***report***: codes of the [**reports**](/docs/client/program/report) in the database. If set to *disabled* the report is not available for the group members.
  - ***menu***: codes of the [**menu shortcuts**](/docs/client/settings/uimenu) in the database. If set to <i>disabled</i> the menu shortcut is not available for the group members.
  - ***trans***: bank, cash, delivery, formula, inventory, invoice, offer, order, production, rent, waybill, worksheet. The settings apply only to the specified subtype.
- **Filter**
  - ***all***: there are no restrictions
  - ***disabled***: cannot be selected or will not be even displayed
  - ***readonly***: can be opened only as read-only
  - ***update***: values ​​can be changed, but new value creation is not allowed, also members of the group can not delete them
- **Supervisor**: related to multi-level access verification, not used in the current version. When set, it regulates the access to forms *State* field.

### Data filter
Regulates the access to data in DOCUMENT, PAYMENT, STOCK CONTROL menus, as follows:
- ***all***: full access
- ***own***: the user can only see the data created by himself
- ***usergroup***: the user can only see the data created by the members of the user group he belongs to
