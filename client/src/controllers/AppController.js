import { APP_MODULE, TOAST_TYPE, MODAL_EVENT, SIDE_VISIBILITY } from '../config/enums.js'
import { getSql, request, saveToDisk } from './Utils.js'

export const UserFilter = {
  usergroup_filter: {
    transitem: ["t.cruser_id","in",[{select:["id"], from:"employee", 
      where:["usergroup","=","?"]}]],
    transmovement: ["t.cruser_id","in",[{select:["id"], from:"employee", 
      where:["usergroup","=","?"]}]],
    transpayment: ["t.cruser_id","in",[{select:["id"], from:"employee", 
      where:["usergroup","=","?"]}]]
  },
  employee_filter: {
    transitem: ["t.cruser_id","=","?"],
    transmovement: ["t.cruser_id","=","?"],
    transpayment: ["t.cruser_id","=","?"]
  }
}

export class AppController {
  constructor(host) {
    this.host = host;
    this.modules = {}
    
    this.getSql = getSql
    this.request = request
    this.saveToDisk = saveToDisk

    this.createHistory = this.createHistory.bind(this)
    this.currentModule = this.currentModule.bind(this)
    this.getAuditFilter = this.getAuditFilter.bind(this)
    this.getDataFilter = this.getDataFilter.bind(this)
    this.getSetting = this.getSetting.bind(this)
    this.getUserFilter = this.getUserFilter.bind(this)
    this.loadBookmark = this.loadBookmark.bind(this)
    this.msg = this.msg.bind(this)
    this.quickSearch = this.quickSearch.bind(this)
    this.requestData = this.requestData.bind(this)
    this.resultError = this.resultError.bind(this)
    this.saveBookmark = this.saveBookmark.bind(this)
    this.showHelp = this.showHelp.bind(this)
    this.showToast = this.showToast.bind(this)
    this.signOut = this.signOut.bind(this)
    this.tokenLogin = this.tokenLogin.bind(this)
    host.addController(this);
  }

  hostConnected(){
    const { state, setData } = this.host
    this.store = {
      data: state.data,
      setData
    };
    
    this._loadConfig(window.location)
  }

  async _loadConfig(location){
    const { data, setData } = this.store
    const getPath = () => {
      const getParams = (prmString) => {
        const params = {}
        prmString.split('&').forEach(prm => {
          const index = String(prm).indexOf("=")
          const fname = String(prm).substring(0,(index >0)?index:String(prm).length)
          const value = ((index > -1) && (index < String(prm).length)) ? String(prm).substring(index+1) : ""
          params[fname] = value
        });
        return params
      }
      if(location.hash){
        return ["hash", getParams(location.hash.substring(1))]
      }
      if(location.search){
        return ["search", getParams(location.search.substring(1))]
      }
      const path = location.pathname.substring(1).split("/")
      return [path[0], path.slice(1)]
    }

    this.request(`${data[APP_MODULE.LOGIN].server}/config`, {
      method: "GET",
      headers: { "Content-Type": "application/json" }
    }).then(result => {
      let config = {...data.session}
      if(result.locales && (typeof result.locales === "object")){
        config = {
          ...config,
          locales: {
            ...config.locales,
            ...result.locales
          }
        }
      }
      setData("session", config )
      if(localStorage.getItem("lang") && config.locales[localStorage.getItem("lang")] 
        && (localStorage.getItem("lang") !== data.current.lang)){
          setData("current", {
            lang: localStorage.getItem("lang")
          })
      }
    }).catch(error => {
      setData("error", error )
      this.showToast(TOAST_TYPE.ERROR, error.message)
    })

    const [ current, params ] = getPath()
    if(current === "hash"){
      if (params.access_token || params.code){
        this.tokenLogin(params)
        return
      }
    }
    import('../components/Login/client-login.js');
  }

