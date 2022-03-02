import update from 'immutability-helper';
import { Browser } from "./Browser";

import { getText as appGetText , store } from 'config/app';
import { Queries } from 'containers/Controller/Queries'

export default {
  title: "Browser",
  component: Browser,
}

const Template = (args) => <Browser {...args} />
const getText = (key)=>appGetText({ locales: store.session.locales, lang: "en", key: key })

export const Default = Template.bind({});
Default.args = {
  data: {
    vkey: "customer",
    view: "CustomerView",
    show_header: true, 
    show_dropdown: false,
    show_columns: false,
    result: [
      {
        account: null, address: "City1 street 1.", creditlimit: 1000000, custname: "First Customer Co.",
        custnumber: "DMCUST/00001", custtype: "company", discount: 2, export_creditlimit: 1000000,
        export_terms: 8, id: "customer//2", inactive: 0, notax: 0, notes: null,
        row_id: 2, taxnumber: "12345678-1-12", terms: 8
      },
      {
        account: null, address: "City3 street 3.", creditlimit: 0, custname: "Second Customer Name",
        custnumber: "DMCUST/00002", custtype: "private", discount: 6, export_creditlimit: 0,
        export_terms: 1, id: "customer//3", inactive: 0, notax: 0, notes: null,
        row_id: 3, taxnumber: "12121212-1-12", terms: 1
      },
      {
        account: null, address: "City4 street 4.", creditlimit: 0, custname: "Third Customer Foundation",
        custnumber: "DMCUST/00003", custtype: "other", discount: 0, export_creditlimit: 0,
        export_terms: 4, id: "customer//4", inactive: 0, notax: 1, notes: null,
        row_id: 4, taxnumber: "10101010-1-01", terms: 4
      }
    ],
    columns: {
      CustomerView: {
        custnumber: true,
        custname: true,
        address: true,
        notax: true,
        creditlimit: true,
        terms: true
      }
    },
    filters: {
      CustomerView: []
    },
    deffield: [
      { description: "Sample float", fieldname: "sample_customer_float", fieldtype: "float", id: 47 },
      { description: "Sample date", fieldname: "sample_customer_date", fieldtype: "date", id: 48 },
      { description: "Sample valuelist", fieldname: "sample_customer_valuelist", fieldtype: "valuelist", id: 49 },
      { description: "Sample customer", fieldname: "sample_customer_reference", fieldtype: "customer", id: 50 },
      { description: "Sample bool", fieldname: "sample_customer_bool", fieldtype: "bool", id: 60 },
      { description: "Sample integer", fieldname: "sample_customer_integer", fieldtype: "integer", id: 55 },
    ], 
    page: 0,
  },
  keyMap: Queries({ getText: getText }).customer(),
  viewDef: Queries({ getText: getText }).customer().CustomerView,
  className: "light",
  onEvent: undefined,
  getText: getText,
  paginationPage: 2
}

export const HideHeader = Template.bind({});
HideHeader.args = {
  ...Default.args,
  className: "dark",
  data: {
    ...Default.args.data,
    show_header: false
  }
}

export const Columns = Template.bind({});
Columns.args = {
  ...Default.args,
  className: "dark",
  data: {
    ...Default.args.data,
    view: "CustomerContactView",
    show_dropdown: true,
    show_columns: true,
    result: [],
    columns: {
      CustomerContactView: {
        custname: true,
        firstname: true,
        surname: true,
        phone: true
      }
    },
    filters: {
      CustomerContactView: []
    },
  },
  viewDef: Queries({ getText: getText }).customer().CustomerContactView,
}

