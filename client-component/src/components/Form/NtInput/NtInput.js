import { LitElement, html } from 'lit';
import { ifDefined } from 'lit/directives/if-defined.js';
import { styleMap } from 'lit/directives/style-map.js';

import { styles } from './NtInput.styles.js'

export const INPUT_TYPE = {
  TEXT: "text",
  COLOR: "color", 
  FILE: "file", 
  PASSWORD: "password"
}

export class NtInput extends LitElement {
  constructor() {
    super();
    this.id = Math.random().toString(36).slice(2);
    this.type = INPUT_TYPE.TEXT
    this.value = ""
    this.name = undefined
    this.placeholder = undefined
    this.label = ""
    this.disabled = false
    this.readonly = false
    this.autofocus = false
    this.accept = undefined
    this.maxlength = undefined
    this.size = undefined
    this.full = false
    this.style = {}
  }

  static get properties() {
    return {
      id: { type: String },
      name: { type: String, reflect: true },
      type: { 
        type: String, 
        converter: (value) => {
          if(!Object.values(INPUT_TYPE).includes(value)){
            return INPUT_TYPE.TEXT
          }
          return value
        } 
      },
      value: { 
        type: String, reflect: true,
      },
      placeholder: { type: String },
      label: { type: String },
      disabled: { type: Boolean, reflect: true },
      readonly: { type: Boolean, reflect: true },
      autofocus: { type: Boolean, reflect: true },
      accept: { type: String, reflect: true },
      maxlength: { type: Number, reflect: true },
      size: { type: Number, reflect: true },
      full: { type: Boolean },
      style: { type: Object },
    };
  }

  _onInput(e){
    const value = (this.type === INPUT_TYPE.FILE) ? e.target.files : e.target.value
    if(value !== this.value){
      if(this.onChange){
        this.onChange({ 
          value, 
          old: this.value
        })
      }
      this.dispatchEvent(
        new CustomEvent('change', {
          bubbles: true, composed: true,
          detail: {
            value, old: this.value
          }
        })
      );
      this.value = value
    }
    if(this._input.value !== value){
      this._input.value = value
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
    this._input = this.renderRoot.querySelector('input');
  }

  render() {
    return html`<input 
      id="${this.id}"
      name="${ifDefined(this.name)}"
      .type="${this.type}"
      .value="${this.value}"
      placeholder="${ifDefined(this.placeholder)}"
      ?disabled="${this.disabled}"
      ?readonly="${this.readonly}"
      ?autofocus="${this.autofocus}"
      aria-label="${ifDefined(this.label)}"
      accept="${ifDefined(this.accept)}"
      maxlength="${ifDefined(this.maxlength)}"
      size="${ifDefined(this.size)}"
      class="${(this.full)?"full":""}"
      style="${styleMap(this.style)}"
      @input=${this._onInput}
      @keydown=${this._onKeyEvent}
      @keypress=${this._onKeyEvent}
    >`
  }

  static get styles () {
    return [
      styles
    ]
  }
}
