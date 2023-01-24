import { LitElement, html, nothing } from 'lit';

import '../Form/Label/form-label.js'
import '../Form/Icon/form-icon.js'
import '../Modal/Bookmark/modal-bookmark.js'

import { MenuController } from '../../controllers/MenuController.js'

import { styles } from './MenuBar.styles.js'
import { SIDE_VISIBILITY, APP_MODULE, MENU_EVENT } from '../../config/enums.js'

export class MenuBar extends LitElement {
  constructor() {
    super();
    this.scrollTop = false
    this.side = SIDE_VISIBILITY.AUTO
    this.module = APP_MODULE.SEARCH
    this.bookmark = { history: null, bookmark: [] }
    this.selectorPage = 5
    this.onEvent = new MenuController(this)
    this.modalBookmark = this.modalBookmark.bind(this)
  }

  static get properties() {
    return {
      side: { type: String },
      module: { type: String },
      scrollTop: { type: Boolean },
      bookmark: { type: Object },
      selectorPage: { type: Number },
      onEvent: { type: Object },
    };
  }

  static get styles () {
    return [
      styles
    ]
  }

  _onMenuEvent(key, data){
    if(this.onEvent.onMenuEvent){
      this.onEvent.onMenuEvent({ key, data })
    }
    this.dispatchEvent(
      new CustomEvent('menu_event', {
        bubbles: true, composed: true,
        detail: {
          key, data
        }
      })
    );
  }

  modalBookmark(bookmark) {
    return html`<modal-bookmark
      .bookmark="${bookmark || this.bookmark}"
      tabView="bookmark"
      pageSize=${this.selectorPage}
      .onEvent=${this.onEvent}
      .msg=${this.msg}
    ></modal-bookmark>`
  }

  selected(key) {
    if(key === this.module){
      return "selected"
    }
    return ""
  }

