import sinon from 'sinon'
import { expect } from '@open-wc/testing';

import { AppController } from './AppController.js'
import { store as storeConfig } from '../config/app.js'
import { APP_MODULE, SIDE_EVENT, MODAL_EVENT } from '../config/enums.js'

const host = { 
  addController: ()=>{},
  inputBox: (prm)=>(prm),
  state: {
    ...storeConfig,
  },
  setData: sinon.spy(),
}

describe('AppController', () => {
  beforeEach(() => {
    Object.defineProperty(window, 'localStorage', { 
      value: {
        getItem: sinon.spy(),
        setItem: sinon.spy()
      } 
    })
  });

  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('hostConnected', () => {
    const app = new AppController(host)
    app._loadConfig = sinon.spy()
    app.hostConnected()
    expect(app.store).to.exist
  })

  it('_loadConfig', async () => {
    const app = new AppController(host)
    app.store = { data: {...storeConfig}, setData: sinon.spy() }
    app.request = sinon.spy(async () => ({ token: "token", report: { id: 1, code: "code", data: { report_name: "value" } } }))
    await app._loadConfig({ pathname: "/", search: "?session=I23sFMTc2NzA1NDY2Nw&code=ntr_employee_en" })
    sinon.assert.callCount(app.store.setData, 1);

    await app._loadConfig({ pathname: "/", hash: "#code=123&baba=haho&semmi=" })
    sinon.assert.callCount(app.store.setData, 1);


    await app._loadConfig({ pathname: "/", search: "abc" })
    sinon.assert.callCount(app.store.setData, 1);

    await app._loadConfig({ pathname: "abc/abc" })
    sinon.assert.callCount(app.store.setData, 1);

    app.request = sinon.spy(async () => { throw new Error("error"); })
    await app._loadConfig({ pathname: "abc/abc" })
    sinon.assert.callCount(app.store.setData, 1);

  })

  it('currentModule 1', async () => {
    const appHost = {...host}
    const app = new AppController(appHost)
    appHost.app = app
    app.store = { data: {...storeConfig}, setData: sinon.spy() }
    await app.currentModule({
      data: { module: APP_MODULE.TEMPLATE }, 
      content: { fkey: "setTemplate", args: [ { type: 'template', report: { id: 1, code: "code", data: { report_name: "value" } } } ] }
    })
    sinon.assert.callCount(app.store.setData, 3);
    await app.currentModule({
      data: { module: APP_MODULE.TEMPLATE }
    })
    sinon.assert.callCount(app.store.setData, 4);

  })

  it('getSetting', () => {
    const app = new AppController({...host})
    app.store = { data: {...storeConfig}, setData: sinon.spy() }

    Object.defineProperty(window, 'localStorage', { 
      value: {
        getItem: (key)=>(key==="loc") ? "de" : undefined,
      } 
    })
    let result = app.getSetting("loc")
    expect(result).to.equal("de");

    result = app.getSetting("missing")
    expect(result).to.equal("");

    result = app.getSetting("toastTimeout")
    expect(result).to.equal(4);

    result = app.getSetting("ui")
    expect(typeof(result)).to.equal("object");

    Object.defineProperty(window, 'localStorage', { 
      value: {
        getItem: (key)=>(key==="toastTimeout") ? 5 : undefined,
      } 
    })
    result = app.getSetting("ui")
    expect(typeof(result)).to.equal("object");

  })

  it('ms', () => {
    const storeData = {
      ...storeConfig,
      current: {
        ...storeConfig.current,
        lang: "de"
      },
      session: {
        ...storeConfig.session,
        locales: {
          ...storeConfig.session.locales,
          de: {
            key: "de"
          }
        }
      }
    }
    const app = new AppController(host)
    app.store = { data: {...storeData}, setData: sinon.spy() }
    let result = app.msg("", { id: "key" })
    expect(result).to.equal("de");

    result = app.msg("", { id: "login_username" })
    expect(result).to.equal("Username");

  })

  it('requestData', async () => {
    let storeData = {
      ...storeConfig,
      [APP_MODULE.LOGIN]: {
        ...storeConfig[APP_MODULE.LOGIN],
        data: {
          token: "token"
        }
      },
      session: {
        ...storeConfig.session,
        serverURL: "SERVER"
      }
    }
    const app = new AppController(host)
    app.store = { data: {...storeData}, setData: sinon.spy() }
    app.request = sinon.spy(async () => ({ hello: "world" }))
    let options = {
      data: {
        value: "value"
      },
      query: { 
        id: 1 
      }
    }
    let result = await app.requestData("/test", options, false)
    expect(result.hello).to.equal("world");

    storeData = {
      ...storeConfig,
      [APP_MODULE.LOGIN]: {
        ...storeConfig[APP_MODULE.LOGIN],
      },
      session: {
        ...storeConfig.session,
      }
    }
    app.store = { data: {...storeData}, setData: sinon.spy() }
    app.request = sinon.spy(async () => ({ code: 401, message: "error" }))
    app.signOut = sinon.spy()
    options = {
      token: "token",
      headers: {}
    }
    await app.requestData("/test", options, true)
    sinon.assert.callCount(app.signOut, 1);

    app.request = sinon.spy(async () => { throw new Error("error"); })
    options = {}
    result = await app.requestData("/test", options, false)
    expect(result.error.message).to.equal("error");

    result = await app.requestData("/test", options, true)
    expect(result.error.message).to.equal("error");

  })

  it('resultError', () => {
    const app = new AppController(host)
    app.store = { data: {...storeConfig}, setData: sinon.spy() }
    app.showToast = sinon.spy()
    app.resultError({})
    sinon.assert.callCount(app.store.setData, 0);
    sinon.assert.callCount(app.showToast, 1);

    app.resultError({ error: {} })
    sinon.assert.callCount(app.showToast, 2);
    sinon.assert.callCount(app.store.setData, 1);

    app.resultError({ error: { message: "error" } })
    sinon.assert.callCount(app.showToast, 3);
    sinon.assert.callCount(app.store.setData, 2);

  })

  it('showHelp', () => {
    const app = new AppController(host)
    app.store = { data: {...storeConfig}, setData: sinon.spy() }
    app.showHelp("help")

  })

  it('showToast', () => {
    let storeData = {
      ...storeConfig,
      current: {
        ...storeConfig.current,
        toast: {
          show: sinon.spy()
        }
      }
    }
    const app = new AppController(host)
    app.store = { data: {...storeData}, setData: sinon.spy() }
    app.getSetting = sinon.spy(()=>(5))
    app.showToast("type", "value", 10)
    sinon.assert.callCount(app.store.data.current.toast.show, 1);

    storeData = {
      ...storeConfig,
      current: {
        ...storeConfig.current,
      }
    }
    app.store = { data: {...storeData}, setData: sinon.spy() }
    app.showToast("type", "value")
  })

  it('signOut', () => {
    const storeData = {
      ...storeConfig,
    }
    const app = new AppController(host)
    app.store = { data: {...storeData}, setData: sinon.spy() }
    app.signOut()
    sinon.assert.callCount(app.store.setData, 1);
  
    /*
    storeData = {
      ...storeConfig,
      [APP_MODULE.LOGIN]: {
        ...storeConfig[APP_MODULE.LOGIN],
        callback: "/"
      }
    }
    app.store = { data: {...storeData}, setData: sinon.spy() }
    app.signOut()
    sinon.assert.callCount(app.store.setData, 0);
    */
  })

})