import { memo, Fragment } from 'react';
import PropTypes from 'prop-types';

import SideBar from 'components/SideBar/Edit'
import Editor from 'components/Editor/Editor'

export const EditorView = ({
  data, current, login, forms,
  onEvent, getText
}) => {
  return (
    <Fragment>
      <SideBar side={current.side}
        edit={current.edit} module={data} forms={forms}
        newFilter={login.edit_new} auditFilter={login.audit_filter}
        onEvent={onEvent} getText={getText} />
      {(data.current.item)?<div className={`${"page padding-normal"} ${current.theme}`} >
        <Editor caption={data.caption}
          current={data.current} template={data.template}
          dataset={data.dataset} audit={data.audit} lastUpdate={data.lastUpdate}
          onEvent={onEvent} getText={getText} />
      </div>:null}
    </Fragment>
  )
}

EditorView.propTypes = {
  data: PropTypes.shape({
    caption: Editor.propTypes.caption, 
    current: Editor.propTypes.current, 
    template: Editor.propTypes.template, 
    dataset: Editor.propTypes.dataset, 
    audit: Editor.propTypes.audit,
  }).isRequired,
  current: PropTypes.object.isRequired,
  login: PropTypes.object.isRequired,
  forms: PropTypes.object.isRequired,
  onEvent: PropTypes.func,
  getText: PropTypes.func
}

EditorView.defaultProps = {
  data: {
    caption: Editor.defaultProps.caption, 
    current: Editor.defaultProps.current, 
    template: Editor.defaultProps.template, 
    dataset: Editor.defaultProps.dataset, 
    audit: Editor.defaultProps.audit,
  }, 
  current: {},
  login: {},
  forms: {},
  onEvent: undefined,
  getText: undefined
}

export default memo(EditorView, (prevProps, nextProps) => {
  return (
    (prevProps.data === nextProps.data) &&
    (prevProps.current.side === nextProps.current.side) &&
    (prevProps.current.edit === nextProps.current.edit)
  )
})

