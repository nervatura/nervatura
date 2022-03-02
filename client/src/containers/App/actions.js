import { lazy } from 'react'
import update from 'immutability-helper';
import 'whatwg-fetch';
import formatISO from 'date-fns/formatISO'

import Toastify from 'components/Form/Toastify'
import { getText, getSetting } from 'config/app'
import { UserFilter } from 'containers/Controller/Filter'

export const getSql = (engine, _sql) => {
  let prm_count = 0;
  const engine_type = (sql) => {
    switch (engine) {
      case "sqlite":
      case "sqlite3":
        sql = sql.replace(/{CCS}/g, "");
        sql = sql.replace(/{SEP}/g, "||");
        sql = sql.replace(/{CCE}/g, "");
        sql = sql.replace(/{CAS_TEXT}/g, "cast(");
        sql = sql.replace(/{CAE_TEXT}/g, " as text)");
        sql = sql.replace(/{CAS_INT}/g, "cast("); //cast as integer - start
        sql = sql.replace(/{CAE_INT}/g, " as integer)"); //cast as integer - end
        sql = sql.replace(/{CAS_FLOAT}/g, "cast(");
        sql = sql.replace(/{CAE_FLOAT}/g, " as double)"); //" as real)");
        sql = sql.replace(/{CAS_DATE}/g, "");
        sql = sql.replace(/{CASF_DATE}/g, "");
        sql = sql.replace(/{CAE_DATE}/g, "");
        sql = sql.replace(/{CAEF_DATE}/g, "");
        sql = sql.replace(/{FMSF_NUMBER}/g, "");
        sql = sql.replace(/{FMSF_DATE}/g, "");
        sql = sql.replace(/{FMEF_CONVERT}/g, "");
        sql = sql.replace(/{FMS_FLOAT}/g, "");
        sql = sql.replace(/{FME_FLOAT}/g, "");
        sql = sql.replace(/{FMS_INT}/g, "");
        sql = sql.replace(/{FME_INT}/g, "");
        sql = sql.replace(/{FMS_DATE}/g, "substr("); //format to iso date - start
        sql = sql.replace(/{FME_DATE}/g, ",1,10)"); //format to iso date - end
        sql = sql.replace(/{FMS_DATETIME}/g, "substr("); //format to iso datetime - start
        sql = sql.replace(/{FME_DATETIME}/g, ",1,19)"); //format to iso datetime - end
        sql = sql.replace(/{FMS_TIME}/g, "substr(time(");
        sql = sql.replace(/{FME_TIME}/g, "),0,6)");
        sql = sql.replace(/{JOKER}/g, "'%'");
        sql = sql.replace(/{CUR_DATE}/g, "date('now')");
        break;
      case "mysql":
        sql = sql.replace(/{CCS}/g, "concat(");
        sql = sql.replace(/{SEP}/g, ",");
        sql = sql.replace(/{CCE}/g, ")");
        sql = sql.replace(/{CAS_TEXT}/g, "cast(");
        sql = sql.replace(/{CAE_TEXT}/g, " as char)");
        sql = sql.replace(/{CAS_INT}/g, "cast("); //cast as integer - start
        sql = sql.replace(/{CAE_INT}/g, " as signed)"); //cast as integer - end
        sql = sql.replace(/{CAS_FLOAT}/g, "cast(");
        sql = sql.replace(/{CAE_FLOAT}/g, " as decimal)"); //" as decimal)");
        sql = sql.replace(/{CAS_DATE}/g, "cast(");
        sql = sql.replace(/{CASF_DATE}/g, "cast(");
        sql = sql.replace(/{CAE_DATE}/g, " as date)");
        sql = sql.replace(/{CAEF_DATE}/g, " as date)");
        sql = sql.replace(/{FMSF_NUMBER}/g, "");
        sql = sql.replace(/{FMSF_DATE}/g, "");
        sql = sql.replace(/{FMEF_CONVERT}/g, "");
        sql = sql.replace(/{FMS_FLOAT}/g, "replace(format(cast(");
        sql = sql.replace(/{FME_FLOAT}/g, " as decimal(10,2)),2),'.00','')");
        sql = sql.replace(/{FMS_INT}/g, "format(cast(");
        sql = sql.replace(/{FME_INT}/g, " as signed), 0)");
        sql = sql.replace(/{FMS_DATE}/g, "date_format(");
        sql = sql.replace(/{FME_DATE}/g, ", '%Y-%m-%d')");
        sql = sql.replace(/{FMS_DATETIME}/g, "date_format(");
        sql = sql.replace(/{FME_DATETIME}/g, ", '%Y-%m-%dT%H:%i:%s')");
        sql = sql.replace(/{FMS_TIME}/g, "cast(cast(");
        sql = sql.replace(/{FME_TIME}/g, " as time) as char)");
        sql = sql.replace(/{JOKER}/g, "'%'");
        sql = sql.replace(/{CUR_DATE}/g, "current_date");
        break;
      case "postgres":
        sql = sql.replace(/{CCS}/g, "");
        sql = sql.replace(/{SEP}/g, "||");
        sql = sql.replace(/{CCE}/g, "");
        sql = sql.replace(/{CAS_TEXT}/g, "cast(");
        sql = sql.replace(/{CAE_TEXT}/g, " as text)");
        sql = sql.replace(/{CAS_INT}/g, "cast("); //cast as integer - start
        sql = sql.replace(/{CAE_INT}/g, " as integer)"); //cast as integer - end
        sql = sql.replace(/{CAS_FLOAT}/g, "cast(");
        sql = sql.replace(/{CAE_FLOAT}/g, " as float)");
        sql = sql.replace(/{CAS_DATE}/g, "cast(");
        sql = sql.replace(/{CASF_DATE}/g, "cast(");
        sql = sql.replace(/{CAE_DATE}/g, " as date)");
        sql = sql.replace(/{CAEF_DATE}/g, " as date)");
        sql = sql.replace(
          /{FMSF_NUMBER}/g,
          "case when rf_number.fieldname is null then 0 else "
        );
        sql = sql.replace(
          /{FMSF_DATE}/g,
          "case when rf_date.fieldname is null then current_date else "
        );
        sql = sql.replace(/{FMEF_CONVERT}/g, " end ");
        sql = sql.replace(/{FMS_FLOAT}/g, "replace(to_char(cast(");
        sql = sql.replace(
          /{FME_FLOAT}/g,
          " as float),'999,999,990.00'),'.00','')"
        );
        sql = sql.replace(/{FMS_INT}/g, "to_char(cast(");
        sql = sql.replace(/{FME_INT}/g, " as integer), '999,999,999')");
        sql = sql.replace(/{FMS_DATE}/g, "to_char(");
        sql = sql.replace(/{FME_DATE}/g, ", 'YYYY-MM-DD')");
        sql = sql.replace(/{FMS_DATETIME}/g, "to_char(");
        sql = sql.replace(/{FME_DATETIME}/g, ", 'YYYY-MM-DD\"T\"HH24:MI:SS')");
        sql = sql.replace(/{FMS_TIME}/g, "substr(cast(cast(");
        sql = sql.replace(/{FME_TIME}/g, " as time) as text), 0, 6)");
        sql = sql.replace(/{JOKER}/g, "chr(37)");
        sql = sql.replace(/{CUR_DATE}/g, "current_date");
        break;
      case "mssql":
        sql = sql.replace(/{CCS}/g, "");
        sql = sql.replace(/{SEP}/g, "+");
        sql = sql.replace(/{CCE}/g, "");
        sql = sql.replace(/{CAS_TEXT}/g, "cast(");
        sql = sql.replace(/{CAE_TEXT}/g, " as nvarchar)");
        sql = sql.replace(/{CAS_INT}/g, "cast(");
        sql = sql.replace(/{CAE_INT}/g, " as int)");
        sql = sql.replace(/{CAS_FLOAT}/g, "cast(");
        sql = sql.replace(/{CAE_FLOAT}/g, " as real)");
        sql = sql.replace(/{CAS_DATE}/g, "cast(");
        sql = sql.replace(/{CASF_DATE}/g, "");
        sql = sql.replace(/{CAE_DATE}/g, " as date)");
        sql = sql.replace(/{CAEF_DATE}/g, "");
        sql = sql.replace(/{FMSF_NUMBER}/g, "");
        sql = sql.replace(/{FMSF_DATE}/g, "");
        sql = sql.replace(/{FMEF_CONVERT}/g, "");
        sql = sql.replace(/{FMS_FLOAT}/g, "replace(convert(varchar,cast("); // mssql 2012+ format(???,'N2')
        sql = sql.replace(/{FME_FLOAT}/g, " as money),1),'.00','')");
        sql = sql.replace(/{FMS_INT}/g, "replace(convert(varchar,cast(");
        sql = sql.replace(/{FME_INT}/g, " as money),1), '.00','')");
        sql = sql.replace(/{FMS_DATE}/g, "convert(varchar(10),");
        sql = sql.replace(/{FME_DATE}/g, ", 120)");
        sql = sql.replace(/{FMS_DATETIME}/g, "convert(varchar(19),");
        sql = sql.replace(/{FME_DATETIME}/g, ", 120)");
        sql = sql.replace(/{FMS_TIME}/g, "SUBSTRING(cast(cast(");
        sql = sql.replace(/{FME_TIME}/g, " as time) as nvarchar),0,6)");
        sql = sql.replace(/{JOKER}/g, "'%'");
        sql = sql.replace(/{CUR_DATE}/g, "cast(GETDATE() as DATE)");
        break;
      default:
        break;
    }
    return sql;
  };

  const sql_decode = (data, key) => {
    let sql = "";
    if (Array.isArray(data)) {
      let sep = " ",
        start_br = "",
        end_br = "";
      if (data.length > 0) {
        if (["select","select_distinct","union_select","order_by","group_by"].includes(key) ||
          data[0].length === 0
        ) {
          sep = ", ";
        }
      }
      data.forEach((element) => {
        if (typeof element === "undefined" || element === null) {
          element = "null";
        }
        if (element.length === 0) {
          if (key !== "set") {
            start_br = "(";
            end_br = ")";
          }
        } else if (
          data.length === 2 &&
          (element === "and" || element === "or")
        ) {
          sql += element + " (";
          end_br = ")";
        } else if (key && data.length === 1 && typeof data[0] === "object") {
          sql += " (" + sql_decode(element, key) + ")";
        } else {
          sql += sep + sql_decode(element, key);
        }
      });
      if (sep === ", ") {
        sql = sql.substr(2);
      }
      if (key && data.includes("on")) {
        sql = key.replace("_", " ") + sql;
      }
      return start_br + sql.toString().trim() + end_br;
    } else if (typeof data === "object") {
      for (let _key in data) {
        if (_key === "inner_join" || _key === "left_join") {
          sql += " " + sql_decode(data[_key], _key);
        } else {
          sql +=
            " " + _key.replace("_", " ") + " " + sql_decode(data[_key], _key);
        }
      }
      return sql;
    } else {
      if (data.includes("?") && key !== "select") {
        prm_count += 1;
        if(engine === "postgres"){
          data = data.replace("?","$" + prm_count);
        }
      }
      return data;
    }
  };

  if (typeof _sql === "string") {
    return {sql: engine_type(_sql), prmCount: prm_count};
  } else {
    return {sql: engine_type(sql_decode(_sql)), prmCount: prm_count};
  }
};

