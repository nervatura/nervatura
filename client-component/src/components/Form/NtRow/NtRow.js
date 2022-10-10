/* eslint-disable lit-a11y/click-events-have-key-events */
/* eslint-disable no-case-declarations */
import { LitElement, html } from 'lit';
import { styleMap } from 'lit/directives/style-map.js';
// import { ifDefined } from 'lit/directives/if-defined.js';

import '../NtInput/nt-input.js'
import '../NtLabel/nt-label.js'
import '../NtIcon/nt-icon.js'
import '../NtField/nt-field.js'

import { styles } from './NtRow.styles.js'

export class NtRow extends LitElement {
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
    return imgValue
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
            html`<nt-icon iconKey="ToggleOn" width=40 height=32.6 ></nt-icon>`:
            html`<nt-icon iconKey="ToggleOff" width=40 height=32.6 ></nt-icon>`}
          <nt-label value="${name}" class="bold padding-tiny ${(enabled) ? "toggle-on" : ""} " ></nt-label>
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
                <nt-field id="${`field_${name}`}"
                  .field=${this.row} .values=${this.values} .options=${this.options} .data=${this.data}
                  .msg=${this.msg} .onEdit=${this.onEdit} 
                  .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></nt-field>
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
                  <nt-input 
                    id="${`file_${name}`}" type="file" ?full="${true}"
                    .style="${{ "font-size": "12px"}}"
                    .onChange=${
                      (event) => this._onEdit({id,
                        file: true,
                        name, 
                        value: event.value, 
                        extend: false
                    })}></nt-input>
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
                <nt-label 
                  value="${cvalue[1]}" class="bold ${(value) ? "toggle-on" : ""}"
                  leftIcon="${(value) ? "CheckSquare" : "SquareEmpty"}" ></nt-label>
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
                  <nt-field id="${`field_${name}`}"
                    .field=${this.row} .values=${this.values} .options=${this.options} .data=${this.data}
                    .msg=${this.msg} .onEdit=${this.onEdit} 
                    .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></nt-field>
                </div>`:null}
              </div>
              ${(info)?html`<div class="row full padding-small">
                <div class="cell padding-small info leftbar" >
                  ${info}
                </div>
              </div>`:null}
            </div>`)
          }

      case "field":
        return(html`<div id="${this.id}"
          style="${styleMap(this.style)}" class="container-row">
          <div class="cell padding-small hide-small field-cell" >
            <nt-label value="${label}" class="bold" ></nt-label>
          </div>
          <div class="cell padding-small" >
            <div class="hide-medium hide-large" >
              <nt-label value="${label}" class="bold" ></nt-label>
            </div>
            <nt-field id="${`field_${name}`}"
              .field=${this.row} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></nt-field>
          </div>
        </div>`)

      case "reportfield":
        return(html`<div id="${this.id}"
          style="${styleMap(this.style)}" class="cell padding-small s12 m6 l4">
          <div id="${`cb_${name}`}"
            class=${`${"padding-small"} ${(empty !== 'false') ? "report-field" : ""}`} 
            @click="${() => {if(empty !== 'false'){
              this._onEdit({id, name: "selected", value: !selected, extend: false })} }}">
            <nt-label 
              value="${label}" class="bold"
              leftIcon="${(selected) ? "CheckSquare" : "SquareEmpty"}" ></nt-label>
          </div>
          <nt-field id="${`field_${name}`}"
            .field=${this.row} .values=${this.values} .options=${this.options} .data=${this.data}
            .msg=${this.msg} .onEdit=${this.onEdit} 
            .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></nt-field>
        </div>`)

      case "fieldvalue":
        return(html`<div id="${this.id}"
          style="${styleMap(this.style)}" class="container-row">
          <div class="row full">
            <div class="cell container-small">
              <nt-label value="${label}" class="bold" ></nt-label>
            </div>
            <div class="cell align-right container-small" >
              <span id=${`delete_${this.row.fieldname}`}
                class="fieldvalue-delete" 
                @click="${ ()=>this._onEdit({id, name: "fieldvalue_deleted"}) }">
                <nt-icon iconKey="Times" ></nt-icon>
              </span>
            </div>
          </div>
          <div class="row full">
            <div class="cell padding-small s12 m6 l6" >
              <nt-field id="${`field_${this.row.fieldname}`}"
                .field=${this.row} .values=${this.values} .options=${this.options} .data=${this.data}
                .msg=${this.msg} .onEdit=${this.onEdit} 
                .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></nt-field>
            </div>
            <div class="cell padding-small s12 m6 l6" >
              <nt-input 
                id="${`notes_${this.row.fieldname}`}" type="text" 
                name="fieldvalue_notes" ?full="${true}" value="${notes}"
                ?disabled=${(disabled || this.data.audit === 'readonly')}
                .onChange=${
                  (event) => this._onEdit({
                    id, name: "fieldvalue_notes", value: event.value
                })}></nt-input>
            </div>
          </div>
        </div>`)

      case "col2":
        return(html`<div id="${this.id}"
          style="${styleMap(this.style)}" class="container-row">
          <div class="cell padding-small s12 m6 l6" >
            <div>
              <nt-label value="${columns[0].label}" class="bold" ></nt-label>
            </div>
            <nt-field id="${`field_${columns[0].name}`}"
              .field=${columns[0]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></nt-field>
          </div>
          <div class="cell padding-small s12 m6 l6" >
            <div>
              <nt-label value="${columns[1].label}" class="bold" ></nt-label>
            </div>
            <nt-field id="${`field_${columns[1].name}`}"
              .field=${columns[1]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></nt-field>
          </div>
        </div>`)

      case "col3":
        return(html`<div id="${this.id}"
          style="${styleMap(this.style)}" class="container-row">
          <div class="cell padding-small s12 m4 l4" >
            <div>
              <nt-label value="${columns[0].label}" class="bold" ></nt-label>
            </div>
            <nt-field id="${`field_${columns[0].name}`}"
              .field=${columns[0]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></nt-field>
          </div>
          <div class="cell padding-small s12 m4 l4" >
            <div>
              <nt-label value="${columns[1].label}" class="bold" ></nt-label>
            </div>
            <nt-field id="${`field_${columns[1].name}`}"
              .field=${columns[1]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></nt-field>
          </div>
          <div class="cell padding-small s12 m4 l4" >
            <div>
              <nt-label value="${columns[2].label}" class="bold" ></nt-label>
            </div>
            <nt-field id="${`field_${columns[2].name}`}"
              .field=${columns[2]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></nt-field>
          </div>
        </div>`)

      case "col4":
        return(html`<div id="${this.id}"
          style="${styleMap(this.style)}" class="container-row">
          <div class="cell padding-small s12 m3 l3" >
            <div>
              <nt-label value="${columns[0].label}" class="bold" ></nt-label>
            </div>
            <nt-field id="${`field_${columns[0].name}`}"
              .field=${columns[0]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></nt-field>
          </div>
          <div class="cell padding-small s12 m3 l3" >
            <div>
              <nt-label value="${columns[1].label}" class="bold" ></nt-label>
            </div>
            <nt-field id="${`field_${columns[1].name}`}"
              .field=${columns[1]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></nt-field>
          </div>
          <div class="cell padding-small s12 m3 l3" >
            <div>
              <nt-label value="${columns[2].label}" class="bold" ></nt-label>
            </div>
            <nt-field id="${`field_${columns[2].name}`}"
              .field=${columns[2]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></nt-field>
          </div>
          <div class="cell padding-small s12 m3 l3" >
            <div>
              <nt-label value="${columns[3].label}" class="bold" ></nt-label>
            </div>
            <nt-field id="${`field_${columns[3].name}`}"
              .field=${columns[3]} .values=${this.values} .options=${this.options} .data=${this.data}
              .msg=${this.msg} .onEdit=${this.onEdit} 
              .onEvent=${this.onEvent} .onSelector=${this.onSelector} ></nt-field>
          </div>
        </div>`)

      default:
        return null;
    }
  }
}