import{a as e,i as t,b as i,A as l}from"./module-1JqPnLJw.js";import"./module-tEN-oy1p.js";import{S as a,a as s,d as o,B as n,b as d,r,D as c,I as h,E as p,e as u,P as m,c as v,M as g,f}from"./main-DoduxAVE.js";import{i as b,t as $,e as y,o as x,a as _}from"./module-BCutSLl1.js";import"./module-NpGsdRHB.js";const w=e`
:host {
  font-family: var(--font-family);
  font-size: var(--font-size);
  color: rgba(var(--accent-1c), 0.85);
  fill: rgba(var(--accent-1c), 0.85);
}
div {
  box-sizing: border-box;
}
@keyframes animateleft{ 
  from { left:-300px; opacity:0; }
  to{ left:0; opacity:1; }
}
.sidebar{
  position: fixed;
  height: 100%;
  width: var(--menu-side-width);
  left: 0;
  animation: animateleft 0.4s;
  background-color: rgba(var(--accent-1), 0.95);
  z-index: 3;
  overflow-x: hidden;
  overflow-y: auto;
  font-weight: bold;
  padding: 5px 4px 50px 2px;
  box-shadow:0 2px 5px 0 rgba(0,0,0,0.16),0 2px 10px 0 rgba(0,0,0,0.12);
}
.hide{ 
  display: none; 
}
.show{ 
  display: block!important; 
}
.row {
  display: table;
}
.cell {
  display: table-cell;
  vertical-align: middle;
}
.full, .half {
  width: 100%; 
}
.separator {
  height: 10px;
  width: 100%;
  border: none;
  margin: 0;
}
.panel-group {
  padding: 2px 8px 3px;
}
.container { 
  padding: 8px 16px;
}
.state-label {
  font-size: 18px;
  width: 100%;
  padding: 8px 16px;
  background-color: rgb(var(--accent-1));
  vertical-align: middle;
}
@media (max-width:600px){
  .container { 
    padding: 4px 8px; 
  }
}
@media only screen and (min-width: 601px){
  .half { 
    width:49.99999% 
  }
}
@media (min-width:769px){
  .sidebar{
    display: block!important;
  }
}
@media (max-width:768px){
  .sidebar{
    display: none;
  }
}
*::-webkit-scrollbar {
  width: 10px;
  height: 5px;
  background-color: transparent;
  visibility: hidden;
}
*::-webkit-scrollbar-track {
  background-color: rgba(var(--functional-green), .05);
  border-radius: 8px;
}
*::-webkit-scrollbar-thumb {
  background-color: rgba(var(--functional-green), .30);
  border-radius: 8px;
}
*::-webkit-scrollbar-thumb:active,
*::-webkit-scrollbar-thumb:hover {
  background-color: rgba(var(--functional-green), .20)
}
`;customElements.define("sidebar-template",class extends t{constructor(){super(),this.msg=e=>e,this.side=a.AUTO,this.theme=s.LIGHT,this.templateKey="",this.dirty=!1}static get properties(){return{side:{type:String,reflect:!0},templateKey:{type:String},dirty:{type:Boolean}}}static get styles(){return[w]}_onSideEvent(e,t){this.onEvent&&this.onEvent.onSideEvent&&this.onEvent.onSideEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("side_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}itemMenu({id:e,selected:t,eventValue:l,label:a,iconKey:s,full:d}){return i`<form-button 
        id="${e}" label="${a}"
        ?full="${void 0===d||d}" ?selected="${void 0!==t&&t}"
        align=${o.LEFT}
        .style="${{"border-radius":"0","border-color":"rgba(var(--accent-1c), 0.2)"}}"
        icon="${s}" type="${n.PRIMARY}"
        @click=${()=>this._onSideEvent(...l)} 
      >${a}</form-button>`}formItems(){const e=[];return e.push(this.itemMenu({id:"cmd_theme",selected:!0,eventValue:[d.THEME,{}],label:this.msg("",{id:"menu_theme"}),iconKey:this.theme===s.DARK?"Sun":"Moon",full:!1})),e.push(i`<hr id="back_sep" class="separator" />`),["_blank","_sample"].includes(this.templateKey)||(e.push(i`<hr id="tmp_sep_2" class="separator" />`),e.push(this.itemMenu({id:"cmd_save",selected:this.dirty,eventValue:[d.SAVE,!0],label:this.msg("",{id:"template_save"}),iconKey:"Check"})),e.push(this.itemMenu({id:"cmd_create",eventValue:[d.CREATE_REPORT,{}],label:this.msg("",{id:"template_create_from"}),iconKey:"Sitemap"})),e.push(this.itemMenu({id:"cmd_delete",eventValue:[d.DELETE,{}],label:this.msg("",{id:"label_delete"}),iconKey:"Times"}))),e.push(i`<hr id="tmp_sep_3" class="separator" />`),e.push(this.itemMenu({id:"cmd_blank",eventValue:[d.CHECK,{value:d.BLANK}],label:this.msg("",{id:"template_new_blank"}),iconKey:"Plus"})),e.push(this.itemMenu({id:"cmd_sample",eventValue:[d.CHECK,{value:d.SAMPLE}],label:this.msg("",{id:"template_new_sample"}),iconKey:"Plus"})),e.push(i`<hr id="tmp_sep_4" class="separator" />`),e.push(this.itemMenu({id:"cmd_print",eventValue:[d.REPORT_SETTINGS,{value:"PREVIEW"}],label:this.msg("",{id:"label_print"}),iconKey:"Eye"})),e.push(this.itemMenu({id:"cmd_json",eventValue:[d.REPORT_SETTINGS,{value:"JSON"}],label:this.msg("",{id:"template_export_json"}),iconKey:"Code"})),e.push(i`<hr id="tmp_sep_5" class="separator" />`),e.push(this.itemMenu({id:"cmd_help",eventValue:[d.HELP,{value:"editor"}],label:this.msg("",{id:"label_help"}),iconKey:"QuestionCircle"})),e.push(this.itemMenu({id:"cmd_logout",eventValue:[d.LOGOUT,{}],label:this.msg("",{id:"menu_logout"}),iconKey:"Exit"})),e}render(){return i`<div class="sidebar ${"auto"!==this.side?this.side:""}" >
    ${this.formItems()}
    </div>`}});const E=(e,t)=>{const i=e._$AN;if(void 0===i)return!1;for(const e of i)e._$AO?.(t,!1),E(e,t);return!0},S=e=>{let t,i;do{if(void 0===(t=e._$AM))break;i=t._$AN,i.delete(e),e=t}while(0===i?.size)},k=e=>{for(let t;t=e._$AM;e=t){let i=t._$AN;if(void 0===i)t._$AN=i=new Set;else if(i.has(e))break;i.add(e),P(t)}};function T(e){void 0!==this._$AN?(S(this),this._$AM=e,k(this)):this._$AM=e}function C(e,t=!1,i=0){const l=this._$AH,a=this._$AN;if(void 0!==a&&0!==a.size)if(t)if(Array.isArray(l))for(let e=i;e<l.length;e++)E(l[e],!1),S(l[e]);else null!=l&&(E(l,!1),S(l));else E(this,e)}const P=e=>{e.type==$.CHILD&&(e._$AP??=C,e._$AQ??=T)};class A extends b{constructor(){super(...arguments),this._$AN=void 0}_$AT(e,t,i){super._$AT(e,t,i),k(this),this.isConnected=e._$AU}_$AO(e,t=!0){e!==this.isConnected&&(this.isConnected=e,e?this.reconnected?.():this.disconnected?.()),t&&(E(this,e),S(this))}setValue(e){if(r(this._$Ct))this._$Ct._$AI(e,this);else{const t=[...this._$Ct._$AH];t[this._$Ci]=e,this._$Ct._$AI(t,this,0)}}disconnected(){}reconnected(){}}const z=new WeakMap,I=y(class extends A{render(e){return l}update(e,[t]){const i=t!==this.G;return i&&void 0!==this.G&&this.rt(void 0),(i||this.lt!==this.ct)&&(this.G=t,this.ht=e.options?.host,this.rt(this.ct=e.element)),l}rt(e){if(this.isConnected||(e=void 0),"function"==typeof this.G){const t=this.ht??globalThis;let i=z.get(t);void 0===i&&(i=new WeakMap,z.set(t,i)),void 0!==i.get(this.G)&&this.G.call(this.ht,void 0),i.set(this.G,e),void 0!==e&&this.G.call(this.ht,e)}else this.G.value=e}get lt(){return"function"==typeof this.G?z.get(this.ht??globalThis)?.get(this.G):this.G?.value}disconnected(){this.lt===this.ct&&this.rt(void 0)}reconnected(){this.rt(this.ct)}}),O=e`
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
      name="${x(this.name)}"
      type="number"
      .value="${String(this.value)}"
      ?disabled="${this.disabled}"
      ?readonly="${this.readonly}"
      ?autofocus="${this.autofocus}"
      min="${x(this.min)}"
      max="${x(this.max)}"
      aria-label="${x(this.label)}"
      class="${this.full?"full":""}"
      style="${_(this.style)}"
      @input=${this._onInput}
      @blur=${this._onBlur}
      @keydown=${this._onKeyEvent}
      @keypress=${this._onKeyEvent}
    >`}static get styles(){return[O]}});const M=e`
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
`,N={[c.TIME]:5,[c.DATE]:10,[c.DATETIME]:16};customElements.define("form-datetime",class extends t{constructor(){super(),this.id=Math.random().toString(36).slice(2),this.name=void 0,this.value="",this.type=c.DATE,this.label=void 0,this.isnull=!0,this.picker=!1,this.disabled=!1,this.readonly=!1,this.autofocus=!1,this.full=!1,this.style={}}static get properties(){return{id:{type:String},name:{type:String,reflect:!0},value:{type:String,reflect:!0},type:{type:String,converter:e=>Object.values(c).includes(e)?e:c.DATE},label:{type:String},isnull:{type:Boolean},picker:{type:Boolean},disabled:{type:Boolean,reflect:!0},readonly:{type:Boolean,reflect:!0},autofocus:{type:Boolean,reflect:!0},full:{type:Boolean},style:{type:Object}}}_defaultValue(){const e=`${(new Date).toISOString().slice(0,10)}T${(new Date).toLocaleTimeString("en",{hour12:!1}).replace("24","00").slice(0,5)}`;switch(this.type){case c.DATE:return String(e).split("T")[0];case c.TIME:return String(e).split("T")[1].split(".")[0];default:return e}}_onInput(e){const t=e=>{const t=this.type!==c.DATE?`${e}:00`:e;t!==this.value&&(this.onChange&&this.onChange({value:t,old:this.value}),this.dispatchEvent(new CustomEvent("change",{bubbles:!0,composed:!0,detail:{value:t,old:this.value}})),this.value=t),this._input.value!==t&&(this._input.value=t)};if(e.target.value!==this.value){if(""===e.target.value&&!this.isnull)return void t(this._defaultValue());t(e.target.value)}}_onBlur(){this._input.value=this.value}_onKeyEvent(e){const t=()=>{this.onEnter&&(this.onEnter({value:this.value}),this.dispatchEvent(new CustomEvent("enter",{bubbles:!0,composed:!0,detail:{value:this.value}})))};"keydown"!==e.type&&"keypress"!==e.type||e.stopPropagation(),"keydown"===e.type&&13===e.keyCode&&(e.preventDefault(),t()),this.readonly||"keypress"!==e.type||13!==e.keyCode||t()}_onFocus(){this.picker&&this._input.showPicker()}firstUpdated(){this._input=this.renderRoot.querySelector("input")}render(){let e=this.value;return e.length>N[this.type]&&(e=e.slice(0,N[this.type])),i`<input 
      id="${this.id}"
      name="${x(this.name)}"
      .type="${this.type}"
      .value="${e}"
      ?disabled="${this.disabled}"
      ?readonly="${this.readonly}"
      ?autofocus="${this.autofocus}"
      aria-label="${x(this.label)}" 
      style="${_(this.style)}"
      class="${this.full?"full":""}"
      @input=${this._onInput}
      @blur=${this._onBlur}
      @keydown=${this._onKeyEvent}
      @keypress=${this._onKeyEvent}
      @focus=${this._onFocus}
    >`}static get styles(){return[M]}});const R=e`
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
`;customElements.define("form-field",class extends t{constructor(){super(),this.id=Math.random().toString(36).slice(2),this.field={},this.values={},this.options={},this.data={dataset:{},current:{},audit:"all"},this.style={},this.msg=e=>e}static get properties(){return{id:{type:String},field:{type:Object},values:{type:Object},options:{type:Object},data:{type:Object},style:{type:Object}}}static get styles(){return[R]}_onChange({value:e,item:t,event_type:i,editName:l,fieldMap:a}){const s={id:this.field.id||1,name:l,event_type:i,value:e,extend:!(!a||!a.extend),refnumber:t&&t.label?t.label:this.field.link_label,label_field:a?a.label_field:void 0,item:t};this.onEdit&&this.onEdit(s),this.dispatchEvent(new CustomEvent("change",{bubbles:!0,composed:!0,detail:{...s}}))}_onEvent(e,t){this.onEvent&&this.onEvent(e,t),this.dispatchEvent(new CustomEvent("event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}_onSelector(e,t,i){this.onSelector&&this.onSelector(e,t,i),this.dispatchEvent(new CustomEvent("selector",{bubbles:!0,composed:!0,detail:{type:e,filter:t,callback:i}}))}lnkValue({fieldMap:e,fieldName:t,value:i}){if(void 0===this.values[this.field.name])return[this.data.current[e.source]?this.data.current[e.source].filter(t=>t.ref_id===this.values.id&&t[e.value]===this.field.name)[0]:this.data.dataset[e.source].filter(t=>t.ref_id===this.values.id&&t[e.value]===this.field.name)[0],!1];const l="id"===t&&""===i?null:i;return[this.data.current[e.source]?this.data.current[e.source].filter(t=>t[e.value]===l)[0]:this.data.dataset[e.source].filter(t=>t[e.value]===l)[0],!0]}getOppositeValue(e){return this.options.opposite&&parseFloat(e)<0?String(e).replace("-",""):this.options.opposite&&parseFloat(e)>0?`-${e}`:e}selectorInit({fieldMap:e,value:t,editName:i}){let l={value:this.values.value||"",filter:this.field.description||"",text:this.field.description||"",type:this.field.fieldtype,ntype:["transitem","transmovement","transpayment"].includes(this.field.fieldtype)?"trans":this.field.fieldtype,ttype:null,id:this.values.value||null,table:{name:"fieldvalue",fieldname:"value",id:this.field.id},fieldMap:e,editName:i};if(e){let i;if(l={...l,value:t||"",type:e.seltype,filter:String(l.text).split(" | ")[0],table:{name:e.table,fieldname:e.fieldname},ntype:e.lnktype,ttype:""!==t?""===e.transtype?this.values.transtype:e.transtype:l.ttype,id:""!==t?t:l.id},!0===e.extend||"extend"===e.table)i=this.data.current.extend;else{let t;t=void 0!==this.data.current[e.table]&&Array.isArray(this.data.current[e.table])?this.data.current[e.table].filter(e=>e.id===this.values.id):this.data.dataset[e.table].filter(e=>e.id===this.values.id),i=t[0]}void 0!==i?(l=void 0!==this.values[e.label_field]?{...l,text:this.values[e.label_field]||"",filter:this.values[e.label_field]||""}:void 0!==i[e.label_field]&&null!==i[e.label_field]?{...l,text:i[e.label_field],filter:i[e.label_field]}:{...l,text:"",filter:""},void 0!==i[e.fieldname]&&""===l.value&&null!==i[e.fieldname]&&(l={...l,ntype:e.lnktype,ttype:e.transtype,id:i[e.fieldname]}),"trans"===e.lnktype&&void 0!==i.transtype&&(l=void 0!==e.lnkid?{...l,id:i[e.lnkid]}:void 0!==i[e.fieldname]?{...l,id:i[e.fieldname]}:{...l,id:l.value})):l={...l,text:"",filter:""}}return this.selector=l,l}setSelector(e,t){let i={...this.selector,text:"",id:null,filter:t||""};if(e){const t=e.id.split("/");i={...i,text:e.label||e.item.lslabel},i={...i,id:parseInt(t[2],10),ttype:t[1]},"trans"===t[0]&&""!==t[1]&&e.trans_id&&(i={...i,id:e.trans_id})}i={...i,value:i.id||""},this.selector=i,this._onChange({value:i.id,item:e,event_type:"change",editName:i.editName,fieldMap:i.fieldMap})}editName(){return this.field.map?this.field.map.extend&&this.field.map.text?this.field.map.text:this.field.map.fieldname?this.field.map.fieldname:this.field.name:this.field.name}_onTextInput(e){this._onChange({value:e.target.value,event_type:"change",editName:this.editName(),fieldMap:this.field.map})}render(){let{datatype:e}=this.field,t=this.field.disabled||"readonly"===this.data.audit,a=this.field.name,s=this.values[this.field.name];const o=this.field.map||null,d=this.editName(),r=!("true"!==this.field.empty&&!0!==this.field.empty);switch("reportfield"!==this.field.rowtype&&"fieldvalue"!==this.field.rowtype||(s=this.values.value),null==s&&(s=this.field.default?this.field.default:""),"fieldvalue"===e&&(e=this.values.datatype),e){case"password":case"color":return i`<form-input 
          id="${this.id}" name="${a}" type="${e}" 
          value="${s||""}" label="${a}"
          ?full="${!0}"
          .style="${this.style}" 
          ?disabled="${t}"
          .onChange=${e=>this._onChange({value:e.value,event_type:"change",editName:d,fieldMap:o})}></form-input>`;case"date":case"datetime":if(o)if(o.extend)s=this.data.current.extend[o.text],a=o.text;else{const e=this.lnkValue({fieldMap:o,fieldName:a,value:s});void 0!==e[0]&&(s=e[0][o.text],t=e[1]?e[1]:t)}return i`<form-datetime id="${this.id}"
          name="${a}" label="${a}"
          .style="${this.style}" .isnull="${r}"
          type="${"datetime"===e?c.DATETIME:c.DATE}"
          .value="${s}"
          .onChange="${e=>{this._onChange({value:e.value,event_type:"change",editName:d,fieldMap:o})}}"
          ?disabled="${t}"></form-datetime>`;case"bool":case"flip":const m=t?"toggle-disabled":"";return[1,"1","true",!0].includes(s)?i`<div id="${this.id}"
            name="${a}" style="${_(this.style)}" 
            class="${`toggle toggle-on ${m}`}"
            @click="${t?null:()=>this._onChange({value:"fieldvalue_value"!==this.field.name&&0,event_type:"change",editName:d,fieldMap:o})}">
            <form-icon iconKey="ToggleOn" width=40 height=32.6 ></form-icon>
          </div>`:i`<div id="${this.id}"
          name="${a}" style="${_(this.style)}" 
          class="${`toggle toggle-off ${m}`}"
          @click="${t?null:()=>this._onChange({value:"fieldvalue_value"===this.field.name||1,event_type:"change",editName:d,fieldMap:o})}">
          <form-icon iconKey="ToggleOff" width=40 height=32.6 ></form-icon>
        </div>`;case"label":return null;case"select":this.field.extend&&(s=this.data.current.extend[this.field.name]||"");const v=[];return o?this.data.dataset[o.source].forEach(e=>{let t=e[o.text];void 0!==o.label&&(t=this.msg(`${o.label}_${t}`,{id:`${o.label}_${t}`})),v.push({value:String(e[o.value]),text:t})}):this.field.options.forEach(e=>{let t=e[1];this.msg(t,{id:t})&&(t=this.msg(t,{id:t})),void 0!==this.field.olabel&&(t=this.msg(`${this.field.olabel}_${e[1]}`,{id:`${this.field.olabel}_${e[1]}`})),v.push({value:String(e[0]),text:t})}),i`<form-select id="${this.id}" ?full="${!0}"
            name="${a}" label="${a}"
            .style="${this.style}"
            ?disabled="${t}" 
            .onChange=${e=>{const t=Number.isNaN(parseInt(e.value,10))?e.value:parseInt(e.value,10);this._onChange({value:t,event_type:"change",editName:d,fieldMap:o})}}
            .options=${v} 
            .isnull="${r}" value="${s}" ></form-select>`;case"valuelist":return i`<form-select id="${this.id}" ?full="${!0}"
          name="${a}" label="${a}"
          .style="${this.style}"
          ?disabled="${t}"
          .onChange=${e=>this._onChange({value:e.value,event_type:"change",editName:d,fieldMap:o})}
          .options=${this.field.description.map(e=>({value:e,text:e}))} 
          .isnull="${!1}" value="${s}" ></form-select>`;case"link":let g=this.values;const f=this.lnkValue({fieldMap:o,fieldName:a,value:s});void 0!==f[0]&&(g=f[0],f[0][o.text]&&(s=f[0][o.text]));let b=s;return void 0!==o.label_field&&void 0!==g[o.label_field]&&(b=g[o.label_field]),i`<div 
          name="${a}" style="${_(this.style)}" class="link" >
          <span id=${`link_${o.lnktype}_${a}`} class="link-text" 
            @click="${()=>this._onEvent(p.CHECK_EDITOR,[{ntype:o.lnktype,ttype:o.transtype,id:s},u.LOAD_EDITOR])}" >${b}</span>
        </div>`;case"selector":let{selector:$}=this;const y=[],w=o&&(!0===o.extend||"extend"===o.table)&&this.data.current.extend&&this.data.current.extend.seltype;return $?w&&$.type!==this.data.current.extend.seltype&&($.text=this.data.current.extend[o.label_field],$.type=this.data.current.extend.seltype,$.filter=$.text,$.ntype=this.data.current.extend.seltype,$.ttype=this.data.current.extend.transtype,$.id=this.data.current.extend.ref_id):$=this.selectorInit({fieldMap:o,value:s,editName:d}),t||y.push(i`<div id="sel_show" class="cell search-col">
            <form-button id="${`sel_show_${a}`}" 
              label="${this.msg("",{id:"label_search"})}"
              .style="${{padding:"5px 8px 7px"}}"
              icon="Search" type="${n.BORDER}"
              @click=${()=>this._onSelector($.type,$.filter,(...e)=>this.setSelector(...e))} 
            ></form-button>
          </div>`),r&&y.push(i`<div id="sel_delete" class="cell times-col">
            <form-button id="${`sel_delete_${a}`}" 
              label="${this.msg("",{id:"label_delete"})}"
              .style="${{padding:"5px 8px 7px"}}"
              ?disabled="${t}" icon="Times" type="${n.BORDER}"
              @click=${t?null:()=>this.setSelector()} ></form-button>
          </div>`),y.push(i`<div id="sel_text" class="link">
          ${""!==$.text?i`<span 
            id=${`sel_link_${a}`}
            class="link-text"
            @click="${()=>this._onEvent(p.CHECK_TRANSTYPE,[{ntype:$.ntype,ttype:$.ttype,id:$.id},u.LOAD_EDITOR])}" >${$.text}</span>`:null}
        </div>`),i`<div id="${this.id}" 
          name="${a}" style="${_(this.style)}"
          class="row full" >${y}</div>`;case"button":return i`<form-button id="${this.id}" 
          name="${a}" type="${n.BORDER}"
          .style="${{padding:"7px 8px",...this.style}}"
          ?disabled="${t}" 
          label="${this.field.title?this.field.title:l}" 
          ?full="${this.field.full}"
          ?autofocus="${this.field.focus||!1}"
          icon="${this.field.icon}"
          @click=${t?null:()=>this._onChange({value:a,item:{},event_type:"click",editName:d,fieldMap:o})} >${this.field.title?this.field.title:l}</form-button>`;case"percent":case"integer":case"float":if(o)if(o.extend)s=this.data.current.extend[o.text],a=o.text;else{const e=this.lnkValue({fieldMap:o,fieldName:a,value:s});void 0!==e[0]&&(s=e[0][o.text],t=e[1]?e[1]:t)}return""===s&&(s=0),void 0!==this.field.opposite&&(s=this.getOppositeValue(s)),i`<form-number id="${this.id}" name="${a}"
          ?integer="${!("float"===e)}" 
          ?full="${!0}" .style="${this.style}"
          value="${s||0}" 
          ?disabled="${t}" label="${a}"
          min="${x(this.field.min)}" 
          max="${x(this.field.max)}" 
          .onChange=${e=>{this._onChange({value:this.field.opposite?this.getOppositeValue(e.value):e.value,event_type:"change",editName:d,fieldMap:o}),this.event_type="change"}}
          .onBlur=${e=>{this._onChange({value:this.field.opposite?parseFloat(this.getOppositeValue(e.value)):parseFloat(e.value),event_type:"change"===this.event_type?"blur":null,editName:d,fieldMap:o}),this.event_type="blur"}}
        ></form-number>`;default:if(o)if(o.extend)s=this.data.current.extend[o.text],a=o.text;else{const e=this.lnkValue({fieldMap:o,fieldName:a,value:s});void 0!==e[0]&&(s=e[0][o.text],t=e[1]?e[1]:t,void 0!==o.label&&(s=this.msg(`${o.label}_${s}`,{id:`${o.label}_${s}`})))}return"notes"===e||"text"===e?i`<textarea id="${this.id}" name="${a}"
            class=${"full"} label="${a}" style="${_(this.style)}"
            rows=${x(this.field.rows?this.field.rows:void 0)}
            @input="${this._onTextInput}"
            .value="${s||""}"
            ?disabled="${t}" ></textarea>`:i`<form-input 
          id="${this.id}" .id="${`${this.id}_input`}" name="${a}" 
          type="${h.TEXT}" 
          value="${s||""}" label="${a}"
          ?full="${!0}"
          .style="${this.style}" 
          maxlength="${x(this.field.length?this.field.length:void 0)}"
          size="${x(this.field.length?this.field.length:void 0)}"
          ?disabled="${t}"
          .onChange=${e=>this._onChange({value:e.value,event_type:"change",editName:d,fieldMap:o})}></form-input>`}}});const D=e`
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
`;customElements.define("form-row",class extends t{constructor(){super(),this.id=Math.random().toString(36).slice(2),this.row={},this.values={},this.options={},this.data={dataset:{},current:{},audit:"all"},this.style={},this.msg=e=>e}static get properties(){return{id:{type:String},row:{type:Object},values:{type:Object},options:{type:Object},data:{type:Object},style:{type:Object}}}static get styles(){return[D]}_onEdit(e){this.onEdit&&this.onEdit(e),this.dispatchEvent(new CustomEvent("edit",{bubbles:!0,composed:!0,detail:{...e}}))}_onTextInput(e){this._onEdit({id:this.row.id,name:this.row.name,value:e.target.value})}imgValue(){let e=this.values[this.row.name]||"";return""!==e&&null!==e&&"data:image"!==e.toString().substr(0,10)&&void 0!==this.data.dataset[e]&&(e=this.data.dataset[e]),this.safeImageSrc(e)}safeImageSrc(e){if(!e||"string"!=typeof e)return"";const t=e.trim();if(""===t)return"";const i=t.toLowerCase();if(i.startsWith("blob:"))return t;if(i.startsWith("https:")||i.startsWith("http:"))return t;if(i.startsWith("data:image/")){const e=i.slice(11);if(["png","jpeg","jpg","gif","webp","ico","bmp"].some(t=>e.startsWith(t+";")||e.startsWith(t+",")))return t}return""}flipItem(){const{id:e,name:t,datatype:l,info:a}=this.row,s=void 0!==this.values[t],o=i`<div id="${`checkbox_${t}`}"
      class="report-field ${s?"toggle-on":"toggle-off"}"
      @click="${()=>this._onEdit({id:e,selected:!0,datatype:l,defvalue:this.row.default,name:t,value:!s,extend:!1})}">
      ${s?i`<form-icon iconKey="ToggleOn" width=40 height=32.6 ></form-icon>`:i`<form-icon iconKey="ToggleOff" width=40 height=32.6 ></form-icon>`}
      <form-label value="${t}" class="bold padding-tiny ${s?"toggle-on":""} " ></form-label>
    </div>`;switch(l){case"text":return i`<div id="${this.id}"
          style="${_(this.style)}" class="container-row">
          <div class="row full">
            <div class="cell padding-small" >
              ${o}
            </div>
          </div>
          ${s?i`<div class="row full"><div class="cell padding-small" >
            <form-field id="${`field_${t}`}"
              .field=${this.row} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div></div>`:null}
          ${a?i`<div class="row full padding-small">
            <div class="cell padding-small info leftbar" >
              ${a}
            </div>
          </div>`:null}
        </div>`;case"image":return i`<div id="${this.id}"
            style="${_(this.style)}" class="container-row">
            <div class="row full">
              <div class="cell padding-small" >
                ${o}
              </div>
              ${s?i`<div class="cell padding-small" >
              <form-input 
                id="${`file_${t}`}" 
                type="${h.FILE}" ?full="${!0}"
                label="${this.labelAdd}"
                .style="${{"font-size":"12px"}}"
                .onChange=${i=>this._onEdit({id:e,file:!0,name:t,value:i.value,extend:!1})}></form-input>
              </div>`:null}
            </div>
            ${s?i`<div class="row full"><div class="cell padding-small" >
              <textarea id="${`input_${t}`}"
                class=${"full"} rows=5 .value="${this.imgValue()}"
                @input="${this._onTextInput}" ></textarea>
              <div class="full padding-normal center" >
                <img src="${this.imgValue()}" alt="" />
              </div>
            </div></div>`:null}
            ${a?i`<div class="row full padding-small">
              <div class="cell padding-small info leftbar" >
                ${a}
              </div>
            </div>`:null}
          </div>`;case"checklist":const l=this.values[t]||"",n=[];return this.row.values.forEach((a,s)=>{const o=a.split("|"),d=l.indexOf(o[0])>-1;n.push(i`<div id="${`checklist_${t}_${s}`}"
            key={index}
            class="cell padding-small report-field"
            @click=${()=>this._onEdit({id:e,checklist:!0,name:t,checked:!d,value:o[0],extend:!1})}>
            <form-label 
              value="${o[1]}" class="bold ${d?"toggle-on":""}"
              leftIcon="${d?"CheckSquare":"SquareEmpty"}" ></form-label>
          </div>`)}),i`<div id="${this.id}"
            style="${_(this.style)}" class="container-row">
          <div class="row full">
            <div class="cell padding-small" >
              ${o}
            </div>
          </div>
          ${s?i`<div class="row full padding-small">
            <div class="cell padding-small toggle" >
              ${n}
            </div>
          </div>`:null}
          ${a?i`<div class="row full padding-small">
            <div class="cell padding-small info leftbar" >
              ${a}
            </div>
          </div>`:null}
        </div>`;default:return i`<div id="${this.id}"
          style="${_(this.style)}" class="container-row">
          <div class="row full">
            <div class="cell padding-small half" >
              ${o}
            </div>
            ${s?i`<div class="cell padding-small half" >
              <form-field id="${`field_${t}`}"
                .field=${this.row} .values=${this.values} .options=${this.options} .data=${this.data}
                .msg=${this.msg} .onEdit=${this.onEdit} 
                .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
            </div>`:null}
          </div>
          ${a?i`<div class="row full padding-small">
            <div class="cell padding-small info leftbar" >
              ${a}
            </div>
          </div>`:null}
        </div>`}}render(){const{id:e,rowtype:t,label:l,columns:a,name:s,disabled:o,notes:n,selected:d,empty:r,datatype:c,info:p}=this.row;switch(t){case"label":return i`<div id="${this.id}" style="${_(this.style)}" 
          class="container-row label-row">
          <div class="cell padding-small" >${this.values[s]||l}</div>
        </div>`;case"flip":return this.flipItem();case"field":return i`<div id="${this.id}"
          style="${_(this.style)}" class="container-row">
          <div class="cell padding-small hide-small field-cell" >
            <form-label value="${l}" class="bold" ></form-label>
          </div>
          <div class="cell padding-small" >
            <div class="hide-medium hide-large" >
              <form-label value="${l}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${s}`}"
              .field=${this.row} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
        </div>`;case"reportfield":return i`<div id="${this.id}"
          style="${_(this.style)}" class="cell padding-small s12 m6 l4">
          <div id="${`cb_${s}`}"
            class=${"padding-small "+("false"!==r?"report-field":"")} 
            @click="${()=>{"false"!==r&&this._onEdit({id:e,name:"selected",value:!d,extend:!1})}}">
            <form-label 
              value="${l}" class="bold"
              leftIcon="${d?"CheckSquare":"SquareEmpty"}" ></form-label>
          </div>
          <form-field id="${`field_${s}`}"
            .field=${this.row} .values=${this.values} .options=${this.options} .data=${this.data}
            .msg=${this.msg} .onEdit=${this.onEdit} 
            .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
        </div>`;case"fieldvalue":return i`<div id="${this.id}"
          style="${_(this.style)}" class="container-row">
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
                id="${`notes_${this.row.fieldname}`}" type="${h.TEXT}"
                label="${this.msg("",{id:"fnote_view"})}"
                name="fieldvalue_notes" ?full="${!0}" value="${n}"
                ?disabled=${o||"readonly"===this.data.audit}
                .onChange=${t=>this._onEdit({id:e,name:"fieldvalue_notes",value:t.value})}></form-input>
            </div>
          </div>
        </div>`;case"col2":return i`<div id="${this.id}"
          style="${_(this.style)}" class="container-row">
          <div class="cell padding-small s12 m6 l6" >
            <div>
              <form-label value="${a[0].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${a[0].name}`}"
              .field=${a[0]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
          <div class="cell padding-small s12 m6 l6" >
            <div>
              <form-label value="${a[1].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${a[1].name}`}"
              .field=${a[1]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
        </div>`;case"col3":return i`<div id="${this.id}"
          style="${_(this.style)}" class="container-row">
          <div class="cell padding-small s12 m4 l4" >
            <div>
              <form-label value="${a[0].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${a[0].name}`}"
              .field=${a[0]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
          <div class="cell padding-small s12 m4 l4" >
            <div>
              <form-label value="${a[1].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${a[1].name}`}"
              .field=${a[1]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
          <div class="cell padding-small s12 m4 l4" >
            <div>
              <form-label value="${a[2].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${a[2].name}`}"
              .field=${a[2]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
        </div>`;case"col4":return i`<div id="${this.id}"
          style="${_(this.style)}" class="container-row">
          <div class="cell padding-small s12 m3 l3" >
            <div>
              <form-label value="${a[0].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${a[0].name}`}"
              .field=${a[0]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
          <div class="cell padding-small s12 m3 l3" >
            <div>
              <form-label value="${a[1].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${a[1].name}`}"
              .field=${a[1]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
          <div class="cell padding-small s12 m3 l3" >
            <div>
              <form-label value="${a[2].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${a[2].name}`}"
              .field=${a[2]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
          <div class="cell padding-small s12 m3 l3" >
            <div>
              <form-label value="${a[3].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${a[3].name}`}"
              .field=${a[3]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
        </div>`}return null}});const B=e`
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
`;customElements.define("form-pagination",class extends t{constructor(){super(),this.id=Math.random().toString(36).slice(2),this.name=void 0,this.pageIndex=0,this.pageSize=5,this.pageCount=0,this.canPreviousPage=!1,this.canNextPage=!1,this.hidePageSize=!1,this.style={}}static get properties(){return{id:{type:String},name:{type:String},pageIndex:{type:Number},pageSize:{type:Number},pageCount:{type:Number},canPreviousPage:{type:Boolean},canNextPage:{type:Boolean},hidePageSize:{type:Boolean},style:{type:Object}}}static get styles(){return[B]}_onEvent(e,t,i){i&&(this.onEvent&&this.onEvent(e,t),this.dispatchEvent(new CustomEvent("pagination",{bubbles:!0,composed:!0,detail:{key:e,value:t}})))}render(){return i`<div id="${this.id}" name="${x(this.name)}"
      class="row" style="${_(this.style)}"
    >
      <div class="cell padding-small" >
        <form-button id="pagination_btn_first" 
          .style="${{padding:"6px 6px 7px","font-size":"15px",margin:"1px 0 2px"}}"
          ?disabled="${!this.canPreviousPage}" label="1"
          @click=${()=>this._onEvent("gotoPage",1,this.canPreviousPage)} type="${n.BORDER}" >1</form-button>
        <form-button id="pagination_btn_previous" 
          .style="${{padding:"5px 6px 8px","font-size":"15px",margin:"1px 0 2px"}}"
          ?disabled="${!this.canPreviousPage}" label="&#10094;"
          @click=${()=>this._onEvent("previousPage",this.pageIndex-1,this.canPreviousPage)} type="${n.BORDER}" >&#10094;</form-button>
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
          @click=${()=>this._onEvent("nextPage",this.pageIndex+1,this.canNextPage)} type="${n.BORDER}" >&#10095;</form-button>
        <form-button id="pagination_btn_last" 
          .style="${{padding:"6px 6px 7px","font-size":"15px",margin:"1px 0 2px"}}"
          ?disabled="${!this.canNextPage}" label="${this.pageCount}"
          @click=${()=>this._onEvent("gotoPage",this.pageCount,this.canNextPage)} type="${n.BORDER}" >${this.pageCount}</form-button>
      </div>
      ${this.hidePageSize?"":i`<div class="cell padding-small hide-small" >
        <form-select id="pagination_page_size"
          label="Size"
          .style="${{padding:"7px"}}" ?disabled="${0===this.pageCount}"
          .onChange=${e=>this._onEvent("setPageSize",Number(e.value),this.pageCount>0)}
          .options=${[5,10,20,50,100].map(e=>({value:String(e),text:String(e)}))} 
          .isnull="${!1}" value="${this.pageSize}" >
        </form-select>
      </div>`}
    </div>`}});const V=e`
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
`;customElements.define("form-list",class extends t{constructor(){super(),this.id=Math.random().toString(36).slice(2),this.name=void 0,this.rows=[],this.pagination=m.TOP,this.currentPage=1,this.pageSize=10,this.hidePaginatonSize=!1,this.listFilter=!1,this.filterPlaceholder=void 0,this.filterValue="",this.labelAdd="",this.addIcon="Plus",this.editIcon="Edit",this.deleteIcon="Times",this.style={}}static get properties(){return{id:{type:String},name:{type:String},rows:{type:Array},pagination:{type:String},currentPage:{type:Number},pageSize:{type:Number},hidePaginatonSize:{type:Boolean},listFilter:{type:Boolean},filterPlaceholder:{type:String},filterValue:{type:String},labelAdd:{type:String},addIcon:{type:String},editIcon:{type:String},deleteIcon:{type:String},style:{type:Object}}}static get styles(){return[V]}connectedCallback(){super.connectedCallback(),this.currentPage>Math.ceil(this.rows.length/this.pageSize)&&(this.currentPage=Math.ceil(this.rows.length/this.pageSize)),this.currentPage<1&&(this.currentPage=1)}_onPagination(e,t){if("setPageSize"===e)return this.currentPage=1,void(this.pageSize=t);this.currentPage=t,this.onCurrentPage&&this.onCurrentPage(t)}_onEdit(e,t,i){e.stopPropagation(),this.onEdit&&this.onEdit(t,i),this.dispatchEvent(new CustomEvent("edit",{bubbles:!0,composed:!0,detail:{rowData:t,index:i}}))}_onDelete(e,t,i){e.stopPropagation(),this.onDelete&&this.onDelete(t,i),this.dispatchEvent(new CustomEvent("delete",{bubbles:!0,composed:!0,detail:{rowData:t,index:i}}))}_onAddItem(){this.onAddItem&&this.onAddItem({}),this.dispatchEvent(new CustomEvent("add_item",{bubbles:!0,composed:!0,detail:{}}))}_onFilterChange(e){this.filterValue=e.value,this.dispatchEvent(new CustomEvent("filter_change",{bubbles:!0,composed:!0,detail:{...e}}))}filterRows(){let e=this.rows;return this.listFilter&&""!==this.filterValue&&(e=e.filter(e=>((e,t)=>String(e.lslabel).toLowerCase().indexOf(t)>-1||String(e.lsvalue).toLowerCase().indexOf(t)>-1)(e,String(this.filterValue).toLowerCase()))),e}renderRows(e,t){let a=e;if(this.pagination!==m.NONE&&t>1){const e=(this.currentPage-1)*this.pageSize,t=this.currentPage*this.pageSize;a=a.slice(e,t)}return a.map((e,t)=>i`<li class="list-row border-bottom">
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
      </li>`)}render(){const e=this.filterRows(),t=Math.ceil(e.length/this.pageSize),a=t>1&&(this.pagination===m.TOP||this.pagination===m.ALL),s=t>1&&(this.pagination===m.BOTTOM||this.pagination===m.ALL);return i`<div class="responsive" >
      ${this.listFilter||a?i`<div>
        ${a?i`<div>
          <form-pagination id="${`${this.id}_top_pagination`}"
            pageIndex=${this.currentPage} pageSize=${this.pageSize} pageCount=${t} 
            ?canPreviousPage=${this.currentPage>1} 
            ?canNextPage=${this.currentPage<t} 
            ?hidePageSize=${this.hidePaginatonSize}
            .onEvent=${(e,t)=>this._onPagination(e,t)} ></form-pagination>
        </div>`:l}
        ${this.listFilter?i`<div class="row full" >
          <div class="cell" >
            <form-input id="filter" type="${h.TEXT}" 
              .style="${{"border-radius":0,margin:"1px 0 2px"}}"
              label="${x(this.filterPlaceholder)}"
              placeholder="${x(this.filterPlaceholder)}"
              value="${this.filterValue}" ?full="${!0}"
              .onChange=${e=>this._onFilterChange({value:e.value,old:this.filterValue})}
            ></form-input>
          </div>
          ${this.onAddItem?i`<div class="cell" style="${_({width:"20px"})}" >
            <form-button id="btn_add" icon="${this.addIcon}"
              label="${this.labelAdd}"
              .style="${{padding:"8px 16px","border-radius":0,margin:"1px 0 2px 1px"}}"
              @click=${()=>this._onAddItem()} type="${n.BORDER}"
            >${this.labelAdd}
            </form-button>
          </div>`:l}
        </div>`:l}
      </div>`:l}
      <ul id="${this.id}" name="${x(this.name)}"
        class="list" style="${_(this.style)}" >
        ${this.renderRows(e,t)}
      </ul>
    </div>
    ${s?i`<div>
        <form-pagination id="${`${this.id}_bottom_pagination`}"
          pageIndex=${this.currentPage} pageSize=${this.pageSize} pageCount=${t} 
          ?canPreviousPage=${this.currentPage>1} 
          ?canNextPage=${this.currentPage<t} 
          ?hidePageSize=${this.hidePaginatonSize}
          .onEvent=${(e,t)=>this._onPagination(e,t)} ></form-pagination>
    </div>`:l}
    `}});const K=e`
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
`;customElements.define("form-table",class extends t{constructor(){super(),this.id=Math.random().toString(36).slice(2),this.name=void 0,this.rows=[],this.fields={},this.rowKey="id",this.pagination=m.TOP,this.currentPage=1,this.pageSize=10,this.hidePaginatonSize=!1,this.tableFilter=!1,this.filterPlaceholder=void 0,this.filterValue="",this.labelYes="YES",this.labelNo="NO",this.labelAdd="",this.addIcon="Plus",this.tablePadding=void 0,this.sortCol=void 0,this.sortAsc=!0,this.style={}}static get properties(){return{id:{type:String},name:{type:String},rowKey:{type:String},rows:{type:Array},fields:{type:Object},pagination:{type:String},currentPage:{type:Number},pageSize:{type:Number},hidePaginatonSize:{type:Boolean},tableFilter:{type:Boolean},filterPlaceholder:{type:String},filterValue:{type:String},labelYes:{type:String},labelNo:{type:String},labelAdd:{type:String},addIcon:{type:String},tablePadding:{type:String},sortCol:{type:String,attribute:!1},sortAsc:{type:Boolean,attribute:!1},style:{type:Object}}}static get styles(){return[K]}connectedCallback(){super.connectedCallback(),this.rows||(this.rows=[]),0===Object.keys(this.fields).length&&this.rows&&Array.isArray(this.rows)&&this.rows.length>0&&Object.keys(this.rows[0]).forEach(e=>{this.fields[e]={fieldtype:"string",label:e}}),this.currentPage>Math.ceil(this.rows.length/this.pageSize)&&(this.currentPage=Math.ceil(this.rows.length/this.pageSize)),this.currentPage<1&&(this.currentPage=1)}_onPagination(e,t){if("setPageSize"===e)return this.currentPage=1,void(this.pageSize=t);this.currentPage=t,this.onCurrentPage&&this.onCurrentPage(t)}_onSort(e){e.stopPropagation();const t=e.target.dataset.sort;this.sortCol&&this.sortCol===t&&(this.sortAsc=!this.sortAsc),this.sortCol=t,this.rows.sort((e,t)=>e[this.sortCol]<t[this.sortCol]?this.sortAsc?1:-1:this.sortAsc?-1:1)}_onRowSelected(e,t,i){e.stopPropagation(),t.disabled||(this.onRowSelected&&this.onRowSelected(t,i),this.dispatchEvent(new CustomEvent("row_selected",{bubbles:!0,composed:!0,detail:{row:t,index:i}})))}_onEditCell(e,t,i,l){e.stopPropagation(),this.onEditCell&&this.onEditCell(t,i,l),this.dispatchEvent(new CustomEvent("edit_cell",{bubbles:!0,composed:!0,detail:{fieldname:t,resultValue:i,rowData:l}}))}_onAddItem(){this.onAddItem&&this.onAddItem({}),this.dispatchEvent(new CustomEvent("add_item",{bubbles:!0,composed:!0,detail:{}}))}_onFilterChange(e){this.filterValue=e.value,this.dispatchEvent(new CustomEvent("filter_change",{bubbles:!0,composed:!0,detail:{...e}}))}columns(){const e=(e,t,l)=>i`<div class="number-cell">
        <span class="cell-label">${t}</span>
        <span style="${_(l)}" >${new Intl.NumberFormat("default").format(e)}</span>
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
      `,a=(e,t,l)=>i`
        <span class="cell-label">${t}</span>
        <span style="${_(l)}" >${e}</span>
      `;let s=[];return Object.keys(this.fields).forEach(o=>{if(this.fields[o].columnDef)s=[...s,{id:o,Header:o,headerStyle:{},cellStyle:{},...this.fields[o].columnDef}];else{const n={id:o,Header:this.fields[o].label||o,headerStyle:{},cellStyle:{}};switch(this.fields[o].fieldtype){case"number":n.headerStyle.textAlign="right",n.Cell=({row:t,value:i})=>{const l={};return this.fields[o].format&&(l.fontWeight="bold",t.edited?l.textDecoration="line-through":l.color=0!==i?"red":"green"),e(i,this.fields[o].label,l)};break;case"datetime":case"date":case"time":n.Cell=({value:e})=>((e,t,l)=>{let a="";const s=new Date(e);if(s instanceof Date&&!Number.isNaN(s.valueOf()))switch(l){case"date":a=new Intl.DateTimeFormat("default",{year:"numeric",month:"2-digit",day:"2-digit"}).format(s);break;case"time":a=new Intl.DateTimeFormat("default",{hour:"2-digit",minute:"2-digit",hour12:!1}).format(s);break;default:a=new Intl.DateTimeFormat("default",{year:"numeric",month:"2-digit",day:"2-digit",hour:"2-digit",minute:"2-digit",hour12:!1}).format(s)}return i`<span class="cell-label">${t}</span><span>${a}</span>`})(e,this.fields[o].label,this.fields[o].fieldtype);break;case"bool":n.Cell=({value:e})=>t(e,this.fields[o].label);break;case"deffield":n.Cell=({row:i,value:s})=>{switch(i.fieldtype){case"bool":return t(s,this.fields[o].label);case"integer":case"float":return e(s,this.fields[o].label,{});case"customer":case"tool":case"product":case"trans":case"transitem":case"transmovement":case"transpayment":case"project":case"employee":case"place":case"urlink":return l(i.export_deffield_value,this.fields[o].label,i.fieldtype,i[o],i);default:return a(s,this.fields[o].label,{})}};break;default:n.Cell=({row:e,value:t})=>{const i={};return e[`${o}_color`]&&(i.color=e[`${o}_color`]),Object.keys(e).includes(`export_${o}`)?l(e[`export_${o}`],this.fields[o].label,o,e[o],e):a(t,this.fields[o].label,i)}}this.tablePadding&&(n.headerStyle.padding=this.tablePadding,n.cellStyle.padding=this.tablePadding),this.fields[o].verticalAlign&&(n.cellStyle.verticalAlign=this.fields[o].verticalAlign),this.fields[o].textAlign&&(n.cellStyle.textAlign=this.fields[o].textAlign),s=[...s,n]}}),s}renderHeader(e){return i`<thead><tr>
      ${e.map(e=>i`<th 
        data-sort="${e.id}" 
        class="sort ${this.sortCol===e.id?this.sortAsc?"sort-asc":"sort-desc":"sort-none"}"
        style="${_(e.headerStyle)}"
        @click=${this._onSort} >${e.Header}</th>`)}
    </tr></thead>`}filterRows(){let e=this.rows;const t=(e,t)=>{let i=!1;return Object.keys(this.fields).forEach(l=>{String(e[l]).toLowerCase().indexOf(t)>-1&&(i=!0)}),i};return this.tableFilter&&""!==this.filterValue&&(e=e.filter(e=>t(e,String(this.filterValue).toLowerCase()))),e}renderRows(e,t,l){let a=t;if(this.pagination!==m.NONE&&l>1){const e=(this.currentPage-1)*this.pageSize,t=this.currentPage*this.pageSize;a=a.slice(e,t)}return a.map((t,l)=>i`<tr id="${`row_${t[this.rowKey]||l}`}"
        class="${t.disabled?"cursor-disabled":this.onRowSelected?"cursor-pointer":""}"
        @click=${e=>this._onRowSelected(e,t,l)}
      >${e.map(e=>i`<td style="${_(e.cellStyle)}">${e.Cell?e.Cell({row:t,value:t[e.id]}):t[e.id]}</td>`)}</tr>`)}render(){const e=this.columns(),t=this.filterRows(),a=Math.ceil(t.length/this.pageSize),s=a>1&&(this.pagination===m.TOP||this.pagination===m.ALL),o=a>1&&(this.pagination===m.BOTTOM||this.pagination===m.ALL);return i`<div class="responsive" >
      ${this.tableFilter||s?i`<div>
        ${s?i`<div>
          <form-pagination id="${`${this.id}_top_pagination`}"
            pageIndex=${this.currentPage} pageSize=${this.pageSize} pageCount=${a} 
            ?canPreviousPage=${this.currentPage>1} 
            ?canNextPage=${this.currentPage<a} 
            ?hidePageSize=${this.hidePaginatonSize}
            .onEvent=${(e,t)=>this._onPagination(e,t)} ></form-pagination>
        </div>`:l}
        ${this.tableFilter?i`<div class="row full" >
          <div class="cell" >
            <form-input id="filter" type="${h.TEXT}" 
              .style="${{"border-radius":0,margin:"1px 0 2px"}}"
              label="${x(this.filterPlaceholder)}"
              placeholder="${x(this.filterPlaceholder)}"
              value="${this.filterValue}" ?full="${!0}"
              .onChange=${e=>this._onFilterChange({value:e.value,old:this.filterValue})}
            ></form-input>
          </div>
          ${this.onAddItem?i`<div class="cell" style="${_({width:"20px"})}" >
            <form-button id="btn_add" icon="${this.addIcon}"
              label="${this.labelAdd}"
              .style="${{padding:"8px 16px","border-radius":0,margin:"1px 0 2px 1px"}}"
              @click=${()=>this._onAddItem()} type="${n.BORDER}"
            >${this.labelAdd}</form-button>
          </div>`:l}
        </div>`:l}
      </div>`:l}
      <div class="table-wrap" >
        <table id="${this.id}" name="${x(this.name)}"
          class="ui-table" style="${_(this.style)}" >
          ${this.renderHeader(e)}
          <tbody>${this.renderRows(e,t,a)}</tbody>
        </table>
      </div>
    </div>
    ${o?i`<div>
        <form-pagination id="${`${this.id}_bottom_pagination`}"
          pageIndex=${this.currentPage} pageSize=${this.pageSize} pageCount=${a} 
          ?canPreviousPage=${this.currentPage>1} 
          ?canNextPage=${this.currentPage<a} 
          ?hidePageSize=${this.hidePaginatonSize}
          .onEvent=${(e,t)=>this._onPagination(e,t)} ></form-pagination>
    </div>`:l}
    `}});const L=e`
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
.panel {
  background-color: rgba(var(--base-2), 1);
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  box-shadow:0 2px 5px 0 rgba(0,0,0,0.16),0 2px 10px 0 rgba(0,0,0,0.12);
  border-radius: 4px;
  max-width: 800px;
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
.section-container {
  padding: 16px;
  display: table;
  width: 100%;
}
.section-container-small {
  display: table;
  width: 100%;
  padding: 8px;
}
.border {
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
}
.padding-small { 
  padding: 4px 8px; 
}
.third{ 
  width: 33.33333%;
  vertical-align: top;
}
.mapBox {
  width: 100%;
  text-align: center;
  padding: 10px 5px;
  box-sizing: border-box;
  margin-top: 5px;
	margin-bottom: 5px;
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  background-color: rgba(var(--neutral-1), 0.1);
}
.reportMap {
	border: 1px solid #666; 
	background-color: white;
	width: 110px; 
	height: 165px;
}
.separator {
  height: 10px;
  width: 100%;
  border: none;
  margin: 0;
}
.report-title {
  width: 100%;
  border: 1.5px solid rgb(var(--functional-yellow));
  margin-top: -1px;
}
.template-row {
  display: table;
  width: 100%;
  border-left: 0.5px solid rgba(var(--neutral-1), 0.2);
  border-right: 0.5px solid rgba(var(--neutral-1), 0.2);
}
.report-title-label {
  color: rgb(var(--functional-yellow));
  fill: rgb(var(--functional-yellow));
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
textarea {
  font-family: var(--font-family);
  font-size: 12px;
  width: 100%;
  border-radius: 3px;
  overflow: auto;
  padding: 8px;
  margin-top: 8px;
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
.meta-title-row {
  display: table;
  border-bottom: 0.5px solid rgba(var(--neutral-1), 0.2);
  margin: auto;
}
.meta-title-cell {
  display: table-cell;
  vertical-align: top;
  padding: 8px;
  font-size: 12px;
}
.bold {
  font-weight: bold;
}
.meta-title-sources {
  display: table;
  width: 100%;
  padding: 16px 0 0px;
}
.meta-sources {
  display: table;
  width: 100%;
  padding: 8px 0;
}
.meta-sources-name {
  border-top: 0.5px solid rgba(var(--neutral-1), 0.2);
  border-bottom: 0.5px solid rgba(var(--neutral-1), 0.2);
  font-weight: bold;
  font-style: italic;
}
.meta-sources-cell {
  display: table-cell;
  vertical-align: top;
  font-size: 10px;
}
@media (max-width:600px){
  .section-container { 
    padding: 8px; 
  }
  .section-container-small{
    padding: 4px; 
  }
  .padding-small { 
    padding: 2px 4px; 
  }
  .third{ 
    width: 100%;
    display: block;
  }
  .meta-title-row{ 
    display:block; 
    width:100%;
  }
  .meta-title-cell {
    display:block; 
    width:100%;
    padding: 4px;
  }
}
`,j=e=>Object.getOwnPropertyNames(e)[0];customElements.define("template-editor",class extends t{constructor(){super(),this.msg=e=>e,this.id=Math.random().toString(36).slice(2),this.data={title:"",tabView:"template",template:{meta:{},report:{},header:[],details:[],footer:[],sources:{},data:{}},current:{},current_data:null,dataset:[]},this.paginationPage=10,this.onEvent={}}static get properties(){return{id:{type:String},data:{type:Object},paginationPage:{type:Number},onEvent:{type:Object}}}static get styles(){return[L]}_onTemplateEvent(e,t){this.onEvent.onTemplateEvent&&this.onEvent.onTemplateEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("template_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}canvasCallback(e){e&&this._onTemplateEvent(v.CREATE_MAP,{mapRef:e})}tabButton(e,t){const{tabView:l}=this.data;return i`<form-button 
      id="${`btn_${e}`}"
      .style="${{"border-radius":0}}" icon="${t}"
      label="${this.msg("",{id:`template_label_${e}`})}"
      @click=${()=>this._onTemplateEvent(v.CHANGE_TEMPLATE,{key:"tabView",value:e})} 
      type="${l===e?n.PRIMARY:""}"
      ?full="${!0}" ?selected="${l===e}" >
      ${this.msg("",{id:`template_label_${e}`})}</form-button>`}navButton(e,t,l,a,s,d){const r=void 0===d?o.LEFT:d,c=void 0===s||s;return i`<form-button 
      id="${`btn_${e}`}"
      .style="${{"border-radius":0}}" icon="${a}"
      label="${this.msg("",{id:l})}" align="${r}"
      @click=${()=>this._onTemplateEvent(...t)} 
      type="${n.BORDER}" ?full="${c}" >${this.msg("",{id:l})}</form-button>`}setListIcon(e,t){const{template:i,current:l}=this.data,a=e=>void 0!==t?t:"object"==typeof e&&Array.isArray(e)&&e.length>0?e.length:0;if(l.item===e)return{selected:!0,icon:"Tag",color:"green",badge:a(e)};if(l.parent===e||i[l.section]===e)return{selected:!0,icon:"Check",color:"",badge:a(e)};let s="InfoCircle";return Array.isArray(e)&&e.length>0&&(s="Plus"),{selected:!1,icon:s,color:"",badge:0}}mapButton(e,t,l,a,s){const d=s>0?s:void 0;return i`<form-button 
      id="${`btn_${e}`}"
      .style="${{"border-radius":0}}" icon="${l.icon}"
      ?selected="${""!==l.color}"
      label="${a}" align="${o.LEFT}"
      @click=${()=>this._onTemplateEvent(v.SET_CURRENT,[{tmp_id:t}])}
      badge="${x(d)}"
      type="${n.BORDER}" ?full="${!0}" >${a}</form-button>`}createSubList(e){const{template:t,current:l}=this.data;for(let a=0;a<t[l.section].length;a+=1){const s=j(t[l.section][a]);let o=t[l.section][a][s];const n=`tmp_${l.section}_${a.toString()}_${s}`;if(["row","datagrid"].includes(s)&&(o=o.columns),null===l.parent){const t=this.setListIcon(o);e.push(i`<div key={mkey}>
          ${this.mapButton(n,n,t,s.toUpperCase(),a+1)}
        </div>`)}else if(l.item===o||l.parent===o){const t=this.setListIcon(o,a+1);if(e.push(i`<div key={mkey}>
            ${this.mapButton(n,n,t,s.toUpperCase(),t.badge)}
          </div>`),["row","datagrid"].includes(l.type)||["row","datagrid"].includes(l.parent_type))for(let t=0;t<o.length;t+=1){const n=j(o[t]),d=o[t][n],r=`tmp_${l.section}_${a.toString()}_${s}_${t.toString()}_${n}`,c=this.setListIcon(d);e.push(i`<div key={skey}>
                ${this.mapButton(r,r,c,n.toUpperCase(),c.badge,`primary ${L.badgeBlack}`)}
              </div>`)}}}}createMapList(){const{template:e}=this.data,t=[];return["report","header","details","footer"].forEach(l=>{const a=this.setListIcon(e[l]);a.selected&&"report"!==l&&t.push(i`<hr id="${`sep_${l}_0`}" class="separator" />`),t.push(i`<div key={mkey}>
        ${this.mapButton(`tmp_${l}`,`tmp_${l}`,a,l.toUpperCase(),a.badge)}
      </div>`),a.selected&&(t.push(i`<hr id="${`sep_${l}_1`}" class="separator" />`),"report"!==l&&(this.createSubList(t),t.push(i`<hr id="${`sep_${l}_2`}" class="separator" />`)))}),t}dataTitle(e,t,l){return i`<div class="panel-title">
      <div class="cell">
        <form-label class="title-cell"
          value="${t}"
        ></form-label>
      </div>
      <div class="cell align-right" >
        <span id="${e}" class="close-icon" 
          @click="${()=>this._onTemplateEvent(...l)}">
          <form-icon iconKey="Times" ></form-icon>
        </span>
      </div>
    </div>`}dataText(e,t,l,a){return i`<textarea id="${`${e}_value`}"
      rows=${l} .value="${t}"
      @input="${e=>this._onTemplateEvent(v.EDIT_DATA_ITEM,{...a,value:e.target.value})}"
    ></textarea>`}tableFields(){const{current_data:e}=this.data;return{...e.fields,edit:{columnDef:{id:"delete",Header:"",headerStyle:{},Cell:({row:e})=>i`<form-icon id=${`delete_${e._index}`}
            iconKey="Times" width=19 height=27.6
            .style=${{cursor:"pointer",fill:"rgb(var(--functional-red))"}}
            @click=${t=>{t.stopPropagation(),this._onTemplateEvent(v.DELETE_DATA_ITEM,{_index:e._index})}}
          ></form-icon>`,cellStyle:{width:40,padding:"4px 8px 3px 8px"}}}}}render(){const{title:e,tabView:t,template:a,current:s,current_data:n,dataset:d}=this.data,r=(e,t)=>({map_edit:{data:!1,report:!1,header:!1,footer:!1,details:!1},map_insert:{header:!0,footer:!0,details:!0,row:!0,datagrid:!0}}[t][e]);return i`<div class="panel" >
      <div class="panel-title">
        <div class="cell">
          <form-label class="title-cell"
            value="${e}" leftIcon="TextHeight"
          ></form-label>
        </div>
      </div>
      <div class="section-container" >
        <div class="row full">
          <div class="cell third">${this.tabButton("template","Tags")}</div>
          <div class="cell third">${this.tabButton("data","Database")}</div>
          <div class="cell third">${this.tabButton("meta","InfoCircle")}</div>
        </div>
        ${"template"===t?i`<div class="section-container-small border" >
          <div class="cell padding-small third" >
            ${this.navButton("previous",[v.GO_PREVIOUS,[]],"label_previous","ArrowLeft")}
            <div class="mapBox" >
              <canvas ${I(this.canvasCallback)} class="reportMap" ></canvas>
            </div>
            ${this.navButton("next",[v.GO_NEXT,[]],"label_next","ArrowRight",!0,o.RIGHT)}
          </div>
          <div class="cell padding-small third" >
            ${this.createMapList()}
          </div>
          <div class="cell padding-small third" >
            ${!1!==r(s.type,"map_edit")?i`<div>
                ${this.navButton("move_up",[v.MOVE_UP,[]],"label_move_up","ArrowUp")}
                ${this.navButton("move_down",[v.MOVE_DOWN,[]],"label_move_down","ArrowDown")}
                ${this.navButton("delete_item",[v.DELETE_ITEM,[]],"label_delete","Times")}
                <hr class="separator" />
              </div>`:l}
            ${r(s.type,"map_insert")?i`<div>
                ${this.navButton("add_item",[v.ADD_ITEM,s.add_item||""],"label_add_item","Plus")}
                <form-select id="sel_add_item" label="" 
                  ?full=${!1} .isnull="${!0}" value="${s.add_item||""}"
                  .onChange=${e=>this._onTemplateEvent(v.CHANGE_CURRENT,{key:"add_item",value:e.value})}
                  .options=${{header:["row","vgap","hline"],details:["row","vgap","hline","html","datagrid"],footer:["row","vgap","hline"],row:["cell","image","barcode","separator"],datagrid:["column"]}[s.type].map(e=>({value:e,text:e.toUpperCase()}))}  
                ></form-select>
              </div>`:l}
          </div>
        </div>
        <div class="report-title padding-small" >
          <form-label class="report-title-label"
            value="${s.type.toUpperCase()}" leftIcon="Tag"
          ></form-label>
        </div>
        ${s.form.rows.map((e,t)=>i`<div class="template-row" >
          <form-row id=${`row_${t}`}
            .row=${e} 
            .values=${["row","datagrid"].includes(s.type)?s.item_base:s.item}
            .options=${s.form.options}
            .data=${{audit:"all",current:s,dataset:a.data}}
            .onEdit=${e=>this._onTemplateEvent(v.EDIT_ITEM,e)}
            .msg=${this.msg}
          ></form-row>
        </div>`)}`:l}
        ${"data"===t?i`<div class="section-container border" >
          ${n&&"string"===n.type?i`<div class="row full section-small">
              ${this.dataTitle("data_string",n.name,[v.SET_CURRENT_DATA,null])}
              ${this.dataText(n.name,a.data[n.name],15,{})}
          </div>`:l}
          ${n&&"list"===n.type&&n.item?i`<div class="row full section-small">
              ${this.dataTitle("data_list_item",n.item,[v.SET_CURRENT_DATA_ITEM,null])}
              ${this.dataText(n.item,a.data[n.name][n.item],10,{})}
          </div>`:l}
          ${n&&"table"===n.type&&n.item?i`<div class="row full section-small">
              ${this.dataTitle("data_table_item",`${n.name} - ${String(n.item._index+1)}`,[v.SET_CURRENT_DATA_ITEM,null])}
              ${Object.keys(n.fields).map(e=>i`<div class="row full">
                <div class="padding-small">
                  <form-label value="${e}"></form-label>
                </div>
                ${this.dataText(e,n.item[e],2,{field:e,_index:n.item._index})}
              </div>`)}
          </div>`:l}
          ${n&&"list"===n.type&&!n.item?i`<div class="row full section-small">
              ${this.dataTitle("data_list",n.name,[v.SET_CURRENT_DATA,null])}
              <form-list id="data_list_items"
                .rows=${n.items} ?listFilter=${!0}
                filterPlaceholder=${this.msg("",{id:"placeholder_filter"})}
                .onAddItem=${()=>this._onTemplateEvent(v.SET_CURRENT_DATA_ITEM,void 0)}
                labelAdd=${this.msg("",{id:"label_new"})}
                pageSize=${this.paginationPage} pagination="${m.TOP}" 
                .onEdit=${e=>this._onTemplateEvent(v.SET_CURRENT_DATA_ITEM,e.lslabel)}
                .onDelete=${e=>this._onTemplateEvent(v.DELETE_DATA_ITEM,{key:e.lslabel})}
              ></form-list>
          </div>`:l}
          ${n&&"table"===n.type&&!n.item?i`<div class="row full section-small">
              ${this.dataTitle("data_table",n.name,[v.SET_CURRENT_DATA,null])}
              <form-table id="data_table_items"
                .fields=${this.tableFields()} .rows=${n.items} ?tableFilter=${!0}
                filterPlaceholder="${this.msg("",{id:"placeholder_filter"})}"
                .onAddItem=${()=>this._onTemplateEvent(v.SET_CURRENT_DATA_ITEM,void 0)}
                .onRowSelected=${e=>this._onTemplateEvent(v.SET_CURRENT_DATA_ITEM,e)} 
                labelAdd=${this.msg("",{id:"label_new"})}  
                pageSize=${this.paginationPage} pagination="${m.TOP}"
              ></form-table>
          </div>`:l}
          ${n?l:i`<div class="row full section-small">
              <form-list id="data_list_items"
                .rows=${d} ?listFilter=${!0}
                filterPlaceholder=${this.msg("",{id:"placeholder_filter"})}
                .onAddItem=${()=>this._onTemplateEvent(v.ADD_TEMPLATE_DATA,[])}
                labelAdd=${this.msg("",{id:"label_new"})}
                pageSize=${this.paginationPage} pagination="${m.TOP}" 
                .onEdit=${e=>this._onTemplateEvent(v.SET_CURRENT_DATA,{name:e.lslabel,type:e.lsvalue})}
                .onDelete=${e=>this._onTemplateEvent(v.DELETE_DATA,e.lslabel)}
              ></form-list>
          </div>`}
        </div>`:l}
        ${"meta"===t?i`<div class="section-container-small border" >
          <div class="cell padding-small" >
            <div class="meta-title-row" >
              ${Object.keys(a.meta).filter(e=>["report_key","report_name","description"].includes(e)).map(e=>i`<div class="meta-title-cell" >
                <div class="bold">${e}</div>
                <div>${a.meta[e]}</div>
              </div>`)}
            </div>
            <div class="meta-title-row" >
              ${Object.keys(a.meta).filter(e=>!["report_key","report_name","description"].includes(e)).map(e=>i`<div class="meta-title-cell" >
                <div class="bold">${e}</div>
                <div>${a.meta[e]}</div>
              </div>`)}
            </div>
            <div class="meta-title-sources" >
              <form-label value="${this.msg("",{id:"template_data_sources"})}"></form-label>
            </div>
            ${Object.keys(a.sources).map(e=>i`<div class="meta-sources" >
              <div class="cell padding-small">
                <div class="meta-sources-name padding-small">
                  ${e}
                </div>
                ${Object.keys(a.sources[e]).map(t=>i`<div class="row" >
                  <div class="meta-sources-cell padding-small bold" >${t}:</div>
                  <div class="meta-sources-cell padding-small" >${a.sources[e][t]}</div>
                </div>`)}
              </div>
            </div>`)}
          </div>
        </div>`:l}
      </div>
    </div>`}});const F=e`
@keyframes animatezoom{from{transform:scale(0)} to{transform:scale(1)}}
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
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  box-shadow:0 4px 10px 0 rgba(0,0,0, 0.2),0 4px 20px 0 rgba(0,0,0, 0.19);
  background-color: rgba(var(--base-2), 1);
  margin: 0 auto;
  animation: animatezoom 0.6s;
  width: 100%;
  max-width: 400px;
  min-width: 280px;
}
.panel {
  box-shadow:0 2px 5px 0 rgba(0,0,0,0.16),0 2px 10px 0 rgba(0,0,0,0.12);
  border-radius: 4px;
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
.section {
  padding: 16px 0; 
}
.section-row {
  display: table;
  width: 100%;
  padding: 0 16px;
}
.half { 
  width:100% 
}
.padding-small { 
  padding: 4px 8px; 
}
.buttons {
  background-color: rgb(var(--base-1));
  border-top: 0.5px solid rgba(var(--neutral-1), 0.2);
}
.input {
  font-weight: bold;
  font-size: 14px!important;
}

.info {
  font-size: 14px!important;
  padding-top: 8px;
}
@media (max-width:600px){
  .section-row { 
    padding: 0 8px 8px; 
  }
  .padding-small { 
    padding: 2px 4px; 
  }
}
@media only screen and (min-width: 601px){
  .dialog {
    min-width: 400px;
  }
  .half { 
    width:49.99999% 
  }
}
`;customElements.define("modal-inputbox",class extends t{constructor(){super(),this.title="",this.message="",this.infoText=void 0,this.value="",this.labelCancel="Cancel",this.labelOK="OK",this.defaultOK=!1,this.showValue=!1,this.values={}}static get properties(){return{title:{type:String},message:{type:String},infoText:{type:String},value:{type:String,reflect:!0},labelOK:{type:String},labelCancel:{type:String},defaultOK:{type:Boolean},showValue:{type:Boolean},values:{type:Object}}}static get styles(){return[F]}_onModalEvent(e){const t={value:(this.renderRoot.querySelector("#input_value")||{}).value,values:this.values};this.onEvent&&this.onEvent.onModalEvent&&this.onEvent.onModalEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("modal_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}render(){return i`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label value="${this.title}" class="title-cell" ></form-label>
            </div>
          </div>
          <div class="section" >
            <div class="section-row" >
              <div class="input">${this.message}</div>
              ${this.infoText?i`<div class="info">${this.infoText}</div>`:l}
              ${this.showValue?i`<div class="info">
                  <form-input id="input_value" type="${h.TEXT}" label="${this.title}"
                    value="${this.value}" ?full="${!0}"
                    .onEnter=${()=>this._onModalEvent(g.OK)}
                  ></form-input>
                </div>`:l}
            </div>
          </div>
          <div class="section buttons" >
            <div class="section-row" >
              <div class="cell padding-small half" >
                <form-button id="btn_cancel" icon="Times"
                  @click=${()=>this._onModalEvent(g.CANCEL)} 
                  ?full="${!0}" label="${this.labelCancel}"
                >${this.labelCancel}</form-button>
              </div>
              <div class="cell padding-small half" >
                <form-button id="btn_ok" icon="Check"
                  @click=${()=>this._onModalEvent(g.OK)} 
                  ?autofocus="${!this.showValue&&this.defaultOK}"
                  type="${n.PRIMARY}" ?full="${!0}" label="${this.labelOK}"
                >${this.labelOK}</form-button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>`}});const U=e`
@keyframes animatezoom{from{transform:scale(0)} to{transform:scale(1)}}
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
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  box-shadow:0 4px 10px 0 rgba(0,0,0, 0.2),0 4px 20px 0 rgba(0,0,0, 0.19);
  background-color: rgba(var(--base-2), 1);
  margin: 0 auto;
  animation: animatezoom 0.6s;
  width: 100%;
  max-width: 400px;
  min-width: 280px;
}
.panel {
  box-shadow:0 2px 5px 0 rgba(0,0,0,0.16),0 2px 10px 0 rgba(0,0,0,0.12);
  border-radius: 4px;
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
.section {
  padding: 16px 0; 
}
.section-row {
  display: table;
  width: 100%;
  padding: 0 8px;
}
.half { 
  width:100% 
}
.padding-small { 
  padding: 4px 8px; 
}
.buttons {
  background-color: rgb(var(--base-1));
  border-top: 0.5px solid rgba(var(--neutral-1), 0.2);
}
.close-icon {
  fill: rgba(var(--accent-1c), 0.85);
  cursor: pointer;
}
.close-icon form-icon:hover {
  fill: rgb(var(--functional-red));
}
textarea {
  font-family: var(--font-family);
  font-size: 12px;
  width: 100%;
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
@media (max-width:600px){
  .section-row { 
    padding: 0 4px 8px; 
  }
  .padding-small { 
    padding: 2px 4px; 
  }
}
@media only screen and (min-width: 601px){
  .dialog {
    min-width: 400px;
  }
  .half { 
    width:49.99999% 
  }
}
`;customElements.define("modal-template",class extends t{constructor(){super(),this.msg=e=>e,this.type=f.TEXT,this.name="",this.columns=""}static get properties(){return{type:{type:String},name:{type:String},columns:{type:String}}}static get styles(){return[U]}_onModalEvent(e,t){this.onEvent&&this.onEvent.onModalEvent&&this.onEvent.onModalEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("modal_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}_onValueChange(e,t){this[e]=t}_onTextInput(e){this._onValueChange("columns",e.target.value)}render(){const{type:e,name:t,columns:a}=this;return i`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label leftIcon="Plus"
                value="${this.msg("",{id:"template_label_new_data"})}" 
                class="title-cell" ></form-label>
            </div>
            <div class="cell align-right" >
              <span id=${"closeIcon"} class="close-icon" 
                @click="${()=>this._onModalEvent(g.CANCEL,{})}">
                <form-icon iconKey="Times" ></form-icon>
              </span>
            </div>
          </div>
          <div class="section" >
            <div class="section-row" >
              <div class="cell padding-small half" >
                <div>
                  <form-label
                    value="${this.msg("",{id:"template_data_type"})}" 
                  ></form-label>
                </div>
                <form-select id="type" 
                  label="${this.msg("",{id:"template_data_type"})}" 
                  ?full=${!0} .isnull="${!1}" value="${e}"
                  .onChange=${e=>this._onValueChange("type",e.value)}
                  .options=${Object.keys(f).map(e=>({value:f[e],text:e}))}  
                ></form-select>
              </div>
              <div class="cell padding-small half" >
                <div>
                  <form-label
                    value="${this.msg("",{id:"template_data_name"})}" 
                  ></form-label>
                </div>
                <form-input id="name"
                  type="${h.TEXT}"
                  label="${this.msg("",{id:"template_data_name"})}" 
                  .onChange=${e=>this._onValueChange("name",e.value)}
                  value="${t}" ?full=${!0}
                ></form-input>
              </div>
            </div>
            ${e===f.TABLE?i`<div class="section-row" >
              <div class="cell padding-small" >
                <div>
                  <form-label
                    value="${this.msg("",{id:"template_data_columns"})}" 
                  ></form-label>
                </div>
                  <textarea id="columns"
                    rows=3 .value="${a}"
                    @input="${this._onTextInput}"
                  ></textarea>
              </div>
            </div>`:l}
          </div>
          <div class="section buttons" >
            <div class="section-row" >
              <div class="cell padding-small half" >
                <form-button id="btn_cancel" icon="Times"
                  @click=${()=>this._onModalEvent(g.CANCEL,{})} 
                  ?full="${!0}" label="${this.msg("",{id:"msg_cancel"})}"
                >${this.msg("",{id:"msg_cancel"})}</form-button>
              </div>
              <div class="cell padding-small half" >
                <form-button id="btn_ok" icon="Check"
                  @click=${()=>this._onModalEvent(g.OK,{value:{name:t,type:e,columns:a}})} 
                  type="${n.PRIMARY}" ?full="${!0}" 
                  label="${this.msg("",{id:"msg_ok"})}"
                >${this.msg("",{id:"msg_ok"})}</form-button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>`}});const H=e`
:host {
  font-family: var(--font-family);
  font-size: var(--font-size);
  color: var(--text-1);
  fill: var(--text-1);
}
div {
  box-sizing: border-box;
}
.page {
  transition: margin-left .4s;
  margin-left: var(--menu-side-width);
  padding: 8px;
}
@media (max-width:768px){
  .page{
    margin-left: 0px;
  }
}
@media (max-width:600px){
  .page {
    padding: 4px;
  }
}
`;customElements.define("client-template",class extends t{constructor(){super(),this.msg=e=>e,this.id=Math.random().toString(36).slice(2),this.side=a.AUTO,this.data={},this.paginationPage=10,this.theme=s.LIGHT,this.onEvent={},this.modalTemplate=this.modalTemplate.bind(this)}static get properties(){return{id:{type:String},side:{type:String},data:{type:Object},paginationPage:{type:Number},onEvent:{type:Object}}}static get styles(){return[H]}connectedCallback(){super.connectedCallback(),this.onEvent.setModule(this)}modalTemplate({type:e,name:t,columns:l,onEvent:a}){return i`<modal-template
      type="${e}"
      name="${t}"
      columns="${l}"
      .onEvent=${a} .msg=${this.msg}
    ></modal-template>`}render(){const{side:e,data:t,paginationPage:l,theme:a}=this;return i`<sidebar-template
      id="${this.id}" side="${e}"
      templateKey="${t.key}" ?dirty="${t.dirty}" theme="${a}"
      .onEvent=${this.onEvent} .msg=${this.msg}
    ></sidebar-template>
      <div class="page">
        <template-editor
          .data="${t}"
          .onEvent=${this.onEvent} .msg=${this.msg}
          paginationPage="${l}"
        ></template-editor>
      </div>`}});
