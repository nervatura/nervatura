import { LitElement, html, nothing } from 'lit';
import { styleMap } from 'lit/directives/style-map.js';
import { ifDefined } from 'lit/directives/if-defined.js';

import '../Input/form-input.js'
import '../NumberInput/form-number.js'
import '../Button/form-button.js'
import '../Icon/form-icon.js'
import '../Select/form-select.js'
import '../DateTime/form-datetime.js'

import { styles } from './Field.styles.js'
import { BUTTON_TYPE, DATETIME_TYPE, INPUT_TYPE } from '../../../config/enums.js'

export class Field extends LitElement {
  constructor() {
    super();
    this.id = Math.random().toString(36).slice(2);
    this.field = {};
    this.values = {};
    this.options = {};
    this.data = {
      dataset: {}, 
      current: {}, 
      audit: "all",
    };
    this.style = {};
    this.msg = (defValue) => defValue
  }

  static get properties() {
    return {
      id: { type: String },
      field: { type: Object },
      values: { type: Object },
      options: { type: Object },
      data: { type: Object },
      style: { type: Object },
    };
  }

  static get styles () {
    return [
      styles
    ]
  }

  _onChange({ value, item, event_type, editName, fieldMap }) {
    const props = {
      id: this.field.id || 1,
      name: editName,
      event_type,
      value, 
      extend: !!((fieldMap && fieldMap.extend)),
      refnumber: (item && item.label) ? item.label : this.field.link_label,
      label_field: (fieldMap) ? fieldMap.label_field : undefined,
      item
    }
    if(this.onEdit){
      this.onEdit(props)
    }
    this.dispatchEvent(
      new CustomEvent('change', {
        bubbles: true, composed: true,
        detail: {
          ...props
        }
      })
    );
  }

  _onEvent(key, data){
    if(this.onEvent){
      this.onEvent(key, data)
    }
    this.dispatchEvent(
      new CustomEvent('event', {
        bubbles: true, composed: true,
        detail: {
          key, data
        }
      })
    );
  }

  _onSelector(type, filter, callback){
    if(this.onSelector){
      this.onSelector(type, filter, callback)
    }
    this.dispatchEvent(
      new CustomEvent('selector', {
        bubbles: true, composed: true,
        detail: {
          type, filter, callback
        }
      })
    );
  }

  lnkValue({fieldMap, fieldName, value}) {
    if (typeof this.values[this.field.name] === "undefined") {
      return [(this.data.current[fieldMap.source]) ?
        this.data.current[fieldMap.source].filter(item => 
          (item.ref_id === this.values.id) && (item[fieldMap.value] === this.field.name))[0] :
        this.data.dataset[fieldMap.source].filter(item => 
          (item.ref_id === this.values.id) && (item[fieldMap.value] === this.field.name))[0], false]
    }
    const svalue = ((fieldName === "id") && (value === "")) ? null : value
    return [(this.data.current[fieldMap.source]) ?
      this.data.current[fieldMap.source].filter(item => 
        (item[fieldMap.value] === svalue))[0] :
      this.data.dataset[fieldMap.source].filter(item => 
        (item[fieldMap.value] === svalue))[0], true]
  }

  getOppositeValue(value) {
    if (this.options.opposite && (parseFloat(value)<0)) {
      return String(value).replace("-","");
    }
    if (this.options.opposite && (parseFloat(value)>0)) {
      return `-${value}`;
    }
    return value;  
  }

