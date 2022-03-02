import PropTypes from 'prop-types';

import styles from './Main.module.css';

import Row from 'components/Form/Row'

export const Main = ({ 
  current, template, dataset, audit, className,
  getText, onEvent,
  ...props 
}) => {  
  const editItem = (options) => onEvent("editItem", [options])
  const onSelector = (...params) => onEvent("onSelector", [...params])
  return (
    <div {...props} className={`${styles.formPanel} ${"border"} ${className}`} >
      {template.rows.map((row, index) => <Row key={index} row={row} 
        values={current.item} options={template.options}
        data={{ audit: audit, current: current, dataset: dataset }}
        getText={getText} onEdit={editItem} onEvent={onEvent} onSelector={onSelector} />)}
      {((current.type === "report") && (current.fieldvalue.length>0))?
        <div className="row full">
        {current.fieldvalue.map((row, index) => <Row key={index} row={row} 
          values={row} options={template.options}
          data={{ audit: audit, current: current, dataset: dataset }}
          getText={getText} onEdit={editItem} onEvent={onEvent} onSelector={onSelector} />)}
      </div>:null}
    </div>
  )
}

Main.propTypes = {
  current: PropTypes.object.isRequired, 
  template: PropTypes.object.isRequired, 
  dataset: PropTypes.object.isRequired, 
  audit: PropTypes.string.isRequired,
  className: PropTypes.string,
  onEvent: PropTypes.func,
  getText: PropTypes.func,
}

Main.defaultProps = {
  current: {}, 
  template: {}, 
  dataset: {}, 
  audit: "",
  className: "",
  onEvent: undefined,
  getText: undefined,
}

export default Main;