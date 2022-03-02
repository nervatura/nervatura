import { Fragment } from 'react';
import PropTypes from 'prop-types';

import Icon from 'components/Form/Icon'
import Label from 'components/Form/Label'
import Button from 'components/Form/Button'

import Main from 'components/Editor/Main'
import Meta from 'components/Editor/Meta'
import Note from 'components/Editor/Note'
import Item from 'components/Editor/Item'
import View from 'components/Editor/View'

import styles from './Editor.module.css';

export const Editor = ({ 
  caption, current, template, dataset, audit, lastUpdate, className,
  getText, onEvent,
  ...props 
}) => {  
  const tabButton = (viewKey, textKey, iconKey, itemCount) => {
    return <div className="row full" >
      <div className="cell" >
        <Button id={"btn_"+viewKey} key={viewKey}
          className={` ${styles.tabButton} ${"full secondary-title"}`} 
          onClick={()=>onEvent(
            "changeCurrentData",["view", (current.view === viewKey) ? "" : viewKey ]
          )}
          value={<div className="row full" >
            <div className="cell" >
              <Label value={getText(textKey)} 
                leftIcon={<Icon iconKey={iconKey} />} iconWidth="20px" />
            </div>
            {(typeof itemCount !== "undefined")?<div className="cell align-right" >
              <span className={`${styles.badge}`} >{itemCount}</span>
            </div>:null} 
          </div>}
        />
      </div>
    </div>
  }
  const mainLabel = () => {
    let label = current.item[template.options.title_field]
    if(current.type === "printqueue"){
      label = template.options.title_field;
    } else if(current.item.id === null){
      label = getText("label_new")+" "+template.options.title;
    }
    return label
  }
  const metafieldCount = () => {
    return current.fieldvalue.filter(fv => {
      let _deffield = dataset.deffield.filter((df) => (df.fieldname === fv.fieldname))[0]
      return ((_deffield.visible === 1) && (fv.deleted === 0))
    }).length
  }
  return (
    <div {...props} className={`${styles.width800} ${className}`}>
      <div className={`${styles.panel}`} >
        <div className={`${styles.panelTitle} ${"primary"}`}>
          <Label value={caption} 
            leftIcon={<Icon iconKey={template.options.icon} />} iconWidth="20px" />
        </div>
        {(current.form)?
        <div className="section container" >
          <Item current={current} dataset={dataset} audit={audit}
            onEvent={onEvent} getText={getText} />
        </div>:
        <div className="section container" >
          <Fragment >
            {tabButton("form", mainLabel(), template.options.icon)}
            {(current.view === "form")?<Main 
              current={current} template={template} dataset={dataset} audit={audit}
              onEvent={onEvent} getText={getText} />:null}                  
          </Fragment>

          {((current.item.id !== null || template.options.search_form) 
            && typeof dataset.fieldvalue !== "undefined" 
            && current.item !== null && template.options.fieldvalue === true)?
            <Fragment >
              {tabButton("fieldvalue", "fields_view", template.options.icon, metafieldCount())}
              {(current.view === "fieldvalue")?<Meta 
                current={current} dataset={dataset} audit={audit}
                onEvent={onEvent} getText={getText} />:null}
            </Fragment>:null}

          {((current.item.id !== null) && (typeof current.item.fnote !== "undefined") && 
            (template.options.pattern === true)) ? 
            <Fragment>
              {tabButton("fnote", "fnote_view", "Comment")}
              {(current.view === "fnote")?<Note 
                value={current.item.fnote}
                lastUpdate={lastUpdate}
                patternId={current.template} 
                patterns={dataset.pattern} 
                readOnly={(audit === "readonly")} 
                onEvent={onEvent} getText={getText} />:null}
            </Fragment>: null}

          {Object.keys(template.view).filter(
            (vname)=>(template.view[vname].view_audit !== "disabled")).map(
              (vname) => <Fragment key={vname} >
                {tabButton(vname, template.view[vname].title, 
                  template.view[vname].icon, 
                  dataset[template.view[vname].data].length)}
                {(current.view === vname)?<View viewName={vname}
                  current={current} template={template} dataset={dataset} audit={audit}
                  onEvent={onEvent} getText={getText} />:null}
              </Fragment>
            )}
        </div>}
      </div>
    </div>
  )
}

Editor.propTypes = {
  caption: PropTypes.string.isRequired, 
  current: PropTypes.object.isRequired, 
  template: PropTypes.object.isRequired, 
  dataset: PropTypes.object.isRequired, 
  audit: PropTypes.string.isRequired,
  lastUpdate: PropTypes.number,
  className: PropTypes.string,
  onEvent: PropTypes.func,
  getText: PropTypes.func,
}

Editor.defaultProps = {
  caption: "", 
  current: {}, 
  template: {}, 
  dataset: {}, 
  audit: "",
  lastUpdate: undefined,
  className: "",
  onEvent: undefined,
  getText: undefined,
}

export default Editor;