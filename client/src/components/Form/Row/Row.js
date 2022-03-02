import PropTypes from 'prop-types';

import Icon from 'components/Form/Icon'
import Label from 'components/Form/Label'
import Field from 'components/Form/Field'
import Input from 'components/Form/Input'

import styles from './Row.module.css';

export const Row = ({ 
  row, values, options, data,
  className, getText, 
  onEdit, onEvent, onSelector,
  ...props 
}) => {
  const { id, rowtype, label, columns, name, disabled, notes, selected, empty, 
    datatype, info } = row
  const imgValue = () => {
    let img_value = values[name] || ""
    if (img_value!=="" && img_value!==null) {
      if (img_value.toString().substr(0,10)!=="data:image") {
        if (typeof data.dataset[img_value]!=="undefined") {
          img_value = data.dataset[img_value]
        }
      }
    }
    return img_value
  }
  switch (rowtype) {

    case "label":
      return (<div {...props} 
          className={`${className} ${"row full padding-small section-small border-bottom"} ${styles.labelRow}`}
        >
        <div className="cell padding-small" >{values[name] || label}</div>
      </div>)

    case "flip":
      const enabled = (typeof values[name] !== "undefined")
      const checkbox = <div id={"checkbox_"+name}
        className={` ${styles.reportField}`}
        onClick={(event) => onEdit({
          id: id,
          selected: true,
          datatype: datatype,
          defvalue: row.default,
          name: name, 
          value: !enabled, 
          extend: false
        })}>
        {(enabled)?
          <Icon iconKey="ToggleOn" className={`${styles.toggleOn}`} width={40} height={32.6} />:
          <Icon iconKey="ToggleOff" className={`${styles.toggleOff}`} width={40} height={32.6} />}
        <Label className={`${"bold padding-tiny"} ${(enabled)?styles.toggleOn:""}`} value={name} />
      </div>

      switch (datatype) {
        case "text":
          return(<div {...props} 
              className={`${className} ${"row full padding-small section-small border-bottom"}`}
            >
            <div className="row full">
              <div className={`${"cell padding-small"}`} >
                {checkbox}
              </div>
            </div>
            {(enabled)?<div className="row full"><div className={`${"cell padding-small"}`} >
              <Field id={"field_"+name}
                field={row} values={values} options={options} data={data}
                getText={getText} onEdit={onEdit} onEvent={onEvent} onSelector={onSelector} />
            </div></div>:null}
            {(info)?<div className="row full padding-small">
              <div className={`${"cell padding-small info"} ${styles.leftbar}`} >
                {info}
              </div>
            </div>:null}
          </div>)
        
        case "image":
          return(
            <div {...props} 
              className={`${className} ${"row full padding-small section-small border-bottom"}`}
            >
            <div className="row full">
              <div className={`${"cell padding-small"}`} >
                {checkbox}
              </div>
              {(enabled)?<div className={`${"cell padding-small"}`} >
                <input id={"file_"+name} 
                  className={`${"full small"} ${styles.inputStyle} `}
                  type="file"
                  onChange={(event) => onEdit({
                    id: id,
                    file: true,
                    name: name, 
                    value: event.target.files, 
                    extend: false
                  })} />
              </div>:null}
            </div>
            {(enabled)?<div className="row full"><div className={`${"cell padding-small"}`} >
              <textarea id={"input_"+name}
                className={`${"full small"} ${styles.textareaStyle}`} 
                value={imgValue()} rows={5}
                onChange={(event) => onEdit({
                  id: id, 
                  name: name, 
                  value: event.target.value
                })} />
              <div className="full padding-normal center" >
                <img src={imgValue()} alt="" />
              </div>
            </div></div>:null}
            {(info)?<div className="row full padding-small">
              <div className={`${"cell padding-small info"} ${styles.leftbar}`} >
                {info}
              </div>
            </div>:null}
          </div>)

        case "checklist":
          let cb_value = values[name] || ""
          let checklist = []
          row.values.forEach((element, index) => {
            let cvalue = element.split("|")
            const value = (cb_value.indexOf(cvalue[0])>-1) ? true : false
            checklist.push(<div id={"checklist_"+name+"_"+index}
              key={index}
              className={` ${"cell padding-small"} ${styles.reportField}`}
              onClick={(event) => onEdit({
                id: id,
                checklist: true,
                name: name,
                checked: !value,
                value: cvalue[0],
                extend: false
              })}>
              <Label className={`${"bold"} ${(value)?styles.toggleOn:""}`} value={cvalue[1]} 
                leftIcon={(value)
                  ?<Icon iconKey="CheckSquare" className={`${styles.toggleOn}`}  />
                  :<Icon iconKey="SquareEmpty" />} />
            </div>)
          });
          return(
            <div {...props} 
              className={`${className} ${"row full padding-small section-small border-bottom"}`}
            >
            <div className="row full">
              <div className={`${"cell padding-small"}`} >
                {checkbox}
              </div>
            </div>
            {(enabled)?<div className="row full padding-small">
              <div className={`${"cell padding-small toggle"} ${styles.toggle}`} >
                {checklist}
              </div>
            </div>:null}
            {(info)?<div className="row full padding-small">
              <div className={`${"cell padding-small info"} ${styles.leftbar}`} >
                {info}
              </div>
            </div>:null}
          </div>)
      
        default:
          return(<div {...props} 
            className={`${className} ${"row full padding-small section-small border-bottom"}`}>
            <div className="row full">
              <div className={`${"cell padding-small half"}`} >
                {checkbox}
              </div>
              {(enabled)?<div className={`${"cell padding-small half"}`} >
                <Field id={"field_"+name}
                  field={row} values={values} options={options} data={data}
                  getText={getText} onEdit={onEdit} onEvent={onEvent} onSelector={onSelector} />
              </div>:null}
            </div>
            {(info)?<div className="row full padding-small">
              <div className={`${"cell padding-small info"} ${styles.leftbar}`} >
                {info}
              </div>
            </div>:null}
          </div>)
      }

    case "field":
      return(<div {...props} 
        className={`${className} ${"row full padding-small section-small border-bottom"}`}>
        <div className={`${"cell padding-small hide-small"} ${styles.fieldCell}`} >
          <Label className="bold" value={label} />
        </div>
        <div className={`${"cell padding-small"}`} >
          <div className={`${"hide-medium hide-large"}`} >
            <Label className="bold" value={label} />
          </div>
          <Field id={"field_"+name}
            field={row} values={values} options={options} data={data}
            getText={getText} onEdit={onEdit} onEvent={onEvent} onSelector={onSelector} />
        </div>
      </div>)
    
    case "reportfield":
      return(<div {...props} 
        className={`${className} ${"cell padding-small s12 m6 l4"}`}>
        <div id={"cb_"+name}
          className={`${"padding-small"} ${(empty !== 'false')?styles.reportField:""}`} 
          onClick={() => {if(empty !== 'false'){
            onEdit({id: id, name: "selected", value: !selected, extend: false })} }}>
          <Label className="bold" value={label} 
            leftIcon={(selected)?<Icon iconKey="CheckSquare" />:<Icon iconKey="SquareEmpty" />} />
        </div>
        <Field id={"field_"+name}
          field={row} values={values} options={options} data={data}
          getText={getText} onEdit={onEdit} onEvent={onEvent} onSelector={onSelector} />
      </div>)

    case "fieldvalue":
      return(<div {...props} 
        className={`${className} ${"row full padding-small section-small border-bottom"}`}>
        <div className="row full">
          <div className="cell container-small">
            <Label className="bold" value={label} />
          </div>
          <div className="cell align-right container-small" >
            <span id={"delete_"+row.fieldname}
              className={`${styles.fieldvalueDelete}`} 
              onClick={ ()=>onEdit({ 
                id: id, name: "fieldvalue_deleted"}) }><Icon iconKey="Times" /></span>
          </div>
        </div>
        <div className="row full">
          <div className={`${"cell padding-small s12 m6 l6"}`} >
            <Field id={"field_"+row.fieldname}
              field={row} values={values} options={options} data={data}
              getText={getText} onEdit={onEdit} onEvent={onEvent} onSelector={onSelector} />
          </div>
          <div className={`${"cell padding-small s12 m6 l6"}`} >
            <Input id={'notes_'+row.fieldname} name="fieldvalue_notes" type="text" 
              value={notes} className="full" 
              onChange={(value) => onEdit({
                id: id, name: "fieldvalue_notes", value: value})}
              disabled={(disabled || data.audit === 'readonly') ? 'disabled' : ''}/>
          </div>
        </div>
      </div>)
    
    case "col2":
      return(<div {...props} 
        className={`${className} ${"row full padding-small section-small border-bottom"}`}>
        <div className={`${"cell padding-small s12 m6 l6"}`} >
          <div>
            <Label className="bold" value={columns[0].label} />
          </div>
          <Field id={"field_"+columns[0].name}
            field={columns[0]} values={values} options={options} data={data}
            getText={getText} onEdit={onEdit} onEvent={onEvent} onSelector={onSelector} />
        </div>
        <div className={`${"cell padding-small s12 m6 l6"}`} >
          <div>
            <Label className="bold" value={columns[1].label} />
          </div>
          <Field id={"field_"+columns[1].name}
            field={columns[1]} values={values} options={options} data={data}
            getText={getText} onEdit={onEdit} onEvent={onEvent} onSelector={onSelector} />
        </div>
      </div>)
    
    case "col3":
      return(<div {...props} 
        className={`${className} ${"row full padding-small section-small border-bottom"}`}>
        <div className={`${"cell padding-small s12 m4 l4"}`} >
          <div>
            <Label className="bold" value={columns[0].label} />
          </div>
          <Field id={"field_"+columns[0].name}
            field={columns[0]} values={values} options={options} data={data}
            getText={getText} onEdit={onEdit} onEvent={onEvent} onSelector={onSelector} />
        </div>
        <div className={`${"cell padding-small s12 m4 l4"}`} >
          <div>
            <Label className="bold" value={columns[1].label} />
          </div>
          <Field id={"field_"+columns[1].name}
            field={columns[1]} values={values} options={options} data={data}
            getText={getText} onEdit={onEdit} onEvent={onEvent} onSelector={onSelector} />
        </div>
        <div className={`${"cell padding-small s12 m4 l4"}`} >
          <div>
            <Label className="bold" value={columns[2].label} />
          </div>
          <Field id={"field_"+columns[2].name}
            field={columns[2]} values={values} options={options} data={data}
            getText={getText} onEdit={onEdit} onEvent={onEvent} onSelector={onSelector} />
        </div>
      </div>)

    case "col4":
      return(<div {...props} 
        className={`${className} ${"row full padding-small section-small border-bottom"}`}>
        <div className={`${"cell padding-small s12 m3 l3"}`} >
          <div>
            <Label className="bold" value={columns[0].label} />
          </div>
          <Field id={"field_"+columns[0].name}
            field={columns[0]} values={values} options={options} data={data}
            getText={getText} onEdit={onEdit} onEvent={onEvent} onSelector={onSelector} />
        </div>
        <div className={`${"cell padding-small s12 m3 l3"}`} >
          <div>
            <Label className="bold" value={columns[1].label} />
          </div>
          <Field id={"field_"+columns[1].name}
            field={columns[1]} values={values} options={options} data={data}
            getText={getText} onEdit={onEdit} onEvent={onEvent} onSelector={onSelector} />
        </div>
        <div className={`${"cell padding-small s12 m3 l3"}`} >
          <div>
            <Label className="bold" value={columns[2].label} />
          </div>
          <Field id={"field_"+columns[2].name}
            field={columns[2]} values={values} options={options} data={data}
            getText={getText} onEdit={onEdit} onEvent={onEvent} onSelector={onSelector} />
        </div>
        <div className={`${"cell padding-small s12 m3 l3"}`} >
          <div>
            <Label className="bold" value={columns[3].label} />
          </div>
          <Field id={"field_"+columns[3].name}
            field={columns[3]} values={values} options={options} data={data}
            getText={getText} onEdit={onEdit} onEvent={onEvent} onSelector={onSelector} />
        </div>
      </div>)
    
    default:
      return null;
  }
}

Row.propTypes = {
  row: PropTypes.object.isRequired,
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

Row.defaultProps = {
  row: {},
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

export default Row;