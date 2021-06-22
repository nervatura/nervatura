(self.webpackChunknervatura_docs=self.webpackChunknervatura_docs||[]).push([[8031],{3905:function(e,t,r){"use strict";r.d(t,{Zo:function(){return s},kt:function(){return p}});var n=r(67294);function i(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function a(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function o(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?a(Object(r),!0).forEach((function(t){i(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):a(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function u(e,t){if(null==e)return{};var r,n,i=function(e,t){if(null==e)return{};var r,n,i={},a=Object.keys(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||(i[r]=e[r]);return i}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(i[r]=e[r])}return i}var c=n.createContext({}),l=function(e){var t=n.useContext(c),r=t;return e&&(r="function"==typeof e?e(t):o(o({},t),e)),r},s=function(e){var t=l(e.components);return n.createElement(c.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},f=n.forwardRef((function(e,t){var r=e.components,i=e.mdxType,a=e.originalType,c=e.parentName,s=u(e,["components","mdxType","originalType","parentName"]),f=l(r),p=i,h=f["".concat(c,".").concat(p)]||f[p]||d[p]||a;return r?n.createElement(h,o(o({ref:t},s),{},{components:r})):n.createElement(h,o({ref:t},s))}));function p(e,t){var r=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var a=r.length,o=new Array(a);o[0]=f;var u={};for(var c in t)hasOwnProperty.call(t,c)&&(u[c]=t[c]);u.originalType=e,u.mdxType="string"==typeof e?e:i,o[1]=u;for(var l=2;l<a;l++)o[l]=r[l];return n.createElement.apply(null,o)}return n.createElement.apply(null,r)}f.displayName="MDXCreateElement"},68942:function(e,t,r){"use strict";r.r(t),r.d(t,{frontMatter:function(){return u},contentTitle:function(){return c},metadata:function(){return l},toc:function(){return s},default:function(){return f}});var n=r(22122),i=r(19756),a=(r(67294),r(3905)),o=["components"],u={id:"numberdef",title:"DOCUMENT NUMBERING",sidebar_label:"Document numbering",hide_table_of_contents:!1},c=void 0,l={unversionedId:"numberdef",id:"numberdef",isDocsHomePage:!1,title:"DOCUMENT NUMBERING",description:"Overview",source:"@site/docs/numberdef.md",sourceDirName:".",slug:"/numberdef",permalink:"/nervatura/docs/numberdef",version:"current",frontMatter:{id:"numberdef",title:"DOCUMENT NUMBERING",sidebar_label:"Document numbering",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"Default settings",permalink:"/nervatura/docs/setting"},next:{title:"Access rights",permalink:"/nervatura/docs/usergroup"}},s=[{value:"Overview",id:"overview",children:[]},{value:"Input fields",id:"input-fields",children:[{value:"Code",id:"code",children:[]},{value:"Prefix",id:"prefix",children:[]},{value:"Year",id:"year",children:[]},{value:"Separator",id:"separator",children:[]},{value:"Lenght",id:"lenght",children:[]},{value:"Value",id:"value",children:[]}]}],d={toc:s};function f(e){var t=e.components,r=(0,i.Z)(e,o);return(0,a.kt)("wrapper",(0,n.Z)({},d,r,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h2",{id:"overview"},"Overview"),(0,a.kt)("p",null,"Format settings of unique identifiers of documents (eg. orders, invoices, cash receipts, delivery notes etc.) and other resources (eg. customers, products, employees, etc.). If a new item is created the identifiers will follow the rule which was set here."),(0,a.kt)("h2",{id:"input-fields"},"Input fields"),(0,a.kt)("h3",{id:"code"},"Code"),(0,a.kt)("p",null,"A unique identifier for a certain set of rules. Its value can not be changed."),(0,a.kt)("h3",{id:"prefix"},"Prefix"),(0,a.kt)("p",null,"The text prefix of the identifier. It can be any length, but usage of special characters,  spaces in the text is not recommended."),(0,a.kt)("h3",{id:"year"},"Year"),(0,a.kt)("p",null,"If selected, the created identifier will contain the year. The number is not formed automatically from the current date, but can be set in the ",(0,a.kt)("a",{parentName:"p",href:"setting"},(0,a.kt)("strong",{parentName:"a"},"DEFAULT SETTINGS"))," section,\nin the ",(0,a.kt)("strong",{parentName:"p"},(0,a.kt)("em",{parentName:"strong"},"business year"))," field. This is due to avoid any technical issues resulting\nfrom changes in different fiscal or calendar years."),(0,a.kt)("h3",{id:"separator"},"Separator"),(0,a.kt)("p",null,"The separator character in the identifier. Default value: /"),(0,a.kt)("h3",{id:"lenght"},"Lenght"),(0,a.kt)("p",null,"The value field is arranged in such length to the right and filled with zeros."),(0,a.kt)("h3",{id:"value"},"Value"),(0,a.kt)("p",null,"The current status of the counter, the next sequence number will be one value higher than this one. It is possible to re-set the counter, but the uniqueness must be ensured in all cases!"))}f.isMDXComponent=!0}}]);