import { LitElement, html, nothing } from 'lit';

import '../../SideBar/Setting/sidebar-setting.js'
import '../Form/setting-form.js'
import '../View/setting-view.js'

import '../../Modal/Audit/modal-audit.js'
import '../../Modal/Menu/modal-menu.js'

import { styles } from './Setting.styles.js'
import { SIDE_VISIBILITY } from '../../../config/enums.js'

export class Setting extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.id = Math.random().toString(36).slice(2);
    this.side = SIDE_VISIBILITY.AUTO
    this.data = {
      caption: "", 
      icon: "",
      view: {},
      actions: {},
      current: undefined,
      audit: "", 
      dataset: {},
      type: ""
    } 
    this.auditFilter = {}
    this.username = ""
    this.paginationPage = 10
    this.onEvent = {}
    this.modalAudit = this.modalAudit.bind(this)
    this.modalMenu = this.modalMenu.bind(this)
  }

  static get properties() {
    return {
      id: { type: String },
      side: { type: String },
      data: { type: Object },
      auditFilter: { type: Object },
      username: { type: String },
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

  modalAudit({
    idKey, usergroup, nervatype, subtype, inputfilter, supervisor,
    typeOptions, subtypeOptions, inputfilterOptions, onEvent
  }){
    return html`<modal-audit
      idKey="${idKey}"
      usergroup="${usergroup}"
      nervatype="${nervatype}"
      subtype="${subtype}"
      inputfilter="${inputfilter}"
      supervisor="${supervisor}"
      .typeOptions="${typeOptions}"
      .subtypeOptions="${subtypeOptions}"
      .inputfilterOptions="${inputfilterOptions}"
      .onEvent=${onEvent} .msg=${this.msg}
    ></modal-audit>`
  }

  modalMenu({
    idKey, menu_id, fieldname, description, orderby, fieldtype, fieldtypeOptions, onEvent
  }){
    return html`<modal-menu
      idKey="${idKey}"
      menu_id="${menu_id}"
      fieldname="${fieldname}"
      description="${description}"
      orderby="${orderby}"
      fieldtype="${fieldtype}"
      .fieldtypeOptions="${fieldtypeOptions}"
      .onEvent=${onEvent} .msg=${this.msg}
    ></modal-menu>`
  }

  render() {
    const { side, data, auditFilter, username, paginationPage } = this
    return html`<sidebar-setting
      id="${this.id}" side="${side}"
      username="${username}" .auditFilter="${auditFilter}"
      .module="${data}"
      .onEvent=${this.onEvent} .msg=${this.msg}
    ></sidebar-setting>
      <div class="page">
        ${(data.current) ? html`<setting-form
          .data="${data}"
          .onEvent=${this.onEvent} .msg=${this.msg}
          paginationPage="${paginationPage}"
        ></setting-form>` : 
        (data.view) ? html`<setting-view
          .data="${data}"
          paginationPage="${paginationPage}"
          .onEvent=${this.onEvent} .msg=${this.msg}
        ></setting-view>` : nothing}
      </div>`
  }
}