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
    app.request = sinon.spy(async () => ({ locales: { de: {} } }))
    app.tokenLogin = sinon.spy()
    await app._loadConfig({ pathname: "/", hash: "abc" })
    sinon.assert.callCount(app.store.setData, 1);

    Object.defineProperty(window, 'localStorage', { 
      value: {
        getItem: sinon.spy(()=>("de")),
        setItem: sinon.spy()
      } 
    })
    await app._loadConfig({ pathname: "/", hash: "#code=123&baba=haho&semmi=" })
    sinon.assert.callCount(app.store.setData, 3);

    await app._loadConfig({ pathname: "/", search: "abc" })
    sinon.assert.callCount(app.store.setData, 5);

    await app._loadConfig({ pathname: "abc/abc" })
    sinon.assert.callCount(app.store.setData, 7);

    app.request = sinon.spy(async () => { throw new Error("error"); })
    await app._loadConfig({ pathname: "abc/abc" })
    sinon.assert.callCount(app.store.setData, 7);

  })

  it('createHistory', async () => {
    const app = new AppController(host)
    let storeData = {
      ...storeConfig,
      [APP_MODULE.EDIT]: {
        ...storeConfig[APP_MODULE.EDIT],
        current: {
          type: "trans",
          transtype: "invoice",
          item: {
            id: 5,
            transnumber: "DMINV/00001"
          }
        },
        template: {
          options: {
            title: "INVOICE",
            title_field: "transnumber"
          }
        }
      },
      [APP_MODULE.LOGIN]: {
        ...storeConfig[APP_MODULE.LOGIN],
        data: {
          employee: {
            id: 1
          }
        }
      }
    }
    app.store = { data: storeData, setData: sinon.spy() }
    app.requestData = sinon.spy(async () => ({}))
    app.resultError = sinon.spy()
    app.msg = sinon.spy(()=>(""))
    app.getSetting = sinon.spy(()=>("5"))
    await app.createHistory("save")
    sinon.assert.callCount(app.store.setData, 1);

    storeData = {
      ...storeConfig,
      [APP_MODULE.EDIT]: {
        ...storeConfig[APP_MODULE.EDIT],
        current: {
          type: "customer",
          item: {
            id: 1,
            custnumber: "CUST/00001"
          }
        },
        template: {
          options: {
            title: "CUSTOMER",
            title_field: "custnumber"
          }
        }
      },
      [APP_MODULE.LOGIN]: {
        ...storeConfig[APP_MODULE.LOGIN],
        data: {
          employee: {
            id: 1
          }
        }
      },
      [APP_MODULE.BOOKMARK]: {
        ...storeConfig[APP_MODULE.BOOKMARK],
        history: {
          employee_id: 1,
          section: "history",
          cfgroup: "2022-01-03T22:27:00+02:00",
          cfname: 1,
          cfvalue: "[{\"datetime\":\"2022-01-03T22:21:32+02:00\",\"type\":\"save\",\"ntype\":\"trans\",\"transtype\":\"invoice\",\"id\":5,\"title\":\"INVOICE | DMINV/00001\"}]",
        }
      }
    }

    app.store = { data: storeData, setData: sinon.spy() }
    await app.createHistory("save")
    sinon.assert.callCount(app.store.setData, 1);

    storeData = {
      ...storeConfig,
      [APP_MODULE.EDIT]: {
        ...storeConfig[APP_MODULE.EDIT],
        current: {
          type: "customer",
          item: {
            id: 1,
          }
        },
        template: {
          options: {
            title: "CUSTOMER",
          }
        }
      },
      [APP_MODULE.LOGIN]: {
        ...storeConfig[APP_MODULE.LOGIN],
        data: {
          employee: {
            id: 1
          }
        }
      },
      [APP_MODULE.BOOKMARK]: {
        ...storeConfig[APP_MODULE.BOOKMARK],
        history: {
          employee_id: 1,
          section: "history",
          cfgroup: "2022-01-03T22:27:00+02:00",
          cfname: 1,
          cfvalue: "[{},{},{},{},{},{},{},{},{},{}]",
        }
      }
    }

    app.store = { data: storeData, setData: sinon.spy() }
    app.requestData = sinon.spy(async () => ({ error: {} }))
    await app.createHistory("save")
    sinon.assert.callCount(app.resultError, 1);
  })

  it('currentModule 1', async () => {
    const appHost = {...host}
    const app = new AppController(appHost)
    appHost.app = app
    app.store = { data: {...storeConfig}, setData: sinon.spy() }
    await app.currentModule({
      data: { module: APP_MODULE.SETTING }, 
      content: { fkey: "checkSetting", args: [ { username: "admin" }, SIDE_EVENT.PASSWORD_FORM ] }
    })
    sinon.assert.callCount(app.store.setData, 2);
    await app.currentModule({
      data: { module: APP_MODULE.SETTING }
    })
    sinon.assert.callCount(app.store.setData, 3);

    await app.currentModule({
      data: { module: APP_MODULE.EDIT }, 
    })
    sinon.assert.callCount(app.store.setData, 4);

  })

  it('currentModule 2', async () => {
    const appHost = {...host}
    const app = new AppController(appHost)
    appHost.app = app
    app.store = { data: {...storeConfig}, setData: sinon.spy() }
    await app.currentModule({
      data: { module: APP_MODULE.EDIT }, 
    })
    sinon.assert.callCount(app.store.setData, 1);

  })

  it('currentModule 3', async () => {
    const appHost = {...host}
    const app = new AppController(appHost)
    appHost.app = app
    app.store = { data: {...storeConfig}, setData: sinon.spy() }
    await app.currentModule({
      data: { module: APP_MODULE.TEMPLATE }, 
    })
    sinon.assert.callCount(app.store.setData, 1);

    await app.currentModule({
      data: { module: APP_MODULE.SEARCH }, 
    })
    sinon.assert.callCount(app.store.setData, 2);

    await app.currentModule({
      data: { module: APP_MODULE.SEARCH }, 
    })
    sinon.assert.callCount(app.store.setData, 3);

  })

  it('getAuditFilter', () => {
    const storeData = {
      ...storeConfig,
      [APP_MODULE.LOGIN]: {
        ...storeConfig[APP_MODULE.LOGIN],
        data: {
          audit: [
            {
              inputfilter: 108,
              inputfilterName: 'update',
              nervatype: 10,
              nervatypeName: 'customer',
              subtype: null,
              subtypeName: null,
              supervisor: 1
            },
            {
              inputfilter: 107,
              inputfilterName: 'readonly',
              nervatype: 31,
              nervatypeName: 'trans',
              subtype: 62,
              subtypeName: 'inventory',
              supervisor: 0
            },
            {
              inputfilter: 106,
              inputfilterName: 'disabled',
              nervatype: 28,
              nervatypeName: 'report',
              subtype: 6,
              subtypeName: null,
              supervisor: 0
            },
            {
              inputfilter: 106,
              inputfilterName: 'disabled',
              nervatype: 18,
              nervatypeName: 'menu',
              subtype: 1,
              subtypeName: 'nextNumber',
              supervisor: 0
            }
          ]
        }
      }
    }
    const app = new AppController({...host})
    app.store = { data: {...storeData}, setData: sinon.spy() }
    let audit = app.getAuditFilter("trans", "inventory")
    expect(audit[0]).to.equal("readonly");

    audit = app.getAuditFilter("menu", "nextNumber")
    expect(audit[0]).to.equal("disabled");

    audit = app.getAuditFilter("report", 6)
    expect(audit[0]).to.equal("disabled");

    audit = app.getAuditFilter("customer")
    expect(audit[0]).to.equal("update");

    audit = app.getAuditFilter("product")
    expect(audit[0]).to.equal("all");

  })

  it('getDataFilter', () => {
    let storeData = {
      ...storeConfig,
      [APP_MODULE.LOGIN]: {
        ...storeConfig[APP_MODULE.LOGIN],
        data: {
          audit: []
        }
      }
    }
    const app = new AppController({...host})
    app.store = { data: {...storeData}, setData: sinon.spy() }
    let result = app.getDataFilter("transitem", [])
    expect(result.length).to.equal(0);

    result = app.getDataFilter("transpayment", [])
    expect(result.length).to.equal(0);

    result = app.getDataFilter("transmovement", [], "")
    expect(result.length).to.equal(0);

    result = app.getDataFilter("transmovement", [], "InventoryView")
    expect(result.length).to.equal(0);

    storeData = {
      ...storeConfig,
      [APP_MODULE.LOGIN]: {
        ...storeConfig[APP_MODULE.LOGIN],
        data: {
          audit: [
            { inputfilterName: 'disabled', nervatypeName: 'trans', subtypeName: 'offer', supervisor: 0 },
            { inputfilterName: 'disabled', nervatypeName: 'trans', subtypeName: 'order', supervisor: 0 },
            { inputfilterName: 'disabled', nervatypeName: 'trans', subtypeName: 'worksheet', supervisor: 0 },
            { inputfilterName: 'disabled', nervatypeName: 'trans', subtypeName: 'rent', supervisor: 0 },
            { inputfilterName: 'disabled', nervatypeName: 'trans', subtypeName: 'invoice', supervisor: 0 },
            { inputfilterName: 'disabled', nervatypeName: 'trans', subtypeName: 'bank', supervisor: 0 },
            { inputfilterName: 'disabled', nervatypeName: 'trans', subtypeName: 'cash', supervisor: 0 },
            { inputfilterName: 'disabled', nervatypeName: 'trans', subtypeName: 'delivery', supervisor: 0 },
            { inputfilterName: 'disabled', nervatypeName: 'trans', subtypeName: 'inventory', supervisor: 0 },
          ]
        }
      }
    }
    app.store = { data: {...storeData}, setData: sinon.spy() }
    result = app.getDataFilter("transitem", [])
    expect(result.length).to.equal(10);

    result = app.getDataFilter("transpayment", [])
    expect(result.length).to.equal(4);

    result = app.getDataFilter("transmovement", [], "")
    expect(result.length).to.equal(4);

    result = app.getDataFilter("transmovement", [], "InventoryView")
    expect(result.length).to.equal(0);

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

  it('getUserFilter', () => {
    let storeData = {
      ...storeConfig,
      [APP_MODULE.LOGIN]: {
        ...storeConfig[APP_MODULE.LOGIN],
        data: {
          employee: {
            id: 1, empnumber: 'admin', username: 'admin', usergroup: 114,
            usergroupName: 'admin'
          },
          audit: [],
        }
      }
    }
    const app = new AppController({...host})
    app.store = { data: {...storeData}, setData: sinon.spy() }
    let filter = app.getUserFilter("customer")
    expect(filter.params.length).to.equal(0);

    storeData = {
      ...storeData,
      [APP_MODULE.LOGIN]: {
        ...storeData[APP_MODULE.LOGIN],
        data: {
          ...storeData[APP_MODULE.LOGIN].data,
          transfilterName: "usergroup"
        }
      }
    }
    app.store = { data: {...storeData}, setData: sinon.spy() }
    filter = app.getUserFilter("customer")
    expect(filter.params.length).to.equal(0);

    filter = app.getUserFilter("transitem")
    expect(filter.params.length).to.equal(1);

    storeData = {
      ...storeData,
      [APP_MODULE.LOGIN]: {
        ...storeData[APP_MODULE.LOGIN],
        data: {
          ...storeData[APP_MODULE.LOGIN].data,
          transfilterName: "own"
        }
      }
    }
    app.store = { data: {...storeData}, setData: sinon.spy() }
    filter = app.getUserFilter("customer")
    expect(filter.params.length).to.equal(0);

    filter = app.getUserFilter("transitem")
    expect(filter.params.length).to.equal(1);

  })

  it('loadBookmark', async () => {
    const app = new AppController(host)
    app.store = { data: {...storeConfig}, setData: sinon.spy() }
    app.requestData = sinon.spy(async () => ([{ section: "history" },{ section: "bookmark" }]))
    app.resultError = sinon.spy()
    const result = await app.loadBookmark({ token:"token", user_id: 1, callback: ()=>{} })
    expect(result.bookmark.length).to.equal(1);

  })

  it('loadBookmark', async () => {
    const app = new AppController(host)
    app.store = { data: {...storeConfig}, setData: sinon.spy() }
    app.resultError = sinon.spy()
    app.requestData = sinon.spy(async () => ([]))
    const result = await app.loadBookmark({ token:"token", user_id: 1 })
    expect(result.history).to.equal(null);

  })

  it('loadBookmark', async () => {
    const app = new AppController(host)
    app.store = { data: {...storeConfig}, setData: sinon.spy() }
    app.resultError = sinon.spy()
    app.requestData = sinon.spy(async () => ({ error: {} }))
    const result = await app.loadBookmark({ token:"token", user_id: 1 })
    expect(result).to.equal(null);

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

  it('quickSearch', async () => {
    const { Quick } = await import('./Quick.js');
    const storeData = {
      ...storeConfig,
      [APP_MODULE.LOGIN]: {
        ...storeConfig[APP_MODULE.LOGIN],
        data: {
          employee: {
            id: 1, empnumber: 'admin', username: 'admin', usergroup: 114,
            usergroupName: 'admin'
          },
          audit: [],
          transfilterName: "usergroup"
        }
      }
    }
    const app = new AppController(host)
    app.modules = {
      quick: {...Quick}
    }
    app.store = { data: {...storeData}, setData: sinon.spy() }
    app.requestData = sinon.spy(async () => ({}))
    let result = await app.quickSearch("customer", "")
    expect(result).to.is.exist;

    result = await app.quickSearch("transitem", "item")
    expect(result).to.is.exist;

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
    app.request = sinon.spy(async () => ({ code: 401 }))
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

  it('saveBookmark', async () => {
    const storeData = {
      ...storeConfig,
      [APP_MODULE.SEARCH]: {
        ...storeConfig[APP_MODULE.SEARCH],
        vkey: "customer",
        filters: {
          CustomerView: [
          ],
        },
        columns: {
          CustomerView: {
            custnumber: true,
            custname: true,
            address: true,
          },
        },
        view: "CustomerView",
      },
      [APP_MODULE.LOGIN]: {
        ...storeConfig[APP_MODULE.LOGIN],
        data: {
          employee: {
            id: 1
          }
        }
      },
    }
    const app = new AppController(host)
    app.store = { 
      data: {...storeData}, 
      setData: sinon.spy((key, data) => {
        if(data && data.modalForm){
          data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL })
          data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: { value: "value" } })
        }
      }) 
    }
    app.resultError = sinon.spy()
    app.requestData = sinon.spy(async () => ([]))
    app.loadBookmark = sinon.spy()
    await app.saveBookmark(['browser', 'Customer Data'])
    sinon.assert.callCount(app.loadBookmark, 1);

  })

  it('saveBookmark', async () => {
    let storeData = {
      ...storeConfig,
      [APP_MODULE.EDIT]: {
        ...storeConfig[APP_MODULE.EDIT],
        current: {
          type: "trans",
          transtype: "invoice",
          item: {
            id: 5,
            transnumber: "DMINV/00001",
            transdate: "2020-12-10",
          }
        },
        dataset: {
          trans: [
            {
              custname: "First Customer Co."
            }
          ]
        }
      },
      [APP_MODULE.LOGIN]: {
        ...storeConfig[APP_MODULE.LOGIN],
        data: {
          employee: {
            id: 1
          }
        }
      },
    }
    const app = new AppController(host)
    app.store = { 
      data: {...storeData}, 
      setData: sinon.spy((key, data) => {
        if(data && data.modalForm){
          data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: { value: "value" } })
        }
      }) 
    }
    app.resultError = sinon.spy()
    app.requestData = sinon.spy(async () => ({ error: {} }))
    app.loadBookmark = sinon.spy()
    await app.saveBookmark(['editor', 'trans', 'transnumber'])
    sinon.assert.callCount(app.resultError, 1);

    storeData = {
      ...storeData,
      [APP_MODULE.EDIT]: {
        ...storeData[APP_MODULE.EDIT],
        current: {
          type: "trans",
          transtype: "receipt",
          item: {
            id: 5,
            transnumber: "DMINV/00001",
            transdate: "2020-12-10",
          }
        },
        dataset: {
          trans: [
            {
              custname: null
            }
          ]
        }
      }
    }
    app.store = { 
      data: {...storeData}, 
      setData: sinon.spy((key, data) => {
        if(data && data.modalForm){
          data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: { value: "value" } })
        }
      }) 
    }
    app.resultError = sinon.spy()
    app.requestData = sinon.spy(async () => ({}))
    app.loadBookmark = sinon.spy()
    await app.saveBookmark(['editor', 'trans', 'transnumber'])
    sinon.assert.callCount(app.loadBookmark, 1);

    storeData = {
      ...storeData,
      [APP_MODULE.EDIT]: {
        ...storeData[APP_MODULE.EDIT],
        current: {
          type: "customer",
          transtype: "",
          item: {
            id: 2,
            custnumber: "DMCUST/00001",
            custname: "First Customer Co.",
          }
        },
        dataset: {
          trans: [
            {
              custname: null
            }
          ]
        }
      }
    }
    app.store = { 
      data: {...storeData}, 
      setData: sinon.spy((key, data) => {
        if(data && data.modalForm){
          data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: { value: "value" } })
        }
      }) 
    }
    app.resultError = sinon.spy()
    app.requestData = sinon.spy(async () => ({}))
    app.loadBookmark = sinon.spy()
    await app.saveBookmark(['editor', 'customer', 'custname', 'custnumber'])
    sinon.assert.callCount(app.loadBookmark, 1);

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

  it('tokenLogin', async () => {
    const appHost = {...host}
    const app = new AppController(appHost)
    appHost.app = app
    app.store = { data: {...storeConfig}, setData: sinon.spy() }
    app.resultError = sinon.spy()
    app.resultError = sinon.spy()
    app.currentModule = sinon.spy()
    app.requestData = sinon.spy(async () => ({ error: {} }))
    await app.tokenLogin({ code: "code" })
    sinon.assert.callCount(app.resultError, 1);

    await app.tokenLogin({ access_token: "access_token" })
    sinon.assert.callCount(app.resultError, 2);
  })

})