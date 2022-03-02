import { useState } from 'react';
import PropTypes from 'prop-types';
import update from 'immutability-helper';

import { getSetting } from 'config/app'
import 'styles/style.css';
import styles from './Selector.module.css';

import Label from 'components/Form/Label'
import Button from 'components/Form/Button'
import Input from 'components/Form/Input'
import Icon from 'components/Form/Icon'
import Table from 'components/Form/Table'

export const Selector = ({
  view, columns, result, filter,
  selectorPage, paginationPage, currentPage, className,
  getText, onClose, onSelect, onSearch, onCurrentPage, 
  ...props 
}) => {
  const [ state, setState ] = useState({
    filter: filter,
  })
  let fields = {
    view: { columnDef: { 
      id: "view",
      Header: "",
      headerStyle: {},
      Cell: ({ row, value }) => {
        if(row.original.deleted === 1)
          return <Icon iconKey="ExclamationTriangle" className={styles.exclamation} />
        return <Icon iconKey="CaretRight" width={9} height={24} />
      },
      cellStyle: { width: 25, padding: "7px 2px 3px 8px" }
    }}
  }
  columns.forEach(field => {
    fields = update(fields, {$merge: {
      [field[0]]: {
        fieldtype:'string', 
        label: getText(view+"_"+field[0])
      }
    }})
  });
  const selectorView = () => {
    return(
      <div className={`${styles.panel} ${className}`} >
        <div className={`${styles.panelTitle} ${"primary"}`}>
          {(onClose)?<div className="row full">
            <div className="cell">
              <Label value={getText("search_"+view)} leftIcon={<Icon iconKey="Search" />} iconWidth="20px" />
            </div>
            <div className={`${"cell align-right"} ${styles.closeIcon}`}>
              <Icon id="closeIcon" iconKey="Times" onClick={onClose} />
            </div>
          </div>:
          <Label 
            value={getText("quick_search")+": "+getText("search_"+view)} />}
        </div>
        <div className="section" >
          <div className="row full container section-small-bottom" >
            <div className="cell" >
              <Input id="filter"
                type="text" className="full" 
                placeholder={getText("placeholder_search")} 
                autoFocus={true}
                value={state.filter} 
                onEnter={()=>onSearch(state.filter)} 
                onChange={(value)=>setState({ ...state, filter: value })} 
              />
            </div>
            <div className={`${"cell"} ${styles.searchCol}`} >
              <Button id="btn_search" className={`${"full medium"}`} 
                onClick={()=>onSearch(state.filter)}
                value={<Label 
                  value={getText("label_search")} leftIcon={<Icon iconKey="Search" />} center />}
              />
            </div>
          </div>
          <div className="row full container section-small-bottom" >
            <Table fields={fields} rows={result}
              filterPlaceholder={getText("placeholder_search")} 
              paginationPage={(onClose) ? selectorPage : paginationPage} 
              paginationTop={true} hidePaginatonSize={true}
              currentPage={currentPage} onCurrentPage={onCurrentPage}
              onRowSelected={(row)=>onSelect(row, state.filter)} />
          </div>
        </div>
      </div>
    )
  }
  if(onClose){
    return(
      <div className={`${"modal"} ${styles.modal}`} >
        <div className={`${"dialog"} ${styles.dialog}`} {...props} >
          {selectorView()}
        </div>
      </div>
    )
  }
  return selectorView()
}

Selector.propTypes = {
  /**
   * Selector type
   */
  view: PropTypes.string.isRequired,
  /**
   * Table columns def. 
   */ 
  columns: PropTypes.array.isRequired,
  /**
   * Result data rows 
   */ 
  result: PropTypes.array.isRequired,
  /**
   * Filter value init
   */
  filter: PropTypes.string.isRequired,
  /**
   * Modal(small size) pagination row number / page
   */
  selectorPage: PropTypes.number.isRequired, 
  /**
   * Pagination row number / page
   */
  paginationPage: PropTypes.number.isRequired,
  /**
   * Current pagination page number
   */
  currentPage: PropTypes.number.isRequired,
  className: PropTypes.string.isRequired,
  /**
   * Paginator select handle 
   */
  onCurrentPage: PropTypes.func,
  /**
   * Localization
   */
  getText: PropTypes.func, 
  /**
   * Close form handle (modal style)
   */ 
  onClose: PropTypes.func,
  /**
   * Set filter value
   */
  onSearch: PropTypes.func,
  /**
   * onSelect row handle 
   */ 
  onSelect: PropTypes.func,
}

Selector.defaultProps = {
  view: "",
  columns: [],
  result: [],
  filter: "",
  selectorPage: getSetting("selectorPage"),
  paginationPage: getSetting("paginationPage"),
  currentPage: 0,
  className: "",
  getText: undefined,
  onClose: undefined,
  onSearch: undefined,
  onSelect: undefined,
  onCurrentPage: undefined
}

export default Selector;