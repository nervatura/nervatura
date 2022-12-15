import { LitElement, html, nothing } from 'lit';
import { ref, createRef } from 'lit/directives/ref.js';
import { unsafeHTML } from 'lit/directives/unsafe-html.js';

import '../../Form/Button/form-button.js'
import '../../Form/Select/form-select.js'

import { styles } from './Note.styles.js'
import { BUTTON_TYPE, EDIT_EVENT } from '../../../config/enums.js'

export class Note extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.id = Math.random().toString(36).slice(2);
    this.value = ""
    this.patternId = undefined
    this.patterns = []
    this.readOnly = false
    this.bold = false
    this.italic = false
    this.onEvent = {}
    this.editorRef = createRef()
  }

  static get properties() {
    return {
      id: { type: String },
      value: { type: String },
      patternId: { type: Number },
      patterns: { type: Array },
      readOnly: { type: Boolean },
      bold: { type: Boolean },
      italic: { type: Boolean },
      onEvent: { type: Object }
    };
  }

  static get styles () {
    return [
      styles
    ]
  }

  connectedCallback() {
    super.connectedCallback();
    this._value = this.value
  }

  disconnectedCallback() {
    this._onEditEvent(EDIT_EVENT.EDIT_ITEM, { 
      name: "fnote", 
      value: this._value
    })
    super.disconnectedCallback();
  }

  _onEditEvent(key, data){
    if(this.onEvent.onEditEvent){
      this.onEvent.onEditEvent({ key, data })
    }
    this.dispatchEvent(
      new CustomEvent('edit_event', {
        bubbles: true, composed: true,
        detail: {
          key, data
        }
      })
    );
  }

  _onInput(){
    this._value = String(this.editorRef.value.innerHTML).split("-->")[1]
  }

  _onContentState () {
    this.bold = document.queryCommandState("bold")
    this.italic = document.queryCommandState("italic")
  }

  _setContentState (state) {
    document.execCommand(state, false)
    this.editorRef.value.focus()
    this._onContentState()
  }

  _onLostFocus () {
    this.bold = false
    this.italic = false
  }

  render() {
    return html`<div id="${this.id}" class="panel" >
      ${(!this.readOnly) ? html`<div class="actionbar padding-small">
        <div class="cell padding-tiny">
          <form-button id="btn_pattern_default" 
            icon="Home" type="${BUTTON_TYPE.BORDER}"
            label="${this.msg("", { id: "pattern_default" })}"
            .style=${{ "padding": "8px 12px" }}
            @click=${()=>this._onEditEvent(EDIT_EVENT.SET_PATTERN, { key: "default" })}
          ></form-button>
          <form-button id="btn_pattern_load" 
            icon="Download" type="${BUTTON_TYPE.BORDER}"
            label="${this.msg("", { id: "pattern_load" })}"
            .style=${{ "padding": "8px 12px" }}
            @click=${()=>this._onEditEvent(EDIT_EVENT.SET_PATTERN, { key: "load", ref: this })}
          ></form-button>
          <form-button id="btn_pattern_save" 
            icon="Upload" type="${BUTTON_TYPE.BORDER}"
            label="${this.msg("", { id: "pattern_save" })}"
            .style=${{ "padding": "8px 12px" }}
            @click=${()=>this._onEditEvent(EDIT_EVENT.SET_PATTERN, { key: "save", text: this._value })}
          ></form-button>
        </div>
        <div class="cell padding-tiny">
          <form-button id="btn_pattern_new" 
            icon="Plus" type="${BUTTON_TYPE.BORDER}"
            label="${this.msg("", { id: "pattern_new" })}"
            .style=${{ "padding": "8px 12px" }}
            @click=${()=>this._onEditEvent(EDIT_EVENT.SET_PATTERN, { key: "new" })}
          ></form-button>
          <form-button id="btn_pattern_delete" 
            icon="Times" type="${BUTTON_TYPE.BORDER}"
            label="${this.msg("", { id: "pattern_delete" })}"
            .style=${{ "padding": "8px 12px" }}
            @click=${()=>this._onEditEvent(EDIT_EVENT.SET_PATTERN, { key: "delete" })}
          ></form-button>
        </div>
        <div class="cell mobile">
          <div class="cell padding-tiny">
            <form-select id="sel_pattern"
              label="${this.msg("", { id: "title_pattern" })}"
              .onChange=${
                (value) => this._onEditEvent(EDIT_EVENT.CHANGE , { fieldname: "template", value: value.value } )
              }
              .options=${
                this.patterns.map( pattern => ({ value: String(pattern.id), 
                  text: pattern.description+((pattern.defpattern === 1)?"*":"") 
                }))
              } 
              .isnull="${true}" value="${(this.patternId) ? String(this.patternId) : ""}" >
            </form-select>
          </div>
          <div class="cell padding-tiny">
            <form-button id="btn_bold" 
              type="${BUTTON_TYPE.BORDER}"
              label="B" ?selected=${this.bold}
              .style=${{ "padding": "8px 12px" }}
              @click=${()=>this._setContentState("bold")}
            >B</form-button>
            <form-button id="btn_italic" 
              type="${BUTTON_TYPE.BORDER}"
              label="I" ?selected=${this.italic}
              .style=${{ "padding": "8px 12px", "font-style": "italic" }}
              @click=${()=>this._setContentState("italic")}
            >I</form-button>
          </div>
        </div>
      </div>` : nothing}
      <div class="rtf-editor" >
        <div id="editor" ${ref(this.editorRef)}
          class="editor-content"
          @input=${this._onInput} 
          @keyup=${this._onContentState} @mouseup=${this._onContentState} @blur=${this._onLostFocus}
          contentEditable="${!this.readOnly}" >${unsafeHTML(this.value)}</div>
      </div>
    </div>`
  }
}