import PropTypes from 'prop-types';

import styles from './Setting.module.css';

import Icon from 'components/Form/Icon'
import Label from 'components/Form/Label'
import Button from 'components/Form/Button'

export const SIDE_VISIBILITY = {
  AUTO: "auto",
  SHOW: "show",
  HIDE: "hide"
}

export const Setting = ({ 
  side, auditFilter, module, username, className,
  getText, onEvent,
  ...props 
}) => {
  const { group_key, current, panel, dirty, type } = module
  const itemMenu = (keyValue, classValue, eventValue, labelValue) => {
    return <Button id={keyValue} key={keyValue}
      className={classValue}
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
  const menuItems = (options)=>{
    let panels = []

    panels.push(
      itemMenu("cmd_back",
        `${"medium"} ${styles.itemButton} ${styles.selected}`, 
        ["settingBack",[]],
        <Label value={getText("label_back")} 
          leftIcon={<Icon iconKey="Reply" />} iconWidth="20px"  />
      )
    )
    panels.push(<div key="back_sep" className={styles.separator} />)

    if (options.save !== false) {
      panels.push(
        itemMenu("cmd_save",
          `${"full medium"} ${styles.itemButton} ${(dirty)?styles.selected:""}`,
          ["settingSave",[]],
          <Label value={getText("label_save")} 
            leftIcon={<Icon iconKey="Check" />} iconWidth="20px"  />
        )
      )
    }
    if ((options.delete !== false) && current.form && (current.form.id !== null)) {
      panels.push(
        itemMenu("cmd_delete",
          `${"full medium"} ${styles.itemButton}`, 
          ["deleteSetting", [current.form]],
          <Label value={getText("label_delete")} 
            leftIcon={<Icon iconKey="Times" />} iconWidth="20px"  />
        )
      )
    }
    if ((options.new !== false) && current.form && (current.form.id !== null)) {
      panels.push(
        itemMenu("cmd_new",
          `${"full medium"} ${styles.itemButton}`, 
          ["checkSetting", [{ type: type, id: null }, 'LOAD_SETTING']],
        <Label value={getText("label_new")} 
          leftIcon={<Icon iconKey="Plus" />} iconWidth="20px"  />
        )
      )
    }
    if (typeof options.help !== "undefined") {
      panels.push(<div key="help_sep" className={styles.separator} />)
      panels.push(
        itemMenu("cmd_help",
          `${"full medium"} ${styles.itemButton}`, 
          ["showHelp", [options.help]],
          <Label value={getText("label_help")} 
            leftIcon={<Icon iconKey="QuestionCircle" />} iconWidth="20px"  />
        )
      )
    }

    return panels
  }
  if(current && panel){
    return(
      <div {...props}
        className={`${styles.sidebar} ${((side !== "auto")? side : "")} ${className}`} >
        {menuItems(panel)}
      </div>
    )
  } else {
    return (
      <div {...props}
        className={`${styles.sidebar} ${((side !== "auto")? side : "")} ${className}`} >
        {(auditFilter.setting[0]!=="disabled")?
          <div className="row full">
            {groupMenu("group_admin", 
              `${"full medium"} ${(group_key === "group_admin")?styles.selectButton:styles.groupButton}`, 
              "group_admin",
              <Label value={getText("title_admin")} 
                leftIcon={<Icon iconKey="ExclamationTriangle" />} iconWidth="20px"  />
            )}
            {(group_key === "group_admin")?<div className={`${"row full"} ${styles.panelGroup}`} >
              {itemMenu("cmd_dbsettings", 
                `${"full medium primary"} ${styles.panelButton}`, 
                ["loadSetting", [{ type: 'setting' }]],
                <Label value={getText("title_dbsettings")} 
                  leftIcon={<Icon iconKey="Cog" />} iconWidth="20px"  />
              )}
              {itemMenu("cmd_numberdef", 
                `${"full medium primary"} ${styles.panelButton}`, 
                ["loadSetting", [{ type: 'numberdef' }]],
                <Label value={getText("title_numberdef")} 
                  leftIcon={<Icon iconKey="ListOl" />} iconWidth="20px"  />
              )}
              {(auditFilter.audit[0]!=="disabled")?
                itemMenu("cmd_usergroup", 
                  `${"full medium primary"} ${styles.panelButton}`, 
                  ["loadSetting", [{ type: 'usergroup' }]],
                  <Label value={getText("title_usergroup")} 
                    leftIcon={<Icon iconKey="Key" />} iconWidth="20px"  />
                ):null}
              {itemMenu("cmd_ui_menu", 
                `${"full medium primary"} ${styles.panelButton}`, 
                ["loadSetting", [{ type: 'ui_menu' }]],
                <Label value={getText("title_menucmd")} 
                  leftIcon={<Icon iconKey="Share" />} iconWidth="20px"  />
              )}
              {itemMenu("cmd_log", 
                `${"full medium primary"} ${styles.panelButton}`, 
                ["loadSetting", [{ type: 'log' }]],
                <Label value={getText("title_log")} 
                 leftIcon={<Icon iconKey="InfoCircle" />} iconWidth="20px"  />
              )}
            </div>:null}
          </div>:null}
        {(auditFilter.setting[0]!=="disabled")?
          <div className="row full">
            {groupMenu("group_database", 
              `${"full medium"} ${(group_key === "group_database")?styles.selectButton:styles.groupButton}`, 
              "group_database",
              <Label value={getText("title_database")} 
                leftIcon={<Icon iconKey="Database" />} iconWidth="20px"  />
            )}
            {(group_key === "group_database")?<div className={`${"row full"} ${styles.panelGroup}`} >
              {itemMenu("cmd_deffield", 
                `${"full medium primary"} ${styles.panelButton}`, 
                ["loadSetting", [{ type: 'deffield' }]],
                <Label value={getText("title_deffield")} 
                  leftIcon={<Icon iconKey="Tag" />} iconWidth="20px"  />
              )}
              {itemMenu("cmd_groups", 
                `${"full medium primary"} ${styles.panelButton}`, 
                ["loadSetting", [{ type: 'groups' }]],
                <Label value={getText("title_groups")} 
                  leftIcon={<Icon iconKey="Th" />} iconWidth="20px"  />
              )}
              {itemMenu("cmd_place", 
                `${"full medium primary"} ${styles.panelButton}`, 
                ["loadSetting", [{ type: 'place' }]],
                <Label value={getText("title_place")} 
                  leftIcon={<Icon iconKey="Map" />} iconWidth="20px"  />
              )}
              {itemMenu("cmd_currency", 
                `${"full medium primary"} ${styles.panelButton}`, 
                ["loadSetting", [{ type: 'currency' }]],
                <Label value={getText("title_currency")} 
                  leftIcon={<Icon iconKey="Dollar" />} iconWidth="20px"  />
              )}
              {itemMenu("cmd_tax", 
                `${"full medium primary"} ${styles.panelButton}`, 
                ["loadSetting", [{ type: 'tax' }]],
                <Label value={getText("title_tax")} 
                  leftIcon={<Icon iconKey="Ticket" />} iconWidth="20px"  />
              )}
              {itemMenu("cmd_company", 
                `${"full medium primary"} ${styles.panelButton}`, 
                ["companyForm", []],
                <Label value={getText("title_company")} 
                  leftIcon={<Icon iconKey="Home" />} iconWidth="20px"  />
              )}
              {itemMenu("cmd_template", 
                `${"full medium primary"} ${styles.panelButton}`, 
                ["loadSetting", [{ type: 'template' }]],
                <Label value={getText("title_report_editor")} 
                  leftIcon={<Icon iconKey="TextHeight" />} iconWidth="20px"  />
              )}
            </div>:null}
          </div>:null}
        <div className="row full">
          {groupMenu("group_user",
            `${"full medium"} ${(group_key === "group_user")?styles.selectButton:styles.groupButton}`, 
            "group_user",
            <Label value={getText("title_user")} 
              leftIcon={<Icon iconKey="Desktop" />} iconWidth="20px"  />
          )}
          {(group_key === "group_user")?<div className={`${"row full"} ${styles.panelGroup}`} >
            {itemMenu("cmd_program", 
              `${"full medium primary"} ${styles.panelButton}`, 
              ["setProgramForm",[]],
              <Label value={getText("title_program")} 
                leftIcon={<Icon iconKey="Keyboard" />} iconWidth="20px"  />
            )}
            {itemMenu("cmd_password", 
              `${"full medium primary"} ${styles.panelButton}`, 
              ["setPassword", [username]],
              <Label value={getText("title_password")} 
                leftIcon={<Icon iconKey="Lock" />} iconWidth="20px"  />
            )}
          </div>:null}
        </div>
      </div>
    )
  }
}

Setting.propTypes = {
  /**
   * SideBar visibility
   */
  side: PropTypes.oneOf(Object.values(SIDE_VISIBILITY)).isRequired,
  auditFilter: PropTypes.object.isRequired,
  username: PropTypes.string,
  module: PropTypes.shape({
    current: PropTypes.object, 
    dirty: PropTypes.bool,
    panel: PropTypes.object,  
    group_key: PropTypes.string
  }).isRequired,
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

Setting.defaultProps = {
  side: SIDE_VISIBILITY.AUTO,
  module: {
    current: {}, 
    dirty: false,
    panel: {},  
    group_key: ""
  }, 
  auditFilter: {},
  username: undefined,
  className: "",  
  onEvent: undefined,
  getText: undefined,
}

export default Setting;