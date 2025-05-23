import{a as e,i as t,E as i,x as o}from"./module-D9Io0VgM.js";import{o as l,a as s}from"./module-DEu6bRyP.js";import{e as n,I as a}from"./main-D0db8xQQ.js";const r=e`
button {
  font-family: var(--font-family);
  font-size: var(--font-size);
  color: var(--text-1);
  fill: var(--text-1);
  border-color: rgba(var(--neutral-1), 0.2);
  background-color: rgba(var(--neutral-1), 0.1);
  font-weight: bold;
  display: inline-block;
  padding: 8px 16px;
  vertical-align: middle;
  overflow: hidden;
  text-decoration: none;
  cursor:pointer;
  white-space:nowrap;
  -webkit-touch-callout:none;
  -webkit-user-select:none;
  -khtml-user-select:none;
  -moz-user-select:none;
  -ms-user-select:none;
  user-select:none;
  border-radius: 3px;
  border-width: 1px;
  transition: 0.1s all ease-out;
  box-sizing: border-box;
  vertical-align: middle;
}
button[button-type='primary'] {
  color: rgba(var(--accent-1c), 0.85);
  fill: rgba(var(--accent-1c), 0.85);
  border-color: rgba(var(--neutral-1), 0.2);
  background-color: rgb(var(--accent-1));
}
button[button-type='border'] {
  background: none;
  border-width: 1px;
  border-style: solid;
  border-color: rgba(var(--neutral-1), 0.3);
}
button[disabled] {
  pointer-events: none;
  opacity: 0.3;
}
button:not(:active):hover {
  background-color: rgba(var(--neutral-1), 0.15);
}
button[button-type='primary']:not(:active):hover {
  background-color: rgb(var(--accent-1b));
}
button[button-type='border']:not(:active):hover {
  color: rgb(var(--accent-1c));
  fill: rgb(var(--accent-1c));
  border-color: rgb(var(--accent-1b));
}
button:disabled:hover {
  box-shadow:none;
}
button:focus, button:hover {
  color: rgb(var(--functional-green));
  fill: rgb(var(--functional-green));
}
.selected {
  color: rgb(var(--functional-yellow))!important;
  fill: rgb(var(--functional-yellow))!important;
}
button[button-type='border'].selected {
  border-color: rgb(var(--functional-yellow))!important;
}
.right {
  float: right;
}
.badge {
  font-weight: bold;
  font-size: 14px;
  color: rgb(var(--accent-1));
  fill: rgb(var(--accent-1));
  background-color: rgb(var(--accent-1c));
  display: inline-block;
  padding: 1px 7px;
  text-align: center;
  border-radius: 50%;
}
.selected-badge {
  background-color: rgb(var(--functional-yellow));
  color: rgb(var(--accent-1c));
}
.small {
  font-size: 12px;
  padding: 6px 8px!important;
}
.full{
  width: 100%;
}
.left {
  text-align: left;
}
.center {
  text-align: center
}
.right {
  text-align: right
}
@media (max-width:600px){
  .hidelabel slot { 
    display: none;
  }
  .hidelabel {
    padding: 8px 12px;
  }
}
`;customElements.define("form-button",class extends t{constructor(){super(),this.id=Math.random().toString(36).slice(2),this.type=void 0,this.name=void 0,this.align=n.CENTER,this.icon=void 0,this.label="",this.disabled=!1,this.autofocus=!1,this.full=!1,this.small=!1,this.selected=!1,this.hidelabel=!1,this.badge=void 0,this.style={}}static get properties(){return{id:{type:String},name:{type:String,reflect:!0},type:{type:String},align:{type:String},label:{type:String},icon:{type:String},disabled:{type:Boolean,reflect:!0},autofocus:{type:Boolean,reflect:!0},full:{type:Boolean},small:{type:Boolean},selected:{type:Boolean},hidelabel:{type:Boolean},badge:{type:Number},style:{type:Object}}}_onClick(e){e.stopPropagation(),this.disabled||(this.onClick&&this.onClick(e),this.dispatchEvent(new CustomEvent("click",{bubbles:!0,composed:!0,detail:{id:this.id}})))}_onKeyEvent(e){const t=()=>{this.onEnter&&(this.onEnter({value:this.value}),this.dispatchEvent(new CustomEvent("enter",{bubbles:!0,composed:!0,detail:{id:this.id}})))};"keydown"!==e.type&&"keypress"!==e.type||e.stopPropagation(),"keydown"===e.type&&13===e.keyCode&&(e.preventDefault(),t()),this.readonly||"keypress"!==e.type||13!==e.keyCode||t()}render(){return o`<button 
      id="${this.id}"
      name="${l(this.name)}"
      button-type="${l(this.type)}"
      ?disabled="${this.disabled}"
      ?autofocus="${this.autofocus}"
      aria-label="${l(this.label)}"
      title="${l(this.label)}"
      class=${`${["small","full","selected","hidelabel"].filter((e=>this[e])).join(" ")} ${this.align}`}
      style="${s(this.style)}"
      @click=${this._onClick}
      @keydown=${this._onKeyEvent}
      @keypress=${this._onKeyEvent}>
        ${this.icon&&this.align!==n.RIGHT?o`<form-icon iconKey="${this.icon}" width=20 ></form-icon>`:i}
        <slot id="value"></slot>
        ${this.icon&&this.align===n.RIGHT?o`<form-icon iconKey="${this.icon}" width=20 ></form-icon>`:i}
        ${this.badge?o`<span class="right" ><span class="${"badge "+(this.selected?"selected-badge":"")}" >${this.badge}</span></span>`:i}
      </button>`}static get styles(){return[r]}});const c=e`
select {
  font-family: var(--font-family);
  font-size: 13px;
  border-radius: 3px;
  padding: 8px;
  display: block;
  color: var(--text-1);
  background-color: rgba(var(--base-4), 1);
  border: 1px solid rgba(var(--neutral-1), 0.2);
  box-sizing: border-box;
}
select:disabled {
  opacity: 0.5;
}
select:focus, select:hover:not(:disabled) {
  color: rgb(var(--functional-green))!important;
}
option {
  font-size: 14px;
  border: none;
}
option:disabled {
  opacity: 0.5;
}
.full{
  width: 100%;
}
`;customElements.define("form-select",class extends t{constructor(){super(),this.id=Math.random().toString(36).slice(2),this.value="",this.name=void 0,this.options=[],this.isnull=!0,this.label="",this.disabled=!1,this.autofocus=!1,this.full=!1,this.style={}}static get properties(){return{id:{type:String},name:{type:String,reflect:!0},value:{type:String,reflect:!0},options:{type:Array},isnull:{type:Boolean},label:{type:String},disabled:{type:Boolean,reflect:!0},autofocus:{type:Boolean,reflect:!0},full:{type:Boolean},style:{type:Object}}}_onInput(e){e.target.value!==this.value&&(this.onChange&&this.onChange({value:e.target.value,old:this.value}),this.dispatchEvent(new CustomEvent("change",{bubbles:!0,composed:!0,detail:{value:e.target.value,old:this.value}})),this.value=e.target.value),this._select.value!==e.target.value&&(this._select.value=e.target.value)}_onKeyEvent(e){const t=()=>{this.onEnter&&(this.onEnter({value:this.value}),this.dispatchEvent(new CustomEvent("enter",{bubbles:!0,composed:!0,detail:{value:this.value}})))};"keydown"!==e.type&&"keypress"!==e.type||e.stopPropagation(),"keydown"===e.type&&13===e.keyCode&&(e.preventDefault(),t()),this.readonly||"keypress"!==e.type||13!==e.keyCode||t()}firstUpdated(){this._select=this.renderRoot.querySelector("select")}render(){const e=this.options.map(((e,t)=>o`<option
      key=${t} value=${e.value} 
      ?selected=${e.value===this.value} >${e.text}</option>`));return this.isnull&&e.unshift(o`<option
        ?selected=${""===this.value}
        key="-1" value="" ></option>`),o`<select 
      id="${this.id}"
      name="${l(this.name)}"
      .value="${this.value}"
      ?disabled="${this.disabled}"
      ?autofocus="${this.autofocus}"
      aria-label="${l(this.label)}"
      class="${this.full?"full":""}"
      style="${s(this.style)}"
      @input=${this._onInput}
      @keydown=${this._onKeyEvent}
      @keypress=${this._onKeyEvent}
      >${e}</select>`}static get styles(){return[c]}});const d=e`
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
input[type=color]{
  height: 34px;
  border: none;
  padding: 0px;
  cursor: pointer;
}
input::placeholder, input::-ms-input-placeholder {
  opacity: 0.5;
}
input:disabled {
  opacity: 0.5;
}
input:focus, input:hover:not(:disabled) {
  color: rgb(var(--functional-green))!important;
}
.full{
  width: 100%;
}
`;customElements.define("form-input",class extends t{constructor(){super(),this.id=Math.random().toString(36).slice(2),this.type=a.TEXT,this.value="",this.name=void 0,this.placeholder=void 0,this.label="",this.disabled=!1,this.readonly=!1,this.autofocus=!1,this.accept=void 0,this.maxlength=void 0,this.size=void 0,this.full=!1,this.style={}}static get properties(){return{id:{type:String},name:{type:String,reflect:!0},type:{type:String,converter:e=>Object.values(a).includes(e)?e:a.TEXT},value:{type:String,reflect:!0},placeholder:{type:String},label:{type:String},disabled:{type:Boolean,reflect:!0},readonly:{type:Boolean,reflect:!0},autofocus:{type:Boolean,reflect:!0},accept:{type:String,reflect:!0},maxlength:{type:Number,reflect:!0},size:{type:Number,reflect:!0},full:{type:Boolean},style:{type:Object}}}_onInput(e){const t=this.type===a.FILE?e.target.files:e.target.value;t!==this.value&&(this.onChange&&this.onChange({value:t,old:this.value}),this.dispatchEvent(new CustomEvent("change",{bubbles:!0,composed:!0,detail:{value:t,old:this.value}})),this.value=t),this._input.value!==t&&(this._input.value=t)}_onKeyEvent(e){const t=()=>{this.onEnter&&(this.onEnter({value:this.value}),this.dispatchEvent(new CustomEvent("enter",{bubbles:!0,composed:!0,detail:{value:this.value}})))};"keydown"!==e.type&&"keypress"!==e.type||e.stopPropagation(),"keydown"===e.type&&13===e.keyCode&&(e.preventDefault(),t()),this.readonly||"keypress"!==e.type||13!==e.keyCode||t()}firstUpdated(){this._input=this.renderRoot.querySelector("input")}render(){return o`<input 
      id="${this.id}"
      name="${l(this.name)}"
      .type="${this.type}"
      .value="${this.value}"
      placeholder="${l(this.placeholder)}"
      ?disabled="${this.disabled}"
      ?readonly="${this.readonly}"
      ?autofocus="${this.autofocus}"
      aria-label="${l(this.label)}"
      accept="${l(this.accept)}"
      maxlength="${l(this.maxlength)}"
      size="${l(this.size)}"
      class="${this.full?"full":""}"
      style="${s(this.style)}"
      @input=${this._onInput}
      @keydown=${this._onKeyEvent}
      @keypress=${this._onKeyEvent}
    >`}static get styles(){return[d]}});const h=e`
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
    line-height: 1;
  }
  .label_info_left {
    padding: 0 0 0 3px;
  }
  .label_info_right {
    text-align: right;
    padding: 0 3px 0 0;
  }
  .label_icon_left {
    text-align: center;
    width: auto;
  }
  .label_icon_right {
    text-align: right;
    width: auto;
  }
  .centered {
    margin: auto;
  }
  .bold {
    font-weight: bold;
  }
`;customElements.define("form-label",class extends t{constructor(){super(),this.id=void 0,this.value="",this.centered=!1,this.leftIcon=void 0,this.rightIcon=void 0,this.style={},this.iconStyle={}}static get properties(){return{id:{type:String},value:{type:String},centered:{type:Boolean},leftIcon:{type:String},rightIcon:{type:String},style:{type:Object},iconStyle:{type:Object}}}render(){return this.leftIcon?o`<div 
        id="${l(this.id)}" 
        class="row ${this.centered?"centered":""}">
        <div class="cell label_icon_left">
          ${o`<form-icon 
            iconKey="${this.leftIcon}" width=20
            color="${l(this.iconStyle.color)}"
            .style="${this.iconStyle}"></form-icon>`}
        </div>
        <div class="cell label_info_left bold"
          style="${s(this.style)}" >${this.value}</div>
      </div>`:this.rightIcon?o`<div 
        id="${l(this.id)}" class="row full">
        <div class="cell label_info_right bold"
          style="${s(this.style)}" >${this.value}</div>
        <div class="cell label_icon_right">
          ${o`<form-icon 
            iconKey="${this.rightIcon}" width=20
            color="${l(this.iconStyle.color)}"
            .style="${this.iconStyle}"></form-icon>`}
        </div>
      </div>`:o`<span 
      id="${l(this.id)}" class="bold" 
      style="${s(this.style)}">${this.value}</span>`}static get styles(){return[h]}});
