import PropTypes from 'prop-types';

import 'styles/style.css';
import styles from './Total.module.css';

import Label from 'components/Form/Label'
import Button from 'components/Form/Button'
import Input from 'components/Form/Input'
import Icon from 'components/Form/Icon'

export const Total = ({
  total, className,
  getText, onClose,
  ...props 
}) => {
  const formatNumber = (number) => {
    return number.toString().replace(/(\d)(?=(\d{3})+(?!\d))/g, '$1,')
  }
  return(
    <div className={`${"modal"} ${styles.modal}`} >
      <div className={`${"dialog"} ${styles.dialog}`} {...props} >
        <div className={`${styles.panel} ${className}`} >
          <div className={`${styles.panelTitle} ${"primary"}`}>
            <div className="row full">
              <div className="cell">
                <Label value={getText("browser_total")} 
                  leftIcon={<Icon iconKey="InfoCircle" />} iconWidth="20px" />
              </div>
              <div className={`${"cell align-right"} ${styles.closeIcon}`}>
                <Icon id="closeIcon" iconKey="Times" onClick={onClose} />
              </div>
            </div>
          </div>
          <div className="row full container section-small">
            {Object.keys(total.totalFields).map(fieldname => 
              <div key={fieldname} className="row full mobile">
                <div className="cell bold padding-tiny">
                  <Label className="bold" value={total.totalLabels[fieldname]} />
                </div>
                <div className="cell padding-tiny right mobile">
                  <Input className={`${"align-right bold"} ${styles.maxInput}`} 
                    value={formatNumber(total.totalFields[fieldname])} 
                    disabled="disabled" />
                </div>
              </div>)}
          </div>
          <div className={`${"row full section container-small secondary-title"}`}>
            <div className={`${"row full"}`}>
              <div className={`${"cell padding-small"}`} >
                <Button id="btn_ok"
                  className={`${"full primary"}`} 
                  onClick={onClose}
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

Total.propTypes = {
  total: PropTypes.shape({
    totalFields: PropTypes.object.isRequired,
    totalLabels: PropTypes.object.isRequired,
    count: PropTypes.number.isRequired,
  }).isRequired,
  className: PropTypes.string.isRequired,
  /**
   * Localization
   */
  getText: PropTypes.func,
  /**
    * Close form handle (modal style)
    */ 
  onClose: PropTypes.func,
}

Total.defaultProps = {
  total: {
    totalFields: {},
    totalLabels: {},
    count: 0
  },
  className: "",
  getText: undefined,
  onClose: undefined,
}

export default Total;