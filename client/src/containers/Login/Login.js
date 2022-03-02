import { memo } from 'react';
import PropTypes from 'prop-types';

import Login from 'components/Login'

export const LoginPage = ({
  session, data, current, locales,
  changeData, getText, onLogin, setTheme, setLocale
}) => {
  return <Login {...data} 
    theme={current.theme} lang={current.lang} 
    version={session.version} locales={locales} configServer={session.configServer}
    changeData={changeData} getText={getText} onLogin={onLogin} setTheme={setTheme} setLocale={setLocale} />
}

LoginPage.propTypes = {
  data: PropTypes.shape({
    username: Login.propTypes.username,
    password: Login.propTypes.password,
    database: Login.propTypes.database,
    server: Login.propTypes.server,
  }).isRequired,
  session: PropTypes.object.isRequired,
  locales: PropTypes.object.isRequired,
  current: PropTypes.object.isRequired,
  changeData: PropTypes.func, 
  getText: PropTypes.func, 
  onLogin: PropTypes.func, 
  setTheme: PropTypes.func, 
  setLocale: PropTypes.func,
}

LoginPage.defaultProps = {
  data: {
    username: Login.defaultProps.username,
    password: Login.defaultProps.password,
    database: Login.defaultProps.database,
    server: Login.defaultProps.server,
  },
  session: {},
  locales: {},
  current: {},
  changeData: undefined, 
  getText: undefined, 
  onLogin: undefined, 
  setTheme: undefined, 
  setLocale: undefined,
}

export default memo(LoginPage, (prevProps, nextProps) => {
  return (
    (prevProps.data === nextProps.data) &&
    (prevProps.current.theme === nextProps.current.theme) &&
    (prevProps.current.lang === nextProps.current.lang) &&
    (prevProps.locales === nextProps.locales)
  )
})

