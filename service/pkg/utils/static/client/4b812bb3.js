import{i as e,s as t,y as i,b as a}from"./4e7ea0c6.js";import"./5be76710.js";import"./81d721ef.js";import"./2788f2e7.js";import{I as l,M as o,B as s,h as n,P as r,b as d,i as c,d as m,E as h,S as u}from"./8ccbd206.js";import"./9479f488.js";import"./95ec07a4.js";const p=e`
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
`;customElements.define("modal-inputbox",class extends t{constructor(){super(),this.title="",this.message="",this.infoText=void 0,this.value="",this.labelCancel="Cancel",this.labelOK="OK",this.defaultOK=!1,this.showValue=!1,this.values={}}static get properties(){return{title:{type:String},message:{type:String},infoText:{type:String},value:{type:String,reflect:!0},labelOK:{type:String},labelCancel:{type:String},defaultOK:{type:Boolean},showValue:{type:Boolean},values:{type:Object}}}static get styles(){return[p]}_onModalEvent(e){const t={value:(this.renderRoot.querySelector("#input_value")||{}).value,values:this.values};this.onEvent&&this.onEvent.onModalEvent&&this.onEvent.onModalEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("modal_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}render(){return i`<div class="modal">
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
              ${this.infoText?i`<div class="info">${this.infoText}</div>`:a}
              ${this.showValue?i`<div class="info">
                  <form-input id="input_value" type="${l.TEXT}" label="${this.title}"
                    value="${this.value}" ?full="${!0}"
                    .onEnter=${()=>this._onModalEvent(o.OK)}
                  ></form-input>
                </div>`:a}
            </div>
          </div>
          <div class="section buttons" >
            <div class="section-row" >
              <div class="cell padding-small half" >
                <form-button id="btn_cancel" icon="Times"
                  @click=${()=>this._onModalEvent(o.CANCEL)} 
                  ?full="${!0}" label="${this.labelCancel}"
                >${this.labelCancel}</form-button>
              </div>
              <div class="cell padding-small half" >
                <form-button id="btn_ok" icon="Check"
                  @click=${()=>this._onModalEvent(o.OK)} 
                  ?autofocus="${!this.showValue&&this.defaultOK}"
                  type="${s.PRIMARY}" ?full="${!0}" label="${this.labelOK}"
                >${this.labelOK}</form-button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>`}});const b=e`
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
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  box-shadow:0 4px 10px 0 rgba(0,0,0, 0.2),0 4px 20px 0 rgba(0,0,0, 0.19);
  background-color: rgba(var(--base-2), 1);
  margin: 0 auto;
  animation: animatezoom 0.6s;
  width: 100%;
  max-width: 600px;
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
  padding: 0 16px 8px;
}
.half { 
  width:100% 
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
  .half { 
    width:49.99999% 
  }
}
`;customElements.define("modal-bookmark",class extends t{constructor(){super(),this.bookmark={history:null,bookmark:[]},this.tabView=n.BOOKMARK,this.pageSize=5,this.onEvent={}}static get properties(){return{bookmark:{type:Object},tabView:{type:String},pageSize:{type:Number},onEvent:{type:Object}}}static get styles(){return[b]}_onModalEvent(e,t){this.onEvent.onModalEvent&&this.onEvent.onModalEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("modal_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}_onTabView(e){this.tabView=e}setBookmark(){return this.bookmark.bookmark.map((e=>{const t=JSON.parse(e.cfvalue),i={bookmark_id:e.id,id:t.id,cfgroup:e.cfgroup,ntype:t.ntype,transtype:"trans"===t.ntype?t.transtype:null,vkey:t.vkey,view:t.view,filters:t.filters,columns:t.columns,lslabel:e.cfname,lsvalue:new Intl.DateTimeFormat("default",{year:"numeric",month:"2-digit",day:"2-digit"}).format(new Date(t.date))};return"editor"===e.cfgroup&&("trans"===t.ntype?i.lsvalue+=` | ${this.msg(`title_${t.transtype}`,{id:`title_${t.transtype}`})} | ${t.info}`:i.lsvalue+=` | ${this.msg(`title_${t.ntype}`,{id:`title_${t.ntype}`})} | ${t.info}`),"browser"===e.cfgroup&&(i.lsvalue+=` | ${this.msg(`browser_${t.vkey}`,{id:`browser_${t.vkey}`})}`),i}))}setHistory(){if(this.bookmark.history&&this.bookmark.history.cfvalue){return JSON.parse(this.bookmark.history.cfvalue).map((e=>({id:e.id,lslabel:e.title,type:e.type,lsvalue:`${new Intl.DateTimeFormat("default",{year:"numeric",month:"2-digit",day:"2-digit",hour:"2-digit",minute:"2-digit",hour12:!1}).format(new Date(e.datetime))} | ${this.msg(`label_${e.type}`,{id:`label_${e.type}`})}`,ntype:e.ntype,transtype:e.transtype})))}return[]}connectedCallback(){super.connectedCallback(),this.bookmarkList=this.setBookmark(),this.historyList=this.setHistory()}render(){return i`<div class="modal">
      <div class="dialog">
        <div class="panel">
          <div class="panel-title">
            <div class="cell" >
              <form-label 
                value="${this.msg("Nervatura Bookmark",{id:"title_bookmark"})}" 
                class="title-cell" leftIcon="Star" >
              </form-label>
            </div>
            <div class="cell align-right" >
              <span id=${"closeIcon"} class="close-icon" 
                @click="${()=>this._onModalEvent(o.CANCEL,{})}">
                <form-icon iconKey="Times" ></form-icon>
              </span>
            </div>
          </div>
          <div class="section" >
            <div class="section-row" >
              <div class="cell half" >
                <form-button id="btn_bookmark"
                  .style="${{"border-radius":0}}" icon="Star"
                  label="${this.msg("",{id:"title_bookmark_list"})}"
                  @click=${()=>this._onTabView(n.BOOKMARK)} 
                  type="${this.tabView===n.BOOKMARK?s.PRIMARY:""}"
                  ?full="${!0}" ?selected="${n.BOOKMARK===this.tabView}" >
                  ${this.msg("",{id:"title_bookmark_list"})}</form-button>
              </div>
              <div class="cell half" >
                <form-button id="btn_history"
                  .style="${{"border-radius":0}}" icon="History"
                  label="${this.msg("",{id:"title_history"})}"
                  @click=${()=>this._onTabView(n.HISTORY)} 
                  type="${this.tabView===n.HISTORY?s.PRIMARY:""}" 
                  ?full="${!0}" ?selected="${n.HISTORY===this.tabView}" >
                  ${this.msg("",{id:"title_history"})}</form-button>
              </div>
            </div>
            <div class="section-row" >
              <form-list id="bookmark_list"
                .rows="${"bookmark"===this.tabView?this.bookmarkList:this.historyList}"
                pagination="${r.TOP}"
                pageSize="${this.pageSize}"
                ?listFilter="${!0}"
                ?hidePaginatonSize="${!0}"
                filterPlaceholder="${this.msg("Filter",{id:"placeholder_filter"})}"
                editIcon="${"bookmark"===this.tabView?"Star":"History"}"
                .onEdit=${e=>this._onModalEvent(o.SELECTED,{view:this.tabView,row:e})}
                .onDelete=${"bookmark"===this.tabView?e=>this._onModalEvent(o.DELETE,{bookmark_id:e.bookmark_id}):null}
              ></form-list>
            </div>
          </div>
        </div>
      </div>
    </div>`}});class v{constructor(e){this.host=e,this.deleteBookmark=this.deleteBookmark.bind(this),this.onMenuEvent=this.onMenuEvent.bind(this),this.onModalEvent=this.onModalEvent.bind(this),e.addController(this)}async deleteBookmark({bookmark_id:e}){const{inputBox:t}=this.host.app.host,{data:i,setData:a}=this.host.app.store,{requestData:l,resultError:s,loadBookmark:n,msg:r}=this.host.app,{modalBookmark:c}=this.host,m=i[d.LOGIN],h=t({title:r("",{id:"msg_warning"}),message:r("",{id:"msg_delete_text"}),infoText:r("",{id:"msg_delete_info"}),onEvent:{onModalEvent:async({key:t})=>{if(a("current",{modalForm:null}),t===o.CANCEL)return a("current",{modalForm:c()});const i=await l("/ui_userconfig",{method:"DELETE",query:{id:e}});if(i&&i.error)return s(i);const r=await n({user_id:m.data.employee.id});return a("current",{modalForm:c(r)})}}});a("current",{modalForm:h})}onMenuEvent({key:e,data:t}){const{setData:i}=this.host.app.store,{current:a,setting:l}=this.host.app.store.data,{signOut:o,showHelp:s,currentModule:n}=this.host.app,{modalBookmark:r}=this.host;switch(e){case c.SIDEBAR:i("current",{side:a.side===m.SHOW?m.HIDE:m.SHOW});break;case c.MODULE:switch(t.value){case d.LOGIN:o();break;case d.HELP:s("");break;case d.BOOKMARK:i("current",{modalForm:r()});break;default:const e={module:t.value,menu:"",side:m.HIDE};let a=null;t.value!==d.SETTING||l.group_key||(i(d.SETTING,{group_key:"group_admin"}),a={fkey:"checkSetting",args:[{type:"setting"},u.LOAD_SETTING]}),n({data:{...e},content:a})}break;case c.SCROLL:window.scrollTo(0,0)}}onModalEvent({key:e,data:t}){const{currentModule:i}=this.host.app,{setData:a}=this.host.app.store,{search:l}=this.host.app.store.data;switch(e){case o.CANCEL:a("current",{modalForm:null});break;case o.SELECTED:if("bookmark"===t.view&&"browser"===t.row.cfgroup){const e={...l,filters:{...l.filters,[t.row.view]:t.row.filters},columns:{...l.columns,[t.row.view]:t.row.columns}};i({data:{module:d.SEARCH,modalForm:null},content:{fkey:"showBrowser",args:[t.row.vkey,t.row.view,e]}})}else i({data:{module:d.EDIT,modalForm:null},content:{fkey:"checkEditor",args:[{ntype:t.row.ntype,ttype:t.row.transtype,id:t.row.id},h.LOAD_EDITOR]}});break;case o.DELETE:this.deleteBookmark(t)}}}const g=e`
:host {
  font-family: var(--font-family);
  font-size: var(--font-size);
}
div {
  box-sizing: border-box;
}
.shadow {
  box-shadow: 0 2px 5px 0 rgba(0,0,0,0.16),0 2px 10px 0 rgba(0,0,0,0.12);
}
.menubar {
  display: table;
  width: 100%;
  height: var(--menu-top-height);
  padding: 0px 8px;
  background-color: rgb(var(--accent-1));
  font-size: 14px;
  overflow:hidden;
  -webkit-touch-callout: none; -webkit-user-select: none; -khtml-user-select: none; 
  -moz-user-select: none; -ms-user-select: none; user-select: none;
}
.cell {
  display: table-cell;
  vertical-align: middle;
}
.menuitem {
  font-size: 17px;
  text-align: center;
  padding: 8px 6px;
  font-weight: bold;
  vertical-align: middle;
  width: auto;
  display: table-cell;
  cursor: pointer;
  -webkit-touch-callout:none;
  -webkit-user-select:none;
  -khtml-user-select:none;
  -moz-user-select:none;
  -ms-user-select:none;
  user-select:none;
}
.exit:hover {
  color: rgb(var(--functional-red))!important;
  fill: rgb(var(--functional-red))!important;
}
.selected {
  color: rgb(var(--functional-yellow))!important;
  fill: rgb(var(--functional-yellow))!important;
}
.menu-label {
  color: rgba(var(--accent-1c), 0.85);
  fill: rgba(var(--accent-1c), 0.85);
}
.menu-label:hover {
  color: rgb(var(--functional-yellow));
  fill: rgb(var(--functional-yellow));
}
.right { 
  float: right; 
}
.container { 
  padding: 0px 4px; 
}
@media (min-width:769px){
  .sidebar{
    display: none;
  }
}
@media (max-width:600px){
  .hide-small { 
    display: none!important; 
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
}
`;customElements.define("client-menubar",class extends t{constructor(){super(),this.scrollTop=!1,this.side=m.AUTO,this.module=d.SEARCH,this.bookmark={history:null,bookmark:[]},this.selectorPage=5,this.onEvent=new v(this),this.modalBookmark=this.modalBookmark.bind(this)}static get properties(){return{side:{type:String},module:{type:String},scrollTop:{type:Boolean},bookmark:{type:Object},selectorPage:{type:Number},onEvent:{type:Object}}}static get styles(){return[g]}_onMenuEvent(e,t){this.onEvent.onMenuEvent&&this.onEvent.onMenuEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("menu_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}modalBookmark(e){return i`<modal-bookmark
      .bookmark="${e||this.bookmark}"
      tabView="bookmark"
      pageSize=${this.selectorPage}
      .onEvent=${this.onEvent}
      .msg=${this.msg}
    ></modal-bookmark>`}selected(e){return e===this.module?"selected":""}render(){return i`<div class="menubar ${this.scrollTop?"shadow":""}" >
      <div class="cell">
        <div id="mnu_sidebar"
          class="menuitem sidebar" 
          @click=${()=>this._onMenuEvent(c.SIDEBAR)}>
          ${this.side===m.SHOW?i`<form-label 
                value="${this.msg("Hide",{id:"menu_hide"})}" class="selected exit"
                leftIcon="Close" ></form-label>`:i`<form-label class="menu-label"
                value="${this.msg("Menu",{id:"menu_side"})}"
                leftIcon="Bars" .iconStyle="${{width:"24px",height:"24px"}}" ></form-label>`}
        </div>
        <div id="mnu_search_large" 
          class="hide-small hide-medium menuitem"
          @click=${()=>this._onMenuEvent(c.MODULE,{value:d.SEARCH})}>
          <form-label class="menu-label ${this.selected(d.SEARCH)}"
            value="${this.msg("Search",{id:"menu_search"})}"
            leftIcon="Search" ></form-label>
        </div>
        <div id="mnu_edit_large" 
          class="hide-small hide-medium menuitem"
          @click=${()=>this._onMenuEvent(c.MODULE,{value:d.EDIT})}>
          <form-label class="menu-label ${this.selected(d.EDIT)}"
            value="${this.msg("Edit",{id:"menu_edit"})}"
            leftIcon="Edit" ></form-label>
        </div>
        <div id="mnu_setting_large" 
          class="hide-small hide-medium menuitem"
          @click=${()=>this._onMenuEvent(c.MODULE,{value:d.SETTING})}>
          <form-label class="menu-label ${this.selected(d.SETTING)}"
            value="${this.msg("Setting",{id:"menu_setting"})}"
            leftIcon="Cog" ></form-label>
        </div>
        <div id="mnu_bookmark_large" 
          class="hide-small hide-medium menuitem"
          @click=${()=>this._onMenuEvent(c.MODULE,{value:d.BOOKMARK})}>
          <form-label class="menu-label ${this.selected(d.BOOKMARK)}"
            value="${this.msg("Bookmark",{id:"menu_bookmark"})}"
            leftIcon="Star" ></form-label>
        </div>
        <div id="mnu_help_large" 
          class="hide-small hide-medium menuitem"
          @click=${()=>this._onMenuEvent(c.MODULE,{value:d.HELP})}>
          <form-label class="menu-label ${this.selected(d.HELP)}"
            value="${this.msg("Help",{id:"menu_help"})}"
            leftIcon="QuestionCircle" ></form-label>
        </div>
        <div id="mnu_logout_large" 
          class="hide-small hide-medium menuitem"
          @click=${()=>this._onMenuEvent(c.MODULE,{value:d.LOGIN})}>
          <form-label class="menu-label exit"
            value="${this.msg("Logout",{id:"menu_logout"})}"
            leftIcon="Exit" ></form-label>
        </div>
        ${this.scrollTop?i`<div id="mnu_scroll" class="menuitem" 
            @click=${()=>this._onMenuEvent(c.SIDEBAR)}>
            <span class="menu-label" ><form-icon iconKey="HandUp" ></form-icon></span>
          </div>`:a}
      </div>
      <div class="cell container">
        <div id="mnu_help_medium" 
          class="right hide-large menuitem"
          @click=${()=>this._onMenuEvent(c.MODULE,{value:d.HELP})}>
          <span class="hide-small"><form-label class="menu-label ${this.selected(d.HELP)}"
            value="${this.msg("Help",{id:"menu_help"})}"
            leftIcon="QuestionCircle" ></form-label></span>
          <span class="menu-label hide-medium ${this.selected(d.HELP)}" ><form-icon iconKey="QuestionCircle" ></form-icon></span>
        </div>
        <div id="mnu_bookmark_medium" 
          class="right hide-large menuitem"
          @click=${()=>this._onMenuEvent(c.MODULE,{value:d.BOOKMARK})}>
          <span class="hide-small"><form-label class="menu-label ${this.selected(d.BOOKMARK)}"
            value="${this.msg("Bookmark",{id:"menu_bookmark"})}"
            leftIcon="Star" ></form-label></span>
          <span class="menu-label hide-medium ${this.selected(d.BOOKMARK)}" ><form-icon iconKey="Star" ></form-icon></span>
        </div>
        <div id="mnu_setting_medium" 
          class="right hide-large menuitem"
          @click=${()=>this._onMenuEvent(c.MODULE,{value:d.SETTING})}>
          <span class="hide-small"><form-label class="menu-label ${this.selected(d.SETTING)}"
            value="${this.msg("Setting",{id:"menu_setting"})}"
            leftIcon="Cog" ></form-label></span>
          <span class="menu-label hide-medium ${this.selected(d.SETTING)}" ><form-icon iconKey="Cog" ></form-icon></span>
        </div>
        <div id="mnu_edit_medium" 
          class="right hide-large menuitem"
          @click=${()=>this._onMenuEvent(c.MODULE,{value:d.EDIT})}>
          <span class="hide-small"><form-label class="menu-label ${this.selected(d.EDIT)}"
            value="${this.msg("Edit",{id:"menu_edit"})}"
            leftIcon="Edit" ></form-label></span>
          <span class="menu-label hide-medium ${this.selected(d.EDIT)}" ><form-icon iconKey="Edit" ></form-icon></span>
        </div>
        <div id="mnu_search_medium" 
          class="right hide-large menuitem"
          @click=${()=>this._onMenuEvent(c.MODULE,{value:d.SEARCH})}>
          <span class="hide-small"><form-label class="menu-label ${this.selected(d.SEARCH)}"
            value="${this.msg("Search",{id:"menu_search"})}"
            leftIcon="Search" ></form-label></span>
          <span class="menu-label hide-medium ${this.selected(d.SEARCH)}" ><form-icon iconKey="Search" ></form-icon></span>
        </div>
      </div>
      <div id="mnu_logout_medium" class="hide-large menuitem" style="width: 24px;"
        @click=${()=>this._onMenuEvent(c.MODULE,{value:d.LOGIN})}>
        <span class="menu-label exit"><form-icon iconKey="Exit" width=24 height=24 ></form-icon></span>
      </div>
    </div>`}});
