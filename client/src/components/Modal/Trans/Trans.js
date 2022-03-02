import { useState } from 'react';
import PropTypes from 'prop-types';

import 'styles/style.css';
import styles from './Trans.module.css';

import Label from 'components/Form/Label'
import Button from 'components/Form/Button'
import Select from 'components/Form/Select'
import Icon from 'components/Form/Icon'

export const Trans = ({
  baseTranstype, transtype, direction, doctypes, directions, refno, 
  nettoDiv, netto, fromDiv, from, elementCount, className,
  getText, onClose, onCreate,
  ...props 
}) => {
  const [ state, setState ] = useState({
    transtype: transtype,
    direction: direction,
    refno: refno,
    netto: netto,
    from: from,
    nettoDiv: nettoDiv,
    fromDiv: fromDiv
  })
  const setTranstype = (value) => {
    let nettoDiv = state.nettoDiv
    let fromDiv = state.fromDiv
    if(["invoice","receipt"].includes(value) && ["order","rent","worksheet"].includes(baseTranstype)){
        nettoDiv = true
        if(elementCount===0){
          fromDiv = true
        }
    } else {
      nettoDiv = false
      fromDiv = false
    }
    setState({ ...state, transtype: value, nettoDiv: nettoDiv, fromDiv: fromDiv })
  }
  const typeOptions = doctypes.map(dt => { return { value: dt, text: dt } })
  const dirOptions = directions.map(dir => { return { value: dir, text: dir } })
  return(
    <div className={`${"modal"} ${styles.modal}`} >
      <div className={`${"dialog"} ${styles.dialog}`} {...props} >
        <div className={`${styles.panel} ${className}`} >
          <div className={`${styles.panelTitle} ${"primary"}`}>
            <div className="row full">
              <div className="cell">
                <Label value={getText("msg_create_title")} 
                  leftIcon={<Icon iconKey="FileText" />} iconWidth="20px" />
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
                  <Label className="bold" value={getText("msg_create_new")} />
                </div>
              </div>
            </div>
            <div className="row full">
              <div className={`${"cell half padding-small"}`} >
                <Select id="transtype" 
                  className="full" value={state.transtype} 
                  onChange={ (value)=>setTranstype(value) }
                  options={typeOptions} />
              </div>
              <div className={`${"cell half padding-small"}`} >
                <Select id="direction" 
                  className="full" value={state.direction} 
                  onChange={ (value)=>setState({ ...state, direction: value }) }
                  options={dirOptions} />
              </div>
            </div>
            <div className="row full">
              <div className={`${"cell padding-small"}`} >
                <div id="refno"
                  className={`${"padding-small"} ${styles.editCol}`} 
                  onClick={()=>setState({ ...state, refno: !state.refno })}>
                  <Label className="bold" value={getText("msg_create_setref")} 
                    leftIcon={<Icon iconKey={(state.refno)?"CheckSquare":"SquareEmpty"} />} />
                </div>
              </div>
            </div>
            {(state.nettoDiv)?<div className="row full">
              <div className={`${"cell padding-small"}`} >
                <div id="netto"
                  className={`${"padding-small"} ${styles.editCol}`} 
                  onClick={()=>setState({ ...state, netto: !state.netto })}>
                  <Label className="bold" value={getText("msg_create_deduction")} 
                    leftIcon={<Icon iconKey={(state.netto)?"CheckSquare":"SquareEmpty"} />} />
                </div>
              </div>
            </div>:null}
            {(state.fromDiv)?<div className="row full">
              <div className={`${"cell padding-small"}`} >
                <div id="from"
                  className={`${"padding-small"} ${styles.editCol}`} 
                  onClick={()=>setState({ ...state, from: !state.from })}>
                  <Label className="bold" value={getText("msg_create_delivery")} 
                    leftIcon={<Icon iconKey={(state.from)?"CheckSquare":"SquareEmpty"} />} />
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
                <Button id="btn_create" 
                  className={`${"full primary"}`} 
                  onClick={ ()=>onCreate({ 
                    newTranstype: state.transtype, 
                    newDirection: state.direction, 
                    refno: state.refno, 
                    fromInventory: (state.from && state.fromDiv), 
                    nettoQty: (state.netto && state.nettoDiv)
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

Trans.propTypes = {
  baseTranstype: PropTypes.string.isRequired,
  transtype: PropTypes.string.isRequired, 
  direction: PropTypes.string.isRequired, 
  doctypes: PropTypes.array.isRequired, 
  directions: PropTypes.array.isRequired, 
  refno: PropTypes.bool.isRequired, 
  nettoDiv: PropTypes.bool.isRequired, 
  netto: PropTypes.bool.isRequired, 
  fromDiv: PropTypes.bool.isRequired, 
  from: PropTypes.bool.isRequired, 
  elementCount: PropTypes.number.isRequired,
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
   * Create a new document
   */
  onCreate: PropTypes.func
}

Trans.defaultProps = {
  baseTranstype: "",
  transtype: "",
  direction: "",
  doctypes: [],
  directions: [],
  refno: true, 
  nettoDiv: false, 
  netto: true, 
  fromDiv: false, 
  from: false,
  elementCount: 0,
  className: "",
  getText: undefined,
  onClose: undefined,
  onCreate: undefined
}

export default Trans;