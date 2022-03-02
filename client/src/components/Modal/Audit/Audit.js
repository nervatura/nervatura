import { useState } from 'react';
import PropTypes from 'prop-types';

import 'styles/style.css';
import styles from './Audit.module.css';

import Label from 'components/Form/Label'
import Button from 'components/Form/Button'
import Select from 'components/Form/Select'
import Icon from 'components/Form/Icon'

export const Audit = ({
  idKey, usergroup, nervatype, subtype, inputfilter, supervisor,
  typeOptions, subtypeOptions, inputfilterOptions, className,
  getText, onClose, onAudit,
  ...props 
}) => {
  const getSubtype = (ntype, ivalue) => {
    if(ivalue){
      return String(ivalue)
    }
    const typeName = typeOptions.filter(item => (item.value === ntype))[0].text
    if(["trans", "report","menu"].includes(typeName)){
      const stype = subtypeOptions.filter(item => (item.type === typeName))[0]
      if(stype){
        return stype.value
      }
    }
    return undefined
  }
  const [ state, setState ] = useState({
    nervatype: String(nervatype),
    subtype: getSubtype(String(nervatype), subtype),
    inputfilter: String(inputfilter),
    supervisor: supervisor
  })
  const nervatypeName = typeOptions.filter(item => (item.value === state.nervatype))[0].text
  const isSubtype = ["trans", "report","menu"].includes(typeOptions.filter(item => (item.value === state.nervatype))[0].text)
  return(
    <div className={`${"modal"} ${styles.modal}`} >
      <div className={`${"dialog"} ${styles.dialog}`} {...props} >
        <div className={`${styles.panel} ${className}`} >
          <div className={`${styles.panelTitle} ${"primary"}`}>
            <div className="row full">
              <div className="cell">
                <Label value={getText("title_usergroup")} 
                  leftIcon={<Icon iconKey="Key" />} iconWidth="20px" />
              </div>
              <div className={`${"cell align-right"} ${styles.closeIcon}`}>
                <Icon id="closeIcon" iconKey="Times" onClick={onClose} />
              </div>
            </div>
          </div>
          <div className="section-small">
            <div className="row full container-small section-small">
              <div className="row full">
                <div className={`${"cell padding-small"}`} >
                  <div>
                    <Label className="bold" value={getText("audit_nervatype")} />
                  </div>
                  <Select id="nervatype"
                    className="full" value={state.nervatype}
                    disabled={(idKey)?"disabled":""}
                    onChange={(value)=>setState({ ...state, 
                      nervatype: value, 
                      subtype: getSubtype(value,null) 
                    })}
                    options={typeOptions} />
                </div>
              </div>
            </div>
            {(isSubtype)?<div className="row full container-small">
              <div className="row full">
                <div className={`${"cell padding-small"}`} >
                  <div>
                    <Label className="bold" value={getText("audit_subtype")} />
                  </div>
                  <Select id="subtype"
                    className="full" value={state.subtype}
                    onChange={ (value)=>setState({ ...state, subtype: value }) }
                    options={subtypeOptions.filter(item => (item.type === nervatypeName))} />
                </div>
              </div>
            </div>:null}
            <div className="row full container-small">
              <div className="row full">
                <div className={`${"cell padding-small"}`} >
                  <div>
                    <Label className="bold" value={getText("audit_inputfilter")} />
                  </div>
                  <Select id="inputfilter"
                    className="full" value={state.inputfilter}
                    onChange={ (value)=>setState({ ...state, inputfilter: value }) }
                    options={inputfilterOptions} />
                </div>
              </div>
            </div>
            <div className="row full container-small">
              <div className="row">
                <div className={`${"cell padding-small"}`} >
                  <div id="supervisor"
                    className={`${"padding-small"} ${styles.reportField}`} 
                    onClick={ ()=>setState({ ...state, supervisor: (state.supervisor===1)?0:1 })}>
                    <Label className="bold" value={getText("audit_supervisor")} 
                      leftIcon={<Icon id={"check_"+state.supervisor} iconKey={(state.supervisor===1)?"CheckSquare":"SquareEmpty"} />} />
                  </div>
                </div>
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
                  disabled={(isSubtype && !state.subtype)?"disabled":""}
                  onClick={()=>onAudit({
                    id: idKey, usergroup: usergroup,
                    nervatype: parseInt(state.nervatype,10),
                    subtype: (state.subtype)?parseInt(state.subtype,10):null,
                    inputfilter: parseInt(state.inputfilter,10),
                    supervisor: state.supervisor
                  })}
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

Audit.propTypes = {
  idKey: PropTypes.number,
  usergroup: PropTypes.number.isRequired,
  nervatype: PropTypes.number.isRequired,
  subtype: PropTypes.number,
  inputfilter: PropTypes.number.isRequired,
  supervisor: PropTypes.number.isRequired,
  typeOptions: PropTypes.array.isRequired, 
  subtypeOptions: PropTypes.array.isRequired, 
  inputfilterOptions: PropTypes.array.isRequired,
  className: PropTypes.string.isRequired,
  /**
   * Localization
   */
  getText: PropTypes.func,
  /**
   * Close form handle (modal style)
   */ 
  onClose: PropTypes.func,
  onAudit: PropTypes.func
}

Audit.defaultProps = {
  idKey: null,
  usergroup: 0,
  nervatype: undefined,
  subtype: null,
  inputfilter: undefined,
  supervisor: 0,
  typeOptions: [],
  subtypeOptions: [],
  inputfilterOptions: [],
  className: "",
  getText: undefined,
  onClose: undefined,
  onAudit: undefined
}

export default Audit;