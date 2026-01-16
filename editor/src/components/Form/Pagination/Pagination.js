import { LitElement, html } from 'lit';
import { styleMap } from 'lit/directives/style-map.js';
import { ifDefined } from 'lit/directives/if-defined.js';

import { styles } from './Pagination.styles.js'
import { BUTTON_TYPE } from '../../../config/enums.js'

import '../NumberInput/form-number.js'
import '../Button/form-button.js'
import '../Select/form-select.js'

export class Pagination extends LitElement {
  constructor() {
    super();
    this.id = Math.random().toString(36).slice(2);
    this.name = undefined;
    this.pageIndex = 0; 
    this.pageSize = 5; 
    this.pageCount = 0;  
    this.canPreviousPage = false;
    this.canNextPage = false;
    this.hidePageSize = false;
    this.style = {};
  }

  static get properties() {
    return {
      id: { type: String },
      name: { type: String },
      pageIndex: { type: Number },
      pageSize: { type: Number },
      pageCount: { type: Number },
      canPreviousPage: { type: Boolean },
      canNextPage: { type: Boolean },
      hidePageSize: { type: Boolean },
      style: { type: Object },
    };
  }

  static get styles () {
    return [
      styles
    ]
  }

  _onEvent(key, value, enabled) {
    if(enabled){
      if(this.onEvent){
        this.onEvent(key, value);
      }
      this.dispatchEvent(new CustomEvent('pagination', {
        bubbles: true, composed: true,
        detail: {
          key, value
        }
      }));
    }
  }

  render() {
    return html`<div id="${this.id}" name="${ifDefined(this.name)}"
      class="row" style="${styleMap(this.style)}"
    >
      <div class="cell padding-small" >
        <form-button id="pagination_btn_first" 
          .style="${{ padding: "6px 6px 7px", "font-size": "15px", margin: "1px 0 2px" }}"
          ?disabled="${!this.canPreviousPage}" label="1"
          @click=${()=>this._onEvent("gotoPage", 1, this.canPreviousPage)} type="${BUTTON_TYPE.BORDER}" >1</form-button>
        <form-button id="pagination_btn_previous" 
          .style="${{ padding: "5px 6px 8px", "font-size": "15px", margin: "1px 0 2px" }}"
          ?disabled="${!this.canPreviousPage}" label="&#10094;"
          @click=${()=>this._onEvent("previousPage", this.pageIndex-1, this.canPreviousPage)} type="${BUTTON_TYPE.BORDER}" >&#10094;</form-button>
      </div>
      <div class="cell" >
        <form-number id="pagination_input_goto" ?integer="${true}" 
          .style="${{ "padding": "7px", width: "60px", "font-weight": "bold" }}"
          value="${this.pageIndex}" ?disabled="${this.pageCount === 0}"
          min="1" max="${this.pageCount}" label="Page"
          .onChange=${
            (value) => this._onEvent("gotoPage", value.value, (this.pageCount > 0)) 
          }
        ></form-number>
      </div>
      <div class="cell padding-small" >
        <form-button id="pagination_btn_next" 
          .style="${{ padding: "5px 6px 8px", "font-size": "15px", margin: "1px 0 2px" }}"
          ?disabled="${!this.canNextPage}" label="&#10095;"
          @click=${()=>this._onEvent("nextPage", this.pageIndex+1, this.canNextPage)} type="${BUTTON_TYPE.BORDER}" >&#10095;</form-button>
        <form-button id="pagination_btn_last" 
          .style="${{ padding: "6px 6px 7px", "font-size": "15px", margin: "1px 0 2px" }}"
          ?disabled="${!this.canNextPage}" label="${this.pageCount}"
          @click=${()=>this._onEvent("gotoPage", this.pageCount, this.canNextPage)} type="${BUTTON_TYPE.BORDER}" >${this.pageCount}</form-button>
      </div>
      ${(!this.hidePageSize)?html`<div class="cell padding-small hide-small" >
        <form-select id="pagination_page_size"
          label="Size"
          .style="${{ padding: "7px" }}" ?disabled="${(this.pageCount === 0)}"
          .onChange=${
            (value) => this._onEvent("setPageSize", Number(value.value), (this.pageCount > 0) )
          }
          .options=${[5, 10, 20, 50, 100].map((pageSize) => ({ value: String(pageSize), text: String(pageSize) }))} 
          .isnull="${false}" value="${this.pageSize}" >
        </form-select>
      </div>`:``}
    </div>`
  }
}