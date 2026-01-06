import { LitElement, html, nothing } from 'lit';
import { ifDefined } from 'lit/directives/if-defined.js';
import { ref } from 'lit/directives/ref.js';

import '../../Form/Label/form-label.js'
import '../../Form/Button/form-button.js'
import '../../Form/Select/form-select.js'
import '../../Form/Row/form-row.js'
import '../../Form/List/form-list.js'
import '../../Form/Table/form-table.js'

import { styles } from './Editor.styles.js'
import { TEMPLATE_EVENT, TEXT_ALIGN, BUTTON_TYPE, PAGINATION_TYPE } from '../../../config/enums.js'

const getElementType = (element) => Object.getOwnPropertyNames(element)[0]

export class Editor extends LitElement {
  constructor() {
    super();
    /* c8 ignore next 1 */
    this.msg = (defValue) => defValue
    this.id = Math.random().toString(36).slice(2);
    this.data = {
      title: "", 
      tabView: "template", 
      template: {
        meta: {},
        report: {},
        header: [],
        details: [],
        footer: [],
        sources: {},
        data: {}
      }, 
      current: {}, 
      current_data: null, 
      dataset: []
    }
    this.paginationPage = 10
    this.onEvent = {}
  }

  static get properties() {
    return {
      id: { type: String },
      data: { type: Object },
      paginationPage: { type: Number },
      onEvent: { type: Object }
    };
  }

  static get styles () {
    return [
      styles
    ]
  }

  _onTemplateEvent(key, data){
    if(this.onEvent.onTemplateEvent){
      this.onEvent.onTemplateEvent({ key, data })
    }
    this.dispatchEvent(
      new CustomEvent('template_event', {
        bubbles: true, composed: true,
        detail: {
          key, data
        }
      })
    );
  }

  canvasCallback(canvasRef) {
    if(canvasRef){
      this._onTemplateEvent(TEMPLATE_EVENT.CREATE_MAP, { mapRef: canvasRef })
    }
  }

  tabButton(key, icon) {
    const { tabView } = this.data
    return html`<form-button 
      id="${`btn_${key}`}"
      .style="${{ "border-radius": 0 }}" icon="${icon}"
      label="${this.msg("", { id: `template_label_${key}` })}"
      @click=${()=>this._onTemplateEvent(TEMPLATE_EVENT.CHANGE_TEMPLATE, { key: "tabView", value: key })} 
      type="${(tabView === key) ? BUTTON_TYPE.PRIMARY : ""}"
      ?full="${true}" ?selected="${(tabView === key)}" >
      ${this.msg("", { id: `template_label_${key}` })}</form-button>`
  }

  navButton(key, event, label, icon, full, align) {
    const btnAlign = (typeof(align) === "undefined") ? TEXT_ALIGN.LEFT : align
    const btnFull = (typeof(full) === "undefined") ? true : full
    return html`<form-button 
      id="${`btn_${key}`}"
      .style="${{ "border-radius": 0 }}" icon="${icon}"
      label="${this.msg("", { id: label })}" align="${btnAlign}"
      @click=${()=>this._onTemplateEvent(...event)} 
      type="${BUTTON_TYPE.BORDER}" ?full="${btnFull}" >${this.msg("", { id: label })}</form-button>`
  }

  setListIcon(item, index) {
    const { template, current } = this.data
    const getBadge = (items) => {
      if (typeof index!=="undefined") {
        return index
      }
      if ((typeof items==="object") && (Array.isArray(items)) && (items.length > 0)) {
        return items.length
      }
      return 0
    }
    if (current.item===item) {
      return {
        selected: true, icon: "Tag", color: "green", 
        badge: getBadge(item, index)
      };
    }
    if (current.parent===item || template[current.section]===item) {
      return {
        selected: true, icon: "Check", color: "", 
        badge: getBadge(item, index) 
      };
    }
    let icon = "InfoCircle"
    if (Array.isArray(item)) {
      if (item.length>0) {
        icon = "Plus"
      }
    }
    return { selected: false, icon, color: "", badge: 0 };
  }

  mapButton (key, tmp_id, info, label, badge) {
    const badgeValue = (badge > 0) ? badge : undefined
    return html`<form-button 
      id="${`btn_${key}`}"
      .style="${{ "border-radius": 0 }}" icon="${info.icon}"
      ?selected="${(info.color!=="")}"
      label="${label}" align="${TEXT_ALIGN.LEFT}"
      @click=${()=>this._onTemplateEvent(TEMPLATE_EVENT.SET_CURRENT, [{ tmp_id }])}
      badge="${ifDefined(badgeValue)}"
      type="${BUTTON_TYPE.BORDER}" ?full="${true}" >${label}</form-button>`
  }

