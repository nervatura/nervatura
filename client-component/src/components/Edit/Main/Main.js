import { LitElement, html, nothing } from 'lit';

import '../../Form/Row/form-row.js'

import { styles } from './Main.styles.js'
import { EDIT_EVENT } from '../../../config/enums.js'

export class Main extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.id = Math.random().toString(36).slice(2);
    this.audit = ""
    this.current = {} 
    this.template = {}
    this.dataset = {}
    this.onEvent = {}
  }

  static get properties() {
    return {
      id: { type: String },
      audit: { type: String },
      current: { type: Object },
      template: { type: Object },
      dataset: { type: Object },
      onEvent: { type: Object }
    };
  }

  static get styles () {
    return [
      styles
    ]
  }

  _onEditEvent(key, data){
    if(this.onEvent.onEditEvent){
      this.onEvent.onEditEvent({ key, data })
    }
    this.dispatchEvent(
      new CustomEvent('edit_event', {
        bubbles: true, composed: true,
        detail: {
          key, data
        }
      })
    );
  }

  render() {
    const { audit, current, dataset, template } = this
    return html`<div id="${this.id}" class="panel" >
      ${template.rows.map((row, index) => html`<form-row
        id=${`row_${index}`}
        .row="${row}"
        .values="${current.item}"
        .options="${template.options}"
        .data="${{ audit, current, dataset }}"
        .onEdit=${(data)=>this._onEditEvent(EDIT_EVENT.EDIT_ITEM, data )}
        .onEvent=${(...args) => this._onEditEvent(...args)}
        .onSelector=${(...args)=>this._onEditEvent(EDIT_EVENT.SELECTOR, [...args] )}
        .msg=${this.msg}
      ></form-row>`)}
      ${((current.type === "report") && (current.fieldvalue.length > 0)) ?
        html`<div class="row full">
          ${current.fieldvalue.map((row, index) => html`<form-row
            id=${`row_${index}`}
            .row="${row}"
            .values="${row}"
            .options="${template.options}"
            .data="${{ audit, current, dataset }}"
            .onEdit=${(data)=>this._onEditEvent(EDIT_EVENT.EDIT_ITEM, data )}
            .onEvent=${(...args) => this._onEditEvent(...args)}
            .onSelector=${(...args)=>this._onEditEvent(EDIT_EVENT.SELECTOR, [...args] )}
            .msg=${this.msg}
          ></form-row>`)}
        </div>` : nothing }
    </div>`
  }
}