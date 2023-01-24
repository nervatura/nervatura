import { LitElement, html, nothing } from 'lit';

import '../../Form/Button/form-button.js'

import { styles } from './Setting.styles.js'
import { SIDE_VISIBILITY, SIDE_EVENT, BUTTON_TYPE, TEXT_ALIGN } from '../../../config/enums.js'

export class Setting extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.side = SIDE_VISIBILITY.AUTO
    this.username = undefined
    this.module = {
      current: {}, 
      dirty: false,
      panel: {},  
      group_key: ""
    }
    this.auditFilter = {} 
  }

  static get properties() {
    return {
      side: { type: String, reflect: true },
      username: { type: String },  
      module: { type: Object }, 
      auditFilter: { type: Object },
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

  itemMenu({id, selected, eventValue, label, iconKey, full, color}){
    const btnSelected = (typeof(selected) === "undefined") ? false : selected
    const btnFull = (typeof(full) === "undefined") ? true : full
    const style = {
      "border-radius": "0",
      "border-color": "rgba(var(--accent-1c), 0.2)",
    }
    if(color){
      style.color = color
      style.fill = color
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

  formItems(_options){
    const { dirty, current, type } = this.module
    const options = _options
    const panels = []

    panels.push(
      this.itemMenu({
        id: "cmd_back",
        selected: true, 
        eventValue: [SIDE_EVENT.BACK, {}],
        label: this.msg("", { id: "label_back" }), 
        iconKey: "Reply", full: false, 
      })
    )
    panels.push(html`<hr id="back_sep" class="separator" />`)

    if (options.save !== false) {
      panels.push(
        this.itemMenu({
          id: "cmd_save",
          selected: dirty,
          eventValue: [SIDE_EVENT.SAVE, {}],
          label: this.msg("", { id: "label_save" }), 
          iconKey: "Check"
        })
      )
    }

    if ((options.delete !== false) && current.form && (current.form.id !== null)) {
      panels.push(
        this.itemMenu({
          id: "cmd_delete", 
          eventValue: [SIDE_EVENT.DELETE, { value: current.form }],
          label: this.msg("", { id: "label_delete" }), 
          iconKey: "Times"
        })
      )
    }

    if ((options.new !== false) && current.form && (current.form.id !== null)) {
      panels.push(
        this.itemMenu({
          id: "cmd_new", 
          eventValue: [SIDE_EVENT.CHECK, [{ type, id: null }, SIDE_EVENT.LOAD_SETTING]],
          label: this.msg("", { id: "label_new" }), 
          iconKey: "Plus"
        })
      )
    }

    if (typeof options.help !== "undefined") {
      panels.push(html`<hr id="help_sep" class="separator" />`)
      panels.push(
        this.itemMenu({
          id: "cmd_help", 
          eventValue: [SIDE_EVENT.HELP, { value: options.help }],
          label: this.msg("", { id: "label_help" }), 
          iconKey: "QuestionCircle"
        })
      )
    }

    return panels
  }

  menuItems(groupKey){
    const mnu_items = []
    if (this.auditFilter.setting[0] !== "disabled"){
      mnu_items.push(html`<div class="row full">
        ${this.itemMenu({
          id: "group_admin_group", 
          selected: (groupKey === "group_admin"),
          eventValue: [SIDE_EVENT.CHANGE, { fieldname: "group_key", value: "group_admin" }],
          label: this.msg("", { id: "title_admin" }), 
          iconKey: "ExclamationTriangle"
        })}
        ${(groupKey === "group_admin") ? html`<div class="row full panel-group" >
          ${this.itemMenu({
            id: "cmd_dbsettings",  
            eventValue: [SIDE_EVENT.LOAD_SETTING, {type: 'setting'}],
            label: this.msg("", { id: `title_dbsettings` }), 
            iconKey: "Cog", color: "rgb(var(--functional-blue))"
          })}
          ${this.itemMenu({
            id: "cmd_numberdef",  
            eventValue: [SIDE_EVENT.LOAD_SETTING, {type: 'numberdef'}],
            label: this.msg("", { id: `title_numberdef` }), 
            iconKey: "ListOl", color: "rgb(var(--functional-blue))"
          })}
          ${(this.auditFilter.audit[0] !== "disabled") ? this.itemMenu({
            id: "cmd_usergroup",  
            eventValue: [SIDE_EVENT.LOAD_SETTING, {type: 'usergroup'}],
            label: this.msg("", { id: `title_usergroup` }), 
            iconKey: "Key", color: "rgb(var(--functional-blue))"
          }) : nothing}
          ${this.itemMenu({
            id: "cmd_ui_menu",  
            eventValue: [SIDE_EVENT.LOAD_SETTING, {type: 'ui_menu'}],
            label: this.msg("", { id: `title_menucmd` }), 
            iconKey: "Share", color: "rgb(var(--functional-blue))"
          })}
          ${this.itemMenu({
            id: "cmd_log",  
            eventValue: [SIDE_EVENT.LOAD_SETTING, {type: 'log'}],
            label: this.msg("", { id: `title_log` }), 
            iconKey: "InfoCircle", color: "rgb(var(--functional-blue))"
          })}
        </div>` : nothing }
      </div>`)

      mnu_items.push(html`<div class="row full">
        ${this.itemMenu({
          id: "group_database_group", 
          selected: (groupKey === "group_database"),
          eventValue: [SIDE_EVENT.CHANGE, { fieldname: "group_key", value: "group_database" }],
          label: this.msg("", { id: "title_database" }), 
          iconKey: "Database"
        })}
        ${(groupKey === "group_database") ? html`<div class="row full panel-group" >
          ${this.itemMenu({
            id: "cmd_deffield",  
            eventValue: [SIDE_EVENT.LOAD_SETTING, {type: 'deffield'}],
            label: this.msg("", { id: `title_deffield` }), 
            iconKey: "Tag", color: "rgb(var(--functional-blue))"
          })}
          ${this.itemMenu({
            id: "cmd_groups",  
            eventValue: [SIDE_EVENT.LOAD_SETTING, {type: 'groups'}],
            label: this.msg("", { id: `title_groups` }), 
            iconKey: "Th", color: "rgb(var(--functional-blue))"
          })}
          ${this.itemMenu({
            id: "cmd_place",  
            eventValue: [SIDE_EVENT.LOAD_SETTING, {type: 'place'}],
            label: this.msg("", { id: `title_place` }), 
            iconKey: "Map", color: "rgb(var(--functional-blue))"
          })}
          ${this.itemMenu({
            id: "cmd_currency",  
            eventValue: [SIDE_EVENT.LOAD_SETTING, {type: 'currency'}],
            label: this.msg("", { id: `title_currency` }), 
            iconKey: "Dollar", color: "rgb(var(--functional-blue))"
          })}
          ${this.itemMenu({
            id: "cmd_tax",  
            eventValue: [SIDE_EVENT.LOAD_SETTING, {type: 'tax'}],
            label: this.msg("", { id: `title_tax` }), 
            iconKey: "Ticket", color: "rgb(var(--functional-blue))"
          })}
          ${this.itemMenu({
            id: "cmd_company",  
            eventValue: [SIDE_EVENT.CHECK, [{ ntype: "customer", ttype: null, id: 1 }]],
            label: this.msg("", { id: `title_company` }), 
            iconKey: "Home", color: "rgb(var(--functional-blue))"
          })}
          ${this.itemMenu({
            id: "cmd_template",  
            eventValue: [SIDE_EVENT.LOAD_SETTING, {type: 'template'}],
            label: this.msg("", { id: `title_report_editor` }), 
            iconKey: "TextHeight", color: "rgb(var(--functional-blue))"
          })}
      </div>` : nothing }
    </div>`)
    }

    mnu_items.push(html`<div class="row full">
        ${this.itemMenu({
          id: "group_user_group", 
          selected: (groupKey === "group_user"),
          eventValue: [SIDE_EVENT.CHANGE, { fieldname: "group_key", value: "group_user" }],
          label: this.msg("", { id: "title_user" }), 
          iconKey: "Desktop"
        })}
        ${(groupKey === "group_user") ? html`<div class="row full panel-group" >
          ${this.itemMenu({
            id: "cmd_program",  
            eventValue: [SIDE_EVENT.PROGRAM_SETTING, {}],
            label: this.msg("", { id: `title_program` }), 
            iconKey: "Keyboard", color: "rgb(var(--functional-blue))"
          })}
          ${this.itemMenu({
            id: "cmd_password",  
            eventValue: [SIDE_EVENT.PASSWORD, {username: this.username}],
            label: this.msg("", { id: `title_password` }), 
            iconKey: "Lock", color: "rgb(var(--functional-blue))"
          })}
      </div>` : nothing }
    </div>`)

    return mnu_items
  }

  render() {
    const { group_key, current, panel } = this.module
    return html`<div class="sidebar ${(this.side !== "auto") ? this.side : ""}" >
    ${(current && panel) ? this.formItems(panel) : this.menuItems(group_key)}
    </div>`
  }

}