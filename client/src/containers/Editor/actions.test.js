import { queryByAttribute, fireEvent } from '@testing-library/react'
import ReactDOM from 'react-dom';
import update from 'immutability-helper';
import printJS from 'print-js'

import { editorActions } from './actions'
import { appActions, saveToDisk, getSql } from 'containers/App/actions'
import { Forms } from 'containers/Controller/Forms'
import { InitItem, Validator } from 'containers/Controller/Validator'
import { store as app_store, getSetting  } from 'config/app'

import { Default as InvoiceData, PrintQueue, Report, Item } from 'components/Editor/Editor/Editor.stories'

jest.mock("containers/App/actions");
jest.mock("containers/Controller/Validator");
jest.mock("print-js");

const getById = queryByAttribute.bind(null, 'id');

const store = update(app_store, {$merge: {
  login: {
    data: {
      employee: {
        id: 1, empnumber: 'admin', username: 'admin', usergroup: 114,
        usergroupName: 'admin'
      },
      groups: [
        { id: 1, groupname: "nervatype", groupvalue: "type" }
      ],
      audit: [
        { nervatypeName: "report", subtype: 1, inputfilterName: "disabled" }
      ]
    }
  }
}})

describe('editorActions', () => {

  beforeEach(() => {
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ([1])),
      resultError: jest.fn(),
      showToast: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    getSql.mockReturnValue({
      sql: "",
      prmCount: 1
    })
    InitItem.mockReturnValue(jest.fn( () => ({}) ))
    Validator.mockReturnValue(jest.fn( () => ({}) ))
    saveToDisk.mockReturnValue()
    printJS.mockReturnValue()
    global.URL.createObjectURL = jest.fn();
  });

  afterEach(() => {
    jest.clearAllMocks();
  });

  it('round', () => {
    const setData = jest.fn()
    let value = editorActions(store, setData).round(2.333, 1)
    expect(value).toBe(2.3)
    value = editorActions(store, setData).round(2.666, 2)
    expect(value).toBe(2.67)
    value = editorActions(store, setData).round(2.333)
    expect(value).toBe(2)
    value = editorActions(store, setData).round("abc")
    expect(value).toBeNaN()
  })

  it('reportPath', () => {
    const setData = jest.fn()
    let it_store = update(store, {edit: {$merge:{
      current: {
        type: "nervatype",
        item: {
          id: 2
        }
      }
    }}})
    let value = editorActions(it_store, setData).reportPath({ 
      template: "template", orient: "orient", size: "size", type: "type", nervatype: "nervatype", id: 1
    })
    expect(value).toBe("/report?reportkey=template&orientation=orient&size=size&output=type&nervatype=nervatype&filters[@id]=1")
    value = editorActions(it_store, setData).reportPath({ 
      template: "template", orient: "orient", size: "size", type: "type"
    })
    expect(value).toBe("/report?reportkey=template&orientation=orient&size=size&output=type&nervatype=nervatype&filters[@id]=2")
    value = editorActions(it_store, setData).reportPath({ 
      template: "template", orient: "orient", size: "size", type: "type", nervatype: "nervatype", filters: "abc=123"
    })
    expect(value).toBe("/report?reportkey=template&orientation=orient&size=size&output=type&abc=123")
  })

  it('addPrintQueue', () => {
    const setData = jest.fn()
    let it_store = update(store, {edit: {$merge:{
      current: {
        type: "type",
        item: {
          id: 1
        }
      },
      dataset: {
        report: [
          { id: 1, reportkey: "reportkey" }
        ]
      }
    }}})
    editorActions(it_store, setData).addPrintQueue("reportkey", 1)
    expect(setData).toHaveBeenCalledTimes(0)

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ({ error: {} })),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    editorActions(it_store, setData).addPrintQueue("reportkey", 1)
    expect(setData).toHaveBeenCalledTimes(0)
  })

  it('reportOutput', () => {
    const setData = jest.fn()
    let it_store = update(store, {edit: {$merge:{
      current: {
        type: "type",
        item: {
          id: 1
        }
      },
      dataset: {
        report: [
          { id: 1, reportkey: "reportkey" }
        ]
      }
    }}})
    editorActions(it_store, setData).reportOutput(
      { type: "printqueue", template: "reportkey", copy: 1 })
    expect(setData).toHaveBeenCalledTimes(0)

    editorActions(it_store, setData).reportOutput(
      { type: "xml", template: "reportkey", copy: 1, title: "title" })
    expect(setData).toHaveBeenCalledTimes(0)

    editorActions(it_store, setData).reportOutput(
      { type: "print", template: "reportkey" })
    expect(setData).toHaveBeenCalledTimes(0)

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ({ error: {} })),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    editorActions(it_store, setData).reportOutput(
      { type: "print", template: "reportkey" })
    expect(setData).toHaveBeenCalledTimes(0)

  })

  it('searchQueue', async () => {
    const setData = jest.fn()
    let it_store = update(store, {edit: {$merge:{
      printqueue: {}
    }}})
    await editorActions(it_store, setData).searchQueue()
    expect(setData).toHaveBeenCalledTimes(1)

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ({ error: {} })),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    await editorActions(it_store, setData).searchQueue()
    expect(setData).toHaveBeenCalledTimes(1)
  })

  it('exportQueueAll', () => {
    let setData = jest.fn((key, data, callback)=>{
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(data.modalForm, container);

        // onOK
        const btn_ok = getById(container, 'btn_ok')
        btn_ok.dispatchEvent(new MouseEvent('click', {bubbles: true}));

        // onCancel
        const btn_cancel = getById(container, 'btn_cancel')
        btn_cancel.dispatchEvent(new MouseEvent('click', {bubbles: true}));

      }
      if(callback){callback()}
    })
    let it_store = update(store, {edit: {$merge:{
      printqueue: {},
      current: {
        type: "type",
        item: {
          id: 1
        }
      },
      dataset: {
        items: []
      }
    }}})
    editorActions(it_store, setData).exportQueueAll()
    expect(setData).toHaveBeenCalledTimes(0)

    it_store = update(store, {edit: {$merge:{
      printqueue: {},
      current: {
        type: "type",
        item: {
          id: 1, mode: "print"
        }
      },
      dataset: {
        items: [
          { id: 1 }
        ]
      }
    }}})
    editorActions(it_store, setData).exportQueueAll()
    expect(setData).toHaveBeenCalledTimes(0)

    it_store = update(store, {edit: {$merge:{
      printqueue: {},
      current: {
        type: "type",
        item: {
          id: 1
        }
      },
      dataset: {
        items: [
          { id: 1 }
        ]
      }
    }}})
    editorActions(it_store, setData).exportQueueAll()
    expect(setData).toHaveBeenCalledTimes(3)

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path) => {
        if(String(path).startsWith("/report")){
          return { error: {} }
        }
        return {}
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    editorActions(it_store, setData).exportQueueAll()
    expect(setData).toHaveBeenCalledTimes(6)

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path) => {
        if(String(path).startsWith("/ui_printqueue")){
          return { error: {} }
        }
        return {}
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    editorActions(it_store, setData).exportQueueAll()
    expect(setData).toHaveBeenCalledTimes(9)
  
  })

  it('createReport', () => {
    const setData = jest.fn()
    let it_store = update(store, {edit: {$merge:{
      current: {
        item: {},
        fieldvalue: [
          { datatype: "date", empty: "false", id: 1, label: "Date",
            name: "posdate", rowtype: "reportfield", selected: true, value: "2022-03-14" 
          }
        ]
      }
    }}})
    editorActions(it_store, setData).createReport("xml")
    expect(setData).toHaveBeenCalledTimes(0)

    it_store = update(store, {edit: {$merge:{
      current: {
        item: {},
        fieldvalue: [
          { datatype: "date", empty: "false", id: 1, label: "Date",
            name: "posdate", rowtype: "reportfield", selected: false, value: "2022-03-14" 
          }
        ]
      }
    }}})
    editorActions(it_store, setData).createReport("csv")
    expect(setData).toHaveBeenCalledTimes(0)

    editorActions(it_store, setData).createReport("print")
    expect(setData).toHaveBeenCalledTimes(0)

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ({ error: {} })),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    editorActions(it_store, setData).createReport("print")
    expect(setData).toHaveBeenCalledTimes(0)

  })
  
  it('reportSettings', () => { 
    let setData = jest.fn((key, data, callback)=>{
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(data.modalForm, container);

        // onOK
        const btn_print = getById(container, 'btn_print')
        btn_print.dispatchEvent(new MouseEvent('click', {bubbles: true}));

        // onCancel
        const btn_cancel = getById(container, 'closeIcon')
        btn_cancel.dispatchEvent(new MouseEvent('click', {bubbles: true}));

      }
      if(callback){callback()}
    })
    let it_store = update(store, {edit: {$merge:{
      current: {
        type: "trans",
        transtype: "invoice",
        item: {
          direction: 1, transnumber: "transnumber"
        }
      },
      template: {
        options: {
          title_field: "transnumber"
        }
      },
      dataset: {
        groups: [
          { id: 1, groupname: "direction", groupvalue: "out" }
        ],
        settings: [
          { fieldname: "default_trans_invoice_out_report", value: "value" }
        ],
        report: [
          { id: 1, direction: 1, reportkey: "reportkey", repname: "repname" }
        ]
      }
    }}})
    editorActions(it_store, setData).reportSettings()
    expect(setData).toHaveBeenCalledTimes(3)

    it_store = update(it_store, {login: {data: {$merge:{
      audit: []
    }}}})
    editorActions(it_store, setData).reportSettings()
    expect(setData).toHaveBeenCalledTimes(6)

    it_store = update(it_store, {edit: {$merge:{
      dataset: {
        groups: [
          { id: 1, groupname: "direction", groupvalue: "out" }
        ],
        settings: [
          { fieldname: "default_trans_invoice_out_report", value: "value" }
        ],
        report: [
          { id: 1, direction: 2, reportkey: "reportkey", repname: "repname" }
        ]
      }
    }}})
    editorActions(it_store, setData).reportSettings()
    expect(setData).toHaveBeenCalledTimes(9)

    it_store = update(it_store, {edit: {$merge:{
      current: {
        type: "customer",
        item: {
          custnumber: "custnumber"
        }
      },
      template: {
        options: {
          title_field: "custnumber"
        }
      },
    }}})
    editorActions(it_store, setData).reportSettings()
    expect(setData).toHaveBeenCalledTimes(12)
  
    it_store = update(it_store, {login: {data: {$merge:{
      audit: [
        { nervatypeName: "report", subtype: 1, inputfilterName: "disabled" }
      ]
    }}}})
    editorActions(it_store, setData).reportSettings()
    expect(setData).toHaveBeenCalledTimes(14)
    
  })

  it('setEditor', () => {
    const setData = jest.fn()

    editorActions(store, setData).setEditor(
      {}, {}, { current: { type: "type" }, dataset: {} } )
    expect(setData).toHaveBeenCalledTimes(0)

    const forms = Forms({ getText: jest.fn() })
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    let options = {
      ntype: "trans", ttype: "invoice", id: 5,
      item: {
        custname: "First Customer Co.", deleted: 0, id: "trans/invoice/5",
        label: "DMINV/00001", transnumber: "DMINV/00001", transtype: "invoice-out",
      },
    }
    let form = forms["invoice"](it_store.edit.current.item, it_store.edit)
    editorActions(it_store, setData).setEditor(options, form)
    expect(setData).toHaveBeenCalledTimes(2)

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ([1])),
      resultError: jest.fn(),
      showToast: jest.fn(),
      getAuditFilter: jest.fn(() => (["disabled",1])),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    editorActions(it_store, setData).setEditor(options, form)
    expect(setData).toHaveBeenCalledTimes(2)

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ([1])),
      resultError: jest.fn(),
      showToast: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    it_store = update(it_store, {edit: {dataset: {$merge:{
      trans: [
        update(InvoiceData.args.current.item, {$merge: {
          id: null, transcast: "cancellation"
        }})
      ],
      pattern: undefined,
      fieldvalue: undefined
    }}}})
    editorActions(it_store, setData).setEditor(options, form )
    expect(setData).toHaveBeenCalledTimes(4)

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ([1])),
      resultError: jest.fn(),
      showToast: jest.fn(),
      getAuditFilter: jest.fn(() => (["readonly",1])),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    it_store = update(it_store, {edit: {dataset: {$merge:{
      trans: [
        update(InvoiceData.args.current.item, {$merge: {
          id: null, transcast: undefined, closed: 1
        }})
      ],
      pattern: []
    }}}})
    form = update(form,{options: {$merge: {
      extend: "item"
    }}})
    editorActions(it_store, setData).setEditor(options, form )
    expect(setData).toHaveBeenCalledTimes(6)

    it_store = update(it_store, {edit: {dataset: {$merge:{
      trans: [
        update(InvoiceData.args.current.item, {$merge: {
          id: 5, deleted: 1
        }})
      ],
      pattern: []
    }}}})
    form = update(form,{options: {$merge: {
      extend: "item"
    }}})
    editorActions(it_store, setData).setEditor(options, form )
    expect(setData).toHaveBeenCalledTimes(8)
    form = update(form,{options: {$merge: {
      extend: "cancel_link"
    }}})
    editorActions(it_store, setData).setEditor(options, form )
    expect(setData).toHaveBeenCalledTimes(10)
    form = update(form,{options: {$merge: {
      extend: "extend"
    }}})
    editorActions(it_store, setData).setEditor(options, form )
    expect(setData).toHaveBeenCalledTimes(12)

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ([1])),
      resultError: jest.fn(),
      showToast: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    it_store = update(store, {edit: {$merge:{
      ...PrintQueue.args
    }}})
    options = {
      ntype: "printqueue", ttype: null, id: null
    }
    form = forms["printqueue"]({}, {}, getSetting("ui"))
    editorActions(it_store, setData).setEditor(options, form )
    expect(setData).toHaveBeenCalledTimes(14)
    it_store = update(it_store, {edit: {$merge:{
      printqueue: it_store.edit.current.item
    }}})
    editorActions(it_store, setData).setEditor(options, form )
    expect(setData).toHaveBeenCalledTimes(16)

    it_store = update(store, {edit: {$merge:{
      ...Report.args
    }}})
    it_store = update(it_store, {edit: {dataset: {report: {0: {$merge:{
      report: JSON.stringify({
        fields: {
          date_1: {
            selected: true, wheretype: "in", fieldtype: "date", description: "description", value: "2022-01-16"
          },
          date_2: {
            selected: false, wheretype: "in", fieldtype: "date", description: "description"
          },
          date_3: {
            selected: false, wheretype: "where", fieldtype: "date", description: "description"
          },
          date_4: {
            selected: false, wheretype: "where", fieldtype: "date", description: "description", defvalue: 3
          },
          bool: {
            selected: false, wheretype: "where", fieldtype: "bool", description: "description", value: 0
          },
          float: {
            selected: false, wheretype: "where", fieldtype: "float", description: "description", defvalue: 20
          },
          integer_1: {
            selected: false, wheretype: "where", fieldtype: "integer", description: "description"
          },
          integer_2: {
            selected: false, wheretype: "where", fieldtype: "integer", description: "description", value: 100
          },
          valuelist_1: {
            selected: false, wheretype: "where", fieldtype: "valuelist", description: "description", valuelist: "a|b|c"
          },
          valuelist_2: {
            selected: true, wheretype: "in", fieldtype: "valuelist", description: "description", valuelist: "a|b|c", defvalue: "a"
          },
          string: {
            selected: false, wheretype: "where", fieldtype: "string", description: "description", value: ""
          },
        }
      })
    }}}}}})
    options = {
      ntype: "report", ttype: "", id: 7,
      item: it_store.edit.current.item,
    }
    form = forms["report"](it_store.edit.current.item, {}, getSetting("ui"))
    editorActions(it_store, setData).setEditor(options, form )
    expect(setData).toHaveBeenCalledTimes(18)

    it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    it_store = update(it_store, {edit: {dataset: {$merge:{
      shipping_items: [
        { description: "Wheel", item_id: 1, partname: "Wheel", partnumber: "DMPROD/00005",
          pgroup: "0", product: "DMPROD/00005 | Wheel", product_id: "5", qty: "40", unit: "piece" },
        { description: "Door", item_id: 2, partname: "Door", partnumber: "DMPROD/00006",
          pgroup: "0", product: "DMPROD/00006 | Door", product_id: "6", qty: "60", unit: "piece" },
        { description: "Paint", item_id: 3, partname: "Paint", partnumber: "DMPROD/00007",
          pgroup: "0", product: "DMPROD/00007 | Paint", product_id: "7", qty: "50", unit: "liter" }
      ],
      transitem_shipping: [
        { id: "1-5", item_product: "Wheel", movement_product: "Wheel", product_id: 5, sqty: "30" },
        { id: "2-6", item_product: "Door", movement_product: "Door", product_id: 6, sqty: "50" }
      ]
    }}}})
    options = {
      ntype: "trans", ttype: "order", id: 1, shipping: true 
    }
    form = update(forms["invoice"](it_store.edit.current.item, it_store.edit), {
      view: {$merge: {
        setting: {}
      }}
    })
    editorActions(it_store, setData).setEditor(options, form )
    expect(setData).toHaveBeenCalledTimes(20)

    it_store = update(it_store, {edit: {
      current: {$merge: {
        shippingdate: "2021.21.21",
        shipping_place_id: 123
      }},
      dataset: {$merge:{
        trans: [
          update(InvoiceData.args.current.item, {$merge: {
            direction: 69
          }})
        ],
        shiptemp: [
          { id: "1-5", item_product: "Wheel", movement_product: "Wheel", product_id: 5, sqty: "30" }
        ]
      }}
    }})
    editorActions(it_store, setData).setEditor(options, form )
    expect(setData).toHaveBeenCalledTimes(22)

  })

  it('loadEditor', async () => {
    let setData = jest.fn((key, data, callback)=>{
      if(callback){callback()}
    })
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    it_store = update(it_store, {edit: {dataset: {$merge:{
      shipping_items: [
        { description: "Wheel", item_id: 1, partname: "Wheel", partnumber: "DMPROD/00005",
          pgroup: "0", product: "DMPROD/00005 | Wheel", product_id: "5", qty: "40", unit: "piece" },
        { description: "Door", item_id: 2, partname: "Door", partnumber: "DMPROD/00006",
          pgroup: "0", product: "DMPROD/00006 | Door", product_id: "6", qty: "60", unit: "piece" },
        { description: "Paint", item_id: 3, partname: "Paint", partnumber: "DMPROD/00007",
          pgroup: "0", product: "DMPROD/00007 | Paint", product_id: "7", qty: "50", unit: "liter" }
      ],
      transitem_shipping: [
        { id: "1-5", item_product: "Wheel", movement_product: "Wheel", product_id: 5, sqty: "30" },
        { id: "2-6", item_product: "Door", movement_product: "Door", product_id: 6, sqty: "50" }
      ]
    }}}})
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => (it_store.edit.dataset)),
      resultError: jest.fn(),
      showToast: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    InitItem.mockImplementation((data, setData)=>{
      return (params) => {
        if (Object.keys(params.dataset).length === 0){
          return null
        }
        return update(it_store.edit.current.item, {$merge: {
          id: null
        }})
      }
    })
    let params = {
      ntype: "trans", ttype: "invoice", id: 5
    }
    await editorActions(it_store, setData).loadEditor(params)
    expect(setData).toHaveBeenCalledTimes(3)

    params = {
      ntype: "trans", ttype: "invoice", id: null, shipping: true
    }
    getSql.mockImplementation((engine, _sql)=>{
      return{ sql: "", prmCount: (_sql.from === "fieldvalue") ? 1 : 0 }})
    await editorActions(it_store, setData).loadEditor(params)
    expect(setData).toHaveBeenCalledTimes(6)

    params = {
      ntype: "trans", ttype: "formula", id: 18, cb_key: "LOAD_EDITOR"
    }
    await editorActions(it_store, setData).loadEditor(params)
    expect(setData).toHaveBeenCalledTimes(7)

    params = {
      ntype: "trans", ttype: "delivery", id: null
    }
    await editorActions(it_store, setData).loadEditor(params)
    expect(setData).toHaveBeenCalledTimes(10)

    it_store = update(store, {edit: {$merge:{
      ...PrintQueue.args
    }}})
    params = {
      ntype: "printqueue", ttype: null, id: null
    }
    InitItem.mockImplementation((data, setData)=>{
      return (params) => {
        return update(it_store.edit.current.item, {$merge: {
          id: null
        }})
      }
    })
    await editorActions(it_store, setData).loadEditor(params)
    expect(setData).toHaveBeenCalledTimes(13)

    it_store = update(store, {edit: {$merge:{
      ...Report.args
    }}})
    params = {
      ntype: "report", ttype: null, id: 7
    }
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ({ error: {} })),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    await editorActions(it_store, setData).loadEditor(params)
    expect(setData).toHaveBeenCalledTimes(13)

  })

  it('setEditorItem', () => {
    let setData = jest.fn()
    let it_store = update(store, {edit: {$merge:{
      ...Item.args
    }}})
    it_store = update(it_store, {edit: {$merge:{
      panel: it_store.edit.template.options.panel
    }}})
    let options = {
      fkey: "item", id: 18
    }
    editorActions(it_store, setData).setEditorItem(options)
    expect(setData).toHaveBeenCalledTimes(1)

    it_store = update(it_store, {edit: {panel: {$merge:{
      state: "normal"
    }}}})
    options = {
      fkey: "item", id: null
    }
    editorActions(it_store, setData).setEditorItem(options)
    expect(setData).toHaveBeenCalledTimes(2)

    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        invoice_link: [
          { id: 5 }
        ],
        payment_link: [
          { id: 5 }
        ],
        price: [
          { id: 5 }
        ]
      }}
    }})
    options = {
      fkey: "invoice_link", id: 5
    }
    editorActions(it_store, setData).setEditorItem(options)
    expect(setData).toHaveBeenCalledTimes(3)
    options = {
      fkey: "invoice_link", id: null
    }
    editorActions(it_store, setData).setEditorItem(options)
    expect(setData).toHaveBeenCalledTimes(4)

    options = {
      fkey: "payment_link", id: 5
    }
    editorActions(it_store, setData).setEditorItem(options)
    expect(setData).toHaveBeenCalledTimes(5)

    options = {
      fkey: "payment_link", id: null, link_field: "link_field"
    }
    editorActions(it_store, setData).setEditorItem(options)
    expect(setData).toHaveBeenCalledTimes(6)

    options = {
      fkey: "price", id: 5
    }
    editorActions(it_store, setData).setEditorItem(options)
    expect(setData).toHaveBeenCalledTimes(7)
    options = {
      fkey: "price", id: null
    }
    editorActions(it_store, setData).setEditorItem(options)
    expect(setData).toHaveBeenCalledTimes(8)

  })

  it('newFieldvalue', () => {
    let setData = jest.fn()
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    editorActions(it_store, setData).newFieldvalue("")
    expect(setData).toHaveBeenCalledTimes(0)

    it_store = update(it_store, {login: {data: {$merge:{
      groups: it_store.edit.dataset.groups
    }}}})
    // valuelist
    editorActions(it_store, setData).newFieldvalue("trans_transcast")
    
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ({ error: {} })),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    // transitem
    editorActions(it_store, setData).newFieldvalue("trans_transitem_link")
    // float
    editorActions(it_store, setData).newFieldvalue("link_qty")
    // date
    editorActions(it_store, setData).newFieldvalue("sample_customer_date")
    // string
    editorActions(it_store, setData).newFieldvalue("trans_custinvoice_custname")
    // time
    editorActions(it_store, setData).newFieldvalue("sample_time")
    // bool
    editorActions(it_store, setData).newFieldvalue("sample_bool")

  })

  it('deleteEditor', () => {
    const setData = jest.fn((key, data, callback)=>{
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(data.modalForm, container);

        // onOK
        const btn_ok = getById(container, 'btn_ok')
        btn_ok.dispatchEvent(new MouseEvent('click', {bubbles: true}));

        // onCancel
        const btn_cancel = getById(container, 'btn_cancel')
        btn_cancel.dispatchEvent(new MouseEvent('click', {bubbles: true}));
      }
      if(callback){callback()}
    })
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ({ state: [{ sco: 0 }] })),
      resultError: jest.fn(),
      showToast: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    editorActions(it_store, setData).deleteEditor()
    expect(setData).toHaveBeenCalledTimes(3);

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(options.method === "DELETE"){
          return { error: {} }
        }
        return { state: [{ sco: 0 }] }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    editorActions(it_store, setData).deleteEditor()
    expect(setData).toHaveBeenCalledTimes(6);

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { state: [{ sco: 1 }] }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    editorActions(it_store, setData).deleteEditor()
    expect(setData).toHaveBeenCalledTimes(9);

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    editorActions(it_store, setData).deleteEditor()
    expect(setData).toHaveBeenCalledTimes(12);

    it_store = update(it_store, {edit: {current: {$merge:{
      type: "event"
    }}}})
    editorActions(it_store, setData).deleteEditor()
    expect(setData).toHaveBeenCalledTimes(15);

    it_store = update(it_store, {edit: {current: {item: {$merge:{
      id: null
    }}}}})
    editorActions(it_store, setData).deleteEditor()
    expect(setData).toHaveBeenCalledTimes(20);

  })

  it('deleteEditorItem', () => {
    const setData = jest.fn((key, data, callback)=>{
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(data.modalForm, container);

        // onOK
        const btn_ok = getById(container, 'btn_ok')
        btn_ok.dispatchEvent(new MouseEvent('click', {bubbles: true}));

        // onCancel
        const btn_cancel = getById(container, 'btn_cancel')
        btn_cancel.dispatchEvent(new MouseEvent('click', {bubbles: true}));
      }
      if(callback){callback()}
    })
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => (it_store.edit.dataset)),
      resultError: jest.fn(),
      showToast: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    let params = { fkey: "item", id: 18 }
    editorActions(it_store, setData).deleteEditorItem(params)
    expect(setData).toHaveBeenCalledTimes(3);

    params = { fkey: "item", id: 18, prompt: true, table: "item" }
    editorActions(it_store, setData).deleteEditorItem(params)
    expect(setData).toHaveBeenCalledTimes(3);

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    params = { fkey: "item", id: 18, prompt: true }
    editorActions(it_store, setData).deleteEditorItem(params)
    expect(setData).toHaveBeenCalledTimes(3);

    params = { fkey: "item", id: null, prompt: true }
    editorActions(it_store, setData).deleteEditorItem(params)
    expect(setData).toHaveBeenCalledTimes(3);

  })

  it('setFieldvalue', () => {
    const setData = jest.fn()
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    let values = editorActions(it_store, setData).setFieldvalue(it_store.edit.current.fieldvalue, 
      "trans_transcast", it_store.edit.current.item.id, null, "edit")
    expect(values[1].value).toBe("edit")

    values = editorActions(it_store, setData).setFieldvalue(it_store.edit.current.fieldvalue, 
      "trans_transcast", it_store.edit.current.item.id, null, null)
    expect(values.length).toBe(it_store.edit.current.fieldvalue.length)

    InitItem.mockImplementation((data, setData)=>{
      return (params) => {
        return {
          id: null, fieldname: null, ref_id: null, value: null, notes: null, deleted: 0
        }
      }
    })
    values = editorActions(it_store, setData).setFieldvalue(it_store.edit.current.fieldvalue, 
      "test", it_store.edit.current.item.id, null, "new")
    expect(values[values.length-1].value).toBe("new")

    values = editorActions(it_store, setData).setFieldvalue(it_store.edit.current.fieldvalue, 
      "test", it_store.edit.current.item.id, "default", null)
    expect(values[values.length-1].value).toBe("default")

  })

  it('tableValues', () => {
    const setData = jest.fn()
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    InitItem.mockImplementation((data, setData)=>{
      return (params) => {
        return {
          id: null, fieldname: null, ref_id: null, value: null, notes: null, deleted: 0
        }
      }
    })
    const values = editorActions(it_store, setData).tableValues("test", { id: 1, missing: "value" })
    expect(values.id).toBe(1)
  })

  it('saveEditorForm', async () => {
    const setData = jest.fn()
    InitItem.mockReturnValue(jest.fn( () => ({
      id: null, value: null
    }) ))
    Validator.mockReturnValue(jest.fn( () => ({}) ))
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    it_store = update(it_store, {login: {data: {$merge:{
      groups: it_store.edit.dataset.groups
    }}}})
    it_store = update(it_store, {edit: {current: {$merge:{
      form_datatype: "item",
      form: { id: null, value: "value"}
    }}}})
    let values = await editorActions(it_store, setData).saveEditorForm()
    expect(values).toBeDefined()

    it_store = update(it_store, {edit: {current: {$merge:{
      form_datatype: "movement",
      form_type: "movement",
      transtype: "inventory",
      form: { id: 1, value: "value" }
    }}}})
    values = await editorActions(it_store, setData).saveEditorForm()
    expect(values).toBeDefined()

    it_store = update(it_store, {edit: {
      current: {$merge:{
        transtype: "delivery",
        form: { id: 1, value: "value", qty: 1, product_id: 1, notes: "", place_id: 1 }
      }},
      dataset: {$merge:{
        movement_transfer: []
      }}
    }})
    values = await editorActions(it_store, setData).saveEditorForm()
    expect(values).toBeDefined()

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if((path === "/link") || ((path === "/movement") && (options.data.length === 3))){
          return { error: {} }
        }
        return {}
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    values = await editorActions(it_store, setData).saveEditorForm()
    expect(values).toBeNull()

    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        movement_transfer: [
          { id: 1, ref_id: 1 },
          { id: 2, ref_id: 1 },
          { id: 3, ref_id: 1 }
        ],
        movement: [
          { id: 1, ref_id: 1, place_id: 2 },
          { id: 2, ref_id: 1, place_id: 2 }
        ]
      }}
    }})
    values = await editorActions(it_store, setData).saveEditorForm()
    expect(values).toBeNull()

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return {}
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    values = await editorActions(it_store, setData).saveEditorForm()
    expect(values).toBeDefined()

    it_store = update(it_store, {edit: {current: {$merge:{
      form_datatype: "price",
      form_type: "price",
      form: { id: 1, value: "value" },
      price_link_customer: null,
      price_customer_id: null
    }}}})
    values = await editorActions(it_store, setData).saveEditorForm()
    expect(values).toBeDefined()

    it_store = update(it_store, {edit: {current: {$merge:{
      form_datatype: "price",
      form_type: "price",
      form: { id: 1, value: "value" },
      price_link_customer: null,
      price_customer_id: 1
    }}}})
    values = await editorActions(it_store, setData).saveEditorForm()
    expect(values).toBeDefined()

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if((path === "/link") && (options.method === "POST")){
          return { error: {} }
        }
        return {}
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    values = await editorActions(it_store, setData).saveEditorForm()
    expect(values).toBeNull()

    it_store = update(it_store, {edit: {current: {$merge:{
      form_datatype: "price",
      form_type: "price",
      form: { id: 1, value: "value" },
      price_link_customer: 1,
      price_customer_id: null
    }}}})
    values = await editorActions(it_store, setData).saveEditorForm()
    expect(values).toBeDefined()

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if((path === "/link") && (options.method === "DELETE")){
          return { error: {} }
        }
        return {}
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    values = await editorActions(it_store, setData).saveEditorForm()
    expect(values).toBeNull()

    it_store = update(it_store, {edit: {current: {$merge:{
      form_datatype: "price",
      form_type: "invoice_link",
      form: { id: 1, value: "value" },
      invoice_link_fieldvalue: [
        { id: 1, ref_id: null },
        { id: 2, ref_id: 1 }
      ]
    }}}})
    values = await editorActions(it_store, setData).saveEditorForm()
    expect(values).toBeDefined()

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/fieldvalue"){
          return { error: {} }
        }
        return {}
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    values = await editorActions(it_store, setData).saveEditorForm()
    expect(values).toBeNull()

    it_store = update(it_store, {edit: {current: {$merge:{
      form_datatype: "error",
    }}}})
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/error"){
          return { error: {} }
        }
        return {}
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    values = await editorActions(it_store, setData).saveEditorForm()
    expect(values).toBeNull()

    Validator.mockReturnValue(jest.fn( () => ({ error: {} }) ))
    values = await editorActions(it_store, setData).saveEditorForm()
    expect(values).toBeNull()

  })

  it('checkSubtype', () => {
    const setData = jest.fn()
    let value = editorActions(store, setData).checkSubtype("customer", null, { custtype: 1 })
    expect(value).toBeTruthy()

    value = editorActions(store, setData).checkSubtype("customer", 1, { custtype: 1 })
    expect(value).toBeTruthy()

    value = editorActions(store, setData).checkSubtype("missing", 1, { custtype: 1 })
    expect(value).toBeTruthy()
  })

  it('saveEditor', async () => {
    const setData = jest.fn()
    InitItem.mockReturnValue(jest.fn( () => {
      return update(InvoiceData.args.current.item, {$merge: {
        id: null
      }})
    } ))
    Validator.mockReturnValue(jest.fn( () => {
      return update(InvoiceData.args.current.item, {$merge: {
        id: null
      }})
    } ))
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view"){
          return { fields: [{ 
            trans_custinvoice_compname: "compname", trans_custinvoice_custname: "custname"
          }] }
        }
        return [1,2]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    it_store = update(it_store, {edit: {current: {item: {$merge:{
      id: null
    }}}}})
    let values = await editorActions(it_store, setData).saveEditor()
    expect(values).toBeDefined()

    it_store = update(it_store, {edit: {current: {item: {$merge:{
      id: 5
    }}}}})
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view"){
          return { fields: [] }
        }
        if(path === "/fieldvalue"){
          return { error: {} }
        }
        return [1,2]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    values = await editorActions(it_store, setData).saveEditor()
    expect(values).toBeNull()

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view"){
          return { error: {} }
        }
        return [1,2]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    values = await editorActions(it_store, setData).saveEditor()
    expect(values).toBeNull()

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    values = await editorActions(it_store, setData).saveEditor()
    expect(values).toBeNull()

    Validator.mockReturnValue(jest.fn( () => {
      return { error: {} }
    } ))
    values = await editorActions(it_store, setData).saveEditor()
    expect(values).toBeNull()

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return [1,2]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    Validator.mockReturnValue(jest.fn( () => {
      return update(InvoiceData.args.current.item, {$merge: {
        id: null
      }})
    } ))
    it_store = update(it_store, {edit: {current: {$merge:{
      transtype: "worksheet"
    }}}})
    values = await editorActions(it_store, setData).saveEditor()
    expect(values).toBeDefined()

    it_store = update(it_store, {edit: {current: {$merge:{
      transtype: "rent"
    }}}})
    values = await editorActions(it_store, setData).saveEditor()
    expect(values).toBeDefined()

    it_store = update(it_store, {edit: {
      current: {$merge:{
        transtype: "delivery"
      }},
      dataset: {$merge:{
        movement: [
          { id: 1, ref_id: 1, place_id: 2, shippingdate: "2021-12-10" },
          { id: 2, ref_id: 1, place_id: 2, shippingdate: "2021-12-11" }
        ]
      }}
    }})
    values = await editorActions(it_store, setData).saveEditor()
    expect(values).toBeDefined()

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/movement"){
          return { error: {} }
        }
        return [1,2]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    values = await editorActions(it_store, setData).saveEditor()
    expect(values).toBeNull()

    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        movement: []
      }}
    }})
    values = await editorActions(it_store, setData).saveEditor()
    expect(values).toBeDefined()

    it_store = update(it_store, {edit: {
      current: {$merge:{
        transtype: "offer"
      }}
    }})
    values = await editorActions(it_store, setData).saveEditor()
    expect(values).toBeDefined()

    it_store = update(store, {edit: {$merge:{
      ...PrintQueue.args
    }}})
    values = await editorActions(it_store, setData).saveEditor()
    expect(values).toBeDefined()

    it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    it_store = update(it_store, {edit: {
      current: {$merge:{
        transtype: "waybill",
        extend: {
          seltype: "employee",
          ref_id: null,
          refnumber: "DMEMP/00001",
          transtype: "",
        }
      }},
      dataset: {$merge:{
        translink: [
          { id: 1 }
        ]
      }}
    }})
    values = await editorActions(it_store, setData).saveEditor()
    expect(values).toBeDefined()

    it_store = update(it_store, {edit: {
      current: {$merge:{
        extend: {
          seltype: "employee",
          ref_id: 5,
          refnumber: "DMEMP/00001",
          transtype: "",
        }
      }}
    }})
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/link"){
          return { error: {} }
        }
        return [1,2]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    values = await editorActions(it_store, setData).saveEditor()
    expect(values).toBeNull()

    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        translink: []
      }}
    }})
    values = await editorActions(it_store, setData).saveEditor()
    expect(values).toBeDefined()

    it_store = update(it_store, {edit: {
      current: {$merge:{
        extend: {
          seltype: "transitem",
          ref_id: 5,
          refnumber: "DMEMP/00001",
          transtype: "",
        }
      }}
    }})
    it_store = update(it_store, {login: {data: {$merge:{
      groups: it_store.edit.dataset.groups
    }}}})
    values = await editorActions(it_store, setData).saveEditor()
    expect(values).toBeDefined()

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return [1,2]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        translink: [
          { id: 1 }
        ]
      }}
    }})
    values = await editorActions(it_store, setData).saveEditor()
    expect(values).toBeDefined()

    it_store = update(it_store, {edit: {
      current: {$merge:{
        transtype: "formula",
        extend: {
          id: null, trans_id: null, shippingdate: "2021-12-01T00:00:00Z", product_id: 4,
          product: "DMPROD/00004 | Car", movetype: 92, tool_id: null, qty: 5,
          place_id: null, shared: 0, notes: null, deleted: 0, description: "Car",
          partnumber: "DMPROD/00004",
        }
      }}
    }})
    values = await editorActions(it_store, setData).saveEditor()
    expect(values).toBeDefined()

    it_store = update(it_store, {edit: {
      current: {$merge:{
        transtype: "production"
      }}
    }})
    values = await editorActions(it_store, setData).saveEditor()
    expect(values).toBeDefined()

    it_store = update(it_store, {edit: {
      current: {$merge:{
        transtype: "cash"
      }}
    }})
    values = await editorActions(it_store, setData).saveEditor()
    expect(values).toBeDefined()

  })

  it('calcFormula', () => {
    const setData = jest.fn((key, data, callback)=>{
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(data.modalForm, container);

        // onOK
        const btn_ok = getById(container, 'btn_ok')
        btn_ok.dispatchEvent(new MouseEvent('click', {bubbles: true}));

        // onCancel
        const btn_cancel = getById(container, 'btn_cancel')
        btn_cancel.dispatchEvent(new MouseEvent('click', {bubbles: true}));
      }
      if(callback){callback()}
    })
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        movement_head: [
          { deleted: 0, description: 'Car', id: 40, movetype: 89, notes: 'demo',
            partnumber: 'DMPROD/00004', place_id: 3, product: 'DMPROD/00004 | Car',
            product_id: 4, qty: 2, shared: 1, shippingdate: '2021-12-02T00:00:00',
            tool_id: null, trans_id: 20} ],
        formula_head: [
          { id: 18, qty: 5, transnumber: 'DMFRM/00001' },
          { id: 19, qty: 5, transnumber: 'DMFRM/00002' } ],
        movement: [
          { deleted: 0, description: 'Wheel', id: 41, movetype: 89,
            notes: 'demo', opposite_qty: '8', partnumber: 'DMPROD/00005',
            place_id: 4, planumber: 'material', product: 'DMPROD/00005 | Wheel',
            product_id: 5, qty: -8, shared: 0, shippingdate: '2021-12-01T00:00:00',
            tool_id: null, trans_id: 20, unit: 'piece' } ]
      }}
    }})
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view"){
          if(options.data.length > 1){
            return { error: {} }
          }
          return { formula: [
            { deleted: 0, id: 31, movetype: 91, notes: null, place_id: 4, product_id: 5,
              qty: 20, shared: 0, shippingdate: "2021-12-01T00:00:00", tool_id: null, trans_id: 18 },
            { deleted: 0, id: 31, movetype: 91, notes: null, place_id: null, product_id: 5,
              qty: 20, shared: 1, shippingdate: "2021-12-01T00:00:00", tool_id: null, trans_id: 18 }
          ] }
        }
        return [1,2]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    editorActions(it_store, setData).calcFormula(18)
    expect(setData).toHaveBeenCalledTimes(3);

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view"){
          return { formula: [
            { deleted: 0, id: 31, movetype: 91, notes: null, place_id: 4, product_id: 5,
              qty: 20, shared: 0, shippingdate: "2021-12-01T00:00:00", tool_id: null, trans_id: 18 },
            { deleted: 0, id: 31, movetype: 91, notes: null, place_id: null, product_id: 5,
              qty: 20, shared: 1, shippingdate: "2021-12-01T00:00:00", tool_id: null, trans_id: 18 }
          ] }
        }
        if(options.method === "POST"){
          return { error: {} }
        }
        return [1,2]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    editorActions(it_store, setData).calcFormula(18)
    expect(setData).toHaveBeenCalledTimes(6);

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view"){
          return { formula: [
            { deleted: 0, id: 31, movetype: 91, notes: null, place_id: 4, product_id: 5,
              qty: 20, shared: 0, shippingdate: "2021-12-01T00:00:00", tool_id: null, trans_id: 18 },
            { deleted: 0, id: 31, movetype: 91, notes: null, place_id: null, product_id: 5,
              qty: 20, shared: 1, shippingdate: "2021-12-01T00:00:00", tool_id: null, trans_id: 18 }
          ] }
        }
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    editorActions(it_store, setData).calcFormula(18)
    expect(setData).toHaveBeenCalledTimes(9);

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    editorActions(it_store, setData).calcFormula(18)
    expect(setData).toHaveBeenCalledTimes(12);

  })

  it('createTrans', async () => {
    const setData = jest.fn()
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        payment: [],
      }}
    }})
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view"){
          return { fields: [{ 
            trans_custinvoice_compname: "compname", trans_custinvoice_custname: "custname"
          }] }
        }
        return [1,2]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    await editorActions(it_store, setData).createTrans({
      cmdtype: "copy", transcast: "normal",
    })
    expect(setData).toHaveBeenCalledTimes(1);

    it_store = update(it_store, {edit: {
      dataset: {trans: {0: {$merge:{
        duedate: null
      }}}}
    }})
    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        settings: [],
      }}
    }})
    await editorActions(it_store, setData).createTrans({
      cmdtype: "copy"
    })
    expect(setData).toHaveBeenCalledTimes(2);

    it_store = update(it_store, {edit: {
      dataset: {trans: {0: {$merge:{
        transtype: 56
      }}}}
    }})
    await editorActions(it_store, setData).createTrans({
      cmdtype: "copy", transcast: "normal",
    })
    expect(setData).toHaveBeenCalledTimes(3);

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/function"){
          return { error: {} }
        }
        return [1,2]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    let value = await editorActions(it_store, setData).createTrans({
      cmdtype: "copy", transcast: "normal",
    })
    expect(value).toBeNull()

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/trans"){
          return { error: {} }
        }
        return [1,2]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    value = await editorActions(it_store, setData).createTrans({
      cmdtype: "copy", transcast: "normal",
    })
    expect(value).toBeNull()

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view") { 
          return { error: {} }
        }
        return [1,2]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    it_store = update(it_store, {edit: {
      dataset: {trans: {0: {$merge:{
        transtype: 55
      }}}}
    }})
    value = await editorActions(it_store, setData).createTrans({
      cmdtype: "copy", transcast: "normal",
    })
    expect(value).toBeNull()

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view") { 
          return { fields: [{ 
            trans_custinvoice_compname: "compname", trans_custinvoice_custname: "custname"
          }] }
        }
        if(path === "/fieldvalue") {
          return { error: {} }
        }
        return [1,2]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    await editorActions(it_store, setData).createTrans({
      cmdtype: "copy", transcast: "normal",
    })
    expect(value).toBeNull()

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view") { 
          return { fields: [] }
        }
        return [1,2]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    it_store = update(it_store, {edit: {
      current: {$merge:{
        fieldvalue: []
      }}
    }})
    await editorActions(it_store, setData).createTrans({
      cmdtype: "copy", transcast: "normal",
    })
    expect(setData).toHaveBeenCalledTimes(4);

    it_store = update(it_store, {edit: {
      dataset: {trans: {0: {$merge:{
        transcast: "cancellation"
      }}}}
    }})
    value = await editorActions(it_store, setData).createTrans({
      cmdtype: "copy", transcast: "normal",
    })
    expect(value).toBeFalsy()

    it_store = update(it_store, {edit: {
      dataset: {trans: {0: {$merge:{
        transtype: 56, direction: 69
      }}}}
    }})
    value = await editorActions(it_store, setData).createTrans({
      cmdtype: "copy", transcast: "normal",
    })
    expect(value).toBeFalsy()

    it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        payment: [
          { id: 123, deleted: 0}
        ]
      }}
    }})
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view") { 
          return { fields: [{ 
            trans_custinvoice_compname: "compname", trans_custinvoice_custname: "custname"
          }] }
        }
        if(path === "/link") {
          return { error: {} }
        }
        return [1,2]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    value = await editorActions(it_store, setData).createTrans({
      cmdtype: "copy", transcast: "amendment",
    })
    expect(value).toBeNull()

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view") { 
          return { fields: [{ 
            trans_custinvoice_compname: "compname", trans_custinvoice_custname: "custname"
          }] }
        }
        if(path === "/item") {
          return { error: {} }
        }
        return [1,2]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    value = await editorActions(it_store, setData).createTrans({
      cmdtype: "copy",
    })
    expect(value).toBeNull()

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view") { 
          return { fields: [{ 
            trans_custinvoice_compname: "compname", trans_custinvoice_custname: "custname"
          }] }
        }
        if(path === "/payment") {
          return { error: {} }
        }
        return [1,2]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    it_store = update(it_store, {edit: {
      dataset: {trans: {0: {$merge:{
        transtype: 63
      }}}}
    }})
    value = await editorActions(it_store, setData).createTrans({
      cmdtype: "copy",
    })
    expect(value).toBeNull()

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view") { 
          return { fields: [{ 
            trans_custinvoice_compname: "compname", trans_custinvoice_custname: "custname"
          }] }
        }
        return [1,2]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    value = await editorActions(it_store, setData).createTrans({
      cmdtype: "copy", transcast: "cancellation",
    })
    expect(value).toBeFalsy()

    it_store = update(it_store, {edit: {
      dataset: {trans: {0: {$merge:{
        deleted: 1
      }}}}
    }})
    value = await editorActions(it_store, setData).createTrans({
      cmdtype: "copy", transcast: "amendment",
    })
    expect(value).toBeFalsy()

    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        cancel_link: [{ transnumber: "transnumber" }],
      }}
    }})
    value = await editorActions(it_store, setData).createTrans({
      cmdtype: "copy", transcast: "cancellation",
    })
    expect(value).toBeFalsy()

    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        cancel_link: [],
        payment: [
          { id: 123, deleted: 0}
        ],
        movement: [
          { id: 1, item_id: 1, ref_id: 1, deleted: 0 },
          { id: 1, item_id: 1, ref_id: 1, deleted: 1 },
          { id: 1, item_id: null, ref_id: null, deleted: 0 },
          { id: 1, item_id: null, ref_id: 1, deleted: 0 },
        ]
      }}
    }})
    await editorActions(it_store, setData).createTrans({
      cmdtype: "copy", transcast: "cancellation",
    })
    expect(setData).toHaveBeenCalledTimes(5);

    it_store = update(it_store, {edit: {
      dataset: {trans: {0: {$merge:{
        transtype: 61
      }}}}
    }})
    await editorActions(it_store, setData).createTrans({
      cmdtype: "copy", transcast: "cancellation",
    })
    expect(setData).toHaveBeenCalledTimes(6);

    await editorActions(it_store, setData).createTrans({
      cmdtype: "copy", transcast: "normal"
    })
    expect(setData).toHaveBeenCalledTimes(7);

    it_store = update(it_store, {edit: {
      dataset: {trans: {0: {$merge:{
        direction: 70
      }}}}
    }})
    await editorActions(it_store, setData).createTrans({
      cmdtype: "create", transcast: "normal"
    })
    expect(setData).toHaveBeenCalledTimes(8);

    it_store = update(it_store, {edit: {
      dataset: {trans: {0: {$merge:{
        transtype: 64
      }}}}
    }})
    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        movement_head: [
          { deleted: 0, description: 'Car', id: 40, movetype: 89, notes: 'demo',
            partnumber: 'DMPROD/00004', place_id: 3, product: 'DMPROD/00004 | Car',
            product_id: 4, qty: 2, shared: 1, shippingdate: '2021-12-02T00:00:00',
            tool_id: null, trans_id: 20} ],
        formula_head: [
          { id: 18, qty: 5, transnumber: 'DMFRM/00001' },
          { id: 19, qty: 5, transnumber: 'DMFRM/00002' } ],
        movement: [
          { deleted: 0, description: 'Wheel', id: 41, movetype: 89,
            notes: 'demo', opposite_qty: '8', partnumber: 'DMPROD/00005',
            place_id: 4, planumber: 'material', product: 'DMPROD/00005 | Wheel',
            product_id: 5, qty: -8, shared: 0, shippingdate: '2021-12-01T00:00:00',
            tool_id: null, trans_id: 20, unit: 'piece' } ]
      }}
    }})
    await editorActions(it_store, setData).createTrans({
      cmdtype: "create", transcast: "normal"
    })
    expect(setData).toHaveBeenCalledTimes(9);

    it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        payment: [
          { id: 123, deleted: 0}
        ]
      }}
    }})
    await editorActions(it_store, setData).createTrans({
      cmdtype: "copy", transcast: "amendment",
    })
    expect(setData).toHaveBeenCalledTimes(10);

    it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    it_store = update(it_store, {edit: {
      dataset: {trans: {0: {$merge:{
        transtype: 57
      }}}}
    }})
    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        payment: [
          { id: 123, deleted: 0},
          { id: 123, deleted: 1}
        ]
      }}
    }})
    await editorActions(it_store, setData).createTrans({
      cmdtype: "create", transcast: "normal", new_transtype: "invoice", new_direction: "out",
      refno: true, from_inventory: false, netto_qty: true,
    })
    expect(setData).toHaveBeenCalledTimes(11);

    it_store = update(it_store, {login: {
      data: {$merge:{
        audit: [
          { nervatypeName: "trans", subtypeName: "invoice", inputfilterName: "disabled" }
        ]
      }}
    }})
    value = await editorActions(it_store, setData).createTrans({
      cmdtype: "create", transcast: "normal", new_transtype: "invoice", new_direction: "out",
      refno: true, from_inventory: false, netto_qty: true,
    })
    expect(value).toBeFalsy()

    it_store = update(it_store, {
      login: {
        data: {$merge:{
          audit: [
            { nervatypeName: "trans", subtypeName: "invoice", inputfilterName: "all" }
          ]
        }}
      },
      edit: {
        dataset: {$merge: {
          transitem_invoice: [
            { product_id: 2, qty: 1, deposit: 0 },
            { product_id: 1, qty: 1, deposit: 1 },
            { product_id: 3, qty: 1, deposit: 0 }
          ],
          transitem_shipping: [
            { id: "19-1", item_product: "Big product", movement_product: "Big product", product_id: 1, sqty: "-2" },
            { id: "200-3", item_product: "Nice product", movement_product: "Nice product", product_id: 3, sqty: "-1" },
            { id: "20-3", item_product: "Nice product", movement_product: "Nice product", product_id: 3, sqty: "-3" },
            { id: "201-3", item_product: "Nice product", movement_product: "Nice product", product_id: 3, sqty: "-1" },
            { id: "205-3", item_product: "Nice product", movement_product: "Nice product", product_id: 3, sqty: "-1" }
          ]
        }}
      }
    })
    InitItem.mockImplementation((data, setData)=>{
      return (params) => {
        if(params.tablename === "item"){
          return it_store.edit.dataset.item[0]
        }
        return {
          id: null, fieldname: null, ref_id: null, value: null, notes: null, deleted: 0
        }
      }
    })
    await editorActions(it_store, setData).createTrans({
      cmdtype: "create", transcast: "normal", new_transtype: "invoice", new_direction: "out",
      refno: true, from_inventory: false, netto_qty: true,
    })
    expect(setData).toHaveBeenCalledTimes(12);

    await editorActions(it_store, setData).createTrans({
      cmdtype: "create", transcast: "normal", new_transtype: "invoice", new_direction: "out",
      refno: true, from_inventory: true, netto_qty: true,
    })
    expect(setData).toHaveBeenCalledTimes(13);

    it_store = update(it_store, {edit: {
      dataset: {trans: {0: {$merge:{
        direction: 69
      }}}}
    }})
    await editorActions(it_store, setData).createTrans({
      cmdtype: "create", transcast: "normal", new_transtype: "invoice", new_direction: "out",
      refno: true, from_inventory: true, netto_qty: true,
    })
    expect(setData).toHaveBeenCalledTimes(14);

    it_store = update(it_store, {
      edit: {
        dataset: {$merge: {
          item: []
        }}
      }
    })
    value = await editorActions(it_store, setData).createTrans({
      cmdtype: "create", transcast: "normal", new_transtype: "invoice", new_direction: "out",
      refno: true, from_inventory: false, netto_qty: true,
    })
    expect(setData).toHaveBeenCalledTimes(15);

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view"){
          return { fields: [{ 
            trans_custinvoice_compname: "compname", trans_custinvoice_custname: "custname"
          }] }
        }
        if((path === "/link") && (options.data.length > 1)) { 
          return { error: {} }
        }
        return [1,2]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        cancel_link: [],
        transitem_invoice: [
          { product_id: 2, qty: 1, deposit: 0 },
          { product_id: 1, qty: 1, deposit: 1 },
          { product_id: 3, qty: 1, deposit: 0 }
        ],
        payment: [
          { id: 123, deleted: 0}
        ],
        movement: [
          { id: 1, item_id: 1, ref_id: 1, deleted: 0 },
          { id: 1, item_id: 1, ref_id: 1, deleted: 1 },
          { id: 1, item_id: null, ref_id: null, deleted: 0 },
          { id: 1, item_id: null, ref_id: 1, deleted: 0 },
        ]
      }}
    }})
    value = await editorActions(it_store, setData).createTrans({
      cmdtype: "copy",
    })
    expect(value).toBeNull()

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view"){
          return { fields: [{ 
            trans_custinvoice_compname: "compname", trans_custinvoice_custname: "custname"
          }] }
        }
        if(path === "/movement") { 
          return { error: {} }
        }
        return [1,2]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    value = await editorActions(it_store, setData).createTrans({
      cmdtype: "copy",
    })
    expect(value).toBeNull()

  })

  it('createTransOptions', () => {
    const setData = jest.fn((key, data, callback)=>{
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(data.modalForm, container);

        // onOK
        const btn_ok = getById(container, 'btn_create')
        btn_ok.dispatchEvent(new MouseEvent('click', {bubbles: true}));

        // onCancel
        const btn_cancel = getById(container, 'btn_cancel')
        btn_cancel.dispatchEvent(new MouseEvent('click', {bubbles: true}));
      }
      if(callback){callback()}
    })
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        element_count: [
          { pec: 0 }
        ],
        payment: []
      }}
    }})
    editorActions(it_store, setData).createTransOptions()
    expect(setData).toHaveBeenCalledTimes(3);

    it_store = update(it_store, {edit: {
      current: {$merge:{
        transtype: "offer"
      }}
    }})
    editorActions(it_store, setData).createTransOptions()
    expect(setData).toHaveBeenCalledTimes(6);

    it_store = update(it_store, {edit: {
      current: {$merge:{
        transtype: "order"
      }}
    }})
    editorActions(it_store, setData).createTransOptions()
    expect(setData).toHaveBeenCalledTimes(9);

    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        element_count: [
          { pec: 1 }
        ]
      }}
    }})
    editorActions(it_store, setData).createTransOptions()
    expect(setData).toHaveBeenCalledTimes(12);

    it_store = update(it_store, {edit: {
      current: {$merge:{
        transtype: "worksheet"
      }}
    }})
    editorActions(it_store, setData).createTransOptions()
    expect(setData).toHaveBeenCalledTimes(15);

    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        element_count: [
          { pec: 0 }
        ]
      }}
    }})
    editorActions(it_store, setData).createTransOptions()
    expect(setData).toHaveBeenCalledTimes(18);

    it_store = update(it_store, {edit: {
      current: {$merge:{
        transtype: "rent"
      }}
    }})
    editorActions(it_store, setData).createTransOptions()
    expect(setData).toHaveBeenCalledTimes(21);

    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        element_count: [
          { pec: 1 }
        ]
      }}
    }})
    editorActions(it_store, setData).createTransOptions()
    expect(setData).toHaveBeenCalledTimes(24);

    it_store = update(it_store, {edit: {
      current: {$merge:{
        transtype: "receipt"
      }}
    }})
    editorActions(it_store, setData).createTransOptions()
    expect(setData).toHaveBeenCalledTimes(27);

  })

  it('checkEditor', () => {
    let setData = jest.fn((key, data, callback)=>{
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(data.modalForm, container);

        // onOK
        const btn_ok = getById(container, 'btn_ok')
        btn_ok.dispatchEvent(new MouseEvent('click', {bubbles: true}));

        // onCancel
        const btn_cancel = getById(container, 'btn_cancel')
        btn_cancel.dispatchEvent(new MouseEvent('click', {bubbles: true}));
      }
      if(callback){callback()}
    })
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    editorActions(it_store, setData).checkEditor({},"")
    expect(setData).toHaveBeenCalledTimes(0);

    it_store = update(it_store, {edit: {$merge:{
        dirty: true
    }}})
    editorActions(it_store, setData).checkEditor({ ntype: "printqueue" },"LOAD_EDITOR")
    expect(setData).toHaveBeenCalledTimes(3);

    setData = jest.fn((key, data, callback)=>{
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(data.modalForm, container);

        // onOK
        const btn_ok = getById(container, 'btn_ok')
        btn_ok.dispatchEvent(new MouseEvent('click', {bubbles: true}));
      }
      if(callback){callback()}
    })
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view"){
          return { fields: [{ 
            trans_custinvoice_compname: "compname", trans_custinvoice_custname: "custname"
          }] }
        }
        return it_store.edit.dataset
      }),
      getAuditFilter: jest.fn(() => (["all",1])),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    it_store = update(it_store, {login: {data: {$merge:{
      groups: it_store.edit.dataset.groups
    }}}})
    it_store = update(it_store, {edit: {current: {$merge:{
      form_datatype: "item",
      form: { id: null, value: "value"}
    }}}})
    it_store = update(it_store, {edit: {$merge:{
      dirty: false,
      form_dirty: true
    }}})
    editorActions(it_store, setData).checkEditor(
      { ntype: "trans", ttype: "invoice", id: 5, form: "item", params: { action: "" } }, "FORM_ACTIONS")
    expect(setData).toHaveBeenCalledTimes(2);

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    it_store = update(it_store, {edit: {$merge:{
      dirty: false,
      form_dirty: false,
      panel: it_store.edit.template.options.panel
    }}})
    editorActions(it_store, setData).checkEditor({fkey: "item", id: 5}, "SET_EDITOR_ITEM")
    expect(setData).toHaveBeenCalledTimes(3);

    editorActions(it_store, setData).checkEditor({fieldname: "sample_bool"}, "NEW_FIELDVALUE")
    expect(setData).toHaveBeenCalledTimes(3);

    editorActions(it_store, setData).checkEditor({cmdtype: "copy"}, "CREATE_TRANS")
    expect(setData).toHaveBeenCalledTimes(3);

    setData = jest.fn((key, data, callback)=>{
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(data.modalForm, container);

        // onCancel
        const btn_cancel = getById(container, 'btn_cancel')
        btn_cancel.dispatchEvent(new MouseEvent('click', {bubbles: true}));
      }
      if(callback){callback()}
    })
    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        element_count: [
          { pec: 0 }
        ],
        payment: []
      }}
    }})
    editorActions(it_store, setData).checkEditor({}, "CREATE_TRANS_OPTIONS")
    expect(setData).toHaveBeenCalledTimes(2);

    setData = jest.fn((key, data, callback)=>{
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(data.modalForm, container);

        // onOK
        const btn_ok = getById(container, 'btn_formula')
        if(btn_ok){
          btn_ok.dispatchEvent(new MouseEvent('click', {bubbles: true}));
        }

        // onCancel
        const btn_cancel = getById(container, 'btn_cancel')
        btn_cancel.dispatchEvent(new MouseEvent('click', {bubbles: true}));
      }
      if(callback){callback()}
    })
    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        movement_head: [
          { deleted: 0, description: 'Car', id: 40, movetype: 89, notes: 'demo',
            partnumber: 'DMPROD/00004', place_id: 3, product: 'DMPROD/00004 | Car',
            product_id: 4, qty: 2, shared: 1, shippingdate: '2021-12-02T00:00:00',
            tool_id: null, trans_id: 20} ],
        formula_head: [
          { id: 18, qty: 5, transnumber: 'DMFRM/00001' },
          { id: 19, qty: 5, transnumber: 'DMFRM/00002' } ],
        movement: [
          { deleted: 0, description: 'Wheel', id: 41, movetype: 89,
            notes: 'demo', opposite_qty: '8', partnumber: 'DMPROD/00005',
            place_id: 4, planumber: 'material', product: 'DMPROD/00005 | Wheel',
            product_id: 5, qty: -8, shared: 0, shippingdate: '2021-12-01T00:00:00',
            tool_id: null, trans_id: 20, unit: 'piece' } ]
      }}
    }})
    editorActions(it_store, setData).checkEditor({ formula: "19" }, "LOAD_FORMULA")
    expect(setData).toHaveBeenCalledTimes(5);

  })

  it('checkTranstype', async () => {
    let setData = jest.fn((key, data, callback)=>{
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(data.modalForm, container);

        // onOK
        const btn_ok = getById(container, 'btn_ok')
        btn_ok.dispatchEvent(new MouseEvent('click', {bubbles: true}));

        // onCancel
        const btn_cancel = getById(container, 'btn_cancel')
        btn_cancel.dispatchEvent(new MouseEvent('click', {bubbles: true}));
      }
      if(callback){callback()}
    })
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    it_store = update(it_store, {login: {data: {$merge:{
      groups: it_store.edit.dataset.groups
    }}}})
    await editorActions(it_store, setData).checkTranstype({fieldname: "sample_bool"},"NEW_FIELDVALUE")
    expect(setData).toHaveBeenCalledTimes(0);

    await editorActions(it_store, setData).checkTranstype({ ntype: "trans", ttype: null, fieldname: "sample_bool"},"NEW_FIELDVALUE")
    expect(setData).toHaveBeenCalledTimes(0);

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view"){
          return { transtype: [{ 
            groupvalue: "invoice"
          }] }
        }
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    await editorActions(it_store, setData).checkTranstype(
      { ntype: "trans", ttype: null, id: 1, fieldname: "sample_bool"},"NEW_FIELDVALUE")
    expect(setData).toHaveBeenCalledTimes(0);

  })

  it('showStock', async () => {
    let setData = jest.fn((key, data, callback)=>{
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(data.modalForm, container);

        // onOK
        const btn_ok = getById(container, 'btn_ok')
        btn_ok.dispatchEvent(new MouseEvent('click', {bubbles: true}));
      }
      if(callback){callback()}
    })
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { stock: [
          { id: 1, partnumber: "partnumber", description: "description", unit: "unit",
            warehouse: "warehouse", sqty: 1 }
        ] }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    await editorActions(it_store, setData).showStock({ partnumber: "partnumber", partname: "partname" })
    expect(setData).toHaveBeenCalledTimes(2);

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { stock: [] }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    await editorActions(it_store, setData).showStock({ partnumber: "partnumber", partname: "partname" })
    expect(setData).toHaveBeenCalledTimes(2);

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    await editorActions(it_store, setData).showStock({ partnumber: "partnumber", partname: "partname" })
    expect(setData).toHaveBeenCalledTimes(2);

  })

  it('exportQueue', async () => {
    const setData = jest.fn()
    let it_store = update(store, {edit: {$merge:{
      ...PrintQueue.args
    }}})
    await editorActions(it_store, setData).exportQueue({ 
      id: 1, ref_id: 1, reportkey: "reportkey", refnumber: "refnumber", copies: 1, typename: "typename" })
    expect(setData).toHaveBeenCalledTimes(0);

  })

  it('setFormActions', () => {
    const setData = jest.fn((key, data, callback)=>{
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(data.modalForm, container);

        // onOK
        const btn_ok = getById(container, 'btn_ok')
        if(btn_ok){
          btn_ok.dispatchEvent(new MouseEvent('click', {bubbles: true}));
        }

        // onCancel
        const btn_cancel = getById(container, 'btn_cancel')
        btn_cancel.dispatchEvent(new MouseEvent('click', {bubbles: true}));
      }
      if(callback){callback()}
    })
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    editorActions(it_store, setData).setFormActions({ params: { action: "loadEditor" } })
    expect(setData).toHaveBeenCalledTimes(0);

    editorActions(it_store, setData).setFormActions({ 
      params: { action: "loadEditor", ntype: "trans", ttype: "invoice", id: 5 } })
    expect(setData).toHaveBeenCalledTimes(0);

    it_store = update(it_store, {edit: {$merge:{
      panel: it_store.edit.template.options.panel
    }}})
    editorActions(it_store, setData).setFormActions({ 
      params: { action: "newEditorItem", fkey: "item" } })
    expect(setData).toHaveBeenCalledTimes(1);

    editorActions(it_store, setData).setFormActions({ 
      params: { action: "editEditorItem", fkey: "item", row: { id: 5 } } })
    expect(setData).toHaveBeenCalledTimes(2);

    editorActions(it_store, setData).setFormActions({ 
      params: { action: "deleteEditorItem", fkey: "item", table: "item", row: { id: 5 } } })
    expect(setData).toHaveBeenCalledTimes(5);

    editorActions(it_store, setData).setFormActions({ 
      params: { action: "loadShipping" } })
    expect(setData).toHaveBeenCalledTimes(5);

    editorActions(it_store, setData).setFormActions({ 
      params: { action: "loadShipping", ntype: "trans", ttype: "order", id: 5 } })
    expect(setData).toHaveBeenCalledTimes(5);

    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        shiptemp: [
          { id: 1 }
        ],
        shipping_items: []
      }}
    }})
    editorActions(it_store, setData).setFormActions({ 
      params: { action: "addShippingRow" }, row: { edited: true } })
    expect(setData).toHaveBeenCalledTimes(5);
    editorActions(it_store, setData).setFormActions({ 
      params: { action: "addShippingRow" }, 
      row: { edited: false, item_id: 1, product_id: 1, product: "product", 
        partnumber: "partnumber", partname: "partname", unit: "unit",
        diff: 0, qty: 1, tqty: 1 } })
    expect(setData).toHaveBeenCalledTimes(7);

    editorActions(it_store, setData).setFormActions({ 
      params: { action: "showShippingStock" }, row: { product_id: 1, partnumber: "partnumber", partname: "partname" } })
    expect(setData).toHaveBeenCalledTimes(7);

    editorActions(it_store, setData).setFormActions({ 
      params: { action: "editShippingRow" }, 
      row: { id: 1, item_id: 1, product_id: 1, product: "product", 
        partnumber: "partnumber", partname: "partname", unit: "unit", batch_no: "batch_no",
        diff: 0, qty: 1, tqty: 1, oqty: 1 } })
    expect(setData).toHaveBeenCalledTimes(11);

    editorActions(it_store, setData).setFormActions({ 
      params: { action: "deleteShippingRow" }, row: { id: 1 } })
    expect(setData).toHaveBeenCalledTimes(13);
    
    it_store = update(store, {edit: {$merge:{
      ...PrintQueue.args
    }}})
    editorActions(it_store, setData).setFormActions({ 
      params: { action: "exportQueueItem" }, row: { 
        id: 1, ref_id: 1, reportkey: "reportkey", refnumber: "refnumber", copies: 1, typename: "typename" } })
    expect(setData).toHaveBeenCalledTimes(13);

    editorActions(it_store, setData).setFormActions({ params: { action: "" } })
    expect(setData).toHaveBeenCalledTimes(13);

  })

  it('getTransFilter', () => {
    const setData = jest.fn()
    let it_store = update(store, {login: {data: {$merge:{
      transfilterName: "usergroup"
    }}}})
    let result = editorActions(it_store, setData).getTransFilter({ where: [] }, [])
    expect(result[1].length).toBe(1)

    it_store = update(store, {login: {data: {$merge:{
      transfilterName: "own"
    }}}})
    result = editorActions(it_store, setData).getTransFilter({ where: [] }, [])
    expect(result[1].length).toBe(1)

    it_store = update(store, {login: {data: {$merge:{
      transfilterName: "all"
    }}}})
    result = editorActions(it_store, setData).getTransFilter({ where: [] }, [])
    expect(result[1].length).toBe(0)
  })

  it('prevTransNumber', async () => {
    const setData = jest.fn()
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view"){
          return { prev: [{ 
            id: 5
          }] }
        }
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    await editorActions(it_store, setData).prevTransNumber()
    expect(setData).toHaveBeenCalledTimes(1);

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view"){
          return { prev: [{ 
            id: null
          }] }
        }
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    it_store = update(it_store, {edit: {
      current: {$merge: {
        transtype: "waybill"
      }}
    }})
    await editorActions(it_store, setData).prevTransNumber()
    expect(setData).toHaveBeenCalledTimes(1);

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    await editorActions(it_store, setData).prevTransNumber()
    expect(setData).toHaveBeenCalledTimes(1);

    it_store = update(it_store, {edit: {
      current: {$merge: {
        type: "customer"
      }}
    }})
    await editorActions(it_store, setData).prevTransNumber()
    expect(setData).toHaveBeenCalledTimes(1);

  })

  it('nextTransNumber', async () => {
    const setData = jest.fn()
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view"){
          return { next: [{ 
            id: 5
          }] }
        }
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    await editorActions(it_store, setData).nextTransNumber()
    expect(setData).toHaveBeenCalledTimes(1);

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view"){
          return { next: [{ 
            id: null
          }] }
        }
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    it_store = update(it_store, {edit: {
      current: {$merge: {
        transtype: "waybill"
      }}
    }})
    await editorActions(it_store, setData).nextTransNumber()
    expect(setData).toHaveBeenCalledTimes(2);

    it_store = update(it_store, {edit: {
      current: {$merge: {
        transtype: "delivery"
      }}
    }})
    await editorActions(it_store, setData).nextTransNumber()
    expect(setData).toHaveBeenCalledTimes(2);

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    await editorActions(it_store, setData).nextTransNumber()
    expect(setData).toHaveBeenCalledTimes(2);

  })

  it('createShipping', async () => {
    const setData = jest.fn()
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    it_store = update(it_store, {edit: {
      current: {$merge: {
        shipping_place_id: null
      }},
    }})
    editorActions(it_store, setData).createShipping()
    expect(setData).toHaveBeenCalledTimes(0)

    it_store = update(it_store, {edit: {current: {item: {$merge:{
      shippingdate: "2021-12-01T00:00:00",
    }}}}})
    it_store = update(it_store, {edit: {
      current: {$merge: {
        shipping_place_id: 1,
        shippingdate: "2021-12-01T00:00:00",
        direction: "out"
      }},
      dataset: {$merge:{
        shiptemp: [
          { id: "1-5", item_product: "Wheel", movement_product: "Wheel", product_id: 5, sqty: "30" }
        ],
        delivery_pattern: [
          { defpattern: 1, notes: "" }
        ]
      }}
    }})
    await editorActions(it_store, setData).createShipping()
    expect(setData).toHaveBeenCalledTimes(1)

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/link"){
          return { error: {} }
        }
        return [1]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    it_store = update(it_store, {edit: {
      current: {$merge: {
        direction: "in"
      }},
      dataset: {$merge:{
        shiptemp: [
          { id: "1-5", item_product: "Wheel", movement_product: "Wheel", product_id: 5, sqty: "30" }
        ],
        delivery_pattern: []
      }}
    }})
    let result = await editorActions(it_store, setData).createShipping()
    expect(result).toBeNull()

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/movement"){
          return { error: {} }
        }
        return [1]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    result = await editorActions(it_store, setData).createShipping()
    expect(result).toBeNull()

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/trans"){
          return { error: {} }
        }
        return [1]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    result = await editorActions(it_store, setData).createShipping()
    expect(result).toBeNull()

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/function"){
          return { error: {} }
        }
        return [1]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    result = await editorActions(it_store, setData).createShipping()
    expect(result).toBeNull()

    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        shiptemp: [],
        delivery_pattern: []
      }}
    }})
    await editorActions(it_store, setData).createShipping()
    expect(setData).toHaveBeenCalledTimes(1)

  })

  it('exportEvent', () => {
    const setData = jest.fn()
    const eventExport = jest.fn()
    saveToDisk.mockImplementation((data, setData)=>{
      eventExport()
    })
    let it_store = update(store, {edit: {current: {$merge:{
      item: {
        id: null, calnumber: "calnumber", 
        nervatype: null, ref_id: null, 
        uid: null, eventgroup: null, fromdate: null, todate: null, subject: null, 
        place: null, description: null, deleted: 0
      }
    }}}})
    editorActions(it_store, setData).exportEvent()
    expect(eventExport).toHaveBeenCalledTimes(1)

    it_store = update(store, {edit: {
      current: {$merge:{
        item: {
          id: null, calnumber: "calnumber", 
          nervatype: null, ref_id: null, 
          uid: "uid", eventgroup: 123, 
          fromdate: "2021-12-01T00:00:00", todate: "2021-12-01T00:00:00", subject: "subject", 
          place: "place", description: "description", deleted: 0
        }
      }},
      dataset: {$merge: {
        eventgroup: []
      }}
    }})
    editorActions(it_store, setData).exportEvent()
    expect(eventExport).toHaveBeenCalledTimes(2)

    it_store = update(store, {edit: {
      current: {$merge:{
        item: {
          id: null, calnumber: "calnumber", 
          nervatype: null, ref_id: null, 
          uid: "uid", eventgroup: 123, 
          fromdate: "2021-12-01T00:00:00", todate: "2021-12-01T00:00:00", subject: "subject", 
          place: "place", description: "description", deleted: 0
        }
      }},
      dataset: {$merge: {
        eventgroup: [
          { id: 123, groupvalue: "value" }
        ]
      }}
    }})
    editorActions(it_store, setData).exportEvent()
    expect(eventExport).toHaveBeenCalledTimes(3)
  })

  it('calcPrice', () => {
    const setData = jest.fn()
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    let item = editorActions(it_store, setData).calcPrice("netamount",
      { netamount: 10, qty: 1, discount: 0 }
    )
    expect(item.amount).toBe(10)

    it_store = update(it_store, {edit: {
      current: {item: {$merge:{
        curr: "MTG"
      }}}
    }})
    item = editorActions(it_store, setData).calcPrice("netamount",
      { netamount: 10, qty: 0, discount: 0 }
    )
    expect(item.amount).toBe(10)

    item = editorActions(it_store, setData).calcPrice("amount",
      { amount: 10, qty: 1, discount: 0, tax_id: 2 }
    )
    expect(item.netamount).toBe(9.52)

    item = editorActions(it_store, setData).calcPrice("amount",
      { amount: 10, qty: 0, discount: 0, tax_id: 2 }
    )
    expect(item.netamount).toBe(0)

    item = editorActions(it_store, setData).calcPrice("fxprice",
      { fxprice: 10, qty: 1, discount: 0 }
    )
    expect(item.amount).toBe(10)

    item = editorActions(it_store, setData).calcPrice("",
      { fxprice: 10, qty: 1, discount: 0 }
    )
    expect(item.amount).toBe(10)

  })

  it('editItem', async () => {
    const setData = jest.fn((key, data, callback)=>{
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(data.modalForm, container);

        // onOK
        const btn_ok = getById(container, 'btn_ok')
        if(btn_ok){
          btn_ok.dispatchEvent(new MouseEvent('click', {bubbles: true}));
        }

        // onCancel
        const btn_cancel = getById(container, 'btn_cancel')
        btn_cancel.dispatchEvent(new MouseEvent('click', {bubbles: true}));
      }
      if(callback){callback()}
    })
    let it_store = update(store, {edit: {$merge:{
      ...Item.args
    }}})
    await editorActions(it_store, setData).editItem({
      name: "description", value: "description"
    })
    expect(setData).toHaveBeenCalledTimes(1)

    await editorActions(it_store, setData).editItem({
      name: "product_id", value: 1, item: it_store.edit.current.form
    })
    expect(setData).toHaveBeenCalledTimes(2)

    appActions.mockReturnValue({
      requestData: jest.fn(async (path, options) => {
        return {
          price: 1, discount: 0
        }
      }),
      resultError: jest.fn(),
    })
    it_store = update(it_store, {edit: {
      current: {form: {$merge:{
        qty: 0
      }}}
    }})
    await editorActions(it_store, setData).editItem({
      name: "product_id", value: 1, item: it_store.edit.current.form, event_type: "blur"
    })
    expect(setData).toHaveBeenCalledTimes(3)

    appActions.mockReturnValue({
      requestData: jest.fn(async (path, options) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
    })
    it_store = update(it_store, {edit: {
      current: {form: {$merge:{
        qty: 0
      }}}
    }})
    await editorActions(it_store, setData).editItem({
      name: "product_id", value: 1, item: it_store.edit.current.form, event_type: "blur"
    })
    expect(setData).toHaveBeenCalledTimes(3)

    appActions.mockReturnValue({
      requestData: jest.fn(async (path, options) => {
        return {
          price: 1, discount: 0
        }
      }),
      resultError: jest.fn(),
    })
    await editorActions(it_store, setData).editItem({
      name: "qty", value: 1, event_type: "blur"
    })
    expect(setData).toHaveBeenCalledTimes(4)
    it_store = update(it_store, {edit: {
      current: {form: {$merge:{
        fxprice: 0
      }}}
    }})
    await editorActions(it_store, setData).editItem({
      name: "qty", value: 1
    })
    expect(setData).toHaveBeenCalledTimes(5)

    appActions.mockReturnValue({
      requestData: jest.fn(async (path, options) => {
        return {}
      }),
      resultError: jest.fn(),
    })
    await editorActions(it_store, setData).editItem({
      name: "qty", value: 1
    })
    expect(setData).toHaveBeenCalledTimes(6)

    appActions.mockReturnValue({
      requestData: jest.fn(async (path, options) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
    })
    await editorActions(it_store, setData).editItem({
      name: "qty", value: 1
    })
    expect(setData).toHaveBeenCalledTimes(6)

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return {
          price: 1, discount: 0
        }
      }),
      resultError: jest.fn(),
    })
    await editorActions(it_store, setData).editItem({
      name: "fxprice", value: 10
    })
    expect(setData).toHaveBeenCalledTimes(7)

    await editorActions(it_store, setData).editItem({
      name: "tax_id", value: 10, event_type: "blur"
    })
    expect(setData).toHaveBeenCalledTimes(8)

    await editorActions(it_store, setData).editItem({
      name: "discount", value: 10, event_type: "blur"
    })
    expect(setData).toHaveBeenCalledTimes(9)

    await editorActions(it_store, setData).editItem({
      name: "amount", value: 10, event_type: "blur"
    })
    expect(setData).toHaveBeenCalledTimes(10)

    await editorActions(it_store, setData).editItem({
      name: "amount", value: 10
    })
    expect(setData).toHaveBeenCalledTimes(11)

    await editorActions(it_store, setData).editItem({
      name: "netamount", value: 10, event_type: "blur"
    })
    expect(setData).toHaveBeenCalledTimes(12)

    await editorActions(it_store, setData).editItem({
      name: "netamount", value: 10
    })
    expect(setData).toHaveBeenCalledTimes(13)

    it_store = update(it_store, {edit: {
      current: {$merge:{
        form_type: "price"
      }}
    }})
    await editorActions(it_store, setData).editItem({
      name: "curr", value: "EUR"
    })
    expect(setData).toHaveBeenCalledTimes(14)

    it_store = update(it_store, {edit: {
      current: {$merge:{
        form_type: "discount"
      }}
    }})
    await editorActions(it_store, setData).editItem({
      name: "customer_id", value: 1
    })
    expect(setData).toHaveBeenCalledTimes(15)

    it_store = update(it_store, {edit: {
      current: {$merge:{
        form_type: "invoice_link",
        invoice_link_fieldvalue: [],
        invoice_link: [
          {}
        ]
      }}
    }})
    await editorActions(it_store, setData).editItem({
      name: "ref_id_1", value: 1, item: { curr: "EUR" }
    })
    expect(setData).toHaveBeenCalledTimes(16)

    await editorActions(it_store, setData).editItem({
      name: "link_qty", value: 1
    })
    expect(setData).toHaveBeenCalledTimes(17)

    await editorActions(it_store, setData).editItem({
      name: "ref_id_2", value: 1
    })
    expect(setData).toHaveBeenCalledTimes(18)

    it_store = update(it_store, {edit: {
      current: {$merge:{
        form_type: "payment_link",
        payment_link_fieldvalue: [],
        payment_link: [
          {}
        ]
      }}
    }})
    await editorActions(it_store, setData).editItem({
      name: "ref_id_2", value: 1, item: { curr: "EUR" }
    })
    expect(setData).toHaveBeenCalledTimes(19)

    await editorActions(it_store, setData).editItem({
      name: "link_qty", value: 1
    })
    expect(setData).toHaveBeenCalledTimes(20)

    await editorActions(it_store, setData).editItem({
      name: "ref_id_1", value: 1
    })
    expect(setData).toHaveBeenCalledTimes(21)

    it_store = update(it_store, {edit: {
      current: {$merge:{
        form_type: "default",
      }}
    }})
    await editorActions(it_store, setData).editItem({
      name: "default", value: "default"
    })
    expect(setData).toHaveBeenCalledTimes(22)

    it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    await editorActions(it_store, setData).editItem({
      name: "paiddate", value: "2021-12-01"
    })
    expect(setData).toHaveBeenCalledTimes(23)

    await editorActions(it_store, setData).editItem({
      name: "closed", value: 1
    })
    expect(setData).toHaveBeenCalledTimes(29)

    await editorActions(it_store, setData).editItem({
      name: "closed", value: 0
    })
    expect(setData).toHaveBeenCalledTimes(30)

    await editorActions(it_store, setData).editItem({
      name: "direction", value: 69
    })
    expect(setData).toHaveBeenCalledTimes(31)

    it_store = update(it_store, {edit: {
      current: {$merge:{
        transtype: "cash",
      }}
    }})
    await editorActions(it_store, setData).editItem({
      name: "direction", value: 69
    })
    expect(setData).toHaveBeenCalledTimes(32)

    it_store = update(it_store, {edit: {
      current: {$merge:{
        transtype: "waybill",
        extend: {
          seltype: "customer"
        }
      }}
    }})
    await editorActions(it_store, setData).editItem({
      name: "seltype", value: 69
    })
    expect(setData).toHaveBeenCalledTimes(33)

    await editorActions(it_store, setData).editItem({
      name: "ref_id", value: 69, refnumber: "refnumber"
    })
    expect(setData).toHaveBeenCalledTimes(34)

    it_store = update(it_store, {edit: {
      current: {$merge:{
        extend: {
          seltype: "employee"
        }
      }}
    }})
    await editorActions(it_store, setData).editItem({
      name: "ref_id", value: 69, refnumber: "refnumber", item: { transtype: "order" }
    })
    expect(setData).toHaveBeenCalledTimes(35)

    it_store = update(it_store, {edit: {
      current: {$merge:{
        extend: {
          seltype: "transitem"
        }
      }}
    }})
    await editorActions(it_store, setData).editItem({
      name: "ref_id", value: 69, refnumber: "refnumber"
    })
    expect(setData).toHaveBeenCalledTimes(36)

    it_store = update(it_store, {edit: {
      current: {$merge:{
        extend: {
          seltype: ""
        }
      }}
    }})
    await editorActions(it_store, setData).editItem({
      name: "ref_id", value: 69, refnumber: "refnumber"
    })
    expect(setData).toHaveBeenCalledTimes(37)

    await editorActions(it_store, setData).editItem({
      name: "trans_wsdistance", value: 69
    })
    await editorActions(it_store, setData).editItem({
      name: "trans_wsrepair", value: 69
    })
    await editorActions(it_store, setData).editItem({
      name: "trans_wstotal", value: 69
    })
    await editorActions(it_store, setData).editItem({
      name: "trans_reholiday", value: 69
    })
    await editorActions(it_store, setData).editItem({
      name: "trans_rebadtool", value: 69
    })
    await editorActions(it_store, setData).editItem({
      name: "trans_reother", value: 69
    })
    await editorActions(it_store, setData).editItem({
      name: "trans_wsnote", value: 69
    })
    await editorActions(it_store, setData).editItem({
      name: "trans_rentnote", value: 69
    })
    expect(setData).toHaveBeenCalledTimes(45)

    await editorActions(it_store, setData).editItem({
      name: "shippingdate", value: "shippingdate"
    })
    await editorActions(it_store, setData).editItem({
      name: "shipping_place_id", value: 1
    })
    expect(setData).toHaveBeenCalledTimes(47)

    await editorActions(it_store, setData).editItem({
      name: "fnote", value: "", extend: true
    })
    expect(setData).toHaveBeenCalledTimes(48)

    await editorActions(it_store, setData).editItem({
      name: "default", value: ""
    })
    expect(setData).toHaveBeenCalledTimes(49)

    it_store = update(it_store, {edit: {
      current: {$merge:{
        type: "default"
      }}
    }})
    await editorActions(it_store, setData).editItem({
      name: "default", value: ""
    })
    expect(setData).toHaveBeenCalledTimes(50)

    it_store = update(it_store, {edit: {
      current: {$merge:{
        type: "printqueue",
      }}
    }})
    it_store = update(it_store, {
      edit: {$merge:{
        printqueue: {},
        audit: "readonly"
      }}
    })
    await editorActions(it_store, setData).editItem({
      name: "name", value: "value"
    })
    expect(setData).toHaveBeenCalledTimes(51)

    it_store = update(it_store, {edit: {
      current: {$merge:{
        type: "report",
      }}
    }})
    await editorActions(it_store, setData).editItem({
      name: "selected", value: 1
    })
    expect(setData).toHaveBeenCalledTimes(52)

    await editorActions(it_store, setData).editItem({
      name: "selected", value: 1, id: 100
    })
    expect(setData).toHaveBeenCalledTimes(53)

    await editorActions(it_store, setData).editItem({
      name: "name", value: 1
    })
    expect(setData).toHaveBeenCalledTimes(54)

    await editorActions(it_store, setData).editItem({
      name: "trans_transcast", value: 1
    })
    expect(setData).toHaveBeenCalledTimes(55)

    await editorActions(it_store, setData).editItem({
      name: "customer_id", value: 69, refnumber: "refnumber", extend: false, label_field: "customer"
    })
    expect(setData).toHaveBeenCalledTimes(56)

    await editorActions(it_store, setData).editItem({
      name: "customer_id", value: 69, extend: false, label_field: "customer"
    })
    expect(setData).toHaveBeenCalledTimes(57)

    await editorActions(it_store, setData).editItem({
      name: "customer_id", value: 69, extend: false
    })
    expect(setData).toHaveBeenCalledTimes(58)

    it_store = update(it_store, {edit: {
      template: {options: {$merge:{
        extend: true
      }}}
    }})
    await editorActions(it_store, setData).editItem({
      name: "customer_id", value: 69, extend: true
    })
    expect(setData).toHaveBeenCalledTimes(59)

    await editorActions(it_store, setData).editItem({
      name: "fieldvalue_value", value: "value", id: 123
    })
    expect(setData).toHaveBeenCalledTimes(60)

    it_store = update(it_store, {
      edit: {$merge:{
        audit: "all"
      }}
    })
    await editorActions(it_store, setData).editItem({
      name: "fieldvalue_value", value: "value", id: 59
    })
    expect(setData).toHaveBeenCalledTimes(61)

    await editorActions(it_store, setData).editItem({
      name: "fieldvalue_deleted", value: true, id: 59
    })
    expect(setData).toHaveBeenCalledTimes(62)

  })

  it('setPattern', () => {
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/pattern"){
          return [1]
        }
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    let setData = jest.fn((key, data, callback)=>{
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(data.modalForm, container);

        const input_value = getById(container, 'input_value')
        if(input_value){
          fireEvent.change(input_value, {target: {value: "value"}})
        }

        // onOK
        const btn_ok = getById(container, 'btn_ok')
        if(btn_ok){
          btn_ok.dispatchEvent(new MouseEvent('click', {bubbles: true}));
        }

        // onCancel
        const btn_cancel = getById(container, 'btn_cancel')
        btn_cancel.dispatchEvent(new MouseEvent('click', {bubbles: true}));
      }
      if(callback){callback()}
    })
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    it_store = update(it_store, {edit: {
      dataset: {$merge: {
        pattern: [
          { id: 1, description: "first pattern", transtype: 55, 
            notes: "pattern text", defpattern: 0, deleted: 0 },
          { id: 2, description: "default pattern", transtype: 55, 
            notes: null, defpattern: 1, deleted: 0 },
        ]
      }}
    }})
    editorActions(it_store, setData).setPattern({ key: "default" })
    expect(setData).toHaveBeenCalledTimes(3)

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    editorActions(it_store, setData).setPattern({ key: "save" })
    expect(setData).toHaveBeenCalledTimes(6)

    editorActions(it_store, setData).setPattern({ key: "delete" })
    expect(setData).toHaveBeenCalledTimes(9)

    editorActions(it_store, setData).setPattern({ key: "load" })
    expect(setData).toHaveBeenCalledTimes(10)

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/pattern"){
          return []
        }
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    editorActions(it_store, setData).setPattern({ key: "new" })
    expect(setData).toHaveBeenCalledTimes(13)

    it_store = update(it_store, {edit: {
      dataset: {$merge: {
        pattern: []
      }}
    }})
    editorActions(it_store, setData).setPattern({ key: "new" })
    expect(setData).toHaveBeenCalledTimes(16)

    editorActions(it_store, setData).setPattern({ key: "load" })
    expect(setData).toHaveBeenCalledTimes(16)

    editorActions(it_store, setData).setPattern({ key: "save" })
    expect(setData).toHaveBeenCalledTimes(19)

    editorActions(it_store, setData).setPattern({ key: "delete" })
    expect(setData).toHaveBeenCalledTimes(22)

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/pattern"){
          return [1]
        }
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    editorActions(it_store, setData).setPattern({ key: "new" })
    expect(setData).toHaveBeenCalledTimes(25)

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    editorActions(it_store, setData).setPattern({ key: "new" })
    expect(setData).toHaveBeenCalledTimes(28)

    setData = jest.fn((key, data, callback)=>{
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(data.modalForm, container);

        // onOK
        const btn_ok = getById(container, 'btn_ok')
        if(btn_ok){
          btn_ok.dispatchEvent(new MouseEvent('click', {bubbles: true}));
        }
      }
      if(callback){callback()}
    })
    editorActions(it_store, setData).setPattern({ key: "new" })
    expect(setData).toHaveBeenCalledTimes(2)

    it_store = update(it_store, {edit: {
      current: {$merge:{
        template: ""
      }}
    }})
    editorActions(it_store, setData).setPattern({ key: "default" })
    expect(setData).toHaveBeenCalledTimes(2)

  })

  it('shippingAddAll', () => {
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
      getAuditFilter: jest.fn(() => (["disabled",1])),
    })
    const setData = jest.fn()
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    it_store = update(it_store, {edit: {
      dataset: {$merge: {
        shiptemp: [],
        shipping_items_: [
          { description: "Wheel", item_id: 1, partname: "Wheel", partnumber: "DMPROD/00005",
            pgroup: "0", product: "DMPROD/00005 | Wheel", product_id: "5", 
            qty: "40", unit: "piece", tqty: 1 },
          { description: "Wheel", item_id: 1, partname: "Wheel", partnumber: "DMPROD/00005",
            pgroup: "0", product: "DMPROD/00005 | Wheel", product_id: "5", 
            qty: "40", unit: "piece", tqty: 1, edited: true, diff: 0 }
        ]
      }}
    }})
    editorActions(it_store, setData).shippingAddAll()
    expect(setData).toHaveBeenCalledTimes(0)

  })

  it('setPassword', () => {
    const setData = jest.fn()
    const it_store = update(store, {edit: {
      current: {$merge: {
        type: "employee"
      }},
      dataset: {$merge: {
        employee: [
          { username: "admin" }
        ],
      }}
    }})
    editorActions(it_store, setData).setPassword("admin")
    expect(setData).toHaveBeenCalledTimes(1)

    editorActions(it_store, setData).setPassword()
    expect(setData).toHaveBeenCalledTimes(2)
  })

  it('setLink', () => {
    let setData = jest.fn((key, data, callback)=>{
      if((key === "current") && data.side ){
        if(callback){
          callback()
        }
      }
    })
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
      getAuditFilter: jest.fn(() => (["disabled",1])),
    })
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    it_store = update(it_store, {edit: {$merge:{
      panel: it_store.edit.template.options.panel
    }}})
    it_store = update(it_store, {edit: {
      current: {$merge: {
        transtype: "cash",
        extend: {
          id: 1
        },
        form: {
          id: 1
        }
      }},
      dataset: {$merge:{
        payment_link: [],
      }}
    }})
    editorActions(it_store, setData).setLink("payment_link", "ref_id_1")
    expect(setData).toHaveBeenCalledTimes(2)

    it_store = update(it_store, {edit: {
      current: {$merge: {
        transtype: "invoice",
      }},
    }})
    editorActions(it_store, setData).setLink("payment_link", "ref_id_1")
    expect(setData).toHaveBeenCalledTimes(4)

  })

  it('editorBack', () => {
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
      getAuditFilter: jest.fn(() => (["disabled",1])),
    })
    const setData = jest.fn()
    let it_store = update(store, {edit: {$merge:{
      ...Item.args
    }}})
    it_store = update(it_store, {login: {data: {$merge:{
      groups: it_store.edit.dataset.groups
    }}}})
    editorActions(it_store, setData).editorBack()
    expect(setData).toHaveBeenCalledTimes(0)

    it_store = update(it_store, {edit: {current: {$merge:{
      form: undefined,
      item: {
        nervatype: 10,
        ref_id: 2
      }
    }}}})
    editorActions(it_store, setData).editorBack()
    expect(setData).toHaveBeenCalledTimes(0)

    it_store = update(it_store, {edit: {current: {$merge:{
      form_type: "transitem_shipping",
      type: "trans",
      transtype: "order"
    }}}})
    editorActions(it_store, setData).editorBack()
    expect(setData).toHaveBeenCalledTimes(0)
    
  })

  it('editorSave', () => {
    InitItem.mockReturnValue(jest.fn( () => {
      return update(InvoiceData.args.current.item, {$merge: {
        id: null
      }})
    } ))
    Validator.mockReturnValue(jest.fn( () => {
      return update(InvoiceData.args.current.item, {$merge: {
        id: null
      }})
    } ))
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/view"){
          return { fields: [{ 
            trans_custinvoice_compname: "compname", trans_custinvoice_custname: "custname"
          }] }
        }
        return [1,2]
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    const setData = jest.fn()
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    editorActions(it_store, setData).editorSave()
    expect(setData).toHaveBeenCalledTimes(0)

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    it_store = update(store, {edit: {$merge:{
      ...Item.args
    }}})
    editorActions(it_store, setData).editorSave()
    expect(setData).toHaveBeenCalledTimes(0)

  })

  it('editorDelete', () => {
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    const setData = jest.fn()
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    editorActions(it_store, setData).editorDelete()
    expect(setData).toHaveBeenCalledTimes(1)

    it_store = update(store, {edit: {$merge:{
      ...Item.args
    }}})
    editorActions(it_store, setData).editorDelete()
    expect(setData).toHaveBeenCalledTimes(2)

  })

  it('editorNew', () => {
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "trans/order/1"})
        }
      }),
    })
    const setData = jest.fn()
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    editorActions(it_store, setData).editorNew({ntype: "trans", ttype: "invoice"})
    expect(setData).toHaveBeenCalledTimes(0)

    editorActions(it_store, setData).editorNew({})
    expect(setData).toHaveBeenCalledTimes(0)

    editorActions(it_store, setData).editorNew({ ttype: "shipping" })
    expect(setData).toHaveBeenCalledTimes(0)

  })

  it('transCopy', () => {
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path, options) => {
        if(path === "/pattern"){
          return [1]
        }
        return { error: {} }
      }),
      resultError: jest.fn(),
      showToast: jest.fn(),
      createHistory: jest.fn(),
      onSelector: jest.fn((selectorType, selectorFilter, setSelector) => {
        if(setSelector){
          setSelector({id: "1/1/1"})
        }
      }),
    })
    let setData = jest.fn((key, data, callback)=>{
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(data.modalForm, container);

        const input_value = getById(container, 'input_value')
        if(input_value){
          fireEvent.change(input_value, {target: {value: "value"}})
        }

        // onOK
        const btn_ok = getById(container, 'btn_ok')
        if(btn_ok){
          btn_ok.dispatchEvent(new MouseEvent('click', {bubbles: true}));
        }

        // onCancel
        const btn_cancel = getById(container, 'btn_cancel')
        btn_cancel.dispatchEvent(new MouseEvent('click', {bubbles: true}));
      }
      if(callback){callback()}
    })
    let it_store = update(store, {edit: {$merge:{
      ...InvoiceData.args
    }}})
    editorActions(it_store, setData).transCopy()
    expect(setData).toHaveBeenCalledTimes(3)

    it_store = update(it_store, {edit: {
      dataset: {$merge:{
        element_count: [
          { pec: 0 }
        ]
      }}
    }})
    editorActions(it_store, setData).transCopy("create")
    expect(setData).toHaveBeenCalledTimes(5)

  })

})