import update from 'immutability-helper';

import formatISO from 'date-fns/formatISO'
import addDays from 'date-fns/addDays'
import parseISO from 'date-fns/parseISO'
import isEqual from 'date-fns/isEqual'
import format from 'date-fns/format'

import printJS from 'print-js'

import { appActions, getSql, saveToDisk, guid } from 'containers/App/actions'
import dataset from 'containers/Controller/Dataset'
import { Sql } from 'containers/Controller/Sql'
import { Forms } from 'containers/Controller/Forms'
import { InitItem, Validator } from 'containers/Controller/Validator'
import { getSetting } from 'config/app'

import InputBox from 'components/Modal/InputBox'
import Formula from 'components/Modal/Formula'
import Shipping from 'components/Modal/Shipping'
import Stock from 'components/Modal/Stock'
import Trans from 'components/Modal/Trans'
import Report from 'components/Modal/Report'

export const editorActions = (data, setData) => {
  const app = appActions(data, setData)
  const initItem = InitItem(data, setData)
  const validator = Validator(data, setData)

  const forms = Forms({ getText: app.getText })
  const sql = Sql({ getText: app.getText })
  
  const round = (n,dec) => {
    n = parseFloat(n);
    if(!isNaN(n)){
      if(!dec) dec= 0;
      let factor= Math.pow(10,dec);
      return Math.floor(n*factor+((n*factor*10)%10>=5?1:0))/factor;
    } else {
      return n
    }
  }

  const reportPath = (params) => {
    let query = new URLSearchParams()
    query.append("reportkey", params.template)
    query.append("orientation", params.orient)
    query.append("size", params.size)
    query.append("output", params.type)
    if(params.filters){
      return `/report?${query.toString()}&${params.filters}`  
    }
    query.append("nervatype", params.nervatype||data.edit.current.type)
    return `/report?${query.toString()}&filters[@id]=${params.id||data.edit.current.item.id}`
  }

  const addPrintQueue = async (reportkey, copy) => {
    const report = data.edit.dataset.report.filter((item)=>(item.reportkey === reportkey))[0]
    const ntype = data.login.data.groups.filter(
      (item)=>((item.groupname === "nervatype") && (item.groupvalue === data.edit.current.type)))[0]
    const values = {
      "nervatype": ntype.id, 
      "ref_id": data.edit.current.item.id, 
      "qty": parseInt(copy), 
      "employee_id": data.login.data.employee.id, 
      "report_id": report.id
    }
    const options = { method: "POST", data: [values] }
    const result = await app.requestData("/ui_printqueue", options)
    if(result.error){
      return app.resultError(result)
    }
    app.showToast({ type: "success", autoClose: true,
      title: app.getText("msg_successful"), 
      message: app.getText("report_add_groups") })
  }

  const reportOutput = async (params) => {
    if(params.type === "printqueue"){
      return addPrintQueue(params.template, params.copy)
    }
    const result = await app.requestData(reportPath(params), {})
    if(result && result.error){
      app.resultError(result)
      return false
    }
    const resultUrl = URL.createObjectURL(result, {type : (params.type === "pdf") ? "application/pdf" : "application/xml; charset=UTF-8"})
    if(params.type === "print"){
      printJS({
        printable: resultUrl,
        type: 'pdf',
        base64: false,
      })
    } else {
      let filename = params.title+"_"+formatISO(new Date(), { representation: 'date' })+"."+params.type
      filename = filename.split("/").join("_")
      saveToDisk(resultUrl, filename)
    }
    return true
  }

  const searchQueue = async () => {
    let edit = update(data.edit, {})
    const params = { 
      method: "POST", 
      data: [{ 
        key: "items",
        text: getSql(data.login.data.engine, sql.printqueue.items(edit.printqueue)).sql,
        values: []
      }]
    }
    let view = await app.requestData("/view", params)
    if(view.error){
      return app.resultError(view)
    }
    edit = update(edit, {dataset: {$merge: {
      items: view.items
    }}})
    setData("edit", edit)
  }

  const exportQueueAll = () => {
    const options = data.edit.current.item
    if (data.edit.dataset.items.length > 0){
      if (options.mode === "print") {
        return app.showToast({ type: "error",
          title: app.getText("msg_warning"), 
          message: app.getText("ms_export_invalid")+" "+app.getText("printqueue_mode_print") })
      }
      setData("current", { modalForm: 
        <InputBox 
          title={app.getText("msg_warning")}
          message={app.getText("label_export_all_selected")}
          infoText={app.getText("msg_delete_info")+" "+app.getText("ms_continue_warning")}
          defaultOK={true}
          labelOK={app.getText("msg_ok")}
          labelCancel={app.getText("msg_cancel")}
          onCancel={() => {
            setData("current", { modalForm: null })
          }}
          onOK={(value) => {
            setData("current", { modalForm: null }, async () => {
              for (let index = 0; index < data.edit.dataset.items.length; index++) {
                const item = data.edit.dataset.items[index];
                let result = await reportOutput({
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
                  result = await app.requestData(
                    "/ui_printqueue", { method: "DELETE", query: { id: item.id } })
                  if(result && result.error){
                    return app.resultError(result)
                  }
                }
              }
              searchQueue()
            })
          }}
        /> 
      })
    }
  }

  const createReport = async (output) => {
    let _filters = [];
    data.edit.current.fieldvalue.forEach((rfdata) => {
      if (rfdata.selected) {
        if(rfdata.fieldtype === "bool"){
          _filters.push(`filters[${rfdata.name}]=${(rfdata.fieldtype)?"1":"0"}`)
        } else {
          _filters.push(`filters[${rfdata.name}]=${rfdata.value}`)
        }
      }
    })
    const report = data.edit.current.item
    const params = {
      type: "auto",
      template: report.reportkey, 
      title: report.reportkey,
      orient: report.orientation, 
      size: report.size,
      filters: _filters.join("&")
    }
    switch (output) {      
      case "xml":
        params.type = "xml"
        params.ctype = "application/xml; charset=UTF-8"
        break;
      
      case "csv":
        params.ctype = "text/csv; charset=UTF-8"
        break;

      default:
        params.ctype = "application/pdf"
        break;
    }
    const result = await app.requestData(reportPath(params), {})
    if(result && result.error){
      return app.resultError(result)
    }
    let resultUrl
    if(output === "csv"){
      var blob = new Blob([result], { type: 'text/csv;charset=utf-8;' })
      resultUrl = URL.createObjectURL(blob)
      output = "csv"
    } else {
      resultUrl = URL.createObjectURL(result, {type : params.ctype})
    }
    if(output === "print"){
      printJS({
        printable: resultUrl,
        type: 'pdf',
        base64: false,
      })
    } else {
      let filename = params.title+"_"+formatISO(new Date(), { representation: 'date' })+"."+output
      filename = filename.split("/").join("_")
      return saveToDisk(resultUrl, filename)
    }
  }

  const reportSettings = () => {
    const direction = (data.edit.current.type === "trans") ? 
      data.edit.dataset.groups.filter(
        item => (item.id === data.edit.current.item.direction))[0].groupvalue : "out"
    let defTemplate = (data.edit.current.type === "trans") ?
      data.edit.dataset.settings.filter(
        item => (item.fieldname === "default_trans_"+data.edit.current.transtype+"_"+direction+"_report"))[0] :
      data.edit.dataset.settings.filter(
        item => (item.fieldname === "default_"+data.edit.current.type+"_report"))[0]
    let templates = []
    data.edit.dataset.report.forEach(template => {
      let audit = data.login.data.audit.filter(item => (
        (item.nervatypeName === "report") && (item.subtype === template.id)))[0]
      if(audit){
        audit= audit.inputfilterName
      } else {
        audit = "all"
      }
      if (audit !== "disabled") {
        if (data.edit.current.type==="trans") {
          if (data.edit.current.item.direction === template.direction) {
            templates.push({
              value: template.reportkey, text: template.repname
            })
          }
        } else {
          templates.push({
            value: template.reportkey, text: template.repname
          })
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
    setData("current", { modalForm: 
      <Report
        title={data.edit.current.item[data.edit.template.options.title_field]}
        template={defTemplate}
        templates={templates}
        orient={getSetting("page_orient")}
        size={getSetting("page_size")}
        copy={1}
        getText={app.getText}
        onClose={()=>setData("current", { modalForm: null })}
        onOutput={(params) => {
          setData("current", { modalForm: null }, ()=>{
            reportOutput(params)
          })
        }}
      />,
      side: "hide" 
    })
  }

  const setEditor = (options, form, iedit) => {
    let edit = update({}, {$set: iedit||data.edit })
    if ((typeof edit.dataset[edit.current.type] === "undefined") || 
      (edit.dataset[edit.current.type].length===0)) {
      app.showToast({ type: "error", autoClose: false,
        title: app.getText("msg_warning"), 
        message: app.getText("msg_editor_invalid") })
      return false;
    }
    
    edit = update(edit, {$merge: {
      template: form,
      panel: form.options.panel,
      caption: form.options.title,
      audit: app.getAuditFilter(edit.current.type, edit.current.transtype)[0]
    }})
    if (edit.audit==="disabled") {
      return false
    }

    if (edit.dataset[edit.current.type][0].id === null) {
      edit = update(edit, { current: {$merge: {
        item: edit.dataset[edit.current.type][0]
      }}})
      if (form.options.search_form) {
        edit = update(edit, {$merge: {
          title_field: form.options.title_field
        }})
      } else {
        if((edit.audit==="all") || (edit.audit==="update")){
          edit = update(edit, {$merge: {
            dirty: true
          }})
        }
        edit = update(edit, {$merge: {
          title_field: app.getText("label_new")+" "+form.options.title
        }})
      }
      if (typeof form.options.extend !== "undefined") {
        edit = update(edit, {current: {$merge: {
          extend: initItem({tablename: form.options.extend, 
            dataset: edit.dataset, current: edit.current})
        }}})
      }
    } else {
      edit = update(edit, {current: {$merge: {
        item: update(initItem({tablename: edit.current.type, 
          dataset: edit.dataset, current: edit.current}), {
          $merge: edit.dataset[edit.current.type][0]})
      }}})
      if (typeof form.options.extend !== "undefined") {
        edit = update(edit, {current: {$merge: {
          extend: initItem({tablename: form.options.extend, 
            dataset: edit.dataset, current: edit.current})
        }}})
        if (typeof edit.dataset[form.options.extend] !== "undefined") {
          if (edit.dataset[form.options.extend].length > 0) {
            edit = update(edit, {current: {$merge: {
              extend: update(edit.current.extend, {
                $merge: edit.dataset[form.options.extend][0]})
            }}})
          }
        }
      }
    }

    edit = update(edit, {current: {$merge: {
      state: "normal",
      page: 0
    }}})
    if (edit.current.type === "trans") {
      if (typeof edit.dataset.trans[0].transcast !== "undefined") {
        if (edit.dataset.trans[0].transcast === "cancellation") {
          edit = update(edit, {current: {$merge: {
            state: "cancellation"
          }}})
        }
      }
      if (edit.dataset.pattern){
        const template = edit.dataset.pattern.filter((item) => (item.defpattern === 1))[0]
        edit = update(edit, {current: {$merge: {
          template: (template) ? template.id : "" 
        }}})
      }
    }
    if (edit.current.state === "normal" && edit.current.item.deleted === 1) {
      edit = update(edit, {current: {$merge: {
        state: "deleted"
      }}})
    } else if (edit.current.item.closed === 1) {
      edit = update(edit, {current: {$merge: {
        state: "closed"
      }}})
    }

    edit = update(edit, {current: {$merge: {
      fieldvalue: edit.dataset.fieldvalue || []
    }}})

    Object.keys(edit.template.view).forEach(vname => {
      edit = update(edit, {template: {view: { [vname]: {$merge: {
        view_audit: "all"
      }}}}})
      if (vname === "setting") {
        edit = update(edit, {template: {view: { [vname]: {$merge: {
          view_audit: "disabled"
        }}}}})
      } else {
        edit = update(edit, {template: {view: { [vname]: {$merge: {
          view_audit: app.getAuditFilter(form.view[vname].audit_type || vname, 
            form.view[vname].audit_transtype || null)[0]
        }}}}})
      }
    });

    if (edit.current.type === "report") {
      const template = JSON.parse(edit.dataset.report[0].report)
      Object.keys(template.fields).forEach((fieldname, index) => {
        const rfdata = template.fields[fieldname]
        const selected = (rfdata.selected)?
          rfdata.selected:
          (rfdata.wheretype === 'in')?true:false
        let tfrow = update({}, {$set: {
          id: index+1,
          rowtype: "reportfield", 
          datatype: rfdata.fieldtype,
          name: fieldname, 
          label: rfdata.description, 
          selected: selected,
          empty: (rfdata.wheretype === 'in') ? 'false' : 'true',
          value: rfdata.value
        }})
        switch (rfdata.fieldtype) {
          case "bool":
            break;
          case "valuelist":
            tfrow = update(tfrow, {$merge: {
              description: (rfdata.wheretype !== "in") ? 
                rfdata.valuelist.split("|").unshift("") : rfdata.valuelist.split("|")
            }})
            break;
          case "date":
            if(typeof(tfrow.value) === "undefined"){
              if (rfdata.defvalue) {
                tfrow = update(tfrow, {$merge: {
                  value: formatISO(addDays(new Date(), parseInt(rfdata.defvalue,10)), { representation: 'date' })
                }})
              } else if (rfdata.wheretype === "in") {
                tfrow = update(tfrow, {$merge: {
                  value: formatISO(new Date(), { representation: 'date' })
                }})
              } else {
                tfrow = update(tfrow, {$merge: {
                  value: ""
                }})
              }
            }
            break;
          case "integer":
          case "float":
            if(typeof(tfrow.value) === "undefined"){
              tfrow = update(tfrow, {$merge: {
                value: (rfdata.defvalue && rfdata.defvalue !== "") ? rfdata.defvalue : "0"
              }})
            }
            break;
          default:
            tfrow = update(tfrow, {$merge: {
              datatype: "string"
            }})
            break;
        }
        if (typeof tfrow.value === "undefined") {
          tfrow = update(tfrow, {$merge: {
            value: (rfdata.defvalue) ? rfdata.defvalue : ""
          }})
        }
        edit = update(edit, {current: {fieldvalue: {
          $push: [tfrow]
        }}})
      });
    }

    if(options.shipping){
      edit = update(edit, {current: {$merge: {
        form_type: "transitem_shipping",
        direction: edit.dataset.groups.filter((group)=> {
          return (group.id === edit.current.item.direction)
        })[0].groupvalue
      }}})
      if (typeof edit.dataset.shiptemp === "undefined") {
        edit = update(edit, {dataset: {$merge: {
          shiptemp: []
        }}})
      }
      edit = update(edit, {current: {item: {$merge: {
        delivery_type: app.getText("delivery_"+edit.current.direction)
      }}}})
      if(edit.current.shippingdate){
        edit = update(edit, {current: {item: {$merge: {
          shippingdate: edit.current.shippingdate
        }}}})
      } else {
        edit = update(edit, {current: {$merge: {
          shippingdate: formatISO(new Date(), { representation: 'date' })+"T00:00:00"
        }}})
        edit = update(edit, {current: {item: {$merge: {
          shippingdate: formatISO(new Date(), { representation: 'date' })+"T00:00:00"
        }}}})
      }
      if(edit.current.shipping_place_id){
        edit = update(edit, {current: {item: {$merge: {
          shipping_place_id: edit.current.shipping_place_id
        }}}})
      } else{
        edit = update(edit, {current: {$merge: {
          shipping_place_id: null
        }}})
        edit = update(edit, {current: {item: {$merge: {
          shipping_place_id: null
        }}}})
      }

      edit = update(edit, {dataset: {$merge: {
        shipping_items_: []
      }}})
      edit.dataset.shipping_items.forEach((item, index) => {
        let oitem = update(item, {$merge: {
          id: index+1,
          qty: parseFloat(item.qty)
        }})
        const mitem = edit.dataset.transitem_shipping.filter((item)=> {
          return (item.id === oitem.item_id+"-"+oitem.product_id)
        })[0]
        if (typeof mitem !== "undefined") {
          const tqty = (edit.current.direction === "out")?-parseFloat(mitem.sqty) : parseFloat(mitem.sqty)
          oitem = update(oitem, {$merge: {
            tqty: tqty,
            diff: parseFloat(oitem.qty) - tqty
          }})
        } else {
          oitem = update(oitem, {$merge: {
            qty: parseFloat(oitem.qty),
            tqty: 0,
            diff: parseFloat(oitem.qty)
          }})
        }
        const sitem = edit.dataset.shiptemp.filter((item)=> {
          return (item.id === oitem.item_id+"-"+oitem.product_id)
        })[0]
        if (typeof sitem !== "undefined") {
          oitem = update(oitem, {$merge: {
            edited: true
          }})
        }
        edit = update(edit, {dataset: {shipping_items_: {
          $push: [oitem]
        }}})
      });
    }

    if (edit.current.type === "printqueue") {
      if (typeof edit.printqueue === "undefined") {
        edit = update(edit, {$merge: {
          printqueue: edit.current.item
        }})
      } else {
        edit = update(edit, {current: {item: {$merge: {
          ...edit.printqueue
        }}}})
      }
    }

    edit = update(edit, {panel: {$merge: {
      form: true,
      state: edit.current.state
    }}})
    if (edit.panel.state !== "normal") {
      edit = update(edit, {$merge: {
        audit: "readonly"
      }})
    }
    if(edit.audit === "readonly") {
      edit = update(edit, {panel: {$merge: {
        save: false, link: false, delete: false, new: false,
        pattern: false, password: false, formula: false
      }}})  
      if (edit.panel.state !== "deleted") {
        edit = update(edit, {panel: {$merge: {
          trans: false
        }}})
      }
    }
    if (edit.audit !== "all") {
      edit = update(edit, {panel: {$merge: {
        delete: false,
        new: false
      }}})
      if (edit.panel.state !== "deleted") {
        edit = update(edit, {panel: {$merge: {
          trans: false
        }}})
      }
    }
    if (edit.panel.state === "deleted") {
      edit = update(edit, {panel: {$merge: {
        copy: false, 
        create: false
      }}})
    }
    edit = update(edit, {current: {$merge: {
      view: options.form||'form'
    }}})
    setData("edit", edit)
    setData("current", { module: "edit", edit: true, side: "hide" })
  }

  const loadEditor = async (params) => {
    let { ntype, ttype, id } = params;
    let edit = update({}, {$set: {
      dataset: { },
      current: { type: ntype, transtype: ttype },
      dirty: false,
      form_dirty: false,
      preview: null
    }})
    let proitem;
    if (id===null) {
      proitem = initItem({tablename: ntype, transtype: ttype, 
        dataset: edit.dataset, current: edit.current});
    };
    let views = []
    dataset[ntype](ttype).forEach(info => {
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
        if (typeof sql[ntype][info.sqlKey] !== "undefined") {
          _sql = sql[ntype][info.sqlKey](ntype)
        } else if (typeof sql[ntype][info.infoName] !== "undefined") {
          _sql = sql[ntype][info.infoName](ntype)
        } else {
          _sql = sql["all"][info.infoName](ntype)
        }
      }
      const sqlInfo = getSql(data.login.data.engine, _sql)
      if( (id !== null) || (sqlInfo.prmCount === 0)){
        views = update(views, {$push: [{
          key: info.infoName,
          text: sqlInfo.sql,
          values: ((sqlInfo.prmCount>0) && (id !== null)) ? Array(sqlInfo.prmCount).fill(id) : []
        }]})
      } else {
        edit = update(edit, { dataset: {$merge: {
          [info.infoName]: []
        }}})
      }
    })

    if(views.length > 0){
      if (ntype !== "report") {
        dataset["report"]().forEach(info => {
          let view = {
            key: info.infoName,
            sql: sql.report[info.infoName](ntype),
            values: []
          }
          if(info.infoName === "report"){
            if(ntype !== "printqueue"){
              view.values.push(data.login.data.employee.usergroup)
              view.values.push(edit.current.type)
            }
          } else if(info.infoName === "message"){
            view.values.push(edit.current.type)
            view.values.push(edit.current.type)
          } else {
            view.values.push(edit.current.type)
          }
          if (edit.current.type ==="trans") {
            const _where = ["and","r.transtype","=",[[],{select:["id"], from:"groups", 
                where:[["groupname","=","'transtype'"],["and","groupvalue","=","?"]]}]]
            if(info.infoName === "message"){
              view.sql.where[0][2][0].from[0][0][0].where[0].push(_where)
              view.sql.where[0][2][0].from[0][0][1].where[0].push(_where)
              view.values.splice(1, 0, edit.current.transtype)
            } else {
              view.sql.where.push(_where)
            }
            view.values.push(edit.current.transtype)
          }
          views = update(views, {$push: [{
            key: view.key,
            text: getSql(data.login.data.engine, view.sql).sql,
            values: view.values
          }]})
        })
      }

      let options = { method: "POST", data: views }
      let view = await app.requestData("/view", options)
      if(view.error){
        return app.resultError(view)
      }
      edit = update(edit, {dataset: {
        $merge: view
      }})
      if (id===null) {
        if (proitem === null) {
          proitem = initItem({tablename: ntype, transtype: ttype, 
            dataset: edit.dataset, current: edit.current});
        }
        if (ttype === "delivery") {
          proitem = update(proitem, {$merge: {
            direction: edit.dataset.groups.filter((group)=> {
              return ((group.groupname === "direction") && (group.groupvalue === "transfer"))
            })[0].id
          }})
        }
        edit = update(edit, { dataset: {$merge: {
          [ntype]: [proitem]
        }}})
      }
      setData("edit", edit, ()=>{
        if (!params.cb_key || (params.cb_key ==="SET_EDITOR")) {
          if (ntype==="trans") {
            if(params.shipping){
              return setEditor(params, forms["shipping"](edit.dataset[ntype][0], edit, getSetting("ui")), edit)
            } else {
              return setEditor(params, forms[ttype](edit.dataset[ntype][0], edit, getSetting("ui")), edit)
            }
          } else {
            return setEditor(params, forms[ntype](edit.dataset[ntype][0], edit, getSetting("ui")), edit)
          }
        }
      })
    } else {
      edit = update(edit, { dataset: {$merge: {
        [ntype]: [initItem({ tablename: ntype, transtype: ttype, 
          dataset: edit.dataset, current: edit.current })]
      }}})
      setData("edit", edit, ()=>{
        if (!params.cb_key || (params.cb_key ==="SET_EDITOR")) {
          return setEditor(params, forms[ntype](edit.dataset[ntype][0], edit, getSetting("ui")), edit)
        }
      })
    }
  }

  const setEditorItem = (options) => {
    let edit = update(data.edit, {})
    let dkey = forms[options.fkey]({}, edit, getSetting("ui")).options.data
    if (typeof dkey === "undefined") {
      dkey = options.fkey
    }
    edit = update(edit, {$merge: {
      form_dirty: false
    }})
    edit = update(edit, {current: {$merge: {
      form_type: options.fkey,
      form_datatype: dkey
    }}})
    edit = update(edit, {current: {$merge: {
      form: initItem({tablename: dkey, dataset: edit.dataset, current: edit.current})
    }}})

    if (options.id!==null) {
      edit = update(edit, {current: {$merge: {
        form: edit.dataset[options.fkey].filter((item)=> {
          return (item.id === parseInt(options.id,10))
        })[0]
      }}})
    } else {
      edit = update(edit, {current: {$merge: {
        form_dirty: true
      }}})
    }
    edit = update(edit, {current: {$merge: {
      form_template: forms[options.fkey](edit.current.form, edit, getSetting("ui"))
    }}})

    switch (options.fkey) {
      case "price":
      case "discount":
        edit = update(edit, {current: {$merge: {
          price_link_customer: null,
          price_customer_id: null
        }}})
        if (options.id!==null) {
          edit = update(edit, {current: {$merge: {
            price_link_customer: edit.current.form.link_customer,
            price_customer_id: edit.current.form.customer_id
          }}})
        }
        break;
      
      case "invoice_link":
        edit = update(edit, {current: {$merge: {
          invoice_link_fieldvalue: edit.dataset.invoice_link_fieldvalue
        }}})
        let invoice_props = { 
          id: edit.current.form.id, ref_id_1: "", transnumber: "", curr: ""
        }
        let invoice_link = edit.dataset.invoice_link.filter((item)=> {
          return (item.id === edit.current.form.id)
        })[0]
        if (typeof invoice_link !== "undefined") {
          invoice_props = invoice_link
        }
        edit = update(edit, {current: {$merge: {
          invoice_link: [invoice_props]
        }}})
        break;
      
      case "payment_link":
        edit = update(edit, {current: {$merge: {
          payment_link_fieldvalue: edit.dataset.payment_link_fieldvalue
        }}})
        let payment_props = { 
          id: edit.current.form.id, ref_id_2: "", transnumber: "", curr: ""
        }
        let payment_link = edit.dataset.payment_link.filter((item)=> {
          return (item.id === edit.current.form.id)
        })[0]
        if (typeof payment_link !== "undefined") {
          payment_props = payment_link
        }
        edit = update(edit, {current: {$merge: {
          payment_link: [payment_props]
        }}})
        if(options.link_field){
          edit = update(edit, {current: {form: {$merge: {
            [options.link_field]: options.link_id
          }}}})
        }
        break;
      
      default:
        break;}
    
    let panel = update(edit.current.form_template.options.panel, {$merge: {
      form: true,
      state: edit.current.state
    }})
    if (edit.panel.state !== "normal") {
      edit = update(edit, {$merge: {
        audit: "readonly"
      }})
    }
    if (edit.audit === "readonly") {
      panel = update(panel, {$merge: {
        save: false,
      }})
    }
    if (edit.audit !== "all") {
      panel = update(panel, {$merge: {
        link: false,
        delete: false,
        new: false
      }})
    }
    edit = update(edit, {$merge: {
      panel: panel
    }})
    setData("edit", edit)
  }

  const newFieldvalue = (_fieldname) => {
    const updateFieldvalue = async (item) => {
      const options = { method: "POST", data: [item] }
      const result = await app.requestData("/fieldvalue", options)
      if(result.error){
        return app.resultError(result)
      }
      loadEditor({
        ntype: data.edit.current.type, 
        ttype: data.edit.current.transtype, 
        id: data.edit.current.item.id, 
        form: "fieldvalue", form_id: result[0]
      })
    }

    if (_fieldname!=="") {
      const deffield = data.edit.dataset.deffield.filter((item) => (item.fieldname === _fieldname))[0]
      const _fieldtype = data.login.data.groups.filter((item) => (item.id === deffield.fieldtype))[0].groupvalue
      let item = update(initItem({tablename: "fieldvalue"}), {$merge: {
        id: null,
        fieldname: deffield.fieldname
      }})
      let _selector = false;
      switch (_fieldtype) {
        case "bool":
          item = update(item, {$merge: {
            value: "false"
          }})
          break;

        case "date":
          item = update(item, {$merge: {
            value: formatISO(new Date(), { representation: 'date' })
          }})
          break;

        case "time":
          item = update(item, {$merge: {
            value: "00:00"
          }})
          break;

        case "float":
        case "integer":
          item = update(item, {$merge: {
            value: "0"
          }})
          break;

        case "valuelist":
          item = update(item, {$merge: {
            value: deffield.valuelist.split("|")[0]
          }})
          break;

        case "customer":
        case "tool":
        case "trans":
        case "transitem":
        case "transmovement":
        case "transpayment":
        case "product":
        case "project":
        case "employee":
        case "place":
          _selector = true;
          break;

        default:
          break;
      }
      if (_selector) {
        app.onSelector(_fieldtype, "", (row, filter)=>{
          const params = row.id.split("/")
          item = update(item, {$merge: {
            value: String(parseInt(params[2],10))
          }})
          updateFieldvalue(item)
        })
      } else {
        updateFieldvalue(item)
      }
    } else {
      return app.showToast({ type: "error", title: app.getText("msg_warning"), 
          message: app.getText("fields_deffield_missing") })
    }
  }

  const deleteEditor = () => {
    const clearEditor = () => {
      setData("edit", { dataset: {}, current: {}, dirty: false, form_dirty: false })
      setData("current", { module: "search" })
    }
    const deleteData = async () => {
      const result = await app.requestData("/"+data.edit.current.type, 
        { method: "DELETE", query: { id: data.edit.current.item.id } })
      if(result && result.error){
        return app.resultError(result)
      }
      await app.createHistory("delete")
      clearEditor()
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
            if (data.edit.current.item.id === null) {
              clearEditor()
            } else {
              if (typeof sql[data.edit.current.type]["delete_state"] !== "undefined") {
                const sqlInfo = getSql(data.login.data.engine, sql[data.edit.current.type]["delete_state"]())
                const params = { 
                  method: "POST", 
                  data: [{ 
                    key: "state",
                    text: sqlInfo.sql,
                    values: Array(sqlInfo.prmCount).fill(data.edit.current.item.id)
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
            }
          })
        }}
      />,
      side: "hide"
    })
  }

  const deleteEditorItem = (params) => {
    const reLoad = () => {
      loadEditor({
        ntype: data.edit.current.type, 
        ttype: data.edit.current.transtype, 
        id: data.edit.current.item.id, 
        form: params.fkey
      })
    }
    const deleteItem = async () => {
      if (params.id === null) {
        reLoad()
      } else {
        const table = (!params.table) ? params.fkey : params.table
        const result = await app.requestData(
          "/"+table, { method: "DELETE", query: { id: params.id } })
        if(result && result.error){
          return app.resultError(result)
        }
        await app.createHistory("save")
        reLoad()
      }
    }

    if(params.prompt){
      deleteItem()
    } else {
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
            setData("current", { modalForm: null }, ()=>{
              deleteItem()
            })
          }}
        />,
        side: "hide"
      })
    }
  }

  const setFieldvalue = (recordset, fieldname, ref_id, defvalue, value) => {
    const fieldvalue_idx = recordset.findIndex((item)=>((item.ref_id === ref_id)&&(item.fieldname === fieldname)))
    if (fieldvalue_idx === -1) {
      const fieldvalue = update(initItem({tablename: "fieldvalue", current: data.edit.current}), {$merge: {
        fieldname: fieldname,
        ref_id: ref_id,
        value: ((typeof value === "undefined") || (value === null)) ? defvalue : value
      }})
      recordset = update(recordset, {$push: [fieldvalue]})
    } else if(value) {
      recordset = update(recordset, { [fieldvalue_idx]: {$merge: {
        value: value
      }}})
    }
    return recordset
  }

  const tableValues = (type, item) => {
    let values = {}
    const baseValues = initItem({tablename: type, 
      dataset: data.edit.dataset, current: data.edit.current})
    for (const key in item) {
      if (baseValues.hasOwnProperty(key)) {
        values[key] = item[key]
      }
    }
    return values
  }

  const saveEditorForm = async () => {
    let edit = update(data.edit, {})

    let values = tableValues(edit.current.form_datatype, edit.current.form)
    values = await validator(edit.current.form_datatype, values)
    if(values.error){
      app.resultError(values)
      return null
    }
    
    let result = await app.requestData("/"+edit.current.form_datatype, { method: "POST", data: [values] })
    if(result.error){
      app.resultError(result)
      return null
    }
    if (edit.current.form.id === null) {
      edit = update(edit, {current: { form: {$merge: {
        id: result[0]
      }}}})
    }
    edit = update(edit, {$merge: {
      form_dirty: false
    }})
    await app.createHistory("save")
   
    switch (edit.current.form_type) {
      case "movement":
        if (edit.current.transtype === "delivery") {
          let movements = []; let mlink = null;
          let movement = edit.dataset.movement_transfer.filter(
            item => (item.id === edit.current.form.id))[0]
          if (typeof movement !== "undefined") {
            movement = edit.dataset.movement.filter(
              item => (item.id === movement.ref_id))[0]
            movement = update(
              initItem({tablename: "movement", dataset: edit.dataset, current: edit.current}), 
              {$merge: movement})
          } else {
            movement = update(
              initItem({tablename: "movement", dataset: edit.dataset, current: edit.current}),
              {$merge: {
                place_id: edit.current.item.place_id
              }}
            )
            mlink = update(
              initItem({tablename: "link", dataset: edit.dataset, current: edit.current}),
              {$merge: {
                nervatype_1: data.login.data.groups.filter(
                  (item)=>((item.groupname === "nervatype") && (item.groupvalue === "movement")))[0].id,
                nervatype_2: data.login.data.groups.filter(
                  (item)=>((item.groupname === "nervatype") && (item.groupvalue === "movement")))[0].id,
                ref_id_2: edit.current.form.id
              }}
            )
          }
          movement = update(movement, {$merge: {
            product_id: edit.current.form.product_id,
            qty: -(edit.current.form.qty),
            notes: edit.current.form.notes
          }})
          movements.push(tableValues("movement", movement));
          edit.dataset.movement_transfer.forEach((mvt) => {
            movement = update(
              initItem({tablename: "movement", dataset: edit.dataset, current: edit.current}), 
              {$merge: tableValues("movement", mvt)})
            if ((movement.id !== edit.current.form.id) &&
              (movement.place_id !== edit.current.form.place_id)){
                movement = update(movement, {$merge: {
                  place_id: edit.current.form.place_id
                }})
                movements.push(movement)
            }
          });
          let result = await app.requestData("/movement", { method: "POST", data: movements })
          if(result.error){
            app.resultError(result)
            return null
          }
          if (mlink !== null) {
            mlink = update(mlink, {$merge: {
              ref_id_1: result[0]
            }})
            result = await app.requestData("/link", { method: "POST", data: [mlink] })
            if(result.error){
              app.resultError(result)
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
            //delete link
            result = await app.requestData("/link", 
              { method: "DELETE", query: { id: edit.current.price_link_customer } })
            if(result && result.error){
              app.resultError(result)
              return null
            }
          } else {
            let clink = update(
              initItem({tablename: "link", dataset: edit.dataset, current: edit.current}),
              {$merge: {
                id: edit.current.price_link_customer, //update or insert
                nervatype_1: data.login.data.groups.filter(
                  (item)=>((item.groupname === "nervatype") && (item.groupvalue === "price")))[0].id,
                ref_id_1: edit.current.form.id,
                nervatype_2: data.login.data.groups.filter(
                  (item)=>((item.groupname === "nervatype") && (item.groupvalue === "customer")))[0].id,
                ref_id_2: edit.current.price_customer_id
              }}
            )
            result = await app.requestData("/link", { method: "POST", data: [clink] })
            if(result.error){
              app.resultError(result)
              return null
            }
          }
        }
        break;
        
      case "invoice_link":
      case "payment_link":
        let rsname = edit.current.form_type+"_fieldvalue"
        for (let i=0; i < edit.current[rsname].length; i++) {
          if (edit.current[rsname][i].ref_id === null) {
            edit = update(edit, {
              current:{
                [rsname]: {
                  [i]: {$merge: {
                    ref_id: edit.current.form.id
                  }}
                }
              }
            })
          }
        }
        let flist = { link_qty:"0", link_rate: "1" }
        let fvalues = edit.current[rsname]
        Object.keys(flist).forEach(function(fieldname) { 
          fvalues = setFieldvalue(fvalues, 
            fieldname, edit.current.form.id, flist[fieldname])
        });
        edit = update(edit, {
          current: {$merge: {
            [rsname]: fvalues
          }}
        })
        result = await app.requestData("/fieldvalue", { method: "POST", data: fvalues })
        if(result.error){
          app.resultError(result)
          return null
        }
        break;

      default:
        break;
    }
    
    return edit
  }

  const checkSubtype = (type, subtype, item) => {
    if(subtype !== null){
      switch (type) {
        case "customer":
          return (subtype === item.custtype)
        case "place":
          return (subtype === item.placetype)
        case "product":
          return (subtype === item.protype)
        case "tool":
          return (subtype === item.toolgroup)
        case "trans":
          return (subtype === item.transtype)
        default:
          break
      }
    }
    return true
  }

  const saveEditor = async () => {
    let edit = update(data.edit, {})

    let values = tableValues(edit.current.type, edit.current.item)
    values = await validator(edit.current.type, values)
    if(values.error){
      return app.resultError(values)
    }

    if (edit.current.item.id === null && edit.dataset.deffield) {
      edit.dataset.deffield.forEach((deffield) => {
        if(deffield.addnew === 1){
          const subtype = checkSubtype(edit.current.type,
            deffield.subtype, edit.current.item);
          const item = edit.current.fieldvalue.filter(item => (item.fieldname === deffield.fieldname))[0]
          if(!item && subtype){
            const fieldtype = edit.dataset.groups.filter(item => (item.id === deffield.fieldtype))[0].groupvalue
            switch (fieldtype) {
              case "bool":
              case "integer":
              case "float":
                edit = update(edit, { current: {$merge: {
                  fieldvalue: setFieldvalue(edit.current.fieldvalue, 
                    deffield.fieldname, null, null, 0)
                }}})
                break;
              case "valuelist":
                edit = update(edit, { current: {$merge: {
                  fieldvalue: setFieldvalue(edit.current.fieldvalue, 
                    deffield.fieldname, null, null, deffield.valuelist.split("|")[0])
                }}})
                break;
              default:
            }
          }
        }
      });
    }

    let result = await app.requestData("/"+edit.current.type, { method: "POST", data: [values] })
    if(result.error){
      app.resultError(result)
      return null
    }
    if (edit.current.item.id === null) {
      edit = update(edit, {current: { item: {$merge: {
        id: result[0]
      }}}})
    }
    edit = update(edit, {$merge: {
      dirty: false
    }})
    await app.createHistory("save")

    if (typeof edit.current.extend !== "undefined") {
      if (edit.current.extend.ref_id === null) {
        edit = update(edit, {current: { extend: {$merge: {
          ref_id: edit.current.item.id
        }}}})
      }
      if (edit.current.extend.trans_id === null) {
        edit = update(edit, {current: { extend: {$merge: {
          trans_id: edit.current.item.id
        }}}})
      }
    }
    if (typeof edit.current.fieldvalue !== "undefined") {
      for (let i=0; i < edit.current.fieldvalue.length; i++) {
        if (edit.current.fieldvalue[i].ref_id === null) {
          edit = update(edit, {current: { fieldvalue: { [i]: {$merge: {
            ref_id: edit.current.item.id
          }}}}})
        }
      }
    }

    if (edit.current.type === "trans") {
      edit = update(edit, { current: {$merge: {
        fieldvalue: setFieldvalue(edit.current.fieldvalue, 
          "trans_transcast", edit.current.item.id, null, "normal")
      }}})
      switch (edit.current.transtype) {
        case "invoice":
          const params = { 
            method: "POST", 
            data: [{ 
              key: "fields",
              text: getSql(data.login.data.engine, sql.trans.invoice_customer()).sql,
              values: [edit.current.item.customer_id]
            }]
          }
          let view = await app.requestData("/view", params)
          if(view.error){
            return app.resultError(view)
          }
          if (view.fields.length > 0) {
            Object.keys(view.fields[0]).forEach((fieldname) => {
              edit = update(edit, { current: {$merge: {
                fieldvalue: setFieldvalue(edit.current.fieldvalue, 
                  fieldname, edit.current.item.id, null, view.fields[0][fieldname])
              }}})
            })
          }
          break;
        case "worksheet":
          let wlist = {trans_wsdistance:0, trans_wsrepair:0, trans_wstotal:0, trans_wsnote:""};
          Object.keys(wlist).forEach((fieldname) => {
            edit = update(edit, { current: {$merge: {
              fieldvalue: setFieldvalue(edit.current.fieldvalue, 
                fieldname, edit.current.item.id, null, wlist[fieldname])
            }}})
          })
          break;
        case "rent":
          let rlist = {trans_reholiday:0, trans_rebadtool:0, trans_reother:0, trans_rentnote:""};
          Object.keys(rlist).forEach((fieldname) => {
            edit = update(edit, { current: {$merge: {
              fieldvalue: setFieldvalue(edit.current.fieldvalue, 
                fieldname, edit.current.item.id, null, rlist[fieldname])
            }}})
          })
          break;
        case "delivery":
          let movements = [];
          edit.dataset.movement.forEach((mvt) => {
            let movement = update(
              initItem({tablename: "movement", dataset: edit.dataset, current: edit.current}),
              {$merge: mvt })
            if (!isEqual(parseISO(movement.shippingdate), parseISO(edit.current.item.transdate))){
              movement = update(movement, {$merge: {
                shippingdate: formatISO(parseISO(edit.current.item.transdate))
              }})
              movements.push(movement)
            }
          })
          if (movements.length > 0) {
            result = await app.requestData("/movement", { method: "POST", data: movements })
            if(result.error){
              app.resultError(result)
              return null
            }
          }
          break;
        default:
          break;
      }
    }

    if (edit.current.fieldvalue.length > 0) {
      result = await app.requestData("/fieldvalue", { method: "POST", data: edit.current.fieldvalue })
      if(result.error){
        app.resultError(result)
        return null
      }
      for (let index = 0; index < result.length; index++) {
        if(!edit.current.fieldvalue[index].id){
          edit = update(edit, { current: { fieldvalue: { [index]: {$merge: {
            id: result[index]
          }}}}})
        }
      }
    }

    if (typeof edit.current.extend !== "undefined") {
      let ptype = String(edit.template.options.extend).split("_")[0]
      let extend = update(edit.current.extend, {})
      switch (edit.current.transtype) {
        case "waybill":
          ptype = extend.seltype
          if (extend.seltype === "transitem") {
            ptype = "link"
            extend = initItem({tablename: "link", dataset: edit.dataset, current: edit.current})
            if (edit.dataset.translink.length > 0) {
              extend = update(extend, {$merge: edit.dataset.translink[0]})
            } else {
              extend = update(extend, {$merge: {
                nervatype_1: data.login.data.groups.filter(
                  (item)=>((item.groupname === "nervatype") && (item.groupvalue === "trans")))[0].id,
                nervatype_2: data.login.data.groups.filter(
                  (item)=>((item.groupname === "nervatype") && (item.groupvalue === "trans")))[0].id,
                ref_id_1: edit.current.item.id
              }})
            }
            extend = update(extend, {$merge: {
              ref_id_2: edit.current.extend.ref_id
            }})
          } else {
            extend = null;
            if (edit.dataset.translink.length > 0) {
              result = await app.requestData("/link", 
                { method: "DELETE", query: { id: edit.dataset.translink[0].id } })
              if(result && result.error){
                app.resultError(result)
                return null
              }
            }
          }
          break;
        case "formula":
          extend = update(extend, {$unset: ["product"]})
          extend = update(extend, {$merge: {
            shippingdate: formatISO(parseISO(edit.current.item.transdate))
          }})
          break;
        case "production":
          extend = update(extend, {$unset: ["product"]})
          extend = update(extend, {$merge: {
            shippingdate: edit.current.item.duedate,
            place_id: edit.current.item.place_id
          }})
          break;
        default:
      }
      if (extend !== null) {
        result = await app.requestData("/"+ptype, { method: "POST", data: [extend] })
        if(result.error){
          app.resultError(result)
          return null
        }
        if(extend.id === null){
          extend = update(extend, {$merge: {
            id: result[0]
          }})
        }
        edit = update(edit, { current: {$merge: {
          extend: extend
        }}})
      }
    }
    
    return edit
  }

  const calcFormula = (formula_id) => {
    setData("current", { modalForm: 
      <InputBox 
        title={app.getText("msg_warning")}
        message={app.getText("ms_load_formula")}
        infoText={app.getText("msg_delete_info")+" "+app.getText("ms_continue_warning")}
        defaultOK={true}
        labelOK={app.getText("msg_ok")}
        labelCancel={app.getText("msg_cancel")}
        onCancel={() => {
          setData("current", { modalForm: null })
        }}
        onOK={(value) => {
          setData("current", { modalForm: null }, async ()=>{
            const params = { 
              method: "POST", 
              data: [{ 
                key: "formula",
                text: getSql(data.login.data.engine, sql.trans.formula_items()).sql,
                values: [formula_id]
              }]
            }
            let view = await app.requestData("/view", params)
            if(view.error){
              return app.resultError(view)
            }
            let production_qty = data.edit.dataset.movement_head[0].qty
            let production_place = data.edit.dataset.movement_head[0].place_id
            let formula_qty = data.edit.dataset.formula_head.filter(item => (item.id === formula_id))[0].qty
            let items = [] 
            view.formula.forEach(fitem => {
              let item = update(
                initItem({tablename: "movement", dataset: data.edit.dataset, current: data.edit.current}), 
                {$merge: {
                  product_id: fitem.product_id,
                  place_id: (fitem.place_id === null) ? production_place : fitem.place_id,
                  qty: (fitem.shared === 1) ? 
                    -Math.ceil(production_qty/formula_qty) : 
                    -round((production_qty/formula_qty)*fitem.qty,2)
                }})
              items.push(item);
            })
            for (let index = 0; index < data.edit.dataset.movement.length; index++) {
              const result = await app.requestData(
                "/movement", { method: "DELETE", query: { id: data.edit.dataset.movement[index].id } })
              if(result && result.error){
                return app.resultError(result)
              }
            }
            const result = await app.requestData("/movement", { method: "POST", data: items })
            if(result.error){
              return app.resultError(result)
            }
            setData("current", { modalForm: null }, ()=>{
              loadEditor({
                ntype: data.edit.current.type, 
                ttype: data.edit.current.transtype, 
                id: data.edit.current.item.id, 
                form: "movement"
              })
            })
          })
        }}
      /> 
    })
  }

  const createTrans = async (options) => {

    const check_refnumber = (params) => {
      if(params.transtype === "waybill"){
        return "link";
      }else if(params.transtype==="delivery" && (params.transcast === "normal") && 
        (params.direction==="in" || params.direction==="out")){
        return "";
      } else if (params.cmdtype === "copy" && params.transcast === "normal") {
        return "";
      } else if ((params.transcast !== "normal") || params.refno){
        return "reflink";
      } else {
        return "refnumber";
      }
    }

    let base_trans = update(data.edit.dataset.trans[0], {});
    //set base data
    let transtype = data.edit.dataset.groups.filter(
      (item) => (item.id === base_trans.transtype)
    )[0].groupvalue
    let transtype_id = base_trans.transtype;
    let direction = data.edit.dataset.groups.filter(
      (item) => (item.id === base_trans.direction)
    )[0].groupvalue
    let direction_id = base_trans.direction;  
    if (typeof options.new_transtype !== "undefined" && typeof options.new_direction !== "undefined") {
      transtype = options.new_transtype;
      let audit = data.login.data.audit.filter(item => (
        (item.nervatypeName === "trans") && (item.subtypeName === transtype)))[0]
      if (typeof audit !== "undefined") {
        if (audit.item.inputfilterName==="disabled"){
          app.showToast({ type: "info",
            title: app.getText("msg_warning"), 
            message: app.getText("msg_create_disabled_err")+" "+transtype })
          return false;
        }
      }
      transtype_id = data.edit.dataset.groups.filter(
        (item) => ((item.groupname === "transtype") && (item.groupvalue === transtype))
      )[0].id
      direction = options.new_direction;
      direction_id = data.edit.dataset.groups.filter(
        (item) => ((item.groupname === "direction") && (item.groupvalue === direction))
      )[0].id
    }

    //to check some things...
    if ((transtype==="receipt" || transtype==="worksheet") && direction === "in") {
      app.showToast({ type: "info",
        title: app.getText("msg_warning"), 
        message: app.getText("msg_input_invalid")+" in" })
      return false;
    }
    if (base_trans.transcast==="cancellation") {
      app.showToast({ type: "info",
        title: app.getText("msg_warning"), 
        message: app.getText("msg_create_cancellation_err1") })
      return false;
    }
    if (options.transcast==="cancellation" && (transtype==="invoice" || transtype==="receipt") 
      && base_trans.deleted===0) {
        app.showToast({ type: "info",
          title: app.getText("msg_warning"), 
          message: app.getText("msg_create_cancellation_err2") })
        return false;
    }
    if (options.transcast==="cancellation" && (data.edit.dataset.cancel_link.length > 0)) {
      app.showToast({ type: "info",
        title: app.getText("msg_warning"), 
        message: app.getText("msg_create_cancellation_err3"+data.edit.dataset.cancel_link[0].transnumber) })
      return false;
    }
    if (options.transcast==="amendment" && base_trans.deleted===1) {
      app.showToast({ type: "info",
        title: app.getText("msg_warning"), 
        message: app.getText("msg_create_amendment_err") })
      return false;
    }

    //creat trans data from the original          
    let values = update({},{$set: { 
      id: null, 
      transtype: transtype_id, 
      transnumber: null, 
      crdate: formatISO(new Date(), { representation: 'date' }), 
      transdate: formatISO(new Date(), { representation: 'date' }), 
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
      transtate: data.edit.dataset.transtate.filter(
        item => ((item.groupname === "transtate") && (item.groupvalue === "ok")))[0].id,
      closed: 0, deleted: 0, 
      direction: direction_id, 
      cruser_id: data.login.data.employee.id,
      trans_transcast: options.transcast || "normal"
    }})
    if (base_trans.duedate !== null) {
      values.duedate = formatISO(new Date(), { representation: 'date' })+"T00:00:00";
    }
    if (transtype === "invoice" && direction === "out") {
      let default_deadline = data.edit.dataset.settings.filter((group)=> {
        return (group.fieldname === "default_deadline")
      })[0]
      if (typeof default_deadline !== "undefined") {
        values.duedate = formatISO(addDays(new Date(), parseInt(default_deadline.value,10)), { representation: 'date' })+"T00:00:00"
      }
    } else if (transtype === "receipt") {
      values.customer_id = null;
    }
    const _refnum = check_refnumber({
      transtype: transtype, transcast: options.transcast, direction: direction, 
      cmdtype: options.cmdtype, refno: options.refno})
    if((_refnum === "refnumber") || (_refnum === "reflink")){
      values.ref_transnumber = base_trans.transnumber
    }
    
    let nextnumber = transtype+"_"+direction;
    if (transtype === "waybill" || transtype === "cash") {
      nextnumber = transtype;
    }
    let params = { method: "POST", 
      data: {
        key: "nextNumber",
        values: {
          numberkey: nextnumber, 
          step: true
        }
      }
    }
    let result = await app.requestData("/function", params)
    if(result.error){
      app.resultError(result)
      return null
    }
    if (options.transcast === "cancellation") {
      values.transnumber = result+"/C";
      if (transtype !== "delivery" && transtype !== "inventory") {
        values.deleted = 1;
      }
      values.transdate = base_trans.transdate;
      values.duedate = base_trans.duedate;
    } else if (options.transcast === "amendment") {
      values.transnumber = result+"/A";
    } else {
      values.transnumber = result
    }

    result = await app.requestData("/trans", { method: "POST", data: [values] })
    if(result.error){
      app.resultError(result)
      return null
    }
    values.id = result[0];

    let fieldvalue = [];
    data.edit.current.fieldvalue.forEach((cfield) => {
      if(cfield.fieldname !== "trans_transcast"){
        let deffield = data.edit.dataset.deffield.filter(
          item => (item.fieldname === cfield.fieldname)
        )[0]
        let subtype = checkSubtype("trans", deffield.subtype, values);
        if ((cfield.deleted===0 && deffield.visible===1 && subtype) 
          || (cfield.deleted===0 && subtype && options.cmdtype === "copy")){
          let field = tableValues("fieldvalue", cfield)
          field.id = null; 
          field.ref_id = values.id; 
          fieldvalue.push(field);
        } 
      }
    })
    if (transtype === "invoice") {
      const params = { 
        method: "POST", 
        data: [{ 
          key: "fields",
          text: getSql(data.login.data.engine, sql.trans.invoice_customer()).sql,
          values: [values.customer_id]
        }]
      }
      let view = await app.requestData("/view", params)
      if(view.error){
        return app.resultError(view)
      }
      if (view.fields.length > 0) {
        Object.keys(view.fields[0]).forEach((fieldname) => {
          fieldvalue = setFieldvalue(fieldvalue, fieldname, values.id, view.fields[0][fieldname]) 
        })
      }
    }

    if(fieldvalue.length > 0){
      result = await app.requestData("/fieldvalue", { method: "POST", data: fieldvalue })
      if(result.error){
        app.resultError(result)
        return null
      }
    }

    if((_refnum === "link") || (_refnum === "reflink")){
      const link = update(initItem({tablename: "link"}), {$merge: {
        nervatype_1: data.edit.dataset.groups.filter(
          (item)=>((item.groupname === "nervatype") && (item.groupvalue === "trans")))[0].id,
        ref_id_1: values.id,
        nervatype_2: data.edit.dataset.groups.filter(
          (item)=>((item.groupname === "nervatype") && (item.groupvalue === "trans")))[0].id,
        ref_id_2: base_trans.id
      }})
      result = await app.requestData("/link", { method: "POST", data: [link] })
      if(result.error){
        app.resultError(result)
        return null
      }
    }

    let items = [];
    if (transtype==="invoice" || transtype==="receipt") {
      
      const get_product_qty = (items, product_id, deposit) => {
        let retvalue = 0;
        items.forEach((item) => {
          if ((item.product_id === product_id) && (item.deposit === deposit)){
            retvalue += item.qty;
          }
        });
        return retvalue;
      }
      
      const recalc_item = (item, rate, digit) => {
        item.netamount = round(item.fxprice*(1-item.discount/100)*item.qty, digit);
        item.vatamount = round(item.fxprice*(1-item.discount/100)*item.qty*rate, digit);
        item.amount = round(item.netamount+item.vatamount, digit);
        return item;
      }
      
      let products = {};
      if (options.from_inventory && data.edit.dataset.transitem_invoice) {
        //create from order,worksheet and rent, on base the delivery rows
        data.edit.dataset.transitem_shipping.forEach(inv_item => {
          const item = data.edit.dataset.item.filter(
            oitem => (oitem.id === inv_item.id)
          )[0]
          if (typeof item!=="undefined") {
            let iqty = inv_item.sqty;
            if(data.edit.dataset.groups.filter(
              group => (group.id === base_trans.direction))[0].groupvalue === "out"){
                iqty = -inv_item.sqty
            }
            if (item.deleted===0 && iqty>0) {
              if (!Object.keys(products).includes(item.product_id)){
                iqty -= get_product_qty(data.edit.dataset.transitem_invoice, 
                  item.product_id, 0);
                products[item.product_id] = true;
              }
              if (iqty !== 0){
                let sitem = tableValues("item", item)
                sitem.qty = iqty;
                sitem = recalc_item(sitem, item.rate, base_trans.digit);
                items.push(sitem);
              }
            }
          }
        });
      } else {
        data.edit.dataset.item.forEach(base_item => {
          if (base_item.deleted===0) {
            if (options.netto_qty && data.edit.dataset.transitem_invoice) {
              //create from order,worksheet and rent, on base the invoice rows
              let iqty = base_item.qty;
              if (!Object.keys(products).includes(base_item.product_id)){
                iqty -= get_product_qty(data.edit.dataset.transitem_invoice, 
                  base_item.product_id, 0);
                products[base_item.product_id] = true;
              }
              if (iqty !== 0){
                let sitem = tableValues("item", base_item)
                sitem.qty = iqty;
                sitem = recalc_item(sitem, base_item.rate, base_trans.digit);
                items.push(sitem);
              }
            } else {
              items.push(tableValues("item", base_item))
            }
          }
        });
      }
              
      //put to deposit rows
      items.forEach(item => {
        if (item.deposit === 1) {
          let dqty = get_product_qty(data.edit.dataset.transitem_invoice, 
            item.product_id, 1);
          if (dqty !== 0) {
            let sitem = tableValues("item", item)
            sitem.qty = -dqty;
            items.unshift(sitem);
          }
        }
      });
    } else {
      data.edit.dataset.item.forEach(item => {
        if (item.deleted===0) {
          items.push(tableValues("item", item))
        }
      });
    }
    
    items.forEach(item => {
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
        let sitem = tableValues("item", item)
        sitem.qty = -sitem.qty;
        sitem.netamount = -sitem.netamount;
        sitem.vatamount = -sitem.vatamount;
        sitem.amount = -sitem.amount;
        items.push(sitem);
      }
    });

    if (items.length > 0) {
      result = await app.requestData("/item", { method: "POST", data: items })
      if(result.error){
        app.resultError(result)
        return null
      }
    }

    let payments = [];
    data.edit.dataset.payment.forEach((base_payment) => {
      if (base_payment.deleted===0) {
        let payment = tableValues("payment", base_payment)
        payment.id = null;
        payment.trans_id = values.id;
        if (options.transcast === "cancellation") {
          payment.amount = -payment.amount;
        }
        payments.push(payment);
      }
    });
    if (payments.length > 0) {
      result = await app.requestData("/payment", { method: "POST", data: payments })
      if(result.error){
        app.resultError(result)
        return null
      }
    }

    let movements = []; let reflinks = [];
    if (transtype === "formula" || transtype === "production") {
      let movement = tableValues("movement", data.edit.dataset.movement_head[0])
      movement.id = null; 
      movement.trans_id = values.id;
      movements.push(movement);
    }
    let base_movements = data.edit.dataset.movement || [];
    base_movements.forEach((bmt) => {
      if (bmt.deleted === 0) {
        if(bmt.item_id || bmt.ref_id){
          reflinks.push({
            id:bmt.id, 
            item_id: bmt.item_id, 
            ref_id: bmt.ref_id 
          });
        }
        let movement = tableValues("movement", bmt)
        movement.id = null; 
        movement.trans_id = values.id;
        if (options.transcast==="cancellation") {
          movement.qty = -movement.qty;
        }
        movements.push(movement);
      }
    });
    if (movements.length > 0) {
      result = await app.requestData("/movement", { method: "POST", data: movements })
      if(result.error){
        app.resultError(result)
        return null
      }
      let links = [];
      let nt_movement = data.edit.dataset.groups.filter(
        (item)=>((item.groupname === "nervatype") && (item.groupvalue === "movement")))[0].id
      let nt_item = data.edit.dataset.groups.filter(
        (item)=>((item.groupname === "nervatype") && (item.groupvalue === "item")))[0].id
      for (let li=0; li < reflinks.length; li++) {
        let ilink = update(initItem({tablename: "link"}), {$merge: {
          nervatype_1: nt_movement
        }})
        if (reflinks[li].item_id !== null) { 
          ilink.ref_id_1 = result[li].id;
          ilink.nervatype_2 = nt_item;
          ilink.ref_id_2 = reflinks[li].item_id;
          links.push(ilink);
        } else if (reflinks[li].ref_id !== null) {
          ilink.ref_id_1 = result[data.edit.dataset.movement.findIndex(
            item => (item.id === reflinks[li].ref_id)
          )].id
          ilink.nervatype_2 = nt_movement;
          ilink.ref_id_2 = result[data.edit.dataset.movement.findIndex(
            item => (item.id === reflinks[li].id)
          )].id
          links.push(ilink);
        }
      }
      if (links.length > 0) {
        result = await app.requestData("/link", { method: "POST", data: links })
        if(result.error){
          app.resultError(result)
          return null
        }
      }
    }

    await app.createHistory("save")
    loadEditor({ ntype: "trans", ttype: transtype, id: values.id })
  }

  const createTransOptions = () => {
    let edit = update(data.edit, {})
    let options = {
      directions: ["in","out"],
      baseTranstype: edit.current.transtype,
      transtype: edit.current.transtype,
      direction: edit.dataset.groups.filter((group)=> {
          return (group.id === edit.current.item.direction)
        })[0].groupvalue,
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
    setData("current", { modalForm: 
      <Trans 
        {...options}
        getText={app.getText}
        onClose={()=>setData("current", { modalForm: null })}
        onCreate={(result) => {
          setData("current", { modalForm: null }, ()=>{
            createTrans({
              cmdtype: "create", transcast: "normal", 
              new_transtype: result.newTranstype, 
              new_direction: result.newDirection, 
              refno: result.refno, 
              from_inventory: result.fromInventory, 
              netto_qty: result.nettoQty
            })
          });
        }}
      />,
      side: "hide"
    })
  }

  const checkEditor = (options, cbKeyTrue, cbKeyFalse) => {
    const cbNext = (cbKey) =>{
      switch (cbKey) {
        case "LOAD_EDITOR":
          loadEditor(options);
          break;
        case "SET_EDITOR_ITEM":
          setEditorItem(options);
          break;
        case "LOAD_FORMULA":
          setData("current", { modalForm: 
            <Formula 
              formula=""
              partnumber={data.edit.dataset.movement_head[0].partnumber}
              description={data.edit.dataset.movement_head[0].description}
              formulaValues={data.edit.dataset.formula_head.map(
                formula => { return { value: String(formula.id), text: formula.transnumber } }
              )}
              getText={app.getText}
              onClose={()=>setData("current", { modalForm: null })}
              onFormula={(formula_id) => {
                setData("current", { modalForm: null }, ()=>{
                  calcFormula(formula_id)
                })
              }}
            />,
            side: "hide"
          })
          break;
        case "NEW_FIELDVALUE":
          newFieldvalue(options.fieldname)
          break;
        case "CREATE_TRANS":
          createTrans(options)
          break;
        case "CREATE_TRANS_OPTIONS":
          createTransOptions()
          break;
        case "FORM_ACTIONS":
          setFormActions(options)
          break;
        default:
          break;
      }
    }
    if ((data.edit.dirty === true && data.edit.current.item) ||
      (data.edit.dirty === true && (data.edit.current.type==="template")) || 
      (data.edit.form_dirty === true && data.edit.current.form)) {
        setData("current", { modalForm: 
          <InputBox 
            title={app.getText("msg_warning")}
            message={app.getText("msg_dirty_text")}
            infoText={app.getText("msg_delete_info")}
            defaultOK={true}
            labelOK={app.getText("msg_save")}
            labelCancel={app.getText("msg_cancel")}
            onCancel={() => {
              setData("current", { modalForm: null }, ()=>{
                if (cbKeyFalse) {
                  setData("edit", { dirty: false, form_dirty: false }, ()=>{
                    cbNext(cbKeyFalse)
                  })
                } else {
                  cbNext(cbKeyTrue)
                }
              })
            }}
            onOK={(value) => {
              setData("current", { modalForm: null }, async ()=>{
                let edit = false
                if (data.edit.form_dirty) {
                  edit = await saveEditorForm()
                } else {
                  if (data.edit.current.type==="template"){
                    
                  } else {
                    edit = await saveEditor()
                  }
                }
                if(edit){
                  return setData("edit", edit, ()=>{
                    cbNext(cbKeyTrue)
                  })
                }
                return cbNext(cbKeyFalse)
              })
            }}
          />,
          side: "hide"
        })
    } else if (cbKeyFalse) {
      cbNext(cbKeyFalse);
    } else {
      cbNext(cbKeyTrue);
    }
  }
  
  const checkTranstype = async (options, cbKeyTrue, cbKeyFalse) => {
    if ((options.ntype==="trans" || options.ntype==="transitem" || 
      options.ntype==="transmovement" || options.ntype==="transpayment") && 
      options.ttype===null) {
        const params = { 
          method: "POST", 
          data: [{ 
            key: "transtype",
            text: getSql(data.login.data.engine, {
              select:["groupvalue"], from:"groups g", 
              inner_join:["trans t","on",[["g.id","=","t.transtype"],["and","t.id","=","?"]]]}).sql,
            values: [options.id] 
          }]
        }
        let view = await app.requestData("/view", params)
        if(view.error){
          return app.resultError(view)
        }
        checkEditor({ntype:"trans", ttype:view.transtype[0].groupvalue, id:options.id}, cbKeyTrue, cbKeyFalse)
    } else {
      checkEditor(options, cbKeyTrue, cbKeyFalse)
    }
  }

  const showStock = async (options) => {
    const params = { 
      method: "POST", 
      data: [{ 
        key: "stock",
        text: getSql(data.login.data.engine, sql.trans.shipping_stock()).sql,
        values: [options.product_id] 
      }]
    }
    let view = await app.requestData("/view", params)
    if(view.error){
      return app.resultError(view)
    }
    if (view.stock.length === 0){
      app.showToast({ type: "info",
        title: app.getText("msg_warning"), 
        message: app.getText("ms_no_stock") })
    }
    setData("current", { modalForm: 
      <Stock 
        partnumber={options.partnumber}
        partname={options.partname}
        rows={view.stock}
        getText={app.getText}
        onClose={()=>setData("current", { modalForm: null })}
      /> 
    })
  }
  
  const exportQueue = async (edit, item) => {
    const options = edit.current.item
    await reportOutput({
      type: options.mode, 
      template: item.reportkey, 
      title: item.refnumber,
      orient: options.orientation, 
      size: options.size, 
      copy: item.copies,
      nervatype: item.typename,
      id: item.ref_id
    })
    deleteEditorItem({
      fkey: "items", table: "ui_printqueue", id: item.id, prompt: true
    })
  }

  const setFormActions = (options, editData) => {

    const row = options.row || {}
    let edit = editData || data.edit
    switch (options.params.action) {
      case "loadEditor":
        checkEditor({
          ntype: options.params.ntype || edit.current.type, 
          ttype: options.params.ttype || edit.current.transtype, 
          id: row.id || null }, 
          'LOAD_EDITOR')
        break;
      
      case "newEditorItem":
        checkEditor({fkey: options.params.fkey, id: null}, 'SET_EDITOR_ITEM')
        break;
      
      case "editEditorItem":
        setEditorItem({fkey: options.params.fkey, id: row.id})
        break;
      
      case "deleteEditorItem":
        deleteEditorItem({fkey: options.params.fkey, table: options.params.table, id: row.id})
        break;

      case "loadShipping":
        checkEditor({
          ntype: options.params.ntype || edit.current.type, 
          ttype: options.params.ttype || edit.current.transtype, 
          id: options.params.id || edit.current.item.id, 
          shipping: true}, 'LOAD_EDITOR')
        break;
      
      case "addShippingRow":
        if (row.edited !== true) {
          edit = update(edit, {dataset: { shiptemp: {$push: [
            { 
              "id": row.item_id+"-"+row.product_id, 
              "item_id": row.item_id, "product_id": row.product_id,  
              "product": row.product, "partnumber": row.partnumber,
              "partname": row.partname, "unit": row.unit, 
              "batch_no":"", "qty":row.diff, "diff":0,
              "oqty":row.qty, "tqty":row.tqty
            }
          ]}}})
          setEditor({ shipping: true, form:"shipping_items" }, edit.template, edit)
        }
        break;
      
      case "showShippingStock":
        showStock({ 
          product_id: row.product_id, 
          partnumber: row.partnumber, 
          partname: row.partname
        })
        break;

      case "editShippingRow":
        setData("current", { modalForm: 
          <Shipping 
            partnumber={row.partnumber}
            description={row.product}
            unit={row.unit} batch_no={row.batch_no} qty={row.qty}
            getText={app.getText}
            onClose={()=>setData("current", { modalForm: null })}
            onShipping={(batch_no, qty) => {
              setData("current", { modalForm: null }, ()=>{
                const index = edit.dataset.shiptemp.findIndex(item => (item.id === row.id))
                edit = update(edit, { dataset: {shiptemp: { [index]: {$merge: {
                  batch_no: batch_no,
                  qty: qty,
                  diff: row.oqty - (row.tqty + qty)
                }}}}})
                setData("edit", edit)
              })
            }}
          /> 
        })
        break;
      
      case "deleteShippingRow":
        const index = edit.dataset.shiptemp.findIndex(item => (item.id === row.id))
        edit = update(edit, { dataset: {shiptemp: {
          $splice: [[index, 1]]
        }} })
        setEditor({shipping: true, form:"shiptemp_items"}, edit.template, edit)
        break;

      case "exportQueueItem":
        exportQueue(edit, row)
        break;
    
      default:
        break;
    }
  }

  const getTransFilter = (_sql, values) => {
    switch (data.login.data.transfilterName) {
      case "usergroup":
        _sql.where.push(
          ["and","cruser_id","in",[{
            select:["id"], from:"employee", 
            where:["usergroup","=","?"]
          }]])
        values.push(data.login.data.employee.usergroup)
        break;
      case "own":
        _sql.where.push(
          ["and","cruser_id","=","?"]
        )
        values.push(data.login.data.employee.id)
        break;
      default:
        break;
    }
    return [_sql, values]
  }

  const prevTransNumber = async () => {
    if (data.edit.current.type !== "trans" || data.edit.current.item.id === null) {
      return
    }
    const transtype = data.edit.current.transtype
    const direction = data.edit.dataset.groups.filter(
      item => (item.id === data.edit.current.item.direction))[0].groupvalue
    let _sql = {
      select:["max(id) as id"], from:"trans", where:[["transtype","=","?"]]
    }
    let values = [data.edit.current.item.transtype]
    if (transtype !== "cash" && transtype !== "waybill") {
      _sql.where.push(["and","direction","=","?"])
      values.push(data.edit.current.item.direction)
    }
    if (data.edit.current.item.id !== null) {
      _sql.where.push(["and","id","<","?"])
      values.push(data.edit.current.item.id)
    }
    if ((transtype === "invoice" && direction === "out") || 
      (transtype === "receipt" && direction === "out")|| (transtype === "cash")) {
      
    } else {
      _sql.where.push(["and","deleted","=","0"])
    }
    const filter = getTransFilter(_sql, values)
    const params = { 
      method: "POST", 
      data: [{ 
        key: "prev",
        text: getSql(data.login.data.engine, filter[0]).sql,
        values: filter[1]
      }]
    }
    let view = await app.requestData("/view", params)
    if(view.error){
      return app.resultError(view)
    }
    if (view.prev[0].id !== null){
      checkEditor({ntype: "trans", ttype: transtype, id: view.prev[0].id}, 'LOAD_EDITOR')
    }
  }

  const nextTransNumber = async () => {
    const transtype = data.edit.current.transtype
    const direction = data.edit.dataset.groups.filter(
      item => (item.id === data.edit.current.item.direction))[0].groupvalue
    let _sql = {
      select:["min(id) as id"], from:"trans", 
      where:[["transtype","=","?"],["and","id",">","?"]]
    }
    let values = [data.edit.current.item.transtype, data.edit.current.item.id]
    if (transtype !== "cash" && transtype !== "waybill") {
      _sql.where.push(["and","direction","=","?"])
      values.push(data.edit.current.item.direction)
    }
    if ((transtype === "invoice" && direction === "out") || 
      (transtype === "receipt" && direction === "out") || (transtype === "cash")) {} 
    else {
      _sql.where.push(["and","deleted","=","0"])
    }
    const filter = getTransFilter(_sql, values)
    const params = { 
      method: "POST", 
      data: [{ 
        key: "next",
        text: getSql(data.login.data.engine, filter[0]).sql,
        values: filter[1]
      }]
    }
    let view = await app.requestData("/view", params)
    if(view.error){
      return app.resultError(view)
    }
    if (view.next[0].id === null) {
      if (transtype==="delivery" && direction!=="transfer") {
        return
      } else {
        checkEditor({ntype: "trans", ttype: transtype, id: null}, 'LOAD_EDITOR')
      }
    } else {
      checkEditor({ntype: "trans", ttype: transtype, id: view.next[0].id}, 'LOAD_EDITOR')
    }
  }

  const createShipping = async () => {
    if (data.edit.current.shipping_place_id === null) {
      return app.showToast({ type: "error",
        title: app.getText("msg_warning"), 
        message: app.getText("msg_required")+" "+app.getText("inventory_warehouse") })
    }
    if (data.edit.dataset.shiptemp.length > 0){
      let delivery_head = update(initItem({tablename: "trans", 
          dataset: data.edit.dataset, current: data.edit.current}), {$merge: {
        transtype: data.edit.dataset.groups.filter(
          (item) => ((item.groupname === "transtype") && (item.groupvalue === "delivery"))
        )[0].id,
        direction: data.edit.dataset[data.edit.current.type][0].direction,
        transdate: formatISO(parseISO(data.edit.current.item.shippingdate), { representation: 'date' }),
        duedate: null, curr: null, paidtype: null
      }})
      let pattern = data.edit.dataset.delivery_pattern.filter(
          (item) => (item.defpattern === 1)
        )[0]
      if (typeof pattern !== "undefined") {
        delivery_head = update(delivery_head, {$merge: {
          fnote: pattern.notes
        }})
      }
      
      let params = { method: "POST", 
        data: {
          key: "nextNumber",
          values: {
            numberkey: "delivery_"+data.edit.current.direction, 
            step: true
          }
        }
      }
      let result = await app.requestData("/function", params)
      if(result.error){
        app.resultError(result)
        return null
      }
      delivery_head = update(delivery_head, {$merge: {
        transnumber: result
      }})

      result = await app.requestData("/trans", { method: "POST", data: [delivery_head] })
      if(result.error){
        app.resultError(result)
        return null
      }
      delivery_head.id = result[0];
      await app.createHistory("save")

      let movements = [];
      data.edit.dataset.shiptemp.forEach((shiptemp) => {
        let movement = update(
          initItem({tablename: "movement", dataset: data.edit.dataset, current: data.edit.current}),
          {$merge: {
            trans_id: delivery_head.id,
            shippingdate: formatISO(parseISO(data.edit.current.shippingdate)),
            product_id: shiptemp.product_id,
            place_id: data.edit.current.shipping_place_id,
            notes: shiptemp.batch_no,
            qty: (data.edit.current.direction === "out") ? -(shiptemp.qty) : shiptemp.qty
        }})
        movements.push(movement);
      });
      result = await app.requestData("/movement", { method: "POST", data: movements })
      if(result.error){
        app.resultError(result)
        return null
      }

      let links = [];
      let nervatype_movement = data.edit.dataset.groups.filter(
        (item)=>((item.groupname === "nervatype") && (item.groupvalue === "movement")))[0].id
      let nervatype_item = data.edit.dataset.groups.filter(
        (item)=>((item.groupname === "nervatype") && (item.groupvalue === "item")))[0].id
      result.forEach((movement_id, index) => {
        let link = update(initItem({tablename: "link", dataset: data.edit.dataset, current: data.edit.current}), {
          $merge: {
            nervatype_1: nervatype_movement,
            ref_id_1: movement_id,
            nervatype_2: nervatype_item,
            ref_id_2: data.edit.dataset.shiptemp[index].item_id
        }})
        links.push(link);
      });
      result = await app.requestData("/link", { method: "POST", data: links })
      if(result.error){
        app.resultError(result)
        return null
      }

      loadEditor({
        ntype: data.edit.current.type, 
        ttype: data.edit.current.transtype, 
        id: data.edit.current.item.id, 
        shipping: true
      })

    }
  }

  const exportEvent = () => {
    const event = data.edit.current.item;
    let eventStr = 
    `BEGIN:VCALENDAR\nPRODID:-//nervatura.com/NONSGML Nervatura Calendar//EN\nVERSION:2.0\nBEGIN:VEVENT\nUID:${
      (event.uid !== null)?event.uid:guid()}`
    if (event.fromdate !== null){
      eventStr+=`\nDTSTART:${format(parseISO(event.fromdate), "yyyyMMdd'T'HHmmss")}`
    }
    if (event.todate !== null){
      eventStr+=`\nDTEND:${format(parseISO(event.todate), "yyyyMMdd'T'HHmmss")}`
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
      let eventgroup = data.edit.dataset.eventgroup.filter(item => (item.id === event.eventgroup))[0]
      if (typeof eventgroup !== "undefined") {
        eventStr+=`\nCATEGORY:${eventgroup.groupvalue}`
      }
    }
    eventStr+=`\nEND:VEVENT\nEND:VCALENDAR`

    const filename = event.calnumber.replace(/\//g, "_")+".ics";
    let icsUrl = URL.createObjectURL(new Blob([eventStr], 
      {type : 'text/ics;charset=utf-8;'}));
    saveToDisk(icsUrl, filename);
  }

  const loadPrice = async (trans, item) => {
    const options = { method: "POST", 
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
    return app.requestData("/function", options)
  }

  const calcPrice = (_calcmode, item) => {
    
    let rate = data.edit.dataset.tax.filter(tax => (tax.id === parseInt(item.tax_id,10)))[0]
    rate = (typeof rate !== "undefined") ? rate.rate : 0
    let digit = data.edit.dataset.currency.filter(currency => 
      (currency.curr === data.edit.current.item.curr))[0]
    digit = (typeof digit !== "undefined") ? digit.digit : 2
    
    let netAmount = 0; let vatAmount = 0; let amount = 0; let fxPrice = 0;
    switch(_calcmode) {
      case "fxprice":
        fxPrice = parseFloat(item.fxprice)
        netAmount = round(fxPrice*(1-parseFloat(item.discount)/100)*parseFloat(item.qty),parseInt(digit,10))
        vatAmount = round(fxPrice*(1-parseFloat(item.discount)/100)*parseFloat(item.qty)*parseFloat(rate),parseInt(digit,10))
        amount = round(netAmount+vatAmount, parseInt(digit,10))
        break;
        
      case "netamount":
        netAmount = parseFloat(item.netamount)
        if (parseFloat(item.qty)!==0) {
          fxPrice = round(netAmount/(1-parseFloat(item.discount)/100)/parseFloat(item.qty),parseInt(digit,10))
          vatAmount = round(netAmount*parseFloat(rate),parseInt(digit,10))
        }
        amount = round(netAmount+vatAmount,parseInt(digit,10))
        break;

      case "amount":
        amount = parseFloat(item.amount)
        if (parseFloat(item.qty)!==0) {
          netAmount = round(amount/(1+parseFloat(rate)),parseInt(digit,10))
          vatAmount = round(amount-netAmount,parseInt(digit,10))
          fxPrice = round(netAmount/(1-parseFloat(item.discount)/100)/parseFloat(item.qty),parseInt(digit,10))
        }
        break;
      default:
    }
    return update(item, {$merge: {
      fxprice: fxPrice,
      netamount: netAmount,
      vatamount: vatAmount,
      amount: amount
    }})
  }

  const editItem = async (options, editData) => {
    let edit = update({}, {$set: editData})
    if((options.name === "fieldvalue_value") || (options.name === "fieldvalue_notes") || (options.name === "fieldvalue_deleted")){
      const fieldvalue_idx = edit.current.fieldvalue.findIndex((item)=>(item.id === options.id))
      if( (fieldvalue_idx > -1) && ((edit.audit==="all") || (edit.audit==="update"))){
        edit = update(edit, {$merge: {
          dirty: true,
        }})
        edit = update(edit, { current: { fieldvalue: { [fieldvalue_idx]: {$merge: {
          [options.name.split("_")[1]]: (options.name === "fieldvalue_deleted") ? 1 : options.value.toString()
        }}}}})
      }
    } else if (edit.current.form) {
      edit = update(edit, {$merge: {
        form_dirty: true
      }})
      if (typeof edit.current.form[options.name] !== "undefined") {
        edit = update(edit, {current: {form: {$merge: {
          [options.name]: options.value
        }}}})
      }
      switch (edit.current.form_type) {
        case "item":
          if (options.name === "product_id" && (typeof options.item !== "undefined")) {
            edit = update(edit, {current: {form: {$merge: {
              description: options.item.description,
              unit: options.item.unit,
              tax_id: parseInt(options.item.tax_id,10)
            }}}})
            if (edit.current.form.qty === 0) {
              edit = update(edit, {current: {form: {$merge: {
                qty: 1
              }}}})
            }
            const price = await loadPrice(edit.current.item, edit.current.form)
            if(price.error){
              return app.resultError(price)
            }
            edit = update(edit, {current: {form: {$merge: {
              fxprice: !isNaN(parseFloat(price.price)) ? parseFloat(price.price) : 0,
              discount: !isNaN(parseFloat(price.discount)) ? parseFloat(price.discount) : 0
            }}}})
            if(options.event_type === "blur"){
              edit = update(edit, {current: {$merge: {
                form : calcPrice("fxprice", edit.current.form)
              }}})
            }
          } else {
            switch(options.name) {
              case "qty":
                if (parseFloat(edit.current.form.fxprice) === 0) {
                  const price = await loadPrice(edit.current.item, edit.current.form)
                  if(price.error){
                    return app.resultError(price)
                  }
                  edit = update(edit, {current: {form: {$merge: {
                    fxprice: !isNaN(parseFloat(price.price)) ? parseFloat(price.price) : 0,
                    discount: !isNaN(parseFloat(price.discount)) ? parseFloat(price.discount) : 0
                  }}}})
                }
                if(options.event_type === "blur"){
                  edit = update(edit, {current: {$merge: {
                    form : calcPrice("fxprice", edit.current.form)
                  }}})
                }
                break;
              case "fxprice":
              case "tax_id":
              case "discount":
                if(options.event_type === "blur"){
                  edit = update(edit, {current: {$merge: {
                    form : calcPrice("fxprice", edit.current.form)
                  }}})
                }
                break;
              case "amount":
                if(options.event_type === "blur"){
                  edit = update(edit, {current: {$merge: {
                    form : calcPrice("amount", edit.current.form)
                  }}})
                }
                break;
              case "netamount":
                if(options.event_type === "blur"){
                  edit = update(edit, {current: {$merge: {
                    form : calcPrice("netamount", edit.current.form)
                  }}})
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
            edit = update(edit, {current: {$merge: {
              price_customer_id: options.value
            }}})
          }
          break;

        case "invoice_link":
          if (options.name === "ref_id_1" && (typeof options.item !== "undefined")) {
            edit = update(edit, {current: {invoice_link: {
              0: {$merge: {
                curr: options.item.curr
              }}
            }}})
          } else if ((options.name === "link_qty") || (options.name === "link_rate")) {
            edit = update(edit, { current: {$merge: {
              invoice_link_fieldvalue: setFieldvalue(edit.current.invoice_link_fieldvalue, 
                options.name, edit.current.form.id, null, options.value)
            }}})
          }
          break;

        case "payment_link":
          if (options.name === "ref_id_2" && (typeof options.item !== "undefined")) {
            edit = update(edit, {current: {payment_link: {
              0: {$merge: {
                curr: options.item.curr
              }}
            }}})
          } else if ((options.name === "link_qty") || (options.name === "link_rate")) {
            edit = update(edit, { current: {$merge: {
              payment_link_fieldvalue: setFieldvalue(edit.current.payment_link_fieldvalue, 
                options.name, edit.current.form.id, null, options.value)
            }}})
          }
          break;

        default:
          break;
      }
    } else {
      if ((typeof edit.current.item[options.name] !== "undefined") && (options.extend === false)) {
        edit = update(edit, {current: {item: {$merge: {
          [options.name]: options.value
        }}}})
        if(options.label_field){
          edit = update(edit, {current: {item: {$merge: {
            [options.label_field]: options.refnumber || null
          }}}})
        }
      } else if ((typeof edit.template.options.extend !== "undefined") && (options.extend === true)) {
        edit = update(edit, {current: {extend: {$merge: {
          [options.name]: options.value
        }}}})
      }
      if((edit.audit==="all") || (edit.audit==="update")){
        edit = update(edit, {$merge: {
          dirty: true
        }})
      }

      switch (edit.current.type){
        case "report":
          edit = update(edit, {$merge: {
            dirty: false
          }})
          if(options.name === "selected"){
            const fieldvalue_idx = edit.current.fieldvalue.findIndex((item)=>(item.id === options.id))
            if(fieldvalue_idx > -1){
              edit = update(edit, { current: { fieldvalue: { [fieldvalue_idx]: {$merge: {
                selected: options.value
              }}}}})
            }
          } else {
            const fieldvalue_idx = edit.current.fieldvalue.findIndex((item)=>(item.name === options.name))
            if(fieldvalue_idx > -1){
              edit = update(edit, { current: { fieldvalue: { [fieldvalue_idx]: {$merge: {
                value: options.value
              }}}}})
            }
          }
          break;

        case "printqueue":
          edit = update(edit, {$merge: {
            dirty: false
          }})
          edit = update(edit, { printqueue: {$merge: {
              [options.name]: options.value
          }}})
          break;

        case "trans":
          switch (options.name) {
            case "closed":
              if (options.value === 1) {
                setData("current", { modalForm: 
                  <InputBox 
                    title={app.getText("msg_warning")}
                    message={app.getText("msg_close_text")}
                    infoText={app.getText("msg_delete_info")}
                    labelOK={app.getText("msg_ok")}
                    labelCancel={app.getText("msg_cancel")}
                    onCancel={() => {
                      setData("current", { modalForm: null }, ()=>{
                        edit = update(edit, { current: { item: {$merge: {
                          [options.name]: 0
                        }}}})
                        setData("edit", edit)
                      })
                    }}
                    onOK={(value) => {
                      setData("current", { modalForm: null }, ()=>{
                        edit = update(edit, { current: {$merge: {
                          closed: 1
                        }}})
                        setData("edit", edit)
                      })
                    }}
                  /> 
                })
              }
              break;
            case "paiddate":
              edit = update(edit, { current: { item: {$merge: {
                transdate: options.value
              }}}})
              break;
            case "direction":
              if(edit.current.transtype === "cash"){
                const direction = edit.dataset.groups.filter((item)=>(item.id === options.value))[0].groupvalue
                edit = update(edit, { template: { options: {$merge: {
                  opposite: (direction === "out")
                }}}})
              }
              break;        
            case "seltype":
              edit = update(edit, { current: { 
                extend: {$merge: {
                  seltype: options.value,
                  ref_id: null,
                  refnumber: ""
                }},
                item: {$merge: {
                  customer_id: null,
                  employee_id: null,
                  ref_transnumber: null
                }}
              }})
              break;
            case "ref_id":
              edit = update(edit, { current: { extend: {$merge: {
                  refnumber: options.refnumber,
                  ntype: (edit.current.extend.seltype === "transitem") ? "trans" : edit.current.extend.seltype,
                  transtype: (options.item && options.item.transtype) ? options.item.transtype.split("-")[0] : ""
              }}}})
              switch (edit.current.extend.seltype){
                case "customer":
                  edit = update(edit, { current: { item: {$merge: {
                    customer_id: options.value
                  }}}})
                  break;
                case "employee":
                  edit = update(edit, { current: { item: {$merge: {
                    employee_id: options.value
                  }}}})
                  break;
                case "transitem":
                  edit = update(edit, { current: { item: {$merge: {
                    ref_transnumber: options.refnumber,
                  }}}})
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
              edit = update(edit, { current: {$merge: {
                fieldvalue: setFieldvalue(edit.current.fieldvalue, 
                  options.name, edit.current.item.id, null, options.value)
              }}})
              break;
            case "shippingdate":
            case "shipping_place_id":
              edit = update(edit, {$merge: {
                dirty: false
              }})
              edit = update(edit, { current: {$merge: {
                  [options.name]: options.value
              }}})
              break;
            case "fnote":
              edit = update(edit, { current: { item: {$merge: {
                fnote: options.value
              }}}})
              break;
            default:
              break;
          }
          break;
        default:
          break;
      }
    }
    setData("edit", edit)
  }

  const setPattern = ( options, editData ) => {
    const { key } = options
    const updatePattern = async (values) => {
      const options = { method: "POST", data: values }
      let result = await app.requestData("/pattern", options)
      if(result.error){
        return app.resultError(result)
      }
      checkEditor({ntype: editData.current.type, 
        ttype: editData.current.transtype, 
        id: editData.current.item.id, form:"fnote"}, 'LOAD_EDITOR')
    }
    const patternBox = {
      default: {
        title: app.getText("msg_warning"),
        message: app.getText("msg_pattern_default"),
        infoText: undefined,
        value: "",
        showValue: false,
        defaultOK: true,
        ok: (value) => {
          let pattern = update(editData.dataset.pattern, {})
          pattern.forEach((element, index) => {
            pattern = update(pattern, {
              [index]: { $merge: {
                defpattern: (element.id === parseInt(editData.current.template,10)) ? 1 : 0
              }}
            })
          });
          updatePattern(pattern)
        }
      },
      save: {
        title: app.getText("msg_warning"),
        message: app.getText("msg_pattern_save"),
        infoText: undefined,
        value: "",
        showValue: false,
        defaultOK: true,
        ok: (value) => {
          let pattern = editData.dataset.pattern.filter((item) => 
            (item.id === parseInt(editData.current.template,10) ))[0]
          if(pattern){
            pattern = update(pattern, {$merge: {
              notes: editData.current.item.fnote
            }})
            updatePattern([pattern])
          }
        }
      },
      new: {
        title: app.getText("msg_pattern_new"),
        message: app.getText("msg_pattern_name"),
        infoText: undefined,
        value:"", 
        showValue: true,
        defaultOK: false,
        ok: async (value) => {
          if(value !== ""){
            let result = await app.requestData("/pattern", {
              query: {
                filter: "description;==;"+value
              }
            })
            if(result.error){
              return app.resultError(result)
            }
            if(result.length > 0){
              return app.showToast({ type: "error", title: app.getText("msg_warning"), 
                message: app.getText("msg_value_exists") })
            }
            const pattern = update(initItem({tablename: "pattern", current: editData.current}), {$merge: {
              description: value,
              defpattern: (editData.dataset.pattern.length === 0) ? 1 : 0
            }})
            updatePattern([pattern])
          }
        }
      },
      delete: {
        title: app.getText("msg_warning"),
        message: app.getText("msg_delete_text"),
        infoText: app.getText("msg_delete_info"),
        value: "",
        showValue: false,
        defaultOK: false,
        ok: (value) => {
          let pattern = editData.dataset.pattern.filter((item) => 
            (item.id === parseInt(editData.current.template,10) ))[0]
          if(pattern){
            pattern = update(pattern, {$merge: {
              deleted: 1,
              defpattern: 0
            }})
            updatePattern([pattern])
          }
        }
      }
    }
    if(key !== "new"){
      if(!editData.current.template || (editData.current.template === "")){
        return app.showToast({ type: "error", title: app.getText("msg_warning"), 
          message: app.getText("msg_pattern_missing") })
      }
    }
    if(key === "load"){
      const pattern = editData.dataset.pattern.filter((item) => 
        (item.id === parseInt(editData.current.template,10) ))[0]
      if(pattern){
        let edit = update(editData, {
          current: { item: { $merge: {
            fnote: pattern.notes
          }}}
        })
        edit = update(edit, {$merge: {
          dirty: true,
          lastUpdate: new Date().getTime()
        }})
        setData("edit", edit)
      }
    } else{
      setData("current", { modalForm: 
        <InputBox 
          title={patternBox[key].title}
          message={patternBox[key].message}
          infoText={patternBox[key].infoText}
          value={patternBox[key].value}
          showValue={patternBox[key].showValue}
          defaultOK={patternBox[key].defaultOK}
          labelOK={app.getText("msg_ok")}
          labelCancel={app.getText("msg_cancel")}
          onCancel={() => {
            setData("current", { modalForm: null })
          }}
          onOK={(value) => {
            setData("current", { modalForm: null }, ()=>{
              patternBox[key].ok(value)
            })
          }}
        /> 
      })
    }
  }

  return {
    round: round,
    createReport: createReport,
    exportQueueAll: exportQueueAll,
    searchQueue: searchQueue,
    reportOutput: reportOutput,
    reportSettings: reportSettings,
    checkEditor: checkEditor,
    checkTranstype: checkTranstype,
    loadEditor: loadEditor,
    setEditor: setEditor,
    setFormActions: setFormActions,
    deleteEditorItem: deleteEditorItem,
    deleteEditor: deleteEditor,
    prevTransNumber: prevTransNumber,
    nextTransNumber: nextTransNumber,
    setFieldvalue: setFieldvalue,
    saveEditorForm: saveEditorForm,
    saveEditor: saveEditor,
    createShipping: createShipping,
    exportEvent: exportEvent,
    editItem: editItem,
    setPattern: setPattern
  }
}