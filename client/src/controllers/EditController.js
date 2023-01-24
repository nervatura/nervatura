import { APP_MODULE, MODAL_EVENT, EDIT_EVENT, SIDE_EVENT, EDITOR_EVENT, TOAST_TYPE, 
  SIDE_VISIBILITY, ACTION_EVENT } from '../config/enums.js'

export const checkSubtype = (type, subtype, item) => {
  const result = {
    customer: (subtype === item.custtype),
    place: (subtype === item.placetype),
    product: (subtype === item.protype),
    tool: (subtype === item.toolgroup),
    trans: (subtype === item.transtype)
  }
  if(subtype === null)
    return true
  return (typeof result[type] === "undefined") ? true : result[type]
}

const roundAmount = (amount, digit) => Math.round(amount*(10**digit))/(10**digit)

export class EditController {
  constructor(host) {
    this.host = host
    this.app = host.app
    this.store = host.app.store
    this.module = {}
    
    this.addPrintQueue = this.addPrintQueue.bind(this)
    this.calcFormula = this.calcFormula.bind(this)
    this.calcPrice = this.calcPrice.bind(this)
    this.checkEditor = this.checkEditor.bind(this)
    this.checkTranstype = this.checkTranstype.bind(this)
    this.createReport = this.createReport.bind(this)
    this.createShipping = this.createShipping.bind(this)
    this.createTrans = this.createTrans.bind(this)
    this.createTransOptions = this.createTransOptions.bind(this)
    this.deleteEditor = this.deleteEditor.bind(this)
    this.deleteEditorItem = this.deleteEditorItem.bind(this)
    this.editItem = this.editItem.bind(this)
    this.exportEvent = this.exportEvent.bind(this)
    this.exportQueueAll = this.exportQueueAll.bind(this)
    this.exportQueue = this.exportQueue.bind(this)
    this.getTransFilter = this.getTransFilter.bind(this)
    this.loadEditor = this.loadEditor.bind(this)
    this.newFieldvalue = this.newFieldvalue.bind(this)
    this.nextTransNumber = this.nextTransNumber.bind(this)
    this.onEditEvent = this.onEditEvent.bind(this)
    this.onSelector = this.onSelector.bind(this)
    this.onSideEvent = this.onSideEvent.bind(this)
    this.prevTransNumber = this.prevTransNumber.bind(this)
    this.reportOutput = this.reportOutput.bind(this)
    this.reportPath = this.reportPath.bind(this)
    this.reportSettings = this.reportSettings.bind(this)
    this.saveEditor = this.saveEditor.bind(this)
    this.saveEditorForm = this.saveEditorForm.bind(this)
    this.searchQueue = this.searchQueue.bind(this)
    this.setEditor = this.setEditor.bind(this)
    this.setEditorItem = this.setEditorItem.bind(this)
    this.setFieldvalue = this.setFieldvalue.bind(this)
    this.setFormActions = this.setFormActions.bind(this)
    this.setModule = this.setModule.bind(this)
    this.setLink = this.setLink.bind(this)
    this.setPassword = this.setPassword.bind(this)
    this.setPattern = this.setPattern.bind(this)
    this.shippingAddAll = this.shippingAddAll.bind(this)
    this.showStock = this.showStock.bind(this)
    this.tableValues = this.tableValues.bind(this)
    this.transCopy = this.transCopy.bind(this)
    host.addController(this);
  }

  setModule(moduleRef){
    this.module = moduleRef
  }

  async addPrintQueue(reportkey, copy) {
    const { requestData, resultError, showToast, msg } = this.app
    const { current, dataset } = this.store.data[APP_MODULE.EDIT]
    const login = this.store.data[APP_MODULE.LOGIN].data
    const report = dataset.report.filter((item)=>(item.reportkey === reportkey))[0]
    const ntype = login.groups.filter(
      (item)=>((item.groupname === "nervatype") && (item.groupvalue === current.type)))[0]
    const values = {
      "nervatype": ntype.id, 
      "ref_id": current.item.id, 
      "qty": parseInt(copy, 10), 
      "employee_id": login.employee.id, 
      "report_id": report.id
    }
    const options = { method: "POST", data: [values] }
    const result = await requestData("/ui_printqueue", options)
    if(result.error){
      return resultError(result)
    }
    showToast(TOAST_TYPE.SUCCESS, msg("", { id: "report_add_groups" }))
    return true
  }

  calcFormula(formula_id) {
    const { inputBox } = this.host
    const { setData, data } = this.store
    const { sql, initItem } = this.app.modules
    const { getSql, requestData, resultError, msg } = this.app
    const  modalForm = inputBox({ 
      title: msg("", { id: "msg_warning" }),
      message: msg("", { id: "ms_load_formula" }),
      infoText: `${msg("", { id: "msg_delete_info" }) } ${  msg("", { id: "ms_continue_warning" })}`,
      defaultOK: true,
      onEvent: {
        onModalEvent: async (modalResult) => {
          setData("current", { modalForm: null })
          if (modalResult.key === MODAL_EVENT.OK) {
            const params = { 
              method: "POST", 
              data: [{ 
                key: "formula",
                text: getSql(data[APP_MODULE.LOGIN].data.engine, sql.trans.formula_items()).sql,
                values: [formula_id]
              }]
            }
            const view = await requestData("/view", params)
            if(view.error){
              return resultError(view)
            }
            const production_qty = data[APP_MODULE.EDIT].dataset.movement_head[0].qty
            const production_place = data[APP_MODULE.EDIT].dataset.movement_head[0].place_id
            const formula_qty = data[APP_MODULE.EDIT].dataset.formula_head.filter(item => (item.id === formula_id))[0].qty
            let items = [] 
            view.formula.forEach(fitem => {
              const item = {
                ...initItem({tablename: "movement", dataset: data[APP_MODULE.EDIT].dataset, current: data[APP_MODULE.EDIT].current}),
                product_id: fitem.product_id,
                place_id: (fitem.place_id === null) ? production_place : fitem.place_id,
                qty: (fitem.shared === 1) 
                  ? -Math.ceil(production_qty/formula_qty) 
                  : -(roundAmount((production_qty/formula_qty)*fitem.qty, 2))
              }
              items = [...items, item]
            })
            for (let index = 0; index < data[APP_MODULE.EDIT].dataset.movement.length; index += 1) {
              // eslint-disable-next-line no-await-in-loop
              const result = await requestData(
                "/movement", { method: "DELETE", query: { id: data.edit.dataset.movement[index].id } })
              if(result && result.error){
                return resultError(result)
              }
            }
            const result = await requestData("/movement", { method: "POST", data: items })
            if(result.error){
              return resultError(result)
            }
            this.loadEditor({
              ntype: data[APP_MODULE.EDIT].current.type, 
              ttype: data[APP_MODULE.EDIT].current.transtype, 
              id: data[APP_MODULE.EDIT].current.item.id, 
              form: "movement"
            })
          }
          return true
        }
      }
    })
    return setData("current", { modalForm, side: SIDE_VISIBILITY.HIDE })
  }

  calcPrice(_calcmode, item) {
    const { data } = this.store
    let rate = data[APP_MODULE.EDIT].dataset.tax.filter(tax => (tax.id === parseInt(item.tax_id,10)))[0]
    rate = (typeof rate !== "undefined") ? rate.rate : 0
    let digit = data[APP_MODULE.EDIT].dataset.currency.filter(currency => 
      (currency.curr === data[APP_MODULE.EDIT].current.item.curr))[0]
    digit = (typeof digit !== "undefined") ? digit.digit : 2
    
    let netAmount = 0; let vatAmount = 0; let amount = 0; let fxPrice = 0;
    switch(_calcmode) {
        
      case "netamount":
        netAmount = parseFloat(item.netamount)
        if (parseFloat(item.qty)!==0) {
          fxPrice = roundAmount(netAmount/(1-parseFloat(item.discount)/100)/parseFloat(item.qty),parseInt(digit,10))
          vatAmount = roundAmount(netAmount*parseFloat(rate),parseInt(digit,10))
        }
        amount = roundAmount(netAmount+vatAmount,parseInt(digit,10))
        break;

      case "amount":
        amount = parseFloat(item.amount)
        if (parseFloat(item.qty)!==0) {
          netAmount = roundAmount(amount/(1+parseFloat(rate)),parseInt(digit,10))
          vatAmount = roundAmount(amount-netAmount,parseInt(digit,10))
          fxPrice = roundAmount(netAmount/(1-parseFloat(item.discount)/100)/parseFloat(item.qty),parseInt(digit,10))
        }
        break;

      case "fxprice":
      default:
        fxPrice = parseFloat(item.fxprice)
        netAmount = roundAmount(fxPrice*(1-parseFloat(item.discount)/100)*parseFloat(item.qty),parseInt(digit,10))
        vatAmount = roundAmount(fxPrice*(1-parseFloat(item.discount)/100)*parseFloat(item.qty)*parseFloat(rate),parseInt(digit,10))
        amount = roundAmount(netAmount+vatAmount, parseInt(digit,10))
        break;
    }
    return {...item,
      fxprice: fxPrice,
      netamount: netAmount,
      vatamount: vatAmount,
      amount
    }
  }

  checkEditor(options, cbKeyTrue) {
    const { modalFormula } = this.module
    const { inputBox } = this.host
    const { msg } = this.app
    const { setData, data } = this.store
    const cbNext = (cbKey) =>{
      switch (cbKey) {
        case EDITOR_EVENT.LOAD_EDITOR:
          this.loadEditor(options);
          break;

        case EDITOR_EVENT.SET_EDITOR_ITEM:
          this.setEditorItem(options);
          break;

        case EDITOR_EVENT.LOAD_FORMULA:
          const modalForm = modalFormula({
            formula: options.formula,
            partnumber: data[APP_MODULE.EDIT].dataset.movement_head[0].partnumber,
            description: data[APP_MODULE.EDIT].dataset.movement_head[0].description,
            formulaValues: data[APP_MODULE.EDIT].dataset.formula_head.map(
              formula => ({ value: String(formula.id), text: formula.transnumber })
            ),
            onEvent: {
              onModalEvent: async (modalResult) => {
                setData("current", { modalForm: null })
                if(modalResult.key === MODAL_EVENT.OK){
                  this.calcFormula(modalResult.data.value)
                }
              }
            }
          })
          setData("current", { modalForm, side: SIDE_VISIBILITY.HIDE })
          break;

        case EDITOR_EVENT.NEW_FIELDVALUE:
          this.newFieldvalue(options.fieldname)
          break;

        case EDITOR_EVENT.CREATE_TRANS:
          this.createTrans(options)
          break;

        case EDITOR_EVENT.CREATE_TRANS_OPTIONS:
          this.createTransOptions()
          break;

        case EDITOR_EVENT.FORM_ACTION:
          this.setFormActions(options)
          break;

        default:
          break;
      }
    }
    if ((data[APP_MODULE.EDIT].dirty === true && data[APP_MODULE.EDIT].current.item) || 
      (data[APP_MODULE.EDIT].form_dirty === true && data[APP_MODULE.EDIT].current.form)) {
        const  modalForm = inputBox({ 
          title: msg("", { id: "msg_warning" }),
          message: msg("", { id: "msg_dirty_text" }),
          infoText: msg("", { id: "msg_delete_info" }),
          defaultOK: true,
          onEvent: {
            onModalEvent: async (modalResult) => {
              setData("current", { modalForm: null })
              if (modalResult.key === MODAL_EVENT.OK) {
                const edit = (data[APP_MODULE.EDIT].form_dirty) 
                  ? await this.saveEditorForm()
                  : await this.saveEditor()
                if(edit){
                  setData(APP_MODULE.EDIT, edit)
                  cbNext(cbKeyTrue)
                }
              } else {
                cbNext(cbKeyTrue)
              }
            }
          }
        })
        return setData("current", { modalForm, side: SIDE_VISIBILITY.HIDE })
    }
    return cbNext(cbKeyTrue);
  }

  async checkTranstype(options, cbKeyTrue) {
    const { requestData, resultError, getSql } = this.app
    const login = this.store.data[APP_MODULE.LOGIN].data
    if ((options.ntype==="trans" || options.ntype==="transitem" || 
      options.ntype==="transmovement" || options.ntype==="transpayment") && 
      options.ttype===null) {
        const params = { 
          method: "POST", 
          data: [{ 
            key: "transtype",
            text: getSql(login.engine, {
              select:["groupvalue"], from:"groups g", 
              inner_join:["trans t","on",[["g.id","=","t.transtype"],["and","t.id","=","?"]]]}).sql,
            values: [options.id] 
          }]
        }
        const view = await requestData("/view", params)
        if(view.error){
          return resultError(view)
        }
        return this.checkEditor({...options, ntype:"trans", ttype:view.transtype[0].groupvalue, id:options.id}, cbKeyTrue)
    }
    return this.checkEditor(options, cbKeyTrue)
  }

  async createReport(_output) {
    let output = _output
    const { requestData, resultError, saveToDisk } = this.app
    const { current } = this.store.data[APP_MODULE.EDIT]
    let _filters = [];
    current.fieldvalue.forEach((rfdata) => {
      if (rfdata.selected) {
        _filters = [..._filters, `filters[${rfdata.name}]=${rfdata.value}`]
      }
    })
    const report = current.item
    const params = {
      type: (output === "xml") ? "xml" : "auto",
      ctype: (output === "xml") 
        ? "application/xml; charset=UTF-8" 
        : (output === "csv") 
          ? "text/csv; charset=UTF-8" 
          : "application/pdf",
      template: report.reportkey, 
      title: report.reportkey,
      orient: report.orientation, 
      size: report.size,
      filters: _filters.join("&")
    }
    const result = await requestData(this.reportPath(params), {})
    if(result && result.error){
      return resultError(result)
    }
    let resultUrl
    if(output === "csv"){
      const blob = new Blob([result], { type: 'text/csv;charset=utf-8;' })
      resultUrl = URL.createObjectURL(blob)
      output = "csv"
    } else {
      resultUrl = URL.createObjectURL(result, {type : params.ctype})
    }
    if(output === "print"){
      return window.open(resultUrl,"_blank")
    }
    let filename = `${params.title}_${new Date().toISOString().split("T")[0]}.${output}`
    filename = filename.split("/").join("_")
    return saveToDisk(resultUrl, filename)
  }