  render() {
    return html`<div class="menubar ${(this.scrollTop) ? "shadow" : ""}" >
      <div class="cell">
        <div id="mnu_sidebar"
          class="menuitem sidebar" 
          @click=${() => this._onMenuEvent( MENU_EVENT.SIDEBAR )}>
          ${(this.side === SIDE_VISIBILITY.SHOW)
            ? html`<form-label 
                value="${this.msg("Hide", { id: "menu_hide" })}" class="selected exit"
                leftIcon="Close" ></form-label>`
            : html`<form-label class="menu-label"
                value="${this.msg("Menu", { id: "menu_side" })}"
                leftIcon="Bars" .iconStyle="${{ width: "24px", height: "24px" }}" ></form-label>`
          }
        </div>
        <div id="mnu_search_large" 
          class="hide-small hide-medium menuitem"
          @click=${() => this._onMenuEvent( MENU_EVENT.MODULE, { value: APP_MODULE.SEARCH } )}>
          <form-label class="menu-label ${this.selected(APP_MODULE.SEARCH)}"
            value="${this.msg("Search", { id: "menu_search" })}"
            leftIcon="Search" ></form-label>
        </div>
        <div id="mnu_edit_large" 
          class="hide-small hide-medium menuitem"
          @click=${() => this._onMenuEvent( MENU_EVENT.MODULE, { value: APP_MODULE.EDIT } )}>
          <form-label class="menu-label ${this.selected(APP_MODULE.EDIT)}"
            value="${this.msg("Edit", { id: "menu_edit" })}"
            leftIcon="Edit" ></form-label>
        </div>
        <div id="mnu_setting_large" 
          class="hide-small hide-medium menuitem"
          @click=${() => this._onMenuEvent( MENU_EVENT.MODULE, { value: APP_MODULE.SETTING } )}>
          <form-label class="menu-label ${this.selected(APP_MODULE.SETTING)}"
            value="${this.msg("Setting", { id: "menu_setting" })}"
            leftIcon="Cog" ></form-label>
        </div>
        <div id="mnu_bookmark_large" 
          class="hide-small hide-medium menuitem"
          @click=${() => this._onMenuEvent( MENU_EVENT.MODULE, { value: APP_MODULE.BOOKMARK } )}>
          <form-label class="menu-label ${this.selected(APP_MODULE.BOOKMARK)}"
            value="${this.msg("Bookmark", { id: "menu_bookmark" })}"
            leftIcon="Star" ></form-label>
        </div>
        <div id="mnu_help_large" 
          class="hide-small hide-medium menuitem"
          @click=${() => this._onMenuEvent( MENU_EVENT.MODULE, { value: APP_MODULE.HELP } )}>
          <form-label class="menu-label ${this.selected(APP_MODULE.HELP)}"
            value="${this.msg("Help", { id: "menu_help" })}"
            leftIcon="QuestionCircle" ></form-label>
        </div>
        <div id="mnu_logout_large" 
          class="hide-small hide-medium menuitem"
          @click=${() => this._onMenuEvent( MENU_EVENT.MODULE, { value: APP_MODULE.LOGIN } )}>
          <form-label class="menu-label exit"
            value="${this.msg("Logout", { id: "menu_logout" })}"
            leftIcon="Exit" ></form-label>
        </div>
        ${(this.scrollTop)
          ? html`<div id="mnu_scroll" class="menuitem" 
            @click=${() => this._onMenuEvent( MENU_EVENT.SIDEBAR )}>
            <span class="menu-label" ><form-icon iconKey="HandUp" ></form-icon></span>
          </div>`
          : nothing
        }
      </div>
      <div class="cell container">
        <div id="mnu_help_medium" 
          class="right hide-large menuitem"
          @click=${() => this._onMenuEvent( MENU_EVENT.MODULE, { value: APP_MODULE.HELP } )}>
          <span class="hide-small"><form-label class="menu-label ${this.selected(APP_MODULE.HELP)}"
            value="${this.msg("Help", { id: "menu_help" })}"
            leftIcon="QuestionCircle" ></form-label></span>
          <span class="menu-label hide-medium ${this.selected(APP_MODULE.HELP)}" ><form-icon iconKey="QuestionCircle" ></form-icon></span>
        </div>
        <div id="mnu_bookmark_medium" 
          class="right hide-large menuitem"
          @click=${() => this._onMenuEvent( MENU_EVENT.MODULE, { value: APP_MODULE.BOOKMARK } )}>
          <span class="hide-small"><form-label class="menu-label ${this.selected(APP_MODULE.BOOKMARK)}"
            value="${this.msg("Bookmark", { id: "menu_bookmark" })}"
            leftIcon="Star" ></form-label></span>
          <span class="menu-label hide-medium ${this.selected(APP_MODULE.BOOKMARK)}" ><form-icon iconKey="Star" ></form-icon></span>
        </div>
        <div id="mnu_setting_medium" 
          class="right hide-large menuitem"
          @click=${() => this._onMenuEvent( MENU_EVENT.MODULE, { value: APP_MODULE.SETTING } )}>
          <span class="hide-small"><form-label class="menu-label ${this.selected(APP_MODULE.SETTING)}"
            value="${this.msg("Setting", { id: "menu_setting" })}"
            leftIcon="Cog" ></form-label></span>
          <span class="menu-label hide-medium ${this.selected(APP_MODULE.SETTING)}" ><form-icon iconKey="Cog" ></form-icon></span>
        </div>
        <div id="mnu_edit_medium" 
          class="right hide-large menuitem"
          @click=${() => this._onMenuEvent( MENU_EVENT.MODULE, { value: APP_MODULE.EDIT } )}>
          <span class="hide-small"><form-label class="menu-label ${this.selected(APP_MODULE.EDIT)}"
            value="${this.msg("Edit", { id: "menu_edit" })}"
            leftIcon="Edit" ></form-label></span>
          <span class="menu-label hide-medium ${this.selected(APP_MODULE.EDIT)}" ><form-icon iconKey="Edit" ></form-icon></span>
        </div>
        <div id="mnu_search_medium" 
          class="right hide-large menuitem"
          @click=${() => this._onMenuEvent( MENU_EVENT.MODULE, { value: APP_MODULE.SEARCH } )}>
          <span class="hide-small"><form-label class="menu-label ${this.selected(APP_MODULE.SEARCH)}"
            value="${this.msg("Search", { id: "menu_search" })}"
            leftIcon="Search" ></form-label></span>
          <span class="menu-label hide-medium ${this.selected(APP_MODULE.SEARCH)}" ><form-icon iconKey="Search" ></form-icon></span>
        </div>
      </div>
      <div id="mnu_logout_medium" class="hide-large menuitem" style="width: 24px;"
        @click=${() => this._onMenuEvent( MENU_EVENT.MODULE, { value: APP_MODULE.LOGIN } )}>
        <span class="menu-label exit"><form-icon iconKey="Exit" width=24 height=24 ></form-icon></span>
      </div>
    </div>`
  }

}