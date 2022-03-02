import PropTypes from 'prop-types';

import styles from './Select.module.css';

export const Select = ({ 
  options, placeholder, className, onChange,
  ...props 
}) => {
  let values = options.map((item,index)=><option
    className={`${styles.optionStyle}`} 
    key={index} value={item.value} >{item.text}</option>)
  if(typeof placeholder !== "undefined"){
    values.unshift(<option className={` ${styles.optionStyle} ${styles.optionPlaceholder}` }
      key="placeholder" value="" >{placeholder}</option>)
  }
  return <select {...props} className={`${className} ${styles.selectStyle}`}
    onChange={(event) => onChange(event.target.value)} >{values}</select>
}

Select.propTypes = {
  /**
   * Select elements
   */
  options: PropTypes.arrayOf(PropTypes.shape({
    value: PropTypes.string.isRequired,
    text: PropTypes.string.isRequired
  })),
  /**
   * Select placeholder text
   */
  placeholder: PropTypes.string,
  /**
   * Input style
   */
  className: PropTypes.string,
  /**
  * Value change event
  */
   onChange: PropTypes.func,
}

Select.defaultProps = {
  options: [],
  placeholder: undefined,
  className: "",
  onChange: undefined
}

export default Select;