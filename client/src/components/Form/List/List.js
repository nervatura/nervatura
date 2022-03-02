import { useState, useMemo } from 'react';
import PropTypes from 'prop-types';
import { useTable, usePagination } from 'react-table'

import 'styles/style.css';
import styles from './List.module.css';

import { getSetting } from 'config/app'
import Icon from 'components/Form/Icon'
import Button from 'components/Form/Button'
import Label from 'components/Form/Label'
import Input from 'components/Form/Input'
import Pagination from 'components/Form/Pagination'

export const ListView = ({
  rows, labelAdd, addIcon, editIcon, deleteIcon,
  listFilter, filterPlaceholder, filterValue,
  currentPage, paginationPage, paginationTop, paginatonScroll, hidePaginatonSize,
  onEdit, onAddItem, onDelete, onCurrentPage,
  ...props 
}) => {
  const [ state, setState ] = useState({
    filter: filterValue,
  })

  const columns = useMemo(() => [ { accessor: "list" } ],[])
  const data = useMemo(() => {
    const getValidRow = (row, filter)=>{
      if(String(row.lslabel).toLowerCase().indexOf(filter)>-1 || 
        String(row.lsvalue).toLowerCase().indexOf(filter)>-1){
          return true
      } else {
        return false
      }
    }
    if(state.filter !== ""){
      let _rows = []; let _filter = String(state.filter).toLowerCase()
      rows.forEach(function(frow) {
        if(getValidRow(frow,_filter)){
          _rows.push(frow);
        }
      });
      return _rows;
    } else {
      return rows
    }
  } , [rows, state.filter])

  const { prepareRow, page, canPreviousPage, canNextPage, pageCount,
    gotoPage, nextPage, previousPage, setPageSize,
    state: { pageIndex, pageSize },
  } = useTable(
    { columns, data,
      initialState: { pageIndex: currentPage, pageSize: paginationPage },
    },
    usePagination
  )

  const onPagination = (key, args) => {
    const pevents = {
      gotoPage: gotoPage, nextPage: nextPage, previousPage: previousPage, setPageSize: setPageSize
    }
    pevents[key](...args)
    if(paginatonScroll){
      window.scrollTo(0,0);
    }
    if(onCurrentPage && (key !== "setPageSize")){
      onCurrentPage(args[0])
    }
  }

  const showPaginator = (pageCount > 1)
  return (
    <div {...props}>
      {(listFilter || (showPaginator && paginationTop))?<div>
        {(showPaginator && paginationTop) ?
          <Pagination pageIndex={pageIndex} pageSize={pageSize} pageCount={pageCount} 
            canPreviousPage={canPreviousPage} canNextPage={canNextPage} hidePageSize={hidePaginatonSize}
            onEvent={onPagination} />:null}
        {(listFilter) ? <div className="row full">
          <div className="cell" >
            <Input id="filter" type="text" className={styles.filterInput}
              placeholder={filterPlaceholder} value={state.filter}
              onChange={(value) => setState({ ...state, filter: value })} />
          </div>
          {(onAddItem)?<div className="cell" style={{width:20}} >
            <Button id="btn_add" 
              className={`${"border-button"} ${styles.addButton}`}
              value={<Label className="addLabel" 
                leftIcon={addIcon} value={labelAdd} />} 
              onClick={(event)=>onAddItem(event)} />
          </div>:null}
        </div>: null}
      </div>:null}
      <ul className={`${"list"} ${styles.list}`} >
        {page.map((row, index) => {
          prepareRow(row)
          return <li 
          key={index}
          className={`${"border-bottom"} ${styles.listRow}`} >
          {(onEdit)?<div id={`row_edit_${index}`}
            className={`${styles.editCell}`} onClick={()=>onEdit(row.original)} >
            {editIcon}
          </div>:null}
          <div id={`row_item_${index}`}
            className={`${styles.valueCell} ${(onEdit)?styles.cursor:""}`} 
            onClick={()=>(onEdit)?onEdit(row.original):null}>
            <div className={`${"border-bottom"} ${styles.label}`} >
              <span>{row.original.lslabel}</span>
            </div>
            <div className={`${styles.value}`} >
              <span>{row.original.lsvalue}</span>
            </div>
          </div>
          {(onDelete)?<div id={`row_delete_${index}`}
            className={`${styles.deleteCell}`} onClick={()=>onDelete(row.original)} >
            {deleteIcon}
          </div>:null}
        </li>})}
      </ul>
      {(showPaginator && !paginationTop) ? <div className="padding-tiny">
        <Pagination pageIndex={pageIndex} pageSize={pageSize} pageCount={pageCount} 
          canPreviousPage={canPreviousPage} canNextPage={canNextPage} hidePageSize={hidePaginatonSize}
          onEvent={onPagination} /></div>:null}
    </div>
  )
}

ListView.propTypes = {
  /**
   * List rows
   */
  rows: PropTypes.arrayOf(PropTypes.shape({
    lslabel: PropTypes.string.isRequired,
    lsvalue: PropTypes.string.isRequired,
  })).isRequired,
  /**
   * Current pagination page number
   */
  currentPage: PropTypes.number.isRequired,
  /**
   * Pagination row number / page
   */
  paginationPage: PropTypes.number.isRequired,
  /**
   * at the top or bottom of the list
   */
  paginationTop: PropTypes.bool.isRequired,
  /**
   * Scroll to top after pagination change
   */
  paginatonScroll: PropTypes.bool.isRequired,
  hidePaginatonSize: PropTypes.bool.isRequired,
  /**
   * Show/hide filter input 
   */ 
  listFilter: PropTypes.bool.isRequired,
  /**
   * Filter input placeholder 
   */ 
  filterPlaceholder: PropTypes.string,
  /**
   * Filter init value
   */
  filterValue: PropTypes.string,
  /**
   * Add button label
   */
  labelAdd: PropTypes.string.isRequired,
  /**
   * Icon element (svg)
   */   
  addIcon: PropTypes.any.isRequired,
  /**
   * Icon element (svg)
   */
  editIcon: PropTypes.any.isRequired,
  /**
   * Icon element (svg)
   */ 
  deleteIcon: PropTypes.any.isRequired,
  /**
   * List row click handle
   */
  onEdit: PropTypes.func,
  /**
   * Add new row handle
   */
  onAddItem: PropTypes.func,
  /**
   * Delete row handle 
   */ 
  onDelete: PropTypes.func,
  /**
   * Paginator select handle 
   */ 
  onCurrentPage: PropTypes.func,
}

ListView.defaultProps = {
  rows: [],
  currentPage: 0,
  paginationPage: getSetting("paginationPage"),
  paginationTop: true,
  paginatonScroll: false,
  hidePaginatonSize: false,
  listFilter: true,
  filterPlaceholder: undefined,
  filterValue: "",
  labelAdd: "",
  addIcon: <Icon iconKey="Plus" />,
  editIcon: <Icon iconKey="Edit" width={24} height={21.3} />,
  deleteIcon: <Icon iconKey="Times" width={19} height={27.6} />,
  onEdit: undefined,
  onAddItem: undefined,
  onDelete: undefined,
  onCurrentPage: undefined
}

export default ListView;