import PropTypes from 'prop-types';

import styles from './Search.module.css';

import Icon from 'components/Form/Icon'
import Label from 'components/Form/Label'
import Button from 'components/Form/Button'

export const SIDE_VISIBILITY = {
  AUTO: "auto",
  SHOW: "show",
  HIDE: "hide"
}

export const Search = ({ 
  side, groupKey, auditFilter, className,
  getText, onEvent,
  ...props 
}) => {
  const groupButton = (key) => {
    if(key === groupKey){
      return styles.selectButton
    }
    return styles.groupButton
  }
  const searchGroup = (key) => {
    return(
      <div key={key} className="row full">
        <Button id={"btn_group_"+key}
          className={`${"full medium"} ${groupButton(key)}`} 
          onClick={()=>onEvent("changeData", ["group_key", key])}
          value={<Label value={getText("search_"+key)} 
            leftIcon={<Icon iconKey="FileText" />} iconWidth="20px" />}
        />
        {(groupKey === key)?<div className={`${"row full"} ${styles.panelGroup}`} >
          <Button id={"btn_view_"+key}
            className={`${"full medium"} ${styles.panelButton}`} 
            onClick={()=>onEvent("quickView", [key])}
            value={<Label value={getText("quick_search")} 
              leftIcon={<Icon iconKey="Bolt" />} iconWidth="20px" />}
          />
          <Button id={"btn_browser_"+key}
            className={`${"full medium"} ${styles.panelButton}`} 
            onClick={()=>onEvent("showBrowser", [key])} 
            value={<Label value={getText("browser_"+key)} 
              leftIcon={<Icon iconKey="Search" />} iconWidth="20px" />}
          />
        </div>:null}
      </div>
    )
  }
  return (
    <div {...props} 
      className={`${styles.sidebar} ${((side !== "auto")? side : "")} ${className}`} >
      {searchGroup("transitem")}
      {((auditFilter.trans.bank[0]!=="disabled") || (auditFilter.trans.cash[0]!=="disabled"))?
        searchGroup("transpayment"):null}
      {((auditFilter.trans.delivery[0]!=="disabled") || (auditFilter.trans.inventory[0]!=="disabled") 
        || (auditFilter.trans.waybill[0]!=="disabled") || (auditFilter.trans.production[0]!=="disabled")
        || (auditFilter.trans.formula[0]!=="disabled"))?
        searchGroup("transmovement"):null}
      
      <div className={styles.separator} />
      {["customer","product","employee","tool","project"].map(key => {
        if(auditFilter[key][0] !== "disabled") {
          return searchGroup(key)}
        return null
      })}

      <div className={styles.separator} />
      <Button id="btn_report"
        className={`${"full medium"} ${groupButton("report")}`} 
        onClick={()=>{
          onEvent("changeData", ["group_key", "report"]); 
          onEvent("quickView",["report"])
        }}
        value={<Label value={getText("search_report")} 
          leftIcon={<Icon iconKey="ChartBar" />} iconWidth="20px"  />}
      />
      <Button id="btn_office"
        className={`${"full medium"} ${groupButton("office")}`} 
        onClick={()=>onEvent("changeData", ["group_key", "office"])} 
        value={<Label value={getText("search_office")} 
          leftIcon={<Icon iconKey="Inbox" />} iconWidth="20px"  />}
      />
      {(groupKey === "office")?<div className={`${"row full"} ${styles.panelGroup}`} >
        <Button id="btn_printqueue"
          className={`${"full medium primary"} ${styles.panelButton}`} 
          onClick={()=>onEvent("checkEditor", [{ ntype: "printqueue", ttype: null, id: null}])} 
          value={<Label value={getText("title_printqueue")} 
            leftIcon={<Icon iconKey="Print" />} iconWidth="20px"  />}
        />
        <Button id="btn_rate"
          className={`${"full medium primary"} ${styles.panelButton}`} 
          onClick={()=>onEvent("showBrowser",["rate"])} 
          value={<Label value={getText("title_rate")} 
            leftIcon={<Icon iconKey="Globe" />} iconWidth="20px"  />}
        />
        <Button id="btn_servercmd"
          className={`${"full medium primary"} ${styles.panelButton}`} 
          onClick={()=>onEvent("quickView",["servercmd"])} 
          value={<Label value={getText("title_servercmd")} 
            leftIcon={<Icon iconKey="Share" />} iconWidth="20px"  />}
        />
      </div>:null}
    </div>
  )
}

Search.propTypes = {
  /**
   * SideBar visibility
   */
  side: PropTypes.oneOf(Object.values(SIDE_VISIBILITY)).isRequired,
  /**
   * Selected menu group
   */
  groupKey: PropTypes.string,
  /**
   * menu rules 
   */ 
  auditFilter: PropTypes.object.isRequired, 
  className: PropTypes.string, 
  /**
   * Menu selection handle
   */
  onEvent: PropTypes.func,
  /**
   * Localization
   */
  getText: PropTypes.func,
}

Search.defaultProps = {
  side: SIDE_VISIBILITY.AUTO,
  groupKey: "", 
  auditFilter: {}, 
  className: "",  
  onEvent: undefined,
  getText: undefined,
}

export default Search;