import{z as e,j as t,T as a,i as r,s,x as i}from"./module-ORVyXmTn.js";import{e as o,i as n,o as l,a as c}from"./module-bL_R6UIF.js";import{l as d}from"./module-r-Be6lM0.js";const{D:p}=e,u=(e,t)=>void 0===t?void 0!==e?._$litType$:e?._$litType$===t,h=e=>void 0===e.strings,g=()=>document.createComment(""),m=(e,t,a)=>{const r=e._$AA.parentNode,s=void 0===t?e._$AB:t._$AA;if(void 0===a){const t=r.insertBefore(g(),s),i=r.insertBefore(g(),s);a=new p(t,i,e,e.options)}else{const t=a._$AB.nextSibling,i=a._$AM,o=i!==e;if(o){let t;a._$AQ?.(e),a._$AM=e,void 0!==a._$AP&&(t=e._$AU)!==i._$AU&&a._$AP(t)}if(t!==s||o){let e=a._$AA;for(;e!==t;){const t=e.nextSibling;r.insertBefore(e,s),e=t}}}return a},_={},E=(e,t=_)=>e._$AH=t,T=e=>e._$AH,f=e=>(e=>null!=e?._$litType$?.h)(e)?e._$litType$.h:e.strings,A=o(class extends n{constructor(e){super(e),this.tt=new WeakMap}render(e){return[e]}update(e,[r]){const s=u(this.et)?f(this.et):null,i=u(r)?f(r):null;if(null!==s&&(null===i||s!==i)){const r=T(e).pop();let i=this.tt.get(s);if(void 0===i){const e=document.createDocumentFragment();i=t(a,e),i.setConnected(!1),this.tt.set(s,i)}E(i,[r]),m(i,void 0,r)}if(null!==i){if(null===s||s!==i){const t=this.tt.get(i);if(void 0!==t){const a=T(t).pop();(e=>{e._$AR()})(e),m(e,void 0,a),E(e,[a])}}this.et=r}else this.et=void 0;return this.render(r)}});var b=r`
html, :host {
  --font-family: "Noto Sans";
  --font-size: 14px;
  --menu-top-height: 43.5px;
  --menu-side-width: 250px;
  --light: 255, 255, 255;
  --dark: 0, 0, 2;
}
html, :host, *[theme="light"] {
  --neutral-1: 0, 0, 0;
  --neutral-2: 255, 255, 255;
  --accent-1: 0, 28, 50;
  --accent-1b: 0, 71, 93;
  --accent-1c: 255, 255, 255;
  --base-0: 255, 255, 255;
  --base-1: 235, 235, 235;
  --base-2: 245, 245, 245;
  --base-3: 255, 255, 255;
  --base-4: 255, 255, 255;
  --functional-blue: 20, 120, 220;
  --functional-red: 210, 105, 125;
  --functional-yellow: 220, 168, 40;
  --functional-green: 50, 168, 40;
  --text-1: rgba(0, 0, 0, .90);
  --text-2: rgba(0, 0, 0, .60);
  --text-3: rgba(0, 0, 0, .20);
  --shadow-1: 0 2px 8px rgba(0,0,0,.1), 0 1px 4px rgba(0,0,0,.05);
}
*[theme="dark"] {
  --neutral-1: 255, 255, 255;
  --neutral-2: 0, 0, 0;
  --accent-1: 0, 28, 50;
  --accent-1b: 0, 71, 93;
  --accent-1c: 255, 255, 255;
  --base-0: 0, 0, 2;
  --base-1: 15, 15, 15;
  --base-2: 25, 25, 25;
  --base-3: 35, 35, 35;
  --base-4: 45, 45, 45;
  --functional-blue: 20, 120, 220;
  --functional-red: 210, 105, 125;
  --functional-yellow: 220, 160, 40;
  --functional-green: 40, 160, 40;
  --text-1: rgba(255, 255, 255, .90);
  --text-2: rgba(255, 255, 255, .60);
  --text-3: rgba(255, 255, 255, .20);
  --shadow-1: 0 2px 8px rgba(0,0,0,.2), 0 1px 4px rgba(0,0,0,.15);
}
`;const v={LIGHT:"light",DARK:"dark"},S={LOGIN:"login",SEARCH:"search",EDIT:"edit",SETTING:"setting",HELP:"help",BOOKMARK:"bookmark",TEMPLATE:"template"},y={DEFAULT:"default",PRIMARY:"primary",BORDER:"border"},I={LEFT:"left",CENTER:"center",RIGHT:"right"},D={TEXT:"text",COLOR:"color",FILE:"file",PASSWORD:"password"},w={DATE:"date",TIME:"time",DATETIME:"datetime-local"},O={INFO:"info",ERROR:"error",SUCCESS:"success"},k={AUTO:"auto",SHOW:"show",HIDE:"hide"},C={EDIT:"edit",NEW:"new"},R={SIDEBAR:"sidebar",MODULE:"module",SCROLL:"scroll"},M={BACK:"back",CHANGE:"change",QUICK:"quick",BROWSER:"browser",CHECK:"check",PREV_NUMBER:"prev_number",NEXT_NUMBER:"next_number",SAVE:"save",DELETE:"delete",NEW:"new",COPY:"copy",LINK:"link",PASSWORD:"password",SHIPPING_ADD_ALL:"shipping_add_all",SHIPPING_CREATE:"shipping_create",REPORT_SETTINGS:"report_settings",SEARCH_QUEUE:"search_queue",EXPORT_QUEUE_ALL:"export_queue_all",CREATE_REPORT:"create_report",EXPORT_EVENT:"export_event",SAVE_BOOKMARK:"save_bookmark",LOAD_SETTING:"load_setting",PROGRAM_SETTING:"program_setting",PASSWORD_FORM:"password_form",HELP:"help",BLANK:"blank",SAMPLE:"sample"},x={CHANGE:"change",LOGIN:"login",THEME:"theme",LANG:"lang"},N={CANCEL:"cancel",DELETE:"delete",SELECTED:"selected",OK:"ok",SEARCH:"search",CURRENT_PAGE:"current_page"},F={TOP:"top",BOTTOM:"bottom",ALL:"all",NONE:"none"},$={BOOKMARK:"bookmark",HISTORY:"history"},L={CHANGE:"change",BROWSER_VIEW:"browser_view",BOOKMARK_SAVE:"bookmark_save",EXPORT_RESULT:"export_result",SHOW_HELP:"show_help",SHOW_BROWSER:"show_browser",ADD_FILTER:"add_filter",SHOW_TOTAL:"show_total",SET_COLUMNS:"set_columns",EDIT_FILTER:"edit_filter",DELETE_FILTER:"delete_filter",SET_FORM_ACTION:"set_form_action",EDIT_CELL:"edit_cell",CURRENT_PAGE:"current_page"},P=[["===","EQUAL"],["==N","IS NULL"],["!==","NOT EQUAL"],[">==",">="],["<==","<="]],q={CHANGE:"change",CHECK_EDITOR:"check_editor",CHECK_TRANSTYPE:"check_transtype",EDIT_ITEM:"edit_item",SET_PATTERN:"set_pattern",SELECTOR:"selector",FORM_ACTION:"form_action"},U={LOAD_EDITOR:"load_editor",SET_EDITOR:"set_editor",SET_EDITOR_ITEM:"set_editor_item",LOAD_FORMULA:"load_formula",NEW_FIELDVALUE:"new_fieldvalue",CREATE_TRANS:"create_trans",CREATE_TRANS_OPTIONS:"create_trans_options",FORM_ACTION:"form_action"},G={LOAD_EDITOR:"load_editor",NEW_EDITOR_ITEM:"new_editor_item",EDIT_EDITOR_ITEM:"edit_editor_item",DELETE_EDITOR_ITEM:"delete_editor_item",LOAD_SHIPPING:"load_shipping",ADD_SHIPPING_ROW:"add_shipping_row",SHOW_SHIPPING_STOCK:"show_shipping_stock",EDIT_SHIPPING_ROW:"edit_shipping_row",DELETE_SHIPPING_ROW:"delete_shipping_row",EXPORT_QUEUE_ITEM:"export_queue_item",NEW_ITEM:"new_item",EDIT_ITEM:"edit_item",DELETE_ITEM:"delete_item",EDIT_AUDIT:"edit_audit",EDIT_MENU_FIELD:"edit_menu_field",DELETE_ITEM_ROW:"delete_item_row"},H={EDIT_ITEM:"edit_item",FORM_ACTION:"form_action",CURRENT_PAGE:"current_page"},B={TEXT:"string",LIST:"list",TABLE:"table"},j={ADD_ITEM:"add_item",CHANGE_TEMPLATE:"change_template",CHANGE_CURRENT:"change_current",GO_PREVIOUS:"go_previous",GO_NEXT:"go_next",CREATE_MAP:"create_map",SET_CURRENT:"set_current",MOVE_UP:"move_up",MOVE_DOWN:"move_down",DELETE_ITEM:"delete_item",EDIT_ITEM:"edit_item",EDIT_DATA_ITEM:"edit_data_item",SET_CURRENT_DATA:"set_current_data",SET_CURRENT_DATA_ITEM:"set_current_data_item",ADD_TEMPLATE_DATA:"add_template_data",DELETE_DATA:"delete_data",DELETE_DATA_ITEM:"delete_data_item"},K=r`
html, :host {
  display: block;
  font-family: var(--font-family);
  font-size: var(--font-size);
  color: var(--text-1);
  fill: var(--text-1);
  width: 100%;
  height: 100%;
  margin: 0;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
.main {
  background-color: rgb(var(--base-0));
  width: 100%;
  position: absolute;
  left: 0;
  top: var(--menu-top-height);
}
.client-menubar {
  z-index: 5;
  position: fixed;
  width: 100%;
  top: 0;
}
*::-webkit-scrollbar {
  width: 10px;
  height: 5px;
  background-color: transparent;
  visibility: hidden;
}
*::-webkit-scrollbar-track {
  background-color: rgba(var(--accent-1), .05);
  border-radius: 8px;
}
*::-webkit-scrollbar-thumb {
  background-color: rgba(var(--accent-1), .60);
  border-radius: 8px;
}
*::-webkit-scrollbar-thumb:active,
*::-webkit-scrollbar-thumb:hover {
  background-color: rgba(var(--accent-1), .20)
}
`;class V{constructor(e,t){this.host=e,this._data=t,e.addController(this)}get data(){return this._data}set data(e){const{key:t,value:a,update:r}=e;this._data[t]&&"object"==typeof a&&(this._data[t]={...this._data[t],...a},!1!==r&&this.host.requestUpdate())}}const W=(e,t)=>{let a=0;const r=t=>{let a=t;switch(e){case"sqlite":case"sqlite3":a=a.replace(/{CCS}/g,""),a=a.replace(/{SEP}/g,"||"),a=a.replace(/{CCE}/g,""),a=a.replace(/{CAS_TEXT}/g,"cast("),a=a.replace(/{CAE_TEXT}/g," as text)"),a=a.replace(/{CAS_INT}/g,"cast("),a=a.replace(/{CAE_INT}/g," as integer)"),a=a.replace(/{CAS_FLOAT}/g,"cast("),a=a.replace(/{CAE_FLOAT}/g," as double)"),a=a.replace(/{CAS_DATE}/g,""),a=a.replace(/{CASF_DATE}/g,""),a=a.replace(/{CAE_DATE}/g,""),a=a.replace(/{CAEF_DATE}/g,""),a=a.replace(/{FMSF_NUMBER}/g,""),a=a.replace(/{FMSF_DATE}/g,""),a=a.replace(/{FMEF_CONVERT}/g,""),a=a.replace(/{FMS_INT}/g,""),a=a.replace(/{FME_INT}/g,""),a=a.replace(/{FMS_DATE}/g,"substr("),a=a.replace(/{FME_DATE}/g,",1,10)"),a=a.replace(/{FMS_DATETIME}/g,"substr("),a=a.replace(/{FME_DATETIME}/g,",1,19)"),a=a.replace(/{FMS_TIME}/g,"substr(time("),a=a.replace(/{FME_TIME}/g,"),0,6)"),a=a.replace(/{JOKER}/g,"'%'"),a=a.replace(/{CUR_DATE}/g,"date('now')");break;case"mysql":a=a.replace(/{CCS}/g,"concat("),a=a.replace(/{SEP}/g,","),a=a.replace(/{CCE}/g,")"),a=a.replace(/{CAS_TEXT}/g,"cast("),a=a.replace(/{CAE_TEXT}/g," as char)"),a=a.replace(/{CAS_INT}/g,"cast("),a=a.replace(/{CAE_INT}/g," as signed)"),a=a.replace(/{CAS_FLOAT}/g,"cast("),a=a.replace(/{CAE_FLOAT}/g," as decimal)"),a=a.replace(/{CAS_DATE}/g,"cast("),a=a.replace(/{CASF_DATE}/g,"cast("),a=a.replace(/{CAE_DATE}/g," as date)"),a=a.replace(/{CAEF_DATE}/g," as date)"),a=a.replace(/{FMSF_NUMBER}/g,""),a=a.replace(/{FMSF_DATE}/g,""),a=a.replace(/{FMEF_CONVERT}/g,""),a=a.replace(/{FMS_INT}/g,"format(cast("),a=a.replace(/{FME_INT}/g," as signed), 0)"),a=a.replace(/{FMS_DATE}/g,"date_format("),a=a.replace(/{FME_DATE}/g,", '%Y-%m-%d')"),a=a.replace(/{FMS_DATETIME}/g,"date_format("),a=a.replace(/{FME_DATETIME}/g,", '%Y-%m-%dT%H:%i:%s')"),a=a.replace(/{FMS_TIME}/g,"cast(cast("),a=a.replace(/{FME_TIME}/g," as time) as char)"),a=a.replace(/{JOKER}/g,"'%'"),a=a.replace(/{CUR_DATE}/g,"current_date");break;case"postgres":a=a.replace(/{CCS}/g,""),a=a.replace(/{SEP}/g,"||"),a=a.replace(/{CCE}/g,""),a=a.replace(/{CAS_TEXT}/g,"cast("),a=a.replace(/{CAE_TEXT}/g," as text)"),a=a.replace(/{CAS_INT}/g,"cast("),a=a.replace(/{CAE_INT}/g," as integer)"),a=a.replace(/{CAS_FLOAT}/g,"cast("),a=a.replace(/{CAE_FLOAT}/g," as float)"),a=a.replace(/{CAS_DATE}/g,"cast("),a=a.replace(/{CASF_DATE}/g,"cast("),a=a.replace(/{CAE_DATE}/g," as date)"),a=a.replace(/{CAEF_DATE}/g," as date)"),a=a.replace(/{FMSF_NUMBER}/g,"case when rf_number.fieldname is null then 0 else "),a=a.replace(/{FMSF_DATE}/g,"case when rf_date.fieldname is null then current_date else "),a=a.replace(/{FMEF_CONVERT}/g," end "),a=a.replace(/{FMS_INT}/g,"to_char(cast("),a=a.replace(/{FME_INT}/g," as integer), '999,999,999')"),a=a.replace(/{FMS_DATE}/g,"to_char("),a=a.replace(/{FME_DATE}/g,", 'YYYY-MM-DD')"),a=a.replace(/{FMS_DATETIME}/g,"to_char("),a=a.replace(/{FME_DATETIME}/g,", 'YYYY-MM-DD\"T\"HH24:MI:SS')"),a=a.replace(/{FMS_TIME}/g,"substr(cast(cast("),a=a.replace(/{FME_TIME}/g," as time) as text), 0, 6)"),a=a.replace(/{JOKER}/g,"chr(37)"),a=a.replace(/{CUR_DATE}/g,"current_date");break;case"mssql":a=a.replace(/{CCS}/g,""),a=a.replace(/{SEP}/g,"+"),a=a.replace(/{CCE}/g,""),a=a.replace(/{CAS_TEXT}/g,"cast("),a=a.replace(/{CAE_TEXT}/g," as nvarchar)"),a=a.replace(/{CAS_INT}/g,"cast("),a=a.replace(/{CAE_INT}/g," as int)"),a=a.replace(/{CAS_FLOAT}/g,"cast("),a=a.replace(/{CAE_FLOAT}/g," as real)"),a=a.replace(/{CAS_DATE}/g,"cast("),a=a.replace(/{CASF_DATE}/g,""),a=a.replace(/{CAE_DATE}/g," as date)"),a=a.replace(/{CAEF_DATE}/g,""),a=a.replace(/{FMSF_NUMBER}/g,""),a=a.replace(/{FMSF_DATE}/g,""),a=a.replace(/{FMEF_CONVERT}/g,""),a=a.replace(/{FMS_INT}/g,"replace(convert(varchar,cast("),a=a.replace(/{FME_INT}/g," as money),1), '.00','')"),a=a.replace(/{FMS_DATE}/g,"convert(varchar(10),"),a=a.replace(/{FME_DATE}/g,", 120)"),a=a.replace(/{FMS_DATETIME}/g,"convert(varchar(19),"),a=a.replace(/{FME_DATETIME}/g,", 120)"),a=a.replace(/{FMS_TIME}/g,"SUBSTRING(cast(cast("),a=a.replace(/{FME_TIME}/g," as time) as nvarchar),0,6)"),a=a.replace(/{JOKER}/g,"'%'"),a=a.replace(/{CUR_DATE}/g,"cast(GETDATE() as DATE)")}return a},s=(t,r)=>{let i=t,o="";if(Array.isArray(i)){let e=" ",t="",a="";return i.length>0&&(["select","select_distinct","union_select","order_by","group_by"].includes(r)||0===i[0].length)&&(e=", "),i.forEach((n=>{let l=n;null==l&&(l="null"),0===l.length?"set"!==r&&(t="(",a=")"):2!==i.length||"and"!==l&&"or"!==l?r&&1===i.length&&"object"==typeof i[0]?o+=` (${s(l,r)})`:o+=e+s(l,r):(o+=`${l} (`,a=")")})),", "===e&&(o=o.substr(2)),r&&i.includes("on")&&(o=r.replace("_"," ")+o),t+o.toString().trim()+a}return"object"==typeof i?(Object.keys(i).forEach((e=>{o+="inner_join"===e||"left_join"===e?` ${s(i[e],e)}`:` ${e.replace("_"," ")} ${s(i[e],e)}`})),o):(i.includes("?")&&"select"!==r&&(a+=1,"postgres"===e&&(i=i.replace("?",`$${a}`))),i)};return"string"==typeof t?{sql:r(t),prmCount:a}:{sql:r(s(t)),prmCount:a}},z=async(e,t)=>(e=>{if(401===e.status)return{code:401,message:"Unauthorized"};if(400===e.status)return e.json();if(204===e.status||205===e.status)return null;switch(e.headers.get("content-type").split(";")[0]){case"application/pdf":case"application/xml":case"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":return e.blob();case"application/json":return e.json();case"text/plain":case"text/csv":return e.text();default:return e}})((e=>{if(e.status>=200&&e.status<300||400===e.status||401===e.status)return e;const t=new Error(e.statusText);throw t.response=e,t})(await fetch(e,t))),X=(e,t)=>{const a=document.createElement("a");a.href=e,a.download=t||e,document.body.appendChild(a),a.click()},Y={usergroup_filter:{transitem:["t.cruser_id","in",[{select:["id"],from:"employee",where:["usergroup","=","?"]}]],transmovement:["t.cruser_id","in",[{select:["id"],from:"employee",where:["usergroup","=","?"]}]],transpayment:["t.cruser_id","in",[{select:["id"],from:"employee",where:["usergroup","=","?"]}]]},employee_filter:{transitem:["t.cruser_id","=","?"],transmovement:["t.cruser_id","=","?"],transpayment:["t.cruser_id","=","?"]}};class J{constructor(e){this.host=e,this.modules={},this.getSql=W,this.request=z,this.saveToDisk=X,this.createHistory=this.createHistory.bind(this),this.currentModule=this.currentModule.bind(this),this.getAuditFilter=this.getAuditFilter.bind(this),this.getDataFilter=this.getDataFilter.bind(this),this.getSetting=this.getSetting.bind(this),this.getUserFilter=this.getUserFilter.bind(this),this.loadBookmark=this.loadBookmark.bind(this),this.msg=this.msg.bind(this),this.quickSearch=this.quickSearch.bind(this),this.requestData=this.requestData.bind(this),this.resultError=this.resultError.bind(this),this.saveBookmark=this.saveBookmark.bind(this),this.showHelp=this.showHelp.bind(this),this.showToast=this.showToast.bind(this),this.signOut=this.signOut.bind(this),this.tokenLogin=this.tokenLogin.bind(this),e.addController(this)}hostConnected(){const{state:e,setData:t}=this.host;this.store={data:e.data,setData:t},this._loadConfig(window.location)}async _loadConfig(e){const{data:t,setData:a}=this.store;this.request(`${t[S.LOGIN].server}/config`,{method:"GET",headers:{"Content-Type":"application/json"}}).then((e=>{let r={...t.session};e.locales&&"object"==typeof e.locales&&(r={...r,locales:{...r.locales,...e.locales}}),a("session",r),localStorage.getItem("lang")&&r.locales[localStorage.getItem("lang")]&&localStorage.getItem("lang")!==t.current.lang&&a("current",{lang:localStorage.getItem("lang")})})).catch((e=>{a("error",e),this.showToast(O.ERROR,e.message)}));const[r,s]=(()=>{const t=e=>{const t={};return e.split("&").forEach((e=>{const a=String(e).indexOf("="),r=String(e).substring(0,a>0?a:String(e).length),s=a>-1&&a<String(e).length?String(e).substring(a+1):"";t[r]=s})),t};if(e.hash)return["hash",t(e.hash.substring(1))];if(e.search)return["search",t(e.search.substring(1))];const a=e.pathname.substring(1).split("/");return[a[0],a.slice(1)]})();"hash"!==r||!s.access_token&&!s.code?import("./module-A7P2DJ_v.js"):this.tokenLogin(s)}async createHistory(e){const{data:t,setData:a}=this.store;let r={datetime:`${(new Date).toISOString().slice(0,10)}T${(new Date).toLocaleTimeString("en",{hour12:!1}).replace("24","00")}`,type:e,type_title:this.msg("",{id:`label_${e}`}),ntype:t[S.EDIT].current.type,transtype:t[S.EDIT].current.transtype||"",id:t[S.EDIT].current.item.id},s="trans"===r.ntype?`${t[S.EDIT].template.options.title} | ${t[S.EDIT].current.item[t[S.EDIT].template.options.title_field]}`:t[S.EDIT].template.options.title;"trans"!==r.ntype&&void 0!==t[S.EDIT].template.options.title_field&&(s+=` | ${t[S.EDIT].current.item[t[S.EDIT].template.options.title_field]}`),r={...r,title:s};const i={...t.bookmark};let o={};if(i.history){o={...i.history,cfgroup:`${(new Date).toISOString().slice(0,10)}T${(new Date).toLocaleTimeString("en",{hour12:!1}).replace("24","00")}`};let e=JSON.parse(o.cfvalue);e.unshift(r),e.length>parseInt(this.getSetting("history"),10)&&(e=e.slice(0,parseInt(this.getSetting("history"),10))),o={...o,cfname:e.length,cfvalue:JSON.stringify(e)}}else o={...o,employee_id:t.login.data.employee.id,section:"history",cfgroup:`${(new Date).toISOString().slice(0,10)}T${(new Date).toLocaleTimeString("en",{hour12:!1}).replace("24","00")}`,cfname:1,cfvalue:JSON.stringify([r])};const n={method:"POST",data:[o]},l=await this.requestData("/ui_userconfig",n);return l.error?this.resultError(l):a("bookmark",{history:o})}async currentModule({data:e,content:t}){const{setData:a}=this.store,r={forms:async()=>{const{Forms:e}=await import("./module-0Mj2q8HM.js"),{Dataset:t}=await import("./module-GIo78Gda.js"),{Sql:a}=await import("./module-DaIb_Znz.js"),{InitItem:r,Validator:s}=await import("./module-_yGp1sO4.js");this.modules.forms=e(this),this.modules.dataSet={...t},this.modules.initItem=r(this),this.modules.sql=a(this),this.modules.validator=s(this)},quick:async()=>{const{Quick:e}=await import("./module-BHgGaoC6.js"),{Queries:t}=await import("./module-bnHj77YI.js");this.modules.quick={...e},this.modules.queries=t(this)},search:async()=>{const{SearchController:e}=await import("./module-xKTUziSw.js");await import("./module-WNH4Rd3J.js"),await import("./module--n2Fxrja.js"),this.modules.quick||await r.quick(),this.modules.search=new e(this.host)},edit:async()=>{const{EditController:e}=await import("./module-kLSFPyys.js");await import("./module-0JE8O3kq.js"),this.modules.forms||await r.forms(),this.modules.edit=new e(this.host)},setting:async()=>{const{SettingController:e}=await import("./module-w8YmkvAJ.js");await import("./module-Eb2IWf8n.js"),this.modules.forms||await r.forms(),this.modules.setting=new e(this.host)},template:async()=>{const{TemplateController:e}=await import("./module-qc2ked0c.js");await import("./module-cE0lINgO.js"),this.modules.template=new e(this.host)}};this.modules[e.module]||await r[e.module](),a("current",{...e}),t&&this.modules[e.module]&&this.modules[e.module][t.fkey](...t.args)}getAuditFilter(e,t){const a=this.store.data[S.LOGIN];let r,s=["all",1];switch(e){case"trans":case"menu":r=a.data.audit.filter((a=>a.nervatypeName===e&&a.subtypeName===t))[0];break;case"report":r=a.data.audit.filter((a=>a.nervatypeName===e&&a.subtype===t))[0];break;default:r=a.data.audit.filter((t=>t.nervatypeName===e))[0]}return void 0!==r&&(s=[r.inputfilterName,r.supervisor]),s}getDataFilter(e,t,a){let r=t;return"transitem"===e&&("disabled"===this.getAuditFilter("trans","offer")[0]&&(r=r.concat(["and",["tg.groupvalue","<>","'offer'"]])),"disabled"===this.getAuditFilter("trans","order")[0]&&(r=r.concat(["and",["tg.groupvalue","<>","'order'"]])),"disabled"===this.getAuditFilter("trans","worksheet")[0]&&(r=r.concat(["and",["tg.groupvalue","<>","'worksheet'"]])),"disabled"===this.getAuditFilter("trans","rent")[0]&&(r=r.concat(["and",["tg.groupvalue","<>","'rent'"]])),"disabled"===this.getAuditFilter("trans","invoice")[0]&&(r=r.concat(["and",["tg.groupvalue","<>","'invoice'"]]))),"transpayment"===e&&("disabled"===this.getAuditFilter("trans","bank")[0]&&(r=r.concat(["and",["tg.groupvalue","<>","'bank'"]])),"disabled"===this.getAuditFilter("trans","cash")[0]&&(r=r.concat(["and",["tg.groupvalue","<>","'cash'"]]))),"transmovement"===e&&"InventoryView"!==a&&("disabled"===this.getAuditFilter("trans","delivery")[0]&&(r=r.concat(["and",["tg.groupvalue","<>","'delivery'"]])),"disabled"===this.getAuditFilter("trans","inventory")[0]&&(r=r.concat(["and",["tg.groupvalue","<>","'inventory'"]]))),r}getSetting(e){const{ui:t}=this.store.data;if("ui"===e){const e={...t};return Object.keys(e).forEach((t=>{localStorage.getItem(t)&&(e[t]=localStorage.getItem(t))})),e}return localStorage.getItem(e)||t[e]||""}getUserFilter(e){const t=this.store.data[S.LOGIN],a={params:[],where:[]};return"usergroup"===t.data.transfilterName&&void 0!==Y.usergroup_filter[e]&&(a.where=["and",Y.usergroup_filter[e]],a.params=[t.data.employee.usergroup]),"own"===t.data.transfilterName&&void 0!==Y.employee_filter[e]&&(a.where=["and",Y.employee_filter[e]],a.params=[t.data.employee.id]),a}async loadBookmark(e){const{setData:t}=this.store,a=await this.requestData("/ui_userconfig",{token:e.token,query:{filter:`employee_id;==;${e.user_id}`}});if(a.error)return this.resultError(a),null;const r=a.filter((e=>"bookmark"===e.section)),s=a.filter((e=>"history"===e.section))[0]||null;return t(S.BOOKMARK,{bookmark:r,history:s}),{bookmark:r,history:s}}msg(e,t){let a=e;const{locales:r}=this.store.data.session,{lang:s}=this.store.data.current;return r[s]&&r[s][t.id]?a=r[s][t.id]:"en"!==s&&r.en[t.id]&&(a=r.en[t.id]),a}async quickSearch(e,t){const{quick:a}=this.modules,r=this.store.data[S.LOGIN],s=a[e](String(r.data.employee.usergroup));let i={...s.sql},o=[],n=[];if(""!==t){const e="{CCS}{JOKER}{SEP}lower(?){SEP}{JOKER}{CCE} ";s.columns.forEach(((a,r)=>{n.push([0!==r?"or":"",[`lower(${a[1]})`,"like",e]]),o.push(t)})),n=["and",[n]]}n=this.getDataFilter(e,n),n.length>0&&(i={...i,where:[...i.where,n]});const l=this.getUserFilter(e);l.where.length>0&&(i={...i,where:[...i.where,l.where]},o=o.concat(l.params));const c={method:"POST",data:[{key:"result",text:W(r.data.engine,i).sql,values:o}]};return this.requestData("/view",c)}async requestData(e,t,a){const{data:r,setData:s}=this.store;let i=t;try{a||s("current",{request:!0});let t="SERVER"===r.session.serverURL?r.session.apiPath+e:r.login.server+e;const o=r.login.data?r.login.data.token:i.token||"";if(i.headers||(i={...i,headers:{}}),i={...i,headers:{...i.headers,"Content-Type":"application/json"}},""!==o&&(i={...i,headers:{...i.headers,Authorization:`Bearer ${o}`}}),i.data&&(i={...i,body:JSON.stringify(i.data)}),i.query){const e=new URLSearchParams;Object.keys(i.query).forEach((t=>{e.append(t,i.query[t])})),t+=`?${e.toString()}`}const n=await this.request(t,i);return a||s("current",{request:!1}),n&&n.code?(401===n.code&&this.signOut(),{error:{message:n.message},data:null}):n}catch(e){return a||s("current",{request:!1}),{error:{message:e.message},data:null}}}resultError(e){const{setData:t}=this.store;return e.error&&t("error",e.error),e.error&&e.error.message?this.showToast(O.ERROR,e.error.message):this.showToast(O.ERROR,this.msg("Internal Server Error",{id:"error_internal"})),!1}saveBookmark(e){const{inputBox:t}=this.host,{data:a,setData:r}=this.store,s=a[S.LOGIN],i=a[S.SEARCH],o=a[S.EDIT],n=t({title:this.msg("",{id:"msg_bookmark_new"}),message:this.msg("",{id:"msg_bookmark_name"}),value:"browser"===e[0]?e[1]:o.current.item[e[2]],showValue:!0,onEvent:{onModalEvent:async t=>{if(r("current",{modalForm:null}),t.key===N.OK&&""!==t.data.value){let a={employee_id:s.data.employee.id,section:"bookmark",cfgroup:e[0]};a="browser"===e[0]?{...a,cfname:t.data.value,cfvalue:JSON.stringify({date:(new Date).toISOString().split("T")[0],vkey:i.vkey,view:i.view,filters:i.filters[i.view],columns:i.columns[i.view]})}:{...a,cfname:t.data.value,cfvalue:JSON.stringify({date:(new Date).toISOString().split("T")[0],ntype:o.current.type,transtype:o.current.transtype,id:o.current.item.id,info:"trans"===o.current.type?null!==o.dataset.trans[0].custname?o.dataset.trans[0].custname:o.current.item.transdate:o.current.item[e[3]]})};const r={method:"POST",data:[a]},n=await this.requestData("/ui_userconfig",r);if(n.error)return this.resultError(n);this.loadBookmark({user_id:s.data.employee.id})}return!0}}});r("current",{modalForm:n,side:k.HIDE})}showHelp(e){const{data:t}=this.store,a=document.createElement("a");a.setAttribute("href",t.session.helpPage+e),a.setAttribute("target","_blank"),document.body.appendChild(a),a.click()}showToast(e,t,a){const{current:r}=this.store.data,s=void 0!==a?a:this.getSetting("toastTimeout");r.toast&&r.toast.show({type:e,value:t,timeout:s})}signOut(){const{data:e,setData:t}=this.store;e[S.LOGIN].callback?window.location.replace(e[S.LOGIN].callback):t(S.LOGIN,{data:null,token:null})}async tokenLogin(e){const{LoginController:t}=await import("./module-gbRhGoFt.js"),a=new t(this.host);e.access_token&&a.tokenValidation(e),e.code&&a.setCodeToken(e)}}const Q={session:{version:"5.2.4",locales:d,serverURL:"SERVER",apiPath:"/api",engines:["sqlite","sqlite3","mysql","postgres","mssql"],service:["dev","5.2.2","5.2.3","5.2.4"],helpPage:"https://nervatura.github.io/nervatura/docs/client/"},ui:{toastTimeout:4,paginationPage:10,selectorPage:5,history:5,timeIntervals:15,searchSubtract:3,export_sep:";",page_size:"a4",page_orient:"portrait",printqueue_mode:[["print","printqueue_mode_print"],["pdf","printqueue_mode_pdf"],["xml","printqueue_mode_xml"]],printqueue_type:[["customer","title_customer"],["product","title_product"],["employee","title_employee"],["tool","title_tool"],["project","title_project"],["order","title_order"],["offer","title_offer"],["invoice","title_invoice"],["receipt","title_receipt"],["rent","title_rent"],["worksheet","title_worksheet"],["delivery","title_delivery"],["inventory","title_inventory"],["waybill","title_waybill"],["production","title_production"],["formula","title_formula"],["bank","title_bank"],["cash","title_cash"]],report_orientation:[["portrait","report_portrait"],["landscape","report_landscape"]],report_size:[["a3","A3"],["a4","A4"],["a5","A5"],["letter","Letter"],["legal","Legal"]]},current:{home:S.SEARCH,module:S.LOGIN,side:k.AUTO,lang:localStorage.getItem("lang")&&d[localStorage.getItem("lang")]?localStorage.getItem("lang"):"en",theme:localStorage.getItem("theme")||v.LIGHT},login:{username:localStorage.getItem("username")||"",database:localStorage.getItem("database")||"",server:localStorage.getItem("server")&&""!==localStorage.getItem("server")?localStorage.getItem("server"):"nervatura.github.io"!==window.location.hostname?window.location.origin+"/api":""},search:{seltype:"selector",group_key:"transitem",result:[],vkey:null,qview:"transitem",qfilter:"",filters:{},columns:{},browser_filter:!0},edit:{dataset:{},current:{},dirty:!1,form_dirty:!1},setting:{dirty:!1,result:[]},template:{dirty:!1},bookmark:{history:null,bookmark:[]}},Z=r`
.modal {
  z-index: 10;
  position: fixed;
  left: 0;
  top: 0;
  width:100%;
  height: 100%;
  overflow: auto;
  background-color: rgba(0, 0, 0, 0.129);
  padding: 10px 5px;
}
.middle {
  z-index: 20;
  margin: 0px;
  position:absolute;
  top:50%;
  left:50%; 
  background: #222222;
  padding: 0px 30px;
  border-radius: 5px;
  border: 1px solid var(--text-1);
  opacity: 0.75;
  transform:translate(-50%,-50%);
  -ms-transform:translate(-50%,-50%);
}
@keyframes lds-roller {
  0% {
    transform: rotate(0deg);
}
  100% {
    transform: rotate(360deg);
}
}
.loading {
  margin: 2em auto;
  position: relative;
  width: 64px;
  height: 64px;
}
.loading div {
  animation: lds-roller 1.2s cubic-bezier(0.5, 0, 0.5, 1) infinite;
  transform-origin: 32px 32px;
}
.loading div:after {
  content: " ";
  display: block;
  position: absolute;
  width: 7px;
  height: 7px;
  border-radius: 50%;
  background: rgb(var(--functional-blue));
  margin: -3px 0 0 -3px;
}
.loading div:nth-child(1) {
  animation-delay: -0.036s;
}
.loading div:nth-child(1):after {
  top: 50px;
  left: 50px;
}
.loading div:nth-child(2) {
  animation-delay: -0.072s;
}
.loading div:nth-child(2):after {
  top: 54px;
  left: 45px;
}
.loading div:nth-child(3) {
  animation-delay: -0.108s;
}
.loading div:nth-child(3):after {
  top: 57px;
  left: 39px;
}
.loading div:nth-child(4) {
  animation-delay: -0.144s;
}
.loading div:nth-child(4):after {
  top: 58px;
  left: 32px;
}
.loading div:nth-child(5) {
  animation-delay: -0.18s;
}
.loading div:nth-child(5):after {
  top: 57px;
  left: 25px;
}
.loading div:nth-child(6) {
  animation-delay: -0.216s;
}
.loading div:nth-child(6):after {
  top: 54px;
  left: 19px;
}
.loading div:nth-child(7) {
  animation-delay: -0.252s;
}
.loading div:nth-child(7):after {
  top: 50px;
  left: 14px;
}
.loading div:nth-child(8) {
  animation-delay: -0.288s;
}
.loading div:nth-child(8):after {
  top: 45px;
  left: 10px;
}
`;customElements.define("form-spinner",class extends s{render(){return i`
    <div class="modal" >
      <div class="middle" >
        <div class="loading">
          <div></div>
          <div></div>
          <div></div>
          <div></div>
          <div></div>
          <div></div>
          <div></div>
          <div></div>
        </div>
      </div>
    </div>`}static get styles(){return[Z]}});const ee=r`
:host {
	font-family: var(--font-family);
	font-size: var(--font-size);
	color: var(--text-1);
	fill: var(--text-1);
}
div {
	--toast-background: rgba(var(--functional-yellow), 1);
}
div[type="error"] {
	--toast-background: rgba(var(--functional-red), 1);
}
div[type="success"] {
	--toast-background: rgba(var(--functional-green), 1);
}
div {
	top: 20px;
	right: 20px;
	position: fixed;
	z-index: 10001;
	contain: layout;
	max-width: 330px;
	box-shadow: 0 5px 5px -3px rgba(0, 0, 0, 0.2), 0 8px 10px 1px rgba(0, 0, 0, 0.14), 0 3px 14px 2px rgba(0, 0, 0, 0.12);
	border-left: 3px solid var(--text-1);
	display: flex;
	align-items: center;
	word-break: break-word;
	font-size: 14px;
	line-height: 20px;
	padding: 15px;
	transition: transform 0.3s, opacity 0.4s;
	opacity: 1;
	transform: translate3d(0, 0, 0);
	background: var(--toast-background);
	border-radius: 5px;
	cursor: pointer;
}
.icon {
	margin-right: 10px;
	width: 32px;
}
`;customElements.define("form-toast",class extends s{constructor(){super(),this.id=Math.random().toString(36).slice(2),this.name=void 0,this.type=O.INFO,this.hidden=!0,this.timeout=4,this.style={},this._iconMap={info:"InfoCircle",error:"ExclamationTriangle",success:"CheckSquare"}}static get properties(){return{id:{type:String},name:{type:String,reflect:!0},type:{type:String},hidden:{type:Boolean},timeout:{type:Number},style:{type:Object}}}connectedCallback(){super.connectedCallback(),this.setData&&this.setData("current",{toast:this},!1)}disconnectedCallback(){this.setData&&this.setData("current",{toast:null},!1),super.disconnectedCallback()}show({type:e,value:t,timeout:a}){this.hidden&&(this.value=t||this.value,this.type=e||this.type,this.timeout=void 0!==a?a:this.timeout,this.hidden=!1,this.timeout>0&&(this.timeoutVar=setTimeout((()=>{this&&!this.hidden&&(this.hidden=!0)}),1e3*this.timeout)))}close(){this.hidden||(clearTimeout(this.timeoutVar),this.hidden=!0)}render(){return this.hidden?a:i`<div 
        id="${this.id}"
        type="${this.type}"
        name="${l(this.name)}"
        @click=${this.close}
        style="${c(this.style)}" >
          <span class="icon">
            ${i`<form-icon 
              iconKey="${this._iconMap[this.type]||"InfoCircle"}"
              .style=${{margin:"auto"}} 
              width=32 height=32 ></form-icon>`}
          </span>
          <slot id="value">${l(this.value)}</slot>
      </div>`}static get styles(){return[ee]}});customElements.define("nervatura-client",class extends s{static get properties(){return{}}static get styles(){return[b,K]}constructor(){super(),this.state=new V(this,Q),this.app=new J(this),this.inputBox=this.inputBox.bind(this),this.setData=this.setData.bind(this)}connectedCallback(){super.connectedCallback()}disconnectedCallback(){super.disconnectedCallback()}_onScroll(){const{current:e}=this.state.data,t=document.body.scrollTop>100||document.documentElement.scrollTop>100;e.scrollTop!==t&&this.setData("current",{scrollTop:t})}setData(e,t,a){this.state.data={key:e,value:t,update:a}}inputBox({title:e,message:t,infoText:a,value:r,defaultOK:s,showValue:o,labelCancel:n,labelOK:c,onEvent:d}){return i`<modal-inputbox
      title="${l(e)}"
      message="${l(t)}"
      infoText="${l(a)}"
      value="${l(r)}"
      labelCancel="${n||this.app.msg("",{id:"msg_cancel"})}"
      labelOK="${c||this.app.msg("",{id:"msg_ok"})}"
      ?defaultOK="${s||!1}"
      ?showValue="${o||!1}"
      .onEvent=${d}
    ></modal-inputbox>`}_protector(){const{data:e}=this.state,{current:t,session:r}=this.state.data;return e[S.LOGIN].data?i`
      <div class="client-menubar">
        <client-menubar id="menuBar"
          side="${t.side}"
          module="${t.module}"
          ?scrollTop="${t.scrollTop}"
          .bookmark="${e[S.BOOKMARK]}"
          selectorPage=${this.app.getSetting("selectorPage")}
          .app="${this.app}" .msg="${this.app.msg}"
        ></client-menubar>
      </div>
      <div theme="${t.theme}" class="main">
        ${A(t.module===S.SEARCH?i`<client-search
          id="search" .data=${e[S.SEARCH]} side="${t.side}"
          .auditFilter="${e[S.LOGIN].data.audit_filter}"
          paginationPage=${this.app.getSetting("paginationPage")}
          .onEvent=${this.app.modules.search} 
          .msg="${this.app.msg}" 
          .quick="${{...this.app.modules.quick}}" .queries="${{...this.app.modules.queries}}"
        ></client-search>`:a)}
        ${A(t.module===S.EDIT?i`<client-edit
          id="edit" .data=${e[S.EDIT]} 
          side="${t.side}"
          .auditFilter="${e[S.LOGIN].data.audit_filter}"
          .newFilter="${e[S.LOGIN].data.edit_new}"
          paginationPage=${this.app.getSetting("paginationPage")}
          selectorPage=${this.app.getSetting("selectorPage")}
          .onEvent=${this.app.modules.edit}
          .msg="${this.app.msg}" .forms="${{...this.app.modules.forms}}"
        ></client-edit>`:a)}
        ${A(t.module===S.SETTING?i`<client-setting
          id="setting" .data=${e[S.SETTING]} 
          side="${t.side}"
          .auditFilter="${e[S.LOGIN].data.audit_filter}"
          username="${e[S.LOGIN].username}"
          paginationPage=${this.app.getSetting("paginationPage")}
          .onEvent=${this.app.modules.setting}
          .msg="${this.app.msg}" .forms="${{...this.app.modules.forms}}"
        ></client-setting>`:a)}
        ${A(t.module===S.TEMPLATE?i`<client-template
          id="template" .data=${e[S.TEMPLATE]} 
          side="${t.side}"
          paginationPage=${this.app.getSetting("paginationPage")}
          .onEvent=${this.app.modules.template}
          .msg="${this.app.msg}"
        ></client-template>`:a)}
        ${t.modalForm?t.modalForm:a}
      </div>`:i`
      <client-login id="Login"
        version="${r.version}"
        serverURL="${r.serverURL}"
        .locales="${{...r.locales}}"
        lang="${t.lang}"
        theme="${t.theme}"
        .data="${{...e[S.LOGIN]}}"
        .app="${this.app}" .msg="${this.app.msg}"
       >
      </client-login>
    `}render(){const{current:e}=this.state.data;return i`<style> :host { background-color: rgb(var(--${e.theme})); } </style>
      <form-toast id="appToast" 
        .setData="${this.setData}" ></form-toast>
      ${this._protector()}
      ${A(e.request?i`<form-spinner></form-spinner>`:a)}
    `}});export{v as A,y as B,w as D,U as E,D as I,x as L,N as M,F as P,M as S,O as T,G as a,S as b,L as c,k as d,I as e,P as f,$ as g,A as h,R as i,q as j,C as k,H as l,j as m,B as n,h as o};
