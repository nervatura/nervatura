import { useState } from 'react';
import PropTypes from 'prop-types';
import update from 'immutability-helper';

import styles from './Browser.module.css';
import { getSetting } from 'config/app'

import Icon from 'components/Form/Icon'
import Label from 'components/Form/Label'
import Button from 'components/Form/Button'
import Table from 'components/Form/Table';
import Select from 'components/Form/Select'
import Input from 'components/Form/Input'
import DateTime from 'components/Form/DateTime'

export const Browser = ({ 
  data, keyMap, viewDef, className,
  paginationPage, dateFormat, timeFormat, filter_opt_1, filter_opt_2,
  getText, onEvent,
  ...props 
}) => {
  const { vkey, view, show_header, show_dropdown, show_columns,
    result, columns, filters, deffield, page } = data
  const [ state, setState ] = useState({
    dropdown: show_dropdown,
    header: show_header,
    columns: show_columns
  })
  const checkTotalFields = (fields, deffield) => {
    let retval = { totalFields: {}, totalLabels: {}, count: 0 }
    if (deffield && Object.keys(fields).includes("deffield_value")) {
      deffield.filter((df)=>((df.fieldtype==="integer")||(df.fieldtype==="float")))
      .forEach((df) => {
         retval = update(retval, { 
           totalFields: { $merge: { [df.fieldname]: 0 }},
           totalLabels: { $merge: { [df.fieldname]: df.description }}
         })
       }
      )
    } else {
      Object.keys(fields).filter((fieldname)=>(
        ((fields[fieldname].fieldtype==="integer")||(fields[fieldname].fieldtype==="float"))
          &&(fields[fieldname].calc !== "avg")))
      .forEach((fieldname) => {
         retval = update(retval, { 
           totalFields: { $merge: { [fieldname]: 0 }},
           totalLabels: { $merge: { [fieldname]: fields[fieldname].label }}
         })
       }
      )
    }
    retval = update(retval, { $merge: {
      count: Object.keys(retval.totalFields).length
    }})
    return retval
  }
  let fields = {
    view: { columnDef: { 
      id: "view",
      Header: "",
      headerStyle: {},
      Cell: ({ row, value }) => {
        if(!viewDef.readonly){
          return <Icon id={"edit_"+row.original["id"]}
            iconKey="Edit" width={24} height={21.3} 
            onClick={ ()=>onEvent("onEdit",['id', row.original["id"], row.original]) }
            className={styles.editCol} />
        }
        return <Icon iconKey="CaretRight" width={9} height={24} />
      },
      cellStyle: { width: 25, padding: "7px 3px 3px 8px" }
    }}
  }
  Object.keys(viewDef.fields).forEach((fieldname) => {
    if(columns[view][fieldname]){
      switch (viewDef.fields[fieldname].fieldtype) {
        case "float":
        case "integer":
          fields = update(fields, {$merge: {
            [fieldname]: { fieldtype:'number', label: viewDef.fields[fieldname].label }
          }})
          break;

        case "bool":
          fields = update(fields, {$merge: {
            [fieldname]: { fieldtype:'bool', label: viewDef.fields[fieldname].label }
          }})
          break;

        case "string":
        default:
          if(fieldname === "deffield_value"){
            fields = update(fields, {$merge: {
              [fieldname]: { fieldtype:'deffield', label: viewDef.fields[fieldname].label }
            }})
          } else {
            fields = update(fields, {$merge: {
              [fieldname]: { fieldtype:'string', label: viewDef.fields[fieldname].label }
            }})
          }
          break;
      }
    }
  })
  const exportFields = () => {
    let fields = {}
    Object.keys(viewDef.fields).filter(fieldname => (columns[view][fieldname]===true)).forEach(fieldname => {
      fields[fieldname] = viewDef.fields[fieldname]
    })
    return fields
  }
  const totalFields = checkTotalFields(viewDef.fields, deffield)
  return (
    <div {...props} className={`${className}`} 
      onClick={ (event)=>{
        setState({ ...state, dropdown: false })
      }}>
      <div className={`${styles.panel}`} >
        <div className={`${styles.panelTitle} ${"primary"}`}>
          <Label value={getText("browser_"+vkey)} />
        </div>
        <div className="section container" >
          <div className="row full" >
            <div className="cell" >
              <Button id="btn_header"
                className={`${"full primary"}`} 
                onClick={(event)=>{
                  event.stopPropagation();
                  onEvent("changeData",["show_header", !state.header])
                  setState({ ...state, header: !state.header })
                }} 
                value={<Label value={viewDef.label} 
                  leftIcon={<Icon iconKey="Filter" />} />}
              />
            </div>
          </div>
          {(state.header)?<div className={`${styles.filterPanel} ${"border"}`} >
            <div className="row full" >
              <div className="cell top" >
                <Button id="btn_search" 
                  className={`${"border-button"} ${styles.barButton}`} 
                  onClick={ ()=>onEvent("browserView",[]) } 
                  value={<Label className="hide-small" value={getText("browser_search")} 
                    leftIcon={<Icon iconKey="Search" height={18} width={18} />} />}
                />
              </div>
              <div className="cell align-right" >
                <Button id="btn_bookmark"
                  className={`${"border-button small-button"} ${styles.barButton}`} 
                  onClick={()=>onEvent("bookmarkSave",[])} 
                  value={<Label value={getText("browser_bookmark")} 
                    leftIcon={<Icon iconKey="Star" height={14} width={15.75} />} />}
                />
                <Button id="btn_export"
                  className={`${"border-button small-button"} ${styles.barButton}`} 
                  disabled={(result.length === 0)?"disabled":""}
                  onClick={ ()=>onEvent("exportResult",[exportFields()]) }
                  value={<Label value={getText("browser_export")} 
                    leftIcon={<Icon iconKey="Download" height={14} width={14} />} />}
                />
                <Button id="btn_help"
                  className={`${"border-button small-button"} ${styles.barButton}`} 
                  onClick={()=>onEvent("showHelp",["program/browser"])} 
                  value={<Label value={getText("browser_help")} 
                    leftIcon={<Icon iconKey="QuestionCircle" height={14} width={14} />} />}
                />
              </div>
            </div>
            <div className="row full section-small-top" >
              <div className="cell" >
                <div className={`${styles.dropdownBox}`} >
                  <Button id="btn_views"
                    className={`${"border-button"} ${styles.barButton} ${(state.dropdown)?styles.selected:""}`} 
                    onClick={ (event)=>{
                      event.stopPropagation();
                      setState({ ...state, dropdown: !state.dropdown })
                    }} 
                    value={<Label className="hide-small" value={getText("browser_views")} 
                      leftIcon={<Icon iconKey="Eye" height={18} width={20.25} />} />}
                  />
                  {(state.dropdown)?<div className={`${styles.dropdownContent} ${"border dialog"}`} >
                    {Object.keys(keyMap).map((vname, index) => (vname !== "options")?
                      <div id={"view_"+vname}
                        key={index} onClick={ ()=>onEvent("showBrowser",[vkey, vname]) }
                        className={`${styles.dropItem} ${(vname === view)?styles.active:null}`} >
                        <Label value={keyMap[vname].label} 
                          leftIcon={(vname === view)?<Icon iconKey="Check" />:<Icon iconKey="Eye" />} />
                      </div>:null
                    )}
                  </div>:null}
                </div>
                <Button id="btn_columns"
                  className={`${"border-button"} ${styles.barButton}`}
                  onClick={ (event)=>{
                    event.stopPropagation();
                    onEvent("changeData",["show_columns", !state.columns])
                    setState({ ...state, columns: !state.columns })
                  }} 
                  value={<Label className="hide-small" value={getText("browser_columns")} 
                    leftIcon={<Icon iconKey="Columns" height={18} width={18} />} />}
                />
                <Button id="btn_filter"
                  className={`${"border-button"} ${styles.barButton}`} 
                  onClick={()=>onEvent("addFilter",[])} 
                  value={<Label className="hide-small" value={getText("browser_filter")} 
                    leftIcon={<Icon iconKey="Plus" height={18} width={15.75} />} />}
                />
                <Button id="btn_total"
                  className={`${"border-button"} ${styles.barButton}`} 
                  disabled={((totalFields.count === 0)||(result.length === 0))?"disabled":""}
                  onClick={ ()=>onEvent("showTotal",[viewDef.fields, totalFields]) } 
                  value={<Label className="hide-small" value={getText("browser_total")} 
                    leftIcon={<Icon iconKey="InfoCircle" height={18} width={18} />} />}
                />
              </div>
            </div>
            {(state.columns)?<div className={`${styles.colBox} ${"border"}`} >
              {Object.keys(viewDef.fields).map(fieldname =>
                <div id={"col_"+fieldname} key={fieldname} className={`${"cell padding-tiny tiny left"} 
                  ${(columns[view][fieldname]===true)?styles.selectCol:styles.editCol}`}
                  onClick={()=>onEvent("setColumns",[fieldname, !(columns[view][fieldname]===true)])} >
                  <Label value={viewDef.fields[fieldname].label} 
                    leftIcon={(columns[view][fieldname]===true)?<Icon iconKey="CheckSquare" />:<Icon iconKey="SquareEmpty" />} />
                </div>
              )}              
            </div>:null}
            {filters[view].map((filter, index) => <div key={index} className="section-small-top" >
              <div className="cell" >
                <Select id={"filter_name_"+index}
                  value={filter.fieldname} 
                  onChange={(value)=>onEvent("editFilter",[index, "fieldname", value]) }
                  options={Object.keys(viewDef.fields).filter(
                    (fieldname)=> (fieldname !== "id") && (fieldname !== "_id")
                  ).flatMap((fieldname) => {
                    if(fieldname === "deffield_value"){
                      return deffield.map((df) => {
                        return { value: df.fieldname, text: getText(df.fieldname, df.description) }
                      })
                    }
                    return { value: fieldname, text: viewDef.fields[fieldname].label } 
                  })} />
              </div>
              <div className="cell" >
                <Select id={"filter_type_"+index}
                  value={filter.filtertype} 
                  onChange={(value)=>onEvent("editFilter",[index, "filtertype", value]) }
                  options={(["date","float","integer"].includes(filter.fieldtype)?filter_opt_2:filter_opt_1).map(
                    (item)=>{ return { value: item[0], text: item[1] }
                  })} />
              </div>
              <div className="cell mobile" >
                {(filter.filtertype !== "==N")?<div className="cell" >
                  {(filter.fieldtype === "bool")?<Select id={"filter_value_"+index}
                    value={filter.value} 
                    onChange={(value)=>onEvent("editFilter",[index, "value", value]) }
                    options={[
                      { value: "0", text: getText("label_no") }, 
                      { value: "1", text: getText("label_yes") }
                    ]} />:null}
                  {((filter.fieldtype === "integer")||(filter.fieldtype === "float"))
                    ?<Input id={"filter_value_"+index} value={filter.value} 
                      onChange={(value)=>onEvent("editFilter",[index, "value", value]) }
                      type={(filter.fieldtype === "float") ? "number" : filter.fieldtype} 
                      className="align-right" />:null}
                  {(filter.fieldtype === "date")?<DateTime id={"filter_value_"+index}
                    value={filter.value} 
                    dateTime={false} isEmpty={false}
                    onChange={(value)=>onEvent("editFilter",[index, "value", value]) } />:null}
                  {(filter.fieldtype === "string")
                    ?<Input id={"filter_value_"+index} value={filter.value} 
                       onChange={(value)=>onEvent("editFilter",[index, "value", value]) } />:null}
                </div>:null}
                <div className="cell" > 
                  <Button id={"btn_delete_filter_"+index}
                    className={` ${"border-button"} ${styles.filterDelete}`} 
                    onClick={ ()=>onEvent("deleteFilter",[index]) } 
                    value={<Icon iconKey="Times" />}
                  />
                </div>
              </div>
            </div>)}
          </div>:null}
          <div className="row full section-small-top" >
            <div className={`${"row full border"}`} >
              <div className={`${"cell"} ${styles.resultTitle}`} >
                {result.length} <Label value={getText("browser_result")} />
              </div>
              {(viewDef.actions_new)?<div className={`${"cell"} ${styles.resultTitlePlus}`}>
                <Button id="btn_actions_new"
                  className={`${"small-button"}`} 
                  onClick={ ()=>onEvent("setFormActions",[{ params: viewDef.actions_new, row: undefined }]) } 
                  value={<Icon iconKey="Plus" />}
                />
              </div>:null}
            </div>
          </div>
          <div className="row full" >
            <Table rowKey="row_id"
              fields={fields} rows={result}
              filterPlaceholder={getText("placeholder_filter")}
              labelYes={getText("label_yes")} labelNo={getText("label_no")}
              dateFormat={dateFormat} timeFormat={timeFormat} 
              paginationPage={paginationPage} paginationTop={true}
              onEditCell={(fieldname, value, row)=>onEvent("onEdit",[fieldname, value, row])} 
              currentPage={page} onCurrentPage={(page)=>onEvent("onPage",[page])} 
              />
          </div>
        </div>
      </div>
    </div>
  )
}

