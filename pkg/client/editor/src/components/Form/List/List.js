import { LitElement, html, nothing } from 'lit';
import { styleMap } from 'lit/directives/style-map.js';
import { ifDefined } from 'lit/directives/if-defined.js';

import '../Icon/form-icon.js';
import '../Input/form-input.js'
import '../Button/form-button.js'
import '../Pagination/form-pagination.js'

import { styles } from './List.styles.js'
import { PAGINATION_TYPE, BUTTON_TYPE, INPUT_TYPE } from '../../../config/enums.js'

export class List extends LitElement {
  constructor() {
    super();
    this.id = Math.random().toString(36).slice(2);
    this.name = undefined;
    this.rows = [];
    this.pagination = PAGINATION_TYPE.TOP;
    this.currentPage = 1;
    this.pageSize = 10;
    this.hidePaginatonSize = false;
    this.listFilter = false;
    this.filterPlaceholder = undefined;
    this.filterValue = "";
    this.labelAdd = "";
    this.addIcon = "Plus";
    this.editIcon = "Edit";
    this.deleteIcon = "Times";
    this.style = {};
  }

  static get properties() {
    return {
      id: { type: String },
      name: { type: String },
      rows: { type: Array },
      pagination: { type: String },
      currentPage: { type: Number },
      pageSize: { type: Number },
      hidePaginatonSize: { type: Boolean },
      listFilter: { type: Boolean },
      filterPlaceholder: { type: String },
      filterValue: { type: String },
      labelAdd: { type: String },
      addIcon: { type: String },
      editIcon: { type: String },
      deleteIcon: { type: String },
      style: { type: Object },
    };
  }

  static get styles () {
    return [
      styles
    ]
  }

  connectedCallback() {
    super.connectedCallback();

    if(this.currentPage > Math.ceil(this.rows.length/this.pageSize)){
      this.currentPage = Math.ceil(this.rows.length/this.pageSize)
    }
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
    if(this.onCurrentPage){
      this.onCurrentPage(value)
    }
  }

  _onEdit(e, rowData, index) {
    e.stopPropagation();
    if(this.onEdit){
      this.onEdit(rowData, index);
    }
    this.dispatchEvent(new CustomEvent('edit', {
      bubbles: true, composed: true,
      detail: {
        rowData, index
      }
    }));
  }

  _onDelete(e, rowData, index) {
    e.stopPropagation();
    if(this.onDelete){
      this.onDelete(rowData, index);
    }
    this.dispatchEvent(new CustomEvent('delete', {
      bubbles: true, composed: true,
      detail: {
        rowData, index
      }
    }));
  }

  _onAddItem() {
    if(this.onAddItem){
      this.onAddItem({});
    }
    this.dispatchEvent(new CustomEvent('add_item', {
      bubbles: true, composed: true,
      detail: {}
    }));
  }

  _onFilterChange(data){
    this.filterValue = data.value
    this.dispatchEvent(
      new CustomEvent('filter_change', {
        bubbles: true, composed: true,
        detail: {
          ...data
        }
      })
    );
  }

  filterRows() {
    let pageRows = this.rows
    const getValidRow = (row, filter)=>{
      if(String(row.lslabel).toLowerCase().indexOf(filter)>-1 || 
        String(row.lsvalue).toLowerCase().indexOf(filter)>-1){
          return true
      }
      return false
    }
    if(this.listFilter && (this.filterValue !== "")){
      pageRows = pageRows.filter(row => getValidRow(row, String(this.filterValue).toLowerCase()))
    }
    return pageRows
  }

