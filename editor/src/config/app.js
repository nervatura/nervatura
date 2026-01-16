/* c8 ignore start */
import * as locales from './locales.js';

import { APP_MODULE, SIDE_VISIBILITY, APP_THEME } from './enums.js'

const publicHost = "nervatura.github.io"
const basePath = "/api/v6"

const version = "__VERSION__"
const serverURL = "__SERVER__"
// const locales = await import('./locales.js');

// Default read and write application context data
export const store = {
  session: {
    version,
    locales,
    serverURL,
    apiPath: "/api/v6",
    helpPage: "/docs/"
  },
  ui: {
    toastTimeout: 4, // sec
    paginationPage: 10,
    timeIntervals: 15,
    export_sep: ";",
    page_size: "a4",
    page_orient: "portrait",
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
    home: APP_MODULE.TEMPLATE, module: APP_MODULE.LOGIN, side: SIDE_VISIBILITY.AUTO,
    lang: (localStorage.getItem("lang") && locales[localStorage.getItem("lang")]) ? localStorage.getItem("lang") : "en",
    theme: localStorage.getItem("theme") || APP_THEME.LIGHT
  },
  login: { 
    username: localStorage.getItem("username") || "",
    database: localStorage.getItem("database") || "",
    code: localStorage.getItem("code") || "",
    server: 
      (!localStorage.getItem("server") || (localStorage.getItem("server") === ""))
      ? (window.location.hostname !== publicHost) ? window.location.origin+basePath : ""
      : localStorage.getItem("server")
  },
  template: { dirty: false }
}
