import { LitElement, html, nothing } from 'lit';

import '../../Form/Row/form-row.js'
import '../../Form/Button/form-button.js'
import '../../Form/Select/form-select.js'
import '../../Form/Pagination/form-pagination.js'

import { styles } from './Meta.styles.js'
import { EDIT_EVENT, BUTTON_TYPE, EDITOR_EVENT } from '../../../config/enums.js'

export class Meta extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.id = Math.random().toString(36).slice(2);
    this.audit = ""
    this.current = {} 
    this.dataset = {}
    this.pageSize = 5
    this.onEvent = {}
  }

  static get properties() {
    return {
      id: { type: String },
      audit: { type: String },
      current: { type: Object },
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

  connectedCallback() {
    super.connectedCallback();

    this.currentPage = this.current.page || 1
    if(this.currentPage < 1) {
      this.currentPage = 1
    }
  }

  _onPagination(key, value) {
    if(key === "setPageSize"){
      this.currentPage = 1
      this.pageSize = value
      return
    }
    this.currentPage = value
    this.requestUpdate()
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

  prepareRows() {
    const { current, dataset } = this
    let fieldvalue_list = []
    current.fieldvalue.forEach(fieldvalue => {
      const _deffield = dataset.deffield.filter((df) => (df.fieldname === fieldvalue.fieldname))[0]
      if ((_deffield.visible === 1) && (fieldvalue.deleted === 0)) {
        const _fieldtype = dataset.groups.filter((group) => (group.id === _deffield.fieldtype ))[0].groupvalue
        let _description = fieldvalue.value;
        let _datatype = _fieldtype;
        if(["customer", "tool", "trans", "transitem", "transmovement", "transpayment", 
          "product", "project", "employee", "place"].includes(_fieldtype)){
            const item = dataset.deffield_prop.filter((df) => (
              (df.ftype === _fieldtype) && (df.id === parseInt(fieldvalue.value,10))))[0]
            if(item){
              _description = item.description;}
            _datatype = "selector";
        }
        if(_fieldtype === "urlink"){
          _datatype = "text";
        }
        if(_fieldtype === "valuelist"){
          _description = _deffield.valuelist.split("|");
        }
        fieldvalue_list = [...fieldvalue_list, { 
          rowtype: 'fieldvalue',
          id: fieldvalue.id, name: 'fieldvalue_value', 
          fieldname: fieldvalue.fieldname, 
          value: fieldvalue.value, notes: fieldvalue.notes||'',
          label: _deffield.description, description: _description, 
          disabled: !!_deffield.readonly,
          fieldtype: _fieldtype, datatype: _datatype
        }]
      }
    })
    return fieldvalue_list
  }

  renderRows(rows, pageCount) {
    const { audit, current, dataset } = this
    let pageRows = rows
    if(pageCount > 1){
      const currentPage = (this.currentPage > pageCount) ? pageCount : this.currentPage
      const start = (currentPage-1)*this.pageSize;
      const end = currentPage*this.pageSize;
      // pageRows = pageRows.filter((row, index) => (index >= start && index < end))
      pageRows = pageRows.slice(start, end)
    }
    return pageRows
      .map((row, index)=> html`<form-row
        id=${`row_${index}`}
        .row="${row}"
        .values="${row}"
        .options="${{}}"
        .data="${{ audit, current, dataset }}"
        .onEdit=${(data)=>this._onEditEvent(EDIT_EVENT.EDIT_ITEM, data )}
        .onEvent=${(...args) => this._onEditEvent(...args)}
        .onSelector=${(...args)=>this._onEditEvent(EDIT_EVENT.SELECTOR, [...args] )}
        .msg=${this.msg}
      ></form-row>`);
  }

  deffields() {
    const { dataset, current } = this
    const ntype_id = dataset.groups.filter((group) => (
      (group.groupname === "nervatype") && (group.groupvalue === current.type )))[0].id
    if (current.type === "trans") {
      return dataset.deffield.filter((df) => (
        (df.nervatype === ntype_id) && (df.visible === 1))).filter( (df) => (
        (df.subtype === current.item.transtype) || (df.subtype === null)) ).map(
          (df)=>({value: df.fieldname, text: df.description }))
    } 
      return dataset.deffield.filter((df) => (
        (df.nervatype === ntype_id) && (df.visible === 1))).map(
          (df)=>({value: df.fieldname, text: df.description }))
    
  }

  render() {
    const { audit, current } = this
    const rows = this.prepareRows()
    const pageCount = Math.ceil(rows.length/this.pageSize)
    return html`<div id="${this.id}" class="panel" >
      ${((audit !== 'readonly')||(pageCount > 1)) ? html`<div class="container-row">
        ${(audit !== 'readonly') ? html`<div class="cell mobile">
          <div class="cell padding-small" >
            <form-select id="sel_deffield"
              label="${this.msg("", { id: "fields_view" })}"
              .onChange=${
                (value) => this._onEditEvent(EDIT_EVENT.CHANGE , { fieldname: "deffield", value: value.value } )
              }
              .options=${this.deffields()} 
              .isnull="${true}" value="${current.deffield||""}" >
            </form-select>
          </div>
          ${(current.deffield && (current.deffield !== "")) ? html`<div class="cell" >
            <form-button id="btn_new" 
              label="${this.msg("", { id: "label_new" })}"
              .style="${{ padding: "6px 16px" }}"
              icon="Plus" type="${BUTTON_TYPE.BORDER}"
              @click=${()=>this._onEditEvent(EDIT_EVENT.CHECK_EDITOR , [{fieldname: current.deffield}, EDITOR_EVENT.NEW_FIELDVALUE] )} 
            >${this.msg("", { id: "label_new" })}</form-button>
          </div>` : nothing }
        </div>` : nothing }
        ${(pageCount > 1) ? html`<div class="paginator-cell">
          <form-pagination id="${`${this.id}_top_pagination`}"
            pageIndex=${this.currentPage} pageSize=${this.pageSize} pageCount=${pageCount} 
            ?canPreviousPage=${(this.currentPage > 1)} 
            ?canNextPage=${(this.currentPage < pageCount)} 
            ?hidePageSize=${this.hidePaginatonSize}
            .onEvent=${(key, value) => this._onPagination(key, value)} ></form-pagination>
        </div>`:nothing}
      </div>` : nothing }
      ${this.renderRows(rows, pageCount)}
    </div>`
  }
}