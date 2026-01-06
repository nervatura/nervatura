import { APP_MODULE, SIDE_EVENT, TEMPLATE_EVENT, SIDE_VISIBILITY, MODAL_EVENT, TOAST_TYPE, APP_THEME } from '../config/enums.js'
import { templateElements } from '../components/Template/Editor/Elements.js'

const getDataset = (data) => {
  let dataset = [];
  Object.keys(data).forEach((dskey) => {
    if (typeof data[dskey] === "string"){
      dataset = [...dataset, { lslabel: dskey, lsvalue: "string" }]
    } else if (Array.isArray(data[dskey])){
        dataset = [...dataset, { lslabel: dskey, lsvalue: "table" }] 
    } else {
      dataset = [...dataset, { lslabel: dskey, lsvalue: "list" }]
    }
  });
  return dataset;
}

export const getElementType = (element) => {
  if (Object.getOwnPropertyNames(element).length>0) {
    return Object.getOwnPropertyNames(element)[0];
  } 
  return null;
}

export const getDataList = (data)=>{
  let datalist = [];
  Object.keys(data).forEach((key) => {
    datalist = [...datalist, { lslabel: key, lsvalue: data[key] }]
  });
  return datalist;
}

export const getDataTable = (data)=>{
  const table = { fields: {}, items: []};
  if(data.length > 0){
    Object.keys(data[0]).forEach((key) => {
      table.fields[key] = { fieldtype: 'string', label: key }
    })
    for (let index = 0; index < data.length; index += 1) {
      const item = {...data[index],
        _index: index
      }
      table.items = [...table.items, item] 
    } 
  }
  return table;
}

export class TemplateController {
  constructor(host) {
    this.host = host
    this.app = host.app
    this.store = host.app.store
    this.module = {}

    this.addItem = this.addItem.bind(this)
    this.addTemplateData = this.addTemplateData.bind(this)
    this.checkTemplate = this.checkTemplate.bind(this)
    this.createMap = this.createMap.bind(this)
    this.createTemplate = this.createTemplate.bind(this)
    this.deleteData = this.deleteData.bind(this)
    this.deleteDataItem = this.deleteDataItem.bind(this)
    this.deleteItem = this.deleteItem.bind(this)
    this.deleteTemplate = this.deleteTemplate.bind(this)
    this.editDataItem = this.editDataItem.bind(this)
    this.editItem = this.editItem.bind(this)
    this.exportTemplate = this.exportTemplate.bind(this)
    this.goNext = this.goNext.bind(this)
    this.goPrevious = this.goPrevious.bind(this)
    this.moveDown = this.moveDown.bind(this)
    this.moveUp = this.moveUp.bind(this)
    this.onTemplateEvent = this.onTemplateEvent.bind(this)
    this.onSideEvent = this.onSideEvent.bind(this)
    this.saveTemplate = this.saveTemplate.bind(this)
    this.setCurrent = this.setCurrent.bind(this)
    this.setCurrentData = this.setCurrentData.bind(this)
    this.setCurrentDataItem = this.setCurrentDataItem.bind(this)
    this.setTemplate = this.setTemplate.bind(this)
    this.showPreview = this.showPreview.bind(this)
    this.setModule = this.setModule.bind(this)
    host.addController(this);
  }

  setModule(moduleRef){
    this.module = moduleRef
  }

  addItem(value) {
    const { data } = this.store
    const { current } = data[APP_MODULE.TEMPLATE]
    if (value !== "") {
      const ename = value.toString().toLowerCase();
      const element = {}; element[ename] = {};
      const id = `${current.id}_${current.item.length.toString()}_${ename}`;
      if (ename==="datagrid" || ename==="row") {
        element[ename].columns=[];
      }
      current.item.push(element)
      this.setCurrent({ tmp_id: id, set_dirty: true})
    }
  }

  addTemplateData() {
    const { modalTemplate } = this.module
    const { setData } = this.store
    const menuForm = modalTemplate({
      type: "string", name: "", columns: "",
      onEvent: {
        onModalEvent: async (modalResult) => {
          setData("current", { modalForm: null })
          if(modalResult.key === MODAL_EVENT.OK){
            this.setCurrentData({ name: "new", type: "new", values: {...modalResult.data.value} })
          }
        }
      }
    })
    setData("current", { modalForm: menuForm })
  }

