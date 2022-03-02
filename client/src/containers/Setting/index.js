import React, { useContext, useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import update from 'immutability-helper';

import AppStore from 'containers/App/context'
import { appActions } from 'containers/App/actions'
import { settingActions } from './actions'
import SettingMemo, { SettingView } from './Setting';

const Setting = (props) => {
  const { data, setData } = useContext(AppStore);
  const app = appActions(data, setData)
  const setting = settingActions(data, setData)

  const [state] = useState(update(props, {$merge: {
    username: data.login.username,
    login: data.login.data,
  }}))

  state.data = update(state.data, {$merge: { ...data[state.key] }})
  state.current = update(state.current, {$merge: { ...data.current }})

  useEffect(() => {
    if(state.current && state.current.content){
      const content = state.current.content
      setData("current", { content: null }, () => {
        setting.checkSetting(content, content.nextKey || "LOAD_SETTING")
      })
    }
  }, [setData, setting, state]);

  state.getText = (key, defValue) => {
    return app.getText(key, defValue)
  }

  state.onEvent = (fname, params) => {
    if(setting[fname]){
      return setting[fname](...params)  
    }
    if(app[fname]){
      return app[fname](...params)  
    }
    state[fname](...params)
  }

  state.changeData = (fieldname, value) => {
    setData("setting", { [fieldname]: value })
  }

  state.companyForm = () => {
    setData("current", { module: "edit", content: { ntype: "customer", ttype: null, id: 1 } })
  }

  if(state.data.type){
    return <SettingMemo {...state} />
  }
  return <div />

}

Setting.propTypes = {
  key: PropTypes.string.isRequired,
  ...SettingView.propTypes,
  changeData: PropTypes.func,
  companyForm: PropTypes.func
}

Setting.defaultProps = {
  key: "setting",
  ...SettingView.defaultProps,
  changeData: undefined,
  companyForm: undefined
}

export default Setting;
