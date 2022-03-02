import { queryByAttribute } from '@testing-library/react'
import ReactDOM from 'react-dom';
import update from 'immutability-helper';

import { store as app_store  } from 'config/app'
import { getSql, appActions } from 'containers/App/actions'
import { InitItem, Validator } from 'containers/Controller/Validator'
import { settingActions } from './actions'

jest.mock("containers/App/actions");
jest.mock("containers/Controller/Validator");
const getById = queryByAttribute.bind(null, 'id');

const store = update(app_store, {$merge: {
  login: {
    data: {
      employee: {
        id: 1, empnumber: 'admin', username: 'admin', usergroup: 114,
        usergroupName: 'admin'
      },
      engine: "sqlite3"
    }
  },
  setting: {
    type: "setting",
    ntype: "setting",
    dataset: {
      transfilter: [{ id: 1, groupvalue: "all" }],
      usergroup_view: [{ id: 1, transfilter: null }],
      setting_view: [{ id: 1, fieldtype: "urlink" }]
    }
  }
}})

describe('settingActions', () => {
  beforeEach(() => {
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ([1])),
      resultError: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      showToast: jest.fn()
    })
    getSql.mockReturnValue({
      sql: "",
      prmCount: 1
    })
    InitItem.mockReturnValue(
      ()=>({fieldname: "value"})
    )
    Validator.mockReturnValue(
      ()=>({})
    )
    Object.defineProperty(global.window, 'open', { value: jest.fn() });
    Object.defineProperty(global.window, 'localStorage', {
      value: {
        getItem: jest.fn(),
        setItem: jest.fn()
      }
    });
  });
  
  afterEach(() => {
    jest.clearAllMocks();
  });

  it('tableValues', () => {
    const setData = jest.fn((key, data, callback)=>{
      if(callback){callback()}
    })
    let it_store = update(store, {setting: {$merge:{
      current: {}
    }}})
    let values = settingActions(it_store, setData).tableValues("type", { fieldname: "value", missing: "values" })
    expect(Object.keys(values).length).toBe(1)
  })

  it('setSettingForm', () => {
    const setData = jest.fn((key, data, callback)=>{
      if(callback){callback()}
    })
    let it_store = update(store, {setting: {$merge:{
      audit: "readonly"
    }}})
    settingActions(it_store, setData).setSettingForm(null)
    expect(setData).toHaveBeenCalledTimes(1)
    it_store = update(store, {setting: {$merge:{
      audit: "all"
    }}})
    settingActions(it_store, setData).setSettingForm(null)
    expect(setData).toHaveBeenCalledTimes(2)
    it_store = update(store, {setting: {$merge:{
      type: "usergroup",
      audit: "all"
    }}})
    settingActions(it_store, setData).setSettingForm(1, it_store.setting)
    expect(setData).toHaveBeenCalledTimes(3)
    it_store = update(store, {setting: {$merge:{
      ntype: "fieldvalue",
    }}})
    settingActions(it_store, setData).setSettingForm(1)
    expect(setData).toHaveBeenCalledTimes(4)
    it_store = update(it_store, {setting: {dataset: {$merge:{
      setting_view: [{ id: 1, fieldtype: "valuelist", valuelist:"a|b|c", notes: "" }]
    }}}})
    settingActions(it_store, setData).setSettingForm(1)
    expect(setData).toHaveBeenCalledTimes(5)
  })

  it('setPasswordForm', () => {
    const setData = jest.fn()
    settingActions(store, setData).setPasswordForm("username")
    expect(setData).toHaveBeenCalledTimes(2)
  })

  it('changePassword', () => {
    const setData = jest.fn()
    //missing username
    let it_store = update(store, {setting: {$merge:{
      current: {
        form: {
          username: ""
        }
      }
    }}})
    settingActions(it_store, setData).changePassword()
    // not match
    it_store = update(store, {setting: {$merge:{
      current: {
        form: {
          username: "username",
          password_1: "1",
          password_2: "2"
        }
      }
    }}})
    settingActions(it_store, setData).changePassword()
    // ok
    it_store = update(store, {setting: {$merge:{
      current: {
        form: {
          username: "username",
          password_1: "password",
          password_2: "password"
        }
      }
    }}})
    settingActions(it_store, setData).changePassword()
    //error
    appActions.mockReturnValue({
      requestData: jest.fn(async () => ({ error: {} })),
      resultError: jest.fn(),
    })
    settingActions(it_store, setData).changePassword()
  })

  it('setProgramForm', () => {
    const setData = jest.fn()
    settingActions(store, setData).setProgramForm()
    expect(setData).toHaveBeenCalledTimes(2)
  })

  it('setSettingData', () => {
    const setData = jest.fn()
    settingActions(store, setData).setSettingData({ type: "setting", dataset: {} })
    expect(setData).toHaveBeenCalledTimes(2)
    settingActions(store, setData).setSettingData({ 
      type: "log", dataset: { log_view: [] }, id: 1 
    })
    expect(setData).toHaveBeenCalledTimes(5)
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ({})),
      resultError: jest.fn(),
      getAuditFilter: jest.fn(() => (["update",1])),
      showToast: jest.fn()
    })
    settingActions(store, setData).setSettingData({ type: "usergroup", dataset: {}, id: 1 })
    expect(setData).toHaveBeenCalledTimes(8)
  })

  
  it('loadSetting', () => {
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ({ placetype: [ { id: 1, groupvalue: "warehouse" } ] })),
      resultError: jest.fn(),
      getAuditFilter: jest.fn(() => (["update",1])),
      showToast: jest.fn()
    })
    const setData = jest.fn()
    settingActions(store, setData).loadSetting({ type: "setting" })
    settingActions(store, setData).loadSetting({ type: "usergroup" })
    settingActions(store, setData).loadSetting({ type: "usergroup", id: null })
    settingActions(store, setData).loadSetting({ type: "template", id: 1 })
    settingActions(store, setData).loadSetting({ type: "place", id: null })
    settingActions(store, setData).loadSetting({ type: "log" })
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ({ error: {} })),
      resultError: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      showToast: jest.fn()
    })
    getSql.mockReturnValue({
      sql: "",
      prmCount: 0
    })
    settingActions(store, setData).loadSetting({ type: "setting", id: 1 })
  })

  it('saveSetting', async () => {
    const setData = jest.fn()
    let it_store = update(store, {
      setting: {$merge:{
        current: {
          form: {
            id: null,
            fieldname: "value",
            transfilter: 1,
            translink: null
          }
        }
      }},
      login: {$merge: {
        data: {
          groups: [
            { id: 1, groupvalue: "all" }
          ]
        }
      }}
    })
    let setting = await settingActions(it_store, setData).saveSetting()
    expect(setting).toBeDefined()
    it_store = update(it_store, {setting: {$merge:{
      type: "usergroup",
      ntype: "usergroup"
    }}})
    setting = await settingActions(it_store, setData).saveSetting()
    expect(setting).toBeDefined()
    it_store = update(it_store, {
      setting: {$merge:{
        current: {
          form: {
            id: 1,
            fieldname: "value",
            transfilter: 1,
            translink: 2
          }
        }
      }},
      login: {$merge: {
        data: {
          groups: [
            { id: 1, groupname: "transfilter", groupvalue: "update" },
            { id: 2, groupname: "nervatype", groupvalue: "groups" },
          ]
        }
      }}
    })
    setting = await settingActions(it_store, setData).saveSetting()
    expect(setting).toBeDefined()
    Validator.mockReturnValue(
      ()=>({ error: {} })
    )
    setting = await settingActions(it_store, setData).saveSetting()
    Validator.mockReturnValue(
      ()=>({})
    )
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ({ error: {} })),
      resultError: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      showToast: jest.fn()
    })
    setting = await settingActions(it_store, setData).saveSetting()
    expect(setting).toBeNull()
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path) => {
        if(path === "/link"){
          return { error: {} }
        }
        return {}
      }),
      resultError: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      showToast: jest.fn()
    })
    setting = await settingActions(it_store, setData).saveSetting()
    expect(setting).toBeNull()
  })

  it('deleteSetting', () => {
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
    let it_store = update(store, {setting: {$merge:{
    }}})
    settingActions(it_store, setData).deleteSetting({ id: 1 })
    expect(setData).toHaveBeenCalledTimes(3);

    //delete_state
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path) => {
        if(path === "/view"){
          return { state: [{ sco: 1 }]  }
        }
        return {}
      }),
      resultError: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      showToast: jest.fn()
    })
    it_store = update(it_store, {setting: {$merge:{
      type: "currency"
    }}})
    settingActions(it_store, setData).deleteSetting({ id: 1 })
    expect(setData).toHaveBeenCalledTimes(6);

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path) => {
        if(path === "/view"){
          return { state: [{ sco: 0 }]  }
        }
        return {}
      }),
      resultError: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      showToast: jest.fn()
    })
    settingActions(it_store, setData).deleteSetting({ id: 1 })
    expect(setData).toHaveBeenCalledTimes(9);

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path) => {
        if(path === "/view"){
          return { error: {} }
        }
        return {}
      }),
      resultError: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      showToast: jest.fn()
    })
    settingActions(it_store, setData).deleteSetting({ id: 1 })
    expect(setData).toHaveBeenCalledTimes(12);

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path) => {
        if(path === "/view"){
          return { state: [{ sco: 0 }]  }
        }
        return { error: {} }
      }),
      resultError: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      showToast: jest.fn()
    })
    it_store = update(it_store, {setting: {$merge:{
      type: "usergroup"
    }}})
    settingActions(it_store, setData).deleteSetting({ id: 1 })
    expect(setData).toHaveBeenCalledTimes(15);
  })

  it('loadLog', async () => {
    const setData = jest.fn()
    let it_store = update(store, {setting: {$merge:{
      view: {},
      current: {
        form: {
          id: null,
          fromdate: "2022-02-12",
          todate: "",
          empnumber: "",
          logstate: "update",
          nervatype: ""
        }
      }
    }}})
    await settingActions(it_store, setData).loadLog()
    expect(setData).toHaveBeenCalledTimes(0);

    it_store = update(store, {setting: {$merge:{
      view: {},
      current: {
        form: {
          id: null,
          fromdate: "2022-02-12",
          todate: "2022-02-12",
          empnumber: "2022-02-12",
          logstate: "update",
          nervatype: "customer"
        }
      }
    }}})
    await settingActions(it_store, setData).loadLog()
    expect(setData).toHaveBeenCalledTimes(1);

    it_store = update(store, {setting: {$merge:{
      view: {},
      current: {
        form: {
          id: null,
          fromdate: "",
          todate: "",
          empnumber: "",
          logstate: "login",
          nervatype: ""
        }
      }
    }}})
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path) => {
        return { error: {} }
      }),
      resultError: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      showToast: jest.fn()
    })
    await settingActions(it_store, setData).loadLog()
    expect(setData).toHaveBeenCalledTimes(1);

  })

  it('checkSetting', () => {
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
    let it_store = update(store, {
      setting: {$merge:{
        current: {
          form: {
            id: null,
            fieldname: "value",
            transfilter: 1,
            translink: null
          }
        },
        dirty: true
      }},
      login: {$merge: {
        data: {
          groups: [
            { id: 1, groupvalue: "all" }
          ]
        }
      }}
    })
    settingActions(it_store, setData).checkSetting({ type: "setting", id: null }, 'LOAD_SETTING')
    expect(setData).toHaveBeenCalledTimes(4);
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ({ error: {} })),
      resultError: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      showToast: jest.fn()
    })
    settingActions(it_store, setData).checkSetting({ type: "setting", id: null }, 'LOAD_SETTING')
    expect(setData).toHaveBeenCalledTimes(8);

    settingActions(store, setData).checkSetting({ username: "admin" }, 'PASSWORD_FORM')
    expect(setData).toHaveBeenCalledTimes(10);
  })

  it('setViewActions', () => {
    //newItem
    let setData = jest.fn()
    settingActions(store, setData).setViewActions({ action: "newItem" })
    expect(setData).toHaveBeenCalledTimes(0);

    //editItem - default
    settingActions(store, setData).setViewActions({ action: "editItem" }, { id: 1 })
    expect(setData).toHaveBeenCalledTimes(1);

    //editItem - template
    let it_store = update(store, {setting: {$merge:{
      type: "template"
    }}})
    settingActions(it_store, setData).setViewActions({ action: "editItem" }, { id: 1 })
    expect(setData).toHaveBeenCalledTimes(1);

    //editItem - place
    it_store = update(store, {setting: {$merge:{
      type: "place"
    }}})
    settingActions(it_store, setData).setViewActions({ action: "editItem" }, { id: 1 })
    expect(setData).toHaveBeenCalledTimes(2);
    settingActions(it_store, setData).setViewActions({ action: "editItem" })
    expect(setData).toHaveBeenCalledTimes(3);

    //editItem - usergroup
    it_store = update(store, {setting: {$merge:{
      type: "usergroup"
    }}})
    settingActions(it_store, setData).setViewActions({ action: "editItem" }, { id: 1 })
    expect(setData).toHaveBeenCalledTimes(3);
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ({ error: {} })),
      resultError: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      showToast: jest.fn()
    })
    settingActions(it_store, setData).setViewActions({ action: "editItem" }, { id: 1 })
    expect(setData).toHaveBeenCalledTimes(3);

    //editItem - ui_menu
    it_store = update(store, {setting: {$merge:{
      type: "ui_menu"
    }}})
    settingActions(it_store, setData).setViewActions({ action: "editItem" }, { id: 1 })
    expect(setData).toHaveBeenCalledTimes(3);
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ({})),
      resultError: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      showToast: jest.fn()
    })
    settingActions(it_store, setData).setViewActions({ action: "editItem" }, { id: null })
    expect(setData).toHaveBeenCalledTimes(3);

    //deleteItem
    settingActions(store, setData).setViewActions({ action: "deleteItem" }, { id: 1 })
    expect(setData).toHaveBeenCalledTimes(4);

    //editAudit
    setData = jest.fn((key, data, callback)=>{
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
    it_store = update(store, {setting: {$merge:{
      type: "usergroup",
      ntype: "groups",
      audit: "all",
      current: {
        form: {
          deleted: 0, description: 'Demo', groupname: 'usergroup', groupvalue: 'demo',
          id: 105, inactive: 0, lslabel: 'demo', lsvalue: 'Demo', transfilter: 112, translink: null
        }
      },
      dataset: {
        nervatype: [
          {
            deleted: 0, description: null, groupname: 'nervatype', groupvalue: 'audit',
            id: 6, inactive: 0
          },
          {
            deleted: 0, description: null, groupname: 'nervatype', groupvalue: 'customer',
            id: 10, inactive: 0
          },
        ],
        transtype: [
          {
            deleted: 0, description: null, groupname: 'transtype', groupvalue: 'bank',
            id: 66, inactive: 0
          }
        ],
        inputfilter: [
          {
            deleted: 0, description: null, groupname: 'inputfilter', groupvalue: 'all',
            id: 109, inactive: 0
          }
        ],
        menukey: [
          {
            id: 1, menukey: 'nextNumber'
          }
        ],
        reportkey: [
          {
            id: 1, reportkey: 'csv_custpos_en'
          }
        ],
        usergroup_view: [
          {
            deleted: 0, description: 'Demo', groupname: 'usergroup', groupvalue: 'demo',
            id: 105, inactive: 0, lslabel: 'demo', lsvalue: 'Demo', transfilter: 112, translink: null
          }
        ]
      }
    }}})
    settingActions(it_store, setData).setViewActions(
      { action: "editAudit" }, 
      { id: 1, usergroup: 1, nervatype: 6, inputfilter: 109, supervisor: 0, subtype: null })
    expect(setData).toHaveBeenCalledTimes(2);
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ({ error: {} })),
      resultError: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      showToast: jest.fn()
    })
    settingActions(it_store, setData).setViewActions(
      { action: "editAudit" }, 
      { id: null, usergroup: 1, nervatype: 6, inputfilter: 109, supervisor: 0, subtype: null })
    expect(setData).toHaveBeenCalledTimes(4);

    //editMenuField
    it_store = update(store, {setting: {$merge:{
      type: "ui_menu",
      ntype: "ui_menu",
      audit: "all",
      current: {
        form: {
          address: null, description: 'Server function example', funcname: 'nextNumber', icon: null,
          id: 1, lslabel: 'nextNumber', lsvalue: 'Server function example', menukey: 'nextNumber',
          method: 130, modul: null
        }
      },
      dataset: {
        fieldtype: [
          {
            deleted: 0, description: null, groupname: 'fieldtype', groupvalue: 'string',
            id: 38, inactive: 0
          }
        ],
        ui_menufields: [
          {
            description: 'Code (e.g. custnumber)', fieldname: 'numberkey', fieldtype: 38,
            fieldtype_name: 'string', id: 1, menu_id: 1, orderby: 0
          }
        ],
        ui_menu_view: [
          {
            address: null, description: 'Server function example', funcname: 'nextNumber', icon: null,
            id: 1, lslabel: 'nextNumber', lsvalue: 'Server function example',
            menukey: 'nextNumber', method: 130, modul: null
          }
        ],
        method: [
          {
            deleted: 0, description: 'GET', groupname: 'method', groupvalue: 'get',
            id: 129, inactive: 0
          }
        ]
      }
    }}})
    settingActions(it_store, setData).setViewActions(
      { action: "editMenuField" }, 
      { id: 1, description: "Code (e.g. custnumber)", fieldname: "numberkey",
        fieldtype: 38, fieldtype_name: "string", menu_id: 1, orderby: 0 })
    expect(setData).toHaveBeenCalledTimes(6);
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ({})),
      resultError: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      showToast: jest.fn()
    })
    InitItem.mockReturnValue(
      ()=>({description: "description", fieldname: "fieldname"})
    )
    settingActions(it_store, setData).setViewActions(
      { action: "editMenuField" }, 
      { id: null, description: "description", fieldname: "fieldname",
        fieldtype: 38, fieldtype_name: "string", menu_id: 1, orderby: 0 })
    expect(setData).toHaveBeenCalledTimes(8);

    //deleteItemRow
    settingActions(it_store, setData).setViewActions({ action: "deleteItemRow", table: "table" }, { id: 1 })
    expect(setData).toHaveBeenCalledTimes(11);
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ({ error: {} })),
      resultError: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      showToast: jest.fn()
    })
    settingActions(it_store, setData).setViewActions({ action: "deleteItemRow", table: "table" }, { id: 1 })
    expect(setData).toHaveBeenCalledTimes(14);

    settingActions(it_store, setData).setViewActions({ action: "missing" })
  })
  
  it('editItem', () => {
    const setData = jest.fn()
    let it_store = update(store, {setting: {$merge:{
      current: {
        form: {},
        fieldvalue: {}
      },
      audit: "all"
    }}})
    settingActions(it_store, setData).editItem({ name: "fieldname", value: "value" })
    expect(setData).toHaveBeenCalledTimes(1)

    it_store = update(it_store, {setting: {$merge:{
      audit: "readonly"
    }}})
    settingActions(it_store, setData).editItem({ name: "fieldvalue_value", value: "value" })
    expect(setData).toHaveBeenCalledTimes(2)

    it_store = update(it_store, {setting: {$merge:{
      type: "program"
    }}})
    settingActions(it_store, setData).editItem({ name: "fieldname", value: "value" })
    expect(setData).toHaveBeenCalledTimes(3)

    it_store = update(store, {setting: {$merge:{
      view: {},
      current: {
        form: {
          id: null,
          fromdate: "2022-02-12",
          todate: "",
          empnumber: "",
          logstate: "update",
          nervatype: ""
        }
      }
    }}})
    settingActions(it_store, setData).editItem({ name: "log_search" })
    expect(setData).toHaveBeenCalledTimes(4)
  })

  it('setPassword', () => {
    const setData = jest.fn()
    settingActions(store, setData).setPassword("admin")
    expect(setData).toHaveBeenCalledTimes(2)
    const it_store = update(store, {edit: {$merge:{
      current: {
        type: "employee"
      },
      dataset: {
        employee: [
          { username: "admin" }
        ]
      }
    }}})
    settingActions(it_store, setData).setPassword()
    expect(setData).toHaveBeenCalledTimes(4)
  })

  it('settingSave', () => {
    const setData = jest.fn((key, data, callback)=>{
      if(callback){callback()}
    })
    let it_store = update(store, {
      setting: {$merge:{
        current: {
          form: {
            id: null,
            fieldname: "value",
            transfilter: 1,
            translink: null
          }
        },
        dataset: {
          setting_view: [{ id: 1 }]
        }
      }},
      login: {$merge: {
        data: {
          groups: [
            { id: 1, groupvalue: "all" }
          ]
        }
      }}
    })
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path) => {
        if(path === "/view"){
          return { 
            setting_view: [{ id: 1 }]
          }
        }
        return [1]
      }),
      resultError: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      showToast: jest.fn()
    })
    settingActions(it_store, setData).settingSave()
    expect(setData).toHaveBeenCalledTimes(1)

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(() => ({ error: {} })),
      resultError: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      showToast: jest.fn()
    })
    settingActions(it_store, setData).settingSave()
    expect(setData).toHaveBeenCalledTimes(2)

    it_store = update(store, {
      setting: {$merge:{
        type: "password",
        current: {
          form: {
            username: ""
          }
        }
      }},
    })
    settingActions(it_store, setData).settingSave()
    expect(setData).toHaveBeenCalledTimes(3)
  })

  it('settingBack', () => {
    const setData = jest.fn((key, data, callback)=>{
      if(callback){callback()}
    })
    let it_store = update(store, {
      setting: {$merge:{
        current: {
          form: {
            id: null,
            fieldname: "value",
            transfilter: 1,
            translink: null
          }
        },
        dataset: {
          setting_view: [{ id: 1 }]
        }
      }},
      login: {$merge: {
        data: {
          groups: [
            { id: 1, groupvalue: "all" }
          ]
        }
      }}
    })
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async (path) => {
        if(path === "/view"){
          return { 
            setting_view: [{ id: 1 }]
          }
        }
        return [1]
      }),
      resultError: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      showToast: jest.fn()
    })
    settingActions(it_store, setData).settingBack("setting")
    expect(setData).toHaveBeenCalledTimes(0)
    settingActions(it_store, setData).settingBack()
    expect(setData).toHaveBeenCalledTimes(0)

    it_store = update(store, {
      setting: {$merge:{
        type: "password"
      }}
    })
    settingActions(it_store, setData).settingBack()
    expect(setData).toHaveBeenCalledTimes(1)
  })

})