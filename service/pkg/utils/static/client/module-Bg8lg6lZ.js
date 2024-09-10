import{i as e,h as t,k as i}from"./module-Cq_Zev4P.js";import"./module-S4DWhBjj.js";import"./module-CbpjHWOZ.js";import{A as l,I as a,L as o,B as n}from"./main-zdcxVATk.js";import{LoginController as s}from"./module-BEegrjbo.js";import"./module-C1fQGCtD.js";const r=e`
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
`;customElements.define("client-login",class extends t{constructor(){super(),this.version="",this.serverURL="",this.theme=l.LIGHT,this.lang="en",this.locales=[],this.data={},this.onEvent=new s(this)}static get properties(){return{version:{type:String},serverURL:{type:String},theme:{type:String},lang:{type:String},locales:{type:Object,attribute:!1},data:{type:Object,attribute:!1},onEvent:{type:Object}}}_onPageEvent(e,t){this.onEvent.onPageEvent&&this.onEvent.onPageEvent({key:e,data:t}),this.dispatchEvent(new CustomEvent("page_event",{bubbles:!0,composed:!0,detail:{key:e,data:t}}))}render(){const{username:e,password:t,database:s,server:r}=this.data,d=Object.keys(this.locales).map((e=>({value:e,text:this.locales[e][e]||e}))),m=!e||0===String(e).length||!s||0===String(s).length;return i`
      <div class="modal" theme="${this.theme}" >
        <div class="middle">
          <div class="dialog">
            <div class="row title">
              <div class="cell title-cell" >
                <span>${this.msg("Nervatura Client",{id:"title_login"})}</span>
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
                  <form-input id="username" type="${a.TEXT}" 
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
                    type="${a.PASSWORD}" 
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
                  <form-input id="database" type="${a.TEXT}"
                    label="${this.msg("",{id:"login_database"})}"
                    value="${s}" ?full="${!0}"
                    .onChange=${e=>this._onPageEvent(o.CHANGE,{fieldname:"database",value:e.value})}
                    .onEnter=${()=>this._onPageEvent(o.LOGIN,this.data)}
                  ></form-input>
                </div>
              </div>
              ${"SERVER"!==this.serverURL?i`
                <div class="row full section-small-bottom" >
                   <div class="cell container" >
                     <div class="section-small" >
                       <form-label value="${this.msg("Server URL",{id:"login_server"})}" ></form-label>
                     </div>
                     <form-input id="server" type="${a.TEXT}" 
                        value="${r}" ?full="${!0}"
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
                    @click=${()=>this._onPageEvent(o.THEME,this.theme===l.DARK?l.LIGHT:l.DARK)} 
                    type="${n.BORDER}" >${this.theme===l.DARK?i`<form-icon iconKey="Sun" width=18 height=18 ></form-icon>`:i`<form-icon iconKey="Moon" width=18 height=18 ></form-icon>`}</form-button>
                </div>
                <div class="cell" >
                  <form-select id="lang" label="${this.msg("Login",{id:"login_lang"})}"
                    .onChange=${e=>this._onPageEvent(o.LANG,e.value)}
                    .options=${d} .isnull="${!1}" value="${this.lang}" ></form-select>
                </div>
              </div>
              <div class="cell container section-small align-right mobile" >
                <form-button id="login" ?autofocus="${!0}"
                  ?disabled="${m}" 
                  label="${this.msg("Login",{id:"login_login"})}"
                  @click=${m?null:()=>this._onPageEvent(o.LOGIN,this.data)} 
                  type="${n.PRIMARY}" ?full="${!0}" 
                >${this.msg("",{id:"login_login"})}
                </form-button>
              </div>
            </div>
          </div>
        </div>
      </div>
    `}static get styles(){return[r]}});
