import { APP_MODULE, SETTING_EVENT, SIDE_EVENT, SIDE_VISIBILITY, MODAL_EVENT, TOAST_TYPE, ACTION_EVENT } from '../config/enums.js'

export class SettingController {
  constructor(host) {
    this.host = host
    this.app = host.app
    this.store = host.app.store
    this.module = {}

    this.changePassword = this.changePassword.bind(this)
    this.checkSetting = this.checkSetting.bind(this)
    this.deleteSetting = this.deleteSetting.bind(this)
    this.editItem = this.editItem.bind(this)
    this.loadLog = this.loadLog.bind(this)
    this.loadSetting = this.loadSetting.bind(this)
    this.onSettingEvent = this.onSettingEvent.bind(this)
    this.onSideEvent = this.onSideEvent.bind(this)
    this.saveSetting = this.saveSetting.bind(this)
    this.setProgramForm = this.setProgramForm.bind(this)
    this.setSettingData = this.setSettingData.bind(this)
    this.setSettingForm = this.setSettingForm.bind(this)
    this.setFormActions = this.setFormActions.bind(this)
    this.tableValues = this.tableValues.bind(this)
    host.addController(this);
  }

  setModule(moduleRef){
    this.module = moduleRef
  }

  async changePassword() {
    const { requestData, resultError, showToast, msg } = this.app
    const { data } = this.store
    const { username, password_1, password_2 } = data[APP_MODULE.SETTING].current.form
    if (username === "" || username === null) {
      return showToast(TOAST_TYPE.ERROR, msg("", { id: "ms_password_username" }))
    }
    if (password_1 !== password_2) {
      return showToast(TOAST_TYPE.ERROR, msg("", { id: "ms_password_pswerr" }))
    }
    const options = {
      method: "POST",
      data: { 
        username,
        password: password_1,
        confirm: password_2
      } 
    }
    const result = await requestData("/auth/password", options)
    if(result && result.error){
      return resultError(result)
    }
    return showToast(TOAST_TYPE.SUCCESS, msg("", { id: "msg_password_ok" }))
  }

  checkSetting(options, cbKeyTrue) {
    const { inputBox } = this.host
    const { msg } = this.app
    const { setData, data } = this.store
    const nextKeys = {
      [SIDE_EVENT.LOAD_SETTING]: ()=>this.loadSetting(options),
      // SETTING_FORM: ()=>setSettingForm(options.id),
      [SIDE_EVENT.PASSWORD_FORM]: ()=>this.setPasswordForm(options.username)
    }
    if (data[APP_MODULE.SETTING].dirty === true) {
        const  modalForm = inputBox({ 
          title: msg("", { id: "msg_warning" }),
          message: msg("", { id: "msg_dirty_text" }),
          infoText: msg("", { id: "msg_delete_info" }),
          defaultOK: true,
          onEvent: {
            onModalEvent: async (modalResult) => {
              setData("current", { modalForm: null })
              if (modalResult.key === MODAL_EVENT.OK) {
                const setting = await this.saveSetting()
                if(setting){
                  setData(APP_MODULE.SETTING, setting)
                  nextKeys[cbKeyTrue]()
                }
              } else {
                nextKeys[cbKeyTrue]()
              }
            }
          }
        })
        return setData("current", { modalForm, side: SIDE_VISIBILITY.HIDE })
    }
    return nextKeys[cbKeyTrue]()
  }

  deleteSetting(item) {
    const { inputBox } = this.host
    const { data, setData } = this.store
    const { sql } = this.app.modules
    const { resultError, requestData, getSql, showToast, msg } = this.app
    
    const deleteData = async () => {
      const path = (data[APP_MODULE.SETTING].type === "usergroup") ? "/groups" : `/${data[APP_MODULE.SETTING].type}`
      const result = await requestData(path, 
        { method: "DELETE", query: { id: item.id } })
      if(result && result.error){
        resultError(result)
      } else {
        this.loadSetting({type: data[APP_MODULE.SETTING].type})
      }
    }
    const  modalForm = inputBox({ 
      title: msg("", { id: "msg_warning" }),
      message: msg("", { id: "msg_delete_text" }),
      infoText: msg("", { id: "msg_delete_info" }),
      onEvent: {
        onModalEvent: async (modalResult) => {
          setData("current", { modalForm: null })
          if (modalResult.key === MODAL_EVENT.OK) {
            if (typeof sql[data[APP_MODULE.SETTING].type].delete_state !== "undefined") {
              const sqlInfo = getSql(data[APP_MODULE.LOGIN].data.engine, sql[data[APP_MODULE.SETTING].type].delete_state())
              const params = { 
                method: "POST", 
                data: [{ 
                  key: "state",
                  text: sqlInfo.sql,
                  values: Array(sqlInfo.prmCount).fill(item.id)
                }]
              }
              const view = await requestData("/view", params)
              if(view.error){
                resultError(view)
              } else if (view.state[0].sco > 0) {
                showToast(TOAST_TYPE.ERROR, msg("", { id: "msg_integrity_err" }))
              } else {
                deleteData()
              }
            } else {
              deleteData()
            }
          }
        }
      }
    })
    setData("current", { modalForm, side: SIDE_VISIBILITY.HIDE })
  }

