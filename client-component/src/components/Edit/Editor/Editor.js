import { LitElement, html, nothing } from 'lit';
import { ifDefined } from 'lit/directives/if-defined.js';

import '../../Form/Label/form-label.js'
import '../../Form/Button/form-button.js'

import '../Main/edit-main.js'
import '../Meta/edit-meta.js'
import '../Item/edit-item.js'
import '../Note/edit-note.js'
import '../View/edit-view.js'

import { styles } from './Editor.styles.js'
import { EDIT_EVENT, TEXT_ALIGN } from '../../../config/enums.js'

export class Editor extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.id = Math.random().toString(36).slice(2);
    this.caption = ""
    this.audit = "" 
    this.current = {}
    this.template = {}
    this.dataset = {}
    this.paginationPage = 10
    this.selectorPage = 5
    this.onEvent = {}
  }

  static get properties() {
    return {
      id: { type: String },
      caption: { type: String },
      audit: { type: String },
      current: { type: Object },
      template: { type: Object },
      dataset: { type: Object },
      paginationPage: { type: Number },
      selectorPage: { type: Number },
      onEvent: { type: Object }
    };
  }

  static get styles () {
    return [
      styles
    ]
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

  _onNoteBlur(event){
    if( event.target._value !==  event.target.value){
      this.onEvent.onEditEvent({ 
        key: EDIT_EVENT.EDIT_ITEM, data: { name: "fnote", value: event.target._value } })
    }
  }

  render() {
    const { caption, audit, template, current, dataset } = this
    const tabButton = (viewKey, textKey, iconKey, itemCount) => html`
      <div class="row full" >
        <div class="cell" >
          <form-button id="${`btn_${viewKey}`}" 
            icon="${iconKey}" ?full="${true}"
            label="${this.msg(textKey, { id: textKey })}" align=${TEXT_ALIGN.LEFT}
            .style=${{ "border-radius": 0, "margin-top": "2px" }}
            badge="${ifDefined(itemCount)}"
            @click=${()=>this._onEditEvent(EDIT_EVENT.CHANGE, { 
              fieldname: "view", value: (current.view === viewKey) ? "" : viewKey })}
          >${this.msg(textKey, { id: textKey })}</form-button>
        </div>
      </div>`

    const mainLabel = () => {
      let label = current.item[template.options.title_field]
      if(current.type === "printqueue"){
        label = template.options.title_field;
      } else if(current.item.id === null){
        label = `${this.msg("", { id: "label_new" })} ${template.options.title}`;
      }
      return label
    }

    const metafieldCount = () => current.fieldvalue.filter(fv => {
        const _deffield = dataset.deffield.filter((df) => (df.fieldname === fv.fieldname))[0]
        return ((_deffield.visible === 1) && (fv.deleted === 0))
      }).length

    return html`<div class="panel" >
      <div class="panel-title">
        <div class="cell">
          <form-label class="title-cell"
            value="${caption}" leftIcon="${template.options.icon}"
          ></form-label>
        </div>
      </div>
      ${(current.form)?
        html`<div class="section-container" >
          <edit-item
            audit="${audit}" .current="${current}" .dataset="${dataset}"
            .onEvent=${this.onEvent} .msg=${this.msg}
          ></edit-item>
        </div>`:
        html`<div class="section-container" >
          ${tabButton("form", mainLabel(), template.options.icon)}
          ${(current.view === "form") ? html`<edit-main
              audit="${audit}" .current="${current}" .template="${template}" .dataset="${dataset}"
              .onEvent=${this.onEvent} .msg=${this.msg}
            ></edit-main>` : nothing}

          ${((current.item.id !== null || template.options.search_form) 
            && typeof this.dataset.fieldvalue !== "undefined" 
            && current.item !== null && template.options.fieldvalue === true)
            ? html`${tabButton("fieldvalue", "fields_view", template.options.icon, metafieldCount())}${
              (current.view === "fieldvalue") ? html`<edit-meta audit="${audit}" 
              .current="${current}" .dataset="${dataset}" 
              .onEvent=${this.onEvent} .msg=${this.msg} pageSize=${this.selectorPage}
              ></edit-meta>` : nothing }`
            : nothing }

          ${((current.item.id !== null) && (typeof current.item.fnote !== "undefined") && 
            (template.options.pattern === true)) 
            ? html`${tabButton("fnote", "fnote_view", "Comment")}${
              (current.view === "fnote") ? html`<edit-note id="editor_note"
                value="${current.item.fnote}" patternId="${ifDefined(current.template)}"
                .patterns="${dataset.pattern}" ?readOnly="${(audit === "readonly")}"
                .onEvent=${this.onEvent} .msg=${this.msg}
                @blur=${this._onNoteBlur}
              ></edit-note>` : nothing }`
            : nothing }

          ${Object.keys(template.view).filter(
            (vname)=>(template.view[vname].view_audit !== "disabled")).map(
              (vname) => html`${tabButton(vname, template.view[vname].title, 
                template.view[vname].icon, 
                dataset[template.view[vname].data].length)}${
              (current.view === vname) ? html`<edit-view
                viewName=${vname} .current=${current} .template=${template} 
                .dataset=${dataset} audit=${audit}
                .onEvent=${this.onEvent} .msg=${this.msg} pageSize=${this.paginationPage}
              ></edit-view>` : nothing }`
          )}
        </div>`
      }
    </div>`
  }
}