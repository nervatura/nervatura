import { ListView } from "./List";
import Icon from 'components/Form/Icon'

export default {
  title: "Form/List",
  component: ListView
}

const Template = (args) => <ListView {...args} />

export const Default = Template.bind({});
Default.args = {
  className: "light",
  rows: [
    { lslabel: "Label 1", lsvalue: "Value row 1"},
    { lslabel: "Label 2", lsvalue: "Value row 2", id: 123},
    { lslabel: "Label 3", lsvalue: "Value row 3"},
    { lslabel: "Label 4", lsvalue: "Value row 6"},
    { lslabel: "Label 5", lsvalue: "Value row 6"},
    { lslabel: "Label 6", lsvalue: "Value row 6"},
    { lslabel: "Label 7", lsvalue: "Value row 7"},
    { lslabel: "Label 8", lsvalue: "Value row 8"},
    { lslabel: "Label 9", lsvalue: "Value row 9"}
  ],
  listFilter: false,
  paginationPage: 10,
  onEdit: undefined,
  onAddItem: null,
  onDelete: null
}

export const TopPagination = Template.bind({});
TopPagination.args = {
  ...Default.args,
  className: "dark",
  paginationPage: 5,
  paginationTop: true,
  paginatonScroll: true,
  hidePaginatonSize: false,
  listFilter: false,
  labelAdd: "Add new",
  onEdit: undefined,
  onAddItem: undefined,
  onDelete: undefined
}

export const BottomPagination = Template.bind({});
BottomPagination.args = {
  ...Default.args,
  paginationPage: 5,
  currentPage: 1,
  paginationTop: false,
  paginatonScroll: false,
  hidePaginatonSize: true,
  listFilter: true,
  filterPlaceholder: "Placeholder text",
  onEdit: undefined,
  onAddItem: undefined,
  onDelete: undefined,
  addIcon: <Icon iconKey="User" />,
  editIcon: <Icon iconKey="Check" />,
  deleteIcon: <Icon iconKey="Close" />,
}

export const Filtered = Template.bind({});
Filtered.args = {
  ...Default.args,
  listFilter: true,
  filterValue: "6",
  labelAdd: "Add new",
  onEdit: undefined,
  onAddItem: undefined,
  onDelete: undefined
}