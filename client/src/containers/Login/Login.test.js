import { render, fireEvent, queryByAttribute } from '@testing-library/react';
import update from 'immutability-helper';

import { AppProvider } from 'containers/App/context'
import { store as app_store  } from 'config/app'

import Login from './index';

import { appActions, getSql } from 'containers/App/actions'
jest.mock("containers/App/actions");

const getById = queryByAttribute.bind(null, 'id');

describe('<Login />', () => {

  beforeEach(() => {
    getSql.mockReturnValue({
      sql: "",
      prmCount: 1
    })
  });
  
  afterEach(() => {
    jest.clearAllMocks();
  });

  it('renders without crashing', () => {
    appActions.mockReturnValue({
      getText: jest.fn(),
    })
    let store_data = update(app_store, {
      current: {$merge: {
        theme: "light"
      }}
    })
    const setData = jest.fn()

    const { container, rerender } = render(
      <AppProvider value={{ data: store_data, setData: setData }}>
        <Login id="test_login" />
      </AppProvider>)
    expect(getById(container, 'test_login')).toBeDefined();

    const username = getById(container, 'username')
    fireEvent.change(username, {target: {value: "username"}})
    expect(setData).toHaveBeenCalledTimes(1);

    const sb_lang = getById(container, 'lang')
    fireEvent.change(sb_lang, {target: {value: "jp"}})
    expect(setData).toHaveBeenCalledTimes(2);

    const cmd_theme = getById(container, 'theme')
    fireEvent.click(cmd_theme)
    expect(setData).toHaveBeenCalledTimes(3);
    
    rerender(
      <AppProvider value={{ data: store_data, setData: setData }}>
        <Login id="test_login" />
      </AppProvider>)
  });

  it('onLogin error', () => {
    appActions.mockReturnValue({
      getText: jest.fn(),
      requestData: jest.fn(async () => ({ error: {} })),
      resultError: jest.fn(),
    })
    let store_data = update(app_store, {
      current: {$merge: {
        theme: "dark"
      }},
      login: {$merge: {
        username: "admin",
        database: "demo"
      }}
    })
    const setData = jest.fn()
    const { container } = render(
      <AppProvider value={{ data: store_data, setData: setData }}>
        <Login id="test_login" />
      </AppProvider>)
    
    const cmd_theme = getById(container, 'theme')
    fireEvent.click(cmd_theme)
    expect(setData).toHaveBeenCalledTimes(1);

    const cmd_login = getById(container, 'login')
    fireEvent.click(cmd_login)
    expect(setData).toHaveBeenCalledTimes(1);
  });

  it('onLogin engine_error', () => {
    appActions.mockReturnValue({
      getText: jest.fn(),
      requestData: jest.fn(async () => ({ token: "token", engine: "engine_error" })),
      resultError: jest.fn(),
    })
    let store_data = update(app_store, {
      login: {$merge: {
        username: "admin",
        database: "demo"
      }}
    })
    const setData = jest.fn()
    const { container } = render(
      <AppProvider value={{ data: store_data, setData: setData }}>
        <Login id="test_login" />
      </AppProvider>)
    
    const cmd_login = getById(container, 'login')
    fireEvent.click(cmd_login)
    expect(setData).toHaveBeenCalledTimes(0);
  });

  it('onLogin version_error', () => {
    appActions.mockReturnValue({
      getText: jest.fn(),
      requestData: jest.fn(async () => ({ token: "token", engine: "sqlite", version: "version_error" })),
      resultError: jest.fn(),
    })
    let store_data = update(app_store, {
      login: {$merge: {
        username: "admin",
        database: "demo"
      }}
    })
    const setData = jest.fn()
    const { container } = render(
      <AppProvider value={{ data: store_data, setData: setData }}>
        <Login id="test_login" />
      </AppProvider>)
    
    const cmd_login = getById(container, 'login')
    fireEvent.click(cmd_login)
    expect(setData).toHaveBeenCalledTimes(0);
  });

  it('onLogin loginData error 1', () => {
    appActions.mockReturnValue({
      getText: jest.fn(),
      requestData: jest.fn(async (path, options) => {
        if(String(path).endsWith("/auth/login")){
          return {
            token: "token", engine: "sqlite", version: "dev"
          }
        }
        return { error: {} }
      }),
      resultError: jest.fn(),
    })
    let store_data = update(app_store, {
      login: {$merge: {
        username: "admin",
        database: "demo"
      }}
    })
    const setData = jest.fn()
    const { container } = render(
      <AppProvider value={{ data: store_data, setData: setData }}>
        <Login id="test_login" />
      </AppProvider>)
    
    const cmd_login = getById(container, 'login')
    fireEvent.click(cmd_login)
    expect(setData).toHaveBeenCalledTimes(0);
  });

  it('onLogin loginData error 2', () => {
    appActions.mockReturnValue({
      getText: jest.fn(),
      requestData: jest.fn(async (path, options) => {
        if(String(path).endsWith("/auth/login")){
          return {
            token: "token", engine: "sqlite", version: "dev"
          }
        }
        if(options.data[0].key === "employee"){
          return {
            employee: [{ usergroup: 0 }], menuCmds: [], menuFields: [], userlogin: [], groups: []
          }
        }
        return { error: {} }
      }),
      resultError: jest.fn(),
    })
    let store_data = update(app_store, {
      login: {$merge: {
        username: "admin",
        database: "demo"
      }}
    })
    const setData = jest.fn()
    const { container } = render(
      <AppProvider value={{ data: store_data, setData: setData }}>
        <Login id="test_login" />
      </AppProvider>)
    
    const cmd_login = getById(container, 'login')
    fireEvent.click(cmd_login)
    expect(setData).toHaveBeenCalledTimes(0);
  });

  it('onLogin userLog error', () => {
    appActions.mockReturnValue({
      getText: jest.fn(),
      requestData: jest.fn(async (path, options) => {
        if(String(path).endsWith("/auth/login")){
          return {
            token: "token", engine: "sqlite", version: "dev"
          }
        }
        if(options.data[0].key === "employee"){
          return {
            employee: [{ usergroup: 0 }], menuCmds: [], menuFields: [], 
            userlogin: [{ value: "true" }], 
            groups: [{ id: 1, groupname: "transfilter", groupvalue: "all" }]
          }
        }
        if(options.data[0].key === "audit"){
          return {
            audit: [], transfilter: []
          }
        }
        return { error: {} }
      }),
      resultError: jest.fn(),
    })
    let store_data = update(app_store, {
      login: {$merge: {
        username: "admin",
        database: "demo"
      }}
    })
    const setData = jest.fn()
    const { container } = render(
      <AppProvider value={{ data: store_data, setData: setData }}>
        <Login id="test_login" />
      </AppProvider>)
    
    const cmd_login = getById(container, 'login')
    fireEvent.click(cmd_login)
    expect(setData).toHaveBeenCalledTimes(0);
  });

  it('onLogin success', () => {
    appActions.mockReturnValue({
      getText: jest.fn(),
      requestData: jest.fn(async (path, options) => {
        if(String(path).endsWith("/auth/login")){
          return {
            token: "token", engine: "sqlite", version: "dev"
          }
        }
        if(options.data[0].key === "employee"){
          return {
            employee: [{ id: 1, usergroup: 0 }], menuCmds: [], menuFields: [], 
            userlogin: [{ value: "false" }], 
            groups: [{ id: 1, groupname: "transfilter", groupvalue: "all" }]
          }
        }
        if(options.data[0].key === "audit"){
          return {
            audit: [
              { nervatypeName: "trans", subtypeName: "invoice", inputfilterName: "update", supervisor: 0 },
              { nervatypeName: "trans", subtypeName: "worksheet", inputfilterName: "disabled", supervisor: 0 },
              { nervatypeName: "tool", subtypeName: null, inputfilterName: "disabled", supervisor: 0 },
            ], 
            transfilter: [{ id: 1, transfilterName: "update" }]
          }
        }
        return {}
      }),
      resultError: jest.fn(),
      loadBookmark: jest.fn(({user_id, callback})=>{ 
        if(callback){callback()} 
      }),
    })
    let store_data = update(app_store, {
      login: {$merge: {
        username: "admin",
        database: "demo"
      }}
    })
    const setData = jest.fn()
    const { container } = render(
      <AppProvider value={{ data: store_data, setData: setData }}>
        <Login id="test_login" />
      </AppProvider>)
    
    const cmd_login = getById(container, 'login')
    fireEvent.click(cmd_login)
    expect(setData).toHaveBeenCalledTimes(0);
  });

  it('onLogin success and log', () => {
    appActions.mockReturnValue({
      getText: jest.fn(),
      requestData: jest.fn(async (path, options) => {
        if(String(path).endsWith("/auth/login")){
          return {
            token: "token", engine: "sqlite", version: "dev"
          }
        }
        if(options.data[0].key === "employee"){
          return {
            employee: [{ id: 1, usergroup: 0 }], menuCmds: [], menuFields: [], 
            userlogin: [{ value: "true" }], 
            groups: [{ id: 1, groupname: "transfilter", groupvalue: "all" }]
          }
        }
        if(options.data[0].key === "audit"){
          return {
            audit: [
              { nervatypeName: "trans", subtypeName: "invoice", inputfilterName: "update", supervisor: 0 },
              { nervatypeName: "trans", subtypeName: "worksheet", inputfilterName: "disabled", supervisor: 0 },
              { nervatypeName: "tool", subtypeName: null, inputfilterName: "disabled", supervisor: 0 },
            ], 
            transfilter: [{ id: 1, transfilterName: "update" }]
          }
        }
        return {}
      }),
      resultError: jest.fn(),
      loadBookmark: jest.fn(({user_id, callback})=>{ 
        if(callback){callback()} 
      }),
    })
    let store_data = update(app_store, {
      login: {$merge: {
        username: "admin",
        database: "demo"
      }}
    })
    const setData = jest.fn()
    const { container } = render(
      <AppProvider value={{ data: store_data, setData: setData }}>
        <Login id="test_login" />
      </AppProvider>)
    
    const cmd_login = getById(container, 'login')
    fireEvent.click(cmd_login)
    expect(setData).toHaveBeenCalledTimes(0)
  });

});