import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './nt-list.js';

import { PAGINATION_TYPE } from './NtList.js'

export default {
  title: 'Form/NtList',
  component: 'nt-list',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: ["light", "dark"],
    },
    pagination: {
      control: 'select',
      options: Object.values(PAGINATION_TYPE),
    },
    onEdit: {
      name: "onEdit",
      description: "onEdit click handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onEdit" 
    },
    onDelete: {
      name: "onDelete",
      description: "onDelete click handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onDelete" 
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
  id, theme, name, rows,  
  pagination, currentPage, pageSize, hidePaginatonSize, listFilter, filterPlaceholder, filterValue,
  labelAdd, addIcon, editIcon, deleteIcon, style,
  onEdit, onDelete, onAddItem, onCurrentPage
}) {
  const component = html`<nt-list
    id="${id}"
    name="${name}"
    .rows="${rows}"
    pagination="${pagination}"
    currentPage="${currentPage}"
    pageSize="${pageSize}"
    ?listFilter="${listFilter}"
    ?hidePaginatonSize="${hidePaginatonSize}"
    filterPlaceholder="${filterPlaceholder}"
    filterValue="${filterValue}"
    labelAdd="${labelAdd}"
    addIcon="${addIcon}"
    editIcon="${editIcon}"
    deleteIcon="${deleteIcon}"
    .style="${style}"
    .onEdit=${onEdit}
    .onDelete=${onDelete}
    .onAddItem=${onAddItem}
    .onCurrentPage=${onCurrentPage}
  ></nt-list>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "test_list",
  theme: "light",
  name: undefined,
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
  pagination: PAGINATION_TYPE.NONE,
  currentPage: 1,
  pageSize: 5,
  hidePaginatonSize: false,
  listFilter: false,
  filterPlaceholder: undefined,
  filterValue: "",
  labelAdd: "",
  addIcon: "Plus",
  editIcon: "Edit",
  deleteIcon: "Times",
  style: {},
  onEdit: null,
  onAddItem: null,
  onDelete: null,
}

export const TopPagination = Template.bind({});
TopPagination.args = {
  ...Default.args,
  theme: "dark",
  pagination: PAGINATION_TYPE.TOP,
  currentPage: 10,
  hidePaginatonSize: false,
  listFilter: false,
  labelAdd: "Add new",
}

export const BottomPagination = Template.bind({});
BottomPagination.args = {
  ...Default.args,
  pagination: PAGINATION_TYPE.BOTTOM,
  currentPage: 0,
  hidePaginatonSize: true,
  listFilter: true,
  filterPlaceholder: "Placeholder text",
  addIcon: "User",
  editIcon: "Check",
  deleteIcon: "Close"
}

export const Filtered = Template.bind({});
Filtered.args = {
  ...Default.args,
  pagination: PAGINATION_TYPE.ALL,
  listFilter: true,
  filterValue: "6",
  labelAdd: "Add new",
  onEdit: undefined,
  onAddItem: undefined,
  onDelete: undefined,
}