  editItem(options) {
    const { setData, data } = this.store
    let settings = {...data[APP_MODULE.SETTING]}
    if(settings.type === "program"){
      settings = {
        ...settings,
        current: {
          ...settings.current,
          form: {
            ...settings.current.form,
            [options.name]: options.value
          }
        }
      }
      localStorage.setItem(options.name, options.value)
    } else if(options.name === "log_search"){
      this.loadLog()
    } else {
      if((settings.audit==="all") || (settings.audit==="update")){
        settings = {
          ...settings,
          dirty: true
        }
      }
      if((options.name === "fieldvalue_value") || (options.name === "fieldvalue_notes")){
        settings = {
          ...settings,
          current: {
            ...settings.current,
            fieldvalue: {
              ...settings.current.fieldvalue,
              [options.name]: options.value
            },
            form: {
              ...settings.current.form,
              [options.name.split("_")[1]]: options.value.toString()
            }
          }
        }
      } else {
        settings = {
          ...settings,
          current: {
            ...settings.current,
            form: {
              ...settings.current.form,
              [options.name]: options.value
            }
          }
        }
      }
    }
    setData(APP_MODULE.SETTING, settings)
  }

  async loadLog() {
    const { sql } = this.app.modules
    const { getSql, requestData, resultError, showToast, msg } = this.app
    const { setData, data } = this.store
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
    let setting = {...data[APP_MODULE.SETTING],
      view: {
        ...data[APP_MODULE.SETTING].view,
        result: []
      }
    }
    const _log = sql.log.result()
    let paramList = []
    if (!["login", "logout"].includes(setting.current.form.logstate)){
      _log.inner_join = [..._log.inner_join, 
        ["groups nt","on",[["l.nervatype","=","nt.id"],["and","nt.groupvalue","=","?"]]]
      ]
      paramList = [...paramList, setting.current.form.nervatype]
      if(!nervatype_filter[setting.current.form.nervatype]){
        return showToast(TOAST_TYPE.ERROR, `${msg("", { id: "msg_required" })} ${msg("", { id: "log_nervatype" })}`)
      }
      _log.inner_join = [..._log.inner_join, nervatype_filter[setting.current.form.nervatype].inner_join]
      _log.select[4] = nervatype_filter[setting.current.form.nervatype].select
    }
    _log.where = [..._log.where, ["ls.groupvalue","=","?"]]
    paramList = [...paramList, setting.current.form.logstate]
    
    if (setting.current.form.empnumber && setting.current.form.empnumber !== "") {
      _log.where = [..._log.where,
        ["and","lower(e.empnumber)","like",`{CCS}{JOKER}{SEP}lower('${setting.current.form.empnumber}'){SEP}{JOKER}{CCE}`]
      ]
    }
    
    if (setting.current.form.fromdate && setting.current.form.fromdate !== "") {
      _log.where = [..._log.where,
        ["and","{FMS_DATE}l.crdate{FME_DATE}",">=","?"]
      ]
      paramList = [...paramList, setting.current.form.fromdate]
    }
    
    if (setting.current.form.todate && setting.current.form.todate !== "") {
      _log.where = [..._log.where,
        ["and","{FMS_DATE}l.crdate{FME_DATE}","<=","?"]
      ]
      paramList = [...paramList, setting.current.form.todate]
    }

    const params = { 
      method: "POST", 
      data: [{ 
        key: "log",
        text: getSql(data.login.data.engine, _log).sql,
        values: paramList
      }]
    }
    const view = await requestData("/view", params)
    if(view.error){
      return resultError(view)
    
    }
    setting = {...setting,
      view: {
        ...setting.view,
        result: view.log
      }
    }
    setData(APP_MODULE.SETTING, setting)
    return true
  }

