import { Editor } from './Editor';

import { getText as appGetText, store, getSetting } from 'config/app';
import { Forms } from 'containers/Controller/Forms'

export default {
  title: "Editor/Editor",
  component: Editor,
}

const Template = (args) => <Editor {...args} />
const getText = (key)=>appGetText({ locales: store.session.locales, lang: "en", key: key })

export const Default = Template.bind({});
Default.args = {
  caption: "INVOICE", 
  current: {
    type: "trans",
    transtype: "invoice",
    item: { id: 5, transtype: 55, direction: 68, transnumber: "DMINV/00001", ref_transnumber: "DMORD/00003",
      crdate: "2022-01-12", transdate: "2021-12-10", duedate: "2021-12-20T00:00:00",
      customer_id: 2, employee_id: null, department: 138, project_id: null, place_id: null,
      paidtype: 123, curr: "EUR", notax: 0, paid: 0, acrate: 0, notes: null, intnotes: null,
      fnote: "<p>A long and <strong><em>rich text</em></strong> at the bottom of the invoice...</p><p>Can be multiple lines ...</p>",
      transtate: 93, cruser_id: 1, closed: 0, deleted: 0,
      amount: "849", balance: "0", cust_notax: 0, custname: "First Customer Co.",
      digit: 2, discount: 2, empnumber: null, expense: "0", income: "0", netamount: "720",
      planumber: null, pronumber: null, target_place: null, target_planumber: null,
      terms: 8, transcast: "normal", vatamount: "129"
    },
    state: "normal",
    page: 0,
    template: 1,
    fieldvalue: [
      { deleted: 0, fieldname: "trans_custinvoice_custname", id: 54,
        notes: null, ref_id: 5, value: "First Customer Co." },
      { deleted: 0, fieldname: "trans_transcast", id: 59, 
        notes: null, ref_id: 5, value: "normal" },
      { deleted: 0, fieldname: "trans_transitem_link",
        id: 99, notes: null, ref_id: 5, value: "5" },
      { deleted: 0, fieldname: "4e451b7f-72d1-b19c-7cbe-2c80495b5a8e",
        id: 100, notes: null, ref_id: 5, value: "blue" },
      { deleted: 0, fieldname: "2b0bd752-2f00-cbdb-a4ee-6741d890c8a6",
        id: 101, notes: null, ref_id: 5, value: "nervatura.github.io/" }
    ],
    view: "form"
  }, 
  template: Forms({ getText: getText })["invoice"](
    { id: 5, direction: 68, transcast: "normal", deleted: 0 },
    { dataset: { 
        translink: [ 
          { deleted: 0, id: 2, nervatype_1: 31, nervatype_2: 31, ref_id_1: 5,
            ref_id_2: 3, transnumber: "DMORD/00003", transtype: "order" }
        ],
        cancel_link: [], 
        groups: [
          { deleted: 0, description: null, groupname: "direction", groupvalue: "out", id: 68, inactive: 0 }
        ] 
      } 
    }), 
  dataset: {
    deffield: [
      { addnew: 0, deleted: 0, description: "Ref.No.", fieldname: "trans_transitem_link",
        fieldtype: 46, id: 1, nervatype: 31, readonly: 0, subtype: null, valuelist: null, visible: 1 },
      { addnew: 1, deleted: 0, description: "customer invoice company name",
        fieldname: "trans_custinvoice_compname", fieldtype: 38, id: 3,
        nervatype: 31, readonly: 1, subtype: 55, valuelist: null, visible: 0 },
      { addnew: 0, deleted: 0, description: "link qty value", fieldname: "link_qty",
        fieldtype: 36, id: 17, nervatype: 16, readonly: 0, subtype: null, valuelist: null, visible: 0 },
      { addnew: 0, deleted: 0, description: "link rate value", fieldname: "link_rate",
        fieldtype: 36, id: 18, nervatype: 16, readonly: 0, subtype: null, valuelist: null, visible: 0 },
      { addnew: 1, deleted: 0, description: "rent bad machine", fieldname: "trans_rebadtool",
        fieldtype: 36, id: 14, nervatype: 31, readonly: 1, subtype: 60, valuelist: null, visible: 0 },
      { addnew: 1, deleted: 0, description: "transaction special state", fieldname: "trans_transcast",
        fieldtype: 39, id: 2, nervatype: 31, readonly: 1, subtype: null,
        valuelist: "normal|cancellation|amendment", visible: 0 },
      { addnew: 1, deleted: 0, description: "customer invoice customer name", fieldname: "trans_custinvoice_custname",
        fieldtype: 38, id: 6, nervatype: 31, readonly: 1, subtype: 55, valuelist: null, visible: 0 },
      { addnew: 0, deleted: 0, description: "Test urlink", fieldname: "2b0bd752-2f00-cbdb-a4ee-6741d890c8a6",
        fieldtype: 42, id: 61, nervatype: 31, readonly: 0, subtype: null, valuelist: null, visible: 1 },
      { addnew: 0, deleted: 0, description: "Test valuelist", fieldname: "4e451b7f-72d1-b19c-7cbe-2c80495b5a8e",
        fieldtype: 39, id: 60, nervatype: 31, readonly: 0, subtype: null, valuelist: "red|blue|green", visible: 1 },
      { addnew: 0, deleted: 0, description: "Sample date", fieldname: "sample_customer_date",
        fieldtype: 34, id: 48, nervatype: 10, readonly: 1, subtype: null, valuelist: null, visible: 1 },
      { addnew: 0, deleted: 0, description: "Sample float", fieldname: "sample_customer_float",
        fieldtype: 36, id: 47, nervatype: 10, readonly: 0, subtype: null, valuelist: null, visible: 1 },
    ],
    deffield_prop: [
      { description: "DMINV/00001", ftype: "transitem", id: 5 }
    ],
    fieldvalue: [
      { deleted: 0, fieldname: "trans_custinvoice_custname",
        id: 54, notes: null, ref_id: 5, value: "First Customer Co." },
      { deleted: 0, fieldname: "trans_transcast", 
        id: 59, notes: null, ref_id: 5, value: "normal" },
      { deleted: 0, fieldname: "trans_transitem_link",
        id: 99, notes: null, ref_id: 5, value: "5" },
      { deleted: 0, fieldname: "4e451b7f-72d1-b19c-7cbe-2c80495b5a8e",
        id: 100, notes: null, ref_id: 5, value: "blue" },
      { deleted: 0, fieldname: "2b0bd752-2f00-cbdb-a4ee-6741d890c8a6",
        id: 101, notes: null, ref_id: 5, value: "nervatura.github.io/" }
    ],
    pattern: [
      { defpattern: 1, deleted: 0, description: "first template", id: 1,
        notes: "<p>A long and <strong><em>rich text</em></strong> at the bottom of the invoice...</p><p>Can be multiple lines ...</p><p>Can be multiple lines ...</p><p>Can be multiple lines ...</p>",
        transtype: 55 },
      { defpattern: 0, deleted: 0, description: "default pattern", id: 2, notes: null, transtype: 55 }
    ],
    item: [
      { actionprice: 0, amount: 144, deleted: 0, deposit: 0, description: "Very good work!", discount: 0,
        fxprice: 120, id: 18, netamount: 120, ownstock: 0, partnumber: "DMPROD/00002",
        product_id: 2, qty: 1, rate: 0.2, tax_id: 5, trans_id: 5, unit: "hour", vatamount: 24 },
      { actionprice: 0, amount: 600, deleted: 0, deposit: 0, description: "Big product", discount: 0,
        fxprice: 166.67, id: 19, netamount: 500, ownstock: 0, partnumber: "DMPROD/00001", product_id: 1,
        qty: 3, rate: 0.2, tax_id: 5, trans_id: 5, unit: "piece", vatamount: 100 },
      { actionprice: 0, amount: 105, deleted: 0, deposit: 0, description: "Nice product", discount: 0,
        fxprice: 20, id: 20, netamount: 100, ownstock: 0, partnumber: "DMPROD/00003", product_id: 3,
        qty: 5, rate: 0.05, tax_id: 2, trans_id: 5, unit: "piece", vatamount: 5 }
    ],
    invoice_link: [
      { curr: "EUR", deleted: 0, id: 6, lslabel: "2021-12-20 | Bank | 849.0",
        lsvalue: "DMPMT/00001 ~ 2 | bank-transfer",
        nervatype_1: 22, nervatype_2: 31, ref_id_1: 2, ref_id_2: 5,
        trans_id: 10, transnumber: "DMPMT/00001 ~ 2", transtype: "bank" }
    ],
    tool_movement: [],
    groups: [
      { deleted: 0, description: null, groupname: "direction", groupvalue: "in",
        id: 69, inactive: 0 },
      { deleted: 0, description: null, groupname: "direction", groupvalue: "out",
        id: 68, inactive: 0 },
      { deleted: 0, description: null, groupname: "direction", groupvalue: "transfer",
        id: 70, inactive: 0 },
      { deleted: 0, description: null, groupname: "fieldtype", groupvalue: "bool",
        id: 33, inactive: 0 },
      { deleted: 0, description: null, groupname: "fieldtype", groupvalue: "customer",
        id: 44, inactive: 0 },
      { deleted: 0, description: null, groupname: "fieldtype", groupvalue: "date",
        id: 34, inactive: 0 },
      { deleted: 0, description: null, groupname: "fieldtype", groupvalue: "float",
        id: 36, inactive: 0 },
      { deleted: 0, description: null, groupname: "fieldtype", groupvalue: "string",
        id: 38, inactive: 0 },
      { deleted: 0, description: null, groupname: "fieldtype", groupvalue: "place",
        id: 52, inactive: 0 },
      { deleted: 0, description: null, groupname: "fieldtype", groupvalue: "product",
        id: 49, inactive: 0 },
      { deleted: 0, description: null, groupname: "fieldtype", groupvalue: "transitem",
        id: 46, inactive: 0 },
      { deleted: 0, description: null, groupname: "fieldtype", groupvalue: "transmovement",
        id: 47, inactive: 0 },
      { deleted: 0, description: null, groupname: "fieldtype", groupvalue: "transpayment",
        id: 48, inactive: 0 },
      { deleted: 0, description: null, groupname: "fieldtype", groupvalue: "urlink",
        id: 42, inactive: 0 },
      { deleted: 0, description: null, groupname: "fieldtype", groupvalue: "valuelist",
        id: 39, inactive: 0 },
      { deleted: 0, description: null, groupname: "nervatype", groupvalue: "customer",
        id: 10, inactive: 0 },
      { deleted: 0, description: null, groupname: "nervatype", groupvalue: "item",
        id: 15, inactive: 0 },
      { deleted: 0, description: null, groupname: "nervatype", groupvalue: "link",
        id: 16, inactive: 0 },
      { deleted: 0, description: null, groupname: "nervatype", groupvalue: "trans",
        id: 31, inactive: 0 },
    ],
    direction: [
      { deleted: 0, description: null, groupname: "direction",
        groupvalue: "in", id: 69, inactive: 0 },
      { deleted: 0, description: null, groupname: "direction",
        groupvalue: "out", id: 68, inactive: 0 }
    ],
    transtate: [
      { deleted: 0, description: null, groupname: "transtate", groupvalue: "back", id: 95, inactive: 0 },
      { deleted: 0, description: null, groupname: "transtate", groupvalue: "new", id: 94, inactive: 0 },
      { deleted: 0, description: null, groupname: "transtate", groupvalue: "ok", id: 93, inactive: 0 }
    ],
    currency: [
      { cround: 0, curr: "EUR", defrate: 0, description: "euro", digit: 2, id: 1 },
      { cround: 0, curr: "USD", defrate: 0, description: "dollar", digit: 2, id: 2 }
    ],
    paidtype: [
      { deleted: 0, description: "cash", groupname: "paidtype", groupvalue: "cash", id: 122, inactive: 0 },
      { deleted: 0, description: "credit card", groupname: "paidtype", groupvalue: "credit_card", id: 124, inactive: 0 },
      { deleted: 0, description: "transfer", groupname: "paidtype", groupvalue: "transfer", id: 123, inactive: 0 }
    ],
    department: [
      { deleted: 0, description: "Sample logistics department", groupname: "department",
        groupvalue: "logistics", id: 139, inactive: 0 },
      { deleted: 0, description: "Sample production department", groupname: "department",
        groupvalue: "production", id: 140, inactive: 0 },
      { deleted: 0, description: "Sample sales department", groupname: "department",
        groupvalue: "sales", id: 138, inactive: 0 }
    ],
    translink: [
      { deleted: 0, id: 2, nervatype_1: 31, nervatype_2: 31, ref_id_1: 5,
        ref_id_2: 3, transnumber: "DMORD/00003", transtype: "order" }
    ],
    cancel_link: [],
    trans: [
      { id: 5, transtype: 55, direction: 68, transnumber: "DMINV/00001", ref_transnumber: "DMORD/00003",
        crdate: "2022-01-12", transdate: "2021-12-10", duedate: "2021-12-20T00:00:00",
        customer_id: 2, employee_id: null, department: 138, project_id: null, place_id: null,
        paidtype: 123, curr: "EUR", notax: 0, paid: 0, acrate: 0, notes: null, intnotes: null,
        fnote: "<p>A long and <strong><em>rich text</em></strong> at the bottom of the invoice...</p><p>Can be multiple lines ...</p>",
        transtate: 93, cruser_id: 1, closed: 0, deleted: 0,
        amount: "849", balance: "0", cust_notax: 0, custname: "First Customer Co.",
        digit: 2, discount: 2, empnumber: null, expense: "0", income: "0", netamount: "720",
        planumber: null, pronumber: null, target_place: null, target_planumber: null,
        terms: 8, transcast: "normal", vatamount: "129"
      }
    ],
    tax: [
      { description: "VAT 5%", id: 2, inactive: 0, rate: 0.05, taxcode: "5%" },
      { description: "VAT 20%", id: 5, inactive: 0, rate: 0.2, taxcode: "20%" },
    ]
  }, 
  audit: "all",
  className: "light",
  onEvent: undefined,
  getText: getText
}

