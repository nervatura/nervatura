import { LitElement, html, css } from 'lit';
import { styleMap } from 'lit/directives/style-map.js';

import variablesCSS from '../../config/variables.js';
import { APP_THEME } from '../../config/enums.js'

export class StoryContainer extends LitElement {
  constructor() {
    super();
    this.theme = APP_THEME.LIGHT
    this.style = {}
  }

  static get properties() {
    return {
      theme: { type: String },
      style: { type: Object }
    };
  }

  render() {
    return html`<div theme="${this.theme}" style="${styleMap(this.style)}"
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
      padding: 0px;
      background-color: rgb(var(--base-1));
    }
  `
]
