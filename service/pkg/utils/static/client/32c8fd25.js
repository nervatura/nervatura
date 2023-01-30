import{i as e,s as t,y as i,b as a,x as l}from"./4e7ea0c6.js";import{a as n,i as o,t as s,e as r,l as d}from"./81d721ef.js";import"./885c7aee.js";import{d as c,k as p,B as m,S as v,T as u,E as h,h as b,P as g,b as f,M as $,I as y}from"./52a11ade.js";import"./a08248bd.js";import"./b825c8e3.js";import{e as x,n as _}from"./aa6beff6.js";import"./76f14517.js";import"./1b30e9af.js";import"./95ec07a4.js";const E=e`
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
`;customElements.define("sidebar-edit",class extends t{constructor(){super(),this.msg=e=>e,this.side=c.AUTO,this.view=p.EDIT,this.module={current:{},form_dirty:!1,dirty:!1,panel:{},dataset:{},group_key:""},this.newFilter=[],this.auditFilter={},this.forms={}}static get properties(){return{side:{type:String,reflect:!0},view:{type:String,reflect:!0},newFilter:{type:Array},auditFilter:{type:Object},module:{type:Object},forms:{type:Object}}}static get styles(){return[E]}_onSideEvent(e,t){this.onEvent&&this.onEvent.onSideEvent&&this.onEvent.onSideEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("side_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}itemMenu({id:e,selected:t,eventValue:a,label:l,iconKey:n,full:o,disabled:s,align:r,color:d}){const c=void 0!==t&&t,p=void 0===o||o,v=void 0!==s&&s,h=void 0===r?u.LEFT:r,b={"border-radius":"0","border-color":"rgba(var(--accent-1c), 0.2)"};return d&&(b.color=d,b.fill=d),i`<form-button 
        id="${e}" label="${l}"
        ?full="${p}" ?disabled="${v}" ?selected="${c}"
        align=${h}
        .style="${b}"
        icon="${n}" type="${m.PRIMARY}"
        @click=${()=>this._onSideEvent(...a)} 
      >${l}</form-button>`}editItems(e){const{current:t,dirty:a,form_dirty:l,dataset:o}=this.module,s=void 0===e?{}:e,r=[];if((!0===s.back||t.form)&&(r.push(this.itemMenu({id:"cmd_back",selected:!0,eventValue:[v.BACK,{}],label:this.msg("",{id:"label_back"}),iconKey:"Reply",full:!1})),r.push(i`<hr id="back_sep" class="separator" />`)),!0===s.arrow&&(r.push(this.itemMenu({id:"cmd_arrow_left",eventValue:[v.PREV_NUMBER,{}],label:this.msg("",{id:"label_previous"}),iconKey:"ArrowLeft"})),r.push(this.itemMenu({id:"cmd_arrow_right",eventValue:[v.NEXT_NUMBER,{}],label:this.msg("",{id:"label_next"}),iconKey:"ArrowRight",align:u.RIGHT})),r.push(i`<hr id="arrow_sep" class="separator" />`)),s.state&&"normal"!==s.state){const e="deleted"===s.state?"rgb(var(--functional-red))":"cancellation"===s.state?"rgb(var(--functional-yellow))":"rgba(var(--accent-1c), 0.85)",t=["closed","readonly"].includes(s.state)?"Lock":"ExclamationTriangle";r.push(i`<div key="cmd_state" class="state-label" >
        <form-icon iconKey="${t}" .style="${{fill:e}}" ></form-icon>
        <span style="${n({color:e,fill:e,"vertical-align":"middle"})}" >${this.msg("",{id:`label_${s.state}`})}</span>
      </div>`),r.push(i`<hr id="state_sep" class="separator" />`)}return!1!==s.save&&r.push(this.itemMenu({id:"cmd_save",selected:!!(t.form&&l||!t.form&&a),eventValue:[v.SAVE,{}],label:this.msg("",{id:"label_save"}),iconKey:"Check"})),!1!==s.delete&&"normal"===s.state&&r.push(this.itemMenu({id:"cmd_delete",eventValue:[v.DELETE,{}],label:this.msg("",{id:"label_delete"}),iconKey:"Times"})),!1===s.new||"normal"!==s.state||t.form||r.push(this.itemMenu({id:"cmd_new",eventValue:[v.NEW,[{}]],label:this.msg("",{id:"label_new"}),iconKey:"Plus"})),!0===s.trans&&(r.push(i`<hr id="trans_sep" class="separator" />`),!1!==s.copy&&r.push(this.itemMenu({id:"cmd_copy",eventValue:[v.COPY,{value:"normal"}],label:this.msg("",{id:"label_copy"}),iconKey:"Copy"})),!1!==s.create&&r.push(this.itemMenu({id:"cmd_create",eventValue:[v.COPY,{value:"create"}],label:this.msg("",{id:"label_create"}),iconKey:"Sitemap"})),!0===s.corrective&&"normal"===s.state&&r.push(this.itemMenu({id:"cmd_corrective",eventValue:[v.COPY,{value:"amendment"}],label:this.msg("",{id:"label_corrective"}),iconKey:"Share"})),!0===s.cancellation&&"cancellation"!==s.state&&r.push(this.itemMenu({id:"cmd_cancellation",eventValue:[v.COPY,{value:"cancellation"}],label:this.msg("",{id:"label_cancellation"}),iconKey:"Undo"})),!0===s.formula&&r.push(this.itemMenu({id:"cmd_formula",eventValue:[v.CHECK,[{},h.LOAD_FORMULA]],label:this.msg("",{id:"label_formula"}),iconKey:"Magic"}))),!0===s.link&&r.push(this.itemMenu({id:"cmd_link",eventValue:[v.LINK,{type:s.link_type,field:s.link_field}],label:s.link_label,iconKey:"Link"})),!0===s.password&&r.push(this.itemMenu({id:"cmd_password",eventValue:[v.PASSWORD,{}],label:this.msg("",{id:"title_password"}),iconKey:"Lock"})),!0===s.shipping&&(r.push(this.itemMenu({id:"cmd_shipping_all",eventValue:[v.SHIPPING_ADD_ALL,{}],label:this.msg("",{id:"shipping_all_label"}),iconKey:"Plus"})),r.push(this.itemMenu({id:"cmd_shipping_create",selected:o.shiptemp&&o.shiptemp.length>0,eventValue:[v.SHIPPING_CREATE,{}],label:this.msg("",{id:"shipping_create_label"}),iconKey:"Check"}))),!0===s.more&&(r.push(i`<hr id="more_sep_1" class="separator" />`),!1!==s.report&&r.push(this.itemMenu({id:"cmd_report",eventValue:[v.REPORT_SETTINGS,{}],label:this.msg("",{id:"label_report"}),iconKey:"ChartBar"})),!0===s.search&&r.push(this.itemMenu({id:"cmd_search",eventValue:[v.SEARCH_QUEUE,{}],label:this.msg("",{id:"label_search"}),iconKey:"Search"})),!0===s.export_all&&"normal"===s.state&&r.push(this.itemMenu({id:"cmd_export_all",eventValue:[v.EXPORT_QUEUE_ALL,{}],label:this.msg("",{id:"label_export_all"}),iconKey:"Download"})),!0===s.print&&r.push(this.itemMenu({id:"cmd_print",eventValue:[v.CREATE_REPORT,{value:"print"}],label:this.msg("",{id:"label_print"}),iconKey:"Print"})),!0===s.export_pdf&&"normal"===s.state&&r.push(this.itemMenu({id:"cmd_export_pdf",eventValue:[v.CREATE_REPORT,{value:"pdf"}],label:this.msg("",{id:"label_export_pdf"}),iconKey:"Download"})),!0===s.export_xml&&"normal"===s.state&&r.push(this.itemMenu({id:"cmd_export_xml",eventValue:[v.CREATE_REPORT,{value:"xml"}],label:this.msg("",{id:"label_export_xml"}),iconKey:"Code"})),!0===s.export_csv&&"normal"===s.state&&r.push(this.itemMenu({id:"cmd_export_csv",eventValue:[v.CREATE_REPORT,{value:"csv"}],label:this.msg("",{id:"label_export_csv"}),iconKey:"Download"})),!0===s.export_event&&"normal"===s.state&&r.push(this.itemMenu({id:"cmd_export_event",eventValue:[v.EXPORT_EVENT,{}],label:this.msg("",{id:"label_export_event"}),iconKey:"Calendar"})),r.push(i`<hr id="more_sep_2" class="separator" />`),!1!==s.bookmark&&"normal"===s.state&&r.push(this.itemMenu({id:"cmd_bookmark",eventValue:[v.SAVE_BOOKMARK,{value:s.bookmark}],label:this.msg("",{id:"label_bookmark"}),iconKey:"Star"})),!1!==s.help&&r.push(this.itemMenu({id:"cmd_help",eventValue:[v.HELP,{value:s.help}],label:this.msg("",{id:"label_help"}),iconKey:"QuestionCircle"}))),!0!==s.more&&void 0!==s.help&&(r.push(i`<hr id="help_sep" class="separator" />`),r.push(this.itemMenu({id:"cmd_help",eventValue:[v.HELP,{value:s.help}],label:this.msg("",{id:"label_help"}),iconKey:"QuestionCircle"}))),r}newItems(){const{group_key:e}=this.module,t=[];return this.newFilter[0].length>0&&t.push(i`<div class="row full">
        ${this.itemMenu({id:"new_transitem_group",selected:"new_transitem"===e,eventValue:[v.CHANGE,{fieldname:"group_key",value:"new_transitem"}],label:this.msg("",{id:"search_transitem"}),iconKey:"FileText"})}
        ${"new_transitem"===e?i`<div class="row full panel-group" >
          ${this.newFilter[0].map((e=>"all"===this.auditFilter.trans[e][0]?this.itemMenu({id:e,eventValue:[v.NEW,[{ntype:"trans",ttype:e}]],label:this.msg("",{id:`title_${e}`}),iconKey:"FileText",color:"rgb(var(--functional-blue))"}):a))}
        </div>`:a}
      </div>`),this.newFilter[1].length>0&&t.push(i`<div class="row full">
        ${this.itemMenu({id:"new_transpayment_group",selected:"new_transpayment"===e,eventValue:[v.CHANGE,{fieldname:"group_key",value:"new_transpayment"}],label:this.msg("",{id:"search_transpayment"}),iconKey:"Money"})}
        ${"new_transpayment"===e?i`<div class="row full panel-group" >
          ${this.newFilter[1].map((e=>"all"===this.auditFilter.trans[e][0]?this.itemMenu({id:e,eventValue:[v.NEW,[{ntype:"trans",ttype:e}]],label:this.msg("",{id:`title_${e}`}),iconKey:"Money",color:"rgb(var(--functional-blue))"}):a))}
        </div>`:a}
      </div>`),this.newFilter[2].length>0&&t.push(i`<div class="row full">
        ${this.itemMenu({id:"new_transmovement_group",selected:"new_transmovement"===e,eventValue:[v.CHANGE,{fieldname:"group_key",value:"new_transmovement"}],label:this.msg("",{id:"search_transmovement"}),iconKey:"Truck"})}
        ${"new_transmovement"===e?i`<div class="row full panel-group" >
          ${this.newFilter[2].map((e=>"all"===this.auditFilter.trans[e][0]?"delivery"===e?[this.itemMenu({id:"shipping",eventValue:[v.NEW,[{ntype:"trans",ttype:"shipping"}]],label:this.msg("",{id:`title_${e}`}),iconKey:this.forms[e]().options.icon,color:"rgb(var(--functional-blue))"}),this.itemMenu({id:e,eventValue:[v.NEW,[{ntype:"trans",ttype:e}]],label:this.msg("",{id:"title_transfer"}),iconKey:this.forms[e]().options.icon,color:"rgb(var(--functional-blue))"})]:this.itemMenu({id:e,eventValue:[v.NEW,[{ntype:"trans",ttype:e}]],label:this.msg("",{id:`title_${e}`}),iconKey:this.forms[e]().options.icon,color:"rgb(var(--functional-blue))"}):a))}
        </div>`:a}
      </div>`),this.newFilter[3].length>0&&t.push(i`<div class="row full">
        ${this.itemMenu({id:"new_resources_group",selected:"new_resources"===e,eventValue:[v.CHANGE,{fieldname:"group_key",value:"new_resources"}],label:this.msg("",{id:"title_resources"}),iconKey:"Wrench"})}
        ${"new_resources"===e?i`<div class="row full panel-group" >
          ${this.newFilter[3].map((e=>"all"===this.auditFilter[e][0]?this.itemMenu({id:e,eventValue:[v.NEW,[{ntype:e,ttype:null}]],label:this.msg("",{id:`title_${e}`}),iconKey:this.forms[e]().options.icon,color:"rgb(var(--functional-blue))"}):a))}
        </div>`:a}
      </div>`),t}render(){const{current:e,panel:t}=this.module;return i`<div class="sidebar ${"auto"!==this.side?this.side:""}" >
      ${e.form||"transitem_shipping"===e.form_type?a:i`<div class="row full container">
          <div class="cell half">
            ${this.itemMenu({id:"state_new",selected:!(this.view===p.EDIT&&e.item),eventValue:[v.CHANGE,{fieldname:"side_view",value:"new"}],label:this.msg("",{id:"label_new"}),iconKey:"Plus",align:u.CENTER})}
          </div>
          <div class="cell half">
            ${this.itemMenu({id:"state_edit",selected:!(this.view!==p.EDIT||!e.item),eventValue:[v.CHANGE,{fieldname:"side_view",value:"edit"}],label:this.msg("",{id:"label_edit"}),iconKey:"Edit",disabled:!e.item,align:u.CENTER})}
          </div>
        </div>`}
      ${this.view===p.EDIT&&e.form||this.view===p.EDIT&&e.item?this.editItems(t):this.newItems()}
     </div>`}});const w=e`
:host {
  font-family: var(--font-family);
  font-size: var(--font-size);
  color: var(--text-1);
  fill: var(--text-1);
}
div {
  box-sizing: border-box;
}
.panel {
  width: 100%;
  background-color: rgba(var(--base-2), 1);
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  border-top: none;
  border-bottom: none;
}
.row {
  display: table;
}
.full { 
  width: 100%; 
}
`;customElements.define("edit-main",class extends t{constructor(){super(),this.msg=e=>e,this.id=Math.random().toString(36).slice(2),this.audit="",this.current={},this.template={},this.dataset={},this.onEvent={}}static get properties(){return{id:{type:String},audit:{type:String},current:{type:Object},template:{type:Object},dataset:{type:Object},onEvent:{type:Object}}}static get styles(){return[w]}_onEditEvent(e,t){this.onEvent.onEditEvent&&this.onEvent.onEditEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("edit_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}render(){const{audit:e,current:t,dataset:l,template:n}=this;return i`<div id="${this.id}" class="panel" >
      ${n.rows.map(((a,o)=>i`<form-row
        id=${`row_${o}`}
        .row="${a}"
        .values="${t.item}"
        .options="${n.options}"
        .data="${{audit:e,current:t,dataset:l}}"
        .onEdit=${e=>this._onEditEvent(b.EDIT_ITEM,e)}
        .onEvent=${(...e)=>this._onEditEvent(...e)}
        .onSelector=${(...e)=>this._onEditEvent(b.SELECTOR,[...e])}
        .msg=${this.msg}
      ></form-row>`))}
      ${"report"===t.type&&t.fieldvalue.length>0?i`<div class="row full">
          ${t.fieldvalue.map(((a,o)=>i`<form-row
            id=${`row_${o}`}
            .row="${a}"
            .values="${a}"
            .options="${n.options}"
            .data="${{audit:e,current:t,dataset:l}}"
            .onEdit=${e=>this._onEditEvent(b.EDIT_ITEM,e)}
            .onEvent=${(...e)=>this._onEditEvent(...e)}
            .onSelector=${(...e)=>this._onEditEvent(b.SELECTOR,[...e])}
            .msg=${this.msg}
          ></form-row>`))}
        </div>`:a}
    </div>`}});const k=e`
:host {
  font-family: var(--font-family);
  font-size: var(--font-size);
  color: var(--text-1);
  fill: var(--text-1);
}
div {
  box-sizing: border-box;
}
.panel {
  width: 100%;
  background-color: rgba(var(--base-2), 1);
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  border-top: none;
  border-bottom: none;
}
.row {
  display: table;
}
.cell {
  display: table-cell;
  vertical-align: middle;
}
.full { 
  width: 100%; 
}
.container-row {
  display: table;
  width: 100%;
  padding: 8px;
  border-bottom: 1px solid rgba(var(--neutral-1), 0.2);
}
.paginator-cell {
  display: table-cell;
  vertical-align: middle;
  padding-left: 8px;
  float: right;
}
.padding-small { 
  padding: 4px 8px; 
}
@media (max-width:600px){
  .mobile, .paginator-cell { 
    display: block; 
    width: 100%; 
  }
  .padding-small { 
    padding: 2px 4px; 
  }
}
`;customElements.define("edit-meta",class extends t{constructor(){super(),this.msg=e=>e,this.id=Math.random().toString(36).slice(2),this.audit="",this.current={},this.dataset={},this.pageSize=5,this.onEvent={}}static get properties(){return{id:{type:String},audit:{type:String},current:{type:Object},dataset:{type:Object},pageSize:{type:Number},onEvent:{type:Object}}}static get styles(){return[k]}connectedCallback(){super.connectedCallback(),this.currentPage=this.current.page||1,this.currentPage<1&&(this.currentPage=1)}_onPagination(e,t){if("setPageSize"===e)return this.currentPage=1,void(this.pageSize=t);this.currentPage=t,this.requestUpdate()}_onEditEvent(e,t){this.onEvent.onEditEvent&&this.onEvent.onEditEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("edit_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}prepareRows(){const{current:e,dataset:t}=this;let i=[];return e.fieldvalue.forEach((e=>{const a=t.deffield.filter((t=>t.fieldname===e.fieldname))[0];if(1===a.visible&&0===e.deleted){const l=t.groups.filter((e=>e.id===a.fieldtype))[0].groupvalue;let n=e.value,o=l;if(["customer","tool","trans","transitem","transmovement","transpayment","product","project","employee","place"].includes(l)){const i=t.deffield_prop.filter((t=>t.ftype===l&&t.id===parseInt(e.value,10)))[0];i&&(n=i.description),o="selector"}"urlink"===l&&(o="text"),"valuelist"===l&&(n=a.valuelist.split("|")),i=[...i,{rowtype:"fieldvalue",id:e.id,name:"fieldvalue_value",fieldname:e.fieldname,value:e.value,notes:e.notes||"",label:a.description,description:n,disabled:!!a.readonly,fieldtype:l,datatype:o}]}})),i}renderRows(e,t){const{audit:a,current:l,dataset:n}=this;let o=e;if(t>1){const e=this.currentPage>t?t:this.currentPage,i=(e-1)*this.pageSize,a=e*this.pageSize;o=o.slice(i,a)}return o.map(((e,t)=>i`<form-row
        id=${`row_${t}`}
        .row="${e}"
        .values="${e}"
        .options="${{}}"
        .data="${{audit:a,current:l,dataset:n}}"
        .onEdit=${e=>this._onEditEvent(b.EDIT_ITEM,e)}
        .onEvent=${(...e)=>this._onEditEvent(...e)}
        .onSelector=${(...e)=>this._onEditEvent(b.SELECTOR,[...e])}
        .msg=${this.msg}
      ></form-row>`))}deffields(){const{dataset:e,current:t}=this,i=e.groups.filter((e=>"nervatype"===e.groupname&&e.groupvalue===t.type))[0].id;return"trans"===t.type?e.deffield.filter((e=>e.nervatype===i&&1===e.visible)).filter((e=>e.subtype===t.item.transtype||null===e.subtype)).map((e=>({value:e.fieldname,text:e.description}))):e.deffield.filter((e=>e.nervatype===i&&1===e.visible)).map((e=>({value:e.fieldname,text:e.description})))}render(){const{audit:e,current:t}=this,l=this.prepareRows(),n=Math.ceil(l.length/this.pageSize);return i`<div id="${this.id}" class="panel" >
      ${"readonly"!==e||n>1?i`<div class="container-row">
        ${"readonly"!==e?i`<div class="cell mobile">
          <div class="cell padding-small" >
            <form-select id="sel_deffield"
              label="${this.msg("",{id:"fields_view"})}"
              .onChange=${e=>this._onEditEvent(b.CHANGE,{fieldname:"deffield",value:e.value})}
              .options=${this.deffields()} 
              .isnull="${!0}" value="${t.deffield||""}" >
            </form-select>
          </div>
          ${t.deffield&&""!==t.deffield?i`<div class="cell" >
            <form-button id="btn_new" 
              label="${this.msg("",{id:"label_new"})}"
              .style="${{padding:"6px 16px"}}"
              icon="Plus" type="${m.BORDER}"
              @click=${()=>this._onEditEvent(b.CHECK_EDITOR,[{fieldname:t.deffield},h.NEW_FIELDVALUE])} 
            >${this.msg("",{id:"label_new"})}</form-button>
          </div>`:a}
        </div>`:a}
        ${n>1?i`<div class="paginator-cell">
          <form-pagination id="${`${this.id}_top_pagination`}"
            pageIndex=${this.currentPage} pageSize=${this.pageSize} pageCount=${n} 
            ?canPreviousPage=${this.currentPage>1} 
            ?canNextPage=${this.currentPage<n} 
            ?hidePageSize=${this.hidePaginatonSize}
            .onEvent=${(e,t)=>this._onPagination(e,t)} ></form-pagination>
        </div>`:a}
      </div>`:a}
      ${this.renderRows(l,n)}
    </div>`}});const C=e`
:host {
  font-family: var(--font-family);
  font-size: var(--font-size);
  color: var(--text-1);
  fill: var(--text-1);
}
div {
  box-sizing: border-box;
}
.panel {
  width: 100%;
  background-color: rgba(var(--base-2), 1);
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  border-top: none;
  border-bottom: none;
}
.row {
  display: table;
}
.full { 
  width: 100%; 
}
.title-cell {
  display: table-cell;
  font-weight: bold;
  padding: 8px 16px;
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  background-color: rgba(var(--neutral-1), 0.1);
}
.title-pre {
  padding-right: 6px;
  opacity: 0.5;
}
@media (max-width:600px){
  .title-cell {
    padding: 4px 8px;
  }
}
`;customElements.define("edit-item",class extends t{constructor(){super(),this.msg=e=>e,this.id=Math.random().toString(36).slice(2),this.audit="",this.current={},this.dataset={},this.onEvent={}}static get properties(){return{id:{type:String},audit:{type:String},current:{type:Object},dataset:{type:Object},onEvent:{type:Object}}}static get styles(){return[C]}_onEditEvent(e,t){this.onEvent.onEditEvent&&this.onEvent.onEditEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("edit_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}render(){const{audit:e,current:t,dataset:a}=this;return i`<div class="row full">
      <div class="title-cell" >
        <span class="title-pre" >${null===t.form.id?this.msg("",{id:"label_new"}):String(t.form.id)}</span>
        <span>${t.form_template.options.title}</span>
      </div>
    </div>
    <div id="${this.id}" class="panel" >
      ${t.form_template.rows.map(((l,n)=>i`<form-row
        id=${`row_${n}`}
        .row="${l}"
        .values="${t.form}"
        .options="${t.form_template.options}"
        .data="${{audit:e,current:t,dataset:a}}"
        .onEdit=${e=>this._onEditEvent(b.EDIT_ITEM,e)}
        .onEvent=${(...e)=>this._onEditEvent(...e)}
        .onSelector=${(...e)=>this._onEditEvent(b.SELECTOR,[...e])}
        .msg=${this.msg}
      ></form-row>`))}
    </div>`}});class S extends o{constructor(e){if(super(e),this.it=a,e.type!==s.CHILD)throw Error(this.constructor.directiveName+"() can only be used in child bindings")}render(e){if(e===a||null==e)return this._t=void 0,this.it=e;if(e===l)return e;if("string"!=typeof e)throw Error(this.constructor.directiveName+"() called with a non-string value");if(e===this.it)return this._t;this.it=e;const t=[e];return t.raw=t,this._t={_$litType$:this.constructor.resultType,strings:t,values:[]}}}S.directiveName="unsafeHTML",S.resultType=1;const T=r(S),M=e`
:host {
  font-family: var(--font-family);
  font-size: var(--font-size);
  color: var(--text-1);
  fill: var(--text-1);
}
div {
  box-sizing: border-box;
}
.panel {
  width: 100%;
  background-color: rgba(var(--base-0), 1);
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  border-top: none;
  border-bottom: none;
}
.actionbar {
  width: 100%;
  background-color: rgba(var(--base-2), 1);
}
.cell {
  display: table-cell;
  vertical-align: middle;
}
.padding-tiny { 
  padding: 4px 12px 4px 4px; 
}
.padding-small { 
  padding: 4px 8px; 
}
.rtf-editor {
  width: 100%;
  min-height: 65px;
  background-color: rgba(var(--base-0), 1);
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  border-left: none;
  border-right: none;
}
.editor-content {
  height: 300px;
  outline: 0;
  overflow-y: auto;
  padding: 10px; 
}
@media (max-width:600px){
  .padding-tiny { 
    padding: 4px 12px 4px 2px; 
  }
  .padding-small { 
    padding: 4px 4px; 
  }
  .mobile{ 
    display: block; 
    width: 100%; 
  }
}
`;customElements.define("edit-note",class extends t{constructor(){super(),this.msg=e=>e,this.id=Math.random().toString(36).slice(2),this.value="",this.patternId=void 0,this.patterns=[],this.readOnly=!1,this.bold=!1,this.italic=!1,this.onEvent={},this.editorRef=x()}static get properties(){return{id:{type:String},value:{type:String},patternId:{type:Number},patterns:{type:Array},readOnly:{type:Boolean},bold:{type:Boolean},italic:{type:Boolean},onEvent:{type:Object}}}static get styles(){return[M]}connectedCallback(){super.connectedCallback(),this._value=this.value}disconnectedCallback(){this._onEditEvent(b.EDIT_ITEM,{name:"fnote",value:this._value}),super.disconnectedCallback()}_onEditEvent(e,t){this.onEvent.onEditEvent&&this.onEvent.onEditEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("edit_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}_onInput(){this._value=String(this.editorRef.value.innerHTML).split("--\x3e")[1]}_onContentState(){this.bold=document.queryCommandState("bold"),this.italic=document.queryCommandState("italic")}_setContentState(e){document.execCommand(e,!1),this.editorRef.value.focus(),this._onContentState()}_onLostFocus(){this.bold=!1,this.italic=!1}render(){return i`<div id="${this.id}" class="panel" >
      ${this.readOnly?a:i`<div class="actionbar padding-small">
        <div class="cell padding-tiny">
          <form-button id="btn_pattern_default" 
            icon="Home" type="${m.BORDER}"
            label="${this.msg("",{id:"pattern_default"})}"
            .style=${{padding:"8px 12px"}}
            @click=${()=>this._onEditEvent(b.SET_PATTERN,{key:"default"})}
          ></form-button>
          <form-button id="btn_pattern_load" 
            icon="Download" type="${m.BORDER}"
            label="${this.msg("",{id:"pattern_load"})}"
            .style=${{padding:"8px 12px"}}
            @click=${()=>this._onEditEvent(b.SET_PATTERN,{key:"load",ref:this})}
          ></form-button>
          <form-button id="btn_pattern_save" 
            icon="Upload" type="${m.BORDER}"
            label="${this.msg("",{id:"pattern_save"})}"
            .style=${{padding:"8px 12px"}}
            @click=${()=>this._onEditEvent(b.SET_PATTERN,{key:"save",text:this._value})}
          ></form-button>
        </div>
        <div class="cell padding-tiny">
          <form-button id="btn_pattern_new" 
            icon="Plus" type="${m.BORDER}"
            label="${this.msg("",{id:"pattern_new"})}"
            .style=${{padding:"8px 12px"}}
            @click=${()=>this._onEditEvent(b.SET_PATTERN,{key:"new"})}
          ></form-button>
          <form-button id="btn_pattern_delete" 
            icon="Times" type="${m.BORDER}"
            label="${this.msg("",{id:"pattern_delete"})}"
            .style=${{padding:"8px 12px"}}
            @click=${()=>this._onEditEvent(b.SET_PATTERN,{key:"delete"})}
          ></form-button>
        </div>
        <div class="cell mobile">
          <div class="cell padding-tiny">
            <form-select id="sel_pattern"
              label="${this.msg("",{id:"title_pattern"})}"
              .onChange=${e=>this._onEditEvent(b.CHANGE,{fieldname:"template",value:e.value})}
              .options=${this.patterns.map((e=>({value:String(e.id),text:e.description+(1===e.defpattern?"*":"")})))} 
              .isnull="${!0}" value="${this.patternId?String(this.patternId):""}" >
            </form-select>
          </div>
          <div class="cell padding-tiny">
            <form-button id="btn_bold" 
              type="${m.BORDER}"
              label="B" ?selected=${this.bold}
              .style=${{padding:"8px 12px"}}
              @click=${()=>this._setContentState("bold")}
            >B</form-button>
            <form-button id="btn_italic" 
              type="${m.BORDER}"
              label="I" ?selected=${this.italic}
              .style=${{padding:"8px 12px","font-style":"italic"}}
              @click=${()=>this._setContentState("italic")}
            >I</form-button>
          </div>
        </div>
      </div>`}
      <div class="rtf-editor" >
        <div id="editor" ${_(this.editorRef)}
          class="editor-content"
          @input=${this._onInput} 
          @keyup=${this._onContentState} @mouseup=${this._onContentState} @blur=${this._onLostFocus}
          contentEditable="${!this.readOnly}" >${T(this.value)}</div>
      </div>
    </div>`}});const O=e`
:host {
  font-family: var(--font-family);
  font-size: var(--font-size);
  color: var(--text-1);
  fill: var(--text-1);
}
div {
  box-sizing: border-box;
}
.panel {
  width: 100%;
  background-color: rgba(var(--base-2), 1);
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  border-top: none;
  border-bottom: none;
}
.row {
  display: table;
}
.cell {
  display: table-cell;
  vertical-align: middle;
}
.full { 
  width: 100%; 
}
.container-row {
  display: table;
  width: 100%;
  border-bottom: 1px solid rgba(var(--neutral-1), 0.2);
}
.total-cell {
  display: table-cell;
  vertical-align: middle;
  padding: 8px 16px;
  text-align: right;
}
.total-label {
  font-size: 13px;
  white-space: nowrap;
  font-weight: bold;
}
.total-value {
  font-size: 13px;
  white-space: nowrap;
  font-weight: bold;
  padding: 4px 6px;
  background-color: rgba(var(--base-3), 1);
  border: 1px solid rgba(var(--neutral-1), 0.2);
  border-radius: 3px;
  margin-left: 8px;
}
@media (max-width:600px){
  .mobile { 
    display: block; 
    width: 100%; 
  }
  .total-cell { 
    padding: 4px 8px;; 
  }
}
`;customElements.define("edit-view",class extends t{constructor(){super(),this.msg=e=>e,this.id=Math.random().toString(36).slice(2),this.viewName="",this.audit="",this.current={},this.template={},this.dataset={},this.pageSize=10,this.onEvent={}}static get properties(){return{id:{type:String},viewName:{type:String},audit:{type:String},current:{type:Object},template:{type:Object},dataset:{type:Object},pageSize:{type:Number},onEvent:{type:Object}}}static get styles(){return[O]}_onEditEvent(e,t){this.onEvent.onEditEvent&&this.onEvent.onEditEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("edit_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}render(){const{viewName:e,audit:t,template:l,dataset:n,current:o,pageSize:s}=this,r=l.view[e],c=n[r.data]||[],p=void 0===r.edited||r.edited;let m=r.actions;void 0===m&&(m={new:{action:f.NEW_EDITOR_ITEM,fkey:e},edit:{action:f.EDIT_EDITOR_ITEM,fkey:e},delete:{action:f.DELETE_EDITOR_ITEM,fkey:e}}),"all"!==t&&(m={...m,new:null,delete:null});const v=void 0!==r.edit_icon?[r.edit_icon,void 0,void 0]:["Edit",24,21.3],u=void 0!==r.delete_icon?[r.delete_icon,void 0,void 0]:["Times",19,27.6],h=void 0!==r.new_icon?r.new_icon:"Plus",$=void 0!==r.new_label?r.new_label:this.msg("",{id:"label_new"});let y={};return"table"===r.type&&(p&&(m.edit||m.delete)&&(y={...y,edit:{columnDef:{id:"edit",Header:"",headerStyle:{},Cell:({row:e})=>{const t=null!==m.edit?i`<form-icon id=${`edit_${e.id}`}
                  iconKey=${v[0]} width=${v[1]} height=${v[2]}
                  .style=${{cursor:"pointer",fill:"rgb(var(--functional-green))"}}
                  @click=${t=>{t.stopPropagation(),this._onEditEvent(b.FORM_ACTION,{params:m.edit,row:e,ref:this})}}
                ></form-icon>`:void 0,a=null!==m.delete?i`<form-icon id=${`delete_${e.id}`}
                  iconKey=${u[0]} width=${u[1]} height=${u[2]}
                  .style=${{cursor:"pointer",fill:"rgb(var(--functional-red))","margin-left":null!==m.edit?"8px":"0"}}
                  @click=${t=>{t.stopPropagation(),this._onEditEvent(b.FORM_ACTION,{params:m.delete,row:e})}}
                ></form-icon>`:void 0;return i`${d(t)}${d(a)}`},cellStyle:{width:30,padding:"4px 3px 3px 8px"}}}}),y={...y,...r.fields}),i`<div id="${this.id}" class="panel" >
      ${r.total?i`<div class="container-row">
        <div class="total-cell">
          <span class="total-label" >${`${r.total[Object.keys(r.total)[0]]}:`}</span>
          <span class="total-value" >${new Intl.NumberFormat("default").format(n[o.type][0][Object.keys(r.total)[0]])}</span>
        </div>
        <div class="total-cell">
          <span class="total-label" >${`${r.total[Object.keys(r.total)[1]]}:`}</span>
          <span class="total-value" >${new Intl.NumberFormat("default").format(n[o.type][0][Object.keys(r.total)[1]])}</span>
        </div>
        <div class="total-cell">
          <span class="total-label" >${`${r.total[Object.keys(r.total)[2]]}:`}</span>
          <span class="total-value" >${new Intl.NumberFormat("default").format(n[o.type][0][Object.keys(r.total)[2]])}</span>
        </div>
      </div>`:a}
      <div class="row full" >
        ${"table"===r.type?i`<form-table id="view_table"
          .onAddItem=${p&&m.new?()=>this._onEditEvent(b.FORM_ACTION,{params:m.new}):void 0}
          .fields=${y} .rows=${c} ?tableFilter=${!0}
          filterPlaceholder="${this.msg("",{id:"placeholder_filter"})}"
          labelYes=${this.msg("",{id:"label_yes"})} labelNo=${this.msg("",{id:"label_no"})} 
          labelAdd=${$} addIcon=${h} 
          pageSize=${s} pagination="${g.TOP}"
        ></form-table>`:i`<form-list id="view_list"
          .rows=${c} labelAdd=${$} addIcon=${h}
          editIcon=${v[0]} deleteIcon=${u[0]} ?listFilter=${!0} 
          filterPlaceholder=${this.msg("",{id:"placeholder_filter"})}
          pageSize=${s} pagination="${g.TOP}"
          .onEdit=${p&&m.edit?e=>this._onEditEvent(b.FORM_ACTION,{params:m.edit,row:e}):void 0}
          .onDelete=${p&&m.delete?e=>this._onEditEvent(b.FORM_ACTION,{params:m.delete,row:e}):void 0}
          .onAddItem=${p&&m.new?()=>this._onEditEvent(b.FORM_ACTION,{params:m.new}):void 0}
        ></form-list>`}
      </div>
    </div>`}});const R=e`
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
}
.section-container::after, .section-container::before { 
  content: ""; 
  display: table; 
  clear: both; 
}
@media (max-width:600px){
  .section-container { 
    padding: 8px; 
  }
}
`;customElements.define("edit-editor",class extends t{constructor(){super(),this.msg=e=>e,this.id=Math.random().toString(36).slice(2),this.caption="",this.audit="",this.current={},this.template={},this.dataset={},this.paginationPage=10,this.selectorPage=5,this.onEvent={}}static get properties(){return{id:{type:String},caption:{type:String},audit:{type:String},current:{type:Object},template:{type:Object},dataset:{type:Object},paginationPage:{type:Number},selectorPage:{type:Number},onEvent:{type:Object}}}static get styles(){return[R]}_onEditEvent(e,t){this.onEvent.onEditEvent&&this.onEvent.onEditEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("edit_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}_onNoteBlur(e){e.target._value!==e.target.value&&this.onEvent.onEditEvent({key:b.EDIT_ITEM,data:{name:"fnote",value:e.target._value}})}render(){const{caption:e,audit:t,template:l,current:n,dataset:o}=this,s=(e,t,a,l)=>i`
      <div class="row full" >
        <div class="cell" >
          <form-button id="${`btn_${e}`}" 
            icon="${a}" ?full="${!0}"
            label="${this.msg(t,{id:t})}" align=${u.LEFT}
            .style=${{"border-radius":0,"margin-top":"2px"}}
            badge="${d(l)}"
            @click=${()=>this._onEditEvent(b.CHANGE,{fieldname:"view",value:n.view===e?"":e})}
          >${this.msg(t,{id:t})}</form-button>
        </div>
      </div>`;return i`<div class="panel" >
      <div class="panel-title">
        <div class="cell">
          <form-label class="title-cell"
            value="${e}" leftIcon="${l.options.icon}"
          ></form-label>
        </div>
      </div>
      ${n.form?i`<div class="section-container" >
          <edit-item
            audit="${t}" .current="${n}" .dataset="${o}"
            .onEvent=${this.onEvent} .msg=${this.msg}
          ></edit-item>
        </div>`:i`<div class="section-container" >
          ${s("form",(()=>{let e=n.item[l.options.title_field];return"printqueue"===n.type?e=l.options.title_field:null===n.item.id&&(e=`${this.msg("",{id:"label_new"})} ${l.options.title}`),e})(),l.options.icon)}
          ${"form"===n.view?i`<edit-main
              audit="${t}" .current="${n}" .template="${l}" .dataset="${o}"
              .onEvent=${this.onEvent} .msg=${this.msg}
            ></edit-main>`:a}

          ${null===n.item.id&&!l.options.search_form||void 0===this.dataset.fieldvalue||null===n.item||!0!==l.options.fieldvalue?a:i`${s("fieldvalue","fields_view",l.options.icon,n.fieldvalue.filter((e=>1===o.deffield.filter((t=>t.fieldname===e.fieldname))[0].visible&&0===e.deleted)).length)}${"fieldvalue"===n.view?i`<edit-meta audit="${t}" 
              .current="${n}" .dataset="${o}" 
              .onEvent=${this.onEvent} .msg=${this.msg} pageSize=${this.selectorPage}
              ></edit-meta>`:a}`}

          ${null!==n.item.id&&void 0!==n.item.fnote&&!0===l.options.pattern?i`${s("fnote","fnote_view","Comment")}${"fnote"===n.view?i`<edit-note id="editor_note"
                value="${n.item.fnote}" patternId="${d(n.template)}"
                .patterns="${o.pattern}" ?readOnly="${"readonly"===t}"
                .onEvent=${this.onEvent} .msg=${this.msg}
                @blur=${this._onNoteBlur}
              ></edit-note>`:a}`:a}

          ${Object.keys(l.view).filter((e=>"disabled"!==l.view[e].view_audit)).map((e=>i`${s(e,l.view[e].title,l.view[e].icon,o[l.view[e].data].length)}${n.view===e?i`<edit-view
                viewName=${e} .current=${n} .template=${l} 
                .dataset=${o} audit=${t}
                .onEvent=${this.onEvent} .msg=${this.msg} pageSize=${this.paginationPage}
              ></edit-view>`:a}`))}
        </div>`}
    </div>`}});const z=e`
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
`;customElements.define("modal-formula",class extends t{constructor(){super(),this.msg=e=>e,this.formula="",this.formulaValues=[],this.partnumber="",this.description=""}static get properties(){return{formula:{type:String},formulaValues:{type:Array},partnumber:{type:String},description:{type:String}}}static get styles(){return[z]}_onModalEvent(e,t){this.onEvent&&this.onEvent.onModalEvent&&this.onEvent.onModalEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("modal_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}_onFormulaChange(e){this.formula=e}render(){const{formula:e,partnumber:t,formulaValues:a,description:l}=this;return i`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label leftIcon="Magic"
                value="${this.msg("",{id:"label_formula"})}" 
                class="title-cell" ></form-label>
            </div>
            <div class="cell align-right" >
              <span id=${"closeIcon"} class="close-icon" 
                @click="${()=>this._onModalEvent($.CANCEL,{})}">
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
                      value="${this.msg("",{id:"product_partnumber"})}" 
                    ></form-label>
                  </div>
                  <form-input type="${y.TEXT}"
                    label="${t}" .style=${{"font-weight":"bold"}}
                    value="${t}" ?disabled=${!0} ?full=${!0}
                  ></form-input>
                </div>
              </div>
              <div class="row full">
                <div class="cell padding-small" >
                  <form-input type="${y.TEXT}"
                    label="${l}" value="${l}" 
                    ?disabled=${!0} ?full=${!0}
                  ></form-input>
                </div>
              </div>
              <div class="row full">
                <div class="cell padding-small" >
                  <form-select id="sel_formula" label="${this.msg("",{id:"label_formula"})}"
                    .onChange=${e=>this._onFormulaChange(e.value)}
                    .options=${a} .isnull="${!0}" value="${e}" 
                  ></form-select>
                </div>
              </div>
            </div>
          </div>
          <div class="section buttons" >
            <div class="section-row" >
              <div class="cell padding-small half" >
                <form-button id="btn_cancel" icon="Times"
                  @click=${()=>this._onModalEvent($.CANCEL,{})} 
                  ?full="${!0}" label="${this.msg("",{id:"msg_cancel"})}"
                >${this.msg("",{id:"msg_cancel"})}</form-button>
              </div>
              <div class="cell padding-small half" >
                <form-button id="btn_ok" icon="Check"
                  @click=${()=>this._onModalEvent($.OK,{value:parseInt(e,10)})} 
                  ?disabled="${""===e}"
                  type="${m.PRIMARY}" ?full="${!0}" 
                  label="${this.msg("",{id:"msg_ok"})}"
                >${this.msg("",{id:"msg_ok"})}</form-button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>`}});const I=e`
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
`;customElements.define("modal-trans",class extends t{constructor(){super(),this.msg=e=>e,this.baseTranstype="",this.transtype="",this.direction="",this.doctypes=[],this.directions=[],this.refno=!0,this.nettoDiv=!1,this.netto=!0,this.fromDiv=!1,this.from=!1,this.elementCount=0}static get properties(){return{baseTranstype:{type:String},transtype:{type:String},direction:{type:String},doctypes:{type:Array},directions:{type:Array},refno:{type:Boolean},nettoDiv:{type:Boolean},netto:{type:Boolean},fromDiv:{type:Boolean},from:{type:Boolean},elementCount:{type:Number}}}static get styles(){return[I]}_onModalEvent(e,t){this.onEvent&&this.onEvent.onModalEvent&&this.onEvent.onModalEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("modal_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}_onValueChange(e,t){this[e]=t}_setTranstype(e){const{baseTranstype:t,elementCount:i}=this;["invoice","receipt"].includes(e)&&["order","rent","worksheet"].includes(t)?(this.nettoDiv=!0,0===i&&(this.fromDiv=!0)):(this.nettoDiv=!1,this.fromDiv=!1),this.transtype=e}render(){const{transtype:e,direction:t,refno:l,from:n,fromDiv:o,netto:s,nettoDiv:r,doctypes:d,directions:c}=this,p=d.map((e=>({value:e,text:e}))),v=c.map((e=>({value:e,text:e})));return i`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label leftIcon="FileText"
                value="${this.msg("",{id:"msg_create_title"})}" 
                class="title-cell" ></form-label>
            </div>
            <div class="cell align-right" >
              <span id=${"closeIcon"} class="close-icon" 
                @click="${()=>this._onModalEvent($.CANCEL,{})}">
                <form-icon iconKey="Times" ></form-icon>
              </span>
            </div>
          </div>
          <div class="section" >
            <div class="section-row" >
              <div class="cell padding-small" >
                <form-label
                  value="${this.msg("",{id:"msg_create_new"})}" 
                ></form-label>
              </div>
            </div>
            <div class="section-row" >
              <div class="cell padding-small half" >
                <form-select id="transtype" label="transtype" ?full=${!0}
                  .onChange=${e=>this._setTranstype(e.value)}
                  .options=${p} .isnull="${!1}" value="${e}" 
                ></form-select>
              </div>
              <div class="cell padding-small half" >
                <form-select id="direction" label="direction" ?full=${!0}
                  .onChange=${e=>this._onValueChange("direction",e.value)}
                  .options=${v} .isnull="${!1}" value="${t}" 
                ></form-select>
              </div>
            </div>
            <div class="section-row" >
              <div class="cell padding-small" >
                <form-label id="refno"
                  value="${this.msg("",{id:"msg_create_setref"})}"
                  leftIcon="${l?"CheckSquare":"SquareEmpty"}"
                  .style=${{cursor:"pointer"}} .iconStyle=${{cursor:"pointer"}}
                  @click=${()=>this._onValueChange("refno",!l)}
                ></form-label>
              </div>
            </div>
            ${r?i`<div class="section-row" >
              <div class="cell padding-small" >
                <form-label id="netto"
                  value="${this.msg("",{id:"msg_create_deduction"})}"
                  leftIcon="${s?"CheckSquare":"SquareEmpty"}"
                  .style=${{cursor:"pointer"}} .iconStyle=${{cursor:"pointer"}}
                  @click=${()=>this._onValueChange("netto",!s)}
                ></form-label>
              </div>
            </div>`:a}
            ${o?i`<div class="section-row" >
              <div class="cell padding-small" >
                <form-label id="from"
                  value="${this.msg("",{id:"msg_create_delivery"})}"
                  leftIcon="${n?"CheckSquare":"SquareEmpty"}"
                  .style=${{cursor:"pointer"}} .iconStyle=${{cursor:"pointer"}}
                  @click=${()=>this._onValueChange("from",!n)}
                ></form-label>
              </div>
            </div>`:a}
          </div>
          <div class="section buttons" >
            <div class="section-row" >
              <div class="cell padding-small half" >
                <form-button id="btn_cancel" icon="Times"
                  @click=${()=>this._onModalEvent($.CANCEL,{})} 
                  ?full="${!0}" label="${this.msg("",{id:"msg_cancel"})}"
                >${this.msg("",{id:"msg_cancel"})}</form-button>
              </div>
              <div class="cell padding-small half" >
                <form-button id="btn_ok" icon="Check"
                  @click=${()=>this._onModalEvent($.OK,{newTranstype:e,newDirection:t,refno:l,fromInventory:n&&o,nettoQty:s&&r})} 
                  type="${m.PRIMARY}" ?full="${!0}" 
                  label="${this.msg("",{id:"msg_ok"})}"
                >${this.msg("",{id:"msg_ok"})}</form-button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>`}});const P=e`
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
  max-width: 450px;
  min-width: 300px;
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
.padding-label { 
  padding: 0px 0px 16px; 
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
  .padding-label { 
    padding: 0px 0px 8px;
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
`;customElements.define("modal-stock",class extends t{constructor(){super(),this.msg=e=>e,this.partnumber="",this.partname="",this.rows=[],this.selectorPage=5}static get properties(){return{partnumber:{type:String},partname:{type:String},rows:{type:Array},selectorPage:{type:Number}}}static get styles(){return[P]}_onModalEvent(e,t){this.onEvent&&this.onEvent.onModalEvent&&this.onEvent.onModalEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("modal_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}render(){const{partnumber:e,partname:t,rows:a,selectorPage:l}=this,n={warehouse:{fieldtype:"string",label:this.msg("",{id:"delivery_place"})},batch_no:{fieldtype:"string",label:this.msg("",{id:"movement_batchnumber"})},description:{fieldtype:"string",label:this.msg("",{id:"product_description"})},sqty:{fieldtype:"number",label:this.msg("",{id:"shipping_stock"})}};return i`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label leftIcon="Book"
                value="${this.msg("",{id:"shipping_stocks"})}" 
                class="title-cell" ></form-label>
            </div>
            <div class="cell align-right" >
              <span id=${"closeIcon"} class="close-icon" 
                @click="${()=>this._onModalEvent($.CANCEL,{})}">
                <form-icon iconKey="Times" ></form-icon>
              </span>
            </div>
          </div>
          <div class="section" >
            <div class="section-row" >
              <div class="cell padding-label" >
                <div>
                  <form-label value="${e}" ></form-label>
                </div>
                <div>
                  <form-label .style=${{"font-weight":"normal"}}
                    value="${t}" 
                  ></form-label>
                </div>
              </div>
            </div>
            <div class="section-row" >
              <form-table id="selector_result"
                .rows="${a}"
                .fields="${n}"
                pagination="${g.TOP}"
                pageSize="${l}"
                ?tableFilter="${!0}" 
                filterPlaceholder="${this.msg("",{id:"placeholder_filter"})}"
                ?hidePaginatonSize="${!0}"
              ></form-table>
            </div>
          </div>
          <div class="section buttons" >
            <div class="section-row" >
              <div class="cell padding-small" >
                <form-button id="btn_ok" icon="Check"
                  label="${this.msg("",{id:"msg_ok"})}"
                  @click=${()=>this._onModalEvent($.CANCEL,{})} 
                  ?autofocus="${!0}"
                  type="${m.PRIMARY}" ?full="${!0}" 
                >${this.msg("",{id:"msg_ok"})}</form-button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>`}});const A=e`
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
`;customElements.define("modal-shipping",class extends t{constructor(){super(),this.partnumber="",this.description="",this.unit="",this.batch_no="",this.qty=0}static get properties(){return{partnumber:{type:String},description:{type:String},unit:{type:String},batch_no:{type:String},qty:{type:Number}}}static get styles(){return[A]}_onModalEvent(e,t){this.onEvent&&this.onEvent.onModalEvent&&this.onEvent.onModalEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("modal_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}_onValueChange(e,t){this[e]=t}render(){const{partnumber:e,description:t,unit:a,batch_no:l,qty:n}=this;return i`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label leftIcon="Truck"
                value="${this.msg("",{id:"shipping_movement_product"})}" 
                class="title-cell" ></form-label>
            </div>
            <div class="cell align-right" >
              <span id=${"closeIcon"} class="close-icon" 
                @click="${()=>this._onModalEvent($.CANCEL,{})}">
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
                      value="${this.msg("",{id:"product_partnumber"})}" 
                    ></form-label>
                  </div>
                  <form-input type="${y.TEXT}"
                    label="${this.msg("",{id:"product_partnumber"})}" 
                    .style=${{"font-weight":"bold"}}
                    value="${e}" ?disabled=${!0} ?full=${!0}
                  ></form-input>
                </div>
              </div>
              <div class="row full">
                <div class="cell padding-small" >
                  <div>
                    <form-label
                      value="${this.msg("",{id:"product_description"})}" 
                    ></form-label>
                  </div>
                  <form-input type="${y.TEXT}"
                    label="${this.msg("",{id:"product_description"})}" 
                    value="${t}" ?disabled=${!0} ?full=${!0}
                  ></form-input>
                </div>
              </div>
              <div class="row full">
                <div class="cell padding-small" >
                  <div>
                    <form-label
                      value="${this.msg("",{id:"product_unit"})}" 
                    ></form-label>
                  </div>
                  <form-input type="${y.TEXT}"
                    label="${this.msg("",{id:"product_unit"})}" 
                    value="${a}" ?disabled=${!0} ?full=${!0}
                  ></form-input>
                </div>
              </div>
              <div class="row full">
                <div class="cell padding-small" >
                  <div>
                    <form-label
                      value="${this.msg("",{id:"movement_batchnumber"})}" 
                    ></form-label>
                  </div>
                  <form-input id="batch_no" type="${y.TEXT}"
                    label="${this.msg("",{id:"movement_batchnumber"})}"
                    value="${l}" ?autofocus=${!0}
                    .onChange=${e=>this._onValueChange("batch_no",e.value)}
                    ?full=${!0}
                  ></form-input>
                </div>
              </div>
              <div class="row full">
                <div class="cell padding-small" >
                  <div>
                    <form-label
                      value="${this.msg("",{id:"movement_qty"})}" 
                    ></form-label>
                  </div>
                  <form-number id="qty" 
                    label="${this.msg("",{id:"movement_qty"})}"
                    ?integer="${!1}" value="${n}"
                    .onChange=${e=>this._onValueChange("qty",e.value)}
                  ></form-number>
                </div>
              </div>
            </div>
          </div>
          <div class="section buttons" >
            <div class="section-row" >
              <div class="cell padding-small half" >
                <form-button id="btn_cancel" icon="Times"
                  @click=${()=>this._onModalEvent($.CANCEL,{})} 
                  ?full="${!0}" label="${this.msg("",{id:"msg_cancel"})}"
                >${this.msg("",{id:"msg_cancel"})}</form-button>
              </div>
              <div class="cell padding-small half" >
                <form-button id="btn_ok" icon="Check"
                  @click=${()=>this._onModalEvent($.OK,{batch_no:l,qty:n})} 
                  type="${m.PRIMARY}" ?full="${!0}" 
                  label="${this.msg("",{id:"msg_ok"})}"
                >${this.msg("",{id:"msg_ok"})}</form-button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>`}});const N=e`
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
.label-padding {
  padding-bottom: 5px;
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
`;customElements.define("modal-report",class extends t{constructor(){super(),this.title="",this.template="",this.templates=[],this.report_orientation=[],this.report_size=[],this.orient="portrait",this.size="a4",this.copy=1}static get properties(){return{title:{type:String},templates:{type:Array},report_orientation:{type:Array},report_size:{type:Array},template:{type:String},orient:{type:String},size:{type:String},copy:{type:Number}}}static get styles(){return[N]}_onModalEvent(e,t){const{template:i,orient:a,size:l,copy:n,title:o}=this,s={type:t,template:i,orient:a,size:l,copy:n,title:o};this.onEvent&&this.onEvent.onModalEvent&&this.onEvent.onModalEvent({key:e,data:s}),this.dispatchEvent(new CustomEvent("modal_event",{bubbles:!0,composed:!0,detail:{key:e,data:s}}))}_onValueChange(e,t){this[e]=t}render(){const{title:e,template:t,templates:a,orient:l,size:n,copy:o,report_size:s,report_orientation:r}=this,d=r.map((e=>({value:e[0],text:this.msg("",{id:e[1]})}))),c=s.map((e=>({value:e[0],text:e[1]})));return i`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label leftIcon="ChartBar"
                value="${e}" 
                class="title-cell" ></form-label>
            </div>
            <div class="cell align-right" >
              <span id=${"closeIcon"} class="close-icon" 
                @click="${()=>this._onModalEvent($.CANCEL,"")}">
                <form-icon iconKey="Times" ></form-icon>
              </span>
            </div>
          </div>
          <div class="section" >
            <div class="section-row" >
              <div class="row full">
                <div class="cell padding-small" >
                  <div class="label-padding">
                    <form-label
                      value="${this.msg("",{id:"msg_template"})}" 
                    ></form-label>
                  </div>
                  <form-select id="template" 
                    label="${this.msg("",{id:"msg_template"})}"
                    .onChange=${e=>this._onValueChange("template",e.value)}
                    .options=${a} .isnull="${!0}" value="${t}" 
                  ></form-select>
                </div>
              </div>
            </div>
            <div class="section-row" >
              <div class="row full">
                <div class="cell padding-small" >
                  <div class="label-padding" >
                    <form-label
                      value="${this.msg("",{id:"msg_report_prop"})}" 
                    ></form-label>
                  </div>
                  <div class="cell" >
                    <form-select id="orient" 
                      label="${l}"
                      .onChange=${e=>this._onValueChange("orient",e.value)}
                      .options=${d} .isnull="${!1}" value="${l}" 
                    ></form-select>
                  </div>
                  <div class="cell" >
                    <form-select id="size" 
                      label="${n}"
                      .onChange=${e=>this._onValueChange("size",e.value)}
                      .options=${c} .isnull="${!1}" value="${n}" 
                    ></form-select>
                  </div>
                  <div class="cell" >
                    <form-number id="copy" 
                      label="${o}" .style=${{width:"60px"}}
                      ?integer="${!0}" value="${o}"
                      .onChange=${e=>this._onValueChange("copy",e.value)}
                    ></form-number>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="section buttons" >
            <div class="section-row" >
              <div class="cell padding-small half" >
                <form-button id="btn_print"
                  @click=${()=>this._onModalEvent($.OK,"print")}
                  ?disabled="${""===t}"
                  type="${m.PRIMARY}" ?full="${!0}" 
                  label="${this.msg("",{id:"msg_print"})}"
                >${this.msg("",{id:"msg_print"})}</form-button>
              </div>
              <div class="cell padding-small half" >
                <form-button id="btn_pdf"
                  @click=${()=>this._onModalEvent($.OK,"pdf")} 
                  ?disabled="${""===t}"
                  type="${m.PRIMARY}" ?full="${!0}" 
                  label="${this.msg("",{id:"msg_export_pdf"})}"
                >${this.msg("",{id:"msg_export_pdf"})}</form-button>
              </div>
            </div>
            <div class="section-row" >
              <div class="cell padding-small half" >
                <form-button id="btn_xml"
                  @click=${()=>this._onModalEvent($.OK,"xml")}
                  ?disabled="${""===t}"
                  type="${m.PRIMARY}" ?full="${!0}" 
                  label="${this.msg("",{id:"msg_export_xml"})}"
                >${this.msg("",{id:"msg_export_xml"})}</form-button>
              </div>
              <div class="cell padding-small half" >
                <form-button id="btn_printqueue"
                  @click=${()=>this._onModalEvent($.OK,"printqueue")} 
                  ?disabled="${""===t}"
                  type="${m.PRIMARY}" ?full="${!0}" 
                  label="${this.msg("",{id:"msg_printqueue"})}"
                >${this.msg("",{id:"msg_printqueue"})}</form-button>
              </div>
            </div>
          </div>
        </div>
      </div>
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
`;customElements.define("client-edit",class extends t{constructor(){super(),this.msg=e=>e,this.id=Math.random().toString(36).slice(2),this.side=c.AUTO,this.data={},this.auditFilter={},this.newFilter=[],this.forms={},this.paginationPage=10,this.selectorPage=5,this.onEvent={},this.modalFormula=this.modalFormula.bind(this),this.modalReport=this.modalReport.bind(this),this.modalSelector=this.modalSelector.bind(this),this.modalShipping=this.modalShipping.bind(this),this.modalStock=this.modalStock.bind(this),this.modalTrans=this.modalTrans.bind(this)}static get properties(){return{id:{type:String},side:{type:String},data:{type:Object},auditFilter:{type:Object},newFilter:{type:Array},forms:{type:Object},paginationPage:{type:Number},selectorPage:{type:Number},onEvent:{type:Object}}}static get styles(){return[V]}connectedCallback(){super.connectedCallback(),this.onEvent.setModule(this)}modalFormula({formula:e,formulaValues:t,partnumber:a,description:l,onEvent:n}){return i`<modal-formula 
      formula="${e}" partnumber="${a}" description="${l}"
      .formulaValues=${t} 
      .onEvent=${n} .msg=${this.msg}
    ></modal-formula>`}modalReport({title:e,template:t,copy:a,orient:l,size:n,templates:o,report_size:s,report_orientation:r,onEvent:d}){return i`<modal-report
      title="${e}"
      template="${t}"
      copy="${a}"
      orient="${l}"
      size="${n}"
      .templates=${o}
      .report_size=${s}
      .report_orientation=${r}
      .onEvent=${d}
      .msg=${this.msg}
    ></modal-report>`}modalSelector({view:e,columns:t,result:a,filter:l,onEvent:n}){return i`<modal-selector
      ?isModal="${!0}"
      view="${e}"
      .columns=${t}
      .result=${a}
      filter="${l}"
      .onEvent=${n}
      .msg=${this.msg}
    ></modal-selector>`}modalShipping({unit:e,batch_no:t,qty:a,partnumber:l,description:n,onEvent:o}){return i`<modal-shipping
      unit="${e}"
      batch_no="${t}"
      qty="${a}"
      partnumber="${l}"
      description="${n}"
      .onEvent=${o}
      .msg=${this.msg}
    ></modal-shipping>`}modalStock({partnumber:e,partname:t,rows:a,selectorPage:l,onEvent:n}){return i`<modal-stock
      partnumber="${e}"
      partname="${t}"
      selectorPage="${l}"
      .rows="${a}"
      .onEvent=${n}
      .msg=${this.msg}
    ></modal-stock>`}modalTrans({baseTranstype:e,transtype:t,direction:a,doctypes:l,directions:n,refno:o,nettoDiv:s,netto:r,fromDiv:d,from:c,elementCount:p,onEvent:m}){return i`<modal-trans
      baseTranstype="${e}"
      transtype="${t}"
      direction="${a}"
      .doctypes="${l}"
      .directions="${n}"
      ?refno="${o}"
      ?nettoDiv="${s}"
      ?netto="${r}"
      ?fromDiv="${d}"
      ?from="${c}"
      elementCount="${p}"
      .onEvent=${m}
      .msg=${this.msg}
    ></modal-trans>`}render(){const{side:e,data:t,newFilter:l,auditFilter:n}=this;return i`<sidebar-edit
      id="${this.id}" side="${e}" view="${t.side_view}"
      .newFilter="${l}" .auditFilter="${n}"
      .module="${t}" .forms="${this.forms}"
      .onEvent=${this.onEvent} .msg=${this.msg}
    ></sidebar-edit>
      <div class="page">
        ${t.current.item?i`<edit-editor
          caption="${t.caption}" 
          .current=${t.current} .template=${t.template}
          .dataset=${t.dataset} audit="${t.audit}"
          .onEvent=${this.onEvent} .msg=${this.msg}
        ></edit-editor>`:a}
      </div>`}});
