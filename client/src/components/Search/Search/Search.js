import { LitElement, html } from 'lit';
import { cache } from 'lit/directives/cache.js';

import '../../SideBar/Search/sidebar-search.js'
import '../../Modal/Selector/modal-selector.js'
import '../Browser/search-browser.js'

import '../../Modal/Total/modal-total.js'
import '../../Modal/Server/modal-server.js'

import { styles } from './Search.styles.js'
import { SIDE_VISIBILITY } from '../../../config/enums.js'

export class Search extends LitElement {
  constructor() {
    super();
    this.data = {}
    this.side = SIDE_VISIBILITY.AUTO
    this.auditFilter = {}
    this.queries = {}
    this.quick = {}
    this.paginationPage = 10
    this.onEvent = {}
    this.modalServer = this.modalServer.bind(this)
    this.modalTotal = this.modalTotal.bind(this)
  }

  static get properties() {
    return {
      data: { type: Object },
      side: { type: String },
      auditFilter: { type: Object },
      paginationPage: { type: Number },
      queries: { type: Object },
      quick: { type: Object },
      onEvent: { type: Object },
    };
  }

  static get styles () {
    return [
      styles
    ]
  }

  connectedCallback() {
    super.connectedCallback();
    this.onEvent.setModule(this)
  }

  modalServer({cmd, fields, values, onEvent}){
    return html`<modal-server 
      .cmd=${cmd} .fields=${fields} .values=${values} 
      .onEvent=${onEvent} .msg=${this.msg}
    ></modal-server>`
  }

  modalTotal(total){
    return html`<modal-total 
      .total=${total} .onEvent=${this.onEvent} .msg=${this.msg}
    ></modal-total>`
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
        ${cache((this.data.seltype === "browser") 
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
            .onEvent=${this.onEvent}
            .msg=${this.msg}
          ></modal-selector>`)}
      </div>`
  }
}