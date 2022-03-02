import { queryByAttribute } from '@testing-library/react'
import ReactDOM from 'react-dom';
import update from 'immutability-helper';

import { store as app_store  } from 'config/app'
import { getSql, appActions, request } from 'containers/App/actions'
import { searchActions } from './actions'

jest.mock("containers/App/actions");
const getById = queryByAttribute.bind(null, 'id');

const store = update(app_store, {$merge: {
  login: {
    data: {
      employee: {
        id: 1, empnumber: 'admin', username: 'admin', usergroup: 114,
        usergroupName: 'admin'
      },
      menuCmds: [
        { address: null, description: 'Server function example', funcname: 'nextNumber',
          icon: null, id: 1, menukey: 'nextNumber',method: 130, methodName: 'post', modul: null },
        { address: 'https://www.google.com', description: 'Internet URL example', funcname: 'search',
          icon: null, id: 2, menukey: 'google', method: 129, methodName: 'get', modul: null },
        { address: '', description: 'Server URL example', funcname: 'search',
          icon: null, id: 3, menukey: 'server', method: 129, methodName: 'get', modul: null },
        { address: "external", description: 'Server function example', funcname: 'nextNumber',
          icon: null, id: 4, menukey: 'nextNumber',method: 130, methodName: 'post', modul: null },
      ],
      menuFields: [
        { description: 'Code (e.g. custnumber)', fieldname: 'numberkey', fieldtype: 38,
          fieldtypeName: 'string', id: 1, menu_id: 1, orderby: 0 },
        { description: 'Stepping', fieldname: 'step', fieldtype: 33, fieldtypeName: 'bool',
          id: 2, menu_id: 1, orderby: 1 },
        { description: 'google search', fieldname: 'q', fieldtype: 38, fieldtypeName: 'string',
          id: 3, menu_id: 2, orderby: 0 },
        { description: 'Number', fieldname: 'number', fieldtype: 66, fieldtypeName: 'float',
          id: 4, menu_id: 1, orderby: 2 },
      ]
    }
  }
}})

