import { Form } from "./Form";

import { getText, store } from 'config/app';

export default {
  title: "Setting/Form",
  component: Form,
}

const Template = (args) => <Form {...args} />

export const Default = Template.bind({});
Default.args = {
  data: {
    caption: "DEFAULT SETTINGS",
    icon: "Cog",
    current: {
      form: { 
        description: "default currency", fieldname: "default_currency", fieldtype: "string",
        id: 6, lslabel: "default currency", lsvalue: "EUR", notes: null,
        value: "EUR", valuelist: null 
      },
      fieldvalue: {
        rowtype: "fieldvalue", id: 6, fieldname: "default_currency", fieldvalue_value: "EUR",
        fieldvalue_notes: "", label: "default currency", description: "EUR",
        disabled: "false", fieldtype: "string", datatype: "string"
      },
      template: {
        options: {
          icon: "Cog",
          data: "fieldvalue",
          title: "DEFAULT SETTINGS",
          panel: { page: "setting", delete: false, new: false, more: false, help: "setting" }
        },
        view: {
          setting: {
            type: "list",
            actions: {
              new: null,
              edit: { action: "editItem" },
              delete: null
            }
          }
        },
        rows: [
          { rowtype: "field", name: "fieldname", label: "Fieldname", datatype: "string", disabled: true },
          { rowtype: "field", name: "label", label: "Description", datatype: "string", disabled: true },
          { rowtype: "field", name: "fieldvalue_value", label: "Value", datatype: "fieldvalue" },
          { rowtype: "field", name: "fieldvalue_notes", label: "Other data", datatype: "text" }
        ]
      }
    },
    audit: "all", 
    dataset: {
      setting_view: [
        { 
          description: "default currency", fieldname: "default_currency", fieldtype: "string",
          id: 6, lslabel: "default currency", lsvalue: "EUR", notes: null,
          value: "EUR", valuelist: null 
        }
      ]
    },
    type: "setting", 
    view: {
      type: "list",
      result: [
        { description: "default currency", fieldname: "default_currency", fieldtype: "string",
          id: 6, lslabel: "default currency", lsvalue: "EUR", notes: null,
          value: "EUR", valuelist: null },
      ],
      fields: null,
    }
  },
  className: "light",
  onEvent: undefined,
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key })
}

export const Items = Template.bind({});
Items.args = {
  ...Default.args,
  data: {
    caption: "MENU SHORTCUTS",
    icon: "Share",
    current: {
      form: {
        address: null, description: "Server function example", funcname: "nextNumber", icon: null,
          id: 1, lslabel: "nextNumber", lsvalue: "Server function example",
          menukey: "nextNumber", method: 130, modul: null
      },
      template: {
        options: {
          icon: "Share",
          data: "ui_menu",
          title: "MENU SHORTCUTS",
          panel: { page: "setting", more: false, help: "menu" }
        },
        view: {
          setting: {
            type: "list",
            actions: {
              new: { action: "newItem" },
              edit: { action: "editItem" },
              delete: { action: "deleteItem" }
            }
          },
          items: {
            type: "table",
            data: "ui_menufields",
            actions: {
              new: { action: "editMenuField" },
              edit: { action: "editMenuField" },
              delete: { action: "deleteItemRow", table: "ui_menufields" }
            },
            fields: {
              fieldname: { fieldtype: "string", label: "Name" },
              description: { fieldtype: "string", label: "Description" },
              fieldtype_name: { fieldtype: "string", label: "Type" },
              orderby: { fieldtype: "number", label: "Order" }
            }
          }
        },
        rows: [
          {
            rowtype: "col2",
            columns: [
              { name: "menukey", label: "Menukey", datatype: "string", disabled: true },
              { name: "description", label: "Description", datatype: "string" }
            ]
          },
          {
            rowtype: "col3",
            columns: [
              { name: "method", label: "Method", datatype: "select",
                map: { source: "method", value: "id", text: "groupvalue"}},
              { name: "modul", label: "Modul", datatype: "string" },
              { name: "icon", label: "Icon", datatype: "string" }
            ]
          },
          {
            rowtype: "col2",
            columns: [
              { name: "funcname", label: "Funcname", datatype: "string" },
              { name: "address", label: "Address", datatype: "string" }
            ]
          }
        ]
      }
    },
    audit: "all", 
    dataset: {
      fieldtype: [
        { deleted: 0, description: null, groupname: "fieldtype", groupvalue: "bool", id: 33, inactive: 0 },
        { deleted: 0, description: null, groupname: "fieldtype", groupvalue: "date", id: 34, inactive: 0  },
        { deleted: 0, description: null, groupname: "fieldtype", groupvalue: "float", id: 36, inactive: 0 },
        { deleted: 0, description: null, groupname: "fieldtype", groupvalue: "integer", id: 37, inactive: 0 },
        { deleted: 0, description: null, groupname: "fieldtype", groupvalue: "string", id: 38, inactive: 0 }
      ],
      method: [
        { deleted: 0, description: "GET", groupname: "method", groupvalue: "get", id: 129, inactive: 0 },
        { deleted: 0, description: "POST", groupname: "method", groupvalue: "post", id: 130, inactive: 0 }
      ],
      ui_menu_view: [
        { address: null, description: "Server function example", funcname: "nextNumber", icon: null,
          id: 1, lslabel: "nextNumber", lsvalue: "Server function example",
          menukey: "nextNumber", method: 130, modul: null },
      ],
      ui_menufields: [
        { description: "Code (e.g. custnumber)", fieldname: "numberkey", fieldtype: 38,
          fieldtype_name: "string", id: 1, menu_id: 1, orderby: 0 },
        { description: "Stepping", fieldname: "step", fieldtype: 33,
          fieldtype_name: "bool", id: 2, menu_id: 1, orderby: 1 }
      ]
    },
    type: "ui_menu", 
    view: {
      type: "list",
      result: [
        {
          address: null, description: "Server function example", funcname: "nextNumber", icon: null,
          id: 1, lslabel: "nextNumber", lsvalue: "Server function example", 
          menukey: "nextNumber", method: 130, modul: null
        },
      ],
      fields: null
    },
  }
}

