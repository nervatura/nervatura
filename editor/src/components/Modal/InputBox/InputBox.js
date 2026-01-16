import { LitElement, html, nothing } from 'lit';

import '../../Form/Label/form-label.js'
import '../../Form/Button/form-button.js'
import '../../Form/Input/form-input.js'

import { styles } from './InputBox.styles.js'
import { MODAL_EVENT, BUTTON_TYPE, INPUT_TYPE } from '../../../config/enums.js'

export class InputBox extends LitElement {
  constructor() {
    super();
    this.title = ""
    this.message = ""
    this.infoText = undefined
    this.value = ""
    this.labelCancel = "Cancel"
    this.labelOK = "OK"
    this.defaultOK = false
    this.showValue = false
    this.values = {}
  }

  static get properties() {
    return {
      title: { type: String },
      message: { type: String },
      infoText: { type: String },
      value: { type: String, reflect: true },
      labelOK: { type: String },
      labelCancel: { type: String },
      defaultOK: { type: Boolean },
      showValue: { type: Boolean },
      values: { type: Object }
    };
  }

  static get styles () {
    return [
      styles
    ]
  }

  _onModalEvent(key){
    const input_value = this.renderRoot.querySelector('#input_value') || {}
    const data = {
      value: input_value.value,
      values: this.values
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
  
  render() {
    return html`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label value="${this.title}" class="title-cell" ></form-label>
            </div>
          </div>
          <div class="section" >
            <div class="section-row" >
              <div class="input">${this.message}</div>
              ${(this.infoText)?
                html`<div class="info">${this.infoText}</div>` : nothing}
              ${(this.showValue)?
                html`<div class="info">
                  <form-input id="input_value" type="${INPUT_TYPE.TEXT}" label="${this.title}"
                    value="${this.value}" ?full="${true}"
                    .onEnter=${()=>this._onModalEvent(MODAL_EVENT.OK)}
                  ></form-input>
                </div>` : nothing}
            </div>
          </div>
          <div class="section buttons" >
            <div class="section-row" >
              <div class="cell padding-small half" >
                <form-button id="btn_cancel" icon="Times"
                  @click=${()=>this._onModalEvent(MODAL_EVENT.CANCEL)} 
                  ?full="${true}" label="${this.labelCancel}"
                >${this.labelCancel}</form-button>
              </div>
              <div class="cell padding-small half" >
                <form-button id="btn_ok" icon="Check"
                  @click=${()=>this._onModalEvent(MODAL_EVENT.OK)} 
                  ?autofocus="${(this.showValue) ? false : this.defaultOK}"
                  type="${BUTTON_TYPE.PRIMARY}" ?full="${true}" label="${this.labelOK}"
                >${this.labelOK}</form-button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>`
  }
}