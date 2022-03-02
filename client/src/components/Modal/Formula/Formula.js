import { useState } from 'react';
import PropTypes from 'prop-types';

import 'styles/style.css';
import styles from './Formula.module.css';

import Label from 'components/Form/Label'
import Button from 'components/Form/Button'
import Input from 'components/Form/Input'
import Select from 'components/Form/Select'
import Icon from 'components/Form/Icon'

export const Formula = ({
  formula, formulaValues, partnumber, description, className,
  getText, onClose, onFormula,
  ...props 
}) => {
  const [ state, setState ] = useState({
    formula: formula,
  })
  return(
    <div className={`${"modal"} ${styles.modal}`} >
      <div className={`${"dialog"} ${styles.dialog}`} {...props} >
        <div className={`${styles.panel} ${className}`} >
          <div className={`${styles.panelTitle} ${"primary"}`}>
            <div className="row full">
              <div className="cell">
                <Label value={getText("label_formula")} leftIcon={<Icon iconKey="Magic" />} iconWidth="20px" />
              </div>
              <div className={`${"cell align-right"} ${styles.closeIcon}`}>
                <Icon id="closeIcon" iconKey="Times" onClick={onClose} />
              </div>
            </div>
          </div>
          <div className="row full container-small section-small">
            <div className="row full">
              <div className={`${"cell padding-small"}`} >
                <div>
                  <Label className="bold" value={getText("product_partnumber")} />
                </div>
                <Input className="full" value={partnumber}
                  disabled="disabled" />
              </div>
            </div>
            <div className="row full">
              <div className={`${"cell padding-small"}`} >
                <Input className="full" value={description}
                  disabled="disabled" />
              </div>
            </div>
            <div className="row full">
              <div className={`${"cell padding-small"}`} >
                <Select id="formula"
                  className="full" value={state.formula} placeholder=""
                  onChange={ (value)=>setState({ ...state, formula: value }) }
                  options={formulaValues} />
              </div>
            </div>
          </div>
          <div className={`${"row full section container-small secondary-title"}`}>
            <div className={`${"row full"}`}>
              <div className={`${"cell padding-small half"}`} >
                <Button id="btn_cancel"
                  className={`${"full"} ${styles.closeIcon} `} 
                  disabled={(state.formula==="")?"disabled":""}
                  onClick={ ()=>onClose() } 
                  value={<Label center value={getText("msg_cancel")} leftIcon={<Icon iconKey="Times" />} iconWidth="20px"  />}
                />
              </div>
              <div className={`${"cell padding-small half"}`} >
                <Button id="btn_formula"
                  className={`${"full primary"}`} 
                  disabled={(state.formula==="")?"disabled":""}
                  onClick={ ()=>onFormula(parseInt(state.formula,10)) } 
                  value={<Label center value={getText("msg_ok")} leftIcon={<Icon iconKey="Check" />} iconWidth="20px"  />}
                />
              </div>
            </div>
          </div> 
        </div>
      </div>
    </div>
  )
}

Formula.propTypes = {
  /**
   * Default formula ID
   */
  formula: PropTypes.string.isRequired,
  /**
   * Product formula values
   */
  formulaValues: PropTypes.array.isRequired,
  /**
   * Product public key
   */
  partnumber: PropTypes.string.isRequired,
  /**
   * Product description
   */
  description: PropTypes.string.isRequired, 
  className: PropTypes.string.isRequired,
  /**
   * Localization
   */
  getText: PropTypes.func,
  /**
    * Close form handle (modal style)
    */ 
  onClose: PropTypes.func,
  /**
   * Loading selected formula
   */
  onFormula: PropTypes.func
}

Formula.defaultProps = {
  formula: "",
  formulaValues: [],
  partnumber: "",
  description: "",
  className: "",
  getText: undefined,
  onClose: undefined,
  onFormula: undefined
}

export default Formula;