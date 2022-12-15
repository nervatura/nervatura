import { LitElement, html, nothing } from 'lit';

import '../../Form/Button/form-button.js'

import { styles } from './Search.styles.js'
import { SIDE_VISIBILITY, SIDE_EVENT, BUTTON_TYPE, TEXT_ALIGN } from '../../../config/enums.js'

export class Search extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.side = SIDE_VISIBILITY.AUTO
    this.groupKey = ""
    this.auditFilter = {}
  }

  static get properties() {
    return {
      side: { type: String, reflect: true },
      groupKey: { type: String },
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

  getButtonStyle(stype, key) {
    if(stype === "group"){
      return { 
        "text-align": "left", "border-radius": "0", 
         "color": (key === this.groupKey) ? "rgb(var(--functional-yellow))" : "",
         "fill": (key === this.groupKey) ? "rgb(var(--functional-yellow))" : "",
         "border-color": "rgba(var(--accent-1c), 0.2)"
      }
    }
    return { 
      "text-align": "left", "border-radius": "0",
      "color": "rgb(var(--functional-blue))",
      "fill": "rgb(var(--functional-blue))",
      "border-color": "rgba(var(--accent-1c), 0.2)"
    }
  }

  searchGroup(key) {
    return(
      html`<div class="row full">
        <form-button id="${`btn_group_${key}`}" 
          label="${this.msg(``, { id: `search_${key}` })}"
          ?full="${true}" align=${TEXT_ALIGN.LEFT}
          .style="${this.getButtonStyle("group", key)}"
          icon="FileText" type="${BUTTON_TYPE.PRIMARY}"
          @click=${()=>this._onSideEvent( SIDE_EVENT.CHANGE, { fieldname: "group_key", value: key } )} 
        >${this.msg(``, { id: `search_${key}` })}</form-button>
        ${(this.groupKey === key) ? html`<div class="row full panel-group" >
          <form-button id="${`btn_view_${key}`}" 
            label="${this.msg("", { id: "quick_search" })}"
            ?full="${true}" align=${TEXT_ALIGN.LEFT}
            .style="${this.getButtonStyle("panel")}"
            icon="Bolt" type="${BUTTON_TYPE.PRIMARY}"
            @click=${()=>this._onSideEvent( SIDE_EVENT.QUICK, { value: key } )} 
          >${this.msg("Quick Search", { id: "quick_search" })}</form-button>
          <form-button id="${`btn_browser_${key}`}" 
            label="${this.msg(``, { id: `browser_${key}` })}"
            ?full="${true}" align=${TEXT_ALIGN.LEFT}
            .style="${this.getButtonStyle("panel")}"
            icon="Search" type="${BUTTON_TYPE.PRIMARY}"
            @click=${()=>this._onSideEvent( SIDE_EVENT.BROWSER, { value: key } )} 
          >${this.msg(``, { id: `browser_${key}` })}</form-button>
        </div>` : nothing}
      </div>`
    )
  }

  render() {
    return html`<div class="sidebar ${(this.side !== "auto") ? this.side : ""}" >
      ${this.searchGroup("transitem")}
      ${((this.auditFilter.trans.bank[0] !== "disabled") || (this.auditFilter.trans.cash[0] !== "disabled"))?
        this.searchGroup("transpayment") : nothing}
      ${((this.auditFilter.trans.delivery[0] !== "disabled") || (this.auditFilter.trans.inventory[0] !== "disabled") 
        || (this.auditFilter.trans.waybill[0] !== "disabled") || (this.auditFilter.trans.production[0] !== "disabled")
        || (this.auditFilter.trans.formula[0] !== "disabled"))?
        this.searchGroup("transmovement") : nothing}
      
      <hr class="separator" />
      ${["customer","product","employee","tool","project"].map(key => {
        if(this.auditFilter[key][0] !== "disabled") {
          return this.searchGroup(key)}
        return nothing
      })}

      <hr class="separator" />
      <form-button id="btn_report" 
        label="${this.msg(``, { id: `search_report` })}"
        ?full="${true}" align=${TEXT_ALIGN.LEFT}
        .style="${this.getButtonStyle("group", "report")}"
        icon="ChartBar" type="${BUTTON_TYPE.PRIMARY}"
        @click=${()=>{
          this._onSideEvent( SIDE_EVENT.CHANGE, { fieldname: "group_key", value: "report" } )
          this._onSideEvent( SIDE_EVENT.QUICK, { value: "report" } )
        }} 
      >${this.msg(``, { id: `search_report` })}</form-button>
      <form-button id="btn_office" 
        label="${this.msg(``, { id: `search_office` })}"
        ?full="${true}" align=${TEXT_ALIGN.LEFT}
        .style="${this.getButtonStyle("group", "office")}"
        icon="Inbox" type="${BUTTON_TYPE.PRIMARY}"
        @click=${()=>this._onSideEvent( SIDE_EVENT.CHANGE, { fieldname: "group_key", value: "office" } )} 
      >${this.msg(``, { id: `search_office` })}</form-button>
      ${(this.groupKey === "office") ? html`<div class="row full panel-group" >
        <form-button id="btn_printqueue" 
          label="${this.msg("", { id: "title_printqueue" })}"
          ?full="${true}" align=${TEXT_ALIGN.LEFT}
          .style="${this.getButtonStyle("panel")}"
          icon="Print" type="${BUTTON_TYPE.PRIMARY}"
          @click=${()=>this._onSideEvent( SIDE_EVENT.CHECK, { ntype: "printqueue", ttype: null, id: null } )} 
        >${this.msg("", { id: "title_printqueue" })}</form-button>
        <form-button id="btn_rate" 
          label="${this.msg("", { id: "title_rate" })}"
          ?full="${true}" align=${TEXT_ALIGN.LEFT}
          .style="${this.getButtonStyle("panel")}"
          icon="Globe" type="${BUTTON_TYPE.PRIMARY}"
          @click=${()=>this._onSideEvent( SIDE_EVENT.BROWSER, { value: "rate" } )} 
        >${this.msg("", { id: "title_rate" })}</form-button>
        <form-button id="btn_servercmd" 
          label="${this.msg("", { id: "title_servercmd" })}"
          ?full="${true}" align=${TEXT_ALIGN.LEFT}
          .style="${this.getButtonStyle("panel")}"
          icon="Share" type="${BUTTON_TYPE.PRIMARY}"
          @click=${()=>this._onSideEvent( SIDE_EVENT.QUICK, { value: "servercmd" } )} 
        >${this.msg("", { id: "title_servercmd" })}</form-button>
      </div>` : nothing}
    </div>`
  }

}