import PropTypes from 'prop-types';

import styles from './Login.module.css';

import Label from 'components/Form/Label'
import Input from 'components/Form/Input'
import Select from 'components/Form/Select'
import Button from 'components/Form/Button'
import Icon from 'components/Form/Icon'

export const Login = ({ 
  theme, version, locales, configServer, lang, 
  username, password, database, server,
  onLogin, setTheme, setLocale, getText, changeData,
  ...props 
}) => {
  return (
    <div className={styles.modal}>
      <div className={styles.middle}>
        <div {...props} className={`${styles.dialog} ${theme}`} >
          <div className={`${"row primary"} ${styles.title}`} >
            <div className="cell" >
              <Label value={getText("title_login")} />
            </div>
            <div className={`${"cell"} ${styles.version}`} >
              <Label value={"v"+version} />
            </div>
          </div>
          <div className="row full section-small" >
            <div className="row full section-small" >
              <div className="row full" >
                <div className="padding-normal s12 m4 l4" >
                  <Label value={getText("login_username")} className="bold" />
                </div>
                <div className="container s12 m8 l8" >
                  <Input id="username" type="text" 
                    className={`${"full"} ${styles.inputMargin}`} value={username}
                    onChange={(value)=>changeData("username", value)} />
                </div>
              </div>
              <div className="row full" >
                <div className="padding-normal s12 m4 l4" >
                  <Label value={getText("login_password")} className="bold" />
                </div>
                <div className="container s12 m8 l8" >
                  <Input id="password" type="password" 
                    className={`${"full"} ${styles.inputMargin}`} value={password}
                    onChange={(value)=>changeData("password", value)} 
                    onEnter={onLogin} />
                </div>
              </div>
            </div>
            <div className="row full section-small" >
              <div className="row full" >
                <div className="padding-normal s12 m4 l4" >
                  <Label value={getText("login_database")} className="bold" />
                </div>
                <div className="container s12 m8 l8" >
                  <Input id="database" type="text" 
                    className={`${"full"} ${styles.inputMargin}`} value={database}
                    onChange={(value)=>changeData("database", value)} 
                    onEnter={onLogin} />
                </div>
              </div>
              {(!configServer)?<div className="row full" >
                <div className="padding-normal full" >
                  <Label value={getText("login_server")} className="bold" />
                </div>
                <div className="container full" >
                  <Input id="server" type="text" 
                    className={`${"full"} ${styles.inputMargin}`} value={server}
                    onChange={(value)=>changeData("server", value)} />
                </div>
              </div>:null}
            </div>
          </div>
          <div className={`${"row full section-small secondary-title"}`} >
            <div className="container section-small s6 m6 l6" >
              <Button id="theme" className={`${"border-button"} ${styles.borderButton}`}
                value={(theme === "dark")
                  ?<Icon iconKey="Sun" width={18} height={18} />:<Icon iconKey="Moon" width={18} height={18} />} 
                onClick={setTheme} />
              <Select id="lang" value={lang}
                options={Object.keys(locales).map(key => {
                  return {
                    value: key, text: locales[key][key] || key 
                  }
                })} 
                onChange={(value)=>setLocale(value)} />
            </div>
            <div className="container section-small s6 m6 l6" >
              <Button id="login" autoFocus
                disabled={((!username || (String(username).length===0) 
                  || !database || (String(database).length===0))?"disabled":"")} 
                className="primary full" label={getText("login_login")}
                onClick={onLogin} />
            </div>
          </div>  
        </div>
      </div>
    </div>
  )
}

Login.propTypes = {
  /**
   * UI theme name
   */ 
  theme: PropTypes.oneOf([ "light", "dark" ]).isRequired,
  /**
   * App version string
   */
  version: PropTypes.string.isRequired,
  /**
   * Language options
   */
  locales: PropTypes.object.isRequired,
  /**
   * Current language
   */
  lang: PropTypes.string.isRequired,
  /**
   * Show / hide server url input value
   */
  configServer: PropTypes.bool.isRequired,
  /**
   * Login username value
   */
  username: PropTypes.string.isRequired,
  /**
   * Password username value
   */
  password: PropTypes.string.isRequired,
  /**
   * Login database value
   */
  database: PropTypes.string.isRequired,
  /**
   * Server URL
   */
  server: PropTypes.string.isRequired,
  /**
   * Login handler
   */
  onLogin: PropTypes.func,
  /**
   * Set light/dark theme
   */
  setTheme: PropTypes.func,
  /**
   * Change language
   */
  setLocale: PropTypes.func,
  /**
   * Input value change handler
   */
  changeData: PropTypes.func,
  /**
   * Localization
   */
  getText: PropTypes.func,
}

Login.defaultProps = {
  theme: "light",
  version: "",
  locales: {},
  lang: "en",
  configServer: false,
  username: "",
  password: "",
  database: "",
  server: "",
  changeData: undefined,
  getText: undefined,
  onLogin: undefined,
  setTheme: undefined,
  setLocale: undefined
}

export default Login;