  createSubList(maplist) {
    const { template, current } = this.data
    for(let index = 0; index < template[current.section].length; index += 1) {
      const etype = getElementType(template[current.section][index]);
      let item = template[current.section][index][etype];
      const mkey = `tmp_${current.section}_${index.toString()}_${etype}`
      if (["row","datagrid"].includes(etype)) {
        item = item.columns;
      }
      if (current.parent===null) {
        const pinfo = this.setListIcon(item);
        maplist.push(html`<div key={mkey}>
          ${this.mapButton(mkey, mkey, pinfo, etype.toUpperCase(), index+1)}
        </div>`)
      } else if ((current.item===item) || (current.parent===item)) {
          const cinfo = this.setListIcon(item, index+1)
          maplist.push(html`<div key={mkey}>
            ${this.mapButton(mkey, mkey, cinfo, etype.toUpperCase(), cinfo.badge)}
          </div>`)
          if (["row","datagrid"].includes(current.type) || ["row","datagrid"].includes(current.parent_type)) {
            for(let i2 = 0; i2 < item.length; i2 += 1) {
              const subtype = getElementType(item[i2]);
              const subitem = item[i2][subtype];
              const skey = `tmp_${current.section}_${index.toString()}_${etype}_${i2.toString()}_${subtype}`
              const sinfo = this.setListIcon(subitem);
              maplist.push(html`<div key={skey}>
                ${this.mapButton(skey, skey, sinfo, subtype.toUpperCase(), sinfo.badge, `${"primary"} ${styles.badgeBlack}`)}
              </div>`)
            }
          }
        }
    }
  }

  createMapList() {
    const { template } = this.data
    const maplist = [];
    ["report", "header", "details", "footer"].forEach(mkey => {
      const info = this.setListIcon(template[mkey]);
      if (info.selected && (mkey !== "report")) {
        maplist.push(html`<hr id="${`sep_${mkey}_0`}" class="separator" />`)
      }
      maplist.push(html`<div key={mkey}>
        ${this.mapButton(`tmp_${mkey}`, `tmp_${mkey}`, info, mkey.toUpperCase(), info.badge)}
      </div>`)
      if (info.selected) {
        maplist.push(html`<hr id="${`sep_${mkey}_1`}" class="separator" />`)
        if(mkey !== "report"){
          this.createSubList(maplist)
          maplist.push(html`<hr id="${`sep_${mkey}_2`}" class="separator" />`)
        }
      }
    })
    return maplist;
  }

  dataTitle(key, title, event){
    return html`<div class="panel-title">
      <div class="cell">
        <form-label class="title-cell"
          value="${title}"
        ></form-label>
      </div>
      <div class="cell align-right" >
        <span id="${key}" class="close-icon" 
          @click="${ ()=>this._onTemplateEvent(...event) }">
          <form-icon iconKey="Times" ></form-icon>
        </span>
      </div>
    </div>`
  }

  dataText(key, value, rows, params) {
    return html`<textarea id="${`${key}_value`}"
      rows=${rows} .value="${value}"
      @input="${(event)=>this._onTemplateEvent(TEMPLATE_EVENT.EDIT_DATA_ITEM, { ...params, value: event.target.value })}"
    ></textarea>`
  }

  tableFields() {
    const { current_data } = this.data
    return {...current_data.fields,
      edit: { columnDef: { 
        id: "delete",
        Header: "",
        headerStyle: {},
        Cell: ({ row }) => html`<form-icon id=${`delete_${row._index}`}
            iconKey="Times" width=19 height=27.6
            .style=${{ cursor: "pointer", fill: "rgb(var(--functional-red))" }}
            @click=${ (event)=>{
              event.stopPropagation();
              this._onTemplateEvent(TEMPLATE_EVENT.DELETE_DATA_ITEM , { _index: row._index })
            }}
          ></form-icon>`,
        cellStyle: { width: 40, padding: "4px 8px 3px 8px" }
      }}
    }
  }

