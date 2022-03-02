---
title: Data browser
type: docs
weight: 20
bookToC: true
---

## Operations

### VIEWS
Switching between data views.

### COLUMNS
Show or hide columns.

### FILTER
Add filters to the query:
- Unlimited number of conditions can be added, which are related to each other with 
an "and" restrictive condition
- Condition set for a certain data view is kept till the program is exited.
- For filtering all columns in the data view can be selected, regardless of whether they are currently displayed or not
- All data types can be filtered with EQUAL, IS NULL, NOT EQUAL relations, as well as the numbers and dates with EQUAL TO OR GREATER and EQUAL TO OR LESS conditions.
- In case of text the search is not case sensitive, the % sign can be used as a joker character
- Filters can be removed by pressing the delete button at the end of the row

{{< hint info >}}

Example for the joker character usage: ***%product*** (all products which have in their name the word product, anywhere)

{{< /hint >}}

### TOTAL
Quick summary of the result rows of number columns. If there are no number type columns in the data view then it is inactive.

### SEARCH
Start search according to the specified conditions.

### BOOKMARK
Set a bookmark. It helps to save any number of pre-defined filters related to the data view. Later can be loaded from bookmarks at any time. The bookmark besides the selection criteria keeps also the selection of displayed columns.

### EXPORT
Export data in CSV (comma-separeted value) files. The result of search will be exported including all columns, regardless whether originally those were displayed on the screen or not. Further information: [**DATA EXPORT**](/docs/client/program/export)

{{< hint info >}}

The filter part of the screen can be closed or opened by clicking on its header.

{{< /hint >}}

{{< hint info >}}

The Filter field in the results section will filter only among displayed result rows and will not query further data from the database!

{{< /hint >}}

{{< hint info >}}

The number of result rows displayed on a page can be set in the [**PROGRAM SETTINGS**](/docs/client/program/psetting) section!

{{< /hint >}}

