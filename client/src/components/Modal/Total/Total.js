import { LitElement, html } from 'lit';

import '../../Form/Label/form-label.js'
import '../../Form/Button/form-button.js'
import '../../Form/Icon/form-icon.js'
import '../../Form/Input/form-input.js'


import { styles } from './Total.style.js'
import { MODAL_EVENT, BUTTON_TYPE, INPUT_TYPE } from '../../../config/enums.js'

export class Total extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.total = {
      totalFields: {},
      totalLabels: {},
      count: 0
    }
  }

  static get properties() {
    return {
      total: { type: Object }
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
  
  render() {
    return html`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label leftIcon="InfoCircle"
                value="${this.msg("Total", { id: "browser_total" })}" 
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
              ${Object.keys(this.total.totalFields).map(fieldname => 
                html`<div class="table-row full">
                  <div class="cell padding-tiny mobile">
                    <form-label
                      value="${this.total.totalLabels[fieldname]}" 
                    ></form-label>
                  </div>
                  <div class="cell padding-tiny mobile">
                    <form-input type="${INPUT_TYPE.TEXT}"
                      label="${this.total.totalLabels[fieldname]}"
                      .style=${{ "font-weight": "bold", "text-align": "right" }}
                      value="${new Intl.NumberFormat('default').format(this.total.totalFields[fieldname])}"
                      ?disabled=${true} ?full=${true}
                    ></form-input>
                  </div>
                </div>`)}
            </div>
          </div>
          <div class="section buttons" >
            <div class="section-row" >
              <div class="cell padding-small" >
                <form-button id="btn_ok" icon="Check"
                  label="${this.msg("", { id: "msg_ok" })}"
                  @click=${()=>this._onModalEvent(MODAL_EVENT.CANCEL, {})} 
                  ?autofocus="${true}"
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