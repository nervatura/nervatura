/* c8 ignore start */
import * as locales from './locales.js';

import { APP_MODULE, SIDE_VISIBILITY, APP_THEME } from './enums.js'

const publicHost = "nervatura.github.io"
const basePath = "/api"

const version = "__VERSION__"
const serverURL = "__SERVER__"
// const locales = await import('./locales.js');

// Default read and write application context data
export const store = {
  session: {
    version,
    locales,
    serverURL,
    apiPath: "/api",
    engines: ["sqlite", "sqlite3", "mysql", "postgres", "mssql"],
    service: ["dev", "5.1.7", "5.1.8", "5.1.9"],
    helpPage: "https://nervatura.github.io/nervatura/docs/client/"
  },
  ui: {
    toastTimeout: 4, // sec
    paginationPage: 10,
    selectorPage: 5,
    history: 5,
    timeIntervals: 15,
    searchSubtract: 3, // months
    export_sep: ";",
    page_size: "a4",
    page_orient: "portrait",
    printqueue_mode: [
      ["print", "printqueue_mode_print"],
      ["pdf", "printqueue_mode_pdf"],
      ["xml", "printqueue_mode_xml"]
    ],
    printqueue_type: [
      ["customer", "title_customer"],
      ["product", "title_product"],
      ["employee", "title_employee"],
      ["tool", "title_tool"],
      ["project", "title_project"],
      ["order", "title_order"],
      ["offer", "title_offer"],
      ["invoice", "title_invoice"],
      ["receipt", "title_receipt"],
      ["rent", "title_rent"],
      ["worksheet", "title_worksheet"],
      ["delivery", "title_delivery"],
      ["inventory", "title_inventory"],
      ["waybill", "title_waybill"],
      ["production", "title_production"],
      ["formula", "title_formula"],
      ["bank", "title_bank"],
      ["cash", "title_cash"]
    ],
    report_orientation: [
      ["portrait", "report_portrait"],
      ["landscape", "report_landscape"]
    ],
    report_size: [
      ["a3", "A3"], ["a4", "A4"],
      ["a5", "A5"], ["letter", "Letter"],
      ["legal", "Legal"]
    ]
  },
  current: { 
    home: APP_MODULE.SEARCH, module: APP_MODULE.LOGIN, side: SIDE_VISIBILITY.AUTO,
    lang: (localStorage.getItem("lang") && locales[localStorage.getItem("lang")]) ? localStorage.getItem("lang") : "en",
    theme: localStorage.getItem("theme") || APP_THEME.LIGHT
  },
  login: { 
    username: localStorage.getItem("username") || "",
    database: localStorage.getItem("database") || "",
    server: 
      (!localStorage.getItem("server") || (localStorage.getItem("server") === ""))
      ? (window.location.hostname !== publicHost) ? window.location.origin+basePath : ""
      : localStorage.getItem("server")
  },
  search: { seltype: "selector", group_key: "transitem", result: [], vkey: null, qview: "transitem", qfilter: "", 
    filters: {}, columns: {}, browser_filter: true },
  edit: { dataset: {}, current: {}, dirty: false, form_dirty: false },
  setting: { dirty: false, result: [] },
  template: { dirty: false }, 
  bookmark: { history: null, bookmark: [] }
}
