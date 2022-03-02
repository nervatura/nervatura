import PropTypes from 'prop-types';

import styles from './Item.module.css';

import Row from 'components/Form/Row'
import Label from 'components/Form/Label'

export const Item = ({ 
  current, dataset, audit, className,
  getText, onEvent,
  ...props 
}) => {  
  const editItem = (options) => onEvent("editItem", [options])
  const onSelector = (...params) => onEvent("onSelector", [...params])
  return (
    <div {...props} className={`${className}`} >
      <div className="row full" >
        <div className={`${"cell padding-normal border secondary-title"} ${styles.itemTitle}` }>
          <Label className={` ${styles.itemTitlePre}` } 
            value={(current.form.id === null) ? getText("label_new") : String(current.form.id)} />
          <Label value={current.form_template.options.title} />
        </div>
      </div>
      <div className={`${styles.formPanel} ${"border"}`} >
        {current.form_template.rows.map((row, index) =>
          <Row key={index} row={row} 
            values={current.form} options={current.form_template.options}
            data={{
              audit: audit,
              current: current,
              dataset: dataset,
            }}
            getText={getText} onEdit={editItem} onEvent={onEvent} onSelector={onSelector}
          />
        )}
      </div>
    </div>
  )
}

Item.propTypes = {
  current: PropTypes.object.isRequired,  
  dataset: PropTypes.object.isRequired, 
  audit: PropTypes.string.isRequired,
  className: PropTypes.string,
  onEvent: PropTypes.func,
  getText: PropTypes.func,
}

Item.defaultProps = {
  current: {},  
  dataset: {}, 
  audit: "",
  className: "",
  onEvent: undefined,
  getText: undefined,
}

export default Item;