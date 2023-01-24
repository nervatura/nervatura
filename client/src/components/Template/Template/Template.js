import { LitElement, html } from 'lit';

import '../../SideBar/Template/sidebar-template.js'
import '../Editor/template-editor.js'
import '../../Modal/Template/modal-template.js'

import { styles } from './Template.styles.js'
import { SIDE_VISIBILITY } from '../../../config/enums.js'

export class Template extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.id = Math.random().toString(36).slice(2);
    this.side = SIDE_VISIBILITY.AUTO
    this.data = {} 
    this.paginationPage = 10
    this.onEvent = {}
    this.modalTemplate = this.modalTemplate.bind(this)
  }

  static get properties() {
    return {
      id: { type: String },
      side: { type: String },
      data: { type: Object },
      paginationPage: { type: Number },
      onEvent: { type: Object },
    };
  }

  static get styles () {
    return [
      styles
    ]
  }

  connectedCallback() {
    super.connectedCallback();
    this.onEvent.setModule(this)
  }

  modalTemplate({
    type, name, columns, onEvent
  }){
    return html`<modal-template
      type="${type}"
      name="${name}"
      columns="${columns}"
      .onEvent=${onEvent} .msg=${this.msg}
    ></modal-template>`
  }

  render() {
    const { side, data, paginationPage } = this
    return html`<sidebar-template
      id="${this.id}" side="${side}"
      templateKey="${data.key}" ?dirty="${data.dirty}"
      .onEvent=${this.onEvent} .msg=${this.msg}
    ></sidebar-template>
      <div class="page">
        <template-editor
          .data="${data}"
          .onEvent=${this.onEvent} .msg=${this.msg}
          paginationPage="${paginationPage}"
        ></template-editor>
      </div>`
  }
}