import { LitElement, html, nothing } from 'lit';

import '../../Form/Label/form-label.js'
import '../../Form/Button/form-button.js'
import '../../Form/Select/form-select.js'

import { styles } from './Trans.styles.js'
import { MODAL_EVENT, BUTTON_TYPE } from '../../../config/enums.js'

export class Trans extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.baseTranstype = ""
    this.transtype = ""
    this.direction = ""
    this.doctypes = []
    this.directions = []
    this.refno = true
    this.nettoDiv = false 
    this.netto = true 
    this.fromDiv = false 
    this.from = false
    this.elementCount = 0
  }

  static get properties() {
    return {
      baseTranstype: { type: String },
      transtype: { type: String }, 
      direction: { type: String }, 
      doctypes: { type: Array }, 
      directions: { type: Array }, 
      refno: { type: Boolean }, 
      nettoDiv: { type: Boolean }, 
      netto: { type: Boolean }, 
      fromDiv: { type: Boolean }, 
      from: { type: Boolean }, 
      elementCount: { type: Number },
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

  _setTranstype(value) {
    const { baseTranstype, elementCount } = this
    if(["invoice","receipt"].includes(value) && ["order","rent","worksheet"].includes(baseTranstype)){
        this.nettoDiv = true
        if(elementCount===0){
          this.fromDiv = true
        }
    } else {
      this.nettoDiv = false
      this.fromDiv = false
    }
    this.transtype = value
  }

  render() {
    const { transtype, direction, refno, from, fromDiv, netto, nettoDiv, doctypes, directions } = this
    const typeOptions = doctypes.map(dt => ({ value: dt, text: dt }))
    const dirOptions = directions.map(dir => ({ value: dir, text: dir }))
    return html`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label leftIcon="FileText"
                value="${this.msg("", { id: "msg_create_title" })}" 
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
              <div class="cell padding-small" >
                <form-label
                  value="${this.msg("", { id: "msg_create_new" })}" 
                ></form-label>
              </div>
            </div>
            <div class="section-row" >
              <div class="cell padding-small half" >
                <form-select id="transtype" label="transtype" ?full=${true}
                  .onChange=${(value) => this._setTranstype(value.value)}
                  .options=${typeOptions} .isnull="${false}" value="${transtype}" 
                ></form-select>
              </div>
              <div class="cell padding-small half" >
                <form-select id="direction" label="direction" ?full=${true}
                  .onChange=${(value) => this._onValueChange("direction", value.value)}
                  .options=${dirOptions} .isnull="${false}" value="${direction}" 
                ></form-select>
              </div>
            </div>
            <div class="section-row" >
              <div class="cell padding-small" >
                <form-label id="refno"
                  value="${this.msg("", { id: "msg_create_setref" })}"
                  leftIcon="${(refno)?"CheckSquare":"SquareEmpty"}"
                  .style=${{ cursor: "pointer" }} .iconStyle=${{ cursor: "pointer" }}
                  @click=${() => this._onValueChange("refno", !refno)}
                ></form-label>
              </div>
            </div>
            ${(nettoDiv) ? html`<div class="section-row" >
              <div class="cell padding-small" >
                <form-label id="netto"
                  value="${this.msg("", { id: "msg_create_deduction" })}"
                  leftIcon="${(netto)?"CheckSquare":"SquareEmpty"}"
                  .style=${{ cursor: "pointer" }} .iconStyle=${{ cursor: "pointer" }}
                  @click=${() => this._onValueChange("netto", !netto)}
                ></form-label>
              </div>
            </div>` : nothing}
            ${(fromDiv) ? html`<div class="section-row" >
              <div class="cell padding-small" >
                <form-label id="from"
                  value="${this.msg("", { id: "msg_create_delivery" })}"
                  leftIcon="${(from)?"CheckSquare":"SquareEmpty"}"
                  .style=${{ cursor: "pointer" }} .iconStyle=${{ cursor: "pointer" }}
                  @click=${() => this._onValueChange("from", !from)}
                ></form-label>
              </div>
            </div>` : nothing}
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
                    newTranstype: transtype, 
                    newDirection: direction, 
                    refno, 
                    fromInventory: (from && fromDiv), 
                    nettoQty: (netto && nettoDiv)
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