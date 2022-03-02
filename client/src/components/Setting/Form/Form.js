import { Fragment } from 'react';
import PropTypes from 'prop-types';
import update from 'immutability-helper';

import Icon from 'components/Form/Icon'
import Label from 'components/Form/Label'
import Table from 'components/Form/Table';
import Row from 'components/Form/Row';

import styles from './Form.module.css';
import { getSetting } from 'config/app'

export const Form = ({ 
  data, className,
  paginationPage, dateFormat, timeFormat,
  getText, onEvent,
  ...props 
}) => {
  const { caption, icon, current, audit, dataset, type, view } = data
  let fields = {}
  if((typeof current.template.view.items !== "undefined") && (current.form.id !== null)){
    if(current.template.view.items.actions.edit){
      fields = update(fields, {$merge: {
        edit: { columnDef: { 
          id: "edit",
          Header: "",
          headerStyle: {},
          Cell: ({ row, value }) => {
            const ecol = <div 
              className={`${"cell"} ${styles.editCol}`} >
              <Icon id={"edit_"+row.original["id"]}
                iconKey="Edit" width={24} height={21.3} 
                onClick={(event)=>{
                  event.stopPropagation();
                  onEvent("setViewActions", [current.template.view.items.actions.edit, row.original])
                }}
                className={styles.editCol} />
            </div>
            const dcol = (current.template.view.items.actions.delete)?<div 
              className={`${"cell"} ${styles.deleteCol}`} >
              <Icon id={"delete_"+row.original["id"]}
                iconKey="Times" width={19} height={27.6} 
                onClick={(event)=>{
                  event.stopPropagation();
                  onEvent("setViewActions", [current.template.view.items.actions.delete, row.original])
                }}
                className={styles.deleteCol} />
            </div>:null
            return <Fragment>{ecol}{dcol}</Fragment>
          },
          cellStyle: { width: 30, padding: "7px 3px 3px 8px" }
        }}
      }})
    }
    fields = update(fields, {$merge: {...current.template.view.items.fields}})
  }
  const editItem = (options) => onEvent("editItem", [options])
  return (
    <div {...props} className={`${styles.width800} ${className}`}>
      <div className={`${styles.panel}`} >
        <div className={`${styles.panelTitle} ${"primary"}`}>
          <Label value={caption} 
            leftIcon={<Icon iconKey={icon} />} iconWidth="20px" />
        </div>
        <div className={`${"section"} ${styles.settingPanel}`} >
          <div className="row full container section-small-bottom" >
          <Fragment >
            <div className={`${"border"} ${styles.formPanel}`} >
              {current.template.rows.map((row, index) =>
                <Row key={index} row={row} 
                  values={current.fieldvalue || current.form}
                  options={current.template.options}
                  data={{
                    audit: audit,
                    current: current,
                    dataset: dataset
                  }} 
                  getText={getText} onEdit={editItem}
                />
              )}
            </div>
            {(((typeof current.template.view.items !== "undefined") && (current.form.id !== null))
              || (type === "log"))?
              <Table rowKey="id"
                onAddItem={(current.template.view.items && current.template.view.items.actions.new) 
                  ? ()=>onEvent("setViewActions", [current.template.view.items.actions.new]) : null}
                labelAdd={getText("label_new")}
                fields={(type === "log") ? view.fields : fields} 
                rows={(type === "log") ? view.result : dataset[current.template.view.items.data]} 
                tableFilter={true}
                filterPlaceholder={getText("placeholder_filter")}
                labelYes={getText("label_yes")} labelNo={getText("label_no")}
                dateFormat={dateFormat} timeFormat={timeFormat} 
                paginationPage={paginationPage} paginationTop={true}/>:null}
          </Fragment>
          </div>
        </div>
      </div>
    </div>
  )
}

Form.propTypes = {
  data: PropTypes.shape({
    caption: PropTypes.string.isRequired,
    icon: PropTypes.string.isRequired,
    current: PropTypes.object.isRequired,
    audit: PropTypes.string.isRequired, 
    dataset: PropTypes.object.isRequired,
    type: PropTypes.string.isRequired, 
    view: PropTypes.object.isRequired,
  }).isRequired,
  paginationPage: PropTypes.number.isRequired, 
  dateFormat: PropTypes.string, 
  timeFormat: PropTypes.string,
  className: PropTypes.string,
  onEvent: PropTypes.func,
  getText: PropTypes.func,
}

Form.defaultProps = {
  data: {
    caption: "",
    icon: "",
    current: {},
    audit: "", 
    dataset: {},
    type: "", 
    view: {},
  },
  paginationPage: getSetting("paginationPage"),
  dateFormat: getSetting("dateFormat"),
  timeFormat: getSetting("timeFormat"),
  className: "",
  onEvent: undefined,
  getText: undefined,
}

export default Form;