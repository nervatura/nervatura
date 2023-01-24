import { LitElement, html } from 'lit';

import { styles } from './Spinner.style.js'

export class Spinner extends LitElement {
  render() {
    return html`
    <div class="modal" >
      <div class="middle" >
        <div class="loading">
          <div></div>
          <div></div>
          <div></div>
          <div></div>
          <div></div>
          <div></div>
          <div></div>
          <div></div>
        </div>
      </div>
    </div>`;
  }

  static get styles () {
    return [
      styles
    ]
  }
}

