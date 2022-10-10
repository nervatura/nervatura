import { LitElement, html } from 'lit';
import { ifDefined } from 'lit/directives/if-defined.js';
import { styleMap } from 'lit/directives/style-map.js';

import { styles } from './NtButton.style.js'

export const BUTTON_TYPE = {
  PRIMARY: "primary",
  SECONDARY: "secondary", 
  BORDER: "border", 
}

export class NtButton extends LitElement {
  constructor() {
    super();
    this.id = Math.random().toString(36).slice(2);
    this.type = BUTTON_TYPE.PRIMARY
    this.name = undefined
    this.label = ""
    this.disabled = false
    this.autofocus = false
    this.full = false
    this.small = false
    this.style = {}
  }

  static get properties() {
    return {
      id: { type: String },
      name: { type: String, reflect: true },
      type: { 
        type: String, 
        converter: (value) => {
          if(!Object.values(BUTTON_TYPE).includes(value)){
            return BUTTON_TYPE.PRIMARY
          }
          return value
        } 
      },
      label: { type: String },
      disabled: { type: Boolean, reflect: true },
      autofocus: { type: Boolean, reflect: true },
      full: { type: Boolean },
      small: { type: Boolean },
      style: { type: Object }
    };
  }

  _onClick (e) {
    e.stopPropagation();
    if(!this.disabled){
      if(this.onClick){
        this.onClick(e)
      }
      this.dispatchEvent(
        new CustomEvent('click', {
          bubbles: true, composed: true,
          detail: {
            id: this.id
          }
        })
      );
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
              id: this.id
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

  render() {
    return html`<button 
      id="${this.id}"
      name="${ifDefined(this.name)}"
      button-type="${this.type}"
      ?disabled="${this.disabled}"
      ?autofocus="${this.autofocus}"
      aria-label="${ifDefined(this.label)}"
      class=${`${(this.small)?`small`:``} ${(this.full)?"full":""}`}
      style="${styleMap(this.style)}"
      @click=${this._onClick}
      @keydown=${this._onKeyEvent}
      @keypress=${this._onKeyEvent}
      ><slot id="value"></slot></button>`;
  }

  static get styles () {
    return [
      styles
    ]
  }
}
