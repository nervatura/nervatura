(self.webpackChunknervatura_docs=self.webpackChunknervatura_docs||[]).push([[5489],{3905:function(e,t,n){"use strict";n.d(t,{Zo:function(){return s},kt:function(){return p}});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function l(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function u(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var o=r.createContext({}),d=function(e){var t=r.useContext(o),n=t;return e&&(n="function"==typeof e?e(t):l(l({},t),e)),n},s=function(e){var t=d(e.components);return r.createElement(o.Provider,{value:t},e.children)},c={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},f=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,i=e.originalType,o=e.parentName,s=u(e,["components","mdxType","originalType","parentName"]),f=d(n),p=a,h=f["".concat(o,".").concat(p)]||f[p]||c[p]||i;return n?r.createElement(h,l(l({ref:t},s),{},{components:n})):r.createElement(h,l({ref:t},s))}));function p(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=n.length,l=new Array(i);l[0]=f;var u={};for(var o in t)hasOwnProperty.call(t,o)&&(u[o]=t[o]);u.originalType=e,u.mdxType="string"==typeof e?e:a,l[1]=u;for(var d=2;d<i;d++)l[d]=n[d];return r.createElement.apply(null,l)}return r.createElement.apply(null,n)}f.displayName="MDXCreateElement"},7465:function(e,t,n){"use strict";n.r(t),n.d(t,{frontMatter:function(){return u},contentTitle:function(){return o},metadata:function(){return d},toc:function(){return s},default:function(){return f}});var r=n(2122),a=n(9756),i=(n(7294),n(3905)),l=["components"],u={id:"setting",title:"DEFAULT SETTINGS",sidebar_label:"Default settings",hide_table_of_contents:!1},o=void 0,d={unversionedId:"setting",id:"setting",isDocsHomePage:!1,title:"DEFAULT SETTINGS",description:"Overview",source:"@site/docs/setting.md",sourceDirName:".",slug:"/setting",permalink:"/nervatura/docs/setting",tags:[],version:"current",frontMatter:{id:"setting",title:"DEFAULT SETTINGS",sidebar_label:"Default settings",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"Event",permalink:"/nervatura/docs/event"},next:{title:"Document numbering",permalink:"/nervatura/docs/numberdef"}},s=[{value:"Overview",id:"overview",children:[]},{value:"Some important settings",id:"some-important-settings",children:[{value:"business year",id:"business-year",children:[]},{value:"default bank, default petty cash, default warehouse",id:"default-bank-default-petty-cash-default-warehouse",children:[]},{value:"default currency",id:"default-currency",children:[]},{value:"default deadline",id:"default-deadline",children:[]},{value:"default paidtype",id:"default-paidtype",children:[]},{value:"default taxcode",id:"default-taxcode",children:[]},{value:"default unit",id:"default-unit",children:[]}]}],c={toc:s};function f(e){var t=e.components,n=(0,a.Z)(e,l);return(0,i.kt)("wrapper",(0,r.Z)({},c,n,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h2",{id:"overview"},"Overview"),(0,i.kt)("p",null,"The database settings are not necessarily related to Nervatura Touch program. ",(0,i.kt)("strong",{parentName:"p"},"Could affect the server's settings, the method of data storage or even operation of other programs!"),(0,i.kt)("br",null),"\nFor example, the default values set here will be used by the ",(0,i.kt)("a",{parentName:"p",href:"/api"},(0,i.kt)("strong",{parentName:"a"},"Nervatura API"))," functions as well."),(0,i.kt)("h2",{id:"some-important-settings"},"Some important settings"),(0,i.kt)("h3",{id:"business-year"},"business year"),(0,i.kt)("p",null,"Set the current fiscal year. This as used by ",(0,i.kt)("a",{parentName:"p",href:"numberdef"},(0,i.kt)("strong",{parentName:"a"},"DOCUMENT NUMBERING"))," as the value of the Year"),(0,i.kt)("h3",{id:"default-bank-default-petty-cash-default-warehouse"},"default bank, default petty cash, default warehouse"),(0,i.kt)("p",null,"These will be the default values for new input forms. The value for ",(0,i.kt)("a",{parentName:"p",href:"place"},(0,i.kt)("strong",{parentName:"a"},"Place No."))," should be set in the field."),(0,i.kt)("h3",{id:"default-currency"},"default currency"),(0,i.kt)("p",null,"This will be the default value for new input forms. The field should include the ",(0,i.kt)("a",{parentName:"p",href:"currency"},(0,i.kt)("strong",{parentName:"a"},"currency"))," code."),(0,i.kt)("h3",{id:"default-deadline"},"default deadline"),(0,i.kt)("p",null,"Number of days. This will be the basis for calculating the default ",(0,i.kt)("a",{parentName:"p",href:"document#duedate"},(0,i.kt)("strong",{parentName:"a"},"duedate"))," field in case of a new invoice. The value can be effected by the ",(0,i.kt)("a",{parentName:"p",href:"customer"},(0,i.kt)("strong",{parentName:"a"},"CUSTOMER"))," settings of the invoice."),(0,i.kt)("h3",{id:"default-paidtype"},"default paidtype"),(0,i.kt)("p",null,"This will be the default value for new input forms. A valid element of the ",(0,i.kt)("a",{parentName:"p",href:"groups"},(0,i.kt)("strong",{parentName:"a"},"paidtype"))," group should be set in the field."),(0,i.kt)("h3",{id:"default-taxcode"},"default taxcode"),(0,i.kt)("p",null,"This will be the default value for new line items. The field should include the ",(0,i.kt)("em",{parentName:"p"},"code")," field of the ",(0,i.kt)("strong",{parentName:"p"},"TAX"),"."),(0,i.kt)("h3",{id:"default-unit"},"default unit"),(0,i.kt)("p",null,"Setting applies to a default unit field of a new ",(0,i.kt)("a",{parentName:"p",href:"product"},(0,i.kt)("strong",{parentName:"a"},"PRODUCT"))))}f.isMDXComponent=!0}}]);