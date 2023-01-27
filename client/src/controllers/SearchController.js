import { APP_MODULE, MODAL_EVENT, BROWSER_EVENT, SIDE_EVENT, SIDE_VISIBILITY, TOAST_TYPE, EDITOR_EVENT } from '../config/enums.js'

export const getFilterWhere = (filter) => {
  switch (filter.filtertype) {
    case "==N":
      if(filter.fieldtype === "string"){
        return ["and", [ [filter.sqlstr, "like", "''"], ["or", filter.sqlstr, "is null"]]]
      }
      return ["and", filter.sqlstr, "is null"]
    
    case "!==":
      if(filter.fieldtype === "string"){
        return ["and", [`lower(${filter.sqlstr})`, "not like", "{CCS}{JOKER}{SEP}lower(?){SEP}{JOKER}{CCE}"]]
      }
      return ["and", [filter.sqlstr, "<>", "?"]]
    
    case ">==":
      return ["and", [filter.sqlstr, ">=", "?"]]
    
    case "<==":
      return ["and", [filter.sqlstr, "<=", "?"]]

    case "===":
    default:
      if(filter.fieldtype === "string"){
        return ["and", [`lower(${filter.sqlstr})`, "like", "{CCS}{JOKER}{SEP}lower(?){SEP}{JOKER}{CCE}"]]
      }
      return ["and", [filter.sqlstr, "=", "?"]]
  }
}

export const defaultFilterValue = (fieldtype) => {
  if(fieldtype === "date"){
    return new Date().toISOString().split("T")[0]
  }
  if(["bool", "integer", "float"].includes(fieldtype)){
    return 0
  }
  return ""
}

export class SearchController {
  constructor(host) {
    this.host = host;
    this.app = host.app
    this.store = host.app.store
    this.module = {}
    
    this.addFilter = this.addFilter.bind(this)
    this.browserView = this.browserView.bind(this)
    this.deleteFilter = this.deleteFilter.bind(this)
    this.editCell = this.editCell.bind(this)
    this.editFilter = this.editFilter.bind(this)
    this.editRow = this.editRow.bind(this)
    this.exportResult = this.exportResult.bind(this)
    this.onBrowserEvent = this.onBrowserEvent.bind(this)
    this.onModalEvent = this.onModalEvent.bind(this)
    this.onSideEvent = this.onSideEvent.bind(this)
    this.quickSearch = this.quickSearch.bind(this)
    this.saveBookmark = this.saveBookmark.bind(this)
    this.setColumns = this.setColumns.bind(this)
    this.setModule = this.setModule.bind(this)
    this.showBrowser = this.showBrowser.bind(this)
    this.showServerCmd = this.showServerCmd.bind(this)
    this.showTotal = this.showTotal.bind(this)
    host.addController(this);
  }

  setModule(moduleRef){
    this.module = moduleRef
  }

  addFilter() {
    const { queries } = this.app.modules
    const { setData } = this.store
    const { vkey, view, filters } = this.store.data[APP_MODULE.SEARCH]
    const viewDef = queries[vkey]()[view]
    const frow = viewDef.fields[Object.keys(viewDef.fields)[0]]
    const _filters = {...filters,
      [view]: [...filters[view], {
        id: new Date().getTime().toString(),
        fieldtype: frow.fieldtype,
        fieldname: Object.keys(viewDef.fields)[0],
        sqlstr: frow.sqlstr,
        wheretype: frow.wheretype,
        filtertype: "===",
        value: defaultFilterValue(frow.fieldtype)
      }]
    }
    setData(APP_MODULE.SEARCH, { filters: _filters })
  }

