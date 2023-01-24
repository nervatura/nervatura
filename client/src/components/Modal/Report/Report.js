import { LitElement, html } from 'lit';

import '../../Form/Label/form-label.js'
import '../../Form/Button/form-button.js'
import '../../Form/Select/form-select.js'
import '../../Form/NumberInput/form-number.js'

import { styles } from './Report.styles.js'
import { MODAL_EVENT, BUTTON_TYPE } from '../../../config/enums.js'

export class Report extends LitElement {
  constructor() {
    super();
    this.title = ""
    this.template = ""
    this.templates = []
    this.report_orientation = []
    this.report_size = []
    this.orient = "portrait"
    this.size = "a4" 
    this.copy = 1
  }

  static get properties() {
    return {
      title: { type: String },
      templates: { type: Array },
      report_orientation: { type: Array },
      report_size: { type: Array },
      template: { type: String },
      orient: { type: String },
      size: { type: String },
      copy: { type: Number },
    };
  }

  static get styles () {
    return [
      styles
    ]
  }

  _onModalEvent(key, value){
    const { template, orient, size, copy, title } = this
    const data = {
      type: value, template, orient, size, copy, title
    }
    if(this.onEvent && this.onEvent.onModalEvent){
      this.onEvent.onModalEvent({ key, data })
    }
    this.dispatchEvent(
      new CustomEvent('modal_event', {
        bubbles: true, composed: true,
        detail: {
          key, data
        }
      })
    );
  }

  _onValueChange(key, value){
    this[key] = value
  }
  
  render() {
    const { title, template, templates, orient, size, copy, report_size, report_orientation } = this
    const _report_orientation= report_orientation.map(item => ({ value: item[0], text: this.msg("", { id: item[1] }) }))
    const _report_size = report_size.map(item => ({ value: item[0], text: item[1] }))
    return html`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label leftIcon="ChartBar"
                value="${title}" 
                class="title-cell" ></form-label>
            </div>
            <div class="cell align-right" >
              <span id=${`closeIcon`} class="close-icon" 
                @click="${ ()=>this._onModalEvent(MODAL_EVENT.CANCEL, "") }">
                <form-icon iconKey="Times" ></form-icon>
              </span>
            </div>
          </div>
          <div class="section" >
            <div class="section-row" >
              <div class="row full">
                <div class="cell padding-small" >
                  <div class="label-padding">
                    <form-label
                      value="${this.msg("", { id: "msg_template" })}" 
                    ></form-label>
                  </div>
                  <form-select id="template" 
                    label="${this.msg("", { id: "msg_template" })}"
                    .onChange=${(event) => this._onValueChange("template", event.value)}
                    .options=${templates} .isnull="${true}" value="${template}" 
                  ></form-select>
                </div>
              </div>
            </div>
            <div class="section-row" >
              <div class="row full">
                <div class="cell padding-small" >
                  <div class="label-padding" >
                    <form-label
                      value="${this.msg("", { id: "msg_report_prop" })}" 
                    ></form-label>
                  </div>
                  <div class="cell" >
                    <form-select id="orient" 
                      label="${orient}"
                      .onChange=${(event) => this._onValueChange("orient", event.value)}
                      .options=${_report_orientation} .isnull="${false}" value="${orient}" 
                    ></form-select>
                  </div>
                  <div class="cell" >
                    <form-select id="size" 
                      label="${size}"
                      .onChange=${(event) => this._onValueChange("size", event.value)}
                      .options=${_report_size} .isnull="${false}" value="${size}" 
                    ></form-select>
                  </div>
                  <div class="cell" >
                    <form-number id="copy" 
                      label="${copy}" .style=${{ width: "60px" }}
                      ?integer="${true}" value="${copy}"
                      .onChange=${(event) => this._onValueChange("copy", event.value)}
                    ></form-number>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="section buttons" >
            <div class="section-row" >
              <div class="cell padding-small half" >
                <form-button id="btn_print"
                  @click=${()=>this._onModalEvent(MODAL_EVENT.OK, "print" )}
                  ?disabled="${(template==="")}"
                  type="${BUTTON_TYPE.PRIMARY}" ?full="${true}" 
                  label="${this.msg("", { id: "msg_print" })}"
                >${this.msg("", { id: "msg_print" })}</form-button>
              </div>
              <div class="cell padding-small half" >
                <form-button id="btn_pdf"
                  @click=${()=>this._onModalEvent(MODAL_EVENT.OK, "pdf" )} 
                  ?disabled="${(template==="")}"
                  type="${BUTTON_TYPE.PRIMARY}" ?full="${true}" 
                  label="${this.msg("", { id: "msg_export_pdf" })}"
                >${this.msg("", { id: "msg_export_pdf" })}</form-button>
              </div>
            </div>
            <div class="section-row" >
              <div class="cell padding-small half" >
                <form-button id="btn_xml"
                  @click=${()=>this._onModalEvent(MODAL_EVENT.OK, "xml" )}
                  ?disabled="${(template==="")}"
                  type="${BUTTON_TYPE.PRIMARY}" ?full="${true}" 
                  label="${this.msg("", { id: "msg_export_xml" })}"
                >${this.msg("", { id: "msg_export_xml" })}</form-button>
              </div>
              <div class="cell padding-small half" >
                <form-button id="btn_printqueue"
                  @click=${()=>this._onModalEvent(MODAL_EVENT.OK, "printqueue" )} 
                  ?disabled="${(template==="")}"
                  type="${BUTTON_TYPE.PRIMARY}" ?full="${true}" 
                  label="${this.msg("", { id: "msg_printqueue" })}"
                >${this.msg("", { id: "msg_printqueue" })}</form-button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>`
  }
}