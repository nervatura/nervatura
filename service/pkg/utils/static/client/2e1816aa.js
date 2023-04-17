import{i as e,s as t,x as i,a as l}from"./3cff389a.js";import{l as a,o as s}from"./7854850a.js";import"./4318ffa4.js";import"./5217a594.js";import{P as n,I as o,B as r}from"./64e23944.js";const d=e`
:host {
	font-family: var(--font-family);
	font-size: var(--font-size);
	color: var(--text-1);
	fill: var(--text-1);
}
div {
  box-sizing: border-box;
}
.row {
  display: table;
}
.full {
  width: 100%; 
}
.cell {
  display: table-cell;
  vertical-align: middle;
}
.responsive {
  display: grid; 
  width: 100%;
}
.list {
  overflow-x: auto;
  width: 100%;
  border: 1px solid rgba(var(--neutral-1), 0.2);
  border-bottom: none;
  list-style-type: none;
  padding: 0;
  margin: 0;
  box-sizing: border-box;
}
.list li:nth-child(odd) {
  background-color: rgba(var(--functional-yellow),0.1);
}
.list-row {
  display: table;
  width: 100%;
}
.border-bottom {
  border-bottom: 1px solid rgba(var(--neutral-1), 0.2);
}
.edit-cell {
  display: table-cell;
  vertical-align: middle;
  text-align: center;
  width: 45px;
  cursor: pointer;
}
.edit-cell:hover {
  fill: rgb(var(--functional-green));
}
.value-cell {
  display: table-cell;
  vertical-align: middle;
}
.cursor {
  cursor: pointer;
}
.label {
  width: 100%;
  font-weight: bold;
  padding: 5px 8px 2px;
}
.value {
  width: 100%;
  padding: 2px 8px 5px;
}
.delete-cell {
  display: table-cell;
  vertical-align: middle;
  text-align: center;
  width: 45px;
  cursor: pointer;
}
.delete-cell:hover {
  fill: rgb(var(--functional-red));
}
`;customElements.define("form-list",class extends t{constructor(){super(),this.id=Math.random().toString(36).slice(2),this.name=void 0,this.rows=[],this.pagination=n.TOP,this.currentPage=1,this.pageSize=10,this.hidePaginatonSize=!1,this.listFilter=!1,this.filterPlaceholder=void 0,this.filterValue="",this.labelAdd="",this.addIcon="Plus",this.editIcon="Edit",this.deleteIcon="Times",this.style={}}static get properties(){return{id:{type:String},name:{type:String},rows:{type:Array},pagination:{type:String},currentPage:{type:Number},pageSize:{type:Number},hidePaginatonSize:{type:Boolean},listFilter:{type:Boolean},filterPlaceholder:{type:String},filterValue:{type:String},labelAdd:{type:String},addIcon:{type:String},editIcon:{type:String},deleteIcon:{type:String},style:{type:Object}}}static get styles(){return[d]}connectedCallback(){super.connectedCallback(),this.currentPage>Math.ceil(this.rows.length/this.pageSize)&&(this.currentPage=Math.ceil(this.rows.length/this.pageSize)),this.currentPage<1&&(this.currentPage=1)}_onPagination(e,t){if("setPageSize"===e)return this.currentPage=1,void(this.pageSize=t);this.currentPage=t,this.onCurrentPage&&this.onCurrentPage(t)}_onEdit(e,t,i){e.stopPropagation(),this.onEdit&&this.onEdit(t,i),this.dispatchEvent(new CustomEvent("edit",{bubbles:!0,composed:!0,detail:{rowData:t,index:i}}))}_onDelete(e,t,i){e.stopPropagation(),this.onDelete&&this.onDelete(t,i),this.dispatchEvent(new CustomEvent("delete",{bubbles:!0,composed:!0,detail:{rowData:t,index:i}}))}_onAddItem(){this.onAddItem&&this.onAddItem({}),this.dispatchEvent(new CustomEvent("add_item",{bubbles:!0,composed:!0,detail:{}}))}_onFilterChange(e){this.filterValue=e.value,this.dispatchEvent(new CustomEvent("filter_change",{bubbles:!0,composed:!0,detail:{...e}}))}filterRows(){let e=this.rows;return this.listFilter&&""!==this.filterValue&&(e=e.filter((e=>((e,t)=>String(e.lslabel).toLowerCase().indexOf(t)>-1||String(e.lsvalue).toLowerCase().indexOf(t)>-1)(e,String(this.filterValue).toLowerCase())))),e}renderRows(e,t){let a=e;if(this.pagination!==n.NONE&&t>1){const e=(this.currentPage-1)*this.pageSize,t=this.currentPage*this.pageSize;a=a.slice(e,t)}return a.map(((e,t)=>i`<li class="list-row border-bottom">
        ${this.onEdit?i`<div 
          id="${`row_edit_${t}`}" class="edit-cell" 
          @click=${i=>this._onEdit(i,e,t)} >
          <form-icon iconKey="${this.editIcon}" ></form-icon>
        </div>`:l}
        <div id="${`row_item_${t}`}"
          class="value-cell ${this.onEdit?"cursor":""}"
          @click=${this.onEdit?i=>this._onEdit(i,e,t):null}>
          <div class="border-bottom label" >
            <span>${e.lslabel}</span>
          </div>
          <div class="value" >
            <span>${e.lsvalue}</span>
          </div>
        </div>
        ${this.onDelete?i`<div 
          id="${`row_delete_${t}`}" class="delete-cell" 
          @click=${i=>this._onDelete(i,e,t)} >
          <form-icon iconKey="${this.deleteIcon}" ></form-icon>
        </div>`:l}
      </li>`))}render(){const e=this.filterRows(),t=Math.ceil(e.length/this.pageSize),d=t>1&&(this.pagination===n.TOP||this.pagination===n.ALL),h=t>1&&(this.pagination===n.BOTTOM||this.pagination===n.ALL);return i`<div class="responsive" >
      ${this.listFilter||d?i`<div>
        ${d?i`<div>
          <form-pagination id="${`${this.id}_top_pagination`}"
            pageIndex=${this.currentPage} pageSize=${this.pageSize} pageCount=${t} 
            ?canPreviousPage=${this.currentPage>1} 
            ?canNextPage=${this.currentPage<t} 
            ?hidePageSize=${this.hidePaginatonSize}
            .onEvent=${(e,t)=>this._onPagination(e,t)} ></form-pagination>
        </div>`:l}
        ${this.listFilter?i`<div class="row full" >
          <div class="cell" >
            <form-input id="filter" type="${o.TEXT}" 
              .style="${{"border-radius":0,margin:"1px 0 2px"}}"
              label="${a(this.filterPlaceholder)}"
              placeholder="${a(this.filterPlaceholder)}"
              value="${this.filterValue}" ?full="${!0}"
              .onChange=${e=>this._onFilterChange({value:e.value,old:this.filterValue})}
            ></form-input>
          </div>
          ${this.onAddItem?i`<div class="cell" style="${s({width:"20px"})}" >
            <form-button id="btn_add" icon="${this.addIcon}"
              label="${this.labelAdd}"
              .style="${{padding:"8px 16px","border-radius":0,margin:"1px 0 2px 1px"}}"
              @click=${()=>this._onAddItem()} type="${r.BORDER}"
            >${this.labelAdd}
            </form-button>
          </div>`:l}
        </div>`:l}
      </div>`:l}
      <ul id="${this.id}" name="${a(this.name)}"
        class="list" style="${s(this.style)}" >
        ${this.renderRows(e,t)}
      </ul>
    </div>
    ${h?i`<div>
        <form-pagination id="${`${this.id}_bottom_pagination`}"
          pageIndex=${this.currentPage} pageSize=${this.pageSize} pageCount=${t} 
          ?canPreviousPage=${this.currentPage>1} 
          ?canNextPage=${this.currentPage<t} 
          ?hidePageSize=${this.hidePaginatonSize}
          .onEvent=${(e,t)=>this._onPagination(e,t)} ></form-pagination>
    </div>`:l}
    `}});
