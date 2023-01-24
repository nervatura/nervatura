import { LitElement, html, nothing } from 'lit';
import { ifDefined } from 'lit/directives/if-defined.js';

import '../../Form/Label/form-label.js'
import '../../Form/Table/form-table.js'
import '../../Form/Row/form-row.js'
import '../../Form/Icon/form-icon.js'

import { styles } from './Form.styles.js'
import { SETTING_EVENT, PAGINATION_TYPE } from '../../../config/enums.js'

export class Form extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.id = Math.random().toString(36).slice(2);
    this.data = {
      caption: "",
      icon: "",
      current: {},
      audit: "", 
      dataset: {},
      type: "", 
      view: {},
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
    const { caption, icon, current, audit, dataset, type, view } = this.data
    let fields = {}
    if((typeof current.template.view.items !== "undefined") && (current.form.id !== null)){
      if(current.template.view.items.actions.edit){
        fields = {...fields,
          edit: { columnDef: { 
            id: "edit",
            Header: "",
            headerStyle: {},
            Cell: ({ row }) => {
              const ecol = html`<form-icon id=${`edit_${row.id}`}
                iconKey="Edit" width=24 height=21.3
                .style=${{ cursor: "pointer", fill: "rgb(var(--functional-green))" }}
                @click=${ (event)=>{
                  event.stopPropagation();
                  this._onSettingEvent(SETTING_EVENT.FORM_ACTION , { params: current.template.view.items.actions.edit, row, ref: this })
                }}
              ></form-icon>`
              const dcol = (current.template.view.items.actions.delete) ? html`<form-icon id=${`delete_${row.id}`}
                  iconKey="Times" width=19 height=27.6
                  .style=${{ cursor: "pointer", fill: "rgb(var(--functional-red))", 
                    "margin-left": "8px" }}
                  @click=${ (event)=>{
                    event.stopPropagation();
                    this._onSettingEvent(SETTING_EVENT.FORM_ACTION , { params: current.template.view.items.actions.delete, row })
                  }}
                ></form-icon>` : undefined
              return html`${ecol}${ifDefined(dcol)}`
            },
            cellStyle: { width: 30, padding: "4px 3px 3px 8px" }
          }}
        }
      }
      fields = { ...fields, ...current.template.view.items.fields }
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
          <div id="${this.id}" class="form-panel" >
            ${current.template.rows.map((row, index) => html`<form-row
              id=${`row_${index}`}
              .row="${row}"
              .values="${current.fieldvalue || current.form}"
              .options="${current.template.options}"
              .data="${{ audit, current, dataset }}"
              .onEdit=${(data)=>this._onSettingEvent(SETTING_EVENT.EDIT_ITEM, data )}
              .msg=${this.msg}
            ></form-row>`)}
          </div>
          ${(((typeof current.template.view.items !== "undefined") && (current.form.id !== null))
            || (type === "log")) ? html`<form-table 
            id="form_view" rowKey="id"
            .rows="${(type === "log") ? view.result : dataset[current.template.view.items.data]}"
            .fields="${(type === "log") ? view.fields : fields}"
            filterPlaceholder="${this.msg("Filter", { id: "placeholder_filter" })}"
            labelYes="${this.msg("YES", { id: "label_yes" })}"
            labelNo="${this.msg("NO", { id: "label_no" })}"
            pagination="${PAGINATION_TYPE.TOP}"
            pageSize="${this.paginationPage}"
            ?tableFilter="${true}"
            ?hidePaginatonSize="${false}"
            .onAddItem=${(current.template.view.items && current.template.view.items.actions.new) 
              ? () => this._onSettingEvent(SETTING_EVENT.FORM_ACTION , { params: current.template.view.items.actions.new } ) : undefined}
            labelAdd=${this.msg("", { id: "label_new" })}
          ></form-table>` : nothing }
        </div>
      </div>
    </div>`
  }
}