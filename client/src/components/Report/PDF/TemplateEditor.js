import { Fragment, useCallback } from 'react';
import PropTypes from 'prop-types';
import update from 'immutability-helper';

import Icon from 'components/Form/Icon'
import Label from 'components/Form/Label'
import Button from 'components/Form/Button'
import Table from 'components/Form/Table';
import List from 'components/Form/List';
import Select from 'components/Form/Select'
import Row from 'components/Form/Row'
import { getSetting } from 'config/app'

import styles from './TemplateEditor.module.css';

export const TemplateEditor = ({ 
  data, paginationPage, className,
  getText, onEvent,
  ...props 
}) => {
  const { title, tabView, template, current, current_data, dataset } = data
  const mapRef = useCallback(map => {
    if (map) {
      onEvent("createMap",[map])
    }
  }, [onEvent])

  const getMapCtr = (type, key) => {
    const keyMap = {
      map_edit: { data: false, report: false, header: false, footer: false, details: false },
      map_insert: { header: true, footer: true, details: true, row: true, datagrid: true }
    }
    return keyMap[key][type]
  }

  const getElementType = (element) => {
    return Object.getOwnPropertyNames(element)[0];
  }

  const getBadge = (items, index) => {
    if (typeof index!=="undefined") {
      return index
    }
    if ((typeof items==="object") && (Array.isArray(items)) && (items.length > 0)) {
      return items.length
    }
    return 0
  }

  const setListIcon = (item, index) => {
    if (current.item===item) {
      return {
        selected: true, icon: "Tag", color: "green", 
        badge: getBadge(item, index)
      };
    }
    if (current.parent===item || template[current.section]===item) {
      return {
        selected: true, icon: "Check", color: "", 
        badge: getBadge(item, index) 
      };
    }
    let icon = "InfoCircle"
    if (Array.isArray(item)) {
      if (item.length>0) {
        icon = "Plus"
      }
    }
    return { selected: false, icon: icon, color: "", badge: 0 };
  }

  const mapButton = (key, tmp_id, info, label, badge, color) => {
    return <Button id={`btn_${key}`}
      onClick={() => onEvent("setCurrent",[{ tmp_id: tmp_id }]) } 
      className={`${styles.mapButton} ${"full border-button"} ${(info.color!=="")?styles.green:""}`}
      value={
        <div className="row full" >
          <div className="cell" >
            <Label value={label} 
              leftIcon={<Icon iconKey={info.icon} />} iconWidth="20px" />
          </div>
          {(badge > 0)?<div className="cell align-right" >
            <span className={color} >{badge}</span>
          </div>:null}
        </div>}
      />
  }

  const createSubList = (maplist) => {
    for(let index = 0; index < template[current.section].length; index++) {
      const etype = getElementType(template[current.section][index]);
      let item = template[current.section][index][etype];
      const mkey = "tmp_"+current.section+"_"+index.toString()+"_"+etype
      if (["row","datagrid"].includes(etype)) {
        item = item.columns;
      }
      if (current.parent===null) {
        const pinfo = setListIcon(item);
        maplist.push(<div key={mkey}>
          {mapButton(mkey, mkey, pinfo, etype.toUpperCase(), index+1, `${"primary"} ${styles.badgeBlack}`)}
        </div>)
      } else {
        if ((current.item===item) || (current.parent===item)) {
          const cinfo = setListIcon(item, index+1)
          maplist.push(<div key={mkey}>
            {mapButton(mkey, mkey, cinfo, etype.toUpperCase(), cinfo.badge, `${"primary"} ${styles.badgeBlack}`)}
          </div>)
          if (["row","datagrid"].includes(current.type) || ["row","datagrid"].includes(current.parent_type)) {
            for(let i2 = 0; i2 < item.length; i2++) {
              let subtype = getElementType(item[i2]);
              let subitem = item[i2][subtype];
              const skey = "tmp_"+current.section+"_"+index.toString()+"_"+etype+"_"+i2.toString()+"_"+subtype
              const sinfo = setListIcon(subitem);
              maplist.push(<div key={skey}>
                {mapButton(skey, skey, sinfo, subtype.toUpperCase(), sinfo.badge, `${"primary"} ${styles.badgeBlack}`)}
              </div>)
            }
          }
        }
      }
    };

  }

  const createMapList = () => {
    let maplist = [];
    ["report", "header", "details", "footer"].forEach(mkey => {
      const info = setListIcon(template[mkey]);
      if (info.selected && (mkey !== "report")) {
        maplist.push(<div key={"sep_"+mkey+"_0"} className={styles.separator} />)
      }
      maplist.push(<div key={mkey}>
        {mapButton("tmp_"+mkey, "tmp_"+mkey, info, mkey.toUpperCase(), info.badge, `${"warning"} ${styles.badgeOrange}`)}
      </div>)
      if (info.selected) {
        maplist.push(<div key={"sep_"+mkey+"_1"} className={styles.separator} />)
        if(mkey !== "report"){
          createSubList(maplist)
          maplist.push(<div key={"sep_"+mkey+"_2"} className={styles.separator} />)
        }
      }
    })
    return maplist;
  }

  const reportElements = {
    header: ["row", "vgap", "hline"],
    details: ["row", "vgap", "hline", "html", "datagrid"],
    footer: ["row", "vgap", "hline"],
    row: ["cell", "image", "barcode", "separator"],
    datagrid: ["column"]
  }

  const tableFields = () => {
    let fields = update(current_data.fields, {$merge: {
      delete: { columnDef: {
        id: "delete",
        Header: "",
        headerStyle: {},
        Cell: ({ row, value }) => {
          return (<div 
            className={`${"cell"} ${styles.deleteCol}`} >
            <Icon id={"delete_"+row.original["_index"]}
              iconKey="Times" width={19} height={27.6} 
              onClick={(event)=>{
                event.stopPropagation();
                onEvent("deleteDataItem",[{ _index: row.original._index }])
              }}
              className={styles.deleteCol} />
          </div>)
        },
        cellStyle: { width: 40, padding: "7px 8px 3px 8px" }
      }}
    }})
    return fields
  }

  const tabButton = (key, icon) => {
    return <Button id={`btn_${key}`} className={`${"full"} ${styles.tabButton} ${
      (tabView === key)?styles.selected:""} ${(tabView === key)?"primary":""}`} 
      onClick={()=>onEvent("changeTemplateData",[{key: "tabView", value: key}])}
      value={<Label value={getText(`template_label_${key}`)} 
        leftIcon={<Icon iconKey={icon} />} iconWidth="25px" />}
    />
  }

  const closeIcon = (key, event) => {
    return <Icon id={`btn_${key}`} iconKey="Times" onClick={()=>onEvent(event,[null])} />
  }

  const navButton = (key, event, label, icon, full, left) => {
    let labelProps = {
      value: getText(label), iconWidth: "20px"
    }
    if(left){
      labelProps.leftIcon= <Icon iconKey={icon} />
    } else {
      labelProps.rightIcon= <Icon iconKey={icon} />
    }
    return <Button id={`btn_${key}`} 
      onClick={() => onEvent(...event) } 
      className={`${styles.mapButton} ${"border-button"} ${full}`}
      value={<Label {...labelProps} />}
    />
  }

  const dataText = (key, value, rows, params) => {
    return <textarea id={`${key}_value`}
      className={`${"full"} ${styles.textareaStyle} `} 
      value={value} rows={rows}
      onChange={ (event)=>onEvent("editDataItem",[update({value: event.target.value}, {$merge: {...params}})]) } />
  }
  
  return (
    <div {...props} className={`${styles.width800} ${className}`}>
      <div className={`${styles.panel}`} >
        <div className={`${styles.panelTitle} ${"primary"}`}>
          <Label value={title} 
            leftIcon={<Icon iconKey="TextHeight" />} iconWidth="20px" />
        </div>
        <div className={`${"section"} ${styles.settingPanel}`} >
          <div className="row full container section-small-bottom" >
            <div className={`${styles.viewPanel}`} >
              <div className="row full" >
                <div className="cell third" >
                  {tabButton("template", "Tags")}
                </div>
                <div className={`${"cell third"}`} >
                  {tabButton("data", "Database")}
                </div>
                <div className={`${"cell third"}`} >
                  {tabButton("meta", "InfoCircle")}
                </div>
              </div>
              {(tabView === "template")?
              <Fragment >
                <div className="row full border section container-small" >
                  <div className="cell padding-small third" >
                    {navButton("previous", ["goPrevious",[]], "label_previous", "ArrowLeft", "full", true)}
                    <div key="map_box" className={`${"secondary-title border"} ${styles.mapBox}`} >
                      <canvas ref={mapRef} className={`${styles.reportMap}`} />
                    </div>
                    {navButton("next", ["goNext",[]], "label_next", "ArrowRight", "full", false)}
                  </div>
                  <div className="cell padding-small third" >
                    {createMapList()}
                  </div>
                  <div className="cell padding-small third" >
                    {(getMapCtr(current.type, "map_edit") !== false)?<div>
                      {navButton("move_up", ["moveUp",[]], "label_move_up", "ArrowUp", "full", true)}
                      {navButton("move_down", ["moveDown",[]], "label_move_down", "ArrowDown", "full", true)}
                      {navButton("delete_item", ["deleteItem",[]], "label_delete", "Times", "full", true)}
                      <div className={styles.separator} />
                    </div>:null}
                    {(getMapCtr(current.type, "map_insert"))?<div>
                      {navButton("add_item", ["addItem",[current.add_item||""]], "label_add_item", "Plus", "", true)}
                      <Select id="sel_add_item"
                        value={current.add_item||""} placeholder=""
                        onChange={(value)=>onEvent("changeCurrentData",[{key: "add_item", value: value}])}
                        options={reportElements[current.type].map(
                          (item)=>{ return { value: item, text: item.toUpperCase() }
                        })} />
                    </div>:null}
                  </div>
                </div>
                <div className={`${styles.title} ${"padding-small"}`}>
                  <Label value={current.type.toUpperCase()} 
                    leftIcon={<Icon iconKey="Tag" />} iconWidth="20px" />
                </div>
                {current.form.rows.map((row, index) =><div 
                  className={`${"row full border"} ${styles.templateRow}`} key={index} ><Row
                  row={row} 
                  values={(["row","datagrid"].includes(current.type)) ? current.item_base : current.item}
                  options={current.form.options}
                  data={{
                    audit: "all",
                    current: current,
                    dataset: template.data,
                  }}
                  getText={getText}
                  onEdit={(options)=>onEvent("editItem",[options])}
                /></div>)}
              </Fragment>:null}
              {(tabView === "data")?<div className="row full border padding-normal" >
              {(current_data && (current_data.type === "string"))?
                <div className="row full section-small">
                  <div className={`${styles.panelTitle} ${"primary"}`}>
                    <div className="row full">
                      <div className="cell">
                        <Label value={current_data.name} />
                      </div>
                      <div className={`${"cell align-right"} ${styles.closeIcon}`}>
                        {closeIcon("data_string", "setCurrentData")}
                      </div>
                    </div>
                  </div>
                  {dataText(current_data.name, template.data[current_data.name], 15, {})}
                </div>:null}
              {(current_data && (current_data.type === "list") && current_data.item)?
                <div className="row full section-small">
                  <div className={`${styles.panelTitle} ${"primary"}`}>
                    <div className="row full">
                      <div className="cell">
                        <Label value={current_data.item} />
                      </div>
                      <div className={`${"cell align-right"} ${styles.closeIcon}`}>
                        {closeIcon("data_list_item", "setCurrentDataItem")}
                      </div>
                    </div>
                  </div>
                  {dataText(current_data.item, template.data[current_data.name][current_data.item], 10, {})}
                </div>:null}
              {(current_data && (current_data.type === "table") && current_data.item)?
                <div className="row full section-small">
                  <div className={`${styles.panelTitle} ${"primary"}`}>
                    <div className="row full">
                      <div className="cell">
                        <Label value={current_data.name+" - "+String(current_data.item._index+1)} />
                      </div>
                      <div className={`${"cell align-right"} ${styles.closeIcon}`}>
                        {closeIcon("data_table_item", "setCurrentDataItem")}
                      </div>
                    </div>
                  </div>
                  {Object.keys(current_data.fields).map((field, index) => <div key={index} className="row full">
                    <div className="padding-small">
                      <Label className="bold" value={field} />
                    </div>
                    {dataText(field, current_data.item[field], 2, {field: field, _index: current_data.item._index})}
                  </div>)}
                </div>:null}
              {(current_data && (current_data.type === "list") && !current_data.item)?
                <div className="section-small" >
                  <div className={`${styles.panelTitle} ${"primary"}`}>
                    <div className="row full">
                      <div className="cell">
                        <Label value={current_data.name} />
                      </div>
                      <div className={`${"cell align-right"} ${styles.closeIcon}`}>
                        {closeIcon("data_list", "setCurrentData")}
                      </div>
                    </div>
                  </div>
                  <List rows={current_data.items}
                    listFilter={true} filterPlaceholder={getText("placeholder_filter")}
                    onAddItem={()=>onEvent("setCurrentDataItem",[undefined])} labelAdd={getText("label_new")}
                    paginationPage={paginationPage} paginationTop={true} 
                    onEdit={(row)=>onEvent("setCurrentDataItem",[row.lslabel])} 
                    onDelete={(row)=>onEvent("deleteDataItem",[{key: row.lslabel}])} 
                  />
                </div>:null}
                {(current_data && (current_data.type === "table") && !current_data.item)?
                <div className="section-small" >
                  <div className={`${styles.panelTitle} ${"primary"}`}>
                    <div className="row full">
                      <div className="cell">
                        <Label value={current_data.name} />
                      </div>
                      <div className={`${"cell align-right"} ${styles.closeIcon}`}>
                        {closeIcon("data_table", "setCurrentData")}
                      </div>
                    </div>
                  </div>
                  <Table rowKey="_index"
                    fields={tableFields()} rows={current_data.items} tableFilter={true}
                    filterPlaceholder={getText("placeholder_filter")}
                    onAddItem={()=>onEvent("setCurrentDataItem",[undefined])}
                    onRowSelected={(row)=>onEvent("setCurrentDataItem",[row])}
                    labelAdd={getText("label_new")} 
                    paginationPage={paginationPage} paginationTop={true}
                  />
                </div>:null}
              {(!current_data)?
                <div className="section-small" >
                  <List rows={dataset}
                    listFilter={true} filterPlaceholder={getText("placeholder_filter")}
                    onAddItem={()=>onEvent("addTemplateData",[])} labelAdd={getText("label_new")}
                    paginationPage={paginationPage} paginationTop={true} 
                    onEdit={(row)=>onEvent("setCurrentData",[{ name: row.lslabel, type: row.lsvalue }])} 
                    onDelete={(row)=>onEvent("deleteData",[row.lslabel])} 
                  />
                </div>:null}
              </div>:null}
              {(tabView === "meta")?<div className="row full border padding-normal" >
                <div className="cell section-small">
                  <div className="row centered border-top border-bottom mobile">
                    {Object.keys(template.meta).map(mkey => <div key={mkey} className="cell padding-small mobile top" >
                      <div className="small bold">{mkey}</div>
                      <div className="small">{template.meta[mkey]}</div>
                    </div>)}
                  </div>
                  <div className="section-small-top" >
                    <Label className="bold" value={getText("template_data_sources")} />
                  </div>
                  {Object.keys(template.sources).map(skey => 
                    <div key={skey} className="row section-small full" >
                      <div className="cell padding-small">
                        <div className="border-top border-bottom padding-small" >
                          <Label className="bold italic" value={skey} />
                        </div>
                        {Object.keys(template.sources[skey]).map(sql => 
                        <div key={sql} className="row">
                           <div className="cell padding-small top tiny bold">{sql}:</div>
                           <div className="cell padding-small tiny">{template.sources[skey][sql]}</div>
                        </div>)}
                      </div>
                    </div>)}
                </div>
              </div>:null}
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

TemplateEditor.propTypes = { 
  data: PropTypes.shape({
    title: PropTypes.string.isRequired, 
    tabView: PropTypes.string.isRequired, 
    template: PropTypes.shape({
      meta: PropTypes.object,
      report: PropTypes.object.isRequired,
      header: PropTypes.array,
      details: PropTypes.array.isRequired,
      footer: PropTypes.array,
      sources: PropTypes.object,
      data: PropTypes.object
    }).isRequired, 
    current: PropTypes.object.isRequired, 
    current_data: PropTypes.object, 
    dataset: PropTypes.array.isRequired
  }).isRequired, 
  paginationPage: PropTypes.number.isRequired,
  className: PropTypes.string,
  onEvent: PropTypes.func,
  getText: PropTypes.func,
}

TemplateEditor.defaultProps = {
  data: {
    title: "", 
    tabView: "template", 
    template: {
      meta: {},
      report: {},
      header: [],
      details: [],
      footer: [],
      sources: {},
      data: {}
    }, 
    current: {}, 
    current_data: null, 
    dataset: []
  },
  paginationPage: getSetting("paginationPage"),
  className: "",
  onEvent: undefined,
  getText: undefined,
}

export default TemplateEditor;