  selectorInit({ fieldMap, value, editName }) {
    let selector = {
      value: this.values.value || "",
      filter: this.field.description || "",
      text: this.field.description || "",
      type: this.field.fieldtype,
      ntype: (["transitem", "transmovement", "transpayment"].includes(this.field.fieldtype)) ? "trans" : this.field.fieldtype,
      ttype: null,
      id: this.values.value || null,
      table: { name:"fieldvalue", fieldname: "value", id: this.field.id },
      fieldMap, editName
    }
    if (fieldMap) {
      selector = {
        ...selector,
        value: value || "",
        type: fieldMap.seltype, 
        filter: String(selector.text).split(" | ")[0],
        table:{ name: fieldMap.table, fieldname: fieldMap.fieldname },
        ntype: fieldMap.lnktype, 
        ttype: (value !== "") 
          ? (fieldMap.transtype === "") 
            ? this.values.transtype 
            : fieldMap.transtype 
          : selector.ttype, 
        id: (value !== "") ? value : selector.id
      }
      let reftable;
      if (fieldMap.extend === true || fieldMap.table === "extend") {
        reftable = this.data.current.extend
      } else {
        let _reftable;
        if((typeof this.data.current[fieldMap.table] !== "undefined") && Array.isArray(this.data.current[fieldMap.table])){
          _reftable = this.data.current[fieldMap.table].filter(item => (item.id === this.values.id))
        } else {
          _reftable = this.data.dataset[fieldMap.table].filter(item => (item.id === this.values.id))
        }
        reftable = _reftable[0]
      }
      if (typeof reftable !== "undefined") {
        if (typeof this.values[fieldMap.label_field] !== "undefined") {
          selector = {
            ...selector,
            text: this.values[fieldMap.label_field]||"", 
            filter: this.values[fieldMap.label_field]||""
          }
        } else if (typeof reftable[fieldMap.label_field] !== "undefined" && 
          reftable[fieldMap.label_field] !== null) {
            selector = {
              ...selector,
              text: reftable[fieldMap.label_field], 
              filter: reftable[fieldMap.label_field]
            }
        } else {
          selector = {
            ...selector,
            text: "", 
            filter: ""
          }
        }
        if (typeof reftable[fieldMap.fieldname] !== "undefined" && 
          selector.value==="" && reftable[fieldMap.fieldname] !== null) {
            selector = {
              ...selector,
              ntype: fieldMap.lnktype, 
              ttype: fieldMap.transtype, 
              id: reftable[fieldMap.fieldname]
            }
        }
        if (fieldMap.lnktype === "trans" && typeof reftable.transtype !== "undefined") {
          if (typeof fieldMap.lnkid !== "undefined") {
            selector = {
              ...selector,
              id: reftable[fieldMap.lnkid]
            }
          } else if (typeof reftable[fieldMap.fieldname] !== "undefined") {
            selector = {
              ...selector,
              id: reftable[fieldMap.fieldname]
            }
          } 
          else {
            selector = {
              ...selector,
              id: selector.value
            }
          }
        }
      } else {
        selector = {
          ...selector,
          text: "", 
          filter: ""
        }
      }
    }
    this.selector = selector
    return selector
  }

  setSelector(row, filter) {
    let selector = {
      ...this.selector,
      text: "",
      id: null,
      filter: filter || ""
    }
    if (row){
      const params = row.id.split("/")
      selector = {
        ...selector,
        text: row.label || row.item.lslabel
      }
      selector = {
        ...selector,
        id: parseInt(params[2],10),
        ttype: params[1]
      }
      if((params[0] === "trans") && (params[1] !== "")){
        if(row.trans_id){
          selector = {
            ...selector,
            id: row.trans_id
          }
        }
      }
    }
    selector = {
      ...selector,
      value: selector.id || ""
    }
    this.selector = selector
    this._onChange({value: selector.id, item: row, event_type: "change", 
      editName: selector.editName, fieldMap: selector.fieldMap})
  }

  editName(){
    return (this.field.map)
      ? (this.field.map.extend && this.field.map.text) 
        ? this.field.map.text 
        : (this.field.map.fieldname) 
          ? this.field.map.fieldname : this.field.name
      : this.field.name
  }

  _onTextInput(e){
    this._onChange({
      value: e.target.value, event_type: "change", 
      editName: this.editName(), fieldMap: this.field.map 
    })
  }