  async browserView() {
    const { queries } = this.app.modules
    const { getDataFilter, getUserFilter, getSql, requestData, resultError } = this.app
    const { setData, data } = this.store
    const { vkey, view, filters } = this.store.data[APP_MODULE.SEARCH]
    const query = queries[vkey]()[view]
    let _sql = { ...query.sql }
    let params = []
    let _where = []
    filters[view].filter((filter)=>(filter.wheretype === "where")).forEach(filter => {
      _where.push(getFilterWhere(filter))
      if(filter.filtertype !== "==N"){
        params.push(filter.value)
      }
    });
    if(_where.length > 0){
      _sql = {
        ..._sql,
        where: [..._sql.where, ..._where]
      }
    }

    _where = []
    filters[view].filter((filter)=>(filter.wheretype === "having")).forEach(filter => {
      _where.push(getFilterWhere(filter))
      if(filter.filtertype !== "==N"){
        params.push(filter.value)
      }
    });
    if(_where.length > 0){
      _sql = {
        ..._sql,
        having: [..._sql.having, ..._where]
      }
    }

    _where = getDataFilter(vkey, [], view)
    if(_where.length > 0){
      _sql = {
        ..._sql,
        where: [..._sql.where, _where]
      }
    }

    const userFilter = getUserFilter(vkey)
    if(userFilter.where.length > 0){
      _sql = {
        ..._sql,
        where: [..._sql.where, userFilter.where]
      }
      params = params.concat(userFilter.params)
    }

    const views = [
      { key: "result",
        text: getSql(data[APP_MODULE.LOGIN].data.engine, _sql).sql,
        values: params 
      }
    ]
    const options = { method: "POST", data: views }
    const oview = await requestData("/view", options)
    if(oview.error){
      return resultError(oview)
    }
    setData(APP_MODULE.SEARCH, { result: oview.result, dropdown: "", page: 1 })
    return true
  }

  deleteFilter(index) {
    const { setData } = this.store
    const { view, filters } = this.store.data[APP_MODULE.SEARCH]
    const _filters = {
      ...filters,
      [view]: [...filters[view].slice(0,index),...filters[view].slice(index+1)]
    }
    setData(APP_MODULE.SEARCH, { filters: _filters })
  }

  editCell({ fieldname, value, row }) {
    const { currentModule } = this.app
    const params = value.split("/")
    let options = { 
      ntype: params[0], 
      ttype: params[1], 
      id: params[2] 
    }
    if(fieldname === "id"){
      options = {
        ...options,
        form: row.form,
        form_id: row.form_id
      }
    }
    currentModule({ 
      data: { module: APP_MODULE.EDIT }, 
      content: { fkey: "checkEditor", args: [options, EDITOR_EVENT.LOAD_EDITOR] }
    })
  }

  editFilter({index, fieldname, value}) {
    const { queries } = this.app.modules
    const { setData, data } = this.store
    const { vkey, view, filters } = this.store.data[APP_MODULE.SEARCH]

    const viewDef = queries[vkey]()[view]
    const _filters = { ...filters }
    if((fieldname === "filtertype") || (fieldname === "value")){
      _filters[view][index] = {
        ..._filters[view][index],
        [fieldname]: value
      }
    }
    if(fieldname === "fieldname"){
      if (Object.keys(viewDef.fields).includes(value)) {
        const frow = viewDef.fields[value]
        _filters[view][index] = {
          ..._filters[view][index],
          fieldname: value, fieldtype: frow.fieldtype,
          sqlstr: frow.sqlstr, wheretype: frow.wheretype, filtertype: "===",
          value: defaultFilterValue(frow.fieldtype)
        }
      } else {
        const deffield = data.deffield.filter((df)=>(df.fieldname === value))[0]
        const deftype = {
          bool: {
            fieldtype: "bool",
            sqlstr: "fg.groupvalue='bool' and case when fv.value='true' then 1 else 0 end "
          },
          integer: {
            fieldtype: "integer",
            sqlstr: "fg.groupvalue='integer' and {FMSF_NUMBER} {CAS_INT}fv.value {CAE_INT} {FMEF_CONVERT} "
          },
          float: {
            fieldtype: "float",
            sqlstr: "fg.groupvalue='float' and {FMSF_NUMBER} {CAS_FLOAT}fv.value {CAE_FLOAT} {FMEF_CONVERT} "
          },
          date: {
            fieldtype: "date",
            sqlstr: "fg.groupvalue='date' and {FMSF_DATE} {CASF_DATE}fv.value{CAEF_DATE} {FMEF_CONVERT} "
          },
          customer: {
            fieldtype: "string",
            sqlstr: "rf_customer.custname "
          },
          tool: {
            fieldtype: "string",
            sqlstr: "rf_tool.serial "
          },
          product: {
            fieldtype: "string",
            sqlstr: "rf_product.partnumber "
          },
          trans: {
            fieldtype: "string",
            sqlstr: "rf_trans.transnumber "
          },
          transitem: {
            fieldtype: "string",
            sqlstr: "rf_trans.transnumber "
          },
          transmovement: {
            fieldtype: "string",
            sqlstr: "rf_trans.transnumber "
          },
          transpayment: {
            fieldtype: "string",
            sqlstr: "rf_trans.transnumber "
          },
          project: {
            fieldtype: "string",
            sqlstr: "rf_project.pronumber "
          },
          employee: {
            fieldtype: "string",
            sqlstr: "rf_employee.empnumber "
          },
          place: {
            fieldtype: "string",
            sqlstr: "rf_place.planumber "
          }
        }
        let fieldtype = "string"; let sqlstr = "fv.value ";
        if(deftype[deffield.fieldtype]){
          fieldtype = deftype[deffield.fieldtype].fieldtype
          sqlstr = deftype[deffield.fieldtype].sqlstr
        }
        _filters[view][index] = {
          ..._filters[view][index],
          fieldname: deffield.fieldname,
          fieldlimit: ["and","fv.fieldname","=",`'${deffield.fieldname}'`],
          fieldtype, sqlstr,
          wheretype: "where", filtertype: "===", value: ""
        }
      }
    }
    setData(APP_MODULE.SEARCH, { filters: _filters })
  }

