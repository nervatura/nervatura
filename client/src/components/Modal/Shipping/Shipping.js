import { useState } from 'react';
import PropTypes from 'prop-types';

import 'styles/style.css';
import styles from './Shipping.module.css';

import Label from 'components/Form/Label'
import Button from 'components/Form/Button'
import Input from 'components/Form/Input'
import Icon from 'components/Form/Icon'

export const Shipping = ({
  partnumber, description, unit, batch_no, qty, className,
  getText, onClose, onShipping,
  ...props 
}) => {
  const [ state, setState ] = useState({
    batch_no: batch_no,
    qty: qty
  })
  return(
    <div className={`${"modal"} ${styles.modal}`} >
      <div className={`${"dialog"} ${styles.dialog}`} {...props} >
        <div className={`${styles.panel} ${className}`} >
          <div className={`${styles.panelTitle} ${"primary"}`}>
            <div className="row full">
              <div className="cell">
                <Label value={getText("shipping_movement_product")} 
                  leftIcon={<Icon iconKey="Truck" />} iconWidth="20px" />
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
                <div>
                  <Label className="bold" value={getText("product_description")} />
                </div>
                <Input className="full" value={description}
                  disabled="disabled" />
              </div>
            </div>
            <div className="row full">
              <div className={`${"cell padding-small"}`} >
                <div>
                  <Label className="bold" value={getText("product_unit")} />
                </div>
                <Input className="full" value={unit}
                  disabled="disabled" />
              </div>
            </div>
            <div className="row full">
              <div className={`${"cell padding-small"}`} >
                <div>
                  <Label className="bold" value={getText("movement_batchnumber")} />
                </div>
                <Input id="batch_no"
                  className="full" value={state.batch_no} autoFocus={true}
                  onChange={ (value)=>setState({ ...state, batch_no: value }) } />
              </div>
            </div>
            <div className="row full">
              <div className={`${"cell padding-small"}`} >
                <div>
                  <Label className="bold" value={getText("movement_qty")} />
                </div>
                <Input id="qty" className="full align-right" 
                  value={state.qty} type="number"
                  onChange={ (value)=>setState({ ...state, qty: value }) } />
              </div>
            </div>
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
                  onClick={ ()=>onShipping(state.batch_no, parseFloat(state.qty)) }
                  value={<Label center value={getText("msg_ok")} 
                    leftIcon={<Icon iconKey="Check" />} iconWidth="20px" />}
                />
              </div>
            </div>
          </div> 
        </div>
      </div>
    </div>
  )
}

Shipping.propTypes = {
  /**
   * Product public key
   */
  partnumber: PropTypes.string.isRequired,
   /**
    * Product description
    */
  description: PropTypes.string.isRequired,
  /**
   * Production unit
   */
  unit: PropTypes.string.isRequired,
  batch_no: PropTypes.string.isRequired,
  /**
   * Shipping quantity 
   */ 
  qty: PropTypes.number.isRequired,
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
    * Shipping handle
    */
  onShipping: PropTypes.func
}

Shipping.defaultProps = {
  partnumber: "",
  description: "", 
  unit: "",
  batch_no: "", 
  qty: 0,
  className: "",
  getText: undefined,
  onClose: undefined,
  onShipping: undefined
}

export default Shipping;