import { LitElement, html } from 'lit';
// import { msg } from '@lit/localize';

import '../SideBar/Search/sidebar-search.js'
import '../Modal/Selector/modal-selector.js'
import '../Browser/search-browser.js'

import { styles } from './Search.styles.js'
import { SIDE_VISIBILITY } from '../../config/enums.js'

export class Search extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.data = {}
    this.side = SIDE_VISIBILITY.AUTO
    this.queries = {}
    this.quick = {}
    this.auditFilter = {}
    this.paginationPage = 10
    this.onEvent = {}
  }

  static get properties() {
    return {
      data: { type: Object },
      side: { type: String },
      queries: { type: Object },
      quick: { type: Object },
      auditFilter: { type: Object },
      paginationPage: { type: Number },
      onEvent: { type: Object },
    };
  }

  static get styles () {
    return [
      styles
    ]
  }

  modalServer({cmd, fields, values, onEvent}){
    return html`<modal-server 
      .cmd=${cmd} .fields=${fields} .values=${values} 
      .onEvent=${onEvent} .msg=${this.msg}
    ></modal-server>`
  }

  _onModalEvent(data){
    if(this.onEvent && this.onEvent.onModalEvent){
      this.onEvent.onModalEvent({ ...data, ref: this })
    }
  }

  render() {
    return html`<sidebar-search
      id="sidebar"
      side="${this.side}"
      groupKey="${this.data.group_key}"
      .auditFilter="${this.auditFilter}"
      .onEvent=${this.onEvent}
      .msg=${this.msg}
    ></sidebar-search>
      <div class="page">
        ${(this.data.seltype === "browser") 
        ? html`<search-browser
            id="${`browser_${this.data.vkey}`}"
            .data=${this.data}
            .keyMap=${this.queries[this.data.vkey]()}
            .viewDef="${this.queries[this.data.vkey]()[this.data.view]}"
            paginationPage=${this.paginationPage}
            .onEvent=${this.onEvent}
            .msg=${this.msg}
          ></search-browser>` 
        : html`<modal-selector
            id="${`selector_${this.data.qview}`}"
            ?isModal="${false}"
            view="${this.data.qview}"
            .columns=${this.quick[this.data.qview]().columns}
            .result=${this.data.result}
            filter="${this.data.qfilter}"
            paginationPage=${this.paginationPage}
            currentPage=${this.data.page}
            .onEvent=${{ onModalEvent: (...args)=>this._onModalEvent(...args) }}
            .msg=${this.msg}
          ></modal-selector>`}
      </div>`
  }
}