  async loadSetting(options) {
    const { dataSet, sql } = this.app.modules
    const { getSql, requestData, resultError, currentModule } = this.app
    const { data } = this.store
    let setting = {...options,
      dataset: {},
      dirty: false
    }
    if(dataSet[setting.type]){
      let views = []
      dataSet[setting.type]().forEach(info => {
        let _sql = {}
        if(info.infoType === "table"){
          _sql = { select:["*"], from: info.classAlias }
          if(info.where){
            _sql.where = info.where
          }
          if(info.order){
            _sql.order_by = info.order
          }
        } else if (typeof sql[setting.type][info.infoName] !== "undefined") {
          _sql = sql[setting.type][info.infoName](setting.type)
        } else {
          _sql = sql.all[info.infoName](setting.type)
        }
        const sqlInfo = getSql(data[APP_MODULE.LOGIN].data.engine, _sql)
        if( (setting.id !== null) || (sqlInfo.prmCount === 0)){
          views = [...views, {
            key: info.infoName,
            text: sqlInfo.sql,
            values: ((sqlInfo.prmCount>0) && (setting.id !== null)) ? Array(sqlInfo.prmCount).fill(setting.id) : []
          }]
        } else {
          setting = {
            ...setting,
            dataset: {
              ...setting.dataset,
              [info.infoName]: []
            }
          }
        }
      })

      const params = { method: "POST", data: views }
      const view = await requestData("/view", params)
      if(view.error){
        return resultError(view)
      }
      setting = {
        ...setting,
        dataset: {
          ...setting.dataset,
          ...view
        }
      }
      if ((setting.type === "template") && (typeof setting.id !== "undefined")) { 
        return currentModule({ 
          data: { module: APP_MODULE.TEMPLATE, side: SIDE_VISIBILITY.HIDE }, 
          content: { 
            fkey: "setTemplate", 
            args: [setting]
          } 
        })
      }
    }
    return this.setSettingData(setting)
  }

  onSettingEvent({key, data}){
    const { setData } = this.store
    switch (key) {
      case SETTING_EVENT.CURRENT_PAGE:
        setData(APP_MODULE.SETTING, {
          page: data.value 
        })
        break;

      case SETTING_EVENT.FORM_ACTION:
        this.setFormActions(data)
        break;

      case SETTING_EVENT.EDIT_ITEM:
        this.editItem(data)
        break;

      default:
        break;
    }
    return true
  }

  async onSideEvent({key, data}){
    const { showHelp, currentModule } = this.app
    const { setData } = this.store
    const storeSetting = this.store.data[APP_MODULE.SETTING]
    const storeEdit = this.store.data[APP_MODULE.EDIT]
    switch (key) {
      case SIDE_EVENT.BACK:
        let backType = storeSetting.type
        if(backType === "password"){
          backType = "setting"
          setData(APP_MODULE.SETTING, { group_key: "group_admin" })
          this.loadSetting({ type: 'setting' })
        } else {
          this.checkSetting({ type: backType }, SIDE_EVENT.LOAD_SETTING)
        }
        break;

      case SIDE_EVENT.CHANGE:
        setData(APP_MODULE.SETTING, {
          [data.fieldname]: data.value 
        })
        break;
      
      case SIDE_EVENT.SAVE:
        setData("current", { side: SIDE_VISIBILITY.HIDE })
        if(storeSetting.type === "password"){
          this.changePassword()
        } else {
          const result = await this.saveSetting()
          if(result){
            this.loadSetting({type: result.type, id: result.current.form.id})
          }
        }
        break;

      case SIDE_EVENT.DELETE:
        this.deleteSetting(data.value)
        break;

      case SIDE_EVENT.CHECK:
        if(data[0].ntype){
          currentModule({ 
            data: { module: APP_MODULE.EDIT, side: SIDE_VISIBILITY.HIDE }, 
            content: { 
              fkey: "checkEditor", 
              args: [{ ntype: "customer", ttype: null, id: 1 }, ACTION_EVENT.LOAD_EDITOR] 
            } 
          })
        } else {
          this.checkSetting(...data)
        }
        break;

      case SIDE_EVENT.LOAD_SETTING:
        this.loadSetting(data)
        break;

      case SIDE_EVENT.PROGRAM_SETTING:
        this.setProgramForm()
        break;

      case SIDE_EVENT.PASSWORD:
        let username = data.username
        if(!username && storeEdit.current){
          username = storeEdit.dataset[storeEdit.current.type][0].username
        }
        this.checkSetting({ username }, SIDE_EVENT.PASSWORD_FORM)
        break;

      case SIDE_EVENT.HELP:
        showHelp(data.value)
        break;

      default:
        break;
    }
    return true
  }

