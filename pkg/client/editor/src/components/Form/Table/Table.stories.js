import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './form-table.js';
import '../Button/form-button.js';

import { PAGINATION_TYPE, BUTTON_TYPE, APP_THEME } from '../../../config/enums.js'

export default {
  title: 'Form/Table',
  component: 'form-table',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
    },
    pagination: {
      control: 'select',
      options: Object.values(PAGINATION_TYPE),
    },
    onRowSelected: {
      name: "onRowSelected",
      description: "onRowSelected click handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onRowSelected" 
    },
    onEditCell: {
      name: "onEditCell",
      description: "onEditCell click handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onEditCell" 
    },
    onAddItem: {
      name: "onAddItem",
      description: "onAddItem click handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onAddItem" 
    },
    onCurrentPage: {
      name: "onCurrentPage",
      description: "onCurrentPage click handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onCurrentPage" 
    },
  }
};

export function Template({ 
  id, theme, name, rowKey, rows, fields, 
  pagination, currentPage, pageSize, hidePaginatonSize, tableFilter, filterPlaceholder, filterValue,
  labelYes, labelNo, labelAdd, addIcon, tablePadding, style,
  onRowSelected, onEditCell, onAddItem, onCurrentPage
}) {
  const component = html`<form-table
    id="${id}"
    name="${name}"
    rowKey="${rowKey}"
    .rows="${rows}"
    .fields="${fields}"
    pagination="${pagination}"
    currentPage="${currentPage}"
    pageSize="${pageSize}"
    ?tableFilter="${tableFilter}"
    ?hidePaginatonSize="${hidePaginatonSize}"
    filterPlaceholder="${filterPlaceholder}"
    filterValue="${filterValue}"
    labelYes="${labelYes}"
    labelNo="${labelNo}"
    labelAdd="${labelAdd}"
    addIcon="${addIcon}"
    tablePadding="${tablePadding}"
    .style="${style}"
    .onRowSelected=${onRowSelected}
    .onEditCell=${onEditCell}
    .onAddItem=${onAddItem}
    .onCurrentPage=${onCurrentPage}
  ></form-table>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "test_table",
  theme: APP_THEME.LIGHT,
  name: undefined,
  rowKey: "id",
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
    { id: 4, name: "Name4", "levels": 401234.345, valid: 0, 
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
    name: { fieldtype:'string', label: "Name", textAlign: "left" },
    valid: { fieldtype:'bool', label: "Valid" },
    from: { fieldtype:'date', label: "From" },
    start: { fieldtype:'time' },
    stamp: { fieldtype:'datetime', label: "Stamp" },
    levels: { fieldtype: 'number', label: "Levels", format: true, verticalAlign: "middle" },
    deffield: { fieldtype: 'deffield', label: "Deffield" },
    editor: { columnDef: { 
      id: "editor",
      Header: "",
      headerStyle: {},
      Cell: ({ row }) => html`<form-button 
        type="${BUTTON_TYPE.PRIMARY}" ?full="${true}" ?disabled="${row.disabled}" ?small="${true}" >Hello</form-button>`,
      cellStyle: {}
    }},
    id: { columnDef: {
      id: "id",
      cellStyle: {
        color: "red"
      }
    }}
  },
  pagination: PAGINATION_TYPE.NONE,
  currentPage: 1,
  pageSize: 10,
  hidePaginatonSize: false,
  tableFilter: false,
  filterPlaceholder: undefined,
  filterValue: "",
  labelYes: "YES",
  labelNo: "NO",
  labelAdd: "",
  addIcon: "Plus",
  tablePadding: undefined,
  style: {}
};

export const TopPagination = Template.bind({});
TopPagination.args = {
  ...Default.args,
  fields: {},
  rows: [
    {"name":"Fluffy","age":9,"breed":"calico","gender":"male"},
    {"name":"Luna","age":10,"breed":"long hair","gender":"female"},
    {"name":"Cracker","age":8,"breed":"fat","gender":"male"},
    {"name":"Pig","age":6,"breed":"calico","gender":"female"},
    {"name":"Robin","age":7,"breed":"long hair","gender":"male"},
    {"name":"Sammy","age":13,"breed":"fat","gender":"male"},
    {"name":"Aliece","age":9,"breed":"long hair","gender":"female"},
    {"name":"Mehatable","age":5,"breed":"calico","gender":"female"},
    {"name":"Scorpia","age":6,"breed":"long hair","gender":"female"},
    {"name":"Zoomies","age":1,"breed":"fat","gender":"male"},
    {"name":"Zues","age":5,"breed":"long hair","gender":"male"},
    {"name":"Lord Kittybottom","age":9,"breed":"calico","gender":"male"},
    {"name":"Princess Furball","age":5,"breed":"calico","gender":"female"},
    {"name":"Delerium","age":4,"breed":"fat","gender":"female"}
  ],
  pagination: PAGINATION_TYPE.TOP,
  currentPage: 2,
  pageSize: 5,
}

export const BottomPagination = Template.bind({});
BottomPagination.args = {
  ...TopPagination.args,
  theme: APP_THEME.DARK,
  pagination: PAGINATION_TYPE.BOTTOM,
  currentPage: 10,
  tableFilter: true,
  filterPlaceholder: "Placeholder text",
  addIcon: "Check",
  tablePadding: "16px"
}

export const Filtered = Template.bind({});
Filtered.args = {
  ...TopPagination.args,
  pagination: PAGINATION_TYPE.ALL,
  currentPage: 1,
  tableFilter: true,
  filterValue: "lo",
  labelAdd: "Add new",
}