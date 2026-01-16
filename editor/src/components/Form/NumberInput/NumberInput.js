import { LitElement, html } from 'lit';
import { ifDefined } from 'lit/directives/if-defined.js';
import { styleMap } from 'lit/directives/style-map.js';

import { styles } from './NumberInput.styles.js'

export class NumberInput extends LitElement {
  constructor() {
    super();
    this.id = Math.random().toString(36).slice(2)
    this.name = undefined
    this.value = 0
    this.integer = false
    this.label = undefined
    this.max = undefined
    this.min = undefined
    this.disabled = false
    this.readonly = false
    this.autofocus = false
    this.full = false
    this.style = {}    
  }

  static get properties() {
    return {
      id: { type: String },
      name: { type: String, reflect: true },
      value: { 
        type: Number, reflect: true,
      },
      integer: { type: Boolean },
      label: { type: String },
      max: { type: Number },
      min: { type: Number },
      disabled: { type: Boolean, reflect: true },
      readonly: { type: Boolean, reflect: true },
      autofocus: { type: Boolean, reflect: true },
      full: { type: Boolean },
      style: { type: Object },
    };
  }

  _onInput(e){
    const onChange = (value) => {
      if(value !== this.value){
        if(this.onChange){
          this.onChange({ value, old: this.value})
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
      if(this._input.value !== String(value)){
        this._input.value = String(value)
      }
    }
    if(e.target.valueAsNumber !== this.value){
      let value = e.target.valueAsNumber
      if(Number.isNaN(value)){
        value = 0
      }
      if((typeof(this.min) !== "undefined") && (value < this.min)){
        value = this.min
      }
      if((typeof(this.max) !== "undefined") && (value > this.max)){
        value = this.max
      }
      if(this.integer){
        value = Math.floor(value)
      }
      onChange(value)
    }
  }

  _onBlur(){
    this._input.value = String(this.value)
    if(this.onBlur){
      this.onBlur({ value: this.value })
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
      type="number"
      .value="${String(this.value)}"
      ?disabled="${this.disabled}"
      ?readonly="${this.readonly}"
      ?autofocus="${this.autofocus}"
      min="${ifDefined(this.min)}"
      max="${ifDefined(this.max)}"
      aria-label="${ifDefined(this.label)}"
      class="${(this.full)?"full":""}"
      style="${styleMap(this.style)}"
      @input=${this._onInput}
      @blur=${this._onBlur}
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
