import sinon from 'sinon'
import { expect } from '@open-wc/testing';

import { InitItem, Validator } from './Validator.js'
import { store as storeConfig } from '../config/app.js'
import { APP_MODULE } from '../config/enums.js'
import { Default as InvoiceData } from '../components/Edit/Editor/Editor.stories.js'

const store = {
  data: {
    ...storeConfig,
    [APP_MODULE.LOGIN]: {
      ...storeConfig[APP_MODULE.LOGIN],
      data: {
        engine: "sqlite",
        menuCmds: [],
        employee: {
          id: 1, empnumber: 'admin', username: 'admin', usergroup: 114,
          usergroupName: 'admin'
        },
        groups: [
          ...InvoiceData.args.dataset.groups
        ],
        audit: [
          { nervatypeName: "report", subtype: 1, inputfilterName: "disabled" }
        ]
      }
    },
    [APP_MODULE.EDIT]: {
      ...storeConfig[APP_MODULE.EDIT],
      ...InvoiceData.args,
    }
  },
  setData: sinon.spy(),
}
const app = {
  store,
  msg: (value, prop)=>prop.key,
  requestData: sinon.spy(async (path) => {
    if(path === "/view"){
      return { check: [{ recnum: 0 }] }
    }
    return "012345"
  }),
  getSetting: (key)=>storeConfig[key],
  getSql: () =>({
    sql: "",
    prmCount: 1
  }),
}

const testApp = (key, values, _app) => ({
  ...(_app||app),
  store: {
    ...(_app||app).store,
    data: {
      ...(_app||app).store.data,
      [APP_MODULE.EDIT]: {
        ...(_app||app).store.data[APP_MODULE.EDIT],
        [key]: {
          ...(_app||app).store.data[APP_MODULE.EDIT][key],
          ...values
        }
      }
    }
  }
})