  render() {
    const { title, tabView, template, current, current_data, dataset } = this.data
    const getMapCtr = (type, key) => {
      const keyMap = {
        map_edit: { data: false, report: false, header: false, footer: false, details: false },
        map_insert: { header: true, footer: true, details: true, row: true, datagrid: true }
      }
      return keyMap[key][type]
    }
    const reportElements = {
      header: ["row", "vgap", "hline"],
      details: ["row", "vgap", "hline", "html", "datagrid"],
      footer: ["row", "vgap", "hline"],
      row: ["cell", "image", "barcode", "separator"],
      datagrid: ["column"]
    }
    return html`<div class="panel" >
      <div class="panel-title">
        <div class="cell">
          <form-label class="title-cell"
            value="${title}" leftIcon="TextHeight"
          ></form-label>
        </div>
      </div>
      <div class="section-container" >
        <div class="row full">
          <div class="cell third">${this.tabButton("template", "Tags")}</div>
          <div class="cell third">${this.tabButton("data", "Database")}</div>
          <div class="cell third">${this.tabButton("meta", "InfoCircle")}</div>
        </div>
        ${(tabView === "template") ? html`<div class="section-container-small border" >
          <div class="cell padding-small third" >
            ${this.navButton("previous", [TEMPLATE_EVENT.GO_PREVIOUS,[]], "label_previous", "ArrowLeft")}
            <div class="mapBox" >
              <canvas ${ref(this.canvasCallback)} class="reportMap" ></canvas>
            </div>
            ${this.navButton("next", [TEMPLATE_EVENT.GO_NEXT,[]], "label_next", "ArrowRight", true, TEXT_ALIGN.RIGHT)}
          </div>
          <div class="cell padding-small third" >
            ${this.createMapList()}
          </div>
          <div class="cell padding-small third" >
            ${(getMapCtr(current.type, "map_edit") !== false) 
              ? html`<div>
                ${this.navButton("move_up", [TEMPLATE_EVENT.MOVE_UP,[]], "label_move_up", "ArrowUp")}
                ${this.navButton("move_down", [TEMPLATE_EVENT.MOVE_DOWN,[]], "label_move_down", "ArrowDown")}
                ${this.navButton("delete_item", [TEMPLATE_EVENT.DELETE_ITEM,[]], "label_delete", "Times")}
                <hr class="separator" />
              </div>` : nothing}
            ${(getMapCtr(current.type, "map_insert")) 
              ? html`<div>
                ${this.navButton("add_item", [TEMPLATE_EVENT.ADD_ITEM, current.add_item||"" ], "label_add_item", "Plus")}
                <form-select id="sel_add_item" label="" 
                  ?full=${false} .isnull="${true}" value="${current.add_item||""}"
                  .onChange=${(event) => this._onTemplateEvent(TEMPLATE_EVENT.CHANGE_CURRENT, { key: "add_item", value: event.value })}
                  .options=${reportElements[current.type].map(
                    (item)=>({ value: item, text: item.toUpperCase() }))}  
                ></form-select>
              </div>` : nothing}
          </div>
        </div>
        <div class="report-title padding-small" >
          <form-label class="report-title-label"
            value="${current.type.toUpperCase()}" leftIcon="Tag"
          ></form-label>
        </div>
        ${current.form.rows.map((row, index) =>html`<div class="template-row" >
          <form-row id=${`row_${index}`}
            .row=${row} 
            .values=${(["row","datagrid"].includes(current.type)) ? current.item_base : current.item}
            .options=${current.form.options}
            .data=${{ audit: "all", current, dataset: template.data }}
            .onEdit=${(data) => this._onTemplateEvent(TEMPLATE_EVENT.EDIT_ITEM, data)}
            .msg=${this.msg}
          ></form-row>
        </div>`)}` : nothing }
        ${(tabView === "data") ? html`<div class="section-container border" >
          ${(current_data && (current_data.type === "string"))?
            html`<div class="row full section-small">
              ${this.dataTitle("data_string", current_data.name, [TEMPLATE_EVENT.SET_CURRENT_DATA, null ])}
              ${this.dataText(current_data.name, template.data[current_data.name], 15, {})}
          </div>` : nothing}
          ${(current_data && (current_data.type === "list") && current_data.item)?
            html`<div class="row full section-small">
              ${this.dataTitle("data_list_item", current_data.item, [TEMPLATE_EVENT.SET_CURRENT_DATA_ITEM, null ])}
              ${this.dataText(current_data.item, template.data[current_data.name][current_data.item], 10, {})}
          </div>` : nothing}
          ${(current_data && (current_data.type === "table") && current_data.item)?
            html`<div class="row full section-small">
              ${this.dataTitle("data_table_item", `${current_data.name} - ${String(current_data.item._index+1)}`, [TEMPLATE_EVENT.SET_CURRENT_DATA_ITEM, null ])}
              ${Object.keys(current_data.fields).map((field) => html`<div class="row full">
                <div class="padding-small">
                  <form-label value="${field}"></form-label>
                </div>
                ${this.dataText(field, current_data.item[field], 2, {field, _index: current_data.item._index})}
              </div>`)}
          </div>` : nothing}
          ${(current_data && (current_data.type === "list") && !current_data.item)?
            html`<div class="row full section-small">
              ${this.dataTitle("data_list", current_data.name, [TEMPLATE_EVENT.SET_CURRENT_DATA, null ])}
              <form-list id="data_list_items"
                .rows=${current_data.items} ?listFilter=${true}
                filterPlaceholder=${this.msg("", { id: "placeholder_filter" })}
                .onAddItem=${() => this._onTemplateEvent(TEMPLATE_EVENT.SET_CURRENT_DATA_ITEM , undefined )}
                labelAdd=${this.msg("", { id: "label_new" })}
                pageSize=${this.paginationPage} pagination="${PAGINATION_TYPE.TOP}" 
                .onEdit=${(row) => this._onTemplateEvent(TEMPLATE_EVENT.SET_CURRENT_DATA_ITEM , row.lslabel )}
                .onDelete=${(row) => this._onTemplateEvent(TEMPLATE_EVENT.DELETE_DATA_ITEM , {key: row.lslabel} )}
              ></form-list>
          </div>` : nothing}
          ${(current_data && (current_data.type === "table") && !current_data.item)?
            html`<div class="row full section-small">
              ${this.dataTitle("data_table", current_data.name, [TEMPLATE_EVENT.SET_CURRENT_DATA, null ])}
              <form-table id="data_table_items"
                .fields=${this.tableFields()} .rows=${current_data.items} ?tableFilter=${true}
                filterPlaceholder="${this.msg("", { id: "placeholder_filter" })}"
                .onAddItem=${() => this._onTemplateEvent(TEMPLATE_EVENT.SET_CURRENT_DATA_ITEM , undefined )}
                .onRowSelected=${(row) => this._onTemplateEvent(TEMPLATE_EVENT.SET_CURRENT_DATA_ITEM , row )} 
                labelAdd=${this.msg("", { id: "label_new" })}  
                pageSize=${this.paginationPage} pagination="${PAGINATION_TYPE.TOP}"
              ></form-table>
          </div>` : nothing}
          ${(!current_data)?
            html`<div class="row full section-small">
              <form-list id="data_list_items"
                .rows=${dataset} ?listFilter=${true}
                filterPlaceholder=${this.msg("", { id: "placeholder_filter" })}
                .onAddItem=${() => this._onTemplateEvent(TEMPLATE_EVENT.ADD_TEMPLATE_DATA , [] )}
                labelAdd=${this.msg("", { id: "label_new" })}
                pageSize=${this.paginationPage} pagination="${PAGINATION_TYPE.TOP}" 
                .onEdit=${(row) => this._onTemplateEvent(TEMPLATE_EVENT.SET_CURRENT_DATA , { name: row.lslabel, type: row.lsvalue } )}
                .onDelete=${(row) => this._onTemplateEvent(TEMPLATE_EVENT.DELETE_DATA , row.lslabel )}
              ></form-list>
          </div>` : nothing}
        </div>` : nothing }
        ${(tabView === "meta") ? html`<div class="section-container-small border" >
          <div class="cell padding-small" >
            <div class="meta-title-row" >
              ${Object.keys(template.meta).filter(mkey => ["report_key", "report_name", "description"].includes(mkey)).map(mkey => html`<div class="meta-title-cell" >
                <div class="bold">${mkey}</div>
                <div>${template.meta[mkey]}</div>
              </div>`)}
            </div>
            <div class="meta-title-row" >
              ${Object.keys(template.meta).filter(mkey => !["report_key", "report_name", "description"].includes(mkey)).map(mkey => html`<div class="meta-title-cell" >
                <div class="bold">${mkey}</div>
                <div>${template.meta[mkey]}</div>
              </div>`)}
            </div>
            <div class="meta-title-sources" >
              <form-label value="${this.msg("", { id: "template_data_sources" })}"></form-label>
            </div>
            ${Object.keys(template.sources).map(skey => html`<div class="meta-sources" >
              <div class="cell padding-small">
                <div class="meta-sources-name padding-small">
                  ${skey}
                </div>
                ${Object.keys(template.sources[skey]).map(sql => html`<div class="row" >
                  <div class="meta-sources-cell padding-small bold" >${sql}:</div>
                  <div class="meta-sources-cell padding-small" >${template.sources[skey][sql]}</div>
                </div>`)}
              </div>
            </div>`)}
          </div>
        </div>` : nothing }
      </div>
    </div>`
  }
}