  render() {
    let { datatype } = this.field
    let disabled = (this.field.disabled || this.data.audit === 'readonly')
    let fieldName = this.field.name;
    let value = this.values[this.field.name]
    const fieldMap = this.field.map || null
    const editName = this.editName()
    const empty = !!(((this.field.empty === "true") || (this.field.empty === true)))
    if((this.field.rowtype === "reportfield") || (this.field.rowtype === "fieldvalue")){
      value = this.values.value
    }
    if ((typeof value==="undefined") || value === null){
      value = (this.field.default) ? this.field.default : ""
    }
    if (datatype === "fieldvalue"){
      datatype = this.values.datatype
    }

    switch (datatype) {
      case "password":
      case "color":
        return html`<form-input 
          id="${this.id}" name="${fieldName}" type="${datatype}" 
          value="${value||""}" label="${fieldName}"
          ?full="${true}"
          .style="${this.style}" 
          ?disabled="${disabled}"
          .onChange=${
            (event) => this._onChange({value: event.value, event_type: "change", editName, fieldMap}) 
          }></form-input>`

      case "date":
      case "datetime":
        if (fieldMap) {
          if (fieldMap.extend) {
            value = this.data.current.extend[fieldMap.text]
            fieldName = fieldMap.text;
          } else {
            const lnkDate = this.lnkValue({fieldMap, fieldName, value})
            if (typeof lnkDate[0] !== "undefined") {
              value = lnkDate[0][fieldMap.text]
              disabled = (lnkDate[1]) ? lnkDate[1] : disabled
            }
          }
        }
        return html`<form-datetime id="${this.id}"
          name="${fieldName}" label="${fieldName}"
          .style="${this.style}" .isnull="${empty}"
          type="${(datatype === "datetime") ? DATETIME_TYPE.DATETIME : DATETIME_TYPE.DATE}"
          .value="${value}"
          .onChange="${(date) => {
            this._onChange({value: date.value, event_type: "change", editName, fieldMap})
          }}"
          ?disabled="${disabled}"></form-datetime>`

      case "bool":
      case "flip":
        const toggleDisabled = (disabled) ? "toggle-disabled" : ""
        if([1,"1","true",true].includes(value)){
          return html`<div id="${this.id}"
            name="${fieldName}" style="${styleMap(this.style)}" 
            class="${`toggle toggle-on ${toggleDisabled}`}"
            @click="${(!disabled)?
              ()=>this._onChange({
                value: (this.field.name === 'fieldvalue_value') ? false : 0,
                event_type: "change", editName, fieldMap
              }):null}">
            <form-icon iconKey="ToggleOn" width=40 height=32.6 ></form-icon>
          </div>`
        }
        return html`<div id="${this.id}"
          name="${fieldName}" style="${styleMap(this.style)}" 
          class="${`toggle toggle-off ${toggleDisabled}`}"
          @click="${(!disabled)?
            ()=>this._onChange({
              value: (this.field.name === 'fieldvalue_value') ? true : 1,
              event_type: "change", editName, fieldMap
            }):null}">
          <form-icon iconKey="ToggleOff" width=40 height=32.6 ></form-icon>
        </div>`

      case "label":
        return null;

      case "select":
        if (this.field.extend) {
          value = this.data.current.extend[this.field.name]||"";
        }
        const selectOptions = []
        if (fieldMap) {
          this.data.dataset[fieldMap.source].forEach((element) => {
            let _label = element[fieldMap.text]
            if (typeof fieldMap.label !== "undefined") {
              _label = this.msg(`${fieldMap.label}_${_label}`, { id: `${fieldMap.label}_${_label}` });
            }
            selectOptions.push({ value: String(element[fieldMap.value]), text: _label })
          });
        } else {
          this.field.options.forEach((element) => {
            let _label = element[1]
            if(this.msg(_label, { id: _label })){
              _label = this.msg("", { id: _label })
            }
            if (typeof this.field.olabel !== "undefined") {
              _label = this.msg(`${this.field.olabel}_${element[1]}`, { id: `${this.field.olabel}_${element[1]}` });
            }
            selectOptions.push({ value: String(element[0]), text: _label })
          });
        }
        return html`<form-select id="${this.id}" ?full="${true}"
            name="${fieldName}" label="${fieldName}"
            .style="${this.style}"
            ?disabled="${disabled}" 
            .onChange=${(v) => {
              const _value = Number.isNaN(parseInt(v.value,10)) ? v.value : parseInt(v.value,10)
              this._onChange({value: _value, event_type: "change", editName, fieldMap})
            }}
            .options=${selectOptions} 
            .isnull="${empty}" value="${value}" ></form-select>`

      case "valuelist":
        return html`<form-select id="${this.id}" ?full="${true}"
          name="${fieldName}" label="${fieldName}"
          .style="${this.style}"
          ?disabled="${disabled}"
          .onChange=${
            (v) => this._onChange({value: v.value, event_type: "change", editName, fieldMap}) 
          }
          .options=${this.field.description.map((v) => ({ value: v, text: v }))} 
          .isnull="${false}" value="${value}" ></form-select>`

      case "link":
        let litem = this.values;
        const lnkLink = this.lnkValue({fieldMap, fieldName, value})
        if (typeof lnkLink[0] !== "undefined") {
          litem = lnkLink[0]
          if(lnkLink[0][fieldMap.text]){
            value = lnkLink[0][fieldMap.text];
          }
        }
        let llabel = value;
        if (typeof fieldMap.label_field !== "undefined") {
          if (typeof litem[fieldMap.label_field] !== "undefined") {
            llabel = litem[fieldMap.label_field];
          }
        }
        return html`<div 
          name="${fieldName}" style="${styleMap(this.style)}" class="link" >
          <span id=${`link_${fieldMap.lnktype}_${fieldName}`} class="link-text" 
            @click="${()=>this._onEvent("checkEditor", [
              { ntype: fieldMap.lnktype, 
                ttype: fieldMap.transtype, 
                id: value 
              }, 
              "LOAD_EDITOR"
            ])
            }" >${llabel}</span>
        </div>`

      case "selector":
        let { selector } = this
        const columns = []
        const reInit = (fieldMap && (fieldMap.extend === true || fieldMap.table === "extend") 
          && this.data.current.extend && this.data.current.extend.seltype)
        if(!selector){
          selector = this.selectorInit({ fieldMap, value, editName })
        } else if(reInit && (selector.type !== this.data.current.extend.seltype)) {
          selector.text = this.data.current.extend[fieldMap.label_field]
          selector.type = this.data.current.extend.seltype
          selector.filter = selector.text
          selector.ntype = this.data.current.extend.seltype
          selector.ttype = this.data.current.extend.transtype
          selector.id = this.data.current.extend.ref_id
        }
        if(!disabled){
          columns.push(html`<div id="sel_show" class="cell search-col">
            <form-button id="${`sel_show_${fieldName}`}" 
              label="${this.msg("", { id: "label_search" })}"
              .style="${{ padding: "5px 8px 7px" }}"
              icon="Search" type="${BUTTON_TYPE.BORDER}"
              @click=${()=>this._onSelector(selector.type, selector.filter, (...args) => this.setSelector(...args))} 
            ></form-button>
          </div>`)
        }
        if (empty) {
          columns.push(html`<div id="sel_delete" class="cell times-col">
            <form-button id="${`sel_delete_${fieldName}`}" 
              label="${this.msg("", { id: "label_delete" })}"
              .style="${{ padding: "5px 8px 7px" }}"
              ?disabled="${disabled}" icon="Times" type="${BUTTON_TYPE.BORDER}"
              @click=${(!disabled) ? ()=>this.setSelector() : null} ></form-button>
          </div>`)
        }
        columns.push(html`<div id="sel_text" class="link">
          ${(selector.text !== "") ? html`<span 
            id=${`sel_link_${fieldName}`}
            class="link-text"
            @click="${()=>this._onEvent("checkTranstype", [
              { ntype: selector.ntype, 
                ttype: selector.ttype, 
                id: selector.id }, 
              'LOAD_EDITOR'
            ])}" >${selector.text}</span>`:null}
        </div>`)
        return html`<div id="${this.id}" 
          name="${fieldName}" style="${styleMap(this.style)}"
          class="row full" >${columns}</div>`

      case "button":
        return html`<form-button id="${this.id}" 
          name="${fieldName}" type="${BUTTON_TYPE.BORDER}"
          .style="${{ padding: "7px 8px", ...this.style }}"
          ?disabled="${disabled}" 
          label="${(this.field.title)?(this.field.title):nothing}" 
          ?full="${this.field.full}"
          ?autofocus="${this.field.focus || false}"
          icon="${this.field.icon}"
          @click=${(!disabled) ? ()=>this._onChange({ 
            value: fieldName, item: {}, event_type: "click", editName, fieldMap 
          }) : null} >${(this.field.title)?(this.field.title):nothing}</form-button>`

      case "percent":
      case "integer":
      case "float":
        if (fieldMap) {
          if (fieldMap.extend) {
            value = this.data.current.extend[fieldMap.text];
            fieldName = fieldMap.text;
          } else {
            const lnkNumber = this.lnkValue({fieldMap, fieldName, value})
            if (typeof lnkNumber[0] !== "undefined") {
              value = lnkNumber[0][fieldMap.text];
              disabled = (lnkNumber[1]) ? lnkNumber[1] : disabled
            }
          }
        }
        if (value === ""){ value = 0 }
        if (typeof this.field.opposite !== "undefined") {
          value = this.getOppositeValue(value) 
        }
        return html`<form-number id="${this.id}" name="${fieldName}"
          ?integer="${!(datatype === "float")}" 
          ?full="${true}" .style="${this.style}"
          value="${value||0}" 
          ?disabled="${disabled}" label="${fieldName}"
          min="${ifDefined(this.field.min)}" 
          max="${ifDefined(this.field.max)}" 
          .onChange=${
            (event) => {
              this._onChange({
                value: (this.field.opposite) ? this.getOppositeValue(event.value): event.value, 
                event_type: "change", editName, fieldMap
              })
              this.event_type = "change"
            }
          }
          .onBlur=${
            (event) => {
              this._onChange({
                value: (this.field.opposite) ? parseFloat(this.getOppositeValue(event.value)): parseFloat(event.value),
                event_type: (this.event_type === "change") ? "blur" : null, editName, fieldMap
              })
              this.event_type = "blur"
            }
          }
        ></form-number>`

      case "notes":
      case "text":
      case "string":
      default:
        if (fieldMap) {
          if (fieldMap.extend) {
            value = this.data.current.extend[fieldMap.text];
            fieldName = fieldMap.text;
          } else {
            const lnkString = this.lnkValue({fieldMap, fieldName, value})
            if (typeof lnkString[0] !== "undefined") {
              value = lnkString[0][fieldMap.text];
              disabled = (lnkString[1]) ? lnkString[1] : disabled
              if (typeof fieldMap.label !== "undefined") {
                value = this.msg(`${fieldMap.label}_${value}`, { id: `${fieldMap.label}_${value}` });
              }
            }
          }
        }
        if((datatype === "notes") || datatype === "text"){
          return html`<textarea id="${this.id}" name="${fieldName}"
            class=${`full`} label="${fieldName}" style="${styleMap(this.style)}"
            rows=${ifDefined((this.field.rows ) ? this.field.rows: undefined)}
            @input="${this._onTextInput}"
            .value="${value||""}"
            ?disabled="${disabled}" ></textarea>`
        }
        return html`<form-input 
          id="${this.id}" .id="${`${this.id}_input`}" name="${fieldName}" 
          type="${INPUT_TYPE.TEXT}" 
          value="${value||""}" label="${fieldName}"
          ?full="${true}"
          .style="${this.style}" 
          maxlength="${ifDefined((this.field.length) ? this.field.length : undefined)}"
          size="${ifDefined((this.field.length) ? this.field.length : undefined)}"
          ?disabled="${disabled}"
          .onChange=${
            (event) => this._onChange({value: event.value, event_type: "change", editName, fieldMap}) 
          }></form-input>`
    }
  }
}