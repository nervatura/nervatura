import { APP_MODULE, LOGIN_PAGE_EVENT } from '../config/enums.js'

export class LoginController {
  constructor(host, app) {
    this.host = host;
    this.store = app.store;
    this.app = app;
    host.addController(this);
  }

  async userLog(loginData) {
    const options = { 
      method: "POST", token: loginData.token, 
      data: [
        { keys: { 
          employee_id: loginData.employee.empnumber,
          logstate: "login"
        }}
      ] 
    }
    const result = await this.app.requestData("/log", options)
    return result
  }
  
  async loginData(params) {
    let data = {
      ...this.store.data[APP_MODULE.LOGIN],
      token: params.token, engine: params.engine 
    }
    let views = [
      { key: "employee",
        text: this.app.getSql(params.engine, { 
          select: ["e.*", "ug.groupvalue as usergroupName", "dp.groupvalue as departmentName"], 
          from: "employee e",
          inner_join: ["groups ug", "on", ["e.usergroup", "=", "ug.id"]],
          left_join: ["groups dp", "on", ["e.department", "=", "dp.id"]], 
          where: ["username", "=", "?"] }).sql,
        values: [data.username] },
      { key: "menuCmds",
        text: this.app.getSql(params.engine, { 
          select: ["m.*", "st.groupvalue as methodName"], 
          from: "ui_menu m",
          inner_join: ["groups st", "on", ["m.method", "=", "st.id"]], }).sql,
        values: [] },
      { key: "menuFields",
        text: this.app.getSql(params.engine, { 
          select: ["mf.*", "ft.groupvalue as fieldtypeName"], 
          from: "ui_menufields mf",
          inner_join: ["groups ft", "on", ["mf.fieldtype", "=", "ft.id"]],
          order_by: ["menu_id", "orderby"] }).sql,
        values: [] },
      { key: "userlogin",
        text: this.app.getSql(params.engine, {
          select: ["value"], from: "fieldvalue",
          where: [["ref_id", "is", "null"], ["and", "fieldname", "=", "'log_login'"]]
        }).sql,
        values: [] },
      { key: "groups",
        text: this.app.getSql(params.engine, {
          select: ["*"], from: "groups",
          where: ["groupname", "in", [[], "'usergroup'", "'nervatype'", "'transtype'", "'inputfilter'",
            "'transfilter'", "'department'", "'logstate'", "'fieldtype'", "'service'"]]
        }).sql,
        values: [] }
    ]
    let options = { method: "POST", token: params.token, data: views }
    let view = await this.app.requestData("/view", options)
    if(view.error){
      return view
    }

    data = {
      ...data, ...view,
      employee: view.employee[0],
      userlogin: (view.userlogin.length>0) ? view.userlogin[0].value : "false"
    }

    views = [
      { key: "audit",
        text: this.app.getSql(params.engine, { 
          select: ["au.nervatype", "nt.groupvalue as nervatypeName", "au.subtype",
            "case when nt.groupvalue = 'trans' then st.groupvalue else m.menukey end as subtypeName",
            "au.inputfilter", "ip.groupvalue as inputfilterName", "au.supervisor"],
          from: "ui_audit au",
          inner_join: [
            ["groups nt", "on", ["au.nervatype", "=", "nt.id"]],
            ["groups ip", "on", ["au.inputfilter", "=", "ip.id"]]],
          left_join: [
            ["groups st", "on", ["au.subtype", "=", "st.id"]],
            ["ui_menu m", "on", ["au.subtype", "=", "m.id"]]],
          where: ["au.usergroup", "=", "?"] 
        }).sql,
        values: [data.employee.usergroup] },
      { key: "transfilter",
        text: this.app.getSql(params.engine, {
          select: ["ref_id_2 as transfilter", "g.groupvalue as transfilterName"], 
          from: "link",
          inner_join: ["groups g", "on", ["link.ref_id_2", "=", "g.id"]],
          where: [["ref_id_1", "=", "?"], ["and", "link.deleted", "=", "0"],
          ["and", "nervatype_1", "in",
            [{
              select: ["id"], from: "groups",
              where: [["groupname", "=", "'nervatype'"], ["and", "groupvalue", "=", "'groups'"]]
            }]],
          ["and", "nervatype_2", "in",
            [{
              select: ["id"], from: "groups",
              where: [["groupname", "=", "'nervatype'"], ["and", "groupvalue", "=", "'groups'"]]
            }]]]
        }).sql,
        values: [data.employee.usergroup] 
      }
    ]

    options = { method: "POST", token: params.token, data: views }
    view = await this.app.requestData("/view", options)
    if(view.error){
      return view
    }
    data = {
      ...data,
      transfilter: (view.transfilter.length>0) ? view.transfilter[0].transfilter : null,
      audit: view.audit
    }
    if(data.transfilter === null){
      const transfilter = data.groups.filter((group)=> ((group.groupname === "transfilter") && (group.groupvalue === "all")))[0]
      data = {
        ...data,
        transfilter: transfilter.id,
        transfilterName: "all"
      }
    } else {
      data = {
        ...data,
        transfilterName: view.transfilter[0].transfilterName
      }
    }

    data = {
      ...data,
      audit_filter: {trans:{}, menu:{}, report:{}},
      edit_new: [[],[],[],[]]
    }
    const trans = [
      ["offer",0],["order",0],["worksheet",0],["rent",0],["invoice",0],["receipt",0],
      ["bank",1],["cash",1],
      ["delivery",2],["inventory",2],["waybill",2],["production",2],["formula",2]]
    trans.forEach((transtype) => {
      const audit = data.audit.filter((item)=> ((item.nervatypeName === "trans") && (item.subtypeName === transtype[0])))[0]
      data = {
        ...data,
        audit_filter: { 
          ...data.audit_filter,
          trans: { 
            ...data.audit_filter.trans,
            [transtype[0]]: (audit) ? 
              [audit.inputfilterName, audit.supervisor] : ["all",1]
          }
        }
      }
      if (data.audit_filter.trans[transtype[0]][0] !== "disabled"){
        data = {
          ...data,
          edit_new: {
            ...data.edit_new,
            [transtype[1]]: [...data.edit_new[transtype[1]], transtype[0]]
          }
        }
      }
    });

    const nervatype = ["customer","product","employee","tool","project","setting","audit"]
    nervatype.forEach((ntype) => {
      const audit = data.audit.filter((item)=> ((item.nervatypeName === ntype) && (item.subtypeName === null)))[0]
      data = {
        ...data,
        audit_filter: {
          ...data.audit_filter,
          [ntype]: (audit) ? 
            [audit.inputfilterName, audit.supervisor] : ["all",1]
        }
      }
      if (data.audit_filter[ntype][0] !== "disabled" && 
        ntype !== "setting" && ntype !== "audit"){
          data = {
            ...data,
            edit_new: {
              ...data.edit_new,
              3: [...data.edit_new[3], ntype]
            }
          }
        }
    });

    return data
  }

