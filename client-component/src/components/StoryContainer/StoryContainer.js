import { LitElement, html, css } from 'lit';

import variablesCSS from '../../config/variables.js';

export class StoryContainer extends LitElement {
  constructor() {
    super();
    this.theme = "light"
  }

  static get properties() {
    return {
      theme: { 
        type: String, 
        converter: (value) => {
          if(!["light", "dark"].includes(value)){
            return "light"
          }
          return value
        } 
      }
    };
  }

  render() {
    return html`<div theme="${this.theme}"
      class=${`container`}>
        <slot></slot>
      </div>`;
  }
}

StoryContainer.styles = [
  variablesCSS,
  css`
    :host {
      display: block;
      font-family: var(--font-family);
      font-size: var(--font-size);
    }
    div {
      box-sizing: border-box;
    }
    .container {
      display: table;
      width: 100%;
      height: 100%;
      padding: 10px;
      background-color: rgb(var(--base-1));
    }
  `
]
