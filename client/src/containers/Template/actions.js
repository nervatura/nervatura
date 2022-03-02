import update from 'immutability-helper';
import printJS from 'print-js'
import format from 'date-fns/format'

import { appActions, saveToDisk } from 'containers/App/actions'
import InputBox from 'components/Modal/InputBox'
import TemplateData from 'components/Modal/Template'
import { getSetting } from 'config/app'
import { templateElements } from 'containers/Template/Template'
import { InitItem } from 'containers/Controller/Validator'

export const templateActions = (data, setData) => {
  const app = appActions(data, setData)

  const getElementType = (element) => {
    if (Object.getOwnPropertyNames(element).length>0) {
      return Object.getOwnPropertyNames(element)[0];
    } else {
      return null;
    }
  }
 
  const getDataset = (data)=>{
    let dataset = [];
    Object.keys(data).forEach((dskey) => {
      if (typeof data[dskey] === "string"){
        dataset.push({ lslabel: dskey, lsvalue: "string" })
      } else {
        if (Array.isArray(data[dskey])){
          dataset.push({ lslabel: dskey, lsvalue: "table" }) 
        } else {
          dataset.push({ lslabel: dskey, lsvalue: "list" })
        }
      }
    });
    return dataset;
  }

  const setCurrent = (options, template) => {
    const item = options.tmp_id.split("_");
    let setting = update(template || data.template, {$merge: {
      current: {
        id: options.tmp_id,
        section: item[1]
      }
    }})
    if(item.length === 2){
      setting = update(setting, {current: {$merge: {
        type: item[1],
        item: setting.template[setting.current.section],
        index: null,
        parent: null,
        parent_type: null,
        parent_index: null
      }}})
    }
    if(item.length === 4){
      setting = update(setting, {current: {$merge: {
        type: item[3],
        index: parseInt(item[2],10),
        parent: setting.template[setting.current.section],
        parent_type: setting.current.section,
        parent_index: null
      }}})
      if (["row","datagrid"].includes(setting.current.type)) {
        setting = update(setting, {current: {$merge: {
          item: setting.template[setting.current.section][parseInt(item[2],10)][item[3]].columns,
          item_base: setting.template[setting.current.section][parseInt(item[2],10)][item[3]]
        }}})
      } else {
        setting = update(setting, {current: {$merge: {
          item: setting.template[setting.current.section][parseInt(item[2],10)][item[3]]
        }}})
      }
    }
    if(item.length === 6){
      setting = update(setting, {current: {$merge: {
        type: item[5],
        item: setting.template[setting.current.section][parseInt(item[2],10)][item[3]].columns[parseInt(item[4],10)][item[5]],
        index: parseInt(item[4],10),
        parent: setting.template[setting.current.section][parseInt(item[2],10)][item[3]].columns,
        parent_type: item[3],
        parent_index: parseInt(item[2],10)
      }}})
    }
    setting = update(setting, {current: {$merge: {
      form: templateElements({ getText: app.getText })[setting.current.type]
    }}})
    if(options.set_dirty && !["_blank","_sample"].includes(setting.key)){
      setting = update(setting, {$merge: {
        dirty: true
      }})
    }
    setData("template", setting)
  }

  const createMap = (cv) => {
    const { template, current } = data.template
    let cont = cv.getContext('2d')
    let cell_color = "#CCCCCC"; let row_color = "#FFFF00"; let sel_color = "#00EE00";
    let cell_size = 8; let cell_pad = 1; let page_pad = 3; let rows = [];
    let def_height = 165;

    cv.height = page_pad;
    let sections = ["header","details","footer"];
    for(let s = 0; s < sections.length; s++) {
      for(let i = 0; i < template[sections[s]].length; i++) {
        let row = {};
        row.type = getElementType(template[sections[s]][i]);
        let item = template[sections[s]][i][row.type];
        if (row.type==="row" || row.type==="datagrid") {
          row.cols = template[sections[s]][i][row.type].columns.length;
          row.selected = (item.columns===current.item || item.columns===current.parent);
          row.selcol = -1;
          if (row.selected) {
            for(let c = 0; c < row.cols; c++) {
              let cname = getElementType(template[sections[s]][i][row.type].columns[c]);
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
    for(let r = 0; r < rows.length; r++) {
      if(rows[r].type === "row"){
        coldif = (cv.width - (rows[r].cols*(cell_size+cell_pad)+2*page_pad))/rows[r].cols;
        for(let cr = 0; cr < rows[r].cols; cr++) {
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
        for(let cc = 0; cc < rows[r].cols; cc++) {
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
  }

  const goPrevious = () => {
    const getPrevItemId = () => {
      //tmp_section_index_type_subindex_subtype
      const { current, template } = data.template
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
          return {tmp_id: "tmp_"+section+"_"+index.toString()+"_"+etype+"_"+(subindex-1).toString()+"_"+subtype}
        }
        return {tmp_id: "tmp_"+section+"_"+index.toString()+"_"+etype}
      }
      if (index!==null) {
        if (index>0) {
          etype = getElementType(template[section][index-1]);
          if (etype==="row" || etype==="datagrid") {
            subindex = template[section][index-1][etype].columns.length;
            if (subindex>0) {
              subtype = getElementType(template[section][index-1][etype].columns[subindex-1]);
              return {tmp_id: "tmp_"+section+"_"+(index-1).toString()+"_"+etype+"_"+(subindex-1).toString()+"_"+subtype}
            } else {
              return {tmp_id: "tmp_"+section+"_"+(index-1).toString()+"_"+etype}
            }
          } else {
            return {tmp_id: "tmp_"+section+"_"+(index-1).toString()+"_"+etype}
          }
        }
        return {tmp_id: "tmp_"+section}
      }
      let sections = ["report","header","details","footer"];
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
            return {tmp_id: "tmp_"+section+"_"+(index-1).toString()+"_"+etype+"_"+(subindex-1).toString()+"_"+subtype}
          } else {
            return {tmp_id: "tmp_"+section+"_"+(index-1).toString()+"_"+etype}
          }
        } else {
          return {tmp_id: "tmp_"+section+"_"+(index-1).toString()+"_"+etype}
        }
      } else {
        return {tmp_id: "tmp_"+section}
      }
    }
    setCurrent(getPrevItemId())
  }

  const goNext = () => {
    const getNextItemId = () => {
      //tmp_section_index_type_subindex_subtype
      const { current, template } = data.template
      let section = current.section;
      let index = current.parent_index;
      let subindex = current.index;
      if (current.parent_index===null) {
        index = current.index; subindex = null;
      }
      let etype; let subtype;
      let sections = ["report","header","details","footer"];
      if (subindex!==null) {
        etype = getElementType(template[section][index]);
        if (subindex < template[section][index][etype].columns.length-1) {
          subtype = getElementType(template[section][index][etype].columns[subindex+1]);
          return{tmp_id: "tmp_"+section+"_"+index.toString()+"_"+etype+"_"+(subindex+1).toString()+"_"+subtype}
        }
      }
      if (index!==null) {
        if (subindex===null) {
          etype = getElementType(template[section][index]);
          if (etype==="row" || etype==="datagrid") {
            if (template[section][index][etype].columns.length>0) {
              subtype = getElementType(template[section][index][etype].columns[0]);
              return {tmp_id: "tmp_"+section+"_"+index.toString()+"_"+etype+"_0_"+subtype}
            }
          }
        }
        if (index < template[section].length-1) {
          etype = getElementType(template[section][index+1]);
          return {tmp_id: "tmp_"+section+"_"+(index+1).toString()+"_"+etype}
        }
        if (section==="footer") {
          if (subindex!==null) {
            subtype = getElementType(template[section][index][etype].columns[subindex]);
            return {tmp_id: "tmp_"+section+"_"+(index).toString()+"_"+etype+"_"+(subindex).toString()+"_"+subtype}
          } else {
            return {tmp_id: "tmp_"+section+"_"+(index).toString()+"_"+etype}
          }
        } else {
          section = sections[sections.indexOf(section)+1];
        }
      }
      if (template[section].length>0) {
        etype = getElementType(template[section][0]);
        return {tmp_id: "tmp_"+section+"_0_"+etype}
      } else {
        if (section!=="footer") {
          section = sections[sections.indexOf(section)+1];
        }
        return {tmp_id: "tmp_"+section}
      }
    }
    setCurrent(getNextItemId())
  }

  const moveDown = () => {
    const { current } = data.template;
    if (current.parent!==null && current.index!==null) {
      if (current.index<current.parent.length-1) {
        let next_item = current.parent[current.index+1];
        current.parent[current.index+1] = current.parent[current.index];
        current.parent[current.index] = next_item;
        let id = "tmp_"+current.section+"_";
        if (current.parent_index!==null) {
          id += current.parent_index.toString()+"_"+current.parent_type+"_";
        }
        id += (current.index+1).toString()+"_"+current.type;
        setCurrent({ tmp_id: id, set_dirty: true})
      }
    }
  }

  const moveUp = () => {
    const { current } = data.template;
    if (current.parent!==null && current.index!==null) {
      if (current.index>0) {
        let prev_item = current.parent[current.index-1];
        current.parent[current.index-1] = current.parent[current.index];
        current.parent[current.index] = prev_item;
        //tmp_section_index_type_subindex_subtype
        let id = "tmp_"+current.section+"_";
        if (current.parent_index!==null) {
          id += current.parent_index.toString()+"_"+current.parent_type+"_";
        }
        id += (current.index-1).toString()+"_"+current.type;
        setCurrent({ tmp_id: id, set_dirty: true})
      }
    }
  }

  const deleteItem = () => {
    const { current } = data.template;
    if (current.parent!==null && current.index!==null) {
      current.parent.splice(current.index, 1);
      let id = "tmp_"+current.section;
      if (current.parent_index!==null) {
        id += "_"+current.parent_index.toString()+"_"+current.parent_type;
      }
      setCurrent({ tmp_id: id, set_dirty: true})
    }
  }

  const addItem = (value) => {
    const { current } = data.template;
    if (value !== "") {
      let ename = value.toString().toLowerCase();
      let element = {}; element[ename] = {};
      let id = current.id+"_"+current.item.length.toString()+"_"+ename;
      if (ename==="datagrid" || ename==="row") {
        element[ename].columns=[];
      }
      current.item.push(element);
      setCurrent({ tmp_id: id, set_dirty: true})
    }
  }

  const editItem = (options) => {
    let setting = update(data.template, {})
    const id = setting.current.id
    const itemId = id.split("_");

    const setItemBase = (key, value) => {
      if(value === null){
        setting = update(setting, {current: {item_base: {
          $unset: [key]
        }}})
      } else {
        setting = update(setting, {current: {item_base: {$merge: {
          [key]: value
        }}}})
      }
      setting = update(setting, { template: {
        [setting.current.section]: {[parseInt(itemId[2],10)]: {[itemId[3]]: { 
          $set: setting.current.item_base
      }}}}})
    }

    const setItem = (key, value) => {
      if(value === null){
        setting = update(setting, {current: {item: {
          $unset: [key]
        }}})
      } else {
        setting = update(setting, {current: {item: {$merge: {
          [key]: value
        }}}})
      }
      switch (itemId.length) {
        case 2:
          setting = update(setting, {template: {[setting.current.section]: {
            $set: setting.current.item
          }}})
          break;

        case 4:
          setting = update(setting, {template: {[setting.current.section]: { [parseInt(itemId[2],10)]: {[itemId[3]]: {
            $set: setting.current.item
          }}}}})
          break;

        case 6:
        default:
          setting = update(setting, {template: {[setting.current.section]: { [parseInt(itemId[2],10)]: {[itemId[3]]: {columns: {[parseInt(itemId[4],10)]: {[itemId[5]]: {
            $set: setting.current.item
          }}}}}}}})
          break;
      }
    }

    if(!["_blank","_sample"].includes(setting.key)){
      setting = update(setting, {$merge: {
        dirty: true
      }})
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
      } else {
        if(setting.current.item_base){
          setItemBase(options.name, null)
        } else {
          setItem(options.name, null)
        }
      }
    } else if(options.file){
      if (options.value.length > 0) {
        let file = options.value[0]
        let fileReader = new FileReader();
        /* istanbul ignore next */
        fileReader.onload = function(event) {
          setItem(options.name, event.target.result)
          setCurrent({tmp_id: id}, setting)
        }
        fileReader.readAsDataURL(file);
      }
    } else if(options.checklist){
      let ovalue = ((setting.current.item_base) ? 
        setting.current.item_base[options.name] : 
        setting.current.item[options.name]) || ""
      let value = options.value
      if(options.checked){
        if(value !== "1"){
          if((ovalue !== "1")){
            value = ovalue + value
          }
        }
      } else {
        if(setting.current.item_base){
          value = ovalue.replace(value,"")
        } else {
          value = ovalue.replace(value,"")
        }
      }
      if(setting.current.item_base){
        setItemBase(options.name, value)
      } else {
        setItem(options.name, value)
      }
    } else {
      if(setting.current.item_base){
        setItemBase(options.name, options.value)
      } else {
        setItem(options.name, options.value)
      }
    }

    if(!options.file){
      setCurrent({tmp_id: id}, setting)
    }
  }

  const exportTemplate = () => {
    setData("current", { side: "hide"}, ()=>{
      const xtempl = JSON.stringify(data.template.template)
      let fUrl = URL.createObjectURL(new Blob([xtempl], 
        {type : 'text/json;charset=utf-8;'}));
      saveToDisk(fUrl, data.template.key+".json")
    })
  }

  const setTemplate = (options) => {
    let blank = {
      id: undefined,
      key: options.type,
      title: "Nervatura Report",
      template: {
        meta: {
          reportkey: options.type,
          nervatype: "", transtype: "", direction: "", 
          repname: "Nervatura Report", description: "", 
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
      const sample = require('../../config/sample.json')
      blank = update(blank, {
        template: {$merge: {
          ...sample
        }},
        dataset: {$set: getDataset(sample.data)}
      })
    }
    if (options.type === "template"){
      blank = update(blank, {$merge: {
        id: options.dataset.template[0].id,
        key: options.dataset.template[0].reportkey,
        title: options.dataset.template[0].repname,
        template: JSON.parse(options.dataset.template[0].report),
        dbtemp: options.dataset.template[0]
      }})
      blank = update(blank, {$merge: {
        dataset: getDataset(blank.template.data)
      }})
    }
    setCurrent({tmp_id: "tmp_report"}, blank)
    setData("current", { module: "template", side: "hide" })
  }

  const saveTemplate = async (warning) => {
    const updateData = async ()=>{
      let setting = update(data.template, {})
      let values = { 
        id: setting.id,
        report: JSON.stringify(setting.template)
      }
      let result = await app.requestData("/ui_report", { method: "POST", data: [values] })
      if(result.error){
        app.resultError(result)
        return null
      }
      setting = update(setting, {$merge: {
        dirty: false
      }})
      return setting
    }
    if(warning){
      setData("current", { modalForm: 
        <InputBox 
          title={app.getText("template_label_template")}
          message={app.getText("msg_dirty_info")}
          infoText={app.getText("msg_delete_info")}
          labelOK={app.getText("msg_ok")}
          labelCancel={app.getText("msg_cancel")} defaultOK={true}
          onCancel={() => {
            setData("current", { modalForm: null })
          }}
          onOK={(value) => {
            setData("current", { modalForm: null }, async ()=>{
              const result = await updateData()
              if(result){
                setData("template", result)
              }
            })
          }}
        />,
        side: "hide"
      })
    } else {
      return await updateData()
    }
  }

  const deleteTemplate = async () => {
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
            const result = await app.requestData("/ui_report", 
              { method: "DELETE", query: { id: data.template.id } })
            if(result && result.error){
              return app.resultError(result)
            } 
            setData("current", { module: "setting", content: { type: 'template' } })
          })
        }}
      />,
      side: "hide"
    })
  }

  const deleteData = (dskey) => {
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
            let setting = update(data.template, {template: {data: {
              $unset: [dskey]
            }}})
            if(!["_blank","_sample"].includes(setting.key)){
              setting = update(setting, {$merge: {
                dirty: true
              }})
            }
            setting = update(setting, { $merge: {
              dataset: getDataset(setting.template.data)
            }})
            setData("template", setting)
          })
        }}
      /> 
    })
  }
  
  const getDataList = (data)=>{
    let datalist = [];
    Object.keys(data).forEach((key) => {
      datalist.push({ lslabel: key, lsvalue: data[key] }) 
    });
    return datalist;
  }

  const getDataTable = (data)=>{
    let table = { fields: {}, items: []};
    if(data.length > 0){
      Object.keys(data[0]).forEach((key) => {
        table.fields[key] = { fieldtype: 'string', label: key }
      })
      for (var index = 0; index < data.length; index++) {
        const item = update(data[index], {$merge: {
          _index: index
        }})
        table.items.push(item) 
      } 
    }
    return table;
  }

  const setCurrentData = (cdata) => {
    let setting = update(data.template, {})

    const setCData=(values) => {
      if(!["_blank","_sample"].includes(setting.key)){
        setting = update(setting, {$merge: {
          dirty: true
        }})
      }
      setting = update(setting, { $merge: {
        current_data: values
      }})
      setData("template", setting)
    }

    if(cdata){
      switch (cdata.type) {
        case "new":
          if (Object.keys(setting.template.data).includes(cdata.values.name) || cdata.values.name==="") {
            app.showToast({ type: "error",
              title: app.getText("msg_warning"), 
              message: app.getText("msg_value_exists") })
          } else if((cdata.values.type === "table") && (cdata.values.columns === "")) {
            app.showToast({ type: "error",
              title: app.getText("msg_warning"), 
              message: app.getText("template_missing_columns") })
          } else {
            let values = update({}, {$set: {
              name: cdata.values.name, 
              type: cdata.values.type 
            }})
            switch (values.type) {
              case "string":
                setting = update(setting, { template: {data: {$merge: {
                  [values.name]: ""
                }}}})
                break;

              case "list":
                setting = update(setting, { template: {data: {$merge: {
                  [values.name]: {}
                }}}})
                values.items = getDataList({})
                break;
              
              case "table":
                setting = update(setting, { template: {data: {$merge: {
                  [values.name]: []
                }}}})
                let columns = cdata.values.columns.split(",")
                let item = {}
                for(let i = 0; i < columns.length; i++) {
                  item[String(columns[i]).trim()] = ""
                }
                setting = update(setting, { template: {data: {
                  [values.name]: {$push: [item]}
                }}})
                const table_data = getDataTable([item])
                values.items = table_data.items
                values.fields = table_data.fields
                break;

                default:
                  break;
            }
            setting = update(setting, {$merge: {
              dataset: getDataset(setting.template.data)
            }})
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

  const setCurrentDataItem = (value) => {
    let setting = update(data.template, {})

    const setItem=(item) => {
      if(!["_blank","_sample"].includes(setting.key)){
        setting = update(setting, {$merge: {
          dirty: true
        }})
      }
      setting = update(setting, {current_data: { $merge: {
        item: item
      }}})
      setData("template", setting)
    }
    
    const newList=() => {
      setData("current", { modalForm: 
        <InputBox 
          title={app.getText("msg_input_title")}
          message={app.getText("msg_new_fieldname")}
          value="" showValue={true}
          labelOK={app.getText("msg_ok")}
          labelCancel={app.getText("msg_cancel")}
          onCancel={() => {
            setData("current", { modalForm: null })
          }}
          onOK={(value) => {
            setData("current", { modalForm: null }, async ()=>{
              if (value !== "") { 
                if (Object.keys(setting.template.data[setting.current_data.name]).includes(value)) {
                  app.showToast({ type: "error",
                    title: app.getText("msg_warning"), 
                    message: app.getText("msg_value_exists") })
                } else {
                  setting = update(setting, { template: {data: {[setting.current_data.name]: {$merge: {
                    [value]: ""
                  }}}}})
                  setting = update(setting, { current_data: {$merge: {
                    items: getDataList(setting.template.data[setting.current_data.name])
                  }}})
                  setItem(value)
                }
              }
            })
          }}
        /> 
      })
    }

    const newTable=() => {      
      let item = update(setting.template.data[setting.current_data.name][0], {$merge:{
        _index: setting.template.data[setting.current_data.name].length
      }})
      Object.keys(item).forEach((fieldname) => {
        if(fieldname !== "_index"){
          item[fieldname] = ""
        }
      })
      setting = update(setting, { template: {data: {
        [setting.current_data.name]: {$push: [item]
      }}}})
      setting = update(setting, { current_data: {$merge: {
        items: getDataTable(setting.template.data[setting.current_data.name]).items
      }}})
      setItem(item)
    }

    if(typeof value === "undefined"){
      if(setting.current_data.type === "list"){
        newList()
      } else if(setting.current_data.type === "table"){
        newTable()
      }
    } else {
      setItem(value)
    }

  }

  const deleteDataItem = (options) => {
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
            let setting = update(data.template, {})
  
            switch (setting.current_data.type) {
              case "list":
                setting = update(setting, {template: {data: {[setting.current_data.name]: {
                  $unset: [options.key]
                }}}})
                setting = update(setting, { current_data: {$merge: {
                  items: getDataList(setting.template.data[setting.current_data.name])
                }}})
                break;
    
              case "table":
              default:
                if (setting.template.data[setting.current_data.name].length===1){
                  setting = update(setting, {template: {data: {
                    $unset: [setting.current_data.name]
                  }}})
                  setting = update(setting, {$merge: {
                    current_data: null,
                    dataset: getDataset(setting.template.data)
                  }})
                } else {
                  setting = update(setting, {template: {data: {[setting.current_data.name]: {
                    $splice: [[options._index, 1]]
                  }}}})
                  setting = update(setting, { current_data: {$merge: {
                    items: getDataTable(setting.template.data[setting.current_data.name]).items
                  }}})
                }
                break;
            }
            
            setting = update(setting, { $merge: {
              dataset: getDataset(setting.template.data)
            }})
  
            if(!["_blank","_sample"].includes(setting.key)){
              setting = update(setting, {$merge: {
                dirty: true
              }})
            }
            setData("template", setting)
          })
        }}
      /> 
    })
  }

  const editDataItem = (options) => {
    let setting = update(data.template, {})
    if(!["_blank","_sample"].includes(setting.key)){
      setting = update(setting, {$merge: {
        dirty: true
      }})
    }
    switch (setting.current_data.type) {
      case "string":
        setting = update(setting, { template: {data: {$merge: {
          [setting.current_data.name]: options.value
        }}}})
        break;
      
      case "list":
        setting = update(setting, { template: {data: {[setting.current_data.name]: {$merge: {
          [setting.current_data.item]: options.value
        }}}}})
        setting = update(setting, { current_data: {$merge: {
          items: getDataList(setting.template.data[setting.current_data.name])
        }}})
        break;

      default:
        setting = update(setting, { current_data: {item: {$merge: {
          [options.field]: options.value
        }}}})
        setting = update(setting, { template: {data: {
          [setting.current_data.name]: {[options._index]: {$merge: {
          [options.field]: options.value
        }}}}}})
        setting = update(setting, { current_data: {$merge: {
          items: getDataTable(setting.template.data[setting.current_data.name]).items
        }}})
        break;
    }
    setData("template", setting)
  }

  const addTemplateData = () => {
    setData("current", { modalForm: 
      <TemplateData
        name="" type="string" columns=""
        getText={app.getText}
        onClose={() => {
          setData("current", { modalForm: null })
        }}
        onData={(values) => {
          setData("current", { modalForm: null }, async ()=>{
            setCurrentData({ name: "new", type: "new", values: {...values} })
          })
        }}
      /> 
    })
  }

  const showPreview = (orient) => {
    const loadPreview = async (data) => {
      let result = await app.requestData("/report", { method: "POST", data: data })
      if(result.error){
        app.resultError(result)
        return null
      }
      printJS({
        printable: URL.createObjectURL(result, {type : "application/pdf"}),
        type: 'pdf',
        base64: false,
      })
    }
    let setting = update(data.template, {})
    const params = {
      reportkey: setting.key,
      orientation: orient || getSetting("page_orient"),
      size: getSetting("page_size"),
      output: "auto",
      title: setting.title,
      template: JSON.stringify(setting.template)
    }
    if(["_blank","_sample"].includes(setting.key)){
      params.reportkey = ""
      loadPreview(params)
    } else {
      params.nervatype = setting.template.meta.nervatype
      if(setting.preview && (setting.docnumber !== "")){
        params.refnumber = setting.docnumber
        loadPreview(params)
      } else {
        setData("current", { modalForm: 
          <InputBox 
            title={app.getText("template_preview_data")}
            message={app.getText("template_preview_input").replace("docname",params.nervatype)}
            value={setting.docnumber} showValue={true}
            labelOK={app.getText("msg_ok")}
            labelCancel={app.getText("msg_cancel")}
            onCancel={() => {
              setData("current", { modalForm: null })
            }}
            onOK={(value) => {
              setData("current", { modalForm: null }, async ()=>{
                if(value !== ""){
                  const template = update(setting.template, {$merge: {
                    docnumber: value
                  }})
                  setData("template", { template: template })
                  params.refnumber = value
                  loadPreview(params)
                }
              })
            }}
          />,
          side: "hide"
        })
      }
    }
  }

  const changeTemplateData = (options) => {
    const { key, value } = options
    let setting = update(data.template, {$merge: {
      [key]: value
    }})
    setData("template", setting)
  }

  const changeCurrentData = (options) => {
    const { key, value } = options
    let setting = update(data.template, {current: {$merge: {
      [key]: value
    }}})
    setData("template", setting)
  }

  const checkTemplate = (options, nextKey) => {
    const cbNext = {
      NEW_BLANK: () => setTemplate({ type: "_blank" }),
      NEW_SAMPLE: () => setTemplate({ type: "_sample" }),
      LOAD_SETTING: () => setData("current", { module: "setting", content: { type: 'template' }, side: "hide" })
    }
    if (data.template.dirty === true) {
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
                return cbNext[nextKey]()
              })
            })
          }}
          onOK={() => {
            setData("current", { modalForm: null }, async ()=>{
              const setting = await saveTemplate()
              if(setting){
                return setData("template", setting, ()=>{
                  return cbNext[nextKey]()
                })
              }
            })
          }}
        /> 
      })
    }
    return cbNext[nextKey]()
  }

  const createTemplate = () => {
    const tableValues = (type, item) => {
      let values = {}
      const baseValues = InitItem(data, setData)({tablename: type, 
        dataset: data.setting.dataset, current: data.setting.current})
      for (const key in item) {
        if (baseValues.hasOwnProperty(key)) {
          values[key] = item[key]
        }
      }
      return values
    }
    let reportkey = data.template.template.meta.nervatype;
    if (reportkey === "trans") {
      reportkey = data.template.template.meta.transtype+"_"+data.template.template.meta.direction;
    }
    reportkey += "_"+format(new Date(),"yyyyMMddHHmm")
    setData("current", { modalForm: 
      <InputBox 
        title={app.getText("template_label_new")}
        message={reportkey}
        value={data.template.repname} showValue={true}
        labelOK={app.getText("msg_ok")}
        labelCancel={app.getText("msg_cancel")}
        onCancel={() => {
          setData("current", { modalForm: null })
        }}
        onOK={(value) => {
          setData("current", { modalForm: null }, async ()=>{
            const template = update(data.template.template, {
              meta: {$merge: {
                reportkey: reportkey,
                repname: value
              }}
            })
            let values = update(tableValues("report", data.template.dbtemp), {$merge: {
              id: null,
              reportkey: reportkey,
              repname: value,
              report: JSON.stringify(template)
            }})
            values = update(values, {
              $unset: ["orientation", "size"]
            })
            let result = await app.requestData("/ui_report", { method: "POST", data: [values] })
            if(result.error){
              return app.resultError(result)
            }
            setData("current", { module: "setting", content: { type: "template", id: result[0] }, side: "hide" })
          })
        }}
      />,
      side: "hide"
    })
  }

  return {
    addItem: addItem,
    addTemplateData: addTemplateData,
    changeCurrentData: changeCurrentData,
    changeTemplateData: changeTemplateData,
    checkTemplate: checkTemplate,
    createMap: createMap,
    createTemplate: createTemplate,
    deleteData: deleteData,
    deleteDataItem: deleteDataItem,
    deleteItem: deleteItem,
    deleteTemplate: deleteTemplate,
    editDataItem: editDataItem,
    editItem: editItem,
    exportTemplate: exportTemplate,
    getElementType: getElementType,
    getDataset: getDataset,
    goNext: goNext,
    goPrevious: goPrevious,
    moveDown: moveDown, 
    moveUp: moveUp, 
    saveTemplate: saveTemplate,
    setCurrentDataItem: setCurrentDataItem,
    setCurrentData: setCurrentData,
    setCurrent: setCurrent,
    setTemplate: setTemplate,
    showPreview: showPreview,
  }
}