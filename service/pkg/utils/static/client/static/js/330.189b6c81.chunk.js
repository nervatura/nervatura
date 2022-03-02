"use strict";(self.webpackChunknervatura_client=self.webpackChunknervatura_client||[]).push([[330],{3331:function(e,r,a){a.d(r,{Z:function(){return d}});var t=a(1413),n=a(4925),s="Select_selectStyle__kV5uh",l="Select_optionStyle__hc7Dx",o="Select_optionPlaceholder__1tEB7",i=a(6417),u=["options","placeholder","className","onChange"],c=function(e){var r=e.options,a=e.placeholder,c=e.className,d=e.onChange,m=(0,n.Z)(e,u),g=r.map((function(e,r){return(0,i.jsx)("option",{className:"".concat(l),value:e.value,children:e.text},r)}));return"undefined"!==typeof a&&g.unshift((0,i.jsx)("option",{className:" ".concat(l," ").concat(o),value:"",children:a},"placeholder")),(0,i.jsx)("select",(0,t.Z)((0,t.Z)({},m),{},{className:"".concat(c," ").concat(s),onChange:function(e){return d(e.target.value)},children:g}))};c.defaultProps={options:[],placeholder:void 0,className:"",onChange:void 0};var d=c},7718:function(e,r,a){a.r(r),a.d(r,{default:function(){return O}});var t=a(4942),n=a(5861),s=a(1413),l=a(885),o=a(7757),i=a.n(o),u=a(7313),c=a(4754),d=a.n(c),m=a(7559),g=a(9155),p=a(4925),f="Login_inputMargin__Iu818",v="Login_borderButton__BmzVE",h="Login_title__vcg8P",x="Login_version__pQpMb",_="Login_modal__nr6qg",y="Login_middle__6d6jf",b="Login_dialog__bOFEw",j=a(8907),N=a(6360),k=a(3331),w=a(4721),Z=a(5750),T=a(6417),S=["theme","version","locales","configServer","lang","username","password","database","server","onLogin","setTheme","setLocale","getText","changeData"],L=function(e){var r=e.theme,a=e.version,t=e.locales,n=e.configServer,l=e.lang,o=e.username,i=e.password,u=e.database,c=e.server,d=e.onLogin,m=e.setTheme,g=e.setLocale,L=e.getText,P=e.changeData,D=(0,p.Z)(e,S);return(0,T.jsx)("div",{className:_,children:(0,T.jsx)("div",{className:y,children:(0,T.jsxs)("div",(0,s.Z)((0,s.Z)({},D),{},{className:"".concat(b," ").concat(r),children:[(0,T.jsxs)("div",{className:"row primary".concat(" ",h),children:[(0,T.jsx)("div",{className:"cell",children:(0,T.jsx)(j.Z,{value:L("title_login")})}),(0,T.jsx)("div",{className:"cell".concat(" ",x),children:(0,T.jsx)(j.Z,{value:"v"+a})})]}),(0,T.jsxs)("div",{className:"row full section-small",children:[(0,T.jsxs)("div",{className:"row full section-small",children:[(0,T.jsxs)("div",{className:"row full",children:[(0,T.jsx)("div",{className:"padding-normal s12 m4 l4",children:(0,T.jsx)(j.Z,{value:L("login_username"),className:"bold"})}),(0,T.jsx)("div",{className:"container s12 m8 l8",children:(0,T.jsx)(N.Z,{id:"username",type:"text",className:"full".concat(" ",f),value:o,onChange:function(e){return P("username",e)}})})]}),(0,T.jsxs)("div",{className:"row full",children:[(0,T.jsx)("div",{className:"padding-normal s12 m4 l4",children:(0,T.jsx)(j.Z,{value:L("login_password"),className:"bold"})}),(0,T.jsx)("div",{className:"container s12 m8 l8",children:(0,T.jsx)(N.Z,{id:"password",type:"password",className:"full".concat(" ",f),value:i,onChange:function(e){return P("password",e)},onEnter:d})})]})]}),(0,T.jsxs)("div",{className:"row full section-small",children:[(0,T.jsxs)("div",{className:"row full",children:[(0,T.jsx)("div",{className:"padding-normal s12 m4 l4",children:(0,T.jsx)(j.Z,{value:L("login_database"),className:"bold"})}),(0,T.jsx)("div",{className:"container s12 m8 l8",children:(0,T.jsx)(N.Z,{id:"database",type:"text",className:"full".concat(" ",f),value:u,onChange:function(e){return P("database",e)},onEnter:d})})]}),n?null:(0,T.jsxs)("div",{className:"row full",children:[(0,T.jsx)("div",{className:"padding-normal full",children:(0,T.jsx)(j.Z,{value:L("login_server"),className:"bold"})}),(0,T.jsx)("div",{className:"container full",children:(0,T.jsx)(N.Z,{id:"server",type:"text",className:"full".concat(" ",f),value:c,onChange:function(e){return P("server",e)}})})]})]})]}),(0,T.jsxs)("div",{className:"row full section-small secondary-title",children:[(0,T.jsxs)("div",{className:"container section-small s6 m6 l6",children:[(0,T.jsx)(w.Z,{id:"theme",className:"border-button".concat(" ",v),value:"dark"===r?(0,T.jsx)(Z.Z,{iconKey:"Sun",width:18,height:18}):(0,T.jsx)(Z.Z,{iconKey:"Moon",width:18,height:18}),onClick:m}),(0,T.jsx)(k.Z,{id:"lang",value:l,options:Object.keys(t).map((function(e){return{value:e,text:t[e][e]||e}})),onChange:function(e){return g(e)}})]}),(0,T.jsx)("div",{className:"container section-small s6 m6 l6",children:(0,T.jsx)(w.Z,{id:"login",autoFocus:!0,disabled:o&&0!==String(o).length&&u&&0!==String(u).length?"":"disabled",className:"primary full",label:L("login_login"),onClick:d})})]})]}))})})};L.defaultProps={theme:"light",version:"",locales:{},lang:"en",configServer:!1,username:"",password:"",database:"",server:"",changeData:void 0,getText:void 0,onLogin:void 0,setTheme:void 0,setLocale:void 0};var P=L,D=function(e){var r=e.session,a=e.data,t=e.current,n=e.locales,l=e.changeData,o=e.getText,i=e.onLogin,u=e.setTheme,c=e.setLocale;return(0,T.jsx)(P,(0,s.Z)((0,s.Z)({},a),{},{theme:t.theme,lang:t.lang,version:r.version,locales:n,configServer:r.configServer,changeData:l,getText:o,onLogin:i,setTheme:u,setLocale:c}))};D.defaultProps={data:{username:P.defaultProps.username,password:P.defaultProps.password,database:P.defaultProps.database,server:P.defaultProps.server},session:{},locales:{},current:{},changeData:void 0,getText:void 0,onLogin:void 0,setTheme:void 0,setLocale:void 0};var C=(0,u.memo)(D,(function(e,r){return e.data===r.data&&e.current.theme===r.current.theme&&e.current.lang===r.current.lang&&e.locales===r.locales})),$=function(e){var r=(0,u.useContext)(m.ZP),a=r.data,o=r.setData,c=(0,g.xZ)(a,o),p=(0,u.useState)(d()(e,{$merge:{session:a.session}})),f=(0,l.Z)(p,1)[0];return f.data=d()(f.data,{$merge:(0,s.Z)({},a[f.key])}),f.current=d()(f.current,{$merge:(0,s.Z)({},a.current)}),f.locales=d()(f.locales,{$merge:(0,s.Z)({},a.session.locales)}),f.getText=function(e,r){return c.getText(e,r)},f.userLog=function(){var e=(0,n.Z)(i().mark((function e(r){var a;return i().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return a={method:"POST",token:r.token,data:[{keys:{employee_id:r.employee.empnumber,logstate:"login"}}]},e.next=3,c.requestData("/log",a);case 3:return e.abrupt("return",e.sent);case 4:case"end":return e.stop()}}),e)})));return function(r){return e.apply(this,arguments)}}(),f.loginData=function(){var e=(0,n.Z)(i().mark((function e(r){var a,n,s,l,o;return i().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return a=d()(f.data,{$set:{token:r.token,engine:r.engine}}),n=[{key:"employee",text:(0,g.aT)(r.engine,{select:["e.*","ug.groupvalue as usergroupName","dp.groupvalue as departmentName"],from:"employee e",inner_join:["groups ug","on",["e.usergroup","=","ug.id"]],left_join:["groups dp","on",["e.department","=","dp.id"]],where:["username","=","?"]}).sql,values:[f.data.username]},{key:"menuCmds",text:(0,g.aT)(r.engine,{select:["m.*","st.groupvalue as methodName"],from:"ui_menu m",inner_join:["groups st","on",["m.method","=","st.id"]]}).sql,values:[]},{key:"menuFields",text:(0,g.aT)(r.engine,{select:["mf.*","ft.groupvalue as fieldtypeName"],from:"ui_menufields mf",inner_join:["groups ft","on",["mf.fieldtype","=","ft.id"]],order_by:["menu_id","orderby"]}).sql,values:[]},{key:"userlogin",text:(0,g.aT)(r.engine,{select:["value"],from:"fieldvalue",where:[["ref_id","is","null"],["and","fieldname","=","'log_login'"]]}).sql,values:[]},{key:"groups",text:(0,g.aT)(r.engine,{select:["*"],from:"groups",where:["groupname","in",[[],"'usergroup'","'nervatype'","'transtype'","'inputfilter'","'transfilter'","'department'","'logstate'","'fieldtype'","'service'"]]}).sql,values:[]}],s={method:"POST",token:r.token,data:n},e.next=5,c.requestData("/view",s);case 5:if(!(l=e.sent).error){e.next=8;break}return e.abrupt("return",l);case 8:return a=d()(a,{$merge:l}),a=d()(a,{$merge:{employee:a.employee[0],userlogin:a.userlogin.length>0?a.userlogin[0].value:"false"}}),n=[{key:"audit",text:(0,g.aT)(r.engine,{select:["au.nervatype","nt.groupvalue as nervatypeName","au.subtype","case when nt.groupvalue = 'trans' then st.groupvalue else m.menukey end as subtypeName","au.inputfilter","ip.groupvalue as inputfilterName","au.supervisor"],from:"ui_audit au",inner_join:[["groups nt","on",["au.nervatype","=","nt.id"]],["groups ip","on",["au.inputfilter","=","ip.id"]]],left_join:[["groups st","on",["au.subtype","=","st.id"]],["ui_menu m","on",["au.subtype","=","m.id"]]],where:["au.usergroup","=","?"]}).sql,values:[a.employee.usergroup]},{key:"transfilter",text:(0,g.aT)(r.engine,{select:["ref_id_2 as transfilter","g.groupvalue as transfilterName"],from:"link",inner_join:["groups g","on",["link.ref_id_2","=","g.id"]],where:[["ref_id_1","=","?"],["and","link.deleted","=","0"],["and","nervatype_1","in",[{select:["id"],from:"groups",where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'groups'"]]}]],["and","nervatype_2","in",[{select:["id"],from:"groups",where:[["groupname","=","'nervatype'"],["and","groupvalue","=","'groups'"]]}]]]}).sql,values:[a.employee.usergroup]}],s={method:"POST",token:r.token,data:n},e.next=14,c.requestData("/view",s);case 14:if(!(l=e.sent).error){e.next=17;break}return e.abrupt("return",l);case 17:return null===(a=d()(a,{$merge:{transfilter:l.transfilter.length>0?l.transfilter[0].transfilter:null,audit:l.audit}})).transfilter?(o=a.groups.filter((function(e){return"transfilter"===e.groupname&&"all"===e.groupvalue}))[0],a=d()(a,{$merge:{transfilter:o.id,transfilterName:"all"}})):a=d()(a,{$merge:{transfilterName:l.transfilter[0].transfilterName}}),a=d()(a,{$merge:{audit_filter:{trans:{},menu:{},report:{}},edit_new:[[],[],[],[]]}}),[["offer",0],["order",0],["worksheet",0],["rent",0],["invoice",0],["receipt",0],["bank",1],["cash",1],["delivery",2],["inventory",2],["waybill",2],["production",2],["formula",2]].forEach((function(e){var r=a.audit.filter((function(r){return"trans"===r.nervatypeName&&r.subtypeName===e[0]}))[0];"disabled"!==(a=d()(a,{audit_filter:{trans:{$merge:(0,t.Z)({},e[0],r?[r.inputfilterName,r.supervisor]:["all",1])}}})).audit_filter.trans[e[0]][0]&&(a=d()(a,{edit_new:(0,t.Z)({},e[1],{$push:[e[0]]})}))})),["customer","product","employee","tool","project","setting","audit"].forEach((function(e){var r=a.audit.filter((function(r){return r.nervatypeName===e&&null===r.subtypeName}))[0];"disabled"!==(a=d()(a,{audit_filter:{$merge:(0,t.Z)({},e,r?[r.inputfilterName,r.supervisor]:["all",1])}})).audit_filter[e][0]&&"setting"!==e&&"audit"!==e&&(a=d()(a,{edit_new:{3:{$push:[e]}}}))})),e.abrupt("return",a);case 25:case"end":return e.stop()}}),e)})));return function(r){return e.apply(this,arguments)}}(),f.onLogin=(0,n.Z)(i().mark((function e(){var r,t,n,s;return i().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return r={method:"POST",data:{username:f.data.username,password:f.data.password,database:f.data.database}},e.next=3,c.requestData("/auth/login",r);case 3:if(!(t=e.sent).token||!t.engine){e.next=28;break}if(a.session.engines.includes(t.engine)){e.next=7;break}return e.abrupt("return",c.resultError({error:{message:c.getText("login_engine_err")}}));case 7:if(a.session.service.includes(t.version)){e.next=9;break}return e.abrupt("return",c.resultError({error:{message:c.getText("login_version_err")}}));case 9:return e.next=11,f.loginData(t);case 11:if(!(n=e.sent).error){e.next=14;break}return e.abrupt("return",c.resultError(n));case 14:if("t"!==n.userlogin&&"true"!==n.userlogin){e.next=20;break}return e.next=17,f.userLog(n);case 17:if(!(s=e.sent).error){e.next=20;break}return e.abrupt("return",c.resultError(s));case 20:o("current",{module:"search"}),o(f.key,{data:n}),localStorage.setItem("database",f.data.database),localStorage.setItem("username",f.data.username),localStorage.setItem("server",f.data.server),c.loadBookmark({user_id:n.employee.id,token:t.token}),e.next=29;break;case 28:c.resultError(t);case 29:case"end":return e.stop()}}),e)}))),f.changeData=function(e,r){o(f.key,(0,t.Z)({},e,r))},f.setTheme=function(){var e="light"===f.current.theme?"dark":"light";o("current",{theme:e}),localStorage.setItem("theme",e)},f.setLocale=function(e){o("current",{lang:e}),localStorage.setItem("lang",e)},(0,T.jsx)(C,(0,s.Z)({},f))};$.defaultProps=(0,s.Z)((0,s.Z)({key:"login"},D.defaultProps),{},{data:(0,s.Z)((0,s.Z)({},D.defaultProps.data),{},{data:{}}),userLog:void 0,loginData:void 0});var O=$},4925:function(e,r,a){function t(e,r){if(null==e)return{};var a,t,n=function(e,r){if(null==e)return{};var a,t,n={},s=Object.keys(e);for(t=0;t<s.length;t++)a=s[t],r.indexOf(a)>=0||(n[a]=e[a]);return n}(e,r);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(t=0;t<s.length;t++)a=s[t],r.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(e,a)&&(n[a]=e[a])}return n}a.d(r,{Z:function(){return t}})}}]);