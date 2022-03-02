import React, { useContext, useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import update from 'immutability-helper';

import AppStore from 'containers/App/context'
import { appActions } from 'containers/App/actions'
import { editorActions } from 'containers/Editor/actions'
import { Forms } from 'containers/Controller/Forms'
import InputBox from 'components/Modal/InputBox'

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
      return editor[fname](...params, state.data)  
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

  state.editState = () => {
    setData("current", { edit: !state.current.edit })
  }

  state.onSelector = (selectorType, selectorFilter, setSelector) => {
    app.onSelector(selectorType, selectorFilter, setSelector)
  }

  state.editorBack = () =>{
    if(data.edit.current.form){
      editor.checkEditor({
        ntype: data.edit.current.type, 
        ttype: data.edit.current.transtype, 
        id: data.edit.current.item.id,
        form: data.edit.current.form_type}, 'LOAD_EDITOR')
    } else {
      if(data.edit.current.form_type === "transitem_shipping"){
        editor.checkEditor({
          ntype: data.edit.current.type, 
          ttype: data.edit.current.transtype, 
          id: data.edit.current.item.id,
          form: data.edit.current.form_type}, 'LOAD_EDITOR')
      } else {
        let reftype = state.login.groups.filter((item)=> {
          return (item.id === data.edit.current.item.nervatype)
        })[0].groupvalue
        editor.checkEditor({ntype: reftype, 
          ttype: null, id: data.edit.current.item.ref_id,
          form: data.edit.current.type}, 'LOAD_EDITOR')
      }
    }
  }

  state.saveEditor = async () => {
    let edit = null
    if(data.edit.current.form){
      edit = await editor.saveEditorForm()
    } else {
      edit = await editor.saveEditor()
    }
    if(edit){
      editor.loadEditor({
        ntype: edit.current.type, 
        ttype: edit.current.transtype, 
        id: edit.current.item.id,
        form: edit.current.view
      })
    }
  }

  state.editorDelete = () => {
    if(data.edit.current.form){
      editor.deleteEditorItem({
        fkey: data.edit.current.form_type, 
        table: data.edit.current.form_datatype, 
        id: data.edit.current.form.id
      })
    } else {
      editor.deleteEditor()
    }
  }

  state.editorNew = (params) =>{
    if(params.ttype === "shipping"){
      app.onSelector("transitem_delivery", "", (row, filter)=>{
        const params = row.id.split("/")
        editor.checkEditor({ 
          ntype: params[0], ttype: params[1], id: parseInt(params[2],10), 
          shipping: true
        }, 'LOAD_EDITOR')
      })
    /*
    } else if(data.edit.current.form){
      editor.checkEditor({
        fkey: params.fkey || data.edit.current.form_type, 
        id: null}, 'SET_EDITOR_ITEM')
    */
    } else {
      editor.checkEditor({
        ntype: params.ntype || data.edit.current.type, 
        ttype: params.ttype || data.edit.current.transtype, 
        id: null}, 'LOAD_EDITOR')
    }
  }

  state.transCopy = (ctype) => {
    if (ctype === "create") {
      editor.checkEditor({}, "CREATE_TRANS_OPTIONS");
    } else {
      setData("current", { modalForm: 
        <InputBox 
          title={app.getText("msg_warning")}
          message={app.getText("msg_copy_text")}
          infoText={app.getText("msg_delete_info")}
          defaultOK={true}
          labelOK={app.getText("msg_ok")}
          labelCancel={app.getText("msg_cancel")}
          onCancel={() => {
            setData("current", { modalForm: null })
          }}
          onOK={(value) => {
            setData("current", { modalForm: null }, () => {
              editor.checkEditor({ cmdtype: "copy", transcast: ctype }, "CREATE_TRANS");
            })
          }}
        />,
        side: "hide"
      })
    }
  };

  state.setLink = (type, field) =>{
    setData("current", { side: "hide" }, ()=>{
      let link_id = (data.edit.current.transtype === "cash") ? 
      data.edit.current.extend.id : data.edit.current.form.id;
      editor.checkEditor(
        { fkey: type, id: null, link_field: field, link_id: link_id }, 'SET_EDITOR_ITEM')
    })    
  }

  state.setPassword = (username) =>{
    if(!username && data.edit.current){
      username = data.edit.dataset[data.edit.current.type][0].username
    }
    setData("current", { module: "setting", content: { username: username, nextKey: "PASSWORD_FORM" }, side: "hide" })
  }

  state.shippingAddAll = () => {
    let edit = update({}, {$set: data.edit})
    edit.dataset.shipping_items_.forEach(sitem => {
      if (sitem.diff !== 0 && sitem.edited !== true) {
        edit = update(edit, {dataset: { shiptemp: {$push: [{ 
          "id": sitem.item_id+"-"+sitem.product_id,
          "item_id": sitem.item_id, 
          "product_id": sitem.product_id,  
          "product": sitem.product, 
          "partnumber": sitem.partnumber,
          "partname": sitem.partname, 
          "unit": sitem.unit, 
          "batch_no":"", 
          "qty":sitem.diff, 
          "diff":0,
          "oqty": sitem.qty, 
          "tqty": sitem.tqty
        }]}}})
      }
    });
    editor.setEditor({shipping: true, form:"shiptemp_items"}, edit.template, edit)
  }

  return <EditorMemo {...state} />

}

Editor.propTypes = {
  key: PropTypes.string.isRequired,
  ...EditorView.propTypes,
  changeCurrentData: PropTypes.func,
  onPaginationSelect: PropTypes.func,
  onSelector: PropTypes.func,
}

Editor.defaultProps = {
  key: "edit",
  ...EditorView.defaultProps,
  changeCurrentData: undefined,
  onPaginationSelect: undefined,
  onSelector: undefined,
}

export default Editor;