  async createShipping() {
    const { initItem } = this.app.modules
    const { requestData, resultError, showToast, msg, createHistory } = this.app
    const { dataset, current } = this.store.data[APP_MODULE.EDIT]
    if (current.shipping_place_id === null) {
      return showToast(TOAST_TYPE.ERROR, 
        `${msg("", { id: "msg_required" })} ${msg("", { id: "inventory_warehouse" })}`)
    }
    if (dataset.shiptemp.length > 0){
      let delivery_head = {...initItem({tablename: "trans", dataset, current}),
        transtype: dataset.groups.filter(
          (item) => ((item.groupname === "transtype") && (item.groupvalue === "delivery"))
        )[0].id,
        direction: dataset[current.type][0].direction,
        transdate: new Date(current.item.shippingdate).toISOString().split("T")[0],
        duedate: null, curr: null, paidtype: null
      }
      const pattern = dataset.delivery_pattern.filter(
          (item) => (item.defpattern === 1)
        )[0]
      if (typeof pattern !== "undefined") {
        delivery_head = {...delivery_head,
          fnote: pattern.notes
        }
      }
      
      const params = { method: "POST", 
        data: {
          key: "nextNumber",
          values: {
            numberkey: `delivery_${current.direction}`, 
            step: true
          }
        }
      }
      let result = await requestData("/function", params)
      if(result.error){
        resultError(result)
        return null
      }
      delivery_head = {...delivery_head,
        transnumber: result
      }

      result = await requestData("/trans", { method: "POST", data: [delivery_head] })
      if(result.error){
        resultError(result)
        return null
      }
      delivery_head.id = result[0];
      await createHistory("save")

      let movements = [];
      dataset.shiptemp.forEach((shiptemp) => {
        const movement = {...initItem({tablename: "movement", dataset, current}),
          trans_id: delivery_head.id,
          shippingdate: `${new Date(current.shippingdate).toISOString().slice(0,10)}T${new Date(current.shippingdate).toLocaleTimeString("en",{hour12: false}).replace("24","00")}`,
          product_id: shiptemp.product_id,
          place_id: current.shipping_place_id,
          notes: shiptemp.batch_no,
          qty: (current.direction === "out") ? -(shiptemp.qty) : shiptemp.qty
        }
        movements=[...movements, movement];
      });
      result = await requestData("/movement", { method: "POST", data: movements })
      if(result.error){
        resultError(result)
        return null
      }

      let links = [];
      const nervatype_movement = dataset.groups.filter(
        (item)=>((item.groupname === "nervatype") && (item.groupvalue === "movement")))[0].id
      const nervatype_item = dataset.groups.filter(
        (item)=>((item.groupname === "nervatype") && (item.groupvalue === "item")))[0].id
      result.forEach((movement_id, index) => {
        const link = {...initItem({tablename: "link", dataset, current}),
          nervatype_1: nervatype_movement,
          ref_id_1: movement_id,
          nervatype_2: nervatype_item,
          ref_id_2: dataset.shiptemp[index].item_id
        }
        links=[...links, link];
      });
      result = await requestData("/link", { method: "POST", data: links })
      if(result.error){
        resultError(result)
        return null
      }

      return this.loadEditor({
        ntype: current.type, 
        ttype: current.transtype, 
        id: current.item.id, 
        shipping: true
      })

    }
    return true
  }

  async createTrans(options) {
    const { data } = this.store
    const { initItem } = this.app.modules
    const { resultError, showToast, requestData, createHistory, msg } = this.app

    const check_refnumber = (params) => {
      if(params.transtype === "waybill"){
        return "link";
      }if(params.transtype==="delivery" && (params.transcast === "normal") && 
        (params.direction==="in" || params.direction==="out")){
        return "";
      } if (params.cmdtype === "copy" && params.transcast === "normal") {
        return "";
      } if ((params.transcast !== "normal") || params.refno){
        return "reflink";
      } 
      return "refnumber";
    }

    const base_trans = {...data[APP_MODULE.EDIT].dataset.trans[0]};
    // set base data
    let transtype = data[APP_MODULE.EDIT].dataset.groups.filter(
      (item) => (item.id === base_trans.transtype)
    )[0].groupvalue
    let transtype_id = base_trans.transtype;
    let direction = data[APP_MODULE.EDIT].dataset.groups.filter(
      (item) => (item.id === base_trans.direction)
    )[0].groupvalue
    let direction_id = base_trans.direction;  
    if (typeof options.new_transtype !== "undefined" && typeof options.new_direction !== "undefined") {
      transtype = options.new_transtype;
      const audit = data[APP_MODULE.LOGIN].data.audit.filter(item => (
        (item.nervatypeName === "trans") && (item.subtypeName === transtype)))[0]
      if (typeof audit !== "undefined") {
        if (audit.inputfilterName==="disabled"){
          showToast(TOAST_TYPE.INFO, `${msg("", { id: "msg_editor_invalid" })} ${transtype}`)
          return false;
        }
      }
      transtype_id = data[APP_MODULE.EDIT].dataset.groups.filter(
        (item) => ((item.groupname === "transtype") && (item.groupvalue === transtype))
      )[0].id
      direction = options.new_direction;
      direction_id = data[APP_MODULE.EDIT].dataset.groups.filter(
        (item) => ((item.groupname === "direction") && (item.groupvalue === direction))
      )[0].id
    }

    // to check some things...
    if ((transtype==="receipt" || transtype==="worksheet") && direction === "in") {
      showToast(TOAST_TYPE.INFO, `${msg("", { id: "msg_input_invalid" })} in`)
      return false;
    }
    if (base_trans.transcast==="cancellation") {
      showToast(TOAST_TYPE.INFO, msg("", { id: "msg_create_cancellation_err1" }))
      return false;
    }
    if (options.transcast==="cancellation" && ["invoice", "receipt"].includes(transtype) && base_trans.deleted===0) {
      showToast(TOAST_TYPE.INFO, msg("", { id: "msg_create_cancellation_err2" }))  
      return false;
    }
    if (options.transcast==="cancellation" && (data[APP_MODULE.EDIT].dataset.cancel_link.length > 0)) {
      showToast(TOAST_TYPE.INFO, `${msg("", { id: "msg_create_cancellation_err3" })} ${data[APP_MODULE.EDIT].dataset.cancel_link[0].transnumber}`)
      return false;
    }
    if (options.transcast==="amendment" && base_trans.deleted===1) {
      showToast(TOAST_TYPE.INFO, msg("", { id: "msg_create_amendment_err" }))
      return false;
    }

    // creat trans data from the original          
    const values = { 
      id: null, 
      transtype: transtype_id, 
      transnumber: null, 
      crdate: new Date().toISOString().split("T")[0], 
      transdate: new Date().toISOString().split("T")[0], 
      duedate: null,
      customer_id: base_trans.customer_id, 
      employee_id: base_trans.employee_id,
      department: base_trans.department, 
      project_id: base_trans.project_id,
      place_id: base_trans.place_id, 
      paidtype: base_trans.paidtype, 
      curr: base_trans.curr,
      notax: base_trans.notax, 
      paid: 0, 
      acrate: base_trans.acrate, 
      notes: base_trans.notes,
      intnotes: base_trans.intnotes, 
      fnote: base_trans.fnote,
      transtate: data[APP_MODULE.EDIT].dataset.transtate.filter(
        item => ((item.groupname === "transtate") && (item.groupvalue === "ok")))[0].id,
      closed: 0, deleted: 0, 
      direction: direction_id, 
      cruser_id: data[APP_MODULE.LOGIN].data.employee.id,
      trans_transcast: options.transcast || "normal"
    }
    if (base_trans.duedate !== null) {
      values.duedate = `${new Date().toISOString().split("T")[0]}T00:00:00`;
    }
    if (transtype === "invoice" && direction === "out") {
      const default_deadline = data[APP_MODULE.EDIT].dataset.settings.filter((group)=> (group.fieldname === "default_deadline"))[0]
      if (typeof default_deadline !== "undefined") {
        const today = new Date()
        values.duedate = `${new Date(today.setDate((today.getDate() + parseInt(default_deadline.value,10)))).toISOString().split("T")[0]}T00:00:00`
      }
    }
    if (transtype === "receipt") {
      values.customer_id = null;
    }
    const _refnum = check_refnumber({
      transtype, transcast: options.transcast, direction, 
      cmdtype: options.cmdtype, refno: options.refno})
    if((_refnum === "refnumber") || (_refnum === "reflink")){
      values.ref_transnumber = base_trans.transnumber
    }
    
    let nextnumber = `${transtype}_${direction}`;
    if (transtype === "waybill" || transtype === "cash") {
      nextnumber = transtype;
    }
    const params = { method: "POST", 
      data: {
        key: "nextNumber",
        values: {
          numberkey: nextnumber, 
          step: true
        }
      }
    }
    let result = await requestData("/function", params)
    if(result.error){
      resultError(result)
      return null
    }
    if (options.transcast === "cancellation") {
      values.transnumber = `${result}/C`;
      if (transtype !== "delivery" && transtype !== "inventory") {
        values.deleted = 1;
      }
      values.transdate = base_trans.transdate;
      values.duedate = base_trans.duedate;
    } else if (options.transcast === "amendment") {
      values.transnumber = `${result}/A`;
    } else {
      values.transnumber = result
    }

    result = await requestData("/trans", { method: "POST", data: [values] })
    if(result.error){
      resultError(result)
      return null
    }
    values.id = result[0];

    let fieldvalue = [];
    data[APP_MODULE.EDIT].current.fieldvalue.forEach((cfield) => {
      if((cfield.fieldname !== "trans_transcast") && (transtype !== "invoice")){
        const deffield = data[APP_MODULE.EDIT].dataset.deffield.filter(
          item => (item.fieldname === cfield.fieldname)
        )[0]
        const subtype = checkSubtype("trans", deffield.subtype, values);
        if ((cfield.deleted===0 && deffield.visible===1 && subtype) 
          || (cfield.deleted===0 && subtype && options.cmdtype === "copy")){
          const field = this.tableValues("fieldvalue", cfield)
          field.id = null; 
          field.ref_id = values.id;
          fieldvalue = [...fieldvalue, field]
        } 
      }
    })

    if(fieldvalue.length > 0){
      result = await requestData("/fieldvalue", { method: "POST", data: fieldvalue })
      if(result.error){
        resultError(result)
        return null
      }
    }

    if((_refnum === "link") || (_refnum === "reflink")){
      const link = {
        ...initItem({tablename: "link"}),
        nervatype_1: data[APP_MODULE.EDIT].dataset.groups.filter(
          (item)=>((item.groupname === "nervatype") && (item.groupvalue === "trans")))[0].id,
        ref_id_1: values.id,
        nervatype_2: data[APP_MODULE.EDIT].dataset.groups.filter(
          (item)=>((item.groupname === "nervatype") && (item.groupvalue === "trans")))[0].id,
        ref_id_2: base_trans.id
      }
      result = await requestData("/link", { method: "POST", data: [link] })
      if(result.error){
        resultError(result)
        return null
      }
    }

    let items = [];
    if (transtype==="invoice" || transtype==="receipt") {
      
      const get_product_qty = (_items, product_id, deposit) => {
        let retvalue = 0;
        _items.forEach((item) => {
          if ((item.product_id === product_id) && (item.deposit === deposit)){
            retvalue += item.qty;
          }
        });
        return retvalue;
      }
      
      const recalc_item = (_item, rate, digit) => {
        const item = _item
        item.netamount = roundAmount(item.fxprice*(1-item.discount/100)*item.qty, digit)
        item.vatamount = roundAmount(item.fxprice*(1-item.discount/100)*item.qty*rate, digit)
        item.amount = roundAmount(item.netamount+item.vatamount, digit)
        return item;
      }
      
      const products = {};
      if (options.from_inventory && data[APP_MODULE.EDIT].dataset.transitem_invoice) {
        // create from order,worksheet and rent, on base the delivery rows
        data[APP_MODULE.EDIT].dataset.transitem_shipping.forEach(inv_item => {
          const item = data[APP_MODULE.EDIT].dataset.item.filter(
            oitem => (`${oitem.id}-${oitem.product_id}` === inv_item.id)
          )[0]
          if (typeof item!=="undefined") {
            const dir = data[APP_MODULE.EDIT].dataset.groups.filter(
              group => (group.id === base_trans.direction))[0].groupvalue
            let iqty = (dir === "out") ? -parseFloat(inv_item.sqty) : parseFloat(inv_item.sqty);
            if (item.deleted===0 && iqty>0) {
              if (!Object.keys(products).includes(String(item.product_id))){
                iqty -= get_product_qty(data[APP_MODULE.EDIT].dataset.transitem_invoice, 
                  item.product_id, 0);
                products[item.product_id] = true;
              }
              if (iqty !== 0){
                let sitem = this.tableValues("item", item)
                sitem.qty = iqty;
                sitem = recalc_item(sitem, item.rate, base_trans.digit);
                items = [...items, sitem]
              }
            }
          }
        });
      } else {
        data[APP_MODULE.EDIT].dataset.item.forEach(base_item => {
          if (base_item.deleted===0) {
            if (options.netto_qty && data[APP_MODULE.EDIT].dataset.transitem_invoice) {
              // create from order,worksheet and rent, on base the invoice rows
              let iqty = base_item.qty;
              if (!Object.keys(products).includes(String(base_item.product_id))){
                iqty -= get_product_qty(data[APP_MODULE.EDIT].dataset.transitem_invoice, 
                  base_item.product_id, 0);
                products[base_item.product_id] = true;
              }
              if (iqty !== 0){
                let sitem = this.tableValues("item", base_item)
                sitem.qty = iqty;
                sitem = recalc_item(sitem, base_item.rate, base_trans.digit);
                items = [...items, sitem]
              }
            } else {
              items = [...items, this.tableValues("item", base_item)]
            }
          }
        });
      }
              
      // put to deposit rows
      const _items = [...items]
      _items.forEach(item => {
        if (item.deposit === 1) {
          const dqty = get_product_qty(data[APP_MODULE.EDIT].dataset.transitem_invoice, 
            item.product_id, 1);
          if (dqty !== 0) {
            const sitem = this.tableValues("item", item)
            sitem.qty = -dqty;
            items = [sitem, ...items]
          }
        }
      });
    } else {
      data[APP_MODULE.EDIT].dataset.item.forEach(item => {
        if (item.deleted===0) {
          items = [...items, this.tableValues("item", item)]
        }
      });
    }
    
    const _items = [...items]
    _items.forEach(_item => {
      const item = _item
      item.id = null;
      item.trans_id = values.id;
      item.ownstock = 0;
      if (transtype!=="invoice" && transtype!=="receipt"){
        item.deposit = 0;
      }
      if (options.transcast === "cancellation") {
        item.qty = -item.qty;
        item.netamount = -item.netamount;
        item.vatamount = -item.vatamount;
        item.amount = -item.amount;
      }
      if (options.transcast==="amendment") {
        const sitem = this.tableValues("item", item)
        sitem.qty = -sitem.qty;
        sitem.netamount = -sitem.netamount;
        sitem.vatamount = -sitem.vatamount;
        sitem.amount = -sitem.amount;
        items = [...items, sitem]
      }
    });

    if (items.length > 0) {
      result = await requestData("/item", { method: "POST", data: items })
      if(result.error){
        resultError(result)
        return null
      }
    }

    let payments = [];
    data[APP_MODULE.EDIT].dataset.payment.forEach((base_payment) => {
      if (base_payment.deleted===0) {
        const payment = this.tableValues("payment", base_payment)
        payment.id = null;
        payment.trans_id = values.id;
        if (options.transcast === "cancellation") {
          payment.amount = -payment.amount;
        }
        payments = [...payments, payment]
      }
    });
    if (payments.length > 0) {
      result = await requestData("/payment", { method: "POST", data: payments })
      if(result.error){
        resultError(result)
        return null
      }
    }

    let movements = []; const reflinks = [];
    if (transtype === "formula" || transtype === "production") {
      const movement = this.tableValues("movement", data[APP_MODULE.EDIT].dataset.movement_head[0])
      movement.id = null; 
      movement.trans_id = values.id;
      movements = [...movements, movement]
    }
    const base_movements = data[APP_MODULE.EDIT].dataset.movement || [];
    base_movements.forEach((bmt) => {
      if (bmt.deleted === 0) {
        if(bmt.item_id || bmt.ref_id){
          reflinks.push({
            id:bmt.id, 
            item_id: bmt.item_id, 
            ref_id: bmt.ref_id 
          });
        }
        const movement = this.tableValues("movement", bmt)
        movement.id = null; 
        movement.trans_id = values.id;
        if (options.transcast==="cancellation") {
          movement.qty = -movement.qty;
        }
        movements = [...movements, movement]
      }
    });
    if (movements.length > 0) {
      result = await requestData("/movement", { method: "POST", data: movements })
      if(result.error){
        resultError(result)
        return null
      }
      let links = [];
      const nt_movement = data[APP_MODULE.EDIT].dataset.groups.filter(
        (item)=>((item.groupname === "nervatype") && (item.groupvalue === "movement")))[0].id
      const nt_item = data[APP_MODULE.EDIT].dataset.groups.filter(
        (item)=>((item.groupname === "nervatype") && (item.groupvalue === "item")))[0].id
      for (let li=0; li < reflinks.length; li += 1) {
        const ilink = {...initItem({tablename: "link"}),
          nervatype_1: nt_movement
        }
        if (reflinks[li].item_id !== null) { 
          ilink.ref_id_1 = result[li].id;
          ilink.nervatype_2 = nt_item;
          ilink.ref_id_2 = reflinks[li].item_id;
          links = [...links, ilink]
        } else {
          ilink.ref_id_1 = result[data[APP_MODULE.EDIT].dataset.movement.findIndex(
            item => (item.id === reflinks[li].ref_id)
          )].id
          ilink.nervatype_2 = nt_movement;
          ilink.ref_id_2 = result[data[APP_MODULE.EDIT].dataset.movement.findIndex(
            item => (item.id === reflinks[li].id)
          )].id
          links = [...links, ilink]
        }
      }
      if (links.length > 0) {
        result = await requestData("/link", { method: "POST", data: links })
        if(result.error){
          resultError(result)
          return null
        }
      }
    }

    await createHistory("save")
    return this.loadEditor({ ntype: "trans", ttype: transtype, id: values.id })
  }

