import { APP_MODULE } from '../config/enums.js'

export const Validator = (app) => {
  const { data } = app.store
  const { getSql, requestData, msg } = app
  const getText = (key) => msg("", { id: key })
  return async (typeName, _values) => {
    const values = _values
    
    const checkUniqueKey = async (whereFilter, whereValue) => {
      const params = { 
        method: "POST", 
        data: [{ 
          key: "check",
          text: getSql(data[APP_MODULE.LOGIN].data.engine, 
            {select:["count(*) as recnum"], from:typeName, where:[whereFilter]}).sql,
          values: whereValue
        }]
      }
      const view = await requestData("/view", params)
      if(view.error){
        return view.error.message || getText("error_internal")
      }
      if(view.check[0].recnum > 0){
        return getText("msg_value_exists")
      }
      return ""
    }

    const nextNumber = async (numberkey, fieldname) => {
      const options = { method: "POST", 
        data: {
          key: "nextNumber",
          values: {
            numberkey, 
            step: true
          }
        }
      }
      const result = await requestData("/function", options)
      if(result.error){
        return result.error.message || getText("error_internal")
      }
      values[fieldname] = result
      return ""
    }

    let msg_err = ""

    if (typeof values.id==="undefined") {
      values.id = null
    }
    switch (typeName) {
      case "address":
        break;
        
      case "barcode":
        if (values.code === null || values.code === "") {
          msg_err = `${getText("msg_required")  } ${ getText("barcode_code")}`
        } else if (values.barcodetype === null || values.barcodetype === "") {
          msg_err = `${getText("msg_required")  } ${ getText("barcode_barcodetype")}`
        } else if(values.id === null) { 
          msg_err = await checkUniqueKey(["code","=","?"], [values.code])
        }
        break;
      
      case "contact":
        break;
        
      case "currency":
        if (values.curr === null || values.curr === "") {
          msg_err = `${getText("msg_required")  } ${ getText("currency_curr")}`
        } else if (values.description === null || values.description === "") {
          msg_err = `${getText("msg_required")  } ${ getText("currency_description")}`
        } else if(values.id === null) {
          msg_err = await checkUniqueKey(["curr","=","?"], [values.curr])
        }
        break;
      
      case "customer":
        if (values.custname === null || values.custname === "") {
          msg_err = `${getText("msg_required")  } ${ getText("customer_custname")}`
        } else if (values.custtype === null || values.custtype === "") {
          msg_err = `${getText("msg_required")  } ${ getText("customer_custtype")}`
        } else if(values.id === null && (values.custnumber !== null && values.custnumber !== "")) {
          msg_err = await checkUniqueKey(["custnumber","=","?"], [values.custnumber])
        } else if(values.custnumber === null) {
          msg_err = await nextNumber("custnumber", "custnumber")
        }
        break;
      
      case "deffield":
        if (values.nervatype === null || values.nervatype === "") {
          msg_err = `${getText("msg_required")  } ${ getText("deffield_nervatype")}`
        } else if (values.fieldtype === null || values.fieldtype === "") {
          msg_err = `${getText("msg_required")  } ${ getText("deffield_fieldtype")}`
        } else if (values.description === null || values.description === "") {
          msg_err = `${getText("msg_required")  } ${ getText("deffield_description")}`
        }
        break;
        
      case "employee":
        if (data[APP_MODULE.EDIT].current.extend.surname === null || 
          data[APP_MODULE.EDIT].current.extend.surname === "") {
          msg_err = `${getText("msg_required")  } ${ getText("contact_surname")}`
        } else if (values.usergroup === null || values.usergroup === "") {
          msg_err = `${getText("msg_required")  } ${ getText("employee_usergroup")}`
        } else if(values.id === null && (values.empnumber !== null && values.empnumber !== "")) {
          msg_err = await checkUniqueKey(["empnumber","=","?"], [values.empnumber])
        } else if(values.empnumber === null) {
          msg_err = await nextNumber("empnumber", "empnumber")
        }
        break;
        
      case "event":
        if(values.calnumber === null) {
          msg_err = await nextNumber("calnumber", "calnumber")
        }
        break;
        
      case "fieldvalue":
        break;
      case "formula":
        break;
        
      case "groups":
        if (values.groupvalue === null || values.groupvalue === "") {
          msg_err = `${getText("msg_required")  } ${ getText("groups_groupvalue")}`
        } else if (values.groupname === null || values.groupname === "") {
          msg_err = `${getText("msg_required")  } ${ getText("groups_groupname")}`
        } else if (values.groupname === "usergroup" && 
          (values.description === "" || values.description === null)) {
          msg_err = `${getText("msg_required")  } ${ getText("groups_description")}`
        } else if (values.groupname === "usergroup" && 
          (values.transfilter === "" || values.transfilter === null)) {
          msg_err = `${getText("msg_required")  } ${ getText("groups_transfilter")}`
        } else if(values.id === null) {
          msg_err = await checkUniqueKey(
            [["groupname","=","?"], ["and","groupvalue","=","?"]], [values.groupname, values.groupvalue])
        }
        break;
        
      case "item":
        if (values.product_id === null || values.product_id === "") {
          msg_err = `${getText("msg_required")  } ${ getText("product_partnumber")}`
        } else if (values.description === null || values.description === "") {
          msg_err = `${getText("msg_required")  } ${ getText("item_description")}`
        } else if (values.tax_id === null || values.tax_id === "") {
          msg_err = `${getText("msg_required")  } ${ getText("item_taxcode")}`
        }
        break;
        
      case "link":
        if (values.ref_id_1 === null || values.ref_id_1 === "" || 
          values.ref_id_2 === null || values.ref_id_2 === "") {
          msg_err = `${getText("msg_required")  } ${ getText("document_ref_transnumber")}`
        }
        break;
        
      case "log":
        break;
      
      case "movement":          
        switch (data[APP_MODULE.EDIT].current.transtype) {
          case "delivery":
            const direction = data[APP_MODULE.EDIT].dataset.groups.filter(
              item => (item.id === data[APP_MODULE.EDIT].current.item.direction))[0].groupvalue
            if (values.place_id === null || values.place_id === "") {
              if (direction === "transfer") {
                msg_err = `${getText("msg_required")  } ${ getText("movement_target")}`
              } else {
                msg_err = `${getText("msg_required")  } ${ getText("movement_place")}`
              }
            } else if ((direction === "transfer") && 
              (data[APP_MODULE.EDIT].current.item.place_id === values.place_id)) {
              msg_err = `${getText("msg_required")  } ${ getText("ms_diff_warehouse_err")}`
            } else if (values.product_id === null || values.product_id === "") {
              msg_err = `${getText("msg_required")  } ${ getText("product_description")}`
            }
            break;
          
          case "inventory":
            if (values.product_id === null || values.product_id === "") {
              msg_err = `${getText("msg_required")  } ${ getText("product_description")}`
            }
            break;
              
          case "production":
            if (values.product_id === null || values.product_id === "") {
              msg_err = `${getText("msg_required")  } ${ getText("product_description")}`
            } else if (values.place_id === null || values.place_id === "") {
              msg_err = `${getText("msg_required")  } ${ getText("movement_place")}`
            }
            break;
          
          case "formula":
            if (values.product_id === null || values.product_id === "") {
              msg_err = `${getText("msg_required")  } ${ getText("product_description")}`
            }
            break;
            
          case "waybill":
            if (values.tool_id === null || values.tool_id === "") {
              msg_err = `${getText("msg_required")  } ${ getText("tool_serial")}`
            }
            break;
          
          default:
            break;}            
        break;
        
      case "numberdef":
        break;
      case "pattern":
        break;
      case "payment":
        break;
        
      case "place":
        if ((values.id === null) && 
          (parseInt(values.placetype,10) !== data[APP_MODULE.EDIT].dataset.placetype.filter(
          item => (item.groupvalue === "warehouse"))[0].id)) {
          values.curr = data[APP_MODULE.EDIT].dataset.currency[0].curr;}
        if (values.description === null || values.description === "") {
          msg_err = `${getText("msg_required")  } ${ getText("place_description")}`
        } else if (values.placetype === null || values.placetype === "") {
          msg_err = `${getText("msg_required")  } ${ getText("place_placetype")}`
        } else if(values.planumber === null) {
          msg_err = await nextNumber("planumber", "planumber")
        }
        break;
      
      case "price":
        if (values.validfrom === null || values.validfrom === "") {
          msg_err = `${getText("msg_required")  } ${ getText("price_validfrom")}`
        } else if (values.curr === null || values.curr === "") {
          msg_err = `${getText("msg_required")  } ${ getText("price_curr")}`
        } else if (((values.calcmode === null) || (values.calcmode === "")) && (values.discount !== null)) {
          msg_err = `${getText("msg_required")  } ${ getText("price_calcmode")}`
        }
        break;
        
      case "product":
        if (values.description === null || values.description === "") {
          msg_err = `${getText("msg_required")  } ${ getText("product_description")}`
        } else if (values.protype === null || values.protype === "") {
          msg_err = `${getText("msg_required")  } ${ getText("product_protype")}`
        } else if (values.unit === null || values.unit === "") {
          msg_err = `${getText("msg_required")  } ${ getText("product_unit")}`
        } else if (values.tax === null || values.tax === "") {
          msg_err = `${getText("msg_required")  } ${ getText("product_tax")}`
        } else if(values.id === null && (values.partnumber !== null && values.partnumber !== "")) {
          msg_err = await checkUniqueKey(["partnumber","=","?"], [values.partnumber])
        } else if(values.partnumber === null) {
          msg_err = await nextNumber("partnumber", "partnumber")
        }
        break;
      
      case "project":
        if (values.description === null || values.description === "") {
          msg_err = `${getText("msg_required")  } ${ getText("project_description")}`
        } else if(values.id === null && (values.pronumber !== null && values.pronumber !== "")) {
          msg_err = await checkUniqueKey(["pronumber","=","?"], [values.pronumber])
        } else if(values.pronumber === null) {
          msg_err = await nextNumber("pronumber", "pronumber")
        }
        break;
        
      case "rate":
        if (values.ratetype === null || values.ratetype === "") {
          msg_err = `${getText("msg_required")  } ${ getText("rate_ratetype")}`
        } else if (values.ratedate === null || values.ratedate === "") {
          msg_err = `${getText("msg_required")  } ${ getText("rate_ratedate")}`
        } else if (values.curr === null || values.curr === "") {
          msg_err = `${getText("rate_curr")  } ${ getText("rate_ratedate")}`
        }
        break;
        
      case "tax":
        if (values.taxcode === null || values.taxcode === "") {
          msg_err = `${getText("msg_required")  } ${ getText("tax_taxcode")}`
        } else if (values.description === null || values.description === "") {
          msg_err = `${getText("msg_required")  } ${ getText("tax_description")}`
        } else if(values.id === null) {
          msg_err = await checkUniqueKey(["taxcode","=","?"], [values.taxcode])
        }
        break;
      
      case "tool":
        if (values.product_id === null || values.product_id === "") {
          msg_err = `${getText("msg_required")  } ${ getText("product_partnumber")}`
        } else if (values.description === null || values.description === "") {
          msg_err = `${getText("msg_required")  } ${ getText("tool_description")}`
        } else if(values.id === null && (values.serial !== null && values.serial !== "")) {
          msg_err = await checkUniqueKey(["serial","=","?"], [values.serial])
        } else if(values.serial === null) {
          msg_err = await nextNumber("serial", "serial")
        }
        break;
        
      case "trans":
        const transtype = data[APP_MODULE.EDIT].dataset.groups.filter(
          item => (item.id === parseInt(values.transtype,10)))[0].groupvalue
        const direction = data[APP_MODULE.EDIT].dataset.groups.filter(
          item => (item.id === parseInt(values.direction,10)))[0].groupvalue
        if (["offer", "order", "worksheet", "rent", "invoice"].includes(transtype) && 
          (values.customer_id === null || values.customer_id === "")) {
          msg_err = `${getText("msg_required")  } ${ getText("customer_custname")}`
        } else if (transtype==="cash" && (values.place_id === null || values.place_id === "")) {
          msg_err = `${getText("msg_required")  } ${ getText("payment_place_cash")}`
        } else if (transtype==="bank" && (values.place_id === null || values.place_id === "")) {
          msg_err = `${getText("msg_required")  } ${ getText("payment_place_bank")}`
        } else if ((transtype==="inventory" || transtype==="production" || 
          (transtype==="delivery" && direction === "transfer")) && 
          (values.place_id === null || values.place_id === "")) {
          msg_err = `${getText("msg_required")  } ${ getText("movement_place")}`
        } else if (transtype==="production" && (values.duedate === null || values.duedate === "")) {
          msg_err = `${getText("msg_required")  } ${ getText("production_duedate")}`
        } else if ((transtype==="production" || transtype==="formula") && 
          (data[APP_MODULE.EDIT].current.extend.product_id === null || 
            data[APP_MODULE.EDIT].current.extend.product_id === "")) {
          msg_err = `${getText("msg_required")  } ${ getText("product_partnumber")}`
        } else if (transtype==="waybill" && (data[APP_MODULE.EDIT].current.extend.ref_id === null || 
            data[APP_MODULE.EDIT].current.extend.ref_id === "")) {
          msg_err = `${getText("msg_required")  } ${ getText("waybill_reference")}`
        } else if(values.transnumber === null) {
          if (transtype === "waybill" || transtype === "cash") {
            msg_err = await nextNumber(transtype, "transnumber")
          } else {
            msg_err = await nextNumber(`${transtype}_${direction}`, "transnumber")
          }
        }
        break;
      
      case "ui_menu":
        if (values.menukey === null || values.menukey === "") {
          msg_err = `${getText("msg_required")  } ${ getText("menucmd_menukey")}`
        } else if (values.description === null || values.description === "") {
          msg_err = `${getText("msg_required")  } ${ getText("menucmd_description")}`
        } else if (values.method === null || values.method === "") {
          msg_err = `${getText("msg_required")  } ${ getText("menucmd_method")}`
        } else if(values.id === null) {
          msg_err = await checkUniqueKey(["menukey","=","?"], [values.menukey])
        }
        break;
      
      default:
        break;
    }
    if (msg_err !== "") {
      return { error: { message: msg_err }}
    }
    return values
  }
}

