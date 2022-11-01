import { LitElement, html } from 'lit';
import { styleMap } from 'lit/directives/style-map.js';
import { ifDefined } from 'lit/directives/if-defined.js';

import '../Icon/form-icon.js';
import '../Input/form-input.js'
import '../Button/form-button.js'
import '../Pagination/form-pagination.js'

import { styles } from './Table.styles.js'
import { PAGINATION_TYPE, BUTTON_TYPE, INPUT_TYPE } from '../../../config/enums.js'

export class Table extends LitElement {
  constructor() {
    super();
    this.id = Math.random().toString(36).slice(2);
    this.name = undefined;
    this.rows = [];
    this.fields = {};
    this.rowKey = "id";
    this.pagination = PAGINATION_TYPE.TOP;
    this.currentPage = 1;
    this.pageSize = 10;
    this.hidePaginatonSize = false;
    this.tableFilter = false;
    this.filterPlaceholder = undefined;
    this.filterValue = "";
    this.labelYes = "YES";
    this.labelNo = "NO";
    this.labelAdd = "";
    this.addIcon = "Plus";
    this.tablePadding = undefined;
    this.sortCol = undefined;
    this.sortAsc = true;
    this.style = {};
  }

  static get properties() {
    return {
      id: { type: String },
      name: { type: String },
      rowKey: { type: String },
      rows: { type: Array },
      fields: { type: Object },
      pagination: { type: String },
      currentPage: { type: Number },
      pageSize: { type: Number },
      hidePaginatonSize: { type: Boolean },
      tableFilter: { type: Boolean },
      filterPlaceholder: { type: String },
      filterValue: { type: String },
      labelYes: { type: String },
      labelNo: { type: String },
      labelAdd: { type: String },
      addIcon: { type: String },
      tablePadding: {type: String },
      sortCol: { type: String, attribute: false },
      sortAsc: { type: Boolean, attribute: false },
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

    if((Object.keys(this.fields).length === 0) && this.rows && Array.isArray(this.rows) && (this.rows.length > 0) ){
      Object.keys(this.rows[0]).forEach(field => {
        this.fields[field] = { fieldtype:'string', label: field }
      })
    }
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

  _onSort(e) {
    e.stopPropagation();
    
    const thisSort = e.target.dataset.sort;
    if(this.sortCol && this.sortCol === thisSort) this.sortAsc = !this.sortAsc;
    this.sortCol = thisSort;
    this.rows.sort((a, b) => {
      if(a[this.sortCol] < b[this.sortCol]) return this.sortAsc?1:-1;
      return this.sortAsc?-1:1;
    });
    // this.requestUpdate();
  }

  _onRowSelected(e, row, index) {
    e.stopPropagation();
    if(!row.disabled){
      if(this.onRowSelected){
        this.onRowSelected(row, index);
      }
      this.dispatchEvent(new CustomEvent('row_selected', {
        bubbles: true, composed: true,
        detail: {
          row, index
        }
      }));
    }
  }

  _onEditCell(e, fieldname, resultValue, rowData) {
    e.stopPropagation();
    if(this.onEditCell){
      this.onEditCell(fieldname, resultValue, rowData);
    }
    this.dispatchEvent(new CustomEvent('edit_cell', {
      bubbles: true, composed: true,
      detail: {
        fieldname, resultValue, rowData
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

  columns() {
    const numberCell = (value, label, style) => html`<div class="number-cell">
        <span class="cell-label">${label}</span>
        <span style="${styleMap(style)}" >${new Intl.NumberFormat('default').format(value)}</span>
      </div>`

    const dateCell = (value, label, dateType) => {
      let fmtValue = ""
      const dateValue = new Date(value)
      if (dateValue instanceof Date && !Number.isNaN(dateValue.valueOf())) {
        switch (dateType) {
          case "date":
            fmtValue = new Intl.DateTimeFormat(
              'default', {
                year: 'numeric', month: '2-digit', day: '2-digit'
              }
            ).format(dateValue)
            break;
          
          case "time":
            fmtValue = new Intl.DateTimeFormat(
              'default', {
                hour: '2-digit', minute: '2-digit', hour12: false
              }
            ).format(dateValue)  
            break;
        
          default:
            fmtValue = new Intl.DateTimeFormat(
              'default', {
                year: 'numeric', month: '2-digit', day: '2-digit', 
                hour: '2-digit', minute: '2-digit', hour12: false
              }
            ).format(dateValue) 
            break;
        }
      }
      return html`<span class="cell-label">${label}</span><span>${fmtValue}</span>`
    }

    const boolCell = (value, label) => {
      if((value === 1) || (value === "true") || (value === true)){
        return html`
          <span class="cell-label">${label}</span>
          <form-icon iconKey="CheckSquare" ></form-icon>
          <span class="middle"> ${this.labelYes}</span>
        `
      }
      return html`
        <span class="cell-label">${label}</span>
        <form-icon iconKey="SquareEmpty" ></form-icon>
        <span class="middle"> ${this.labelNo}</span>
      `
    }

    const linkCell = (value, label, fieldname, resultValue, rowData) => html`
        <span class="cell-label">${label}</span>
        <span id=${`link_${rowData[this.rowKey]}`} class="link-cell"
          @click=${(event)=>this._onEditCell(event, fieldname, resultValue, rowData)} 
          >${value}</span>
      `

    const stringCell = (value, label, style) => html`
        <span class="cell-label">${label}</span>
        <span style="${styleMap(style)}" >${value}</span>
      `

    let cols = []
    Object.keys(this.fields).forEach((fieldname) => {
      if(this.fields[fieldname].columnDef){
        cols = [...cols, {
          id: fieldname,
          Header: fieldname,
          headerStyle: {},
          cellStyle: {},
          ...this.fields[fieldname].columnDef
        }]
      } else {
        const coldef = {
          id: fieldname,
          Header: this.fields[fieldname].label || fieldname,
          headerStyle: {},
          cellStyle: {}
        }
        switch (this.fields[fieldname].fieldtype) {
          case "number":
            coldef.headerStyle.textAlign = "right"
            coldef.Cell = ({ row, value }) => {
              const style = {}
              if(this.fields[fieldname].format){
                style.fontWeight = "bold";
                if(row.edited){
                  style.textDecoration = "line-through";
                } else if(value !== 0){
                  style.color = "red"
                } else {
                  style.color = "green"
                }
              }
              return numberCell(value, this.fields[fieldname].label, style)
            }
            break;
          
          case "datetime":
          case "date":
          case "time":
            coldef.Cell = ({ value }) => dateCell(
              value, this.fields[fieldname].label, this.fields[fieldname].fieldtype
            )
            break;
          
          case "bool":
            coldef.Cell = ({ value }) => boolCell(value, this.fields[fieldname].label)
            break;

          case "deffield":
            coldef.Cell = ({ row, value }) => {
              switch (row.fieldtype) {
                case "bool":
                  return boolCell(value, this.fields[fieldname].label)

                case "integer":
                case "float":
                  return numberCell(value, this.fields[fieldname].label, {})

                case "customer":
                case "tool":
                case "product":
                case "trans": 
                case "transitem":
                case "transmovement": 
                case "transpayment":
                case "project":
                case "employee":
                case "place":
                case "urlink":
                  return linkCell(
                    row.export_deffield_value, this.fields[fieldname].label, 
                    row.fieldtype, row[fieldname], row)

                default:
                  return stringCell(value, this.fields[fieldname].label, {})
              }
            };
            break;

          default:
            coldef.Cell = ({ row, value }) => {
              const style = {}
              if(row[`${fieldname}_color`]){
                style.color = row[`${fieldname}_color`]
              }
              if(Object.keys(row).includes(`export_${fieldname}`)){
                return linkCell(
                  row[`export_${fieldname}`], this.fields[fieldname].label, fieldname, 
                  row[fieldname], row
                )
              }
              return stringCell(value, this.fields[fieldname].label, style)
            }
        }
        if(this.tablePadding){
          coldef.headerStyle.padding = this.tablePadding
          coldef.cellStyle.padding = this.tablePadding
        }
        if(this.fields[fieldname].verticalAlign) {
          coldef.cellStyle.verticalAlign = this.fields[fieldname].verticalAlign
        }
        if(this.fields[fieldname].textAlign) {
          coldef.cellStyle.textAlign = this.fields[fieldname].textAlign
        }
        cols = [...cols, coldef]
      }
    })
    return cols
  }

  renderHeader(cols) {
    return html`<thead><tr>
      ${cols.map(col => html`<th 
        data-sort="${col.id}" 
        class="sort ${
          (this.sortCol === col.id)
          ? (this.sortAsc)
            /* c8 ignore next 1 */
            ? "sort-asc" : "sort-desc"
          : "sort-none"
        }"
        style="${styleMap(col.headerStyle)}"
        @click=${this._onSort} >${col.Header}</th>`)}
    </tr></thead>`
  }

  filterRows() {
    let pageRows = this.rows
    const getValidRow = (row, filter)=>{
      let find = false;
      Object.keys(this.fields).forEach((fieldname) => {
        if(String(row[fieldname]).toLowerCase().indexOf(filter)>-1){
          find = true;
        }
      });
      return find;
    }
    if(this.tableFilter && (this.filterValue !== "")){
      pageRows = pageRows.filter(row => getValidRow(row, String(this.filterValue).toLowerCase()))
    }
    return pageRows
  }

  renderRows(cols, rows, pageCount) {
    let pageRows = rows
    if((this.pagination !== PAGINATION_TYPE.NONE) && (pageCount > 1)){
      const start = (this.currentPage-1)*this.pageSize;
      const end = this.currentPage*this.pageSize;
      // pageRows = pageRows.filter((row, index) => (index >= start && index < end))
      pageRows = pageRows.slice(start, end)
    }
    return pageRows
      .map((row, index)=> html`<tr id="${`row_${row[this.rowKey]}`}"
        class="${(row.disabled) ? "cursor-disabled" : (this.onRowSelected) ? "cursor-pointer" : ""}"
        @click=${(event)=>this._onRowSelected(event, row, index)}
      >${cols.map(col => html`<td style="${styleMap(col.cellStyle)}">${
        (col.Cell) ? col.Cell({ row, value: row[col.id]}) : row[col.id]}</td>`)}</tr>`);
  }

  render() {
    const cols = this.columns()
    const rows = this.filterRows()
    const pageCount = Math.ceil(rows.length/this.pageSize)
    const topPagination = ((pageCount > 1) && ((this.pagination === PAGINATION_TYPE.TOP) || (this.pagination === PAGINATION_TYPE.ALL)))
    const bottomPagination = ((pageCount > 1) && ((this.pagination === PAGINATION_TYPE.BOTTOM) || (this.pagination === PAGINATION_TYPE.ALL)))
    return html`<div class="responsive" >
      ${(this.tableFilter || topPagination) ?
      html`<div>
        ${topPagination ?
        html`<div>
          <form-pagination id="${`${this.id}_top_pagination`}"
            pageIndex=${this.currentPage} pageSize=${this.pageSize} pageCount=${pageCount} 
            ?canPreviousPage=${(this.currentPage > 1)} 
            ?canNextPage=${(this.currentPage < pageCount)} 
            ?hidePageSize=${this.hidePaginatonSize}
            .onEvent=${(key, value) => this._onPagination(key, value)} ></form-pagination>
        </div>`:``}
        ${(this.tableFilter) ?
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
            >${this.labelAdd}</form-button>
          </div>`:``}
        </div>`:``}
      </div>`:``}
      <div class="table-wrap" >
        <table id="${this.id}" name="${ifDefined(this.name)}"
          class="ui-table" style="${styleMap(this.style)}" >
          ${this.renderHeader(cols)}
          <tbody>${this.renderRows(cols, rows, pageCount)}</tbody>
        </table>
      </div>
    </div>
    ${bottomPagination ?
      html`<div>
        <form-pagination id="${`${this.id}_bottom_pagination`}"
          pageIndex=${this.currentPage} pageSize=${this.pageSize} pageCount=${pageCount} 
          ?canPreviousPage=${(this.currentPage > 1)} 
          ?canNextPage=${(this.currentPage < pageCount)} 
          ?hidePageSize=${this.hidePaginatonSize}
          .onEvent=${(key, value) => this._onPagination(key, value)} ></form-pagination>
    </div>`:``}
    `
  }

}