  editRow(row) {
    const { currentModule } = this.app
    const params = row.id.split("/")
    const options = { 
      ntype: params[0], 
      ttype: params[1], 
      id: parseInt(params[2],10), 
      item: row 
    }
    if (options.ntype === "servercmd") {
      this.showServerCmd(options.id)
    } else {
      currentModule({ 
        data: { module: APP_MODULE.EDIT }, 
        content: { fkey: "checkEditor", args: [options, EDITOR_EVENT.LOAD_EDITOR] }
      })
    }
  }

  exportResult(fields) {
    const { getSetting, saveToDisk } = this.app
    const { result, view } = this.store.data[APP_MODULE.SEARCH]
    const filename = `${view}.csv`
    let data = ""
    const labels = Object.keys(fields).map(fieldname => fields[fieldname].label)
    data += `${labels.join(getSetting("export_sep"))  }\n`
    result.forEach(row => {
      const cols = Object.keys(fields).map(fieldname => (typeof row[`export_${fieldname}`] !== "undefined") ? row[`export_${fieldname}`] : row[fieldname])
      data += `${cols.join(getSetting("export_sep"))  }\n`
    });
    const csvUrl = URL.createObjectURL(new Blob([data], {type : 'text/csv;charset=utf-8;'}))
    saveToDisk(csvUrl, filename)
  }

  onBrowserEvent({key, data}){
    const { showHelp, currentModule } = this.app
    const { setData } = this.store
    switch (key) {
      case BROWSER_EVENT.CHANGE:
        setData(APP_MODULE.SEARCH, {
          [data.fieldname]: data.value 
        })
        break;

      case BROWSER_EVENT.ADD_FILTER:
        this.addFilter()
        break;

      case BROWSER_EVENT.BOOKMARK_SAVE:
        this.saveBookmark()
        break;

      case BROWSER_EVENT.BROWSER_VIEW:
        this.browserView()
        break;

      case BROWSER_EVENT.CURRENT_PAGE:
        setData(APP_MODULE.SEARCH, {
          page: data.value 
        })
        break;

      case BROWSER_EVENT.DELETE_FILTER:
        this.deleteFilter(data.value)
        break;

      case BROWSER_EVENT.EDIT_FILTER:
        this.editFilter(data)
        break;

      case BROWSER_EVENT.EXPORT_RESULT:
        this.exportResult(data.value)
        break;

      case BROWSER_EVENT.EDIT_CELL:
        this.editCell(data)
        break;

      case BROWSER_EVENT.SET_COLUMNS:
        this.setColumns(data.fieldname, data.value)
        break;

      case BROWSER_EVENT.SET_FORM_ACTION:
        currentModule({ 
          data: { module: APP_MODULE.EDIT }, 
          content: { fkey: "checkEditor", args: [data, EDITOR_EVENT.FORM_ACTION] }
        })
        break;

      case BROWSER_EVENT.SHOW_BROWSER:
        this.showBrowser(data.value, data.vname)
        break;

      case BROWSER_EVENT.SHOW_HELP:
        showHelp(data.value)
        break;

      case BROWSER_EVENT.SHOW_TOTAL:
        this.showTotal(data)
        break;
    
      default:
        break;
    }
  }

  onModalEvent({key, data}){
    const { setData } = this.store
    switch (key) {
      case MODAL_EVENT.CANCEL:
        setData("current", { modalForm: null })
        break;
      
      case MODAL_EVENT.SEARCH:
        this.quickSearch(data.value)
        break;

      case MODAL_EVENT.SELECTED:
        this.editRow(data.value)
        break;

      case MODAL_EVENT.CURRENT_PAGE:
        setData(APP_MODULE.SEARCH, {
          page: data.value 
        })
        break;
    
      default:
        break;
    }
  }

