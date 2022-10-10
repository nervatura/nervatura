/* c8 ignore start */
import * as locales from './locales.js';

const publicHost = "nervatura.github.io"
const basePath = "/api"

const version = "__VERSION__"
const serverURL = "__SERVER__"
/*
const locales = {
  en: { key: "en", en: "English" }
}
*/

// Default read and write application context data
export const store = {
  session: {
    version,
    locales,
    serverURL,
    proxy: "http://localhost:5000",
    apiPath: "/api",
    engines: ["sqlite", "sqlite3", "mysql", "postgres", "mssql"],
    service: ["dev", "5.0.2", "5.0.3", "5.0.4"],
    helpPage: "https://nervatura.github.io/nervatura/docs/client/"
  },
  ui: {
    toastTime: 7000,
    paginationPage: 10,
    selectorPage: 5,
    history: 5,
    timeIntervals: 15,
    searchSubtract: 3, // months
    filter_opt_1: [["===","EQUAL"],["==N","IS NULL"],["!==","NOT EQUAL"]],
    filter_opt_2: [["===","EQUAL"],["==N","IS NULL"],["!==","NOT EQUAL"],[">==",">="],["<==","<="]],
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
    home: "search", module: "login", side: "auto",
    lang: (localStorage.getItem("lang") && locales[localStorage.getItem("lang")]) ? localStorage.getItem("lang") : "en",
    theme: localStorage.getItem("theme") || "light"
  },
  login: { 
    username: localStorage.getItem("username") || "",
    database: localStorage.getItem("database") || "",
    value: 123,
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

/* c8 ignore stop */

/*

export const getSetting = (key) => {
  switch (key) {    
    case "ui":
      let values = update({}, {$set: store.ui})
      for (const ikey in values) {
        if(localStorage.getItem(ikey)){
          values[ikey] = localStorage.getItem(ikey)
        }
      }
      return values

    default:
      return localStorage.getItem(key) || store.ui[key] || "";
  }
}

*/