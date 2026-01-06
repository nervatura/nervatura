import sinon from 'sinon'

import { LoginController } from './LoginController.js'
import { store as storeConfig } from '../config/app.js'
import { LOGIN_PAGE_EVENT, APP_MODULE } from '../config/enums.js'

const host = { 
  addController: ()=>{}, 
}
const store = {
  data: {
    ...storeConfig
  },
  setData: ()=>{},
}
const app = {
  store,
  showToast: sinon.spy(),
  requestData: () => ({ value: "OK" }),
  getSql: () =>({
    sql: "",
    prmCount: 1
  }),
  currentModule: sinon.spy(),
  msg: (value)=>value,
}

describe('LoginController', () => {

  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('onLogin', async () => {
    // error
    let testApp = {
      ...app,
      requestData: () => ({ error: {} }),
      resultError: sinon.spy()
    }
    // debugger;
    let login = new LoginController({...host, app: testApp})
    await login.onLogin()
    sinon.assert.callCount(testApp.resultError, 1);

    testApp = {
      ...testApp,
      requestData: () => ({ token: "token", report: { id: 1, code: "code", data: { report_name: "value" } } }),
      store: {
        ...testApp.store,
        setData: sinon.spy(),
        data: {
          ...testApp.store.data,
          [APP_MODULE.LOGIN]: {
            ...testApp.store.data[APP_MODULE.LOGIN],
            username: "username",
            password: "password",
            database: "database",
            code: "code",
          }
        }
      }
    }
    login = new LoginController({...host, app: testApp})
    await login.onLogin()
    sinon.assert.callCount(testApp.store.setData, 1);

    testApp = {
      ...testApp,
      requestData: () => ({ error: {}, token: "token", report: { id: 1, code: "code", data: { report_name: "value" } } }),
      store: {
        ...testApp.store,
        setData: sinon.spy(),
        data: {
          ...testApp.store.data,
          [APP_MODULE.LOGIN]: {
            ...testApp.store.data[APP_MODULE.LOGIN],
            username: "username",
            password: "password",
            database: "database",
            code: "code",
          }
        }
      }
    }
    login = new LoginController({...host, app: testApp})
    await login.onLogin()
    sinon.assert.callCount(testApp.store.setData, 1);

  })

  it('onPageEvent', () => {
    const testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy()
      },
      requestData: () => ({ error: {} }),
      resultError: sinon.spy(),
    }
    const login = new LoginController({...host, app: testApp})
    login.onPageEvent({ key: LOGIN_PAGE_EVENT.CHANGE, data: { fieldname: "fieldname", value: "value" } })
    sinon.assert.callCount(testApp.store.setData, 1);

    login.onPageEvent({ key: LOGIN_PAGE_EVENT.THEME, data: "theme" })
    sinon.assert.callCount(testApp.store.setData, 2);

    login.onPageEvent({ key: LOGIN_PAGE_EVENT.LANG, data: "lang" })
    sinon.assert.callCount(testApp.store.setData, 3);

    login.onPageEvent({ key: LOGIN_PAGE_EVENT.LOGIN, data: {} })
    login.onPageEvent({ key: "missing", data: {} })

  })

})