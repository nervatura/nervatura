import { LitElement, html, css } from 'lit';
import { ifDefined } from 'lit/directives/if-defined.js';
import { styleMap } from 'lit/directives/style-map.js';

import { IconData } from './IconData.js'

export class NtIcon extends LitElement {
  constructor() {
    super();
    this.id = undefined
    this.iconKey = "ExclamationTriangle"
    this.width = undefined
    this.height = undefined
    this.color = undefined
    this.style = ""
  }

  static get properties() {
    return {
      id: { type: String },
      iconKey: { 
        type: String,
        converter: (value) => {
          if(!Object.keys(IconData).includes(value)){
            return Object.keys(IconData)[0]
          }
          return value
        }
      },
      width: { type: Number },
      height: { type: Number },
      color: { type: String },
      style: { type: Object },
    };
  }

  render() {
    const style = (typeof(this.type) === "object") 
      ? styleMap(this.style) : this.style || ""
    return html`<svg xmlns="http://www.w3.org/2000/svg" 
      viewBox=${IconData[this.iconKey].viewBox}
      id="${ifDefined(this.id)}"
      width=${this.width || IconData[this.iconKey].width} 
      height=${this.height || IconData[this.iconKey].height}
      style="${style}">
    <g fill="${ifDefined(this.color)}">
      <path d=${IconData[this.iconKey].path}></path>
    </g>
  </svg>`;
  }

  static get styles () {
    return [
      css`
        svg {
          vertical-align: middle;
        }
        svg:hover {
          fill: var(--icon-hover);
        }
      `
    ]
  }
}