  renderRows(rows, pageCount) {
    let pageRows = rows
    if((this.pagination !== PAGINATION_TYPE.NONE) && (pageCount > 1)){
      const start = (this.currentPage-1)*this.pageSize;
      const end = this.currentPage*this.pageSize;
      // pageRows = pageRows.filter((row, index) => (index >= start && index < end))
      pageRows = pageRows.slice(start, end)
    }
    return pageRows
      .map((row, index)=> html`<li class="list-row border-bottom">
        ${(this.onEdit) ? html`<div 
          id="${`row_edit_${index}`}" class="edit-cell" 
          @click=${(event)=>this._onEdit(event, row, index)} >
          <form-icon iconKey="${this.editIcon}" ></form-icon>
        </div>` : nothing}
        <div id="${`row_item_${index}`}"
          class="value-cell ${(this.onEdit) ? "cursor" : ""}"
          @click=${(this.onEdit) ? (event)=>this._onEdit(event, row, index) : null}>
          <div class="border-bottom label" >
            <span>${row.lslabel}</span>
          </div>
          <div class="value" >
            <span>${row.lsvalue}</span>
          </div>
        </div>
        ${(this.onDelete) ? html`<div 
          id="${`row_delete_${index}`}" class="delete-cell" 
          @click=${(event)=>this._onDelete(event, row, index)} >
          <form-icon iconKey="${this.deleteIcon}" ></form-icon>
        </div>` : nothing}
      </li>`);
  }

  render() {
    const rows = this.filterRows()
    const pageCount = Math.ceil(rows.length/this.pageSize)
    const topPagination = ((pageCount > 1) && ((this.pagination === PAGINATION_TYPE.TOP) || (this.pagination === PAGINATION_TYPE.ALL)))
    const bottomPagination = ((pageCount > 1) && ((this.pagination === PAGINATION_TYPE.BOTTOM) || (this.pagination === PAGINATION_TYPE.ALL)))
    return html`<div class="responsive" >
      ${(this.listFilter || topPagination) ?
      html`<div>
        ${topPagination ?
        html`<div>
          <form-pagination id="${`${this.id}_top_pagination`}"
            pageIndex=${this.currentPage} pageSize=${this.pageSize} pageCount=${pageCount} 
            ?canPreviousPage=${(this.currentPage > 1)} 
            ?canNextPage=${(this.currentPage < pageCount)} 
            ?hidePageSize=${this.hidePaginatonSize}
            .onEvent=${(key, value) => this._onPagination(key, value)} ></form-pagination>
        </div>`:nothing}
        ${(this.listFilter) ?
        html`<div class="row full" >
          <div class="cell" >
            <form-input id="filter" type="${INPUT_TYPE.TEXT}" 
              .style="${{ "border-radius": 0, margin: "1px 0 2px" }}"
              label="${ifDefined(this.filterPlaceholder)}"
              placeholder="${ifDefined(this.filterPlaceholder)}"
              value="${this.filterValue}" ?full="${true}"
              .onChange=${
                (event) => this._onFilterChange({ value: event.value, old: this.filterValue }) 
              }
            ></form-input>
          </div>
          ${(this.onAddItem)?html`<div class="cell" style="${styleMap({ width: "20px" })}" >
            <form-button id="btn_add" icon="${this.addIcon}"
              label="${this.labelAdd}"
              .style="${{ padding: "8px 16px", "border-radius": 0, margin: "1px 0 2px 1px" }}"
              @click=${()=>this._onAddItem()} type="${BUTTON_TYPE.BORDER}"
            >${this.labelAdd}
            </form-button>
          </div>`:nothing}
        </div>`:nothing}
      </div>`:nothing}
      <ul id="${this.id}" name="${ifDefined(this.name)}"
        class="list" style="${styleMap(this.style)}" >
        ${this.renderRows(rows, pageCount)}
      </ul>
    </div>
    ${bottomPagination ?
      html`<div>
        <form-pagination id="${`${this.id}_bottom_pagination`}"
          pageIndex=${this.currentPage} pageSize=${this.pageSize} pageCount=${pageCount} 
          ?canPreviousPage=${(this.currentPage > 1)} 
          ?canNextPage=${(this.currentPage < pageCount)} 
          ?hidePageSize=${this.hidePaginatonSize}
          .onEvent=${(key, value) => this._onPagination(key, value)} ></form-pagination>
    </div>`:nothing}
    `
  }

}