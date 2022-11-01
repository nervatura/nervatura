import { LitElement, html, nothing } from 'lit';
import { ifDefined } from 'lit/directives/if-defined.js';

import variablesCSS from '../../config/variables.js';
import { APP_MODULE, TOAST_TYPE } from '../../config/enums.js'
import { styles } from './NervaturaClient.styles.js'

import { StateController } from '../../controllers/StateController.js'
import { AppController } from '../../controllers/AppController.js'
import { LoginController } from '../../controllers/LoginController.js'
import { MenuController } from '../../controllers/MenuController.js'
import { SearchController } from '../../controllers/SearchController.js'
import { Queries } from '../../controllers/Queries.js'
import { Quick } from '../../controllers/Quick.js'

import { store as storeConfig } from '../../config/app.js'

import '../Form/Spinner/form-spinner.js'
import '../Form/Toast/form-toast.js'
import '../Login/client-login.js'
import '../MenuBar/client-menubar.js'
import '../Search/client-search.js'

const InputBox = ({ 
  title, message, infoText, value, defaultOK, showValue, labelCancel, labelOK, onEvent 
}) => html`<modal-inputbox
    title="${ifDefined(title)}"
    message="${ifDefined(message)}"
    infoText="${ifDefined(infoText)}"
    value="${ifDefined(value)}"
    labelCancel="${ifDefined(labelCancel)}"
    labelOK="${ifDefined(labelOK)}"
    ?defaultOK="${defaultOK||false}"
    ?showValue="${showValue||false}"
    .onEvent=${onEvent}
  ></modal-inputbox>`

export class NervaturaClient extends LitElement {

  static get properties() {
    return {}
  }

  static get styles() {
    return [
      variablesCSS,
      styles
    ];
  }

  constructor() {
    super();
    this.state = new StateController(this, storeConfig)
    this.inputBox = InputBox
    this.queries= Queries({ getText: (key) => this.msg(key,{ id: key }) })
    this.quick= {...Quick}
  }

  connectedCallback() {
    super.connectedCallback();
    // window.addEveListener("scroll", this._onScroll.bind(this), {passive: true});
    // window.addEveListener('resize', this._onResize.bind(this));
    this._loadConfig()
  }

  disconnectedCallback() {
    // window.removeEveListener("scroll", this._onScroll.bind(this));
    // window.removeEveListener('resize', this._onResize.bind(this));
    super.disconnectedCallback();
  }
  
  _onResize() {
    const { current } = this.state.data
    if((current.clientHeight !== window.innerHeight) || 
      (current.clientWidth !== window.innerWidth)){
        this.setData("current", {
          clientHeight: window.innerHeight, clientWidth: window.innerWidth
        })
    }
  }

  _onScroll() {
    const { current } = this.state.data
    const scrollTop = ((document.body.scrollTop > 100) || (document.documentElement.scrollTop > 100))
    if(current.scrollTop !== scrollTop){
      this.setData("current", {
        scrollTop
      })
    }
  }

  getSetting(key) {
    const { ui } = this.state.data
    switch (key) {    
      case "ui":
        const values = {...ui}
        Object.keys(values).forEach(ikey => {
          if(localStorage.getItem(ikey)){
            values[ikey] = localStorage.getItem(ikey)
          }
        });
        return values
  
      default:
        return localStorage.getItem(key) || ui[key] || "";
    }
  }

