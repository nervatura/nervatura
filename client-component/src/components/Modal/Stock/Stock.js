import { LitElement, html } from 'lit';

import '../../Form/Label/form-label.js'
import '../../Form/Button/form-button.js'
import '../../Form/Icon/form-icon.js'
import '../../Form/Table/form-table.js'

import { styles } from './Stock.style.js'
import { MODAL_EVENT, BUTTON_TYPE, PAGINATION_TYPE } from '../../../config/enums.js'

export class Stock extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.partnumber = "" 
    this.partname = "" 
    this.rows = []
    this.selectorPage = 5
  }

  static get properties() {
    return {
      partnumber: { type: String }, 
      partname: { type: String }, 
      rows: { type: Array }, 
      selectorPage: { type: Number },
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
    const { partnumber, partname, rows, selectorPage } = this
    const fields = {
      warehouse: { fieldtype:"string", label: this.msg("", { id: "delivery_place" }) },
      batch_no: { fieldtype:"string", label: this.msg("", { id: "movement_batchnumber" }) },
      description: { fieldtype:"string", label: this.msg("", { id: "product_description" }) },
      sqty: { fieldtype:"number", label: this.msg("", { id: "shipping_stock" }) }
    }
    return html`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label leftIcon="Book"
                value="${this.msg("", { id: "shipping_stocks" })}" 
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
              <div class="cell padding-label" >
                <div>
                  <form-label value="${partnumber}" ></form-label>
                </div>
                <div>
                  <form-label .style=${{ "font-weight": "normal" }}
                    value="${partname}" 
                  ></form-label>
                </div>
              </div>
            </div>
            <div class="section-row" >
              <form-table id="selector_result"
                .rows="${rows}"
                .fields="${fields}"
                pagination="${PAGINATION_TYPE.TOP}"
                pageSize="${selectorPage}"
                ?tableFilter="${true}" 
                filterPlaceholder="${this.msg("", { id: "placeholder_filter" })}"
                ?hidePaginatonSize="${true}"
              ></form-table>
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