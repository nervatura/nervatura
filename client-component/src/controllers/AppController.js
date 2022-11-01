import { APP_MODULE, TOAST_TYPE, MODAL_EVENT } from '../config/enums.js'
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
  constructor(host, store) {
    this.host = host;
    this.store = store
    this.getSql = getSql
    this.request = request
    this.saveToDisk = saveToDisk
    host.addController(this);
  }

  signOut() {
    if(this.store.data[APP_MODULE.LOGIN].callback){
      window.location.replace(this.store.data[APP_MODULE.LOGIN].callback)
      return
    }
    this.store.setData(APP_MODULE.LOGIN, { 
      data: null, token: null 
    })
  }

  showHelp(key) {
    const { data } = this.store
    const element = document.createElement("a")
    element.setAttribute("href", data.session.helpPage+key)
    element.setAttribute("target", "_blank")
    document.body.appendChild(element)
    element.click()
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
        data.session.proxy+data.session.apiPath+path : data.login.server+path
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
      
      const result = await request(url, options)
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
    const { setData, showToast, msg } = this.store
    if(result.error){
      setData("error", result.error )
    }
    if(result.error && result.error.message){
      showToast(TOAST_TYPE.ERROR, result.error.message)
    } else {
      showToast(TOAST_TYPE.ERROR, 
        msg("Internal Server Error", { id: "error_internal" }) )
    }
  }

  async loadBookmark(params) {
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
    this.store.setData(APP_MODULE.BOOKMARK, { bookmark, history })
    return { bookmark, history }
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

  async quickSearch(qview, qfilter) {
    const { quick } = this.host
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

  saveBookmark(params){
    const { inputBox } = this.host
    const { data, setData, msg } = this.store
    const login = data[APP_MODULE.LOGIN]
    const search = data[APP_MODULE.SEARCH]
    const edit = data[APP_MODULE.EDIT]
    const  modalForm = inputBox({ 
      title: msg("", { id: "msg_bookmark_new" }), 
      message: msg("", { id: "msg_bookmark_name" }),  
      value: (params[0] === "browser") ? params[1] : edit.current.item[params[2]],  
      showValue: true,
      labelCancel: msg("", { id: "msg_cancel" }),
      labelOK: msg("", { id: "msg_ok" }),
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
    setData("current", { modalForm })
  }

}