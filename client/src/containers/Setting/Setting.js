import { memo, Fragment } from 'react';
import PropTypes from 'prop-types';

import SideBar from 'components/SideBar/Setting'
import Form from 'components/Setting/Form'
import View from 'components/Setting/View'

export const SettingView = ({
  data, current, login, username, paginationPage,
  onEvent, getText
}) => {
  return (
    <Fragment>
      <SideBar side={current.side} 
        module={data} auditFilter={login.audit_filter} username={username}
        onEvent={onEvent} getText={getText} />
      <div className={`${"page padding-normal"} ${current.theme}`} >
        {(data.current)?
        <Form data={data} onEvent={onEvent} getText={getText} />: 
        <View data={data} paginationPage={paginationPage} onEvent={onEvent} getText={getText} />}
      </div>
    </Fragment>
  )
}

SettingView.propTypes = {
  data: PropTypes.shape({
    caption: PropTypes.string.isRequired, 
    icon: PropTypes.string.isRequired,
    view: PropTypes.object.isRequired,
    actions: PropTypes.object.isRequired,
    current: PropTypes.object,
    audit: PropTypes.string.isRequired, 
    dataset: PropTypes.object.isRequired,
    type: PropTypes.string.isRequired,
  }).isRequired,
  current: PropTypes.object.isRequired,
  login: PropTypes.object.isRequired,
  username: PropTypes.string.isRequired,
  paginationPage: PropTypes.number,
  onEvent: PropTypes.func,
  getText: PropTypes.func
}

SettingView.defaultProps = {
  data: {
    caption: "", 
    icon: "",
    view: {},
    actions: {},
    current: undefined,
    audit: "", 
    dataset: {},
    type: "",
  }, 
  current: {},
  login: {},
  username: "",
  paginationPage: undefined,
  onEvent: undefined,
  getText: undefined
}

export default memo(SettingView, (prevProps, nextProps) => {
  return (
    (prevProps.data === nextProps.data) &&
    (prevProps.current.side === nextProps.current.side)
  )
})
