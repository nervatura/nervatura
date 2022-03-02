import PropTypes from 'prop-types';

import { toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

toast.configure({});

export const TOAST_TYPE = {
  ERROR: "error",
  WARNING: "warning",
  SUCCESS: "success",
  INFO: "info",
  DEFAULT: "default"
}

export const Toastify = ({ 
  toastType, message, toastTime 
}) => {
  const autoClose = (toastTime === 0) ? false : toastTime
  switch (toastType) {
    case "error":
      return toast.error(message, {
        theme: "colored",
        autoClose: autoClose,
      });
    
    case "warning":
      return toast.warning(message, {
        theme: "colored",
        autoClose: autoClose,
      });
    
    case "success":
      return toast.success(message, {
        theme: "colored",
        autoClose: autoClose,
      });
    
    case "info":
      return toast.info(message, {
        theme: "colored",
        autoClose: autoClose,
      });
  
    default:
      return toast(message, { autoClose: autoClose })
  }
}

Toastify.propTypes = {
  toastType: PropTypes.oneOf(Object.values(TOAST_TYPE)).isRequired, 
  message: PropTypes.string.isRequired, 
  toastTime: PropTypes.number.isRequired
}

Toastify.defaultProps = {
  toastType: TOAST_TYPE.DEFAULT, 
  message: "", 
  toastTime: 0
}

export default Toastify;