export const saveToDisk = (fileUrl, fileName) => {
  const element = document.createElement("a")
  element.href = fileUrl
  element.download = fileName || fileUrl
  document.body.appendChild(element) // Required for this to work in FireFox
  element.click()
}

export const guid = () => {
  const _p8 = (s) => {
    let p = (Math.random().toString(16)+"000000000").substr(2,8);
    return s ? "-" + p.substr(0,4) + "-" + p.substr(4,4) : p ;
  }
  return _p8() + _p8(true) + _p8(true) + _p8();
}

export const request = async (url, options) => {
  const parseJSON = (response) => {
    if (response.status === 401)
      return { code: 401, message: "Unauthorized" }
    if (response.status === 400)
      return response.json()
    if (response.status === 204 || response.status === 205) {
      return null;
    }
    switch (response.headers.get('content-type').split(";")[0]) {
      case "application/pdf":
      case "application/xml":
      case "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":
        return response.blob()

      case "application/json":
        return response.json()
      
      case "text/plain":
      case "text/csv":
        return response.text()
    
      default:
        return response
    }
  }

  const checkStatus = (response) => {
    if ((response.status >= 200 && response.status < 300) || (response.status === 400) || (response.status === 401)) {
      return response;
    }

    const error = new Error(response.statusText);
    error.response = response;
    throw error;
  }

  const fetch_response = await fetch(url, options);
  const response = checkStatus(fetch_response);
  return parseJSON(response);
}

