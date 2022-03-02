import { useState } from 'react';
import PropTypes from 'prop-types';

import { store } from 'config/app'
import 'styles/style.css';
import styles from './Report.module.css';

import Label from 'components/Form/Label'
import Button from 'components/Form/Button'
import Input from 'components/Form/Input'
import Select from 'components/Form/Select'
import Icon from 'components/Form/Icon'

export const Report = ({
  title, template, templates, orient, size, copy, className,
  getText, onClose, onOutput,
  ...props 
}) => {
  const [ state, setState ] = useState({
    template: template,
    orient: orient,
    size: size,
    copy: copy
  })
  const reportOutput = (otype) => {
    onOutput({ 
      type: otype, template: state.template, title: title,
      orient: state.orient, size: state.size, copy: state.copy 
    })
  }
  const report_orientation= store.ui.report_orientation.map(item => { 
    return { value: item[0], text: getText(item[1]) } 
  })
  const report_size = store.ui.report_size.map(item => { 
    return { value: item[0], text: item[1] } 
  })
  return(
    <div className={`${"modal"} ${styles.modal}`} >
      <div className={`${"dialog"} ${styles.dialog}`} {...props} >
        <div className={`${styles.panel} ${className}`} >
          <div className={`${styles.panelTitle} ${"primary"}`}>
            <div className="row full">
              <div className="cell">
                <Label value={title} leftIcon={<Icon iconKey="ChartBar" />} iconWidth="20px" />
              </div>
              <div className={`${"cell align-right"} ${styles.closeIcon}`}>
                <Icon id="closeIcon" iconKey="Times" onClick={onClose} />
              </div>
            </div>
          </div>
          <div className="row full container-small section-small">
            <div className="row full">
              <div className={`${"cell padding-small"}`} >
                <div>
                  <Label className="bold" text="msg_template" />
                </div>
                <Select id="template"
                  className="full" value={state.template}
                  onChange={ (value)=>setState({ ...state, template: value }) }
                  options={templates} />
              </div>
            </div>
            <div className="row full">
              <div className={`${"cell padding-small"}`} >
                <div>
                  <Label className="bold" text="msg_report_prop" />
                </div>
                <Select id="orient"
                  value={state.orient}
                  onChange={ (value)=>setState({ ...state, orient: value }) }
                  options={report_orientation} />
                <Select id="size"
                  value={state.size}
                  onChange={ (value)=>setState({ ...state, size: value }) }
                  options={report_size} />
                <Input id="copy"
                  className={`${styles.copyInput}`} 
                  type="integer" value={state.copy} 
                  onChange={ (value)=>setState({ ...state, copy: value }) } />
              </div>
            </div> 
          </div>
          <div className={`${"row full section container-small secondary-title"}`}>
            <div className={`${"row full"}`}>
              <div className={`${"cell padding-small half"}`} >
                <Button id="btn_print"
                  className={`${"full primary"}`} 
                  disabled={(template==="")?"disabled":""}
                  onClick={()=>reportOutput("print")}
                  label={getText("msg_print")} />
              </div>
              <div className={`${"cell padding-small half"}`} >
                <Button id="btn_pdf"
                  className={`${"full primary"}`} 
                  disabled={(template==="")?"disabled":""}
                  onClick={()=>reportOutput("pdf")}
                  label={getText("msg_export_pdf")} />
              </div>
            </div>
            <div className={`${"row full"}`}>
              <div className={`${"cell padding-small half"}`} >
                <Button id="btn_xml"
                  className={`${"full primary"}`} 
                  disabled={(template==="")?"disabled":""}
                  onClick={()=>reportOutput("xml")} 
                  label={getText("msg_export_xml")} />
              </div>
              <div className={`${"cell padding-small half"}`} >
                <Button id="btn_printqueue"
                  className={`${"full primary"}`} 
                  disabled={(template==="")?"disabled":""}
                  onClick={()=>reportOutput("printqueue")}
                  label={getText("msg_printqueue")} />
              </div>
            </div>
          </div> 
        </div>
      </div>
    </div>
  )
}

Report.propTypes = {
  /**
   * Form title
   */
  title: PropTypes.string.isRequired,
  /**
   * Default report template
   */
  template: PropTypes.string.isRequired,
  /**
   * Report templates
   */
  templates: PropTypes.array.isRequired,
  /**
   * Default page orientation
   */
  orient: PropTypes.string.isRequired,
  /**
   * Default page size
   */
  size: PropTypes.string.isRequired,
  /**
   * Default copy value 
   */ 
  copy: PropTypes.number.isRequired,
  className: PropTypes.string.isRequired,
  /**
   * Localization
   */
  getText: PropTypes.func, 
   /**
    * Close form handle (modal style)
    */ 
  onClose: PropTypes.func,
  /**
   * Selected output type handle
   */
  onOutput: PropTypes.func,
}

Report.defaultProps = {
  title: "",
  template: "",
  templates: [],
  orient: "portrait",
  size: "a4", 
  copy: 1,
  className: "",
  getText: undefined, 
  onClose: undefined,
  onOutput: undefined,
}

export default Report;