  async saveSetting() {
    const { initItem, validator } = this.app.modules
    const { resultError, requestData } = this.app
    const { data } = this.store

    let setting = {...data[APP_MODULE.SETTING]}

    let values = this.tableValues(setting.ntype, setting.current.form)
    values = await validator(setting.ntype, values)
    if(values.error){
      return resultError(values)
    }

    let result = await requestData(`/${setting.ntype}`, { method: "POST", data: [values] })
    if(result.error){
      resultError(result)
      return null
    }
    if (setting.current.form.id === null) {
      setting = {...setting,
        current: {...setting.current,
          form: {...setting.current.form,
            id: result[0]
          }
        }
      }
    }
    setting = {...setting,
      dirty: false
    }
    if (setting.type === "usergroup") {
      const transfilter_name = data[APP_MODULE.LOGIN].data.groups.filter(
        item => (item.id === setting.current.form.transfilter))[0].groupvalue
      if((setting.current.form.translink !== null) || (transfilter_name !== "all")){
        const link = {
          ...initItem({tablename: "link", dataset: setting.dataset, current: setting.current}),
          id: setting.current.form.translink,
          nervatype_1: data[APP_MODULE.LOGIN].data.groups.filter(
            (item)=>((item.groupname === "nervatype") && (item.groupvalue === "groups")))[0].id,
          ref_id_1: setting.current.form.id,
          nervatype_2: data[APP_MODULE.LOGIN].data.groups.filter(
            (item)=>((item.groupname === "nervatype") && (item.groupvalue === "groups")))[0].id,
          ref_id_2: setting.current.form.transfilter
        }
        result = await requestData("/link", { method: "POST", data: [link] })
        if(result.error){
          resultError(result)
          return null
        }
      }
    }
    return setting
  }