export const Log = Template.bind({});
Log.args = {
  ...Default.args,
  data: {
    caption: "DATABASE LOG",
    icon: "InfoCircle",
    current: {
      form: {
        id: null, fromdate: "2022-02-02", todate: "", empnumber: "", logstate: "update", nervatype: ""
      },
      template: {
        options: {
          title: "DATABASE LOG",
          title_field: "",
          edited: false,
          icon: "InfoCircle",
          panel: {}
        },
        view: {
          setting: {
            type: "table",
            actions: {
              new: null,
              edit: null,
              delete: null
            },
            fields: {
              crdate: { fieldtype: "date", label: "Date" },
              empnumber: { fieldtype: "string", label: "Employee No." },
              logstate: { fieldtype: "string", label: "State" },
              nervatype: { fieldtype: "string", label: "Type" },
              refnumber: { fieldtype: "string", label: "Doc.No./Description" }
            }
          }
        },
        rows: [
          {
            rowtype: "col3",
            columns: [
              { name: "fromdate", label: "Start Date", datatype: "date" },
              { name: "todate", label: "End Date", datatype: "date", empty: true },
              { name: "empnumber", label: "Employee No.", datatype: "string" }
            ]
          },
          {
            rowtype: "col3",
            columns: [
              { name: "logstate", label: "State", datatype: "select", empty: false,
                options: [
                  [ "update", "update"], [ "closed", "closed" ], [ "deleted", "deleted" ],
                  [ "print", "print" ], [ "login", "login" ], [ "logout", "logout" ]
                ]},
              { name: "nervatype", label: "Type", datatype: "select", default: "",
                options: [
                  [ "", "" ], [ "customer", "customer" ], [ "employee", "employee" ],
                  [ "event", "event" ], [ "place", "place" ], [ "product", "product" ],
                  [ "project", "project" ], [ "tool", "tool" ], [ "trans", "trans" ]
                ]},
              { name: "log_search", title: "Search", label: "", focus: true,
                class: "full", icon: "Search", datatype: "button" }
            ]
          }
        ]
      }
    },
    audit: "all", 
    dataset: {},
    type: "log", 
    view: {
      type: "table",
      fields: {
        crdate: { fieldtype: "date", label: "Date" },
        empnumber: { fieldtype: "string", label: "Employee No." },
        logstate: { fieldtype: "string", label: "State" },
        nervatype: { fieldtype: "string", label: "Type" },
        refnumber: { fieldtype: "string", label: "Doc.No./Description" }
      }
    },
  }
}