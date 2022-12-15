import sinon from 'sinon'
import { expect } from '@open-wc/testing';

import { EditController, checkSubtype } from './EditController.js'
import { store as storeConfig } from '../config/app.js'
import { SIDE_EVENT, APP_MODULE, EDIT_EVENT, MODAL_EVENT, EDITOR_EVENT, ACTION_EVENT } from '../config/enums.js'
import { Forms } from './Forms.js'
import { Quick } from './Quick.js'
import { Sql } from './Sql.js'
import { Dataset } from './Dataset.js'
import { Default as InvoiceData, Item, PrintQueue, Report } from '../components/Edit/Editor/Editor.stories.js'

const host = { 
  addController: ()=>{},
  inputBox: (prm)=>(prm)
}
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
          { id: 1, groupname: "nervatype", groupvalue: "type" }
        ],
        audit: [
          { nervatypeName: "report", subtype: 1, inputfilterName: "disabled" }
        ]
      }
    }
  },
  setData: sinon.spy(),
}
const app = {
  store,
  modules: {
    forms: Forms({ msg: (key)=> key }),
    sql: Sql({ msg: (key)=> key }),
    quick: {...Quick},
    dataSet: {...Dataset}
  },
  msg: (value)=>value,
  requestData: () => ({}),
  resultError: sinon.spy(),
  loadBookmark: () => ({}),
  showHelp: sinon.spy(),
  showToast: sinon.spy(),
  getSetting: sinon.spy(),
  getSql: () =>({
    sql: "",
    prmCount: 1
  }),
  getDataFilter: ()=>[[]],
  getUserFilter: ()=>({ where: [[]], params: [[]] }),
  currentModule: sinon.spy(),
}

