import { LitElement, html } from 'lit';

import '../../Form/Label/form-label.js'
import '../../Form/Button/form-button.js'
import '../../Form/Icon/form-icon.js'
import '../../Form/Row/form-row.js'


import { styles } from './Server.styles.js'
import { MODAL_EVENT, BUTTON_TYPE } from '../../../config/enums.js'

export class Server extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.cmd = {} 
    this.fields = []
    this.values = {}
  }

  static get properties() {
    return {
      cmd: { type: Object },
      fields: { type: Array },
      values: { type: Object }
    };
  }

  static get styles () {
    return [
      styles
    ]
  }

  _onModalEvent(key, data){
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

  _onEdit(data){
    this.values[data.name] = data.value
    this.requestUpdate()
  }
  
  render() {
    return html`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label leftIcon="Share"
                value="${this.cmd.description}" 
                class="title-cell" ></form-label>
            </div>
            <div class="cell align-right" >
              <span id=${`closeIcon`} class="close-icon" 
                @click="${ ()=>this._onModalEvent(MODAL_EVENT.CANCEL, {}) }">
                <form-icon iconKey="Times" ></form-icon>
              </span>
            </div>
          </div>
          ${this.fields.map((field, index) =>
            html`<form-row id=${`row_${index}`}
              .row=${{
                rowtype: "field", 
                name: field.fieldname,
                datatype: field.fieldtypeName,
                label: field.description 
              }} 
              .values=${{...this.values}}
              .options=${{}}
              .data=${{
                audit: "all",
                current: {},
                dataset: {},  
              }}
              .onEdit=${(data)=>{
                this._onEdit(data)
              }}
              .msg=${this.msg}
            ></form-row>`
          )}
          <div class="section buttons" >
            <div class="section-row" >
              <div class="cell padding-small half" >
                <form-button id="btn_cancel" icon="Times"
                  label="${this.msg("", { id: "msg_cancel" })}"
                  @click=${()=>this._onModalEvent(MODAL_EVENT.CANCEL, {})} 
                  ?full="${true}" 
                >${this.msg("", { id: "msg_cancel" })}</form-button>
              </div>
              <div class="cell padding-small half" >
                <form-button id="btn_ok" icon="Check"
                  label="${this.msg("", { id: "msg_ok" })}"
                  @click=${()=>this._onModalEvent(MODAL_EVENT.OK, {
                    cmd: this.cmd, fields: this.fields, 
                    values: {...this.values}
                  })} 
                  type="${BUTTON_TYPE.PRIMARY}" ?full="${true}" 
                >${this.msg("", { id: "msg_ok" })}</form-button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>`
  }
}