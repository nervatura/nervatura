import { LitElement, html } from 'lit';

import '../../Form/Row/form-row.js'

import { styles } from './Item.styles.js'
import { EDIT_EVENT } from '../../../config/enums.js'

export class Item extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.id = Math.random().toString(36).slice(2);
    this.audit = ""
    this.current = {} 
    this.dataset = {}
    this.onEvent = {}
  }

  static get properties() {
    return {
      id: { type: String },
      audit: { type: String },
      current: { type: Object },
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
    const { audit, current, dataset } = this
    return html`<div class="row full">
      <div class="title-cell" >
        <span class="title-pre" >${(current.form.id === null) 
          ? this.msg("", { id: "label_new" }) : String(current.form.id)}</span>
        <span>${current.form_template.options.title}</span>
      </div>
    </div>
    <div id="${this.id}" class="panel" >
      ${current.form_template.rows.map((row, index) => html`<form-row
        id=${`row_${index}`}
        .row="${row}"
        .values="${current.form}"
        .options="${current.form_template.options}"
        .data="${{ audit, current, dataset }}"
        .onEdit=${(data)=>this._onEditEvent(EDIT_EVENT.EDIT_ITEM, data )}
        .onEvent=${(...args) => this._onEditEvent(...args)}
        .onSelector=${(...args)=>this._onEditEvent(EDIT_EVENT.SELECTOR, [...args] )}
        .msg=${this.msg}
      ></form-row>`)}
    </div>`
  }
}