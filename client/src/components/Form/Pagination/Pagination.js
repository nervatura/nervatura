import PropTypes from 'prop-types';

import styles from './Pagination.module.css';

import Button from 'components/Form/Button'
import { Input, INPUT_TYPE } from 'components/Form/Input/Input'
import Select from 'components/Form/Select'

export const Pagination = ({
  pageIndex, pageSize, pageCount, canPreviousPage, canNextPage, hidePageSize,
  className, onEvent,
  ...props 
}) => {
  return(
    <div {...props} className={`${"row"} ${className}`} >
      <div className="cell padding-small" >
        <Button id="btn_first" 
          className={`${"border-button"} ${styles.buttonStyle}`}
          label={"1"} disabled={!canPreviousPage}
          onClick={()=>onEvent("gotoPage",[0])} />
        <Button id="btn_previous" 
          className={`${"border-button"} ${styles.buttonStyle}`}
          value={<span>&#10094;</span>} disabled={!canPreviousPage}
          onClick={()=>onEvent("previousPage",[pageIndex-1])} />
      </div>
      <div className="cell" >
        <Input id="input_goto" className={`${styles.inputStyle}`}
          type={INPUT_TYPE.INTEGER} value={pageIndex + 1} disabled={(pageCount === 0)}
          onChange={(value)=>onEvent("gotoPage",[value-1])} />
      </div>
      <div className="cell padding-small" >
        <Button id="btn_next" 
          className={`${"border-button"} ${styles.buttonStyle}`}
          value={<span>&#10095;</span>} disabled={!canNextPage}
          onClick={()=>onEvent("nextPage",[pageIndex+1])} />
        <Button id="btn_last" 
          className={`${"border-button"} ${styles.buttonStyle}`}
          label={String(pageCount)} disabled={!canNextPage}
          onClick={()=>onEvent("gotoPage",[pageCount - 1])} />
      </div>
      {(!hidePageSize)?<div className="cell padding-small hide-small" >
        <Select id="sel_page_size" 
          value={String(pageSize)} className={`${styles.pageSize}`}
          disabled={(pageCount === 0)}
          onChange={(value)=>onEvent("setPageSize",[Number(value)])}
          options={[5, 10, 20, 50, 100].map((pageSize) => {
            return { value: String(pageSize), text: String(pageSize) }
          })}
          />
      </div>:null}
    </div>
  )
}

Pagination.propTypes = {
  pageIndex: PropTypes.number.isRequired, 
  pageSize: PropTypes.number.isRequired, 
  pageCount: PropTypes.number.isRequired,  
  canPreviousPage: PropTypes.bool.isRequired, 
  canNextPage: PropTypes.bool.isRequired,
  hidePageSize: PropTypes.bool.isRequired,
  className: PropTypes.string.isRequired,
  onEvent: PropTypes.func,
}

Pagination.defaultProps = {
  pageIndex: 0, 
  pageSize: 5, 
  pageCount: 0,  
  canPreviousPage: false,
  canNextPage: false,
  hidePageSize: false,
  className: "",
  onEvent: undefined,
}

export default Pagination;