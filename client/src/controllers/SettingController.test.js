import sinon from 'sinon'
import { expect } from '@open-wc/testing';

import { SettingController } from './SettingController.js'
import { store as storeConfig } from '../config/app.js'
import { SIDE_EVENT, APP_MODULE, SETTING_EVENT, MODAL_EVENT, ACTION_EVENT } from '../config/enums.js'
import { Forms } from './Forms.js'
import { Sql } from './Sql.js'
import { Dataset } from './Dataset.js'
import { Default as SettingData } from '../components/Setting/Form/Form.stories.js'

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
          { id: 1, groupname: "nervatype", groupvalue: "all" }
        ]
      }
    }
  },
  setData: sinon.spy(),
}
const app = {
  store,
  modules: {
    forms: Forms({ msg: (key)=> key, getSetting: (key)=>storeConfig.ui[key] }),
    sql: Sql({ msg: (key)=> key }),
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
  currentModule: sinon.spy(),
}

describe('SettingController', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('onSideEvent', async () => {
    let testApp = {
      ...app,
      showHelp: sinon.spy(),
      currentModule: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy(),
      },
    }

    let setting = new SettingController({...host, app: testApp})
    setting.onSideEvent({ key: SIDE_EVENT.CHANGE, data: { fieldname: "", value: "" } })
    sinon.assert.callCount(testApp.store.setData, 1);

    setting.checkSetting = sinon.spy()
    setting.onSideEvent({ key: SIDE_EVENT.BACK, data: {} })
    sinon.assert.callCount(setting.checkSetting, 1);

    setting.loadSetting = sinon.spy()
    setting.saveSetting = sinon.spy(async () => ({current: { form: { id: 1 } }}))
    await setting.onSideEvent({ key: SIDE_EVENT.SAVE, data: {} })
    sinon.assert.callCount(setting.loadSetting, 1);

    setting.deleteSetting = sinon.spy()
    setting.onSideEvent({ key: SIDE_EVENT.DELETE, data: { value: 1 } })
    sinon.assert.callCount(setting.deleteSetting, 1);

    setting.onSideEvent({ key: SIDE_EVENT.CHECK, data: [{}] })
    sinon.assert.callCount(setting.checkSetting, 2);
    setting.onSideEvent({ key: SIDE_EVENT.CHECK, data: [{ ntype: "type" }] })
    sinon.assert.callCount(testApp.currentModule, 1);

    setting.onSideEvent({ key: SIDE_EVENT.LOAD_SETTING, data: {} })
    sinon.assert.callCount(setting.loadSetting, 2);

    setting.setProgramForm = sinon.spy()
    setting.onSideEvent({ key: SIDE_EVENT.PROGRAM_SETTING, data: {} })
    sinon.assert.callCount(setting.setProgramForm, 1);

    setting.onSideEvent({ key: SIDE_EVENT.PASSWORD, data: { username: "admin" } })
    sinon.assert.callCount(setting.checkSetting, 3);

    setting.onSideEvent({ key: SIDE_EVENT.HELP, data: { value: "help" } })
    sinon.assert.callCount(testApp.showHelp, 1);

    setting.onSideEvent({ key: "", data: {} })

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
              type: "employee"
            },
            dataset: {
              employee: [
                { username: "admin" }
              ]
            }
          },
          [APP_MODULE.SETTING]: {
            ...testApp.store.data[APP_MODULE.SETTING],
            type: "password"
          }
        }
      }
    }
    setting = new SettingController({...host, app: testApp})
    setting.loadSetting = sinon.spy()
    setting.onSideEvent({ key: SIDE_EVENT.BACK, data: {} })
    sinon.assert.callCount(setting.loadSetting, 1);

    setting.changePassword = sinon.spy()
    setting.onSideEvent({ key: SIDE_EVENT.SAVE, data: {} })
    sinon.assert.callCount(setting.changePassword, 1);

    setting.checkSetting = sinon.spy()
    setting.onSideEvent({ key: SIDE_EVENT.PASSWORD, data: {} })
    sinon.assert.callCount(setting.checkSetting, 1);

  })

  it('onSettingEvent', async () => {
    const testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy(),
      },
    }

    const setting = new SettingController({...host, app: testApp})
    setting.onSettingEvent({ key: SETTING_EVENT.CURRENT_PAGE, data: { value: "" } })
    sinon.assert.callCount(testApp.store.setData, 1);

    setting.setFormActions = sinon.spy()
    setting.onSettingEvent({ key: SETTING_EVENT.FORM_ACTION, data: {} })
    sinon.assert.callCount(setting.setFormActions, 1);

    setting.editItem = sinon.spy()
    setting.onSettingEvent({ key: SETTING_EVENT.EDIT_ITEM, data: {} })
    sinon.assert.callCount(setting.editItem, 1);

    setting.onSettingEvent({ key: "", data: {} })

  })

  it('changePassword', async () => {
    let testApp = {
      ...app,
      showToast: sinon.spy(),
      requestData: () => ({}),
      resultError: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.SETTING]: {
            ...app.store.data[APP_MODULE.SETTING],
            current: {...app.store.data[APP_MODULE.SETTING].current,
              form: {
                username: ""
              }
            }
          }
        }
      },
    }

    // missing username
    let setting = new SettingController({...host, app: testApp})
    await setting.changePassword()
    sinon.assert.callCount(testApp.showToast, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.SETTING]: {
            ...testApp.store.data[APP_MODULE.SETTING],
            current: {
              ...testApp.store.data[APP_MODULE.SETTING].current,
              form: {
                username: "username",
                password_1: "1",
                password_2: "2"
              }
            }
          }
        }
      }
    }

    // not match
    setting = new SettingController({...host, app: testApp})
    await setting.changePassword()
    sinon.assert.callCount(testApp.showToast, 2);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.SETTING]: {
            ...testApp.store.data[APP_MODULE.SETTING],
            current: {
              ...testApp.store.data[APP_MODULE.SETTING].current,
              form: {
                username: "username",
                password_1: "password",
                password_2: "password"
              }
            }
          }
        }
      }
    }
    // ok
    setting = new SettingController({...host, app: testApp})
    await setting.changePassword()
    sinon.assert.callCount(testApp.showToast, 3);

    testApp = {
      ...testApp,
      requestData: () => ({ error: {} }),
    }
    // error
    setting = new SettingController({...host, app: testApp})
    await setting.changePassword()
    sinon.assert.callCount(testApp.resultError, 1);

  })

  it('checkSetting', async () => {
    const testApp = {
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
          [APP_MODULE.SETTING]: {
            ...app.store.data[APP_MODULE.SETTING],
            current: {...app.store.data[APP_MODULE.SETTING].current,
              form: {
                id: null,
                fieldname: "value",
                transfilter: 1,
                translink: null
              }
            },
            dirty: true
          }
        }
      },
    }

    const setting = new SettingController({...host, app: testApp})
    setting.loadSetting = sinon.spy()
    setting.saveSetting = sinon.spy(async () => ({}))
    await setting.checkSetting({ type: "setting", id: null }, SIDE_EVENT.LOAD_SETTING)
    sinon.assert.callCount(setting.loadSetting, 2);

    setting.saveSetting = sinon.spy(async () => ({ error: {} }))
    setting.setPasswordForm = sinon.spy()
    await setting.checkSetting({ username: "admin" }, SIDE_EVENT.PASSWORD_FORM)
    sinon.assert.callCount(setting.setPasswordForm, 2);

  })

  it('deleteSetting', async () => {
    let testApp = {
      ...app,
      showToast: sinon.spy(),
      requestData: sinon.spy(async () => ({ state: [{ sco: 0 }] })),
      resultError: sinon.spy(),
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
          [APP_MODULE.SETTING]: {
            ...app.store.data[APP_MODULE.SETTING],
            type: "usergroup"
          }
        }
      },
    }

    let setting = new SettingController({...host, app: testApp})
    setting.loadSetting = sinon.spy()
    await setting.deleteSetting({ id: 1 })
    sinon.assert.callCount(testApp.requestData, 2);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
        if(path === "/view"){
          return { state: [{ sco: 0 }]  }
        }
        return { error: {} }
      }),
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.SETTING]: {
            ...testApp.store.data[APP_MODULE.SETTING],
            type: "currency"
          }
        }
      }
    }
    setting = new SettingController({...host, app: testApp})
    await setting.deleteSetting({ id: 1 })
    sinon.assert.callCount(testApp.requestData, 2);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ state: [{ sco: 1 }] })),
      store: {
        ...testApp.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: {} })
          }
        })
      }
    }
    setting = new SettingController({...host, app: testApp})
    await setting.deleteSetting({ id: 1 })
    sinon.assert.callCount(testApp.showToast, 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} })),
      resultError: sinon.spy(),
    }
    setting = new SettingController({...host, app: testApp})
    await setting.deleteSetting({ id: 1 })
    sinon.assert.callCount(testApp.resultError, 1)

    testApp = {
      ...testApp,
      resultError: sinon.spy(),
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.SETTING]: {
            ...testApp.store.data[APP_MODULE.SETTING],
            type: "setting"
          }
        }
      }
    }
    setting = new SettingController({...host, app: testApp})
    await setting.deleteSetting({ id: 1 })
    sinon.assert.callCount(testApp.resultError, 1)

  })

  it('editItem', () => {
    let testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.SETTING]: {
            ...app.store.data[APP_MODULE.SETTING],
            current: {
              form: {},
              fieldvalue: {}
            },
            audit: "all"
          }
        }
      },
    }

    let setting = new SettingController({...host, app: testApp})
    setting.editItem({ name: "fieldname", value: "value" })
    sinon.assert.callCount(testApp.store.setData, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.SETTING]: {
            ...testApp.store.data[APP_MODULE.SETTING],
            audit: "readonly"
          }
        }
      }
    }
    setting = new SettingController({...host, app: testApp})
    setting.editItem({ name: "fieldvalue_value", value: "value" })
    sinon.assert.callCount(testApp.store.setData, 2);

    setting.loadLog = sinon.spy()
    setting.editItem({ name: "log_search" })
    sinon.assert.callCount(setting.loadLog, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.SETTING]: {
            ...testApp.store.data[APP_MODULE.SETTING],
            type: "program"
          }
        }
      }
    }
    setting = new SettingController({...host, app: testApp})
    setting.editItem({ name: "fieldname", value: "value" })
    sinon.assert.callCount(testApp.store.setData, 4);

  })

  it('loadLog', async () => {
    let testApp = {
      ...app,
      showToast: sinon.spy(),
      requestData: sinon.spy(async () => ({})),
      resultError: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.SETTING]: {
            ...app.store.data[APP_MODULE.SETTING],
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
          }
        }
      },
    }

    let setting = new SettingController({...host, app: testApp})
    await setting.loadLog()
    sinon.assert.callCount(testApp.showToast, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.SETTING]: {
            ...testApp.store.data[APP_MODULE.SETTING],
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
          }
        }
      }
    }
    setting = new SettingController({...host, app: testApp})
    await setting.loadLog()
    sinon.assert.callCount(testApp.store.setData, 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} })),
      resultError: sinon.spy(),
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.SETTING]: {
            ...testApp.store.data[APP_MODULE.SETTING],
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
          }
        }
      }
    }
    setting = new SettingController({...host, app: testApp})
    await setting.loadLog()
    sinon.assert.callCount(testApp.resultError, 1);

  })

  it('loadSetting', async () => {
    let testApp = {
      ...app,
      requestData: sinon.spy(async () => ({ placetype: [ { id: 1, groupvalue: "warehouse" } ] })),
      resultError: sinon.spy(),
      currentModule: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy(),
      },
    }

    let setting = new SettingController({...host, app: testApp})
    setting.setSettingData = sinon.spy()

    await setting.loadSetting({ type: "setting" })
    sinon.assert.callCount(setting.setSettingData, 1);

    await setting.loadSetting({ type: "usergroup" })
    sinon.assert.callCount(setting.setSettingData, 2);

    await setting.loadSetting({ type: "usergroup", id: null })
    sinon.assert.callCount(setting.setSettingData, 3);

    await setting.loadSetting({ type: "template", id: 1 })
    sinon.assert.callCount(testApp.currentModule, 1);

    await setting.loadSetting({ type: "place", id: null })
    sinon.assert.callCount(setting.setSettingData, 4);

    await setting.loadSetting({ type: "log" })
    sinon.assert.callCount(setting.setSettingData, 5);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} })),
      resultError: sinon.spy(),
      getSql: () =>({
        sql: "",
        prmCount: 0
      })
    }
    setting = new SettingController({...host, app: testApp})
    await setting.loadSetting({ type: "setting", id: 1 })
    sinon.assert.callCount(testApp.resultError, 1);

  })

  it('saveSetting', async () => {
    let testApp = {
      ...app,
      modules: {
        ...app.modules,
        initItem: ()=>({}),
        validator: () => ({})
      },
      requestData: sinon.spy(async () => ({})),
      resultError: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.SETTING]: {
            ...app.store.data[APP_MODULE.SETTING],
            current: {
              form: {
                id: null,
                fieldname: "value",
                transfilter: 1,
                translink: null
              }
            }
          }
        }
      },
    }

    let setting = new SettingController({...host, app: testApp})
    let result = await setting.saveSetting()
    expect(result).to.exist;

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.SETTING]: {
            ...testApp.store.data[APP_MODULE.SETTING],
            type: "usergroup",
            ntype: "usergroup"
          }
        }
      }
    }
    setting = new SettingController({...host, app: testApp})
    result = await setting.saveSetting()
    expect(result).to.exist;

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.SETTING]: {
            ...testApp.store.data[APP_MODULE.SETTING],
            current: {
              form: {
                id: 1,
                fieldname: "value",
                transfilter: 1,
                translink: 2
              }
            }
          },
          [APP_MODULE.LOGIN]: {
            ...testApp.store.data[APP_MODULE.LOGIN],
            data: {
              groups: [
                { id: 1, groupname: "transfilter", groupvalue: "update" },
                { id: 2, groupname: "nervatype", groupvalue: "groups" },
              ]
            }
          }
        }
      }
    }
    setting = new SettingController({...host, app: testApp})
    result = await setting.saveSetting()
    expect(result).to.exist;

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} })),
      resultError: sinon.spy(),
    }
    setting = new SettingController({...host, app: testApp})
    result = await setting.saveSetting()
    sinon.assert.callCount(testApp.resultError, 1);

    testApp = {
      ...testApp,
      requestData: sinon.spy(async (path) => {
        if(path === "/link"){
          return { error: {} }
        }
        return {}
      }),
      resultError: sinon.spy(),
    }
    setting = new SettingController({...host, app: testApp})
    result = await setting.saveSetting()
    sinon.assert.callCount(testApp.resultError, 1);

    testApp = {
      ...testApp,
      modules: {
        ...app.modules,
        initItem: ()=>({}),
        validator: () => ({ error: {} })
      },
      resultError: sinon.spy(),
    }
    setting = new SettingController({...host, app: testApp})
    result = await setting.saveSetting()
    sinon.assert.callCount(testApp.resultError, 1);

  })

  it('setPasswordForm', () => {
    const testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy(),
      },
    }

    const setting = new SettingController({...host, app: testApp})
    setting.setPasswordForm("username")
    sinon.assert.callCount(testApp.store.setData, 1);

  })

  it('setProgramForm', () => {
    const testApp = {
      ...app,
      getSetting: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy(),
      },
    }

    const setting = new SettingController({...host, app: testApp})
    setting.setProgramForm()
    sinon.assert.callCount(testApp.store.setData, 1);

  })

  it('setSettingData', () => {
    let testApp = {
      ...app,
      modules: {
        ...app.modules,
        initItem: ()=>({}),
      },
      getAuditFilter: sinon.spy(() => (["all",1])),
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.SETTING]: {
            ...app.store.data[APP_MODULE.SETTING],
            type: "setting",
          }
        }
      },
    }

    let setting = new SettingController({...host, app: testApp})
    setting.setSettingForm = sinon.spy()
    setting.setSettingData({ type: "setting", dataset: {} })
    sinon.assert.callCount(testApp.store.setData, 1);

    setting.setSettingData({ 
      type: "log", dataset: { log_view: [] }, id: 1 
    })
    sinon.assert.callCount(setting.setSettingForm, 1);

    testApp = {
      ...app,
      modules: {
        ...app.modules,
        initItem: ()=>({}),
      },
      getAuditFilter: sinon.spy(() => (["update",1])),
    }
    setting = new SettingController({...host, app: testApp})
    setting.setFormActions = sinon.spy()
    setting.setSettingData({ type: "usergroup", dataset: {}, id: 1 })
    sinon.assert.callCount(setting.setFormActions, 1);

  })

  it('setSettingForm', () => {
    let testApp = {
      ...app,
      modules: {
        ...app.modules,
        initItem: ()=>({}),
      },
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.SETTING]: {
            ...app.store.data[APP_MODULE.SETTING],
            ...SettingData.args.data,
            audit: "readonly"
          }
        }
      },
    }

    let setting = new SettingController({...host, app: testApp})
    setting.setSettingForm(null)
    sinon.assert.callCount(testApp.store.setData, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.SETTING]: {
            ...testApp.store.data[APP_MODULE.SETTING],
            audit: "all"
          }
        }
      }
    }
    setting = new SettingController({...host, app: testApp})
    setting.setSettingForm(null)
    sinon.assert.callCount(testApp.store.setData, 2);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.SETTING]: {
            ...testApp.store.data[APP_MODULE.SETTING],
            audit: "all"
          }
        }
      }
    }
    setting = new SettingController({...host, app: testApp})
    setting.setSettingForm(1, testApp.store.data.setting)
    sinon.assert.callCount(testApp.store.setData, 3);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.SETTING]: {
            ...testApp.store.data[APP_MODULE.SETTING],
            type: "setting",
            ntype: "fieldvalue",
            dataset: {
              setting_view: [{ id: 1, fieldtype: "urlink", valuelist:"", notes: "" }]
            }
          }
        }
      }
    }
    setting = new SettingController({...host, app: testApp})
    setting.setSettingForm(1)
    sinon.assert.callCount(testApp.store.setData, 4);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.SETTING]: {
            ...testApp.store.data[APP_MODULE.SETTING],
            dataset: {
              setting_view: [{ id: 1, fieldtype: "valuelist", valuelist:"a|b|c", notes: "" }]
            }
          }
        }
      }
    }
    setting = new SettingController({...host, app: testApp})
    setting.setSettingForm(1)
    sinon.assert.callCount(testApp.store.setData, 5);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.SETTING]: {
            ...testApp.store.data[APP_MODULE.SETTING],
            type: "usergroup",
            ntype: null,
            dataset: {
              usergroup_view: [{ id: 1, transfilter: null }],
              transfilter: [
                { id: 1, groupname: "transfilter", groupvalue: "all" },
              ]
            }
          }
        }
      }
    }
    setting = new SettingController({...host, app: testApp})
    setting.setSettingForm(1, testApp.store.data.setting)
    sinon.assert.callCount(testApp.store.setData, 6);

  })

  it('setFormActions', async () => {
    let testApp = {
      ...app,
      modules: {
        ...app.modules,
        initItem: ()=>({}),
      },
      requestData: sinon.spy(async () => ({})),
      resultError: sinon.spy(),
      currentModule: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.SETTING]: {
            ...app.store.data[APP_MODULE.SETTING],
            type: "template"
          }
        }
      },
    }

    let setting = new SettingController({...host, app: testApp})
    setting.checkSetting = sinon.spy()
    setting.deleteSetting = sinon.spy()
    setting.setFormActions({ params: { action: ACTION_EVENT.NEW_ITEM } })
    sinon.assert.callCount(setting.checkSetting, 1);

    setting.setFormActions({ params: { action: ACTION_EVENT.DELETE_ITEM } })
    sinon.assert.callCount(setting.deleteSetting, 1)

    setting.setFormActions({ params: { action: ACTION_EVENT.EDIT_ITEM, row: { id: 1 } } })
    sinon.assert.callCount(setting.checkSetting, 2)

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.SETTING]: {
            ...testApp.store.data[APP_MODULE.SETTING],
            type: "place"
          }
        }
      }
    }
    setting = new SettingController({...host, app: testApp})
    setting.setFormActions({ params: { action: ACTION_EVENT.EDIT_ITEM, row: { id: 1 } } })
    sinon.assert.callCount(testApp.currentModule, 1)

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.SETTING]: {
            ...testApp.store.data[APP_MODULE.SETTING],
            type: "usergroup"
          }
        }
      }
    }
    setting = new SettingController({...host, app: testApp})
    setting.setSettingForm = sinon.spy()
    await setting.setFormActions({ params: { action: ACTION_EVENT.EDIT_ITEM, row: { id: 1 } } })
    sinon.assert.callCount(setting.setSettingForm, 1)

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} })),
      resultError: sinon.spy(),
    }
    setting = new SettingController({...host, app: testApp})
    await setting.setFormActions({ params: { action: ACTION_EVENT.EDIT_ITEM, row: { id: 1 } } })
    sinon.assert.callCount(testApp.resultError, 1)

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({})),
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.SETTING]: {
            ...testApp.store.data[APP_MODULE.SETTING],
            type: "ui_menu"
          }
        }
      }
    }
    setting = new SettingController({...host, app: testApp})
    setting.setSettingForm = sinon.spy()
    await setting.setFormActions({ params: { action: ACTION_EVENT.EDIT_ITEM, row: { id: 1 } } })
    sinon.assert.callCount(setting.setSettingForm, 1)

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} })),
      resultError: sinon.spy(),
    }
    setting = new SettingController({...host, app: testApp})
    await setting.setFormActions({ params: { action: ACTION_EVENT.EDIT_ITEM, row: { id: 1 } } })
    sinon.assert.callCount(testApp.resultError, 1)

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({})),
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.SETTING]: {
            ...testApp.store.data[APP_MODULE.SETTING],
            type: "setting"
          }
        }
      }
    }
    setting = new SettingController({...host, app: testApp})
    setting.setSettingForm = sinon.spy()
    await setting.setFormActions({ params: { action: ACTION_EVENT.EDIT_ITEM, row: { id: 1 } } })
    sinon.assert.callCount(setting.setSettingForm, 1)

    testApp = {
      ...testApp,
      modules: {
        ...app.modules,
        initItem: ()=>({ id: null, fieldname: null, }),
      },
      requestData: sinon.spy(async () => ({})),
      store: {
        ...testApp.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: { value: { id: 1 } } })
          }
        }),
        data: {
          ...testApp.store.data,
          [APP_MODULE.SETTING]: {
            ...testApp.store.data[APP_MODULE.SETTING],
            current: {
              form: {
                id: 1
              }
            },
            dataset: {
              nervatype: [
                { id: 1, groupvalue: "customer" }
              ],
              inputfilter: [
                { id: 1, groupvalue: "all" }
              ],
              transtype: [
                { id: 1, groupvalue: "trans" }
              ],
              reportkey: [
                { id: 1, groupvalue: "report" }
              ],
              menukey: [
                { id: 1, groupvalue: "menu" }
              ]
            }
          }
        }
      }
    }
    setting = new SettingController({...host, app: testApp})
    setting.module.modalAudit = (prm)=>(prm)
    setting.setSettingForm = sinon.spy()
    await setting.setFormActions({ params: { action: ACTION_EVENT.EDIT_AUDIT, row: { id: 1, subtype: null, supervisor: 1 } } })
    sinon.assert.callCount(setting.setSettingForm, 1)

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} })),
      resultError: sinon.spy(),
    }
    setting = new SettingController({...host, app: testApp})
    setting.module.modalAudit = (prm)=>(prm)
    await setting.setFormActions({ params: { action: ACTION_EVENT.EDIT_AUDIT, row: { id: 1, subtype: null, supervisor: 1 } } })
    sinon.assert.callCount(testApp.resultError, 1)

    testApp = {
      ...testApp,
      modules: {
        ...app.modules,
        initItem: ()=>({ id: null, fieldname: null, }),
      },
      requestData: sinon.spy(async () => ({})),
      store: {
        ...testApp.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: { value: { id: 1 } } })
          }
        }),
        data: {
          ...testApp.store.data,
          [APP_MODULE.SETTING]: {
            ...testApp.store.data[APP_MODULE.SETTING],
            current: {
              form: {
                id: 1
              }
            },
            dataset: {
              fieldtype: [
                { id: 1, groupvalue: "string" }
              ],
              ui_menufields: [
                { id: 1, groupvalue: "menu" }
              ]
            }
          }
        }
      }
    }
    setting = new SettingController({...host, app: testApp})
    setting.module.modalMenu = (prm)=>(prm)
    setting.setSettingForm = sinon.spy()
    await setting.setFormActions({ params: { action: ACTION_EVENT.EDIT_MENU_FIELD, row: { id: 1 } } })
    sinon.assert.callCount(setting.setSettingForm, 1)

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} })),
      resultError: sinon.spy(),
    }
    setting = new SettingController({...host, app: testApp})
    setting.module.modalMenu = (prm)=>(prm)
    await setting.setFormActions({ params: { action: ACTION_EVENT.EDIT_MENU_FIELD, row: { id: 1 } } })
    sinon.assert.callCount(testApp.resultError, 1)

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({})),
      store: {
        ...testApp.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: { value: { id: 1 } } })
          }
        }),
      }
    }
    setting = new SettingController({...host, app: testApp})
    setting.setSettingForm = sinon.spy()
    await setting.setFormActions({ params: { action: ACTION_EVENT.DELETE_ITEM_ROW, table: "table" } })
    sinon.assert.callCount(setting.setSettingForm, 1)

    testApp = {
      ...testApp,
      requestData: sinon.spy(async () => ({ error: {} })),
      resultError: sinon.spy(),
    }
    setting = new SettingController({...host, app: testApp})
    setting.module.modalMenu = (prm)=>(prm)
    await setting.setFormActions({ params: { action: ACTION_EVENT.DELETE_ITEM_ROW, table: "table" } })
    sinon.assert.callCount(testApp.resultError, 1)

    setting.setFormActions({ params: { action: "" } })

  })

  it('setModule', () => {
    const setting = new SettingController({...host, app})
    setting.setModule({})
    expect(setting.module).to.exist
  })

})