  createTransOptions() {
    const { modalTrans } = this.module
    const { setData } = this.store
    const edit = {...this.store.data[APP_MODULE.EDIT]}
    const options = {
      directions: ["in","out"],
      baseTranstype: edit.current.transtype,
      transtype: edit.current.transtype,
      direction: edit.dataset.groups.filter((group)=> (group.id === edit.current.item.direction))[0].groupvalue,
      elementCount: parseInt(edit.dataset.element_count[0].pec,10),
      doctypes: ["order","worksheet","rent","invoice","receipt"],
      refno: true, netto: true, from: false,
      nettoDiv: false, fromDiv: false
    }
    
    switch (options.transtype) {
      case "offer":
        options.doctypes = ["offer","order","worksheet","rent"];
        options.transtype = "order";
        options.nettoDiv = false;
        options.fromDiv = false;
        break;
      case "order":
        options.doctypes = ["offer","order","worksheet","rent","invoice","receipt"];
        options.transtype = "invoice";
        options.nettoDiv = true;
        if (options.elementCount===0) {
          options.fromDiv = true;
        } else {
          options.fromDiv = false;
        }
        break;
      case "worksheet":
        options.doctypes = ["offer","order","worksheet","rent","invoice","receipt"];
        options.transtype = "invoice";
        options.nettoDiv = true;
        if (options.elementCount===0) {
          options.fromDiv = true;
        } else {
          options.fromDiv = false;
        }
        break;
      case "rent":
        options.doctypes = ["offer","order","worksheet","rent","invoice","receipt"];
        options.transtype = "invoice";
        options.nettoDiv = true;
        if (options.elementCount===0) {
          options.fromDiv = true;
        } else {
          options.fromDiv = false;
        }
        break;
      case "invoice":
        options.doctypes = ["order","worksheet","rent","invoice","receipt"];
        options.transtype = "order";
        options.nettoDiv = false;
        options.fromDiv = false;
        break;
      default:
        options.transtype = "order";
    }
    const modalForm = modalTrans({
      ...options,
      onEvent: {
        onModalEvent: async ({ key, data }) => {
          setData("current", { modalForm: null })
          if(key === MODAL_EVENT.OK){
            this.createTrans({
              cmdtype: "create", transcast: "normal", 
              new_transtype: data.newTranstype, 
              new_direction: data.newDirection, 
              refno: data.refno, 
              from_inventory: data.fromInventory, 
              netto_qty: data.nettoQty
            })
          }
        }
      }
    })
    setData("current", { modalForm, side: SIDE_VISIBILITY.HIDE })
  }

  deleteEditor() {
    const { inputBox } = this.host
    const { data, setData } = this.store
    const { sql } = this.app.modules
    const { resultError, requestData, createHistory, getSql, showToast, msg, currentModule } = this.app
    const clearEditor = () => {
      setData(APP_MODULE.EDIT, { dataset: {}, current: {}, dirty: false, form_dirty: false })
      currentModule({ data: { module: APP_MODULE.SEARCH } })
    }
    const deleteData = async () => {
      const result = await requestData(`/${data[APP_MODULE.EDIT].current.type}`, 
        { method: "DELETE", query: { id: data[APP_MODULE.EDIT].current.item.id } })
      if(result && result.error){
        return resultError(result)
      }
      await createHistory("delete")
      return clearEditor()
    }
    const  modalForm = inputBox({ 
      title: msg("", { id: "msg_warning" }),
      message: msg("", { id: "msg_delete_text" }),
      infoText: msg("", { id: "msg_delete_info" }),
      onEvent: {
        onModalEvent: async (modalResult) => {
          setData("current", { modalForm: null })
          if (modalResult.key === MODAL_EVENT.OK) {
            if (data[APP_MODULE.EDIT].current.item.id === null) {
              return clearEditor()
            }
            if (typeof sql[data[APP_MODULE.EDIT].current.type].delete_state !== "undefined") {
              const sqlInfo = getSql(data[APP_MODULE.LOGIN].data.engine, sql[data[APP_MODULE.EDIT].current.type].delete_state())
              const params = { 
                method: "POST", 
                data: [{ 
                  key: "state",
                  text: sqlInfo.sql,
                  values: Array(sqlInfo.prmCount).fill(data[APP_MODULE.EDIT].current.item.id)
                }]
              }
              const view = await requestData("/view", params)
              if(view.error){
                return resultError(view)
              }
              if (view.state[0].sco > 0) {
                showToast(TOAST_TYPE.ERROR, msg("", { id: "msg_integrity_err" }))
              } else {
                deleteData()
              }
            } else {
              deleteData()
            }
          }
          return true
        }
      }
    })
    return setData("current", { modalForm, side: SIDE_VISIBILITY.HIDE })
  }

  deleteEditorItem(params) {
    const { inputBox } = this.host
    const { data, setData } = this.store
    const { resultError, requestData, createHistory, msg } = this.app
    const reLoad = () => {
      this.loadEditor({
        ntype: data[APP_MODULE.EDIT].current.type, 
        ttype: data[APP_MODULE.EDIT].current.transtype, 
        id: data[APP_MODULE.EDIT].current.item.id, 
        form: params.fkey
      })
    }
    const deleteItem = async () => {
      if (params.id === null) {
        return reLoad()
      }
      const table = (!params.table) ? params.fkey : params.table
      const result = await requestData(
        `/${table}`, { method: "DELETE", query: { id: params.id } })
      if(result && result.error){
        return resultError(result)
      }
      if(params.callback){
        params.callback()
        return true
      }
      await createHistory("delete")
      return reLoad()
    }

    if(params.prompt){
      return deleteItem()
    }
    const  modalForm = inputBox({ 
      title: msg("", { id: "msg_warning" }),
      message: msg("", { id: "msg_delete_text" }),
      infoText: msg("", { id: "msg_delete_info" }),
      onEvent: {
        onModalEvent: async (modalResult) => {
          setData("current", { modalForm: null })
          if (modalResult.key === MODAL_EVENT.OK) {
            deleteItem()
          }
        }
      }
    })
    return setData("current", { modalForm, side: SIDE_VISIBILITY.HIDE })
  }

