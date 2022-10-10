import { LitElement, html } from 'lit';

import { styles } from './NtSpinner.style.js'

export class NtSpinner extends LitElement {
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

