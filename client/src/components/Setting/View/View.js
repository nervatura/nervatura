import { LitElement, html } from 'lit';

import '../../Form/Label/form-label.js'
import '../../Form/Table/form-table.js'
import '../../Form/List/form-list.js'
import '../../Form/Icon/form-icon.js'

import { styles } from './View.styles.js'
import { SETTING_EVENT, PAGINATION_TYPE } from '../../../config/enums.js'

export class View extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.id = Math.random().toString(36).slice(2);
    this.data = {
      caption: "",
      icon: "",
      view: {},
      actions: {},
    }
    this.paginationPage = 10
    this.onEvent = {}
  }

  static get properties() {
    return {
      id: { type: String },
      data: { type: Object },
      paginationPage: { type: Number },
      onEvent: { type: Object }
    };
  }

  static get styles () {
    return [
      styles
    ]
  }

  _onSettingEvent(key, data){
    if(this.onEvent.onSettingEvent){
      this.onEvent.onSettingEvent({ key, data })
    }
    this.dispatchEvent(
      new CustomEvent('setting_event', {
        bubbles: true, composed: true,
        detail: {
          key, data
        }
      })
    );
  }

  render() {
    const { caption, icon, view, actions, page } = this.data
    let fields = {}
    if(view.type === "table"){
      if(actions.edit){
        fields = {...fields,
          edit: { columnDef: { 
            id: "edit",
            Header: "",
            headerStyle: {},
            Cell: ({ row }) => html`<form-icon id=${`edit_${row.id}`}
                iconKey="Edit" width=24 height=21.3
                .style=${{ cursor: "pointer", fill: "rgb(var(--functional-green))" }}
                @click=${ (event)=>{
                  event.stopPropagation();
                  this._onSettingEvent(SETTING_EVENT.FORM_ACTION , { params: actions.edit, row, ref: this })
                }}
              ></form-icon>`,
            cellStyle: { width: 30, padding: "4px 3px 3px 8px" }
          }}
        }
      }
      fields = { ...fields, ...view.fields }
    }
    return html`<div class="panel" >
      <div class="panel-title">
        <div class="cell">
          <form-label class="title-cell"
            value="${caption}" leftIcon="${icon}"
          ></form-label>
        </div>
      </div>
      <div class="section full">
        <div class="section-row">
          ${(view.type === "table") ? html`<form-table id="view_table"
            .fields=${fields} .rows=${view.result} ?tableFilter=${true}
            filterPlaceholder="${this.msg("", { id: "placeholder_filter" })}"
            .onAddItem=${(actions.new) 
              ? () => this._onSettingEvent(SETTING_EVENT.FORM_ACTION, { params: actions.new } ) : undefined }
            .onRowSelected=${(actions.edit) 
              ? (row) => this._onSettingEvent(SETTING_EVENT.FORM_ACTION, { params: actions.edit, row }) : undefined }
            labelYes=${this.msg("", { id: "label_yes" })} 
            labelNo=${this.msg("", { id: "label_no" })} 
            labelAdd=${this.msg("", { id: "label_new" })}  
            pageSize=${this.paginationPage} pagination="${PAGINATION_TYPE.TOP}"
            currentPage="${page}"
            .onCurrentPage=${(value)=>this._onSettingEvent(SETTING_EVENT.CURRENT_PAGE, { value })}
          ></form-table>` : html`<form-list id="view_list"
            .rows=${view.result} ?listFilter=${true}
            filterPlaceholder=${this.msg("", { id: "placeholder_filter" })}
            .onAddItem=${(actions.new) 
              ? () => this._onSettingEvent(SETTING_EVENT.FORM_ACTION , { params: actions.new } ) : undefined}
            labelAdd=${this.msg("", { id: "label_new" })}
            pageSize=${this.paginationPage} pagination="${PAGINATION_TYPE.TOP}" 
            .onEdit=${(actions.edit) 
              ? (row) => this._onSettingEvent(SETTING_EVENT.FORM_ACTION , { params: actions.edit, row } ) : undefined}
            .onDelete=${(actions.delete) 
              ? (row) => this._onSettingEvent(SETTING_EVENT.FORM_ACTION , { params: actions.delete, row } ) : undefined}
            currentPage="${page}"
            .onCurrentPage=${(value)=>this._onSettingEvent(SETTING_EVENT.CURRENT_PAGE, { value })}
          ></form-list>`}
        </div>
      </div>
    </div>`
  }
}