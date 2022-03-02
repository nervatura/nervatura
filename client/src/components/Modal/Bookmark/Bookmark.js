import { useState } from 'react';
import PropTypes from 'prop-types';

import parseISO from 'date-fns/parseISO'
import format from 'date-fns/format'

import { getSetting } from 'config/app'
import 'styles/style.css';
import styles from './Bookmark.module.css';

import Label from 'components/Form/Label'
import List from 'components/Form/List'
import Button from 'components/Form/Button'
import Icon from 'components/Form/Icon'

export const Bookmark = ({
  bookmark, tabView, paginationPage, dateFormat, timeFormat,
  getText, onClose, onSelect, onDelete, className,
  ...props 
}) => {
  const setBookmark = ()=> bookmark.bookmark.map(item => {
    let bvalue = JSON.parse(item.cfvalue)
    let value = {
      bookmark_id: item.id,
      id: bvalue.id,
      cfgroup: item.cfgroup,
      ntype: bvalue.ntype,
      transtype: (bvalue.ntype === "trans") ? bvalue.transtype : null,
      vkey: bvalue.vkey,
      view: bvalue.view,
      filters: bvalue.filters,
      columns: bvalue.columns,
      lslabel: item.cfname, 
      lsvalue: format(parseISO(bvalue.date), dateFormat)
    }
    if(item.cfgroup === "editor"){
      if (bvalue.ntype==="trans") {
        value.lsvalue += " | " + getText("title_"+bvalue.transtype) + " | " + bvalue.info
      } else {
        value.lsvalue += " | " + getText("title_"+bvalue.ntype) + " | " + bvalue.info
      }
    }
    if(item.cfgroup === "browser"){
      value.lsvalue += " | " + getText("browser_"+bvalue.vkey)
    }
    return value
  })
  const setHistory = ()=> {
    if(bookmark.history && bookmark.history.cfvalue){
      const history_values = JSON.parse(bookmark.history.cfvalue)
      return history_values.map(item => {
        return {
          id: item.id, 
          lslabel: item.title, 
          type: item.type,
          lsvalue: format(parseISO(item.datetime), dateFormat+" "+timeFormat)+" | "+ getText("label_"+item.type,item.type), 
          ntype: item.ntype, transtype: item.transtype
        }
      })
    }
    return []
  }
  const [ state, setState ] = useState({
    tabView: tabView,
    bookmarkList: setBookmark(),
    historyList: setHistory()
  })
  return(
    <div className={`${"modal"} ${styles.modal} ${className}`}  >
        <div className={`${"dialog"} ${styles.dialog}`} {...props} >
          <div className={`${styles.panel}`} >
            <div className={`${styles.panelTitle} ${"primary"}`}>
              <div className="row full">
                <div className="cell">
                  <Label value={getText("title_bookmark")} leftIcon={<Icon iconKey="Star" />} iconWidth="20px" />
                </div>
                <div className={`${"cell align-right"} ${styles.closeIcon}`}>
                  <Icon id="closeIcon" iconKey="Times" onClick={onClose} />
                </div>
              </div>
            </div>
            <div className="section" >
              <div className="row full container section-small-bottom" >
                <div className="cell half" >
                  <Button id="btn_bookmark"
                    className={`${"full"} ${styles.tabButton} ${(state.tabView === "bookmark")?styles.selected:""} ${(state.tabView === "bookmark")?"primary":""}`} 
                    onClick={()=>setState({ ...state, tabView: "bookmark" })} 
                    value={<Label value={getText("title_bookmark_list")} leftIcon={<Icon iconKey="Star" />} />}
                  />
                </div>
                <div className={`${"cell half"}`} >
                  <Button id="btn_history"
                    className={`${"full"} ${styles.tabButton} ${(state.tabView === "history")?styles.selected:""} ${(state.tabView === "history")?"primary":""}`} 
                    onClick={()=>setState({ ...state, tabView: "history" })} 
                    value={<Label value={getText("title_history")} leftIcon={<Icon iconKey="History" />} />}
                  />
                </div>
              </div>
              <div className="row full container section-small-bottom" >
                <List 
                  rows={(state.tabView === "bookmark") ? state.bookmarkList : state.historyList} 
                  editIcon={(state.tabView === "bookmark") ? <Icon iconKey="Star" /> : <Icon iconKey="History" />}
                  listFilter={true} filterPlaceholder={getText("placeholder_filter")}
                  paginationPage={paginationPage} paginationTop={true} hidePaginatonSize={true}
                  onEdit={(row)=>onSelect(state.tabView, row)}  
                  onDelete={(state.tabView === "bookmark") ? (row)=>onDelete(row.bookmark_id) : null} />
              </div>
            </div>
         </div>
      </div>
    </div>
  )
}

Bookmark.propTypes = {
  /**
   * Bookmark/history data
   */
  bookmark: PropTypes.object.isRequired,
  /**
   * Bokkmark view
   */
  tabView: PropTypes.oneOf(["bookmark","history"]).isRequired,
  /**
   * Pagination row number / page
   */
  paginationPage: PropTypes.number.isRequired,
   /**
   * Locale date format
   */
  dateFormat: PropTypes.string,
  /**
   * Locale time format 
   */ 
  timeFormat: PropTypes.string,
  /**
   * onSelect row handle 
   */ 
  onSelect: PropTypes.func,
  /**
   * Delete row handle 
   */ 
  onDelete: PropTypes.func,
  /**
   * Close form handle 
   */ 
  onClose: PropTypes.func,
  /**
   * Localization
   */
  getText: PropTypes.func,
}

Bookmark.defaultProps = {
  bookmark: { history: null, bookmark: [] },
  tabView: "bookmark",
  paginationPage: getSetting("selectorPage"),
  dateFormat: getSetting("dateFormat"),
  timeFormat: getSetting("timeFormat"),
  onSelect: undefined,
  onDelete: undefined,
  onClose: undefined,
  getText: undefined,
}

export default Bookmark;