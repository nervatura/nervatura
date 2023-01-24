import { LitElement, html, nothing } from 'lit';
import { ifDefined } from 'lit/directives/if-defined.js';

import '../../Form/Table/form-table.js'
import '../../Form/List/form-list.js'
import '../../Form/Icon/form-icon.js'

import { styles } from './View.styles.js'
import { EDIT_EVENT, PAGINATION_TYPE, ACTION_EVENT } from '../../../config/enums.js'

export class View extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.id = Math.random().toString(36).slice(2);
    this.viewName = ""
    this.audit = ""
    this.current = {}
    this.template = {} 
    this.dataset = {}
    this.pageSize = 10
    this.onEvent = {}
  }

  static get properties() {
    return {
      id: { type: String },
      viewName: { type: String },
      audit: { type: String },
      current: { type: Object },
      template: { type: Object },
      dataset: { type: Object },
      pageSize: { type: Number },
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
    const { viewName, audit, template, dataset, current, pageSize } = this
    const vtemplate = template.view[viewName]
    const rows = dataset[vtemplate.data] || [];
    const edited = (typeof vtemplate.edited !== "undefined") ? vtemplate.edited : true;
    let actions = vtemplate.actions;
    if (typeof actions === "undefined") {
      actions = {
        new: { action: ACTION_EVENT.NEW_EDITOR_ITEM, fkey: viewName }, 
        edit: { action: ACTION_EVENT.EDIT_EDITOR_ITEM, fkey: viewName }, 
        delete: { action: ACTION_EVENT.DELETE_EDITOR_ITEM, fkey: viewName }
      }
    }
    if (audit !== "all") {
      actions = {...actions,
        new: null,
        delete: null
      }
    }

    const editIcon = (typeof vtemplate.edit_icon !== "undefined") ? 
      [vtemplate.edit_icon, undefined, undefined] : ["Edit", 24, 21.3]
    const deleteIcon = (typeof vtemplate.delete_icon !== "undefined") ? 
      [vtemplate.delete_icon, undefined, undefined] : ["Times", 19, 27.6]
    const addIcon = (typeof vtemplate.new_icon !== "undefined") ? 
      vtemplate.new_icon : "Plus"
    const labelAdd = (typeof vtemplate.new_label !== "undefined") ? 
      vtemplate.new_label : this.msg("", { id: "label_new" })

    let fields = {}
    if(vtemplate.type === "table"){
      if(edited && (actions.edit || actions.delete)){
        fields = {...fields,
          edit: { columnDef: { 
            id: "edit",
            Header: "",
            headerStyle: {},
            Cell: ({ row }) => {
              const ecol = (actions.edit !== null) ? html`<form-icon id=${`edit_${row.id}`}
                  iconKey=${editIcon[0]} width=${editIcon[1]} height=${editIcon[2]}
                  .style=${{ cursor: "pointer", fill: "rgb(var(--functional-green))" }}
                  @click=${ (event)=>{
                    event.stopPropagation();
                    this._onEditEvent(EDIT_EVENT.FORM_ACTION , { params: actions.edit, row, ref: this })
                  }}
                ></form-icon>` : undefined
              const dcol = (actions.delete !== null) ? html`<form-icon id=${`delete_${row.id}`}
                  iconKey=${deleteIcon[0]} width=${deleteIcon[1]} height=${deleteIcon[2]}
                  .style=${{ cursor: "pointer", fill: "rgb(var(--functional-red))", 
                    "margin-left": (actions.edit !== null) ? "8px" : "0" }}
                  @click=${ (event)=>{
                    event.stopPropagation();
                    this._onEditEvent(EDIT_EVENT.FORM_ACTION , { params: actions.delete, row })
                  }}
                ></form-icon>` : undefined
              return html`${ifDefined(ecol)}${ifDefined(dcol)}`
            },
            cellStyle: { width: 30, padding: "4px 3px 3px 8px" }
          }}
        }
      }
      fields = { ...fields, ...vtemplate.fields }
    }
    return html`<div id="${this.id}" class="panel" >
      ${(vtemplate.total) ? html`<div class="container-row">
        <div class="total-cell">
          <span class="total-label" >${`${vtemplate.total[Object.keys(vtemplate.total)[0]]}:`}</span>
          <span class="total-value" >${new Intl.NumberFormat('default').format(dataset[current.type][0][Object.keys(vtemplate.total)[0]])}</span>
        </div>
        <div class="total-cell">
          <span class="total-label" >${`${vtemplate.total[Object.keys(vtemplate.total)[1]]}:`}</span>
          <span class="total-value" >${new Intl.NumberFormat('default').format(dataset[current.type][0][Object.keys(vtemplate.total)[1]])}</span>
        </div>
        <div class="total-cell">
          <span class="total-label" >${`${vtemplate.total[Object.keys(vtemplate.total)[2]]}:`}</span>
          <span class="total-value" >${new Intl.NumberFormat('default').format(dataset[current.type][0][Object.keys(vtemplate.total)[2]])}</span>
        </div>
      </div>` : nothing }
      <div class="row full" >
        ${(vtemplate.type === "table") ? html`<form-table id="view_table"
          .onAddItem=${(edited && actions.new) 
            ? () => this._onEditEvent(EDIT_EVENT.FORM_ACTION , { params: actions.new } ) : undefined }
          .fields=${fields} .rows=${rows} ?tableFilter=${true}
          filterPlaceholder="${this.msg("", { id: "placeholder_filter" })}"
          labelYes=${this.msg("", { id: "label_yes" })} labelNo=${this.msg("", { id: "label_no" })} 
          labelAdd=${labelAdd} addIcon=${addIcon} 
          pageSize=${pageSize} pagination="${PAGINATION_TYPE.TOP}"
        ></form-table>` : html`<form-list id="view_list"
          .rows=${rows} labelAdd=${labelAdd} addIcon=${addIcon}
          editIcon=${editIcon[0]} deleteIcon=${deleteIcon[0]} ?listFilter=${true} 
          filterPlaceholder=${this.msg("", { id: "placeholder_filter" })}
          pageSize=${pageSize} pagination="${PAGINATION_TYPE.TOP}"
          .onEdit=${(edited && actions.edit) 
            ? (row) => this._onEditEvent(EDIT_EVENT.FORM_ACTION , { params: actions.edit, row } ) : undefined}
          .onDelete=${(edited && actions.delete) 
            ? (row) => this._onEditEvent(EDIT_EVENT.FORM_ACTION , { params: actions.delete, row } ) : undefined}
          .onAddItem=${(edited && actions.new) 
            ? () => this._onEditEvent(EDIT_EVENT.FORM_ACTION , { params: actions.new } ) : undefined}
        ></form-list>`}
      </div>
    </div>`
  }
}