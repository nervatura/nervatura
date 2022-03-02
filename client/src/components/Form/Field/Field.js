import { useState } from 'react';
import PropTypes from 'prop-types';
import update from 'immutability-helper';
import formatISO from 'date-fns/formatISO'
import isValid from 'date-fns/isValid'
import parseISO from 'date-fns/parseISO'
import format from 'date-fns/format'

import Icon from 'components/Form/Icon'
import DateTime from 'components/Form/DateTime'
import Label from 'components/Form/Label'
import Button from 'components/Form/Button'
import Select from 'components/Form/Select'
import { Input, INPUT_TYPE } from 'components/Form/Input/Input'

import styles from './Field.module.css';
import { getSetting } from 'config/app'

export const Field = ({ 
  field, values, options, data,
  className, getText, 
  onEdit, onEvent, onSelector,
  ...props 
}) => {
  const [ state, setState ] = useState({
    event_type: undefined,
    selector: undefined
  })
  let { datatype, description } = field
  let disabled = (field.disabled || data.audit === 'readonly')
  let fieldName = field.name;
  let value = values[field.name]
  let fieldMap = field.map || null
  const editName = (fieldMap)
    ? (fieldMap.extend && fieldMap.text) 
      ? fieldMap.text 
      : (fieldMap.fieldname) 
        ? fieldMap.fieldname : fieldName
    : fieldName
  const empty = ((field.empty === "true") || (field.empty === true)) ? true : false

  const onChange = ( {value, item, event_type} ) => {
    onEdit({
      id: field.id || 1,
      name: editName,
      event_type: event_type,
      value: value, 
      extend: (fieldMap && fieldMap.extend) ? true : false,
      refnumber: (item && item.label) ? item.label : field.link_label,
      label_field: (fieldMap) ? fieldMap.label_field : undefined,
      item: item
    })
  }

  const selectorInit = () => {
    let selector = update({}, {$set: {
      value: values.value || "",
      filter: description || "",
      text: description || "",
      type: field.fieldtype,
      ntype: (["transitem", "transmovement", "transpayment"].includes(field.fieldtype)) ? "trans" : field.fieldtype,
      ttype: null,
      id: values.value || null,
      table: { name:"fieldvalue", fieldname: "value", id: field.id }
    }})
    if (fieldMap) {
      selector = update(selector, {$merge: {
        value: value || "",
        type: fieldMap.seltype, 
        filter: String(selector.text).split(" | ")[0],
        table:{ name: fieldMap.table, fieldname: fieldMap.fieldname },
        ntype: fieldMap.lnktype, 
        ttype: (value !== "") 
          ? (fieldMap.transtype === "") 
            ? values.transtype 
            : fieldMap.transtype 
          : selector.ttype, 
        id: (value !== "") ? value : selector.id
      }})
      let reftable;
      if (fieldMap.extend === true || fieldMap.table === "extend") {
        reftable = data.current.extend
      } else {
        if((typeof data.current[fieldMap.table] !== "undefined") && Array.isArray(data.current[fieldMap.table])){
          reftable = data.current[fieldMap.table].filter(item => (item.id === values.id))[0]
        } else {
          reftable = data.dataset[fieldMap.table].filter(item => (item.id === values.id))[0]
        }
      }
      /*      
      if (typeof reftable === "undefined" && data.current[fieldMap.table] && Array.isArray(data.current[fieldMap.table])) {
        reftable = data.current[fieldMap.table].filter(item => 
          (item[fieldMap.fieldname] === selector.value))[0]
      }
      if (typeof reftable === "undefined") {
        reftable = data.dataset[fieldMap.table].filter(item => 
          (item[fieldMap.fieldname] === selector.value))[0]
      }
      if (typeof reftable === "undefined" && data.current[fieldMap.table] && Array.isArray(data.current[fieldMap.table])) {
        reftable = data.current[fieldMap.table].filter(item => 
          (item.id === selector.value))[0]
      }
      if (typeof reftable === "undefined") {
        reftable = data.dataset[fieldMap.table].filter(item => 
          (item.id === selector.value))[0]
      }
      */
      if (typeof reftable !== "undefined") {
        if (typeof values[fieldMap.label_field] !== "undefined") {
          selector = update(selector, {$merge: {
            text: values[fieldMap.label_field]||"", 
            filter: values[fieldMap.label_field]||""
          }})
        } else if (typeof reftable[fieldMap.label_field] !== "undefined" && 
          reftable[fieldMap.label_field] !== null) {
            selector = update(selector, {$merge: {
              text: reftable[fieldMap.label_field], 
              filter: reftable[fieldMap.label_field]
            }})
        } else {
          selector = update(selector, {$merge: {
            text: "", 
            filter: ""
          }})
        }
        if (typeof reftable[fieldMap.fieldname] !== "undefined" && 
          selector.value==="" && reftable[fieldMap.fieldname] !== null) {
            selector = update(selector, {$merge: {
              ntype: fieldMap.lnktype, 
              ttype: fieldMap.transtype, 
              id: reftable[fieldMap.fieldname]
            }})
        }
        if (fieldMap.lnktype === "trans" && typeof reftable.transtype !== "undefined") {
          if (typeof fieldMap.lnkid !== "undefined") {
            selector = update(selector, {$merge: { 
              id: reftable[fieldMap.lnkid]
            }})
          } else if (typeof reftable[fieldMap.fieldname] !== "undefined") {
            selector = update(selector, {$merge: { 
              id: reftable[fieldMap.fieldname]
            }})
          } 
          else {
            selector = update(selector, {$merge: { 
              id: selector.value
            }})
          }
        }
      } else {
        selector = update(selector, {$merge: {
          text: "", 
          filter: ""
        }})
      }
    }
    setState({...state, selector: selector })
    return selector
  }

  const setSelector = (row, filter) => {
    let selector = update(state.selector, {$merge: {
      text: "",
      id: null,
      filter: filter || ""
    }})
    if (row){
      const params = row.id.split("/")
      selector = update(selector, {$merge: {
        text: row.label || row.item.lslabel
      }})
      selector = update(selector, { $merge: {
        id: parseInt(params[2],10),
        ttype: params[1]
      }})
      if((params[0] === "trans") && (params[1] !== "")){
        if(row.trans_id){
          selector = update(selector, { $merge: {
            id: row.trans_id
          }})
        }
      }
    }
    selector = update(selector, {$merge: {
      value: selector.id || ""
    }})
    
    setState({...state, selector: selector })
    onChange({value: selector.id, item: row, event_type: "change"})
  }

  const getOppositeValue = (value) => {
    if (options.opposite && (parseFloat(value)<0)) {
      return String(value).replace("-","");
    } else if (options.opposite && (parseFloat(value)>0)) {
      return "-"+value;
    }
    return value;  
  }

  const lnkValue = () => {
    if (typeof values[field.name] === "undefined") {
      return [(data.current[fieldMap.source]) ?
        data.current[fieldMap.source].filter(item => 
          (item.ref_id === values.id) && (item[fieldMap.value] === field.name))[0] :
        data.dataset[fieldMap.source].filter(item => 
          (item.ref_id === values.id) && (item[fieldMap.value] === field.name))[0], false]
    } else {
      const svalue = ((fieldName === "id") && (value === "")) ? null : value
      return [(data.current[fieldMap.source]) ?
        data.current[fieldMap.source].filter(item => 
          (item[fieldMap.value] === svalue))[0] :
        data.dataset[fieldMap.source].filter(item => 
          (item[fieldMap.value] === svalue))[0], true]
    }
  }

  if((field.rowtype === "reportfield") || (field.rowtype === "fieldvalue")){
    value = values.value
  }
  if ((typeof value==="undefined") || value === null){
    value = (field.default) ? field.default : ""
  }
  if (datatype === "fieldvalue"){
    datatype = values.datatype
    /*
    if(values.datatype){
      datatype = values.datatype 
    } else {
      if (fieldMap) {
        const mitem = data.dataset[fieldMap.source].filter((field)=>(
          field[fieldMap.value] === values.id))[0]
        datatype = mitem.fieldtype
        fieldMap = null
        if (mitem.valuelist !== null) {
          description = mitem.valuelist.split("|");
        }
      } 
    }
    */
  }
  switch (datatype) {
    case "password":
    case "color":
      return <Input {...props} 
        className={`${className} ${"full"}`} 
        name={fieldName} type={datatype} value={value||""} 
        onChange={(value) => onChange({value: value, event_type: "change"})}
        disabled={(disabled) ? 'disabled' : ''}/>

    case "date":
    case "datetime":
      let dateValue = parseISO(value)
      if (fieldMap) {
        if (fieldMap.extend) {
          dateValue = parseISO(data.current.extend[fieldMap.text])
          fieldName = fieldMap.text;
        } else {
          const lnkDate = lnkValue()
          if (typeof lnkDate[0] !== "undefined") {
            dateValue = parseISO(lnkDate[0][fieldMap.text])
            disabled = (lnkDate[1]) ? lnkDate[1] : disabled
          }
        }
      }
      value = isValid(dateValue) ? formatISO(dateValue) : ""
      return <DateTime {...props} 
        className={className} value={value} 
        dateTime={(datatype === "datetime")}
        isEmpty={empty} disabled={(disabled)}
        onChange={(value) => {
          if(value && (datatype === "datetime")){
            onChange({
              value: format(parseISO(value), getSetting("dateFormat")+" "+getSetting("timeFormat")),
              event_type: "change"
            })
          } else {
            onChange({value: value, event_type: "change"})
          }
        }}
        dateFormat={getSetting("dateFormat")}
        timeFormat={getSetting("timeFormat")}
        locale={getSetting("calendar")} />

    case "bool":
    case "flip":
      const toggleDisabled = (disabled)?styles.toggleDisabled:""
      if([1,"1","true",true].includes(value)){
        return <div {...props}
          className={` ${className} ${"toggle"} ${styles.toggle} ${toggleDisabled}`}
          onClick={(!disabled)?
            ()=>onChange({
              value: (field.name === 'fieldvalue_value') ? false : 0,
              event_type: "change"
            }):null}>
          <Icon iconKey="ToggleOn" className={`${styles.toggleOn}`} width={40} height={32.6} />
        </div>
      } else {
        return <div {...props}
          className={` ${className} ${"toggle"} ${styles.toggle} ${toggleDisabled}`}
          onClick={(!disabled)?
            ()=>onChange({
              value: (field.name === 'fieldvalue_value') ? true : 1,
              event_type: "change"
            }):null}>
          <Icon iconKey="ToggleOff" className={`${styles.toggleOff}`} width={40} height={32.6} />
        </div>
      }
    
    case "label":
      return null;

    case "select":
      if (field.extend) {
        value = data.current.extend[field.name]||"";
      }
      let selectOptions = []
      if (fieldMap) {
        data.dataset[fieldMap.source].forEach((element, index) => {
          let _label = element[fieldMap.text]
          if (typeof fieldMap.label !== "undefined") {
            _label = getText(fieldMap.label+"_"+_label);
          }
          selectOptions.push({ value: String(element[fieldMap.value]), text: _label })
        });
      } else {
        field.options.forEach((element, index) => {
          let _label = element[1]
          if(getText(_label)){
            _label = getText(_label)
          }
          if (typeof field.olabel !== "undefined") {
            _label = getText(field.olabel+"_"+element[1]);
          }
          selectOptions.push({ value: String(element[0]), text: _label })
        });
      }
      return <Select {...props} className={`${className} ${"full"}`} 
        name={field.name} value={value} placeholder={(empty)?"":undefined}
        disabled={(disabled) ? 'disabled' : ''}
        onChange={(value) => {
          let _value = isNaN(parseInt(value,10)) ?
            value : parseInt(value,10)
          onChange({value: _value, event_type: "change"})
        }}
        options={selectOptions} />
    
    case "valuelist":
      return <Select {...props} className={`${className} ${"full"}`} 
        name={field.name} value={value}
        disabled={(disabled) ? 'disabled' : ''}
        onChange={(value) => onChange({value: value, event_type: "change"})}
        options={description.map((value, index) => {
          return { value: value, text: value }
        })}
        />
    
    case "link":
      let litem = values;
      const lnkLink = lnkValue()
      if (typeof lnkLink[0] !== "undefined") {
        litem = lnkLink[0]
        if(lnkLink[0][fieldMap.text]){
          value = lnkLink[0][fieldMap.text];
        }
      }
      let llabel = value;
      if (typeof fieldMap.label_field !== "undefined") {
        if (typeof litem[fieldMap.label_field] !== "undefined") {
          llabel = litem[fieldMap.label_field];
        }
      }
      return <div {...props} 
        className={`${className} ${"link"} ${styles.link}`} >
        <span id={"link_"+fieldMap.lnktype+"_"+fieldName} className={`${styles.lnkText}`} 
          onClick={()=>onEvent("checkEditor", [
            { ntype: fieldMap.lnktype, 
              ttype: fieldMap.transtype, 
              id: value 
            }, 
            "LOAD_EDITOR", undefined
          ])
          } >{llabel}</span>
      </div>
    
    case "selector":
      let columns = []
      let selector = state.selector
      const reInit = (fieldMap && (fieldMap.extend === true || fieldMap.table === "extend") 
        && data.current.extend && data.current.extend.seltype)
      if(!selector){
        selector = selectorInit()
      } else if(reInit && (selector.type !== data.current.extend.seltype)) {
        selector.text = data.current.extend[fieldMap.label_field]
        selector.type = data.current.extend.seltype
        selector.filter = selector.text
        selector.ntype = data.current.extend.seltype
        selector.ttype = data.current.extend.transtype
        selector.id = data.current.extend.ref_id
      }
      if(!disabled){
        columns.push(<div key="sel_show" className={` ${"cell"} ${styles.searchCol}`}>
          <Button id={"sel_show_"+fieldName}
            className={`${"border-button"} ${styles.selectorButton}`}
            onClick={()=>onSelector(selector.type, selector.filter, setSelector)}
            value={<Icon iconKey="Search" />}
          />
        </div>)
      }
      if (empty) {
        columns.push(<div key="sel_delete" className={` ${"cell"} ${styles.timesCol}`}>
          <Button id={"sel_delete_"+fieldName}
            className={`${"border-button"} ${styles.selectorButton}`}
            disabled={(disabled) ? 'disabled' : ''}
            onClick={ ()=>setSelector() }
            value={<Icon iconKey="Times" />}
          />
        </div>)
      }
      columns.push(<div key="sel_text" className={`${"link"} ${styles.link}`}>
        {(selector.text !== "")?<span id={"sel_link_"+fieldName}
          className={`${styles.lnkText}`}
          onClick={()=>onEvent("checkTranstype", [
            { ntype: selector.ntype, 
              ttype: selector.ttype, 
              id: selector.id }, 
            'LOAD_EDITOR', undefined
          ])} >{selector.text}</span>:null}
      </div>)
      return <div {...props} 
        className={`${className} ${"row full"}`} >{columns}</div>
    
    case "button":
      return <Button {...props} 
        className={`${className} ${"border-button"} ${styles.selectorButton} ${field.class}`}
        disabled={(disabled) ? 'disabled' : ''}
        autoFocus={field.focus || false}
        onClick={ ()=>onChange({value: fieldName, item: {}, event_type: "click"}) }
        value={<Label value={(field.title)?(field.title):""} 
          leftIcon={(field.icon)?<Icon iconKey={field.icon} />:null} 
          iconWidth={(field.icon)?"20px":null} />}
      />

    case "percent":
    case "integer":
    case "float":
      if (fieldMap) {
        if (fieldMap.extend) {
          value = data.current.extend[fieldMap.text];
          fieldName = fieldMap.text;
        } else {
          const lnkNumber = lnkValue()
          if (typeof lnkNumber[0] !== "undefined") {
            value = lnkNumber[0][fieldMap.text];
            disabled = (lnkNumber[1]) ? lnkNumber[1] : disabled
          }
        }
      }
      if (value === ""){ value = 0 }
      if (typeof field.opposite !== "undefined") {
        value = getOppositeValue(value) 
      }
      return <Input {...props} className={`${className} ${"full"}`}
        name={fieldName} value={value||"0"} 
        minValue={field.min} maxValue={field.max}
        type={(datatype === "float")?INPUT_TYPE.NUMBER:INPUT_TYPE.INTEGER}  
        onChange={(value) => {
          onChange({
            value: (field.opposite) ? getOppositeValue(value): value, 
            event_type: "change"
          })
          setState({...state, event_type: "change" })
        }}
        onBlur={(value) => {
          onChange({
            value: (field.opposite) ? parseFloat(getOppositeValue(value)): parseFloat(value),
            event_type: (state.event_type === "change") ? "blur" : null
          })
          setState({...state, event_type: "blur" })
        }}
        disabled={(disabled) ? 'disabled' : ''}/>
    
    case "notes":
    case "text":
    case "string":
    default:
      if (fieldMap) {
        if (fieldMap.extend) {
          value = data.current.extend[fieldMap.text];
          fieldName = fieldMap.text;
        } else {
          const lnkString = lnkValue()
          if (typeof lnkString[0] !== "undefined") {
            value = lnkString[0][fieldMap.text];
            disabled = (lnkString[1]) ? lnkString[1] : disabled
            if (typeof fieldMap.label !== "undefined") {
              value = getText(fieldMap.label+"_"+value);
            }
          }
        }
      }
      if((datatype === "notes") || datatype === "text"){
        return <textarea {...props} className={`${className} ${"full"} ${styles.textareaStyle}`} 
          name={fieldName} value={value||""}
          rows={(field.rows )?field.rows:null}
          onChange={(event) => onChange({value: event.target.value, event_type: "change"})}
          disabled={(disabled) ? 'disabled' : ''}/>
      }
      return <Input {...props} className={`${className} ${"full"}`} 
        name={fieldName} type="text" value={value||""} 
        maxLength={(field.length)?field.length:null}
        size={(field.length)?field.length:null}
        onChange={(value) => onChange({value: value, event_type: "change"})}
        disabled={(disabled) ? 'disabled' : ''}/>
  }
}

Field.propTypes = {
  field: PropTypes.object.isRequired,
  values: PropTypes.object.isRequired,
  options: PropTypes.object.isRequired,
  data: PropTypes.shape({
    dataset: PropTypes.object, 
    current: PropTypes.object, 
    audit: PropTypes.string.isRequired, 
  }).isRequired,
  className: PropTypes.string,
  getText: PropTypes.func,
  onEdit: PropTypes.func,
  onEvent: PropTypes.func,
  onSelector: PropTypes.func,
}

Field.defaultProps = {
  field: {},
  values: {},
  options: {},
  data: {
    dataset: {}, 
    current: {}, 
    audit: "all",
  },
  className: "",
  getText: undefined,
  onEdit: undefined,
  onEvent: undefined,
  onSelector: undefined,
}

export default Field;