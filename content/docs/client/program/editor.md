---
title: Report Editor
type: docs
weight: 70
bookToC: true
---

## Overview

The REPORT EDITOR is used to modify some of the Nervatura Report definitions and descriptions of **REPORT TEMPLATE**. It defines the layout of the reports: where and how the data should appear. It will not support changes regarding what data should be displayed from the database, but will give the possibility to modify the report with additional information (eg. logo files, texts, language translations). Therefore this tool is not suitable for creating completely new reports, however the existing reports can freely be modified and by using copies additional versions can be created.

It's designed to help to easily customize the samples in the database with a visual tool. You can set your own logo, rename, change or skip certain data. For the same thing (eg. invoice, delivery note) you can create multiple templates, in different languages ​​or with different data content.

{{< hint warning >}}

The **Nervatura Report** enables you to create from your data formatted files with standardized outcome. Currently two formats are supported.

With **PDF** format ready-to-print documents can be created. These are locked, not modifiable files, can be sent forward as email attachments. Their usage is widely spread, supported by most of the devices. Useful for making orders, invoices, delivery notes, etc.

The **XML** format is useful especially in data exchange with other applications, databases. The format is well known, the content can be easily processed.

{{< /hint >}}

{{< hint warning >}}

The list does not include the general reports of the [**REPORT**](/docs/client/program/report) menu. Preparation and modification of those requires specific database skills.

{{< /hint >}}

## Operations

### PRINT PDF
The print/preview of the template with sample data. The program asks for the identifier of concerned template according to its type (eg CUSTOMER DATASHEET template: DMCUST/00001 in the demo database) and fills up the template with the data. For following run offers the identifier that was used last time.

### SAVE TEMPLATE
Save all changes made to the template in the database.

### CREATE FROM
Create a new template based on the existing. The currently edited template will remain unchanged, by copying it a new one is created. The template ID is given by the program, the name can be set by the user. This name will appear in the selection list.

{{< hint info >}}

To delete a template use the delete button at the end of the item lines in the template selection list!

{{< /hint >}}

{{< hint warning >}}

The following items are not necessary for an average user. Would however be useful for developers and advanced users to create new Nervatura Reports.

{{< /hint >}}

### NEW BLANK
A new, blank template creation. The template can not be saved to the database, only export is possible.

### NEW SAMPLE
Creating a new template with sample data. The template can not be saved to the database, only export is possible.

### JSON EXPORT
Export the template to JSON file.

## TEMPLATE

The template structure is shown in hierarchical way (tree) as parent-child relation. When an element of the list is selected then only the highest level (parent) and possibly the subordinates are visible in the list. On the other (editable) page the elements’ potential attributes and their current value is shown. Under the name of the attributes a short help can be read about their possible values.

The template consists of the following four main components:

### REPORT
Includes general information about the whole template: subject, data about the author, default font and colors and margin settings. Children level elements are not included.

### HEADER
Elements placed in this section will appear at the top (header) of each page. The following types can be included: ROW, VGAP, HLINE.

### DETAILS
It constantly fills up the available space with data. If necessary, automatically opens a new page. The header and footer areas are skipped. The following types can be included: ROW, VGAP, HLINE, HTML, DATAGRID.

### FOOTER
Elements placed in this section will appear in the lower part (footer) of each page. The following types can be included: ROW, VGAP, HLINE.

---

Elements to be used:

### ROW
Horizontal logical group. The last element width extends up to the right margin.
- **CELL**: Displays data with format setting options: font, color, background, alignment, frame, etc. The value may be provided directly and also in databind (see DATA) form.
- **IMAGE**: jpg or png image, coded in base64 form. The value may be provided directly and also in databind (see DATA) form. It is possible to select and load an image file from a specific location. In this case the coding to base64 text format is done automatically by the program.
- **BARCODE**: Interleaved 2of5, Code 39, Code 128, EAN 8, EAN 13, QR
- **SEPARATOR**: Vertical separator line.

### VGAP
Vertical gap.

### HLINE
Horizontal line.

### HTML
The text specified here can include Basic HTML format elements (bold, italic, etc.). Experimental!

### DATAGRID
Displays data in table format. The value cannot be provided directly, just in databind (see DATA) form.
- **COLUMN**: displays settings of a table column (field)

---

### Operations
It shows the relative location of the selected element on the page. With the commands here you can change its location, delete it or insert a new element.

- **PREVIOUS, NEXT** <br />
Move the selection to previous or next element.

- **MOVE UP, MOVE DOWN** <br />
Move the selected element at its own level. It allows to modify the order of the child elements. If the element itself also contains child elements then those will be moved along.

- **DELETE, ADD ELEMENT** <br />
To delete the selected element or insert a child element under it.

## DATA

The elements to be used can be classified into two types:
- graphic element or grouping: ROW, SEPARATOR, VGAP, HLINE
- element to display data: CELL, IMAGE, BARCODE, HTML, DATAGRID(COLUMN)

