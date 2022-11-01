import { LitElement, html } from 'lit';
// import { msg } from '@lit/localize';

import '../../Form/Label/form-label.js'
import '../../Form/Icon/form-icon.js'
import '../../Form/Button/form-button.js'
import '../../Form/List/form-list.js'
import '../InputBox/modal-inputbox.js'

import { styles } from './Bookmark.styles.js'
import { MODAL_EVENT, PAGINATION_TYPE, BUTTON_TYPE, BOOKMARK_VIEW } from '../../../config/enums.js'

export class Bookmark extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.bookmark = { history: null, bookmark: [] }
    this.tabView = BOOKMARK_VIEW.BOOKMARK
    this.pageSize = 5;
    this.values = {}
  }

  static get properties() {
    return {
      bookmark: { type: Object },
      tabView: { type: String },
      pageSize: { type: Number },
      values: { type: Object },
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

  _onTabView(view) {
    this.tabView = view
  }

  setBookmark() {
    return this.bookmark.bookmark.map(item => {
      const bvalue = JSON.parse(item.cfvalue)
      const value = {
        bookmark_id: item.id,
        id: bvalue.id,
        cfgroup: item.cfgroup,
        ntype: bvalue.ntype,
        transtype: (bvalue.ntype === "trans") ? bvalue.transtype : null,
        vkey: bvalue.vkey,
        view: bvalue.view,
        filters: bvalue.filters,
        columns: bvalue.columns,
        lslabel: item.cfname, 
        lsvalue: new Intl.DateTimeFormat(
          'default', {
            year: 'numeric', month: '2-digit', day: '2-digit'
          }
        ).format(new Date(bvalue.date))
      }
      if(item.cfgroup === "editor"){
        if (bvalue.ntype==="trans") {
          value.lsvalue += ` | ${  this.msg(`title_${bvalue.transtype}`, { id: `title_${bvalue.transtype}` })  } | ${  bvalue.info}`
        } else {
          value.lsvalue += ` | ${  this.msg(`title_${bvalue.ntype}`, { id: `title_${bvalue.ntype}` })  } | ${  bvalue.info}`
        }
      }
      if(item.cfgroup === "browser"){
        value.lsvalue += ` | ${  this.msg(`browser_${bvalue.vkey}`, { id: `browser_${bvalue.vkey}` })}`
      }
      return value
    })
  }

  setHistory() {
    if(this.bookmark.history && this.bookmark.history.cfvalue){
      const history_values = JSON.parse(this.bookmark.history.cfvalue)
      return history_values.map(item => ({
        id: item.id, 
        lslabel: item.title, 
        type: item.type,
        lsvalue: `${new Intl.DateTimeFormat(
          'default', {
            year: 'numeric', month: '2-digit', day: '2-digit', 
            hour: '2-digit', minute: '2-digit', hour12: false
          }
        ).format(new Date(item.datetime))} | ${ this.msg(`label_${item.type}`, { id: `label_${item.type}` })}`, 
        ntype: item.ntype, transtype: item.transtype
      }))
    }
    return []
  }

  connectedCallback() {
    super.connectedCallback();
    this.bookmarkList = this.setBookmark()
    this.historyList = this.setHistory()
  }

  render() {
    return html`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label 
                value="${this.msg("Nervatura Bookmark", { id: "title_bookmark" })}" 
                class="title-cell" leftIcon="Star" >
              </form-label>
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
              <div class="cell half" >
                <form-button id="btn_bookmark"
                  .style="${{ "border-radius": 0 }}" icon="Star"
                  label="${this.msg("", { id: "title_bookmark_list" })}"
                  @click=${()=>this._onTabView(BOOKMARK_VIEW.BOOKMARK)} 
                  type="${(this.tabView === BOOKMARK_VIEW.BOOKMARK) ? BUTTON_TYPE.PRIMARY : ""}"
                  ?full="${true}" ?selected="${(BOOKMARK_VIEW.BOOKMARK === this.tabView)}" >
                  ${this.msg("", { id: "title_bookmark_list" })}</form-button>
              </div>
              <div class="cell half" >
                <form-button id="btn_history"
                  .style="${{ "border-radius": 0 }}" icon="History"
                  label="${this.msg("", { id: "title_history" })}"
                  @click=${()=>this._onTabView(BOOKMARK_VIEW.HISTORY)} 
                  type="${(this.tabView === BOOKMARK_VIEW.HISTORY) ? BUTTON_TYPE.PRIMARY : ""}" 
                  ?full="${true}" ?selected="${(BOOKMARK_VIEW.HISTORY === this.tabView)}" >
                  ${this.msg("", { id: "title_history" })}</form-button>
              </div>
            </div>
            <div class="section-row" >
              <form-list id="bookmark_list"
                .rows="${(this.tabView === "bookmark") ? this.bookmarkList : this.historyList}"
                pagination="${PAGINATION_TYPE.TOP}"
                pageSize="${this.pageSize}"
                ?listFilter="${true}"
                ?hidePaginatonSize="${true}"
                filterPlaceholder="${this.msg("Filter", { id: "placeholder_filter" })}"
                editIcon="${(this.tabView === "bookmark") ? "Star" : "History"}"
                .onEdit=${(row)=>this._onModalEvent(MODAL_EVENT.SELECTED, { view: this.tabView, row })}
                .onDelete=${
                  (this.tabView === "bookmark") ? (row)=>this._onModalEvent(MODAL_EVENT.DELETE, { bookmark_id: row.bookmark_id, menubar: this.values.menubar }) : null 
                }
              ></form-list>
            </div>
          </div>
        </div>
      </div>
    </div>`
  }
}