  async createHistory(ctype) {
    const { data, setData } = this.store

    let history = {
      datetime: `${new Date().toISOString().slice(0,10)}T${new Date().toLocaleTimeString("en",{hour12: false}).replace("24","00")}`,
      type: ctype, 
      type_title: this.msg("", { id: `label_${ctype}` }),
      ntype: data[APP_MODULE.EDIT].current.type,
      transtype: data[APP_MODULE.EDIT].current.transtype || "",
      id: data[APP_MODULE.EDIT].current.item.id
    }
    let title = (history.ntype === "trans") ?
      `${data[APP_MODULE.EDIT].template.options.title} | ${data[APP_MODULE.EDIT].current.item[data[APP_MODULE.EDIT].template.options.title_field]}` :
      data[APP_MODULE.EDIT].template.options.title
    if ((history.ntype !== "trans") && (typeof data[APP_MODULE.EDIT].template.options.title_field !== "undefined")){
      title += ` | ${data[APP_MODULE.EDIT].current.item[data[APP_MODULE.EDIT].template.options.title_field]}`
    }
    history = {...history,
      title
    }
    const bookmark = {...data.bookmark}
    let userconfig = {}
    if (bookmark.history) {
      userconfig = {...bookmark.history,
        cfgroup: `${new Date().toISOString().slice(0,10)}T${new Date().toLocaleTimeString("en",{hour12: false}).replace("24","00")}`
      }
      let history_values = JSON.parse(userconfig.cfvalue);
      history_values.unshift(history)
      if (history_values.length > parseInt(this.getSetting("history"),10)) {
        history_values = history_values.slice(0, parseInt(this.getSetting("history"),10))
      }
      userconfig = {...userconfig,
        cfname: history_values.length,
        cfvalue: JSON.stringify(history_values)
      }
    } else {
      userconfig = {...userconfig,
        employee_id: data.login.data.employee.id,
        section: "history",
        cfgroup: `${new Date().toISOString().slice(0,10)}T${new Date().toLocaleTimeString("en",{hour12: false}).replace("24","00")}`,
        cfname: 1,
        cfvalue: JSON.stringify([history])
      }
    }
    const options = { method: "POST", data: [userconfig] }
    const result = await this.requestData("/ui_userconfig", options)
    if(result.error){
      return this.resultError(result)
    }
    return setData("bookmark", { history: userconfig})
  }

  async currentModule({data, content}) {
    const { setData } = this.store

    const modules = {
      forms: async ()=>{
        const { Forms } = await import('./Forms.js');
        const { Dataset } = await import('./Dataset.js');
        const { Sql } = await import('./Sql.js');
        const { InitItem, Validator } = await import('./Validator.js');

        this.modules.forms = Forms(this)
        this.modules.dataSet = {...Dataset}
        this.modules.initItem = InitItem(this)
        this.modules.sql = Sql(this)
        this.modules.validator = Validator(this)
      },
      quick: async ()=>{
        const { Quick } = await import('./Quick.js');
        const { Queries } = await import('./Queries.js');

        this.modules.quick = {...Quick}
        this.modules.queries = Queries(this)
      },
      search: async ()=>{
        const { SearchController } = await import('./SearchController.js');
        await import('../components/Search/Search/client-search.js');
        await import('../components/MenuBar/client-menubar.js');

        if(!this.modules.quick){
          await modules.quick()
        }
        this.modules.search = new SearchController(this.host)
      },
      edit: async ()=>{
        const { EditController } = await import('./EditController.js');
        await import('../components/Edit/Edit/client-edit.js');

        if(!this.modules.forms){
          await modules.forms()
        }
        this.modules.edit = new EditController(this.host)
      },
      setting: async ()=>{
        const { SettingController } = await import('./SettingController.js');
        await import('../components/Setting/Setting/client-setting.js');

        if(!this.modules.forms){
          await modules.forms()
        }
        this.modules.setting = new SettingController(this.host)
      },
      template: async ()=>{
        const { TemplateController } = await import('./TemplateController.js');
        await import('../components/Template/Template/client-template.js');

        this.modules.template = new TemplateController(this.host)
      }
    }

    if(!this.modules[data.module]){
      await modules[data.module]()
    }
    setData("current", {...data})
    if(content && this.modules[data.module]){
      this.modules[data.module][content.fkey](...content.args)
    }
  }

  getAuditFilter(nervatype, transtype) {
    const login = this.store.data[APP_MODULE.LOGIN]
    let retvalue = ["all",1]; let audit;
    switch (nervatype) {
      case "trans":
      case "menu":
        audit = login.data.audit.filter((au)=> ((au.nervatypeName === nervatype) && (au.subtypeName === transtype)))[0]
        break;
      case "report":
        audit = login.data.audit.filter((au)=> ((au.nervatypeName === nervatype) && (au.subtype === transtype)))[0]
        break;
      default:
        audit = login.data.audit.filter((au)=> (au.nervatypeName === nervatype))[0]
    }
    if (typeof audit !== "undefined") {
      retvalue = [audit.inputfilterName, audit.supervisor];
    }
    return retvalue;
  }