From the latter group for CELL, IMAGE, BARCODE, HTML elements the data can be defined in two ways. Alternatively, you can simply specify a value for the corresponding attribute of the element (value, src, html) in the template. In this case the report will show exactly this value in the format corresponding to the element.

Most of the data presented in a report however are not like this. Usually we do not know exactly what will be the value of an element, we know only what kind of information we want to be displayed. In these cases you can use the *databind* variables.

{{< hint warning >}}

For example, a customer name will depend on what client the report is prepared for. So, not an exact customer name should be entered in the field, but a code: *head.0.custname*. This is a variable, which will automatically be replaced by the program with a customer name when the data is displayed.

{{< /hint >}}

The program will always check if the given value corresponds to any data. If it does then the data will be used, otherwise the entered value itself.

The data may come from two different sources. One of them is representing the data which are provided automatically by the program. These were defined by using the Nervatura Report, but fall outside the Report Editor’s scope. Other source is the possibility to specify data in the Report Template section, these are then included in its DATA part.

{{< hint info >}}

The easiest way to find out what variables are available for a certain type of template (besides those specified in the DATA section) is to check the original sample templates.

The name of the variables everywhere follows the same logic. Some examples:
- The base data related to a template type, such as a customer's name, customer’s number, an invoice duedate or an employee’s user name all will have an identifier starting with **head**: head.0.custname, head.0.custnumber, head.0.duedate, head.0.username.
- The item lines of orders, invoices, delivery notes begin with name **items**. The first row of the net amount of an invoice will be items.0.netamount, the third row of the quantities on an order will be items.2.qty.
- All contact details of a customer will start with **contact**, the city of the second address will be **address**.1.city.
- The **labels** are special variables, can be found in each template and include the texts of the captions: labels.lb_customer_no, labels.lb_due_date, labels.lb_comment.

{{< /hint >}}

The following types of variables can be defined:

### TEXT
A simple variable name-text binding, where the variable name is entered the connected text will be displayed. The text can be a single word or longer section. For example, the pictures of the IMAGE element are usually also coded with this type in a text format.

The variable name must be unique within the template, with the same variable name two different data types mustn’t be defined!

### LIST
If we group several TEXT variable pairs together and name the group, then we’ll get the LIST type. To an element of the LIST we can refer with list_name.value_name. Typical example of the LIST type data is the general group 
of *labels*.

The variable names in this case must be unique within the LIST, and the LIST name 
should not be given to anything else within the template.

### TABLE
If the LIST types are organized in an orderly sequenced new list and are given a name, we get the TABLE type. Here each row is made up of lists with same column name. An item can be referred to as table_name.list_index.value_name and one of its lines as table_name.list_index.

Most of the preset variables are defined this way. You can create this type also using the Report Editor, in this case the column names should be separated by commas when they are created and they can not be changed later.

### Operations
- **NEW DATA, DELETE**: To create or delete a new databind

{{< hint info >}}

The modified templates can be saved to the database, but we recommend to make copies of original templates and save them under a new name. If you do not want your users to use the original templates, you can restrict their user groups’ access to these in [**ACCESS RIGHTS**](/docs/client/settings/usergroup) section.

{{< /hint >}}

{{< hint info >}}

If you would like to change the captions in a template, please do not modify the **labels** list field names! The program will always display this list with originally specified values.

Instead, create a new list called **labels_de** for example. Add to the list a new field with city name and then enter text Stadt. Replace the template’s **labels.city** value by **labels_de.city**. In the reports of such modified templates instead of City labels the Stadt name will be displayed.

{{< /hint >}}

{{< hint info >}}

The data of logo images in the original templates can be found in the DATA section. The PNG and JPG type images must be specified in base64 text encoding. If you do not know how to encode these pictures in such formats, you can also do the following:<br />
Find the IMAGE element in the template. With the loading button located here please specify the destination of the image file. After loading, the src field will contain the encoded version of the image. You can choose to leave this value here or can also copy the content to the logo (or differently named) variable of the DATA section. In this case, the src field of the IMAGE should be replaced with logo value.

If the template uses the image file only at one place, then there is no difference between the two solutions. If at several places, then the latter solution is more effective because the image data should be stored only once.

{{< /hint >}}

{{< hint info >}}

The default templates are optimized to a smaller 25-30px size logo image file. Bigger images cannot be placed there with a simple replace. In these cases you might need to restructure the other elements of the template as well.

If you would not like to use a logo in the template, it can be deleted according to below instructions.
- select the template in Setting->REPORT EDITOR
- choose the options in following order: Edit->LIST->HEADER->ROW->IMAGE
- select the MAP group and then press the DELETE button
- choose the TEMPLATE group and with SAVE TEMPLATE button save the changes

If you would like to keep also the original version, create a copy by using the CREATE FROM button in TEMPLATE group before making any changes!

{{< /hint >}}