export const New = Template.bind({});
New.args = {
  ...Default.args,
  current: {
    ...Default.args.current,
    item: {
      ...Default.args.current.item,
      id: null,
      fnote: undefined
    }
  },
  template: Forms({ getText: getText })["invoice"](
    { id: null, direction: 68, transcast: "normal", deleted: 0 },
    { dataset: { 
        translink: [],
        cancel_link: [], 
        groups: [
          { deleted: 0, description: null, groupname: "direction", groupvalue: "out", id: 68, inactive: 0 }
        ] 
      } 
    }),
  dataset: {
    ...Default.args.dataset,
    item: [],
    invoice_link: []
  }
}

export const Report = Template.bind({});
Report.args = {
  caption: "REPORT", 
  current: {
    type: "report",
    transtype: "",
    item: { id: 7, reportkey: "ntr_custpos_en", nervatype: 28, transtype: null, direction: null,
      repname: "Payments Due List", description: "Accounts Payable and Receivable", label: "Invoice",
      filetype: 54, report: "", orientation: "portrait", size: "a4", ftype: "pdf" },
    state: "normal",
    page: 0,
    fieldvalue: [
      { id: 1, rowtype: "reportfield", datatype: "date", name: "posdate",
        label: "Date", selected: true, empty: "false", value: "2022-01-16" },
      { id: 2, rowtype: "reportfield", datatype: "string", name: "curr",
        label: "Currency", selected: false, empty: "true", value: "" },
      { id: 3, rowtype: "reportfield", datatype: "date", name: "transdate_from",
        label: "Inv. Date >=", selected: false, empty: "true", value: "" },
      { id: 4, rowtype: "reportfield", datatype: "date", name: "transdate_to",
        label: "Inv. Date <=", selected: false, empty: "true", value: "" },
      { id: 5, rowtype: "reportfield", datatype: "date", name: "duedate_from",
        label: "Due Date >=", selected: false, empty: "true", value: "" },
      { id: 6, rowtype: "reportfield", datatype: "date", name: "duedate_to",
        label: "Due Date <=", selected: false, empty: "true", value: "" },
      { id: 7, rowtype: "reportfield", datatype: "string", name: "customer",
        label: "Customer/Supplier", selected: false, empty: "true", value: "" }
    ],
    view: "form"
  },
  template: Forms({ getText: getText })["report"]({ ftype: "pdf" }, {}, getSetting("ui")),
  dataset: {
    report: [
      { id: 7, reportkey: "ntr_custpos_en", nervatype: 28, transtype: null, direction: null,
        repname: "Payments Due List", description: "Accounts Payable and Receivable", label: "Invoice",
        filetype: 54, report: "", orientation: "portrait", size: "a4", ftype: "pdf" }
    ]
  },
  audit: "all",
  className: "light",
  onEvent: undefined,
  getText: getText
}

