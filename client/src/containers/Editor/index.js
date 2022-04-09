import React, { useContext, useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import update from 'immutability-helper';

import AppStore from 'containers/App/context'
import { appActions } from 'containers/App/actions'
import { editorActions } from 'containers/Editor/actions'
import { Forms } from 'containers/Controller/Forms'

import EditorMemo, { EditorView } from './Editor';

const Editor = (props) => {
  const { data, setData } = useContext(AppStore)
  const app = appActions(data, setData)
  const editor = editorActions(data, setData)

  const [state] = useState(update(props, {$merge: {
    login: data.login.data,
    forms: Forms({ getText: app.getText }),
  }}))

  state.data = update(state.data, {$merge: { ...data[state.key] }})
  state.current = update(state.current, {$merge: { ...data.current }})

  useEffect(() => {
    if(state.current && state.current.content){
      const content = state.current.content
      setData("current", { content: null }, () => {
        editor.checkEditor(content, content.nextKey || "LOAD_EDITOR")
      })
    }
  }, [setData, editor, state]);

  state.getText = (key, defValue) => {
    return app.getText(key, defValue)
  }

  state.onEvent = (fname, params) => {
    params = params || []
    if(state[fname]){
      return state[fname](...params)
    }
    if(editor[fname]){
      return editor[fname](...params)  
    }
    if(app[fname]){
      return app[fname](...params)  
    }
  }

  state.changeData = (fieldname, value) => {
    setData(state.key, { [fieldname]: value })
  }

  state.changeCurrentData = (key, value) => {
    const current = update(state.data.current, {$merge: {
      [key]: value
    }})
    setData(state.key, { current: current })
  }

  return <EditorMemo {...state} />

}

Editor.propTypes = {
  key: PropTypes.string.isRequired,
  ...EditorView.propTypes,
  changeData: PropTypes.func,
  changeCurrentData: PropTypes.func,
}

Editor.defaultProps = {
  key: "edit",
  ...EditorView.defaultProps,
  changeData: undefined,
  changeCurrentData: undefined,
}

export default Editor;