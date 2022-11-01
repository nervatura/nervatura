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

export const MENU_EVENT = {
  SIDEBAR: "sidebar",
  MODULE: "module",
  SCROLL: "scroll",
}

export const SIDE_EVENT = {
  CHANGE: "change",
  SEARCH_QUICK: "search_quick",
  SEARCH_BROWSER: "search_browser",
  CHECK_EDITOR: "check_editor"
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
  SET_FORM_ACTIONS: "set_form_actions",
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