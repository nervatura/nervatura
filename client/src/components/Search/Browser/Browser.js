import { LitElement, html, nothing } from 'lit';

import '../../Form/Label/form-label.js'
import '../../Form/Icon/form-icon.js'
import '../../Form/Button/form-button.js'
import '../../Form/Input/form-input.js'
import '../../Form/NumberInput/form-number.js'
import '../../Form/Table/form-table.js'
import '../../Form/Select/form-select.js'
import '../../Form/DateTime/form-datetime.js'

import { styles } from './Browser.styles.js'
import { BROWSER_EVENT, BROWSER_FILTER, BUTTON_TYPE, DATETIME_TYPE, INPUT_TYPE, PAGINATION_TYPE, TEXT_ALIGN } from '../../../config/enums.js'

export class Browser extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.data = {
      vkey: "",
      view: "",
      show_header: true, 
      show_dropdown: false,
      show_columns: false,
      result: [],
      columns: {},
      filters: {},
      deffield: [], 
      page: 1,
    }
    this.keyMap = {}
    this.viewDef = {
      fields: {},
      label: "",
      readonly: false
    }
    this.paginationPage = 10
    this.onEvent = {}
  }

  static get properties() {
    return {
      data: { type: Object },
      keyMap: { type: Object },
      viewDef: { type: Object },
      paginationPage: { type: Number },
      onEvent: { type: Object }
    };
  }

  static get styles () {
    return [
      styles
    ]
  }

  connectedCallback() {
    super.connectedCallback();

    const { show_header, show_dropdown, show_columns } = this.data
    this.dropdown = show_dropdown
    this.header = show_header
    this.columns = show_columns
  }

  _onBrowserEvent(key, data){
    if(this.onEvent && this.onEvent.onBrowserEvent){
      this.onEvent.onBrowserEvent({ key, data })
    }
    this.dispatchEvent(
      new CustomEvent('browser_event', {
        bubbles: true, composed: true,
        detail: {
          key, data
        }
      })
    );
  }

  _onValueChange(key, value){
    this[key] = value
    this.dispatchEvent(
      new CustomEvent('value_change', {
        bubbles: true, composed: true,
        detail: {
          ...{ key, value }
        }
      })
    );
    this.requestUpdate()
  }

  exportFields() {
    const { view, columns } = this.data
    const fields = {}
    Object.keys(this.viewDef.fields).filter(fieldname => (columns[view][fieldname]===true)).forEach(fieldname => {
      fields[fieldname] = this.viewDef.fields[fieldname]
    })
    return fields
  }

  checkTotalFields() {
    const { fields } = this.viewDef
    const { deffield } = this.data
    let total = { totalFields: {}, totalLabels: {}, count: 0 }
    if (deffield && Object.keys(fields).includes("deffield_value")) {
      deffield.filter((df)=>((df.fieldtype==="integer")||(df.fieldtype==="float")))
      .forEach((df) => {
        total = {
          ...total,
          totalFields: {
            ...total.totalFields,
            [df.fieldname]: 0
          },
          totalLabels: {
            ...total.totalLabels,
            [df.fieldname]: df.description
          }
        }
      })
    } else {
      Object.keys(fields).filter((fieldname)=>(
        ((fields[fieldname].fieldtype==="integer")||(fields[fieldname].fieldtype==="float"))
          &&(fields[fieldname].calc !== "avg")))
      .forEach((fieldname) => {
        total = {
          ...total,
          totalFields: {
            ...total.totalFields,
            [fieldname]: 0
          },
          totalLabels: {
            ...total.totalLabels,
            [fieldname]: fields[fieldname].label
          }
        }
       }
      )
    }
    total = {
      ...total,
      count:  Object.keys(total.totalFields).length
    }
    return total
  }

  fields(){
    const { view, columns } = this.data
    let fields = {
      view: { columnDef: { 
        id: "view",
        Header: "",
        headerStyle: {},
        Cell: ({ row }) => {
          if(!this.viewDef.readonly){
            return html`<form-icon id=${`edit_${row.id}`}
              iconKey="Edit" width=24 height=21.3 
              @click=${ ()=>this._onBrowserEvent(BROWSER_EVENT.EDIT_CELL, { 
                fieldname: "id", value: row.id, row })}
              .style=${{ cursor: "pointer", fill: "rgb(var(--functional-green))" }} ></form-icon>`
          }
          return html`<form-icon iconKey="CaretRight" width=9 height=24 ></form-icon>`
        },
        cellStyle: { width: "25px", padding: "7px 3px 3px 8px" }
      }}
    }
    Object.keys(this.viewDef.fields).forEach((fieldname) => {
      if(columns[view][fieldname]){
        switch (this.viewDef.fields[fieldname].fieldtype) {
          case "float":
          case "integer":
            fields = {
              ...fields,
              [fieldname]: { fieldtype:'number', label: this.viewDef.fields[fieldname].label }
            }
            break;
  
          case "bool":
            fields = {
              ...fields,
              [fieldname]: { fieldtype:'bool', label: this.viewDef.fields[fieldname].label }
            }
            break;
  
          case "string":
          default:
            if(fieldname === "deffield_value"){
              fields = {
                ...fields,
                [fieldname]: { fieldtype:'deffield', label: this.viewDef.fields[fieldname].label }
              }
            } else {
              fields = {
                ...fields,
                [fieldname]: { fieldtype:'string', label: this.viewDef.fields[fieldname].label }
              }
            }
            break;
        }
      }
    })
    return fields
  }

  render() {
    const { 
      vkey, view, result, columns, filters, deffield, page 
    } = this.data

    
    const totalFields = this.checkTotalFields()
    return html`<div @click="${()=>(this.dropdown)?this._onValueChange("dropdown", false):null}">
      <div class="panel">
        <div class="panel-title">
          <div class="cell">
            <form-label 
              value="${this.msg(`browser_${vkey}`, { id: `browser_${vkey}` })}" 
              class="title-cell" >
            </form-label>
          </div>
        </div>
        <div class="panel-container" >
          <div class="row full" >
            <div class="cell" >
              <form-button id="btn_header" 
                icon="Filter" type="${BUTTON_TYPE.PRIMARY}" ?full="${true}"
                label="${this.viewDef.label}" align=${TEXT_ALIGN.LEFT}
                @click=${(event)=>{
                  event.stopPropagation();
                  this._onBrowserEvent(BROWSER_EVENT.CHANGE, { fieldname: "show_header", value: !this.header })
                  this._onValueChange("header", !this.header)
                }}
              >${this.viewDef.label}
              </form-button>
            </div>
          </div>
          ${(this.header) ? html`<div class="filter-panel" >
            <div class="row full" >
              <div class="cell" >
                <form-button id="btn_search" 
                  icon="Search" type="${BUTTON_TYPE.BORDER}"
                  label="${this.msg("", { id: "browser_search" })}"
                  .style=${{ "padding": "8px 12px" }} ?hidelabel=${true}
                  @click=${()=>this._onBrowserEvent(BROWSER_EVENT.BROWSER_VIEW, {})}
                >${this.msg("", { id: "browser_search" })}</form-button>
              </div>
              <div class="cell align-right" >
                <form-button id="btn_bookmark" 
                  icon="Star" type="${BUTTON_TYPE.BORDER}"
                  label="${this.msg("", { id: "browser_bookmark" })}"
                  .style=${{ "padding": "8px 12px" }} ?hidelabel=${true}
                  @click=${()=>this._onBrowserEvent(BROWSER_EVENT.BOOKMARK_SAVE, [])}
                >${this.msg("", { id: "browser_bookmark" })}</form-button>
                <form-button id="btn_export" 
                  icon="Download" type="${BUTTON_TYPE.BORDER}"
                  label="${this.msg("", { id: "browser_export" })}"
                  .style=${{ "padding": "8px 12px" }} ?hidelabel=${true}
                  @click=${()=>this._onBrowserEvent(BROWSER_EVENT.EXPORT_RESULT, { value: this.exportFields() })}
                >${this.msg("", { id: "browser_export" })}</form-button>
                <form-button id="btn_help" 
                  icon="QuestionCircle" type="${BUTTON_TYPE.BORDER}"
                  label="${this.msg("", { id: "browser_help" })}"
                  .style=${{ "padding": "8px 12px" }} ?hidelabel=${true}
                  @click=${()=>this._onBrowserEvent(BROWSER_EVENT.SHOW_HELP, { value: "program/browser" })}
                >${this.msg("", { id: "browser_help" })}</form-button>
              </div>
            </div>
            <div class="row full section-small-top" >
              <div class="cell" >
                <div class="dropdown-box" >
                  <form-button id="btn_views"
                    type="${BUTTON_TYPE.BORDER}"
                    .style=${{ "padding": "8px 12px" }} icon="Eye"
                    label="${this.msg("", { id: "browser_views" })}"
                    ?selected="${(this.dropdown)}" ?hidelabel=${true}
                    @click=${(event)=>{
                      event.stopPropagation();
                      this._onValueChange("dropdown", !this.dropdown)
                    }}
                  >${this.msg("", { id: "browser_views" })}</form-button>
                  ${(this.dropdown) ? html`<div class="dropdown-content" >
                    ${Object.keys(this.keyMap).map((vname) => (vname !== "options")?
                      html`<div id=${`view_${vname}`}
                        @click=${ ()=>this._onBrowserEvent(BROWSER_EVENT.SHOW_BROWSER, { value: vkey, vname }) }
                        class="drop-label" >
                        <form-label class="${(vname === view)?"active":""}"
                          value="${this.keyMap[vname].label}"
                          leftIcon="${(vname === view)?"Check":"Eye"}" ></form-label>
                      </div>`: nothing
                    )}
                  </div>` : nothing}
                </div>
                <form-button id="btn_columns" 
                  type="${BUTTON_TYPE.BORDER}"
                  .style=${{ "padding": "8px 12px" }} icon="Columns" ?hidelabel=${true}
                  label="${this.msg("", { id: "browser_columns" })}"
                  @click=${(event)=>{
                    event.stopPropagation();
                    this._onBrowserEvent(BROWSER_EVENT.CHANGE, { fieldname: "show_columns", value: !this.columns })
                    this._onValueChange("columns", !this.columns)
                  }}
                >${this.msg("", { id: "browser_columns" })}</form-button>
                <form-button id="btn_filter" 
                  icon="Plus" type="${BUTTON_TYPE.BORDER}" ?hidelabel=${true}
                  label="${this.msg("", { id: "browser_filter" })}"
                  .style=${{ "padding": "8px 12px" }}
                  @click=${()=>this._onBrowserEvent(BROWSER_EVENT.ADD_FILTER, {})}
                >${this.msg("", { id: "browser_filter" })}</form-button>
                <form-button id="btn_total" 
                  icon="InfoCircle" type="${BUTTON_TYPE.BORDER}"
                  label="${this.msg("", { id: "browser_total" })}"
                  .style=${{ "padding": "8px 12px" }} ?hidelabel=${true}
                  ?disabled=${!!(((totalFields.count === 0)||(result.length === 0)))}
                  @click=${()=>this._onBrowserEvent(BROWSER_EVENT.SHOW_TOTAL, {
                    fields: this.viewDef.fields, totalFields 
                  })}
                >${this.msg("", { id: "browser_total" })}</form-button>
              </div>
            </div>
            ${(this.columns) ? html`<div class="col-box">
              ${Object.keys(this.viewDef.fields).map(fieldname =>
                html`<div id=${`col_${fieldname}`} 
                  class="cell col-cell base-col ${
                    (columns[view][fieldname]===true) ? "select-col" : "edit-col"}"
                  @click=${()=>this._onBrowserEvent(BROWSER_EVENT.SET_COLUMNS, { fieldname, value: !(columns[view][fieldname]===true) })}
                  >
                  <form-label
                    value="${this.viewDef.fields[fieldname].label}"
                    leftIcon="${(columns[view][fieldname]===true)?"CheckSquare":"SquareEmpty"}" ></form-label>
                </div>`
              )}              
            </div>`:nothing}
            ${filters[view].map((filter, index) => html`<div class="section-small-top" >
              <div class="cell" >
                <form-select id=${`filter_name_${index}`}
                  label="${this.msg("", { id: "browser_filter" })}"
                  .onChange=${
                    (value) => this._onBrowserEvent(BROWSER_EVENT.EDIT_FILTER , { index, fieldname: "fieldname", value: value.value } )
                  }
                  .options=${Object.keys(this.viewDef.fields).filter(
                    (fieldname)=> (fieldname !== "id") && (fieldname !== "_id")
                    ).flatMap((fieldname) => {
                      if(fieldname === "deffield_value"){
                        return deffield.map((df) => ({ 
                          value: df.fieldname, 
                          text: this.msg(df.description, { id: df.fieldname }) 
                        }))
                      }
                      return { value: fieldname, text: this.viewDef.fields[fieldname].label } 
                    })
                  } 
                  .isnull="${false}" value="${filter.fieldname}" >
                </form-select>
              </div>
              <div class="cell" >
                <form-select id=${`filter_type_${index}`}
                  label="${this.msg("", { id: "browser_filter" })}"
                  .onChange=${
                    (value) => this._onBrowserEvent(BROWSER_EVENT.EDIT_FILTER , { index, fieldname: "filtertype", value: value.value } )
                  }
                  .options=${
                    (["date","float","integer"].includes(filter.fieldtype) 
                    ? BROWSER_FILTER : BROWSER_FILTER.slice(0,3))
                    .map((item)=>({ value: item[0], text: item[1] }))
                  } 
                  .isnull="${false}" value="${filter.filtertype}" >
                </form-select>
              </div>
              <div class="cell mobile" >
                ${(filter.filtertype !== "==N") ? html`<div class="cell" >
                  ${(filter.fieldtype === "bool") ? html`<form-select 
                    id=${`filter_value_${index}`}
                    label="${this.msg("", { id: "browser_filter" })}"
                    .onChange=${
                      (value) => this._onBrowserEvent(BROWSER_EVENT.EDIT_FILTER , { index, fieldname: "value", value: value.value } )
                    }
                    .options=${[
                        { value: "0", text: this.msg("NO", { id: "label_no" }) }, 
                        { value: "1", text: this.msg("YES", { id: "label_yes" }) }
                    ]} 
                    .isnull="${false}" value="${filter.value}" >
                  </form-select>` : nothing}
                  ${((filter.fieldtype === "integer")||(filter.fieldtype === "float")) ? html`<form-number 
                    id=${`filter_value_${index}`} 
                    label="${this.msg("", { id: "browser_filter" })}"
                    ?integer="${!(filter.fieldtype === "float")}"
                    value="${filter.value}"
                    .onChange=${
                      (value) => this._onBrowserEvent(BROWSER_EVENT.EDIT_FILTER , { 
                        index, fieldname: "value", value: value.value 
                      }) 
                    }
                  ></form-number>` : nothing}
                  ${(filter.fieldtype === "date") ? html`<form-datetime 
                    id=${`filter_value_${index}`}
                    label="${this.msg("", { id: "browser_filter" })}"
                    .isnull="${false}"
                    type="${DATETIME_TYPE.DATE}"
                    .value="${filter.value}"
                    .onChange="${
                      (date) => this._onBrowserEvent(BROWSER_EVENT.EDIT_FILTER , { 
                        index, fieldname: "value", value: date.value 
                      })
                    }"
                  ></form-datetime>` : nothing}
                  ${(filter.fieldtype === "string") ? html`<form-input 
                    id=${`filter_value_${index}`} 
                    label="${this.msg("", { id: "browser_filter" })}"
                    type="${INPUT_TYPE.TEXT}" 
                    value="${filter.value}"
                    .onChange=${
                      (event) => this._onBrowserEvent(BROWSER_EVENT.EDIT_FILTER , { 
                        index, fieldname: "value", value: event.value 
                      }) 
                    }
                  ></form-input>` : nothing}
                </div>`:nothing}
                <div class="cell" >
                  <form-button id="${`btn_delete_filter_${index}`}" 
                    icon="Times" type="${BUTTON_TYPE.BORDER}"
                    label="${this.msg("", { id: "label_delete" })}"
                    .style=${{ "padding": "8px", "border-radius": "3px" }}
                    @click=${()=>this._onBrowserEvent(BROWSER_EVENT.DELETE_FILTER, { value: index })}
                  ></form-button>
                </div>
              </div>
            </div>`)}
          </div>` : nothing}
          <div class="row full section-small-top" >
            <div class="row full border" >
              <div class="cell result-title" >
                ${result.length} <form-label 
                  value="${this.msg("record(s) found", { id: `browser_result` })}" 
                ></form-label>
              </div>
              ${(this.viewDef.actions_new) ? html`<div class="cell result-title-plus" >
                <form-button 
                id="btn_actions_new" 
                icon="Plus" ?small=${true}
                label="${this.msg("", { id: "label_new" })}"
                @click=${()=>this._onBrowserEvent(BROWSER_EVENT.SET_FORM_ACTION, { 
                  params: this.viewDef.actions_new, row: undefined })}
                ></form-button>
              </div>` : nothing}
            </div>
          </div>
          <div class="row full" >
          <form-table 
            id="browser_result" rowKey="row_id"
            .rows="${result}"
            .fields="${this.fields()}"
            filterPlaceholder="${this.msg("Filter", { id: "placeholder_filter" })}"
            labelYes="${this.msg("YES", { id: "label_yes" })}"
            labelNo="${this.msg("NO", { id: "label_no" })}"
            pagination="${PAGINATION_TYPE.TOP}"
            currentPage="${page}"
            pageSize="${this.paginationPage}"
            ?tableFilter="${true}"
            ?hidePaginatonSize="${false}"
            .onEditCell=${(fieldname, value, row)=>this._onBrowserEvent(BROWSER_EVENT.EDIT_CELL, { fieldname, value, row })}
            .onCurrentPage=${(value)=>this._onBrowserEvent(BROWSER_EVENT.CURRENT_PAGE, { value })}
          ></form-table>
          </div>
        </div>
      </div>
    </div>`
  }
}