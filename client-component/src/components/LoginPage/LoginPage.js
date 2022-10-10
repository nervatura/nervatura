import { LitElement, html } from 'lit';
// import { msg } from '@lit/localize';
 
import '../Form/NtButton/nt-button.js'
import '../Form/NtSelect/nt-select.js'
import '../Form/NtInput/nt-input.js'
import '../Form/NtIcon/nt-icon.js'
import '../Form/NtLabel/nt-label.js'

import { styles } from './LoginPage.styles.js'

export class LoginPage extends LitElement {
  constructor() {
    super();
    this.msg = (defValue) => defValue
    this.version = ""
    this.serverURL = ""
    this.theme = "light"
    this.lang = "en"
    this.locales = []
    this.data = {}
  }

  static get properties() {
    return {
      version: { type: String },
      serverURL: { type: String },
      theme: { type: String },
      lang: { type: String },
      locales: { type: Object, attribute: false },
      data: { type: Object, attribute: false },
    };
  }

  _onPageEvent(key, data){
    if(this.onEvent && this.onEvent.onPageEvent){
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
    const { username, password, database, server } = this.data
    const locales = Object.keys(this.locales).map(key => ({value: key, text: this.locales[key][key] || key}))
    const loginDisabled = (!username || (String(username).length===0) || !database || (String(database).length===0))
    return html`
      <div class="modal" theme="${this.theme}" >
        <div class="middle">
          <div class="dialog">
            <div class="row title">
              <div class="cell title-cell" >
                <span>${this.msg("Nervatura Client", {id: "title_login"})}</span>
              </div>
              <div class="cell version-cell" >
                <span>${`v${this.version}`}</span>
              </div>
            </div>
            <div class="row full section-small" >
              <div class="row full section-small" >
                <div class="cell label-cell padding-normal mobile" >
                  <nt-label value="${this.msg("Username", {id: "login_username"})}" ></nt-label>
                </div>
                <div class="cell container mobile" >
                  <nt-input id="username" type="text" 
                    label="${this.msg("Username", {id: "login_username"})}"
                    value="${username}" ?full="${true}"
                    .onChange=${
                      (event) => this._onPageEvent("change", { fieldname: "username", value: event.value }) 
                    }
                  ></nt-input>
                </div>
              </div>
              <div class="row full" >
                <div class="cell label-cell padding-normal mobile" >
                  <nt-label value="${this.msg("Password", {id: "login_password"})}" ></nt-label>
                </div>
                <div class="cell container mobile" >
                  <nt-input id="password" type="password" 
                    label="${this.msg("Password", {id: "login_password"})}"
                    value="${password}" ?full="${true}"
                    .onChange=${
                      (event) => this._onPageEvent("change", { fieldname: "password", value: event.value }) 
                    }
                    .onEnter=${()=>this._onPageEvent("login", this.data)}
                  ></nt-input>
                </div>
              </div>
              <div class="row full section-small" >
                <div class="cell label-cell padding-normal mobile" >
                  <nt-label value="${this.msg("Database", {id: "login_database"})}" ></nt-label>
                </div>
                <div class="cell container mobile" >
                  <nt-input id="database" type="text" 
                    label="${this.msg("Database", {id: "login_database"})}"
                    value="${database}" ?full="${true}"
                    .onChange=${
                      (event) => this._onPageEvent("change", { fieldname: "database", value: event.value }) 
                    }
                    .onEnter=${()=>this._onPageEvent("login", this.data)}
                  ></nt-input>
                </div>
              </div>
              ${(this.serverURL !== "SERVER") ? html`
                <div class="row full section-small-bottom" >
                   <div class="cell container" >
                     <div class="section-small" >
                       <nt-label value="${this.msg("Server URL", {id: "login_server"})}" ></nt-label>
                     </div>
                     <nt-input id="server" type="text" 
                        value="${server}" ?full="${true}"
                        label="${this.msg("Server URL", {id: "login_server"})}"
                        .onChange=${
                          (event) => this._onPageEvent("change", { fieldname: "server", value: event.value }) 
                        }
                      ></nt-input>
                   </div>
                </div>
              ` : ""}
            </div>
            <div class="row full section buttons" >
              <div class="cell section-small mobile" >
                <div class="cell container" >
                  <nt-button id="theme" label="Theme"
                    @click=${()=>this._onPageEvent("theme", (this.theme === "dark") ? "light" : "dark")} 
                    type="border" >${
                      (this.theme === "dark")
                      ? html`<nt-icon iconKey="Sun" width=18 height=18 ></nt-icon>`
                      : html`<nt-icon iconKey="Moon" width=18 height=18 ></nt-icon>`}</nt-button>
                </div>
                <div class="cell" >
                  <nt-select id="lang" label="Lang"
                    .onChange=${
                      (value) => this._onPageEvent("lang", value.value ) 
                    }
                    .options=${locales} .isnull="${false}" value="${this.lang}" ></nt-select>
                </div>
              </div>
              <div class="cell container section-small align-right mobile" >
                <nt-button id="login" ?autofocus="${true}"
                  ?disabled="${loginDisabled}" label="${this.msg("Login", {id: "login_login"})}"
                  @click=${(!loginDisabled) ? ()=>this._onPageEvent("login", this.data) : null} 
                  type="primary" ?full="${true}" >
                ${this.msg("Login", {id: "login_login"})}
                </nt-button>
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