  onSideEvent({key, data}){
    const { currentModule } = this.app
    const { setData } = this.store
    switch (key) {
      case SIDE_EVENT.CHANGE:
        setData(APP_MODULE.SEARCH, {
          [data.fieldname]: data.value 
        })
        break;

      case SIDE_EVENT.BROWSER:
        this.showBrowser(data.value)
        break;

      case SIDE_EVENT.QUICK:
        setData(APP_MODULE.SEARCH, { 
          seltype: "selector",
          result: [], qview: data.value, qfilter: "", page: 1 
        })
        setData("current", { side: SIDE_VISIBILITY.HIDE })
        break;
    
      case SIDE_EVENT.CHECK:
        currentModule({ 
          data: { module: APP_MODULE.EDIT }, 
          content: { fkey: "checkEditor", args: [data, EDITOR_EVENT.LOAD_EDITOR] }
        })
        break;

      default:
        break;
    }
  }

  async quickSearch(filter) {
    const { setData } = this.store
    const { quickSearch, resultError } = this.app
    const { qview } = this.store.data[APP_MODULE.SEARCH]
    const view = await quickSearch(qview, filter)
    if(view.error){
      return resultError(view)
    }
    setData(APP_MODULE.SEARCH, { result: view.result, qfilter: filter, page: 1 })
    return true
  }

  saveBookmark(){
    const { queries } = this.app.modules
    const { saveBookmark } = this.app
    const { vkey, view } = this.store.data[APP_MODULE.SEARCH]
    saveBookmark(["browser", queries[vkey]()[view].label])
  }

  setColumns(fieldname, visible) {
    const { setData } = this.store
    const { view, columns } = this.store.data[APP_MODULE.SEARCH]
    if(visible){
      setData(APP_MODULE.SEARCH, { 
        columns: {
          ...columns,
          [view]: {
            ...columns[view],
            [fieldname]: true
          }
        }
      })
    } else {
      const { [fieldname]: value, ...viewCols } = columns[view];
      setData(APP_MODULE.SEARCH, { 
        columns: {
          ...columns,
          [view]: viewCols
        }
      })
    }
  }

  async showBrowser(vkey, view, searchData){
    const { queries } = this.app.modules
    const { getSql, requestData, resultError } = this.app
    const { setData, data } = this.store
    const search_data = searchData || data[APP_MODULE.SEARCH]
    setData("current", { side: SIDE_VISIBILITY.HIDE })
    let search = {
      ...search_data,
      seltype: "browser",
      vkey, 
      view: (typeof view==="undefined") ? Object.keys(queries[vkey]())[1] : view, 
      result: [], deffield:[],
      show_dropdown: false,
      show_header: (typeof search_data.show_header === "undefined") ? true : search_data.show_header,
      show_columns: (typeof search_data.show_columns === "undefined") ? false : search_data.show_columns,
      page: search_data.page || 1,
    }
    const views = [
      { key: "deffield",
        text: getSql(data[APP_MODULE.LOGIN].data.engine, 
          queries[vkey]().options.deffield_sql).sql,
        values: [] 
      }
    ]
    const options = { method: "POST", data: views }
    const result = await requestData("/view", options)
    if(result.error){
      return resultError(result)
    }
    search = {
      ...search,
      deffield: result.deffield
    }
    if(!search.filters[search.view]){
      search = {
        ...search,
        filters: {
          ...search.filters,
          [search.view]: []
        }
      }
    }
    const viewDef = queries[vkey]()[search.view]
    if (typeof search.columns[search.view] === "undefined") {
      search = {
        ...search,
        columns: {
          ...search.columns,
          [search.view]: {}
        }
      }
      for(let fic = 0; fic < Object.keys(viewDef.columns).length; fic += 1) {
        const fieldname = Object.keys(viewDef.columns)[fic];
        search = {
          ...search,
          columns: {
            ...search.columns,
            [search.view]: {
              ...search.columns[search.view],
              [fieldname]: viewDef.columns[fieldname]
            }
          }
        }
      }
    }
    if (Object.keys(search.columns[search.view]).length === 0) {
      for(let v = 0; v < 3; v += 1) {
        const fieldname = Object.keys(viewDef.fields)[v];
        search = {
          ...search,
          columns: {
            ...search.columns,
            [search.view]: {
              ...search.columns[search.view],
              [fieldname]: true
            }
          }
        }
      }
    }
    setData(APP_MODULE.SEARCH, search)
    return true
  }

