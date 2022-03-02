import PropTypes from 'prop-types';

import 'styles/style.css';
import styles from './Stock.module.css';

import { getSetting } from 'config/app'
import Label from 'components/Form/Label'
import Button from 'components/Form/Button'
import Icon from 'components/Form/Icon'
import Table from 'components/Form/Table'

export const Stock = ({
  partnumber, partname, rows, selectorPage, className,
  onClose, getText,
  ...props 
}) => {
  const fields = {
    warehouse: { fieldtype:"string", label:getText("delivery_place") },
    batch_no: { fieldtype:"string", label:getText("movement_batchnumber") },
    description: { fieldtype:"string", label:getText("product_description") },
    sqty: { fieldtype:"number", label:getText("shipping_stock") }
  }
  return(
    <div className={`${"modal"} ${styles.modal}`} >
      <div className={`${"dialog"} ${styles.dialog}`} {...props} >
        <div className={`${styles.panel} ${className}`} >
          <div className={`${styles.panelTitle} ${"primary"}`}>
            <div className="row full">
              <div className="cell">
                <Label value={getText("shipping_stocks")} 
                  leftIcon={<Icon iconKey="Book" />} iconWidth="20px" />
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
                  <Label className="bold" value={partnumber} />
                </div>
                <div>
                  <Label value={partname} />
                </div>
              </div>
            </div>
            <div className="row full">
              <Table rowKey="id"
                fields={fields} rows={rows} tableFilter={true}
                filterPlaceholder={getText("placeholder_filter")} 
                paginationPage={selectorPage} paginationTop={true}
                hidePaginatonSize={true} />
            </div>
          </div>
          <div className={`${"row full section container-small secondary-title"}`}>
            <div className={`${"row full"}`}>
              <div className={`${"cell padding-small"}`} >
                <Button id="btn_ok"
                  className={`${"full primary"}`} onClick={onClose}
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

Stock.propTypes = {
/**
  * Product public key
  */
partnumber: PropTypes.string.isRequired, 
/**
 * Product description
 */
partname: PropTypes.string.isRequired,
/**
 * Product stock values 
 */ 
rows: PropTypes.array.isRequired,
/**
   * Pagination row number / page
   */
selectorPage: PropTypes.number.isRequired, 
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

Stock.defaultProps = {
  partnumber: "",
  partname: "",
  rows: [],
  selectorPage: getSetting("selectorPage"),
  className: "",
  getText: undefined,
  onClose: undefined,
}

export default Stock;