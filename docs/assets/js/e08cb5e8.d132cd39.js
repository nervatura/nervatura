(self.webpackChunknervatura_docs=self.webpackChunknervatura_docs||[]).push([[8128],{3905:function(e,t,n){"use strict";n.d(t,{Zo:function(){return u},kt:function(){return m}});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},o=Object.keys(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var c=r.createContext({}),d=function(e){var t=r.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},u=function(e){var t=d(e.components);return r.createElement(c.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},s=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,o=e.originalType,c=e.parentName,u=l(e,["components","mdxType","originalType","parentName"]),s=d(n),m=a,f=s["".concat(c,".").concat(m)]||s[m]||p[m]||o;return n?r.createElement(f,i(i({ref:t},u),{},{components:n})):r.createElement(f,i({ref:t},u))}));function m(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var o=n.length,i=new Array(o);i[0]=s;var l={};for(var c in t)hasOwnProperty.call(t,c)&&(l[c]=t[c]);l.originalType=e,l.mdxType="string"==typeof e?e:a,i[1]=l;for(var d=2;d<o;d++)i[d]=n[d];return r.createElement.apply(null,i)}return r.createElement.apply(null,n)}s.displayName="MDXCreateElement"},6807:function(e,t,n){"use strict";n.r(t),n.d(t,{frontMatter:function(){return l},contentTitle:function(){return c},metadata:function(){return d},toc:function(){return u},default:function(){return s}});var r=n(2122),a=n(9756),o=(n(7294),n(3905)),i=["components"],l={id:"waybill",title:"TOOL MOVEMENT",sidebar_label:"Tool movement",hide_table_of_contents:!1},c=void 0,d={unversionedId:"waybill",id:"waybill",isDocsHomePage:!1,title:"TOOL MOVEMENT",description:"Overview",source:"@site/docs/waybill.md",sourceDirName:".",slug:"/waybill",permalink:"/nervatura/docs/waybill",tags:[],version:"current",frontMatter:{id:"waybill",title:"TOOL MOVEMENT",sidebar_label:"Tool movement",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"Inventory control",permalink:"/nervatura/docs/inventory"},next:{title:"Production",permalink:"/nervatura/docs/production"}},u=[{value:"Overview",id:"overview",children:[]},{value:"Input fields",id:"input-fields",children:[{value:"Document No.",id:"document-no",children:[]},{value:"Direction",id:"direction",children:[]},{value:"Creation",id:"creation",children:[]},{value:"State",id:"state",children:[]},{value:"Reference Type",id:"reference-type",children:[]},{value:"Reference",id:"reference",children:[]},{value:"Comment",id:"comment",children:[]},{value:"Internal notes",id:"internal-notes",children:[]}]},{value:"Related data",id:"related-data",children:[{value:"METADATA",id:"metadata",children:[]},{value:"REPORT NOTES",id:"report-notes",children:[]},{value:"DOCUMENT ITEM",id:"document-item",children:[]}]},{value:"Operations",id:"operations",children:[{value:"COPY FROM",id:"copy-from",children:[]},{value:"REPORT",id:"report",children:[]},{value:"BOOKMARK",id:"bookmark",children:[]}]}],p={toc:u};function s(e){var t=e.components,n=(0,a.Z)(e,i);return(0,o.kt)("wrapper",(0,r.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h2",{id:"overview"},"Overview"),(0,o.kt)("p",null,"The movement of ",(0,o.kt)("a",{parentName:"p",href:"product"},(0,o.kt)("strong",{parentName:"a"},"products"))," in warehouses can be tracked with INVENTORY and DELIVERY document types. To track the moves of ",(0,o.kt)("a",{parentName:"p",href:"tool"},(0,o.kt)("strong",{parentName:"a"},"tools"))," the WAYBILL type should be used. It helps to connect  the tool for a certain period to a ",(0,o.kt)("a",{parentName:"p",href:"customer"},(0,o.kt)("strong",{parentName:"a"},"customer")),", ",(0,o.kt)("a",{parentName:"p",href:"employee"},(0,o.kt)("strong",{parentName:"a"},"employee"))," or ",(0,o.kt)("a",{parentName:"p",href:"document"},(0,o.kt)("strong",{parentName:"a"},"document")),"  (orders, worksheets, rental, invoice)."),(0,o.kt)("p",null,"The forms to provide easy handling are designed to enable connecting multiple tools to a customer, employee or document with one data sheet. This allows to easily manage cases like following up on  equipment being handed out to or taken back from your employees, as well as tracking the tools being used for a given project."),(0,o.kt)("h2",{id:"input-fields"},"Input fields"),(0,o.kt)("h3",{id:"document-no"},"Document No."),(0,o.kt)("p",null,"Unique ID, generated at the first data save. The format and value of the next data in row is taken from the ",(0,o.kt)("a",{parentName:"p",href:"numberdef"},(0,o.kt)("strong",{parentName:"a"},"DOCUMENT NUMBERING"))," (code = waybill) data series."),(0,o.kt)("h3",{id:"direction"},"Direction"),(0,o.kt)("p",null,"OUT, IN. Value cannot be changed after first save!"),(0,o.kt)("h3",{id:"creation"},"Creation"),(0,o.kt)("p",null,"Date of creation. Automatic value, cannot be changed."),(0,o.kt)("h3",{id:"state"},"State"),(0,o.kt)("p",null,"Its value and editing possibility is linked to ",(0,o.kt)("a",{parentName:"p",href:"usergroup#supervisor"},(0,o.kt)("strong",{parentName:"a"},"ACCESS RIGHTS"))," setting. Not used in current version."),(0,o.kt)("h3",{id:"reference-type"},"Reference Type"),(0,o.kt)("p",null,"DOCUMENT, CUSTOMER, EMPLOYEE"),(0,o.kt)("h3",{id:"reference"},"Reference"),(0,o.kt)("p",null,"Depending on the Reference Type, an ID for DOCUMENT, CUSTOMER, EMPLOYEE."),(0,o.kt)("h3",{id:"comment"},"Comment"),(0,o.kt)("p",null,"Remarks field."),(0,o.kt)("h3",{id:"internal-notes"},"Internal notes"),(0,o.kt)("p",null,"Internal comments. Text defined in this field will not appear on the document."),(0,o.kt)("h2",{id:"related-data"},"Related data"),(0,o.kt)("h3",{id:"metadata"},(0,o.kt)("a",{parentName:"h3",href:"metadata"},(0,o.kt)("strong",{parentName:"a"},"METADATA"))),(0,o.kt)("p",null,"Unlimited number of supplementary data can be added."),(0,o.kt)("h3",{id:"report-notes"},(0,o.kt)("a",{parentName:"h3",href:"notes"},(0,o.kt)("strong",{parentName:"a"},"REPORT NOTES"))),(0,o.kt)("p",null,"Editable remarks, data for reports."),(0,o.kt)("h3",{id:"document-item"},"DOCUMENT ITEM"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},"Shipping Date"),": Stock release or receipt date."),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},"Serial"),": One of the items of ",(0,o.kt)("a",{parentName:"li",href:"tool"},(0,o.kt)("strong",{parentName:"a"},"TOOL")),"."),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},"Comment"),": Other remarks, data.")),(0,o.kt)("h2",{id:"operations"},"Operations"),(0,o.kt)("h3",{id:"copy-from"},"COPY FROM"),(0,o.kt)("p",null,"Create a new, ",(0,o.kt)("em",{parentName:"p"},"same transaction type")," document on the basis of current document's data. "),(0,o.kt)("p",null,"The dates and information related to creation gets updated and the references from the original document will not be transferred either."),(0,o.kt)("h3",{id:"report"},"REPORT"),(0,o.kt)("p",null,(0,o.kt)("a",{parentName:"p",href:"export"},(0,o.kt)("strong",{parentName:"a"},"DATA EXPORT"))),(0,o.kt)("h3",{id:"bookmark"},"BOOKMARK"),(0,o.kt)("p",null,"Set a bookmark for the record. Later can be loaded from bookmarks at any time."))}s.isMDXComponent=!0}}]);