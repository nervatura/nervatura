import { TableView } from "./Table";
import Button from 'components/Form/Button'
import Icon from 'components/Form/Icon'

export default {
  title: "Form/Table",
  component: TableView
}

const Template = (args) => <TableView {...args} />

export const Default = Template.bind({});
Default.args = {
  className: "light",
  rows: [
    { id: 1, name: "Name1", "levels": 0, valid: "true", 
      from: "2020-06-06", start: "2020-04-23T10:30:00+02:00", stamp: "2020-04-23T10:30:00+02:00",
      name_color: "red", 
      export_deffield_value: "Customer 1", fieldtype: "customer", deffield: 123 },
    { id: 2, name: "Name2", export_name: "Name link", 
      "levels": 20, valid: 1, 
      from: "2020-06-06", start: "2020-04-23T10:30:00+02:00", stamp: "2020-04-23T10:30:00+02:00",
      name_color: "red", edited: true,
      fieldtype: "bool", deffield: "true" },
    { id: 3, name: "Name3", "levels": 40, valid: "false", 
      from: "2020-06-06", start: "2020-04-23T10:30:00+02:00", stamp: "2020-04-23T10:30:00+02:00",
      name_color: "orange", disabled: true,
      fieldtype: "integer", deffield: 123 },
    { id: 4, name: "Name4", "levels": 40, valid: 0, 
      from: "2020-06-06", start: "", stamp: "2020-04-23T10:30:00+02:00",
      name_color: "orange",
      fieldtype: "string", deffield: "value" },
    { id: 5, name: "Name5", "levels": 40, valid: false,
      from: "2020-06-06", start: "2020-04-23T10:30:00+02:00", stamp: "2020-04-23T10:30:00+02:00",
      export_deffield_value: "Customer 2", fieldtype: "customer", deffield: 124 },
    { id: 6, name: "Name6", "levels": 60, valid: true, 
      from: "2020-06-06", start: "2020-04-23T10:30:00+02:00", stamp: "2020-04-23T10:30:00+02:00",
      name_color: "green",
      export_deffield_value: "Customer 7", fieldtype: "customer", deffield: 222 }
  ],
  fields: {
    name: { fieldtype:'string', label: "Name" },
    valid: { fieldtype:'bool', label: "Valid" },
    from: { fieldtype:'date', label: "From" },
    start: { fieldtype:'time', label: "Start" },
    stamp: { fieldtype:'datetime', label: "Stamp" },
    levels: { fieldtype: 'number', label: "Levels", format: true, verticalAlign: "middle" },
    videos: { fieldtype: 'number', textAlign: "center" },
    deffield: { fieldtype: 'deffield', label: "Deffield" },
    editor: { columnDef: { 
      id: "editor",
      Header: "",
      headerStyle: {},
      Cell: ({ row, value }) => {
        return <Button className="primary full" label="Hello" />
      },
      cellStyle: {}
    }}
  },
  tableFilter: false,
  onRowSelected: null,
  onEditCell: null,
  onAddItem: null
}

export const TopPagination = Template.bind({});
TopPagination.args = {
  ...Default.args,
  className: "dark",
  paginationPage: 5,
  paginationTop: true,
  paginatonScroll: true,
  tableFilter: false,
  labelAdd: "Add new",
  tablePadding: 6,
  onRowSelected: undefined,
  onEditCell: undefined,
  onAddItem: undefined
}

export const BottomPagination = Template.bind({});
BottomPagination.args = {
  ...Default.args,
  paginationPage: 5,
  currentPage: 2,
  paginationTop: false,
  paginatonScroll: false,
  tableFilter: true,
  filterPlaceholder: "Placeholder text",
  onRowSelected: undefined,
  onEditCell: undefined,
  onAddItem: undefined,
  addIcon: <Icon iconKey="Check" />,
}

export const Filtered = Template.bind({});
Filtered.args = {
  ...Default.args,
  tableFilter: true,
  filterValue: "40",
  labelAdd: "Add new",
  dateFormat: null,
  timeFormat: null,
  onRowSelected: undefined,
  onEditCell: undefined,
  onAddItem: undefined,
}