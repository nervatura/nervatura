import { LitElement, html } from 'lit';
 
import '../Form/Button/form-button.js'
import '../Form/Select/form-select.js'
import '../Form/Input/form-input.js'
import '../Form/Icon/form-icon.js'
import '../Form/Label/form-label.js'

import { styles } from './Login.styles.js'
import { LOGIN_PAGE_EVENT, BUTTON_TYPE, INPUT_TYPE, APP_THEME } from '../../config/enums.js'
import { LoginController } from '../../controllers/LoginController.js'

export class Login extends LitElement {
  constructor() {
    super();
    this.version = ""
    this.serverURL = ""
    this.theme = APP_THEME.LIGHT
    this.lang = "en"
    this.locales = []
    this.data = {}
    this.onEvent = new LoginController(this)
  }

  static get properties() {
    return {
      version: { type: String },
      serverURL: { type: String },
      theme: { type: String },
      lang: { type: String },
      locales: { type: Object, attribute: false },
      data: { type: Object, attribute: false },
      onEvent: { type: Object },
    };
  }

  _onPageEvent(key, data){
    if(this.onEvent.onPageEvent){
      this.onEvent.onPageEvent({ key, data })
    }
    this.dispatchEvent(
      new CustomEvent('page_event', {
        bubbles: true, composed: true,
        detail: {
          key, data
        }
      })
    );
  }
  
  render() {
    const { username, password, database, server, code } = this.data
    const locales = Object.keys(this.locales).map(key => ({value: key, text: this.locales[key][key] || key}))
    const loginDisabled = (!username || (String(username).length===0) || !database || (String(database).length===0))
    return html`
      <div class="modal" theme="${this.theme}" >
        <div class="middle">
          <div class="dialog">
            <div class="row title">
              <div class="cell title-cell" >
                <span>${this.msg("Nervatura Report Editor", {id: "title_login"})}</span>
              </div>
              <div class="cell version-cell" >
                <span>${`v${this.version}`}</span>
              </div>
            </div>
            <div class="row full section-small" >
              <div class="row full section-small" >
                <div class="cell label-cell padding-normal mobile" >
                  <form-label value="${this.msg("Username", {id: "login_username"})}" ></form-label>
                </div>
                <div class="cell container mobile" >
                  <form-input id="username" type="${INPUT_TYPE.TEXT}" 
                    label="${this.msg("", {id: "login_username"})}"
                    value="${username}" ?full="${true}"
                    .onChange=${
                      (event) => this._onPageEvent(LOGIN_PAGE_EVENT.CHANGE, { fieldname: "username", value: event.value }) 
                    }
                  ></form-input>
                </div>
              </div>
              <div class="row full" >
                <div class="cell label-cell padding-normal mobile" >
                  <form-label value="${this.msg("Password", {id: "login_password"})}" ></form-label>
                </div>
                <div class="cell container mobile" >
                  <form-input id="password" 
                    type="${INPUT_TYPE.PASSWORD}" 
                    label="${this.msg("", {id: "login_password"})}"
                    value="${password}" ?full="${true}"
                    .onChange=${
                      (event) => this._onPageEvent(LOGIN_PAGE_EVENT.CHANGE, { fieldname: "password", value: event.value }) 
                    }
                    .onEnter=${()=>this._onPageEvent(LOGIN_PAGE_EVENT.LOGIN, this.data)}
                  ></form-input>
                </div>
              </div>
              <div class="row full section-small" >
                <div class="cell label-cell padding-normal mobile" >
                  <form-label value="${this.msg("Database", {id: "login_database"})}" ></form-label>
                </div>
                <div class="cell container mobile" >
                  <form-input id="database" type="${INPUT_TYPE.TEXT}"
                    label="${this.msg("", {id: "login_database"})}"
                    value="${database}" ?full="${true}"
                    .onChange=${
                      (event) => this._onPageEvent(LOGIN_PAGE_EVENT.CHANGE, { fieldname: "database", value: event.value }) 
                    }
                    .onEnter=${()=>this._onPageEvent(LOGIN_PAGE_EVENT.LOGIN, this.data)}
                  ></form-input>
                </div>
              </div>
              <div class="row full section-small" >
                <div class="cell label-cell padding-normal mobile" >
                  <form-label value="${this.msg("Template code", {id: "login_code"})}" ></form-label>
                </div>
                <div class="cell container mobile" >
                  <form-input id="code" type="${INPUT_TYPE.TEXT}"
                    label="${this.msg("", {id: "login_code"})}"
                    value="${code}" ?full="${true}"
                    .onChange=${
                      (event) => this._onPageEvent(LOGIN_PAGE_EVENT.CHANGE, { fieldname: "code", value: event.value }) 
                    }
                    .onEnter=${()=>this._onPageEvent(LOGIN_PAGE_EVENT.LOGIN, this.data)}
                  ></form-input>
                </div>
              </div>
              ${(this.serverURL !== "SERVER") ? html`
                <div class="row full section-small-bottom" >
                   <div class="cell container" >
                     <div class="section-small" >
                       <form-label value="${this.msg("Server URL", {id: "login_server"})}" ></form-label>
                     </div>
                     <form-input id="server" type="${INPUT_TYPE.TEXT}" 
                        value="${server}" ?full="${true}"
                        label="${this.msg("", {id: "login_server"})}"
                        .onChange=${
                          (event) => this._onPageEvent(LOGIN_PAGE_EVENT.CHANGE, { fieldname: "server", value: event.value }) 
                        }
                      ></form-input>
                   </div>
                </div>
              ` : ""}
            </div>
            <div class="row full section buttons" >
              <div class="cell section-small mobile" >
                <div class="cell container" >
                  <form-button id="theme" label="Theme"
                    @click=${()=>this._onPageEvent(LOGIN_PAGE_EVENT.THEME, (this.theme === APP_THEME.DARK) ? APP_THEME.LIGHT : APP_THEME.DARK)} 
                    type="${BUTTON_TYPE.BORDER}" >${
                      (this.theme === APP_THEME.DARK)
                      ? html`<form-icon iconKey="Sun" width=18 height=18 ></form-icon>`
                      : html`<form-icon iconKey="Moon" width=18 height=18 ></form-icon>`}</form-button>
                </div>
                <div class="cell" >
                  <form-select id="lang" label="${this.msg("Login", {id: "login_lang"})}"
                    .onChange=${
                      (value) => this._onPageEvent(LOGIN_PAGE_EVENT.LANG, value.value ) 
                    }
                    .options=${locales} .isnull="${false}" value="${this.lang}" ></form-select>
                </div>
              </div>
              <div class="cell container section-small align-right mobile" >
                <form-button id="login" ?autofocus="${true}"
                  ?disabled="${loginDisabled}" 
                  label="${this.msg("Login", {id: "login_login"})}"
                  @click=${(!loginDisabled) ? ()=>this._onPageEvent(LOGIN_PAGE_EVENT.LOGIN, this.data) : null} 
                  type="${BUTTON_TYPE.PRIMARY}" ?full="${true}" 
                >${this.msg("", {id: "login_login"})}
                </form-button>
              </div>
            </div>
          </div>
        </div>
      </div>
    `;
  }

  static get styles () {
    return [
      styles
    ]
  }
}

