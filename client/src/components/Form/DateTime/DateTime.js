import PropTypes from 'prop-types';

import Flatpickr from "react-flatpickr";
import formatISO from 'date-fns/formatISO'
import isValid from 'date-fns/isValid'
import parseISO from 'date-fns/parseISO'
import isEqual from 'date-fns/isEqual'

import { German } from "flatpickr/dist/l10n/de.js"
import { Spanish } from "flatpickr/dist/l10n/es.js"
import { French } from "flatpickr/dist/l10n/fr.js"
import { Italian } from "flatpickr/dist/l10n/it.js"
import { Portuguese } from "flatpickr/dist/l10n/pt.js"

import "flatpickr/dist/themes/dark.css";
import './DateTime.css';

import { getSetting } from 'config/app'

const locales = {
  de: German, es: Spanish, fr: French, it: Italian, pt: Portuguese 
}

const dateStyle = [
  ["yyyy-MM-dd", "Y-m-d"], 
  ["dd-MM-yyyy", "d-m-Y"], 
  ["MM-dd-yyyy", "m-d-Y"]
]

export const DateTime = ({ 
  value, dateTime, isEmpty, showTimeSelectOnly,
  dateFormat, timeFormat, locale, className,
  onChange,
  ...props 
}) => {
  let calDateFormat = dateStyle.filter(value => (value[0] === dateFormat))[0][1]
  if(dateTime){
    calDateFormat += " H:i"
  }
  const selectedDate = () => {
    let dateValue = (value) ? parseISO(value) : null
    if(value && !isValid(dateValue) && showTimeSelectOnly){
      dateValue = new Date(formatISO(new Date(), { representation: 'date' })+"T"+value)
    }
    return dateValue
  } 

  const setValue = ( selectedDate ) => {
    if(onChange){
      if(typeof selectedDate === "undefined"){
        if(isEmpty){
          return onChange(null)
        }
        return onChange(value)
      }
      if(showTimeSelectOnly){
        return onChange(formatISO(selectedDate, { representation: 'time' }))  
      }
      if(!dateTime){
        return onChange(formatISO(selectedDate, { representation: 'date' }))  
      }
      return onChange(formatISO(selectedDate))
    }
  }
  return(
    <Flatpickr {...props} className={`${className}`}
      options={{ 
        allowInput: true, 
        time_24hr: true,
        enableTime: (dateTime || showTimeSelectOnly),
        noCalendar: showTimeSelectOnly,
        locale: (locale !== "en") ? locales[locale] : null,
        dateFormat: calDateFormat,
        onChange: (selectedDates) => {
          if(!isEqual(selectedDates[0], selectedDate())){
            setValue( selectedDates[0] )
          }
        }
      }}
      value={selectedDate()}
    />
  )
}

DateTime.propTypes = {
  /**
   * Calendar selected value
   */
  value: PropTypes.string,
  /**
   * Date or Datetime input
   */ 
  dateTime: PropTypes.bool.isRequired,
  /**
   * Enabled empty (null) value 
   */ 
  isEmpty: PropTypes.bool.isRequired,
  /**
   * Time input 
   */ 
  showTimeSelectOnly: PropTypes.bool.isRequired,
  /**
   * Locale date format
   */
  dateFormat: PropTypes.string,
  /**
    * Locale time format 
    */ 
  timeFormat: PropTypes.string,
  /**
   * Calendar locale
   */
  locale: PropTypes.string,
  /**
   * onChange handle
   */
  onChange: PropTypes.func,
}

DateTime.defaultProps = {
  value: undefined,
  dateTime: true,
  isEmpty: true,
  showTimeSelectOnly: false,
  dateFormat: getSetting("dateFormat"),
  timeFormat: getSetting("timeFormat"),
  locale: getSetting("calendar"),
  onChange: undefined
}

export default DateTime;