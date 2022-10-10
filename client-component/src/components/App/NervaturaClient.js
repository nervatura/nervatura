import { LitElement, html } from 'lit';

import variablesCSS from '../../config/variables.js';
import { styles } from './NervaturaClient.styles.js'

import { StateController } from '../../controllers/StateController.js'
import { AppController } from '../../controllers/AppController.js'
import { LoginController } from '../../controllers/LoginController.js'

import { store as storeConfig } from '../../config/app.js'

import '../LoginPage/login-page.js'
import '../Form/NtSpinner/nt-spinner.js'
import '../Form/NtToast/nt-toast.js'

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
  }

  connectedCallback() {
    super.connectedCallback();
    // window.addEventListener("scroll", this._onScroll.bind(this), {passive: true});
    // window.addEventListener('resize', this._onResize.bind(this));
    this._loadConfig()
  }

  disconnectedCallback() {
    // window.removeEventListener("scroll", this._onScroll.bind(this));
    // window.removeEventListener('resize', this._onResize.bind(this));
    super.disconnectedCallback();
  }
  
  _onResize() {
    if((this.state.data.current.clientHeight !== window.innerHeight) || 
      (this.state.data.current.clientWidth !== window.innerWidth)){
        this.setData("current", {
          ...this.state.data.current,
          clientHeight: window.innerHeight, clientWidth: window.innerWidth
        })
    }
  }

  _onScroll() {
    const scrollTop = ((document.body.scrollTop > 100) || (document.documentElement.scrollTop > 100))
    if(this.state.data.current.scrollTop !== scrollTop){
      this.setData("current", {
        ...this.state.data.current,
        scrollTop
      })
    }
  }

  async _loadConfig(){
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

    let config = {...this.state.data.session}
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
            ...this.state.data.current,
            lang: localStorage.getItem("lang")
          })
        }
      const [ current, params ] = getPath(window.location)
      if(current === "hash"){ 
        const login = new LoginController(this, this.getStore(), new AppController(this, this.getStore()))
        if (params.access_token){
          login.tokenValidation(params)
        }
        if(params.code){
          login.setCodeToken(params)
        }
      }
    } catch (error) {
      this.setData("error", error )
      this.showToast("error", error.message)
    }
  }

  getStore() {
    return {
      data: this.state.data,
      setData: (...args) => this.setData(...args),
      showToast: (...args) => this.showToast(...args),
      msg: (...args) => this.msg(...args)
    }
  }

  setData(key, value, update) {
    this.state.data = { key, value, update }
  }

  showToast(toastType, toastValue) {
    this.setData("current", {
      ...this.state.data.current,
      toastType, toastValue
    })
    if(this.state.data.current.toast){
      this.state.data.current.toast.show()
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
    if(data.login.data){
      return html`
      <div theme="${data.current.theme}" class=${["main"].join(' ')}>
        <span>${data.login.data.token}</span>
      </div>`
    }
    return html`
      <login-page id="loginPage"
        version="${this.state.data.session.version}"
        serverURL="${this.state.data.session.serverURL}"
        .locales="${{...this.state.data.session.locales}}"
        lang="${this.state.data.current.lang}"
        .msg="${(...args)=>this.msg(...args)}"
        theme="${this.state.data.current.theme}"
        .data="${{...this.state.data.login}}"
        .onEvent="${new LoginController(this, this.getStore(), new AppController(this, this.getStore()))}"
       >
      </login-page>
    `;
  }

  render() {
    const { data } = this.state
    return html`
      <nt-toast
        id="appToast"
        type="${this.state.data.current.toastType}"
        .store="${this.getStore()}"
      >${this.state.data.current.toastValue}</nt-toast>
      ${this.protector()}
      ${(data.current.request)?html`<nt-spinner></nt-spinner>`:``}
    `
  }
}

