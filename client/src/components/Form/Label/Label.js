import { LitElement, html } from 'lit';
import { ifDefined } from 'lit/directives/if-defined.js';
import { styleMap } from 'lit/directives/style-map.js';

import '../Icon/form-icon.js';

import { styles } from './Label.style.js'

export class Label extends LitElement {
  constructor() {
    super();
    this.id = undefined
    this.value = ""
    this.centered = false
    this.leftIcon = undefined
    this.rightIcon= undefined
    this.style = {}
    this.iconStyle = {}
  }

  static get properties() {
    return {
      id: { type: String },
      value: { type: String },
      centered: { type: Boolean },
      leftIcon: { type: String },
      rightIcon: { type: String },
      style: { type: Object },
      iconStyle: { type: Object },
    };
  }

  render() {
    if(this.leftIcon){
      return html`<div 
        id="${ifDefined(this.id)}" 
        class="row ${(this.centered)?`centered`:``}">
        <div class="cell label_icon_left">
          ${html`<form-icon 
            iconKey="${this.leftIcon}" width=20
            color="${ifDefined(this.iconStyle.color)}"
            .style="${this.iconStyle}"></form-icon>`}
        </div>
        <div class="cell label_info_left bold"
          style="${styleMap(this.style)}" >${this.value}</div>
      </div>`    
    }
    if(this.rightIcon){
      return html`<div 
        id="${ifDefined(this.id)}" class="row full">
        <div class="cell label_info_right bold"
          style="${styleMap(this.style)}" >${this.value}</div>
        <div class="cell label_icon_right">
          ${html`<form-icon 
            iconKey="${this.rightIcon}" width=20
            color="${ifDefined(this.iconStyle.color)}"
            .style="${this.iconStyle}"></form-icon>`}
        </div>
      </div>`
    }
    return html`<span 
      id="${ifDefined(this.id)}" class="bold" 
      style="${styleMap(this.style)}">${this.value}</span>`
  }

  static get styles () {
    return [
      styles
    ]
  }
}

