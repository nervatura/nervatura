import { LitElement, html, nothing } from 'lit';

import '../../SideBar/Edit/sidebar-edit.js'
import '../Editor/edit-editor.js'
import '../../Modal/Formula/modal-formula.js'
import '../../Modal/Selector/modal-selector.js'
import '../../Modal/Trans/modal-trans.js'
import '../../Modal/Stock/modal-stock.js'
import '../../Modal/Shipping/modal-shipping.js'
import '../../Modal/Report/modal-report.js'

import { styles } from './Edit.styles.js'
import { SIDE_VISIBILITY } from '../../../config/enums.js'

export class Edit extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.id = Math.random().toString(36).slice(2);
    this.side = SIDE_VISIBILITY.AUTO
    this.data = {} 
    this.auditFilter = {}
    this.newFilter = []
    this.forms = {}
    this.paginationPage = 10
    this.selectorPage = 5
    this.onEvent = {}
    this.modalFormula = this.modalFormula.bind(this)
    this.modalReport = this.modalReport.bind(this)
    this.modalSelector = this.modalSelector.bind(this)
    this.modalShipping = this.modalShipping.bind(this)
    this.modalStock = this.modalStock.bind(this)
    this.modalTrans = this.modalTrans.bind(this)
  }

  static get properties() {
    return {
      id: { type: String },
      side: { type: String },
      data: { type: Object },
      auditFilter: { type: Object },
      newFilter: { type: Array },
      forms: { type: Object },
      paginationPage: { type: Number },
      selectorPage: { type: Number },
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

  modalFormula({formula, formulaValues, partnumber, description, onEvent}){
    return html`<modal-formula 
      formula="${formula}" partnumber="${partnumber}" description="${description}"
      .formulaValues=${formulaValues} 
      .onEvent=${onEvent} .msg=${this.msg}
    ></modal-formula>`
  }

  modalReport({ 
    title, template, copy, orient, size, templates, report_size, report_orientation, onEvent 
  }) {
    return html`<modal-report
      title="${title}"
      template="${template}"
      copy="${copy}"
      orient="${orient}"
      size="${size}"
      .templates=${templates}
      .report_size=${report_size}
      .report_orientation=${report_orientation}
      .onEvent=${onEvent}
      .msg=${this.msg}
    ></modal-report>`
  }

  modalSelector({ 
    view, columns, result, filter, onEvent 
  }) {
    return html`<modal-selector
      ?isModal="${true}"
      view="${view}"
      .columns=${columns}
      .result=${result}
      filter="${filter}"
      .onEvent=${onEvent}
      .msg=${this.msg}
    ></modal-selector>`
  }

  modalShipping({ 
    unit, batch_no, qty, partnumber, description, onEvent 
  }){
    return html`<modal-shipping
      unit="${unit}"
      batch_no="${batch_no}"
      qty="${qty}"
      partnumber="${partnumber}"
      description="${description}"
      .onEvent=${onEvent}
      .msg=${this.msg}
    ></modal-shipping>`
  }

  modalStock({ 
    partnumber, partname, rows, selectorPage, onEvent 
  }) {
    return html`<modal-stock
      partnumber="${partnumber}"
      partname="${partname}"
      selectorPage="${selectorPage}"
      .rows="${rows}"
      .onEvent=${onEvent}
      .msg=${this.msg}
    ></modal-stock>`
  }

  modalTrans({
    baseTranstype, transtype, direction, doctypes, directions, 
    refno, nettoDiv, netto, fromDiv, from, elementCount, onEvent
  }){
    return html`<modal-trans
      baseTranstype="${baseTranstype}"
      transtype="${transtype}"
      direction="${direction}"
      .doctypes="${doctypes}"
      .directions="${directions}"
      ?refno="${refno}"
      ?nettoDiv="${nettoDiv}"
      ?netto="${netto}"
      ?fromDiv="${fromDiv}"
      ?from="${from}"
      elementCount="${elementCount}"
      .onEvent=${onEvent}
      .msg=${this.msg}
    ></modal-trans>`
  }

  render() {
    const { side, data, newFilter, auditFilter } = this
    return html`<sidebar-edit
      id="${this.id}" side="${side}" view="${data.side_view}"
      .newFilter="${newFilter}" .auditFilter="${auditFilter}"
      .module="${data}" .forms="${this.forms}"
      .onEvent=${this.onEvent} .msg=${this.msg}
    ></sidebar-edit>
      <div class="page">
        ${(data.current.item) ? html`<edit-editor
          caption="${data.caption}" 
          .current=${data.current} .template=${data.template}
          .dataset=${data.dataset} audit="${data.audit}"
          .onEvent=${this.onEvent} .msg=${this.msg}
        ></edit-editor>` : nothing }
      </div>`
  }
}