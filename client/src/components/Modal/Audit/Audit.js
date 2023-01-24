import { LitElement, html, nothing } from 'lit';

import '../../Form/Label/form-label.js'
import '../../Form/Button/form-button.js'
import '../../Form/Select/form-select.js'

import { styles } from './Audit.styles.js'
import { MODAL_EVENT, BUTTON_TYPE } from '../../../config/enums.js'

export class Audit extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.idKey = null
    this.usergroup = 0
    this.nervatype = undefined
    this.subtype = null
    this.inputfilter = undefined
    this.supervisor = 0
    this.typeOptions = []
    this.subtypeOptions = []
    this.inputfilterOptions = []
  }

  static get properties() {
    return {
      idKey: { type: Object },
      usergroup: { type: Number },
      nervatype: { type: Number },
      subtype: { type: Object },
      inputfilter: { type: Number },
      supervisor: { type: Number },
      typeOptions: { type: Array },
      subtypeOptions: { type: Array },
      inputfilterOptions: { type: Array },
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
    const { 
      idKey, usergroup, nervatype, subtype, inputfilter, supervisor,
      typeOptions, subtypeOptions, inputfilterOptions 
    } = this
    const nervatypeName = typeOptions.filter(item => (item.value === String(nervatype)))[0].text
    const isSubtype = ["trans", "report","menu"].includes(typeOptions.filter(item => (item.value === String(nervatype)))[0].text)
    return html`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label leftIcon="Key"
                value="${this.msg("", { id: "title_usergroup" })}" 
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
                      value="${this.msg("", { id: "audit_nervatype" })}" 
                    ></form-label>
                  </div>
                  <form-select id="nervatype" 
                    label="${this.msg("", {id: "audit_nervatype"})}"
                    ?disabled="${(idKey)}" ?full="${true}"
                    .onChange=${(value) => this._onValueChange("nervatype", parseInt(value.value,10)) }
                    .options=${typeOptions} .isnull="${false}" value="${String(nervatype)}" 
                  ></form-select>
                </div>
              </div>
              ${(isSubtype) ? html`<div class="row full">
                <div class="cell padding-small" >
                  <div>
                    <form-label
                      value="${this.msg("", { id: "audit_subtype" })}" 
                    ></form-label>
                  </div>
                  <form-select id="subtype" 
                    label="${this.msg("", {id: "audit_subtype"})}"
                    .onChange=${(value) => this._onValueChange("subtype", parseInt(value.value,10)) }
                    .options=${subtypeOptions.filter(item => (item.type === nervatypeName))} 
                    .isnull="${false}" value="${String(subtype)}" ?full="${true}"
                  ></form-select>
                </div>
              </div>` : nothing}
              <div class="row full">
                <div class="cell padding-small" >
                  <div>
                    <form-label
                      value="${this.msg("", { id: "audit_inputfilter" })}" 
                    ></form-label>
                  </div>
                  <form-select id="inputfilter" 
                    label="${this.msg("", {id: "audit_inputfilter"})}" ?full="${true}"
                    .onChange=${(value) => this._onValueChange("inputfilter", parseInt(value.value,10)) }
                    .options=${inputfilterOptions} .isnull="${false}" value="${String(inputfilter)}" 
                  ></form-select>
                </div>
              </div>
              <div class="section-row" >
                <div class="cell padding-small" >
                  <form-label id="supervisor"
                    value="${this.msg("", { id: "audit_supervisor" })}"
                    leftIcon="${(supervisor === 1) ? "CheckSquare" : "SquareEmpty"}"
                    .style=${{ cursor: "pointer" }} .iconStyle=${{ cursor: "pointer" }}
                    @click=${() => this._onValueChange("supervisor", (supervisor === 1) ? 0 : 1)}
                  ></form-label>
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
                  @click=${()=>this._onModalEvent(MODAL_EVENT.OK, { 
                    value: { id: idKey, usergroup, nervatype, subtype, inputfilter, supervisor } })} 
                  ?disabled="${(isSubtype && !subtype)}"
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