import { queryByAttribute } from '@testing-library/react'
import ReactDOM from 'react-dom';
import update from 'immutability-helper';
import { templateActions } from './actions'
import { templateElements } from './Template'

import { appActions, saveToDisk } from 'containers/App/actions'
import { getText as appGetText, store as app_store  } from 'config/app'

jest.mock("containers/App/actions");

const getById = queryByAttribute.bind(null, 'id');
const sample_template = require('../../config/sample.json')
const getText = (key)=>appGetText({ locales: app_store.session.locales, lang: "en", key: key })

const store = update(app_store, {$merge: {
  login: {
    data: {
      employee: {
        id: 1, empnumber: 'admin', username: 'admin', usergroup: 114,
        usergroupName: 'admin'
      },
    }
  },
  template: {
    key: "_sample",
    template: {
      ...update(sample_template, {}),
      sources: {
        head: {
          default: "select * from table"
        }
      },
      footer: []
    },
    current: {
      id: "tmp_report",
      section: "report",
      type: "report",
      item: update(sample_template, {}).report,
      index: null,
      parent: null,
      parent_type: null,
      parent_index: null,
      form: templateElements({ getText: getText })["report"]
    }
  }
}})

describe('templateActions', () => {
  
  beforeEach(() => {
    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ([1])),
      resultError: jest.fn(),
      showToast: jest.fn()
    })
    saveToDisk.mockReturnValue()
  });

  afterEach(() => {
    jest.clearAllMocks();
  });

  it('getElementType', () => {
    const setData = jest.fn()
    let values = templateActions(store, setData).getElementType({ row: {} })
    expect(values).toBe("row")
    values = templateActions(store, setData).getElementType({})
    expect(values).toBeNull()
  })

  it('getDataset', () => {
    const sample = update(sample_template, {})
    const setData = jest.fn()
    const dataset = templateActions(store, setData).getDataset(sample.data)
    expect(dataset.length).toBe(6)
  })

  it('setCurrent', () => {
    const setData = jest.fn()
    let it_store = update(store, {template: {$merge:{
      key: "test"
    }}})
    templateActions(it_store, setData).setCurrent({tmp_id: "tmp_details", set_dirty: true})
    expect(setData).toHaveBeenCalledTimes(1)
    templateActions(it_store, setData).setCurrent({tmp_id: "tmp_details_1_row"})
    expect(setData).toHaveBeenCalledTimes(2)
    templateActions(it_store, setData).setCurrent({tmp_id: "tmp_details_0_vgap"})
    expect(setData).toHaveBeenCalledTimes(3)
    templateActions(it_store, setData).setCurrent({tmp_id: "tmp_details_1_row_0_cell"})
    expect(setData).toHaveBeenCalledTimes(4)
  })

  it('createMap', () => {
    const sample = update(sample_template, {})
    const setData = jest.fn()
    let canvas = {
      height: 0, width: 0,
      getContext: () => ({
        fillStyle: "",
        fillRect: jest.fn(),
        beginPath: jest.fn(),
        moveTo: jest.fn(),
        lineTo: jest.fn(),
        stroke: jest.fn(),
      })
    }
    let it_store = update(store, {template: {$merge:{
      current: {
        item: sample.details,
        parent: null,
      }
    }}})
    templateActions(it_store, setData).createMap(canvas)
    it_store = update(it_store, {template: {$merge:{
      current: {
        item: sample.details[1].row.columns[0].cell,
        parent: sample.details[1].row.columns,
      }
    }}})
    templateActions(it_store, setData).createMap(canvas)
    it_store = update(it_store, {template: {$merge:{
      current: {
        item: sample.details[19].datagrid.columns[1].column,
        parent: sample.details[19].datagrid.columns,
      }
    }}})
    templateActions(it_store, setData).createMap(canvas)
    it_store = update(it_store, {template: {$merge:{
      template: {
        meta: {},
        report: {},
        header: [],
        details: [],
        footer: [],
        sources: {},
        data: {}
      },
      current: {}
    }}})
    templateActions(it_store, setData).createMap(canvas)
  })

  it('goPrevious', () => {
    const setData = jest.fn()
    let it_store = update(store, {template: {$merge:{
      current: {
        section: "report",
        index: null,
        parent_index: null,
      }
    }}})
    templateActions(it_store, setData).goPrevious()
    expect(setData).toHaveBeenCalledTimes(1)

    it_store = update(it_store, {template: {$merge:{
      current: {
        section: "header",
        index: null,
        parent_index: null,
      }
    }}})
    templateActions(it_store, setData).goPrevious()
    expect(setData).toHaveBeenCalledTimes(2)

    it_store = update(it_store, {template: {$merge:{
      current: {
        section: "header",
        index: 0,
        parent_index: null,
      }
    }}})
    templateActions(it_store, setData).goPrevious()
    expect(setData).toHaveBeenCalledTimes(3)

    it_store = update(it_store, {template: {$merge:{
      current: {
        section: "header",
        index: 0,
        parent_index: 0,
      }
    }}})
    templateActions(it_store, setData).goPrevious()
    expect(setData).toHaveBeenCalledTimes(4)

    it_store = update(it_store, {template: {$merge:{
      current: {
        section: "header",
        index: 1,
        parent_index: 0,
      }
    }}})
    templateActions(it_store, setData).goPrevious()
    expect(setData).toHaveBeenCalledTimes(5)

    it_store = update(it_store, {template: {$merge:{
      current: {
        section: "header",
        index: 1,
        parent_index: null,
      }
    }}})
    templateActions(it_store, setData).goPrevious()
    expect(setData).toHaveBeenCalledTimes(6)

    it_store = update(it_store, {template: {$merge:{
      current: {
        section: "header",
        index: 2,
        parent_index: null,
      }
    }}})
    templateActions(it_store, setData).goPrevious()
    expect(setData).toHaveBeenCalledTimes(7)

    it_store = update(it_store, {template: {$merge:{
      current: {
        section: "details",
        index: null,
        parent_index: null,
      }
    }}})
    templateActions(it_store, setData).goPrevious()
    expect(setData).toHaveBeenCalledTimes(8)

    it_store = update(it_store, {template: {$merge:{
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
    }}})
    templateActions(it_store, setData).goPrevious()
    expect(setData).toHaveBeenCalledTimes(9)

    it_store = update(it_store, {template: {$merge:{
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
    }}})
    templateActions(it_store, setData).goPrevious()
    expect(setData).toHaveBeenCalledTimes(10)

    it_store = update(it_store, {template: {$merge:{
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
    }}})
    templateActions(it_store, setData).goPrevious()
    expect(setData).toHaveBeenCalledTimes(11)

    it_store = update(it_store, {template: {$merge:{
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
    }}})
    templateActions(it_store, setData).goPrevious()
    expect(setData).toHaveBeenCalledTimes(12)
  })

  it('goNext', () => {
    const setData = jest.fn()
    let it_store = update(store, {template: {$merge:{
      current: {
        section: "report",
        index: null,
        parent_index: null,
      }
    }}})
    templateActions(it_store, setData).goNext()
    expect(setData).toHaveBeenCalledTimes(1)

    it_store = update(it_store, {template: {$merge:{
      current: {
        section: "header",
        index: null,
        parent_index: null,
      }
    }}})
    templateActions(it_store, setData).goNext()
    expect(setData).toHaveBeenCalledTimes(2)

    it_store = update(it_store, {template: {$merge:{
      current: {
        section: "header",
        index: 0,
        parent_index: null,
      }
    }}})
    templateActions(it_store, setData).goNext()
    expect(setData).toHaveBeenCalledTimes(3)

    it_store = update(it_store, {template: {$merge:{
      current: {
        section: "header",
        index: 0,
        parent_index: 0,
      }
    }}})
    templateActions(it_store, setData).goNext()
    expect(setData).toHaveBeenCalledTimes(4)

    it_store = update(it_store, {template: {$merge:{
      current: {
        section: "header",
        index: 2,
        parent_index: 0,
      }
    }}})
    templateActions(it_store, setData).goNext()
    expect(setData).toHaveBeenCalledTimes(5)

    it_store = update(it_store, {template: {$merge:{
      current: {
        section: "footer",
        index: null,
        parent_index: null,
      }
    }}})
    templateActions(it_store, setData).goNext()
    expect(setData).toHaveBeenCalledTimes(6)

    it_store = update(it_store, {template: {$merge:{
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
    }}})
    templateActions(it_store, setData).goNext()
    expect(setData).toHaveBeenCalledTimes(7)

    it_store = update(it_store, {template: {$merge:{
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
    }}})
    templateActions(it_store, setData).goNext()
    expect(setData).toHaveBeenCalledTimes(8)

    it_store = update(it_store, {template: {$merge:{
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
    }}})
    templateActions(it_store, setData).goNext()
    expect(setData).toHaveBeenCalledTimes(9)

    it_store = update(it_store, {template: {$merge:{
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
    }}})
    templateActions(it_store, setData).goNext()
    expect(setData).toHaveBeenCalledTimes(10)

  })

  it('moveDown', () => {
    const sample = update(sample_template, {})
    const setData = jest.fn()
    let it_store = update(store, {template: {$merge:{
      current: {
        section: "details",
        type: "vgap",
        item: sample.details[0].vgap,
        index: 0,
        parent: sample.details,
        parent_type: "details",
        parent_index: null,
      },
    }}})
    templateActions(it_store, setData).moveDown()
    expect(setData).toHaveBeenCalledTimes(1)

    it_store = update(store, {template: {$merge:{
      current: {
        section: "details",
        type: "cell",
        item: sample.details[0].row.columns[0].cell,
        index: 0,
        parent: sample.details[0].row.columns,
        parent_type: "row",
        parent_index: 0,
      },
    }}})
    templateActions(it_store, setData).moveDown()
    expect(setData).toHaveBeenCalledTimes(2)

    it_store = update(store, {template: {$merge:{
      current: {
        section: "details",
        type: "cell",
        item: sample.details[0].row.columns[1].cell,
        index: 1,
        parent: sample.details[0].row.columns,
        parent_type: "row",
        parent_index: 0,
      },
    }}})
    templateActions(it_store, setData).moveDown()
    expect(setData).toHaveBeenCalledTimes(2)

    it_store = update(store, {template: {$merge:{
      current: {
        section: "header",
        type: "row",
        item: sample.header[0].row,
        index: null,
        parent: sample.header,
        parent_type: "header",
        parent_index: null,
      },
    }}})
    templateActions(it_store, setData).moveDown()
    expect(setData).toHaveBeenCalledTimes(2)

  })

  it('moveUp', () => {
    const sample = update(sample_template, {})
    const setData = jest.fn()
    let it_store = update(store, {template: {$merge:{
      current: {
        section: "details",
        type: "vgap",
        item: sample.details[8].vgap,
        index: 8,
        parent: sample.details,
        parent_type: "details",
        parent_index: null,
      },
    }}})
    templateActions(it_store, setData).moveUp()
    expect(setData).toHaveBeenCalledTimes(1)

    it_store = update(store, {template: {$merge:{
      current: {
        section: "details",
        type: "cell",
        item: sample.details[5].row.columns[1].cell,
        index: 1,
        parent: sample.details[5].row.columns,
        parent_type: "row",
        parent_index: 5,
      },
    }}})
    templateActions(it_store, setData).moveUp()
    expect(setData).toHaveBeenCalledTimes(2)

    it_store = update(store, {template: {$merge:{
      current: {
        section: "details",
        type: "vgap",
        item: sample.details[0].vgap,
        index: 0,
        parent: sample.details,
        parent_type: "details",
        parent_index: null,
      },
    }}})
    templateActions(it_store, setData).moveUp()
    expect(setData).toHaveBeenCalledTimes(2)

    it_store = update(store, {template: {$merge:{
      current: {
        section: "header",
        type: "row",
        item: sample.header[0].row,
        index: null,
        parent: sample.header,
        parent_type: "header",
        parent_index: null,
      },
    }}})
    templateActions(it_store, setData).moveUp()
    expect(setData).toHaveBeenCalledTimes(2)

  })

  it('deleteItem', () => {
    const sample = update(sample_template, {})
    const setData = jest.fn()
    let it_store = update(store, {template: {$merge:{
      current: {
        section: "details",
        type: "vgap",
        item: sample.details[0].vgap,
        index: 0,
        parent: sample.details,
        parent_type: "details",
        parent_index: null,
      },
    }}})
    templateActions(it_store, setData).deleteItem()
    expect(setData).toHaveBeenCalledTimes(1)

    it_store = update(store, {template: {$merge:{
      current: {
        section: "details",
        type: "cell",
        item: sample.details[5].row.columns[1].cell,
        index: 1,
        parent: sample.details[5].row.columns,
        parent_type: "row",
        parent_index: 5,
      },
    }}})
    templateActions(it_store, setData).deleteItem()
    expect(setData).toHaveBeenCalledTimes(2)

    it_store = update(store, {template: {$merge:{
      current: {
        section: "header",
        type: "row",
        item: sample.header[0].row,
        index: null,
        parent: sample.header,
        parent_type: "header",
        parent_index: null,
      },
    }}})
    templateActions(it_store, setData).deleteItem()
    expect(setData).toHaveBeenCalledTimes(2)

  })

  it('addItem', () => {
    const sample = update(sample_template, {})
    const setData = jest.fn()
    let it_store = update(store, {template: {$merge:{
      current: {
        section: "details",
        type: "details",
        item: sample.details,
        index: null,
        parent: null,
        parent_type: null,
        parent_index: null,
      },
    }}})
    templateActions(it_store, setData).addItem("row")
    expect(setData).toHaveBeenCalledTimes(1)

    it_store = update(it_store, {template: {$merge:{
      current: {
        section: "details",
        type: "details",
        item: sample.details,
        index: null,
        parent: null,
        parent_type: null,
        parent_index: null,
      },
    }}})
    templateActions(it_store, setData).addItem("vgap")
    expect(setData).toHaveBeenCalledTimes(2)

    templateActions(it_store, setData).addItem("")
    expect(setData).toHaveBeenCalledTimes(2)

  })

  it('editItem', () => {
    const sample = update(sample_template, {})
    const setData = jest.fn()
    let it_store = update(store, {template: {$merge:{
      current: {
        id: "tmp_report",
        section: "report",
        type: "report",
        item: sample.report,
        index: null,
        parent: null,
        parent_type: null,
        parent_index: null,
      },
    }}})
    templateActions(it_store, setData).editItem({ 
      id: 1, name: "title", event_type: "change", value: "value", extend: false 
    })
    expect(setData).toHaveBeenCalledTimes(1)
    templateActions(it_store, setData).editItem({ 
      selected: true, datatype: "string", defvalue: "title", name: "title", value: false, extend: false 
    })
    expect(setData).toHaveBeenCalledTimes(2)
    templateActions(it_store, setData).editItem({ 
      selected: true, datatype: "string", defvalue: "title", name: "title", value: true, extend: false 
    })
    expect(setData).toHaveBeenCalledTimes(3)

    it_store = update(store, {template: {$merge:{
      current: {
        id: "tmp_details_0_vgap",
        section: "details",
        type: "vgap",
        index: 0,
        parent: sample.details,
        parent_type: "details",
        parent_index: null,
        item: sample.details[0].vgap,
      },
    }}})
    templateActions(it_store, setData).editItem({ 
      id: 1, name: "height", event_type: "change", value: 4, extend: false 
    })
    expect(setData).toHaveBeenCalledTimes(4)
    templateActions(it_store, setData).editItem({ 
      selected: true, datatype: "float", defvalue: 0, name: "height", value: false, extend: false 
    })
    expect(setData).toHaveBeenCalledTimes(5)
    templateActions(it_store, setData).editItem({ 
      selected: true, datatype: "float", defvalue: 0, name: "height", value: true, extend: false 
    })
    expect(setData).toHaveBeenCalledTimes(6)

    it_store = update(store, {template: {$merge:{
      current: {
        id: "tmp_details_1_row",
        section: "details",
        type: "row",
        index: 1,
        parent: sample.details,
        parent_type: "details",
        parent_index: null,
        item: sample.details[1].row.columns,
        item_base: sample.details[1].row
      },
    }}})
    templateActions(it_store, setData).editItem({ 
      selected: true, datatype: "string", name: "visible", value: true, extend: false 
    })
    expect(setData).toHaveBeenCalledTimes(7)
    templateActions(it_store, setData).editItem({ 
      id: 1, name: "visible", event_type: "change", value: "value", extend: false 
    })
    expect(setData).toHaveBeenCalledTimes(8)
    templateActions(it_store, setData).editItem({ 
      selected: true, datatype: "string", name: "visible", value: false, extend: false 
    })
    expect(setData).toHaveBeenCalledTimes(9)

    it_store = update(store, {template: {$merge:{
      current: {
        id: "tmp_details_1_row_0_cell",
        section: "details",
        type: "cell",
        item: sample.details[1].row.columns[0].cell,
        index: 0,
        parent: sample.details[1].row.columns[0],
        parent_type: "row",
        parent_index: 1,
      },
    }}})
    templateActions(it_store, setData).editItem({ 
      id: 1, name: "font-style", event_type: "change", value: "italic", extend: false 
    })
    expect(setData).toHaveBeenCalledTimes(10)
    templateActions(it_store, setData).editItem({ 
      selected: true, datatype: "select", defvalue: "", name: "font-style", value: false, extend: false 
    })
    expect(setData).toHaveBeenCalledTimes(11)
    templateActions(it_store, setData).editItem({ 
      selected: true, datatype: "select", defvalue: "", name: "font-style", value: true, extend: false 
    })
    expect(setData).toHaveBeenCalledTimes(12)
    templateActions(it_store, setData).editItem({ 
      checklist: true, name: "border", checked: false, value: "L", extend: false, 
    })
    expect(setData).toHaveBeenCalledTimes(13)
    it_store = update(it_store, {template: {current: {item: {$merge: {
      border: null
    }}}}})
    templateActions(it_store, setData).editItem({ 
      checklist: true, name: "border", checked: true, value: "L", extend: false, 
    })
    expect(setData).toHaveBeenCalledTimes(14)

    it_store = update(store, {template: {$merge:{
      key: "_template",
      current: {
        id: "tmp_details_18_datagrid",
        section: "details",
        type: "datagrid",
        index: 18,
        parent: sample.details,
        parent_type: "details",
        parent_index: null,
        item: sample.details[18].datagrid.columns,
        item_base: sample.details[18].datagrid
      },
    }}})
    templateActions(it_store, setData).editItem({ 
      checklist: true, name: "border", checked: true, value: "T", extend: false 
    })
    expect(setData).toHaveBeenCalledTimes(15)
    it_store = update(it_store, {template: {current: {item_base: {$merge: {
      border: "T"
    }}}}})
    templateActions(it_store, setData).editItem({ 
      checklist: true, name: "border", checked: true, value: "R", extend: false 
    })
    expect(setData).toHaveBeenCalledTimes(16)
    templateActions(it_store, setData).editItem({ 
      checklist: true, name: "border", checked: false, value: "T", extend: false 
    })
    expect(setData).toHaveBeenCalledTimes(17)
    templateActions(it_store, setData).editItem({ 
      checklist: true, name: "border", checked: true, value: "1", extend: false 
    })
    expect(setData).toHaveBeenCalledTimes(18)
    templateActions(it_store, setData).editItem({ 
      selected: true, datatype: "checklist", name: "border", value: false, extend: false 
    })
    expect(setData).toHaveBeenCalledTimes(19)
    templateActions(it_store, setData).editItem({ 
      selected: true, datatype: "checklist", name: "border", value: true, extend: false 
    })
    expect(setData).toHaveBeenCalledTimes(20)
    
    jest.spyOn(global, 'FileReader').mockImplementation(function () {
      this.readAsDataURL = jest.fn();
    });
    
    it_store = update(store, {template: {$merge:{
      key: "_template",
      current: {
        id: "tmp_header_0_row_0_image",
        section: "header",
        type: "image",
        item: sample.header[0].row.columns[0].image,
        index: 0,
        parent: sample.header[0].row.columns[0],
        parent_type: "row",
        parent_index: 0,
      },
    }}})
    templateActions(it_store, setData).editItem({ 
      file: true, name: "src", value: [{}], extend: false 
    })
    expect(setData).toHaveBeenCalledTimes(20)
    templateActions(it_store, setData).editItem({ 
      file: true, name: "src", value: [], extend: false 
    })
    expect(setData).toHaveBeenCalledTimes(20)
  })

  it('exportTemplate', () => {
    global.URL.createObjectURL = jest.fn();
    const setData = jest.fn((key, data, callback)=>{
      if(callback){callback()}
    })
    templateActions(store, setData).exportTemplate()
    expect(setData).toHaveBeenCalledTimes(1)
  })

  it('setTemplate', () => {
    const sample = update(sample_template, {})
    const setData = jest.fn()
    let options = update(store.template, {$merge: {
      type: "template",
      dataset: {
        template: [{
          id: "id", reportkey: "reportkey", repname: "repname",
          report: JSON.stringify(sample)
        }]
      }
    }})
    templateActions(store, setData).setTemplate(options)
    expect(setData).toHaveBeenCalledTimes(2)

    templateActions(store, setData).setTemplate({
      type: "_sample"
    })
    expect(setData).toHaveBeenCalledTimes(4)
  })

  it('saveTemplate', () => {
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
    templateActions(it_store, setData).saveTemplate(true)
    expect(setData).toHaveBeenCalledTimes(3);

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ({ error: {} })),
      resultError: jest.fn(),
      showToast: jest.fn()
    })
    templateActions(it_store, setData).saveTemplate(false)
    expect(setData).toHaveBeenCalledTimes(3);

    templateActions(it_store, setData).saveTemplate(true)
    expect(setData).toHaveBeenCalledTimes(6);

  })

  it('deleteTemplate', () => {
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
    templateActions(it_store, setData).deleteTemplate()
    expect(setData).toHaveBeenCalledTimes(3);

    appActions.mockReturnValue({
      getText: jest.fn((key)=>(key)),
      requestData: jest.fn(async () => ({ error: {} })),
      resultError: jest.fn(),
      showToast: jest.fn()
    })
    templateActions(it_store, setData).deleteTemplate()
    expect(setData).toHaveBeenCalledTimes(6);

  })

  it('deleteData', () => {
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
    let it_store = update(store, {template: {$merge:{
      key: "test"
    }}})
    templateActions(it_store, setData).deleteData("dtkey")
    expect(setData).toHaveBeenCalledTimes(4);

    it_store = update(store, {template: {$merge:{
      key: "_sample"
    }}})
    templateActions(it_store, setData).deleteData("dtkey")
    expect(setData).toHaveBeenCalledTimes(8);
  })

})