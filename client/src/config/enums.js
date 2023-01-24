export const APP_THEME = {
  LIGHT: "light", 
  DARK: "dark"
}

export const APP_MODULE = {
  LOGIN: "login",
  SEARCH: "search",
  EDIT: "edit",
  SETTING: "setting",
  HELP: "help",
  BOOKMARK: "bookmark",
  TEMPLATE: "template"
}

export const BUTTON_TYPE = {
  DEFAULT: "default",
  PRIMARY: "primary", 
  BORDER: "border", 
}

export const TEXT_ALIGN = {
  LEFT: "left",
  CENTER: "center", 
  RIGHT: "right", 
}

export const INPUT_TYPE = {
  TEXT: "text",
  COLOR: "color", 
  FILE: "file", 
  PASSWORD: "password"
}

export const DATETIME_TYPE = { 
  DATE: "date", 
  TIME: "time", 
  DATETIME: "datetime-local"
}

export const TOAST_TYPE = {
  INFO: "info",
  ERROR: "error", 
  SUCCESS: "success"
}

export const SIDE_VISIBILITY = {
  AUTO: "auto",
  SHOW: "show",
  HIDE: "hide"
}

export const SIDE_VIEW = {
  EDIT: "edit",
  NEW: "new"
}

export const MENU_EVENT = {
  SIDEBAR: "sidebar",
  MODULE: "module",
  SCROLL: "scroll",
}

export const SIDE_EVENT = {
  BACK: "back",
  CHANGE: "change",
  QUICK: "quick",
  BROWSER: "browser",
  CHECK: "check",
  PREV_NUMBER: "prev_number",
  NEXT_NUMBER: "next_number",
  SAVE: "save",
  DELETE: "delete",
  NEW: "new",
  COPY: "copy",
  LINK: "link",
  PASSWORD: "password",
  SHIPPING_ADD_ALL: "shipping_add_all",
  SHIPPING_CREATE: "shipping_create",
  REPORT_SETTINGS: "report_settings",
  SEARCH_QUEUE: "search_queue",
  EXPORT_QUEUE_ALL: "export_queue_all",
  CREATE_REPORT: "create_report",
  EXPORT_EVENT: "export_event",
  SAVE_BOOKMARK: "save_bookmark",
  LOAD_SETTING: "load_setting",
  PROGRAM_SETTING: "program_setting",
  PASSWORD_FORM: "password_form",
  HELP: "help",
  BLANK: "blank",
  SAMPLE: "sample"
}

export const LOGIN_PAGE_EVENT = {
  CHANGE: "change",
  LOGIN: "login",
  THEME: "theme",
  LANG: "lang"
}

export const MODAL_EVENT = {
  CANCEL: "cancel",
  DELETE: "delete",
  SELECTED: "selected",
  OK: "ok",
  SEARCH: "search",
  CURRENT_PAGE: "current_page",
}

export const PAGINATION_TYPE = {
  TOP: "top",
  BOTTOM: "bottom", 
  ALL: "all",
  NONE: "none" 
}

export const BOOKMARK_VIEW = {
  BOOKMARK: "bookmark",
  HISTORY: "history"
}

export const BROWSER_EVENT = {
  CHANGE: "change",
  BROWSER_VIEW: "browser_view",
  BOOKMARK_SAVE: "bookmark_save",
  EXPORT_RESULT: "export_result",
  SHOW_HELP: "show_help",
  SHOW_BROWSER: "show_browser",
  ADD_FILTER: "add_filter",
  SHOW_TOTAL: "show_total",
  SET_COLUMNS: "set_columns",
  EDIT_FILTER: "edit_filter",
  DELETE_FILTER: "delete_filter",
  SET_FORM_ACTION: "set_form_action",
  EDIT_CELL: "edit_cell",
  CURRENT_PAGE: "current_page"
}

export const BROWSER_FILTER = [
  ["===", "EQUAL"],
  ["==N", "IS NULL"],
  ["!==", "NOT EQUAL"],
  [">==", ">="],
  ["<==", "<="]
]

export const EDIT_EVENT = {
  CHANGE: "change",
  CHECK_EDITOR: "check_editor",
  CHECK_TRANSTYPE: "check_transtype",
  EDIT_ITEM: "edit_item",
  SET_PATTERN: "set_pattern",
  SELECTOR: "selector",
  FORM_ACTION: "form_action"
}

export const EDITOR_EVENT = {
  LOAD_EDITOR: "load_editor",
  SET_EDITOR: "set_editor",
  SET_EDITOR_ITEM: "set_editor_item",
  LOAD_FORMULA: "load_formula",
  NEW_FIELDVALUE: "new_fieldvalue",
  CREATE_TRANS: "create_trans",
  CREATE_TRANS_OPTIONS: "create_trans_options",
  FORM_ACTION: "form_action"
}

export const ACTION_EVENT = {
  LOAD_EDITOR: "load_editor",
  NEW_EDITOR_ITEM: "new_editor_item",
  EDIT_EDITOR_ITEM: "edit_editor_item",
  DELETE_EDITOR_ITEM: "delete_editor_item",
  LOAD_SHIPPING: "load_shipping",
  ADD_SHIPPING_ROW: "add_shipping_row",
  SHOW_SHIPPING_STOCK: "show_shipping_stock",
  EDIT_SHIPPING_ROW: "edit_shipping_row",
  DELETE_SHIPPING_ROW: "delete_shipping_row",
  EXPORT_QUEUE_ITEM: "export_queue_item",
  NEW_ITEM: "new_item",
  EDIT_ITEM: "edit_item",
  DELETE_ITEM: "delete_item",
  EDIT_AUDIT: "edit_audit",
  EDIT_MENU_FIELD: "edit_menu_field",
  DELETE_ITEM_ROW: "delete_item_row"
}

export const SETTING_EVENT = {
  EDIT_ITEM: "edit_item",
  FORM_ACTION: "form_action",
  CURRENT_PAGE: "current_page",
}

export const TEMPLATE_DATA_TYPE = {
  TEXT: "string",
  LIST: "list",
  TABLE: "table"
}

export const TEMPLATE_EVENT = {
  ADD_ITEM: "add_item",
  CHANGE_TEMPLATE: "change_template",
  CHANGE_CURRENT: "change_current",
  GO_PREVIOUS: "go_previous",
  GO_NEXT: "go_next",
  CREATE_MAP: "create_map",
  SET_CURRENT: "set_current",
  MOVE_UP: "move_up",
  MOVE_DOWN: "move_down",
  DELETE_ITEM: "delete_item",
  EDIT_ITEM: "edit_item",
  EDIT_DATA_ITEM: "edit_data_item",
  SET_CURRENT_DATA: "set_current_data",
  SET_CURRENT_DATA_ITEM: "set_current_data_item",
  ADD_TEMPLATE_DATA: "add_template_data",
  DELETE_DATA: "delete_data",
  DELETE_DATA_ITEM: "delete_data_item"
}