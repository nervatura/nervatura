(self.webpackChunknervatura_docs=self.webpackChunknervatura_docs||[]).push([[4111],{3905:function(e,t,r){"use strict";r.d(t,{Zo:function(){return p},kt:function(){return d}});var n=r(7294);function a(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function o(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function i(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?o(Object(r),!0).forEach((function(t){a(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):o(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function l(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},o=Object.keys(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var s=n.createContext({}),u=function(e){var t=n.useContext(s),r=t;return e&&(r="function"==typeof e?e(t):i(i({},t),e)),r},p=function(e){var t=u(e.components);return n.createElement(s.Provider,{value:t},e.children)},c={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},m=n.forwardRef((function(e,t){var r=e.components,a=e.mdxType,o=e.originalType,s=e.parentName,p=l(e,["components","mdxType","originalType","parentName"]),m=u(r),d=a,f=m["".concat(s,".").concat(d)]||m[d]||c[d]||o;return r?n.createElement(f,i(i({ref:t},p),{},{components:r})):n.createElement(f,i({ref:t},p))}));function d(e,t){var r=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var o=r.length,i=new Array(o);i[0]=m;var l={};for(var s in t)hasOwnProperty.call(t,s)&&(l[s]=t[s]);l.originalType=e,l.mdxType="string"==typeof e?e:a,i[1]=l;for(var u=2;u<o;u++)i[u]=r[u];return n.createElement.apply(null,i)}return n.createElement.apply(null,r)}m.displayName="MDXCreateElement"},5913:function(e,t,r){"use strict";r.r(t),r.d(t,{frontMatter:function(){return l},contentTitle:function(){return s},metadata:function(){return u},toc:function(){return p},default:function(){return m}});var n=r(2122),a=r(9756),o=(r(7294),r(3905)),i=["components"],l={id:"usergroup",title:"ACCESS RIGHTS",sidebar_label:"Access rights",hide_table_of_contents:!1},s=void 0,u={unversionedId:"usergroup",id:"usergroup",isDocsHomePage:!1,title:"ACCESS RIGHTS",description:"Overview",source:"@site/docs/usergroup.md",sourceDirName:".",slug:"/usergroup",permalink:"/nervatura/docs/usergroup",version:"current",frontMatter:{id:"usergroup",title:"ACCESS RIGHTS",sidebar_label:"Access rights",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"Document numbering",permalink:"/nervatura/docs/numberdef"},next:{title:"Menu shortcuts",permalink:"/nervatura/docs/uimenu"}},p=[{value:"Overview",id:"overview",children:[]},{value:"Input fields",id:"input-fields",children:[{value:"Group",id:"group",children:[]},{value:"Data filter",id:"data-filter",children:[]}]}],c={toc:p};function m(e){var t=e.components,r=(0,a.Z)(e,i);return(0,o.kt)("wrapper",(0,n.Z)({},c,r,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h2",{id:"overview"},"Overview"),(0,o.kt)("p",null,"The user access rights in Nervatura are managed through user groups. The access rules are assigned to these groups and are applicable to group members. When the rule is changed, the new settings will automatically be valid for all members.",(0,o.kt)("br",null),"\nEach user must be a member of an access rights group, but can only be part of one of these groups.  This can be set in ",(0,o.kt)("a",{parentName:"p",href:"employee"},(0,o.kt)("strong",{parentName:"a"},"EMPLOYEE"))," usergroup field."),(0,o.kt)("div",{className:"admonition admonition-important alert alert--info"},(0,o.kt)("div",{parentName:"div",className:"admonition-heading"},(0,o.kt)("h5",{parentName:"div"},(0,o.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,o.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"},(0,o.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M7 2.3c3.14 0 5.7 2.56 5.7 5.7s-2.56 5.7-5.7 5.7A5.71 5.71 0 0 1 1.3 8c0-3.14 2.56-5.7 5.7-5.7zM7 1C3.14 1 0 4.14 0 8s3.14 7 7 7 7-3.14 7-7-3.14-7-7-7zm1 3H6v5h2V4zm0 6H6v2h2v-2z"}))),"important")),(0,o.kt)("div",{parentName:"div",className:"admonition-content"},(0,o.kt)("p",{parentName:"div"},(0,o.kt)("strong",{parentName:"p"},"The rights set here will apply also to other user interfaces of Nervatura"),", such as the functions of ",(0,o.kt)("a",{parentName:"p",href:"/api"},(0,o.kt)("strong",{parentName:"a"},"Nervatura API")),"!"))),(0,o.kt)("h2",{id:"input-fields"},"Input fields"),(0,o.kt)("h3",{id:"group"},"Group"),(0,o.kt)("p",null,"Group ID. A unique value, can not be repeated."),(0,o.kt)("p",null,"An user access rights group by default is created with full access, though the scope can be limited by rules. When a new rule is created the following can be set:"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},"Type"),": audit, customer, employee, event, menu, price, product, project, report, setting, tool, trans"),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},"Subtype"),": just in case of ",(0,o.kt)("em",{parentName:"li"},"report"),", ",(0,o.kt)("em",{parentName:"li"},"menu")," or ",(0,o.kt)("em",{parentName:"li"},"trans")," types. New rule appears after first save.",(0,o.kt)("ul",{parentName:"li"},(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},(0,o.kt)("em",{parentName:"strong"},"report")),": codes of the ",(0,o.kt)("a",{parentName:"li",href:"report"},(0,o.kt)("strong",{parentName:"a"},"reports"))," in the database. If set to ",(0,o.kt)("em",{parentName:"li"},"disabled")," the report is not available for the group members."),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},(0,o.kt)("em",{parentName:"strong"},"menu")),": codes of the ",(0,o.kt)("a",{parentName:"li",href:"uimenu"},(0,o.kt)("strong",{parentName:"a"},"menu shortcuts"))," in the database. If set to ",(0,o.kt)("i",null,"disabled")," the menu shortcut is not available for the group members."),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},(0,o.kt)("em",{parentName:"strong"},"trans")),": bank, cash, delivery, formula, inventory, invoice, offer, order, production, rent, waybill, worksheet. The settings apply only to the specified subtype."))),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},"Filter"),(0,o.kt)("ul",{parentName:"li"},(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},(0,o.kt)("em",{parentName:"strong"},"all")),": there are no restrictions"),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},(0,o.kt)("em",{parentName:"strong"},"disabled")),": cannot be selected or will not be even displayed"),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},(0,o.kt)("em",{parentName:"strong"},"readonly")),": can be opened only as read-only"),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},(0,o.kt)("em",{parentName:"strong"},"update")),": values \u200b\u200bcan be changed, but new value creation is not allowed, also members of the group can not delete them"))),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},"Supervisor"),": related to multi-level access verification, not used in the current version. When set, it regulates the access to forms ",(0,o.kt)("em",{parentName:"li"},"State")," field.")),(0,o.kt)("h3",{id:"data-filter"},"Data filter"),(0,o.kt)("p",null,"Regulates the access to data in DOCUMENT, PAYMENT, STOCK CONTROL menus, as follows:"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},(0,o.kt)("em",{parentName:"strong"},"all")),": full access"),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},(0,o.kt)("em",{parentName:"strong"},"own")),": the user can only see the data created by himself"),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},(0,o.kt)("em",{parentName:"strong"},"usergroup")),": the user can only see the data created by the members of the user group he belongs to")))}m.isMDXComponent=!0}}]);