describe('EditController', () => {
  beforeEach(async () => {
    Object.defineProperty(URL, 'createObjectURL', { value: sinon.spy() })
    Object.defineProperty(window, 'open', { value: sinon.spy() })
  });
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('InitItem', () => {
    let values = InitItem(app)({
      tablename: "address", 
      dataset: store.data[APP_MODULE.EDIT].dataset, 
      current: store.data[APP_MODULE.EDIT].current
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "audit"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "barcode"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "contact"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "currency"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "customer"
    })
    expect(values).to.exist

    values = InitItem(testApp("dataset", { custtype: undefined }))({
      tablename: "customer"
    })
    expect(values).to.null

    values = InitItem(app)({
      tablename: "deffield"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "employee"
    })
    expect(values).to.exist
    values = InitItem(testApp("dataset", { usergroup: undefined }))({
      tablename: "employee"
    })
    expect(values).to.null

    values = InitItem(app)({
      tablename: "event"
    })
    expect(values).to.exist
    values = InitItem(testApp("current", { type: "event" }))({
      tablename: "event"
    })
    expect(values).to.exist
    values = InitItem(testApp("current", { item: undefined }))({
      tablename: "event"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "fieldvalue"
    })
    expect(values).to.exist
    values = InitItem(testApp("current", { item: undefined }))({
      tablename: "fieldvalue"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "groups"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "usergroup"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "item"
    })
    expect(values).to.exist

    values = InitItem(testApp("current", { form_type: "invoice_link" }))({
      tablename: "link"
    })
    expect(values).to.exist
    values = InitItem(testApp("current", { form_type: "payment_link" }))({
      tablename: "link"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "log"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "ui_menu"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "ui_menufields"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "movement"
    })
    expect(values).to.exist
    values = InitItem(testApp("dataset", { movement_transfer: [] }, testApp("current", { transtype: "delivery" })))({
      tablename: "movement"
    })
    expect(values).to.exist
    values = InitItem(
      testApp("dataset", { movement_transfer: [{ place_id: 1 }] }, 
      testApp("current", { transtype: "delivery" }))
    )({
      tablename: "movement"
    })
    expect(values).to.exist
    values = InitItem(testApp("current", { transtype: "inventory" }))({
      tablename: "movement"
    })
    expect(values).to.exist
    values = InitItem(testApp("current", { transtype: "production" }))({
      tablename: "movement"
    })
    expect(values).to.exist
    values = InitItem(testApp("current", { transtype: "formula" }))({
      tablename: "movement"
    })
    expect(values).to.exist
    values = InitItem(testApp("current", { transtype: "waybill" }))({
      tablename: "movement"
    })
    expect(values).to.exist

    values = InitItem(testApp("current", { transtype: "formula" }))({
      tablename: "movement_head"
    })
    expect(values).to.exist
    values = InitItem(testApp("current", { transtype: "production" }))({
      tablename: "movement_head"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "numberdef"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "pattern"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "payment"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "place"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "discount"
    })
    expect(values).to.exist
    values = InitItem(testApp("dataset", { settings: [] }))({
      tablename: "price"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "product"
    })
    expect(values).to.exist
    values = InitItem(testApp("dataset", { settings: [] }))({
      tablename: "product"
    })
    expect(values).to.exist
    values = InitItem(testApp("dataset", { protype: undefined }))({
      tablename: "product"
    })
    expect(values).to.null

    values = InitItem(app)({
      tablename: "project"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "printqueue"
    })
    expect(values).to.exist
    values = InitItem(testApp("current", { type: "printqueue" }))({
      tablename: "printqueue"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "rate"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "refvalue"
    })
    expect(values).to.exist
    let _testApp = testApp("current", { transtype: "waybill" })
    values = InitItem(_testApp)({
      tablename: "refvalue"
    })
    expect(values).to.exist
    _testApp = testApp("dataset", { trans: [..._testApp.store.data[APP_MODULE.EDIT].dataset.trans] }, _testApp)
    _testApp.store.data[APP_MODULE.EDIT].dataset.trans[0] = {
      ..._testApp.store.data[APP_MODULE.EDIT].dataset.trans[0],
      customer_id: null,
      employee_id: 1
    }
    values = InitItem(_testApp)({
      tablename: "refvalue"
    })
    expect(values).to.exist
    _testApp.store.data[APP_MODULE.EDIT].dataset.trans[0] = {
      ..._testApp.store.data[APP_MODULE.EDIT].dataset.trans[0],
      customer_id: null,
      employee_id: null
    }
    values = InitItem(_testApp)({
      tablename: "refvalue"
    })
    expect(values).to.exist
    _testApp = testApp("dataset", { translink: [] }, _testApp)
    values = InitItem(_testApp)({
      tablename: "refvalue"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "report"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "tax"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "tool"
    })
    expect(values).to.exist

    values = InitItem(app)({
      tablename: "trans"
    })
    expect(values).to.exist
    _testApp = testApp("dataset", { settings: [], pattern: [] })
    values = InitItem(_testApp)({
      tablename: "trans", transtype: "invoice"
    })
    expect(values).to.exist
    values = InitItem(_testApp)({
      tablename: "trans", transtype: "formula"
    })
    expect(values).to.exist
    values = InitItem(_testApp)({
      tablename: "trans", transtype: "production"
    })
    expect(values).to.exist
    values = InitItem(_testApp)({
      tablename: "trans", transtype: "cash"
    })
    expect(values).to.exist
    values = InitItem(_testApp)({
      tablename: "trans", transtype: "offer"
    })
    expect(values).to.exist
    values = InitItem(_testApp)({
      tablename: "trans", transtype: "order"
    })
    expect(values).to.exist
    values = InitItem(_testApp)({
      tablename: "trans", transtype: "worksheet"
    })
    expect(values).to.exist
    values = InitItem(_testApp)({
      tablename: "trans", transtype: "rent"
    })
    expect(values).to.exist
    values = InitItem(_testApp)({
      tablename: "trans", transtype: "receipt"
    })
    expect(values).to.exist
    values = InitItem(_testApp)({
      tablename: "trans", transtype: "bank"
    })
    expect(values).to.exist
    values = InitItem(_testApp)({
      tablename: "trans", transtype: "inventory"
    })
    expect(values).to.exist
    _testApp = testApp("dataset", { pattern: undefined }, _testApp)
    values = InitItem(_testApp)({
      tablename: "trans"
    })
    expect(values).to.null

    values = InitItem(_testApp)({
      tablename: "missing"
    })
    expect(values).to.false

  })

  it('Validator', async () => {
    let result = Validator(app)(
      "address", {}
    )
    expect(result.error).to.not.exist

    result = await Validator(app)(
      "barcode", { code: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "barcode", { barcodetype: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "barcode", { id: null, code: "code" }
    )
    expect(result.error).to.not.exist
    result = await Validator(app)(
      "barcode", { id: 1 }
    )
    expect(result.error).to.not.exist

    result = await Validator(app)(
      "contact", {}
    )
    expect(result.error).to.not.exist

    result = await Validator(app)(
      "currency", { curr: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "currency", { description: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "currency", { id: null, curr: "curr" }
    )
    expect(result.error).to.not.exist
    result = await Validator(app)(
      "currency", { id: 1 }
    )
    expect(result.error).to.not.exist

    result = await Validator(app)(
      "customer", { custname: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "customer", { custtype: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "customer", { id: null, custnumber: "custnumber" }
    )
    expect(result.error).to.not.exist
    result = await Validator(app)(
      "customer", { id: null, custnumber: null }
    )
    expect(result.error).to.not.exist
    result = await Validator(app)(
      "customer", { id: 1 }
    )
    expect(result.error).to.not.exist

    result = await Validator(app)(
      "deffield", { nervatype: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "deffield", { fieldtype: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "deffield", { description: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "deffield", { id: 1 }
    )
    expect(result.error).to.not.exist

    let testApp_ = testApp("current", { extend: { surname: null } })
    result = await Validator(testApp_)(
      "employee", {}
    )
    expect(result.error).to.exist
    testApp_ = testApp("current", { extend: { surname: "surname"} }, testApp_)
    result = await Validator(testApp_)(
      "employee", { usergroup: null }
    )
    expect(result.error).to.exist
    result = await Validator(testApp_)(
      "employee", { id: null, empnumber: "empnumber" }
    )
    expect(result.error).to.not.exist
    result = await Validator(testApp_)(
      "employee", { id: null, empnumber: null }
    )
    expect(result.error).to.not.exist
    result = await Validator(testApp_)(
      "employee", { id: 1 }
    )
    expect(result.error).to.not.exist

    result = await Validator(app)(
      "event", { calnumber: null }
    )
    expect(result.error).to.not.exist
    result = await Validator(app)(
      "event", { id: 1 }
    )
    expect(result.error).to.not.exist

    result = await Validator(app)(
      "fieldvalue", {}
    )
    expect(result.error).to.not.exist

    result = await Validator(app)(
      "formula", {}
    )
    expect(result.error).to.not.exist

    result = await Validator(app)(
      "groups", { groupvalue: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "groups", { groupname: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "groups", { groupname: "usergroup", description: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "groups", { groupname: "usergroup", description: "description", transfilter: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "groups", { id: null }
    )
    expect(result.error).to.not.exist
    result = await Validator(app)(
      "groups", { id: 1 }
    )
    expect(result.error).to.not.exist

    result = await Validator(app)(
      "item", { product_id: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "item", { description: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "item", { tax_id: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "item", { id: 1 }
    )
    expect(result.error).to.not.exist

    result = await Validator(app)(
      "link", { ref_id_1: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "link", { id: 1 }
    )
    expect(result.error).to.not.exist

    result = await Validator(app)(
      "log", {}
    )
    expect(result.error).to.not.exist

    testApp_ = testApp("current", { transtype: "delivery", item: { direction: 68 } })
    result = await Validator(testApp_)(
      "movement", { place_id: null }
    )
    expect(result.error).to.exist
    result = await Validator(testApp_)(
      "movement", { product_id: null }
    )
    expect(result.error).to.exist
    testApp_ = testApp("current", { transtype: "delivery", item: { direction: 70, place_id: 1 } })
    result = await Validator(testApp_)(
      "movement", { place_id: null }
    )
    expect(result.error).to.exist
    result = await Validator(testApp_)(
      "movement", { place_id: 1 }
    )
    expect(result.error).to.exist
    result = await Validator(testApp_)(
      "movement", {}
    )
    expect(result.error).to.not.exist

    testApp_ = testApp("current", { transtype: "inventory" })
    result = await Validator(testApp_)(
      "movement", { product_id: null }
    )
    expect(result.error).to.exist
    result = await Validator(testApp_)(
      "movement", {}
    )
    expect(result.error).to.not.exist

    testApp_ = testApp("current", { transtype: "production" })
    result = await Validator(testApp_)(
      "movement", { product_id: null }
    )
    expect(result.error).to.exist
    result = await Validator(testApp_)(
      "movement", { place_id: null }
    )
    expect(result.error).to.exist
    result = await Validator(testApp_)(
      "movement", {}
    )
    expect(result.error).to.not.exist

    testApp_ = testApp("current", { transtype: "formula" })
    result = await Validator(testApp_)(
      "movement", { product_id: null }
    )
    expect(result.error).to.exist
    result = await Validator(testApp_)(
      "movement", {}
    )
    expect(result.error).to.not.exist

    testApp_ = testApp("current", { transtype: "waybill" })
    result = await Validator(testApp_)(
      "movement", { tool_id: null }
    )
    expect(result.error).to.exist
    result = await Validator(testApp_)(
      "movement", {}
    )
    expect(result.error).to.not.exist

    testApp_ = testApp("current", { transtype: "invoice" })
    result = await Validator(testApp_)(
      "movement", {}
    )
    expect(result.error).to.not.exist

    result = await Validator(app)(
      "numberdef", {}
    )
    expect(result.error).to.not.exist

    result = await Validator(app)(
      "pattern", {}
    )
    expect(result.error).to.not.exist

    result = await Validator(app)(
      "payment", {}
    )
    expect(result.error).to.not.exist

    result = await Validator(testApp_)(
      "place", { placetype: 126 }
    )
    expect(result.error).to.not.exist
    result = await Validator(testApp_)(
      "place", { id: 1, description: null }
    )
    expect(result.error).to.exist
    result = await Validator(testApp_)(
      "place", { id: 1, placetype: null }
    )
    expect(result.error).to.exist
    result = await Validator(testApp_)(
      "place", { id: 1, planumber: null }
    )
    expect(result.error).to.not.exist

    result = await Validator(app)(
      "price", { validfrom: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "price", { curr: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "price", { calcmode: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "price", {}
    )
    expect(result.error).to.not.exist
 
    result = await Validator(app)(
      "product", { description: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "product", { protype: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "product", { unit: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "product", { tax: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "product", { id: null, partnumber: "partnumber" }
    )
    expect(result.error).to.not.exist
    result = await Validator(app)(
      "product", { id: 1, partnumber: null }
    )
    expect(result.error).to.not.exist
    result = await Validator(app)(
      "product", { id: 1 }
    )
    expect(result.error).to.not.exist

    result = await Validator(app)(
      "project", { description: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "project", { id: null, pronumber: "pronumber" }
    )
    expect(result.error).to.not.exist
    result = await Validator(app)(
      "project", { id: 1, pronumber: null }
    )
    expect(result.error).to.not.exist
    result = await Validator(app)(
      "project", { id: 1 }
    )
    expect(result.error).to.not.exist

    result = await Validator(app)(
      "rate", { ratetype: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "rate", { ratedate: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "rate", { curr: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "rate", {}
    )
    expect(result.error).to.not.exist

    result = await Validator(app)(
      "tax", { taxcode: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "tax", { description: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "tax", { id: null, taxcode: "taxcode" }
    )
    expect(result.error).to.not.exist
    result = await Validator(app)(
      "tax", { id: 1 }
    )
    expect(result.error).to.not.exist

    result = await Validator(app)(
      "tool", { product_id: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "tool", { description: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "tool", { id: null, serial: "serial" }
    )
    expect(result.error).to.not.exist
    result = await Validator(app)(
      "tool", { id: 1, serial: null }
    )
    expect(result.error).to.not.exist
    result = await Validator(app)(
      "tool", { id: 1 }
    )
    expect(result.error).to.not.exist

    result = await Validator(app)(
      "trans", { transtype: 55, direction: 68, customer_id: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "trans", { transtype: 67, direction: 68, place_id: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "trans", { transtype: 66, direction: 68, place_id: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "trans", { id: 1, transtype: 66, direction: 68, place_id: 1 }
    )
    expect(result.error).to.not.exist
    result = await Validator(app)(
      "trans", { transtype: 62, direction: 68, place_id: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "trans", { transtype: 64, direction: 68, place_id: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "trans", { transtype: 61, direction: 70, place_id: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "trans", { transtype: 64, direction: 68, duedate: null }
    )
    expect(result.error).to.exist
    testApp_ = testApp("current", { extend: { product_id: null } })
    result = await Validator(testApp_)(
      "trans", { transtype: 64, direction: 68 }
    )
    expect(result.error).to.exist
    result = await Validator(testApp_)(
      "trans", { transtype: 63, direction: 68 }
    )
    expect(result.error).to.not.exist
    testApp_ = testApp("current", { extend: { product_id: 1, ref_id: null } })
    result = await Validator(testApp_)(
      "trans", { transtype: 65, direction: 68 }
    )
    expect(result.error).to.not.exist
    result = await Validator(testApp_)(
      "trans", { transtype: 63, direction: 68 }
    )
    expect(result.error).to.exist
    result = await Validator(testApp_)(
      "trans", { transtype: 55, direction: 68, transnumber: null }
    )
    expect(result.error).to.not.exist
    result = await Validator(testApp_)(
      "trans", { transtype: 67, direction: 68, transnumber: null }
    )
    expect(result.error).to.not.exist

    result = await Validator(app)(
      "ui_menu", { menukey: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "ui_menu", { description: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "ui_menu", { method: null }
    )
    expect(result.error).to.exist
    result = await Validator(app)(
      "ui_menu", { id: null, menukey: "menukey" }
    )
    expect(result.error).to.not.exist
    result = await Validator(app)(
      "ui_menu", { id: 1 }
    )
    expect(result.error).to.not.exist

    result = await Validator(app)(
      "default", {}
    )
    expect(result.error).to.not.exist

    testApp_ = {
      ...app,
      requestData: sinon.spy(async (path) => {
        if(path === "/view"){
          return { check: [{ recnum: 1 }] }
        }
        return { error: {} }
      }),
    }

    result = await Validator(testApp_)(
      "customer", { id: null, custnumber: "custnumber" }
    )
    expect(result.error).to.exist
    result = await Validator(testApp_)(
      "customer", { id: null, custnumber: null }
    )
    expect(result.error).to.exist

    testApp_ = {
      ...app,
      requestData: sinon.spy(async () => ({ error: {} })),
    }
    result = await Validator(testApp_)(
      "customer", { id: null, custnumber: "custnumber" }
    )
    expect(result.error).to.exist

  })

})