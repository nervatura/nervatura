import { LitElement, html } from 'lit';

import '../../Form/Label/form-label.js'
import '../../Form/Button/form-button.js'
import '../../Form/Input/form-input.js'
import '../../Form/NumberInput/form-number.js'
import '../../Form/Select/form-select.js'

import { styles } from './Menu.styles.js'
import { MODAL_EVENT, BUTTON_TYPE, INPUT_TYPE } from '../../../config/enums.js'

export class Menu extends LitElement {
  constructor() {
    super();
    this.idKey = null
    this.menu_id = 0 
    this.fieldname = ""
    this.description = "" 
    this.fieldtype = 0
    this.fieldtypeOptions = [] 
    this.orderby = 0
  }

  static get properties() {
    return {
      idKey: { type: Object },
      menu_id: { type: Number },
      fieldname: { type: String },
      description: { type: String },
      fieldtype: { type: Number },
      fieldtypeOptions: { type: Array },
      orderby: { type: Number },
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
  
  render() {
    const { idKey, menu_id, fieldname, description, orderby, fieldtype, fieldtypeOptions } = this
    return html`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label leftIcon="Share"
                value="${this.msg("", { id: "title_menucmd" })}" 
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
              <div class="row full">
                <div class="cell padding-small" >
                  <div>
                    <form-label
                      value="${this.msg("", { id: "menufields_fieldname" })}" 
                    ></form-label>
                  </div>
                  <form-input id="fieldname"
                    type="${INPUT_TYPE.TEXT}"
                    label="${this.msg("", { id: "menufields_fieldname" })}" 
                    .onChange=${(event) => this._onValueChange("fieldname", event.value)}
                    value="${fieldname}" ?full=${true}
                  ></form-input>
                </div>
              </div>
              <div class="row full">
                <div class="cell padding-small" >
                  <div>
                    <form-label
                      value="${this.msg("", { id: "menufields_description" })}" 
                    ></form-label>
                  </div>
                  <form-input id="description"
                    type="${INPUT_TYPE.TEXT}"
                    label="${this.msg("", { id: "menufields_description" })}" 
                    .onChange=${(event) => this._onValueChange("description", event.value)}
                    value="${description}" ?full=${true}
                  ></form-input>
                </div>
              </div>
              <div class="row full">
                <div class="cell padding-small half" >
                  <div>
                    <form-label
                      value="${this.msg("", { id: "menufields_fieldtype" })}" 
                    ></form-label>
                  </div>
                  <form-select id="fieldtype" 
                    label="${this.msg("", {id: "menufields_fieldtype"})}" ?full="${true}"
                    .onChange=${(value) => this._onValueChange("fieldtype", parseInt(value.value,10)) }
                    .options=${fieldtypeOptions} .isnull="${false}" value="${String(fieldtype)}" 
                  ></form-select>
                </div>
                <div class="cell padding-small half" >
                  <div>
                    <form-label
                      value="${this.msg("", { id: "menufields_orderby" })}" 
                    ></form-label>
                  </div>
                  <form-number id="orderby" 
                    label="${this.msg("", { id: "menufields_orderby" })}"
                    ?integer="${false}" value="${orderby}"
                    .onChange=${(event) => this._onValueChange("orderby", event.value)}
                  ></form-number>
                </div>
              </div>
            </div>
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
                  ?disabled="${(fieldname === "") || (description === "")}"
                  @click=${()=>this._onModalEvent(MODAL_EVENT.OK, { 
                    value: { id: idKey, menu_id, fieldname, description, fieldtype, orderby } 
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