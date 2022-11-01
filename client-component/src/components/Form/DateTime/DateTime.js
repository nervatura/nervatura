import { LitElement, html } from 'lit';
import { ifDefined } from 'lit/directives/if-defined.js';
import { styleMap } from 'lit/directives/style-map.js';

import { styles } from './DateTime.styles.js'
import { DATETIME_TYPE } from '../../../config/enums.js'

export class DateTime extends LitElement {
  constructor() {
    super();
    this.id = Math.random().toString(36).slice(2)
    this.name = undefined
    this.value = ""
    this.type = DATETIME_TYPE.DATE
    this.label = undefined
    this.isnull = true
    this.picker = false
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
        type: String, reflect: true,
      },
      type: { 
        type: String, 
        converter: (value) => {
          if(!Object.values(DATETIME_TYPE).includes(value)){
            return DATETIME_TYPE.DATE
          }
          return value
        } 
      },
      label: { type: String },
      isnull: { type: Boolean },
      picker: { type: Boolean },
      disabled: { type: Boolean, reflect: true },
      readonly: { type: Boolean, reflect: true },
      autofocus: { type: Boolean, reflect: true },
      full: { type: Boolean },
      style: { type: Object },
    };
  }

  connectedCallback() {
    super.connectedCallback();
    if((this.type === DATETIME_TYPE.DATE) && (this.value.length > 10)){
      this.value = this.value.slice(0,10)
    }
  }

  _defaultValue(){
    const defaultValue = new Date().toISOString()
    switch (this.type) {
      case DATETIME_TYPE.DATE:
        return String(defaultValue).split("T")[0]

      case DATETIME_TYPE.TIME:
        return String(defaultValue).split("T")[1].split(".")[0]

      default:
        return String(defaultValue).substring(0,16);
    }
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
      if(this._input.value !== value){
        this._input.value = value
      }
    }
    if(e.target.value !== this.value){
      if((e.target.value === "") && !this.isnull){
        onChange(this._defaultValue())
        return
      }
      onChange(e.target.value)
    }
  }

  _onBlur(){
    this._input.value = this.value
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

  _onFocus () {
    /* c8 ignore next 3 */
    if(this.picker){
      this._input.showPicker()
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
      ?disabled="${this.disabled}"
      ?readonly="${this.readonly}"
      ?autofocus="${this.autofocus}"
      aria-label="${ifDefined(this.label)}" 
      style="${styleMap(this.style)}"
      class="${(this.full)?"full":""}"
      @input=${this._onInput}
      @blur=${this._onBlur}
      @keydown=${this._onKeyEvent}
      @keypress=${this._onKeyEvent}
      @focus=${this._onFocus}
    >`
  }

  static get styles () {
    return [
      styles
    ]
  }
}