  async onLogin() {
    const { data, msg } = this.store
    const options = {
      method: "POST",
      data: {
        username: data[APP_MODULE.LOGIN].username, password: data[APP_MODULE.LOGIN].password,
        database: data[APP_MODULE.LOGIN].database
      }
    }
    const result = await this.app.requestData("/auth/login", options)
    if(result.token && result.engine ){
      if(!data.session.engines.includes(result.engine)){
        return this.app.resultError({ error: { message: msg("Invalid database type!", { id: "login_engine_err" }) } })
      }
      if(!data.session.service.includes(result.version)){
        return this.app.resultError({ error: { message: msg("Invalid service version!", { id: "login_version_err" }) } })
      }
      const resultData = await this.loginData(result)
      if(resultData.error){
        return this.app.resultError(resultData)
      }
      return this.setLogin(resultData)
    } 
    return this.app.resultError(result)
  }

  async setLogin(loginData){
    const { data, setData } = this.store
    if (loginData.userlogin === "t" || loginData.userlogin === "true") {
      const log = await this.userLog(loginData)
      if(log.error){
        this.app.resultError(log)
        return
      }
    }

    setData("current", {
      module: APP_MODULE.SEARCH
    })
    setData(APP_MODULE.LOGIN, {
      data: loginData
    })
    localStorage.setItem("database", data[APP_MODULE.LOGIN].database);
    localStorage.setItem("username", data[APP_MODULE.LOGIN].username);
    localStorage.setItem("server", data[APP_MODULE.LOGIN].server);
    this.app.loadBookmark({ user_id: loginData.employee.id, token: loginData.token })
  }

  tokenError(err, callback) {
    /* c8 ignore next 3 */
    if(callback){
      return window.location.replace(`${callback}?error=${window.btoa(err.message)}`)
    }
    this.store.setData("current", { 
      request: false 
    })
    return this.app.resultError(err)
  }

  async tokenValidation(params) {
    const options = {
      method: "GET", token: params.access_token,
    }
    const validate = await this.app.requestData("/auth/validate", options)
    if(validate.error){
      return this.tokenError(validate.error, params.callback)
    }
    this.store.setData(APP_MODULE.LOGIN, {
      username: validate.username,
      database: validate.database,
      callback: params.callback,
    } )
    const resultData = await this.loginData({
      token: params.access_token,
      engine: validate.engine,
      version: validate.version
    })
    if(resultData.error){
      return this.app.resultError(resultData)
    }
    window.history.replaceState(null, null, window.location.pathname)
    return this.setLogin(resultData)
  }

  async setCodeToken(params) {
    if(params.callback){
      const options = {
        headers: { "Content-Type": "application/json" },
        method: "POST",
        body: JSON.stringify({
          code: params.code
        })
      }
      try {
        this.store.setData("current", {
          request: true 
        })
        const result = await this.app.request(params.callback, options)
        if(result.access_token){
          return this.tokenValidation({ access_token: result.access_token, callback: result.callback })
        }
        return this.tokenError(result, params.error)
      } catch (err) {
        return this.tokenError(err, params.error)
      }
    } else {
      return this.tokenError({
        id: "error_unauthorized", type: "error", message: "Unauthorized user"
      }, params.error)
    }
  }

  onPageEvent({key, data}){
    switch (key) {
      case LOGIN_PAGE_EVENT.CHANGE:
        this.store.setData(APP_MODULE.LOGIN, {
          [data.fieldname]: data.value 
        })
        break;

      case LOGIN_PAGE_EVENT.THEME:
      case LOGIN_PAGE_EVENT.LANG:
        this.store.setData("current", { 
          [key]: data 
        })
        localStorage.setItem([key], data);
        break;
      
      case LOGIN_PAGE_EVENT.LOGIN:
        this.onLogin()
        break;
    
      default:
        break;
    }
  }
}