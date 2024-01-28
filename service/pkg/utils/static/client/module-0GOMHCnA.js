import{i as e,s as t,x as i,T as l}from"./module-3zU3FH2L.js";import"./module-7N9Fn5lU.js";import{d as a,e as n,B as o,S as s,l as r,P as d,M as p,I as c}from"./main-HhCNuxlS.js";import{o as u}from"./module-JEOuC3n3.js";import"./module-67w4bJ4f.js";import"./module-AQ0TXxQm.js";import"./module-3cipbsnY.js";import"./module-r-Be6lM0.js";const m=e`
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
`;customElements.define("sidebar-setting",class extends t{constructor(){super(),this.msg=e=>e,this.side=a.AUTO,this.username=void 0,this.module={current:{},dirty:!1,panel:{},group_key:""},this.auditFilter={}}static get properties(){return{side:{type:String,reflect:!0},username:{type:String},module:{type:Object},auditFilter:{type:Object}}}static get styles(){return[m]}_onSideEvent(e,t){this.onEvent&&this.onEvent.onSideEvent&&this.onEvent.onSideEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("side_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}itemMenu({id:e,selected:t,eventValue:l,label:a,iconKey:s,full:r,color:d}){const p=void 0!==t&&t,c=void 0===r||r,u={"border-radius":"0","border-color":"rgba(var(--accent-1c), 0.2)"};return d&&(u.color=d,u.fill=d),i`<form-button 
        id="${e}" label="${a}"
        ?full="${c}" ?selected="${p}"
        align=${n.LEFT}
        .style="${u}"
        icon="${s}" type="${o.PRIMARY}"
        @click=${()=>this._onSideEvent(...l)} 
      >${a}</form-button>`}formItems(e){const{dirty:t,current:l,type:a}=this.module,n=e,o=[];return o.push(this.itemMenu({id:"cmd_back",selected:!0,eventValue:[s.BACK,{}],label:this.msg("",{id:"label_back"}),iconKey:"Reply",full:!1})),o.push(i`<hr id="back_sep" class="separator" />`),!1!==n.save&&o.push(this.itemMenu({id:"cmd_save",selected:t,eventValue:[s.SAVE,{}],label:this.msg("",{id:"label_save"}),iconKey:"Check"})),!1!==n.delete&&l.form&&null!==l.form.id&&o.push(this.itemMenu({id:"cmd_delete",eventValue:[s.DELETE,{value:l.form}],label:this.msg("",{id:"label_delete"}),iconKey:"Times"})),!1!==n.new&&l.form&&null!==l.form.id&&o.push(this.itemMenu({id:"cmd_new",eventValue:[s.CHECK,[{type:a,id:null},s.LOAD_SETTING]],label:this.msg("",{id:"label_new"}),iconKey:"Plus"})),void 0!==n.help&&(o.push(i`<hr id="help_sep" class="separator" />`),o.push(this.itemMenu({id:"cmd_help",eventValue:[s.HELP,{value:n.help}],label:this.msg("",{id:"label_help"}),iconKey:"QuestionCircle"}))),o}menuItems(e){const t=[];return"disabled"!==this.auditFilter.setting[0]&&(t.push(i`<div class="row full">
        ${this.itemMenu({id:"group_admin_group",selected:"group_admin"===e,eventValue:[s.CHANGE,{fieldname:"group_key",value:"group_admin"}],label:this.msg("",{id:"title_admin"}),iconKey:"ExclamationTriangle"})}
        ${"group_admin"===e?i`<div class="row full panel-group" >
          ${this.itemMenu({id:"cmd_dbsettings",eventValue:[s.LOAD_SETTING,{type:"setting"}],label:this.msg("",{id:"title_dbsettings"}),iconKey:"Cog",color:"rgb(var(--functional-blue))"})}
          ${this.itemMenu({id:"cmd_numberdef",eventValue:[s.LOAD_SETTING,{type:"numberdef"}],label:this.msg("",{id:"title_numberdef"}),iconKey:"ListOl",color:"rgb(var(--functional-blue))"})}
          ${"disabled"!==this.auditFilter.audit[0]?this.itemMenu({id:"cmd_usergroup",eventValue:[s.LOAD_SETTING,{type:"usergroup"}],label:this.msg("",{id:"title_usergroup"}),iconKey:"Key",color:"rgb(var(--functional-blue))"}):l}
          ${this.itemMenu({id:"cmd_ui_menu",eventValue:[s.LOAD_SETTING,{type:"ui_menu"}],label:this.msg("",{id:"title_menucmd"}),iconKey:"Share",color:"rgb(var(--functional-blue))"})}
          ${this.itemMenu({id:"cmd_log",eventValue:[s.LOAD_SETTING,{type:"log"}],label:this.msg("",{id:"title_log"}),iconKey:"InfoCircle",color:"rgb(var(--functional-blue))"})}
        </div>`:l}
      </div>`),t.push(i`<div class="row full">
        ${this.itemMenu({id:"group_database_group",selected:"group_database"===e,eventValue:[s.CHANGE,{fieldname:"group_key",value:"group_database"}],label:this.msg("",{id:"title_database"}),iconKey:"Database"})}
        ${"group_database"===e?i`<div class="row full panel-group" >
          ${this.itemMenu({id:"cmd_deffield",eventValue:[s.LOAD_SETTING,{type:"deffield"}],label:this.msg("",{id:"title_deffield"}),iconKey:"Tag",color:"rgb(var(--functional-blue))"})}
          ${this.itemMenu({id:"cmd_groups",eventValue:[s.LOAD_SETTING,{type:"groups"}],label:this.msg("",{id:"title_groups"}),iconKey:"Th",color:"rgb(var(--functional-blue))"})}
          ${this.itemMenu({id:"cmd_place",eventValue:[s.LOAD_SETTING,{type:"place"}],label:this.msg("",{id:"title_place"}),iconKey:"Map",color:"rgb(var(--functional-blue))"})}
          ${this.itemMenu({id:"cmd_currency",eventValue:[s.LOAD_SETTING,{type:"currency"}],label:this.msg("",{id:"title_currency"}),iconKey:"Dollar",color:"rgb(var(--functional-blue))"})}
          ${this.itemMenu({id:"cmd_tax",eventValue:[s.LOAD_SETTING,{type:"tax"}],label:this.msg("",{id:"title_tax"}),iconKey:"Ticket",color:"rgb(var(--functional-blue))"})}
          ${this.itemMenu({id:"cmd_company",eventValue:[s.CHECK,[{ntype:"customer",ttype:null,id:1}]],label:this.msg("",{id:"title_company"}),iconKey:"Home",color:"rgb(var(--functional-blue))"})}
          ${this.itemMenu({id:"cmd_template",eventValue:[s.LOAD_SETTING,{type:"template"}],label:this.msg("",{id:"title_report_editor"}),iconKey:"TextHeight",color:"rgb(var(--functional-blue))"})}
      </div>`:l}
    </div>`)),t.push(i`<div class="row full">
        ${this.itemMenu({id:"group_user_group",selected:"group_user"===e,eventValue:[s.CHANGE,{fieldname:"group_key",value:"group_user"}],label:this.msg("",{id:"title_user"}),iconKey:"Desktop"})}
        ${"group_user"===e?i`<div class="row full panel-group" >
          ${this.itemMenu({id:"cmd_program",eventValue:[s.PROGRAM_SETTING,{}],label:this.msg("",{id:"title_program"}),iconKey:"Keyboard",color:"rgb(var(--functional-blue))"})}
          ${this.itemMenu({id:"cmd_password",eventValue:[s.PASSWORD,{username:this.username}],label:this.msg("",{id:"title_password"}),iconKey:"Lock",color:"rgb(var(--functional-blue))"})}
      </div>`:l}
    </div>`),t}render(){const{group_key:e,current:t,panel:l}=this.module;return i`<div class="sidebar ${"auto"!==this.side?this.side:""}" >
    ${t&&l?this.formItems(l):this.menuItems(e)}
    </div>`}});const g=e`
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
.section {
  padding: 16px 0; 
}
.section-row {
  display: table;
  width: 100%;
  padding: 0 16px 8px;
}
.form-panel {
  width: 100%;
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  border-bottom: none;
}
@media (max-width:600px){
  .section-row { 
    padding: 0 8px 8px; 
  }
}
`;customElements.define("setting-form",class extends t{constructor(){super(),this.msg=e=>e,this.id=Math.random().toString(36).slice(2),this.data={caption:"",icon:"",current:{},audit:"",dataset:{},type:"",view:{}},this.paginationPage=10,this.onEvent={}}static get properties(){return{id:{type:String},data:{type:Object},paginationPage:{type:Number},onEvent:{type:Object}}}static get styles(){return[g]}_onSettingEvent(e,t){this.onEvent.onSettingEvent&&this.onEvent.onSettingEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("setting_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}render(){const{caption:e,icon:t,current:a,audit:n,dataset:o,type:s,view:p}=this.data;let c={};return void 0!==a.template.view.items&&null!==a.form.id&&(a.template.view.items.actions.edit&&(c={...c,edit:{columnDef:{id:"edit",Header:"",headerStyle:{},Cell:({row:e})=>{const t=i`<form-icon id=${`edit_${e.id}`}
                iconKey="Edit" width=24 height=21.3
                .style=${{cursor:"pointer",fill:"rgb(var(--functional-green))"}}
                @click=${t=>{t.stopPropagation(),this._onSettingEvent(r.FORM_ACTION,{params:a.template.view.items.actions.edit,row:e,ref:this})}}
              ></form-icon>`,l=a.template.view.items.actions.delete?i`<form-icon id=${`delete_${e.id}`}
                  iconKey="Times" width=19 height=27.6
                  .style=${{cursor:"pointer",fill:"rgb(var(--functional-red))","margin-left":"8px"}}
                  @click=${t=>{t.stopPropagation(),this._onSettingEvent(r.FORM_ACTION,{params:a.template.view.items.actions.delete,row:e})}}
                ></form-icon>`:void 0;return i`${t}${u(l)}`},cellStyle:{width:30,padding:"4px 3px 3px 8px"}}}}),c={...c,...a.template.view.items.fields}),i`<div class="panel" >
      <div class="panel-title">
        <div class="cell">
          <form-label class="title-cell"
            value="${e}" leftIcon="${t}"
          ></form-label>
        </div>
      </div>
      <div class="section full">
        <div class="section-row">
          <div id="${this.id}" class="form-panel" >
            ${a.template.rows.map(((e,t)=>i`<form-row
              id=${`row_${t}`}
              .row="${e}"
              .values="${a.fieldvalue||a.form}"
              .options="${a.template.options}"
              .data="${{audit:n,current:a,dataset:o}}"
              .onEdit=${e=>this._onSettingEvent(r.EDIT_ITEM,e)}
              .msg=${this.msg}
            ></form-row>`))}
          </div>
          ${void 0!==a.template.view.items&&null!==a.form.id||"log"===s?i`<form-table 
            id="form_view" rowKey="id"
            .rows="${"log"===s?p.result:o[a.template.view.items.data]}"
            .fields="${"log"===s?p.fields:c}"
            filterPlaceholder="${this.msg("Filter",{id:"placeholder_filter"})}"
            labelYes="${this.msg("YES",{id:"label_yes"})}"
            labelNo="${this.msg("NO",{id:"label_no"})}"
            pagination="${d.TOP}"
            pageSize="${this.paginationPage}"
            ?tableFilter="${!0}"
            ?hidePaginatonSize="${!1}"
            .onAddItem=${a.template.view.items&&a.template.view.items.actions.new?()=>this._onSettingEvent(r.FORM_ACTION,{params:a.template.view.items.actions.new}):void 0}
            labelAdd=${this.msg("",{id:"label_new"})}
          ></form-table>`:l}
        </div>
      </div>
    </div>`}});const v=e`
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
.section {
  padding: 16px 0; 
}
.section-row {
  display: table;
  width: 100%;
  padding: 0 16px 8px;
}
.form-panel {
  width: 100%;
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  border-bottom: none;
}
@media (max-width:600px){
  .section-row { 
    padding: 0 8px 8px; 
  }
}
`;customElements.define("setting-view",class extends t{constructor(){super(),this.msg=e=>e,this.id=Math.random().toString(36).slice(2),this.data={caption:"",icon:"",view:{},actions:{}},this.paginationPage=10,this.onEvent={}}static get properties(){return{id:{type:String},data:{type:Object},paginationPage:{type:Number},onEvent:{type:Object}}}static get styles(){return[v]}_onSettingEvent(e,t){this.onEvent.onSettingEvent&&this.onEvent.onSettingEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("setting_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}render(){const{caption:e,icon:t,view:l,actions:a,page:n}=this.data;let o={};return"table"===l.type&&(a.edit&&(o={...o,edit:{columnDef:{id:"edit",Header:"",headerStyle:{},Cell:({row:e})=>i`<form-icon id=${`edit_${e.id}`}
                iconKey="Edit" width=24 height=21.3
                .style=${{cursor:"pointer",fill:"rgb(var(--functional-green))"}}
                @click=${t=>{t.stopPropagation(),this._onSettingEvent(r.FORM_ACTION,{params:a.edit,row:e,ref:this})}}
              ></form-icon>`,cellStyle:{width:30,padding:"4px 3px 3px 8px"}}}}),o={...o,...l.fields}),i`<div class="panel" >
      <div class="panel-title">
        <div class="cell">
          <form-label class="title-cell"
            value="${e}" leftIcon="${t}"
          ></form-label>
        </div>
      </div>
      <div class="section full">
        <div class="section-row">
          ${"table"===l.type?i`<form-table id="view_table"
            .fields=${o} .rows=${l.result} ?tableFilter=${!0}
            filterPlaceholder="${this.msg("",{id:"placeholder_filter"})}"
            .onAddItem=${a.new?()=>this._onSettingEvent(r.FORM_ACTION,{params:a.new}):void 0}
            .onRowSelected=${a.edit?e=>this._onSettingEvent(r.FORM_ACTION,{params:a.edit,row:e}):void 0}
            labelYes=${this.msg("",{id:"label_yes"})} 
            labelNo=${this.msg("",{id:"label_no"})} 
            labelAdd=${this.msg("",{id:"label_new"})}  
            pageSize=${this.paginationPage} pagination="${d.TOP}"
            currentPage="${n}"
            .onCurrentPage=${e=>this._onSettingEvent(r.CURRENT_PAGE,{value:e})}
          ></form-table>`:i`<form-list id="view_list"
            .rows=${l.result} ?listFilter=${!0}
            filterPlaceholder=${this.msg("",{id:"placeholder_filter"})}
            .onAddItem=${a.new?()=>this._onSettingEvent(r.FORM_ACTION,{params:a.new}):void 0}
            labelAdd=${this.msg("",{id:"label_new"})}
            pageSize=${this.paginationPage} pagination="${d.TOP}" 
            .onEdit=${a.edit?e=>this._onSettingEvent(r.FORM_ACTION,{params:a.edit,row:e}):void 0}
            .onDelete=${a.delete?e=>this._onSettingEvent(r.FORM_ACTION,{params:a.delete,row:e}):void 0}
            currentPage="${n}"
            .onCurrentPage=${e=>this._onSettingEvent(r.CURRENT_PAGE,{value:e})}
          ></form-list>`}
        </div>
      </div>
    </div>`}});const b=e`
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
`;customElements.define("modal-audit",class extends t{constructor(){super(),this.msg=e=>e,this.idKey=null,this.usergroup=0,this.nervatype=void 0,this.subtype=null,this.inputfilter=void 0,this.supervisor=0,this.typeOptions=[],this.subtypeOptions=[],this.inputfilterOptions=[]}static get properties(){return{idKey:{type:Object},usergroup:{type:Number},nervatype:{type:Number},subtype:{type:Object},inputfilter:{type:Number},supervisor:{type:Number},typeOptions:{type:Array},subtypeOptions:{type:Array},inputfilterOptions:{type:Array}}}static get styles(){return[b]}_onModalEvent(e,t){this.onEvent&&this.onEvent.onModalEvent&&this.onEvent.onModalEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("modal_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}_onValueChange(e,t){this[e]=t}render(){const{idKey:e,usergroup:t,nervatype:a,subtype:n,inputfilter:s,supervisor:r,typeOptions:d,subtypeOptions:c,inputfilterOptions:u}=this,m=d.filter((e=>e.value===String(a)))[0].text,g=["trans","report","menu"].includes(d.filter((e=>e.value===String(a)))[0].text);return i`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label leftIcon="Key"
                value="${this.msg("",{id:"title_usergroup"})}" 
                class="title-cell" ></form-label>
            </div>
            <div class="cell align-right" >
              <span id=${"closeIcon"} class="close-icon" 
                @click="${()=>this._onModalEvent(p.CANCEL,{})}">
                <form-icon iconKey="Times" ></form-icon>
              </span>
            </div>
          </div>
          <div class="section" >
            <div class="section-row" >
              <div class="row full">
                <div class="cell padding-small" >
                  <div>
                    <form-label
                      value="${this.msg("",{id:"audit_nervatype"})}" 
                    ></form-label>
                  </div>
                  <form-select id="nervatype" 
                    label="${this.msg("",{id:"audit_nervatype"})}"
                    ?disabled="${e}" ?full="${!0}"
                    .onChange=${e=>this._onValueChange("nervatype",parseInt(e.value,10))}
                    .options=${d} .isnull="${!1}" value="${String(a)}" 
                  ></form-select>
                </div>
              </div>
              ${g?i`<div class="row full">
                <div class="cell padding-small" >
                  <div>
                    <form-label
                      value="${this.msg("",{id:"audit_subtype"})}" 
                    ></form-label>
                  </div>
                  <form-select id="subtype" 
                    label="${this.msg("",{id:"audit_subtype"})}"
                    .onChange=${e=>this._onValueChange("subtype",parseInt(e.value,10))}
                    .options=${c.filter((e=>e.type===m))} 
                    .isnull="${!1}" value="${String(n)}" ?full="${!0}"
                  ></form-select>
                </div>
              </div>`:l}
              <div class="row full">
                <div class="cell padding-small" >
                  <div>
                    <form-label
                      value="${this.msg("",{id:"audit_inputfilter"})}" 
                    ></form-label>
                  </div>
                  <form-select id="inputfilter" 
                    label="${this.msg("",{id:"audit_inputfilter"})}" ?full="${!0}"
                    .onChange=${e=>this._onValueChange("inputfilter",parseInt(e.value,10))}
                    .options=${u} .isnull="${!1}" value="${String(s)}" 
                  ></form-select>
                </div>
              </div>
              <div class="section-row" >
                <div class="cell padding-small" >
                  <form-label id="supervisor"
                    value="${this.msg("",{id:"audit_supervisor"})}"
                    leftIcon="${1===r?"CheckSquare":"SquareEmpty"}"
                    .style=${{cursor:"pointer"}} .iconStyle=${{cursor:"pointer"}}
                    @click=${()=>this._onValueChange("supervisor",1===r?0:1)}
                  ></form-label>
                </div>
              </div>
            </div>
          </div>
          <div class="section buttons" >
            <div class="section-row" >
              <div class="cell padding-small half" >
                <form-button id="btn_cancel" icon="Times"
                  @click=${()=>this._onModalEvent(p.CANCEL,{})} 
                  ?full="${!0}" label="${this.msg("",{id:"msg_cancel"})}"
                >${this.msg("",{id:"msg_cancel"})}</form-button>
              </div>
              <div class="cell padding-small half" >
                <form-button id="btn_ok" icon="Check"
                  @click=${()=>this._onModalEvent(p.OK,{value:{id:e,usergroup:t,nervatype:a,subtype:n,inputfilter:s,supervisor:r}})} 
                  ?disabled="${g&&!n}"
                  type="${o.PRIMARY}" ?full="${!0}" 
                  label="${this.msg("",{id:"msg_ok"})}"
                >${this.msg("",{id:"msg_ok"})}</form-button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>`}});const h=e`
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
`;customElements.define("modal-menu",class extends t{constructor(){super(),this.idKey=null,this.menu_id=0,this.fieldname="",this.description="",this.fieldtype=0,this.fieldtypeOptions=[],this.orderby=0}static get properties(){return{idKey:{type:Object},menu_id:{type:Number},fieldname:{type:String},description:{type:String},fieldtype:{type:Number},fieldtypeOptions:{type:Array},orderby:{type:Number}}}static get styles(){return[h]}_onModalEvent(e,t){this.onEvent&&this.onEvent.onModalEvent&&this.onEvent.onModalEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("modal_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}_onValueChange(e,t){this[e]=t}render(){const{idKey:e,menu_id:t,fieldname:l,description:a,orderby:n,fieldtype:s,fieldtypeOptions:r}=this;return i`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label leftIcon="Share"
                value="${this.msg("",{id:"title_menucmd"})}" 
                class="title-cell" ></form-label>
            </div>
            <div class="cell align-right" >
              <span id=${"closeIcon"} class="close-icon" 
                @click="${()=>this._onModalEvent(p.CANCEL,{})}">
                <form-icon iconKey="Times" ></form-icon>
              </span>
            </div>
          </div>
          <div class="section" >
            <div class="section-row" >
              <div class="row full">
                <div class="cell padding-small" >
                  <div>
                    <form-label
                      value="${this.msg("",{id:"menufields_fieldname"})}" 
                    ></form-label>
                  </div>
                  <form-input id="fieldname"
                    type="${c.TEXT}"
                    label="${this.msg("",{id:"menufields_fieldname"})}" 
                    .onChange=${e=>this._onValueChange("fieldname",e.value)}
                    value="${l}" ?full=${!0}
                  ></form-input>
                </div>
              </div>
              <div class="row full">
                <div class="cell padding-small" >
                  <div>
                    <form-label
                      value="${this.msg("",{id:"menufields_description"})}" 
                    ></form-label>
                  </div>
                  <form-input id="description"
                    type="${c.TEXT}"
                    label="${this.msg("",{id:"menufields_description"})}" 
                    .onChange=${e=>this._onValueChange("description",e.value)}
                    value="${a}" ?full=${!0}
                  ></form-input>
                </div>
              </div>
              <div class="row full">
                <div class="cell padding-small half" >
                  <div>
                    <form-label
                      value="${this.msg("",{id:"menufields_fieldtype"})}" 
                    ></form-label>
                  </div>
                  <form-select id="fieldtype" 
                    label="${this.msg("",{id:"menufields_fieldtype"})}" ?full="${!0}"
                    .onChange=${e=>this._onValueChange("fieldtype",parseInt(e.value,10))}
                    .options=${r} .isnull="${!1}" value="${String(s)}" 
                  ></form-select>
                </div>
                <div class="cell padding-small half" >
                  <div>
                    <form-label
                      value="${this.msg("",{id:"menufields_orderby"})}" 
                    ></form-label>
                  </div>
                  <form-number id="orderby" 
                    label="${this.msg("",{id:"menufields_orderby"})}"
                    ?integer="${!1}" value="${n}"
                    .onChange=${e=>this._onValueChange("orderby",e.value)}
                  ></form-number>
                </div>
              </div>
            </div>
          </div>
          <div class="section buttons" >
            <div class="section-row" >
              <div class="cell padding-small half" >
                <form-button id="btn_cancel" icon="Times"
                  @click=${()=>this._onModalEvent(p.CANCEL,{})} 
                  ?full="${!0}" label="${this.msg("",{id:"msg_cancel"})}"
                >${this.msg("",{id:"msg_cancel"})}</form-button>
              </div>
              <div class="cell padding-small half" >
                <form-button id="btn_ok" icon="Check"
                  ?disabled="${""===l||""===a}"
                  @click=${()=>this._onModalEvent(p.OK,{value:{id:e,menu_id:t,fieldname:l,description:a,fieldtype:s,orderby:n}})} 
                  type="${o.PRIMARY}" ?full="${!0}" 
                  label="${this.msg("",{id:"msg_ok"})}"
                >${this.msg("",{id:"msg_ok"})}</form-button>
              </div>
            </div>
          </div>
        </div>
      </div>
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
`;customElements.define("client-setting",class extends t{constructor(){super(),this.msg=e=>e,this.id=Math.random().toString(36).slice(2),this.side=a.AUTO,this.data={caption:"",icon:"",view:{},actions:{},current:void 0,audit:"",dataset:{},type:""},this.auditFilter={},this.username="",this.paginationPage=10,this.onEvent={},this.modalAudit=this.modalAudit.bind(this),this.modalMenu=this.modalMenu.bind(this)}static get properties(){return{id:{type:String},side:{type:String},data:{type:Object},auditFilter:{type:Object},username:{type:String},paginationPage:{type:Number},onEvent:{type:Object}}}static get styles(){return[f]}connectedCallback(){super.connectedCallback(),this.onEvent.setModule(this)}modalAudit({idKey:e,usergroup:t,nervatype:l,subtype:a,inputfilter:n,supervisor:o,typeOptions:s,subtypeOptions:r,inputfilterOptions:d,onEvent:p}){return i`<modal-audit
      idKey="${e}"
      usergroup="${t}"
      nervatype="${l}"
      subtype="${a}"
      inputfilter="${n}"
      supervisor="${o}"
      .typeOptions="${s}"
      .subtypeOptions="${r}"
      .inputfilterOptions="${d}"
      .onEvent=${p} .msg=${this.msg}
    ></modal-audit>`}modalMenu({idKey:e,menu_id:t,fieldname:l,description:a,orderby:n,fieldtype:o,fieldtypeOptions:s,onEvent:r}){return i`<modal-menu
      idKey="${e}"
      menu_id="${t}"
      fieldname="${l}"
      description="${a}"
      orderby="${n}"
      fieldtype="${o}"
      .fieldtypeOptions="${s}"
      .onEvent=${r} .msg=${this.msg}
    ></modal-menu>`}render(){const{side:e,data:t,auditFilter:a,username:n,paginationPage:o}=this;return i`<sidebar-setting
      id="${this.id}" side="${e}"
      username="${n}" .auditFilter="${a}"
      .module="${t}"
      .onEvent=${this.onEvent} .msg=${this.msg}
    ></sidebar-setting>
      <div class="page">
        ${t.current?i`<setting-form
          .data="${t}"
          .onEvent=${this.onEvent} .msg=${this.msg}
          paginationPage="${o}"
        ></setting-form>`:t.view?i`<setting-view
          .data="${t}"
          paginationPage="${o}"
          .onEvent=${this.onEvent} .msg=${this.msg}
        ></setting-view>`:l}
      </div>`}});
