import { LitElement, html } from 'lit';
import { styleMap } from 'lit/directives/style-map.js';

import '../Input/form-input.js'
import '../Label/form-label.js'
import '../Icon/form-icon.js'
import '../Field/form-field.js'

import { styles } from './Row.styles.js'
import { INPUT_TYPE } from '../../../config/enums.js'

export class Row extends LitElement {
  constructor() {
    super();
    this.id = Math.random().toString(36).slice(2);
    this.row = {};
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
      row: { type: Object },
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

  _onEdit(props){
    if(this.onEdit){
      this.onEdit(props)
    }
    this.dispatchEvent(
      new CustomEvent('edit', {
        bubbles: true, composed: true,
        detail: {
          ...props
        }
      })
    );
  }

  _onTextInput(e){
    this._onEdit({
      id: this.row.id, 
      name: this.row.name, 
      value: e.target.value
    })
  }

  imgValue() {
    let imgValue = this.values[this.row.name] || ""
    if (imgValue!=="" && imgValue!==null) {
      if (imgValue.toString().substr(0,10)!=="data:image") {
        if (typeof this.data.dataset[imgValue]!=="undefined") {
          imgValue = this.data.dataset[imgValue]
        }
      }
    }
    return this.safeImageSrc(imgValue)
  }

  /**
   * Restricts img src to safe schemes to prevent XSS (e.g. data:image/svg+xml can execute script).
   * Allows: blob:, https:, http:, and data:image/ for non-SVG types (png, jpeg, gif, webp, ico, bmp).
   */
  safeImageSrc(value) {
    if (!value || typeof value !== "string") return ""
    const v = value.trim()
    if (v === "") return ""
    const lower = v.toLowerCase()
    if (lower.startsWith("blob:")) return v
    if (lower.startsWith("https:") || lower.startsWith("http:")) return v
    if (lower.startsWith("data:image/")) {
      const after = lower.slice(11)
      const safeTypes = ["png", "jpeg", "jpg", "gif", "webp", "ico", "bmp"]
      if (safeTypes.some((t) => after.startsWith(t + ";") || after.startsWith(t + ","))) return v
    }
    return ""
  }

  flipItem(){
    const { 
      id, name, datatype, info 
    } = this.row
    const enabled = (typeof this.values[name] !== "undefined")
    const checkbox = html`<div id="${`checkbox_${name}`}"
      class="report-field ${(enabled) ? "toggle-on" : "toggle-off"}"
      @click="${() => this._onEdit({
        id,
        selected: true,
        datatype,
        defvalue: this.row.default,
        name, 
        value: !enabled, 
        extend: false
      })}">
      ${(enabled)?
        html`<form-icon iconKey="ToggleOn" width=40 height=32.6 ></form-icon>`:
        html`<form-icon iconKey="ToggleOff" width=40 height=32.6 ></form-icon>`}
      <form-label value="${name}" class="bold padding-tiny ${(enabled) ? "toggle-on" : ""} " ></form-label>
    </div>`

    switch (datatype) {
      case "text":
        return(html`<div id="${this.id}"
          style="${styleMap(this.style)}" class="container-row">
          <div class="row full">
            <div class="cell padding-small" >
              ${checkbox}
            </div>
          </div>
          ${(enabled) ? html`<div class="row full"><div class="cell padding-small" >
            <form-field id="${`field_${name}`}"
              .field=${this.row} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div></div>`:null}
          ${(info) ? html`<div class="row full padding-small">
            <div class="cell padding-small info leftbar" >
              ${info}
            </div>
          </div>`:null}
        </div>`)

      case "image":
        return(
          html`<div id="${this.id}"
            style="${styleMap(this.style)}" class="container-row">
            <div class="row full">
              <div class="cell padding-small" >
                ${checkbox}
              </div>
              ${(enabled) ? html`<div class="cell padding-small" >
              <form-input 
                id="${`file_${name}`}" 
                type="${INPUT_TYPE.FILE}" ?full="${true}"
                label="${this.labelAdd}"
                .style="${{ "font-size": "12px"}}"
                .onChange=${
                  (event) => this._onEdit({id,
                    file: true,
                    name, 
                    value: event.value, 
                    extend: false
                })}></form-input>
              </div>`:null}
            </div>
            ${(enabled) ? html`<div class="row full"><div class="cell padding-small" >
              <textarea id="${`input_${name}`}"
                class=${`full`} rows=5 .value="${this.imgValue()}"
                @input="${this._onTextInput}" ></textarea>
              <div class="full padding-normal center" >
                <img src="${this.imgValue()}" alt="" />
              </div>
            </div></div>`:null}
            ${(info) ? html`<div class="row full padding-small">
              <div class="cell padding-small info leftbar" >
                ${info}
              </div>
            </div>`:null}
          </div>`)

      case "checklist":
        const cbValue = this.values[name] || ""
        const checklist = []
        this.row.values.forEach((element, index) => {
          const cvalue = element.split("|")
          const value = (cbValue.indexOf(cvalue[0])>-1)
          checklist.push(html`<div id="${`checklist_${name}_${index}`}"
            key={index}
            class="cell padding-small report-field"
            @click=${() => this._onEdit({
              id,
              checklist: true,
              name,
              checked: !value,
              value: cvalue[0],
              extend: false
            })}>
            <form-label 
              value="${cvalue[1]}" class="bold ${(value) ? "toggle-on" : ""}"
              leftIcon="${(value) ? "CheckSquare" : "SquareEmpty"}" ></form-label>
          </div>`)
        });
        return(
          html`<div id="${this.id}"
            style="${styleMap(this.style)}" class="container-row">
          <div class="row full">
            <div class="cell padding-small" >
              ${checkbox}
            </div>
          </div>
          ${(enabled) ? html`<div class="row full padding-small">
            <div class="cell padding-small toggle" >
              ${checklist}
            </div>
          </div>`:null}
          ${(info) ? html`<div class="row full padding-small">
            <div class="cell padding-small info leftbar" >
              ${info}
            </div>
          </div>`:null}
        </div>`)

      default:
        return(html`<div id="${this.id}"
          style="${styleMap(this.style)}" class="container-row">
          <div class="row full">
            <div class="cell padding-small half" >
              ${checkbox}
            </div>
            ${(enabled)?html`<div class="cell padding-small half" >
              <form-field id="${`field_${name}`}"
                .field=${this.row} .values=${this.values} .options=${this.options} .data=${this.data}
                .msg=${this.msg} .onEdit=${this.onEdit} 
                .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
            </div>`:null}
          </div>
          ${(info)?html`<div class="row full padding-small">
            <div class="cell padding-small info leftbar" >
              ${info}
            </div>
          </div>`:null}
        </div>`)
      }
  }

  render() {
    const { 
      id, rowtype, label, columns, name, disabled, notes, selected, empty, datatype, info 
    } = this.row

    switch (rowtype) {

      case "label":
        return html`<div id="${this.id}" style="${styleMap(this.style)}" 
          class="container-row label-row">
          <div class="cell padding-small" >${this.values[name] || label}</div>
        </div>`

      case "flip":
        return this.flipItem()

      case "field":
        return(html`<div id="${this.id}"
          style="${styleMap(this.style)}" class="container-row">
          <div class="cell padding-small hide-small field-cell" >
            <form-label value="${label}" class="bold" ></form-label>
          </div>
          <div class="cell padding-small" >
            <div class="hide-medium hide-large" >
              <form-label value="${label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${name}`}"
              .field=${this.row} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
        </div>`)

      case "reportfield":
        return(html`<div id="${this.id}"
          style="${styleMap(this.style)}" class="cell padding-small s12 m6 l4">
          <div id="${`cb_${name}`}"
            class=${`${"padding-small"} ${(empty !== 'false') ? "report-field" : ""}`} 
            @click="${() => {if(empty !== 'false'){
              this._onEdit({id, name: "selected", value: !selected, extend: false })} }}">
            <form-label 
              value="${label}" class="bold"
              leftIcon="${(selected) ? "CheckSquare" : "SquareEmpty"}" ></form-label>
          </div>
          <form-field id="${`field_${name}`}"
            .field=${this.row} .values=${this.values} .options=${this.options} .data=${this.data}
            .msg=${this.msg} .onEdit=${this.onEdit} 
            .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
        </div>`)

      case "fieldvalue":
        return(html`<div id="${this.id}"
          style="${styleMap(this.style)}" class="container-row">
          <div class="row full">
            <div class="cell container-small">
              <form-label value="${label}" class="bold" ></form-label>
            </div>
            <div class="cell align-right container-small" >
              <span id=${`delete_${this.row.fieldname}`}
                class="fieldvalue-delete" 
                @click="${ ()=>this._onEdit({id, name: "fieldvalue_deleted"}) }">
                <form-icon iconKey="Times" ></form-icon>
              </span>
            </div>
          </div>
          <div class="row full">
            <div class="cell padding-small s12 m6 l6" >
              <form-field id="${`field_${this.row.fieldname}`}"
                .field=${this.row} .values=${this.values} .options=${this.options} .data=${this.data}
                .msg=${this.msg} .onEdit=${this.onEdit} 
                .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
            </div>
            <div class="cell padding-small s12 m6 l6" >
              <form-input 
                id="${`notes_${this.row.fieldname}`}" type="${INPUT_TYPE.TEXT}"
                label="${this.msg("", { id: "fnote_view" })}"
                name="fieldvalue_notes" ?full="${true}" value="${notes}"
                ?disabled=${(disabled || this.data.audit === 'readonly')}
                .onChange=${
                  (event) => this._onEdit({
                    id, name: "fieldvalue_notes", value: event.value
                })}></form-input>
            </div>
          </div>
        </div>`)

      case "col2":
        return(html`<div id="${this.id}"
          style="${styleMap(this.style)}" class="container-row">
          <div class="cell padding-small s12 m6 l6" >
            <div>
              <form-label value="${columns[0].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${columns[0].name}`}"
              .field=${columns[0]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
          <div class="cell padding-small s12 m6 l6" >
            <div>
              <form-label value="${columns[1].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${columns[1].name}`}"
              .field=${columns[1]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
        </div>`)

      case "col3":
        return(html`<div id="${this.id}"
          style="${styleMap(this.style)}" class="container-row">
          <div class="cell padding-small s12 m4 l4" >
            <div>
              <form-label value="${columns[0].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${columns[0].name}`}"
              .field=${columns[0]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
          <div class="cell padding-small s12 m4 l4" >
            <div>
              <form-label value="${columns[1].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${columns[1].name}`}"
              .field=${columns[1]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
          <div class="cell padding-small s12 m4 l4" >
            <div>
              <form-label value="${columns[2].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${columns[2].name}`}"
              .field=${columns[2]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
        </div>`)

      case "col4":
        return(html`<div id="${this.id}"
          style="${styleMap(this.style)}" class="container-row">
          <div class="cell padding-small s12 m3 l3" >
            <div>
              <form-label value="${columns[0].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${columns[0].name}`}"
              .field=${columns[0]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
          <div class="cell padding-small s12 m3 l3" >
            <div>
              <form-label value="${columns[1].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${columns[1].name}`}"
              .field=${columns[1]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
          <div class="cell padding-small s12 m3 l3" >
            <div>
              <form-label value="${columns[2].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${columns[2].name}`}"
              .field=${columns[2]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
          <div class="cell padding-small s12 m3 l3" >
            <div>
              <form-label value="${columns[3].label}" class="bold" ></form-label>
            </div>
            <form-field id="${`field_${columns[3].name}`}"
              .field=${columns[3]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></form-field>
          </div>
        </div>`)

      default:
        break;
    }
    return null
  }
}