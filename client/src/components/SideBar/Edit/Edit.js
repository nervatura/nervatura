import { LitElement, html, nothing } from 'lit';
import { styleMap } from 'lit/directives/style-map.js';

import '../../Form/Icon/form-icon.js'
import '../../Form/Button/form-button.js'

import { styles } from './Edit.styles.js'
import { SIDE_VISIBILITY, SIDE_EVENT, BUTTON_TYPE, SIDE_VIEW, TEXT_ALIGN, EDITOR_EVENT } from '../../../config/enums.js'

export class Edit extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.side = SIDE_VISIBILITY.AUTO
    this.view = SIDE_VIEW.EDIT
    this.module = {
      current: {}, 
      form_dirty: false,
      dirty: false,
      panel: {}, 
      dataset: {}, 
      group_key: ""
    }
    this.newFilter = []
    this.auditFilter = {} 
    this.forms = {}
  }

  static get properties() {
    return {
      side: { type: String, reflect: true },
      view: { type: String, reflect: true },
      newFilter: { type: Array }, 
      auditFilter: { type: Object },  
      module: { type: Object }, 
      forms: { type: Object }
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

  itemMenu({id, selected, eventValue, label, iconKey, full, disabled, align, color}){
    const btnSelected = (typeof(selected) === "undefined") ? false : selected
    const btnFull = (typeof(full) === "undefined") ? true : full
    const btnDisabled = (typeof(disabled) === "undefined") ? false : disabled
    const btnAlign = (typeof(align) === "undefined") ? TEXT_ALIGN.LEFT : align
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
        ?full="${btnFull}" ?disabled="${btnDisabled}" ?selected="${btnSelected}"
        align=${btnAlign}
        .style="${style}"
        icon="${iconKey}" type="${BUTTON_TYPE.PRIMARY}"
        @click=${()=>this._onSideEvent( ...eventValue )} 
      >${label}</form-button>`
    )
  }

  editItems(_options){
    const { current, dirty, form_dirty, dataset } = this.module
    const options = (typeof _options === "undefined") ? {} : _options
    const panels = []

    if (options.back === true || current.form) {
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
    }

    if (options.arrow === true) {
      panels.push(
        this.itemMenu({
          id: "cmd_arrow_left", 
          eventValue: [SIDE_EVENT.PREV_NUMBER, {}],
          label: this.msg("", { id: "label_previous" }), 
          iconKey: "ArrowLeft"
        })
      )
      panels.push(
        this.itemMenu({
          id: "cmd_arrow_right", 
          eventValue: [SIDE_EVENT.NEXT_NUMBER, {}],
          label: this.msg("", { id: "label_next" }), 
          iconKey: "ArrowRight", align: TEXT_ALIGN.RIGHT
        })
      )
      panels.push(html`<hr id="arrow_sep" class="separator" />`)
    }

    if (options.state && options.state !== "normal") {
      const color = (options.state === "deleted") 
        ? "rgb(var(--functional-red))" : (options.state === "cancellation") 
          ? "rgb(var(--functional-yellow))" : "rgba(var(--accent-1c), 0.85)"
      const icon = (["closed", "readonly"].includes(options.state)) 
        ? "Lock" : "ExclamationTriangle"
      panels.push(html`<div key="cmd_state" class="state-label" >
        <form-icon iconKey="${icon}" .style="${{ fill: color }}" ></form-icon>
        <span style="${styleMap({ color, fill: color, "vertical-align": "middle" })}" >${this.msg("", { id: `label_${options.state}` })}</span>
      </div>`)
      panels.push(html`<hr id="state_sep" class="separator" />`)
    }

    if (options.save !== false) {
      panels.push(
        this.itemMenu({
          id: "cmd_save",
          selected: !!(((current.form && form_dirty)||(!current.form && dirty))),
          eventValue: [SIDE_EVENT.SAVE, {}],
          label: this.msg("", { id: "label_save" }), 
          iconKey: "Check"
        })
      )
    }
    if (options.delete !== false && options.state === "normal") {
      panels.push(
        this.itemMenu({
          id: "cmd_delete", 
          eventValue: [SIDE_EVENT.DELETE, {}],
          label: this.msg("", { id: "label_delete" }), 
          iconKey: "Times"
        })
      )
    }
    if (options.new !== false && options.state === "normal" && !current.form) {
      panels.push(
        this.itemMenu({
          id: "cmd_new", 
          eventValue: [SIDE_EVENT.NEW, [{}]],
          label: this.msg("", { id: "label_new" }), 
          iconKey: "Plus"
        })
      )
    }

    if (options.trans === true) {
      panels.push(html`<hr id="trans_sep" class="separator" />`)
      if (options.copy !== false) {
        panels.push(
          this.itemMenu({
            id: "cmd_copy",
            eventValue: [SIDE_EVENT.COPY, { value: "normal" }],
            label: this.msg("", { id: "label_copy" }),
            iconKey: "Copy"
          })
        );
      }
      if (options.create !== false) {
        panels.push(
          this.itemMenu({
            id: "cmd_create",
            eventValue: [SIDE_EVENT.COPY, { value: "create" }],
            label: this.msg("", { id: "label_create" }), 
            iconKey: "Sitemap"
          })
        );
      }
      if (options.corrective === true && options.state === "normal") {
        panels.push(
          this.itemMenu({
            id: "cmd_corrective",
            eventValue: [SIDE_EVENT.COPY, { value: "amendment" }],
            label: this.msg("", { id: "label_corrective" }), 
            iconKey: "Share"
          })
        );
      }
      if (options.cancellation === true && options.state !== "cancellation") {
        panels.push(
          this.itemMenu({
            id: "cmd_cancellation",
            eventValue: [SIDE_EVENT.COPY, { value: "cancellation" }],
            label: this.msg("", { id: "label_cancellation" }), 
            iconKey: "Undo"
          })
        );
      }
      if (options.formula === true) {
        panels.push(
          this.itemMenu({
            id: "cmd_formula", 
            eventValue: [SIDE_EVENT.CHECK, [{}, EDITOR_EVENT.LOAD_FORMULA]],
            label: this.msg("", { id: "label_formula" }), 
            iconKey: "Magic"
          })    
        )
      }
    }

    if (options.link === true) {
      panels.push(
        this.itemMenu({
          id: "cmd_link", 
          eventValue: [SIDE_EVENT.LINK, { type: options.link_type, field: options.link_field }],
          label: options.link_label,
          iconKey: "Link"
        })
      )
    }

    if (options.password === true) {
      panels.push(
        this.itemMenu({
          id: "cmd_password", 
          eventValue: [SIDE_EVENT.PASSWORD, {}],
          label: this.msg("", { id: "title_password" }), 
          iconKey: "Lock"
        })
      )
    }

    if (options.shipping === true) {
      panels.push(
        this.itemMenu({
          id: "cmd_shipping_all", 
          eventValue: [SIDE_EVENT.SHIPPING_ADD_ALL, {}],
          label: this.msg("", { id: "shipping_all_label" }), 
          iconKey: "Plus"
        })
      )
      panels.push(
        this.itemMenu({
          id: "cmd_shipping_create",
          selected: (dataset.shiptemp && (dataset.shiptemp.length > 0)),
          eventValue: [SIDE_EVENT.SHIPPING_CREATE, {}],
          label: this.msg("", { id: "shipping_create_label" }), 
          iconKey: "Check"
        })
      )
    }

    if (options.more === true) {
      panels.push(html`<hr id="more_sep_1" class="separator" />`)
      if (options.report !== false) {
        panels.push(
          this.itemMenu({
            id: "cmd_report",
            eventValue: [SIDE_EVENT.REPORT_SETTINGS, {}],
            label: this.msg("", { id: "label_report" }), 
            iconKey: "ChartBar"
          })
        )
      }
      if (options.search === true) {
        panels.push(
          this.itemMenu({
            id: "cmd_search", 
            eventValue: [SIDE_EVENT.SEARCH_QUEUE, {}],
            label: this.msg("", { id: "label_search" }), 
            iconKey: "Search"
          })
        )
      }
      if (options.export_all === true && options.state === "normal") {
        panels.push(
          this.itemMenu({
            id: "cmd_export_all", 
            eventValue: [SIDE_EVENT.EXPORT_QUEUE_ALL, {}],
            label: this.msg("", { id: "label_export_all" }), 
            iconKey: "Download"
          })
        )
      }
      if (options.print === true) {
        panels.push(
          this.itemMenu({
            id: "cmd_print", 
            eventValue: [SIDE_EVENT.CREATE_REPORT, { value: "print" }],
            label: this.msg("", { id: "label_print" }), 
            iconKey: "Print"
          })
        )
      }
      if (options.export_pdf === true && options.state === "normal") {
        panels.push(
          this.itemMenu({
            id: "cmd_export_pdf", 
            eventValue: [SIDE_EVENT.CREATE_REPORT, { value: "pdf" }],
            label: this.msg("", { id: "label_export_pdf" }), 
            iconKey: "Download"
          })
        )
      }
      if (options.export_xml === true && options.state === "normal") {
        panels.push(
          this.itemMenu({
            id: "cmd_export_xml", 
            eventValue: [SIDE_EVENT.CREATE_REPORT, { value: "xml" }],
            label: this.msg("", { id: "label_export_xml" }), 
            iconKey: "Code"
          })
        )
      }
      if (options.export_csv === true && options.state === "normal") {
        panels.push(
          this.itemMenu({
            id: "cmd_export_csv", 
            eventValue: [SIDE_EVENT.CREATE_REPORT, { value: "csv" }],
            label: this.msg("", { id: "label_export_csv" }), 
            iconKey: "Download"
          })
        )
      }
      if (options.export_event === true && options.state === "normal") {
        panels.push(
          this.itemMenu({
            id: "cmd_export_event", 
            eventValue: [SIDE_EVENT.EXPORT_EVENT, {}],
            label: this.msg("", { id: "label_export_event" }), 
            iconKey: "Calendar"
          })
        )
      }

      panels.push(html`<hr id="more_sep_2" class="separator" />`)
      if (options.bookmark !== false && options.state === "normal") {
        panels.push(
          this.itemMenu({
            id: "cmd_bookmark",
            eventValue: [SIDE_EVENT.SAVE_BOOKMARK, { value: options.bookmark }],
            label: this.msg("", { id: "label_bookmark" }), 
            iconKey: "Star"
          })
        )
      }
      if (options.help !== false) {
        panels.push(
          this.itemMenu({
            id: "cmd_help", 
            eventValue: [SIDE_EVENT.HELP, { value: options.help }],
            label: this.msg("", { id: "label_help" }), 
            iconKey: "QuestionCircle"
          })
        )
      }
    }

    if (options.more !== true && typeof options.help !== "undefined") {
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

  newItems(){
    const { group_key } = this.module
    const mnu_items = []

    if(this.newFilter[0].length > 0){
      mnu_items.push(html`<div class="row full">
        ${this.itemMenu({
          id: "new_transitem_group", 
          selected: (group_key === "new_transitem"),
          eventValue: [SIDE_EVENT.CHANGE, { fieldname: "group_key", value: "new_transitem" }],
          label: this.msg("", { id: "search_transitem" }), 
          iconKey: "FileText"
        })}
        ${(group_key === "new_transitem") ? html`<div class="row full panel-group" >
          ${this.newFilter[0].map(transtype =>{
            if (this.auditFilter.trans[transtype][0] === "all"){ 
              return (
                this.itemMenu({
                  id: transtype,  
                  eventValue: [SIDE_EVENT.NEW, [{ntype: 'trans', ttype: transtype}]],
                  label: this.msg("", { id: `title_${transtype}` }), 
                  iconKey: "FileText", color: "rgb(var(--functional-blue))"
                })
              ) 
            } 
            return nothing
          })}
        </div>` : nothing }
      </div>`)
    }

    if(this.newFilter[1].length > 0){
      mnu_items.push(html`<div class="row full">
        ${this.itemMenu({
          id: "new_transpayment_group",
          selected: (group_key === "new_transpayment"),
          eventValue: [SIDE_EVENT.CHANGE, { fieldname: "group_key", value: "new_transpayment" }],
          label: this.msg("", { id: "search_transpayment" }), 
          iconKey: "Money"
        })}
        ${(group_key === "new_transpayment") ? html`<div class="row full panel-group" >
          ${this.newFilter[1].map(transtype =>{
            if (this.auditFilter.trans[transtype][0] === "all"){ 
              return (
                this.itemMenu({
                  id: transtype,  
                  eventValue: [SIDE_EVENT.NEW, [{ntype: 'trans', ttype: transtype}]],
                  label: this.msg("", { id: `title_${transtype}` }), 
                  iconKey: "Money", color: "rgb(var(--functional-blue))"
                })
              ) 
            }
            return nothing
          })}
        </div>` : nothing }
      </div>`)
    }

    if(this.newFilter[2].length > 0){
      mnu_items.push(html`<div class="row full">
        ${this.itemMenu({
          id: "new_transmovement_group",
          selected: (group_key === "new_transmovement"),
          eventValue: [SIDE_EVENT.CHANGE, { fieldname: "group_key", value: "new_transmovement" }],
          label: this.msg("", { id: "search_transmovement" }), 
          iconKey: "Truck"
        })}
        ${(group_key === "new_transmovement") ? html`<div class="row full panel-group" >
          ${this.newFilter[2].map(transtype => {
            if (this.auditFilter.trans[transtype][0] === "all"){
              if(transtype === "delivery"){
                return ([
                  this.itemMenu({
                    id: "shipping",  
                    eventValue: [SIDE_EVENT.NEW, [{ntype: 'trans', ttype: "shipping"}]],
                    label: this.msg("", { id: `title_${transtype}` }), 
                    iconKey: this.forms[transtype]().options.icon, 
                    color: "rgb(var(--functional-blue))"
                  }),
                  this.itemMenu({
                    id: transtype,  
                    eventValue: [SIDE_EVENT.NEW, [{ntype: 'trans', ttype: transtype}]],
                    label: this.msg("", { id: "title_transfer" }), 
                    iconKey: this.forms[transtype]().options.icon, 
                    color: "rgb(var(--functional-blue))"
                  })
                ])
              }
              return (
                this.itemMenu({
                  id: transtype, 
                  eventValue: [SIDE_EVENT.NEW, [{ntype: 'trans', ttype: transtype}]],
                  label: this.msg("", { id: `title_${transtype}` }), 
                  iconKey: this.forms[transtype]().options.icon, 
                  color: "rgb(var(--functional-blue))"
                })
              )
            } 
            return nothing
          })}
        </div>` : nothing }
      </div>`)
    }

    if(this.newFilter[3].length > 0){
      mnu_items.push(html`<div class="row full">
        ${this.itemMenu({
          id: "new_resources_group",
          selected: (group_key === "new_resources"),
          eventValue: [SIDE_EVENT.CHANGE, { fieldname: "group_key", value: "new_resources" }],
          label: this.msg("", { id: "title_resources" }), 
          iconKey: "Wrench"
        })}
        ${(group_key === "new_resources") ? html`<div class="row full panel-group" >
          ${this.newFilter[3].map(ntype =>{
            if (this.auditFilter[ntype][0] === "all"){ 
              return (
                this.itemMenu({
                  id: ntype,  
                  eventValue: [SIDE_EVENT.NEW, [{ntype, ttype: null}]],
                  label: this.msg("", { id: `title_${ntype}` }), 
                  iconKey: this.forms[ntype]().options.icon,
                  color: "rgb(var(--functional-blue))"
                })
              ) 
            } 
            return nothing
          })}
        </div>` : nothing }
      </div>`)
    }

    return mnu_items
  }

  render() {
    const { current, panel } = this.module
    return html`<div class="sidebar ${(this.side !== "auto") ? this.side : ""}" >
      ${(!current.form && (current.form_type !== "transitem_shipping")) ?
        html`<div class="row full container">
          <div class="cell half">
            ${this.itemMenu({
              id: "state_new",
              selected: !(((this.view === SIDE_VIEW.EDIT) && current.item)),
              eventValue: [SIDE_EVENT.CHANGE, { fieldname: "side_view", value: "new" }],
              label: this.msg("", { id: "label_new" }), 
              iconKey: "Plus", align: TEXT_ALIGN.CENTER
            })}
          </div>
          <div class="cell half">
            ${this.itemMenu({
              id: "state_edit", 
              selected: !!(((this.view === SIDE_VIEW.EDIT) && current.item)),
              eventValue: [SIDE_EVENT.CHANGE, { fieldname: "side_view", value: "edit" }],
              label: this.msg("", { id: "label_edit"}), 
              iconKey: "Edit", disabled: (!current.item), align: TEXT_ALIGN.CENTER
            })}
          </div>
        </div>` : nothing
      }
      ${(((this.view === SIDE_VIEW.EDIT) && current.form) || ((this.view === SIDE_VIEW.EDIT) && current.item))
        ? this.editItems(panel) : this.newItems()}
     </div>`
  }

}