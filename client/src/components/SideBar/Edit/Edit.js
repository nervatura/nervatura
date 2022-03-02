import PropTypes from 'prop-types';

import styles from './Edit.module.css';

import Icon from 'components/Form/Icon'
import Label from 'components/Form/Label'
import Button from 'components/Form/Button'

export const SIDE_VISIBILITY = {
  AUTO: "auto",
  SHOW: "show",
  HIDE: "hide"
}

export const Edit = ({ 
  side, newFilter, auditFilter, edit, module, forms,
  className,
  getText, onEvent,
  ...props 
}) => {
  const { current, form_dirty, dirty, panel, dataset, group_key } = module
  const itemMenu = (keyValue, classValue, eventValue, labelValue, disabledValue) => {
    return <Button id={keyValue} key={keyValue}
      className={classValue} disabled={(disabledValue)?"disabled":""}
      onClick={ ()=>onEvent(...eventValue) }
      value={labelValue}
    />
  }
  const groupMenu = (keyValue, classValue, groupKey, labelValue) => {
    return <Button id={keyValue} key={keyValue}
      className={classValue}
      onClick={ ()=>onEvent("changeData", ["group_key", groupKey]) }
      value={labelValue}
    />
  }
  const editItems = (options)=>{
    if (typeof options === "undefined") {
      options = {}
    }
    let panels = []

    if (options.back === true || current.form) {
      panels.push(
        itemMenu("cmd_back",
          `${"medium"} ${styles.itemButton} ${styles.selected}`, 
          ["editorBack"],
          <Label value={getText("label_back")} 
            leftIcon={<Icon iconKey="Reply" />} iconWidth="20px"  />
        )
      )
      panels.push(<div key="back_sep" className={styles.separator} />)
    }

    if (options.arrow === true) {
      panels.push(
        itemMenu("cmd_arrow_left", 
          `${"full medium"} ${styles.itemButton}`, 
          ["prevTransNumber"],
          <Label value={getText("label_previous")} 
            leftIcon={<Icon iconKey="ArrowLeft" />} iconWidth="20px"  />
        )
      )
      panels.push(
        itemMenu("cmd_arrow_right",
          `${"full medium"} ${styles.itemButton}`, 
          ["nextTransNumber"],
          <Label value={getText("label_next")} 
            rightIcon={<Icon iconKey="ArrowRight" />} iconWidth="20px"  />
        )
      )
      panels.push(<div key="arrow_sep" className={styles.separator} />)
    }

    if (options.state && options.state !== "normal") {
      const color = (options.state === "deleted") 
        ? "red" : (options.state === "cancellation") 
          ? "orange" : "white"
      const icon = (["closed", "readonly"].includes(options.state)) 
        ? "Lock" : "ExclamationTriangle"
      panels.push(<div key="cmd_state" className={`${"full padding-small large"} ${styles.stateLabel}`} >
        <Label value={getText("label_"+options.state)} style={{ color: color }}
          leftIcon={<Icon iconKey={icon} color={color} />} iconWidth="25px"  />
      </div>)
      panels.push(<div key="state_sep" className={styles.separator} />)
    }

    if (options.save !== false) {
      panels.push(
        itemMenu("cmd_save",
          `${"full medium"} ${styles.itemButton} ${((current.form && form_dirty)||(!current.form && dirty))?styles.selected:""}`,
          ["saveEditor"],
          <Label value={getText("label_save")} 
            leftIcon={<Icon iconKey="Check" />} iconWidth="20px"  />
        )
      )
    }
    if (options.delete !== false && options.state === "normal") {
      panels.push(
        itemMenu("cmd_delete",
          `${"full medium"} ${styles.itemButton}`, 
          ["editorDelete"],
          <Label value={getText("label_delete")} 
            leftIcon={<Icon iconKey="Times" />} iconWidth="20px"  />
        )
      )
    }
    if (options.new !== false && options.state === "normal" && !current.form) {
      panels.push(
        itemMenu("cmd_new",
          `${"full medium"} ${styles.itemButton}`, 
          ["editorNew",[{}]],
          <Label value={getText("label_new")} 
            leftIcon={<Icon iconKey="Plus" />} iconWidth="20px"  />
        )
      )
    }

    if (options.trans === true) {
      panels.push(<div key="trans_sep" className={styles.separator} />)
      if (options.copy !== false) {
        panels.push(
          itemMenu("cmd_copy",
            `${"full medium"} ${styles.itemButton}`,
            ["transCopy",["normal"]],
              <Label value={getText("label_copy")} 
                leftIcon={<Icon iconKey="Copy" />} iconWidth="20px" />
          )
        );
      }
      if (options.create !== false) {
        panels.push(
          itemMenu("cmd_create",
            `${"full medium"} ${styles.itemButton}`,
            ["transCopy",["create"]],
            <Label value={getText("label_create")} 
              leftIcon={<Icon iconKey="Sitemap" />} iconWidth="20px" />
          )
        );
      }
      if (options.corrective === true && options.state === "normal") {
        panels.push(
          itemMenu("cmd_corrective",
            `${"full medium"} ${styles.itemButton}`,
            ["transCopy",["amendment"]],
            <Label value={getText("label_corrective")} 
              leftIcon={<Icon iconKey="Share" />} iconWidth="20px" />
          )
        );
      }
      if (options.cancellation === true && options.state !== "cancellation") {
        panels.push(
          itemMenu("cmd_cancellation",
            `${"full medium"} ${styles.itemButton}`,
            ["transCopy",["cancellation"]],
            <Label value={getText("label_cancellation")} 
              leftIcon={<Icon iconKey="Undo" />} iconWidth="20px" />
          )
        );
      }
      if (options.formula === true) {
        panels.push(
          itemMenu("cmd_formula",
            `${"full medium"} ${styles.itemButton}`, 
            ["checkEditor", [{}, 'LOAD_FORMULA', undefined]],
          <Label value={getText("label_formula")} 
            leftIcon={<Icon iconKey="Magic" />} iconWidth="20px"  />
          )    
        )
      }
    }

    if (options.link === true) {
      panels.push(
        itemMenu("cmd_link",
          `${"full medium"} ${styles.itemButton}`, 
          ["setLink",[options.link_type, options.link_field]],
          <Label value={options.link_label} 
            leftIcon={<Icon iconKey="Link" />} iconWidth="20px"  />
        )
      )
    }

    if (options.password === true) {
      panels.push(
        itemMenu("cmd_password",
          `${"full medium"} ${styles.itemButton}`, 
          ["setPassword"],
          <Label value={getText("title_password")} 
            leftIcon={<Icon iconKey="Lock" />} iconWidth="20px"  />
        )
      )
    }

    if (options.shipping === true) {
      panels.push(
        itemMenu("cmd_shipping_all",
          `${"full medium"} ${styles.itemButton}`, 
          ["shippingAddAll"],
          <Label value={getText("shipping_all_label")} 
            leftIcon={<Icon iconKey="Plus" />} iconWidth="20px"  />
        )
      )
      panels.push(
        itemMenu("cmd_shipping_create",
          `${"full medium"} ${styles.itemButton} ${(dataset.shiptemp.length > 0)?styles.selected:""}`,
          ["createShipping"],
          <Label value={getText("shipping_create_label")} 
            leftIcon={<Icon iconKey="Check" />} iconWidth="20px"  />
        )
      )
    }

    if (options.more === true) {
      panels.push(<div key="more_sep_1" className={styles.separator} />)
      if (options.report !== false) {
        panels.push(
          itemMenu("cmd_report",
            `${"full medium"} ${styles.itemButton}`, 
            ["reportSettings", []],
            <Label value={getText("label_report")} 
              leftIcon={<Icon iconKey="ChartBar" />} iconWidth="20px"  />
          )
        )
      }
      if (options.search === true) {
        panels.push(
          itemMenu("cmd_search",
            `${"full medium"} ${styles.itemButton}`, 
            ["searchQueue"],
            <Label value={getText("label_search")} 
              leftIcon={<Icon iconKey="Search" />} iconWidth="20px"  />
          )
        )
      }
      if (options.export_all === true && options.state === "normal") {
        panels.push(
          itemMenu("cmd_export_all",
            `${"full medium"} ${styles.itemButton}`, 
             ["exportQueueAll"],
            <Label value={getText("label_export_all")} 
              leftIcon={<Icon iconKey="Download" />} iconWidth="20px"  />
          )
        )
      }
      if (options.print === true) {
        panels.push(
          itemMenu("cmd_print",
            `${"full medium"} ${styles.itemButton}`, 
            ["createReport",["print"]],
            <Label value={getText("label_print")} 
              leftIcon={<Icon iconKey="Print" />} iconWidth="20px"  />
          )
        )
      }
      if (options.export_pdf === true && options.state === "normal") {
        panels.push(
          itemMenu("cmd_export_pdf",
            `${"full medium"} ${styles.itemButton}`, 
            ["createReport",["pdf"]],
          <Label value={getText("label_export_pdf")} 
            leftIcon={<Icon iconKey="Download" />} iconWidth="20px"  />
          )
        )
      }
      if (options.export_xml === true && options.state === "normal") {
        panels.push(
          itemMenu("cmd_export_xml",
            `${"full medium"} ${styles.itemButton}`, 
            ["createReport",["xml"]],
            <Label value={getText("label_export_xml")} 
              leftIcon={<Icon iconKey="Code" />} iconWidth="20px"  />
          )
        )
      }
      if (options.export_csv === true && options.state === "normal") {
        panels.push(
          itemMenu("cmd_export_csv",
            `${"full medium"} ${styles.itemButton}`, 
            ["createReport",["csv"]],
            <Label value={getText("label_export_csv")} 
              leftIcon={<Icon iconKey="Download" />} iconWidth="20px"  />
          )
        )
      }
      if (options.export_event === true && options.state === "normal") {
        panels.push(
          itemMenu("cmd_export_event",
            `${"full medium"} ${styles.itemButton}`, 
            ["exportEvent"],
            <Label value={getText("label_export_event")} 
              leftIcon={<Icon iconKey="Calendar" />} iconWidth="20px"  />
          )
        )
      }
      panels.push(<div key="more_sep_2" className={styles.separator} />)
      if (options.bookmark !== false && options.state === "normal") {
        panels.push(
          itemMenu("cmd_bookmark",
            `${"full medium"} ${styles.itemButton}`, 
            ["saveBookmark",[options.bookmark]],
            <Label value={getText("label_bookmark")} 
              leftIcon={<Icon iconKey="Star" />} iconWidth="20px"  />
          )
        )
      }
      if (options.help !== false) {
        panels.push(
          itemMenu("cmd_help",
            `${"full medium"} ${styles.itemButton}`, 
            ["showHelp",[options.help]],
            <Label value={getText("label_help")} 
              leftIcon={<Icon iconKey="QuestionCircle" />} iconWidth="20px"  />
          )
        )
      }
    }

    if (options.more !== true && typeof options.help !== "undefined") {
      panels.push(<div key="help_sep" className={styles.separator} />)
      panels.push(
        itemMenu("cmd_help",
          `${"full medium"} ${styles.itemButton}`, 
          ["showHelp",[options.help]],
          <Label value={getText("label_help")} 
            leftIcon={<Icon iconKey="QuestionCircle" />} iconWidth="20px"  />
        )
      )
    }
    
    return panels
  }
  const groupButton = (key) => {
    if(key === group_key){
      return styles.selectButton
    }
    return styles.groupButton
  }
  const newItems = ()=>{
    let mnu_items = []

    if(newFilter[0].length > 0){
      mnu_items.push(<div key="0" className="row full">
        {groupMenu("new_transitem_group", 
          `${"full medium"} ${groupButton("new_transitem")}`, 
          "new_transitem",
          <Label value={getText("search_transitem")} 
            leftIcon={<Icon iconKey="FileText" />} iconWidth="25px"  />
        )}
        {(group_key === "new_transitem")?<div className={`${"row full"} ${styles.panelGroup}`} >
          {newFilter[0].map(transtype =>{
            if (auditFilter.trans[transtype][0] === "all"){ 
              return (
                itemMenu(transtype, 
                  `${"full medium primary"} ${styles.panelButton}`, 
                  ["editorNew",[{ntype: 'trans', ttype: transtype}]],
                  <Label value={getText("title_"+transtype)} 
                    leftIcon={<Icon iconKey="FileText" />} iconWidth="25px"  />
                )
              ) 
            } else { 
              return null 
            }
          })}
        </div>:null}
      </div>)
    }

    if(newFilter[1].length > 0){
      mnu_items.push(<div key="1" className="row full">
        {groupMenu(
          "new_transpayment_group", 
          `${"full medium"} ${groupButton("new_transpayment")}`, 
          "new_transpayment",
          <Label value={getText("search_transpayment")} 
            leftIcon={<Icon iconKey="Money" />} iconWidth="25px"  />
        )}
        {(group_key === "new_transpayment")?<div className={`${"row full"} ${styles.panelGroup}`} >
          {newFilter[1].map(transtype =>{
            if (auditFilter.trans[transtype][0] === "all"){ 
              return (
                itemMenu(transtype, 
                  `${"full medium primary"} ${styles.panelButton}`, 
                  ["editorNew",[{ntype: 'trans', ttype: transtype}]],
                  <Label value={getText("title_"+transtype)} 
                    leftIcon={<Icon iconKey="Money" />} iconWidth="25px"  />
                )
              ) 
            } else { 
              return null 
            }
          })}
        </div>:null}
      </div>)
    }

    if(newFilter[2].length > 0){
      mnu_items.push(<div key="2" className="row full">
        {groupMenu("new_transmovement_group", 
          `${"full medium"} ${groupButton("new_transmovement")}`, 
          "new_transmovement",
          <Label value={getText("search_transmovement")} 
            leftIcon={<Icon iconKey="Truck" />} iconWidth="25px"  />
        )}
        {(group_key === "new_transmovement")?<div className={`${"row full"} ${styles.panelGroup}`} >
          {newFilter[2].map(transtype => {
            if (auditFilter.trans[transtype][0] === "all"){
              if(transtype === "delivery"){
                return ([
                  itemMenu("shipping", 
                    `${"full medium primary"} ${styles.panelButton}`, 
                    ["editorNew",[{ntype: 'trans', ttype: "shipping"}]],
                    <Label value={getText("title_"+transtype)} 
                      leftIcon={<Icon iconKey={forms[transtype]().options.icon} />} iconWidth="25px"  />
                  ),
                  itemMenu(transtype, 
                    `${"full medium primary"} ${styles.panelButton}`, 
                    ["editorNew",[{ntype: 'trans', ttype: transtype}]],
                    <Label value={getText("title_transfer")} 
                      leftIcon={<Icon iconKey={forms[transtype]().options.icon} />} iconWidth="25px"  />
                  )
                ])
              } else {
                return (
                  itemMenu(transtype,
                    `${"full medium primary"} ${styles.panelButton}`, 
                    ["editorNew",[{ntype: 'trans', ttype: transtype}]],
                    <Label value={getText("title_"+transtype)} 
                      leftIcon={<Icon iconKey={forms[transtype]().options.icon} />} iconWidth="25px"  />
                  )
                )
              } 
            } else { 
              return null 
            }
          })}
        </div>:null}
      </div>)
    }

    if(newFilter[3].length > 0){
      mnu_items.push(<div key="3" className="row full">
        {groupMenu("new_resources_group", 
          `${"full medium"} ${groupButton("new_resources")}`, 
          "new_resources",
          <Label value={getText("title_resources")} 
            leftIcon={<Icon iconKey="Wrench" />} iconWidth="25px"  />
        )}
        {(group_key === "new_resources")?<div className={`${"row full"} ${styles.panelGroup}`} >
          {newFilter[3].map(ntype =>{
            if (auditFilter[ntype][0] === "all"){ 
              return (
                itemMenu(ntype, 
                  `${"full medium primary"} ${styles.panelButton}`, 
                  ["editorNew", [{ntype: ntype, ttype: null}]],
                  <Label value={getText("title_"+ntype)} 
                    leftIcon={<Icon iconKey={forms[ntype]().options.icon} />} iconWidth="25px"  />
                )
              ) 
            } else { 
              return null 
            }
          })}
        </div>:null}
      </div>)
    }

    return mnu_items
  }
  return (
    <div {...props}
      className={`${styles.sidebar} ${((side !== "auto")? side : "")} ${className}`} >
      {(!current.form && (current.form_type !== "transitem_shipping"))?
      <div className="row full section-small container">
        <div className="cell half">
          {itemMenu("state_new", 
            `${"full medium"} ${(edit && current.item)?styles.groupButton:styles.selectButton}`,
            ["editState"],
            <Label value={getText("label_new")} 
              leftIcon={<Icon iconKey="Plus" />} iconWidth="20px"  />
          )}
        </div>
        <div className="cell half">
          {itemMenu("state_edit", 
            `${"full medium"} ${(edit && current.item)?styles.selectButton:styles.groupButton}`,
            ["editState"],
            <Label value={getText("label_edit")} 
              leftIcon={<Icon iconKey="Edit" />} iconWidth="20px"  />,
            (!current.item) ? true : false
          )}
        </div>
      </div>:null}
      {((edit && current.form) || (edit && current.item))
        ? editItems(panel) : newItems()}
    </div>
  )
}

Edit.propTypes = {
  /**
   * SideBar visibility
   */
  side: PropTypes.oneOf(Object.values(SIDE_VISIBILITY)).isRequired,
  edit: PropTypes.bool.isRequired, 
  module: PropTypes.shape({
    current: PropTypes.object, 
    form_dirty: PropTypes.bool,
    dirty: PropTypes.bool,
    panel: PropTypes.object, 
    dataset: PropTypes.object, 
    group_key: PropTypes.string
  }).isRequired, 
  newFilter: PropTypes.array.isRequired,
  auditFilter: PropTypes.object.isRequired,
  forms: PropTypes.object.isRequired,
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

Edit.defaultProps = {
  side: SIDE_VISIBILITY.AUTO,
  edit: true, 
  module: {
    current: {}, 
    form_dirty: false,
    dirty: false,
    panel: {}, 
    dataset: {}, 
    group_key: ""
  }, 
  newFilter: [],
  auditFilter: {}, 
  forms: {}, 
  className: "",
  onEvent: undefined,  
  getText: undefined,
}

export default Edit;