export const PrintQueue = Template.bind({});
PrintQueue.args = {
  caption:" REPORT QUEUE",
  current: {
    type: "printqueue",
    transtype: null,
    item: { id: null, nervatype: null, startdate: null, enddate: null,
      transnumber: null, username: null, server: null, mode: "pdf",
      orientation: "portrait", size: "a4" },
    state: "normal",
    page: 0,
    fieldvalue: [],
    view: "form"
  },
  dataset: {
    items: [],
    report: [],
    server_printers: [],
    printqueue: [
      { id: null, nervatype: null, startdate: null, enddate: null,
        transnumber: null, username: null, server: null, mode: "pdf",
        orientation: "portrait", size: "a4" }
    ]
  },
  template: Forms({ getText: getText })["printqueue"]({}, {}, getSetting("ui")),
  audit: "all",
  className: "light",
  onEvent: undefined,
  getText: getText
}

export const Meta = Template.bind({});
Meta.args = {
  ...Default.args,
  current: {
    ...Default.args.current,
    view: "fieldvalue"
  }
}

export const Notes = Template.bind({});
Notes.args = {
  ...Default.args,
  current: {
    ...Default.args.current,
    view: "fnote"
  },
  className: "dark",
}

export const Item = Template.bind({});
Item.args = {
  ...Default.args,
  current: {
    ...Default.args.current,
    form_type: "item",
    form_datatype: "item",
    form: { actionprice: 0, amount: 144, deleted: 0, deposit: 0, description: "Very good work!",
      discount: 0, fxprice: 120, id: 18, netamount: 120, ownstock: 0, partnumber: "DMPROD/00002",
      product_id: 2, qty: 1, rate: 0.2, tax_id: 5, trans_id: 5, unit: "hour", vatamount: 24 },
    form_template: Forms({ getText: getText })["item"]({ id: 18 }, { current: { transtype: "invoice" } })
  },
}

export const View = Template.bind({});
View.args = {
  ...Default.args,
  current: {
    ...Default.args.current,
    view: "item"
  },
}