  getDataFilter(type, where, view) {
    let _where = where
    if(type === "transitem"){
      if (this.getAuditFilter("trans", "offer")[0] === "disabled") {
        _where = _where.concat(["and", ["tg.groupvalue", "<>", "'offer'"]]);
      }
      if (this.getAuditFilter("trans", "order")[0] === "disabled") {
        _where = _where.concat(["and", ["tg.groupvalue", "<>", "'order'"]]);
      }
      if (this.getAuditFilter("trans", "worksheet")[0] === "disabled") {
        _where = _where.concat(["and", ["tg.groupvalue", "<>", "'worksheet'"]]);
      }
      if (this.getAuditFilter("trans", "rent")[0] === "disabled") {
        _where = _where.concat(["and", ["tg.groupvalue", "<>", "'rent'"]]);
      }
      if (this.getAuditFilter("trans", "invoice")[0] === "disabled") {
        _where = _where.concat(["and", ["tg.groupvalue", "<>", "'invoice'"]]);
      }
    }
    if(type === "transpayment"){
      if (this.getAuditFilter("trans", "bank")[0] === "disabled") {
        _where = _where.concat(["and", ["tg.groupvalue", "<>", "'bank'"]]);
      }
      if (this.getAuditFilter("trans", "cash")[0] === "disabled") {
        _where = _where.concat(["and", ["tg.groupvalue", "<>", "'cash'"]]);
      }
    }
    if((type === "transmovement") && (view !== "InventoryView")){
      if (this.getAuditFilter("trans", "delivery")[0] === "disabled") {
        _where = _where.concat(["and", ["tg.groupvalue", "<>", "'delivery'"]]);
      }
      if (this.getAuditFilter("trans", "inventory")[0] === "disabled") {
        _where = _where.concat(["and", ["tg.groupvalue", "<>", "'inventory'"]]);
      }
    }
    return _where;
  }

  getSetting(key) {
    const { ui } = this.store.data
    if(key === "ui"){
      const values = {...ui}
        Object.keys(values).forEach(ikey => {
          if(localStorage.getItem(ikey)){
            values[ikey] = localStorage.getItem(ikey)
          }
        });
        return values
    }
    return localStorage.getItem(key) || ui[key] || "";
  }

  getUserFilter(type) {
    const login = this.store.data[APP_MODULE.LOGIN]
    const filter = { params: [], where: []}
    if(login.data.transfilterName === "usergroup"){
      if (typeof UserFilter.usergroup_filter[type] !== "undefined") {
        filter.where = ["and", UserFilter.usergroup_filter[type]]
        filter.params = [login.data.employee.usergroup]
      }
    }
    if(login.data.transfilterName === "own"){
      if (typeof UserFilter.employee_filter[type] !== "undefined") {
        filter.where = ["and", UserFilter.employee_filter[type]]
        filter.params = [login.data.employee.id]
      }
    }
    return filter;
  }

  async loadBookmark(params) {
    const { setData } = this.store
    const result = await this.requestData("/ui_userconfig", {
      token: params.token,
      query: {
        filter: `employee_id;==;${params.user_id}`
      }
    })
    if(result.error){
      this.resultError(result)
      return null
    }
    const bookmark = result.filter(item => (item.section === "bookmark"))
    const history = result.filter(item => (item.section === "history"))[0]||null
    setData(APP_MODULE.BOOKMARK, { bookmark, history })
    return { bookmark, history }
  }

  msg(defaultValue, props) {
    let value = defaultValue
    const {locales} = this.store.data.session
    const {lang} = this.store.data.current
    if(locales[lang] && locales[lang][props.id]){
      value = locales[lang][props.id]
    } else if((lang !== "en") && locales.en[props.id]) {
      value = locales.en[props.id]
    }
    return value
  }

  async quickSearch(qview, qfilter) {
    const { quick } = this.modules
    const login = this.store.data[APP_MODULE.LOGIN]
    const query = quick[qview](String(login.data.employee.usergroup))
    let _sql = { ...query.sql }
    let params = []; let _where = []
    if(qfilter !== ""){
      const filter = `{CCS}{JOKER}{SEP}lower(?){SEP}{JOKER}{CCE} `
      query.columns.forEach((column, index) => {
        _where.push([((index!==0)?"or":""),[`lower(${column[1]})`,"like", filter]])
        params.push(qfilter)
      });
      _where = ["and",[_where]]
    }
    _where = this.getDataFilter(qview, _where)
    if(_where.length > 0){
      _sql = {
        ..._sql,
        where: [..._sql.where, _where]
      }
    }

    const userFilter = this.getUserFilter(qview)
    if(userFilter.where.length > 0){
      _sql = {
        ..._sql,
        where: [..._sql.where, userFilter.where]
      }
      params = params.concat(userFilter.params)
    }
    
    const views = [
      { key: "result",
        text: getSql(login.data.engine, _sql).sql,
        values: params 
      }
    ]
    const options = { method: "POST", data: views }
    return this.requestData("/view", options)
  }

