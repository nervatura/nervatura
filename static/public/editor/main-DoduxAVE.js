import{j as e,D as t,A as a,a as r,i as s,b as i}from"./module-1JqPnLJw.js";import{e as o,i as n,o as l,a as c}from"./module-BCutSLl1.js";import{l as d}from"./module-NpGsdRHB.js";const{I:p}=e,g=e=>e,h=(e,t)=>void 0!==e?._$litType$,u=e=>void 0===e.strings,m=()=>document.createComment(""),E=(e,t,a)=>{const r=e._$AA.parentNode,s=e._$AB;if(void 0===a){const t=r.insertBefore(m(),s),i=r.insertBefore(m(),s);a=new p(t,i,e,e.options)}else{const t=a._$AB.nextSibling,i=a._$AM,o=i!==e;if(o){let t;a._$AQ?.(e),a._$AM=e,void 0!==a._$AP&&(t=e._$AU)!==i._$AU&&a._$AP(t)}if(t!==s||o){let e=a._$AA;for(;e!==t;){const t=g(e).nextSibling;g(r).insertBefore(e,s),e=t}}}return a},T={},_=(e,t=T)=>e._$AH=t,A=e=>e._$AH,b=e=>(e=>null!=e?._$litType$?.h)(e)?e._$litType$.h:e.strings,f=o(class extends n{constructor(e){super(e),this.et=new WeakMap}render(e){return[e]}update(e,[r]){const s=h(this.it)?b(this.it):null,i=h(r)?b(r):null;if(null!==s&&(null===i||s!==i)){const r=A(e).pop();let i=this.et.get(s);if(void 0===i){const e=document.createDocumentFragment();i=t(a,e),i.setConnected(!1),this.et.set(s,i)}_(i,[r]),E(i,0,r)}if(null!==i){if(null===s||s!==i){const t=this.et.get(i);if(void 0!==t){const a=A(t).pop();(e=>{e._$AR()})(e),E(e,0,a),_(e,[a])}}this.it=r}else this.it=void 0;return this.render(r)}});var v=r`
html, :host {
  --font-family: "Noto Sans";
  --font-size: 14px;
  --menu-top-height: 0px;
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
`;const x={LIGHT:"light",DARK:"dark"},S={LOGIN:"login",TEMPLATE:"template"},C={PRIMARY:"primary",BORDER:"border"},M={LEFT:"left",CENTER:"center",RIGHT:"right"},y={TEXT:"text",COLOR:"color",FILE:"file",PASSWORD:"password"},D={DATE:"date",TIME:"time",DATETIME:"datetime-local"},I={INFO:"info",ERROR:"error"},w={AUTO:"auto",HIDE:"hide"},F={CHANGE:"change",CHECK:"check",SAVE:"save",DELETE:"delete",REPORT_SETTINGS:"report_settings",CREATE_REPORT:"create_report",HELP:"help",BLANK:"blank",SAMPLE:"sample",LOGOUT:"logout",THEME:"theme"},$={CHANGE:"change",LOGIN:"login",THEME:"theme",LANG:"lang"},O={CANCEL:"cancel",OK:"ok"},R={TOP:"top",BOTTOM:"bottom",ALL:"all",NONE:"none"},k={CHECK_EDITOR:"check_editor",CHECK_TRANSTYPE:"check_transtype"},N={LOAD_EDITOR:"load_editor"},L={TEXT:"string",LIST:"list",TABLE:"table"},P={ADD_ITEM:"add_item",CHANGE_TEMPLATE:"change_template",CHANGE_CURRENT:"change_current",GO_PREVIOUS:"go_previous",GO_NEXT:"go_next",CREATE_MAP:"create_map",SET_CURRENT:"set_current",MOVE_UP:"move_up",MOVE_DOWN:"move_down",DELETE_ITEM:"delete_item",EDIT_ITEM:"edit_item",EDIT_DATA_ITEM:"edit_data_item",SET_CURRENT_DATA:"set_current_data",SET_CURRENT_DATA_ITEM:"set_current_data_item",ADD_TEMPLATE_DATA:"add_template_data",DELETE_DATA:"delete_data",DELETE_DATA_ITEM:"delete_data_item"},U=r`
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
`;class q{constructor(e,t){this.host=e,this._data=t,e.addController(this)}get data(){return this._data}set data(e){const{key:t,value:a,update:r}=e;this._data[t]&&"object"==typeof a&&(this._data[t]={...this._data[t],...a},!1!==r&&this.host.requestUpdate())}}const G=(e,t)=>{let a=0;const r=t=>{let a=t;switch(e){case"sqlite":case"sqlite3":a=a.replace(/{CCS}/g,""),a=a.replace(/{SEP}/g,"||"),a=a.replace(/{CCE}/g,""),a=a.replace(/{CAS_TEXT}/g,"cast("),a=a.replace(/{CAE_TEXT}/g," as text)"),a=a.replace(/{CAS_INT}/g,"cast("),a=a.replace(/{CAE_INT}/g," as integer)"),a=a.replace(/{CAS_FLOAT}/g,"cast("),a=a.replace(/{CAE_FLOAT}/g," as double)"),a=a.replace(/{CAS_DATE}/g,""),a=a.replace(/{CASF_DATE}/g,""),a=a.replace(/{CAE_DATE}/g,""),a=a.replace(/{CAEF_DATE}/g,""),a=a.replace(/{FMSF_NUMBER}/g,""),a=a.replace(/{FMSF_DATE}/g,""),a=a.replace(/{FMEF_CONVERT}/g,""),a=a.replace(/{FMS_INT}/g,""),a=a.replace(/{FME_INT}/g,""),a=a.replace(/{FMS_DATE}/g,"substr("),a=a.replace(/{FME_DATE}/g,",1,10)"),a=a.replace(/{FMS_DATETIME}/g,"substr("),a=a.replace(/{FME_DATETIME}/g,",1,19)"),a=a.replace(/{FMS_TIME}/g,"substr(time("),a=a.replace(/{FME_TIME}/g,"),0,6)"),a=a.replace(/{JOKER}/g,"'%'"),a=a.replace(/{CUR_DATE}/g,"date('now')");break;case"mysql":a=a.replace(/{CCS}/g,"concat("),a=a.replace(/{SEP}/g,","),a=a.replace(/{CCE}/g,")"),a=a.replace(/{CAS_TEXT}/g,"cast("),a=a.replace(/{CAE_TEXT}/g," as char)"),a=a.replace(/{CAS_INT}/g,"cast("),a=a.replace(/{CAE_INT}/g," as signed)"),a=a.replace(/{CAS_FLOAT}/g,"cast("),a=a.replace(/{CAE_FLOAT}/g," as decimal)"),a=a.replace(/{CAS_DATE}/g,"cast("),a=a.replace(/{CASF_DATE}/g,"cast("),a=a.replace(/{CAE_DATE}/g," as date)"),a=a.replace(/{CAEF_DATE}/g," as date)"),a=a.replace(/{FMSF_NUMBER}/g,""),a=a.replace(/{FMSF_DATE}/g,""),a=a.replace(/{FMEF_CONVERT}/g,""),a=a.replace(/{FMS_INT}/g,"format(cast("),a=a.replace(/{FME_INT}/g," as signed), 0)"),a=a.replace(/{FMS_DATE}/g,"date_format("),a=a.replace(/{FME_DATE}/g,", '%Y-%m-%d')"),a=a.replace(/{FMS_DATETIME}/g,"date_format("),a=a.replace(/{FME_DATETIME}/g,", '%Y-%m-%dT%H:%i:%s')"),a=a.replace(/{FMS_TIME}/g,"cast(cast("),a=a.replace(/{FME_TIME}/g," as time) as char)"),a=a.replace(/{JOKER}/g,"'%'"),a=a.replace(/{CUR_DATE}/g,"current_date");break;case"postgres":a=a.replace(/{CCS}/g,""),a=a.replace(/{SEP}/g,"||"),a=a.replace(/{CCE}/g,""),a=a.replace(/{CAS_TEXT}/g,"cast("),a=a.replace(/{CAE_TEXT}/g," as text)"),a=a.replace(/{CAS_INT}/g,"cast("),a=a.replace(/{CAE_INT}/g," as integer)"),a=a.replace(/{CAS_FLOAT}/g,"cast("),a=a.replace(/{CAE_FLOAT}/g," as float)"),a=a.replace(/{CAS_DATE}/g,"cast("),a=a.replace(/{CASF_DATE}/g,"cast("),a=a.replace(/{CAE_DATE}/g," as date)"),a=a.replace(/{CAEF_DATE}/g," as date)"),a=a.replace(/{FMSF_NUMBER}/g,"case when rf_number.fieldname is null then 0 else "),a=a.replace(/{FMSF_DATE}/g,"case when rf_date.fieldname is null then current_date else "),a=a.replace(/{FMEF_CONVERT}/g," end "),a=a.replace(/{FMS_INT}/g,"to_char(cast("),a=a.replace(/{FME_INT}/g," as integer), '999,999,999')"),a=a.replace(/{FMS_DATE}/g,"to_char("),a=a.replace(/{FME_DATE}/g,", 'YYYY-MM-DD')"),a=a.replace(/{FMS_DATETIME}/g,"to_char("),a=a.replace(/{FME_DATETIME}/g,", 'YYYY-MM-DD\"T\"HH24:MI:SS')"),a=a.replace(/{FMS_TIME}/g,"substr(cast(cast("),a=a.replace(/{FME_TIME}/g," as time) as text), 0, 6)"),a=a.replace(/{JOKER}/g,"chr(37)"),a=a.replace(/{CUR_DATE}/g,"current_date");break;case"mssql":a=a.replace(/{CCS}/g,""),a=a.replace(/{SEP}/g,"+"),a=a.replace(/{CCE}/g,""),a=a.replace(/{CAS_TEXT}/g,"cast("),a=a.replace(/{CAE_TEXT}/g," as nvarchar)"),a=a.replace(/{CAS_INT}/g,"cast("),a=a.replace(/{CAE_INT}/g," as int)"),a=a.replace(/{CAS_FLOAT}/g,"cast("),a=a.replace(/{CAE_FLOAT}/g," as real)"),a=a.replace(/{CAS_DATE}/g,"cast("),a=a.replace(/{CASF_DATE}/g,""),a=a.replace(/{CAE_DATE}/g," as date)"),a=a.replace(/{CAEF_DATE}/g,""),a=a.replace(/{FMSF_NUMBER}/g,""),a=a.replace(/{FMSF_DATE}/g,""),a=a.replace(/{FMEF_CONVERT}/g,""),a=a.replace(/{FMS_INT}/g,"replace(convert(varchar,cast("),a=a.replace(/{FME_INT}/g," as money),1), '.00','')"),a=a.replace(/{FMS_DATE}/g,"convert(varchar(10),"),a=a.replace(/{FME_DATE}/g,", 120)"),a=a.replace(/{FMS_DATETIME}/g,"convert(varchar(19),"),a=a.replace(/{FME_DATETIME}/g,", 120)"),a=a.replace(/{FMS_TIME}/g,"SUBSTRING(cast(cast("),a=a.replace(/{FME_TIME}/g," as time) as nvarchar),0,6)"),a=a.replace(/{JOKER}/g,"'%'"),a=a.replace(/{CUR_DATE}/g,"cast(GETDATE() as DATE)")}return a},s=(t,r)=>{let i=t,o="";if(Array.isArray(i)){let e=" ",t="",a="";return i.length>0&&(["select","select_distinct","union_select","order_by","group_by"].includes(r)||0===i[0].length)&&(e=", "),i.forEach(n=>{let l=n;null==l&&(l="null"),0===l.length?"set"!==r&&(t="(",a=")"):2!==i.length||"and"!==l&&"or"!==l?r&&1===i.length&&"object"==typeof i[0]?o+=` (${s(l,r)})`:o+=e+s(l,r):(o+=`${l} (`,a=")")}),", "===e&&(o=o.substr(2)),r&&i.includes("on")&&(o=r.replace("_"," ")+o),t+o.toString().trim()+a}return"object"==typeof i?(Object.keys(i).forEach(e=>{o+="inner_join"===e||"left_join"===e?` ${s(i[e],e)}`:` ${e.replace("_"," ")} ${s(i[e],e)}`}),o):(i.includes("?")&&"select"!==r&&(a+=1,"postgres"===e&&(i=i.replace("?",`$${a}`))),i)};return"string"==typeof t?{sql:r(t),prmCount:a}:{sql:r(s(t)),prmCount:a}},H=async(e,t)=>(e=>{if(400===e.status||401===e.status)return e.json();if(204===e.status||205===e.status)return null;switch(e.headers.get("content-type").split(";")[0]){case"application/pdf":case"application/xml":case"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":return e.blob();case"application/json":return e.json();case"text/plain":case"text/csv":return e.text();default:return e}})((e=>{if(e.status>=200&&e.status<300||400===e.status||401===e.status)return e;const t=new Error(e.statusText);throw t.response=e,t})(await fetch(e,t))),j=(e,t)=>{const a=document.createElement("a");a.href=e,a.download=t||e,document.body.appendChild(a),a.click()};class B{constructor(e){this.host=e,this.modules={},this.getSql=G,this.request=H,this.saveToDisk=j,this.currentModule=this.currentModule.bind(this),this.getSetting=this.getSetting.bind(this),this.msg=this.msg.bind(this),this.requestData=this.requestData.bind(this),this.resultError=this.resultError.bind(this),this.showHelp=this.showHelp.bind(this),this.showToast=this.showToast.bind(this),this.signOut=this.signOut.bind(this),e.addController(this)}hostConnected(){const{state:e,setData:t}=this.host;this.store={data:e.data,setData:t},this._loadConfig(window.location)}async _loadConfig(e){const{data:t,setData:a}=this.store,[r,s]=(()=>{const t=e=>{const t={};return e.split("&").forEach(e=>{const a=String(e).indexOf("="),r=String(e).substring(0,a>0?a:String(e).length),s=a>-1&&a<String(e).length?String(e).substring(a+1):"";t[r]=s}),t};if(e.hash)return["hash",t(e.hash.substring(1))];if(e.search)return["search",t(e.search.substring(1))];const a=e.pathname.substring(1).split("/");return[a[0],a.slice(1)]})();s.session&&this.request(`/client/api/template/${s.session}/${s.code}`,{method:"POST",headers:{"Content-Type":"application/x-www-form-urlencoded"}}).then(e=>(a(S.LOGIN,{data:{token:e.token}}),window.history.replaceState(null,null,window.location.pathname),this.currentModule({data:{module:S.TEMPLATE},content:{fkey:"setTemplate",args:[{type:"template",report:e.report}]}}))).catch(e=>{a("error",e),this.showToast(I.ERROR,e.message)}),import("./module-CMtt1ZNs.js")}async currentModule({data:e,content:t}){const{setData:a}=this.store,r={template:async()=>{const{TemplateController:e}=await import("./module-C6KOjx36.js");await import("./module-CaDT3Dha.js"),this.modules.template=new e(this.host)}};this.modules[e.module]||await r[e.module](),a("current",{...e}),t&&this.modules[e.module]&&this.modules[e.module][t.fkey](...t.args)}getSetting(e){const{ui:t}=this.store.data;if("ui"===e){const e={...t};return Object.keys(e).forEach(t=>{localStorage.getItem(t)&&(e[t]=localStorage.getItem(t))}),e}return localStorage.getItem(e)||t[e]||""}msg(e,t){let a=e;const{locales:r}=this.store.data.session,{lang:s}=this.store.data.current;return r[s]&&r[s][t.id]?a=r[s][t.id]:"en"!==s&&r.en[t.id]&&(a=r.en[t.id]),a}async requestData(e,t,a){const{data:r,setData:s}=this.store;let i=t;try{a||s("current",{request:!0});let t="SERVER"===r.session.serverURL?r.session.apiPath+e:r.login.server+e;const o=r.login.data?r.login.data.token:i.token||"";if(i.headers||(i={...i,headers:{}}),i={...i,headers:{...i.headers,"Content-Type":"application/json"}},""!==o&&(i={...i,headers:{...i.headers,Authorization:`Bearer ${o}`}}),i.data&&(i={...i,body:JSON.stringify(i.data)}),i.query){const e=new URLSearchParams;Object.keys(i.query).forEach(t=>{e.append(t,i.query[t])}),t+=`?${e.toString()}`}const n=await this.request(t,i);return a||s("current",{request:!1}),n&&n.code&&n.message?(401===n.code&&this.signOut(),{error:{message:n.message},data:null}):n}catch(e){return a||s("current",{request:!1}),{error:{message:e.message},data:null}}}resultError(e){const{setData:t}=this.store;return e.error&&t("error",e.error),e.error&&e.error.message?this.showToast(I.ERROR,e.error.message):this.showToast(I.ERROR,this.msg("Internal Server Error",{id:"error_internal"})),!1}showHelp(e){const{data:t}=this.store,a=document.createElement("a");a.setAttribute("href",t.session.helpPage+e),a.setAttribute("target","_blank"),document.body.appendChild(a),a.click()}showToast(e,t,a){const{current:r}=this.store.data,s=void 0!==a?a:this.getSetting("toastTimeout");r.toast&&r.toast.show({type:e,value:t,timeout:s})}signOut(){const{data:e,setData:t}=this.store,a=e[S.LOGIN]?.callback;a&&this.isSafeRedirectUrl(a)?window.location.replace(a):t(S.LOGIN,{data:null,token:null})}isSafeRedirectUrl(e){if(!e||"string"!=typeof e)return!1;const t=e.trim();if(""===t||t.startsWith("javascript:")||t.startsWith("data:"))return!1;try{return new URL(t,window.location.origin).origin===window.location.origin}catch{return t.startsWith("/")&&!t.startsWith("//")}}}const z={session:{version:"1.0.0",locales:d,serverURL:"SERVER",apiPath:"/api/v6",helpPage:"/docs/"},ui:{toastTimeout:4,paginationPage:10,timeIntervals:15,export_sep:";",page_size:"a4",page_orient:"portrait",report_orientation:[["portrait","report_portrait"],["landscape","report_landscape"]],report_size:[["a3","A3"],["a4","A4"],["a5","A5"],["letter","Letter"],["legal","Legal"]]},current:{home:S.TEMPLATE,module:S.LOGIN,side:w.AUTO,lang:localStorage.getItem("lang")&&d[localStorage.getItem("lang")]?localStorage.getItem("lang"):"en",theme:localStorage.getItem("theme")||x.LIGHT},login:{username:localStorage.getItem("username")||"",database:localStorage.getItem("database")||"",code:localStorage.getItem("code")||"",server:localStorage.getItem("server")&&""!==localStorage.getItem("server")?localStorage.getItem("server"):"nervatura.github.io"!==window.location.hostname?window.location.origin+"/api/v6":""},template:{dirty:!1}},K=r`
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
    </div>`}static get styles(){return[K]}});const V=r`
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
`;customElements.define("form-toast",class extends s{constructor(){super(),this.id=Math.random().toString(36).slice(2),this.name=void 0,this.type=I.INFO,this.hidden=!0,this.timeout=4,this.style={},this._iconMap={info:"InfoCircle",error:"ExclamationTriangle",success:"CheckSquare"}}static get properties(){return{id:{type:String},name:{type:String,reflect:!0},type:{type:String},hidden:{type:Boolean},timeout:{type:Number},style:{type:Object}}}connectedCallback(){super.connectedCallback(),this.setData&&this.setData("current",{toast:this},!1)}disconnectedCallback(){this.setData&&this.setData("current",{toast:null},!1),super.disconnectedCallback()}show({type:e,value:t,timeout:a}){this.hidden&&(this.value=t||this.value,this.type=e||this.type,this.timeout=void 0!==a?a:this.timeout,this.hidden=!1,this.timeout>0&&(this.timeoutVar=setTimeout(()=>{this&&!this.hidden&&(this.hidden=!0)},1e3*this.timeout)))}close(){this.hidden||(clearTimeout(this.timeoutVar),this.hidden=!0)}render(){return this.hidden?a:i`<div 
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
      </div>`}static get styles(){return[V]}});customElements.define("nervatura-editor",class extends s{static get properties(){return{}}static get styles(){return[v,U]}constructor(){super(),this.state=new q(this,z),this.app=new B(this),this.inputBox=this.inputBox.bind(this),this.setData=this.setData.bind(this)}connectedCallback(){super.connectedCallback()}disconnectedCallback(){super.disconnectedCallback()}_onScroll(){const{current:e}=this.state.data,t=document.body.scrollTop>100||document.documentElement.scrollTop>100;e.scrollTop!==t&&this.setData("current",{scrollTop:t})}setData(e,t,a){this.state.data={key:e,value:t,update:a}}inputBox({title:e,message:t,infoText:a,value:r,defaultOK:s,showValue:o,labelCancel:n,labelOK:c,onEvent:d}){return i`<modal-inputbox
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
      <div theme="${t.theme}" class="main">
        ${f(t.module===S.TEMPLATE?i`<client-template
          id="template" .data=${e[S.TEMPLATE]} 
          side="${t.side}"
          paginationPage=${this.app.getSetting("paginationPage")}
          theme="${t.theme}"
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
      ${f(e.request?i`<form-spinner></form-spinner>`:a)}
    `}});export{S as A,C as B,D,k as E,y as I,$ as L,O as M,R as P,w as S,I as T,x as a,F as b,P as c,M as d,N as e,L as f,u as r};