  showServerCmd(menu_id) {
    const { modalServer } = this.module
    const { request, resultError, requestData, showToast } = this.app
    const { setData } = this.store
    const login = this.store.data[APP_MODULE.LOGIN]
    const { session } = this.store.data
    const menuCmd = login.data.menuCmds.filter(item => (item.id === parseInt(menu_id, 10)))[0]
    if(menuCmd){
      const menuFields = login.data.menuFields.filter(item => (item.menu_id === parseInt(menu_id, 10)))
      const _values = {}
      menuFields.forEach(mfield => {
        switch (mfield.fieldtypeName) {
          case "bool":
            _values[mfield.fieldname] = false
            break;
          case "float":
          case "integer":
            _values[mfield.fieldname] = 0
            break;
          default:
            _values[mfield.fieldname] = ""
            break;
        }
      });
      const modalForm = modalServer({
        cmd: {...menuCmd}, 
        fields: [...menuFields], 
        values: _values,
        onEvent: {
          onModalEvent: async (modalResult) => {
            setData("current", { modalForm: null })
            if(modalResult.key === MODAL_EVENT.OK){
              const options = modalResult.data
              const query = new URLSearchParams();
              const values = {...options.values}
              options.fields.forEach((field) => {
                if (field.fieldtypeName === "bool") {
                  query.append(field.fieldname, (options.values[field.fieldname])?1:0)
                  values[field.fieldname] = (options.values[field.fieldname])?1:0
                } else {
                  query.append(field.fieldname, options.values[field.fieldname])
                }
              })
              if (options.cmd.methodName === "get") {
                let server = options.cmd.address || ""
                /* c8 ignore start */
                if((server === "") && options.cmd.funcname && (options.cmd.funcname !== "")){
                  server = (session.serverURL === "SERVER")?
                    `${session.apiPath}/${options.cmd.funcname}` : 
                    `${login.server}/${options.cmd.funcname}`
                }
                if (server!=="") {
                  window.open(`${server}?${query.toString()}`, '_system')
                }
                /* c8 ignore end */
              } else {
                const params_data = {
                  key: options.cmd.funcname || options.cmd.menukey,
                  values
                }
                let result
                if(options.cmd.address && (options.cmd.address !== "")){
                  try {
                    result = await request(options.cmd.address, {
                      method: "POST",
                      headers: { "Content-Type": "application/json" },
                      body: JSON.stringify(params_data)
                    })
                  /* c8 ignore start */
                  } catch (error) {
                    resultError(error)
                    return null
                  }
                  /* c8 ignore end */
                } else {
                  result = await requestData("/function", {
                    method: "POST", data: params_data
                  })
                }
                if(result.error){
                  resultError(result)
                  return null
                }
                let message = result
                if(typeof result === "object"){
                  message = JSON.stringify(result,null,"  ")
                }
                showToast(TOAST_TYPE.SUCCESS, message, 0)
              }
            }
            return true
          }
        }
      })
      setData("current", { modalForm })
    }
  }

  showTotal({fields, totalFields}){
    const { modalTotal } = this.module
    const { setData } = this.store
    const { result } = this.store.data[APP_MODULE.SEARCH]
    const getValidValue = (value) => {
      if(Number.isNaN(parseFloat(value))) {
        return 0
      }
      return parseFloat(value)
    }
    let total = totalFields
    const isDeffield = Object.keys(fields).includes("deffield_value")
    result.forEach(row => {
      if (isDeffield) {
        if (typeof total.totalFields[row.fieldname] !== "undefined") {
          total = {
            ...total,
            totalFields: {
              ...total.totalFields,
              [row.fieldname]: total.totalFields[row.fieldname] + getValidValue(row.export_deffield_value)
            }
          }
        }
      } else {
        Object.keys(total.totalFields).forEach(fieldname => {
          if (typeof row[`export_${fieldname}`] !== "undefined") {
            total = {
              ...total,
              totalFields: {
                ...total.totalFields,
                [fieldname]: total.totalFields[fieldname] + getValidValue(row[`export_${fieldname}`])
              }
            }
          } else {
            total = {
              ...total,
              totalFields: {
                ...total.totalFields,
                [fieldname]: total.totalFields[fieldname] + getValidValue(row[fieldname])
              }
            }
          }
        });
      }
    })
    setData("current", { modalForm: modalTotal(total) })
  }
}