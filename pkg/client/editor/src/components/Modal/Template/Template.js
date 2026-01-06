import { LitElement, html, nothing } from 'lit';

import '../../Form/Label/form-label.js'
import '../../Form/Button/form-button.js'
import '../../Form/Select/form-select.js'
import '../../Form/Input/form-input.js'
import '../../Modal/InputBox/modal-inputbox.js'

import { styles } from './Template.styles.js'
import { MODAL_EVENT, BUTTON_TYPE, TEMPLATE_DATA_TYPE, INPUT_TYPE } from '../../../config/enums.js'

export class Template extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.type = TEMPLATE_DATA_TYPE.TEXT
    this.name = ""
    this.columns = ""
  }

  static get properties() {
    return {
      type: { type: String },
      name: { type: String }, 
      columns: { type: String }
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

  _onValueChange(key, value){
    this[key] = value
  }

  _onTextInput(event){
    this._onValueChange("columns", event.target.value)
  }

  render() {
    const { type, name, columns } = this
    return html`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label leftIcon="Plus"
                value="${this.msg("", { id: "template_label_new_data" })}" 
                class="title-cell" ></form-label>
            </div>
            <div class="cell align-right" >
              <span id=${`closeIcon`} class="close-icon" 
                @click="${ ()=>this._onModalEvent(MODAL_EVENT.CANCEL, {}) }">
                <form-icon iconKey="Times" ></form-icon>
              </span>
            </div>
          </div>
          <div class="section" >
            <div class="section-row" >
              <div class="cell padding-small half" >
                <div>
                  <form-label
                    value="${this.msg("", { id: "template_data_type" })}" 
                  ></form-label>
                </div>
                <form-select id="type" 
                  label="${this.msg("", { id: "template_data_type" })}" 
                  ?full=${true} .isnull="${false}" value="${type}"
                  .onChange=${(event) => this._onValueChange("type", event.value)}
                  .options=${Object.keys(TEMPLATE_DATA_TYPE).map(key => ({ value: TEMPLATE_DATA_TYPE[key], text: key }))}  
                ></form-select>
              </div>
              <div class="cell padding-small half" >
                <div>
                  <form-label
                    value="${this.msg("", { id: "template_data_name" })}" 
                  ></form-label>
                </div>
                <form-input id="name"
                  type="${INPUT_TYPE.TEXT}"
                  label="${this.msg("", { id: "template_data_name" })}" 
                  .onChange=${(event) => this._onValueChange("name", event.value)}
                  value="${name}" ?full=${true}
                ></form-input>
              </div>
            </div>
            ${(type === TEMPLATE_DATA_TYPE.TABLE) ? html`<div class="section-row" >
              <div class="cell padding-small" >
                <div>
                  <form-label
                    value="${this.msg("", { id: "template_data_columns" })}" 
                  ></form-label>
                </div>
                  <textarea id="columns"
                    rows=3 .value="${columns}"
                    @input="${this._onTextInput}"
                  ></textarea>
              </div>
            </div>` : nothing }
          </div>
          <div class="section buttons" >
            <div class="section-row" >
              <div class="cell padding-small half" >
                <form-button id="btn_cancel" icon="Times"
                  @click=${()=>this._onModalEvent(MODAL_EVENT.CANCEL, {})} 
                  ?full="${true}" label="${this.msg("", { id: "msg_cancel" })}"
                >${this.msg("", { id: "msg_cancel" })}</form-button>
              </div>
              <div class="cell padding-small half" >
                <form-button id="btn_ok" icon="Check"
                  @click=${()=>this._onModalEvent(MODAL_EVENT.OK, { 
                    value: { name, type, columns }
                  })} 
                  type="${BUTTON_TYPE.PRIMARY}" ?full="${true}" 
                  label="${this.msg("", { id: "msg_ok" })}"
                >${this.msg("", { id: "msg_ok" })}</form-button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>`
  }
}