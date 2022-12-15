import { LitElement, html, nothing } from 'lit';

import '../../Form/Label/form-label.js'
import '../../Form/Icon/form-icon.js'
import '../../Form/Button/form-button.js'
import '../../Form/Input/form-input.js'
import '../../Form/Table/form-table.js'

import { styles } from './Selector.styles.js'
import { MODAL_EVENT, PAGINATION_TYPE, INPUT_TYPE } from '../../../config/enums.js'

export class Selector extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.isModal = false
    this.view = ""
    this.columns = []
    this.result = []
    this.filter = ""
    this.selectorPage = 5
    this.paginationPage = 10
    this.currentPage = 1
  }

  static get properties() {
    return {
      isModal: { type: Boolean },
      view: { type: String },
      columns: { type: Array },
      result: { type: Array },
      filter: { type: String },
      selectorPage: { type: Number },
      paginationPage: { type: Number },
      currentPage: { type: Number },
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

  _onFilterChange(data){
    this.filter = data.value
    this.dispatchEvent(
      new CustomEvent('filter_change', {
        bubbles: true, composed: true,
        detail: {
          ...data
        }
      })
    );
  }

  selectorView(){
    let fields = {
      view: { columnDef: { 
        id: "view",
        Header: "",
        headerStyle: {},
        Cell: ({ row }) => {
          if(row.deleted === 1)
            return html`<form-icon iconKey="ExclamationTriangle" .style=${{ "fill": "rgb(var(--functional-yellow))" }} ></form-icon>`
          return html`<form-icon iconKey="CaretRight" width=9 height=24 ></form-icon>`
        },
        cellStyle: { width: "25px", padding: "7px 2px 3px 8px" }
      }}
    }
    this.columns.forEach(field => {
      fields = {
        ...fields,
        [field[0]]: {
          fieldtype:'string', 
          label: this.msg(`${this.view}_${field[0]}`, { id: `${this.view}_${field[0]}` })
        }
      }
    });
    return html`<div class="panel ${(!this.isModal) ? "margin0" : ""}">
      <div class="panel-title">
        <div class="cell" >
          <form-label 
            value=${(this.isModal) 
              ? this.msg(`search_${this.view}`, { id: `search_${this.view}` })
              : `${this.msg("Quick Search", { id: "quick_search" })}: ${this.msg(`search_${this.view}`, { id: `search_${this.view}` })}`}
            class="title-cell" leftIcon=${(this.isModal) ? "Search" : ""} >
          </form-label>
        </div>
        ${(this.isModal) ? html`<div class="cell align-right" >
          <span id=${`closeIcon`} class="close-icon" 
            @click="${ ()=>this._onModalEvent(MODAL_EVENT.CANCEL, {}) }">
            <form-icon iconKey="Times" ></form-icon>
          </span>
        </div>` : nothing }
      </div>
      <div class="section" >
        <div class="section-row" >
          <div class="cell" >
            <form-input id="selector_filter" type="${INPUT_TYPE.TEXT}" 
              label="${this.msg("", { id: "placeholder_search" })}"
              placeholder="${this.msg("", { id: "placeholder_search" })}"
              value="${this.filter}" ?full="${true}" ?autofocus="${true}"
              .onChange=${
                (event) => this._onFilterChange({ value: event.value, old: this.filter }) 
              }
              .onEnter=${()=>this._onModalEvent(MODAL_EVENT.SEARCH, { value: this.filter })}
            ></form-input>
          </div>
          <div class="cell search-col" >
            <form-button id="selector_btn_search" icon="Search"
              label="${this.msg("", { id: "label_search" })}"
              @click=${()=>this._onModalEvent(MODAL_EVENT.SEARCH, { value: this.filter })}
            >${this.msg("", { id: "label_search" })}
            </form-button>
          </div>
        </div>
        <div class="section-row" >
        <form-table id="selector_result"
          .rows="${this.result}"
          .fields="${fields}"
          pagination="${PAGINATION_TYPE.TOP}"
          currentPage="${this.currentPage}"
          pageSize="${(this.isModal) ? this.selectorPage : this.paginationPage}"
          ?tableFilter="${false}"
          ?hidePaginatonSize="${true}"
          .onRowSelected=${(row)=>this._onModalEvent(MODAL_EVENT.SELECTED, { value: row, filter: this.filter })}
          .onCurrentPage=${(value)=>this._onModalEvent(MODAL_EVENT.CURRENT_PAGE, { value })}
        ></form-table>
        </div>
      </div>
    </div>`
  }

  render() {
    if(this.isModal){
      return html`<div class="modal">
        <div class="dialog">
          ${this.selectorView()}
        </div>
      </div>`
    }
    return this.selectorView()
  }
}