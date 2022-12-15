import { LitElement, html } from 'lit';

import '../../Form/Label/form-label.js'
import '../../Form/Button/form-button.js'
import '../../Form/Input/form-input.js'
import '../../Form/Select/form-select.js'

import { styles } from './Formula.styles.js'
import { MODAL_EVENT, BUTTON_TYPE, INPUT_TYPE } from '../../../config/enums.js'

export class Formula extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.formula = "" 
    this.formulaValues = [] 
    this.partnumber = ""
    this.description = ""
  }

  static get properties() {
    return {
      formula: { type: String },
      formulaValues: { type: Array },
      partnumber: { type: String },
      description: { type: String },
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

  _onFormulaChange(value){
    this.formula = value
  }
  
  render() {
    const { formula, partnumber, formulaValues, description } = this
    return html`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label leftIcon="Magic"
                value="${this.msg("", { id: "label_formula" })}" 
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
                    label="${partnumber}" .style=${{ "font-weight": "bold" }}
                    value="${partnumber}" ?disabled=${true} ?full=${true}
                  ></form-input>
                </div>
              </div>
              <div class="row full">
                <div class="cell padding-small" >
                  <form-input type="${INPUT_TYPE.TEXT}"
                    label="${description}" value="${description}" 
                    ?disabled=${true} ?full=${true}
                  ></form-input>
                </div>
              </div>
              <div class="row full">
                <div class="cell padding-small" >
                  <form-select id="sel_formula" label="${this.msg("", {id: "label_formula"})}"
                    .onChange=${(value) => this._onFormulaChange(value.value) }
                    .options=${formulaValues} .isnull="${true}" value="${formula}" 
                  ></form-select>
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
                  @click=${()=>this._onModalEvent(MODAL_EVENT.OK, { value: parseInt(formula, 10) })} 
                  ?disabled="${(formula === "")}"
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