  async editItem(options) {
    const { inputBox } = this.host
    const { data, setData } = this.store
    const { resultError, requestData, msg } = this.app
    const loadPrice = async (trans, item) => {
      const params = { method: "POST", 
        data: {
          key: "getPriceValue",
          values: {
            vendorprice: item.vendorprice, 
            product_id: item.product_id,
            posdate: trans.transdate, 
            curr: trans.curr, 
            qty: item.qty, 
            customer_id: trans.customer_id
          }
        }
      }
      return requestData("/function", params)
    }

    let edit = {...data[APP_MODULE.EDIT]}
    if(["fieldvalue_value", "fieldvalue_notes", "fieldvalue_deleted"].includes(options.name)){
      const fieldvalue_idx = edit.current.fieldvalue.findIndex((item)=>(item.id === options.id))
      if((fieldvalue_idx > -1) && ["all", "update"].includes(edit.audit)){
        edit = {...edit,
          dirty: true,
          current: {
            ...edit.current,
            fieldvalue: [
              ...edit.current.fieldvalue
            ]
          }
        }
        edit.current.fieldvalue[fieldvalue_idx] = {
          ...edit.current.fieldvalue[fieldvalue_idx],
          [options.name.split("_")[1]]: (options.name === "fieldvalue_deleted") ? 1 : options.value.toString()
        }
      }
    } else if (edit.current.form) {
      edit = {...edit,
        form_dirty: true
      }
      if (typeof edit.current.form[options.name] !== "undefined") {
        edit = {...edit, current: {...edit.current, form: {...edit.current.form,
          [options.name]: options.value
        }}}
      }
      switch (edit.current.form_type) {
        case "item":
          if (options.name === "product_id" && (typeof options.item !== "undefined")) {
            edit = {...edit, current: {...edit.current, form: {...edit.current.form,
              description: options.item.description,
              unit: options.item.unit,
              tax_id: parseInt(options.item.tax_id,10)
            }}}
            if (edit.current.form.qty === 0) {
              edit = {...edit, current: {...edit.current, form: {...edit.current.form,
                qty: 1
              }}}
            }
            const price = await loadPrice(edit.current.item, edit.current.form)
            if(price.error){
              return resultError(price)
            }
            edit = {...edit, current: {...edit.current, form: {...edit.current.form,
              fxprice: !Number.isNaN(parseFloat(price.price)) ? parseFloat(price.price) : 0,
              discount: !Number.isNaN(parseFloat(price.discount)) ? parseFloat(price.discount) : 0
            }}}
            edit = {...edit, current: {...edit.current,
              form : this.calcPrice("fxprice", edit.current.form)
            }}
          } else {
            switch(options.name) {
              case "qty":
                if (parseFloat(edit.current.form.fxprice) === 0) {
                  const price = await loadPrice(edit.current.item, edit.current.form)
                  if(price.error){
                    return resultError(price)
                  }
                  edit = {...edit, current: {...edit.current, form: {...edit.current.form,
                    fxprice: !Number.isNaN(parseFloat(price.price)) ? parseFloat(price.price) : 0,
                    discount: !Number.isNaN(parseFloat(price.discount)) ? parseFloat(price.discount) : 0
                  }}}
                }
                if(options.event_type === "blur"){
                  edit = {...edit, current: {...edit.current,
                    form : this.calcPrice("fxprice", edit.current.form)
                  }}
                }
                break;
              case "fxprice":
              case "tax_id":
              case "discount":
                if((options.event_type === "blur") || (options.name === "tax_id")){
                  edit = {...edit, current: {...edit.current,
                    form : this.calcPrice("fxprice", edit.current.form)
                  }}
                }
                break;
              case "amount":
                if(options.event_type === "blur"){
                  edit = {...edit, current: {...edit.current,
                    form : this.calcPrice("amount", edit.current.form)
                  }}
                }
                break;
              case "netamount":
                if(options.event_type === "blur"){
                  edit = {...edit, current: {...edit.current,
                    form : this.calcPrice("netamount", edit.current.form)
                  }}
                }
                break;
              default:
                break;
            }
          }
          break;
        
        case "price":
        case "discount":
          if (options.name === "customer_id") {
            edit = {...edit, current: {...edit.current,
              price_customer_id: options.value
            }}
          }
          break;

        case "invoice_link":
          if (options.name === "ref_id_1" && (typeof options.item !== "undefined")) {
            edit = {...edit, current: {...edit.current,
              price_customer_id: options.value,
              invoice_link: [...edit.current.invoice_link]
            }}
            edit.current.invoice_link[0] = {...edit.current.invoice_link[0],
              curr: options.item.curr
            }
          } else if ((options.name === "link_qty") || (options.name === "link_rate")) {
            edit = {...edit, current: {...edit.current,
              invoice_link_fieldvalue: this.setFieldvalue(edit.current.invoice_link_fieldvalue, 
                options.name, edit.current.form.id, null, options.value)
            }}
          }
          break;

        case "payment_link":
          if (options.name === "ref_id_2" && (typeof options.item !== "undefined")) {
            edit = {...edit, current: {...edit.current,
              payment_link: [...edit.current.payment_link]
            }}
            edit.current.payment_link[0] = {...edit.current.payment_link[0],
              curr: options.item.curr
            }
          } else if ((options.name === "link_qty") || (options.name === "link_rate")) {
            edit = {...edit, current: {...edit.current,
              payment_link_fieldvalue: this.setFieldvalue(edit.current.payment_link_fieldvalue, 
                options.name, edit.current.form.id, null, options.value)
            }}
          }
          break;

        default:
          break;
      }
    } else {
      if ((typeof edit.current.item[options.name] !== "undefined") && (options.extend === false)) {
        edit = {...edit, current: {...edit.current, item: {...edit.current.item,
          [options.name]: options.value
        }}}
        if(options.label_field){
          edit = {...edit, current: {...edit.current, item: {...edit.current.item,
            [options.label_field]: options.refnumber || null
          }}}
        }
      } else if ((typeof edit.template.options.extend !== "undefined") && (options.extend === true)) {
        edit = {...edit, current: {...edit.current, extend: {...edit.current.extend,
          [options.name]: options.value
        }}}
      }
      if((edit.audit==="all") || (edit.audit==="update")){
        edit = {...edit,
          dirty: true
        }
      }

      switch (edit.current.type){
        case "report":
          edit = {...edit,
            dirty: false, 
            current: {...edit.current,
              fieldvalue: [...edit.current.fieldvalue]
            }
          }
          if(options.name === "selected"){
            const fieldvalue_idx = edit.current.fieldvalue.findIndex((item)=>(item.id === options.id))
            if(fieldvalue_idx > -1){
              edit.current.fieldvalue[fieldvalue_idx] = {...edit.current.fieldvalue[fieldvalue_idx],
                selected: options.value
              }
            }
          } else {
            const fieldvalue_idx = edit.current.fieldvalue.findIndex((item)=>(item.name === options.name))
            if(fieldvalue_idx > -1){
              edit.current.fieldvalue[fieldvalue_idx] = {...edit.current.fieldvalue[fieldvalue_idx],
                value: options.value
              }
            }
          }
          break;

        case "printqueue":
          edit = {...edit,
            dirty: false
          }
          edit = {...edit, printqueue: {...edit.printqueue,
            [options.name]: options.value
          }}
          break;

        case "trans":
          switch (options.name) {
            case "closed":
              if (options.value === 1) {
                const  modalForm = inputBox({ 
                  title: msg("", { id: "msg_warning" }),
                  message: msg("", { id: "msg_close_text" }),
                  infoText: msg("", { id: "msg_delete_info" }),
                  onEvent: {
                    onModalEvent: async (modalResult) => {
                      setData("current", { modalForm: null })
                      if (modalResult.key === MODAL_EVENT.OK) {
                        edit = {...edit, current: {...edit.current,
                          closed: 1
                        }}
                      }
                      if (modalResult.key === MODAL_EVENT.CANCEL) {
                        edit = {...edit, current: {...edit.current, item: {...edit.current.item,
                          [options.name]: 0
                        }}}
                      }
                      setData(APP_MODULE.EDIT, edit)
                    }
                  }
                })
                return setData("current", { modalForm })
              }
              break;
            case "paiddate":
              edit = {...edit, current: {...edit.current, item: {...edit.current.item,
                transdate: options.value
              }}}
              break;
            case "direction":
              if(edit.current.transtype === "cash"){
                const direction = edit.dataset.groups.filter((item)=>(item.id === options.value))[0].groupvalue
                edit = {...edit, template: {...edit.template, options: {...edit.template.options,
                  opposite: (direction === "out")
                }}}
              }
              break;        
            case "seltype":
              edit = {...edit, current: {...edit.current, 
                extend: {...edit.current.extend,
                  seltype: options.value,
                  ref_id: null,
                  refnumber: ""
                },
                item: {...edit.current.item,
                  customer_id: null,
                  employee_id: null,
                  ref_transnumber: null
                }
              }}
              break;
            case "ref_id":
              edit = {...edit, current: {...edit.current, extend: {...edit.current.extend,
                refnumber: options.refnumber,
                ntype: (edit.current.extend.seltype === "transitem") ? "trans" : edit.current.extend.seltype,
                transtype: (options.item && options.item.transtype) ? options.item.transtype.split("-")[0] : ""
              }}}
              switch (edit.current.extend.seltype){
                case "customer":
                  edit = {...edit, current: {...edit.current, item: {...edit.current.item,
                    customer_id: options.value
                  }}}
                  break;
                case "employee":
                  edit = {...edit, current: {...edit.current, item: {...edit.current.item,
                    employee_id: options.value
                  }}}
                  break;
                case "transitem":
                  edit = {...edit, current: {...edit.current, item: {...edit.current.item,
                    ref_transnumber: options.refnumber,
                  }}}
                  break;
                default:
                  break;}
              break;
            case "trans_wsdistance":
            case "trans_wsrepair":
            case "trans_wstotal":
            case "trans_reholiday":
            case "trans_rebadtool":
            case "trans_reother":
            case "trans_wsnote":
            case "trans_rentnote":
              edit = {...edit, current: {...edit.current,
                fieldvalue: this.setFieldvalue(edit.current.fieldvalue, 
                  options.name, edit.current.item.id, null, options.value)
              }}
              break;
            case "shippingdate":
            case "shipping_place_id":
              edit = {...edit,
                dirty: false
              }
              edit = {...edit, current: {...edit.current,
                [options.name]: options.value
              }}
              break;
            case "fnote":
              edit = {...edit, current: {...edit.current, item: {...edit.current.item,
                fnote: options.value
              }}}
              break;
            default:
              break;
          }
          break;
        default:
          break;
      }
    }
    return setData(APP_MODULE.EDIT, edit)
  }

  exportEvent() {
    const { saveToDisk } = this.app
    const { dataset, current } = this.store.data[APP_MODULE.EDIT]
    const event = {...current.item};
    const getCalDateTime = (value) => `${new Date(value).toISOString().slice(0,10)}T${new Date(value).toLocaleTimeString("en",{hour12: false}).replace("24","00")}`.replaceAll("-","").replaceAll(":","")
    let eventStr = 
    `BEGIN:VCALENDAR\nPRODID:-//nervatura.com/NONSGML Nervatura Calendar//EN\nVERSION:2.0\nBEGIN:VEVENT\nUID:${
      (event.uid !== null) ? event.uid : `${Math.random().toString(16).slice(2)}-${Math.random().toString(16).slice(2)}`}`
    if (event.fromdate !== null){
      eventStr+=`\nDTSTART:${getCalDateTime(event.fromdate)}`
    }
    if (event.todate !== null){
      eventStr+=`\nDTEND:${getCalDateTime(event.todate)}`
    }
    if (event.subject !== null){
      eventStr+=`\nSUMMARY:${event.subject}`
    }
    if (event.place !== null){
      eventStr+=`\nLOCATION:${event.place}`
    }
    if (event.description !== null){
      eventStr+=`\nDESCRIPTION:${event.description}`
    }
    if (event.eventgroup !== null){
      const eventgroup = dataset.eventgroup.filter(item => (item.id === event.eventgroup))[0]
      if (typeof eventgroup !== "undefined") {
        eventStr+=`\nCATEGORY:${eventgroup.groupvalue}`
      }
    }
    eventStr+=`\nEND:VEVENT\nEND:VCALENDAR`

    const filename = `${event.calnumber.replace(/\//g, "_")}.ics`;
    const icsUrl = URL.createObjectURL(new Blob([eventStr], 
      {type : 'text/ics;charset=utf-8;'}));
    saveToDisk(icsUrl, filename);
  }

  exportQueueAll() {
    const { inputBox } = this.host
    const { showToast, msg } = this.app
    const { setData } = this.store
    const { dataset, current } = this.store.data[APP_MODULE.EDIT]
    const options = {...current.item}
    if (dataset.items.length > 0){
      if (options.mode === "print") {
        return showToast(TOAST_TYPE.ERROR, `${msg("", { id: "ms_export_invalid" })} ${msg("", { id: "printqueue_mode_print" })}`)
      }
      const  modalForm = inputBox({ 
        title: msg("", { id: "msg_warning" }),
        message: msg("", { id: "label_export_all_selected" }),
        infoText: `${msg("", { id: "msg_delete_info" })} ${msg("", { id: "ms_continue_warning" })}`,
        defaultOK: true,
        onEvent: {
          onModalEvent: async (modalResult) => {
            setData("current", { modalForm: null })
            if (modalResult.key === MODAL_EVENT.OK) {
              let result = true
              for (let index = 0; ((index < dataset.items.length) && result); index += 1) {
                // eslint-disable-next-line no-await-in-loop
                result = await this.exportQueue(dataset.items[index], ()=>{})
              }
              this.searchQueue()
            }
          }
        }
      })
      return setData("current", { modalForm })
    }
    return true
  }

  async exportQueue(item, callback) {
    const { current } = this.store.data[APP_MODULE.EDIT]
    const options = {...current.item}
    let result = await this.reportOutput({
      type: options.mode, 
      template: item.reportkey, 
      title: item.refnumber,
      orient: options.orientation, 
      size: options.size, 
      copy: item.copies,
      nervatype: item.typename,
      id: item.ref_id
    })
    if(result){
      result = this.deleteEditorItem({
        fkey: "items", table: "ui_printqueue", id: item.id, prompt: true, 
        callback: (callback) || this.searchQueue
      })
    }
    return result
  }

  getTransFilter(_sql, values) {
    const { data } = this.store
    switch (data[APP_MODULE.LOGIN].data.transfilterName) {
      case "usergroup":
        _sql.where.push(
          ["and","cruser_id","in",[{
            select:["id"], from:"employee", 
            where:["usergroup","=","?"]
          }]])
        values.push(data[APP_MODULE.LOGIN].data.employee.usergroup)
        break;
      case "own":
        _sql.where.push(
          ["and","cruser_id","=","?"]
        )
        values.push(data[APP_MODULE.LOGIN].data.employee.id)
        break;
      default:
        break;
    }
    return [_sql, values]
  }

  async loadEditor(params) {
    const { forms, dataSet, sql, initItem } = this.app.modules
    const { getSql, requestData, resultError, getSetting } = this.app
    const { setData, data } = this.store
    const { ntype, ttype, id } = params;
    let edit = {
      dataset: { },
      current: { type: ntype, transtype: ttype },
      dirty: false,
      form_dirty: false
    }
    let proitem;
    if (id===null) {
      proitem = initItem({tablename: ntype, transtype: ttype, 
        dataset: edit.dataset, current: edit.current});
    };
    let views = []
    dataSet[ntype](ttype).forEach(info => {
      let _sql = {}
      if(info.infoType === "table"){
        _sql = { select:["*"], from: info.classAlias }
        if(info.where){
          _sql.where = info.where
        }
        if(info.order){
          _sql.order_by = info.order
        }
      } else if (typeof sql[ntype][info.sqlKey] !== "undefined") {
          _sql = sql[ntype][info.sqlKey](ntype)
        } else if (typeof sql[ntype][info.infoName] !== "undefined") {
          _sql = sql[ntype][info.infoName](ntype)
        } else {
          _sql = sql.all[info.infoName](ntype)
        }
      const sqlInfo = getSql(data[APP_MODULE.LOGIN].data.engine, _sql)
      if( (id !== null) || (sqlInfo.prmCount === 0)){
        views = [...views, {
          key: info.infoName,
          text: sqlInfo.sql,
          values: ((sqlInfo.prmCount>0) && (id !== null)) ? Array(sqlInfo.prmCount).fill(id) : []
        }]
      } else {
        edit = {
          ...edit,
          dataset: {
            ...edit.dataset,
            [info.infoName]: []
          }
        }
      }
    })

    if (ntype !== "report") {
      dataSet.report().forEach(info => {
        const view = {
          key: info.infoName,
          sql: sql.report[info.infoName](ntype),
          values: []
        }
        if(ntype !== "printqueue"){
          view.values = [...view.values, 
            data[APP_MODULE.LOGIN].data.employee.usergroup, 
            edit.current.type
          ]
        }
        if (edit.current.type ==="trans") {
          const _where = ["and","r.transtype","=",[[],{select:["id"], from:"groups", 
              where:[["groupname","=","'transtype'"],["and","groupvalue","=","?"]]}]]
          view.sql.where = [...view.sql.where, _where]
          view.values = [...view.values, edit.current.transtype]
        }
        views = [...views,{
          key: view.key,
          text: getSql(data[APP_MODULE.LOGIN].data.engine, view.sql).sql,
          values: view.values
        }]
      })
    }

    const options = { method: "POST", data: views }
    const view = await requestData("/view", options)
    if(view.error){
      return resultError(view)
    }
    edit = {
      ...edit,
      dataset: {
        ...edit.dataset,
        ...view
      }
    }
    if (id===null) {
      if (proitem === null) {
        proitem = initItem({tablename: ntype, transtype: ttype, 
          dataset: edit.dataset, current: edit.current});
      }
      if (ttype === "delivery") {
        proitem = {
          ...proitem,
          direction: edit.dataset.groups.filter((group)=> ((group.groupname === "direction") && (group.groupvalue === "transfer")))[0].id
        }
      }
      edit = {
        ...edit,
        dataset: {
          ...edit.dataset,
          [ntype]: [proitem]
        }
      }
    }
    setData(APP_MODULE.EDIT, edit)
    if (!params.cb_key || (params.cb_key === EDITOR_EVENT.SET_EDITOR)) {
      if (ntype === "trans") {
        if(params.shipping){
          return this.setEditor(params, forms.shipping(edit.dataset[ntype][0], edit, getSetting("ui")), edit)
        } 
        return this.setEditor(params, forms[ttype](edit.dataset[ntype][0], edit, getSetting("ui")), edit)
      }
      return this.setEditor(params, forms[ntype](edit.dataset[ntype][0], edit, getSetting("ui")), edit)
    }
    return true
  }

