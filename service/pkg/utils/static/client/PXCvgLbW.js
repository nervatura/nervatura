import{i as e,s as i,x as t,A as a}from"./6kG9gGCM.js";import"./zCqrk7Sg.js";import"./AnyrskTL.js";import"./blk_cvvB.js";import{M as l,I as o,P as s}from"./MSqLDGvE.js";const r=e`
@keyframes animatezoom{from{transform:scale(0)} to{transform:scale(1)}}
:host {
  font-family: var(--font-family);
  font-size: var(--font-size);
}
div {
  box-sizing: border-box;
}
.row {
  display: table;
}
.row::before {
  content: "";
  display: table;
  clear: both;
}
.row::after {
  content: "";
  display: table;
  clear: both;
}
.cell {
  display: table-cell;
  vertical-align: middle;
}
.full { 
  width: 100%; 
}
.modal {
  z-index: 10;
  position: fixed;
  left: 0;
  right: 0;
  top: 0;
  height: 100vh;
  overflow: auto;
  background-color: rgba(25, 25, 25, 0.7);
  padding: 30px 5px;
}
.dialog {
  border-radius: 4px;
  box-shadow:0 4px 10px 0 rgba(0,0,0, 0.2),0 4px 20px 0 rgba(0,0,0, 0.19);
  margin: 0 auto;
  animation: animatezoom 0.6s;
  width: 100%;
  max-width: 600px;
  min-width: 280px;
}
.panel {
  background-color: rgba(var(--base-2), 1);
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  box-shadow:0 2px 5px 0 rgba(0,0,0,0.16),0 2px 10px 0 rgba(0,0,0,0.12);
  border-radius: 4px;
  max-width: 800px;
  margin: auto;
}
.margin0 { 
  margin: 0!important; 
}
.panel-title {
  display: table;
  border-radius: 4px;
  font-weight: bold;
  padding: 8px 16px;
  width: 100%;
  border-bottom: 0.5px solid rgba(var(--neutral-1), 0.2);
  background-color: rgb(var(--accent-1));
}
.title-cell {
  color: rgba(var(--accent-1c), 0.85);
  fill: rgba(var(--accent-1c), 0.85);
}
.align-right { 
  text-align: right; 
}
.close-icon {
  fill: rgba(var(--accent-1c), 0.85);
  cursor: pointer;
}
.close-icon form-icon:hover {
  fill: rgb(var(--functional-red));
}
.section {
  padding: 16px 0; 
}
.section-row {
  display: table;
  width: 100%;
  padding: 0 16px 8px;
}
.search-col {
  width: 100px;
  padding-left: 16px;
}
@media (max-width:600px){
  .section-row { 
    padding: 0 8px 8px; 
  }
}
@media only screen and (min-width: 601px){
  .dialog {
    min-width: 400px;
  }
}
`;customElements.define("modal-selector",class extends i{constructor(){super(),this.msg=e=>e,this.isModal=!1,this.view="",this.columns=[],this.result=[],this.filter="",this.selectorPage=5,this.paginationPage=10,this.currentPage=1}static get properties(){return{isModal:{type:Boolean},view:{type:String},columns:{type:Array},result:{type:Array},filter:{type:String},selectorPage:{type:Number},paginationPage:{type:Number},currentPage:{type:Number}}}static get styles(){return[r]}_onModalEvent(e,i){this.onEvent&&this.onEvent.onModalEvent&&this.onEvent.onModalEvent({key:e,data:i}),this.dispatchEvent(new CustomEvent("modal_event",{bubbles:!0,composed:!0,detail:{key:e,data:i}}))}_onFilterChange(e){this.filter=e.value,this.dispatchEvent(new CustomEvent("filter_change",{bubbles:!0,composed:!0,detail:{...e}}))}selectorView(){let e={view:{columnDef:{id:"view",Header:"",headerStyle:{},Cell:({row:e})=>1===e.deleted?t`<form-icon iconKey="ExclamationTriangle" .style=${{fill:"rgb(var(--functional-yellow))"}} ></form-icon>`:t`<form-icon iconKey="CaretRight" width=9 height=24 ></form-icon>`,cellStyle:{width:"25px",padding:"7px 2px 3px 8px"}}}};return this.columns.forEach((i=>{e={...e,[i[0]]:{fieldtype:"string",label:this.msg(`${this.view}_${i[0]}`,{id:`${this.view}_${i[0]}`})}}})),t`<div class="panel ${this.isModal?"":"margin0"}">
      <div class="panel-title">
        <div class="cell" >
          <form-label 
            value=${this.isModal?this.msg(`search_${this.view}`,{id:`search_${this.view}`}):`${this.msg("Quick Search",{id:"quick_search"})}: ${this.msg(`search_${this.view}`,{id:`search_${this.view}`})}`}
            class="title-cell" leftIcon=${this.isModal?"Search":""} >
          </form-label>
        </div>
        ${this.isModal?t`<div class="cell align-right" >
          <span id=${"closeIcon"} class="close-icon" 
            @click="${()=>this._onModalEvent(l.CANCEL,{})}">
            <form-icon iconKey="Times" ></form-icon>
          </span>
        </div>`:a}
      </div>
      <div class="section" >
        <div class="section-row" >
          <div class="cell" >
            <form-input id="selector_filter" type="${o.TEXT}" 
              label="${this.msg("",{id:"placeholder_search"})}"
              placeholder="${this.msg("",{id:"placeholder_search"})}"
              value="${this.filter}" ?full="${!0}" ?autofocus="${!0}"
              .onChange=${e=>this._onFilterChange({value:e.value,old:this.filter})}
              .onEnter=${()=>this._onModalEvent(l.SEARCH,{value:this.filter})}
            ></form-input>
          </div>
          <div class="cell search-col" >
            <form-button id="selector_btn_search" icon="Search"
              label="${this.msg("",{id:"label_search"})}"
              @click=${()=>this._onModalEvent(l.SEARCH,{value:this.filter})}
            >${this.msg("",{id:"label_search"})}
            </form-button>
          </div>
        </div>
        <div class="section-row" >
        <form-table id="selector_result"
          .rows="${this.result}"
          .fields="${e}"
          pagination="${s.TOP}"
          currentPage="${this.currentPage}"
          pageSize="${this.isModal?this.selectorPage:this.paginationPage}"
          ?tableFilter="${!1}"
          ?hidePaginatonSize="${!0}"
          .onRowSelected=${e=>this._onModalEvent(l.SELECTED,{value:e,filter:this.filter})}
          .onCurrentPage=${e=>this._onModalEvent(l.CURRENT_PAGE,{value:e})}
        ></form-table>
        </div>
      </div>
    </div>`}render(){return this.isModal?t`<div class="modal">
        <div class="dialog">
          ${this.selectorView()}
        </div>
      </div>`:this.selectorView()}});
