import sinon from 'sinon'
import { expect } from '@open-wc/testing';

import { TemplateController, getDataList, getDataTable, getElementType } from './TemplateController.js'
import { store as storeConfig } from '../config/app.js'
import { SIDE_EVENT, APP_MODULE, TEMPLATE_EVENT, MODAL_EVENT } from '../config/enums.js'
import { Default as TemplateData } from '../components/Template/Editor/Editor.stories.js'
import { sample } from '../config/sample.js';

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
    },
    [APP_MODULE.TEMPLATE]: {
      ...TemplateData.args.data
    }
  },
  setData: sinon.spy(),
}
const app = {
  store,
  msg: (value)=>value,
  requestData: () => ({}),
  resultError: sinon.spy(),
  showHelp: sinon.spy(),
  showToast: sinon.spy(),
  getSetting: sinon.spy(),
  currentModule: sinon.spy(),
}
const sampleTemplate = () => JSON.parse(JSON.stringify(sample))

describe('TemplateController', () => {
  beforeEach(async () => {
    Object.defineProperty(URL, 'createObjectURL', { value: sinon.spy() })
    Object.defineProperty(window, 'open', { value: sinon.spy() })
  });

  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('onSideEvent', async () => {
    const testApp = {
      ...app,
      showHelp: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy(),
      },
    }

    const template = new TemplateController({...host, app: testApp})
    template.onSideEvent({ key: SIDE_EVENT.CHANGE, data: { fieldname: "", value: "" } })
    sinon.assert.callCount(testApp.store.setData, 1);

    template.saveTemplate = sinon.spy()
    template.onSideEvent({ key: SIDE_EVENT.SAVE, data: {} })
    sinon.assert.callCount(template.saveTemplate, 1);

    template.createTemplate = sinon.spy()
    template.onSideEvent({ key: SIDE_EVENT.CREATE_REPORT, data: {} })
    sinon.assert.callCount(template.createTemplate, 1);

    template.deleteTemplate = sinon.spy()
    template.onSideEvent({ key: SIDE_EVENT.DELETE, data: {} })
    sinon.assert.callCount(template.deleteTemplate, 1);

    template.checkTemplate = sinon.spy()
    template.onSideEvent({ key: SIDE_EVENT.CHECK, data: {} })
    sinon.assert.callCount(template.checkTemplate, 1);

    template.showPreview = sinon.spy()
    template.onSideEvent({ key: SIDE_EVENT.REPORT_SETTINGS, data: {} })
    sinon.assert.callCount(template.showPreview, 1);

    template.exportTemplate = sinon.spy()
    template.onSideEvent({ key: SIDE_EVENT.REPORT_SETTINGS, data: { value: "JSON" } })
    sinon.assert.callCount(template.exportTemplate, 1);

    template.onSideEvent({ key: SIDE_EVENT.HELP, data: { value: "help" } })
    sinon.assert.callCount(testApp.showHelp, 1);

    template.onSideEvent({ key: "", data: {} })

  })

  it('onTemplateEvent', async () => {
    const testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy(),
      },
    }

    const template = new TemplateController({...host, app: testApp})
    template.addItem = sinon.spy()
    template.onTemplateEvent({ key: TEMPLATE_EVENT.ADD_ITEM, data: {} })
    sinon.assert.callCount(template.addItem, 1);

    template.onTemplateEvent({ key: TEMPLATE_EVENT.CHANGE_TEMPLATE, data: { key: "", value: "" } })
    sinon.assert.callCount(testApp.store.setData, 1);

    template.onTemplateEvent({ key: TEMPLATE_EVENT.CHANGE_CURRENT, data: { key: "", value: "" } })
    sinon.assert.callCount(testApp.store.setData, 2);

    template.goPrevious = sinon.spy()
    template.onTemplateEvent({ key: TEMPLATE_EVENT.GO_PREVIOUS, data: {} })
    sinon.assert.callCount(template.goPrevious, 1);

    template.goNext = sinon.spy()
    template.onTemplateEvent({ key: TEMPLATE_EVENT.GO_NEXT, data: {} })
    sinon.assert.callCount(template.goNext, 1);

    template.createMap = sinon.spy()
    template.onTemplateEvent({ key: TEMPLATE_EVENT.CREATE_MAP, data: {} })
    sinon.assert.callCount(template.createMap, 1);

    template.setCurrent = sinon.spy()
    template.onTemplateEvent({ key: TEMPLATE_EVENT.SET_CURRENT, data: [] })
    sinon.assert.callCount(template.setCurrent, 1);

    template.moveUp = sinon.spy()
    template.onTemplateEvent({ key: TEMPLATE_EVENT.MOVE_UP, data: {} })
    sinon.assert.callCount(template.moveUp, 1);

    template.moveDown = sinon.spy()
    template.onTemplateEvent({ key: TEMPLATE_EVENT.MOVE_DOWN, data: {} })
    sinon.assert.callCount(template.moveDown, 1);

    template.deleteItem = sinon.spy()
    template.onTemplateEvent({ key: TEMPLATE_EVENT.DELETE_ITEM, data: {} })
    sinon.assert.callCount(template.deleteItem, 1);

    template.editItem = sinon.spy()
    template.onTemplateEvent({ key: TEMPLATE_EVENT.EDIT_ITEM, data: {} })
    sinon.assert.callCount(template.editItem, 1);

    template.editDataItem = sinon.spy()
    template.onTemplateEvent({ key: TEMPLATE_EVENT.EDIT_DATA_ITEM, data: {} })
    sinon.assert.callCount(template.editDataItem, 1);

    template.setCurrentData = sinon.spy()
    template.onTemplateEvent({ key: TEMPLATE_EVENT.SET_CURRENT_DATA, data: {} })
    sinon.assert.callCount(template.setCurrentData, 1);

    template.setCurrentDataItem = sinon.spy()
    template.onTemplateEvent({ key: TEMPLATE_EVENT.SET_CURRENT_DATA_ITEM, data: {} })
    sinon.assert.callCount(template.setCurrentDataItem, 1);

    template.addTemplateData = sinon.spy()
    template.onTemplateEvent({ key: TEMPLATE_EVENT.ADD_TEMPLATE_DATA, data: {} })
    sinon.assert.callCount(template.addTemplateData, 1);

    template.deleteData = sinon.spy()
    template.onTemplateEvent({ key: TEMPLATE_EVENT.DELETE_DATA, data: {} })
    sinon.assert.callCount(template.deleteData, 1);

    template.deleteDataItem = sinon.spy()
    template.onTemplateEvent({ key: TEMPLATE_EVENT.DELETE_DATA_ITEM, data: {} })
    sinon.assert.callCount(template.deleteDataItem, 1);

    template.onTemplateEvent({ key: "", data: {} })

  })

  it('addItem', async () => {
    const testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...app.store.data[APP_MODULE.TEMPLATE],
            current: {...app.store.data[APP_MODULE.TEMPLATE].current,
              item: []
            }
          }
        }
      },
    }

    const template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.addItem("cell")
    sinon.assert.callCount(template.setCurrent, 1);

    template.addItem("row")
    sinon.assert.callCount(template.setCurrent, 2);

    template.addItem("")
    sinon.assert.callCount(template.setCurrent, 2);

  })

  it('addTemplateData', async () => {
    const testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: {} })
          }
        })
      },
    }

    const template = new TemplateController({...host, app: testApp})
    template.module.modalTemplate = (prm)=>(prm)
    template.setCurrentData = sinon.spy()
    template.addTemplateData()
    sinon.assert.callCount(template.setCurrentData, 1);

  })

  it('checkTemplate', async () => {
    let testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL })
            //data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: { value: "value" } })
          }
        }),
        data: {
          ...app.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...app.store.data[APP_MODULE.TEMPLATE],
            dirty: true
          }
        }
      },
    }

    let template = new TemplateController({...host, app: testApp})
    template.setTemplate = sinon.spy()
    template.checkTemplate(SIDE_EVENT.BLANK)
    sinon.assert.callCount(template.setTemplate, 1);

    template.checkTemplate(SIDE_EVENT.SAMPLE)
    sinon.assert.callCount(template.setTemplate, 2);

  })

  it('createMap', async () => {
    const app_template = app.store.data[APP_MODULE.TEMPLATE].template
    let testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...app.store.data[APP_MODULE.TEMPLATE],
            current: {
              item: app_template.details,
              parent: null,
            }
          }
        }
      },
    }
    const canvas = {
      height: 0, width: 0,
      getContext: () => ({
        fillStyle: "",
        clearRect: sinon.spy(),
        fillRect: sinon.spy(),
        beginPath: sinon.spy(),
        moveTo: sinon.spy(),
        lineTo: sinon.spy(),
        stroke: sinon.spy(),
      })
    }

    let template = new TemplateController({...host, app: testApp})
    template.createMap(canvas)
    sinon.assert.callCount(testApp.store.setData, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        setData: sinon.spy(),
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              item: app_template.details[1].row.columns[0].cell,
              parent: app_template.details[1].row.columns,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.createMap(canvas)
    sinon.assert.callCount(testApp.store.setData, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        setData: sinon.spy(),
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              item: app_template.details[19].datagrid.columns[1].column,
              parent: app_template.details[19].datagrid.columns,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.createMap(canvas)
    sinon.assert.callCount(testApp.store.setData, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        setData: sinon.spy(),
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            template: {
              meta: {},
              report: {},
              header: [],
              details: [],
              footer: [],
              sources: {},
              data: {}
            },
            current: {},
            mapRef: canvas
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.createMap()
    sinon.assert.callCount(testApp.store.setData, 0);

  })

  it('createTemplate', async () => {
    let testApp = {
      ...app,
      modules: {
        ...app.modules,
        initItem: ()=>({ id: null }),
      },
      resultError: sinon.spy(),
      requestData: sinon.spy(() => ({})),
      currentModule: sinon.spy(),
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
          [APP_MODULE.TEMPLATE]: {
            ...app.store.data[APP_MODULE.TEMPLATE],
            dbtemp: { id: 1, test: "value", data: { report_name: "value" } }
          }
        }
      },
    }

    let template = new TemplateController({...host, app: testApp})
    await template.createTemplate()
    sinon.assert.callCount(testApp.currentModule, 1);

    testApp = {
      ...testApp,
      resultError: sinon.spy(),
      requestData: sinon.spy(() => ({ error: {} })),
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            template: {
              ...testApp.store.data[APP_MODULE.TEMPLATE].template,
              meta: {
                ...testApp.store.data[APP_MODULE.TEMPLATE].template.meta,
                report_type: "trans"
              }
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    await template.createTemplate()
    sinon.assert.callCount(testApp.resultError, 1);

  })

  it('deleteData', () => {
    let testApp = {
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
          [APP_MODULE.TEMPLATE]: {
            ...app.store.data[APP_MODULE.TEMPLATE],
            key: "test"
          }
        }
      },
    }

    let template = new TemplateController({...host, app: testApp})
    template.deleteData()
    sinon.assert.callCount(testApp.store.setData, 4);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            template: {
              ...testApp.store.data[APP_MODULE.TEMPLATE].template,
              key: "_sample"
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.deleteData()
    sinon.assert.callCount(testApp.store.setData, 8);

  })

  it('deleteDataItem', () => {
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
          [APP_MODULE.TEMPLATE]: {
            ...app.store.data[APP_MODULE.TEMPLATE],
            key: "_sample",
            current_data: {
              name: "labels",
              type: "list",
              items: getDataList(sampleTemplate().data.labels)
            }
          }
        }
      },
    }

    let template = new TemplateController({...host, app: testApp})
    template.deleteDataItem({ key: "title" })
    sinon.assert.callCount(testApp.store.setData, 4);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: {} })
          }
        }),
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            key: "template",
            current_data: {
              name: "items",
              type: "table",
              items: getDataTable(sampleTemplate().data.items)
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.deleteDataItem({ _index: 1 })
    sinon.assert.callCount(testApp.store.setData, 4);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: {} })
          }
        }),
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            key: "template",
            current_data: {
              name: "test",
              type: "table",
              items: getDataTable([{ col1: 1, col2: "col2" }])
            },
            template: {
              data: {
                test: [{ col1: 1, col2: "col2" }]
              }
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.deleteDataItem({ _index: 0 })
    sinon.assert.callCount(testApp.store.setData, 4);

  })

  it('deleteItem', () => {
    let testApp = {
      ...app,
      store: {
        ...app.store,
        data: {
          ...app.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...app.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "details",
              type: "vgap",
              item: sampleTemplate().details[0].vgap,
              index: 0,
              parent: sampleTemplate().details,
              parent_type: "details",
              parent_index: null,
            }
          }
        }
      },
    }

    let template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.deleteItem()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "details",
              type: "cell",
              item: sampleTemplate().details[5].row.columns[1].cell,
              index: 1,
              parent: sampleTemplate().details[5].row.columns,
              parent_type: "row",
              parent_index: 5,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.deleteItem()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "header",
              type: "row",
              item: sampleTemplate().header[0].row,
              index: null,
              parent: sampleTemplate().header,
              parent_type: "header",
              parent_index: null,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.deleteItem()
    sinon.assert.callCount(template.setCurrent, 0);

  })

  it('deleteTemplate', async () => {
    let testApp = {
      ...app,
      resultError: sinon.spy(),
      requestData: sinon.spy(() => ({})),
      currentModule: sinon.spy(),
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
          [APP_MODULE.TEMPLATE]: {
            ...app.store.data[APP_MODULE.TEMPLATE],
            dbtemp: { id: 1, code: "code", data: { report_name: "value" } }
          }
        }
      },
    }

    let template = new TemplateController({...host, app: testApp})
    await template.deleteTemplate()
    sinon.assert.callCount(testApp.currentModule, 1);

    testApp = {
      ...testApp,
      resultError: sinon.spy(),
      requestData: sinon.spy(() => ({ error: {} })),
    }
    template = new TemplateController({...host, app: testApp})
    await template.deleteTemplate()
    sinon.assert.callCount(testApp.resultError, 1);

  })

  it('editDataItem', () => {
    let testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...app.store.data[APP_MODULE.TEMPLATE],
            key: "_sample",
            current_data: {
              name: "labels",
              type: "list",
              items: getDataList(sampleTemplate().data.labels),
              item: "title"
            }
          }
        }
      },
    }

    let template = new TemplateController({...host, app: testApp})
    template.editDataItem({ value: "value" })
    sinon.assert.callCount(testApp.store.setData, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        setData: sinon.spy(),
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            key: "template",
            current_data: {
              name: "html_text",
              type: "string",
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.editDataItem({ value: "value" })
    sinon.assert.callCount(testApp.store.setData, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        setData: sinon.spy(),
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current_data: {
              name: "items",
              type: "table",
              items: getDataTable(sampleTemplate().data.items),
              item: {
                text: "Lorem ipsum dolor1", number: "3", date: "2014.01.08", _index: 0,
              }
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.editDataItem({ value: "value", field: "text", _index: 0, })
    sinon.assert.callCount(testApp.store.setData, 1);

  })

  it('editItem', () => {

    let testApp = {
      ...app,
      store: {
        ...app.store,
        data: {
          ...app.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...app.store.data[APP_MODULE.TEMPLATE],
            current: {
              id: "tmp_report",
              section: "report",
              type: "report",
              item: sampleTemplate().report,
              index: null,
              parent: null,
              parent_type: null,
              parent_index: null,
            }
          }
        }
      },
    }

    let template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.editItem({ 
      id: 1, name: "title", event_type: "change", value: "value", extend: false 
    })
    sinon.assert.callCount(template.setCurrent, 1);

    template.editItem({ 
      selected: true, datatype: "string", defvalue: "title", name: "title", value: false, extend: false 
    })
    sinon.assert.callCount(template.setCurrent, 2);

    template.editItem({ 
      selected: true, datatype: "string", defvalue: "title", name: "title", value: true, extend: false 
    })
    sinon.assert.callCount(template.setCurrent, 3);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              id: "tmp_details_0_vgap",
              section: "details",
              type: "vgap",
              index: 0,
              parent: sampleTemplate().details,
              parent_type: "details",
              parent_index: null,
              item: sampleTemplate().details[0].vgap,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.editItem({ 
      id: 1, name: "height", event_type: "change", value: 4, extend: false 
    })
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              id: "tmp_details_0_vgap",
              section: "details",
              type: "vgap",
              index: 0,
              parent: sampleTemplate().details,
              parent_type: "details",
              parent_index: null,
              item: sampleTemplate().details[0].vgap,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.editItem({ 
      selected: true, datatype: "float", defvalue: 0, name: "height", value: false, extend: false 
    })
    sinon.assert.callCount(template.setCurrent, 1);

    template.editItem({ 
      selected: true, datatype: "float", defvalue: 0, name: "height", value: true, extend: false 
    })
    sinon.assert.callCount(template.setCurrent, 2);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              id: "tmp_details_1_row",
              section: "details",
              type: "row",
              index: 1,
              parent: sampleTemplate().details,
              parent_type: "details",
              parent_index: null,
              item: sampleTemplate().details[1].row.columns,
              item_base: sampleTemplate().details[1].row
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.editItem({ 
      selected: true, datatype: "string", name: "visible", value: true, extend: false 
    })
    sinon.assert.callCount(template.setCurrent, 1);

    template.editItem({ 
      id: 1, name: "visible", event_type: "change", value: "value", extend: false 
    })
    sinon.assert.callCount(template.setCurrent, 2);

    template.editItem({ 
      selected: true, datatype: "string", name: "visible", value: false, extend: false 
    })
    sinon.assert.callCount(template.setCurrent, 3);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              id: "tmp_details_1_row_0_cell",
              section: "details",
              type: "cell",
              item: sampleTemplate().details[1].row.columns[0].cell,
              index: 0,
              parent: sampleTemplate().details[1].row.columns[0],
              parent_type: "row",
              parent_index: 1,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.editItem({ 
      id: 1, name: "font-style", event_type: "change", value: "italic", extend: false 
    })
    sinon.assert.callCount(template.setCurrent, 1);

    template.editItem({ 
      selected: true, datatype: "select", defvalue: "", name: "font-style", value: false, extend: false 
    })
    sinon.assert.callCount(template.setCurrent, 2);

    template.editItem({ 
      selected: true, datatype: "select", defvalue: "", name: "font-style", value: true, extend: false 
    })
    sinon.assert.callCount(template.setCurrent, 3);

    template.editItem({ 
      checklist: true, name: "border", checked: false, value: "L", extend: false, 
    })
    sinon.assert.callCount(template.setCurrent, 4);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              ...testApp.store.data[APP_MODULE.TEMPLATE].current,
              item: {
                ...testApp.store.data[APP_MODULE.TEMPLATE].current.item,
                border: null
              }
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.editItem({ 
      checklist: true, name: "border", checked: true, value: "L", extend: false, 
    })
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            key: "_template",
            current: {
              id: "tmp_details_19_datagrid",
              section: "details",
              type: "datagrid",
              index: 19,
              parent: sampleTemplate().details,
              parent_type: "details",
              parent_index: null,
              item: sampleTemplate().details[19].datagrid.columns,
              item_base: sampleTemplate().details[19].datagrid
            },
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.editItem({ 
      checklist: true, name: "border", checked: true, value: "T", extend: false 
    })
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              ...testApp.store.data[APP_MODULE.TEMPLATE].current,
              item_base: {
                ...testApp.store.data[APP_MODULE.TEMPLATE].current.item_base,
                border: "T"
              }
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.editItem({ 
      checklist: true, name: "border", checked: true, value: "R", extend: false 
    })
    sinon.assert.callCount(template.setCurrent, 1);

    template.editItem({ 
      checklist: true, name: "border", checked: false, value: "T", extend: false 
    })
    sinon.assert.callCount(template.setCurrent, 2);

    template.editItem({ 
      checklist: true, name: "border", checked: true, value: "1", extend: false 
    })
    sinon.assert.callCount(template.setCurrent, 3);

    template.editItem({ 
      selected: true, datatype: "checklist", name: "border", value: false, extend: false 
    })
    sinon.assert.callCount(template.setCurrent, 4);

    template.editItem({ 
      selected: true, datatype: "checklist", name: "border", value: true, extend: false 
    })
    sinon.assert.callCount(template.setCurrent, 5);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            key: "_template",
            current: {
              id: "tmp_header_0_row_0_image",
              section: "header",
              type: "image",
              item: sampleTemplate().header[0].row.columns[0].image,
              index: 0,
              parent: sampleTemplate().header[0].row.columns[0],
              parent_type: "row",
              parent_index: 0,
            },
          }
        }
      },
    }

    const stub = sinon.stub(FileReader.prototype, 'readAsDataURL')

    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    
    template.editItem({ 
      file: true, name: "src", value: [{}], extend: false 
    })
    sinon.assert.callCount(template.setCurrent, 0);

    template.editItem({ 
      file: true, name: "src", value: [], extend: false 
    })
    sinon.assert.callCount(template.setCurrent, 0);

    stub.restore();

  })

  it('exportTemplate', async () => {
    const testApp = {
      ...app,
      saveToDisk: sinon.spy(),
      store: {
        ...app.store,
        setData: sinon.spy()
      },
    }

    const template = new TemplateController({...host, app: testApp})
    template.exportTemplate()
    sinon.assert.callCount(testApp.saveToDisk, 1);

  })

  it('goNext', async () => {
    let testApp = {
      ...app,
      store: {
        ...app.store,
        data: {
          ...app.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...app.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "report",
              index: null,
              parent_index: null,
            }
          }
        }
      },
    }

    let template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.goNext()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "header",
              index: null,
              parent_index: null,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.goNext()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "header",
              index: 0,
              parent_index: null,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.goNext()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "header",
              index: 0,
              parent_index: 0,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.goNext()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "header",
              index: 2,
              parent_index: 0,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.goNext()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "footer",
              index: null,
              parent_index: null,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.goNext()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            template: {
              meta: {},
              report: {},
              header: [],
              details: [],
              footer: [
                { row: { columns: [] } }
              ],
              sources: {},
              data: {}
            },
            current: {
              section: "footer",
              index: null,
              parent_index: 0,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.goNext()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            template: {
              meta: {},
              report: {},
              header: [],
              details: [],
              footer: [
                { row: { columns: [ { cell: {} } ] } }
              ],
              sources: {},
              data: {}
            },
            current: {
              section: "footer",
              index: 0,
              parent_index: 0,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.goNext()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            template: {
              meta: {},
              report: {},
              header: [],
              details: [
                { row: { columns: [ { cell: {} } ] } }
              ],
              footer: [],
              sources: {},
              data: {}
            },
            current: {
              section: "details",
              index: 0,
              parent_index: 0,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.goNext()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            template: {
              meta: {},
              report: {},
              header: [],
              details: [
                { vgap: { } }
              ],
              footer: [],
              sources: {},
              data: {}
            },
            current: {
              section: "details",
              index: null,
              parent_index: 0,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.goNext()
    sinon.assert.callCount(template.setCurrent, 1);

  })

  it('goPrevious', async () => {
    let testApp = {
      ...app,
      store: {
        ...app.store,
        data: {
          ...app.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...app.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "report",
              index: null,
              parent_index: null,
            }
          }
        }
      },
    }

    let template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.goPrevious()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "header",
              index: null,
              parent_index: null,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.goPrevious()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "header",
              index: 0,
              parent_index: null,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.goPrevious()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "header",
              index: 0,
              parent_index: 0,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.goPrevious()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "header",
              index: 1,
              parent_index: 0,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.goPrevious()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "header",
              index: 1,
              parent_index: null,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.goPrevious()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "header",
              index: 2,
              parent_index: null,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.goPrevious()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "details",
              index: null,
              parent_index: null,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.goPrevious()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            template: {
              meta: {},
              report: {},
              header: [],
              details: [],
              footer: [],
              sources: {},
              data: {}
            },
            current: {
              section: "details",
              index: null,
              parent_index: null,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.goPrevious()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            template: {
              meta: {},
              report: {},
              header: [],
              details: [
                { row: { columns: [] } }
              ],
              footer: [],
              sources: {},
              data: {}
            },
            current: {
              section: "details",
              index: null,
              parent_index: 1,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.goPrevious()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            template: {
              meta: {},
              report: {},
              header: [
                { row: { columns: [] } }
              ],
              details: [],
              footer: [],
              sources: {},
              data: {}
            },
            current: {
              section: "details",
              index: null,
              parent_index: null,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.goPrevious()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            template: {
              meta: {},
              report: {},
              header: [
                { row: { columns: [ { cell: {} } ] } }
              ],
              details: [],
              footer: [],
              sources: {},
              data: {}
            },
            current: {
              section: "details",
              index: null,
              parent_index: null,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.goPrevious()
    sinon.assert.callCount(template.setCurrent, 1);

  })

  it('moveDown', async () => {
    let testApp = {
      ...app,
      store: {
        ...app.store,
        data: {
          ...app.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...app.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "details",
              type: "vgap",
              item: sampleTemplate().details[0].vgap,
              index: 0,
              parent: sampleTemplate().details,
              parent_type: "details",
              parent_index: null,
            }
          }
        }
      },
    }

    let template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.moveDown()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "details",
              type: "cell",
              item: sampleTemplate().details[1].row.columns[0].cell,
              index: 0,
              parent: sampleTemplate().details[1].row.columns,
              parent_type: "row",
              parent_index: 1,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.moveDown()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "details",
              type: "cell",
              item: sampleTemplate().details[1].row.columns[1].cell,
              index: 1,
              parent: sampleTemplate().details[1].row.columns,
              parent_type: "row",
              parent_index: 1,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.moveDown()
    sinon.assert.callCount(template.setCurrent, 0);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "header",
              type: "row",
              item: sampleTemplate().header[0].row,
              index: null,
              parent: sampleTemplate().header,
              parent_type: "header",
              parent_index: null,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.moveDown()
    sinon.assert.callCount(template.setCurrent, 0);

  })

  it('moveUp', async () => {
    let testApp = {
      ...app,
      store: {
        ...app.store,
        data: {
          ...app.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...app.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "details",
              type: "vgap",
              item: sampleTemplate().details[8].vgap,
              index: 8,
              parent: sampleTemplate().details,
              parent_type: "details",
              parent_index: null,
            }
          }
        }
      },
    }

    let template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.moveUp()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "details",
              type: "cell",
              item: sampleTemplate().details[5].row.columns[1].cell,
              index: 1,
              parent: sampleTemplate().details[5].row.columns,
              parent_type: "row",
              parent_index: 5,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.moveUp()
    sinon.assert.callCount(template.setCurrent, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "details",
              type: "vgap",
              item: sampleTemplate().details[0].vgap,
              index: 0,
              parent: sampleTemplate().details,
              parent_type: "details",
              parent_index: null,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.moveUp()
    sinon.assert.callCount(template.setCurrent, 0);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            current: {
              section: "header",
              type: "row",
              item: sampleTemplate().header[0].row,
              index: null,
              parent: sampleTemplate().header,
              parent_type: "header",
              parent_index: null,
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    template.moveUp()
    sinon.assert.callCount(template.setCurrent, 0);

  })

  it('saveTemplate', async () => {
    let testApp = {
      ...app,
      resultError: sinon.spy(),
      requestData: sinon.spy(() => ({})),
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
          [APP_MODULE.TEMPLATE]: {
            ...app.store.data[APP_MODULE.TEMPLATE],
            dbtemp: { id: 1, code: "code", data: { report_name: "value" } }
          }
        }
      },
    }

    let template = new TemplateController({...host, app: testApp})
    await template.saveTemplate(true)
    sinon.assert.callCount(testApp.store.setData, 3);

    testApp = {
      ...testApp,
      resultError: sinon.spy(),
      requestData: sinon.spy(() => ({ error: {} })),
    }
    template = new TemplateController({...host, app: testApp})
    await template.saveTemplate(true)
    sinon.assert.callCount(testApp.resultError, 1);

  })

  it('setCurrent', () => {
    let testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...app.store.data[APP_MODULE.TEMPLATE],
            key: "test"
          }
        }
      },
    }

    let template = new TemplateController({...host, app: testApp})
    template.setCurrent({tmp_id: "tmp_details", set_dirty: true})
    sinon.assert.callCount(testApp.store.setData, 1);

    template.setCurrent({tmp_id: "tmp_details_1_row"})
    sinon.assert.callCount(testApp.store.setData, 2);

    template.setCurrent({tmp_id: "tmp_details_0_vgap"})
    sinon.assert.callCount(testApp.store.setData, 3);

    template.setCurrent({tmp_id: "tmp_details_1_row_0_cell"})
    sinon.assert.callCount(testApp.store.setData, 4);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            mapRef: {}
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.createMap = sinon.spy()
    template.setCurrent({tmp_id: "tmp_details_1_row_0_cell"})
    sinon.assert.callCount(template.createMap, 1);

  })

  it('setCurrentData', () => {
    let testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...app.store.data[APP_MODULE.TEMPLATE],
            key: "_sample"
          }
        }
      },
    }

    let template = new TemplateController({...host, app: testApp})
    template.setCurrentData(
      { name: "labels", type: "list" }
    )
    sinon.assert.callCount(testApp.store.setData, 1);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            key: "template"
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrentData(
      { name: "html_text", type: "string" }
    )
    sinon.assert.callCount(testApp.store.setData, 2);

    template.setCurrentData(
      { name: "items", type: "table" }
    )
    sinon.assert.callCount(testApp.store.setData, 3);

    template.setCurrentData(
      { name: "new", type: "new",
        values: { name: "new_text", type: "string", columns: "" } 
      }
    )
    sinon.assert.callCount(testApp.store.setData, 4);

    template.setCurrentData(
      { name: "new", type: "new",
        values: { name: "new_list", type: "list", columns: "" } 
      }
    )
    sinon.assert.callCount(testApp.store.setData, 5);

    template.setCurrentData(
      { name: "new", type: "new",
        values: { name: "new_table", type: "table", columns: "col1,col2" } 
      }
    )
    sinon.assert.callCount(testApp.store.setData, 6);

    template.setCurrentData(null)
    sinon.assert.callCount(testApp.store.setData, 7);

    template.setCurrentData(
      { name: "new", type: "new",
        values: { name: "", type: "table", columns: "col1,col2" } 
      }
    )
    sinon.assert.callCount(testApp.store.setData, 7);

    template.setCurrentData(
      { name: "new", type: "new",
        values: { name: "new_table", type: "table", columns: "" } 
      }
    )
    sinon.assert.callCount(testApp.store.setData, 7);

  })

  it('setCurrentDataItem', () => {
    let testApp = {
      ...app,
      resultError: sinon.spy(),
      requestData: sinon.spy(() => ({})),
      store: {
        ...app.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: { value: "" } })
          }
        }),
        data: {
          ...app.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...app.store.data[APP_MODULE.TEMPLATE],
            key: "_sample",
            current_data: {
              name: "labels",
              type: "list",
              items: getDataList(sampleTemplate().data.labels)
            }
          }
        }
      },
    }

    let template = new TemplateController({...host, app: testApp})
    template.setCurrentDataItem("title")
    sinon.assert.callCount(testApp.store.setData, 1);

    template.setCurrentDataItem()
    sinon.assert.callCount(testApp.store.setData, 4);

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: { value: "value" } })
          }
        })
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrentDataItem()
    sinon.assert.callCount(testApp.store.setData, 3)

    testApp = {
      ...testApp,
      showToast: sinon.spy(),
      store: {
        ...testApp.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: { value: "title" } })
          }
        })
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrentDataItem()
    sinon.assert.callCount(testApp.showToast, 1)

    testApp = {
      ...testApp,
      store: {
        ...testApp.store,
        data: {
          ...testApp.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...testApp.store.data[APP_MODULE.TEMPLATE],
            key: "template",
            current_data: {
              name: "items",
              type: "table",
              items: getDataTable(sampleTemplate().data.items)
            }
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    template.setCurrentDataItem()
    sinon.assert.callCount(testApp.store.setData, 3)

  })

  it('setTemplate', async () => {
    const testApp = {
      ...app,
      resultError: sinon.spy(),
      currentModule: sinon.spy()
    }

    let options = {
      ...app.store.data[APP_MODULE.TEMPLATE],
      type: "template",
      report: {
        data: { report_name: "value", template: JSON.stringify(sample) }
      },
      id: 1,
      code: "code"
    }
    const template = new TemplateController({...host, app: testApp})
    template.setCurrent = sinon.spy()
    await template.setTemplate(options)
    sinon.assert.callCount(template.setCurrent, 1);

    options = {
      ...app.store.data[APP_MODULE.TEMPLATE],
      type: "template",
      dataset: {
        template: [{
          id: "id", report_key: "report_key", report_name: "report_name",
          report: ""
        }]
      }
    }
    await template.setTemplate(options)
    sinon.assert.callCount(testApp.currentModule, 1);

    options = {
      ...app.store.data[APP_MODULE.TEMPLATE],
      type: "_sample"
    }
    await template.setTemplate(options)
    sinon.assert.callCount(template.setCurrent, 2);

  })

  it('showPreview', async () => {
    let testApp = {
      ...app,
      resultError: sinon.spy(),
      requestData: sinon.spy(() => ({})),
      store: {
        ...app.store,
        setData: sinon.spy((key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK, data: { value: "test" } })
          }
        }),
        data: {
          ...app.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...app.store.data[APP_MODULE.TEMPLATE],
            key: "_sample"
          }
        }
      },
    }

    let template = new TemplateController({...host, app: testApp})
    await template.showPreview("portrait")
    sinon.assert.callCount(testApp.requestData, 1);

    testApp = {
      ...testApp,
      resultError: sinon.spy(),
      requestData: sinon.spy(() => ({ error: {} })),
      store: {
        ...testApp.store,
        data: {
          ...app.store.data,
          [APP_MODULE.TEMPLATE]: {
            ...app.store.data[APP_MODULE.TEMPLATE],
            key: "template"
          }
        }
      },
    }
    template = new TemplateController({...host, app: testApp})
    await template.showPreview()
    sinon.assert.callCount(testApp.resultError, 1);

  })

  it('setModule', () => {
    const template = new TemplateController({...host, app})
    template.setModule({})
    expect(template.module).to.exist

    getElementType({})
  })

})