import { useCallback, useEffect } from 'react';
import PropTypes from 'prop-types';

import { init } from 'pell';
import 'pell/dist/pell.css'

import styles from './Note.module.css';

import Icon from 'components/Form/Icon'
import Button from 'components/Form/Button'
import Select from 'components/Form/Select'

let noteEditor = null

export const Note = ({ 
  value, patternId, patterns, readOnly, lastUpdate, className,
  getText, onEvent,
  ...props 
}) => {
  
  const editorRef = useCallback(element => {
    if(element && !noteEditor){
      const editor = init({
        element: element,
        onChange: (html) => {
          /* istanbul ignore next */
          onEvent("editItem", [{ name: "fnote", value: html }])
        },
        actions: ['bold', 'italic'],
        classes: {
          actionbar: `border-bottom ${styles.actionbar}`,
          button: `border-button ${styles.barButton}`,
          content: 'pell-content',
          selected: styles.activeStyle
        }
      })
      editor.content.innerHTML = value
      noteEditor = editor
    }
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  useEffect(() => {
    if(noteEditor && (noteEditor.content.innerHTML !== value)){
      noteEditor.content.innerHTML = value
    }
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [lastUpdate])

  useEffect(() => {
    return function cleanup() {
      noteEditor = null
    }
  })
  
  return (
    <div {...props} className={`${styles.formPanel} ${"border"} ${className}`} >
      {(!readOnly)?<div>
        <div className="row full" >
          <div className={`${"cell padding-small"}`} >
            <div className="cell padding-tiny">
              <Button id="btn_pattern_default"
                className={`${"border-button"} ${styles.barButton}`}
                title={getText("pattern_default")}
                onClick={ ()=>onEvent("setPattern",[{ key: "default" }]) }
                value={<Icon iconKey="Home" />}
              />
              <Button id="btn_pattern_load"
                className={`${"border-button"} ${styles.barButton}`}
                title={getText("pattern_load")}
                onClick={ ()=>onEvent("setPattern",[{ key: "load" }]) }
                value={<Icon iconKey="Download" />}
              />
              <Button id="btn_pattern_save"
                className={`${"border-button"} ${styles.barButton}`}
                title={getText("pattern_save")}
                onClick={ ()=>onEvent("setPattern",[{ key: "save"}]) }
                value={<Icon iconKey="Upload" />}
              />
            </div>
            <div className="cell padding-tiny">
              <Button id="btn_pattern_new"
                className={`${"border-button"} ${styles.barButton}`}
                title={getText("pattern_new")}
                onClick={ ()=>onEvent("setPattern",[{ key: "new"}]) }
                value={<Icon iconKey="Plus" />}
              />
              <Button id="btn_pattern_delete"
                className={`${"border-button"} ${styles.barButton}`}
                title={getText("pattern_delete")}
                onClick={ ()=>onEvent("setPattern",[{ key: "delete"}]) }
                value={<Icon iconKey="Times" />}
              />
            </div>
            <div className="cell padding-tiny mobile" >
              <Select id="sel_pattern"
                value={(patternId) ? String(patternId) : ""} placeholder=""
                onChange={ (value)=>onEvent("changeCurrentData",["template", value]) }
                options={patterns.map( pattern => {
                  return { value: String(pattern.id), 
                    text: pattern.description+((pattern.defpattern === 1)?"*":"") 
                }})} />
            </div>
          </div>
        </div>
      </div>:null}
      <div className={`${styles.rtfEditor} ${"rtf"}`} >
        {(!readOnly) 
          ? <div id="editor" className="pell" ref={editorRef} />
          : <div className="pell-content" dangerouslySetInnerHTML={{ __html: value }} />}
      </div>
    </div>
  )
}

Note.propTypes = {
  value: PropTypes.string, 
  patternId: PropTypes.oneOfType([PropTypes.number, PropTypes.string]), 
  patterns: PropTypes.arrayOf(PropTypes.shape({
    id: PropTypes.number.isRequired,
    description: PropTypes.string.isRequired,
    defpattern: PropTypes.number.isRequired
  })).isRequired,
  readOnly: PropTypes.bool.isRequired, 
  lastUpdate: PropTypes.number,
  className: PropTypes.string,
  onEvent: PropTypes.func,
  getText: PropTypes.func,
}

Note.defaultProps = {
  value: "",
  patternId: undefined,
  patterns: [],
  className: "",
  readOnly: false,
  lastUpdate: undefined,
  onEvent: undefined,
  getText: undefined,
}

export default Note;