  newFieldvalue(_fieldname) {
    const { initItem } = this.app.modules
    const { requestData, resultError, showToast, msg } = this.app
    const { data } = this.store
    const updateFieldvalue = async (item) => {
      const options = { method: "POST", data: [item] }
      const result = await requestData("/fieldvalue", options)
      if(result.error){
        return resultError(result)
      }
      return this.loadEditor({
        ntype: data[APP_MODULE.EDIT].current.type, 
        ttype: data[APP_MODULE.EDIT].current.transtype, 
        id: data[APP_MODULE.EDIT].current.item.id, 
        form: "fieldvalue", form_id: result[0]
      })
    }

    if (_fieldname!=="") {
      const deffield = data[APP_MODULE.EDIT].dataset.deffield.filter((item) => (item.fieldname === _fieldname))[0]
      const _fieldtype = data[APP_MODULE.LOGIN].data.groups.filter((item) => (item.id === deffield.fieldtype))[0].groupvalue
      let item = {
        ...initItem({tablename: "fieldvalue"}),
        id: null,
        fieldname: deffield.fieldname
      }
      let _selector = false;
      switch (_fieldtype) {
        case "bool":
          item = {...item,
            value: "false"
          }
          break;

        case "date":
          item = {...item,
            value: new Date().toISOString().split("T")[0]
          }
          break;

        case "time":
          item = {...item,
            value: "00:00"
          }
          break;

        case "float":
        case "integer":
          item = {...item,
            value: "0"
          }
          break;

        case "valuelist":
          item = {...item,
            value: deffield.valuelist.split("|")[0]
          }
          break;

        default:
          break;
      }
      if(["customer", "tool", "trans", "transitem", "transmovement", "transpayment", 
        "product", "project", "employee", "place"].includes(_fieldtype)){
        _selector = true;
      }
      if (_selector) {
        this.onSelector(_fieldtype, "", (row)=>{
          const params = row.id.split("/")
          item = {...item,
            value: String(parseInt(params[2],10))
          }
          updateFieldvalue(item)
        })
      } else {
        updateFieldvalue(item)
      }
      return true
    }
    return showToast(TOAST_TYPE.ERROR, msg("", { id: "fields_deffield_missing" }))
  }

  async nextTransNumber() {
    const { getSql, requestData, resultError } = this.app
    const { data } = this.store
    const { current, dataset } = data[APP_MODULE.EDIT]
    const transtype = current.transtype
    const direction = dataset.groups.filter(
      item => (item.id === current.item.direction))[0].groupvalue
    const _sql = {
      select:["min(id) as id"], from:"trans", 
      where:[["transtype","=","?"],["and","id",">","?"]]
    }
    const values = [current.item.transtype, current.item.id]
    if(!["cash", "waybill"].includes(transtype)){
      _sql.where.push(["and","direction","=","?"])
      values.push(current.item.direction)
    }
    if(!["invoice_out", "receipt_out", "cash_out", "cash_in"].includes(`${transtype}_${direction}`)){
      _sql.where.push(["and","deleted","=","0"])
    }
    const filter = this.getTransFilter(_sql, values)
    const params = { 
      method: "POST", 
      data: [{ 
        key: "next",
        text: getSql(data.login.data.engine, filter[0]).sql,
        values: filter[1]
      }]
    }
    const view = await requestData("/view", params)
    if(view.error){
      return resultError(view)
    }
    if (view.next[0].id === null) {
      if(!["delivery_out", "delivery_in"].includes(`${transtype}_${direction}`)){
        this.checkEditor({ntype: "trans", ttype: transtype, id: null}, EDITOR_EVENT.LOAD_EDITOR)
      }
    } else {
      this.checkEditor({ntype: "trans", ttype: transtype, id: view.next[0].id}, EDITOR_EVENT.LOAD_EDITOR)
    }
    return true
  }

  async prevTransNumber() {
    const { getSql, requestData, resultError } = this.app
    const { data } = this.store
    if (data[APP_MODULE.EDIT].current.type !== "trans" || data[APP_MODULE.EDIT].current.item.id === null) {
      return true
    }
    const transtype = data[APP_MODULE.EDIT].current.transtype
    const direction = data[APP_MODULE.EDIT].dataset.groups.filter(
      item => (item.id === data[APP_MODULE.EDIT].current.item.direction))[0].groupvalue
    const _sql = {
      select:["max(id) as id"], from:"trans", where:[["transtype","=","?"], ["and","id","<","?"]]
    }
    const values = [data[APP_MODULE.EDIT].current.item.transtype, data[APP_MODULE.EDIT].current.item.id]
    if(!["cash", "waybill"].includes(transtype)){
      _sql.where.push(["and","direction","=","?"])
      values.push(data[APP_MODULE.EDIT].current.item.direction)
    }
    if(!["invoice_out", "receipt_out", "cash_out", "cash_in"].includes(`${transtype}_${direction}`)){
      _sql.where.push(["and","deleted","=","0"])
    }
    const filter = this.getTransFilter(_sql, values)
    const params = { 
      method: "POST", 
      data: [{ 
        key: "prev",
        text: getSql(data[APP_MODULE.LOGIN].data.engine, filter[0]).sql,
        values: filter[1]
      }]
    }
    const view = await requestData("/view", params)
    if(view.error){
      return resultError(view)
    }
    if (view.prev[0].id !== null){
      return this.checkEditor({ntype: "trans", ttype: transtype, id: view.prev[0].id}, EDITOR_EVENT.LOAD_EDITOR)
    }
    return true
  }

  async reportOutput(params) {
    const { requestData, resultError, saveToDisk } = this.app
    if(params.type === "printqueue"){
      return this.addPrintQueue(params.template, params.copy)
    }
    const result = await requestData(this.reportPath(params), {})
    if(result && result.error){
      resultError(result)
      return false
    }
    const resultUrl = URL.createObjectURL(result, {type : (params.type === "xml") ? "application/xml; charset=UTF-8" : "application/pdf"})
    if(params.type === "print"){
      window.open(resultUrl,"_blank")
    } else {
      let filename = `${params.title}_${new Date().toISOString().split("T")[0]}.${params.type}`
      filename = filename.split("/").join("_")
      saveToDisk(resultUrl, filename)
    }
    return true
  }

  reportPath(params) {
    const { current } = this.store.data[APP_MODULE.EDIT]
    const query = new URLSearchParams()
    query.append("reportkey", params.template)
    query.append("orientation", params.orient)
    query.append("size", params.size)
    query.append("output", params.type)
    if(params.filters){
      return `/report?${query.toString()}&${params.filters}`  
    }
    query.append("nervatype", params.nervatype||current.type)
    return `/report?${query.toString()}&filters[@id]=${params.id||current.item.id}`
  }

  reportSettings() {
    const { modalReport } = this.module
    const { getSetting } = this.app
    const { setData } = this.store
    const { current, dataset, template } = this.store.data[APP_MODULE.EDIT]
    const login = this.store.data[APP_MODULE.LOGIN].data
    const direction = (current.type === "trans") ? 
      dataset.groups.filter(
        item => (item.id === current.item.direction))[0].groupvalue : "out"
    let defTemplate = (current.type === "trans") ?
      dataset.settings.filter(
        item => (item.fieldname === `default_trans_${current.transtype}_${direction}_report`))[0] :
      dataset.settings.filter(
        item => (item.fieldname === `default_${current.type}_report`))[0]
    let templates = []
    dataset.report.forEach(tpl => {
      let audit = login.audit.filter(item => (
        (item.nervatypeName === "report") && (item.subtype === tpl.id)))[0]
      if(audit){
        audit= audit.inputfilterName
      } else {
        audit = "all"
      }
      if (audit !== "disabled") {
        if (current.type==="trans") {
          if (current.item.direction === tpl.direction) {
            templates = [...templates, {
              value: tpl.reportkey, text: tpl.repname
            }]
          }
        } else {
          templates = [...templates, {
            value: tpl.reportkey, text: tpl.repname
          }]
        }
      }
    })
    if (typeof defTemplate !== "undefined") {
      defTemplate = defTemplate.value
    } else if(templates.length > 0){
      defTemplate = templates[0].value
    } else {
      defTemplate = ""
    }
    const modalForm = modalReport({
      title: current.item[template.options.title_field],
      template: defTemplate,
      templates, 
      orient: getSetting("page_orient"),
      size: getSetting("page_size"),
      report_size: getSetting("report_size"),
      report_orientation: getSetting("report_orientation"),
      copy: 1,
      onEvent: {
        onModalEvent: (modalResult) => {
          setData("current", { modalForm: null })
          if (modalResult.key === MODAL_EVENT.OK) {
            this.reportOutput(modalResult.data)
          }
        }
      }
    })
    setData("current", { modalForm, side: SIDE_VISIBILITY.HIDE })
  }

  onEditEvent({key, data}){
    const { setData } = this.store
    const edit = this.store.data[APP_MODULE.EDIT]
    switch (key) {
      case EDIT_EVENT.CHANGE:
        setData(APP_MODULE.EDIT, {
          current: {
            ...edit.current,
            [data.fieldname]: data.value
          } 
        })
        break;

      case EDIT_EVENT.CHECK_EDITOR:
        this.checkEditor(...data)
        break;

      case EDIT_EVENT.CHECK_TRANSTYPE:
        this.checkTranstype(...data)
        break;

      case EDIT_EVENT.EDIT_ITEM:
        this.editItem(data)
        break;

      case EDIT_EVENT.SET_PATTERN:
        this.setPattern(data)
        break;

      case EDIT_EVENT.SELECTOR:
        this.onSelector(...data)
        break;

      case EDIT_EVENT.FORM_ACTION:
        this.setFormActions(data)
        break;
    
      default:
        break;
    }
  }

  async onSelector(selectorType, selectorFilter, setSelector) {
    const { modalSelector } = this.module
    const { quick } = this.app.modules
    const { quickSearch, resultError } = this.app
    const { setData } = this.store
    let  formProps = {
      view: selectorType, 
      columns: quick[selectorType]().columns,
      result: [],
      filter: selectorFilter,
      onEvent: {
        onModalEvent: async ({ key, data }) => {
          switch (key) {
            case MODAL_EVENT.CANCEL:
              setData("current", { modalForm: null })
              break;

            case MODAL_EVENT.SELECTED:
              setData("current", { modalForm: null })
              setSelector(data.value, data.filter)
              break;
          
            case MODAL_EVENT.SEARCH:
            default:
              const view = await quickSearch(selectorType, data.value)
              if(view.error){
                return resultError(view)
              }
              formProps = {...formProps,
                result: view.result
              }
              setData("current", { modalForm: modalSelector(formProps) })
              break;
          }
          return true
        }
      }
    }
    setData("current", { modalForm: modalSelector(formProps) })
  }

