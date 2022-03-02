import { useState } from 'react';
import PropTypes from 'prop-types';

import 'styles/style.css';
import styles from './Template.module.css';

import Label from 'components/Form/Label'
import Button from 'components/Form/Button'
import Input from 'components/Form/Input'
import Select from 'components/Form/Select'
import Icon from 'components/Form/Icon'

export const DATA_TYPE = {
  TEXT: "string",
  LIST: "list",
  TABLE: "table"
}

export const Template = ({
  type, name, columns, className,
  getText, onClose, onData,
  ...props 
}) => {
  const [ state, setState ] = useState({
    type: type,
    name: name,
    columns: columns
  })
  return(
    <div className={`${"modal"} ${styles.modal}`} >
      <div className={`${"dialog"} ${styles.dialog}`} {...props} >
        <div className={`${styles.panel} ${className}`} >
          <div className={`${styles.panelTitle} ${"primary"}`}>
            <div className="row full">
              <div className="cell">
                <Label value={getText("template_label_new_data")} 
                  leftIcon={<Icon iconKey="Plus" />} iconWidth="20px" />
              </div>
              <div className={`${"cell align-right"} ${styles.closeIcon}`}>
                <Icon id="closeIcon" iconKey="Times" onClick={onClose} />
              </div>
            </div>
          </div>
          <div className="section-small">
            <div className="row full container-small">
              <div className="cell padding-small half">
                <div className="row full">
                  <Label className="bold" value={getText("template_data_type")} />
                </div>
                <Select id="type"
                  className="full" value={state.type} autoFocus={true}
                  onChange={ (value)=>setState({ ...state, type: value }) }
                  options={Object.keys(DATA_TYPE).map(key => {
                    return { value: DATA_TYPE[key], text: key }
                  })} />
              </div>
              <div className="cell padding-small half">
                <div className="row full">
                  <Label className="bold" value={getText("template_data_name")} />
                </div>
                <Input id="name"
                  className="full" value={state.name} 
                  onChange={ (value)=>setState({ ...state, name: value }) } />
              </div>
            </div>
            {(state.type === DATA_TYPE.TABLE)?<div className="row full container-small">
              <div className="row full">
                <div className={`${"cell padding-small"}`} >
                  <div>
                    <Label className="bold" value={getText("template_data_columns")} />
                  </div>
                  <textarea id="columns"
                    className={`${"full"} ${styles.textareaStyle}`} value={state.columns} rows={3}
                    onChange={ (event)=>setState({ ...state, columns: event.target.value }) } />
                </div>
              </div>
            </div>:null}
          </div>
          <div className={`${"row full section container-small secondary-title"}`}>
            <div className={`${"row full"}`}>
              <div className={`${"cell padding-small half"}`} >
                <Button id="btn_cancel"
                  className={`${"full"} ${styles.closeIcon} `}
                  onClick={ ()=>onClose() }
                  value={<Label center value={getText("msg_cancel")} 
                    leftIcon={<Icon iconKey="Times" />} iconWidth="20px"  />}
                />
              </div>
              <div className={`${"cell padding-small half"}`} >
                <Button id="btn_ok"
                  className={`${"full primary"}`}
                  disabled={(state.name==="")?"disabled":""}
                  onClick={ ()=>onData({
                    name: state.name, type: state.type, columns: state.columns
                  }) }
                  value={<Label center value={getText("msg_ok")} 
                    leftIcon={<Icon iconKey="Check" />} iconWidth="20px"  />}
                />
              </div>
            </div>
          </div> 
        </div>
      </div>
    </div>
  )
}

Template.propTypes = {
  type: PropTypes.oneOf(Object.values(DATA_TYPE)).isRequired,
  name: PropTypes.string.isRequired,
  /**
   * Columns names (separated by commas) - TABLE
   */
  columns: PropTypes.string.isRequired,
  className: PropTypes.string.isRequired,
  /**
   * Localization
   */
  getText: PropTypes.func,
  /**
    * Close form handle (modal style)
    */ 
  onClose: PropTypes.func,
  onData: PropTypes.func
}

Template.defaultProps = {
  type: DATA_TYPE.TEXT,
  name: "",
  columns: "",
  className: "",
  getText: undefined,
  onClose: undefined,
  onData: undefined
}

export default Template;