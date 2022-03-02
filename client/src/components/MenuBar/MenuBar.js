import PropTypes from 'prop-types';

import styles from './MenuBar.module.css';

import Icon from 'components/Form/Icon'
import Label from 'components/Form/Label'

export const SIDE_VISIBILITY = {
  AUTO: "auto",
  SHOW: "show",
  HIDE: "hide"
}

export const APP_MODULE = {
  LOGIN: "login",
  SEARCH: "search",
  EDIT: "edit",
  SETTING: "setting",
  HELP: "help",
  BOOKMARK: "bookmark",
  TEMPLATE: "template"
}

export const MenuBar = ({ 
  side, scrollTop, module,
  getText, loadModule, sideBar, setScroll,
  ...props 
}) => {
  const selected = (key) => {
    if(key === module){
      return styles.selected
    }
    return ""
  }
  return (
    <div {...props} className={`${(scrollTop)?styles.shadow:""} ${styles.menubar}`} >
      <div className="cell">
        <div id="mnu_sidebar"
          className={`${styles.menuitem} ${styles.sidebar}`} onClick={() => sideBar() }>
          {(side === "show")?
            <Label className={`${styles.selected} ${styles.exit}`} 
              leftIcon={<Icon iconKey="Close" className={`${styles.selected} ${styles.exit}`} />} 
              value={getText("menu_hide")} />:
            <Label leftIcon={<Icon iconKey="Bars" width={24} height={24} />} 
              value={getText("menu_side")} />}
        </div>
        <div id="mnu_search_large" 
          className={`${"hide-small hide-medium"} ${selected(APP_MODULE.SEARCH)} ${styles.menuitem}`} 
          onClick={() => loadModule(APP_MODULE.SEARCH) } >
          <Label leftIcon={<Icon iconKey="Search" />} value={getText("menu_search")} />
        </div>
        <div id="mnu_edit_large"
          className={`${"hide-small hide-medium"} ${selected(APP_MODULE.EDIT)} ${styles.menuitem}`} 
          onClick={() => loadModule(APP_MODULE.EDIT) } >
          <Label leftIcon={<Icon iconKey="Edit" />} value={getText("menu_edit")} />
        </div>
        <div id="mnu_setting_large"
          className={`${"hide-small hide-medium"} ${selected(APP_MODULE.SETTING)} ${styles.menuitem}`} 
          onClick={() => loadModule(APP_MODULE.SETTING) } >
          <Label leftIcon={<Icon iconKey="Cog" />} value={getText("menu_setting")} />
        </div>
        <div id="mnu_bookmark_large"
          className={`${"hide-small hide-medium"} ${selected(APP_MODULE.BOOKMARK)} ${styles.menuitem}`} 
          onClick={() => loadModule(APP_MODULE.BOOKMARK) } >
          <Label leftIcon={<Icon iconKey="Star" />} value={getText("menu_bookmark")} />
        </div>
        <div id="mnu_help_large"
          className={`${"hide-small hide-medium"} ${selected(APP_MODULE.HELP)} ${styles.menuitem}`} 
          onClick={() => loadModule(APP_MODULE.HELP) } >
          <Label leftIcon={<Icon iconKey="QuestionCircle" />} value={getText("menu_help")} />
        </div>
        <div id="mnu_logout_large"
          className={`${"hide-small hide-medium"} ${styles.menuitem} ${styles.exit}`} 
          onClick={() => loadModule(APP_MODULE.LOGIN) } >
          <Label leftIcon={<Icon iconKey="Exit" />} value={getText("menu_logout")} />
        </div>

        {(scrollTop)?<div id="mnu_scroll" className={`${styles.menuitem}`} 
          onClick={() => setScroll() } >
          <Icon iconKey="HandUp" />
        </div>:null}
      </div>
      <div id="mnu_logout_medium"
        className={`${"right hide-large"} ${styles.menuitem} ${styles.exit}`} 
        onClick={() => loadModule(APP_MODULE.LOGIN) } >
        <Icon iconKey="Exit" width={24} height={24}/>
      </div>
      <div className={`${"cell"} ${"right"} ${styles.container}`}>
        <div id="mnu_help_medium"
          className={`${"right hide-large"} ${selected(APP_MODULE.HELP)} ${styles.menuitem}`} 
          onClick={() => loadModule(APP_MODULE.HELP) } >
          <Label className="hide-small" leftIcon={<Icon iconKey="QuestionCircle" />} value={getText("menu_help")} />
        </div>
        <div id="mnu_bookmark_medium"
          className={`${"right hide-large"} ${selected(APP_MODULE.BOOKMARK)} ${styles.menuitem}`} 
          onClick={() => loadModule(APP_MODULE.BOOKMARK) } >
          <Label className="hide-small" leftIcon={<Icon iconKey="Star" />} value={getText("menu_bookmark")} />
        </div>
        <div id="mnu_setting_medium"
          className={`${"right hide-large"} ${selected(APP_MODULE.SETTING)} ${styles.menuitem}`} 
          onClick={() => loadModule(APP_MODULE.SETTING) } >
          <Label className="hide-small" leftIcon={<Icon iconKey="Cog" />} value={getText("menu_setting")} />
        </div>
        <div id="mnu_edit_medium"
          className={`${"right hide-large"} ${selected(APP_MODULE.EDIT)} ${styles.menuitem}`} 
          onClick={() => loadModule(APP_MODULE.EDIT) } >
          <Label className="hide-small" leftIcon={<Icon iconKey="Edit" />} value={getText("menu_edit")} />
        </div>
        <div id="mnu_search_medium"
          className={`${"right hide-large"} ${selected(APP_MODULE.SEARCH)} ${styles.menuitem}`} 
          onClick={() => loadModule(APP_MODULE.SEARCH) } >
          <Label className="hide-small" leftIcon={<Icon iconKey="Search" />} value={getText("menu_search")} />
        </div>
      </div>
    </div>
  )
}

MenuBar.propTypes = {
  /**
   * SideBar visibility
   */
  side: PropTypes.oneOf(Object.values(SIDE_VISIBILITY)).isRequired,
  /**
   * Show scroll to top button
   */
  scrollTop: PropTypes.bool.isRequired,
  /**
   * Current module key
   */
  module: PropTypes.oneOf(Object.values(APP_MODULE)).isRequired,
  /**
   * Module loading handle 
   */ 
  loadModule: PropTypes.func,
  /**
   * Show/hide sidebar
   */ 
  sideBar: PropTypes.func,
  /**
   * Scroll to top handle
   */ 
  setScroll: PropTypes.func,
  /**
   * Localization
   */
  getText: PropTypes.func,
}

MenuBar.defaultProps = {
  side: SIDE_VISIBILITY.AUTO,
  scrollTop: false,
  module: APP_MODULE.SEARCH, 
  loadModule: undefined, 
  sideBar: undefined, 
  setScroll: undefined,
  getText: undefined,
}


export default MenuBar;