  async _loadConfig(){
    const { session } = this.state.data
    const getPath = (location) => {
      const getParams = (prmString) => {
        const params = {}
        prmString.split('&').forEach(prm => {
          const index = String(prm).indexOf("=")
          const fname = String(prm).substring(0,(index >0)?index:String(prm).length)
          const value = ((index > -1) && (index < String(prm).length)) ? String(prm).substring(index+1) : ""
          params[fname] = value
        });
        return params
      }
      if(location.hash){
        return ["hash", getParams(location.hash.substring(1))]
      }
      if(location.search){
        return ["search", getParams(location.search.substring(1))]
      }
      const path = location.pathname.substring(1).split("/")
      return [path[0], path.slice(1)]
    }

    let config = {...session}
    try {
      const app = new AppController(this, this.getStore())
      const result = await app.request(`${this.state.data.login.server}/config`, {
        method: "GET",
        headers: { "Content-Type": "application/json" }
      })
      if(result.locales && (typeof result.locales === "object")){
        config = {
          ...config,
          locales: {
            ...config.locales,
            ...result.locales
          }
        }
      }
      this.setData("session", config )
      if(localStorage.getItem("lang") && config.locales[localStorage.getItem("lang")] 
        && (localStorage.getItem("lang") !== this.state.data.current.lang)){
          this.setData("current", {
            lang: localStorage.getItem("lang")
          })
        }
      const [ current, params ] = getPath(window.location)
      if(current === "hash"){ 
        const login = new LoginController(this, app)
        if (params.access_token){
          login.tokenValidation(params)
        }
        if(params.code){
          login.setCodeToken(params)
        }
      }
    } catch (error) {
      this.setData("error", error )
      this.showToast(TOAST_TYPE.ERROR, error.message)
    }
  }

  getStore() {
    return {
      data: this.state.data,
      setData: (...args) => this.setData(...args),
      showToast: (...args) => this.showToast(...args),
      msg: (...args) => this.msg(...args),
      getSetting: (...args) => this.getSetting(...args),
    }
  }

  setData(key, value, update) {
    this.state.data = { key, value, update }
  }

  showToast(type, value, toastTimeout) {
    const { current } = this.state.data
    const timeout = (typeof(toastTimeout) !== "undefined") ? toastTimeout : this.getSetting("toastTimeout")
    if(current.toast){
      current.toast.show({
        type, value, timeout
      })
    }
  }

  msg(defaultValue, props) {
    let value = defaultValue
    const {locales} = this.state.data.session
    const {lang} = this.state.data.current
    if(locales[lang] && locales[lang][props.id]){
      value = locales[lang][props.id]
    } else if((lang !== "en") && locales.en[props.id]) {
      value = locales.en[props.id]
    }
    return value
  }

  protector(){
    const { data } = this.state
    const { current, session } = this.state.data
    const app = new AppController(this, this.getStore())
    if(data[APP_MODULE.LOGIN].data){
      return html`
      <div class="client-menubar">
        <client-menubar id="menuBar"
          side="${current.side}"
          module="${current.module}"
          ?scrollTop="${current.scrollTop}"
          .bookmark="${data[APP_MODULE.BOOKMARK]}"
          selectorPage=${this.getSetting("selectorPage")}
          .msg="${(...args)=>this.msg(...args)}"
          .onEvent="${new MenuController(this, app)}"
        ></client-menubar>
      </div>
      <div theme="${current.theme}" class="main">
        ${(current.module === APP_MODULE.SEARCH) ? html`<client-search
          id="search" .data=${data[APP_MODULE.SEARCH]} side="${current.side}"
          .queries="${this.queries}"
          .quick="${this.quick}"
          .auditFilter="${data[APP_MODULE.LOGIN].data.audit_filter}"
          paginationPage=${this.getSetting("paginationPage")}
          .onEvent=${new SearchController(this, app)}
          .msg=${(...args)=>this.msg(...args)}
        ></client-search>` :nothing}
        ${(current.modalForm) ? current.modalForm : nothing}
      </div>`
    }
    return html`
      <client-login id="Login"
        version="${session.version}"
        serverURL="${session.serverURL}"
        .locales="${{...session.locales}}"
        lang="${current.lang}"
        .msg="${(...args)=>this.msg(...args)}"
        theme="${current.theme}"
        .data="${{...data[APP_MODULE.LOGIN]}}"
        .onEvent="${new LoginController(this, app)}"
       >
      </client-login>
    `;
  }

  render() {
    const { current } = this.state.data
    return html`
      <form-toast id="appToast" 
        .store="${this.getStore()}" 
        timeout=${this.getSetting("toastTimeout")} ></form-toast>
      ${this.protector()}
      ${(current.request)?html`<form-spinner></form-spinner>`:``}
    `
  }
}

