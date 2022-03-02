import { useState } from 'react';
import PropTypes from 'prop-types';

import 'styles/style.css';
import styles from './Server.module.css';

import Label from 'components/Form/Label'
import Button from 'components/Form/Button'
import Icon from 'components/Form/Icon'
import Row from 'components/Form/Row'

export const Server = ({
  cmd, fields, values, className,
  getText, onClose, onOK,
  ...props 
}) => {
  const [ state, setState ] = useState({
    ...values
  })
  return(
    <div className={`${"modal"} ${styles.modal}`} >
      <div className={`${"dialog"} ${styles.dialog}`} {...props} >
        <div className={`${styles.panel} ${className}`} >
          <div className={`${styles.panelTitle} ${"primary"}`}>
            <div className="row full">
              <div className="cell">
                <Label value={cmd.description} leftIcon={<Icon iconKey="Share" />} iconWidth="20px" />
              </div>
              <div className={`${"cell align-right"} ${styles.closeIcon}`}>
                <Icon id="closeIcon" iconKey="Times" onClick={onClose} />
              </div>
            </div>
          </div>
          {fields.map((field, index) =>
            <Row key={index}
              row={{
                rowtype: "field", 
                name: field.fieldname,
                datatype: field.fieldtypeName,
                label: field.description 
              }} 
              values={{...state}}
              options={{}}
              data={{
                audit: "all",
                current: {},
                dataset: {},  
              }}
              onEdit={
                (options)=>setState({ ...state, [options.name]: options.value })
              }
              getText={getText}
            />
          )}
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
                  onClick={ ()=>onOK({
                    cmd: cmd, fields: fields, 
                    values: {...state}
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

Server.propTypes = {
  cmd: PropTypes.object.isRequired, 
  fields: PropTypes.array.isRequired, 
  values: PropTypes.object.isRequired,
  className: PropTypes.string.isRequired,
  /**
   * Localization
   */
  getText: PropTypes.func,
  /**
    * Close form handle (modal style)
    */ 
  onClose: PropTypes.func,
  onOK: PropTypes.func
}

Server.defaultProps = {
  cmd: {}, 
  fields: [], 
  values: {},
  className: "",
  getText: undefined,
  onClose: undefined,
  onOK: undefined
}

export default Server;