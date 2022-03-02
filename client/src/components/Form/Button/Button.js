import PropTypes from 'prop-types';
import styles from './Button.module.css';

export const Button = ({ 
  label, value, small, className, onClick,
  ...props 
}) => {
  return <button 
    {...props} onClick={onClick}
    className={`${className} ${styles.buttonStyle} ${(small)?styles.smallButton:""}`} >{(typeof value === "undefined")?label:value}</button>
}

Button.propTypes = {
  /**
   * String label value
   */
  label: PropTypes.string.isRequired,
  /**
   * Component content
   */
  value: PropTypes.object,
  /**
   * Button style
   */
  className: PropTypes.string,
  /**
   * Small button
   */
  small: PropTypes.bool.isRequired,
  /**
   * click event handler
   */
  onClick: PropTypes.func,
}

Button.defaultProps = {
  label: "",
  value: undefined,
  className: "",
  small: false,
  onClick: undefined
}

export default Button;