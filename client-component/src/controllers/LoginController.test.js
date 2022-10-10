// import { expect } from '@open-wc/testing';
import sinon from 'sinon'

import { LoginController } from './LoginController.js'
import { store as storeConfig } from '../config/app.js'

const host = { addController: ()=>{} }
const store = {
  data: {
    ...storeConfig
  },
  setData: ()=>{},
  showToast: ()=>{},
  msg: (value)=>value
}
const app = {
  requestData: () => ({ value: "OK" }),
  getSql: () =>({
    sql: "",
    prmCount: 1
  })
}

describe('LoginPage', () => {
  /*
  it('userLog', async () => {
    const login = new LoginController(host, store, app)
    const result = await login.userLog({
      employee: {
        empnumber: ""
      }
    })
    // debugger;
    expect(result.value).to.equal("OK");
  })
  */

  it('onLogin', async () => {
    // error
    let testApp = {
      ...app,
      requestData: () => ({ error: {} }),
      resultError: sinon.spy()
    }
    let login = new LoginController(host, store, testApp)
    await login.onLogin()
    sinon.assert.callCount(testApp.resultError, 1);

    // engine_error
    testApp = {
      ...testApp,
      requestData: () => ({ token: "token", engine: "engine_error" }),
    }
    login = new LoginController(host, store, testApp)
    await login.onLogin()
    sinon.assert.callCount(testApp.resultError, 2);

    // version_error
    testApp = {
      ...testApp,
      requestData: () => ({ token: "token", engine: "sqlite", version: "version_error" }),
    }
    login = new LoginController(host, store, testApp)
    await login.onLogin()
    sinon.assert.callCount(testApp.resultError, 3);

  })

  it('loginData', async () => {
    // error 1
    let testApp = {
      ...app,
      requestData: (path) => {
        if(String(path).endsWith("/auth/login")){
          return {
            token: "token", engine: "sqlite", version: "dev"
          }
        }
        return { error: {} }
      },
      resultError: sinon.spy()
    }
    let login = new LoginController(host, store, testApp)
    await login.onLogin()
    sinon.assert.callCount(testApp.resultError, 1);

    // error 2
    testApp = {
      ...testApp,
      requestData: (path, options) => {
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
      },
    }
    login = new LoginController(host, store, testApp)
    await login.onLogin()
    sinon.assert.callCount(testApp.resultError, 2);

    // userLog error
    testApp = {
      ...testApp,
      requestData: (path, options) => {
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
      },
    }
    login = new LoginController(host, store, testApp)
    await login.onLogin()
    sinon.assert.callCount(testApp.resultError, 3);

    // success
    testApp = {
      ...testApp,
      requestData: (path, options) => {
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
      },
      loadBookmark: sinon.spy()
    }
    login = new LoginController(host, store, testApp)
    await login.onLogin()
    sinon.assert.callCount(testApp.loadBookmark, 1);

    // success and log
    testApp = {
      ...testApp,
      requestData: (path, options) => {
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
      }
    }
    login = new LoginController(host, store, testApp)
    await login.onLogin()
    sinon.assert.callCount(testApp.loadBookmark, 2);
  })

  it('setCodeToken', async () => {
    let testApp = {
      ...app,
      requestData: (path, options) => {
        if(String(path).endsWith("/auth/validate")){
          return {
            username:"username", database:"database", engine:"engine", version:"version"
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
      },
      resultError: sinon.spy(),
      loadBookmark: sinon.spy()
    }
    let login = new LoginController(host, store, testApp)
    await login.tokenValidation({ code: "code", callback: "/callback" })
    sinon.assert.callCount(testApp.loadBookmark, 1);

    // tokenError error callback
    testApp = {
      ...testApp,
      requestData: () => ({ error: {} }),
    }
    login = new LoginController(host, store, testApp)
    await login.tokenValidation({ code: "code", callback: "/callback" })

    // tokenError error
    testApp = {
      ...testApp,
      requestData: () => ({ error: {} }),
    }
    login = new LoginController(host, store, testApp)
    await login.tokenValidation({ code: "code" })
    sinon.assert.callCount(testApp.resultError, 1);

    // resultData error
    testApp = {
      ...testApp,
      requestData: (path) => {
        if(String(path).endsWith("/auth/validate")){
          return {
            token: "token", engine: "sqlite", version: "dev"
          }
        }
        return { error: {} }
      }
    }
    login = new LoginController(host, store, testApp)
    await login.tokenValidation({ code: "code" })
    sinon.assert.callCount(testApp.resultError, 2);
    
  })

  it('setCodeToken', async () => {
    const testApp = {
      ...app,
      request: ()=>({access_token: "access_token"}),
      requestData: () => ({ error: {} }),
      resultError: sinon.spy(),
      loadBookmark: sinon.spy()
    }
    let login = new LoginController(host, store, testApp)
    await login.setCodeToken({ code: "code", callback: "/callback" })
    sinon.assert.callCount(testApp.resultError, 1);

    // callback error
    login = new LoginController(host, store, testApp)
    await login.setCodeToken({ code: "code" })
    sinon.assert.callCount(testApp.resultError, 2);

  })

  it('onPageEvent', () => {
    const testStore = {
      ...store,
      setData: sinon.spy()
    }
    const testApp = {
      ...app,
      requestData: () => ({ error: {} }),
      resultError: sinon.spy(),
    }
    const login = new LoginController(host, testStore, testApp)
    login.onPageEvent({ key: "change", data: { fieldname: "fieldname", value: "value" } })
    sinon.assert.callCount(testStore.setData, 1);

    login.onPageEvent({ key: "theme", data: "theme" })
    sinon.assert.callCount(testStore.setData, 2);

    login.onPageEvent({ key: "lang", data: "lang" })
    sinon.assert.callCount(testStore.setData, 3);

    login.onPageEvent({ key: "login", data: {} })
    login.onPageEvent({ key: "missing", data: {} })

  })

})