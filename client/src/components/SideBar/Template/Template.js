import { LitElement, html } from 'lit';

import '../../Form/Button/form-button.js'

import { styles } from './Template.styles.js'
import { SIDE_VISIBILITY, SIDE_EVENT, BUTTON_TYPE, TEXT_ALIGN } from '../../../config/enums.js'

export class Template extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.side = SIDE_VISIBILITY.AUTO
    this.templateKey = ""
    this.dirty = false 
  }

  static get properties() {
    return {
      side: { type: String, reflect: true },
      templateKey: { type: String },  
      dirty: { type: Boolean },
    };
  }

  static get styles () {
    return [
      styles
    ]
  }

  _onSideEvent(key, data){
    if(this.onEvent && this.onEvent.onSideEvent){
      this.onEvent.onSideEvent({ key, data })
    }
    this.dispatchEvent(
      new CustomEvent('side_event', {
        bubbles: true, composed: true,
        detail: {
          key, data
        }
      })
    );
  }

  itemMenu({id, selected, eventValue, label, iconKey, full}){
    const btnSelected = (typeof(selected) === "undefined") ? false : selected
    const btnFull = (typeof(full) === "undefined") ? true : full
    const style = {
      "border-radius": "0",
      "border-color": "rgba(var(--accent-1c), 0.2)",
    }
    return(
      html`<form-button 
        id="${id}" label="${label}"
        ?full="${btnFull}" ?selected="${btnSelected}"
        align=${TEXT_ALIGN.LEFT}
        .style="${style}"
        icon="${iconKey}" type="${BUTTON_TYPE.PRIMARY}"
        @click=${()=>this._onSideEvent( ...eventValue )} 
      >${label}</form-button>`
    )
  }

  formItems(){
    const panels = []

    panels.push(
      this.itemMenu({
        id: "cmd_back",
        selected: true, 
        eventValue: [SIDE_EVENT.CHECK, { value: "LOAD_SETTING" }],
        label: this.msg("", { id: "label_back" }), 
        iconKey: "Reply", full: false, 
      })
    )
    panels.push(html`<hr id="back_sep" class="separator" />`)

    if(!["_blank", "_sample"].includes(this.templateKey)){
      panels.push(html`<hr id="tmp_sep_2" class="separator" />`)
      panels.push(
        this.itemMenu({
          id: "cmd_save",
          selected: this.dirty,
          eventValue: [SIDE_EVENT.SAVE, true],
          label: this.msg("", { id: "template_save" }), 
          iconKey: "Check"
        })
      )
      panels.push(
        this.itemMenu({
          id: "cmd_create",
          eventValue: [SIDE_EVENT.CREATE_REPORT, {}],
          label: this.msg("", { id: "template_create_from" }), 
          iconKey: "Sitemap"
        })
      )
      panels.push(
        this.itemMenu({
          id: "cmd_delete",
          eventValue: [SIDE_EVENT.DELETE, {}],
          label: this.msg("", { id: "label_delete" }), 
          iconKey: "Times"
        })
      )
    }

    panels.push(html`<hr id="tmp_sep_3" class="separator" />`)
    panels.push(
      this.itemMenu({
        id: "cmd_blank",
        eventValue: [SIDE_EVENT.CHECK, { value: SIDE_EVENT.BLANK }],
        label: this.msg("", { id: "template_new_blank" }), 
        iconKey: "Plus"
      })
    )
    panels.push(
      this.itemMenu({
        id: "cmd_sample",
        eventValue: [SIDE_EVENT.CHECK, { value: SIDE_EVENT.SAMPLE }],
        label: this.msg("", { id: "template_new_sample" }), 
        iconKey: "Plus"
      })
    )

    panels.push(html`<hr id="tmp_sep_4" class="separator" />`)
    panels.push(
      this.itemMenu({
        id: "cmd_print",
        eventValue: [SIDE_EVENT.REPORT_SETTINGS, { value: "PREVIEW" }],
        label: this.msg("", { id: "label_print" }), 
        iconKey: "Eye"
      })
    )
    panels.push(
      this.itemMenu({
        id: "cmd_json",
        eventValue: [SIDE_EVENT.REPORT_SETTINGS, { value: "JSON" }],
        label: this.msg("", { id: "template_export_json" }), 
        iconKey: "Code"
      })
    )

    panels.push(html`<hr id="tmp_sep_5" class="separator" />`)
    panels.push(
      this.itemMenu({
        id: "cmd_help", 
        eventValue: [SIDE_EVENT.HELP, { value: "program/editor" }],
        label: this.msg("", { id: "label_help" }), 
        iconKey: "QuestionCircle"
      })
    )

    return panels
  }

  render() {
    return html`<div class="sidebar ${(this.side !== "auto") ? this.side : ""}" >
    ${this.formItems()}
    </div>`
  }

}