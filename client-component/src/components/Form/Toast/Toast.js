import { LitElement, html, nothing } from 'lit';
import { ifDefined } from 'lit/directives/if-defined.js';
import { styleMap } from 'lit/directives/style-map.js';

import '../Icon/form-icon.js';
import { styles } from './Toast.styles.js'

import { TOAST_TYPE } from '../../../config/enums.js'

export class Toast extends LitElement {
  constructor() {
    super();
    this.id = Math.random().toString(36).slice(2);
    this.name = undefined;
    this.type = TOAST_TYPE.INFO;
    this.hidden = true;
    this.timeout = 4;
    this.style = {}
    this._iconMap = {
      info: "InfoCircle",
      error: "ExclamationTriangle",
      success: "CheckSquare"
    }
  }

  static get properties() {
    return {
      id: { type: String },
      name: { type: String, reflect: true },
      type: { type: String },
      hidden: { type: Boolean },
      timeout: { type: Number },
      style: { type: Object },
    };
  }

  connectedCallback() {
    super.connectedCallback();
    if(this.store){
      const { setData } = this.store
      setData("current", {
        toast: this
      }, false)
    }
  }

  disconnectedCallback() {
    if(this.store){
      const { setData } = this.store
      setData("current", {
        toast: null
      }, false)
    }
    super.disconnectedCallback();
  }

  show({ type, value, timeout }) {
		if (!this.hidden) return;
    this.value = value || this.value
    this.type = type || this.type
    this.timeout = (typeof(timeout) !== "undefined") ? timeout : this.timeout
    this.hidden = false;
    if(this.timeout > 0){
      this.timeoutVar = setTimeout(() => {
        /* c8 ignore next 3 */
        if (this && !this.hidden) {
          this.hidden = true;
        }
      }, this.timeout * 1000);
    }
	}

	close() {
		if (this.hidden) return;
		clearTimeout(this.timeoutVar);
    this.hidden = true
	}

  render() {
    if(!this.hidden){
      return html`<div 
        id="${this.id}"
        type="${this.type}"
        name="${ifDefined(this.name)}"
        @click=${this.close}
        style="${styleMap(this.style)}" >
          <span class="icon">
            ${html`<form-icon 
              iconKey="${this._iconMap[this.type]||"InfoCircle"}"
              .style=${{ margin: "auto" }} 
              width=32 height=32 ></form-icon>`}
          </span>
          <slot id="value">${ifDefined(this.value)}</slot>
      </div>`
    }
    return nothing
  }

  static get styles () {
    return [
      styles
    ]
  }
}