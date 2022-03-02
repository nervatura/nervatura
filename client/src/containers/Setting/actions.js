import update from 'immutability-helper';

import { appActions, getSql } from 'containers/App/actions'
import { Sql } from 'containers/Controller/Sql'
import dataset from 'containers/Controller/Dataset'
import { Forms } from 'containers/Controller/Forms'
import { InitItem, Validator } from 'containers/Controller/Validator'
import InputBox from 'components/Modal/InputBox'
import Audit from 'components/Modal/Audit'
import Menu from 'components/Modal/Menu'
import { getSetting } from 'config/app'

export const settingActions = (data, setData) => {
  const app = appActions(data, setData)
  const validator = Validator(data, setData)
  const initItem = InitItem(data, setData)

  const forms = Forms({ getText: app.getText })
  const sql = Sql({ getText: app.getText })

  const tableValues = (type, item) => {
    let values = {}
    const baseValues = initItem({tablename: type, 
      dataset: data.setting.dataset, current: data.setting.current})
    for (const key in item) {
      if (baseValues.hasOwnProperty(key)) {
        values[key] = item[key]
      }
    }
    return values
  }

  const setSettingForm = (id, isetting) => {
    let setting = update(isetting||data.setting, {$merge: {
      current: {}
    }})
    if (id!==null) {
      let item = setting.dataset[setting.type+"_view"].filter(item => (item.id === parseInt(id,10)))[0]
      setting = update(setting, {$merge: {
        dirty: false,
        current: { form: item }
      }})
      if(setting.ntype === "fieldvalue"){
        setting = update(setting, { current: {$merge: {
          fieldvalue: {
            rowtype: 'fieldvalue',
            id: item.id, 
            //name: 'fieldvalue_value', 
            fieldname: item.fieldname, 
            fieldvalue_value: item.value, 
            fieldvalue_notes: item.notes || '',
            label: item.lslabel, 
            description: (item.fieldtype === 'valuelist') ? item.valuelist.split("|") : item.value, 
            disabled: 'false',
            fieldtype: item.fieldtype, 
            datatype: (item.fieldtype === 'urlink') ? 'text' : item.fieldtype
          }
        }}})
      }
    } else {
      setting = update(setting, {$merge: {
        dirty: ((setting.audit==="all") || (setting.audit==="update")) ? true : false,
        current: { form: initItem({tablename: setting.type, 
          dataset: setting.dataset, current: setting.current}) }
      }})
    }
    if((setting.type === "usergroup") && (setting.current.form.transfilter === null)){
      setting.current.form.transfilter = setting.dataset.transfilter.filter(
        item => (item.groupvalue === "all"))[0].id
    }
    setting = update(setting, {current : {$merge: {
      template: forms[setting.type](setting.current.form, setting)
    }}})
    setting = update(setting, {$merge: {
      panel: setting.current.template.options.panel
    }})
    if (setting.audit === "readonly") {
      setting.panel.save = false
    }
    if (setting.audit !== "all") {
      setting.panel.delete = false;
      setting.panel.new = false;
    }
    setData("setting", setting)
  }

  const setPasswordForm = (username) => {
    const data = {
      username: username, password_1: "", password_2: "" 
    }
    const form = forms.password(data)
    let setting = {
      type: "password", 
      dataset: {}, 
      current: {
        form: data,
        template: form
      },
      panel: form.options.panel,
      caption: form.options.title,
      icon: form.options.icon,
      filter: "", 
      result: [], 
      dirty: false,
      audit: "all",
      view: {
        type: "password",
        result: []
      }
    }
    setData("setting", setting)
    setData("current", { module: "setting", side: "hide" })
  }

  const changePassword = async () => {
    const { username, password_1, password_2 } = data.setting.current.form
    if (username === "" || username === null) {
      return app.showToast({ type: "error",
        title: app.getText("msg_warning"), 
        message: app.getText("ms_password_username") })
    }
    if (password_1 !== password_2) {
      return app.showToast({ type: "error",
        title: app.getText("msg_warning"), 
        message: app.getText("ms_password_pswerr") })
    }
    const options = {
      method: "POST",
      data: { 
        username: username,
        password: password_1,
        confirm: password_2
      } 
    }
    const result = await app.requestData("/auth/password", options)
    if(result && result.error){
      return app.resultError(result)
    }
    app.showToast({ type: "success",
      title: app.getText("msg_warning"), 
      message: app.getText("msg_password_ok") })
  }

  const setProgramForm = () => {
    const template = forms.program()
    let setting = update({}, {$set: {
      type: "program", 
      dataset: {}, 
      current: {
        form: {
          paginationPage: getSetting("paginationPage"),
          history: getSetting("history"),
          page_size: getSetting("page_size"),
          export_sep: getSetting("export_sep"),
          decimal_sep: getSetting("decimal_sep"),
          dateFormat: getSetting("dateFormat"),
          calendar: getSetting("calendar")
        },
        template: template
      }, 
      filter: "", 
      result: [], 
      dirty: false,
      panel: null,
      caption: template.options.title,
      icon: template.options.icon,
      view: {
        type: "program",
        result: []
      }
    }})
    setData("setting", setting)
    setData("current", { module: "setting" })
  }

  const setSettingData = (options) => {
    const form = forms[options.type]()
    const audit = app.getAuditFilter(options.type)[0]
    let setting = update(options, {$merge: {
      current: (options.type === "log") ? {
        form: initItem({ tablename: "log" }),
        template: form
      } : null,
      panel: null,
      filter: "",
      result: [],
      ntype: form.options.data,
      caption: form.options.title,
      icon: form.options.icon,
      view: {
        type: form.view.setting.type,
        result: options.dataset[options.type+"_view"],
        fields: (form.view.setting.fields) ? form.view.setting.fields : null
      },
      actions: {
        new: (audit !== "all") ? null : form.view.setting.actions.new, 
        edit: form.view.setting.actions.edit, 
        delete: (audit !== "all") ? null : form.view.setting.actions.delete
      },
      audit: audit,
      template: null,
      page: (data.setting.type === options.type) ? data.setting.page || 0 : 0
    }})
    setData("setting", setting)
    setData("current", { module: "setting", side: "hide" })
    if(((options.type === "usergroup") || (options.type === "ui_menu")) && options.id) {
      setViewActions({action: "editItem", setting: setting}, { id: options.id })
    } else if(typeof options.id !== "undefined"){
      setSettingForm(options.id, setting)
    }
  }

  const loadSetting = async (options) => {
    let setting = update(options, {$merge: {
      dataset: {},
      dirty: false
    }})
    if(dataset[setting.type]){
      let views = []
      dataset[setting.type]().forEach(info => {
        let _sql = {}
        if(info.infoType === "table"){
          _sql = { select:["*"], from: info.classAlias }
          if(info.where){
            _sql.where = info.where
          }
          if(info.order){
            _sql.order_by = info.order
          }
        } else {
          if (typeof sql[setting.type][info.infoName] !== "undefined") {
            _sql = sql[setting.type][info.infoName](setting.type)
          } else {
            _sql = sql["all"][info.infoName](setting.type)
          }
        }
        const sqlInfo = getSql(data.login.data.engine, _sql)
        if( (setting.id !== null) || (sqlInfo.prmCount === 0)){
          views = update(views, {$push: [{
            key: info.infoName,
            text: sqlInfo.sql,
            values: ((sqlInfo.prmCount>0) && (setting.id !== null)) ? Array(sqlInfo.prmCount).fill(setting.id) : []
          }]})
        } else {
          setting = update(setting, { dataset: {$merge: {
            [info.infoName]: []
          }}})
        }
      })

      let params = { method: "POST", data: views }
      let view = await app.requestData("/view", params)
      if(view.error){
        return app.resultError(view)
      }
      setting = update(setting, {dataset: {
        $merge: view
      }})
      if ((setting.type === "template") && (typeof setting.id !== "undefined")) {
        return setData("current", { module: "template", content: setting }) 
      }
    }
    setSettingData(setting)
  }

  const saveSetting = async () => {
    let setting = update(data.setting, {})

    let values = tableValues(setting.ntype, setting.current.form)
    values = await validator(setting.ntype, values)
    if(values.error){
      return app.resultError(values)
    }

    let result = await app.requestData("/"+setting.ntype, { method: "POST", data: [values] })
    if(result.error){
      app.resultError(result)
      return null
    }
    if (setting.current.form.id === null) {
      setting = update(setting, {current: { form: {$merge: {
        id: result[0]
      }}}})
    }
    setting = update(setting, {$merge: {
      dirty: false
    }})
    if (setting.type === "usergroup") {
      const transfilter_name = data.login.data.groups.filter(
        item => (item.id === setting.current.form.transfilter))[0].groupvalue
      if((setting.current.form.translink !== null) || (transfilter_name !== "all")){
        const link = update(
          initItem({tablename: "link", dataset: setting.dataset, current: setting.current}),
          {$merge: {
            id: setting.current.form.translink,
            nervatype_1: data.login.data.groups.filter(
              (item)=>((item.groupname === "nervatype") && (item.groupvalue === "groups")))[0].id,
            ref_id_1: setting.current.form.id,
            nervatype_2: data.login.data.groups.filter(
              (item)=>((item.groupname === "nervatype") && (item.groupvalue === "groups")))[0].id,
            ref_id_2: setting.current.form.transfilter
          }}
        )
        result = await app.requestData("/link", { method: "POST", data: [link] })
        if(result.error){
          app.resultError(result)
          return null
        }
      }
    }
    return setting
  }

  const deleteSetting = (item) => {
    const deleteData = async () => {
      const path = (data.setting.type === "usergroup") ? "/groups" : "/"+data.setting.type
      const result = await app.requestData(path, 
        { method: "DELETE", query: { id: item.id } })
      if(result && result.error){
        return app.resultError(result)
      }
      loadSetting({type: data.setting.type})
    }
    setData("current", { modalForm: 
      <InputBox 
        title={app.getText("msg_warning")}
        message={app.getText("msg_delete_text")}
        infoText={app.getText("msg_delete_info")}
        labelOK={app.getText("msg_ok")}
        labelCancel={app.getText("msg_cancel")}
        onCancel={() => {
          setData("current", { modalForm: null })
        }}
        onOK={(value) => {
          setData("current", { modalForm: null }, async () => {
            if (typeof sql[data.setting.type]["delete_state"] !== "undefined") {
              const sqlInfo = getSql(data.login.data.engine, sql[data.setting.type]["delete_state"]())
              const params = { 
                method: "POST", 
                data: [{ 
                  key: "state",
                  text: sqlInfo.sql,
                  values: Array(sqlInfo.prmCount).fill(item.id)
                }]
              }
              let view = await app.requestData("/view", params)
              if(view.error){
                return app.resultError(view)
              }
              if (view.state[0].sco > 0) {
                app.showToast({ type: "error",
                  title: app.getText("msg_warning"), 
                  message: app.getText("msg_integrity_err") })
              } else {
                deleteData()
              }
            } else {
              deleteData()
            }
          })
        }}
      />,
      side: "hide"
    })
  }

  const loadLog = async () => {
    const nervatype_filter = {
      "customer": {
        inner_join: ["customer c","on",["l.ref_id","=","c.id"]],
        select: "c.custnumber as refnumber"
      },
      "employee": {
        inner_join: ["employee em","on",["l.ref_id","=","em.id"]],
        select: "em.empnumber as refnumber"
      },
      "event": {
        inner_join: ["event ev","on",["l.ref_id","=","ev.id"]],
        select: "ev.calnumber as refnumber"
      },
      "place": {
        inner_join: ["place p","on",["l.ref_id","=","p.id"]],
        select: "p.planumber as refnumber"
      },
      "product": {
        inner_join: ["product p","on",["l.ref_id","=","p.id"]],
        select: "p.partnumber as refnumber"
      },
      "project": {
        inner_join: ["project p","on",["l.ref_id","=","p.id"]],
        select: "p.pronumber as refnumber"
      },
      "tool": {
        inner_join: ["tool t","on",["l.ref_id","=","t.id"]],
        select: "t.serial as refnumber"
      },
      "trans": {
        inner_join: ["trans t","on",["l.ref_id","=","t.id"]],
        select: "t.transnumber as refnumber"
      }
    }
    let setting = update(data.setting, {view: {$merge: {
      result: []
    }}})
    let _log = sql.log.result()
    let paramList = []
    if (!["login", "logout"].includes(setting.current.form.logstate)){
      _log.inner_join.push(
        ["groups nt","on",[["l.nervatype","=","nt.id"],
          ["and","nt.groupvalue","=","?"]]])
      paramList.push(setting.current.form.nervatype)
      if(!nervatype_filter[setting.current.form.nervatype]){
        return app.showToast({ type: "error",
          title: app.getText("msg_warning"), 
          message: app.getText("msg_required")+" "+app.getText("log_nervatype") })
      }
      _log.inner_join.push(nervatype_filter[setting.current.form.nervatype].inner_join)
      _log.select[4] = nervatype_filter[setting.current.form.nervatype].select
    }

    _log.where.push(["ls.groupvalue","=","?"])
    paramList.push(setting.current.form.logstate)
    
    if (setting.current.form.empnumber && setting.current.form.empnumber !== "") {
      _log.where.push(["and","lower(e.empnumber)","like","{CCS}{JOKER}{SEP}lower('"+
        setting.current.form.empnumber+"'){SEP}{JOKER}{CCE}"]);
    }
    
    if (setting.current.form.fromdate && setting.current.form.fromdate !== "") {
      _log.where.push(["and","{FMS_DATE}l.crdate{FME_DATE}",">=","?"])
      paramList.push(setting.current.form.fromdate)
    }
    
    if (setting.current.form.todate && setting.current.form.todate !== "") {
      _log.where.push(["and","{FMS_DATE}l.crdate{FME_DATE}","<=","?"]);
      paramList.push(setting.current.form.todate)
    }

    const params = { 
      method: "POST", 
      data: [{ 
        key: "log",
        text: getSql(data.login.data.engine, _log).sql,
        values: paramList
      }]
    }
    let view = await app.requestData("/view", params)
    if(view.error){
      return app.resultError(view)
    }
    setting = update(setting, {view: {$merge: {
      result: view.log
    }}})
    setData("setting", setting)
  }

  const checkSetting = (options, cbKeyTrue) => {
    const nextKeys = {
      LOAD_SETTING: ()=>loadSetting(options),
      //SETTING_FORM: ()=>setSettingForm(options.id),
      PASSWORD_FORM: ()=>setPasswordForm(options.username)
    }
    if (data.setting.dirty === true) {
      return setData("current", { modalForm: 
        <InputBox 
          title={app.getText("msg_warning")}
          message={app.getText("msg_dirty_text")}
          infoText={app.getText("msg_dirty_info")}
          labelOK={app.getText("msg_save")}
          labelCancel={app.getText("msg_cancel")}
          onCancel={() => {
            setData("current", { modalForm: null }, ()=>{
              setData(data.current.module, { dirty: false }, ()=>{
                nextKeys[cbKeyTrue]()
              })
            })
          }}
          onOK={(value) => {
            setData("current", { modalForm: null }, async ()=>{
              const setting = await saveSetting()
              if(setting){
                return setData("setting", setting, ()=>{
                  nextKeys[cbKeyTrue]()
                })
              }
            })
          }}
        />,
        side: "hide"
      })
    }
    nextKeys[cbKeyTrue]()
  }

  const setViewActions = async (params, _row) => {
    const row = _row || {}
    switch (params.action) {
      case "newItem":
        checkSetting({ type: data.setting.type, id: null }, 'LOAD_SETTING')
        break;

      case "editItem":
        switch (data.setting.type) {
          case "template":
            checkSetting({ type: data.setting.type, id: row.id }, 'LOAD_SETTING')
            break;
          
          case "place":
            setData("current", { module: "edit", content: {ntype: data.setting.type, ttype: null, 
              id: row.id || null} })
            break;
         
          case "usergroup":
            const audit_options = { 
              method: "POST", 
              data: [{ 
                key: "audit",
                text: getSql(data.login.data.engine, sql.usergroup.audit()).sql,
                values: [row.id]
              }]
            }
            let audit_view = await app.requestData("/view", audit_options)
            if(audit_view.error){
              return app.resultError(audit_view)
            }
            const audit_setting = update(params.setting||data.setting, {dataset: {$merge: {
              audit: audit_view.audit
            }}})
            setSettingForm(row.id, audit_setting)
            break;
          
          case "ui_menu":
            const menufields_options = { 
              method: "POST", 
              data: [{ 
                key: "menufields",
                text: getSql(data.login.data.engine, sql.ui_menu.ui_menufields()).sql,
                values: [row.id]
              }]
            }
            const menufields_view = await app.requestData("/view", menufields_options)
            if(menufields_view.error){
              return app.resultError(menufields_view)
            }
            const menufields_setting = update(params.setting||data.setting, {dataset: {$merge: {
              ui_menufields: menufields_view.menufields
            }}})
            setSettingForm(row.id, menufields_setting)
            break;

          default:
            setSettingForm(row.id)
            break;
        }
        break;

      case "deleteItem":
        deleteSetting(row)
        break;

      case "editAudit":
        let audit = row
        if(!audit.id){
          audit = update(
            initItem({tablename: "link", dataset: data.setting.dataset, current: data.setting.current}),
            {$merge: {
              usergroup: data.setting.current.form.id,
              nervatype: data.setting.dataset.nervatype.filter(
                (item)=>(item.groupvalue === "customer"))[0].id,
              inputfilter: data.setting.dataset.inputfilter.filter(
                (item)=>(item.groupvalue === "all"))[0].id
            }}
          )
        }
        setData("current", { modalForm: 
          <Audit 
            idKey={audit.id} usergroup={audit.usergroup}
            nervatype={audit.nervatype} subtype={audit.subtype}
            inputfilter={audit.inputfilter} supervisor={audit.supervisor}
            typeOptions={data.setting.dataset.nervatype.map(group => { 
              return { value: String(group.id), text: group.groupvalue } 
            })}
            subtypeOptions={data.setting.dataset.transtype.map(group => { 
              return { value: String(group.id), text: group.groupvalue, type: "trans" } }).concat(
                data.setting.dataset.reportkey.map(report => { 
                  return { value: String(report.id), text: report.reportkey, type: "report" } }),
                data.setting.dataset.menukey.map(menu => { 
                  return { value: String(menu.id), text: menu.menukey, type: "menu" } })
              )
            }
            inputfilterOptions={data.setting.dataset.inputfilter.map(group => { 
              return { value: String(group.id), text: group.groupvalue } 
            })}
            getText={app.getText}
            onClose={()=>setData("current", { modalForm: null })}
            onAudit={async (audit) => {
              let result = await app.requestData("/ui_audit", { 
                method: "POST", data: [tableValues("audit", audit)] })
              if(result.error){
                return app.resultError(result)
              }
              setData("current", { modalForm: null }, ()=>{
                setViewActions({action: "editItem"}, data.setting.current.form)
              })
            }}
          /> 
        })
        break;

      case "editMenuField":
        let menufields = row
        if(!menufields.id){
          menufields = update(
            initItem({tablename: "ui_menufields", dataset: data.setting.dataset, current: data.setting.current}),
            {$merge: {
              menu_id: data.setting.current.form.id,
              fieldtype: data.setting.dataset.fieldtype.filter(
                (item)=>(item.groupvalue === "string"))[0].id,
              orderby: data.setting.dataset.ui_menufields.length
            }}
          )
        }
        setData("current", { modalForm: 
          <Menu 
            idKey={menufields.id} menu_id={menufields.menu_id}
            fieldname={menufields.fieldname} description={menufields.description}
            fieldtype={menufields.fieldtype} orderby={menufields.orderby}
            fieldtypeOptions={data.setting.dataset.fieldtype.map(group => { 
              return { value: String(group.id), text: group.groupvalue } 
            })}
            getText={app.getText}
            onClose={()=>setData("current", { modalForm: null })}
            onMenu={async (menufields) => {
              let result = await app.requestData("/ui_menufields", { 
                method: "POST", data: [tableValues("ui_menufields", menufields)] })
              if(result.error){
                return app.resultError(result)
              }
              setData("current", { modalForm: null }, ()=>{
                setViewActions({action: "editItem"}, data.setting.current.form)
              })
            }}
          /> 
        })
        break;
      
      case "deleteItemRow":
        setData("current", { modalForm: 
          <InputBox 
            title={app.getText("msg_warning")}
            message={app.getText("msg_delete_text")}
            infoText={app.getText("msg_delete_info")}
            labelOK={app.getText("msg_ok")}
            labelCancel={app.getText("msg_cancel")}
            onCancel={() => {
              setData("current", { modalForm: null })
            }}
            onOK={(value) => {
              setData("current", { modalForm: null }, async ()=>{
                const result = await app.requestData(
                  "/"+params.table, { method: "DELETE", query: { id: row.id } })
                if(result && result.error){
                  return app.resultError(result)
                }
                setViewActions({action: "editItem"}, data.setting.current.form)
              })
            }}
          /> 
        })
        break;

      default:
        break;
    }
  }

  const editItem = (options) => {
    let settings = update({}, {$set: data.setting})
    if(settings.type === "program"){
      settings = update(settings, {current: {form: {$merge: {
        [options.name]: options.value
      }}}})
      localStorage.setItem(options.name, options.value)
    } else if(options.name === "log_search"){
      loadLog()
    } else {
      if((settings.audit==="all") || (settings.audit==="update")){
        settings = update(settings, {$merge: {
          dirty: true
        }})
      }
      if((options.name === "fieldvalue_value") || (options.name === "fieldvalue_notes")){
        settings = update(settings, {current: {fieldvalue: {$merge: {
          [options.name]: options.value
        }}}})
        settings = update(settings, {current: {form: {$merge: {
          [options.name.split("_")[1]]: options.value.toString()
        }}}})
      } else {
        settings = update(settings, {current: {form: {$merge: {
          [options.name]: options.value
        }}}})
      }
    }
    setData("setting", settings)
  }

  const setPassword = (username) =>{
    if(!username && data.edit.current){
      username = data.edit.dataset[data.edit.current.type][0].username
    }
    checkSetting({ username: username }, "PASSWORD_FORM")
  }

  const settingSave = () => {
    setData("current", { side: "hide" }, async ()=>{
      if(data.setting.type === "password"){
        changePassword()
      } else {
        const result = await saveSetting()
        if(result){
          loadSetting({type: result.type, id: result.current.form.id})
        }
      }
    })
  }

  const settingBack = (back_type) => {
    if(data.setting.type === "password"){
      return setData(data.current.module, { group_key: "group_admin" }, ()=>{
        loadSetting({ type: 'setting' })
      })
    }
    checkSetting({ type: back_type || data.setting.type }, 'LOAD_SETTING')
  }

  return {
    changePassword: changePassword,
    checkSetting: checkSetting,
    deleteSetting: deleteSetting,
    editItem: editItem,
    loadLog: loadLog,
    loadSetting: loadSetting,
    saveSetting: saveSetting,
    setPasswordForm: setPasswordForm,
    setPassword: setPassword,
    setProgramForm: setProgramForm,
    setSettingForm: setSettingForm,
    settingBack: settingBack,
    settingSave: settingSave,
    setSettingData: setSettingData,
    setViewActions: setViewActions,
    tableValues: tableValues,
  }
}