import PropTypes from 'prop-types';
import styles from './Input.module.css';

import { getSetting, DECIMAL_SEPARATOR } from 'config/app'

export const INPUT_TYPE = {
  TEXT: "text",
  INTEGER: "integer",
  NUMBER: "number",
  COLOR: "color",
  FILE: "file",
  PASSWORD: "password",
}

export const Input = ({ 
  type, separator, value, minValue, maxValue, focus, className,
  onChange, onBlur, onEnter,
  ...props 
}) => {
  return <input {...props} 
    type={[INPUT_TYPE.INTEGER,INPUT_TYPE.NUMBER].includes(type)?INPUT_TYPE.TEXT:type}
    value={([INPUT_TYPE.INTEGER,INPUT_TYPE.NUMBER].includes(type))
      ? String(value).replace(DECIMAL_SEPARATOR.POINT,separator) : value}
    className={`${className} ${styles.inputStyle} ${([INPUT_TYPE.INTEGER,INPUT_TYPE.NUMBER].includes(type))?"align-right":""} `}
    onChange={(event) => {
      event.stopPropagation();
      let inputValue = event.target.value
      if([INPUT_TYPE.INTEGER,INPUT_TYPE.NUMBER].includes(type)){
        inputValue = (type === "number") ? 
          String(event.target.value).replace(new RegExp(`[^0-9${separator}-]`, "g"), "") : 
          String(event.target.value).replace(/[^0-9-]|-(?=.)/g,'')
        if(inputValue === ""){
          inputValue = 0
        }
        if(!String(inputValue).endsWith(separator) 
          || (String(inputValue).endsWith(separator) 
            && (String(inputValue).match(new RegExp(`[${separator}]`, "g")).length > 1)) 
          || String(inputValue).endsWith(separator+separator)){
            inputValue = parseFloat(String(inputValue).replace(separator,DECIMAL_SEPARATOR.POINT))
          if(minValue && (inputValue < minValue)){
            inputValue = minValue
          }
          if(maxValue && (inputValue > maxValue)){
            inputValue = maxValue
          }
        }
      }
      onChange(inputValue)
    }}
    onKeyDown={(event)=>{
      event.stopPropagation();
      if((event.keyCode === 13) && onEnter){
        onEnter(value)
      }
    }}
    onBlur={(event) => {
      event.stopPropagation();
      let inputValue = event.target.value
      if([INPUT_TYPE.INTEGER,INPUT_TYPE.NUMBER].includes(type)){
        inputValue = 0
        if((type === "number") 
          && !isNaN(parseFloat(String(event.target.value).replace(separator,DECIMAL_SEPARATOR.POINT)))){
            inputValue = parseFloat(String(event.target.value).replace(separator,DECIMAL_SEPARATOR.POINT))
        }
        if((type === "integer") && !isNaN(parseInt(event.target.value))){
          inputValue = parseInt(event.target.value,10)
        }
        if(!onBlur &&(String(inputValue) !== event.target.value)){
          onChange(inputValue)
        }
      }
      if(onBlur){
        onBlur(inputValue)
      }
    }}
  />
}

Input.propTypes = {
  /**
    * Input type
    */
  type: PropTypes.oneOf(Object.values(INPUT_TYPE)).isRequired,
  /**
   * Number type decimal separator
   */
  separator: PropTypes.oneOf(Object.values(DECIMAL_SEPARATOR)).isRequired,
  /**
   * Input value
   */
  value: PropTypes.oneOfType([PropTypes.number, PropTypes.string]).isRequired,
  /**
   * Number/integer type minimum value limit
   */
  minValue: PropTypes.number,
  /**
  * Number/integer type maximum value limit
  */
  maxValue: PropTypes.number,
  /**
  * Value change event
  */
  onChange: PropTypes.func,
  /**
  * Lost focus event
  */
  onBlur: PropTypes.func,
  /**
  * Enter key event
  */
  onEnter: PropTypes.func,
  /**
   * Input style
   */
  className: PropTypes.string,
}

Input.defaultProps = {
  type: INPUT_TYPE.TEXT,
  separator: getSetting("decimal_sep"),
  value: "",
  minValue: undefined,
  maxValue: undefined,
  className: "",
  onChange: undefined,
  onBlur: undefined,
  onEnter: undefined,
}

export default Input;