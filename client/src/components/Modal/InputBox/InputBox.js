import { useState } from 'react';
import PropTypes from 'prop-types';

import 'styles/style.css';
import styles from './InputBox.module.css';

import Label from 'components/Form/Label'
import Button from 'components/Form/Button'
import Input from 'components/Form/Input'
import Icon from 'components/Form/Icon'

export const InputBox = ({
  title, message, infoText, value, labelCancel, labelOK, defaultOK, showValue, className,
  onOK, onCancel,
  ...props 
}) => {
  const [ state, setState ] = useState({
    value: value,
  })
  return (
    <div className={`${"modal"} ${styles.modal} ${className}`} >
      <div className={`${"dialog"} ${styles.dialog}`} {...props} >
        <div className={`${styles.panel}`} >
          <div className={`${styles.panelTitle} ${"primary"}`}>
            <div className="row full">
              <div className="cell">
                <Label value={title} />
              </div>
            </div>
          </div>
          <div className="row full container-small section-small">
            <div className="cell padding-normal">
              <div className={`${styles.input}`}>{message}</div>
              {(infoText)?
                <div className={`${"section-small-top"} ${styles.info}`}>
                  {infoText}
                </div>:null}
              {(showValue)?
                <div className={`${"section-small-top"}`}>
                  <Input id="input_value" type="text" className="full" 
                    value={state.value} autoFocus={true}
                    onChange={ 
                      (value) => setState({ ...state, value: value }) 
                    }
                    onEnter={()=>onOK(state.value)} />
                </div>:null}
            </div>
          </div>
          <div className={`${"row full section container-small secondary-title"}`}>
            <div className={`${"row full"}`}>
              <div className={`${"cell padding-small half"}`} >
                <Button id="btn_cancel" className={`${"full"} ${styles.closeIcon} `}
                  autoFocus={(showValue)?false:!defaultOK}
                  onClick={onCancel}
                  value={<Label center value={labelCancel} leftIcon={<Icon iconKey="Times" />} iconWidth="20px"  />}
                />
              </div>
              <div className={`${"cell padding-small half"}`} >
                <Button id="btn_ok" className={`${"full primary"}`}
                  autoFocus={(showValue)?false:defaultOK}
                  onClick={()=>onOK(state.value)}
                  value={<Label center value={labelOK} leftIcon={<Icon iconKey="Check" />} iconWidth="20px"  />} 
                />
              </div>
            </div>
          </div> 
        </div>
      </div>
    </div>
  )
}

InputBox.propTypes = {
  /**
   * Message box title
   */
  title: PropTypes.string.isRequired,
  /**
   * Message text
   */ 
  message: PropTypes.string.isRequired,
  /**
   * Optional info text
   */
  infoText: PropTypes.string,
  /**
   * Default input value
   */ 
  value: PropTypes.string.isRequired,
  /**
   * Cancel button label
   */
  labelCancel: PropTypes.string.isRequired,
  /**
   * OK button label 
   */ 
  labelOK: PropTypes.string.isRequired,
  /** 
   * OK button focus
  */
  defaultOK: PropTypes.bool.isRequired,
  /** 
   * Show/hide input box
  */
  showValue: PropTypes.bool.isRequired,
  className: PropTypes.string.isRequired,
  /**
   * OK button handle
   */
  onOK: PropTypes.func,
  /**
   * Cancel button handle
   */
  onCancel: PropTypes.func,
}

InputBox.defaultProps = {
  title: "",
  message: "",
  infoText: undefined,
  value: "",
  labelCancel: "Cancel",
  labelOK: "OK",
  defaultOK: false,
  showValue: false,
  className: "",
  onOK: undefined,
  onCancel: undefined
}

export default InputBox;