  checkTemplate(cbKeyTrue) {
    const { inputBox } = this.host
    const { msg, currentModule } = this.app
    const { setData, data } = this.store
    const nextKeys = {
      [SIDE_EVENT.BLANK]: () => this.setTemplate({ type: "_blank" }),
      [SIDE_EVENT.SAMPLE]: () => this.setTemplate({ type: "_sample" }),
    }
    if (data[APP_MODULE.TEMPLATE].dirty === true) {
        const  modalForm = inputBox({ 
          title: msg("", { id: "msg_warning" }),
          message: msg("", { id: "msg_dirty_text" }),
          infoText: msg("", { id: "msg_delete_info" }),
          defaultOK: true,
          onEvent: {
            onModalEvent: async (modalResult) => {
              setData("current", { modalForm: null })
              if (modalResult.key === MODAL_EVENT.OK) {
                const setting = await this.saveTemplate()
                if(setting){
                  setData(APP_MODULE.TEMPLATE, setting)
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

  createMap(_cv) {
    const { data, setData } = this.store
    const { template, current, mapRef } = data[APP_MODULE.TEMPLATE]
    const cv = _cv || mapRef
    const cont = cv.getContext('2d')
    cont.clearRect(0, 0, cv.width, cv.height)
    const cell_color = "#CCCCCC"; const row_color = "#FFFF00"; const sel_color = "#00EE00";
    const cell_size = 8; const cell_pad = 1; const page_pad = 3; const rows = [];
    const def_height = 165;

    cv.height = page_pad;
    const sections = ["header","details","footer"];
    for(let s = 0; s < sections.length; s += 1) {
      for(let i = 0; i < template[sections[s]].length; i += 1) {
        const row = {};
        row.type = getElementType(template[sections[s]][i]);
        const item = template[sections[s]][i][row.type];
        if (row.type==="row" || row.type==="datagrid") {
          row.cols = template[sections[s]][i][row.type].columns.length;
          row.selected = (item.columns===current.item || item.columns===current.parent);
          row.selcol = -1;
          if (row.selected) {
            for(let c = 0; c < row.cols; c += 1) {
              const cname = getElementType(template[sections[s]][i][row.type].columns[c]);
              if (template[sections[s]][i][row.type].columns[c][cname]===current.item) {
                row.selcol = c;
              }
            }
          }
          if (row.cols*(cell_size+cell_pad)+2*page_pad > cv.width)
            cv.width = row.cols*(cell_size+cell_pad)+2*page_pad;
        } else {
          row.selected = (item===current.item); row.cols = 1;
        }
        switch (row.type) {
          case "vgap":
            if (template[sections[s]][i][row.type].height>2) {
              row.height = template[sections[s]][i][row.type].height;
            } else {
              row.height = 2;
            }
            cv.height += row.height;
            break;
          case "hline":
            cv.height += 2;
            break;
          case "datagrid":
            cv.height += 2*cell_size + 4*cell_pad;
            break;
          default:
            cv.height += cell_size + cell_pad;
            break;}
        if (template[sections[s]]===current.item || template.report===current.item) {
          row.selected = true;
        }
        rows.push(row);
      }
    }
    cv.height += page_pad;
    if (cv.height<def_height) {
      cv.height = def_height;
    }

    let x = page_pad; let y = page_pad; let coldif = 0;
    for(let r = 0; r < rows.length; r += 1) {
      if(rows[r].type === "row"){
        coldif = (cv.width - (rows[r].cols*(cell_size+cell_pad)+2*page_pad))/rows[r].cols;
        for(let cr = 0; cr < rows[r].cols; cr += 1) {
          if (rows[r].selected) {
            if (rows[r].selcol === cr || rows[r].selcol === -1) {
              cont.fillStyle = sel_color;
            } else {
              cont.fillStyle = row_color;
            }
          } else {
            cont.fillStyle = cell_color;
          }
          cont.fillRect(x, y, cell_size+coldif, cell_size);
          x += cell_size + coldif + cell_pad;
        }
        y += cell_size + cell_pad;
      }
      if(rows[r].type === "datagrid"){
        if (rows[r].selected) {
          cont.fillStyle = sel_color;
        } else {
          cont.fillStyle = cell_color;
        }
        cont.fillRect(x, y, cv.width-2*page_pad-cell_pad, cell_size/2);
        coldif = (cv.width - (rows[r].cols*(cell_size+cell_pad)+2*page_pad))/rows[r].cols;
        for(let cc = 0; cc < rows[r].cols; cc += 1) {
          if (rows[r].selected) {
            if (rows[r].selcol === cc || rows[r].selcol === -1) {
              cont.fillStyle = sel_color;
            } else {
              cont.fillStyle = row_color;
            }
          } else {
            cont.fillStyle = cell_color;
          }
          cont.fillRect(x, y+cell_size/2+cell_pad, cell_size+coldif, cell_size/2);
          cont.fillRect(x, y+cell_size+2*cell_pad, cell_size+coldif, cell_size/2);
          cont.fillRect(x, y+1.5*cell_size+3*cell_pad, cell_size+coldif, cell_size/2);
          x += cell_size + coldif + cell_pad;
        }
        y += 2*cell_size + 4*cell_pad;
      }
      if(rows[r].type === "vgap"){
        if (rows[r].selected) {
          cont.fillStyle = sel_color;
          cont.fillRect(x, y-cell_pad, cv.width-2*page_pad-cell_pad, rows[r].height);
        }
        y += rows[r].height;
      }
      if(rows[r].type === "vgap"){
        if (rows[r].selected) {
          cont.strokeStyle = sel_color;
        } else {
          cont.strokeStyle = cell_color;
        }
        cont.beginPath();
        cont.moveTo(x, y);
        cont.lineTo(cv.width-page_pad-cell_pad, y);
        cont.stroke();
        y += 2;
      }
      if(rows[r].type === "html"){
        if (rows[r].selected) {
          cont.fillStyle = sel_color;
        } else {
          cont.fillStyle = cell_color;
        }
        cont.fillRect(x, y, cv.width-2*page_pad-cell_pad, cell_size);
        y += cell_size + cell_pad;
      }
      x = page_pad;
    }
    if(_cv){
      setData(APP_MODULE.TEMPLATE, { mapRef: cv })
    }
  }

  createTemplate() {
    const { inputBox } = this.host
    const { msg, currentModule, requestData, resultError } = this.app
    const { data, setData } = this.store
    let report_key = String(data[APP_MODULE.TEMPLATE].template.meta.report_type).toLowerCase();
    if (report_key === "trans") {
      report_key = `${String(data[APP_MODULE.TEMPLATE].template.meta.trans_type).toLowerCase()}_${
        String(data[APP_MODULE.TEMPLATE].template.meta.direction).toLowerCase().replace("direction_","")
      }`;
    }
    const timeValue = new Date()
    report_key += `_${new Date(timeValue).toISOString().slice(0,10)}${new Date(timeValue).toLocaleTimeString("en",{hour12: false}).replace("24","00")}`.replaceAll("-","").replaceAll(":","")
    const  modalForm = inputBox({ 
      title: msg("", { id: "template_label_new" }),
      message: report_key,
      value: data[APP_MODULE.TEMPLATE].dbtemp.data.report_name, showValue: true,
      onEvent: {
        onModalEvent: async (modalResult) => {
          setData("current", { modalForm: null })
          if ((modalResult.key === MODAL_EVENT.OK) && (modalResult.data.value !== "")) {
            const template = {...data[APP_MODULE.TEMPLATE].template,
              meta: {...data[APP_MODULE.TEMPLATE].template.meta,
                report_key,
                report_name: modalResult.data.value
              }
            }
            const dbtemp = {...data[APP_MODULE.TEMPLATE].dbtemp,
              id: null,
              code: report_key,
              data: {...data[APP_MODULE.TEMPLATE].dbtemp.data,
                report_name: modalResult.data.value,
                template: JSON.stringify(template)
              }
            }

            const result = await requestData("/config", { method: "POST", data: dbtemp })
            if(result.error){
              return resultError(result)
            }

            dbtemp.id = result.id
            return currentModule({ 
              data: { module: APP_MODULE.TEMPLATE },
              content: { fkey: "setTemplate", args: [ { type: "template", report: dbtemp } ] } 
            })
          }
          return true
        }
      }
    })
    return setData("current", { modalForm, side: SIDE_VISIBILITY.HIDE })
  }

  deleteData(dskey) {
    const { inputBox } = this.host
    const { msg } = this.app
    const { setData, data } = this.store
    const  modalForm = inputBox({ 
      title: msg("", { id: "msg_warning" }),
      message: msg("", { id: "msg_delete_text" }),
      infoText: msg("", { id: "msg_delete_info" }),
      defaultOK: true,
      onEvent: {
        onModalEvent: async (modalResult) => {
          setData("current", { modalForm: null })
          if (modalResult.key === MODAL_EVENT.OK) {
            const {[dskey]:unset, ...rest_data} = data[APP_MODULE.TEMPLATE].template.data
            let setting = {...data[APP_MODULE.TEMPLATE],
              template: {...data[APP_MODULE.TEMPLATE].template,
                data: rest_data
              }
            }
            if(!["_blank","_sample"].includes(setting.key)){
              setting = {...setting,
                dirty: true
              }
            }
            setting = {...setting,
              dataset: [...getDataset(setting.template.data)]
            }
            setData(APP_MODULE.TEMPLATE, setting)
          }
        }
      }
    })
    return setData("current", { modalForm, side: SIDE_VISIBILITY.HIDE })
  }

  deleteDataItem(options) {
    const { inputBox } = this.host
    const { msg } = this.app
    const { setData, data } = this.store
    const  modalForm = inputBox({ 
      title: msg("", { id: "msg_warning" }),
      message: msg("", { id: "msg_delete_text" }),
      infoText: msg("", { id: "msg_delete_info" }),
      defaultOK: true,
      onEvent: {
        onModalEvent: async (modalResult) => {
          setData("current", { modalForm: null })
          if (modalResult.key === MODAL_EVENT.OK) {

            let setting = {...data[APP_MODULE.TEMPLATE]}
  
            switch (setting.current_data.type) {
              case "list":
                const {[options.key]:unset, ...rest_list} = setting.template.data[setting.current_data.name]
                setting = {...setting,
                  template: {...setting.template,
                    data: {...setting.template.data,
                      [setting.current_data.name]: rest_list
                    }
                  }
                }
                setting = {...setting,
                  current_data: {...setting.current_data,
                    items: [...getDataList(setting.template.data[setting.current_data.name])]
                  }
                }
                break;
    
              case "table":
              default:
                if (setting.template.data[setting.current_data.name].length===1){
                  const {[setting.current_data.name]:unset_tbl, ...rest_table} = setting.template.data
                  setting = {...setting,
                    template: {...setting.template,
                      data: rest_table
                    }
                  }
                  setting = {...setting,
                    current_data: null,
                    dataset: [...getDataset(setting.template.data)]
                  }

                } else {
                  setting.template.data[setting.current_data.name] =  [
                    ...setting.template.data[setting.current_data.name].slice(0,options._index),
                    ...setting.template.data[setting.current_data.name].slice(options._index+1)
                  ]
                  setting = {...setting,
                    current_data: {...setting.current_data,
                      items: [...getDataTable(setting.template.data[setting.current_data.name]).items]
                    }
                  }
                }
                break;
            }
            setting = {...setting,
              dataset: [...getDataset(setting.template.data)]
            }
            if(!["_blank","_sample"].includes(setting.key)){
              setting = {...setting,
                dirty: true
              }
            }
            setData(APP_MODULE.TEMPLATE, setting)
          }
        }
      }
    })
    return setData("current", { modalForm, side: SIDE_VISIBILITY.HIDE })
  }

  deleteItem() {
    const { data } = this.store
    const { current } = data[APP_MODULE.TEMPLATE]
    if (current.parent!==null && current.index!==null) {
      current.parent.splice(current.index, 1);
      let id = `tmp_${current.section}`;
      if (current.parent_index!==null) {
        id += `_${current.parent_index.toString()}_${current.parent_type}`;
      }
      this.setCurrent({ tmp_id: id, set_dirty: true})
    }
  }

  async deleteTemplate() {
    const { inputBox } = this.host
    const { msg, requestData, resultError, currentModule } = this.app
    const { setData, data } = this.store
    const  modalForm = inputBox({ 
      title: msg("", { id: "msg_warning" }),
      message: msg("", { id: "msg_delete_text" }),
      infoText: msg("", { id: "msg_delete_info" }),
      onEvent: {
        onModalEvent: async (modalResult) => {
          setData("current", { modalForm: null })
          if (modalResult.key === MODAL_EVENT.OK) {
            const result = await requestData("/config/"+data[APP_MODULE.TEMPLATE].dbtemp.code, 
              { method: "DELETE" })
            if(result && result.error){
              return resultError(result)
            }
            return currentModule({ 
              data: { module: APP_MODULE.TEMPLATE },
              content: { fkey: "setTemplate", args: [ { type: "_sample" } ] } 
            })
          }
          return true
        }
      }
    })
    return setData("current", { modalForm, side: SIDE_VISIBILITY.HIDE })
  }

  editDataItem(options) {
    const { data, setData } = this.store
    let setting = {...data[APP_MODULE.TEMPLATE]}
    if(!["_blank","_sample"].includes(setting.key)){
      setting = {...setting,
        dirty: true
      }
    }
    switch (setting.current_data.type) {
      case "string":
        setting = {...setting,
          template: {...setting.template,
            data: {...setting.template.data,
              [setting.current_data.name]: options.value
            }
          }
        }
        break;
      
      case "list":
        setting = {...setting,
          template: {...setting.template,
            data: {...setting.template.data,
              [setting.current_data.name]: {...setting.template.data[setting.current_data.name],
                [setting.current_data.item]: options.value
              }
            }
          }
        }
        setting = {...setting,
          current_data: {...setting.current_data,
            items: [...getDataList(setting.template.data[setting.current_data.name])]
          }
        }
        break;

      default:
        setting = {...setting,
          current_data: {...setting.current_data,
            item: {...setting.current_data.item,
              [options.field]: options.value
            }
          },
          template: {...setting.template,
            data: {...setting.template.data,
              [setting.current_data.name]: [...setting.template.data[setting.current_data.name]]
            }
          }
        }
        setting.template.data[setting.current_data.name][options._index] = {
          ...setting.template.data[setting.current_data.name][options._index],
          [options.field]: options.value
        }
        setting = {...setting,
          current_data: {...setting.current_data,
            items: [...getDataTable(setting.template.data[setting.current_data.name]).items]
          }
        }
        break;
    }
    setData(APP_MODULE.TEMPLATE, setting)
  }

  editItem(options) {
    const { data } = this.store
    let setting = {...data[APP_MODULE.TEMPLATE]}
    const id = setting.current.id
    const itemId = id.split("_");

    const setItemBase = (key, value) => {
      if(value === null){
        const {[key]:unset, ...item_base} = setting.current.item_base
        setting = {...setting,
          current: {...setting.current,
            item_base
          }
        }
      } else {
        setting = {...setting,
          current: {...setting.current,
            item_base: {...setting.current.item_base,
              [key]: value
            }
          }
        }
      }
      setting.template[setting.current.section][parseInt(itemId[2],10)][itemId[3]] = {
        ...setting.current.item_base,
        columns: [...setting.current.item_base.columns]
      }
    }

    const setItem = (key, value) => {
      if(value === null){
        const {[key]:unset, ...item} = setting.current.item
        setting = {...setting,
          current: {...setting.current,
            item
          }
        }
      } else {
        setting = {...setting,
          current: {...setting.current,
            item: {...setting.current.item,
              [key]: value
            }
          }
        }
      }
      switch (itemId.length) {
        case 2:
          setting.template[setting.current.section] = {...setting.current.item}
          break;

        case 4:
          setting.template[setting.current.section][parseInt(itemId[2],10)][itemId[3]] = {...setting.current.item}
          break;

        case 6:
        default:
          setting.template[setting.current.section][parseInt(itemId[2],10)][itemId[3]].columns[parseInt(itemId[4],10)][itemId[5]] = {...setting.current.item}
          break;
      }
    }

    if(!["_blank","_sample"].includes(setting.key)){
      setting = {...setting,
        dirty: true
      }
    }

    if(options.selected){
      let value = ""
      if(options.value){
        if(options.defvalue){
          value = options.defvalue
        } else {
          switch (options.datatype) {
            case "float":
            case "integer":
              value = 0
              break;
            default:
              value = "";
              break;
          }
        }
        if(setting.current.item_base){
          setItemBase(options.name, value)
        } else {
          setItem(options.name, value)
        }
      } else if(setting.current.item_base){
        setItemBase(options.name, null)
      } else {
        setItem(options.name, null)
      }
    } else if(options.file){
      if (options.value.length > 0) {
        const file = options.value[0]
        const fileReader = new FileReader();
        /* c8 ignore next 4 */
        fileReader.onload = (event) => {
          setItem(options.name, event.target.result)
          this.setCurrent({tmp_id: id}, setting)
        }
        fileReader.readAsDataURL(file);
      }
    } else if(options.checklist){
      const ovalue = ((setting.current.item_base) ? 
        setting.current.item_base[options.name] : 
        setting.current.item[options.name]) || ""
      let value = options.value
      if(options.checked){
        if(value !== "1"){
          if((ovalue !== "1")){
            value = ovalue + value
          }
        }
      } else if(setting.current.item_base){
          value = ovalue.replace(value,"")
      } else {
          value = ovalue.replace(value,"")
      }
      if(setting.current.item_base){
        setItemBase(options.name, value)
      } else {
        setItem(options.name, value)
      }
    } else if(setting.current.item_base){
      setItemBase(options.name, options.value)
    } else {
      setItem(options.name, options.value)
    }

    if(!options.file){
      this.setCurrent({tmp_id: id}, setting)
    }
  }

  exportTemplate() {
    const { saveToDisk } = this.app
    const { data, setData } = this.store
    setData("current", { side: SIDE_VISIBILITY.HIDE })
    const xtempl = JSON.stringify(data[APP_MODULE.TEMPLATE].template)
    const fUrl = URL.createObjectURL(new Blob([xtempl], 
      {type : 'text/json;charset=utf-8;'}));
    saveToDisk(fUrl, `${data[APP_MODULE.TEMPLATE].key}.json`)
  }

  goNext() {
    const { data } = this.store
    const { template, current } = data[APP_MODULE.TEMPLATE]
    const getNextItemId = () => {
      // tmp_section_index_type_subindex_subtype
      let section = current.section;
      let index = current.parent_index;
      let subindex = current.index;
      if (current.parent_index===null) {
        index = current.index; subindex = null;
      }
      let etype; let subtype;
      const sections = ["report","header","details","footer"];
      if (subindex!==null) {
        etype = getElementType(template[section][index]);
        if (subindex < template[section][index][etype].columns.length-1) {
          subtype = getElementType(template[section][index][etype].columns[subindex+1]);
          return{tmp_id: `tmp_${section}_${index.toString()}_${etype}_${(subindex+1).toString()}_${subtype}`}
        }
      }
      if (index!==null) {
        if (subindex===null) {
          etype = getElementType(template[section][index]);
          if (etype==="row" || etype==="datagrid") {
            if (template[section][index][etype].columns.length>0) {
              subtype = getElementType(template[section][index][etype].columns[0]);
              return {tmp_id: `tmp_${section}_${index.toString()}_${etype}_0_${subtype}`}
            }
          }
        }
        if (index < template[section].length-1) {
          etype = getElementType(template[section][index+1]);
          return {tmp_id: `tmp_${section}_${(index+1).toString()}_${etype}`}
        }
        if (section==="footer") {
          if (subindex!==null) {
            subtype = getElementType(template[section][index][etype].columns[subindex]);
            return {tmp_id: `tmp_${section}_${(index).toString()}_${etype}_${(subindex).toString()}_${subtype}`}
          } 
          return {tmp_id: `tmp_${section}_${(index).toString()}_${etype}`}
        } 
        section = sections[sections.indexOf(section)+1];
      }
      if (template[section].length>0) {
        etype = getElementType(template[section][0]);
        return {tmp_id: `tmp_${section}_0_${etype}`}
      } 
      if (section!=="footer") {
        section = sections[sections.indexOf(section)+1];
      }
      return {tmp_id: `tmp_${section}`}
      
    }
    this.setCurrent(getNextItemId())
  }

  goPrevious() {
    const { data } = this.store
    const { template, current } = data[APP_MODULE.TEMPLATE]
    const getPrevItemId = () => {
      // tmp_section_index_type_subindex_subtype
      let section = current.section;
      let index = current.parent_index;
      let subindex = current.index;
      if (current.parent_index===null) {
        index = current.index; subindex = null;
      }
      if (section==="report") {
        return {tmp_id: "tmp_report"}
      }
      let etype; let subtype;
      if (subindex!==null) {
        etype = getElementType(template[section][index]);
        if (subindex>0) {
          subtype = getElementType(template[section][index][etype].columns[subindex-1]);
          return {tmp_id: `tmp_${section}_${index.toString()}_${etype}_${(subindex-1).toString()}_${subtype}`}
        }
        return {tmp_id: `tmp_${section}_${index.toString()}_${etype}`}
      }
      if (index!==null) {
        if (index>0) {
          etype = getElementType(template[section][index-1]);
          if (etype==="row" || etype==="datagrid") {
            subindex = template[section][index-1][etype].columns.length;
            if (subindex>0) {
              subtype = getElementType(template[section][index-1][etype].columns[subindex-1]);
              return {tmp_id: `tmp_${section}_${(index-1).toString()}_${etype}_${(subindex-1).toString()}_${subtype}`}
            } 
            return {tmp_id: `tmp_${section}_${(index-1).toString()}_${etype}`} 
          } 
          return {tmp_id: `tmp_${section}_${(index-1).toString()}_${etype}`}
        }
        return {tmp_id: `tmp_${section}`}
      }
      const sections = ["report","header","details","footer"];
      section = sections[sections.indexOf(section)-1];
      if (section==="report") {
        return {tmp_id: "tmp_report"}
      }
      index = template[section].length;
      if (index>0) {
        etype = getElementType(template[section][index-1]);
        if (etype==="row" || etype==="datagrid") {
          subindex = template[section][index-1][etype].columns.length;
          if (subindex>0) {
            subtype = getElementType(template[section][index-1][etype].columns[subindex-1]);
            return {tmp_id: `tmp_${section}_${(index-1).toString()}_${etype}_${(subindex-1).toString()}_${subtype}`}
          } 
          return {tmp_id: `tmp_${section}_${(index-1).toString()}_${etype}`}
        } 
        return {tmp_id: `tmp_${section}_${(index-1).toString()}_${etype}`}
      } 
      return {tmp_id: `tmp_${section}`}
    }
    this.setCurrent(getPrevItemId())
  }

  moveDown() {
    const { data } = this.store
    const { current } = data[APP_MODULE.TEMPLATE]
    if (current.parent!==null && current.index!==null) {
      if (current.index<current.parent.length-1) {
        const next_item = current.parent[current.index+1];
        current.parent[current.index+1] = current.parent[current.index];
        current.parent[current.index] = next_item;
        let id = `tmp_${current.section}_`;
        if (current.parent_index!==null) {
          id += `${current.parent_index.toString()}_${current.parent_type}_`;
        }
        id += `${(current.index+1).toString()}_${current.type}`;
        this.setCurrent({ tmp_id: id, set_dirty: true})
      }
    }
  }

  moveUp() {
    const { data } = this.store
    const { current } = data[APP_MODULE.TEMPLATE]
    if (current.parent!==null && current.index!==null) {
      if (current.index>0) {
        const prev_item = current.parent[current.index-1];
        current.parent[current.index-1] = current.parent[current.index];
        current.parent[current.index] = prev_item;
        // tmp_section_index_type_subindex_subtype
        let id = `tmp_${current.section}_`;
        if (current.parent_index!==null) {
          id += `${current.parent_index.toString()}_${current.parent_type}_`;
        }
        id += `${(current.index-1).toString()}_${current.type}`;
        this.setCurrent({ tmp_id: id, set_dirty: true})
      }
    }
  }

  async onSideEvent({key, data}){
    const { showHelp, signOut } = this.app
    const { setData } = this.store
    const { theme } = this.store.data.current
    switch (key) {

      case SIDE_EVENT.LOGOUT:
        signOut()
        break;

      case SIDE_EVENT.THEME:
        const newTheme = (theme === APP_THEME.DARK) ? APP_THEME.LIGHT : APP_THEME.DARK
        setData("current", { 
          theme: newTheme
        })
        localStorage.setItem("theme", newTheme);
        break;

      case SIDE_EVENT.CHANGE:
        setData(APP_MODULE.TEMPLATE, {
          [data.fieldname]: data.value 
        })
        break;

      case SIDE_EVENT.SAVE:
        this.saveTemplate(data)
        break;

      case SIDE_EVENT.CREATE_REPORT:
        this.createTemplate()
        break;

      case SIDE_EVENT.DELETE:
        this.deleteTemplate()
        break;

      case SIDE_EVENT.CHECK:
        this.checkTemplate(data.value)
        break;

      case SIDE_EVENT.REPORT_SETTINGS:
        if(data.value === "JSON"){
          this.exportTemplate()
        } else {
          this.showPreview()
        }
        break;

      case SIDE_EVENT.HELP:
        showHelp(data.value)
        break;

      default:
        break;
    }
    return true
  }

  onTemplateEvent({key, data}){
    const { setData } = this.store
    const storeTemplate = this.store.data[APP_MODULE.TEMPLATE]
    switch (key) {

      case TEMPLATE_EVENT.ADD_ITEM:
        this.addItem(data)
        break;

      case TEMPLATE_EVENT.CHANGE_TEMPLATE:
        setData(APP_MODULE.TEMPLATE, {
          [data.key]: data.value
        })
        break;

      case TEMPLATE_EVENT.CHANGE_CURRENT:
        setData(APP_MODULE.TEMPLATE, {...storeTemplate,
          current: {...storeTemplate.current,
            [data.key]: data.value
          }
        })
        break;

      case TEMPLATE_EVENT.GO_PREVIOUS:
        this.goPrevious()
        break;

      case TEMPLATE_EVENT.GO_NEXT:
        this.goNext()
        break;

      case TEMPLATE_EVENT.CREATE_MAP:
        this.createMap(data.mapRef)
        break;

      case TEMPLATE_EVENT.SET_CURRENT:
        this.setCurrent(...data)
        break;

      case TEMPLATE_EVENT.MOVE_UP:
        this.moveUp()
        break;

      case TEMPLATE_EVENT.MOVE_DOWN:
        this.moveDown()
        break;

      case TEMPLATE_EVENT.DELETE_ITEM:
        this.deleteItem()
        break;

      case TEMPLATE_EVENT.EDIT_ITEM:
        this.editItem(data)
        break;

      case TEMPLATE_EVENT.EDIT_DATA_ITEM:
        this.editDataItem(data)
        break;

      case TEMPLATE_EVENT.SET_CURRENT_DATA:
        this.setCurrentData(data)
        break;

      case TEMPLATE_EVENT.SET_CURRENT_DATA_ITEM:
        this.setCurrentDataItem(data)
        break;

      case TEMPLATE_EVENT.ADD_TEMPLATE_DATA:
        this.addTemplateData()
        break;

      case TEMPLATE_EVENT.DELETE_DATA:
        this.deleteData(data)
        break;

      case TEMPLATE_EVENT.DELETE_DATA_ITEM:
        this.deleteDataItem(data)
        break;

      default:
        break;
    }
    return true
  }

  async saveTemplate(warning) {
    const { inputBox } = this.host
    const { requestData, resultError, msg } = this.app
    const { setData, data } = this.store
    const updateData = async ()=>{
      let setting = {...data[APP_MODULE.TEMPLATE]}
      setting.dbtemp.data.template = JSON.stringify(setting.template)
      const values = { 
        config_type: "CONFIG_REPORT",
        data: setting.dbtemp.data
      }
      const result = await requestData("/config/"+setting.dbtemp.code, { method: "PUT", data: values })
      if(result && result.error){
        resultError(result)
        return null
      }
      setting = {...setting,
        dirty: false
      }
      return setting
    }
    if(warning){
      const  modalForm = inputBox({ 
        title: msg("", { id: "template_label_template" }),
        message: msg("", { id: "msg_dirty_info" }),
        infoText: msg("", { id: "msg_delete_info" }),
        defaultOK: true,
        onEvent: {
          onModalEvent: async (modalResult) => {
            setData("current", { modalForm: null })
            if (modalResult.key === MODAL_EVENT.OK) {
              const result = await updateData()
              if(result){
                setData(APP_MODULE.TEMPLATE, result)
              }
            }
          }
        }
      })
      return setData("current", { modalForm, side: SIDE_VISIBILITY.HIDE })
    }
    return updateData()
  }

  setCurrent(options, template) {
    const { setData, data } = this.store
    const { mapRef } = data[APP_MODULE.TEMPLATE]
    const { msg } = this.app
    const item = options.tmp_id.split("_");
    let setting = {...(template || data[APP_MODULE.TEMPLATE]),
      current: {
        id: options.tmp_id,
        section: item[1]
      }
    }
    if(item.length === 2){
      setting = {...setting,
        current: {...setting.current,
          type: item[1],
          item: setting.template[setting.current.section],
          index: null,
          parent: null,
          parent_type: null,
          parent_index: null
        }
      }
    }
    if(item.length === 4){
      setting = {...setting,
        current: {...setting.current,
          type: item[3],
          index: parseInt(item[2],10),
          parent: setting.template[setting.current.section],
          parent_type: setting.current.section,
          parent_index: null
        }
      }
      if (["row","datagrid"].includes(setting.current.type)) {
        setting = {...setting,
          current: {...setting.current,
            item: setting.template[setting.current.section][parseInt(item[2],10)][item[3]].columns,
            item_base: setting.template[setting.current.section][parseInt(item[2],10)][item[3]]
          }
        }
      } else {
        setting = {...setting,
          current: {...setting.current,
            item: setting.template[setting.current.section][parseInt(item[2],10)][item[3]]
          }
        }
      }
    }
    if(item.length === 6){
      setting = {...setting,
        current: {...setting.current,
          type: item[5],
          item: setting.template[setting.current.section][parseInt(item[2],10)][item[3]].columns[parseInt(item[4],10)][item[5]],
          index: parseInt(item[4],10),
          parent: setting.template[setting.current.section][parseInt(item[2],10)][item[3]].columns,
          parent_type: item[3],
          parent_index: parseInt(item[2],10)
        }
      }
    }
    setting = {...setting,
      current: {...setting.current,
        form: templateElements({ msg })[setting.current.type]
      }
    }
    if(options.set_dirty && !["_blank","_sample"].includes(setting.key)){
      setting = {...setting,
        dirty: true
      }
    }
    setData(APP_MODULE.TEMPLATE, setting)
    if(mapRef){
      this.createMap()
    }
  }

  setCurrentData(_cdata) {
    const { showToast, msg } = this.app
    const { data, setData } = this.store
    let setting = {...data[APP_MODULE.TEMPLATE]}
    const cdata = _cdata
    const setCData=(values) => {
      if(!["_blank","_sample"].includes(setting.key)){
        setting = {...setting,
          dirty: true,
        }
      }
      setting = {...setting,
        current_data: values
      }
      setData(APP_MODULE.TEMPLATE, setting)
    }

    if(cdata){
      switch (cdata.type) {
        case "new":
          if (Object.keys(setting.template.data).includes(cdata.values.name) || cdata.values.name==="") {
            showToast(TOAST_TYPE.ERROR, msg("", { id: "msg_value_exists" }))
          } else if((cdata.values.type === "table") && (cdata.values.columns === "")) {
            showToast(TOAST_TYPE.ERROR, msg("", { id: "template_missing_columns" }))
          } else {
            const values = {
              name: cdata.values.name, 
              type: cdata.values.type 
            }
            switch (values.type) {
              case "list":
                setting = {...setting,
                  template: {...setting.template,
                    data: {...setting.template.data,
                      [values.name]: {}
                    }
                  }
                }
                values.items = [...getDataList({})]
                break;
              
              case "table":
                setting = {...setting,
                  template: {...setting.template,
                    data: {...setting.template.data,
                      [values.name]: []
                    }
                  }
                }
                const columns = cdata.values.columns.split(",")
                const item = {}
                for(let i = 0; i < columns.length; i += 1) {
                  item[String(columns[i]).trim()] = ""
                }
                setting.template.data[values.name] = [...setting.template.data[values.name], item]
                const table_data = getDataTable([item])
                values.items = table_data.items
                values.fields = table_data.fields
                break;

              case "string":
              default:
                setting = {...setting,
                  template: {...setting.template,
                    data: {...setting.template.data,
                      [values.name]: ""
                    }
                  }
                }
                break;
            }
            setting = {...setting,
              dataset: [...getDataset(setting.template.data)]
            }
            setCData(values)
          }
          break;

        case "list":
          cdata.items = getDataList(setting.template.data[cdata.name])
          setCData(cdata)
          break;
        
        case "table":
          const table_data = getDataTable(setting.template.data[cdata.name])
          cdata.items = table_data.items
          cdata.fields = table_data.fields
          setCData(cdata)
          break;
      
        default:
          setCData(cdata)
          break;
      }
    } else {
      setCData(cdata)
    }
  }

  setCurrentDataItem(value) {
    const { inputBox } = this.host
    const { showToast, msg } = this.app
    const { data, setData } = this.store
    let setting = {...data[APP_MODULE.TEMPLATE]}

    const setItem=(item) => {
      if(!["_blank","_sample"].includes(setting.key)){
        setting = {...setting,
          dirty: true
        }
      }
      setting = {...setting,
        current_data: {...setting.current_data,
          item
        }
      }
      setData(APP_MODULE.TEMPLATE, setting)
    }

    const newTable=() => {
      const item = {...setting.template.data[setting.current_data.name][0],
        _index: setting.template.data[setting.current_data.name].length
      }   
      Object.keys(item).forEach((fieldname) => {
        if(fieldname !== "_index"){
          item[fieldname] = ""
        }
      })
      setting = {...setting,
        template: {...setting.template,
          data: {...setting.template.data,
            [setting.current_data.name]: [...setting.template.data[setting.current_data.name], item]
          }
        }
      }
      setting = {...setting,
        current_data: {...setting.current_data,
          items: [...getDataTable(setting.template.data[setting.current_data.name]).items]
        }
      }
      setItem(item)
    }

    const newList=() => {
      const  modalForm = inputBox({ 
        title: msg("", { id: "msg_input_title" }),
        message: msg("", { id: "msg_new_fieldname" }),
        value: "", showValue: true,
        onEvent: {
          onModalEvent: async (modalResult) => {
            setData("current", { modalForm: null })
            if ((modalResult.key === MODAL_EVENT.OK) && (modalResult.data.value !== "")) {
              if (Object.keys(setting.template.data[setting.current_data.name]).includes(modalResult.data.value)) {
                showToast(TOAST_TYPE.ERROR, msg("", { id: "msg_value_exists" }))
              } else {
                setting = {...setting,
                  template: {...setting.template,
                    data: {...setting.template.data,
                      [setting.current_data.name]: {...setting.template.data[setting.current_data.name], 
                        [modalResult.data.value]: ""
                      }
                    }
                  }
                }
                setting = {...setting,
                  current_data: {...setting.current_data,
                    items: [...getDataList(setting.template.data[setting.current_data.name])]
                  }
                }
                setItem(modalResult.data.value)
              }
            }
          }
        }
      })
      return setData("current", { modalForm, side: SIDE_VISIBILITY.HIDE })
    }

    if(typeof value === "undefined"){
      if(setting.current_data.type === "list"){
        newList()
      } 
      if(setting.current_data.type === "table"){
        newTable()
      }
    } else {
      setItem(value)
    }
  }

  async setTemplate(options) {
    const { resultError, currentModule } = this.app
    let blank = {
      id: undefined,
      key: options.type,
      title: "Nervatura Report",
      template: {
        meta: {
          report_key: options.type,
          nervatype: "", report_type: "", direction: "", 
          report_name: "Nervatura Report", description: "", 
          filetype:"pdf"
        },
        report: {},
        header: [],
        details: [],
        footer: [],
        sources: {},
        data: {}
      },
      current: {},
      current_data: null,
      dataset: [],
      docnumber: "",
      tabView: "template",
      dirty: false
    }
    if (options.type === "_sample"){
      const { sample } = await import('../config/sample.js');
      blank = {...blank,
        template: {...blank.template,
          ...sample
        },
        dataset: getDataset(sample.data)
      }
    }
    if (options.type === "template"){
      let report_template = {}
      try {
        report_template = JSON.parse(options.report.data.template)
      } catch (err) {
        resultError({ error: { message: err.message } })
        return currentModule({ 
          data: { module: APP_MODULE.TEMPLATE, side: SIDE_VISIBILITY.HIDE },
          content: { fkey: "setTemplate", args: [ { type: "_sample" } ] } 
        })
      }
      blank = {...blank,
        id: options.report.id,
        key: options.report.code,
        title: options.report.data.report_name,
        template: report_template,
        dbtemp: options.report
      }
      blank = {...blank,
        dataset: getDataset(blank.template.data)
      }
    }
    return this.setCurrent({tmp_id: "tmp_report"}, blank)
  }

  showPreview(orient) {
    const { inputBox } = this.host
    const { getSetting, msg, requestData, resultError } = this.app
    const { data, setData } = this.store
    const loadPreview = async (params) => {
      const result = await requestData("/service/function", { method: "POST", data: params })
      if(result.error){
        resultError(result)
      } else {
        const resultUrl = URL.createObjectURL(result, {type : "application/pdf"})
        window.open(resultUrl,"_blank")
      }
    }
    const setting = {...data[APP_MODULE.TEMPLATE]}
    const params = {
      name: "report_get",
      values: {
        report_key: setting.key,
        orientation: orient || getSetting("page_orient"),
        size: getSetting("page_size"),
        title: setting.title,
        template: JSON.stringify(setting.template),
        output: "auto"
      }
    }
    if(["_blank","_sample"].includes(setting.key)){
      params.reportkey = ""
      loadPreview(params)
    } else {
      const  modalForm = inputBox({ 
        title: msg("", { id: "template_preview_data" }),
        message: msg("", { id: "template_preview_input" }).replace("docname", setting.title),
        value: setting.docnumber, showValue: true,
        onEvent: {
          onModalEvent: async (modalResult) => {
            setData("current", { modalForm: null })
            if ((modalResult.key === MODAL_EVENT.OK) && (modalResult.data.value !== "")) {
              const template = {...setting,
                docnumber: modalResult.data.value
              }
              setData(APP_MODULE.TEMPLATE, template )
              params.values.code = modalResult.data.value
              loadPreview(params)
            }
            return true
          }
        }
      })
      setData("current", { modalForm, side: SIDE_VISIBILITY.HIDE })
    }
  }

}