export const appActions = (data, setData) => {

  const getLangText = (key, defValue) => {
    return getText({ 
      locales: data.session.locales,
      lang: data.current.lang, 
      key: key, defaultValue: defValue 
    })
  }

  const showToast = (params) => {
    Toastify({ 
      toastType: params.type, 
      message: params.message,
      toastTime: (params.autoClose === false) ? 0 : getSetting("toastTime")
    })
  }

  const resultError = (result) => {
    if(result.error){
      setData("error", result.error )
    }
    if(result.error && result.error.message){
      showToast({ type: "error", message: result.error.message })
    } else {
      showToast({ type: "error", 
        message: getLangText("error_internal", "Internal Server Error") })
    }
  }

  const signOut = () => {
    setData("login", { data: null, token: null })
  }

  const requestData = async (path, options, silent) => {
    try {
      if (!silent){
        setData("current", { "request": true })
      }
      let url = (data.session.configServer)?
        data.session.proxy+data.session.apiPath+path : data.login.server+path
      const token = (data.login.data) ? data.login.data.token : options.token || ""
      if (!options.headers)
        options = update(options, {$merge: { headers: {} }})
      options = update(options, { 
        headers: {$merge: { "Content-Type": "application/json" }}
      })
      if(token !== ""){
        options = update(options, { 
          headers: {$merge: { "Authorization": "Bearer "+token }} 
        })
      }
      if (options.data){
        options = update(options, { 
          body: {$set: JSON.stringify(options.data)}
        })
      }
      if(options.query){
        let query = new URLSearchParams();
        for (const key in options.query) {
          query.append(key, options.query[key])
        }
        url += "?" + query.toString()
      }
      
      const result = await request(url, options)
      if (!silent) {
        setData("current", { "request": false })
      }
      if(result && result.code){
        if(result.code === 401){
          signOut()
        }
        return { error: { message: result.message }, data: null }
      }
      return result
    } catch (err) {
      if(!silent){
        setData("current", { "request": false })
      }
      return { error: { message: err.message }, data: null }
    }
  }

  const getAuditFilter = (nervatype, transtype) => {
    let retvalue = ["all",1]; let audit;
    switch (nervatype) {
      case "trans":
      case "menu":
        audit = data.login.data.audit.filter((audit)=> {
          return ((audit.nervatypeName === nervatype) && (audit.subtypeName === transtype))
        })[0]
        break;
      case "report":
        audit = data.login.data.audit.filter((audit)=> {
          return ((audit.nervatypeName === nervatype) && (audit.subtype === transtype))
        })[0]
        break;
      default:
        audit = data.login.data.audit.filter((audit)=> {
          return (audit.nervatypeName === nervatype)
        })[0]
    }
    if (typeof audit !== "undefined") {
      retvalue = [audit.inputfilterName, audit.supervisor];}
    return retvalue;
  }

  const createHistory = async (ctype) => {
    let history = update({}, {$set: {
      datetime: formatISO(new Date()),
      type: ctype, 
      type_title: getLangText["label_"+ctype],
      ntype: data.edit.current.type,
      transtype: data.edit.current.transtype || "",
      id: data.edit.current.item.id
    }})
    let title = (history.ntype === "trans") ?
      data.edit.template.options.title+" | "+data.edit.current.item[data.edit.template.options.title_field] :
      data.edit.template.options.title
    if ((history.ntype !== "trans") && (typeof data.edit.template.options.title_field !== "undefined")){
      title += " | "+data.edit.current.item[data.edit.template.options.title_field]
    }
    history = update(history, {$merge: {
      title: title
    }})
    let bookmark = update(data.bookmark, {})
    let userconfig = {}
    if (bookmark.history) {
      userconfig = update(bookmark.history, {$merge: {
        cfgroup: formatISO(new Date())
      }})
      let history_values = JSON.parse(userconfig.cfvalue);
      history_values.unshift(history)
      if (history_values.length > parseInt(getSetting("history"),10)) {
        history_values = history_values.slice(0, parseInt(getSetting("history"),10))
      }
      userconfig = update(userconfig, {$merge: {
        cfname: history_values.length,
        cfvalue: JSON.stringify(history_values)
      }})
    } else {
      userconfig = update(userconfig, {$merge: {
        employee_id: data.login.data.employee.id,
        section: "history",
        cfgroup: formatISO(new Date()),
        cfname: 1,
        cfvalue: JSON.stringify([history])
      }})
    }
    const options = { method: "POST", data: [userconfig] }
    const result = await requestData("/ui_userconfig", options)
    if(result.error){
      return resultError(result)
    }
    setData("bookmark", { history: userconfig})
  }

  const loadBookmark = async (params) => {
    const result = await requestData("/ui_userconfig", {
      token: params.token,
      query: {
        filter: "employee_id;==;"+params.user_id
      }
    })
    if(result.error){
      resultError(result)
      return null
    }
    setData("bookmark", { 
      bookmark: result.filter(item => (item.section === "bookmark")),
      history: result.filter(item => (item.section === "history"))[0]||null
    }, ()=>{
      if(params.callback){
        params.callback()
      }
    })
  }

  const saveBookmark = (params) => {
    const InputBox = lazy(() => import('components/Modal/InputBox'));
    setData("current", { modalForm: 
      <InputBox 
        title={getLangText("msg_bookmark_new")}
        message={getLangText("msg_bookmark_name")}
        value={(params[0] === "browser") ? params[1] : data.edit.current.item[params[2]]}
        showValue={true}
        labelOK={getLangText("msg_ok")}
        labelCancel={getLangText("msg_cancel")}
        onCancel={() => {
          setData("current", { modalForm: null })
        }}
        onOK={(value) => {
          setData("current", { modalForm: null }, async () => {
            if (value !== "") {
              let userconfig = {
                employee_id: data.login.data.employee.id,
                section: "bookmark",
                cfgroup: params[0],
              }
              if((params[0]) === "browser"){
                userconfig = update(userconfig, {$merge: {
                  cfname: value,
                  cfvalue: JSON.stringify({
                    date: formatISO(new Date(), { representation: 'date' }),
                    vkey: data.search.vkey,
                    view: data.search.view,
                    filters: data.search.filters[data.search.view],
                    columns: data.search.columns[data.search.view]
                  })
                }})
              } else {
                userconfig = update(userconfig, {$merge: {
                  cfname: value,
                  cfvalue: JSON.stringify({
                    date: formatISO(new Date(), { representation: 'date' }),
                    ntype: data.edit.current.type,
                    transtype: data.edit.current.transtype,
                    id: data.edit.current.item.id,
                    info: (data.edit.current.type === "trans") 
                      ? (data.edit.dataset.trans[0].custname !== null) 
                        ? data.edit.dataset.trans[0].custname 
                        : data.edit.current.item.transdate 
                      : data.edit.current.item[params[3]]
                  })
                }})
              }
  
              const options = { method: "POST", data: [userconfig] }
              const result = await requestData("/ui_userconfig", options)
              if(result.error){
                return resultError(result)
              }
              loadBookmark({user_id: data.login.data.employee.id})
            }
          })
        }}
      />, side: "hide"
    })
  }

  const showHelp = (key) => {
    const element = document.createElement("a")
    element.setAttribute("href", data.session.helpPage+key)
    element.setAttribute("target", "_blank")
    document.body.appendChild(element)
    element.click()
  }

  const getDataFilter = (type, _where, view) => {
    if(type === "transitem"){
      if (getAuditFilter("trans", "offer")[0] === "disabled") {
        _where = _where.concat(["and", ["tg.groupvalue", "<>", "'offer'"]]);
      }
      if (getAuditFilter("trans", "order")[0] === "disabled") {
        _where = _where.concat(["and", ["tg.groupvalue", "<>", "'order'"]]);
      }
      if (getAuditFilter("trans", "worksheet")[0] === "disabled") {
        _where = _where.concat(["and", ["tg.groupvalue", "<>", "'worksheet'"]]);
      }
      if (getAuditFilter("trans", "rent")[0] === "disabled") {
        _where = _where.concat(["and", ["tg.groupvalue", "<>", "'rent'"]]);
      }
      if (getAuditFilter("trans", "invoice")[0] === "disabled") {
        _where = _where.concat(["and", ["tg.groupvalue", "<>", "'invoice'"]]);
      }
    }
    if(type === "transpayment"){
      if (getAuditFilter("trans", "bank")[0] === "disabled") {
        _where = _where.concat(["and", ["tg.groupvalue", "<>", "'bank'"]]);
      }
      if (getAuditFilter("trans", "cash")[0] === "disabled") {
        _where = _where.concat(["and", ["tg.groupvalue", "<>", "'cash'"]]);
      }
    }
    if((type === "transmovement") && (view !== "InventoryView")){
      if (getAuditFilter("trans", "delivery")[0] === "disabled") {
        _where = _where.concat(["and", ["tg.groupvalue", "<>", "'delivery'"]]);
      }
      if (getAuditFilter("trans", "inventory")[0] === "disabled") {
        _where = _where.concat(["and", ["tg.groupvalue", "<>", "'inventory'"]]);
      }
    }
    return _where;
  }

  const getUserFilter = (type) => {
    let filter = { params: [], where: []}
    if(data.login.data.transfilterName === "usergroup"){
      if (typeof UserFilter.usergroup_filter[type] !== "undefined") {
        filter.where = ["and", UserFilter.usergroup_filter[type]]
        filter.params = [data.login.data.employee.usergroup]
      }
    }
    if(data.login.data.transfilterName === "own"){
      if (typeof UserFilter.employee_filter[type] !== "undefined") {
        filter.where = ["and", UserFilter.employee_filter[type]]
        filter.params = [data.login.data.employee.id]
      }
    }
    return filter;
  }

  const quickSearch = async (qview, qfilter) => {
    const { Quick } = await import('containers/Controller/Quick');
    const query = Quick({ getText: getLangText })[qview](String(data.login.data.employee.usergroup))
    let _sql = update({}, {$set: query.sql})
    let params = []; let _where = []
    if(qfilter !== ""){
      const filter = `{CCS}{JOKER}{SEP}lower(?){SEP}{JOKER}{CCE} `
      query.columns.forEach((column, index) => {
        _where.push([((index!==0)?"or":""),[`lower(${column[1]})`,"like", filter]])
        params.push(qfilter)
      });
      _where = ["and",[_where]]
    }
    _where = getDataFilter(qview, _where)
    if(_where.length > 0){
      _sql = update(_sql, { where: {$push: [_where]}})
    }

    let userFilter = getUserFilter(qview)
    if(userFilter.where.length > 0){
      _sql = update(_sql, { where: {$push: userFilter.where}})
      params = params.concat(userFilter.params)
    }
    
    let views = [
      { key: "result",
        text: getSql(data.login.data.engine, _sql).sql,
        values: params 
      }
    ]
    let options = { method: "POST", data: views }
    return await requestData("/view", options)
  }

  const onSelector = async (selectorType, selectorFilter, setSelector) => {
    const { Quick } = await import('containers/Controller/Quick');
    const Selector = lazy(() => import('components/Modal/Selector'));
    let formProps = {
      view: selectorType, 
      columns: Quick({ getText: getLangText })[selectorType]().columns,
      result: [],
      filter: selectorFilter,
      getText: getLangText, 
      onClose: ()=>setData("current", { modalForm: null }),
      onSelect: (row, filter) => {
        setData("current", { modalForm: null }, ()=>{
          setSelector(row, filter)
        })
      },
      onSearch: async (filter)=>{
        const view = await quickSearch(selectorType, filter)
        if(view.error){
          return resultError(view)
        }
        formProps.result = view.result
        setData("current", { modalForm: <Selector {...formProps} /> })
      }
    }
    setData("current", { modalForm: <Selector {...formProps} /> })
    return formProps
  }

  return {
    createHistory: createHistory,
    getAuditFilter: getAuditFilter,
    getDataFilter: getDataFilter,
    getText: getLangText,
    getUserFilter: getUserFilter,
    loadBookmark: loadBookmark,
    onSelector: onSelector,
    quickSearch: quickSearch,
    requestData: requestData,
    resultError: resultError,
    saveBookmark: saveBookmark,
    showHelp: showHelp,
    showToast: showToast,
    signOut: signOut,
  }
}