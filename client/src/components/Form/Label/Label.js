import PropTypes from 'prop-types';

import 'styles/style.css';
import styles from './Label.module.css';

export const Label = ({ 
  value, leftIcon, rightIcon, className, center, style, iconWidth,
  ...props 
}) => {
  if(leftIcon){
    return(
      <div {...props} className={`${"row"} ${(center)?styles.centered:""}`}>
        <div className={`${"cell"} ${styles.label_icon_left}`}
          style={{ width: iconWidth}} >{leftIcon}</div>
        <div className={`${"cell"} ${styles.label_info_left} ${className}`}
          style={style} >{value}</div>
      </div>
    )
  }
  if(rightIcon){
    return(
      <div {...props} className="row full">
        <div className={`${styles.label_info_right} ${className}`}
          style={style} >{value}</div>
        <div className={`${"cell"} ${styles.label_icon_right}`}
          style={{ width: iconWidth }} >{rightIcon}</div>
      </div>
    )
  }
  return(
    <span {...props} className={className||""} style={style} >{value}</span>
  )
}

Label.propTypes = {
  /**
   * Label value string
   */
  value: PropTypes.string.isRequired,
  /**
   * SVG icon component
   */
  leftIcon: PropTypes.object,
  /**
   * SVG icon component
   */
  rightIcon: PropTypes.object,
  /**
   * Custom global css class name
   */
  className: PropTypes.string,
  /**
   * Style object values
   */
  style: PropTypes.object,
  /**
   * Centered content
   */
  center: PropTypes.bool.isRequired,
  /**
   * Icon column width
   */
  iconWidth: PropTypes.string,
}

Label.defaultProps = {
  value: "",
  leftIcon: undefined,
  rightIcon: undefined,
  className: "",
  style: {},
  center: false,
  iconWidth: "auto",
}

export default Label;