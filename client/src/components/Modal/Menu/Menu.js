import { useState } from 'react';
import PropTypes from 'prop-types';

import 'styles/style.css';
import styles from './Menu.module.css';

import Label from 'components/Form/Label'
import Button from 'components/Form/Button'
import Input from 'components/Form/Input'
import Select from 'components/Form/Select'
import Icon from 'components/Form/Icon'

export const Menu = ({
  idKey, menu_id, fieldname, description, fieldtype, fieldtypeOptions, orderby, className,
  getText, onClose, onMenu,
  ...props 
}) => {
  const [ state, setState ] = useState({
    fieldname: fieldname,
    description: description,
    fieldtype: fieldtype,
    orderby: orderby
  })
  return(
    <div className={`${"modal"} ${styles.modal}`} >
      <div className={`${"dialog"} ${styles.dialog}`} {...props} >
        <div className={`${styles.panel} ${className}`} >
          <div className={`${styles.panelTitle} ${"primary"}`}>
            <div className="row full">
              <div className="cell">
                <Label value={getText("title_menucmd")} 
                  leftIcon={<Icon iconKey="Share" />} iconWidth="20px" />
              </div>
              <div className={`${"cell align-right"} ${styles.closeIcon}`}>
                <Icon id="closeIcon" iconKey="Times" onClick={onClose} />
              </div>
            </div>
          </div>
          <div className="section-small">
            <div className="row full container-small">
              <div className="row full">
                <div className={`${"cell padding-small"}`} >
                  <div>
                    <Label className="bold" value={getText("menufields_fieldname")} />
                  </div>
                  <Input id="fieldname"
                    className="full" value={state.fieldname}
                    onChange={ (value)=>setState({ ...state, fieldname: value }) } />
                </div>
              </div>
            </div>
            <div className="row full container-small">
              <div className="row full">
                <div className={`${"cell padding-small"}`} >
                  <div>
                    <Label className="bold" value={getText("menufields_description")} />
                  </div>
                  <Input id="description"
                    className="full" value={state.description}
                    onChange={ (value)=>setState({ ...state, description: value }) } />
                </div>
              </div>
            </div>
            <div className="row full container-small">
              <div className="cell padding-small half">
                <div className="row full">
                  <Label className="bold" value={getText("menufields_fieldtype")} />
                </div>
                <Select id="fieldtype" 
                  className="full" value={String(state.fieldtype)}
                  onChange={ (value)=>setState({ ...state, fieldtype: parseInt(value,10) }) }
                  options={fieldtypeOptions} />
              </div>
              <div className="cell padding-small half">
                <div className="row full">
                  <Label className="bold" value={getText("menufields_orderby")} />
                </div>
                <Input id="orderby"
                  className="full align-right" value={String(state.orderby)} type="integer"
                  onChange={ (value)=>setState({ ...state, orderby: value }) } />
              </div>
            </div>
          </div>
          <div className={`${"row full section container-small secondary-title"}`}>
            <div className={`${"row full"}`}>
              <div className={`${"cell padding-small half"}`} >
                <Button id="btn_cancel" 
                  className={`${"full"} ${styles.closeIcon} `}
                  onClick={onClose} 
                  value={<Label center value={getText("msg_cancel")} 
                    leftIcon={<Icon iconKey="Times" />} iconWidth="20px"  />}
                />
              </div>
              <div className={`${"cell padding-small half"}`} >
                <Button id="btn_ok"
                  className={`${"full primary"}`}
                  disabled={((state.fieldname === "")||(state.description === ""))?"disabled":""}
                  onClick={ ()=>onMenu({
                    id: idKey, menu_id: menu_id, fieldname: state.fieldname, description: state.description, 
                    fieldtype: state.fieldtype, orderby: state.orderby
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

Menu.propTypes = {
  idKey: PropTypes.number,
  menu_id: PropTypes.number.isRequired, 
  fieldname: PropTypes.string.isRequired, 
  description: PropTypes.string.isRequired, 
  fieldtype: PropTypes.number.isRequired, 
  fieldtypeOptions: PropTypes.array.isRequired, 
  orderby: PropTypes.number.isRequired,
  className: PropTypes.string.isRequired,
  /**
   * Localization
   */
  getText: PropTypes.func,
  /**
    * Close form handle (modal style)
    */ 
  onClose: PropTypes.func,
  onMenu: PropTypes.func
}

Menu.defaultProps = {
  idKey: null,
  menu_id: 0, 
  fieldname: "", 
  description: "", 
  fieldtype: 0, 
  fieldtypeOptions: [], 
  orderby: 0,
  className: "",
  getText: undefined,
  onClose: undefined,
  onMenu: undefined
}

export default Menu;