  async requestData(path, params, silent) {
    const { data, setData } = this.store
    let options = params
    try {
      if (!silent){
        setData("current", {
          request: true 
        })
      }
      let url = (data.session.serverURL === "SERVER")?
        data.session.apiPath+path : data.login.server+path
      const token = (data.login.data) ? data.login.data.token : options.token || ""
      if (!options.headers)
        options = {
          ...options, 
          headers: {}
        }
      options = {
        ...options, 
        headers: {
          ...options.headers,
          "Content-Type": "application/json"
        }
      }
      if(token !== ""){
        options = {
          ...options, 
          headers: {
            ...options.headers,
            "Authorization": `Bearer ${token}`
          }
        }
      }
      if (options.data){
        options = {
          ...options, 
          body: JSON.stringify(options.data) 
        }
      }
      if(options.query){
        const query = new URLSearchParams();
        Object.keys(options.query).forEach(key => {
          query.append(key, options.query[key])
        });
        url += `?${query.toString()}`
      }
      
      const result = await this.request(url, options)
      if (!silent) {
        setData("current", { 
          request: false 
        })
      }
      if(result && result.code){
        if(result.code === 401){
          this.signOut()
        }
        return { error: { message: result.message }, data: null }
      }
      return result
    } catch (err) {
      if(!silent){
        setData("current", { 
          request: false 
        })
      }
      return { error: { message: err.message }, data: null }
    }
  }

  resultError(result) {
    const { setData } = this.store
    if(result.error){
      setData("error", result.error )
    }
    if(result.error && result.error.message){
      this.showToast(TOAST_TYPE.ERROR, result.error.message)
    } else {
      this.showToast(TOAST_TYPE.ERROR, 
        this.msg("Internal Server Error", { id: "error_internal" }) )
    }
    return false
  }

  saveBookmark(params){
    const { inputBox } = this.host
    const { data, setData } = this.store
    const login = data[APP_MODULE.LOGIN]
    const search = data[APP_MODULE.SEARCH]
    const edit = data[APP_MODULE.EDIT]
    const  modalForm = inputBox({ 
      title: this.msg("", { id: "msg_bookmark_new" }), 
      message: this.msg("", { id: "msg_bookmark_name" }),  
      value: (params[0] === "browser") ? params[1] : edit.current.item[params[2]],  
      showValue: true,
      onEvent: {
        onModalEvent: async (modalResult) => {
          setData("current", { modalForm: null })
          if ((modalResult.key === MODAL_EVENT.OK) && (modalResult.data.value !== "")) {
            let userconfig = {
              employee_id: login.data.employee.id,
              section: "bookmark",
              cfgroup: params[0],
            }
            if((params[0]) === "browser"){
              userconfig = {
                ...userconfig,
                cfname: modalResult.data.value,
                cfvalue: JSON.stringify({
                  date: new Date().toISOString().split("T")[0],
                  vkey: search.vkey,
                  view: search.view,
                  filters: search.filters[search.view],
                  columns: search.columns[search.view]
                })
              }
            } else {
              userconfig = {
                ...userconfig,
                cfname: modalResult.data.value,
                cfvalue: JSON.stringify({
                  date: new Date().toISOString().split("T")[0],
                  ntype: edit.current.type,
                  transtype: edit.current.transtype,
                  id: edit.current.item.id,
                  info: (edit.current.type === "trans") 
                    ? (edit.dataset.trans[0].custname !== null) 
                      ? edit.dataset.trans[0].custname 
                      : edit.current.item.transdate 
                    : edit.current.item[params[3]]
                })
              }
            }

            const options = { method: "POST", data: [userconfig] }
            const result = await this.requestData("/ui_userconfig", options)
            if(result.error){
              return this.resultError(result)
            }
            this.loadBookmark({user_id: login.data.employee.id})
          }
          return true
        }
      }
    })
    setData("current", { modalForm, side: SIDE_VISIBILITY.HIDE })
  }

  showHelp(key) {
    const { data } = this.store
    const element = document.createElement("a")
    element.setAttribute("href", data.session.helpPage+key)
    element.setAttribute("target", "_blank")
    document.body.appendChild(element)
    element.click()
  }

  showToast(type, value, toastTimeout) {
    const { current } = this.store.data
    const timeout = (typeof(toastTimeout) !== "undefined") ? toastTimeout : this.getSetting("toastTimeout")
    if(current.toast){
      current.toast.show({
        type, value, timeout
      })
    }
  }

  signOut() {
    const { data, setData } = this.store
    /* c8 ignore next 4 */
    if(data[APP_MODULE.LOGIN].callback){
      window.location.replace(data[APP_MODULE.LOGIN].callback)
      return
    }
    setData(APP_MODULE.LOGIN, { 
      data: null, token: null 
    })
  }

  async tokenLogin(params) {
    const { LoginController } = await import('./LoginController.js');
    const login = new LoginController(this.host)
    if (params.access_token){
      login.tokenValidation(params)
    }
    if(params.code){
      login.setCodeToken(params)
    }
  }
}