  setPasswordForm(username) {
    const { forms } = this.app.modules
    const { setData } = this.store
    const data = {
      username, password_1: "", password_2: "" 
    }
    const form = forms.password(data)
    const setting = {
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
    setData(APP_MODULE.SETTING, setting)
  }

  setProgramForm() {
    const { forms } = this.app.modules
    const { getSetting } = this.app
    const { setData } = this.store
    const template = forms.program()
    const setting = {
      type: "program", 
      dataset: {}, 
      current: {
        form: {
          paginationPage: getSetting("paginationPage"),
          history: getSetting("history"),
          page_size: getSetting("page_size"),
          export_sep: getSetting("export_sep"),
        },
        template
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
    }
    setData(APP_MODULE.SETTING, setting)
  }

  setSettingData(options) {
    const { forms, initItem } = this.app.modules
    const { getAuditFilter } = this.app
    const { data, setData } = this.store
    const form = forms[options.type]()
    const audit = getAuditFilter(options.type)[0]
    const setting = {...options,
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
        result: options.dataset[`${options.type}_view`],
        fields: (form.view.setting.fields) ? form.view.setting.fields : null
      },
      actions: {
        new: (audit !== "all") ? null : form.view.setting.actions.new, 
        edit: form.view.setting.actions.edit, 
        delete: (audit !== "all") ? null : form.view.setting.actions.delete
      },
      audit,
      template: null,
      page: (data[APP_MODULE.SETTING].type === options.type) ? data[APP_MODULE.SETTING].page || 0 : 0
    }
    setData(APP_MODULE.SETTING, setting)
    if(((options.type === "usergroup") || (options.type === "ui_menu")) && options.id) {
      this.setFormActions({ params: {action: ACTION_EVENT.EDIT_ITEM, setting}, row: { id: options.id }})
    } else if(typeof options.id !== "undefined"){
      this.setSettingForm(options.id, setting)
    }
  }

  setSettingForm(id, isetting) {
    const { forms, initItem } = this.app.modules
    const { data, setData } = this.store
    let setting = {...(isetting||data[APP_MODULE.SETTING]),
      current: {}
    }
    if (id!==null) {
      const item = setting.dataset[`${setting.type}_view`].filter(i => (i.id === parseInt(id,10)))[0]
      setting = {...setting,
        dirty: false,
        current: { form: item }
      }
      if(setting.ntype === "fieldvalue"){
        setting = {...setting,
          current: {...setting.current,
            fieldvalue: {
              rowtype: 'fieldvalue',
              id: item.id, 
              // name: 'fieldvalue_value', 
              fieldname: item.fieldname, 
              fieldvalue_value: item.value, 
              fieldvalue_notes: item.notes || '',
              label: item.lslabel, 
              description: (item.fieldtype === 'valuelist') ? item.valuelist.split("|") : item.value, 
              disabled: 'false',
              fieldtype: item.fieldtype, 
              datatype: (item.fieldtype === 'urlink') ? 'text' : item.fieldtype
            }
          }
        }
      }
    } else {
      setting = {...setting,
        dirty: !!(((setting.audit==="all") || (setting.audit==="update"))),
        current: { form: initItem({tablename: setting.type, 
          dataset: setting.dataset, current: setting.current}) }
      }
    }
    if((setting.type === "usergroup") && (setting.current.form.transfilter === null)){
      setting.current.form.transfilter = setting.dataset.transfilter.filter(
        item => (item.groupvalue === "all"))[0].id
    }
    setting = {...setting,
      current: {...setting.current,
        template: forms[setting.type](setting.current.form, setting)
      }
    }
    setting = {...setting,
      panel: setting.current.template.options.panel,
    }
    if (setting.audit === "readonly") {
      setting.panel.save = false
    }
    if (setting.audit !== "all") {
      setting.panel.delete = false;
      setting.panel.new = false;
    }
    setData(APP_MODULE.SETTING, setting)
  }

  async setFormActions(options) {
    const { modalAudit, modalMenu } = this.module
    const { inputBox } = this.host
    const { sql, initItem } = this.app.modules
    const { getSql, requestData, resultError, currentModule, msg } = this.app
    const { setData, data } = this.store
    const row = options.row || {}
    switch (options.params.action) {
      case ACTION_EVENT.NEW_ITEM:
        this.checkSetting({ type: data[APP_MODULE.SETTING].type, id: null }, SIDE_EVENT.LOAD_SETTING)
        break;

      case ACTION_EVENT.EDIT_ITEM:
        switch (data[APP_MODULE.SETTING].type) {
          case "template":
            this.checkSetting({ type: data[APP_MODULE.SETTING].type, id: row.id }, SIDE_EVENT.LOAD_SETTING)
            break;

          case "place":
            currentModule({ 
              data: { module: APP_MODULE.EDIT, side: SIDE_VISIBILITY.HIDE }, 
              content: { 
                fkey: "checkEditor", 
                args: [{ ntype: data[APP_MODULE.SETTING].type, ttype: null, id: row.id || null }, ACTION_EVENT.LOAD_EDITOR] } 
            })
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
            const audit_view = await requestData("/view", audit_options)
            if(audit_view.error){
              resultError(audit_view)
            } else {
              let audit_setting = {
                ...(options.params.setting||data[APP_MODULE.SETTING]),
              }
              audit_setting = {...audit_setting,
                dataset: {
                  ...audit_setting.dataset,
                  audit: audit_view.audit
                }
              }
              this.setSettingForm(row.id, audit_setting)
            }
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
            const menufields_view = await requestData("/view", menufields_options)
            if(menufields_view.error){
              resultError(menufields_view)
            } else {
              let menufields_setting = {
                ...(options.params.setting||data[APP_MODULE.SETTING]),
              }
              menufields_setting = {...menufields_setting,
                dataset: {
                  ...menufields_setting.dataset,
                  ui_menufields: menufields_view.menufields
                }
              }
              this.setSettingForm(row.id, menufields_setting)
            }
            break;
        
          default:
            this.setSettingForm(row.id)
            break;
        }
        break;

      case ACTION_EVENT.DELETE_ITEM:
        this.deleteSetting(row)
        break;

      case ACTION_EVENT.EDIT_AUDIT:
        let audit = row
        if(!audit.id){
          audit = {
            ...initItem({tablename: "link", dataset: data[APP_MODULE.SETTING].dataset, current: data[APP_MODULE.SETTING].current}),
            usergroup: data[APP_MODULE.SETTING].current.form.id,
            nervatype: data[APP_MODULE.SETTING].dataset.nervatype.filter(
              (item)=>(item.groupvalue === "customer"))[0].id,
            inputfilter: data[APP_MODULE.SETTING].dataset.inputfilter.filter(
              (item)=>(item.groupvalue === "all"))[0].id
          }
        }
        const auditForm = modalAudit({
          idKey: audit.id, usergroup: audit.usergroup,
          nervatype: audit.nervatype, subtype: audit.subtype,
          inputfilter: audit.inputfilter,
          supervisor: audit.supervisor,
          typeOptions: data[APP_MODULE.SETTING].dataset.nervatype.map(group => ({ value: String(group.id), text: group.groupvalue })),
          subtypeOptions: data[APP_MODULE.SETTING].dataset.transtype.map(group => ({ value: String(group.id), text: group.groupvalue, type: "trans" })).concat(
            data[APP_MODULE.SETTING].dataset.reportkey.map(report => ({ value: String(report.id), text: report.reportkey, type: "report" })),
            data[APP_MODULE.SETTING].dataset.menukey.map(menu => ({ value: String(menu.id), text: menu.menukey, type: "menu" }))
          ),
          inputfilterOptions: data[APP_MODULE.SETTING].dataset.inputfilter.map(group => ({ value: String(group.id), text: group.groupvalue })),
          onEvent: {
            onModalEvent: async (modalResult) => {
              setData("current", { modalForm: null })
              if(modalResult.key === MODAL_EVENT.OK){
                const result = await requestData("/ui_audit", { 
                  method: "POST", data: [this.tableValues("audit", modalResult.data.value)] })
                if(result.error){
                  resultError(result)
                } else {
                  this.setFormActions({ params: {action: ACTION_EVENT.EDIT_ITEM}, row: data[APP_MODULE.SETTING].current.form})
                }
              }
            }
          }
        })
        setData("current", { modalForm: auditForm, side: SIDE_VISIBILITY.HIDE })
        break;

      case ACTION_EVENT.EDIT_MENU_FIELD:
        let menufields = row
        if(!menufields.id){
          menufields = {
            ...initItem({tablename: "ui_menufields", dataset: data[APP_MODULE.SETTING].dataset, current: data[APP_MODULE.SETTING].current}),
            menu_id: data[APP_MODULE.SETTING].current.form.id,
            fieldtype: data[APP_MODULE.SETTING].dataset.fieldtype.filter(
              (item)=>(item.groupvalue === "string"))[0].id,
            orderby: data[APP_MODULE.SETTING].dataset.ui_menufields.length
          }
        }
        const menuForm = modalMenu({
          idKey: menufields.id, menu_id: menufields.menu_id,
          fieldname: menufields.fieldname, description: menufields.description,
          fieldtype: menufields.fieldtype, orderby: menufields.orderby,
          fieldtypeOptions: data[APP_MODULE.SETTING].dataset.fieldtype.map(group => ({ value: String(group.id), text: group.groupvalue })),
          onEvent: {
            onModalEvent: async (modalResult) => {
              setData("current", { modalForm: null })
              if(modalResult.key === MODAL_EVENT.OK){
                const result = await requestData("/ui_menufields", { 
                  method: "POST", data: [this.tableValues("ui_menufields", modalResult.data.value)] })
                if(result.error){
                  resultError(result)
                } else {
                  this.setFormActions({ params: {action: ACTION_EVENT.EDIT_ITEM}, row: data[APP_MODULE.SETTING].current.form})
                }
              }
            }
          }
        })
        setData("current", { modalForm: menuForm, side: SIDE_VISIBILITY.HIDE })
        break;

      case ACTION_EVENT.DELETE_ITEM_ROW:
        const  modalForm = inputBox({ 
          title: msg("", { id: "msg_warning" }),
          message: msg("", { id: "msg_delete_text" }),
          infoText: msg("", { id: "msg_delete_info" }),
          onEvent: {
            onModalEvent: async (modalResult) => {
              setData("current", { modalForm: null })
              if (modalResult.key === MODAL_EVENT.OK) {
                const result = await requestData(`/${options.params.table}`, { method: "DELETE", query: { id: row.id } })
                if(result && result.error){
                  resultError(result)
                } else {
                  this.setFormActions({ params: {action: ACTION_EVENT.EDIT_ITEM}, row: data[APP_MODULE.SETTING].current.form})
                }
              }
            }
          }
        })
        setData("current", { modalForm, side: SIDE_VISIBILITY.HIDE })
        break;

      default:
        break;
    }
  }

  tableValues(type, item) {
    const { initItem } = this.app.modules
    const { data } = this.store
    const values = {}
    const baseValues = initItem({tablename: type, 
      dataset: data[APP_MODULE.SETTING].dataset, current: data[APP_MODULE.SETTING].current})
    Object.keys(item).forEach(key => {
      if (typeof(baseValues[key]) !== "undefined") {
        values[key] = item[key]
      }
    });
    return values
  }

}