import sinon from 'sinon'
import { expect } from '@open-wc/testing';

import { SearchController, getFilterWhere, defaultFilterValue } from './SearchController.js'
import { store as storeConfig } from '../config/app.js'
import { BROWSER_EVENT, SIDE_EVENT, APP_MODULE, MODAL_EVENT } from '../config/enums.js'
import { Queries } from './Queries.js'

const host = { 
  addController: ()=>{},
  inputBox: (prm)=>(prm),
  queries: Queries({ getText: (key)=>key })
}
const store = {
  data: {
    ...storeConfig,
    [APP_MODULE.LOGIN]: {
      ...storeConfig[APP_MODULE.LOGIN],
      data: {
        engine: "sqlite",
        menuCmds: []
      }
    }
  },
  setData: sinon.spy(),
  showToast: sinon.spy(),
  msg: (value)=>value,
  getSetting: sinon.spy(),
}
const app = {
  store,
  requestData: () => ({}),
  resultError: sinon.spy(),
  loadBookmark: () => ({}),
  showHelp: sinon.spy(),
  getSql: () =>({
    sql: "",
    prmCount: 1
  }),
  getDataFilter: ()=>[[]],
  getUserFilter: ()=>({ where: [[]], params: [[]] })
}

describe('SearchController', () => {
  it('onSideEvent', () => {
    const testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy()
      },
      getSql: () =>({
        sql: "",
        prmCount: 1
      })
    }

    const search = new SearchController(host, testApp)
    search.onSideEvent({ key: SIDE_EVENT.CHANGE, data: { fieldname: "", value: "" } })
    sinon.assert.callCount(testApp.store.setData, 1);

    search.onSideEvent({ key: SIDE_EVENT.SEARCH_BROWSER, data: { value: "customer" } })
    sinon.assert.callCount(testApp.store.setData, 2);

    search.onSideEvent({ key: SIDE_EVENT.SEARCH_QUICK, data: { value: "" } })
    sinon.assert.callCount(testApp.store.setData, 4);

    search.onSideEvent({ key: SIDE_EVENT.CHECK_EDITOR, data: {} })
    sinon.assert.callCount(testApp.store.setData, 5);

    search.onSideEvent({})

  })

  it('onBrowserEvent', async () => {
    let testApp = {
      ...app,
      store: {
        ...app.store,
        data: {
          ...app.store.data,
          [APP_MODULE.SEARCH]: {
            ...app.store.data[APP_MODULE.SEARCH],
            vkey: "transmovement", 
            view: "InventoryView", 
            filters: {
              InventoryView: [
                { wheretype: "where", filtertype: "===" },
                { wheretype: "having", filtertype: "===" }
              ]
            },
            columns: {
              InventoryView: {}
            },
            result: [
              { description: "Big product", export_partnumber: "DMPROD/00001", 
                export_sqty: "2", qt: "", fieldname: "deffield", export_deffield_value: 3 }
            ]
          }
        },
        setData: sinon.spy()
      },
      saveBookmark: sinon.spy(),
      saveToDisk: sinon.spy(),
      showHelp: sinon.spy()
    }

    let search = new SearchController(host, testApp)
    search.onBrowserEvent({ key: BROWSER_EVENT.CHANGE, data: { fieldname: "", value: "" } })
    sinon.assert.callCount(testApp.store.setData, 1);

    search.onBrowserEvent({ key: BROWSER_EVENT.ADD_FILTER, data: {} })
    sinon.assert.callCount(testApp.store.setData, 2);

    search.onBrowserEvent({ key: BROWSER_EVENT.BOOKMARK_SAVE, data: {} })
    sinon.assert.callCount(testApp.saveBookmark, 1);

    await search.onBrowserEvent({ key: BROWSER_EVENT.BROWSER_VIEW, data: {} })
    sinon.assert.callCount(testApp.store.setData, 3);

    search.onBrowserEvent({ key: BROWSER_EVENT.CURRENT_PAGE, data: { value: 1 } })
    sinon.assert.callCount(testApp.store.setData, 4);

    search.onBrowserEvent({ key: BROWSER_EVENT.DELETE_FILTER, data: { value: 0 } })
    sinon.assert.callCount(testApp.store.setData, 5);

    search.onBrowserEvent({ key: BROWSER_EVENT.EDIT_FILTER, data: { index: 0, fieldname: "value", value: "value" } })
    sinon.assert.callCount(testApp.store.setData, 6);

    search.onBrowserEvent({ key: BROWSER_EVENT.EXPORT_RESULT, data: { value: [ "description" ] } })
    sinon.assert.callCount(testApp.saveToDisk, 1);

    search.onBrowserEvent({ key: BROWSER_EVENT.EDIT_CELL, data: { 
      fieldname: "id", value: "ntype/ttype/id", row: { form: "form", form_id: 1 } } })
    sinon.assert.callCount(testApp.store.setData, 7);

    search.onBrowserEvent({ key: BROWSER_EVENT.SET_COLUMNS, data: { 
      fieldname: "description", value: true } })
    sinon.assert.callCount(testApp.store.setData, 8);
    search.onBrowserEvent({ key: BROWSER_EVENT.SET_COLUMNS, data: { 
      fieldname: "description", value: false } })
    sinon.assert.callCount(testApp.store.setData, 9);

    search.onBrowserEvent({ key: BROWSER_EVENT.SET_FORM_ACTIONS, data: {} })
    sinon.assert.callCount(testApp.store.setData, 10);

    await search.onBrowserEvent({ key: BROWSER_EVENT.SHOW_BROWSER, data: { value: "transmovement", view: "InventoryView" } })
    sinon.assert.callCount(testApp.store.setData, 13);

    search.onBrowserEvent({ key: BROWSER_EVENT.SHOW_HELP, data: { value: "" } })
    sinon.assert.callCount(testApp.showHelp, 1);

    search.onBrowserEvent({ key: BROWSER_EVENT.SHOW_TOTAL, 
      data: { 
        fields: { sqty: {} }, 
        totalFields: { totalFields: { sqty: 0, qt: 0 }, totalLabels: { sqty: "Quantity", qt: "NA" }, count: 2 } 
      }, ref: { modalTotal: sinon.spy() } })
    sinon.assert.callCount(testApp.store.setData, 14);

    search.onBrowserEvent({ key: BROWSER_EVENT.SHOW_TOTAL, 
      data: { 
        fields: { deffield_value: {} }, 
        totalFields: { totalFields: { deffield: 0 }, totalLabels: { deffield: "Deffield" }, count: 1 } 
      }, ref: { modalTotal: sinon.spy() } })
    sinon.assert.callCount(testApp.store.setData, 15);

    search.onBrowserEvent({})

    testApp = {
      ...testApp,
      requestData: sinon.spy(()=>({ error: {} })),
      resultError: sinon.spy(),
      store: {
        ...app.store,
        data: {
          ...app.store.data,
          [APP_MODULE.SEARCH]: {
            ...app.store.data[APP_MODULE.SEARCH],
            vkey: "customer", 
            view: "CustomerView",
            filters: {
              CustomerView: [
                { wheretype: "where", filtertype: "===" }
              ]
            },
          },
          deffield: [
            { fieldname: "bool", fieldtype: "bool", sqlstr: "" }
          ]
        }
      }
    }

    search = new SearchController(host, testApp)
    await search.onBrowserEvent({ key: BROWSER_EVENT.BROWSER_VIEW, data: {} })
    sinon.assert.callCount(testApp.resultError, 1);

    search.onBrowserEvent({ key: BROWSER_EVENT.EDIT_FILTER, data: { index: 0, fieldname: "fieldname", value: "bool" } })
    sinon.assert.callCount(testApp.store.setData, 1);

    search.onBrowserEvent({ key: BROWSER_EVENT.EDIT_FILTER, data: { index: 0, fieldname: "fieldname", value: "custname" } })
    sinon.assert.callCount(testApp.store.setData, 2);
  
  })

  it('onModalEvent', async () => {
    let testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy()
      },
      quickSearch: sinon.spy(()=>({ result: {} }))
    }

    let search = new SearchController(host, testApp)
    search.onModalEvent({ key: MODAL_EVENT.CANCEL, data: {} })
    sinon.assert.callCount(testApp.store.setData, 1);

    search.onModalEvent({ key: MODAL_EVENT.CURRENT_PAGE, data: { value: 1 } })
    sinon.assert.callCount(testApp.store.setData, 2);

    await search.onModalEvent({ key: MODAL_EVENT.SEARCH, data: { value: "" } })
    sinon.assert.callCount(testApp.store.setData, 3);

    testApp = {
      ...testApp,
      quickSearch: sinon.spy(()=>({ error: {} })),
      resultError: sinon.spy()
    }

    search = new SearchController(host, testApp)
    await search.onModalEvent({ key: MODAL_EVENT.SEARCH, data: { value: "" } })
    sinon.assert.callCount(testApp.resultError, 1);

    search.onModalEvent({ key: MODAL_EVENT.SELECTED, data: { value: { id: "ntype/ttype/1" } }, ref: { modalServer: (prm)=>({prm}) } })
    sinon.assert.callCount(testApp.store.setData, 4);

    search.onModalEvent({ key: MODAL_EVENT.SELECTED, data: { value: { id: "servercmd//1" } }, ref: { modalServer: (prm)=>({prm}) } })
    sinon.assert.callCount(testApp.store.setData, 4);

    search.onModalEvent({})

  })

  it('hostConnected', async () => {
    const testApp = {
      ...app,
      store: {
        ...app.store,
        data: {
          ...app.store.data,
          current: {
            ...app.store.data.current,
            content: { [APP_MODULE.SEARCH]: [ 
              "transmovement", "InventoryView", 
              { show_header: true, show_columns: true } 
            ] }
          }
        },
        setData: sinon.spy()
      },
      requestData: sinon.spy(()=>({ error: {} })),
      resultError: sinon.spy()
    }

    const search = new SearchController(host, testApp)
    await search.hostConnected()
    sinon.assert.callCount(testApp.resultError, 1);

  })

  it('getFilterWhere', () => {
    
    let result = getFilterWhere({ 
      filtertype: "==N", fieldtype: "string", sqlstr: "" })
    expect(result[1][1][2]).to.equal("is null")
    result = getFilterWhere({ 
      filtertype: "==N", fieldtype: "integer", sqlstr: "" })
    expect(result[2]).to.equal("is null")

    result = getFilterWhere({ 
      filtertype: "!==", fieldtype: "string", sqlstr: "" })
    expect(result[1][2]).to.equal("{CCS}{JOKER}{SEP}lower(?){SEP}{JOKER}{CCE}")
    result = getFilterWhere({ 
      filtertype: "!==", fieldtype: "integer", sqlstr: "" })
    expect(result[1][2]).to.equal("?")

    result = getFilterWhere({ 
      filtertype: ">==", fieldtype: "string", sqlstr: "" })
    expect(result[1][2]).to.equal("?")

    result = getFilterWhere({ 
      filtertype: "<==", fieldtype: "string", sqlstr: "" })
    expect(result[1][2]).to.equal("?")

    result = getFilterWhere({ 
      filtertype: "===", fieldtype: "string", sqlstr: "" })
    expect(result[1][2]).to.equal("{CCS}{JOKER}{SEP}lower(?){SEP}{JOKER}{CCE}")
    result = getFilterWhere({ 
      filtertype: "===", fieldtype: "integer", sqlstr: "" })
    expect(result[1][2]).to.equal("?")
    
  })

  it('defaultFilterValue', () => {
    let result = defaultFilterValue("date")
    expect(result).to.exist
    result = defaultFilterValue("integer")
    expect(result).to.equal(0)
    result = defaultFilterValue("string")
    expect(result).to.equal("")
  })

  it('showServerCmd', () => {
    const testApp = {
      ...app,
      store: {
        ...app.store,
        data: {
          ...app.store.data,
          [APP_MODULE.LOGIN]: {
            ...app.store.data[APP_MODULE.LOGIN],
            data: {
              menuCmds: [
                { address: null, description: "Server function example", funcname: "nextNumber",
                  icon: null, id: 1, menukey: "nextNumber", method: 130, methodName: "post", modul: null },
                { address: "", description: "Internet URL example", funcname: "",
                  icon: null, id: 2, menukey: "google", method: 129, methodName: "get", modul: null },
                { address: "address", description: "Server function example", funcname: null,
                  icon: null, id: 3, menukey: "nextNumber2", method: 130, methodName: "post", modul: null },
              ],
              menuFields: [
                { description: "Code", fieldname: "numberkey", fieldtype: 38, fieldtypeName: "string",
                  id: 1, menu_id: 1, orderby: 0 },
                { description: "Stepping", fieldname: "step", fieldtype: 33, fieldtypeName: "bool",
                  id: 2, menu_id: 1, orderby: 1 },
                { description: "float", fieldname: "float", fieldtype: 34, fieldtypeName: "float",
                  id: 3, menu_id: 1, orderby: 2 },
                { description: "google search", fieldname: "q", fieldtype: 38, fieldtypeName: "string",
                  id: 3, menu_id: 2, orderby: 0 }
              ]
            }
          }
        },
        setData: (key, data) => {
          // debugger;
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL })
            data.modalForm.onEvent.onModalEvent({ 
              key: MODAL_EVENT.OK, 
              data: { 
                values: { numberkey: "", step: true }, 
                cmd: testApp.store.data[APP_MODULE.LOGIN].data.menuCmds[0],
                fields: testApp.store.data[APP_MODULE.LOGIN].data.menuFields 
              } 
            })
            data.modalForm.onEvent.onModalEvent({ 
              key: MODAL_EVENT.OK, 
              data: { 
                values: { numberkey: "", step: false }, 
                cmd: testApp.store.data[APP_MODULE.LOGIN].data.menuCmds[2],
                fields: testApp.store.data[APP_MODULE.LOGIN].data.menuFields 
              } 
            })
            data.modalForm.onEvent.onModalEvent({ 
              key: MODAL_EVENT.OK, 
              data: { 
                values: { q: "" }, 
                cmd: testApp.store.data[APP_MODULE.LOGIN].data.menuCmds[1],
                fields: testApp.store.data[APP_MODULE.LOGIN].data.menuFields 
              } 
            })
          }
        }
      },
      requestData: () => ({}),
      resultError: sinon.spy(),
      request: () => ({ error: {} }),
      showToast: sinon.spy(),
    }
    const search = new SearchController(host, testApp)
    search.showServerCmd(1, { modalServer: (prm)=>(prm)})

  })

})