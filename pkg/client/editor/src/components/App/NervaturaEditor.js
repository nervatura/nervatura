import { LitElement, html, nothing } from 'lit';
import { ifDefined } from 'lit/directives/if-defined.js';
import { cache } from 'lit/directives/cache.js';

import variablesCSS from '../../config/variables.js';
import { APP_MODULE } from '../../config/enums.js'
import { styles } from './NervaturaEditor.styles.js'

import { StateController } from '../../controllers/StateController.js'
import { AppController } from '../../controllers/AppController.js'

import { store as storeConfig } from '../../config/app.js'

import '../Form/Spinner/form-spinner.js'
import '../Form/Toast/form-toast.js'

export class NervaturaEditor extends LitElement {

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
    this.app = new AppController(this)

    this.inputBox = this.inputBox.bind(this)
    this.setData = this.setData.bind(this)
  }

  connectedCallback() {
    super.connectedCallback();
    // window.addEveListener("scroll", this._onScroll.bind(this), {passive: true});
    // window.addEveListener('resize', this._onResize.bind(this));
  }

  disconnectedCallback() {
    // window.removeEveListener("scroll", this._onScroll.bind(this));
    // window.removeEveListener('resize', this._onResize.bind(this));
    super.disconnectedCallback();
  }
  
  /*
  _onResize() {
    const { current } = this.state.data
    if((current.clientHeight !== window.innerHeight) || 
      (current.clientWidth !== window.innerWidth)){
        this.setData("current", {
          clientHeight: window.innerHeight, clientWidth: window.innerWidth
        })
    }
  }
  */

  _onScroll() {
    const { current } = this.state.data
    const scrollTop = ((document.body.scrollTop > 100) || (document.documentElement.scrollTop > 100))
    if(current.scrollTop !== scrollTop){
      this.setData("current", {
        scrollTop
      })
    }
  }

  setData(key, value, update) {
    this.state.data = { key, value, update }
  }

  inputBox({ 
    title, message, infoText, value, defaultOK, showValue, labelCancel, labelOK, onEvent 
  }){
    return html`<modal-inputbox
      title="${ifDefined(title)}"
      message="${ifDefined(message)}"
      infoText="${ifDefined(infoText)}"
      value="${ifDefined(value)}"
      labelCancel="${labelCancel || this.app.msg("", { id: "msg_cancel" })}"
      labelOK="${labelOK || this.app.msg("", { id: "msg_ok" })}"
      ?defaultOK="${defaultOK||false}"
      ?showValue="${showValue||false}"
      .onEvent=${onEvent}
    ></modal-inputbox>`
  }

  _protector(){
    const { data } = this.state
    const { current, session } = this.state.data
    if(data[APP_MODULE.LOGIN].data){
      return html`
      <div theme="${current.theme}" class="main">
        ${cache((current.module === APP_MODULE.TEMPLATE) ? html`<client-template
          id="template" .data=${data[APP_MODULE.TEMPLATE]} 
          side="${current.side}"
          paginationPage=${this.app.getSetting("paginationPage")}
          theme="${current.theme}"
          .onEvent=${this.app.modules.template}
          .msg="${this.app.msg}"
        ></client-template>` : nothing)}
        ${(current.modalForm) ? current.modalForm : nothing}
      </div>`
    }
    return html`
      <client-login id="Login"
        version="${session.version}"
        serverURL="${session.serverURL}"
        .locales="${{...session.locales}}"
        lang="${current.lang}"
        theme="${current.theme}"
        .data="${{...data[APP_MODULE.LOGIN]}}"
        .app="${this.app}" .msg="${this.app.msg}"
       >
      </client-login>
    `;
  }

  render() {
    const { current } = this.state.data
    return html`<style> :host { background-color: rgb(var(--${current.theme})); } </style>
      <form-toast id="appToast" 
        .setData="${this.setData}" ></form-toast>
      ${this._protector()}
      ${cache((current.request) ? html`<form-spinner></form-spinner>` : nothing)}
    `
  }
}