Browser.propTypes = {
  data: PropTypes.shape({
    vkey: PropTypes.string.isRequired,
    view: PropTypes.string.isRequired,
    show_header: PropTypes.bool.isRequired, 
    show_dropdown: PropTypes.bool.isRequired,
    show_columns: PropTypes.bool.isRequired,
    result: PropTypes.array.isRequired,
    columns: PropTypes.object.isRequired,
    filters: PropTypes.object.isRequired,
    deffield: PropTypes.array.isRequired, 
    page: PropTypes.number.isRequired,
  }).isRequired,
  keyMap: PropTypes.object.isRequired,
  viewDef: PropTypes.shape({
    fields: PropTypes.object.isRequired,
    label: PropTypes.string.isRequired,
    actions_new: PropTypes.object,
    readonly: PropTypes.bool
  }).isRequired,
  paginationPage: PropTypes.number.isRequired, 
  dateFormat: PropTypes.string, 
  timeFormat: PropTypes.string, 
  filter_opt_1: PropTypes.array, 
  filter_opt_2: PropTypes.array,
  className: PropTypes.string,
  onEvent: PropTypes.func,
  getText: PropTypes.func,
}

Browser.defaultProps = {
  data: {
    vkey: "",
    view: "",
    show_header: true, 
    show_dropdown: false,
    show_columns: false,
    result: [],
    columns: {},
    filters: {},
    deffield: [], 
    page: 0,
  },
  keyMap: {},
  viewDef: {
    fields: {},
    label: "",
    readonly: false
  },
  paginationPage: getSetting("paginationPage"),
  dateFormat: getSetting("dateFormat"),
  timeFormat: getSetting("timeFormat"),
  filter_opt_1: getSetting("filter_opt_1"), 
  filter_opt_2: getSetting("filter_opt_2"),
  className: "",
  onEvent: undefined,
  getText: undefined,
}

export default Browser;