  async onSideEvent({key, data}){
    const { showHelp, saveBookmark } = this.app
    const { setData } = this.store
    const storeEdit = this.store.data[APP_MODULE.EDIT]
    const loginData = this.store.data[APP_MODULE.LOGIN].data
    switch (key) {
      case SIDE_EVENT.CHANGE:
        setData(APP_MODULE.EDIT, {
          [data.fieldname]: data.value 
        })
        break;

      case SIDE_EVENT.BACK:
        if(storeEdit.current.form){
          return this.checkEditor({
            ntype: storeEdit.current.type, 
            ttype: storeEdit.current.transtype, 
            id: storeEdit.current.item.id,
            form: storeEdit.current.form_type}, EDITOR_EVENT.LOAD_EDITOR)
        }
        if(storeEdit.current.form_type === "transitem_shipping"){
          return this.checkEditor({
            ntype: storeEdit.current.type, 
            ttype: storeEdit.current.transtype, 
            id: storeEdit.current.item.id,
            form: storeEdit.current.form_type}, EDITOR_EVENT.LOAD_EDITOR)
        }
        const reftype = loginData.groups.filter((item)=> (item.id === storeEdit.current.item.nervatype))[0].groupvalue
        this.checkEditor({ntype: reftype, 
          ttype: null, id: storeEdit.current.item.ref_id,
          form: storeEdit.current.type}, EDITOR_EVENT.LOAD_EDITOR)
        break;

      case SIDE_EVENT.CHECK:
        this.checkEditor(...data)
        break;

      case SIDE_EVENT.PREV_NUMBER:
        this.prevTransNumber()
        break;

      case SIDE_EVENT.NEXT_NUMBER:
        this.nextTransNumber()
        break;

      case SIDE_EVENT.SAVE:
        let edit = null
        if(storeEdit.current.form){
          edit = await this.saveEditorForm()
        } else {
          edit = await this.saveEditor()
        }
        if(edit){
          this.loadEditor({
            ntype: edit.current.type, 
            ttype: edit.current.transtype, 
            id: edit.current.item.id,
            form: edit.current.view
          })
        }
        break;

      case SIDE_EVENT.DELETE:
        if(storeEdit.current.form){
          return this.deleteEditorItem({
            fkey: storeEdit.current.form_type, 
            table: storeEdit.current.form_datatype, 
            id: storeEdit.current.form.id
          })
        }
        this.deleteEditor()
        break;

      case SIDE_EVENT.NEW:
        const { ntype, ttype } = data[0]
        if(ttype === "shipping"){
          this.onSelector("transitem_delivery", "", (row)=>{
            const params = row.id.split("/")
            this.checkEditor({ 
              ntype: params[0], ttype: params[1], id: parseInt(params[2],10), 
              shipping: true
            }, EDITOR_EVENT.LOAD_EDITOR)
          })
        } else {
          this.checkEditor({
            ntype: ntype || storeEdit.current.type, 
            ttype: ttype || storeEdit.current.transtype, 
            id: null}, EDITOR_EVENT.LOAD_EDITOR)
        }
        break;

      case SIDE_EVENT.COPY:
        this.transCopy(data.value)
        break;

      case SIDE_EVENT.LINK:
        this.setLink(data.type, data.field)
        break;

      case SIDE_EVENT.PASSWORD:
        this.setPassword()
        break;

      case SIDE_EVENT.SHIPPING_ADD_ALL:
        this.shippingAddAll()
        break;

      case SIDE_EVENT.SHIPPING_CREATE:
        this.createShipping()
        break;

      case SIDE_EVENT.REPORT_SETTINGS:
        this.reportSettings()
        break;

      case SIDE_EVENT.SEARCH_QUEUE:
        this.searchQueue()
        break;

      case SIDE_EVENT.EXPORT_QUEUE_ALL:
        this.exportQueueAll()
        break;

      case SIDE_EVENT.CREATE_REPORT:
        this.createReport(data.value)
        break;

      case SIDE_EVENT.EXPORT_EVENT:
        this.exportEvent()
        break;

      case SIDE_EVENT.SAVE_BOOKMARK:
        saveBookmark([...data.value])
        break;

      case SIDE_EVENT.HELP:
        showHelp(data.value)
        break;

      default:
        break;
    }
    return true
  }

  async saveEditor() {
    const { sql, initItem, validator } = this.app.modules
    const { resultError, createHistory, requestData, getSql } = this.app
    const { data } = this.store

    let edit = {...data[APP_MODULE.EDIT]}
    const newItem = (edit.current.item.id === null)

    let values = this.tableValues(edit.current.type, edit.current.item)
    values = await validator(edit.current.type, values)
    if(values.error){
      resultError(values)
      return null
    }

    let result = await requestData(`/${edit.current.type}`, { method: "POST", data: [values] })
    if(result.error){
      resultError(result)
      return null
    }
    if (edit.current.item.id === null) {
      edit = {...edit, current: {...edit.current, item: {...edit.current.item,
        id: result[0]
      }}}
    }
    edit = {...edit,
      dirty: false
    }
    await createHistory("save")

    if (typeof edit.current.extend !== "undefined") {
      if (edit.current.extend.ref_id === null) {
        edit = {...edit, current: {...edit.current, extend: {...edit.current.extend,
          ref_id: edit.current.item.id
        }}}
      }
      if (edit.current.extend.trans_id === null) {
        edit = {...edit, current: {...edit.current, extend: {...edit.current.extend,
          trans_id: edit.current.item.id
        }}}
      }
    }
    edit = {...edit, current: {...edit.current,
      fieldvalue: [...edit.current.fieldvalue]
    }}
    for (let i=0; i < edit.current.fieldvalue.length; i += 1) {
      if (edit.current.fieldvalue[i].ref_id === null) {
        edit.current.fieldvalue[i] = {...edit.current.fieldvalue[i],
          ref_id: edit.current.item.id
        }
      }
    }

    if (edit.current.type === "trans") {
      switch (edit.current.transtype) {
        case "invoice":
          if(!newItem){
            const params = { 
              method: "POST", 
              data: [{ 
                key: "fields",
                text: getSql(data[APP_MODULE.LOGIN].data.engine, sql.trans.invoice_customer()).sql,
                values: [edit.current.item.customer_id]
              }]
            }
            const view = await requestData("/view", params)
            if(view.error){
              resultError(view)
              return null
            }
            if (view.fields.length > 0) {
              Object.keys(view.fields[0]).forEach((fieldname) => {
                edit = {...edit, current: {...edit.current,
                  fieldvalue: this.setFieldvalue(edit.current.fieldvalue, 
                    fieldname, edit.current.item.id, null, view.fields[0][fieldname])
                }}
              })
            }
          }
          break;
        case "worksheet":
          const wlist = {trans_wsdistance:0, trans_wsrepair:0, trans_wstotal:0, trans_wsnote:""};
          Object.keys(wlist).forEach((fieldname) => {
            edit = {...edit, current: {...edit.current,
              fieldvalue: this.setFieldvalue(edit.current.fieldvalue, 
                fieldname, edit.current.item.id, null, wlist[fieldname])
            }}
          })
          break;
        case "rent":
          const rlist = {trans_reholiday:0, trans_rebadtool:0, trans_reother:0, trans_rentnote:""};
          Object.keys(rlist).forEach((fieldname) => {
            edit = {...edit, current: {...edit.current,
              fieldvalue: this.setFieldvalue(edit.current.fieldvalue, 
                fieldname, edit.current.item.id, null, rlist[fieldname])
            }}
          })
          break;
        case "delivery":
          let movements = [];
          edit.dataset.movement.forEach((mvt) => {
            let movement = {
              ...initItem({tablename: "movement", dataset: edit.dataset, current: edit.current}),
              ...mvt 
            }
            if(new Date(movement.shippingdate).toLocaleDateString() !== new Date(edit.current.item.transdate).toLocaleDateString()){
              movement = {...movement,
                shippingdate: `${new Date(edit.current.item.transdate).toISOString().split("T")[0]}T00:00:00`
              }
              movements = [...movements, movement]
            }
          })
          if (movements.length > 0) {
            result = await requestData("/movement", { method: "POST", data: movements })
            if(result.error){
              resultError(result)
              return null
            }
          }
          break;
        default:
          break;
      }
    }

    if (edit.current.fieldvalue.length > 0) {
      result = await requestData("/fieldvalue", { method: "POST", data: edit.current.fieldvalue })
      if(result.error){
        resultError(result)
        return null
      }
      edit = {...edit, current: {...edit.current,
        fieldvalue: [...edit.current.fieldvalue]
      }}
      for (let index = 0; index < result.length; index += 1) {
        if(!edit.current.fieldvalue[index].id){
          edit.current.fieldvalue[index] = {...edit.current.fieldvalue[index],
            id: result[index]
          }
        }
      }
    }

    if (typeof edit.current.extend !== "undefined") {
      let ptype = String(edit.template.options.extend).split("_")[0]
      let extend = {...edit.current.extend}
      switch (edit.current.transtype) {
        case "waybill":
          ptype = extend.seltype
          if (extend.seltype === "transitem") {
            ptype = "link"
            extend = initItem({tablename: "link", dataset: edit.dataset, current: edit.current})
            if (edit.dataset.translink.length > 0) {
              extend = {...extend, ...edit.dataset.translink[0]}
            } else {
              extend = {...extend,
                nervatype_1: data[APP_MODULE.LOGIN].data.groups.filter(
                  (item)=>((item.groupname === "nervatype") && (item.groupvalue === "trans")))[0].id,
                nervatype_2: data[APP_MODULE.LOGIN].data.groups.filter(
                  (item)=>((item.groupname === "nervatype") && (item.groupvalue === "trans")))[0].id,
                ref_id_1: edit.current.item.id
              }
            }
            extend = {...extend,
              ref_id_2: edit.current.extend.ref_id
            }
          } else {
            extend = null;
            if (edit.dataset.translink.length > 0) {
              result = await requestData("/link", 
                { method: "DELETE", query: { id: edit.dataset.translink[0].id } })
              if(result && result.error){
                resultError(result)
                return null
              }
            }
          }
          break;
        case "formula":
        case "production":
          const {product, ..._extend} = extend
          if(edit.current.transtype === "formula"){
            extend = {..._extend,
              shippingdate: `${new Date(edit.current.item.transdate).toISOString().split("T")[0]}T00:00:00`
            }
          } else {
            extend = {..._extend,
              shippingdate: edit.current.item.duedate,
              place_id: edit.current.item.place_id
            }
          }
          break;
        default:
      }
      if (extend !== null) {
        result = await requestData(`/${ptype}`, { method: "POST", data: [extend] })
        if(result.error){
          resultError(result)
          return null
        }
        if(extend.id === null){
          extend = {...extend,
            id: result[0]
          }
        }
        edit = {...edit, current: {...edit.current,
          extend
        }}
      }
    }
    
    return edit
  }

  async saveEditorForm() {
    const { initItem, validator } = this.app.modules
    const { resultError, createHistory, requestData } = this.app
    const { data } = this.store

    let edit = {...data[APP_MODULE.EDIT]}

    let values = this.tableValues(edit.current.form_datatype, edit.current.form)
    values = await validator(edit.current.form_datatype, values)
    if(values.error){
      resultError(values)
      return null
    }
    
    const fresult = await requestData(`/${edit.current.form_datatype}`, { method: "POST", data: [values] })
    if(fresult.error){
      resultError(fresult)
      return null
    }
    if (edit.current.form.id === null) {
      edit = {...edit, current: {...edit.current, form: {...edit.current.form,
        id: fresult[0]
      }}}
    }
    edit = {...edit,
      form_dirty: false
    }
    await createHistory("save")
   
    switch (edit.current.form_type) {
      case "movement":
        if (edit.current.transtype === "delivery") {
          let movements = []; let mlink = null;
          let movement = edit.dataset.movement_transfer.filter(
            item => (item.id === edit.current.form.id))[0]
          if (typeof movement !== "undefined") {
            movement = edit.dataset.movement.filter(
              item => (item.id === movement.ref_id))[0]
            movement = {
              ...initItem({tablename: "movement", dataset: edit.dataset, current: edit.current}), 
              ...movement
            }
          } else {
            movement = {
              ...initItem({tablename: "movement", dataset: edit.dataset, current: edit.current}),
              place_id: edit.current.item.place_id
            }
            mlink = {
              ...initItem({tablename: "link", dataset: edit.dataset, current: edit.current}),
              nervatype_1: data[APP_MODULE.LOGIN].data.groups.filter(
                (item)=>((item.groupname === "nervatype") && (item.groupvalue === "movement")))[0].id,
              nervatype_2: data[APP_MODULE.LOGIN].data.groups.filter(
                (item)=>((item.groupname === "nervatype") && (item.groupvalue === "movement")))[0].id,
              ref_id_2: edit.current.form.id
            }
          }
          movement = {...movement,
            product_id: edit.current.form.product_id,
            qty: -(edit.current.form.qty),
            notes: edit.current.form.notes
          }
          movements = [...movements, this.tableValues("movement", movement)]
          edit.dataset.movement_transfer.forEach((mvt) => {
            movement = {
              ...initItem({tablename: "movement", dataset: edit.dataset, current: edit.current}), 
              ...this.tableValues("movement", mvt)
            }
            if ((movement.id !== edit.current.form.id) &&
              (movement.place_id !== edit.current.form.place_id)){
                movement = {...movement,
                  place_id: edit.current.form.place_id
                }
                movements = [...movements, movement]
            }
          });
          const mresult = await requestData("/movement", { method: "POST", data: movements })
          if(mresult.error){
            resultError(mresult)
            return null
          }
          if (mlink !== null) {
            mlink = {...mlink,
              ref_id_1: mresult[0]
            }
            const lresult = await requestData("/link", { method: "POST", data: [mlink] })
            if(lresult.error){
              resultError(lresult)
              return null
            }
          }
        }
        break;
      
      case "price":
      case "discount":
        if (edit.current.price_link_customer !== null || 
            edit.current.price_customer_id !== null) {
          if (edit.current.price_customer_id === null) {
            // delete link
            const presult = await requestData("/link", 
              { method: "DELETE", query: { id: edit.current.price_link_customer } })
            if(presult && presult.error){
              resultError(presult)
              return null
            }
          } else {
            const clink = {
              ...initItem({tablename: "link", dataset: edit.dataset, current: edit.current}),
              id: edit.current.price_link_customer, // update or insert
              nervatype_1: data[APP_MODULE.LOGIN].data.groups.filter(
                (item)=>((item.groupname === "nervatype") && (item.groupvalue === "price")))[0].id,
              ref_id_1: edit.current.form.id,
              nervatype_2: data[APP_MODULE.LOGIN].data.groups.filter(
                (item)=>((item.groupname === "nervatype") && (item.groupvalue === "customer")))[0].id,
              ref_id_2: edit.current.price_customer_id
            }
            const plresult = await requestData("/link", { method: "POST", data: [clink] })
            if(plresult.error){
              resultError(plresult)
              return null
            }
          }
        }
        break;
        
      case "invoice_link":
      case "payment_link":
        const rsname = `${edit.current.form_type}_fieldvalue`
        for (let i=0; i < edit.current[rsname].length; i += 1) {
          if (edit.current[rsname][i].ref_id === null) {
            edit = {...edit, current: {...edit.current,
              [rsname]: [...edit.current[rsname]]
            }}
            edit.current[rsname][i] = {...edit.current[rsname][i],
              ref_id: edit.current.form.id
            }
          }
        }
        const flist = { link_qty:"0", link_rate: "1" }
        let fvalues = edit.current[rsname]
        Object.keys(flist).forEach((fieldname) => { 
          fvalues = this.setFieldvalue(fvalues, 
            fieldname, edit.current.form.id, flist[fieldname])
        });
        edit = {...edit, current: {...edit.current,
          [rsname]: fvalues
        }}
        const fpresult = await requestData("/fieldvalue", { method: "POST", data: fvalues })
        if(fpresult.error){
          resultError(fpresult)
          return null
        }
        break;

      default:
        break;
    }
    
    return edit
  }

