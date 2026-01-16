import { LitElement, html } from 'lit';
import { ifDefined } from 'lit/directives/if-defined.js';
import { styleMap } from 'lit/directives/style-map.js';

import { styles } from './Select.styles.js'

export class Select extends LitElement {
  constructor() {
    super();
    this.id = Math.random().toString(36).slice(2);
    this.value = ""
    this.name = undefined
    this.options = []
    this.isnull = true
    this.label = ""
    this.disabled = false
    this.autofocus = false
    this.full = false
    this.style = {}
  }

  static get properties() {
    return {
      id: { type: String },
      name: { type: String, reflect: true },
      value: { 
        type: String, reflect: true,
      },
      options: { type: Array },
      isnull: { type: Boolean },
      label: { type: String },
      disabled: { type: Boolean, reflect: true },
      autofocus: { type: Boolean, reflect: true },
      full: { type: Boolean },
      style: { type: Object }
    };
  }

  _onInput(e){
    if(e.target.value !== this.value){
      if(this.onChange){
        this.onChange({ value: e.target.value, old: this.value})
      }
      this.dispatchEvent(
        new CustomEvent('change', {
          bubbles: true, composed: true,
          detail: {
            value: e.target.value, old: this.value
          }
        })
      );
      this.value = e.target.value
    }
    if(this._select.value !== e.target.value){
      this._select.value = e.target.value
    }
  }

  _onKeyEvent (e) {
    const onEnter = () => {
      if(this.onEnter){
        this.onEnter({ value: this.value })
        this.dispatchEvent(
          new CustomEvent('enter', {
            bubbles: true, composed: true,
            detail: {
              value: this.value
            }
          })
        );
      }
    }
    if (e.type === 'keydown' || e.type === 'keypress') {
      e.stopPropagation();
    }
    // Here we prevent keydown on enter key from modifying the value
    if (e.type === 'keydown' && e.keyCode === 13) {
      e.preventDefault();
      onEnter();
    }
    // Request implicit submit with keypress on enter key
    if (!this.readonly && e.type === 'keypress' && e.keyCode === 13) {
      onEnter();
    }
  }

  firstUpdated() {
    this._select = this.renderRoot.querySelector('select');
  }

  render() {
    const values = this.options.map((item,index)=>html`<option
      key=${index} value=${item.value} 
      ?selected=${(item.value === this.value)} >${item.text}</option>`)
    if(this.isnull){
      values.unshift(html`<option
        ?selected=${(this.value === "")}
        key="-1" value="" ></option>`)
    }
    return html`<select 
      id="${this.id}"
      name="${ifDefined(this.name)}"
      .value="${this.value}"
      ?disabled="${this.disabled}"
      ?autofocus="${this.autofocus}"
      aria-label="${ifDefined(this.label)}"
      class="${(this.full)?"full":""}"
      style="${styleMap(this.style)}"
      @input=${this._onInput}
      @keydown=${this._onKeyEvent}
      @keypress=${this._onKeyEvent}
      >${values}</select>`
  }

  static get styles () {
    return [
      styles
    ]
  }
}

