//import { queryByAttribute } from '@testing-library/react'
//import ReactDOM from 'react-dom';
import update from 'immutability-helper';

import { InitItem, Validator } from './Validator'
import { appActions, getSql } from 'containers/App/actions'
//import { Forms } from 'containers/Controller/Forms'
//import { InitItem, Validator } from 'containers/Controller/Validator'
import { store as app_store  } from 'config/app'

import { Default as InvoiceData } from 'components/Editor/Editor/Editor.stories'

//const getById = queryByAttribute.bind(null, 'id');
jest.mock("containers/App/actions");

const store = update(app_store, {$merge: {
  login: {
    data: {
      employee: {
        id: 1, empnumber: 'admin', username: 'admin', usergroup: 114,
        usergroupName: 'admin'
      },
      groups: InvoiceData.args.dataset.groups
    }
  },
  edit: {
    ...InvoiceData.args
  }
}})

describe('Validator', () => {

  beforeEach(() => {
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view"){
          return { check: [{ recnum: 0 }] }
        }
        return "012345"
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
    })
    getSql.mockReturnValue({
      sql: "",
      prmCount: 1
    })
  })

  afterEach(() => {
    jest.clearAllMocks();
  });

  it('InitItem', () => {
    const setData = jest.fn()
    let values = InitItem(store, setData)({
      tablename: "address", dataset: store.edit.dataset, current: store.edit.current
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "audit"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "barcode"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "contact"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "currency"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "customer"
    })
    expect(values).toBeDefined()
    let it_store = update(store, {edit: {dataset: {$merge:{
      custtype: undefined,
    }}}})
    values = InitItem(it_store, setData)({
      tablename: "customer"
    })
    expect(values).toBeNull()

    values = InitItem(store, setData)({
      tablename: "deffield"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "employee"
    })
    expect(values).toBeDefined()
    it_store = update(store, {edit: {dataset: {$merge:{
      usergroup: undefined,
    }}}})
    values = InitItem(it_store, setData)({
      tablename: "employee"
    })
    expect(values).toBeNull()

    values = InitItem(store, setData)({
      tablename: "event"
    })
    expect(values).toBeDefined()
    it_store = update(store, {edit: {current: {$merge:{
      type: "event",
    }}}})
    values = InitItem(it_store, setData)({
      tablename: "event"
    })
    expect(values).toBeDefined()
    it_store = update(store, {edit: {current: {$merge:{
      item: undefined,
    }}}})
    values = InitItem(it_store, setData)({
      tablename: "event"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "fieldvalue"
    })
    expect(values).toBeDefined()
    it_store = update(store, {edit: {current: {$merge:{
      item: undefined,
    }}}})
    values = InitItem(it_store, setData)({
      tablename: "fieldvalue"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "groups"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "usergroup"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "item"
    })
    expect(values).toBeDefined()

    it_store = update(store, {edit: {current: {$merge:{
      form_type: "invoice_link",
    }}}})
    values = InitItem(it_store, setData)({
      tablename: "link"
    })
    expect(values).toBeDefined()
    it_store = update(store, {edit: {current: {$merge:{
      form_type: "payment_link",
    }}}})
    values = InitItem(it_store, setData)({
      tablename: "link"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "log"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "ui_menu"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "ui_menufields"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "movement"
    })
    expect(values).toBeDefined()
    it_store = update(store, {edit: {
      current: {$merge:{
        transtype: "delivery",
      }},
      dataset: {$merge:{
        movement_transfer: []
      }}
    }})
    values = InitItem(it_store, setData)({
      tablename: "movement"
    })
    expect(values).toBeDefined()
    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        movement_transfer: [
          { place_id: 1 }
        ]
      }}
    }})
    values = InitItem(it_store, setData)({
      tablename: "movement"
    })
    expect(values).toBeDefined()
    it_store = update(store, {edit: {
      current: {$merge:{
        transtype: "inventory",
      }}
    }})
    values = InitItem(it_store, setData)({
      tablename: "movement"
    })
    expect(values).toBeDefined()
    it_store = update(store, {edit: {
      current: {$merge:{
        transtype: "production",
      }}
    }})
    values = InitItem(it_store, setData)({
      tablename: "movement"
    })
    expect(values).toBeDefined()
    it_store = update(store, {edit: {
      current: {$merge:{
        transtype: "formula",
      }}
    }})
    values = InitItem(it_store, setData)({
      tablename: "movement"
    })
    expect(values).toBeDefined()
    it_store = update(store, {edit: {
      current: {$merge:{
        transtype: "waybill",
      }}
    }})
    values = InitItem(it_store, setData)({
      tablename: "movement"
    })
    expect(values).toBeDefined()

    it_store = update(store, {edit: {
      current: {$merge:{
        transtype: "formula",
      }}
    }})
    values = InitItem(it_store, setData)({
      tablename: "movement_head"
    })
    expect(values).toBeDefined()
    it_store = update(store, {edit: {
      current: {$merge:{
        transtype: "production",
      }}
    }})
    values = InitItem(it_store, setData)({
      tablename: "movement_head"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "numberdef"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "pattern"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "payment"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "place"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "discount"
    })
    expect(values).toBeDefined()
    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        settings: []
      }}
    }})
    values = InitItem(it_store, setData)({
      tablename: "price"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "product"
    })
    expect(values).toBeDefined()
    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        settings: []
      }}
    }})
    values = InitItem(it_store, setData)({
      tablename: "product"
    })
    expect(values).toBeDefined()
    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        protype: undefined
      }}
    }})
    values = InitItem(it_store, setData)({
      tablename: "product"
    })
    expect(values).toBeNull()

    values = InitItem(store, setData)({
      tablename: "project"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "printqueue"
    })
    expect(values).toBeDefined()
    it_store = update(store, {edit: {
      current: {$merge:{
        type: "printqueue",
      }}
    }})
    values = InitItem(it_store, setData)({
      tablename: "printqueue"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "rate"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "refvalue"
    })
    expect(values).toBeDefined()
    it_store = update(store, {edit: {
      current: {$merge:{
        transtype: "waybill",
      }}
    }})
    values = InitItem(it_store, setData)({
      tablename: "refvalue"
    })
    expect(values).toBeDefined()
    it_store = update(it_store, {edit: {
      dataset: {trans: {0: {$merge:{
        customer_id: null,
        employee_id: 1
      }}}}
    }})
    values = InitItem(it_store, setData)({
      tablename: "refvalue"
    })
    expect(values).toBeDefined()
    it_store = update(it_store, {edit: {
      dataset: {trans: {0: {$merge:{
        customer_id: null,
        employee_id: null
      }}}}
    }})
    values = InitItem(it_store, setData)({
      tablename: "refvalue"
    })
    expect(values).toBeDefined()
    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        translink: []
      }}
    }})
    values = InitItem(it_store, setData)({
      tablename: "refvalue"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "report"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "tax"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "tool"
    })
    expect(values).toBeDefined()

    values = InitItem(store, setData)({
      tablename: "trans"
    })
    expect(values).toBeDefined()
    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        settings: [],
        pattern: []
      }}
    }})
    values = InitItem(it_store, setData)({
      tablename: "trans", transtype: "invoice"
    })
    expect(values).toBeDefined()
    values = InitItem(it_store, setData)({
      tablename: "trans", transtype: "formula"
    })
    expect(values).toBeDefined()
    values = InitItem(it_store, setData)({
      tablename: "trans", transtype: "production"
    })
    expect(values).toBeDefined()
    values = InitItem(it_store, setData)({
      tablename: "trans", transtype: "cash"
    })
    expect(values).toBeDefined()
    values = InitItem(it_store, setData)({
      tablename: "trans", transtype: "offer"
    })
    expect(values).toBeDefined()
    values = InitItem(it_store, setData)({
      tablename: "trans", transtype: "order"
    })
    expect(values).toBeDefined()
    values = InitItem(it_store, setData)({
      tablename: "trans", transtype: "worksheet"
    })
    expect(values).toBeDefined()
    values = InitItem(it_store, setData)({
      tablename: "trans", transtype: "rent"
    })
    expect(values).toBeDefined()
    values = InitItem(it_store, setData)({
      tablename: "trans", transtype: "receipt"
    })
    expect(values).toBeDefined()
    values = InitItem(it_store, setData)({
      tablename: "trans", transtype: "bank"
    })
    expect(values).toBeDefined()
    values = InitItem(it_store, setData)({
      tablename: "trans", transtype: "inventory"
    })
    expect(values).toBeDefined()
    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        pattern: undefined
      }}
    }})
    values = InitItem(it_store, setData)({
      tablename: "trans"
    })
    expect(values).toBeNull()

    values = InitItem(it_store, setData)({
      tablename: "missing"
    })
    expect(values).toBeFalsy()

  })

  it('Validator', async () => {
    const setData = jest.fn()
    let values = await Validator(store, setData)(
      "address", {}
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "barcode", { code: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "barcode", { barcodetype: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "barcode", { id: null, code: "code" }
    )
    expect(values.error).toBeUndefined()
    values = await Validator(store, setData)(
      "barcode", { id: 1 }
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "contact", {}
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "currency", { curr: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "currency", { description: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "currency", { id: null, curr: "curr" }
    )
    expect(values.error).toBeUndefined()
    values = await Validator(store, setData)(
      "currency", { id: 1 }
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "customer", { custname: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "customer", { custtype: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "customer", { id: null, custnumber: "custnumber" }
    )
    expect(values.error).toBeUndefined()
    values = await Validator(store, setData)(
      "customer", { id: null, custnumber: null }
    )
    expect(values.error).toBeUndefined()
    values = await Validator(store, setData)(
      "customer", { id: 1 }
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "deffield", { nervatype: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "deffield", { fieldtype: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "deffield", { description: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "deffield", { id: 1 }
    )
    expect(values.error).toBeUndefined()

    let it_store = update(store, {edit: {
      current: {$merge:{
        extend: {
          surname: null
        },
      }}
    }})
    values = await Validator(it_store, setData)(
      "employee", {}
    )
    expect(values.error).toBeDefined()
    it_store = update(store, {edit: {
      current: {$merge:{
        extend: {
          surname: "surname"
        },
      }}
    }})
    values = await Validator(it_store, setData)(
      "employee", { usergroup: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(it_store, setData)(
      "employee", { id: null, empnumber: "empnumber" }
    )
    expect(values.error).toBeUndefined()
    values = await Validator(it_store, setData)(
      "employee", { id: null, empnumber: null }
    )
    expect(values.error).toBeUndefined()
    values = await Validator(it_store, setData)(
      "employee", { id: 1 }
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "event", { calnumber: null }
    )
    expect(values.error).toBeUndefined()
    values = await Validator(store, setData)(
      "event", { id: 1 }
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "fieldvalue", {}
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "formula", {}
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "groups", { groupvalue: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "groups", { groupname: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "groups", { groupname: "usergroup", description: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "groups", { groupname: "usergroup", description: "description", transfilter: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "groups", { id: null }
    )
    expect(values.error).toBeUndefined()
    values = await Validator(store, setData)(
      "groups", { id: 1 }
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "item", { product_id: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "item", { description: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "item", { tax_id: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "item", { id: 1 }
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "link", { ref_id_1: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "link", { id: 1 }
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "log", {}
    )
    expect(values.error).toBeUndefined()

    it_store = update(store, {edit: {
      current: {$merge:{
        transtype: "delivery",
        item: {
          direction: 68
        }
      }}
    }})
    values = await Validator(it_store, setData)(
      "movement", { place_id: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(it_store, setData)(
      "movement", { product_id: null }
    )
    expect(values.error).toBeDefined()
    it_store = update(store, {edit: {
      current: {$merge:{
        transtype: "delivery",
        item: {
          direction: 70, place_id: 1
        }
      }}
    }})
    values = await Validator(it_store, setData)(
      "movement", { place_id: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(it_store, setData)(
      "movement", { place_id: 1 }
    )
    expect(values.error).toBeDefined()
    values = await Validator(it_store, setData)(
      "movement", {}
    )
    expect(values.error).toBeUndefined()

    it_store = update(store, {edit: {
      current: {$merge:{
        transtype: "inventory",
      }}
    }})
    values = await Validator(it_store, setData)(
      "movement", { product_id: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(it_store, setData)(
      "movement", {}
    )
    expect(values.error).toBeUndefined()

    it_store = update(store, {edit: {
      current: {$merge:{
        transtype: "production",
      }}
    }})
    values = await Validator(it_store, setData)(
      "movement", { product_id: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(it_store, setData)(
      "movement", { place_id: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(it_store, setData)(
      "movement", {}
    )
    expect(values.error).toBeUndefined()

    it_store = update(store, {edit: {
      current: {$merge:{
        transtype: "formula",
      }}
    }})
    values = await Validator(it_store, setData)(
      "movement", { product_id: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(it_store, setData)(
      "movement", {}
    )
    expect(values.error).toBeUndefined()

    it_store = update(store, {edit: {
      current: {$merge:{
        transtype: "waybill",
      }}
    }})
    values = await Validator(it_store, setData)(
      "movement", { tool_id: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(it_store, setData)(
      "movement", {}
    )
    expect(values.error).toBeUndefined()

    it_store = update(store, {edit: {
      current: {$merge:{
        transtype: "invoice",
      }}
    }})
    values = await Validator(it_store, setData)(
      "movement", {}
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "numberdef", {}
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "pattern", {}
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "payment", {}
    )
    expect(values.error).toBeUndefined()

    values = await Validator(it_store, setData)(
      "place", { placetype: 126 }
    )
    expect(values.error).toBeUndefined()
    values = await Validator(it_store, setData)(
      "place", { id: 1, description: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(it_store, setData)(
      "place", { id: 1, placetype: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(it_store, setData)(
      "place", { id: 1, planumber: null }
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "price", { validfrom: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "price", { curr: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "price", { calcmode: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "price", {}
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "product", { description: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "product", { protype: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "product", { unit: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "product", { tax: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "product", { id: null, partnumber: "partnumber" }
    )
    expect(values.error).toBeUndefined()
    values = await Validator(store, setData)(
      "product", { id: 1, partnumber: null }
    )
    expect(values.error).toBeUndefined()
    values = await Validator(store, setData)(
      "product", { id: 1 }
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "project", { description: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "project", { id: null, pronumber: "pronumber" }
    )
    expect(values.error).toBeUndefined()
    values = await Validator(store, setData)(
      "project", { id: 1, pronumber: null }
    )
    expect(values.error).toBeUndefined()
    values = await Validator(store, setData)(
      "project", { id: 1 }
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "rate", { ratetype: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "rate", { ratedate: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "rate", { curr: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "rate", {}
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "tax", { taxcode: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "tax", { description: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "tax", { id: null, taxcode: "taxcode" }
    )
    expect(values.error).toBeUndefined()
    values = await Validator(store, setData)(
      "tax", { id: 1 }
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "tool", { product_id: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "tool", { description: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "tool", { id: null, serial: "serial" }
    )
    expect(values.error).toBeUndefined()
    values = await Validator(store, setData)(
      "tool", { id: 1, serial: null }
    )
    expect(values.error).toBeUndefined()
    values = await Validator(store, setData)(
      "tool", { id: 1 }
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "trans", { transtype: 55, direction: 68, customer_id: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "trans", { transtype: 67, direction: 68, place_id: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "trans", { transtype: 66, direction: 68, place_id: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "trans", { id: 1, transtype: 66, direction: 68, place_id: 1 }
    )
    expect(values.error).toBeUndefined()
    values = await Validator(store, setData)(
      "trans", { transtype: 62, direction: 68, place_id: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "trans", { transtype: 64, direction: 68, place_id: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "trans", { transtype: 61, direction: 70, place_id: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "trans", { transtype: 64, direction: 68, duedate: null }
    )
    expect(values.error).toBeDefined()
    it_store = update(store, {edit: {
      current: {$merge:{
        extend: {
          product_id: null
        },
      }}
    }})
    values = await Validator(it_store, setData)(
      "trans", { transtype: 64, direction: 68 }
    )
    expect(values.error).toBeDefined()
    values = await Validator(it_store, setData)(
      "trans", { transtype: 63, direction: 68 }
    )
    expect(values.error).toBeUndefined()
    it_store = update(store, {edit: {
      current: {$merge:{
        extend: {
          product_id: 1, ref_id: null
        },
      }}
    }})
    values = await Validator(it_store, setData)(
      "trans", { transtype: 65, direction: 68 }
    )
    expect(values.error).toBeUndefined()
    values = await Validator(it_store, setData)(
      "trans", { transtype: 63, direction: 68 }
    )
    expect(values.error).toBeDefined()
    values = await Validator(it_store, setData)(
      "trans", { transtype: 55, direction: 68, transnumber: null }
    )
    expect(values.error).toBeUndefined()
    values = await Validator(it_store, setData)(
      "trans", { transtype: 67, direction: 68, transnumber: null }
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "ui_menu", { menukey: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "ui_menu", { description: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "ui_menu", { method: null }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "ui_menu", { id: null, menukey: "menukey" }
    )
    expect(values.error).toBeUndefined()
    values = await Validator(store, setData)(
      "ui_menu", { id: 1 }
    )
    expect(values.error).toBeUndefined()

    values = await Validator(store, setData)(
      "default", {}
    )
    expect(values.error).toBeUndefined()

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view"){
          return { check: [{ recnum: 1 }] }
        }
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
    })
    values = await Validator(store, setData)(
      "customer", { id: null, custnumber: "custnumber" }
    )
    expect(values.error).toBeDefined()
    values = await Validator(store, setData)(
      "customer", { id: null, custnumber: null }
    )
    expect(values.error).toBeDefined()

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
    })
    values = await Validator(store, setData)(
      "customer", { id: null, custnumber: "custnumber" }
    )
    expect(values.error).toBeDefined()

  })

})