  async searchQueue() {
    const { sql } = this.app.modules
    const { resultError, requestData, getSql } = this.app
    const { setData, data } = this.store
    let edit = {...data[APP_MODULE.EDIT]}
    
    const params = { 
      method: "POST", 
      data: [{ 
        key: "items",
        text: getSql(data[APP_MODULE.LOGIN].data.engine, sql.printqueue.items(edit.printqueue)).sql,
        values: []
      }]
    }
    const view = await requestData("/view", params)
    if(view.error){
      return resultError(view)
    }
    edit = {...edit, dataset: {...edit.dataset, 
      items: view.items 
    }}
    return setData(APP_MODULE.EDIT, edit)
  }

  setEditor(options, form, iedit) {
    const { initItem } = this.app.modules
    const { resultError, showToast, getAuditFilter, currentModule, msg } = this.app
    const { setData, data } = this.store

    let edit = {...(iedit||data[APP_MODULE.EDIT])}
    if ((typeof edit.dataset[edit.current.type] === "undefined") || 
      (edit.dataset[edit.current.type].length===0)) {
        showToast(TOAST_TYPE.ERROR, msg("", { id: "msg_editor_invalid" }), 0)
      return false;
    }
    
    edit = {
      ...edit,
      template: form,
      panel: form.options.panel,
      caption: form.options.title,
      audit: getAuditFilter(edit.current.type, edit.current.transtype)[0],
      side_view: "edit"
    }
    if (edit.audit==="disabled") {
      return false
    }

    if (edit.dataset[edit.current.type][0].id === null) {
      edit = {
        ...edit,
        current: {
          ...edit.current,
          item: edit.dataset[edit.current.type][0]
        }
      }
      if (form.options.search_form) {
        edit = {
          ...edit,
          title_field: form.options.title_field
        }
      } else {
        if((edit.audit==="all") || (edit.audit==="update")){
          edit = {
            ...edit,
            dirty: true
          }
        }
        edit = {
          ...edit,
          title_field: `${msg("",{ id: "label_new" })} ${form.options.title}`
        }
      }
      if (typeof form.options.extend !== "undefined") {
        edit = {
          ...edit,
          current: {
            ...edit.current,
            extend: initItem({tablename: form.options.extend, 
              dataset: edit.dataset, current: edit.current})
          }
        }
      }
    } else {
      edit = {
        ...edit,
        current: {
          ...edit.current,
          item: {
            ...initItem({
              tablename: edit.current.type, 
              dataset: edit.dataset, current: edit.current
            }),
            ...edit.dataset[edit.current.type][0]
          }
        }
      }
      if (typeof form.options.extend !== "undefined") {
        edit = {
          ...edit,
          current: {
            ...edit.current,
            extend: initItem({tablename: form.options.extend, 
              dataset: edit.dataset, current: edit.current})
          }
        }
        if (typeof edit.dataset[form.options.extend] !== "undefined") {
          if (edit.dataset[form.options.extend].length > 0) {
            edit = {
              ...edit,
              current: {
                ...edit.current,
                extend: {
                  ...edit.current.extend,
                  ...edit.dataset[form.options.extend][0]
                }
              }
            }
          }
        }
      }
    }

    edit = {
      ...edit,
      current: {
        ...edit.current,
        state: "normal",
        page: 0
      }
    }
    if (edit.current.type === "trans") {
      if (typeof edit.dataset.trans[0].transcast !== "undefined") {
        if (edit.dataset.trans[0].transcast === "cancellation") {
          edit = {
            ...edit,
            current: {
              ...edit.current,
              state: "cancellation"
            }
          }
        }
      }
      if (edit.dataset.pattern){
        const template = edit.dataset.pattern.filter((item) => (item.defpattern === 1))[0]
        edit = {
          ...edit,
          current: {
            ...edit.current,
            template: (template) ? template.id : ""
          }
        }
      }
    }
    if (edit.current.state === "normal" && edit.current.item.deleted === 1) {
      edit = {
        ...edit,
        current: {
          ...edit.current,
          state: "deleted"
        }
      }
    } else if (edit.current.item.closed === 1) {
      edit = {
        ...edit,
        current: {
          ...edit.current,
          state: "closed"
        }
      }
    }

    edit = {
      ...edit,
      current: {
        ...edit.current,
        fieldvalue: edit.dataset.fieldvalue || []
      }
    }

    Object.keys(edit.template.view).forEach(vname => {
      edit = { ...edit, template: { ...edit.template, view: { ...edit.template.view,
        [vname]: { ...edit.template.view[vname],
          view_audit: "all"
        }
      }}}
      if (vname === "setting") {
        edit = { ...edit, template: { ...edit.template, view: { ...edit.template.view,
          [vname]: { ...edit.template.view[vname],
            view_audit: "disabled"
          }
        }}}
      } else {
        edit = { ...edit, template: { ...edit.template, view: { ...edit.template.view,
          [vname]: { ...edit.template.view[vname],
            view_audit: getAuditFilter(form.view[vname].audit_type || vname, 
              form.view[vname].audit_transtype || null)[0]
          }
        }}}
      }
    });

    if (edit.current.type === "report") {
      let template = { fields: {} }
      try {
        template = JSON.parse(edit.dataset.report[0].report)
      } catch (err) {
        currentModule({ data: { module: APP_MODULE.SEARCH } })
        return resultError({ error: { message: err.message } })
      }
      Object.keys(template.fields).forEach((fieldname, index) => {
        const rfdata = template.fields[fieldname]
        const selected = (rfdata.selected)?
          rfdata.selected:
          (rfdata.wheretype === 'in')
        let tfrow = {
          id: index+1,
          rowtype: "reportfield", 
          datatype: rfdata.fieldtype,
          name: fieldname, 
          label: rfdata.description, 
          selected,
          empty: (rfdata.wheretype === 'in') ? 'false' : 'true',
          value: rfdata.value
        }
        switch (rfdata.fieldtype) {
          case "bool":
            tfrow = { ...tfrow,
              value: 0
            }
            break;
          case "valuelist":
            tfrow = { ...tfrow,
              description: (rfdata.wheretype !== "in") ? 
                rfdata.valuelist.split("|").unshift("") : rfdata.valuelist.split("|")
            }
            break;
          case "date":
            if(typeof(tfrow.value) === "undefined"){
              if (rfdata.defvalue) {
                const today = new Date()
                tfrow = { ...tfrow,
                  value: new Date(today.setDate((today.getDate() + parseInt(rfdata.defvalue,10)))).toISOString().split("T")[0]
                }
              } else if (rfdata.wheretype === "in") {
                tfrow = { ...tfrow,
                  value: new Date().toISOString().split("T")[0]
                }
              } else {
                tfrow = { ...tfrow,
                  value: ""
                }
              }
            }
            break;
          case "integer":
          case "float":
            if(typeof(tfrow.value) === "undefined"){
              tfrow = { ...tfrow,
                value: (rfdata.defvalue && rfdata.defvalue !== "") ? rfdata.defvalue : 0
              }
            }
            break;
          default:
            tfrow = { ...tfrow,
              datatype: "string"
            }
            break;
        }
        if (typeof tfrow.value === "undefined") {
          tfrow = { ...tfrow,
            value: (rfdata.defvalue) ? rfdata.defvalue : ""
          }
        }
        edit = {
          ...edit,
          current: {
            ...edit.current,
            fieldvalue: [...edit.current.fieldvalue, tfrow]
          }
        }
      });
    }

    if(options.shipping){
      edit = {
        ...edit,
        current: {
          ...edit.current,
          form_type: "transitem_shipping",
          direction: edit.dataset.groups.filter((group)=> (group.id === edit.current.item.direction))[0].groupvalue
        }
      }
      if (typeof edit.dataset.shiptemp === "undefined") {
        edit = {
          ...edit,
          dataset: {
            ...edit.dataset,
            shiptemp: []
          }
        }
      }
      edit = { ...edit, current: { ...edit.current, item: { ...edit.current.item,
        delivery_type: msg("", { id: `delivery_${edit.current.direction}` })
      }}}
      if(edit.current.shippingdate){
        edit = { ...edit, current: { ...edit.current, item: { ...edit.current.item,
          shippingdate: edit.current.shippingdate
        }}}
      } else {
        edit = { ...edit, current: { ...edit.current,
          shippingdate: `${new Date().toISOString().split("T")[0]}T00:00:00`,
          item: { ...edit.current.item,
            shippingdate: `${new Date().toISOString().split("T")[0]}T00:00:00`
          }
        }}
      }
      if(edit.current.shipping_place_id){
        edit = { ...edit, current: { ...edit.current, item: { ...edit.current.item,
          shipping_place_id: edit.current.shipping_place_id
        }}}
      } else{
        edit = { ...edit, current: { ...edit.current,
          shipping_place_id: null,
          item: { ...edit.current.item,
            shipping_place_id: null
          }
        }}
      }

      edit = { ...edit, dataset: { ...edit.dataset, 
        shipping_items_: []
      }}
      edit.dataset.shipping_items.forEach((item, index) => {
        let oitem = {...item,
          id: index+1,
          qty: parseFloat(item.qty)
        }
        const mitem = edit.dataset.transitem_shipping.filter((i)=> (i.id === `${oitem.item_id}-${oitem.product_id}`))[0]
        if (typeof mitem !== "undefined") {
          const tqty = (edit.current.direction === "out")?-parseFloat(mitem.sqty) : parseFloat(mitem.sqty)
          oitem = {...oitem,
            tqty,
            diff: parseFloat(oitem.qty) - tqty
          }
        } else {
          oitem = {...oitem,
            qty: parseFloat(oitem.qty),
            tqty: 0,
            diff: parseFloat(oitem.qty)
          }
        }
        const sitem = edit.dataset.shiptemp.filter((i)=> (i.id === `${oitem.item_id}-${oitem.product_id}`))[0]
        if (typeof sitem !== "undefined") {
          oitem = {...oitem,
            edited: true
          }
        }
        edit = { ...edit, dataset: { ...edit.dataset, 
          shipping_items_: [...edit.dataset.shipping_items_, oitem]
        }}
      });
    }

    if (edit.current.type === "printqueue") {
      if (typeof edit.printqueue === "undefined") {
        edit = {...edit,
          printqueue: edit.current.item
        }
      } else {
        edit = { ...edit, current: { ...edit.current, item: { ...edit.current.item,
          ...edit.printqueue
        }}}
      }
    }

    edit = {...edit, panel: {...edit.panel,
      form: true,
      state: edit.current.state
    }}
    if (edit.panel.state !== "normal") {
      edit = {...edit,
        audit: "readonly"
      }
    }
    if(edit.audit === "readonly") {
      edit = {...edit, panel: {...edit.panel,
        save: false, link: false, delete: false, new: false,
        pattern: false, password: false, formula: false
      }} 
      if (edit.panel.state !== "deleted") {
        edit = {...edit, panel: {...edit.panel,
          trans: false
        }}
      }
    }
    if (edit.audit !== "all") {
      edit = {...edit, panel: {...edit.panel,
        delete: false,
        new: false
      }}
      if (edit.panel.state !== "deleted") {
        edit = {...edit, panel: {...edit.panel,
          trans: false
        }}
      }
    }
    if (edit.panel.state === "deleted") {
      edit = {...edit, panel: {...edit.panel,
        copy: false, 
        create: false
      }}
    }
    edit = { ...edit, current: { ...edit.current,
      view: options.form||'form'
    }}
    setData(APP_MODULE.EDIT, edit)
    return true
  }

  setEditorItem(options) {
    const { forms, initItem } = this.app.modules
    const { getSetting } = this.app
    const { setData, data } = this.store
    let edit = {...data[APP_MODULE.EDIT]}
    let dkey = forms[options.fkey](undefined, edit, getSetting("ui")).options.data
    if (typeof dkey === "undefined") {
      dkey = options.fkey
    }
    edit = {...edit,
      form_dirty: false
    }
    edit = {...edit, current: {...edit.current,
      form_type: options.fkey,
      form_datatype: dkey,
      form: initItem({tablename: dkey, dataset: edit.dataset, current: edit.current})
    }}

    if (options.id!==null) {
      edit = {...edit, current: {...edit.current,
        form: edit.dataset[options.fkey].filter((item)=> (item.id === parseInt(options.id,10)))[0]
      }}
    } else {
      edit = {...edit, current: {...edit.current,
        form_dirty: true
      }}
    }
    edit = {...edit, current: {...edit.current,
      form_template: forms[options.fkey](edit.current.form, edit, getSetting("ui"))
    }}

    switch (options.fkey) {
      case "price":
      case "discount":
        edit = {...edit, current: {...edit.current,
          price_link_customer: null,
          price_customer_id: null
        }}
        if (options.id!==null) {
          edit = {...edit, current: {...edit.current,
            price_link_customer: edit.current.form.link_customer,
            price_customer_id: edit.current.form.customer_id
          }}
        }
        break;
      
      case "invoice_link":
        edit = {...edit, current: {...edit.current,
          invoice_link_fieldvalue: edit.dataset.invoice_link_fieldvalue
        }}
        let invoice_props = { 
          id: edit.current.form.id, ref_id_1: "", transnumber: "", curr: ""
        }
        const invoice_link = edit.dataset.invoice_link.filter((item)=> (item.id === edit.current.form.id))[0]
        if (typeof invoice_link !== "undefined") {
          invoice_props = invoice_link
        }
        edit = {...edit, current: {...edit.current,
          invoice_link: [invoice_props]
        }}
        break;
      
      case "payment_link":
        edit = {...edit, current: {...edit.current,
          payment_link_fieldvalue: edit.dataset.payment_link_fieldvalue
        }}
        let payment_props = { 
          id: edit.current.form.id, ref_id_2: "", transnumber: "", curr: ""
        }
        const payment_link = edit.dataset.payment_link.filter((item)=> (item.id === edit.current.form.id))[0]
        if (typeof payment_link !== "undefined") {
          payment_props = payment_link
        }
        edit = {...edit, current: {
          ...edit.current,
          payment_link: [payment_props]
        }}
        if(options.link_field){
          edit = {...edit, current: {...edit.current, form: {...edit.current.form,
            [options.link_field]: options.link_id
          }}}
        }
        break;
      
      default:
        break;}
    
    let panel = {...edit.current.form_template.options.panel,
      form: true,
      state: edit.current.state
    }
    if (edit.panel.state !== "normal") {
      edit = {...edit,
        audit: "readonly"
      }
    }
    if (edit.audit === "readonly") {
      panel = {...panel,
        save: false,
      }
    }
    if (edit.audit !== "all") {
      panel = {...panel,
        link: false,
        delete: false,
        new: false
      }
    }
    edit = {...edit,
      panel
    }
    setData(APP_MODULE.EDIT, edit)
  }

