---
title: Employee
type: docs
weight: 30
bookToC: true
---

## Input fields

### Employee No.
Unique ID, generated at the first data save. The format and value of the next data in row is taken from the [**DOCUMENT NUMBERING**](/docs/client/settings/numberdef) (code = empnumber) data series.

### Firtsname, Surname
The name of the employee.

### Status, Phone, Mobile, Email
Contact data.

### Start Date, End Date
Employment start and end date.

### Department
An item of the [**department**](/docs/client/settings/groups) group should be set in this field.

### Usergroup
One of the usergroups handled under [**ACCESS RIGHTS**](/docs/client/settings/usergroup). Mandatory.

### Username
Database login name. Should be unique on database level.

{{< hint info >}}

Usernames, passwords and access right settings specified here are also used by other programs of the Nervatura Framework, thus by the [**Nervatura HTTP API**](/docs/service/api)!

{{< /hint >}}

### Comment
Remarks field.

## Related data

### [**METADATA**](/docs/client/settings/metadata)
Unlimited number of supplementary data can be added.

### ADDRESS DATA
**Country, State, Zipcode, City, Street, Comment**<br />
Unlimited number of address data can be added.

### [**EVENT**](/docs/client/resources/event)
Unlimited number of events can be added.

## Operations

### CHANGE PASSWORD
Password change possibility for the user on the form.

{{< hint danger >}}

When the username for an employee is created for first time by default no password will be given by the program. **However we highly recommend to set up a temporary password!**<br />
Users can modify their own password at any time, regardless of their other assigned user rights (under Setting/USER/CHANGE PASSWORD).<br />
The user passwords are stored encrypted in the database. In case a password is forgotten it is only possible to change it, but the original one cannot be decrypted.

{{< /hint >}}

### REPORT
[**DATA EXPORT**](/docs/client/program/export)

### BOOKMARK
Set a bookmark for the record. Later can be loaded from bookmarks at any time.
