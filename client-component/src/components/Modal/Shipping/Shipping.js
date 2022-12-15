import { LitElement, html } from 'lit';

import '../../Form/Label/form-label.js'
import '../../Form/Button/form-button.js'
import '../../Form/Input/form-input.js'
import '../../Form/NumberInput/form-number.js'

import { styles } from './Shipping.styles.js'
import { MODAL_EVENT, BUTTON_TYPE, INPUT_TYPE } from '../../../config/enums.js'

export class Shipping extends LitElement {
  constructor() {
    super();
    this.partnumber = ""
    this.description = ""
    this.unit = ""
    this.batch_no = ""
    this.qty = 0
  }

  static get properties() {
    return {
      partnumber: { type: String },
      description: { type: String },
      unit: { type: String },
      batch_no: { type: String },
      qty: { type: Number },
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
    const { partnumber, description, unit, batch_no, qty } = this
    return html`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label leftIcon="Truck"
                value="${this.msg("", { id: "shipping_movement_product" })}" 
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
                      value="${this.msg("", { id: "product_partnumber" })}" 
                    ></form-label>
                  </div>
                  <form-input type="${INPUT_TYPE.TEXT}"
                    label="${this.msg("", { id: "product_partnumber" })}" 
                    .style=${{ "font-weight": "bold" }}
                    value="${partnumber}" ?disabled=${true} ?full=${true}
                  ></form-input>
                </div>
              </div>
              <div class="row full">
                <div class="cell padding-small" >
                  <div>
                    <form-label
                      value="${this.msg("", { id: "product_description" })}" 
                    ></form-label>
                  </div>
                  <form-input type="${INPUT_TYPE.TEXT}"
                    label="${this.msg("", { id: "product_description" })}" 
                    value="${description}" ?disabled=${true} ?full=${true}
                  ></form-input>
                </div>
              </div>
              <div class="row full">
                <div class="cell padding-small" >
                  <div>
                    <form-label
                      value="${this.msg("", { id: "product_unit" })}" 
                    ></form-label>
                  </div>
                  <form-input type="${INPUT_TYPE.TEXT}"
                    label="${this.msg("", { id: "product_unit" })}" 
                    value="${unit}" ?disabled=${true} ?full=${true}
                  ></form-input>
                </div>
              </div>
              <div class="row full">
                <div class="cell padding-small" >
                  <div>
                    <form-label
                      value="${this.msg("", { id: "movement_batchnumber" })}" 
                    ></form-label>
                  </div>
                  <form-input id="batch_no" type="${INPUT_TYPE.TEXT}"
                    label="${this.msg("", { id: "movement_batchnumber" })}"
                    value="${batch_no}" ?autofocus=${true}
                    .onChange=${(event) => this._onValueChange("batch_no", event.value)}
                    ?full=${true}
                  ></form-input>
                </div>
              </div>
              <div class="row full">
                <div class="cell padding-small" >
                  <div>
                    <form-label
                      value="${this.msg("", { id: "movement_qty" })}" 
                    ></form-label>
                  </div>
                  <form-number id="qty" 
                    label="${this.msg("", { id: "movement_qty" })}"
                    ?integer="${false}" value="${qty}"
                    .onChange=${(event) => this._onValueChange("qty", event.value)}
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
                  @click=${()=>this._onModalEvent(MODAL_EVENT.OK, { batch_no, qty })} 
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