import{a as e,i as t,b as a}from"./module-1JqPnLJw.js";import"./module-tEN-oy1p.js";import"./module-BCutSLl1.js";import{A as l,T as i,L as o,a as s,I as n,B as r}from"./main-DoduxAVE.js";import"./module-NpGsdRHB.js";const d=e`
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
  .modal {
    position: fixed;
    left: 0;
    right: 0;
    top: 0;
    height: 100vh;
    overflow: auto;
    background-color: rgba(25, 25, 25, 0.7);
    padding: 10px 5px;
  }
  .middle {
    margin: 20px 10px 10px;
  }
  .dialog {
    border-radius: 0px;
    border: 0.5px solid rgba(var(--neutral-1), 0.2);
    box-shadow:0 4px 10px 0 rgba(0,0,0, 0.2),0 4px 20px 0 rgba(0,0,0, 0.19);
    background-color: rgba(var(--base-2), 1);
    margin: 0 auto;
    animation: animatezoom 0.6s;
    width: 100%;
    max-width: 400px;
    min-width: 280px;
    top: 0;
    left: 0;
  }
  .title {
    width: 100%;
    border-bottom: 0.5px solid rgba(var(--neutral-1), 0.2);
    background-color: rgb(var(--accent-1));
    color: rgba(var(--accent-1c), 0.85);
    fill: rgba(var(--accent-1c), 0.85);
  }
  .title-cell {
    padding: 8px 16px;
    font-weight: bold;
  }
  .version-cell {
    padding: 8px 16px;
    text-align: right;
    font-size: 13px;
    font-weight: normal;
  }
  .label-cell {
    width: 35%;
  }
  .buttons {
    background-color: rgb(var(--base-1));
    border-top: 0.5px solid rgba(var(--neutral-1), 0.2);
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
  .section { 
    padding-top: 16px!important; 
    padding-bottom: 16px!important; 
  }
  .section-small { 
    padding-top: 8px!important; 
    padding-bottom: 8px!important; 
  }
  .section-small-bottom { 
    padding-bottom: 8px!important; 
  }
  .container { 
    padding: 0.01em 16px; 
  } 
  .container::after, .container::before { 
    content: ""; display: table; clear: both; 
  }
  .align-right { 
    text-align: right; 
  }
  .padding-normal { 
    padding: 8px 16px; 
  }
  @media (max-width:600px){
    .mobile{ 
      display: block; 
      width: 100%; 
    }
    .container { 
      padding: 0px 8px; 
    }
    .padding-normal { 
      padding: 4px 8px; 
    }
  }
  @media only screen and (min-width: 601px){
    .middle {
      margin: 0px;
      position:absolute;
      top:50%;
      left:50%;
      transform:translate(-50%,-50%);
      -ms-transform:translate(-50%,-50%)
    }
    .dialog {
      min-width: 400px;
    }
  }
`;class c{constructor(e){this.host=e,this.onLogin=this.onLogin.bind(this),this.onPageEvent=this.onPageEvent.bind(this),e.addController(this)}async onLogin(){const{data:e,setData:t}=this.host.app.store,{requestData:a,resultError:o,currentModule:s,showToast:n,msg:r}=this.host.app;var d={method:"POST",data:{user_name:e[l.LOGIN].username,password:e[l.LOGIN].password,database:e[l.LOGIN].database}},c=await a("/auth/login",d);if(!c.token)return o(c);var m={token:c.token,content:{fkey:"setTemplate",args:[{type:"_sample"}]}};e[l.LOGIN].code&&(d={method:"GET",token:c.token},(c=await a("/config/"+e[l.LOGIN].code,d)).error?n(i.ERROR,r("",{id:"login_template_code_err"})):m.content={fkey:"setTemplate",args:[{type:"template",report:c}]}),t(l.LOGIN,{data:m}),localStorage.setItem("database",e[l.LOGIN].database),localStorage.setItem("username",e[l.LOGIN].username),localStorage.setItem("server",e[l.LOGIN].server),localStorage.setItem("code",e[l.LOGIN].code),s({data:{module:l.TEMPLATE},content:m.content}),window.history.replaceState(null,null,window.location.pathname)}onPageEvent({key:e,data:t}){const{setData:a}=this.host.app.store;switch(e){case o.CHANGE:a(l.LOGIN,{[t.fieldname]:t.value});break;case o.THEME:case o.LANG:a("current",{[e]:t}),localStorage.setItem([e],t);break;case o.LOGIN:this.onLogin()}}}customElements.define("client-login",class extends t{constructor(){super(),this.version="",this.serverURL="",this.theme=s.LIGHT,this.lang="en",this.locales=[],this.data={},this.onEvent=new c(this)}static get properties(){return{version:{type:String},serverURL:{type:String},theme:{type:String},lang:{type:String},locales:{type:Object,attribute:!1},data:{type:Object,attribute:!1},onEvent:{type:Object}}}_onPageEvent(e,t){this.onEvent.onPageEvent&&this.onEvent.onPageEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("page_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}render(){const{username:e,password:t,database:l,server:i,code:d}=this.data,c=Object.keys(this.locales).map(e=>({value:e,text:this.locales[e][e]||e})),m=!e||0===String(e).length||!l||0===String(l).length;return a`
      <div class="modal" theme="${this.theme}" >
        <div class="middle">
          <div class="dialog">
            <div class="row title">
              <div class="cell title-cell" >
                <span>${this.msg("Nervatura Report Editor",{id:"title_login"})}</span>
              </div>
              <div class="cell version-cell" >
                <span>${`v${this.version}`}</span>
              </div>
            </div>
            <div class="row full section-small" >
              <div class="row full section-small" >
                <div class="cell label-cell padding-normal mobile" >
                  <form-label value="${this.msg("Username",{id:"login_username"})}" ></form-label>
                </div>
                <div class="cell container mobile" >
                  <form-input id="username" type="${n.TEXT}" 
                    label="${this.msg("",{id:"login_username"})}"
                    value="${e}" ?full="${!0}"
                    .onChange=${e=>this._onPageEvent(o.CHANGE,{fieldname:"username",value:e.value})}
                  ></form-input>
                </div>
              </div>
              <div class="row full" >
                <div class="cell label-cell padding-normal mobile" >
                  <form-label value="${this.msg("Password",{id:"login_password"})}" ></form-label>
                </div>
                <div class="cell container mobile" >
                  <form-input id="password" 
                    type="${n.PASSWORD}" 
                    label="${this.msg("",{id:"login_password"})}"
                    value="${t}" ?full="${!0}"
                    .onChange=${e=>this._onPageEvent(o.CHANGE,{fieldname:"password",value:e.value})}
                    .onEnter=${()=>this._onPageEvent(o.LOGIN,this.data)}
                  ></form-input>
                </div>
              </div>
              <div class="row full section-small" >
                <div class="cell label-cell padding-normal mobile" >
                  <form-label value="${this.msg("Database",{id:"login_database"})}" ></form-label>
                </div>
                <div class="cell container mobile" >
                  <form-input id="database" type="${n.TEXT}"
                    label="${this.msg("",{id:"login_database"})}"
                    value="${l}" ?full="${!0}"
                    .onChange=${e=>this._onPageEvent(o.CHANGE,{fieldname:"database",value:e.value})}
                    .onEnter=${()=>this._onPageEvent(o.LOGIN,this.data)}
                  ></form-input>
                </div>
              </div>
              <div class="row full section-small" >
                <div class="cell label-cell padding-normal mobile" >
                  <form-label value="${this.msg("Template code",{id:"login_code"})}" ></form-label>
                </div>
                <div class="cell container mobile" >
                  <form-input id="code" type="${n.TEXT}"
                    label="${this.msg("",{id:"login_code"})}"
                    value="${d}" ?full="${!0}"
                    .onChange=${e=>this._onPageEvent(o.CHANGE,{fieldname:"code",value:e.value})}
                    .onEnter=${()=>this._onPageEvent(o.LOGIN,this.data)}
                  ></form-input>
                </div>
              </div>
              ${"SERVER"!==this.serverURL?a`
                <div class="row full section-small-bottom" >
                   <div class="cell container" >
                     <div class="section-small" >
                       <form-label value="${this.msg("Server URL",{id:"login_server"})}" ></form-label>
                     </div>
                     <form-input id="server" type="${n.TEXT}" 
                        value="${i}" ?full="${!0}"
                        label="${this.msg("",{id:"login_server"})}"
                        .onChange=${e=>this._onPageEvent(o.CHANGE,{fieldname:"server",value:e.value})}
                      ></form-input>
                   </div>
                </div>
              `:""}
            </div>
            <div class="row full section buttons" >
              <div class="cell section-small mobile" >
                <div class="cell container" >
                  <form-button id="theme" label="Theme"
                    @click=${()=>this._onPageEvent(o.THEME,this.theme===s.DARK?s.LIGHT:s.DARK)} 
                    type="${r.BORDER}" >${this.theme===s.DARK?a`<form-icon iconKey="Sun" width=18 height=18 ></form-icon>`:a`<form-icon iconKey="Moon" width=18 height=18 ></form-icon>`}</form-button>
                </div>
                <div class="cell" >
                  <form-select id="lang" label="${this.msg("Login",{id:"login_lang"})}"
                    .onChange=${e=>this._onPageEvent(o.LANG,e.value)}
                    .options=${c} .isnull="${!1}" value="${this.lang}" ></form-select>
                </div>
              </div>
              <div class="cell container section-small align-right mobile" >
                <form-button id="login" ?autofocus="${!0}"
                  ?disabled="${m}" 
                  label="${this.msg("Login",{id:"login_login"})}"
                  @click=${m?null:()=>this._onPageEvent(o.LOGIN,this.data)} 
                  type="${r.PRIMARY}" ?full="${!0}" 
                >${this.msg("",{id:"login_login"})}
                </form-button>
              </div>
            </div>
          </div>
        </div>
      </div>
    `}static get styles(){return[d]}});
