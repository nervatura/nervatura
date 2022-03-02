import { useState, Fragment, useMemo } from 'react';
import update from 'immutability-helper';
import PropTypes from 'prop-types';
import { useTable, usePagination, useSortBy } from 'react-table'

import formatISO from 'date-fns/formatISO'
import isValid from 'date-fns/isValid'
import parseISO from 'date-fns/parseISO'
import format from 'date-fns/format'

import Pagination from 'components/Form/Pagination'
import Icon from 'components/Form/Icon'
import Button from 'components/Form/Button'
import Label from 'components/Form/Label'
import Input from 'components/Form/Input'
import { getSetting } from 'config/app'

import styles from './Table.module.css';

export const TableView = ({
  rowKey, rows, fields,
  dateFormat, timeFormat, className,
  labelYes, labelNo, labelAdd, addIcon, tableFilter, filterPlaceholder, filterValue,
  paginationPage, paginationTop, paginatonScroll, hidePaginatonSize,
  currentPage, onCurrentPage, tablePadding, onRowSelected, onEditCell, onAddItem,
  ...props 
}) => {
  const [ state, setState ] = useState({
    filter: filterValue
  })

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

  const columns = useMemo(() => {
    const numberCell = (value, label, style) => {
      return <div className={styles.numberCell}>
        <span className={styles.cellLabel}>{label}</span>
        <span style={style}>{value}</span>
      </div>
    }

    const dateCell = (value, label, dateType) => {
      let fmtValue = ""
      const dateValue = parseISO(value)
      if (isValid(dateValue)) {
        switch (dateType) {
          case "date":
            if(dateFormat){
              fmtValue = format(dateValue, dateFormat)
            } else {
              fmtValue = formatISO(dateValue, { representation: 'date' })
            }   
            break;
          
          case "time":
            if(timeFormat){
              fmtValue = format(dateValue, timeFormat)
            } else {
              fmtValue = formatISO(dateValue, { representation: 'time' })
            }   
            break;
        
          default:
          if(dateFormat && timeFormat){
              fmtValue = format(dateValue, dateFormat+" "+timeFormat)
            } else {
              fmtValue = formatISO(dateValue)
            } 
            break;
        }
      }
      return <Fragment>
        <span className={styles.cellLabel}>{label}</span>
        <span>{fmtValue}</span>
      </Fragment>
    }

    const boolCell = (value, label) => {
      if((value === 1) || (value === "true") || (value === true)){
        return <Fragment>
          <span className={styles.cellLabel}>{label}</span>
          <Icon iconKey="CheckSquare" className={styles.middle} />
          <span className={styles.middle}> {labelYes}</span>
        </Fragment>
      }
      else {
        return <Fragment>
          <span className={styles.cellLabel}>{label}</span>
          <Icon iconKey="SquareEmpty" className={styles.middle} />
          <span className={styles.middle}> {labelNo}</span>
        </Fragment>
      }
    }

    const linkCell = (value, label, fieldname, resultValue, rowData) => {
      return <Fragment>
        <span className={styles.cellLabel}>{label}</span>
        <span id={"link_"+rowData[rowKey]} className={styles.linkCell} 
          onClick={(onEditCell)?(event)=>{
            event.stopPropagation();
            onEditCell(fieldname, resultValue, rowData);
          }:null} >{value}</span>
      </Fragment>
    }

    const stringCell = (value, label, style) => {
      return <Fragment>
        <span className={styles.cellLabel}>{label}</span>
        <span style={style} >{value}</span>
      </Fragment>
    }

    let cols = []
    Object.keys(fields).forEach((fieldname) => {
      if(fields[fieldname].columnDef){
        cols = update(cols, {$push: [fields[fieldname].columnDef]});
      } else {
        let coldef = {
          accessor: fieldname,
          Header: fields[fieldname].label || "",
          headerStyle: {},
          cellStyle: {}
        }
        switch (fields[fieldname].fieldtype) {
          case "number":
            coldef.headerStyle.textAlign = "right"
            coldef.Cell = ({ row, value }) => {
              let style = {}
              if(fields[fieldname].format){
                style.fontWeight = "bold";
                if(row.original.edited){
                  style.textDecoration = "line-through";
                } else if(value !== 0){
                  style.color = "red"
                } else {
                  style.color = "green"
                }
              }
              return numberCell(value, fields[fieldname].label, style)
            }
            break;
          
          case "datetime":
          case "date":
          case "time":
            coldef.Cell = ({ value }) => {
              return dateCell(value, fields[fieldname].label, fields[fieldname].fieldtype)
            }
            break;
          
          case "bool":
            coldef.Cell = ({ value }) => {
              return boolCell(value, fields[fieldname].label)
            }
            break;

          case "deffield":
            coldef.Cell = ({ row, value }) => {
              switch (row.original.fieldtype) {
                case "bool":
                  return boolCell(value, fields[fieldname].label)

                case "integer":
                case "float":
                  return numberCell(value, fields[fieldname].label, {})

                case "customer":
                case "tool":
                case "product":
                case "trans": 
                case "transitem":
                case "transmovement": 
                case "transpayment":
                case "project":
                case "employee":
                case "place":
                case "urlink":
                  return linkCell(
                    row.original["export_deffield_value"], fields[fieldname].label, 
                    row.original.fieldtype, row.original[fieldname], row.original)

                default:
                  return stringCell(value, fields[fieldname].label, {})
              }
            };
            break;

          default:
            coldef.Cell = ({ row, value }) => {
              let style = {}
              if(row.original[fieldname+"_color"]){
                style.color = row.original[fieldname+"_color"]
              }
              if(Object.keys(row.original).includes("export_"+fieldname)){
                return linkCell(
                  row.original["export_"+fieldname], fields[fieldname].label, fieldname, row.original[fieldname], row.original)
              }
              return stringCell(value, fields[fieldname].label, style)
            }
        }
        if(tablePadding){
          coldef.headerStyle.padding = tablePadding
          coldef.cellStyle.padding = tablePadding
        }
        if(fields[fieldname].verticalAlign) {
          coldef.cellStyle.verticalAlign = fields[fieldname].verticalAlign
        }
        if(fields[fieldname].textAlign) {
          coldef.cellStyle.textAlign = fields[fieldname].textAlign
        }
        cols = update(cols, {$push: [coldef]});
      }
    })
    return cols
  },[fields, dateFormat, timeFormat, labelNo, labelYes, onEditCell, rowKey, tablePadding])

  const data = useMemo(() => {
    const getValidRow = (row, filter)=>{
      let find = false;
      Object.keys(fields).forEach((fieldname) => {
        if(String(row[fieldname]).toLowerCase().indexOf(filter)>-1){
          find = true;
        }
      });
      return find;
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
  }, [fields, rows, state.filter])

  const { getTableProps, getTableBodyProps, headerGroups,
    prepareRow, page, canPreviousPage, canNextPage, pageCount,
    gotoPage, nextPage, previousPage, setPageSize,
    state: { pageIndex, pageSize },
  } = useTable(
    { columns, data,
      initialState: { pageIndex: currentPage, pageSize: paginationPage },
    },
    useSortBy, usePagination
  )

  const showPaginator = (pageCount > 1)
  return(
    <div {...props} className={` ${styles.responsive} ${className}`} >

      {(tableFilter || (showPaginator && paginationTop))?<div>
        {(showPaginator && paginationTop) ?
          <Pagination pageIndex={pageIndex} pageSize={pageSize} pageCount={pageCount} 
            canPreviousPage={canPreviousPage} canNextPage={canNextPage} hidePageSize={hidePaginatonSize}
            onEvent={onPagination} />:null}
        {(tableFilter) ? <div className="row full">
          <div className="cell" >
            <Input id="filter" type="text" className={styles.filterInput}
              placeholder={filterPlaceholder} value={state.filter}
              onChange={(value) => setState({ ...state, filter: value })} />
          </div>
          {(onAddItem)?<div className="cell" style={{width:20}} >
            <Button id="btn_add" 
              className={`${"border-button"} ${styles.addButton}`}
              value={<Label className={styles.addLabel} 
                leftIcon={addIcon} value={labelAdd} />} 
              onClick={(event)=>onAddItem(event)} />
          </div>:null}
        </div>: null}
      </div>:null}

      <div className={styles.tableWrap}>
        <table {...getTableProps()} className={`ui-table ${styles.uiTable}`} >
          <thead>
            {headerGroups.map(headerGroup => (
              <tr {...headerGroup.getHeaderGroupProps()}>
                {headerGroup.headers.map(column => (
                  <th {...column.getHeaderProps(column.getSortByToggleProps())} 
                    style={column.headerStyle} 
                    className={`${styles.sort} ${
                      (column.isSorted)
                        ?(column.isSortedDesc)
                          ? styles.sortDesc 
                          : styles.sortAsc
                        : styles.sortNone
                    }`} >
                    {column.render('Header')}
                  </th>
                ))}
              </tr>
            ))}
          </thead>
          <tbody {...getTableBodyProps()}>
            {page.map(
              (row, i) => {
                prepareRow(row);
                return (
                  <tr {...row.getRowProps()} 
                    className={(row.original.disabled) ? styles.cursorDisabled : (onRowSelected) ? styles.cursorPointer : ""}
                    onClick={(onRowSelected && !row.original.disabled) ? () => onRowSelected(row.original, i) : null}>
                    {row.cells.map(cell => {
                      return (
                        <td {...cell.getCellProps()} style={cell.column.cellStyle} >{cell.render('Cell')}</td>
                      )
                    })}
                  </tr>
                )}
            )}
          </tbody>
        </table>
      </div>

      {(showPaginator && !paginationTop) ? <div>
        <Pagination pageIndex={pageIndex} pageSize={pageSize} pageCount={pageCount} 
          canPreviousPage={canPreviousPage} canNextPage={canNextPage} hidePageSize={hidePaginatonSize}
          onEvent={onPagination} /></div>:null}
      
    </div>
  )
}

TableView.propTypes = {
  /**
   * Table row unique identifier
   */
  rowKey: PropTypes.string.isRequired,
  /**
   * Table data rows
   */
  rows: PropTypes.array.isRequired,
  /**
   * Table columns def.
   */
  fields: PropTypes.object.isRequired,
  /**
   * Locale date format
   */
  dateFormat: PropTypes.string,
  /**
   * Locale time format 
   */ 
  timeFormat: PropTypes.string,
  /**
   * Table header/cell padding value
   */
  tablePadding: PropTypes.number,
  labelYes: PropTypes.string.isRequired,
  labelNo: PropTypes.string.isRequired,
  /**
   * Add button label
   */
  labelAdd: PropTypes.string.isRequired, 
  /**
   * Icon element (svg)
   */   
  addIcon: PropTypes.any.isRequired,
  /**
   * Show/hide filter input 
   */
  tableFilter: PropTypes.bool.isRequired, 
  /**
   * Filter input placeholder 
   */ 
  filterPlaceholder: PropTypes.string, 
  /**
   * Filter init value
   */
  filterValue: PropTypes.string,
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
   * Current pagination page number
   */
  currentPage: PropTypes.number.isRequired,
  /**
   * Paginator select handle 
   */ 
  onCurrentPage: PropTypes.func,
  /**
   * List row click handle
   */ 
  onRowSelected: PropTypes.func,
  /**
   * Table cell click handle
   */
  onEditCell: PropTypes.func,
  /**
   * Add new row handle
   */
  onAddItem: PropTypes.func,
}

TableView.defaultProps = {
  rowKey: "id",
  rows: [],
  fields: undefined,
  dateFormat: getSetting("dateFormat"),
  timeFormat: getSetting("timeFormat"),
  labelYes: "YES",
  labelNo: "NO",
  labelAdd: "",
  addIcon: <Icon iconKey="Plus" />,
  tableFilter: false,
  filterPlaceholder: undefined,
  filterValue: "",
  paginationPage: 10,
  currentPage: 0,
  paginationTop: true,
  paginatonScroll: false,
  hidePaginatonSize: false,
  tablePadding: undefined,
  className: "",
  onRowSelected: undefined,
  onEditCell: undefined,
  onAddItem: undefined
}

export default TableView;