import{b as e,E as t,c as s,M as r,S as i,d as a,T as l}from"./main-nlu71Zxt.js";import"./module-ORVyXmTn.js";import"./module-bL_R6UIF.js";import"./module-r-Be6lM0.js";const o=e=>{switch(e.filtertype){case"==N":return"string"===e.fieldtype?["and",[[e.sqlstr,"like","''"],["or",e.sqlstr,"is null"]]]:["and",e.sqlstr,"is null"];case"!==":return"string"===e.fieldtype?["and",[`lower(${e.sqlstr})`,"not like","{CCS}{JOKER}{SEP}lower(?){SEP}{JOKER}{CCE}"]]:["and",[e.sqlstr,"<>","?"]];case">==":return["and",[e.sqlstr,">=","?"]];case"<==":return["and",[e.sqlstr,"<=","?"]];default:return"string"===e.fieldtype?["and",[`lower(${e.sqlstr})`,"like","{CCS}{JOKER}{SEP}lower(?){SEP}{JOKER}{CCE}"]]:["and",[e.sqlstr,"=","?"]]}},d=e=>"date"===e?(new Date).toISOString().split("T")[0]:["bool","integer","float"].includes(e)?0:"";class n{constructor(e){this.host=e,this.app=e.app,this.store=e.app.store,this.module={},this.addFilter=this.addFilter.bind(this),this.browserView=this.browserView.bind(this),this.deleteFilter=this.deleteFilter.bind(this),this.editCell=this.editCell.bind(this),this.editFilter=this.editFilter.bind(this),this.editRow=this.editRow.bind(this),this.exportResult=this.exportResult.bind(this),this.onBrowserEvent=this.onBrowserEvent.bind(this),this.onModalEvent=this.onModalEvent.bind(this),this.onSideEvent=this.onSideEvent.bind(this),this.quickSearch=this.quickSearch.bind(this),this.saveBookmark=this.saveBookmark.bind(this),this.setColumns=this.setColumns.bind(this),this.setModule=this.setModule.bind(this),this.showBrowser=this.showBrowser.bind(this),this.showServerCmd=this.showServerCmd.bind(this),this.showTotal=this.showTotal.bind(this),e.addController(this)}setModule(e){this.module=e}addFilter(){const{queries:t}=this.app.modules,{setData:s}=this.store,{vkey:r,view:i,filters:a}=this.store.data[e.SEARCH],l=t[r]()[i],o=l.fields[Object.keys(l.fields)[0]],n={...a,[i]:[...a[i],{id:(new Date).getTime().toString(),fieldtype:o.fieldtype,fieldname:Object.keys(l.fields)[0],sqlstr:o.sqlstr,wheretype:o.wheretype,filtertype:"===",value:d(o.fieldtype)}]};s(e.SEARCH,{filters:n})}async browserView(){const{queries:t}=this.app.modules,{getDataFilter:s,getUserFilter:r,getSql:i,requestData:a,resultError:l}=this.app,{setData:d,data:n}=this.store,{vkey:h,view:c,filters:u}=this.store.data[e.SEARCH];let f={...t[h]()[c].sql},p=[],m=[];u[c].filter((e=>"where"===e.wheretype)).forEach((e=>{m.push(o(e)),"==N"!==e.filtertype&&p.push(e.value)})),m.length>0&&(f={...f,where:[...f.where,...m]}),m=[],u[c].filter((e=>"having"===e.wheretype)).forEach((e=>{m.push(o(e)),"==N"!==e.filtertype&&p.push(e.value)})),m.length>0&&(f={...f,having:[...f.having,...m]}),m=s(h,[],c),m.length>0&&(f={...f,where:[...f.where,m]});const E=r(h);E.where.length>0&&(f={...f,where:[...f.where,E.where]},p=p.concat(E.params));const v={method:"POST",data:[{key:"result",text:i(n[e.LOGIN].data.engine,f).sql,values:p}]},w=await a("/view",v);return w.error?l(w):(d(e.SEARCH,{result:w.result,dropdown:"",page:1}),!0)}deleteFilter(t){const{setData:s}=this.store,{view:r,filters:i}=this.store.data[e.SEARCH],a={...i,[r]:[...i[r].slice(0,t),...i[r].slice(t+1)]};s(e.SEARCH,{filters:a})}editCell({fieldname:s,value:r,row:i}){const{currentModule:a}=this.app,l=r.split("/");let o={ntype:l[0],ttype:l[1],id:l[2]};"id"===s&&(o={...o,form:i.form,form_id:i.form_id}),a({data:{module:e.EDIT},content:{fkey:"checkEditor",args:[o,t.LOAD_EDITOR]}})}editFilter({index:t,fieldname:s,value:r}){const{queries:i}=this.app.modules,{setData:a,data:l}=this.store,{vkey:o,view:n,filters:h}=this.store.data[e.SEARCH],c=i[o]()[n],u={...h};if("filtertype"!==s&&"value"!==s||(u[n][t]={...u[n][t],[s]:r}),"fieldname"===s)if(Object.keys(c.fields).includes(r)){const e=c.fields[r];u[n][t]={...u[n][t],fieldname:r,fieldtype:e.fieldtype,sqlstr:e.sqlstr,wheretype:e.wheretype,filtertype:"===",value:d(e.fieldtype)}}else{const e=l.deffield.filter((e=>e.fieldname===r))[0],s={bool:{fieldtype:"bool",sqlstr:"fg.groupvalue='bool' and case when fv.value='true' then 1 else 0 end "},integer:{fieldtype:"integer",sqlstr:"fg.groupvalue='integer' and {FMSF_NUMBER} {CAS_INT}fv.value {CAE_INT} {FMEF_CONVERT} "},float:{fieldtype:"float",sqlstr:"fg.groupvalue='float' and {FMSF_NUMBER} {CAS_FLOAT}fv.value {CAE_FLOAT} {FMEF_CONVERT} "},date:{fieldtype:"date",sqlstr:"fg.groupvalue='date' and {FMSF_DATE} {CASF_DATE}fv.value{CAEF_DATE} {FMEF_CONVERT} "},customer:{fieldtype:"string",sqlstr:"rf_customer.custname "},tool:{fieldtype:"string",sqlstr:"rf_tool.serial "},product:{fieldtype:"string",sqlstr:"rf_product.partnumber "},trans:{fieldtype:"string",sqlstr:"rf_trans.transnumber "},transitem:{fieldtype:"string",sqlstr:"rf_trans.transnumber "},transmovement:{fieldtype:"string",sqlstr:"rf_trans.transnumber "},transpayment:{fieldtype:"string",sqlstr:"rf_trans.transnumber "},project:{fieldtype:"string",sqlstr:"rf_project.pronumber "},employee:{fieldtype:"string",sqlstr:"rf_employee.empnumber "},place:{fieldtype:"string",sqlstr:"rf_place.planumber "}};let i="string",a="fv.value ";s[e.fieldtype]&&(i=s[e.fieldtype].fieldtype,a=s[e.fieldtype].sqlstr),u[n][t]={...u[n][t],fieldname:e.fieldname,fieldlimit:["and","fv.fieldname","=",`'${e.fieldname}'`],fieldtype:i,sqlstr:a,wheretype:"where",filtertype:"===",value:""}}a(e.SEARCH,{filters:u})}editRow(s){const{currentModule:r}=this.app,i=s.id.split("/"),a={ntype:i[0],ttype:i[1],id:parseInt(i[2],10),item:s};"servercmd"===a.ntype?this.showServerCmd(a.id):r({data:{module:e.EDIT},content:{fkey:"checkEditor",args:[a,t.LOAD_EDITOR]}})}exportResult(t){const{getSetting:s,saveToDisk:r}=this.app,{result:i,view:a}=this.store.data[e.SEARCH],l=`${a}.csv`;let o="";const d=Object.keys(t).map((e=>t[e].label));o+=`${d.join(s("export_sep"))}\n`,i.forEach((e=>{const r=Object.keys(t).map((t=>void 0!==e[`export_${t}`]?e[`export_${t}`]:e[t]));o+=`${r.join(s("export_sep"))}\n`}));r(URL.createObjectURL(new Blob([o],{type:"text/csv;charset=utf-8;"})),l)}onBrowserEvent({key:r,data:i}){const{showHelp:a,currentModule:l}=this.app,{setData:o}=this.store;switch(r){case s.CHANGE:o(e.SEARCH,{[i.fieldname]:i.value});break;case s.ADD_FILTER:this.addFilter();break;case s.BOOKMARK_SAVE:this.saveBookmark();break;case s.BROWSER_VIEW:this.browserView();break;case s.CURRENT_PAGE:o(e.SEARCH,{page:i.value});break;case s.DELETE_FILTER:this.deleteFilter(i.value);break;case s.EDIT_FILTER:this.editFilter(i);break;case s.EXPORT_RESULT:this.exportResult(i.value);break;case s.EDIT_CELL:this.editCell(i);break;case s.SET_COLUMNS:this.setColumns(i.fieldname,i.value);break;case s.SET_FORM_ACTION:l({data:{module:e.EDIT},content:{fkey:"checkEditor",args:[i,t.FORM_ACTION]}});break;case s.SHOW_BROWSER:this.showBrowser(i.value,i.vname);break;case s.SHOW_HELP:a(i.value);break;case s.SHOW_TOTAL:this.showTotal(i)}}onModalEvent({key:t,data:s}){const{setData:i}=this.store;switch(t){case r.CANCEL:i("current",{modalForm:null});break;case r.SEARCH:this.quickSearch(s.value);break;case r.SELECTED:this.editRow(s.value);break;case r.CURRENT_PAGE:i(e.SEARCH,{page:s.value})}}onSideEvent({key:s,data:r}){const{currentModule:l}=this.app,{setData:o}=this.store;switch(s){case i.CHANGE:o(e.SEARCH,{[r.fieldname]:r.value});break;case i.BROWSER:this.showBrowser(r.value);break;case i.QUICK:o(e.SEARCH,{seltype:"selector",result:[],qview:r.value,qfilter:"",page:1}),o("current",{side:a.HIDE});break;case i.CHECK:l({data:{module:e.EDIT},content:{fkey:"checkEditor",args:[r,t.LOAD_EDITOR]}})}}async quickSearch(t){const{setData:s}=this.store,{quickSearch:r,resultError:i}=this.app,{qview:a}=this.store.data[e.SEARCH],l=await r(a,t);return l.error?i(l):(s(e.SEARCH,{result:l.result,qfilter:t,page:1}),!0)}saveBookmark(){const{queries:t}=this.app.modules,{saveBookmark:s}=this.app,{vkey:r,view:i}=this.store.data[e.SEARCH];s(["browser",t[r]()[i].label])}setColumns(t,s){const{setData:r}=this.store,{view:i,columns:a}=this.store.data[e.SEARCH];if(s)r(e.SEARCH,{columns:{...a,[i]:{...a[i],[t]:!0}}});else{const{[t]:s,...l}=a[i];r(e.SEARCH,{columns:{...a,[i]:l}})}}async showBrowser(t,s,r){const{queries:i}=this.app.modules,{getSql:l,requestData:o,resultError:d}=this.app,{setData:n,data:h}=this.store,c=r||h[e.SEARCH];n("current",{side:a.HIDE});let u={...c,seltype:"browser",vkey:t,view:void 0===s?Object.keys(i[t]())[1]:s,result:[],deffield:[],show_dropdown:!1,show_header:void 0===c.show_header||c.show_header,show_columns:void 0!==c.show_columns&&c.show_columns,page:c.page||1};const f={method:"POST",data:[{key:"deffield",text:l(h[e.LOGIN].data.engine,i[t]().options.deffield_sql).sql,values:[]}]},p=await o("/view",f);if(p.error)return d(p);u={...u,deffield:p.deffield},u.filters[u.view]||(u={...u,filters:{...u.filters,[u.view]:[]}});const m=i[t]()[u.view];if(void 0===u.columns[u.view]){u={...u,columns:{...u.columns,[u.view]:{}}};for(let e=0;e<Object.keys(m.columns).length;e+=1){const t=Object.keys(m.columns)[e];u={...u,columns:{...u.columns,[u.view]:{...u.columns[u.view],[t]:m.columns[t]}}}}}if(0===Object.keys(u.columns[u.view]).length)for(let e=0;e<3;e+=1){const t=Object.keys(m.fields)[e];u={...u,columns:{...u.columns,[u.view]:{...u.columns[u.view],[t]:!0}}}}return n(e.SEARCH,u),!0}showServerCmd(t){const{modalServer:s}=this.module,{request:i,resultError:a,requestData:o,showToast:d}=this.app,{setData:n}=this.store,h=this.store.data[e.LOGIN],{session:c}=this.store.data,u=h.data.menuCmds.filter((e=>e.id===parseInt(t,10)))[0];if(u){const e=h.data.menuFields.filter((e=>e.menu_id===parseInt(t,10))),f={};e.forEach((e=>{switch(e.fieldtypeName){case"bool":f[e.fieldname]=!1;break;case"float":case"integer":f[e.fieldname]=0;break;default:f[e.fieldname]=""}}));const p=s({cmd:{...u},fields:[...e],values:f,onEvent:{onModalEvent:async e=>{if(n("current",{modalForm:null}),e.key===r.OK){const t=e.data,s=new URLSearchParams,r={...t.values};if(t.fields.forEach((e=>{"bool"===e.fieldtypeName?(s.append(e.fieldname,t.values[e.fieldname]?1:0),r[e.fieldname]=t.values[e.fieldname]?1:0):s.append(e.fieldname,t.values[e.fieldname])})),"get"===t.cmd.methodName){let e=t.cmd.address||"";""===e&&t.cmd.funcname&&""!==t.cmd.funcname&&(e="SERVER"===c.serverURL?`${c.apiPath}/${t.cmd.funcname}`:`${h.server}/${t.cmd.funcname}`),""!==e&&window.open(`${e}?${s.toString()}`,"_system")}else{const e={key:t.cmd.funcname||t.cmd.menukey,values:r};let s;if(t.cmd.address&&""!==t.cmd.address)try{s=await i(t.cmd.address,{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify(e)})}catch(e){return a(e),null}else s=await o("/function",{method:"POST",data:e});if(s.error)return a(s),null;let n=s;"object"==typeof s&&(n=JSON.stringify(s,null,"  ")),d(l.SUCCESS,n,0)}}return!0}}});n("current",{modalForm:p})}}showTotal({fields:t,totalFields:s}){const{modalTotal:r}=this.module,{setData:i}=this.store,{result:a}=this.store.data[e.SEARCH],l=e=>Number.isNaN(parseFloat(e))?0:parseFloat(e);let o=s;const d=Object.keys(t).includes("deffield_value");a.forEach((e=>{d?void 0!==o.totalFields[e.fieldname]&&(o={...o,totalFields:{...o.totalFields,[e.fieldname]:o.totalFields[e.fieldname]+l(e.export_deffield_value)}}):Object.keys(o.totalFields).forEach((t=>{o=void 0!==e[`export_${t}`]?{...o,totalFields:{...o.totalFields,[t]:o.totalFields[t]+l(e[`export_${t}`])}}:{...o,totalFields:{...o.totalFields,[t]:o.totalFields[t]+l(e[t])}}}))})),i("current",{modalForm:r(o)})}}export{n as SearchController,d as defaultFilterValue,o as getFilterWhere};
