import{i as e,r as t,x as i}from"./module-FRmUNWHB.js";import{o as n,a}from"./module-CC7fmSsS.js";import{B as s}from"./main-BEo7670f.js";import"./module-D-7nCm3D.js";const o=e`
:host {
	font-family: var(--font-family);
	font-size: var(--font-size);
	color: var(--text-1);
	fill: var(--text-1);
}
.row {
  display: table;
}
.cell {
  display: table-cell;
  vertical-align: middle;
}
.padding-small { 
  padding: 4px 8px; 
}
@media (max-width:600px){
  .padding-small { 
    padding: 2px 4px; 
  }
  .hide-small { 
    display: none!important; 
  }
}
`,l=e`
input[type=number] {
  font-family: var(--font-family);
  font-size: var(--font-size);
  text-align: right;
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
input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
  /* display: none; <- Crashes Chrome on hover */
  -webkit-appearance: none;
  margin: 0; /* <-- Apparently some margin are still there even though it's hidden */
}
input[type=number] {
  -moz-appearance:textfield; /* Firefox */
}
.full{
  width: 100%;
}
`;customElements.define("form-number",class extends t{constructor(){super(),this.id=Math.random().toString(36).slice(2),this.name=void 0,this.value=0,this.integer=!1,this.label=void 0,this.max=void 0,this.min=void 0,this.disabled=!1,this.readonly=!1,this.autofocus=!1,this.full=!1,this.style={}}static get properties(){return{id:{type:String},name:{type:String,reflect:!0},value:{type:Number,reflect:!0},integer:{type:Boolean},label:{type:String},max:{type:Number},min:{type:Number},disabled:{type:Boolean,reflect:!0},readonly:{type:Boolean,reflect:!0},autofocus:{type:Boolean,reflect:!0},full:{type:Boolean},style:{type:Object}}}_onInput(e){const t=e=>{e!==this.value&&(this.onChange&&this.onChange({value:e,old:this.value}),this.dispatchEvent(new CustomEvent("change",{bubbles:!0,composed:!0,detail:{value:e,old:this.value}})),this.value=e),this._input.value!==String(e)&&(this._input.value=String(e))};if(e.target.valueAsNumber!==this.value){let i=e.target.valueAsNumber;Number.isNaN(i)&&(i=0),void 0!==this.min&&i<this.min&&(i=this.min),void 0!==this.max&&i>this.max&&(i=this.max),this.integer&&(i=Math.floor(i)),t(i)}}_onBlur(){this._input.value=String(this.value),this.onBlur&&this.onBlur({value:this.value})}_onKeyEvent(e){const t=()=>{this.onEnter&&(this.onEnter({value:this.value}),this.dispatchEvent(new CustomEvent("enter",{bubbles:!0,composed:!0,detail:{value:this.value}})))};"keydown"!==e.type&&"keypress"!==e.type||e.stopPropagation(),"keydown"===e.type&&13===e.keyCode&&(e.preventDefault(),t()),this.readonly||"keypress"!==e.type||13!==e.keyCode||t()}firstUpdated(){this._input=this.renderRoot.querySelector("input")}render(){return i`<input 
      id="${this.id}"
      name="${n(this.name)}"
      type="number"
      .value="${String(this.value)}"
      ?disabled="${this.disabled}"
      ?readonly="${this.readonly}"
      ?autofocus="${this.autofocus}"
      min="${n(this.min)}"
      max="${n(this.max)}"
      aria-label="${n(this.label)}"
      class="${this.full?"full":""}"
      style="${a(this.style)}"
      @input=${this._onInput}
      @blur=${this._onBlur}
      @keydown=${this._onKeyEvent}
      @keypress=${this._onKeyEvent}
    >`}static get styles(){return[l]}});customElements.define("form-pagination",class extends t{constructor(){super(),this.id=Math.random().toString(36).slice(2),this.name=void 0,this.pageIndex=0,this.pageSize=5,this.pageCount=0,this.canPreviousPage=!1,this.canNextPage=!1,this.hidePageSize=!1,this.style={}}static get properties(){return{id:{type:String},name:{type:String},pageIndex:{type:Number},pageSize:{type:Number},pageCount:{type:Number},canPreviousPage:{type:Boolean},canNextPage:{type:Boolean},hidePageSize:{type:Boolean},style:{type:Object}}}static get styles(){return[o]}_onEvent(e,t,i){i&&(this.onEvent&&this.onEvent(e,t),this.dispatchEvent(new CustomEvent("pagination",{bubbles:!0,composed:!0,detail:{key:e,value:t}})))}render(){return i`<div id="${this.id}" name="${n(this.name)}"
      class="row" style="${a(this.style)}"
    >
      <div class="cell padding-small" >
        <form-button id="pagination_btn_first" 
          .style="${{padding:"6px 6px 7px","font-size":"15px",margin:"1px 0 2px"}}"
          ?disabled="${!this.canPreviousPage}" label="1"
          @click=${()=>this._onEvent("gotoPage",1,this.canPreviousPage)} type="${s.BORDER}" >1</form-button>
        <form-button id="pagination_btn_previous" 
          .style="${{padding:"5px 6px 8px","font-size":"15px",margin:"1px 0 2px"}}"
          ?disabled="${!this.canPreviousPage}" label="&#10094;"
          @click=${()=>this._onEvent("previousPage",this.pageIndex-1,this.canPreviousPage)} type="${s.BORDER}" >&#10094;</form-button>
      </div>
      <div class="cell" >
        <form-number id="pagination_input_goto" ?integer="${!0}" 
          .style="${{padding:"7px",width:"60px","font-weight":"bold"}}"
          value="${this.pageIndex}" ?disabled="${0===this.pageCount}"
          min="1" max="${this.pageCount}" label="Page"
          .onChange=${e=>this._onEvent("gotoPage",e.value,this.pageCount>0)}
        ></form-number>
      </div>
      <div class="cell padding-small" >
        <form-button id="pagination_btn_next" 
          .style="${{padding:"5px 6px 8px","font-size":"15px",margin:"1px 0 2px"}}"
          ?disabled="${!this.canNextPage}" label="&#10095;"
          @click=${()=>this._onEvent("nextPage",this.pageIndex+1,this.canNextPage)} type="${s.BORDER}" >&#10095;</form-button>
        <form-button id="pagination_btn_last" 
          .style="${{padding:"6px 6px 7px","font-size":"15px",margin:"1px 0 2px"}}"
          ?disabled="${!this.canNextPage}" label="${this.pageCount}"
          @click=${()=>this._onEvent("gotoPage",this.pageCount,this.canNextPage)} type="${s.BORDER}" >${this.pageCount}</form-button>
      </div>
      ${this.hidePageSize?"":i`<div class="cell padding-small hide-small" >
        <form-select id="pagination_page_size"
          label="Size"
          .style="${{padding:"7px"}}" ?disabled="${0===this.pageCount}"
          .onChange=${e=>this._onEvent("setPageSize",Number(e.value),this.pageCount>0)}
          .options=${[5,10,20,50,100].map((e=>({value:String(e),text:String(e)})))} 
          .isnull="${!1}" value="${this.pageSize}" >
        </form-select>
      </div>`}
    </div>`}});
