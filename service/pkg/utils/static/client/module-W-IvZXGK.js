import{a as e,i as t,x as i,E as l}from"./module-CgqbBeKY.js";import{a,o as s}from"./module-pdXndwQj.js";import"./module-C3sfG5TL.js";import"./module-CSXOkCWC.js";import{P as d,I as o,B as n,D as r,j as h,E as c}from"./main-CONOq8Fm.js";const u=e`
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
.table-wrap {
  display: block;
  max-width: 100%;
  overflow-x: auto;
  overflow-y: hidden;
}
.sort {
  cursor: pointer;
  cursor: hand;
}
.sort:after {
  padding-left: 1em;
  font-size: 0.5em;
}
.sort-none:after {
  padding-left: 0.5em;
  font-size: 1em;
}
.sort-asc:after {
  content: '▲';
  font-size: 14px;
  color: rgb(var(--functional-yellow));
}
.sort-desc:after {
  content: '▼';
  font-size: 14px;
  color: rgb(var(--functional-yellow));
}
.sort-order {
  margin-left: 0.5em;
}
.number-cell {
  text-align: right;
}
.link-cell {
  color: rgb(var(--functional-blue));
  cursor: pointer;
}
.link-cell:hover {
  text-decoration: underline;
}
.ui-table {
  overflow-x: auto;
  border-collapse: collapse;
  border-spacing: 0;
  width: 100%;
  font-size: 14px;
  border: 1px solid rgba(var(--neutral-1), 0.2);
  box-sizing: border-box;
}
.ui-table thead th {
  border-bottom: 1px solid rgba(var(--neutral-1),0.2);
}
.ui-table tbody tr:nth-child(odd) {
  background-color: rgba(var(--functional-yellow),0.1);
}
.ui-table tbody th, .ui-table tbody td {
  vertical-align: top;
}
.ui-table tbody tr:hover {
  color: rgb(var(--functional-yellow));
}
.ui-table th:hover {
  color: rgb(var(--functional-yellow));
}
.cell-label { 
	display: none;
  font-weight: bold;
  font-size: 12px;
}
/* Mobile first styles: Begin with the stacked presentation at narrow widths */ 
@media only all {
	/* Hide the table headers */ 
	.ui-table thead td, 
	.ui-table thead th {
		display: none;
    font-size: 13px;
    background-color: rgb(var(--accent-1));
    color: rgba(var(--accent-1c), 0.85);
    fill: rgba(var(--accent-1c), 0.85);
    white-space: nowrap;
	}
	/* Show the table cells as a block level element */ 
	.ui-table td,
	.ui-table th { 
		text-align: left;
		display: block;
    padding: 8px;
	}
	/* Add a fair amount of top margin to visually separate each row when stacked */  
	.ui-table tbody th {
		margin-top: 3em;
	}
	/* Make the label elements a percentage width */ 
	.cell-label { 
		padding: 4px 6px;
		min-width: 30%; 
		display: inline-block;
		margin: -.4em 1em -.1em -.4em;
	}
}
/* Breakpoint to show as a standard table at 560px (35em x 16px) or wider */ 
@media ( min-width: 35em ) {
	/* Show the table header rows */ 
	.ui-table td,
	.ui-table th,
	.ui-table tbody th,
	.ui-table tbody td,
	.ui-table thead td,
	.ui-table thead th {
		display: table-cell;
    margin: 0;
	}
	/* Hide the labels in each cell */ 
  .cell-label { 
		display: none;
	}
}
/* Hack to make IE9 and WP7.5 treat cells like block level elements, scoped to ui-responsive class */ 
/* Applied in a max-width media query up to the table layout breakpoint so we don't need to negate this*/ 
@media ( max-width: 35em ) {
	.ui-table td,
	.ui-table th {
		width: 100%;
		-webkit-box-sizing: border-box;
		-moz-box-sizing: border-box;
		box-sizing: border-box;
		float: left;
		clear: left;
    font-size: 13px;
    padding: 4px 6px;
    line-height: 11.5px;
	}
  .number-cell {
    text-align: left;
  }
}
.cursor-pointer {
  cursor: pointer;
}
.cursor-disabled {
  cursor: not-allowed;
}
.middle {
  vertical-align: middle;
}
`;customElements.define("form-table",class extends t{constructor(){super(),this.id=Math.random().toString(36).slice(2),this.name=void 0,this.rows=[],this.fields={},this.rowKey="id",this.pagination=d.TOP,this.currentPage=1,this.pageSize=10,this.hidePaginatonSize=!1,this.tableFilter=!1,this.filterPlaceholder=void 0,this.filterValue="",this.labelYes="YES",this.labelNo="NO",this.labelAdd="",this.addIcon="Plus",this.tablePadding=void 0,this.sortCol=void 0,this.sortAsc=!0,this.style={}}static get properties(){return{id:{type:String},name:{type:String},rowKey:{type:String},rows:{type:Array},fields:{type:Object},pagination:{type:String},currentPage:{type:Number},pageSize:{type:Number},hidePaginatonSize:{type:Boolean},tableFilter:{type:Boolean},filterPlaceholder:{type:String},filterValue:{type:String},labelYes:{type:String},labelNo:{type:String},labelAdd:{type:String},addIcon:{type:String},tablePadding:{type:String},sortCol:{type:String,attribute:!1},sortAsc:{type:Boolean,attribute:!1},style:{type:Object}}}static get styles(){return[u]}connectedCallback(){super.connectedCallback(),this.rows||(this.rows=[]),0===Object.keys(this.fields).length&&this.rows&&Array.isArray(this.rows)&&this.rows.length>0&&Object.keys(this.rows[0]).forEach(e=>{this.fields[e]={fieldtype:"string",label:e}}),this.currentPage>Math.ceil(this.rows.length/this.pageSize)&&(this.currentPage=Math.ceil(this.rows.length/this.pageSize)),this.currentPage<1&&(this.currentPage=1)}_onPagination(e,t){if("setPageSize"===e)return this.currentPage=1,void(this.pageSize=t);this.currentPage=t,this.onCurrentPage&&this.onCurrentPage(t)}_onSort(e){e.stopPropagation();const t=e.target.dataset.sort;this.sortCol&&this.sortCol===t&&(this.sortAsc=!this.sortAsc),this.sortCol=t,this.rows.sort((e,t)=>e[this.sortCol]<t[this.sortCol]?this.sortAsc?1:-1:this.sortAsc?-1:1)}_onRowSelected(e,t,i){e.stopPropagation(),t.disabled||(this.onRowSelected&&this.onRowSelected(t,i),this.dispatchEvent(new CustomEvent("row_selected",{bubbles:!0,composed:!0,detail:{row:t,index:i}})))}_onEditCell(e,t,i,l){e.stopPropagation(),this.onEditCell&&this.onEditCell(t,i,l),this.dispatchEvent(new CustomEvent("edit_cell",{bubbles:!0,composed:!0,detail:{fieldname:t,resultValue:i,rowData:l}}))}_onAddItem(){this.onAddItem&&this.onAddItem({}),this.dispatchEvent(new CustomEvent("add_item",{bubbles:!0,composed:!0,detail:{}}))}_onFilterChange(e){this.filterValue=e.value,this.dispatchEvent(new CustomEvent("filter_change",{bubbles:!0,composed:!0,detail:{...e}}))}columns(){const e=(e,t,l)=>i`<div class="number-cell">
        <span class="cell-label">${t}</span>
        <span style="${a(l)}" >${new Intl.NumberFormat("default").format(e)}</span>
      </div>`,t=(e,t)=>1===e||"true"===e||!0===e?i`
          <span class="cell-label">${t}</span>
          <form-icon iconKey="CheckSquare" ></form-icon>
          <span class="middle"> ${this.labelYes}</span>
        `:i`
        <span class="cell-label">${t}</span>
        <form-icon iconKey="SquareEmpty" ></form-icon>
        <span class="middle"> ${this.labelNo}</span>
      `,l=(e,t,l,a,s)=>i`
        <span class="cell-label">${t}</span>
        <span id=${`link_${s[this.rowKey]}`} class="link-cell"
          @click=${e=>this._onEditCell(e,l,a,s)} 
          >${e}</span>
      `,s=(e,t,l)=>i`
        <span class="cell-label">${t}</span>
        <span style="${a(l)}" >${e}</span>
      `;let d=[];return Object.keys(this.fields).forEach(a=>{if(this.fields[a].columnDef)d=[...d,{id:a,Header:a,headerStyle:{},cellStyle:{},...this.fields[a].columnDef}];else{const o={id:a,Header:this.fields[a].label||a,headerStyle:{},cellStyle:{}};switch(this.fields[a].fieldtype){case"number":o.headerStyle.textAlign="right",o.Cell=({row:t,value:i})=>{const l={};return this.fields[a].format&&(l.fontWeight="bold",t.edited?l.textDecoration="line-through":l.color=0!==i?"red":"green"),e(i,this.fields[a].label,l)};break;case"datetime":case"date":case"time":o.Cell=({value:e})=>((e,t,l)=>{let a="";const s=new Date(e);if(s instanceof Date&&!Number.isNaN(s.valueOf()))switch(l){case"date":a=new Intl.DateTimeFormat("default",{year:"numeric",month:"2-digit",day:"2-digit"}).format(s);break;case"time":a=new Intl.DateTimeFormat("default",{hour:"2-digit",minute:"2-digit",hour12:!1}).format(s);break;default:a=new Intl.DateTimeFormat("default",{year:"numeric",month:"2-digit",day:"2-digit",hour:"2-digit",minute:"2-digit",hour12:!1}).format(s)}return i`<span class="cell-label">${t}</span><span>${a}</span>`})(e,this.fields[a].label,this.fields[a].fieldtype);break;case"bool":o.Cell=({value:e})=>t(e,this.fields[a].label);break;case"deffield":o.Cell=({row:i,value:d})=>{switch(i.fieldtype){case"bool":return t(d,this.fields[a].label);case"integer":case"float":return e(d,this.fields[a].label,{});case"customer":case"tool":case"product":case"trans":case"transitem":case"transmovement":case"transpayment":case"project":case"employee":case"place":case"urlink":return l(i.export_deffield_value,this.fields[a].label,i.fieldtype,i[a],i);default:return s(d,this.fields[a].label,{})}};break;default:o.Cell=({row:e,value:t})=>{const i={};return e[`${a}_color`]&&(i.color=e[`${a}_color`]),Object.keys(e).includes(`export_${a}`)?l(e[`export_${a}`],this.fields[a].label,a,e[a],e):s(t,this.fields[a].label,i)}}this.tablePadding&&(o.headerStyle.padding=this.tablePadding,o.cellStyle.padding=this.tablePadding),this.fields[a].verticalAlign&&(o.cellStyle.verticalAlign=this.fields[a].verticalAlign),this.fields[a].textAlign&&(o.cellStyle.textAlign=this.fields[a].textAlign),d=[...d,o]}}),d}renderHeader(e){return i`<thead><tr>
      ${e.map(e=>i`<th 
        data-sort="${e.id}" 
        class="sort ${this.sortCol===e.id?this.sortAsc?"sort-asc":"sort-desc":"sort-none"}"
        style="${a(e.headerStyle)}"
        @click=${this._onSort} >${e.Header}</th>`)}
    </tr></thead>`}filterRows(){let e=this.rows;const t=(e,t)=>{let i=!1;return Object.keys(this.fields).forEach(l=>{String(e[l]).toLowerCase().indexOf(t)>-1&&(i=!0)}),i};return this.tableFilter&&""!==this.filterValue&&(e=e.filter(e=>t(e,String(this.filterValue).toLowerCase()))),e}renderRows(e,t,l){let s=t;if(this.pagination!==d.NONE&&l>1){const e=(this.currentPage-1)*this.pageSize,t=this.currentPage*this.pageSize;s=s.slice(e,t)}return s.map((t,l)=>i`<tr id="${`row_${t[this.rowKey]||l}`}"
        class="${t.disabled?"cursor-disabled":this.onRowSelected?"cursor-pointer":""}"
        @click=${e=>this._onRowSelected(e,t,l)}
      >${e.map(e=>i`<td style="${a(e.cellStyle)}">${e.Cell?e.Cell({row:t,value:t[e.id]}):t[e.id]}</td>`)}</tr>`)}render(){const e=this.columns(),t=this.filterRows(),r=Math.ceil(t.length/this.pageSize),h=r>1&&(this.pagination===d.TOP||this.pagination===d.ALL),c=r>1&&(this.pagination===d.BOTTOM||this.pagination===d.ALL);return i`<div class="responsive" >
      ${this.tableFilter||h?i`<div>
        ${h?i`<div>
          <form-pagination id="${`${this.id}_top_pagination`}"
            pageIndex=${this.currentPage} pageSize=${this.pageSize} pageCount=${r} 
            ?canPreviousPage=${this.currentPage>1} 
            ?canNextPage=${this.currentPage<r} 
            ?hidePageSize=${this.hidePaginatonSize}
            .onEvent=${(e,t)=>this._onPagination(e,t)} ></form-pagination>
        </div>`:l}
        ${this.tableFilter?i`<div class="row full" >
          <div class="cell" >
            <form-input id="filter" type="${o.TEXT}" 
              .style="${{"border-radius":0,margin:"1px 0 2px"}}"
              label="${s(this.filterPlaceholder)}"
              placeholder="${s(this.filterPlaceholder)}"
              value="${this.filterValue}" ?full="${!0}"
              .onChange=${e=>this._onFilterChange({value:e.value,old:this.filterValue})}
            ></form-input>
          </div>
          ${this.onAddItem?i`<div class="cell" style="${a({width:"20px"})}" >
            <form-button id="btn_add" icon="${this.addIcon}"
              label="${this.labelAdd}"
              .style="${{padding:"8px 16px","border-radius":0,margin:"1px 0 2px 1px"}}"
              @click=${()=>this._onAddItem()} type="${n.BORDER}"
            >${this.labelAdd}</form-button>
          </div>`:l}
        </div>`:l}
      </div>`:l}
      <div class="table-wrap" >
        <table id="${this.id}" name="${s(this.name)}"
          class="ui-table" style="${a(this.style)}" >
          ${this.renderHeader(e)}
          <tbody>${this.renderRows(e,t,r)}</tbody>
        </table>
      </div>
    </div>
    ${c?i`<div>
        <form-pagination id="${`${this.id}_bottom_pagination`}"
          pageIndex=${this.currentPage} pageSize=${this.pageSize} pageCount=${r} 
          ?canPreviousPage=${this.currentPage>1} 
          ?canNextPage=${this.currentPage<r} 
          ?hidePageSize=${this.hidePaginatonSize}
          .onEvent=${(e,t)=>this._onPagination(e,t)} ></form-pagination>
    </div>`:l}
    `}});const f=e`
input {
  font-family: var(--font-family);
  font-size: var(--font-size);
  border-radius: 3px;
  border-width: 1px;
  padding: 8px;
  display: block;
  color: var(--text-1);
  background-color: rgba(var(--base-4), 1);
  border: 1px solid rgba(var(--neutral-1), 0.2);
  box-sizing: border-box;
}
input:disabled {
  opacity: 0.5;
}
input:focus, input:hover:not(:disabled) {
  color: rgb(var(--functional-green))!important;
}
::-webkit-calendar-picker-indicator {
  filter: invert(0.5);
}
.full{
  width: 100%;
}
`,p={[r.TIME]:5,[r.DATE]:10,[r.DATETIME]:16};customElements.define("form-datetime",class extends t{constructor(){super(),this.id=Math.random().toString(36).slice(2),this.name=void 0,this.value="",this.type=r.DATE,this.label=void 0,this.isnull=!0,this.picker=!1,this.disabled=!1,this.readonly=!1,this.autofocus=!1,this.full=!1,this.style={}}static get properties(){return{id:{type:String},name:{type:String,reflect:!0},value:{type:String,reflect:!0},type:{type:String,converter:e=>Object.values(r).includes(e)?e:r.DATE},label:{type:String},isnull:{type:Boolean},picker:{type:Boolean},disabled:{type:Boolean,reflect:!0},readonly:{type:Boolean,reflect:!0},autofocus:{type:Boolean,reflect:!0},full:{type:Boolean},style:{type:Object}}}_defaultValue(){const e=`${(new Date).toISOString().slice(0,10)}T${(new Date).toLocaleTimeString("en",{hour12:!1}).replace("24","00").slice(0,5)}`;switch(this.type){case r.DATE:return String(e).split("T")[0];case r.TIME:return String(e).split("T")[1].split(".")[0];default:return e}}_onInput(e){const t=e=>{const t=this.type!==r.DATE?`${e}:00`:e;t!==this.value&&(this.onChange&&this.onChange({value:t,old:this.value}),this.dispatchEvent(new CustomEvent("change",{bubbles:!0,composed:!0,detail:{value:t,old:this.value}})),this.value=t),this._input.value!==t&&(this._input.value=t)};if(e.target.value!==this.value){if(""===e.target.value&&!this.isnull)return void t(this._defaultValue());t(e.target.value)}}_onBlur(){this._input.value=this.value}_onKeyEvent(e){const t=()=>{this.onEnter&&(this.onEnter({value:this.value}),this.dispatchEvent(new CustomEvent("enter",{bubbles:!0,composed:!0,detail:{value:this.value}})))};"keydown"!==e.type&&"keypress"!==e.type||e.stopPropagation(),"keydown"===e.type&&13===e.keyCode&&(e.preventDefault(),t()),this.readonly||"keypress"!==e.type||13!==e.keyCode||t()}_onFocus(){this.picker&&this._input.showPicker()}firstUpdated(){this._input=this.renderRoot.querySelector("input")}render(){let e=this.value;return e.length>p[this.type]&&(e=e.slice(0,p[this.type])),i`<input 
      id="${this.id}"
      name="${s(this.name)}"
      .type="${this.type}"
      .value="${e}"
      ?disabled="${this.disabled}"
      ?readonly="${this.readonly}"
      ?autofocus="${this.autofocus}"
      aria-label="${s(this.label)}" 
      style="${a(this.style)}"
      class="${this.full?"full":""}"
      @input=${this._onInput}
      @blur=${this._onBlur}
      @keydown=${this._onKeyEvent}
      @keypress=${this._onKeyEvent}
      @focus=${this._onFocus}
    >`}static get styles(){return[f]}});const v=e`
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
textarea {
  font-family: var(--font-family);
  font-size: var(--font-size);
  border-radius: 3px;
  overflow: auto;
  padding: 8px;
  color: var(--text-1);
  background-color: rgba(var(--base-4), 1);
  border: 1px solid rgba(var(--neutral-1), 0.2);
  box-sizing: border-box;
}
textarea:focus, textarea:hover:not(:disabled) {
  color: rgb(var(--functional-green))!important;
}
textarea:disabled {
  opacity: 0.5;
}
textarea::placeholder, textarea::-ms-input-placeholder {
  opacity: 0.5;
}
.link {
  font-family: var(--font-family);
  font-size: var(--font-size);
  width: 100%;
  border-radius: 3px;
  padding: 6px 8px 5px;
  vertical-align: middle;
  min-height: 35px;
  color: var(--text-1);
  background-color: rgba(var(--base-4), 1);
  border: 1px solid rgba(var(--neutral-1), 0.2);
}
.link-text {
  font-size: 14px;
  cursor: pointer;
  color: rgb(var(--functional-blue));
}
.link-text:hover {
  text-decoration: underline;
}
.search-col {
  width: 38px;
}
.times-col {
  width: 34px;
}
.toggle-on {
  fill: rgb(var(--functional-green))!important;
}
.toggle {
  font-family: var(--font-family);
  font-size: var(--font-size);
  text-align: center;
  cursor: pointer;
  width: 100%;
  border-radius: 3px;
  fill: var(--text-1);
  background-color: rgba(var(--base-4), 1);
  border: 1px solid rgba(var(--neutral-1), 0.2);
}
.toggle:hover:not(:disabled) {
  fill: rgb(var(--functional-green))!important;
}
.toggle-off {
  opacity: 0.5;
}
.toggle-disabled {
  opacity: 0.5;
}
`;customElements.define("form-field",class extends t{constructor(){super(),this.id=Math.random().toString(36).slice(2),this.field={},this.values={},this.options={},this.data={dataset:{},current:{},audit:"all"},this.style={},this.msg=e=>e}static get properties(){return{id:{type:String},field:{type:Object},values:{type:Object},options:{type:Object},data:{type:Object},style:{type:Object}}}static get styles(){return[v]}_onChange({value:e,item:t,event_type:i,editName:l,fieldMap:a}){const s={id:this.field.id||1,name:l,event_type:i,value:e,extend:!(!a||!a.extend),refnumber:t&&t.label?t.label:this.field.link_label,label_field:a?a.label_field:void 0,item:t};this.onEdit&&this.onEdit(s),this.dispatchEvent(new CustomEvent("change",{bubbles:!0,composed:!0,detail:{...s}}))}_onEvent(e,t){this.onEvent&&this.onEvent(e,t),this.dispatchEvent(new CustomEvent("event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}_onSelector(e,t,i){this.onSelector&&this.onSelector(e,t,i),this.dispatchEvent(new CustomEvent("selector",{bubbles:!0,composed:!0,detail:{type:e,filter:t,callback:i}}))}lnkValue({fieldMap:e,fieldName:t,value:i}){if(void 0===this.values[this.field.name])return[this.data.current[e.source]?this.data.current[e.source].filter(t=>t.ref_id===this.values.id&&t[e.value]===this.field.name)[0]:this.data.dataset[e.source].filter(t=>t.ref_id===this.values.id&&t[e.value]===this.field.name)[0],!1];const l="id"===t&&""===i?null:i;return[this.data.current[e.source]?this.data.current[e.source].filter(t=>t[e.value]===l)[0]:this.data.dataset[e.source].filter(t=>t[e.value]===l)[0],!0]}getOppositeValue(e){return this.options.opposite&&parseFloat(e)<0?String(e).replace("-",""):this.options.opposite&&parseFloat(e)>0?`-${e}`:e}selectorInit({fieldMap:e,value:t,editName:i}){let l={value:this.values.value||"",filter:this.field.description||"",text:this.field.description||"",type:this.field.fieldtype,ntype:["transitem","transmovement","transpayment"].includes(this.field.fieldtype)?"trans":this.field.fieldtype,ttype:null,id:this.values.value||null,table:{name:"fieldvalue",fieldname:"value",id:this.field.id},fieldMap:e,editName:i};if(e){let i;if(l={...l,value:t||"",type:e.seltype,filter:String(l.text).split(" | ")[0],table:{name:e.table,fieldname:e.fieldname},ntype:e.lnktype,ttype:""!==t?""===e.transtype?this.values.transtype:e.transtype:l.ttype,id:""!==t?t:l.id},!0===e.extend||"extend"===e.table)i=this.data.current.extend;else{let t;t=void 0!==this.data.current[e.table]&&Array.isArray(this.data.current[e.table])?this.data.current[e.table].filter(e=>e.id===this.values.id):this.data.dataset[e.table].filter(e=>e.id===this.values.id),i=t[0]}void 0!==i?(l=void 0!==this.values[e.label_field]?{...l,text:this.values[e.label_field]||"",filter:this.values[e.label_field]||""}:void 0!==i[e.label_field]&&null!==i[e.label_field]?{...l,text:i[e.label_field],filter:i[e.label_field]}:{...l,text:"",filter:""},void 0!==i[e.fieldname]&&""===l.value&&null!==i[e.fieldname]&&(l={...l,ntype:e.lnktype,ttype:e.transtype,id:i[e.fieldname]}),"trans"===e.lnktype&&void 0!==i.transtype&&(l=void 0!==e.lnkid?{...l,id:i[e.lnkid]}:void 0!==i[e.fieldname]?{...l,id:i[e.fieldname]}:{...l,id:l.value})):l={...l,text:"",filter:""}}return this.selector=l,l}setSelector(e,t){let i={...this.selector,text:"",id:null,filter:t||""};if(e){const t=e.id.split("/");i={...i,text:e.label||e.item.lslabel},i={...i,id:parseInt(t[2],10),ttype:t[1]},"trans"===t[0]&&""!==t[1]&&e.trans_id&&(i={...i,id:e.trans_id})}i={...i,value:i.id||""},this.selector=i,this._onChange({value:i.id,item:e,event_type:"change",editName:i.editName,fieldMap:i.fieldMap})}editName(){return this.field.map?this.field.map.extend&&this.field.map.text?this.field.map.text:this.field.map.fieldname?this.field.map.fieldname:this.field.name:this.field.name}_onTextInput(e){this._onChange({value:e.target.value,event_type:"change",editName:this.editName(),fieldMap:this.field.map})}render(){let{datatype:e}=this.field,t=this.field.disabled||"readonly"===this.data.audit,d=this.field.name,u=this.values[this.field.name];const f=this.field.map||null,p=this.editName(),v=!("true"!==this.field.empty&&!0!==this.field.empty);switch("reportfield"!==this.field.rowtype&&"fieldvalue"!==this.field.rowtype||(u=this.values.value),null==u&&(u=this.field.default?this.field.default:""),"fieldvalue"===e&&(e=this.values.datatype),e){case"password":case"color":return i`<form-input 
          id="${this.id}" name="${d}" type="${e}" 
          value="${u||""}" label="${d}"
          ?full="${!0}"
          .style="${this.style}" 
          ?disabled="${t}"
          .onChange=${e=>this._onChange({value:e.value,event_type:"change",editName:p,fieldMap:f})}></form-input>`;case"date":case"datetime":if(f)if(f.extend)u=this.data.current.extend[f.text],d=f.text;else{const e=this.lnkValue({fieldMap:f,fieldName:d,value:u});void 0!==e[0]&&(u=e[0][f.text],t=e[1]?e[1]:t)}return i`<form-datetime id="${this.id}"
          name="${d}" label="${d}"
          .style="${this.style}" .isnull="${v}"
          type="${"datetime"===e?r.DATETIME:r.DATE}"
          .value="${u}"
          .onChange="${e=>{this._onChange({value:e.value,event_type:"change",editName:p,fieldMap:f})}}"
          ?disabled="${t}"></form-datetime>`;case"bool":case"flip":const m=t?"toggle-disabled":"";return[1,"1","true",!0].includes(u)?i`<div id="${this.id}"
            name="${d}" style="${a(this.style)}" 
            class="${`toggle toggle-on ${m}`}"
            @click="${t?null:()=>this._onChange({value:"fieldvalue_value"!==this.field.name&&0,event_type:"change",editName:p,fieldMap:f})}">
            <form-icon iconKey="ToggleOn" width=40 height=32.6 ></form-icon>
          </div>`:i`<div id="${this.id}"
          name="${d}" style="${a(this.style)}" 
          class="${`toggle toggle-off ${m}`}"
          @click="${t?null:()=>this._onChange({value:"fieldvalue_value"===this.field.name||1,event_type:"change",editName:p,fieldMap:f})}">
          <form-icon iconKey="ToggleOff" width=40 height=32.6 ></form-icon>
        </div>`;case"label":return null;case"select":this.field.extend&&(u=this.data.current.extend[this.field.name]||"");const $=[];return f?this.data.dataset[f.source].forEach(e=>{let t=e[f.text];void 0!==f.label&&(t=this.msg(`${f.label}_${t}`,{id:`${f.label}_${t}`})),$.push({value:String(e[f.value]),text:t})}):this.field.options.forEach(e=>{let t=e[1];this.msg(t,{id:t})&&(t=this.msg(t,{id:t})),void 0!==this.field.olabel&&(t=this.msg(`${this.field.olabel}_${e[1]}`,{id:`${this.field.olabel}_${e[1]}`})),$.push({value:String(e[0]),text:t})}),i`<form-select id="${this.id}" ?full="${!0}"
            name="${d}" label="${d}"
            .style="${this.style}"
            ?disabled="${t}" 
            .onChange=${e=>{const t=Number.isNaN(parseInt(e.value,10))?e.value:parseInt(e.value,10);this._onChange({value:t,event_type:"change",editName:p,fieldMap:f})}}
            .options=${$} 
            .isnull="${v}" value="${u}" ></form-select>`;case"valuelist":return i`<form-select id="${this.id}" ?full="${!0}"
          name="${d}" label="${d}"
          .style="${this.style}"
          ?disabled="${t}"
          .onChange=${e=>this._onChange({value:e.value,event_type:"change",editName:p,fieldMap:f})}
          .options=${this.field.description.map(e=>({value:e,text:e}))} 
          .isnull="${!1}" value="${u}" ></form-select>`;case"link":let g=this.values;const b=this.lnkValue({fieldMap:f,fieldName:d,value:u});void 0!==b[0]&&(g=b[0],b[0][f.text]&&(u=b[0][f.text]));let y=u;return void 0!==f.label_field&&void 0!==g[f.label_field]&&(y=g[f.label_field]),i`<div 
          name="${d}" style="${a(this.style)}" class="link" >
          <span id=${`link_${f.lnktype}_${d}`} class="link-text" 
            @click="${()=>this._onEvent(h.CHECK_EDITOR,[{ntype:f.lnktype,ttype:f.transtype,id:u},c.LOAD_EDITOR])}" >${y}</span>
        </div>`;case"selector":let{selector:x}=this;const w=[],_=f&&(!0===f.extend||"extend"===f.table)&&this.data.current.extend&&this.data.current.extend.seltype;return x?_&&x.type!==this.data.current.extend.seltype&&(x.text=this.data.current.extend[f.label_field],x.type=this.data.current.extend.seltype,x.filter=x.text,x.ntype=this.data.current.extend.seltype,x.ttype=this.data.current.extend.transtype,x.id=this.data.current.extend.ref_id):x=this.selectorInit({fieldMap:f,value:u,editName:p}),t||w.push(i`<div id="sel_show" class="cell search-col">
            <form-button id="${`sel_show_${d}`}" 
              label="${this.msg("",{id:"label_search"})}"
              .style="${{padding:"5px 8px 7px"}}"
              icon="Search" type="${n.BORDER}"
              @click=${()=>this._onSelector(x.type,x.filter,(...e)=>this.setSelector(...e))} 
            ></form-button>
          </div>`),v&&w.push(i`<div id="sel_delete" class="cell times-col">
            <form-button id="${`sel_delete_${d}`}" 
              label="${this.msg("",{id:"label_delete"})}"
              .style="${{padding:"5px 8px 7px"}}"
              ?disabled="${t}" icon="Times" type="${n.BORDER}"
              @click=${t?null:()=>this.setSelector()} ></form-button>
          </div>`),w.push(i`<div id="sel_text" class="link">
          ${""!==x.text?i`<span 
            id=${`sel_link_${d}`}
            class="link-text"
            @click="${()=>this._onEvent(h.CHECK_TRANSTYPE,[{ntype:x.ntype,ttype:x.ttype,id:x.id},c.LOAD_EDITOR])}" >${x.text}</span>`:null}
        </div>`),i`<div id="${this.id}" 
          name="${d}" style="${a(this.style)}"
          class="row full" >${w}</div>`;case"button":return i`<form-button id="${this.id}" 
          name="${d}" type="${n.BORDER}"
          .style="${{padding:"7px 8px",...this.style}}"
          ?disabled="${t}" 
          label="${this.field.title?this.field.title:l}" 
          ?full="${this.field.full}"
          ?autofocus="${this.field.focus||!1}"
          icon="${this.field.icon}"
          @click=${t?null:()=>this._onChange({value:d,item:{},event_type:"click",editName:p,fieldMap:f})} >${this.field.title?this.field.title:l}</form-button>`;case"percent":case"integer":case"float":if(f)if(f.extend)u=this.data.current.extend[f.text],d=f.text;else{const e=this.lnkValue({fieldMap:f,fieldName:d,value:u});void 0!==e[0]&&(u=e[0][f.text],t=e[1]?e[1]:t)}return""===u&&(u=0),void 0!==this.field.opposite&&(u=this.getOppositeValue(u)),i`<form-number id="${this.id}" name="${d}"
          ?integer="${!("float"===e)}" 
          ?full="${!0}" .style="${this.style}"
          value="${u||0}" 
          ?disabled="${t}" label="${d}"
          min="${s(this.field.min)}" 
          max="${s(this.field.max)}" 
          .onChange=${e=>{this._onChange({value:this.field.opposite?this.getOppositeValue(e.value):e.value,event_type:"change",editName:p,fieldMap:f}),this.event_type="change"}}
          .onBlur=${e=>{this._onChange({value:this.field.opposite?parseFloat(this.getOppositeValue(e.value)):parseFloat(e.value),event_type:"change"===this.event_type?"blur":null,editName:p,fieldMap:f}),this.event_type="blur"}}
        ></form-number>`;default:if(f)if(f.extend)u=this.data.current.extend[f.text],d=f.text;else{const e=this.lnkValue({fieldMap:f,fieldName:d,value:u});void 0!==e[0]&&(u=e[0][f.text],t=e[1]?e[1]:t,void 0!==f.label&&(u=this.msg(`${f.label}_${u}`,{id:`${f.label}_${u}`})))}return"notes"===e||"text"===e?i`<textarea id="${this.id}" name="${d}"
            class=${"full"} label="${d}" style="${a(this.style)}"
            rows=${s(this.field.rows?this.field.rows:void 0)}
            @input="${this._onTextInput}"
            .value="${u||""}"
            ?disabled="${t}" ></textarea>`:i`<form-input 
          id="${this.id}" .id="${`${this.id}_input`}" name="${d}" 
          type="${o.TEXT}" 
          value="${u||""}" label="${d}"
          ?full="${!0}"
          .style="${this.style}" 
          maxlength="${s(this.field.length?this.field.length:void 0)}"
          size="${s(this.field.length?this.field.length:void 0)}"
          ?disabled="${t}"
          .onChange=${e=>this._onChange({value:e.value,event_type:"change",editName:p,fieldMap:f})}></form-input>`}}});const m=e`
:host {
	font-family: var(--font-family);
	font-size: var(--font-size);
	color: var(--text-1);
	fill: var(--text-1);
}
div {
  box-sizing: border-box;
}
.half { 
  float:left; 
  width:100% 
}
.s12 { 
  float:left; 
  width:99.99999%; 
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
.bold { 
  font-weight: bold; 
}
.padding-small { 
  padding: 4px 8px; 
}
.padding-tiny { 
  padding: 2px 4px; 
}
.container-small { 
  padding: 0px 8px; 
}
.container-row {
  display: table;
  width: 100%;
  padding: 8px;
  border-bottom: 1px solid rgba(var(--neutral-1), 0.2);
}
.label-row {
  color: rgb(var(--functional-blue));
}
.report-field {
  cursor: pointer;
}
.toggle-on {
  fill: rgb(var(--functional-green))!important;
  color: rgb(var(--functional-green))!important;
}
.toggle {
  font-family: var(--font-family);
  font-size: var(--font-size);
  text-align: center;
  cursor: pointer;
  width: 100%;
  border-radius: 3px;
  fill: var(--text-1);
  background-color: rgba(var(--base-4), 1);
  border: 1px solid rgba(var(--neutral-1), 0.2);
}
.toggle:hover:not(:disabled) {
  fill: rgb(var(--functional-green))!important;
  color: rgb(var(--functional-green))!important;
}
.toggle-off {
  opacity: 0.5;
}
.toggle-disabled {
  opacity: 0.5;
}
.info {
  background-color: rgba(var(--functional-blue),0.2)!important;
}
.leftbar{
  border-left: 8px solid rgba(var(--functional-yellow),0.8)!important;
  padding: 8px;
  font-style: italic;
  font-size: 13px;
}
.center { 
  text-align: center; 
}
.align-right { 
  text-align: right; 
}
textarea {
  font-family: var(--font-family);
  font-size: 12px;
  border-radius: 3px;
  overflow: auto;
  padding: 8px;
  color: var(--text-1);
  background-color: rgba(var(--base-4), 1);
  border: 1px solid rgba(var(--neutral-1), 0.2);
  box-sizing: border-box;
}
textarea:focus, textarea:hover:not(:disabled) {
  color: rgb(var(--functional-green))!important;
}
textarea:disabled {
  opacity: 0.5;
}
textarea::placeholder, textarea::-ms-input-placeholder {
  opacity: 0.5;
}
.field-cell {
  width: 150px;
}
.fieldvalue-delete {
  cursor: pointer;
}

.fieldvalue-delete:hover {
  fill: rgb(var(--functional-red));
}
@media (max-width:600px){
  .container-row {
    padding: 4px;
  }
  .container-small { 
    padding: 0px 4px; 
  }
  .padding-small { 
    padding: 2px 4px; 
  }
  .padding-tiny { 
    padding: 1px 2px; 
  }
  .hide-small { 
    display: none!important; 
  }
}
@media (min-width:601px) {
  .half { 
    width:49.99999% 
  }
  .m3 { 
    float:left; 
    width:24.99999%; 
  }
  .m4 { 
    float:left; 
    width: 33.33333%; 
  }
  .m6 { 
    float:left; 
    width: 49.99999%; 
  }
}
@media (max-width:992px) and (min-width:601px){
  .hide-medium { 
    display: none!important; 
  }
}
@media (min-width:993px) {
  .hide-large { 
    display: none!important; 
  }
  .l3 { 
    float:left; 
    width: 24.99999%; 
  }
  .l4 { 
    float:left; 
    width: 33.33333%; 
  }
  .l6 {
    float:left; 
    width: 49.99999%; 
  }
}
`;customElements.define("form-row",class extends t{constructor(){super(),this.id=Math.random().toString(36).slice(2),this.row={},this.values={},this.options={},this.data={dataset:{},current:{},audit:"all"},this.style={},this.msg=e=>e}static get properties(){return{id:{type:String},row:{type:Object},values:{type:Object},options:{type:Object},data:{type:Object},style:{type:Object}}}static get styles(){return[m]}_onEdit(e){this.onEdit&&this.onEdit(e),this.dispatchEvent(new CustomEvent("edit",{bubbles:!0,composed:!0,detail:{...e}}))}_onTextInput(e){this._onEdit({id:this.row.id,name:this.row.name,value:e.target.value})}imgValue(){let e=this.values[this.row.name]||"";return""!==e&&null!==e&&"data:image"!==e.toString().substr(0,10)&&void 0!==this.data.dataset[e]&&(e=this.data.dataset[e]),e}flipItem(){const{id:e,name:t,datatype:l,info:s}=this.row,d=void 0!==this.values[t],n=i`<div id="${`checkbox_${t}`}"
      class="report-field ${d?"toggle-on":"toggle-off"}"
      @click="${()=>this._onEdit({id:e,selected:!0,datatype:l,defvalue:this.row.default,name:t,value:!d,extend:!1})}">
      ${d?i`<form-icon iconKey="ToggleOn" width=40 height=32.6 ></form-icon>`:i`<form-icon iconKey="ToggleOff" width=40 height=32.6 ></form-icon>`}
      <form-label value="${t}" class="bold padding-tiny ${d?"toggle-on":""} " ></form-label>
    </div>`;switch(l){case"text":return i`<div id="${this.id}"
          style="${a(this.style)}" class="container-row">
          <div class="row full">
            <div class="cell padding-small" >
              ${n}
            </div>
          </div>
          ${d?i`<div class="row full"><div class="cell padding-small" >
            <form-field id="${`field_${t}`}"
              .field=${this.row} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div></div>`:null}
          ${s?i`<div class="row full padding-small">
            <div class="cell padding-small info leftbar" >
              ${s}
            </div>
          </div>`:null}
        </div>`;case"image":return i`<div id="${this.id}"
            style="${a(this.style)}" class="container-row">
            <div class="row full">
              <div class="cell padding-small" >
                ${n}
              </div>
              ${d?i`<div class="cell padding-small" >
              <form-input 
                id="${`file_${t}`}" 
                type="${o.FILE}" ?full="${!0}"
                label="${this.labelAdd}"
                .style="${{"font-size":"12px"}}"
                .onChange=${i=>this._onEdit({id:e,file:!0,name:t,value:i.value,extend:!1})}></form-input>
              </div>`:null}
            </div>
            ${d?i`<div class="row full"><div class="cell padding-small" >
              <textarea id="${`input_${t}`}"
                class=${"full"} rows=5 .value="${this.imgValue()}"
                @input="${this._onTextInput}" ></textarea>
              <div class="full padding-normal center" >
                <img src="${this.imgValue()}" alt="" />
              </div>
            </div></div>`:null}
            ${s?i`<div class="row full padding-small">
              <div class="cell padding-small info leftbar" >
                ${s}
              </div>
            </div>`:null}
          </div>`;case"checklist":const l=this.values[t]||"",r=[];return this.row.values.forEach((a,s)=>{const d=a.split("|"),o=l.indexOf(d[0])>-1;r.push(i`<div id="${`checklist_${t}_${s}`}"
            key={index}
            class="cell padding-small report-field"
            @click=${()=>this._onEdit({id:e,checklist:!0,name:t,checked:!o,value:d[0],extend:!1})}>
            <form-label 
              value="${d[1]}" class="bold ${o?"toggle-on":""}"
              leftIcon="${o?"CheckSquare":"SquareEmpty"}" ></form-label>
          </div>`)}),i`<div id="${this.id}"
            style="${a(this.style)}" class="container-row">
          <div class="row full">
            <div class="cell padding-small" >
              ${n}
            </div>
          </div>
          ${d?i`<div class="row full padding-small">
            <div class="cell padding-small toggle" >
              ${r}
            </div>
          </div>`:null}
          ${s?i`<div class="row full padding-small">
            <div class="cell padding-small info leftbar" >
              ${s}
            </div>
          </div>`:null}
        </div>`;default:return i`<div id="${this.id}"
          style="${a(this.style)}" class="container-row">
          <div class="row full">
            <div class="cell padding-small half" >
              ${n}
            </div>
            ${d?i`<div class="cell padding-small half" >
              <form-field id="${`field_${t}`}"
                .field=${this.row} .values=${this.values} .options=${this.options} .data=${this.data}
                .msg=${this.msg} .onEdit=${this.onEdit} 
                .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
            </div>`:null}
          </div>
          ${s?i`<div class="row full padding-small">
            <div class="cell padding-small info leftbar" >
              ${s}
            </div>
          </div>`:null}
        </div>`}}render(){const{id:e,rowtype:t,label:l,columns:s,name:d,disabled:n,notes:r,selected:h,empty:c,datatype:u,info:f}=this.row;switch(t){case"label":return i`<div id="${this.id}" style="${a(this.style)}" 
          class="container-row label-row">
          <div class="cell padding-small" >${this.values[d]||l}</div>
        </div>`;case"flip":return this.flipItem();case"field":return i`<div id="${this.id}"
          style="${a(this.style)}" class="container-row">
          <div class="cell padding-small hide-small field-cell" >
            <form-label value="${l}" class="bold" ></form-label>
          </div>
          <div class="cell padding-small" >
            <div class="hide-medium hide-large" >
              <form-label value="${l}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${d}`}"
              .field=${this.row} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
        </div>`;case"reportfield":return i`<div id="${this.id}"
          style="${a(this.style)}" class="cell padding-small s12 m6 l4">
          <div id="${`cb_${d}`}"
            class=${"padding-small "+("false"!==c?"report-field":"")} 
            @click="${()=>{"false"!==c&&this._onEdit({id:e,name:"selected",value:!h,extend:!1})}}">
            <form-label 
              value="${l}" class="bold"
              leftIcon="${h?"CheckSquare":"SquareEmpty"}" ></form-label>
          </div>
          <form-field id="${`field_${d}`}"
            .field=${this.row} .values=${this.values} .options=${this.options} .data=${this.data}
            .msg=${this.msg} .onEdit=${this.onEdit} 
            .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
        </div>`;case"fieldvalue":return i`<div id="${this.id}"
          style="${a(this.style)}" class="container-row">
          <div class="row full">
            <div class="cell container-small">
              <form-label value="${l}" class="bold" ></form-label>
            </div>
            <div class="cell align-right container-small" >
              <span id=${`delete_${this.row.fieldname}`}
                class="fieldvalue-delete" 
                @click="${()=>this._onEdit({id:e,name:"fieldvalue_deleted"})}">
                <form-icon iconKey="Times" ></form-icon>
              </span>
            </div>
          </div>
          <div class="row full">
            <div class="cell padding-small s12 m6 l6" >
              <form-field id="${`field_${this.row.fieldname}`}"
                .field=${this.row} .values=${this.values} .options=${this.options} .data=${this.data}
                .msg=${this.msg} .onEdit=${this.onEdit} 
                .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
            </div>
            <div class="cell padding-small s12 m6 l6" >
              <form-input 
                id="${`notes_${this.row.fieldname}`}" type="${o.TEXT}"
                label="${this.msg("",{id:"fnote_view"})}"
                name="fieldvalue_notes" ?full="${!0}" value="${r}"
                ?disabled=${n||"readonly"===this.data.audit}
                .onChange=${t=>this._onEdit({id:e,name:"fieldvalue_notes",value:t.value})}></form-input>
            </div>
          </div>
        </div>`;case"col2":return i`<div id="${this.id}"
          style="${a(this.style)}" class="container-row">
          <div class="cell padding-small s12 m6 l6" >
            <div>
              <form-label value="${s[0].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${s[0].name}`}"
              .field=${s[0]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
          <div class="cell padding-small s12 m6 l6" >
            <div>
              <form-label value="${s[1].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${s[1].name}`}"
              .field=${s[1]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
        </div>`;case"col3":return i`<div id="${this.id}"
          style="${a(this.style)}" class="container-row">
          <div class="cell padding-small s12 m4 l4" >
            <div>
              <form-label value="${s[0].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${s[0].name}`}"
              .field=${s[0]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
          <div class="cell padding-small s12 m4 l4" >
            <div>
              <form-label value="${s[1].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${s[1].name}`}"
              .field=${s[1]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
          <div class="cell padding-small s12 m4 l4" >
            <div>
              <form-label value="${s[2].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${s[2].name}`}"
              .field=${s[2]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
        </div>`;case"col4":return i`<div id="${this.id}"
          style="${a(this.style)}" class="container-row">
          <div class="cell padding-small s12 m3 l3" >
            <div>
              <form-label value="${s[0].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${s[0].name}`}"
              .field=${s[0]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
          <div class="cell padding-small s12 m3 l3" >
            <div>
              <form-label value="${s[1].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${s[1].name}`}"
              .field=${s[1]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
          <div class="cell padding-small s12 m3 l3" >
            <div>
              <form-label value="${s[2].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${s[2].name}`}"
              .field=${s[2]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
          <div class="cell padding-small s12 m3 l3" >
            <div>
              <form-label value="${s[3].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${s[3].name}`}"
              .field=${s[3]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
        </div>`}return null}});