export const Filters = Template.bind({});
Filters.args = {
  ...Default.args,
  data: {
    ...Default.args.data,
    view: "CustomerFieldsView",
    result: [
      { custname: 'Second Customer Name', custnumber: 'DMCUST/00002', deffield_value: 'customer//2',
        export_deffield_value: 'First Customer Co.', fielddef: 'Sample customer',
        fieldname: 'sample_customer_reference', fieldtype: 'customer', form: 'fieldvalue',
        id: 'customer//3', notes: '', row_id: 31 },
      { custname: 'Second Customer Name', custnumber: 'DMCUST/00002', deffield_value: '56789.67',
        export_deffield_value: '56789.67', fielddef: 'Sample float', fieldname: 'sample_customer_float',
        fieldtype: 'float', form: 'fieldvalue', id: 'customer//3', notes: '', row_id: 32 },
      { custname: 'Second Customer Name', custnumber: 'DMCUST/00002', deffield_value: '56789',
        export_deffield_value: '', fielddef: 'Sample integer', fieldname: 'sample_customer_integer',
        fieldtype: 'integer', form: 'fieldvalue', id: 'customer//3', notes: '', row_id: 33 },
    ],
    columns: {
      CustomerFieldsView: {
        custname: true,
        fielddef: true,
        deffield_value: true
      }
    },
    filters: {
      CustomerFieldsView: [
        { id: "1641589935696", fieldtype: "float", fieldname: "sample_customer_float",
          sqlstr: "fg.groupvalue=\"float\" and {FMSF_NUMBER} {CAS_FLOAT}fv.value {CAE_FLOAT} {FMEF_CONVERT} ",
          wheretype: "where", filtertype: "!==", value: 123,
          fieldlimit: ["and", "fv.fieldname", "=", "\"sample_customer_float\""] },
        { id: "1641589951968", fieldtype: "date", fieldname: "sample_customer_date",
          sqlstr: "fg.groupvalue=\"date\" and {FMSF_DATE} {CASF_DATE}fv.value{CAEF_DATE} {FMEF_CONVERT} ",
          wheretype: "where", filtertype: ">==", value: "2022-01-01",
          fieldlimit: ["and", "fv.fieldname", "=", "\"sample_customer_date\""] },
        { id: "1641589963923", fieldtype: "bool", fieldname: "sample_customer_bool",
          sqlstr: "fg.groupvalue=\"bool\" and case when fv.value=\"true\" then 1 else 0 end ",
          wheretype: "where", filtertype: "===", value: "1",
          fieldlimit: ["and", "fv.fieldname", "=", "\"sample_customer_bool\""] },
        { id: "1641589990034", fieldtype: "string", fieldname: "sample_customer_valuelist",
          sqlstr: "fv.value ", wheretype: "where", filtertype: "===", value: "blue",
          fieldlimit: ["and", "fv.fieldname", "=", "\"sample_customer_valuelist\""] },
        { id: "1641602672808", fieldtype: "string", fieldname: "notes",
          sqlstr: "fv.notes ", wheretype: "where", filtertype: "==N", value: "" },
        { id: "1641589935696", fieldtype: "integer", fieldname: "sample_customer_integer",
          sqlstr: "fg.groupvalue=\"integer\" and {FMSF_NUMBER} {CAS_INT}fv.value {CAE_INT} {FMEF_CONVERT} ",
          wheretype: "where", filtertype: "<==", value: 5555,
          fieldlimit: ["and", "fv.fieldname", "=", "\"sample_customer_integer\""] },
      ]
    },
  },
  viewDef: update(Queries({ getText: getText }).customer().CustomerFieldsView, {$merge: {
    readonly: true
  }}),
}

export const HavingFilter = Template.bind({});
HavingFilter.args = {
  ...Default.args,
  data: {
    vkey: "transmovement",
    view: "InventoryView",
    show_header: true, 
    show_dropdown: false,
    show_columns: false,
    result: [],
    columns: {
      InventoryView: {
        warehouse: true,
        partnumber: true,
        unit: true,
        sqty: true
      }
    },
    filters: {
      InventoryView: [
        { id: "1641760171724", fieldtype: "date", fieldname: "posdate", 
          sqlstr: "{CAS_DATE}max(mv.shippingdate){CAE_DATE} ",
          wheretype: "having", filtertype: "==N", value: "2021-01-09" },
        { id: "1641760287189", fieldtype: "float", fieldname: "sqty",
          sqlstr: "sum(mv.qty) ", wheretype: "having", filtertype: ">==", value: 0 },
        { id: "1641760309207", fieldtype: "string", fieldname: "warehouse",
          sqlstr: "pl.description ", wheretype: "where", filtertype: "===", value: "wa" },
        { id: "1641760328875", fieldtype: "string", fieldname: "description",
          sqlstr: "p.description ", wheretype: "where", filtertype: "==N", value: "" }
      ]
    },
    deffield: [
      { description: "Ref.No.", fieldname: "trans_transitem_link", fieldtype: "transitem", id: 1 }
    ], 
    page: 0,
  },
  keyMap: Queries({ getText: getText }).transmovement(),
  viewDef: Queries({ getText: getText }).transmovement().InventoryView,
}

export const FormActions = Template.bind({});
FormActions.args = {
  ...Default.args,
  data: {
    vkey: "rate",
    view: "RateView",
    show_header: true, 
    show_dropdown: false,
    show_columns: false,
    result: [],
    columns: {
      RateView: {
        ratetype: true,
        ratedate: true,
        curr: true,
        ratevalue: true
      }
    },
    filters: {
      RateView: []
    },
    deffield: [], 
    page: 0,
  },
  keyMap: Queries({ getText: getText }).rate(),
  viewDef: Queries({ getText: getText }).rate().RateView,
}