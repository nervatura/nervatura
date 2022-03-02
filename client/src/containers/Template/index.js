import React, { useContext, useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import update from 'immutability-helper';

import AppStore from 'containers/App/context'
import { appActions } from 'containers/App/actions'
import { templateActions } from 'containers/Template/actions'

import TemplateMemo, { TemplateView } from './Template';

const Template = (props) => {
  const { data, setData } = useContext(AppStore)
  const app = appActions(data, setData)
  const template = templateActions(data, setData)

  const [state] = useState(update({}, {$merge: {
    ...props
  }}))

  state.data = update(state.data, {$merge: { ...data[state.key] }})
  state.current = update(state.current, {$merge: { ...data.current }})

  useEffect(() => {
    if(state.current && state.current.content){
      const content = state.current.content
      setData("current", { content: null }, () => {
        template.setTemplate(content)
      })
    }
  }, [setData, template, state]);

  state.getText = (key, defValue) => {
    return app.getText(key, defValue)
  }

  state.onEvent = (fname, params) => {
    if(app[fname]){
      return app[fname](...params)  
    }
    template[fname](...params)  
  }

  if(state.data.template){
    return <TemplateMemo {...state} />
  }
  return <div />

}

Template.propTypes = {
  key: PropTypes.string.isRequired,
  ...TemplateView.propTypes,
}

Template.defaultProps = {
  key: "template",
  ...TemplateView.defaultProps,
}

export default Template;