describe('EditController', () => {
  beforeEach(async () => {
    // global.URL.createObjectURL = sinon.spy();
    Object.defineProperty(URL, 'createObjectURL', { value: sinon.spy() })
    Object.defineProperty(window, 'open', { value: sinon.spy() })
  });
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('onSideEvent', async () => {
    let testApp = {
      ...app,
      showHelp: sinon.spy(),
      saveBookmark: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            current: {
              ...app.store.data[APP_MODULE.EDIT].current,
              item: {
                id: 1,
                nervatype: 1
              }
            }
          }
        }
      },
    }

    let edit = new EditController({...host, app: testApp})
    edit.onSideEvent({ key: SIDE_EVENT.CHANGE, data: { fieldname: "", value: "" } })
    sinon.assert.callCount(testApp.store.setData, 1);

    edit.checkEditor = sinon.spy()
    edit.onSideEvent({ key: SIDE_EVENT.BACK, data: {} })
    sinon.assert.callCount(edit.checkEditor, 1);

    edit.onSideEvent({ key: SIDE_EVENT.CHECK, data: [] })
    sinon.assert.callCount(edit.checkEditor, 2);

    edit.prevTransNumber = sinon.spy()
    edit.onSideEvent({ key: SIDE_EVENT.PREV_NUMBER, data: {} })
    sinon.assert.callCount(edit.prevTransNumber, 1);

    edit.nextTransNumber = sinon.spy()
    edit.onSideEvent({ key: SIDE_EVENT.NEXT_NUMBER, data: {} })
    sinon.assert.callCount(edit.nextTransNumber, 1);

    edit.saveEditor = async () => ({ current: { item: {} } })
    edit.loadEditor = sinon.spy()
    await edit.onSideEvent({ key: SIDE_EVENT.SAVE, data: {} })
    sinon.assert.callCount(edit.loadEditor, 1);

    edit.deleteEditor = sinon.spy()
    edit.onSideEvent({ key: SIDE_EVENT.DELETE, data: {} })
    sinon.assert.callCount(edit.deleteEditor, 1);

    edit.onSideEvent({ key: SIDE_EVENT.NEW, data: [{}] })
    sinon.assert.callCount(edit.checkEditor, 3);

    edit.onSelector = (type, filter, callback)=>{ callback({ id: "ntype/ttype/1" }) }
    edit.onSideEvent({ key: SIDE_EVENT.NEW, data: [{ ttype: "shipping" }] })
    sinon.assert.callCount(edit.checkEditor, 4);

    edit.transCopy = sinon.spy()
    edit.onSideEvent({ key: SIDE_EVENT.COPY, data: {} })
    sinon.assert.callCount(edit.transCopy, 1);

    edit.setLink = sinon.spy()
    edit.onSideEvent({ key: SIDE_EVENT.LINK, data: {} })
    sinon.assert.callCount(edit.setLink, 1);

    edit.setPassword = sinon.spy()
    edit.onSideEvent({ key: SIDE_EVENT.PASSWORD, data: {} })
    sinon.assert.callCount(edit.setPassword, 1);

    edit.shippingAddAll = sinon.spy()
    edit.onSideEvent({ key: SIDE_EVENT.SHIPPING_ADD_ALL, data: {} })
    sinon.assert.callCount(edit.shippingAddAll, 1);

    edit.createShipping = sinon.spy()
    edit.onSideEvent({ key: SIDE_EVENT.SHIPPING_CREATE, data: {} })
    sinon.assert.callCount(edit.createShipping, 1);

    edit.reportSettings = sinon.spy()
    edit.onSideEvent({ key: SIDE_EVENT.REPORT_SETTINGS, data: {} })
    sinon.assert.callCount(edit.reportSettings, 1);

    edit.searchQueue = sinon.spy()
    edit.onSideEvent({ key: SIDE_EVENT.SEARCH_QUEUE, data: {} })
    sinon.assert.callCount(edit.searchQueue, 1);

    edit.exportQueueAll = sinon.spy()
    edit.onSideEvent({ key: SIDE_EVENT.EXPORT_QUEUE_ALL, data: {} })
    sinon.assert.callCount(edit.exportQueueAll, 1);

    edit.createReport = sinon.spy()
    edit.onSideEvent({ key: SIDE_EVENT.CREATE_REPORT, data: {} })
    sinon.assert.callCount(edit.createReport, 1);

    edit.exportEvent = sinon.spy()
    edit.onSideEvent({ key: SIDE_EVENT.EXPORT_EVENT, data: {} })
    sinon.assert.callCount(edit.exportEvent, 1);

    edit.saveBookmark = sinon.spy()
    edit.onSideEvent({ key: SIDE_EVENT.SAVE_BOOKMARK, data: { value: [] } })
    sinon.assert.callCount(testApp.saveBookmark, 1);

    edit.showHelp = sinon.spy()
    edit.onSideEvent({ key: SIDE_EVENT.HELP, data: {} })
    sinon.assert.callCount(testApp.showHelp, 1);

    edit.onSideEvent({ key: "", data: {} })

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              form_type: "transitem_shipping"
            }
          }
        }
      }
    }

    edit = new EditController({...host, app: testApp})
    edit.checkEditor = sinon.spy()
    edit.onSideEvent({ key: SIDE_EVENT.BACK, data: {} })
    sinon.assert.callCount(edit.checkEditor, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              form: {}
            }
          }
        }
      }
    }

    edit = new EditController({...host, app: testApp})
    edit.checkEditor = sinon.spy()
    edit.onSideEvent({ key: SIDE_EVENT.BACK, data: {} })
    sinon.assert.callCount(edit.checkEditor, 1);

    edit.saveEditorForm = sinon.spy(async () => (null))
    await edit.onSideEvent({ key: SIDE_EVENT.SAVE, data: {} })
    sinon.assert.callCount(edit.saveEditorForm, 1);

    edit.deleteEditorItem = sinon.spy(async () => (null))
    edit.onSideEvent({ key: SIDE_EVENT.DELETE, data: {} })
    sinon.assert.callCount(edit.deleteEditorItem, 1);

  })

  it('onEditEvent', () => {
    const testApp = {
      ...app,
      showHelp: sinon.spy(),
      saveBookmark: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy(),
      }
    }

    const edit = new EditController({...host, app: testApp})
    edit.onEditEvent({ key: EDIT_EVENT.CHANGE, data: { fieldname: "fieldname", value: "value" } })
    sinon.assert.callCount(testApp.store.setData, 1);

    edit.checkEditor = sinon.spy()
    edit.onEditEvent({ key: EDIT_EVENT.CHECK_EDITOR, data: [] })
    sinon.assert.callCount(edit.checkEditor, 1);

    edit.checkTranstype = sinon.spy()
    edit.onEditEvent({ key: EDIT_EVENT.CHECK_TRANSTYPE, data: [] })
    sinon.assert.callCount(edit.checkTranstype, 1);

    edit.editItem = sinon.spy()
    edit.onEditEvent({ key: EDIT_EVENT.EDIT_ITEM, data: [] })
    sinon.assert.callCount(edit.editItem, 1);

    edit.setPattern = sinon.spy()
    edit.onEditEvent({ key: EDIT_EVENT.SET_PATTERN, data: [] })
    sinon.assert.callCount(edit.setPattern, 1);

    edit.onSelector = sinon.spy()
    edit.onEditEvent({ key: EDIT_EVENT.SELECTOR, data: [] })
    sinon.assert.callCount(edit.onSelector, 1);

    edit.setFormActions = sinon.spy()
    edit.onEditEvent({ key: EDIT_EVENT.FORM_ACTION, data: [] })
    sinon.assert.callCount(edit.setFormActions, 1);

    edit.onEditEvent({ key: "", data: [] })

  })

  it('onSelector', async () => {
    let testApp = {
      ...app,
      resultError: sinon.spy(),
      quickSearch: sinon.spy(async ()=>({ error: {} })),
      store: {
        ...app.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.SEARCH, data: {} })
          }
        })
      }
    }

    const setSelector = sinon.spy()
    let edit = new EditController({...host, app: testApp})
    edit.module.modalSelector = (prm)=>(prm)
    edit.onSelector("customer", "", setSelector)
    sinon.assert.callCount(testApp.store.setData, 2);

    let modalForm = 1
    testApp = {
      ...testApp,
      quickSearch: sinon.spy(async ()=>({ result: {} })),
      store: {
        ...app.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm && (modalForm === 1)){
            modalForm += 1
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.SELECTED, data: {} })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.SEARCH, data: {} })
          }
        })
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.module.modalSelector = (prm)=>(prm)
    await edit.onSelector("customer", "", setSelector)
    sinon.assert.callCount(setSelector, 1);

  })

  it('addPrintQueue', async () => {
    let testApp = {
      ...app,
      resultError: sinon.spy(),
      showToast: sinon.spy(),
      requestData: sinon.spy(async ()=>({ error: {} })),
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            current: {
              ...app.store.data[APP_MODULE.EDIT].current,
              type: "type",
              item: {
                id: 1
              }
            },
            dataset: {
              ...app.store.data[APP_MODULE.EDIT].dataset,
              report: [
                { id: 1, reportkey: "reportkey" }
              ]
            }
          }
        }
      }
    }

    let edit = new EditController({...host, app: testApp})
    await edit.addPrintQueue("reportkey", 1)
    sinon.assert.callCount(testApp.resultError, 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async ()=>({ })),
    }

    edit = new EditController({...host, app: testApp})
    await edit.addPrintQueue("reportkey", 1)
    sinon.assert.callCount(testApp.showToast, 1);
  })

  it('calcFormula', async () => {
    let testApp = {
      ...app,
      modules: {
        ...app.modules,
        initItem: ()=>({})
      },
      resultError: sinon.spy(),
      showToast: sinon.spy(),
      requestData: sinon.spy(async (path, options) => {
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
      store: {
        ...app.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL, data: {} })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: {} })
          }
        }),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }
    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
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
            }
          }
        }
      }
    }

    let edit = new EditController({...host, app: testApp})
    edit.loadEditor = sinon.spy()
    await edit.calcFormula(18)
    sinon.assert.callCount(testApp.store.setData, 3);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path, options) => {
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
      })
    }
    edit = new EditController({...host, app: testApp})
    await edit.calcFormula(18)
    sinon.assert.callCount(testApp.store.setData, 6);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
        if(path === "/view"){
          return { formula: [
            { deleted: 0, id: 31, movetype: 91, notes: null, place_id: 4, product_id: 5,
              qty: 20, shared: 0, shippingdate: "2021-12-01T00:00:00", tool_id: null, trans_id: 18 },
            { deleted: 0, id: 31, movetype: 91, notes: null, place_id: null, product_id: 5,
              qty: 20, shared: 1, shippingdate: "2021-12-01T00:00:00", tool_id: null, trans_id: 18 }
          ] }
        }
        return { error: {} }
      })
    }
    edit = new EditController({...host, app: testApp})
    await edit.calcFormula(18)
    sinon.assert.callCount(testApp.store.setData, 9);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} }))
    }
    edit = new EditController({...host, app: testApp})
    await edit.calcFormula(18)
    sinon.assert.callCount(testApp.store.setData, 12);

  })

  it('calcPrice', async () => {
    let testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }
    
    let edit = new EditController({...host, app: testApp})
    let item = edit.calcPrice("netamount",
      { netamount: 10, qty: 1, discount: 0 }
    )
    expect(item.amount).to.equal(10)

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              item: {
                ...testApp.store.data[APP_MODULE.EDIT].current.item,
                curr: "MTG"
              }
            }
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    item = edit.calcPrice("netamount",
      { netamount: 10, qty: 0, discount: 0 }
    )
    expect(item.amount).to.equal(10)

    item = edit.calcPrice("amount",
      { amount: 10, qty: 1, discount: 0, tax_id: 2 }
    )
    expect(item.netamount).to.equal(9.52)

    item = edit.calcPrice("amount",
      { amount: 10, qty: 0, discount: 0, tax_id: 2 }
    )
    expect(item.netamount).to.equal(0)

    item = edit.calcPrice("fxprice",
      { fxprice: 10, qty: 1, discount: 0 }
    )
    expect(item.amount).to.equal(10)

    item = edit.calcPrice("",
      { fxprice: 10, qty: 1, discount: 0 }
    )
    expect(item.amount).to.equal(10)

  })

  it('checkEditor', async () => {
    let testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: {} })
          }
        }),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }
    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              movement_head: [
                { description: 'Car', id: 40, partnumber: 'DMPROD/00004'} ],
              formula_head: [
                { id: 18, qty: 5, transnumber: 'DMFRM/00001' },
                { id: 19, qty: 5, transnumber: 'DMFRM/00002' } ],
            }
          }
        }
      }
    }

    let edit = new EditController({...host, app: testApp})
    edit.module.modalFormula = (prm)=>(prm)
    edit.calcFormula = sinon.spy()
    edit.checkEditor({}, EDITOR_EVENT.LOAD_FORMULA)
    sinon.assert.callCount(edit.calcFormula, 1);

    edit.loadEditor = sinon.spy()
    edit.checkEditor({}, EDITOR_EVENT.LOAD_EDITOR)
    sinon.assert.callCount(edit.loadEditor, 1);

    edit.setEditorItem = sinon.spy()
    edit.checkEditor({}, EDITOR_EVENT.SET_EDITOR_ITEM)
    sinon.assert.callCount(edit.setEditorItem, 1);

    edit.newFieldvalue = sinon.spy()
    edit.checkEditor({}, EDITOR_EVENT.NEW_FIELDVALUE)
    sinon.assert.callCount(edit.newFieldvalue, 1);

    edit.createTrans = sinon.spy()
    edit.checkEditor({}, EDITOR_EVENT.CREATE_TRANS)
    sinon.assert.callCount(edit.createTrans, 1);

    edit.createTransOptions = sinon.spy()
    edit.checkEditor({}, EDITOR_EVENT.CREATE_TRANS_OPTIONS)
    sinon.assert.callCount(edit.createTransOptions, 1);

    edit.setFormActions = sinon.spy()
    edit.checkEditor({}, EDITOR_EVENT.FORM_ACTIONS)
    sinon.assert.callCount(edit.setFormActions, 1);

    edit.checkEditor({}, "")

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dirty: true,
          }
        }
      }
    }

    edit = new EditController({...host, app: testApp})
    edit.saveEditor = sinon.spy(async () => ({}))
    edit.loadEditor = sinon.spy()
    edit.checkEditor({}, EDITOR_EVENT.LOAD_EDITOR)
    sinon.assert.callCount(edit.loadEditor, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            form_dirty: true,
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              form: {},
              item: null
            }
          }
        }
      }
    }

    edit = new EditController({...host, app: testApp})
    edit.saveEditorForm = sinon.spy(async () => (null))
    edit.loadEditor = sinon.spy()
    edit.checkEditor({}, EDITOR_EVENT.LOAD_EDITOR)
    sinon.assert.callCount(edit.loadEditor, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            form_dirty: false,
            dirty: false
          }
        }
      }
    }

    edit = new EditController({...host, app: testApp})
    edit.saveEditorForm = sinon.spy(async () => ({}))
    edit.loadEditor = sinon.spy()
    edit.checkEditor({}, EDITOR_EVENT.LOAD_EDITOR)
    sinon.assert.callCount(edit.loadEditor, 1);
  
  })

  it('checkTranstype', async () => {
    let testApp = {
      ...app,
      resultError: sinon.spy(),
      requestData: sinon.spy(async ()=>({ error: {} })),
    }

    let edit = new EditController({...host, app: testApp})
    edit.checkEditor = sinon.spy()
    await edit.checkTranstype({}, EDITOR_EVENT.LOAD_EDITOR)
    sinon.assert.callCount(edit.checkEditor , 1);

    await edit.checkTranstype({ ntype: "trans", ttype: null }, EDITOR_EVENT.LOAD_EDITOR)
    sinon.assert.callCount(testApp.resultError , 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async ()=>({ transtype: [{ groupvalue: "groupvalue" }] })),
    }
    edit = new EditController({...host, app: testApp})
    edit.checkEditor = sinon.spy()
    await edit.checkTranstype({ ntype: "trans", ttype: null }, EDITOR_EVENT.LOAD_EDITOR)
    sinon.assert.callCount(edit.checkEditor , 1);

  })

  it('createReport', async () => {
    let testApp = {
      ...app,
      resultError: sinon.spy(),
      requestData: sinon.spy(async ()=>({})),
      saveToDisk: sinon.spy(),
      store: {
        ...app.store,
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            current: {
              ...app.store.data[APP_MODULE.EDIT].current,
              item: {},
              fieldvalue: [
                { datatype: "date", empty: "false", id: 1, label: "Date",
                  name: "posdate", rowtype: "reportfield", selected: true, value: "2022-03-14" 
                }
              ]
            }
          }
        }
      }
    }

    let edit = new EditController({...host, app: testApp})
    await edit.createReport("xml")
    sinon.assert.callCount(testApp.saveToDisk , 1);

    await edit.createReport("csv")
    sinon.assert.callCount(testApp.saveToDisk , 2);

    await edit.createReport("print")
    sinon.assert.callCount(testApp.saveToDisk , 2);


    testApp = {
      ...testApp,
      resultError: sinon.spy(),
      requestData: sinon.spy(async ()=>({ error: {} }))
    }
    edit = new EditController({...host, app: testApp})
    await edit.createReport("xml")
    sinon.assert.callCount(testApp.resultError , 1);
  })

  it('createShipping', async () => {
    let testApp = {
      ...app,
      modules: {
        ...app.modules,
        initItem: ()=>({})
      },
      requestData: sinon.spy(async ()=>([])),
      resultError: sinon.spy(),
      showToast: sinon.spy(), 
      createHistory: sinon.spy(),
      store: {
        ...app.store,
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              shipping_place_id: null
            },
          }
        }
      }
    }

    let edit = new EditController({...host, app: testApp})
    await edit.createShipping()
    sinon.assert.callCount(testApp.showToast , 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              shipping_place_id: 1,
              shippingdate: "2021-12-01T00:00:00",
              direction: "out",
              item: {
                ...testApp.store.data[APP_MODULE.EDIT].current.item,
                shippingdate: "2021-12-01T00:00:00"
              }
            },
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              shiptemp: [
                { id: "1-5", item_product: "Wheel", movement_product: "Wheel", product_id: 5, sqty: "30" }
              ],
              delivery_pattern: [
                { defpattern: 1, notes: "" }
              ]
            }
          }
        }
      }
    }

    edit = new EditController({...host, app: testApp})
    edit.loadEditor = sinon.spy()
    await edit.createShipping()
    sinon.assert.callCount(edit.loadEditor , 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
        if(path === "/link"){
          return { error: {} }
        }
        return [1]
      }),
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              direction: "in"
            },
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              shiptemp: [
                { id: "1-5", item_product: "Wheel", movement_product: "Wheel", product_id: 5, sqty: "30" }
              ],
              delivery_pattern: []
            }
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.createShipping()
    sinon.assert.callCount(testApp.resultError , 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
        if(path === "/movement"){
          return { error: {} }
        }
        return [1]
      }),
    }
    edit = new EditController({...host, app: testApp})
    await edit.createShipping()
    sinon.assert.callCount(testApp.resultError , 2);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
        if(path === "/trans"){
          return { error: {} }
        }
        return [1]
      }),
    }
    edit = new EditController({...host, app: testApp})
    await edit.createShipping()
    sinon.assert.callCount(testApp.resultError , 3);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
        if(path === "/function"){
          return { error: {} }
        }
        return [1]
      }),
    }
    edit = new EditController({...host, app: testApp})
    await edit.createShipping()
    sinon.assert.callCount(testApp.resultError , 4);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              shiptemp: [],
              delivery_pattern: []
            }
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    const result = await edit.createShipping()
    expect(result).to.true

  })

  it('createTrans', async () => {
    let testApp = {
      ...app,
      modules: {
        ...app.modules,
        initItem: ()=>({})
      },
      requestData: sinon.spy(async (path) => {
        if(path === "/view"){
          return { fields: [{ 
            trans_custinvoice_compname: "compname", trans_custinvoice_custname: "custname"
          }] }
        }
        return [1,2]
      }),
      resultError: sinon.spy(),
      showToast: sinon.spy(), 
      createHistory: sinon.spy(),
      store: {
        ...app.store,
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }
    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              payment: []
            },
          }
        }
      }
    }

    let edit = new EditController({...host, app: testApp})
    edit.loadEditor = sinon.spy()
    await edit.createTrans({
      cmdtype: "copy", transcast: "normal",
    })
    sinon.assert.callCount(edit.loadEditor , 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              settings: [],
              trans: [...testApp.store.data[APP_MODULE.EDIT].dataset.trans]
            },
          }
        }
      }
    }
    testApp.store.data[APP_MODULE.EDIT].dataset.trans[0] = {
      ...testApp.store.data[APP_MODULE.EDIT].dataset.trans[0],
      duedate: null
    }

    edit = new EditController({...host, app: testApp})
    edit.loadEditor = sinon.spy()
    await edit.createTrans({
      cmdtype: "copy"
    })
    sinon.assert.callCount(edit.loadEditor , 1);

    testApp.store.data[APP_MODULE.EDIT].dataset.trans[0] = {
      ...testApp.store.data[APP_MODULE.EDIT].dataset.trans[0],
      transtype: 56
    }

    edit = new EditController({...host, app: testApp})
    edit.loadEditor = sinon.spy()
    await edit.createTrans({
      cmdtype: "copy", transcast: "normal",
    })
    sinon.assert.callCount(edit.loadEditor , 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
        if(path === "/function"){
          return { error: {} }
        }
        return [1,2]
      }),
    }
    edit = new EditController({...host, app: testApp})
    await edit.createTrans({
      cmdtype: "copy", transcast: "normal",
    })
    sinon.assert.callCount(testApp.resultError , 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
        if(path === "/trans"){
          return { error: {} }
        }
        return [1,2]
      }),
    }
    edit = new EditController({...host, app: testApp})
    await edit.createTrans({
      cmdtype: "copy", transcast: "normal",
    })
    sinon.assert.callCount(testApp.resultError , 2);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
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
    }
    edit = new EditController({...host, app: testApp})
    await edit.createTrans({
      cmdtype: "copy", transcast: "normal",
    })
    sinon.assert.callCount(testApp.resultError , 3);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
        if(path === "/view") { 
          return { fields: [] }
        }
        return [1,2]
      }),
    }
    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              fieldvalue: []
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.loadEditor = sinon.spy()
    await edit.createTrans({
      cmdtype: "copy", transcast: "normal",
    })
    sinon.assert.callCount(edit.loadEditor, 1);

    testApp.store.data[APP_MODULE.EDIT].dataset.trans[0] = {
      ...testApp.store.data[APP_MODULE.EDIT].dataset.trans[0],
      transcast: "cancellation"
    }
    edit = new EditController({...host, app: testApp})
    await edit.createTrans({
      cmdtype: "copy", transcast: "normal",
    })
    sinon.assert.callCount(testApp.showToast , 1);

    testApp.store.data[APP_MODULE.EDIT].dataset.trans[0] = {
      ...testApp.store.data[APP_MODULE.EDIT].dataset.trans[0],
      transtype: 56, direction: 69
    }
    edit = new EditController({...host, app: testApp})
    await edit.createTrans({
      cmdtype: "copy", transcast: "normal",
    })
    sinon.assert.callCount(testApp.showToast , 2);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
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
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }
    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              payment: [
                { id: 123, deleted: 0}
              ]
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.createTrans({
      cmdtype: "copy", transcast: "amendment"
    })
    sinon.assert.callCount(testApp.resultError, 4);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
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
    }
    edit = new EditController({...host, app: testApp})
    await edit.createTrans({
      cmdtype: "copy", 
    })
    sinon.assert.callCount(testApp.resultError , 5);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
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
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              trans: [...testApp.store.data[APP_MODULE.EDIT].dataset.trans]
            },
          }
        }
      }
    }
    testApp.store.data[APP_MODULE.EDIT].dataset.trans[0] = {
      ...testApp.store.data[APP_MODULE.EDIT].dataset.trans[0],
      transtype: 63
    }
    edit = new EditController({...host, app: testApp})
    await edit.createTrans({
      cmdtype: "copy", 
    })
    sinon.assert.callCount(testApp.resultError , 6);
    
    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
        if(path === "/view") { 
          return { fields: [{ 
            trans_custinvoice_compname: "compname", trans_custinvoice_custname: "custname"
          }] }
        }
        return [1,2]
      }),
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.createTrans({
      cmdtype: "copy", transcast: "cancellation" 
    })
    sinon.assert.callCount(testApp.showToast , 3);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              trans: [...testApp.store.data[APP_MODULE.EDIT].dataset.trans]
            },
          }
        }
      }
    }
    testApp.store.data[APP_MODULE.EDIT].dataset.trans[0] = {
      ...testApp.store.data[APP_MODULE.EDIT].dataset.trans[0],
      deleted: 1
    }
    edit = new EditController({...host, app: testApp})
    await edit.createTrans({
      cmdtype: "copy", transcast: "amendment" 
    })
    sinon.assert.callCount(testApp.showToast , 4);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              cancel_link: [{ transnumber: "transnumber" }]
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.createTrans({
      cmdtype: "copy", transcast: "cancellation" 
    })
    sinon.assert.callCount(testApp.showToast, 5);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
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
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.loadEditor = sinon.spy()
    await edit.createTrans({
      cmdtype: "copy", transcast: "cancellation" 
    })
    sinon.assert.callCount(edit.loadEditor, 1);

    testApp.store.data[APP_MODULE.EDIT].dataset.trans[0] = {
      ...testApp.store.data[APP_MODULE.EDIT].dataset.trans[0],
      transtype: 61
    }
    edit = new EditController({...host, app: testApp})
    edit.loadEditor = sinon.spy()
    await edit.createTrans({
      cmdtype: "copy", transcast: "cancellation" 
    })
    sinon.assert.callCount(edit.loadEditor, 1);

    await edit.createTrans({
      cmdtype: "copy", transcast: "normal" 
    })
    sinon.assert.callCount(edit.loadEditor, 2);

    testApp.store.data[APP_MODULE.EDIT].dataset.trans[0] = {
      ...testApp.store.data[APP_MODULE.EDIT].dataset.trans[0],
      direction: 70
    }
    edit = new EditController({...host, app: testApp})
    edit.loadEditor = sinon.spy()
    await edit.createTrans({
      cmdtype: "create", transcast: "normal" 
    })
    sinon.assert.callCount(edit.loadEditor, 1);

    testApp.store.data[APP_MODULE.EDIT].dataset.trans[0] = {
      ...testApp.store.data[APP_MODULE.EDIT].dataset.trans[0],
      transtype: 64
    }
    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
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
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.loadEditor = sinon.spy()
    await edit.createTrans({
      cmdtype: "create", transcast: "normal" 
    })
    sinon.assert.callCount(edit.loadEditor, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }
    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              payment: [
                { id: 123, deleted: 0}
              ]
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.loadEditor = sinon.spy()
    await edit.createTrans({
      cmdtype: "copy", transcast: "amendment" 
    })
    sinon.assert.callCount(edit.loadEditor, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }
    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              trans: [...testApp.store.data[APP_MODULE.EDIT].dataset.trans],
              payment: [
                { id: 123, deleted: 0},
                { id: 123, deleted: 1}
              ]
            },
          }
        }
      }
    }
    testApp.store.data[APP_MODULE.EDIT].dataset.trans[0] = {
      ...testApp.store.data[APP_MODULE.EDIT].dataset.trans[0],
      transtype: 57
    }
    edit = new EditController({...host, app: testApp})
    edit.loadEditor = sinon.spy()
    await edit.createTrans({
      cmdtype: "create", transcast: "normal", new_transtype: "invoice", new_direction: "out",
      refno: true, from_inventory: false, netto_qty: true,
    })
    sinon.assert.callCount(edit.loadEditor, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.LOGIN]: {
            ...testApp.store.data[APP_MODULE.LOGIN],
            data: {
              ...testApp.store.data[APP_MODULE.LOGIN].data,
              audit: [
                { nervatypeName: "trans", subtypeName: "invoice", inputfilterName: "disabled" }
              ]
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.createTrans({
      cmdtype: "create", transcast: "normal", new_transtype: "invoice", new_direction: "out",
      refno: true, from_inventory: false, netto_qty: true,
    })
    sinon.assert.callCount(testApp.showToast, 6);

    testApp = {
      ...testApp,
      modules: {
        ...app.modules,
        initItem: (params) => {
          if(params.tablename === "item"){
            return testApp.store.data[APP_MODULE.EDIT].dataset.item[0]
          }
          return {
            id: null, fieldname: null, ref_id: null, value: null, notes: null, deleted: 0
          }
        }
      },
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.LOGIN]: {
            ...testApp.store.data[APP_MODULE.LOGIN],
            data: {
              ...testApp.store.data[APP_MODULE.LOGIN].data,
              audit: [
                { nervatypeName: "trans", subtypeName: "invoice", inputfilterName: "all" }
              ]
            },
          },
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
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
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.loadEditor = sinon.spy()
    await edit.createTrans({
      cmdtype: "create", transcast: "normal", new_transtype: "invoice", new_direction: "out",
      refno: true, from_inventory: false, netto_qty: true,
    })
    sinon.assert.callCount(edit.loadEditor, 1);
    
    await edit.createTrans({
      cmdtype: "create", transcast: "normal", new_transtype: "invoice", new_direction: "out",
      refno: true, from_inventory: true, netto_qty: true,
    })
    sinon.assert.callCount(edit.loadEditor, 2);

    testApp.store.data[APP_MODULE.EDIT].dataset.trans[0] = {
      ...testApp.store.data[APP_MODULE.EDIT].dataset.trans[0],
      direction: 69
    }
    edit = new EditController({...host, app: testApp})
    edit.loadEditor = sinon.spy()
    await edit.createTrans({
      cmdtype: "create", transcast: "normal", new_transtype: "invoice", new_direction: "out",
      refno: true, from_inventory: true, netto_qty: true,
    })
    sinon.assert.callCount(edit.loadEditor, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              item: []
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.loadEditor = sinon.spy()
    await edit.createTrans({
      cmdtype: "create", transcast: "normal", new_transtype: "invoice", new_direction: "out",
      refno: true, from_inventory: false, netto_qty: true,
    })
    sinon.assert.callCount(edit.loadEditor, 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path, options) => {
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
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }
    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
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
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.createTrans({
      cmdtype: "copy", 
    })
    sinon.assert.callCount(testApp.resultError, 7);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
        if(path === "/view"){
          return { fields: [{ 
            trans_custinvoice_compname: "compname", trans_custinvoice_custname: "custname"
          }] }
        }
        if(path === "/movement") { 
          return { error: {} }
        }
        return [1,2]
      })
    }
    edit = new EditController({...host, app: testApp})
    await edit.createTrans({
      cmdtype: "copy", 
    })
    sinon.assert.callCount(testApp.resultError, 8);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
        if(path === "/fieldvalue") { 
          return { error: {} }
        }
        return [1,2]
      }),
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }
    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              trans: [...testApp.store.data[APP_MODULE.EDIT].dataset.trans]
            },
          }
        }
      }
    }
    testApp.store.data[APP_MODULE.EDIT].dataset.trans[0] = {
      ...testApp.store.data[APP_MODULE.EDIT].dataset.trans[0],
      transtype: 57
    }
    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              fieldvalue: [
                { deleted: 0, fieldname: "2b0bd752-2f00-cbdb-a4ee-6741d890c8a6",
                  id: 101, notes: null, ref_id: 5, value: "nervatura.github.io/" },
                { deleted: 1, fieldname: "sample_customer_float", id: 59, name: "deleted_value",
                  notes: null, ref_id: 5, value: "2" },
                { deleted: 0, fieldname: "trans_custinvoice_compname", id: 59, name: "deleted_value",
                  notes: null, ref_id: 5, value: "2" },
                { deleted: 0, fieldname: "link_qty", id: 59, name: "deleted_value",
                  notes: null, ref_id: 5, value: "2" },
              ]
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.createTrans({
      cmdtype: "create" 
    })
    sinon.assert.callCount(testApp.resultError, 9);

  })

  it('createTransOptions', async () => {
    let testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: {} })
          }
        }),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }
    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              element_count: [
                { pec: 0 }
              ],
              payment: []
            },
          }
        }
      }
    }

    let edit = new EditController({...host, app: testApp})
    edit.module.modalTrans = (prm)=>(prm)
    edit.createTrans = sinon.spy()
    edit.createTransOptions()
    sinon.assert.callCount(edit.createTrans, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              transtype: "offer"
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.module.modalTrans = (prm)=>(prm)
    edit.createTrans = sinon.spy()
    edit.createTransOptions()
    sinon.assert.callCount(edit.createTrans, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              transtype: "order"
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.module.modalTrans = (prm)=>(prm)
    edit.createTrans = sinon.spy()
    edit.createTransOptions()
    sinon.assert.callCount(edit.createTrans, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              element_count: [
                { pec: 1 }
              ]
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.module.modalTrans = (prm)=>(prm)
    edit.createTrans = sinon.spy()
    edit.createTransOptions()
    sinon.assert.callCount(edit.createTrans, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              transtype: "worksheet"
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.module.modalTrans = (prm)=>(prm)
    edit.createTrans = sinon.spy()
    edit.createTransOptions()
    sinon.assert.callCount(edit.createTrans, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              element_count: [
                { pec: 0 }
              ]
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.module.modalTrans = (prm)=>(prm)
    edit.createTrans = sinon.spy()
    edit.createTransOptions()
    sinon.assert.callCount(edit.createTrans, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              transtype: "rent"
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.module.modalTrans = (prm)=>(prm)
    edit.createTrans = sinon.spy()
    edit.createTransOptions()
    sinon.assert.callCount(edit.createTrans, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              element_count: [
                { pec: 1 }
              ]
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.module.modalTrans = (prm)=>(prm)
    edit.createTrans = sinon.spy()
    edit.createTransOptions()
    sinon.assert.callCount(edit.createTrans, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              transtype: "receipt"
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.module.modalTrans = (prm)=>(prm)
    edit.createTrans = sinon.spy()
    edit.createTransOptions()
    sinon.assert.callCount(edit.createTrans, 1);

  })

  it('deleteEditor', async () => {
    let testApp = {
      ...app,
      resultError: sinon.spy(),
      showToast: sinon.spy(),
      requestData: sinon.spy(async () => ({ state: [{ sco: 0 }] })),
      getAuditFilter: sinon.spy(() => (["all",1])),
      createHistory: sinon.spy(),
      currentModule: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL, data: {} })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: {} })
          }
        }),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }

    let edit = new EditController({...host, app: testApp})
    await edit.deleteEditor()
    sinon.assert.callCount(testApp.requestData, 2);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path, options) => {
        if(options.method === "DELETE"){
          return { error: {} }
        }
        return { state: [{ sco: 0 }] }
      }),
      resultError: sinon.spy()
    }
    edit = new EditController({...host, app: testApp})
    await edit.deleteEditor()
    sinon.assert.callCount(testApp.requestData, 2);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ state: [{ sco: 1 }] })),
      resultError: sinon.spy()
    }
    edit = new EditController({...host, app: testApp})
    await edit.deleteEditor()
    sinon.assert.callCount(testApp.requestData, 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} })),
      resultError: sinon.spy()
    }
    edit = new EditController({...host, app: testApp})
    await edit.deleteEditor()
    sinon.assert.callCount(testApp.requestData, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              type: "event"
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.deleteEditor()
    sinon.assert.callCount(testApp.requestData, 2);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              item: {
                ...testApp.store.data[APP_MODULE.EDIT].current.item,
                id: null
              }
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.deleteEditor()
    sinon.assert.callCount(testApp.requestData, 2);

  })

  it('deleteEditorItem', async () => {
    let testApp = {
      ...app,
      resultError: sinon.spy(),
      showToast: sinon.spy(),
      requestData: sinon.spy(async () => (testApp.store.data[APP_MODULE.EDIT].dataset)),
      getAuditFilter: sinon.spy(() => (["all",1])),
      createHistory: sinon.spy(),
      currentModule: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL, data: {} })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: {} })
          }
        }),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }

    let edit = new EditController({...host, app: testApp})
    edit.loadEditor = sinon.spy()
    await edit.deleteEditorItem({ fkey: "item", id: 18 })
    sinon.assert.callCount(testApp.createHistory, 1);

    await edit.deleteEditorItem({ fkey: "item", id: 18, prompt: true, table: "item" })
    sinon.assert.callCount(edit.loadEditor, 2);

    const callback = sinon.spy()
    await edit.deleteEditorItem({ fkey: "item", id: 18, prompt: true, table: "item", callback })
    sinon.assert.callCount(callback, 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} })),
      resultError: sinon.spy()
    }
    edit = new EditController({...host, app: testApp})
    edit.loadEditor = sinon.spy()
    await edit.deleteEditorItem({ fkey: "item", id: 18, prompt: true })
    sinon.assert.callCount(testApp.resultError, 1);

    await edit.deleteEditorItem({ fkey: "item", id: null, prompt: true })
    sinon.assert.callCount(edit.loadEditor, 1);

  })

  it('editItem', async () => {
    let testApp = {
      ...app,
      modules: {
        ...app.modules,
        initItem: ()=>({})
      },
      resultError: sinon.spy(),
      requestData: sinon.spy(async () => (testApp.store.data[APP_MODULE.EDIT].dataset)),
      store: {
        ...app.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL, data: {} })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: {} })
          }
        }),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...Item.args,
          }
        }
      }
    }

    let edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "description", value: "description"
    })
    sinon.assert.callCount(testApp.store.setData, 1);

    await edit.editItem({
      name: "product_id", value: 1, item: testApp.store.data[APP_MODULE.EDIT].current.form
    })
    sinon.assert.callCount(testApp.store.setData, 2);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({
          price: 1, discount: 0
        })),
      resultError: sinon.spy(),
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              form: {
                ...testApp.store.data[APP_MODULE.EDIT].current.form,
                qty: 0
              }
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "product_id", value: 1, item: testApp.store.data[APP_MODULE.EDIT].current.form, event_type: "blur"
    })
    sinon.assert.callCount(testApp.store.setData, 3);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} })),
      resultError: sinon.spy()
    }
    edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "product_id", value: 1, item: testApp.store.data[APP_MODULE.EDIT].current.form, event_type: "blur"
    })
    sinon.assert.callCount(testApp.resultError, 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({
          price: 1, discount: 0
        })),
      resultError: sinon.spy()
    }
    edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "qty", value: 1, event_type: "blur"
    })
    sinon.assert.callCount(testApp.store.setData, 4);
    
    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              form: {
                ...testApp.store.data[APP_MODULE.EDIT].current.form,
                fxprice: 0
              }
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "qty", value: 1
    })
    sinon.assert.callCount(testApp.store.setData, 5);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({})),
      resultError: sinon.spy()
    }
    edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "qty", value: 1
    })
    sinon.assert.callCount(testApp.store.setData, 6);

    
    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} })),
      resultError: sinon.spy()
    }
    edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "qty", value: 1
    })
    sinon.assert.callCount(testApp.resultError, 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({
          price: 1, discount: 0
        })),
      resultError: sinon.spy()
    }
    edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "fxprice", value: 10
    })
    sinon.assert.callCount(testApp.store.setData, 7);

    await edit.editItem({
      name: "tax_id", value: 10, event_type: "blur"
    })
    sinon.assert.callCount(testApp.store.setData, 8);

    await edit.editItem({
      name: "discount", value: 10, event_type: "blur"
    })
    sinon.assert.callCount(testApp.store.setData, 9);

    await edit.editItem({
      name: "amount", value: 10, event_type: "blur"
    })
    sinon.assert.callCount(testApp.store.setData, 10);

    await edit.editItem({
      name: "amount", value: 10
    })
    sinon.assert.callCount(testApp.store.setData, 11);

    await edit.editItem({
      name: "netamount", value: 10, event_type: "blur"
    })
    sinon.assert.callCount(testApp.store.setData, 12);

    await edit.editItem({
      name: "netamount", value: 10
    })
    sinon.assert.callCount(testApp.store.setData, 13);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              form_type: "price"
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "curr", value: "EUR"
    })
    sinon.assert.callCount(testApp.store.setData, 14);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              form_type: "discount"
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "customer_id", value: 1
    })
    sinon.assert.callCount(testApp.store.setData, 15);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              form_type: "invoice_link",
              invoice_link_fieldvalue: [],
              invoice_link: [
                {}
              ]
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "ref_id_1", value: 1, item: { curr: "EUR" }
    })
    sinon.assert.callCount(testApp.store.setData, 16);

    await edit.editItem({
      name: "link_qty", value: 1
    })
    sinon.assert.callCount(testApp.store.setData, 17);

    await edit.editItem({
      name: "ref_id_2", value: 1
    })
    sinon.assert.callCount(testApp.store.setData, 18);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              form_type: "payment_link",
              payment_link_fieldvalue: [],
              payment_link: [
                {}
              ]
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "ref_id_2", value: 1, item: { curr: "EUR" }
    })
    sinon.assert.callCount(testApp.store.setData, 19);

    await edit.editItem({
      name: "link_qty", value: 1
    })
    sinon.assert.callCount(testApp.store.setData, 20);

    await edit.editItem({
      name: "ref_id_1", value: 1
    })
    sinon.assert.callCount(testApp.store.setData, 21);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              form_type: "default",
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "default", value: "default"
    })
    sinon.assert.callCount(testApp.store.setData, 22);

    testApp = {
      ...testApp,
      store: {
        ...app.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL, data: {} })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: {} })
          }
        }),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "paiddate", value: "2021-12-01"
    })
    sinon.assert.callCount(testApp.store.setData, 1);

    await edit.editItem({
      name: "closed", value: 1
    })
    sinon.assert.callCount(testApp.store.setData, 6);

    await edit.editItem({
      name: "closed", value: 0
    })
    sinon.assert.callCount(testApp.store.setData, 7);

    await edit.editItem({
      name: "direction", value: 69
    })
    sinon.assert.callCount(testApp.store.setData, 8);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              transtype: "cash",
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "direction", value: 69
    })
    sinon.assert.callCount(testApp.store.setData, 9);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              transtype: "waybill",
              extend: {
                seltype: "customer"
              }
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "seltype", value: 69
    })
    sinon.assert.callCount(testApp.store.setData, 10);

    await edit.editItem({
      name: "ref_id", value: 69, refnumber: "refnumber"
    })
    sinon.assert.callCount(testApp.store.setData, 11);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              extend: {
                seltype: "employee"
              }
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "ref_id", value: 69, refnumber: "refnumber", item: { transtype: "order" }
    })
    sinon.assert.callCount(testApp.store.setData, 12);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              extend: {
                seltype: "transitem"
              }
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "ref_id", value: 69, refnumber: "refnumber"
    })
    sinon.assert.callCount(testApp.store.setData, 13);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              extend: {
                seltype: ""
              }
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "ref_id", value: 69, refnumber: "refnumber"
    })
    sinon.assert.callCount(testApp.store.setData, 14);

    await edit.editItem({
      name: "trans_wsdistance", value: 69
    })
    sinon.assert.callCount(testApp.store.setData, 15);

    await edit.editItem({
      name: "trans_wsrepair", value: 69
    })
    sinon.assert.callCount(testApp.store.setData, 16);

    await edit.editItem({
      name: "trans_wstotal", value: 69
    })
    sinon.assert.callCount(testApp.store.setData, 17);

    await edit.editItem({
      name: "trans_reholiday", value: 69
    })
    sinon.assert.callCount(testApp.store.setData, 18);

    await edit.editItem({
      name: "trans_rebadtool", value: 69
    })
    sinon.assert.callCount(testApp.store.setData, 19);

    await edit.editItem({
      name: "trans_reother", value: 69
    })
    sinon.assert.callCount(testApp.store.setData, 20);

    await edit.editItem({
      name: "trans_wsnote", value: 69
    })
    sinon.assert.callCount(testApp.store.setData, 21);

    await edit.editItem({
      name: "trans_rentnote", value: 69
    })
    sinon.assert.callCount(testApp.store.setData, 22);

    await edit.editItem({
      name: "shippingdate", value: "shippingdate"
    })
    sinon.assert.callCount(testApp.store.setData, 23);

    await edit.editItem({
      name: "shipping_place_id", value: 1
    })
    sinon.assert.callCount(testApp.store.setData, 24);

    await edit.editItem({
      name: "fnote", value: "", extend: true
    })
    sinon.assert.callCount(testApp.store.setData, 25);

    await edit.editItem({
      name: "default", value: ""
    })
    sinon.assert.callCount(testApp.store.setData, 26);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              type: "default"
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "default", value: ""
    })
    sinon.assert.callCount(testApp.store.setData, 27);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            printqueue: {},
            audit: "readonly",
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              type: "printqueue",
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "name", value: "value"
    })
    sinon.assert.callCount(testApp.store.setData, 28);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              type: "report",
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "selected", value: 1
    })
    sinon.assert.callCount(testApp.store.setData, 29);

    await edit.editItem({
      name: "selected", value: 1, id: 100
    })
    sinon.assert.callCount(testApp.store.setData, 30);

    await edit.editItem({
      name: "name", value: 1
    })
    sinon.assert.callCount(testApp.store.setData, 31);

    await edit.editItem({
      name: "trans_transcast", value: 1
    })
    sinon.assert.callCount(testApp.store.setData, 32);

    await edit.editItem({
      name: "customer_id", value: 69, refnumber: "refnumber", extend: false, label_field: "customer"
    })
    sinon.assert.callCount(testApp.store.setData, 33);

    await edit.editItem({
      name: "customer_id", value: 69, extend: false, label_field: "customer"
    })
    sinon.assert.callCount(testApp.store.setData, 34);

    await edit.editItem({
      name: "customer_id", value: 69, extend: false
    })
    sinon.assert.callCount(testApp.store.setData, 35);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            template: {
              ...testApp.store.data[APP_MODULE.EDIT].template,
              options: {
                ...testApp.store.data[APP_MODULE.EDIT].template.options,
                extend: true
              }
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "customer_id", value: 69, extend: true
    })
    sinon.assert.callCount(testApp.store.setData, 36);

    await edit.editItem({
      name: "fieldvalue_value", value: "value", id: 123
    })
    sinon.assert.callCount(testApp.store.setData, 37);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            audit: "all"
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.editItem({
      name: "fieldvalue_value", value: "value", id: 59
    })
    sinon.assert.callCount(testApp.store.setData, 38);

    await edit.editItem({
      name: "fieldvalue_deleted", value: true, id: 59
    })
    sinon.assert.callCount(testApp.store.setData, 39);

  })

  it('exportEvent', () => {
    let testApp = {
      ...app,
      saveToDisk: sinon.spy(),
      store: {
        ...app.store,
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            current: {
              ...app.store.data[APP_MODULE.EDIT].current,
              item: {
                id: null, calnumber: "calnumber", 
                nervatype: null, ref_id: null, 
                uid: null, eventgroup: null, fromdate: null, todate: null, subject: null, 
                place: null, description: null, deleted: 0
              }
            }
          }
        }
      }
    }

    let edit = new EditController({...host, app: testApp})
    edit.exportEvent()
    sinon.assert.callCount(testApp.saveToDisk, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              item: {
                id: null, calnumber: "calnumber", 
                nervatype: null, ref_id: null, 
                uid: "uid", eventgroup: 123, 
                fromdate: "2021-12-01T00:00:00", todate: "2021-12-01T00:00:00", subject: "subject", 
                place: "place", description: "description", deleted: 0
              }
            },
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              eventgroup: []
            }
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.exportEvent()
    sinon.assert.callCount(testApp.saveToDisk, 2);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              item: {
                id: null, calnumber: "calnumber", 
                nervatype: null, ref_id: null, 
                uid: "uid", eventgroup: 123, 
                fromdate: "2021-12-01T00:00:00", todate: "2021-12-01T00:00:00", subject: "subject", 
                place: "place", description: "description", deleted: 0
              }
            },
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              eventgroup: [
                { id: 123, groupvalue: "value" }
              ]
            }
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.exportEvent()
    sinon.assert.callCount(testApp.saveToDisk, 3);

  })

  it('exportQueueAll', async () => {
    let testApp = {
      ...app,
      showToast: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL, data: {} })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: {} })
          }
        }),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
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
          }
        }
      }
    }

    let edit = new EditController({...host, app: testApp})
    const result = edit.exportQueueAll()
    expect(result).to.true

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
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
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.exportQueueAll()
    sinon.assert.callCount(testApp.showToast, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
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
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.exportQueue = sinon.spy(async ()=>(true))
    edit.searchQueue = sinon.spy()
    await edit.exportQueueAll()
    sinon.assert.callCount(edit.searchQueue, 1);

  })

  it('exportQueue', async () => {
    const testApp = {
      ...app,
      store: {
        ...app.store,
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...PrintQueue.args,
          }
        }
      }
    }

    const edit = new EditController({...host, app: testApp})
    edit.reportOutput = sinon.spy(async ()=>({}))
    edit.deleteEditorItem = sinon.spy(()=>({}))
    await edit.exportQueue({ 
      id: 1, ref_id: 1, reportkey: "reportkey", refnumber: "refnumber", copies: 1, typename: "typename" })
    sinon.assert.callCount(edit.reportOutput, 1);
  
  })

  it('getTransFilter', async () => {
    let testApp = {
      ...app,
      store: {
        ...app.store,
        data: {
          ...app.store.data,
          [APP_MODULE.LOGIN]: {
            ...app.store.data[APP_MODULE.LOGIN],
            data: {
              ...app.store.data[APP_MODULE.LOGIN].data,
              transfilterName: "usergroup"
            }
          }
        }
      }
    }

    let edit = new EditController({...host, app: testApp})
    let result = edit.getTransFilter({ where: [] }, [])
    expect(result[1].length).to.equal(1)

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.LOGIN]: {
            ...testApp.store.data[APP_MODULE.LOGIN],
            data: {
              ...testApp.store.data[APP_MODULE.LOGIN].data,
              transfilterName: "own"
            }
          }
        }
      }
    }

    edit = new EditController({...host, app: testApp})
    result = edit.getTransFilter({ where: [] }, [])
    expect(result[1].length).to.equal(1)

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.LOGIN]: {
            ...testApp.store.data[APP_MODULE.LOGIN],
            data: {
              ...testApp.store.data[APP_MODULE.LOGIN].data,
              transfilterName: "all"
            }
          }
        }
      }
    }

    edit = new EditController({...host, app: testApp})
    result = edit.getTransFilter({ where: [] }, [])
    expect(result[1].length).to.equal(0)

  })

  it('loadEditor', async () => {
    let testApp = {
      ...app,
      modules: {
        ...app.modules,
        initItem: (params) => {
          if (Object.keys(params.dataset).length === 0){
            return null
          }
          return {
            ...testApp.store.data[APP_MODULE.EDIT].current.item,
            id: null
          }
        }
      },
      requestData: sinon.spy(async () => (testApp.store.data[APP_MODULE.EDIT].dataset)),
      resultError: sinon.spy(),
      showToast: sinon.spy(), 
      getSetting: (key)=>storeConfig[key],
      getAuditFilter: sinon.spy(() => (["all",1])),
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }
    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
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
            },
          }
        }
      }
    }

    let edit = new EditController({...host, app: testApp})
    edit.setEditor = sinon.spy()
    await edit.loadEditor({
      ntype: "trans", ttype: "invoice", id: 5
    })
    sinon.assert.callCount(edit.setEditor, 1);

    testApp = {
      ...testApp,
      getSql: (engine, _sql)=>({ sql: "", prmCount: (_sql.from === "fieldvalue") ? 1 : 0 })
    }
    edit = new EditController({...host, app: testApp})
    edit.setEditor = sinon.spy()
    await edit.loadEditor({
      ntype: "trans", ttype: "invoice", id: null, shipping: true
    })
    sinon.assert.callCount(edit.setEditor, 1);

    const result = await edit.loadEditor({
      ntype: "trans", ttype: "formula", id: 18, cb_key: "LOAD_EDITOR"
    })
    expect(result).to.true;

    await edit.loadEditor({
      ntype: "trans", ttype: "delivery", id: null
    })
    sinon.assert.callCount(edit.setEditor, 2);

    testApp = {
      ...testApp,
      modules: {
        ...testApp.modules,
        initItem: () => ({
            ...testApp.store.data[APP_MODULE.EDIT].current.item,
            id: null
          })
      },
      store: {
        ...testApp.store,
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...PrintQueue.args
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.setEditor = sinon.spy()
    await edit.loadEditor({
      ntype: "printqueue", ttype: null, id: null
    })
    sinon.assert.callCount(edit.setEditor, 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} })),
      resultError: sinon.spy(),
      store: {
        ...testApp.store,
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...Report.args
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.loadEditor({
      ntype: "report", ttype: null, id: 7
    })
    sinon.assert.callCount(testApp.resultError, 1);

  })

  it('newFieldvalue', async () => {
    let testApp = {
      ...app,
      modules: {
        ...app.modules,
        initItem: ()=>({})
      },
      requestData: sinon.spy(async () => ({})),
      resultError: sinon.spy(),
      showToast: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }
    let edit = new EditController({...host, app: testApp})
    await edit.newFieldvalue("")
    sinon.assert.callCount(testApp.showToast, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.LOGIN]: {
            ...testApp.store.data[APP_MODULE.LOGIN],
            data: {
              ...testApp.store.data[APP_MODULE.LOGIN].data,
              groups: testApp.store.data[APP_MODULE.EDIT].dataset.groups
            }
          }
        }
      }
    }
    // valuelist
    edit = new EditController({...host, app: testApp})
    edit.loadEditor = sinon.spy()
    await edit.newFieldvalue("trans_transcast")
    sinon.assert.callCount(edit.loadEditor, 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} })),
      resultError: sinon.spy(),
    }
    edit = new EditController({...host, app: testApp})
    edit.onSelector = (type, filter, callback)=>{ callback({ id: "ntype/ttype/1" }) }
    // transitem
    await edit.newFieldvalue("trans_transitem_link")
    sinon.assert.callCount(testApp.resultError, 1);

    // float
    await edit.newFieldvalue("link_qty")
    sinon.assert.callCount(testApp.resultError, 2);

    // date
    await edit.newFieldvalue("sample_customer_date")
    sinon.assert.callCount(testApp.resultError, 3);

    // string
    await edit.newFieldvalue("trans_custinvoice_custname")
    sinon.assert.callCount(testApp.resultError, 4);

    // time
    await edit.newFieldvalue("sample_time")
    sinon.assert.callCount(testApp.resultError, 5);

    // bool
    await edit.newFieldvalue("sample_bool")
    sinon.assert.callCount(testApp.resultError, 6);

  })

  it('nextTransNumber', async () => {
    let testApp = {
      ...app,
      requestData: sinon.spy(async (path) => {
        if(path === "/view"){
          return { next: [{ 
            id: 5
          }] }
        }
        return { error: {} }
      }),
      resultError: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }
    let edit = new EditController({...host, app: testApp})
    edit.checkEditor = sinon.spy()
    await edit.nextTransNumber()
    sinon.assert.callCount(edit.checkEditor, 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
        if(path === "/view"){
          return { next: [{ 
            id: null
          }] }
        }
        return { error: {} }
      }),
      resultError: sinon.spy(),
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              transtype: "waybill"
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.checkEditor = sinon.spy()
    await edit.nextTransNumber()
    sinon.assert.callCount(edit.checkEditor, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              transtype: "delivery"
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    const result = await edit.nextTransNumber()
    expect(result).to.true;

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} })),
      resultError: sinon.spy(),
    }
    edit = new EditController({...host, app: testApp})
    await edit.nextTransNumber()
    sinon.assert.callCount(testApp.resultError, 1);

  })

  it('prevTransNumber', async () => {
    let testApp = {
      ...app,
      requestData: sinon.spy(async (path) => {
        if(path === "/view"){
          return { prev: [{ 
            id: 5
          }] }
        }
        return { error: {} }
      }),
      resultError: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }
    let edit = new EditController({...host, app: testApp})
    edit.checkEditor = sinon.spy()
    await edit.prevTransNumber()
    sinon.assert.callCount(edit.checkEditor, 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
        if(path === "/view"){
          return { prev: [{ 
            id: null
          }] }
        }
        return { error: {} }
      }),
      resultError: sinon.spy(),
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              transtype: "waybill"
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    let result = await edit.prevTransNumber()
    expect(result).to.true;

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} })),
      resultError: sinon.spy(),
    }
    edit = new EditController({...host, app: testApp})
    await edit.prevTransNumber()
    sinon.assert.callCount(testApp.resultError, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              type: "customer"
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    result = await edit.prevTransNumber()
    expect(result).to.true;

  })

  it('reportOutput', async () => {
    let testApp = {
      ...app,
      requestData: sinon.spy(async () => ({})),
      resultError: sinon.spy(),
      saveToDisk: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            current: {
              ...app.store.data[APP_MODULE.EDIT].current,
              type: "type",
              item: {
                id: 1
              }
            },
            dataset: {
              ...app.store.data[APP_MODULE.EDIT].dataset,
              report: [
                { id: 1, reportkey: "reportkey" }
              ]
            }
          }
        }
      }
    }
    let edit = new EditController({...host, app: testApp})
    edit.addPrintQueue = sinon.spy()
    await edit.reportOutput({ type: "printqueue", template: "reportkey", copy: 1 })
    sinon.assert.callCount(edit.addPrintQueue, 1);

    await edit.reportOutput({ type: "xml", template: "reportkey", copy: 1, title: "title" })
    sinon.assert.callCount(testApp.saveToDisk, 1);

    await edit.reportOutput({ type: "print", template: "reportkey" })
    sinon.assert.callCount(testApp.saveToDisk, 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} })),
      resultError: sinon.spy(),
    }
    edit = new EditController({...host, app: testApp})
    await edit.reportOutput({ type: "print", template: "reportkey" })
    sinon.assert.callCount(testApp.resultError, 1);

  })

  it('reportPath', () => {
    const testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            current: {
              ...app.store.data[APP_MODULE.EDIT].current,
              type: "nervatype",
              item: {
                id: 2
              }
            }
          }
        }
      }
    }
    const edit = new EditController({...host, app: testApp})
    let result = edit.reportPath({ 
      template: "template", orient: "orient", size: "size", type: "type", nervatype: "nervatype", id: 1
    })
    expect(result).to.equal("/report?reportkey=template&orientation=orient&size=size&output=type&nervatype=nervatype&filters[@id]=1")
    result = edit.reportPath({ 
      template: "template", orient: "orient", size: "size", type: "type"
    })
    expect(result).to.equal("/report?reportkey=template&orientation=orient&size=size&output=type&nervatype=nervatype&filters[@id]=2")
    result = edit.reportPath({ 
      template: "template", orient: "orient", size: "size", type: "type", nervatype: "nervatype", filters: "abc=123"
    })
    expect(result).to.equal("/report?reportkey=template&orientation=orient&size=size&output=type&abc=123")
  })

  it('reportSettings', async () => {
    let testApp = {
      ...app,
      getSetting: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: {} })
          }
        }),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            current: {
              ...app.store.data[APP_MODULE.EDIT].current,
              type: "trans",
              transtype: "invoice",
              item: {
                direction: 1, transnumber: "transnumber"
              }
            },
            template: {
              ...app.store.data[APP_MODULE.EDIT].template,
              options: {
                title_field: "transnumber"
              }
            },
            dataset: {
              ...app.store.data[APP_MODULE.EDIT].dataset,
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
          }
        }
      }
    }

    let edit = new EditController({...host, app: testApp})
    edit.module.modalReport = (prm)=>(prm)
    edit.reportOutput = sinon.spy()
    edit.reportSettings()
    sinon.assert.callCount(edit.reportOutput, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.LOGIN]: {
            ...testApp.store.data[APP_MODULE.LOGIN],
            data: {
              ...testApp.store.data[APP_MODULE.LOGIN].data,
              audit: []
            }
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.module.modalReport = (prm)=>(prm)
    edit.reportOutput = sinon.spy()
    edit.reportSettings()
    sinon.assert.callCount(edit.reportOutput, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              groups: [
                { id: 1, groupname: "direction", groupvalue: "out" }
              ],
              settings: [
                { fieldname: "default_trans_invoice_out_report", value: "value" }
              ],
              report: [
                { id: 1, direction: 2, reportkey: "reportkey", repname: "repname" }
              ]
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.module.modalReport = (prm)=>(prm)
    edit.reportOutput = sinon.spy()
    edit.reportSettings()
    sinon.assert.callCount(edit.reportOutput, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              type: "customer",
              item: {
                custnumber: "custnumber"
              }
            },
            template: {
              ...testApp.store.data[APP_MODULE.EDIT].template,
              options: {
                title_field: "custnumber"
              }
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.module.modalReport = (prm)=>(prm)
    edit.reportOutput = sinon.spy()
    edit.reportSettings()
    sinon.assert.callCount(edit.reportOutput, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.LOGIN]: {
            ...testApp.store.data[APP_MODULE.LOGIN],
            data: {
              ...testApp.store.data[APP_MODULE.LOGIN].data,
              audit: [
                { nervatypeName: "report", subtype: 1, inputfilterName: "disabled" }
              ]
            }
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.module.modalReport = (prm)=>(prm)
    edit.reportOutput = sinon.spy()
    edit.reportSettings()
    sinon.assert.callCount(edit.reportOutput, 1);

  })

  it('saveEditor', async () => {
    let testApp = {
      ...app,
      modules: {
        ...app.modules,
        initItem: () => ({
          ...InvoiceData.args.current.item,
          id: null
        }),
        validator: () => ({
          ...InvoiceData.args.current.item,
          id: null
        })
      },
      requestData: sinon.spy(async (path) => {
        if(path === "/view"){
          return { fields: [{ 
            trans_custinvoice_compname: "compname", trans_custinvoice_custname: "custname"
          }] }
        }
        return [1,2]
      }),
      resultError: sinon.spy(),
      createHistory: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }
    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              item: {
                ...testApp.store.data[APP_MODULE.EDIT].current.item,
                id: null
              }
            },
          }
        }
      }
    }
    let edit = new EditController({...host, app: testApp})
    let result = await edit.saveEditor()
    expect(result).to.exist;

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              item: {
                ...testApp.store.data[APP_MODULE.EDIT].current.item,
                id: 5
              }
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    result = await edit.saveEditor()
    expect(result).to.exist;

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
        if(path === "/view"){
          return { fields: [] }
        }
        if(path === "/fieldvalue"){
          return { error: {} }
        }
        return [1,2]
      }),
      resultError: sinon.spy(),
    }
    edit = new EditController({...host, app: testApp})
    await edit.saveEditor()
    sinon.assert.callCount(testApp.resultError, 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
        if(path === "/view"){
          return { error: {} }
        }
        return [1,2]
      }),
      resultError: sinon.spy(),
    }
    edit = new EditController({...host, app: testApp})
    await edit.saveEditor()
    sinon.assert.callCount(testApp.resultError, 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} })),
      resultError: sinon.spy(),
    }
    edit = new EditController({...host, app: testApp})
    await edit.saveEditor()
    sinon.assert.callCount(testApp.resultError, 1);

    testApp = {
      ...testApp,
      modules: {
        ...testApp.modules,
        validator: sinon.spy(async () => ({ error: {} }))
      },
      resultError: sinon.spy(),
    }
    edit = new EditController({...host, app: testApp})
    await edit.saveEditor()
    sinon.assert.callCount(testApp.resultError, 1);

    testApp = {
      ...testApp,
      modules: {
        ...testApp.modules,
        validator: () => ({
          ...InvoiceData.args.current.item,
          id: null
        })
      },
      requestData: sinon.spy(async () => [1,2,3,4]),
      resultError: sinon.spy(),
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              transtype: "worksheet",
              fieldvalue: [
                { id: null, fieldname: 'trans_wsrepair', ref_id: null,
                  value: 22, notes: null, deleted: 0 }]
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    result = await edit.saveEditor()
    expect(result).to.exist;

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              transtype: "rent",
              fieldvalue: []
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    result = await edit.saveEditor()
    expect(result).to.exist;

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              transtype: "delivery"
            },
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              movement: [
                { id: 1, ref_id: 1, place_id: 2, shippingdate: "2021-12-10" },
                { id: 2, ref_id: 1, place_id: 2, shippingdate: "2021-12-11" }
              ]
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    result = await edit.saveEditor()
    expect(result).to.exist;

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
        if(path === "/movement"){
          return { error: {} }
        }
        return [1,2]
      }),
      resultError: sinon.spy(),
    }
    edit = new EditController({...host, app: testApp})
    await edit.saveEditor()
    sinon.assert.callCount(testApp.resultError, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              movement: []
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    result = await edit.saveEditor()
    expect(result).to.exist;

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              transtype: "offer"
            }
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    result = await edit.saveEditor()
    expect(result).to.exist;

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...PrintQueue.args
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    result = await edit.saveEditor()
    expect(result).to.exist;

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args
          }
        }
      }
    }
    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              transtype: "waybill",
              extend: {
                seltype: "employee",
                ref_id: null,
                refnumber: "DMEMP/00001",
                transtype: "",
              }
            },
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              translink: [
                { id: 1 }
              ]
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    result = await edit.saveEditor()
    expect(result).to.exist;
    
    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
        if(path === "/link"){
          return { error: {} }
        }
        return [1,2]
      }),
      resultError: sinon.spy(),
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              extend: {
                seltype: "employee",
                ref_id: 5,
                refnumber: "DMEMP/00001",
                transtype: "",
              }
            }
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.saveEditor()
    sinon.assert.callCount(testApp.resultError, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              translink: []
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    result = await edit.saveEditor()
    expect(result).to.exist;

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              extend: {
                seltype: "transitem",
                ref_id: 5,
                refnumber: "DMEMP/00001",
                transtype: "",
              }
            },
          },
          [APP_MODULE.LOGIN]: {
            ...testApp.store.data[APP_MODULE.LOGIN],
            data: {
              ...testApp.store.data[APP_MODULE.LOGIN].data,
              groups: testApp.store.data[APP_MODULE.EDIT].dataset.groups
            }
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.saveEditor()
    sinon.assert.callCount(testApp.resultError, 2);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => [1,2]),
      resultError: sinon.spy(),
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              translink: [
                { id: 1 }
              ]
            }
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    result = await edit.saveEditor()
    expect(result).to.exist;

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              transtype: "formula",
              extend: {
                id: null, trans_id: null, shippingdate: "2021-12-01T00:00:00Z", product_id: 4,
                product: "DMPROD/00004 | Car", movetype: 92, tool_id: null, qty: 5,
                place_id: null, shared: 0, notes: null, deleted: 0, description: "Car",
                partnumber: "DMPROD/00004",
              }
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    result = await edit.saveEditor()
    expect(result).to.exist;

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              transtype: "production"
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    result = await edit.saveEditor()
    expect(result).to.exist;

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              transtype: "cash"
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    result = await edit.saveEditor()
    expect(result).to.exist;

  })

  it('saveEditorForm', async () => {
    let testApp = {
      ...app,
      modules: {
        ...app.modules,
        initItem: () => ({
          id: null, value: null
        }),
        validator: () => ({})
      },
      requestData: sinon.spy(async () => ({})),
      resultError: sinon.spy(),
      createHistory: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }
    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              form_datatype: "item",
              form: { id: null, value: "value"}
            },
          },
          [APP_MODULE.LOGIN]: {
            ...testApp.store.data[APP_MODULE.LOGIN],
            data: {
              ...testApp.store.data[APP_MODULE.LOGIN].data,
              groups: testApp.store.data[APP_MODULE.EDIT].dataset.groups
            }
          }
        }
      }
    }
    let edit = new EditController({...host, app: testApp})
    let result = await edit.saveEditorForm()
    expect(result).to.exist;

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              form_datatype: "movement",
              form_type: "movement",
              transtype: "inventory",
              form: { id: 1, value: "value" }
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    result = await edit.saveEditorForm()
    expect(result).to.exist;

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              transtype: "delivery",
              form: { id: 1, value: "value", qty: 1, product_id: 1, notes: "", place_id: 1 }
            },
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              movement_transfer: []
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    result = await edit.saveEditorForm()
    expect(result).to.exist;

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path, options) => {
        if((path === "/link") || ((path === "/movement") && (options.data.length === 3))){
          return { error: {} }
        }
        return {}
      }),
      resultError: sinon.spy(),
    }
    edit = new EditController({...host, app: testApp})
    await edit.saveEditorForm()
    sinon.assert.callCount(testApp.resultError, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              movement_transfer: [
                { id: 1, ref_id: 1 },
                { id: 2, ref_id: 1 },
                { id: 3, ref_id: 1 }
              ],
              movement: [
                { id: 1, ref_id: 1, place_id: 2 },
                { id: 2, ref_id: 1, place_id: 2 }
              ]
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.saveEditorForm()
    sinon.assert.callCount(testApp.resultError, 2);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({})),
      resultError: sinon.spy(),
    }
    edit = new EditController({...host, app: testApp})
    result = await edit.saveEditorForm()
    expect(result).to.exist;

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              form_datatype: "price",
              form_type: "price",
              form: { id: 1, value: "value" },
              price_link_customer: null,
              price_customer_id: null
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    result = await edit.saveEditorForm()
    expect(result).to.exist;

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              form_datatype: "price",
              form_type: "price",
              form: { id: 1, value: "value" },
              price_link_customer: null,
              price_customer_id: 1
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    result = await edit.saveEditorForm()
    expect(result).to.exist;

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path, options) => {
        if((path === "/link") && (options.method === "POST")){
          return { error: {} }
        }
        return {}
      }),
      resultError: sinon.spy(),
    }
    edit = new EditController({...host, app: testApp})
    await edit.saveEditorForm()
    sinon.assert.callCount(testApp.resultError, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              form_datatype: "price",
              form_type: "price",
              form: { id: 1, value: "value" },
              price_link_customer: 1,
              price_customer_id: null
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    result = await edit.saveEditorForm()
    expect(result).to.exist;

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path, options) => {
        if((path === "/link") && (options.method === "DELETE")){
          return { error: {} }
        }
        return {}
      }),
      resultError: sinon.spy(),
    }
    edit = new EditController({...host, app: testApp})
    await edit.saveEditorForm()
    sinon.assert.callCount(testApp.resultError, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              form_datatype: "price",
              form_type: "invoice_link",
              form: { id: 1, value: "value" },
              invoice_link_fieldvalue: [
                { id: 1, ref_id: null },
                { id: 2, ref_id: 1 }
              ]
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    result = await edit.saveEditorForm()
    expect(result).to.exist;

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
        if(path === "/fieldvalue"){
          return { error: {} }
        }
        return {}
      }),
      resultError: sinon.spy(),
    }
    edit = new EditController({...host, app: testApp})
    await edit.saveEditorForm()
    sinon.assert.callCount(testApp.resultError, 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
        if(path === "/error"){
          return { error: {} }
        }
        return {}
      }),
      resultError: sinon.spy(),
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              form_datatype: "error",
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.saveEditorForm()
    sinon.assert.callCount(testApp.resultError, 1);

    testApp = {
      ...testApp,
      modules: {
        ...testApp.modules,
        validator: () => ({ error: {} })
      },
      resultError: sinon.spy(),
    }
    edit = new EditController({...host, app: testApp})
    await edit.saveEditorForm()
    sinon.assert.callCount(testApp.resultError, 1);

  })

  it('searchQueue', async () => {
    let testApp = {
      ...app,
      requestData: sinon.spy(async () => ({})),
      resultError: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            printqueue: {}
          }
        }
      }
    }
    let edit = new EditController({...host, app: testApp})
    await edit.searchQueue()
    sinon.assert.callCount(testApp.store.setData, 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} })),
      resultError: sinon.spy(),
    }
    edit = new EditController({...host, app: testApp})
    await edit.searchQueue()
    sinon.assert.callCount(testApp.resultError, 1);

  })

  it('setEditor', () => {
    let testApp = {
      ...app,
      modules: {
        ...app.modules,
        initItem: () => ({}),
      },
      resultError: sinon.spy(),
      showToast: sinon.spy(),
      getAuditFilter: sinon.spy(() => (["all",1])),
      currentModule: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }
    const forms = Forms({ msg: (key)=> key })
    let edit = new EditController({...host, app: testApp})
    edit.setEditor({}, {}, { current: { type: "type" }, dataset: {} })
    sinon.assert.callCount(testApp.showToast, 1);

    let options = {
      ntype: "trans", ttype: "invoice", id: 5,
      item: {
        custname: "First Customer Co.", deleted: 0, id: "trans/invoice/5",
        label: "DMINV/00001", transnumber: "DMINV/00001", transtype: "invoice-out",
      },
    }
    let form = forms.invoice(testApp.store.data[APP_MODULE.EDIT].current.item, testApp.store.data[APP_MODULE.EDIT])

    edit = new EditController({...host, app: testApp})
    let result = edit.setEditor(options, form)
    expect(result).to.true;

    testApp = {
      ...testApp,
      getAuditFilter: sinon.spy(() => (["disabled",1])),
    }
    edit = new EditController({...host, app: testApp})
    result = edit.setEditor(options, form)
    expect(result).to.false;

    testApp = {
      ...testApp,
      getAuditFilter: sinon.spy(() => (["all",1])),
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              pattern: undefined,
              fieldvalue: undefined,
              trans: [...testApp.store.data[APP_MODULE.EDIT].dataset.trans]
            },
          }
        }
      }
    }
    testApp.store.data[APP_MODULE.EDIT].dataset.trans[0] = {
      ...InvoiceData.args.current.item,
      id: null, transcast: "cancellation"
    }
    edit = new EditController({...host, app: testApp})
    result = edit.setEditor(options, form)
    expect(result).to.true;

    testApp = {
      ...testApp,
      getAuditFilter: sinon.spy(() => (["readonly",1])),
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              pattern: []
            },
          }
        }
      }
    }
    testApp.store.data[APP_MODULE.EDIT].dataset.trans[0] = {
      ...InvoiceData.args.current.item,
      id: null, transcast: undefined, closed: 1
    }
    form = {
      ...form,
      options: {
        ...form.options,
        extend: "item"
      }
    }
    edit = new EditController({...host, app: testApp})
    result = edit.setEditor(options, form)
    expect(result).to.true;

    testApp.store.data[APP_MODULE.EDIT].dataset.trans[0] = {
      ...InvoiceData.args.current.item,
      id: 5, deleted: 1
    }
    edit = new EditController({...host, app: testApp})
    result = edit.setEditor(options, form)
    expect(result).to.true;

    form = {
      ...form,
      options: {
        ...form.options,
        extend: "cancel_link"
      }
    }
    edit = new EditController({...host, app: testApp})
    result = edit.setEditor(options, form)
    expect(result).to.true;

    form = {
      ...form,
      options: {
        ...form.options,
        extend: "extend"
      }
    }
    edit = new EditController({...host, app: testApp})
    result = edit.setEditor(options, form)
    expect(result).to.true;

    testApp = {
      ...testApp,
      getAuditFilter: sinon.spy(() => (["all",1])),
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...PrintQueue.args
          }
        }
      }
    }
    options = {
      ntype: "printqueue", ttype: null, id: null
    }
    form = forms.printqueue({}, {}, storeConfig.ui)
    edit = new EditController({...host, app: testApp})
    result = edit.setEditor(options, form)
    expect(result).to.true;

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            printqueue: testApp.store.data[APP_MODULE.EDIT].current.item
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    result = edit.setEditor(options, form)
    expect(result).to.true;

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...Report.args
          }
        }
      }
    }
    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              report: [...testApp.store.data[APP_MODULE.EDIT].dataset.report]
            },
          }
        }
      }
    }
    testApp.store.data[APP_MODULE.EDIT].dataset.report[0] = {...testApp.store.data[APP_MODULE.EDIT].dataset.report[0],
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
      }
    )}
    options = {
      ntype: "report", ttype: "", id: 7,
      item: testApp.store.data[APP_MODULE.EDIT].current.item,
    }
    form = forms.report(testApp.store.data[APP_MODULE.EDIT].current.item, {}, storeConfig.ui)
    edit = new EditController({...host, app: testApp})
    result = edit.setEditor(options, form)
    expect(result).to.true;

    testApp.store.data[APP_MODULE.EDIT].dataset.report[0] = {
      ...testApp.store.data[APP_MODULE.EDIT].dataset.report[0],
      report: ""
    }
    edit = new EditController({...host, app: testApp})
    edit.setEditor(options, form)
    sinon.assert.callCount(testApp.currentModule, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args
          }
        }
      }
    }
    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
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
            },
          }
        }
      }
    }
    options = {
      ntype: "trans", ttype: "order", id: 1, shipping: true 
    }
    form = {
      ...forms.invoice(testApp.store.data[APP_MODULE.EDIT].current.item, testApp.store.data[APP_MODULE.EDIT], storeConfig.ui),
    }
    form = {...form,
      view: {
        ...form.view,
        setting: {}
      }
    }
    edit = new EditController({...host, app: testApp})
    result = edit.setEditor(options, form)
    expect(result).to.true;

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              shippingdate: "2021.21.21",
              shipping_place_id: 123
            },
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              trans: [...testApp.store.data[APP_MODULE.EDIT].dataset.trans],
              shiptemp: [
                { id: "1-5", item_product: "Wheel", movement_product: "Wheel", product_id: 5, sqty: "30" }
              ]
            },
          }
        }
      }
    }
    testApp.store.data[APP_MODULE.EDIT].dataset.trans[0] = {
      ...InvoiceData.args.current.item,
      direction: 69
    }
    edit = new EditController({...host, app: testApp})
    result = edit.setEditor(options, form)
    expect(result).to.true;

  })

  it('setEditorItem', () => {
    let testApp = {
      ...app,
      modules: {
        ...app.modules,
        initItem: () => ({}),
      },
      getSetting: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...Item.args,
          }
        }
      }
    }
    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            panel: {
              ...testApp.store.data[APP_MODULE.EDIT].template.options.panel,
            }
          }
        }
      }
    }
    let options = {
      fkey: "item", id: 18
    }
    let edit = new EditController({...host, app: testApp})
    edit.setEditorItem(options)
    sinon.assert.callCount(testApp.store.setData, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            panel: {
              ...testApp.store.data[APP_MODULE.EDIT].panel,
              state: "normal"
            }
          }
        }
      }
    }
    options = {
      fkey: "item", id: null
    }
    edit = new EditController({...host, app: testApp})
    edit.setEditorItem(options)
    sinon.assert.callCount(testApp.store.setData, 2);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              invoice_link: [
                { id: 5 }
              ],
              payment_link: [
                { id: 5 }
              ],
              price: [
                { id: 5 }
              ]
            },
          }
        }
      }
    }
    options = {
      fkey: "invoice_link", id: 5
    }
    edit = new EditController({...host, app: testApp})
    edit.setEditorItem(options)
    sinon.assert.callCount(testApp.store.setData, 3);

    options = {
      fkey: "invoice_link", id: null
    }
    edit = new EditController({...host, app: testApp})
    edit.setEditorItem(options)
    sinon.assert.callCount(testApp.store.setData, 4);

    options = {
      fkey: "payment_link", id: 5
    }
    edit = new EditController({...host, app: testApp})
    edit.setEditorItem(options)
    sinon.assert.callCount(testApp.store.setData, 5);

    options = {
      fkey: "payment_link", id: null, link_field: "link_field"
    }
    edit = new EditController({...host, app: testApp})
    edit.setEditorItem(options)
    sinon.assert.callCount(testApp.store.setData, 6);

    options = {
      fkey: "price", id: 5
    }
    edit = new EditController({...host, app: testApp})
    edit.setEditorItem(options)
    sinon.assert.callCount(testApp.store.setData, 7);

    options = {
      fkey: "price", id: null
    }
    edit = new EditController({...host, app: testApp})
    edit.setEditorItem(options)
    sinon.assert.callCount(testApp.store.setData, 8);

  })

  it('setFieldvalue', () => {
    let testApp = {
      ...app,
      modules: {
        ...app.modules,
        initItem: () => ({}),
      },
      store: {
        ...app.store,
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }
    const edit = new EditController({...host, app: testApp})
    let result = edit.setFieldvalue(testApp.store.data[APP_MODULE.EDIT].current.fieldvalue, 
      "trans_transcast", testApp.store.data[APP_MODULE.EDIT].current.item.id, null, "edit")
    expect(result[1].value).to.equal("edit");

    result = edit.setFieldvalue(testApp.store.data[APP_MODULE.EDIT].current.fieldvalue, 
      "trans_transcast", testApp.store.data[APP_MODULE.EDIT].current.item.id, null, null)
    expect(result.length).to.equal(testApp.store.data[APP_MODULE.EDIT].current.fieldvalue.length);

    testApp = {
      ...testApp,
      modules: {
        ...testApp.modules,
        initItem: () => ({
          id: null, fieldname: null, ref_id: null, value: null, notes: null, deleted: 0
        }),
      }
    }
    result = edit.setFieldvalue(testApp.store.data[APP_MODULE.EDIT].current.fieldvalue, 
      "test", testApp.store.data[APP_MODULE.EDIT].current.item.id, null, "new")
    expect(result[result.length-1].value).to.equal("new");

    result = edit.setFieldvalue(testApp.store.data[APP_MODULE.EDIT].current.fieldvalue, 
      "test", testApp.store.data[APP_MODULE.EDIT].current.item.id, "default", null)
    expect(result[result.length-1].value).to.equal("default");

  })

  it('setFormActions', async () => {
    let testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: {} })
          }
        }),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }

    let edit = new EditController({...host, app: testApp})
    edit.checkEditor = sinon.spy()
    edit.setFormActions({ params: { action: ACTION_EVENT.LOAD_EDITOR } })
    sinon.assert.callCount(edit.checkEditor, 1);

    edit.setFormActions({ params: { action: ACTION_EVENT.LOAD_EDITOR, ntype: "trans", ttype: "invoice", id: 5 } })
    sinon.assert.callCount(edit.checkEditor, 2);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            panel: {
              ...testApp.store.data[APP_MODULE.EDIT].template.options.panel,
            }
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.checkEditor = sinon.spy()
    edit.setEditorItem = sinon.spy()
    edit.deleteEditorItem = sinon.spy()
    edit.setFormActions({ 
      params: { action: ACTION_EVENT.NEW_EDITOR_ITEM, fkey: "item" } })
    sinon.assert.callCount(edit.checkEditor, 1);

    edit.setFormActions({ 
      params: { action: ACTION_EVENT.EDIT_EDITOR_ITEM, fkey: "item", row: { id: 5 } } })
    sinon.assert.callCount(edit.setEditorItem, 1);

    edit.setFormActions({ 
      params: { action: ACTION_EVENT.DELETE_EDITOR_ITEM, fkey: "item", table: "item", row: { id: 5 } } })
    sinon.assert.callCount(edit.deleteEditorItem, 1);

    edit.setFormActions({ 
      params: { action: ACTION_EVENT.LOAD_SHIPPING } })
    sinon.assert.callCount(edit.checkEditor, 2);

    edit.setFormActions({ 
      params: { action: ACTION_EVENT.LOAD_SHIPPING, ntype: "trans", ttype: "order", id: 5 } })
    sinon.assert.callCount(edit.checkEditor, 3);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              shiptemp: [
                { id: 1 }
              ],
              shipping_items: []
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.module.modalShipping = (prm)=>(prm)
    edit.setEditor = sinon.spy()
    edit.showStock = sinon.spy()
    edit.exportQueue = sinon.spy()

    edit.setFormActions({ 
      params: { action: ACTION_EVENT.ADD_SHIPPING_ROW, row: { edited: true } } })
    sinon.assert.callCount(edit.setEditor, 1);

    edit.setFormActions({ 
      params: { action: ACTION_EVENT.ADD_SHIPPING_ROW,
        row: { edited: false, item_id: 1, product_id: 1, product: "product", 
          partnumber: "partnumber", partname: "partname", unit: "unit",
          diff: 0, qty: 1, tqty: 1 } } })
    sinon.assert.callCount(edit.setEditor, 2);

    edit.setFormActions({ 
      params: { action: ACTION_EVENT.SHOW_SHIPPING_STOCK,
        row: { product_id: 1, partnumber: "partnumber", partname: "partname" } } })
    sinon.assert.callCount(edit.showStock, 1);

    edit.setFormActions({ 
      params: { action: ACTION_EVENT.EDIT_SHIPPING_ROW,
        row: { id: 1, item_id: 1, product_id: 1, product: "product", 
        partnumber: "partnumber", partname: "partname", unit: "unit", batch_no: "batch_no",
        diff: 0, qty: 1, tqty: 1, oqty: 1 } }, 
      ref: { requestUpdate: sinon.spy() } })
    sinon.assert.callCount(testApp.store.setData, 4);

    edit.setFormActions({ 
      params: { action: ACTION_EVENT.DELETE_SHIPPING_ROW, row: { id: 1 } } })
    sinon.assert.callCount(edit.setEditor, 3);

    edit.setFormActions({ 
      params: { action: ACTION_EVENT.EXPORT_QUEUE_ITEM, row: { 
        id: 1, ref_id: 1, reportkey: "reportkey", refnumber: "refnumber", copies: 1, typename: "typename" } } })
    sinon.assert.callCount(edit.exportQueue, 1);

    edit.setFormActions({ params: { action: "" } })

  })

  it('setLink', async () => {
    let testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            current: {
              ...app.store.data[APP_MODULE.EDIT].current,
              transtype: "cash",
              extend: {
                id: 1
              },
              form: {
                id: 1
              }
            },
            dataset: {
              ...app.store.data[APP_MODULE.EDIT].dataset,
              payment_link: []
            },
          }
        }
      }
    }

    let edit = new EditController({...host, app: testApp})
    edit.checkEditor = sinon.spy()
    edit.setLink("payment_link", "ref_id_1")
    sinon.assert.callCount(edit.checkEditor, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              transtype: "invoice",
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.checkEditor = sinon.spy()
    edit.setLink("payment_link", "ref_id_1")
    sinon.assert.callCount(edit.checkEditor, 1);

  })

  it('setPassword', async () => {
    const testApp = {
      ...app,
      currentModule: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            current: {
              ...app.store.data[APP_MODULE.EDIT].current,
              type: "employee"
            },
            dataset: {
              ...app.store.data[APP_MODULE.EDIT].dataset,
              employee: [
                { username: "admin" }
              ],
            },
          }
        }
      }
    }

    const edit = new EditController({...host, app: testApp})
    edit.setPassword("admin")
    sinon.assert.callCount(testApp.currentModule, 1);

    edit.setPassword()
    sinon.assert.callCount(testApp.currentModule, 2);

  })

  it('setPattern', async () => {
    let testApp = {
      ...app,
      modules: {
        ...app.modules,
        initItem: () => ({}),
      },
      requestData: sinon.spy(async () => ([1])),
      resultError: sinon.spy(),
      showToast: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: { value: "value" } })
          }
        }),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }
    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              pattern: [
                { id: 1, description: "first pattern", transtype: 55, 
                  notes: "pattern text", defpattern: 0, deleted: 0 },
                { id: 2, description: "default pattern", transtype: 55, 
                  notes: null, defpattern: 1, deleted: 0 },
              ]
            },
          }
        }
      }
    }

    let edit = new EditController({...host, app: testApp})
    edit.checkEditor = sinon.spy()
    await edit.setPattern({ key: "default" })
    sinon.assert.callCount(edit.checkEditor, 1);

    await edit.setPattern({ key: "save" })
    sinon.assert.callCount(testApp.requestData, 2);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} })),
      resultError: sinon.spy(),
    }
    edit = new EditController({...host, app: testApp})
    await edit.setPattern({ key: "save" })
    sinon.assert.callCount(testApp.resultError, 1);

    edit = new EditController({...host, app: testApp})
    await edit.setPattern({ key: "delete" })
    sinon.assert.callCount(testApp.resultError, 2);

    edit = new EditController({...host, app: testApp})
    await edit.setPattern({ key: "load", ref: { _value: "" } })
    sinon.assert.callCount(testApp.store.setData, 13);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ([])),
      store: {
        ...testApp.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL, data: { value: "value" } })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: { value: "value" } })
          }
        })
      }
    }
    edit = new EditController({...host, app: testApp})
    edit.checkEditor = sinon.spy()
    await edit.setPattern({ key: "new" })
    sinon.assert.callCount(testApp.requestData, 2);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ([])),
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            dataset: {
              ...testApp.store.data[APP_MODULE.EDIT].dataset,
              pattern: []
            },
          }
        }
      }
    }
    
    edit = new EditController({...host, app: testApp})
    edit.checkEditor = sinon.spy()
    await edit.setPattern({ key: "new" })
    sinon.assert.callCount(testApp.requestData, 2);
   
    const result = await edit.setPattern({ key: "load" })
    expect(result).to.true;

    await edit.setPattern({ key: "save" })
    sinon.assert.callCount(testApp.store.setData, 9);

    await edit.setPattern({ key: "delete" })
    sinon.assert.callCount(testApp.store.setData, 12);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ([1])),
      showToast: sinon.spy(),
      store: {
        ...testApp.store,
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.setPattern({ key: "new" })
    sinon.assert.callCount(testApp.showToast, 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} })),
      resultError: sinon.spy(),
    }
    edit = new EditController({...host, app: testApp})
    await edit.setPattern({ key: "new" })
    sinon.assert.callCount(testApp.resultError, 1);

    testApp = {
      ...testApp,
      showToast: sinon.spy(),
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.EDIT]: {
            ...testApp.store.data[APP_MODULE.EDIT],
            current: {
              ...testApp.store.data[APP_MODULE.EDIT].current,
              template: ""
            },
          }
        }
      }
    }
    edit = new EditController({...host, app: testApp})
    await edit.setPattern({ key: "default" })
    sinon.assert.callCount(testApp.showToast, 1);

  })

  it('shippingAddAll', () => {
    const testApp = {
      ...app,
      store: {
        ...app.store,
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            dataset: {
              ...app.store.data[APP_MODULE.EDIT].dataset,
              shiptemp: [],
              shipping_items_: [
                { description: "Wheel", item_id: 1, partname: "Wheel", partnumber: "DMPROD/00005",
                  pgroup: "0", product: "DMPROD/00005 | Wheel", product_id: "5", 
                  qty: "40", unit: "piece", tqty: 1 },
                { description: "Wheel", item_id: 1, partname: "Wheel", partnumber: "DMPROD/00005",
                  pgroup: "0", product: "DMPROD/00005 | Wheel", product_id: "5", 
                  qty: "40", unit: "piece", tqty: 1, edited: true, diff: 0 }
              ]
            },
          }
        }
      }
    }
    const edit = new EditController({...host, app: testApp})
    edit.setEditor = sinon.spy()
    edit.shippingAddAll()
    sinon.assert.callCount(edit.setEditor, 1);

  })

  it('setPattern', async () => {
    let testApp = {
      ...app,
      requestData: sinon.spy(async () => ({ 
        stock: [
          { id: 1, partnumber: "partnumber", description: "description", unit: "unit",
            warehouse: "warehouse", sqty: 1 }
          ]
      })),
      resultError: sinon.spy(),
      getSetting: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL })
          }
        }),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }
    let edit = new EditController({...host, app: testApp})
    edit.module.modalStock = (prm)=>(prm)
    await edit.showStock({ partnumber: "partnumber", partname: "partname" })
    sinon.assert.callCount(testApp.store.setData, 2);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ stock: [] })),
      resultError: sinon.spy(),
      showToast: sinon.spy(),
    }
    edit = new EditController({...host, app: testApp})
    edit.module.modalStock = (prm)=>(prm)
    await edit.showStock({ partnumber: "partnumber", partname: "partname" })
    sinon.assert.callCount(testApp.showToast, 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} })),
      resultError: sinon.spy(),
    }
    edit = new EditController({...host, app: testApp})
    edit.module.modalStock = (prm)=>(prm)
    await edit.showStock({ partnumber: "partnumber", partname: "partname" })
    sinon.assert.callCount(testApp.resultError, 1);

  })

  it('tableValues', () => {
    const testApp = {
      ...app,
      modules: {
        ...app.modules,
        initItem: () => ({
          id: null, fieldname: null, ref_id: null, value: null, notes: null, deleted: 0
        }),
      },
      store: {
        ...app.store,
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }
    const edit = new EditController({...host, app: testApp})   
    const result = edit.tableValues("test", { id: 1, missing: "value" })
    expect(result.id).to.equal(1);

  })

  it('transCopy', () => {
    const testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: { value: "value" } })
          }
        }),
        data: {
          ...app.store.data,
          [APP_MODULE.EDIT]: {
            ...app.store.data[APP_MODULE.EDIT],
            ...InvoiceData.args,
          }
        }
      }
    }

    const edit = new EditController({...host, app: testApp})
    edit.checkEditor = sinon.spy()
    edit.transCopy()
    sinon.assert.callCount(edit.checkEditor, 1);

    edit.transCopy("create")
    sinon.assert.callCount(edit.checkEditor, 2);

  })

  it('setModule', () => {
    const edit = new EditController({...host, app})
    edit.setModule({})
    expect(edit.module).to.exist
  })

  it('checkSubtype', () => {
    let value = checkSubtype("customer", null, { custtype: 1 })
    expect(value).to.true

    value = checkSubtype("customer", 1, { custtype: 1 })
    expect(value).to.true

    value = checkSubtype("missing", 1, { custtype: 1 })
    expect(value).to.true
  })

})