  setFieldvalue(_recordset, fieldname, ref_id, defvalue, value) {
    const { data } = this.store
    const { initItem } = this.app.modules
    let recordset = [..._recordset]
    const fieldvalue_idx = recordset.findIndex((item)=>((item.ref_id === ref_id)&&(item.fieldname === fieldname)))
    if (fieldvalue_idx === -1) {
      const fieldvalue = {...initItem({tablename: "fieldvalue", current: data[APP_MODULE.EDIT].current}),
        fieldname,
        ref_id,
        value: ((typeof value === "undefined") || (value === null)) ? defvalue : value
      }
      recordset = [...recordset, fieldvalue]
    } else if(value) {
      recordset[fieldvalue_idx] = {...recordset[fieldvalue_idx],
        value
      }
    }
    return recordset
  }

  setFormActions(options) {
    const { modalShipping } = this.module
    const { data, setData } = this.store

    const row = options.row || {}
    let edit = {...data[APP_MODULE.EDIT]}
    switch (options.params.action) {
      case ACTION_EVENT.LOAD_EDITOR:
        this.checkEditor({
          ntype: options.params.ntype || edit.current.type, 
          ttype: options.params.ttype || edit.current.transtype, 
          id: row.id || null }, 
          EDITOR_EVENT.LOAD_EDITOR)
        break;
      
      case ACTION_EVENT.NEW_EDITOR_ITEM:
        this.checkEditor({fkey: options.params.fkey, id: null}, EDITOR_EVENT.SET_EDITOR_ITEM)
        break;
      
      case ACTION_EVENT.EDIT_EDITOR_ITEM:
        this.setEditorItem({fkey: options.params.fkey, id: row.id})
        break;
      
      case ACTION_EVENT.DELETE_EDITOR_ITEM:
        this.deleteEditorItem({fkey: options.params.fkey, table: options.params.table, id: row.id})
        break;

      case ACTION_EVENT.LOAD_SHIPPING:
        this.checkEditor({
          ntype: options.params.ntype || edit.current.type, 
          ttype: options.params.ttype || edit.current.transtype, 
          id: options.params.id || edit.current.item.id, 
          shipping: true}, EDITOR_EVENT.LOAD_EDITOR)
        break;
      
      case ACTION_EVENT.ADD_SHIPPING_ROW:
        if (row.edited !== true) {
          edit.dataset.shiptemp = [...edit.dataset.shiptemp,{ 
            "id": `${row.item_id}-${row.product_id}`, 
            "item_id": row.item_id, "product_id": row.product_id,  
            "product": row.product, "partnumber": row.partnumber,
            "partname": row.partname, "unit": row.unit, 
            "batch_no":"", "qty":row.diff, "diff":0,
            "oqty":row.qty, "tqty":row.tqty
          }]
          this.setEditor({ shipping: true, form:"shipping_items" }, edit.template, edit)
        }
        break;
      
      case ACTION_EVENT.SHOW_SHIPPING_STOCK:
        this.showStock({ 
          product_id: row.product_id, 
          partnumber: row.partnumber, 
          partname: row.partname
        })
        break;

      case ACTION_EVENT.EDIT_SHIPPING_ROW:
        const modalForm = modalShipping({
          partnumber: row.partnumber,
          description: row.product,
          unit: row.unit,
          batch_no: row.batch_no,
          qty: row.qty,
          onEvent: {
            onModalEvent: (modalResult) => {
              setData("current", { modalForm: null })
              if(modalResult.key === MODAL_EVENT.OK){
                const { batch_no, qty } = modalResult.data
                const index = edit.dataset.shiptemp.findIndex(item => (item.id === row.id))
                edit = {...edit, dataset: {...edit.dataset,
                  shiptemp: [...edit.dataset.shiptemp]
                }}
                edit.dataset.shiptemp[index] = {...edit.dataset.shiptemp[index],
                  batch_no, qty,
                  diff: row.oqty - (row.tqty + qty)
                }
                setData(APP_MODULE.EDIT, edit)
                options.ref.requestUpdate()
              }
            }
          }
        })
        setData("current", { modalForm, side: SIDE_VISIBILITY.HIDE })
        break;
      
      case ACTION_EVENT.DELETE_SHIPPING_ROW:
        const index = edit.dataset.shiptemp.findIndex(item => (item.id === row.id))
        edit.dataset.shiptemp =  [...edit.dataset.shiptemp.slice(0,index),...edit.dataset.shiptemp.slice(index+1)]
        this.setEditor({shipping: true, form:"shiptemp_items"}, edit.template, edit)
        break;

      case ACTION_EVENT.EXPORT_QUEUE_ITEM:
        this.exportQueue(row)
        break;
    
      default:
        break;
    }
  }
  
  setLink(type, field) {
    const { setData, data } = this.store
    const { current } = data[APP_MODULE.EDIT]
    setData("current", { side: SIDE_VISIBILITY.HIDE })
    const link_id = (current.transtype === "cash") ? current.extend.id : current.form.id;
    this.checkEditor(
      { fkey: type, id: null, link_field: field, link_id }, EDITOR_EVENT.SET_EDITOR_ITEM)
  }

  setPassword(_username) {
    const { data } = this.store
    const { currentModule } = this.app
    const { current, dataset } = data[APP_MODULE.EDIT]
    let username = _username
    if(!username && current){
      username = dataset[current.type][0].username
    }
    currentModule({ 
      data: { module: APP_MODULE.SETTING, side: SIDE_VISIBILITY.HIDE }, 
      content: { fkey: "checkSetting", args: [ { username }, SIDE_EVENT.PASSWORD_FORM ] } 
    })
  }

  setPattern( options ) {
    const { inputBox } = this.host
    const { initItem } = this.app.modules
    const { requestData, resultError, msg, showToast } = this.app
    const { current, dataset } = this.store.data[APP_MODULE.EDIT]
    const { setData } = this.store
    const { key, text, ref } = options
    const updatePattern = async (values, skipLoading) => {
      const params = { method: "POST", data: values }
      const result = await requestData("/pattern", params)
      if(result.error){
        return resultError(result)
      }
      if(!skipLoading){
        return this.checkEditor({ntype: current.type, 
          ttype: current.transtype, 
          id: current.item.id, form:"fnote"}, EDITOR_EVENT.LOAD_EDITOR)
      }
      return true
    }
    const patternBox = {
      default: {
        title: msg("", { id: "msg_warning" }),
        message: msg("", { id: "msg_pattern_default" }),
        infoText: undefined,
        value: "",
        showValue: false,
        defaultOK: true,
        onEvent: {
          onModalEvent: async (modalResult) => {
            setData("current", { modalForm: null })
            if (modalResult.key === MODAL_EVENT.OK) {
              const pattern = [...dataset.pattern]
              pattern.forEach((element, index) => {
                pattern[index] = {...pattern[index],
                  defpattern: (element.id === parseInt(current.template,10)) ? 1 : 0
                }
              });
              updatePattern(pattern)
            }
          }
        }
      },
      save: {
        title: msg("", { id: "msg_warning" }),
        message: msg("", { id: "msg_pattern_save" }),
        infoText: undefined,
        value: "",
        showValue: false,
        defaultOK: true,
        onEvent: {
          onModalEvent: async (modalResult) => {
            setData("current", { modalForm: null })
            if (modalResult.key === MODAL_EVENT.OK) {
              let pattern = dataset.pattern.filter((item) => 
                (item.id === parseInt(current.template,10) ))[0]
              if(pattern){
                pattern = {...pattern,
                  notes: text
                }
                updatePattern([pattern], true)
              }
            }
          }
        }
      },
      new: {
        title: msg("", { id: "msg_pattern_new" }),
        message: msg("", { id: "msg_pattern_name" }),
        infoText: undefined,
        value:"", 
        showValue: true,
        defaultOK: false,
        onEvent: {
          onModalEvent: async (modalResult) => {
            const { value } = modalResult.data
            setData("current", { modalForm: null })
            if (modalResult.key === MODAL_EVENT.OK) {
              if(value !== ""){
                const result = await requestData("/pattern", {
                  query: {
                    filter: `description;==;${value}`
                  }
                })
                if(result.error){
                  return resultError(result)
                }
                if(result.length > 0){
                  return showToast(TOAST_TYPE.ERROR, msg("", { id: "msg_value_exists" }))
                }
                const pattern = {...initItem({tablename: "pattern", current}),
                  description: value,
                  defpattern: (dataset.pattern.length === 0) ? 1 : 0
                }
                return updatePattern([pattern])
              }
            }
            return true
          }
        }
      },
      delete: {
        title: msg("", { id: "msg_warning" }),
        message: msg("", { id: "msg_delete_text" }),
        infoText: msg("", { id: "msg_delete_info" }),
        value: "",
        showValue: false,
        defaultOK: false,
        onEvent: {
          onModalEvent: async (modalResult) => {
            setData("current", { modalForm: null })
            if (modalResult.key === MODAL_EVENT.OK) {
              let pattern = dataset.pattern.filter((item) => 
                (item.id === parseInt(current.template,10) ))[0]
              if(pattern){
                pattern = {...pattern,
                  deleted: 1,
                  defpattern: 0
                }
                updatePattern([pattern])
              }
            }
          }
        }
      }
    }
    if(key !== "new"){
      if(!current.template || (current.template === "")){
        return showToast(TOAST_TYPE.ERROR, msg("", { id: "msg_pattern_missing" }))
      }
    }
    if(key === "load"){
      const pattern = dataset.pattern.filter((item) => 
        (item.id === parseInt(current.template,10) ))[0]
      if(pattern){
        let edit = {...this.store.data[APP_MODULE.EDIT],
          current: {...this.store.data[APP_MODULE.EDIT].current, 
            item: { ...this.store.data[APP_MODULE.EDIT].current.item,
              fnote: pattern.notes
            }
          }
        }
        edit = {...edit,
          dirty: true,
        }
        ref._value = pattern.notes
        setData(APP_MODULE.EDIT, edit)
      }
      return true
    }
    const modalForm = inputBox({
      title: patternBox[key].title,
      message: patternBox[key].message,
      infoText: patternBox[key].infoText,
      value: patternBox[key].value,
      showValue: patternBox[key].showValue,
      defaultOK: patternBox[key].defaultOK,
      onEvent: patternBox[key].onEvent
    })
    return setData("current", { modalForm })
  }

  shippingAddAll() {
    const { data } = this.store
    let edit = {...data[APP_MODULE.EDIT]}
    edit.dataset.shipping_items_.forEach(sitem => {
      if (sitem.diff !== 0 && sitem.edited !== true) {
        edit = {...edit, dataset: {...edit.dataset, shiptemp: [...edit.dataset.shiptemp, {
          "id": `${sitem.item_id}-${sitem.product_id}`,
          "item_id": sitem.item_id, 
          "product_id": sitem.product_id,  
          "product": sitem.product, 
          "partnumber": sitem.partnumber,
          "partname": sitem.partname, 
          "unit": sitem.unit, 
          "batch_no":"", 
          "qty":sitem.diff, 
          "diff":0,
          "oqty": sitem.qty, 
          "tqty": sitem.tqty
        }]}}
      }
    });
    this.setEditor({shipping: true, form:"shiptemp_items"}, edit.template, edit)
  }

  async showStock (options) {
    const { modalStock } = this.module
    const { sql } = this.app.modules
    const { requestData, resultError, getSql, showToast, getSetting, msg } = this.app
    const { data, setData } = this.store
    const params = { 
      method: "POST", 
      data: [{ 
        key: "stock",
        text: getSql(data[APP_MODULE.LOGIN].data.engine, sql.trans.shipping_stock()).sql,
        values: [options.product_id] 
      }]
    }
    const view = await requestData("/view", params)
    if(view.error){
      return resultError(view)
    }
    if (view.stock.length === 0){
      return showToast(TOAST_TYPE.INFO, msg("", { id: "ms_no_stock" }))
    }
    const modalForm = modalStock({
      partnumber: options.partnumber,
      partname: options.partname,
      rows: view.stock,
      selectorPage: getSetting("selectorPage"),
      onEvent: {
        onModalEvent: () => {
          setData("current", { modalForm: null })
        }
      }
    })
    return setData("current", { modalForm, side: SIDE_VISIBILITY.HIDE })
  }

  tableValues(type, item) {
    const { initItem } = this.app.modules
    const { data } = this.store
    const values = {}
    const baseValues = initItem({tablename: type, 
      dataset: data[APP_MODULE.EDIT].dataset, current: data[APP_MODULE.EDIT].current})
    Object.keys(item).forEach(key => {
      if (typeof(baseValues[key]) !== "undefined") {
        values[key] = item[key]
      }
    });
    return values
  }

  transCopy(ctype) {
    const { inputBox } = this.host
    const { msg } = this.app
    const { setData } = this.store
    if (ctype === "create") {
      return this.checkEditor({}, EDITOR_EVENT.CREATE_TRANS_OPTIONS);
    }
    const  modalForm = inputBox({ 
      title: msg("", { id: "msg_warning" }),
      message: msg("", { id: "msg_copy_text" }),
      infoText: msg("", { id: "msg_delete_info" }),
      defaultOK: true,
      onEvent: {
        onModalEvent: async (modalResult) => {
          setData("current", { modalForm: null })
          if (modalResult.key === MODAL_EVENT.OK) {
            this.checkEditor({ cmdtype: "copy", transcast: ctype }, EDITOR_EVENT.CREATE_TRANS);
          }
        }
      }
    })
    return setData("current", { modalForm, side: SIDE_VISIBILITY.HIDE })
  }

}