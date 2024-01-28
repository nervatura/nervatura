import{i as e,s as t,x as a,T as i}from"./module-3zU3FH2L.js";import"./module-7N9Fn5lU.js";import{d as l,e as s,B as o,S as r,m as d,P as n,n as p,M as c,I as m}from"./main-HhCNuxlS.js";import{o as v}from"./module-JEOuC3n3.js";import{n as b}from"./module-ChuqIslB.js";import"./module-67w4bJ4f.js";import"./module-AQ0TXxQm.js";import"./module-r-Be6lM0.js";import"./module-3cipbsnY.js";const h=e`
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
`;customElements.define("sidebar-template",class extends t{constructor(){super(),this.msg=e=>e,this.side=l.AUTO,this.templateKey="",this.dirty=!1}static get properties(){return{side:{type:String,reflect:!0},templateKey:{type:String},dirty:{type:Boolean}}}static get styles(){return[h]}_onSideEvent(e,t){this.onEvent&&this.onEvent.onSideEvent&&this.onEvent.onSideEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("side_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}itemMenu({id:e,selected:t,eventValue:i,label:l,iconKey:r,full:d}){return a`<form-button 
        id="${e}" label="${l}"
        ?full="${void 0===d||d}" ?selected="${void 0!==t&&t}"
        align=${s.LEFT}
        .style="${{"border-radius":"0","border-color":"rgba(var(--accent-1c), 0.2)"}}"
        icon="${r}" type="${o.PRIMARY}"
        @click=${()=>this._onSideEvent(...i)} 
      >${l}</form-button>`}formItems(){const e=[];return e.push(this.itemMenu({id:"cmd_back",selected:!0,eventValue:[r.CHECK,{value:r.LOAD_SETTING}],label:this.msg("",{id:"label_back"}),iconKey:"Reply",full:!1})),e.push(a`<hr id="back_sep" class="separator" />`),["_blank","_sample"].includes(this.templateKey)||(e.push(a`<hr id="tmp_sep_2" class="separator" />`),e.push(this.itemMenu({id:"cmd_save",selected:this.dirty,eventValue:[r.SAVE,!0],label:this.msg("",{id:"template_save"}),iconKey:"Check"})),e.push(this.itemMenu({id:"cmd_create",eventValue:[r.CREATE_REPORT,{}],label:this.msg("",{id:"template_create_from"}),iconKey:"Sitemap"})),e.push(this.itemMenu({id:"cmd_delete",eventValue:[r.DELETE,{}],label:this.msg("",{id:"label_delete"}),iconKey:"Times"}))),e.push(a`<hr id="tmp_sep_3" class="separator" />`),e.push(this.itemMenu({id:"cmd_blank",eventValue:[r.CHECK,{value:r.BLANK}],label:this.msg("",{id:"template_new_blank"}),iconKey:"Plus"})),e.push(this.itemMenu({id:"cmd_sample",eventValue:[r.CHECK,{value:r.SAMPLE}],label:this.msg("",{id:"template_new_sample"}),iconKey:"Plus"})),e.push(a`<hr id="tmp_sep_4" class="separator" />`),e.push(this.itemMenu({id:"cmd_print",eventValue:[r.REPORT_SETTINGS,{value:"PREVIEW"}],label:this.msg("",{id:"label_print"}),iconKey:"Eye"})),e.push(this.itemMenu({id:"cmd_json",eventValue:[r.REPORT_SETTINGS,{value:"JSON"}],label:this.msg("",{id:"template_export_json"}),iconKey:"Code"})),e.push(a`<hr id="tmp_sep_5" class="separator" />`),e.push(this.itemMenu({id:"cmd_help",eventValue:[r.HELP,{value:"program/editor"}],label:this.msg("",{id:"label_help"}),iconKey:"QuestionCircle"})),e}render(){return a`<div class="sidebar ${"auto"!==this.side?this.side:""}" >
    ${this.formItems()}
    </div>`}});const u=e`
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
  border-top: 0.5px solid rgba(var(--neutral-1), 0.2);
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
`,g=e=>Object.getOwnPropertyNames(e)[0];customElements.define("template-editor",class extends t{constructor(){super(),this.msg=e=>e,this.id=Math.random().toString(36).slice(2),this.data={title:"",tabView:"template",template:{meta:{},report:{},header:[],details:[],footer:[],sources:{},data:{}},current:{},current_data:null,dataset:[]},this.paginationPage=10,this.onEvent={}}static get properties(){return{id:{type:String},data:{type:Object},paginationPage:{type:Number},onEvent:{type:Object}}}static get styles(){return[u]}_onTemplateEvent(e,t){this.onEvent.onTemplateEvent&&this.onEvent.onTemplateEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("template_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}canvasCallback(e){e&&this._onTemplateEvent(d.CREATE_MAP,{mapRef:e})}tabButton(e,t){const{tabView:i}=this.data;return a`<form-button 
      id="${`btn_${e}`}"
      .style="${{"border-radius":0}}" icon="${t}"
      label="${this.msg("",{id:`template_label_${e}`})}"
      @click=${()=>this._onTemplateEvent(d.CHANGE_TEMPLATE,{key:"tabView",value:e})} 
      type="${i===e?o.PRIMARY:""}"
      ?full="${!0}" ?selected="${i===e}" >
      ${this.msg("",{id:`template_label_${e}`})}</form-button>`}navButton(e,t,i,l,r,d){const n=void 0===d?s.LEFT:d,p=void 0===r||r;return a`<form-button 
      id="${`btn_${e}`}"
      .style="${{"border-radius":0}}" icon="${l}"
      label="${this.msg("",{id:i})}" align="${n}"
      @click=${()=>this._onTemplateEvent(...t)} 
      type="${o.BORDER}" ?full="${p}" >${this.msg("",{id:i})}</form-button>`}setListIcon(e,t){const{template:a,current:i}=this.data,l=e=>void 0!==t?t:"object"==typeof e&&Array.isArray(e)&&e.length>0?e.length:0;if(i.item===e)return{selected:!0,icon:"Tag",color:"green",badge:l(e)};if(i.parent===e||a[i.section]===e)return{selected:!0,icon:"Check",color:"",badge:l(e)};let s="InfoCircle";return Array.isArray(e)&&e.length>0&&(s="Plus"),{selected:!1,icon:s,color:"",badge:0}}mapButton(e,t,i,l,r){const n=r>0?r:void 0;return a`<form-button 
      id="${`btn_${e}`}"
      .style="${{"border-radius":0}}" icon="${i.icon}"
      ?selected="${""!==i.color}"
      label="${l}" align="${s.LEFT}"
      @click=${()=>this._onTemplateEvent(d.SET_CURRENT,[{tmp_id:t}])}
      badge="${v(n)}"
      type="${o.BORDER}" ?full="${!0}" >${l}</form-button>`}createSubList(e){const{template:t,current:i}=this.data;for(let l=0;l<t[i.section].length;l+=1){const s=g(t[i.section][l]);let o=t[i.section][l][s];const r=`tmp_${i.section}_${l.toString()}_${s}`;if(["row","datagrid"].includes(s)&&(o=o.columns),null===i.parent){const t=this.setListIcon(o);e.push(a`<div key={mkey}>
          ${this.mapButton(r,r,t,s.toUpperCase(),l+1)}
        </div>`)}else if(i.item===o||i.parent===o){const t=this.setListIcon(o,l+1);if(e.push(a`<div key={mkey}>
            ${this.mapButton(r,r,t,s.toUpperCase(),t.badge)}
          </div>`),["row","datagrid"].includes(i.type)||["row","datagrid"].includes(i.parent_type))for(let t=0;t<o.length;t+=1){const r=g(o[t]),d=o[t][r],n=`tmp_${i.section}_${l.toString()}_${s}_${t.toString()}_${r}`,p=this.setListIcon(d);e.push(a`<div key={skey}>
                ${this.mapButton(n,n,p,r.toUpperCase(),p.badge,`primary ${u.badgeBlack}`)}
              </div>`)}}}}createMapList(){const{template:e}=this.data,t=[];return["report","header","details","footer"].forEach((i=>{const l=this.setListIcon(e[i]);l.selected&&"report"!==i&&t.push(a`<hr id="${`sep_${i}_0`}" class="separator" />`),t.push(a`<div key={mkey}>
        ${this.mapButton(`tmp_${i}`,`tmp_${i}`,l,i.toUpperCase(),l.badge)}
      </div>`),l.selected&&(t.push(a`<hr id="${`sep_${i}_1`}" class="separator" />`),"report"!==i&&(this.createSubList(t),t.push(a`<hr id="${`sep_${i}_2`}" class="separator" />`)))})),t}dataTitle(e,t,i){return a`<div class="panel-title">
      <div class="cell">
        <form-label class="title-cell"
          value="${t}"
        ></form-label>
      </div>
      <div class="cell align-right" >
        <span id="${e}" class="close-icon" 
          @click="${()=>this._onTemplateEvent(...i)}">
          <form-icon iconKey="Times" ></form-icon>
        </span>
      </div>
    </div>`}dataText(e,t,i,l){return a`<textarea id="${`${e}_value`}"
      rows=${i} .value="${t}"
      @input="${e=>this._onTemplateEvent(d.EDIT_DATA_ITEM,{...l,value:e.target.value})}"
    ></textarea>`}tableFields(){const{current_data:e}=this.data;return{...e.fields,edit:{columnDef:{id:"delete",Header:"",headerStyle:{},Cell:({row:e})=>a`<form-icon id=${`delete_${e._index}`}
            iconKey="Times" width=19 height=27.6
            .style=${{cursor:"pointer",fill:"rgb(var(--functional-red))"}}
            @click=${t=>{t.stopPropagation(),this._onTemplateEvent(d.DELETE_DATA_ITEM,{_index:e._index})}}
          ></form-icon>`,cellStyle:{width:40,padding:"4px 8px 3px 8px"}}}}}render(){const{title:e,tabView:t,template:l,current:o,current_data:r,dataset:p}=this.data,c=(e,t)=>({map_edit:{data:!1,report:!1,header:!1,footer:!1,details:!1},map_insert:{header:!0,footer:!0,details:!0,row:!0,datagrid:!0}}[t][e]);return a`<div class="panel" >
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
        ${"template"===t?a`<div class="section-container-small border" >
          <div class="cell padding-small third" >
            ${this.navButton("previous",[d.GO_PREVIOUS,[]],"label_previous","ArrowLeft")}
            <div class="mapBox" >
              <canvas ${b(this.canvasCallback)} class="reportMap" ></canvas>
            </div>
            ${this.navButton("next",[d.GO_NEXT,[]],"label_next","ArrowRight",!0,s.RIGHT)}
          </div>
          <div class="cell padding-small third" >
            ${this.createMapList()}
          </div>
          <div class="cell padding-small third" >
            ${!1!==c(o.type,"map_edit")?a`<div>
                ${this.navButton("move_up",[d.MOVE_UP,[]],"label_move_up","ArrowUp")}
                ${this.navButton("move_down",[d.MOVE_DOWN,[]],"label_move_down","ArrowDown")}
                ${this.navButton("delete_item",[d.DELETE_ITEM,[]],"label_delete","Times")}
                <hr class="separator" />
              </div>`:i}
            ${c(o.type,"map_insert")?a`<div>
                ${this.navButton("add_item",[d.ADD_ITEM,o.add_item||""],"label_add_item","Plus")}
                <form-select id="sel_add_item" label="" 
                  ?full=${!1} .isnull="${!0}" value="${o.add_item||""}"
                  .onChange=${e=>this._onTemplateEvent(d.CHANGE_CURRENT,{key:"add_item",value:e.value})}
                  .options=${{header:["row","vgap","hline"],details:["row","vgap","hline","html","datagrid"],footer:["row","vgap","hline"],row:["cell","image","barcode","separator"],datagrid:["column"]}[o.type].map((e=>({value:e,text:e.toUpperCase()})))}  
                ></form-select>
              </div>`:i}
          </div>
        </div>
        <div class="report-title padding-small" >
          <form-label class="report-title-label"
            value="${o.type.toUpperCase()}" leftIcon="Tag"
          ></form-label>
        </div>
        ${o.form.rows.map(((e,t)=>a`<div class="template-row" >
          <form-row id=${`row_${t}`}
            .row=${e} 
            .values=${["row","datagrid"].includes(o.type)?o.item_base:o.item}
            .options=${o.form.options}
            .data=${{audit:"all",current:o,dataset:l.data}}
            .onEdit=${e=>this._onTemplateEvent(d.EDIT_ITEM,e)}
            .msg=${this.msg}
          ></form-row>
        </div>`))}`:i}
        ${"data"===t?a`<div class="section-container border" >
          ${r&&"string"===r.type?a`<div class="row full section-small">
              ${this.dataTitle("data_string",r.name,[d.SET_CURRENT_DATA,null])}
              ${this.dataText(r.name,l.data[r.name],15,{})}
          </div>`:i}
          ${r&&"list"===r.type&&r.item?a`<div class="row full section-small">
              ${this.dataTitle("data_list_item",r.item,[d.SET_CURRENT_DATA_ITEM,null])}
              ${this.dataText(r.item,l.data[r.name][r.item],10,{})}
          </div>`:i}
          ${r&&"table"===r.type&&r.item?a`<div class="row full section-small">
              ${this.dataTitle("data_table_item",`${r.name} - ${String(r.item._index+1)}`,[d.SET_CURRENT_DATA_ITEM,null])}
              ${Object.keys(r.fields).map((e=>a`<div class="row full">
                <div class="padding-small">
                  <form-label value="${e}"></form-label>
                </div>
                ${this.dataText(e,r.item[e],2,{field:e,_index:r.item._index})}
              </div>`))}
          </div>`:i}
          ${r&&"list"===r.type&&!r.item?a`<div class="row full section-small">
              ${this.dataTitle("data_list",r.name,[d.SET_CURRENT_DATA,null])}
              <form-list id="data_list_items"
                .rows=${r.items} ?listFilter=${!0}
                filterPlaceholder=${this.msg("",{id:"placeholder_filter"})}
                .onAddItem=${()=>this._onTemplateEvent(d.SET_CURRENT_DATA_ITEM,void 0)}
                labelAdd=${this.msg("",{id:"label_new"})}
                pageSize=${this.paginationPage} pagination="${n.TOP}" 
                .onEdit=${e=>this._onTemplateEvent(d.SET_CURRENT_DATA_ITEM,e.lslabel)}
                .onDelete=${e=>this._onTemplateEvent(d.DELETE_DATA_ITEM,{key:e.lslabel})}
              ></form-list>
          </div>`:i}
          ${r&&"table"===r.type&&!r.item?a`<div class="row full section-small">
              ${this.dataTitle("data_table",r.name,[d.SET_CURRENT_DATA,null])}
              <form-table id="data_table_items"
                .fields=${this.tableFields()} .rows=${r.items} ?tableFilter=${!0}
                filterPlaceholder="${this.msg("",{id:"placeholder_filter"})}"
                .onAddItem=${()=>this._onTemplateEvent(d.SET_CURRENT_DATA_ITEM,void 0)}
                .onRowSelected=${e=>this._onTemplateEvent(d.SET_CURRENT_DATA_ITEM,e)} 
                labelAdd=${this.msg("",{id:"label_new"})}  
                pageSize=${this.paginationPage} pagination="${n.TOP}"
              ></form-table>
          </div>`:i}
          ${r?i:a`<div class="row full section-small">
              <form-list id="data_list_items"
                .rows=${p} ?listFilter=${!0}
                filterPlaceholder=${this.msg("",{id:"placeholder_filter"})}
                .onAddItem=${()=>this._onTemplateEvent(d.ADD_TEMPLATE_DATA,[])}
                labelAdd=${this.msg("",{id:"label_new"})}
                pageSize=${this.paginationPage} pagination="${n.TOP}" 
                .onEdit=${e=>this._onTemplateEvent(d.SET_CURRENT_DATA,{name:e.lslabel,type:e.lsvalue})}
                .onDelete=${e=>this._onTemplateEvent(d.DELETE_DATA,e.lslabel)}
              ></form-list>
          </div>`}
        </div>`:i}
        ${"meta"===t?a`<div class="section-container-small border" >
          <div class="cell padding-small" >
            <div class="meta-title-row" >
              ${Object.keys(l.meta).map((e=>a`<div class="meta-title-cell" >
                <div class="bold">${e}</div>
                <div>${l.meta[e]}</div>
              </div>`))}
            </div>
            <div class="meta-title-sources" >
              <form-label value="${this.msg("",{id:"template_data_sources"})}"></form-label>
            </div>
            ${Object.keys(l.sources).map((e=>a`<div class="meta-sources" >
              <div class="cell padding-small">
                <div class="meta-sources-name padding-small">
                  ${e}
                </div>
                ${Object.keys(l.sources[e]).map((t=>a`<div class="row" >
                  <div class="meta-sources-cell padding-small bold" >${t}:</div>
                  <div class="meta-sources-cell padding-small" >${l.sources[e][t]}</div>
                </div>`))}
              </div>
            </div>`))}
          </div>
        </div>`:i}
      </div>
    </div>`}});const f=e`
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
`;customElements.define("modal-template",class extends t{constructor(){super(),this.msg=e=>e,this.type=p.TEXT,this.name="",this.columns=""}static get properties(){return{type:{type:String},name:{type:String},columns:{type:String}}}static get styles(){return[f]}_onModalEvent(e,t){this.onEvent&&this.onEvent.onModalEvent&&this.onEvent.onModalEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("modal_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}_onValueChange(e,t){this[e]=t}_onTextInput(e){this._onValueChange("columns",e.target.value)}render(){const{type:e,name:t,columns:l}=this;return a`<div class="modal">
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
                @click="${()=>this._onModalEvent(c.CANCEL,{})}">
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
                  .options=${Object.keys(p).map((e=>({value:p[e],text:e})))}  
                ></form-select>
              </div>
              <div class="cell padding-small half" >
                <div>
                  <form-label
                    value="${this.msg("",{id:"template_data_name"})}" 
                  ></form-label>
                </div>
                <form-input id="name"
                  type="${m.TEXT}"
                  label="${this.msg("",{id:"template_data_name"})}" 
                  .onChange=${e=>this._onValueChange("name",e.value)}
                  value="${t}" ?full=${!0}
                ></form-input>
              </div>
            </div>
            ${e===p.TABLE?a`<div class="section-row" >
              <div class="cell padding-small" >
                <div>
                  <form-label
                    value="${this.msg("",{id:"template_data_columns"})}" 
                  ></form-label>
                </div>
                  <textarea id="columns"
                    rows=3 .value="${l}"
                    @input="${this._onTextInput}"
                  ></textarea>
              </div>
            </div>`:i}
          </div>
          <div class="section buttons" >
            <div class="section-row" >
              <div class="cell padding-small half" >
                <form-button id="btn_cancel" icon="Times"
                  @click=${()=>this._onModalEvent(c.CANCEL,{})} 
                  ?full="${!0}" label="${this.msg("",{id:"msg_cancel"})}"
                >${this.msg("",{id:"msg_cancel"})}</form-button>
              </div>
              <div class="cell padding-small half" >
                <form-button id="btn_ok" icon="Check"
                  @click=${()=>this._onModalEvent(c.OK,{value:{name:t,type:e,columns:l}})} 
                  type="${o.PRIMARY}" ?full="${!0}" 
                  label="${this.msg("",{id:"msg_ok"})}"
                >${this.msg("",{id:"msg_ok"})}</form-button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>`}});const _=e`
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
`;customElements.define("client-template",class extends t{constructor(){super(),this.msg=e=>e,this.id=Math.random().toString(36).slice(2),this.side=l.AUTO,this.data={},this.paginationPage=10,this.onEvent={},this.modalTemplate=this.modalTemplate.bind(this)}static get properties(){return{id:{type:String},side:{type:String},data:{type:Object},paginationPage:{type:Number},onEvent:{type:Object}}}static get styles(){return[_]}connectedCallback(){super.connectedCallback(),this.onEvent.setModule(this)}modalTemplate({type:e,name:t,columns:i,onEvent:l}){return a`<modal-template
      type="${e}"
      name="${t}"
      columns="${i}"
      .onEvent=${l} .msg=${this.msg}
    ></modal-template>`}render(){const{side:e,data:t,paginationPage:i}=this;return a`<sidebar-template
      id="${this.id}" side="${e}"
      templateKey="${t.key}" ?dirty="${t.dirty}"
      .onEvent=${this.onEvent} .msg=${this.msg}
    ></sidebar-template>
      <div class="page">
        <template-editor
          .data="${t}"
          .onEvent=${this.onEvent} .msg=${this.msg}
          paginationPage="${i}"
        ></template-editor>
      </div>`}});
