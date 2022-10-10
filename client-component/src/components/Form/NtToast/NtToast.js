/* eslint-disable lit-a11y/click-events-have-key-events */
import { LitElement, html } from 'lit';
import { ifDefined } from 'lit/directives/if-defined.js';
import { styleMap } from 'lit/directives/style-map.js';

import '../NtIcon/nt-icon.js';
import { styles } from './NtToast.styles.js'

export const TOAST_TYPE = {
  INFO: "info",
  ERROR: "error", 
  SUCCESS: "success"
}

export class NtToast extends LitElement {
  constructor() {
    super();
    this.id = Math.random().toString(36).slice(2);
    this.name = undefined;
    this.type = TOAST_TYPE.INFO;
    this.hidden = true;
    this.timeout = 3;
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
      type: { 
        type: String, 
        converter: (value) => {
          if(!Object.values(TOAST_TYPE).includes(value)){
            return TOAST_TYPE.INFO
          }
          return value
        } 
      },
      hidden: { type: Boolean },
      timeout: { type: Number },
      style: { type: Object },
    };
  }

  connectedCallback() {
    super.connectedCallback();
    if(this.store){
      const { data, setData } = this.store
      setData("current", {
        ...data.current,
        toast: this
      }, false)
    }
  }

  disconnectedCallback() {
    if(this.store){
      const { data, setData } = this.store
      setData("current", {
        ...data.current,
        toast: null
      }, false)
    }
    super.disconnectedCallback();
  }

  show() {
    /* c8 ignore next 1 */
		if (!this.hidden) return;
    this.style.display = 'block';
    this.timeoutVar = setTimeout(() => {
      /* c8 ignore next 11 */
      this.hidden = !this.hidden;
      this.toggleToastClass();
      this.timeoutVar = setTimeout(() => {
        if (this && !this.hidden) {
          this.hidden = !this.hidden;
          this.timeoutVar = setTimeout(() => {
            this.style.display = 'none';
          }, 300);
          this.toggleToastClass();
        }
      }, this.timeout * 1000);
    }, 30);
	}

	close() {
		if (this.hidden) return;
    /* c8 ignore next 10 */
		clearTimeout(this.timeoutVar);
		setTimeout(() => {
			if (this && !this.hidden) {
				this.hidden = !this.hidden;
				setTimeout(() => {
          this.style.display = 'none';
        }, 300);
				this.toggleToastClass();
			}
		}, 30);
	}

	toggleToastClass() {
		const toast = this.shadowRoot.querySelector('div');
		toast.classList.toggle('show');
	}

  render() {
    return html`<div 
      id="${this.id}"
      type="${this.type}"
      name="${ifDefined(this.name)}"
      @click=${this.close}
      style="${styleMap(this.style)}" >
        <span class="icon">
          ${html`<nt-icon iconKey="${this._iconMap[this.type]}" ></nt-icon>`}
        </span>
        <slot id="value"></slot>
    </div>`
  }

  static get styles () {
    return [
      styles
    ]
  }
}