describe('searchActions', () => {
  beforeEach(() => {
    request.mockReturnValue({})
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ({})),
      resultError: jest.fn(),
      getAuditFilter: jest.fn(() => (["all",1])),
      showToast: jest.fn()
    })
    getSql.mockReturnValue({
      sql: "",
      prmCount: 1
    })
    Object.defineProperty(global.window, 'open', { value: jest.fn() });
  });
  
  afterEach(() => {
    jest.clearAllMocks();
  });

  it('showBrowser', () => {
    const setData = jest.fn((key, data, callback)=>{
      if(callback){callback()}
    })
    let it_store = update(store, {})
    searchActions(it_store, setData).showBrowser("customer", "CustomerView", it_store.search)
    expect(setData).toHaveBeenCalledTimes(1);
    
    it_store = update(it_store, {search: {$merge: {
      show_header: true,
      show_columns: true,
      page: 1,
      columns: {
        CustomerView: {
          custnumber:true, custname:true
        }
      }
    }}})
    searchActions(it_store, setData).showBrowser("customer", undefined, it_store.search)
    expect(setData).toHaveBeenCalledTimes(2);

    it_store = update(it_store, {search: {$merge: {
      show_header: true,
      show_columns: true,
      page: 1,
      columns: {
        CustomerView: {}
      },
      filters: {
        CustomerView: [{}]
      }
    }}})
    searchActions(it_store, setData).showBrowser("customer", undefined, it_store.search)
    expect(setData).toHaveBeenCalledTimes(3);
  })

  it('showBrowser error', () => {
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ({ error: {} })),
      resultError: jest.fn(),
    })
    const setData = jest.fn((key, data, callback)=>{
      if(callback){callback()}
    })
    const it_store = update(store, {})
    searchActions(it_store, setData).showBrowser("customer", "CustomerView")
    expect(setData).toHaveBeenCalledTimes(1);
  })

  it('defaultFilterValue', () => {
    const it_store = update(store, {})
    let result = searchActions(it_store, jest.fn()).defaultFilterValue("date")
    expect(result).toBeDefined()
    result = searchActions(it_store, jest.fn()).defaultFilterValue("integer")
    expect(result).toBe(0)
    result = searchActions(it_store, jest.fn()).defaultFilterValue("string")
    expect(result).toBe("")
  })

  it('getFilterWhere', () => {
    const it_store = update(store, {})
    let result = searchActions(it_store, jest.fn()).getFilterWhere({ 
      filtertype: "==N", fieldtype: "string", sqlstr: "" })
    expect(result[1][1][2]).toBe("is null")
    result = searchActions(it_store, jest.fn()).getFilterWhere({ 
      filtertype: "==N", fieldtype: "integer", sqlstr: "" })
    expect(result[2]).toBe("is null")

    result = searchActions(it_store, jest.fn()).getFilterWhere({ 
      filtertype: "!==", fieldtype: "string", sqlstr: "" })
    expect(result[1][2]).toBe("{CCS}{JOKER}{SEP}lower(?){SEP}{JOKER}{CCE}")
    result = searchActions(it_store, jest.fn()).getFilterWhere({ 
      filtertype: "!==", fieldtype: "integer", sqlstr: "" })
    expect(result[1][2]).toBe("?")

    result = searchActions(it_store, jest.fn()).getFilterWhere({ 
      filtertype: ">==", fieldtype: "string", sqlstr: "" })
    expect(result[1][2]).toBe("?")

    result = searchActions(it_store, jest.fn()).getFilterWhere({ 
      filtertype: "<==", fieldtype: "string", sqlstr: "" })
    expect(result[1][2]).toBe("?")

    result = searchActions(it_store, jest.fn()).getFilterWhere({ 
      filtertype: "===", fieldtype: "string", sqlstr: "" })
    expect(result[1][2]).toBe("{CCS}{JOKER}{SEP}lower(?){SEP}{JOKER}{CCE}")
    result = searchActions(it_store, jest.fn()).getFilterWhere({ 
      filtertype: "===", fieldtype: "integer", sqlstr: "" })
    expect(result[1][2]).toBe("?")
    
  })

  it('showServerCmd', () => {
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
    let it_store = update(store, {})
    searchActions(it_store, setData).showServerCmd(1)
    expect(setData).toHaveBeenCalledTimes(3);

    searchActions(it_store, setData).showServerCmd(2)
    expect(setData).toHaveBeenCalledTimes(6);

    searchActions(it_store, setData).showServerCmd(3)
    expect(setData).toHaveBeenCalledTimes(9);

    searchActions(it_store, setData).showServerCmd(4)
    expect(setData).toHaveBeenCalledTimes(12);

    it_store = update(store, {session: {$merge: {
      configServer: true
    }}})
    searchActions(it_store, setData).showServerCmd(3)
    expect(setData).toHaveBeenCalledTimes(15);

    it_store = update(store, {login: {data: {$merge: {
      menuCmds: [
        { address: null, description: 'Server URL example', funcname: '',
          icon: null, id: 3, menukey: 'server', method: 129, methodName: 'get', modul: null }
      ]
    }}}})
    searchActions(it_store, setData).showServerCmd(3)
    expect(setData).toHaveBeenCalledTimes(18);

    searchActions(it_store, setData).showServerCmd(99)
    expect(setData).toHaveBeenCalledTimes(18);
    
  })

  it('showServerCmd error', () => {
    request.mockImplementation(() => {
      throw new Error();
    })
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ({ error: {} })),
      resultError: jest.fn(),
      showToast: jest.fn()
    })

    const setData = jest.fn((key, data, callback)=>{
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(data.modalForm, container);

        const field_step = getById(container, 'field_step')
        if(field_step){
          field_step.dispatchEvent(new MouseEvent('click', {bubbles: true}))
        }

        // onOK
        const btn_ok = getById(container, 'btn_ok')
        btn_ok.dispatchEvent(new MouseEvent('click', {bubbles: true}));
      }
      if(callback){callback()}
    })
    let it_store = update(store, {})
    searchActions(it_store, setData).showServerCmd(1)
    expect(setData).toHaveBeenCalledTimes(2);

    searchActions(it_store, setData).showServerCmd(4)
    expect(setData).toHaveBeenCalledTimes(4);
  })

  it('showServerCmd string result', () => {
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ("")),
      resultError: jest.fn(),
      showToast: jest.fn()
    })

    const setData = jest.fn((key, data, callback)=>{
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(data.modalForm, container);

        // onOK
        const btn_ok = getById(container, 'btn_ok')
        btn_ok.dispatchEvent(new MouseEvent('click', {bubbles: true}));
      }
      if(callback){callback()}
    })
    let it_store = update(store, {login: {data: {$merge: {
      menuCmds: [
        { address: null, description: 'Server function example', funcname: null,
          icon: null, id: 1, menukey: 'nextNumber',method: 130, methodName: 'post', modul: null }
      ]
    }}}})
    searchActions(it_store, setData).showServerCmd(1)
    expect(setData).toHaveBeenCalledTimes(2);
  })

})