export const InitItem = (app) => {
  const { data } = app.store
  const { getSetting } = app
  return (params) => {
    const dataset = params.dataset || data[APP_MODULE.EDIT].dataset
    const current = params.current || data[APP_MODULE.EDIT].current
    const store = data[APP_MODULE.LOGIN].data
    const config = getSetting("ui")
    switch (params.tablename) {
      case "address":
        return {
          id: null, 
          nervatype: store.groups.filter((group)=> ((group.groupname === "nervatype") && (group.groupvalue === current.type)))[0].id, 
          ref_id: current.item.id, 
          country: null, state: null, zipcode: null, city: null, street: null, notes: null, deleted: 0
        }
          
      case "audit":
        // ui_audit
        return {
          id: null, usergroup: null, nervatype: null, subtype: null, inputfilter: null, supervisor: 1
        }
        
      case "barcode":
        return {
          id: null, code: null, product_id: current.item.id, description: null,
          barcodetype: dataset.barcodetype.filter((group)=> ((group.groupname === "barcodetype") && (group.groupvalue === "CODE_39")))[0].id, 
          qty: 0, defcode: 0
        }
      
      case "contact":
        return {
          id: null,
          nervatype: store.groups.filter((group)=> ((group.groupname === "nervatype") && (group.groupvalue === current.type)))[0].id, 
          ref_id: current.item.id, 
          firstname: null, surname: null, status: null, 
          phone: null, fax: null, mobil: null, email: null, notes: null, deleted: 0
        }
          
      case "currency":
        return {
          id: null, curr: null, description: null, digit: 0, defrate: 0, cround: 0
        }
          
      case "customer":
        if (typeof dataset.custtype !== "undefined") {
          return {
            id: null,
            custtype: dataset.custtype.filter((group)=> (group.groupvalue === "company"))[0].id,  
            custnumber: null, custname: null, taxnumber: null, account: null,
            notax: 0, terms: 0, creditlimit: 0, discount: 0, notes: null, inactive: 0, deleted: 0
          }
        }  
        return null;
          
      case "deffield":
        return {
          id: null, 
          fieldname: `${Math.random().toString(16).slice(2)}-${Math.random().toString(16).slice(2)}`, 
          nervatype: null, subtype: null, fieldtype: null, description: null,
          valuelist:null, addnew: 0, visible: 1, readonly: 0, deleted: 0
        }
        
      case "employee":
        if(dataset.usergroup){
          return {
            id: null,
            empnumber: null, username: null,
            usergroup: dataset.usergroup.filter((group)=> (group.groupvalue === "admin"))[0].id, 
            startdate: new Date().toISOString().split("T")[0], 
            enddate: null, department: null,
            password: null, registration_key: null, inactive: 0, deleted: 0
          }
        }
        return null
        
      case "event":
        let event = {
          id: null, calnumber: null, 
          nervatype: null, ref_id: null, 
          uid: null, eventgroup: null, fromdate: null, todate: null, subject: null, 
          place: null, description: null, deleted: 0
        }
        if (typeof current.item !== "undefined") {
          if (current.type === "event") {
            event = {...event,
              nervatype: current.item.nervatype,
              ref_id: current.item.ref_id
            } 
          } else {
            event = {...event,
              nervatype: store.groups.filter((group)=> ((group.groupname === "nervatype") && (group.groupvalue === current.type)))[0].id,
              ref_id: current.item.id
            }
          }
        }
        return event;
      
      case "fieldvalue":
        let fieldvalue = {
          id: null, fieldname: null, ref_id: null, value: null, notes: null, deleted: 0
        }
        if (typeof current.item !== "undefined") {
          fieldvalue = {...fieldvalue,
            ref_id: current.item.id
          }
        }
        return fieldvalue;
      
      case "groups":
        return {
          id: null, groupname: null, groupvalue: null, description: null, 
          inactive: 0, deleted: 0
        }
      
      case "usergroup":
        // groups
        return {
          id: null, groupname: "usergroup", groupvalue: null, description: null, 
          transfilter: null, inactive: 0, deleted: 0
        }
        
      case "item":
        return {
          id: null, 
          trans_id: current.item.id, 
          product_id: null, unit: null, qty: 0, 
          fxprice: 0, netamount: 0, discount: 0, tax_id: null, 
          vatamount: 0, amount: 0, description: null, deposit: 0, 
          ownstock: 0, actionprice: 0, deleted: 0
        }
        
      case "link":
        let link = { 
          id: null, nervatype_1: null, ref_id_1: null, nervatype_2: null, 
          ref_id_2: null, deleted: 0
        }
        if(current.form_type === "invoice_link"){
          link = {...link,
            nervatype_1: store.groups.filter((group)=> ((group.groupname === "nervatype") && (group.groupvalue === "payment")))[0].id,
            nervatype_2: store.groups.filter((group)=> ((group.groupname === "nervatype") && (group.groupvalue === "trans")))[0].id,
            ref_id_2: current.item.id
          }
        }
        if(current.form_type === "payment_link"){
          link = {...link,
            nervatype_1: store.groups.filter((group)=> ((group.groupname === "nervatype") && (group.groupvalue === "payment")))[0].id,
            nervatype_2: store.groups.filter((group)=> ((group.groupname === "nervatype") && (group.groupvalue === "trans")))[0].id
          }
        }
        return link;
      
      case "log":
        return {
          id: null,
          fromdate: new Date().toISOString().split("T")[0], 
          todate: "", empnumber: "", logstate: "update", nervatype: ""
        }
      
      case "ui_menu":
        return {
          id: null, menukey: null, description: null, modul: null, icon: null, 
          funcname: null, method: dataset.method.filter((group)=> (group.groupvalue === "post"))[0].id, address: null
        }
      
      case "ui_menufields":
        return {
          id: null, menu_id: null, fieldname: "", description: "", 
          fieldtype: null, orderby: 0
        }
          
      case "movement":
        let movement = {
          id: null, trans_id: current.item.id, 
          shippingdate: null, movetype: null, product_id: null,
          tool_id: null, qty: 0, place_id: null, shared: 0, notes: null, deleted: 0
        }
        switch (current.transtype) {
          case "delivery":
            movement = {...movement,
              movetype: dataset.groups.filter((group)=> ((group.groupname === "movetype") && (group.groupvalue === "inventory")))[0].id,
              shippingdate: `${current.item.transdate} 00:00:00`
            }
            if (dataset.movement_transfer.length > 0){
              movement = {...movement,
                place_id: dataset.movement_transfer[0].place_id
              }
            }
            break;
          case "inventory":
            movement = {...movement,
              movetype: dataset.groups.filter((group)=> ((group.groupname === "movetype") && (group.groupvalue === "inventory")))[0].id,
              shippingdate: `${current.item.transdate} 00:00:00`,
              place_id: current.item.place_id
            }
            break;
          case "production":
            movement = {...movement,
              movetype: dataset.groups.filter((group)=> ((group.groupname === "movetype") && (group.groupvalue === "inventory")))[0].id,
              shippingdate: current.item.duedate
            }
            break;
          case "formula":
            movement = {...movement,
              movetype: dataset.groups.filter((group)=> ((group.groupname === "movetype") && (group.groupvalue === "plan")))[0].id,
              shippingdate: `${current.item.transdate} 00:00:00`
            }
            break;
          case "waybill":
            movement = {...movement,
              movetype: dataset.groups.filter((group)=> ((group.groupname === "movetype") && (group.groupvalue === "tool")))[0].id,
              shippingdate: `${current.item.transdate} 00:00:00`
            }
            break;
          default:
            movement = {...movement,
              movetype: dataset.groups.filter((group)=> ((group.groupname === "movetype") && (group.groupvalue === "inventory")))[0].id
            }
        }
        return movement;
      
      case "movement_head":
        // movement
        let movement_head = {
          id: null, trans_id: current.item.id, 
          shippingdate: null, product_id: null, product: "", movetype: null, 
          tool_id: null, qty: 0, place_id: null, shared: 0, notes: null, deleted: 0
        }
        if(current.transtype === "formula"){
          movement_head = {...movement_head,
            movetype: dataset.groups.filter((group)=> ((group.groupname === "movetype") && (group.groupvalue === "head")))[0].id
          }
        }
        if(current.transtype === "production"){
          movement_head = {...movement_head,
            movetype: dataset.groups.filter((group)=> ((group.groupname === "movetype") && (group.groupvalue === "inventory")))[0].id,
            shared: 1
          }
        }
        return movement_head;
          
      case "numberdef":
        return {
          id: null,
          numberkey: null, prefix: null, curvalue: 0, isyear: 1, sep: "/",
          len: 5, description: null, visible: 0, readonly: 0, orderby: 0
        }
        
      case "pattern":
        return {
          id: null,
          transtype: current.item.transtype,
          description: null, notes: "", defpattern: 0, deleted: 0
        }
          
      case "payment":
        return {
          id: null,
          trans_id: current.item.id, 
          paiddate: current.item.transdate, amount: 0, notes: null, deleted: 0
        }
      
      case "place":
        return {
          id: null,
          planumber: null, placetype:null, description: null,
          curr: null, defplace: 0, notes: null, inactive: 0, deleted: 0
        }
        
      case "price":
      case "discount":
        let price =  {
          id: null, product_id: current.item.id,
          validfrom: new Date().toISOString().split("T")[0], 
          validto: null, curr: null, qty: 0,
          pricevalue: 0, discount: null,
          calcmode: dataset.calcmode.filter((group)=> ((group.groupname === "calcmode") && (group.groupvalue === "amo")))[0].id, 
          vendorprice: 0, deleted: 0
        }
        if (params.tablename === "discount") {
          price = {...price,
            discount: 0
          }
        }
        const default_currency = dataset.settings.filter((group)=> (group.fieldname === "default_currency"))[0]
        if (typeof default_currency !== "undefined") {
          price = {...price,
            curr: default_currency.value
          }
        }
        return price;
        
      case "product":
        if(dataset.protype){
          let product = {
            id: null,
            protype: dataset.protype.filter((group)=> (group.groupvalue === "item"))[0].id,
            partnumber: null, description: null, unit: null,
            tax_id: null, notes: null, inactive: 0, webitem: 0, deleted: 0
          }
          const default_unit = dataset.settings.filter((group)=> (group.fieldname === "default_unit"))[0]
          if (typeof default_unit !== "undefined") {
            product = {...product,
              unit: default_unit.value
            }
          }
          const default_taxcode = dataset.settings.filter((group)=> (group.fieldname === "default_taxcode"))[0]
          if (typeof default_taxcode !== "undefined") {
            product = {...product,
              tax_id: dataset.tax.filter((tax)=> (tax.taxcode === default_taxcode.value))[0].id
            }
          } else {
            product = {...product,
              tax_id: dataset.tax.filter((tax)=> (tax.taxcode === "0%"))[0].id
            }
          }
          return product;
        }
        return null
      
      case "project":
        return {
          id: null,
          pronumber: null, description: null, customer_id: null, startdate: null, 
          enddate:null, notes:null, inactive:0, deleted: 0
        }
      
      case "printqueue":
        if ((current.type === "printqueue") && current.item) {
          return {
            id: null, 
            nervatype: current.item.nervatype, 
            startdate: current.item.startdate, 
            enddate: current.item.enddate,
            transnumber: current.item.transnumber, 
            username: current.item.username, 
            server: current.item.server, 
            mode: current.item.mode,
            orientation: current.item.orientation,
            size: current.item.size
          }
        }
        return {
          id: null, nervatype: null, startdate: null, enddate: null,
          transnumber: null, username: null, server: null, mode: "pdf", 
          orientation: config.page_orient, 
          size: config.page_size
        }
      
      case "rate":
        return {
          id: null,
          ratetype: null, ratedate: new Date().toISOString().split("T")[0], 
          curr: null, place_id: null, rategroup: null, ratevalue: 0, deleted: 0
        }
      
      case "refvalue":
        let refvalue = {
          seltype: "transitem", ref_id: null, refnumber: "", transtype: ""
        }
        if (current.transtype === "waybill") {
          const base_trans = dataset.trans[0]
          if (base_trans.customer_id !== null) {
            refvalue = {...refvalue,
              seltype: "customer",
              ref_id: base_trans.customer_id,
              refnumber: base_trans.custname
            }
          } else if (base_trans.employee_id !== null) {
            refvalue = {...refvalue,
              seltype: "employee",
              ref_id: base_trans.employee_id,
              refnumber: base_trans.empnumber
            }
          } else {
            refvalue = {...refvalue,
              seltype: "transitem",
            }
            if (dataset.translink && (dataset.translink.length > 0)) {
              refvalue = {...refvalue,
                ref_id: dataset.translink[0].ref_id_2,
                transtype: dataset.translink[0].transtype,
                refnumber: dataset.translink[0].transnumber
              }
            }
          }
        }
        return refvalue;
      
      case "report":
        // ui_report
        return {
          id: null,
          reportkey: null, nervatype: null, transtype: null, direction: null, repname: null,
          description: null, label: null, filetype: null, report: null,
          orientation: config.page_orient, size: config.page_size
        }
        
      case "tax":
        return {
          id:null,
          taxcode: null, description: null, rate: 0, inactive: 0
        }
      
      case "tool":
        return {
          id: null,
          serial: null, description: null, product_id: null, 
          toolgroup: null, notes: null, inactive: 0, deleted: 0
        }
        
      case "trans":
        const transtype = params.transtype || current.transtype;
        if (typeof dataset.pattern !== "undefined") {
          let trans = {
            id: null,
            transtype: dataset.groups.filter((group)=> ((group.groupname === "transtype") && (group.groupvalue === transtype)))[0].id,
            direction: dataset.groups.filter((group)=> ((group.groupname === "direction") && (group.groupvalue === "out")))[0].id, 
            transnumber: null, ref_transnumber: null, 
            crdate: new Date().toISOString().split("T")[0], 
            transdate: new Date().toISOString().split("T")[0], 
            duedate: null,
            customer_id: null, employee_id: null, department: null, project_id: null,
            place_id: null, paidtype: null, curr: null, notax: 0, paid: 0, acrate: 0, 
            notes: null, intnotes: null, fnote: null,
            transtate: dataset.transtate.filter((group)=> ((group.groupname === "transtate") && (group.groupvalue === "ok")))[0].id,
            cruser_id: store.employee.id, closed: 0, deleted: 0
          }
          const pattern = dataset.pattern.filter((p)=> (p.defpattern === 1))[0]
          if (typeof pattern !== "undefined") {
            trans = {...trans,
              fnote: pattern.notes
            }
          }
          switch (transtype) {
            case "offer":
            case "order":
            case "worksheet":
            case "rent":
            case "invoice":
            case "receipt":
              trans = {...trans,
                duedate: `${new Date().toISOString().split("T")[0]}T00:00:00`
              }
              const def_currency = dataset.settings.filter((group)=> (group.fieldname === "default_currency"))[0]
              if (typeof def_currency !== "undefined") {
                trans = {...trans,
                  curr: def_currency.value
                }
              }
              const default_paidtype = dataset.settings.filter((group)=> (group.fieldname === "default_paidtype"))[0]
              if (typeof default_paidtype !== "undefined") {
                trans = {...trans,
                  paidtype: dataset.paidtype.filter((group)=> (group.groupvalue === default_paidtype.value))[0].id
                }
              }
              break;
            case "bank":
            case "inventory":
            case "formula":
              trans = {...trans,
                direction: dataset.groups.filter((group)=> ((group.groupname === "direction") && (group.groupvalue === "transfer")))[0].id
              }
              break;
            case "production":
              trans = {...trans,
                direction: dataset.groups.filter((group)=> ((group.groupname === "direction") && (group.groupvalue === "transfer")))[0].id,
                duedate: `${new Date().toISOString().split("T")[0]}T00:00:00`
              }
              break;
            default:
              }
          if (transtype === "invoice") {
            const default_deadline = dataset.settings.filter((group)=> (group.fieldname === "default_deadline"))[0]
            if (typeof default_deadline !== "undefined") {
              const today = new Date();
              trans = {...trans,
                duedate: `${new Date(today.setDate((today.getDate() + parseInt(default_deadline.value,10)))).toISOString().split("T")[0]}T00:00:00`
              }
            }
          }    
          return trans;}
        return null;
        
      default:
    }
    return false;
  }
}