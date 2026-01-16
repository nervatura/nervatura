import { LitElement, html, nothing } from 'lit';
import { ifDefined } from 'lit/directives/if-defined.js';
import { styleMap } from 'lit/directives/style-map.js';

import '../Icon/form-icon.js';

import { styles } from './Button.style.js'
import { TEXT_ALIGN } from '../../../config/enums.js'

export class Button extends LitElement {
  constructor() {
    super();
    this.id = Math.random().toString(36).slice(2);
    this.type = undefined
    this.name = undefined
    this.align = TEXT_ALIGN.CENTER
    this.icon = undefined
    this.label = ""
    this.disabled = false
    this.autofocus = false
    this.full = false
    this.small = false
    this.selected = false
    this.hidelabel = false
    this.badge = undefined
    this.style = {}
  }

  static get properties() {
    return {
      id: { type: String },
      name: { type: String, reflect: true },
      type: { type: String },
      align: { type: String },
      label: { type: String },
      icon: { type: String },
      disabled: { type: Boolean, reflect: true },
      autofocus: { type: Boolean, reflect: true },
      full: { type: Boolean },
      small: { type: Boolean },
      selected: { type: Boolean },
      hidelabel: { type: Boolean },
      badge: { type: Number },
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
      button-type="${ifDefined(this.type)}"
      ?disabled="${this.disabled}"
      ?autofocus="${this.autofocus}"
      aria-label="${ifDefined(this.label)}"
      title="${ifDefined(this.label)}"
      class=${`${["small", "full", "selected", "hidelabel"].filter(key => (this[key])).join(" ")} ${this.align}`}
      style="${styleMap(this.style)}"
      @click=${this._onClick}
      @keydown=${this._onKeyEvent}
      @keypress=${this._onKeyEvent}>
        ${(this.icon && (this.align !== TEXT_ALIGN.RIGHT)) 
          ? html`<form-icon iconKey="${this.icon}" width=20 ></form-icon>` : nothing}
        <slot id="value"></slot>
        ${(this.icon && (this.align === TEXT_ALIGN.RIGHT)) 
          ? html`<form-icon iconKey="${this.icon}" width=20 ></form-icon>` : nothing}
        ${(this.badge) 
          ? html`<span class="right" ><span class="${`badge ${(this.selected) ? `selected-badge` : ``}`}" >${this.badge}</span></span>` : nothing}
      </button>`;
  }

  static get styles () {
    return [
      styles
    ]
  }
}
