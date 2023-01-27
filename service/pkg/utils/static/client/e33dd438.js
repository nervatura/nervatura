import{i as e,s as t,y as i,b as l}from"./4e7ea0c6.js";import{d as a,T as o,B as s,S as r,c as n,f as d,D as c,I as p,P as b,M as h,g as v}from"./c1c25c6b.js";import"./2255c85d.js";import"./ddfaa96d.js";import"./81d721ef.js";import"./f6baf16b.js";import"./2b347bbf.js";import"./95ec07a4.js";const u=e`
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
.full {
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
`;customElements.define("sidebar-search",class extends t{constructor(){super(),this.msg=e=>e,this.side=a.AUTO,this.groupKey="",this.auditFilter={}}static get properties(){return{side:{type:String,reflect:!0},groupKey:{type:String},auditFilter:{type:Object}}}static get styles(){return[u]}_onSideEvent(e,t){this.onEvent&&this.onEvent.onSideEvent&&this.onEvent.onSideEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("side_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}getButtonStyle(e,t){return"group"===e?{"text-align":"left","border-radius":"0",color:t===this.groupKey?"rgb(var(--functional-yellow))":"",fill:t===this.groupKey?"rgb(var(--functional-yellow))":"","border-color":"rgba(var(--accent-1c), 0.2)"}:{"text-align":"left","border-radius":"0",color:"rgb(var(--functional-blue))",fill:"rgb(var(--functional-blue))","border-color":"rgba(var(--accent-1c), 0.2)"}}searchGroup(e){return i`<div class="row full">
        <form-button id="${`btn_group_${e}`}" 
          label="${this.msg("",{id:`search_${e}`})}"
          ?full="${!0}" align=${o.LEFT}
          .style="${this.getButtonStyle("group",e)}"
          icon="FileText" type="${s.PRIMARY}"
          @click=${()=>this._onSideEvent(r.CHANGE,{fieldname:"group_key",value:e})} 
        >${this.msg("",{id:`search_${e}`})}</form-button>
        ${this.groupKey===e?i`<div class="row full panel-group" >
          <form-button id="${`btn_view_${e}`}" 
            label="${this.msg("",{id:"quick_search"})}"
            ?full="${!0}" align=${o.LEFT}
            .style="${this.getButtonStyle("panel")}"
            icon="Bolt" type="${s.PRIMARY}"
            @click=${()=>this._onSideEvent(r.QUICK,{value:e})} 
          >${this.msg("Quick Search",{id:"quick_search"})}</form-button>
          <form-button id="${`btn_browser_${e}`}" 
            label="${this.msg("",{id:`browser_${e}`})}"
            ?full="${!0}" align=${o.LEFT}
            .style="${this.getButtonStyle("panel")}"
            icon="Search" type="${s.PRIMARY}"
            @click=${()=>this._onSideEvent(r.BROWSER,{value:e})} 
          >${this.msg("",{id:`browser_${e}`})}</form-button>
        </div>`:l}
      </div>`}render(){return i`<div class="sidebar ${"auto"!==this.side?this.side:""}" >
      ${this.searchGroup("transitem")}
      ${"disabled"!==this.auditFilter.trans.bank[0]||"disabled"!==this.auditFilter.trans.cash[0]?this.searchGroup("transpayment"):l}
      ${"disabled"!==this.auditFilter.trans.delivery[0]||"disabled"!==this.auditFilter.trans.inventory[0]||"disabled"!==this.auditFilter.trans.waybill[0]||"disabled"!==this.auditFilter.trans.production[0]||"disabled"!==this.auditFilter.trans.formula[0]?this.searchGroup("transmovement"):l}
      
      <hr class="separator" />
      ${["customer","product","employee","tool","project"].map((e=>"disabled"!==this.auditFilter[e][0]?this.searchGroup(e):l))}

      <hr class="separator" />
      <form-button id="btn_report" 
        label="${this.msg("",{id:"search_report"})}"
        ?full="${!0}" align=${o.LEFT}
        .style="${this.getButtonStyle("group","report")}"
        icon="ChartBar" type="${s.PRIMARY}"
        @click=${()=>{this._onSideEvent(r.CHANGE,{fieldname:"group_key",value:"report"}),this._onSideEvent(r.QUICK,{value:"report"})}} 
      >${this.msg("",{id:"search_report"})}</form-button>
      <form-button id="btn_office" 
        label="${this.msg("",{id:"search_office"})}"
        ?full="${!0}" align=${o.LEFT}
        .style="${this.getButtonStyle("group","office")}"
        icon="Inbox" type="${s.PRIMARY}"
        @click=${()=>this._onSideEvent(r.CHANGE,{fieldname:"group_key",value:"office"})} 
      >${this.msg("",{id:"search_office"})}</form-button>
      ${"office"===this.groupKey?i`<div class="row full panel-group" >
        <form-button id="btn_printqueue" 
          label="${this.msg("",{id:"title_printqueue"})}"
          ?full="${!0}" align=${o.LEFT}
          .style="${this.getButtonStyle("panel")}"
          icon="Print" type="${s.PRIMARY}"
          @click=${()=>this._onSideEvent(r.CHECK,{ntype:"printqueue",ttype:null,id:null})} 
        >${this.msg("",{id:"title_printqueue"})}</form-button>
        <form-button id="btn_rate" 
          label="${this.msg("",{id:"title_rate"})}"
          ?full="${!0}" align=${o.LEFT}
          .style="${this.getButtonStyle("panel")}"
          icon="Globe" type="${s.PRIMARY}"
          @click=${()=>this._onSideEvent(r.BROWSER,{value:"rate"})} 
        >${this.msg("",{id:"title_rate"})}</form-button>
        <form-button id="btn_servercmd" 
          label="${this.msg("",{id:"title_servercmd"})}"
          ?full="${!0}" align=${o.LEFT}
          .style="${this.getButtonStyle("panel")}"
          icon="Share" type="${s.PRIMARY}"
          @click=${()=>this._onSideEvent(r.QUICK,{value:"servercmd"})} 
        >${this.msg("",{id:"title_servercmd"})}</form-button>
      </div>`:l}
    </div>`}});const f=e`
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
.filter-panel {
  width: 100%;
  padding: 16px;
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  border-top: none;
  border-radius: 0px;
}
.panel-container { 
  padding: 16px; 
}
.align-right { 
  text-align: right; 
}
.dropdown-box {
  position: relative;
  display: inline-block;
}
.section-small-top { 
  padding-top: 8px; 
}
@keyframes opac{
  from{opacity:0} to{opacity:1}
}
.dropdown-content{
  background-color: rgba(var(--base-2), 1);
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  border-top: none;
  animation: opac 0.8s;
  box-shadow:0 2px 3px 0 rgba(0,0,0,0.16),0 2px 10px 0 rgba(0,0,0,0.12);
  position: absolute;
  min-width: 160px;
  margin: 0;
  padding: 4px 8px;
  z-index: 5;
}
.drop-label {
  font-size: 14px;
  padding: 2px 0px;
  font-weight: bold;
  vertical-align: middle;
  white-space: nowrap;
  width: 100%;
  cursor: pointer;
  -webkit-touch-callout:none;
  -webkit-user-select:none;
  -khtml-user-select:none;
  -moz-user-select:none;
  -ms-user-select:none;
  user-select:none;
}
.drop-label form-label:hover {
  color: rgb(var(--functional-yellow));
  fill: rgb(var(--functional-yellow));
}
.active {
  color: rgb(var(--functional-yellow))!important;
  fill: rgb(var(--functional-yellow))!important;
}
.col-box {
  display: table;
  width: 100%;
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  border-radius: 4px;
  white-space: nowrap;
  padding: 6px;
  margin-top: 2px;
}
.base-col form-label{
  font-size: 10px;
}
.select-col form-label{
  color: rgb(var(--functional-green));
  fill: rgb(var(--functional-green));
  font-weight: bold;
}
.edit-col form-label:hover {
  font-weight: normal;
}
.edit-col form-label:hover {
  color: rgb(var(--functional-green));
  fill: rgb(var(--functional-green));
}
.col-cell {
  padding: 2px 4px;
  float: left;
  cursor: pointer;
}
.border {
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
}
.result-title {
  background-color: rgba(var(--accent-1),0.1);
  font-weight: bold;
  padding: 8px 16px;
}
.result-title-plus {
  background-color: rgba(var(--accent-1),0.1);
  vertical-align: middle;
  padding: 0px 16px;
  width: 40px;
  text-align: right;
}
@media (max-width:600px){
  .filter-panel {
    padding: 8px;
  }
  .panel-container { 
    padding: 16px 8px; 
  }
  .col-cell { 
    padding: 1px 2px; 
  }
  .mobile{ 
    display: block; 
    width: 100%; 
  }
}
`;customElements.define("search-browser",class extends t{constructor(){super(),this.msg=e=>e,this.data={vkey:"",view:"",show_header:!0,show_dropdown:!1,show_columns:!1,result:[],columns:{},filters:{},deffield:[],page:1},this.keyMap={},this.viewDef={fields:{},label:"",readonly:!1},this.paginationPage=10,this.onEvent={}}static get properties(){return{data:{type:Object},keyMap:{type:Object},viewDef:{type:Object},paginationPage:{type:Number},onEvent:{type:Object}}}static get styles(){return[f]}connectedCallback(){super.connectedCallback();const{show_header:e,show_dropdown:t,show_columns:i}=this.data;this.dropdown=t,this.header=e,this.columns=i}_onBrowserEvent(e,t){this.onEvent&&this.onEvent.onBrowserEvent&&this.onEvent.onBrowserEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("browser_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}_onValueChange(e,t){this[e]=t,this.dispatchEvent(new CustomEvent("value_change",{bubbles:!0,composed:!0,detail:{key:e,value:t}})),this.requestUpdate()}exportFields(){const{view:e,columns:t}=this.data,i={};return Object.keys(this.viewDef.fields).filter((i=>!0===t[e][i])).forEach((e=>{i[e]=this.viewDef.fields[e]})),i}checkTotalFields(){const{fields:e}=this.viewDef,{deffield:t}=this.data;let i={totalFields:{},totalLabels:{},count:0};return t&&Object.keys(e).includes("deffield_value")?t.filter((e=>"integer"===e.fieldtype||"float"===e.fieldtype)).forEach((e=>{i={...i,totalFields:{...i.totalFields,[e.fieldname]:0},totalLabels:{...i.totalLabels,[e.fieldname]:e.description}}})):Object.keys(e).filter((t=>("integer"===e[t].fieldtype||"float"===e[t].fieldtype)&&"avg"!==e[t].calc)).forEach((t=>{i={...i,totalFields:{...i.totalFields,[t]:0},totalLabels:{...i.totalLabels,[t]:e[t].label}}})),i={...i,count:Object.keys(i.totalFields).length},i}fields(){const{view:e,columns:t}=this.data;let l={view:{columnDef:{id:"view",Header:"",headerStyle:{},Cell:({row:e})=>this.viewDef.readonly?i`<form-icon iconKey="CaretRight" width=9 height=24 ></form-icon>`:i`<form-icon id=${`edit_${e.id}`}
              iconKey="Edit" width=24 height=21.3 
              @click=${()=>this._onBrowserEvent(n.EDIT_CELL,{fieldname:"id",value:e.id,row:e})}
              .style=${{cursor:"pointer",fill:"rgb(var(--functional-green))"}} ></form-icon>`,cellStyle:{width:"25px",padding:"7px 3px 3px 8px"}}}};return Object.keys(this.viewDef.fields).forEach((i=>{if(t[e][i])switch(this.viewDef.fields[i].fieldtype){case"float":case"integer":l={...l,[i]:{fieldtype:"number",label:this.viewDef.fields[i].label}};break;case"bool":l={...l,[i]:{fieldtype:"bool",label:this.viewDef.fields[i].label}};break;default:l="deffield_value"===i?{...l,[i]:{fieldtype:"deffield",label:this.viewDef.fields[i].label}}:{...l,[i]:{fieldtype:"string",label:this.viewDef.fields[i].label}}}})),l}render(){const{vkey:e,view:t,result:a,columns:r,filters:h,deffield:v,page:u}=this.data,f=this.checkTotalFields();return i`<div @click="${()=>this.dropdown?this._onValueChange("dropdown",!1):null}">
      <div class="panel">
        <div class="panel-title">
          <div class="cell">
            <form-label 
              value="${this.msg(`browser_${e}`,{id:`browser_${e}`})}" 
              class="title-cell" >
            </form-label>
          </div>
        </div>
        <div class="panel-container" >
          <div class="row full" >
            <div class="cell" >
              <form-button id="btn_header" 
                icon="Filter" type="${s.PRIMARY}" ?full="${!0}"
                label="${this.viewDef.label}" align=${o.LEFT}
                @click=${e=>{e.stopPropagation(),this._onBrowserEvent(n.CHANGE,{fieldname:"show_header",value:!this.header}),this._onValueChange("header",!this.header)}}
              >${this.viewDef.label}
              </form-button>
            </div>
          </div>
          ${this.header?i`<div class="filter-panel" >
            <div class="row full" >
              <div class="cell" >
                <form-button id="btn_search" 
                  icon="Search" type="${s.BORDER}"
                  label="${this.msg("",{id:"browser_search"})}"
                  .style=${{padding:"8px 12px"}} ?hidelabel=${!0}
                  @click=${()=>this._onBrowserEvent(n.BROWSER_VIEW,{})}
                >${this.msg("",{id:"browser_search"})}</form-button>
              </div>
              <div class="cell align-right" >
                <form-button id="btn_bookmark" 
                  icon="Star" type="${s.BORDER}"
                  label="${this.msg("",{id:"browser_bookmark"})}"
                  .style=${{padding:"8px 12px"}} ?hidelabel=${!0}
                  @click=${()=>this._onBrowserEvent(n.BOOKMARK_SAVE,[])}
                >${this.msg("",{id:"browser_bookmark"})}</form-button>
                <form-button id="btn_export" 
                  icon="Download" type="${s.BORDER}"
                  label="${this.msg("",{id:"browser_export"})}"
                  .style=${{padding:"8px 12px"}} ?hidelabel=${!0}
                  @click=${()=>this._onBrowserEvent(n.EXPORT_RESULT,{value:this.exportFields()})}
                >${this.msg("",{id:"browser_export"})}</form-button>
                <form-button id="btn_help" 
                  icon="QuestionCircle" type="${s.BORDER}"
                  label="${this.msg("",{id:"browser_help"})}"
                  .style=${{padding:"8px 12px"}} ?hidelabel=${!0}
                  @click=${()=>this._onBrowserEvent(n.SHOW_HELP,{value:"program/browser"})}
                >${this.msg("",{id:"browser_help"})}</form-button>
              </div>
            </div>
            <div class="row full section-small-top" >
              <div class="cell" >
                <div class="dropdown-box" >
                  <form-button id="btn_views"
                    type="${s.BORDER}"
                    .style=${{padding:"8px 12px"}} icon="Eye"
                    label="${this.msg("",{id:"browser_views"})}"
                    ?selected="${this.dropdown}" ?hidelabel=${!0}
                    @click=${e=>{e.stopPropagation(),this._onValueChange("dropdown",!this.dropdown)}}
                  >${this.msg("",{id:"browser_views"})}</form-button>
                  ${this.dropdown?i`<div class="dropdown-content" >
                    ${Object.keys(this.keyMap).map((a=>"options"!==a?i`<div id=${`view_${a}`}
                        @click=${()=>this._onBrowserEvent(n.SHOW_BROWSER,{value:e,vname:a})}
                        class="drop-label" >
                        <form-label class="${a===t?"active":""}"
                          value="${this.keyMap[a].label}"
                          leftIcon="${a===t?"Check":"Eye"}" ></form-label>
                      </div>`:l))}
                  </div>`:l}
                </div>
                <form-button id="btn_columns" 
                  type="${s.BORDER}"
                  .style=${{padding:"8px 12px"}} icon="Columns" ?hidelabel=${!0}
                  label="${this.msg("",{id:"browser_columns"})}"
                  @click=${e=>{e.stopPropagation(),this._onBrowserEvent(n.CHANGE,{fieldname:"show_columns",value:!this.columns}),this._onValueChange("columns",!this.columns)}}
                >${this.msg("",{id:"browser_columns"})}</form-button>
                <form-button id="btn_filter" 
                  icon="Plus" type="${s.BORDER}" ?hidelabel=${!0}
                  label="${this.msg("",{id:"browser_filter"})}"
                  .style=${{padding:"8px 12px"}}
                  @click=${()=>this._onBrowserEvent(n.ADD_FILTER,{})}
                >${this.msg("",{id:"browser_filter"})}</form-button>
                <form-button id="btn_total" 
                  icon="InfoCircle" type="${s.BORDER}"
                  label="${this.msg("",{id:"browser_total"})}"
                  .style=${{padding:"8px 12px"}} ?hidelabel=${!0}
                  ?disabled=${!(0!==f.count&&0!==a.length)}
                  @click=${()=>this._onBrowserEvent(n.SHOW_TOTAL,{fields:this.viewDef.fields,totalFields:f})}
                >${this.msg("",{id:"browser_total"})}</form-button>
              </div>
            </div>
            ${this.columns?i`<div class="col-box">
              ${Object.keys(this.viewDef.fields).map((e=>i`<div id=${`col_${e}`} 
                  class="cell col-cell base-col ${!0===r[t][e]?"select-col":"edit-col"}"
                  @click=${()=>this._onBrowserEvent(n.SET_COLUMNS,{fieldname:e,value:!(!0===r[t][e])})}
                  >
                  <form-label
                    value="${this.viewDef.fields[e].label}"
                    leftIcon="${!0===r[t][e]?"CheckSquare":"SquareEmpty"}" ></form-label>
                </div>`))}              
            </div>`:l}
            ${h[t].map(((e,t)=>i`<div class="section-small-top" >
              <div class="cell" >
                <form-select id=${`filter_name_${t}`}
                  label="${this.msg("",{id:"browser_filter"})}"
                  .onChange=${e=>this._onBrowserEvent(n.EDIT_FILTER,{index:t,fieldname:"fieldname",value:e.value})}
                  .options=${Object.keys(this.viewDef.fields).filter((e=>"id"!==e&&"_id"!==e)).flatMap((e=>"deffield_value"===e?v.map((e=>({value:e.fieldname,text:this.msg(e.description,{id:e.fieldname})}))):{value:e,text:this.viewDef.fields[e].label}))} 
                  .isnull="${!1}" value="${e.fieldname}" >
                </form-select>
              </div>
              <div class="cell" >
                <form-select id=${`filter_type_${t}`}
                  label="${this.msg("",{id:"browser_filter"})}"
                  .onChange=${e=>this._onBrowserEvent(n.EDIT_FILTER,{index:t,fieldname:"filtertype",value:e.value})}
                  .options=${(["date","float","integer"].includes(e.fieldtype)?d:d.slice(0,3)).map((e=>({value:e[0],text:e[1]})))} 
                  .isnull="${!1}" value="${e.filtertype}" >
                </form-select>
              </div>
              <div class="cell mobile" >
                ${"==N"!==e.filtertype?i`<div class="cell" >
                  ${"bool"===e.fieldtype?i`<form-select 
                    id=${`filter_value_${t}`}
                    label="${this.msg("",{id:"browser_filter"})}"
                    .onChange=${e=>this._onBrowserEvent(n.EDIT_FILTER,{index:t,fieldname:"value",value:e.value})}
                    .options=${[{value:"0",text:this.msg("NO",{id:"label_no"})},{value:"1",text:this.msg("YES",{id:"label_yes"})}]} 
                    .isnull="${!1}" value="${e.value}" >
                  </form-select>`:l}
                  ${"integer"===e.fieldtype||"float"===e.fieldtype?i`<form-number 
                    id=${`filter_value_${t}`} 
                    label="${this.msg("",{id:"browser_filter"})}"
                    ?integer="${!("float"===e.fieldtype)}"
                    value="${e.value}"
                    .onChange=${e=>this._onBrowserEvent(n.EDIT_FILTER,{index:t,fieldname:"value",value:e.value})}
                  ></form-number>`:l}
                  ${"date"===e.fieldtype?i`<form-datetime 
                    id=${`filter_value_${t}`}
                    label="${this.msg("",{id:"browser_filter"})}"
                    .isnull="${!1}"
                    type="${c.DATE}"
                    .value="${e.value}"
                    .onChange="${e=>this._onBrowserEvent(n.EDIT_FILTER,{index:t,fieldname:"value",value:e.value})}"
                  ></form-datetime>`:l}
                  ${"string"===e.fieldtype?i`<form-input 
                    id=${`filter_value_${t}`} 
                    label="${this.msg("",{id:"browser_filter"})}"
                    type="${p.TEXT}" 
                    value="${e.value}"
                    .onChange=${e=>this._onBrowserEvent(n.EDIT_FILTER,{index:t,fieldname:"value",value:e.value})}
                  ></form-input>`:l}
                </div>`:l}
                <div class="cell" >
                  <form-button id="${`btn_delete_filter_${t}`}" 
                    icon="Times" type="${s.BORDER}"
                    label="${this.msg("",{id:"label_delete"})}"
                    .style=${{padding:"8px","border-radius":"3px"}}
                    @click=${()=>this._onBrowserEvent(n.DELETE_FILTER,{value:t})}
                  ></form-button>
                </div>
              </div>
            </div>`))}
          </div>`:l}
          <div class="row full section-small-top" >
            <div class="row full border" >
              <div class="cell result-title" >
                ${a.length} <form-label 
                  value="${this.msg("record(s) found",{id:"browser_result"})}" 
                ></form-label>
              </div>
              ${this.viewDef.actions_new?i`<div class="cell result-title-plus" >
                <form-button 
                id="btn_actions_new" 
                icon="Plus" ?small=${!0}
                label="${this.msg("",{id:"label_new"})}"
                @click=${()=>this._onBrowserEvent(n.SET_FORM_ACTION,{params:this.viewDef.actions_new,row:void 0})}
                ></form-button>
              </div>`:l}
            </div>
          </div>
          <div class="row full" >
          <form-table 
            id="browser_result" rowKey="row_id"
            .rows="${a}"
            .fields="${this.fields()}"
            filterPlaceholder="${this.msg("Filter",{id:"placeholder_filter"})}"
            labelYes="${this.msg("YES",{id:"label_yes"})}"
            labelNo="${this.msg("NO",{id:"label_no"})}"
            pagination="${b.TOP}"
            currentPage="${u}"
            pageSize="${this.paginationPage}"
            ?tableFilter="${!0}"
            ?hidePaginatonSize="${!1}"
            .onEditCell=${(e,t,i)=>this._onBrowserEvent(n.EDIT_CELL,{fieldname:e,value:t,row:i})}
            .onCurrentPage=${e=>this._onBrowserEvent(n.CURRENT_PAGE,{value:e})}
          ></form-table>
          </div>
        </div>
      </div>
    </div>`}});const m=e`
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
  padding: 0 16px;
}
.padding-small { 
  padding: 4px 8px; 
}
.padding-tiny { 
  padding: 2px 4px; 
}
.buttons {
  background-color: rgb(var(--base-1));
  border-top: 0.5px solid rgba(var(--neutral-1), 0.2);
}
.table-row {
  display: table-row;
}
@media (max-width:600px){
  .section-row { 
    padding: 0 8px 8px; 
  }
  .padding-small { 
    padding: 2px 4px; 
  }
  .padding-tiny { 
    padding: 1px 2px; 
  }
  .mobile{ 
    display: block; 
    width: 100%; 
  }
}
@media only screen and (min-width: 601px){
  .dialog {
    min-width: 400px;
  }
}
`;customElements.define("modal-total",class extends t{constructor(){super(),this.msg=e=>e,this.total={totalFields:{},totalLabels:{},count:0}}static get properties(){return{total:{type:Object}}}static get styles(){return[m]}_onModalEvent(e,t){this.onEvent&&this.onEvent.onModalEvent&&this.onEvent.onModalEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("modal_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}render(){return i`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label leftIcon="InfoCircle"
                value="${this.msg("Total",{id:"browser_total"})}" 
                class="title-cell" ></form-label>
            </div>
            <div class="cell align-right" >
              <span id=${"closeIcon"} class="close-icon" 
                @click="${()=>this._onModalEvent(h.CANCEL,{})}">
                <form-icon iconKey="Times" ></form-icon>
              </span>
            </div>
          </div>
          <div class="section" >
            <div class="section-row" >
              ${Object.keys(this.total.totalFields).map((e=>i`<div class="table-row full">
                  <div class="cell padding-tiny mobile">
                    <form-label
                      value="${this.total.totalLabels[e]}" 
                    ></form-label>
                  </div>
                  <div class="cell padding-tiny mobile">
                    <form-input type="${p.TEXT}"
                      label="${this.total.totalLabels[e]}"
                      .style=${{"font-weight":"bold","text-align":"right"}}
                      value="${new Intl.NumberFormat("default").format(this.total.totalFields[e])}"
                      ?disabled=${!0} ?full=${!0}
                    ></form-input>
                  </div>
                </div>`))}
            </div>
          </div>
          <div class="section buttons" >
            <div class="section-row" >
              <div class="cell padding-small" >
                <form-button id="btn_ok" icon="Check"
                  label="${this.msg("",{id:"msg_ok"})}"
                  @click=${()=>this._onModalEvent(h.CANCEL,{})} 
                  ?autofocus="${!0}"
                  type="${s.PRIMARY}" ?full="${!0}" 
                >${this.msg("",{id:"msg_ok"})}</form-button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>`}});const g=e`
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
`;customElements.define("modal-server",class extends t{constructor(){super(),this.msg=e=>e,this.cmd={},this.fields=[],this.values={}}static get properties(){return{cmd:{type:Object},fields:{type:Array},values:{type:Object}}}static get styles(){return[g]}_onModalEvent(e,t){this.onEvent&&this.onEvent.onModalEvent&&this.onEvent.onModalEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("modal_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}_onEdit(e){this.values[e.name]=e.value,this.requestUpdate()}render(){return i`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label leftIcon="Share"
                value="${this.cmd.description}" 
                class="title-cell" ></form-label>
            </div>
            <div class="cell align-right" >
              <span id=${"closeIcon"} class="close-icon" 
                @click="${()=>this._onModalEvent(h.CANCEL,{})}">
                <form-icon iconKey="Times" ></form-icon>
              </span>
            </div>
          </div>
          ${this.fields.map(((e,t)=>i`<form-row id=${`row_${t}`}
              .row=${{rowtype:"field",name:e.fieldname,datatype:e.fieldtypeName,label:e.description}} 
              .values=${{...this.values}}
              .options=${{}}
              .data=${{audit:"all",current:{},dataset:{}}}
              .onEdit=${e=>{this._onEdit(e)}}
              .msg=${this.msg}
            ></form-row>`))}
          <div class="section buttons" >
            <div class="section-row" >
              <div class="cell padding-small half" >
                <form-button id="btn_cancel" icon="Times"
                  label="${this.msg("",{id:"msg_cancel"})}"
                  @click=${()=>this._onModalEvent(h.CANCEL,{})} 
                  ?full="${!0}" 
                >${this.msg("",{id:"msg_cancel"})}</form-button>
              </div>
              <div class="cell padding-small half" >
                <form-button id="btn_ok" icon="Check"
                  label="${this.msg("",{id:"msg_ok"})}"
                  @click=${()=>this._onModalEvent(h.OK,{cmd:this.cmd,fields:this.fields,values:{...this.values}})} 
                  type="${s.PRIMARY}" ?full="${!0}" 
                >${this.msg("",{id:"msg_ok"})}</form-button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>`}});const $=e`
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
`;customElements.define("client-search",class extends t{constructor(){super(),this.data={},this.side=a.AUTO,this.auditFilter={},this.queries={},this.quick={},this.paginationPage=10,this.onEvent={},this.modalServer=this.modalServer.bind(this),this.modalTotal=this.modalTotal.bind(this)}static get properties(){return{data:{type:Object},side:{type:String},auditFilter:{type:Object},paginationPage:{type:Number},queries:{type:Object},quick:{type:Object},onEvent:{type:Object}}}static get styles(){return[$]}connectedCallback(){super.connectedCallback(),this.onEvent.setModule(this)}modalServer({cmd:e,fields:t,values:l,onEvent:a}){return i`<modal-server 
      .cmd=${e} .fields=${t} .values=${l} 
      .onEvent=${a} .msg=${this.msg}
    ></modal-server>`}modalTotal(e){return i`<modal-total 
      .total=${e} .onEvent=${this.onEvent} .msg=${this.msg}
    ></modal-total>`}render(){return i`<sidebar-search
      id="sidebar"
      side="${this.side}"
      groupKey="${this.data.group_key}"
      .auditFilter="${this.auditFilter}"
      .onEvent=${this.onEvent}
      .msg=${this.msg}
    ></sidebar-search>
      <div class="page">
        ${v("browser"===this.data.seltype?i`<search-browser
            id="${`browser_${this.data.vkey}`}"
            .data=${this.data}
            .keyMap=${this.queries[this.data.vkey]()}
            .viewDef="${this.queries[this.data.vkey]()[this.data.view]}"
            paginationPage=${this.paginationPage}
            .onEvent=${this.onEvent}
            .msg=${this.msg}
          ></search-browser>`:i`<modal-selector
            id="${`selector_${this.data.qview}`}"
            ?isModal="${!1}"
            view="${this.data.qview}"
            .columns=${this.quick[this.data.qview]().columns}
            .result=${this.data.result}
            filter="${this.data.qfilter}"
            paginationPage=${this.paginationPage}
            currentPage=${this.data.page}
            .onEvent=${this.onEvent}
            .msg=${this.msg}
          ></modal-selector>`)}
      </div>`}});
