(self.webpackChunknervatura_docs=self.webpackChunknervatura_docs||[]).push([[9777],{3905:function(e,t,a){"use strict";a.d(t,{Zo:function(){return c},kt:function(){return m}});var n=a(7294);function r(e,t,a){return t in e?Object.defineProperty(e,t,{value:a,enumerable:!0,configurable:!0,writable:!0}):e[t]=a,e}function o(e,t){var a=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),a.push.apply(a,n)}return a}function i(e){for(var t=1;t<arguments.length;t++){var a=null!=arguments[t]?arguments[t]:{};t%2?o(Object(a),!0).forEach((function(t){r(e,t,a[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(a)):o(Object(a)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(a,t))}))}return e}function l(e,t){if(null==e)return{};var a,n,r=function(e,t){if(null==e)return{};var a,n,r={},o=Object.keys(e);for(n=0;n<o.length;n++)a=o[n],t.indexOf(a)>=0||(r[a]=e[a]);return r}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)a=o[n],t.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(e,a)&&(r[a]=e[a])}return r}var d=n.createContext({}),u=function(e){var t=n.useContext(d),a=t;return e&&(a="function"==typeof e?e(t):i(i({},t),e)),a},c=function(e){var t=u(e.components);return n.createElement(d.Provider,{value:t},e.children)},s={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},p=n.forwardRef((function(e,t){var a=e.components,r=e.mdxType,o=e.originalType,d=e.parentName,c=l(e,["components","mdxType","originalType","parentName"]),p=u(a),m=r,h=p["".concat(d,".").concat(m)]||p[m]||s[m]||o;return a?n.createElement(h,i(i({ref:t},c),{},{components:a})):n.createElement(h,i({ref:t},c))}));function m(e,t){var a=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var o=a.length,i=new Array(o);i[0]=p;var l={};for(var d in t)hasOwnProperty.call(t,d)&&(l[d]=t[d]);l.originalType=e,l.mdxType="string"==typeof e?e:r,i[1]=l;for(var u=2;u<o;u++)i[u]=a[u];return n.createElement.apply(null,i)}return n.createElement.apply(null,a)}p.displayName="MDXCreateElement"},9692:function(e,t,a){"use strict";a.r(t),a.d(t,{frontMatter:function(){return l},contentTitle:function(){return d},metadata:function(){return u},toc:function(){return c},default:function(){return p}});var n=a(2122),r=a(9756),o=(a(7294),a(3905)),i=["components"],l={id:"production",title:"PRODUCTION",sidebar_label:"Production",hide_table_of_contents:!1},d=void 0,u={unversionedId:"production",id:"production",isDocsHomePage:!1,title:"PRODUCTION",description:"Overview",source:"@site/docs/production.md",sourceDirName:".",slug:"/production",permalink:"/nervatura/docs/production",version:"current",frontMatter:{id:"production",title:"PRODUCTION",sidebar_label:"Production",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"Tool movement",permalink:"/nervatura/docs/waybill"},next:{title:"Formula",permalink:"/nervatura/docs/formula"}},c=[{value:"Overview",id:"overview",children:[]},{value:"Input fields",id:"input-fields",children:[{value:"Document No.",id:"document-no",children:[]},{value:"Creation",id:"creation",children:[]},{value:"Closed",id:"closed",children:[]},{value:"State",id:"state",children:[]},{value:"Start Date",id:"start-date",children:[]},{value:"End Date",id:"end-date",children:[]},{value:"Product No.",id:"product-no",children:[]},{value:"Reference No.",id:"reference-no",children:[]},{value:"Warehouse",id:"warehouse",children:[]},{value:"Batch No.",id:"batch-no",children:[]},{value:"Quantity",id:"quantity",children:[]},{value:"Comment",id:"comment",children:[]},{value:"Internal notes",id:"internal-notes",children:[]}]},{value:"Related data",id:"related-data",children:[{value:"METADATA",id:"metadata",children:[]},{value:"REPORT NOTES",id:"report-notes",children:[]},{value:"DOCUMENT ITEM",id:"document-item",children:[]}]},{value:"Operations",id:"operations",children:[{value:"COPY FROM",id:"copy-from",children:[]},{value:"LOAD FORMULA",id:"load-formula",children:[]},{value:"REPORT",id:"report",children:[]},{value:"BOOKMARK",id:"bookmark",children:[]}]}],s={toc:c};function p(e){var t=e.components,a=(0,r.Z)(e,i);return(0,o.kt)("wrapper",(0,n.Z)({},s,a,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h2",{id:"overview"},"Overview"),(0,o.kt)("p",null,"PRODUCTION helps to produce new products from raw material stored in the warehouses. The inventory level will decrease by the amount of raw materials used, and the manufactured new product quantity will appear in stock. The data sheet tracks the material usage. If other costs, resource usage (eg. used energy, time spent, tool used etc.) is needed to be tracked, then through additional data it can be linked to ",(0,o.kt)("a",{parentName:"p",href:"document"},(0,o.kt)("strong",{parentName:"a"},"WORKSHEET"))," forms as well."),(0,o.kt)("h2",{id:"input-fields"},"Input fields"),(0,o.kt)("h3",{id:"document-no"},"Document No."),(0,o.kt)("p",null,"Unique ID, generated at the first data save. The format and value of the next data in row is taken from the ",(0,o.kt)("a",{parentName:"p",href:"numberdef"},(0,o.kt)("strong",{parentName:"a"},"DOCUMENT NUMBERING"))," (code = production_transfer) data series."),(0,o.kt)("h3",{id:"creation"},"Creation"),(0,o.kt)("p",null,"Date of creation. Automatic value, cannot be changed."),(0,o.kt)("h3",{id:"closed"},"Closed"),(0,o.kt)("p",null,(0,o.kt)("em",{parentName:"p"},"Technical")," closing of the document."),(0,o.kt)("div",{className:"admonition admonition-caution alert alert--warning"},(0,o.kt)("div",{parentName:"div",className:"admonition-heading"},(0,o.kt)("h5",{parentName:"div"},(0,o.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,o.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"16",height:"16",viewBox:"0 0 16 16"},(0,o.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M8.893 1.5c-.183-.31-.52-.5-.887-.5s-.703.19-.886.5L.138 13.499a.98.98 0 0 0 0 1.001c.193.31.53.501.886.501h13.964c.367 0 .704-.19.877-.5a1.03 1.03 0 0 0 .01-1.002L8.893 1.5zm.133 11.497H6.987v-2.003h2.039v2.003zm0-3.004H6.987V5.987h2.039v4.006z"}))),"caution")),(0,o.kt)("div",{parentName:"div",className:"admonition-content"},(0,o.kt)("p",{parentName:"div"},"If set, document data become read only. ",(0,o.kt)("strong",{parentName:"p"},"Marking is not revocable on the user interface!")))),(0,o.kt)("h3",{id:"state"},"State"),(0,o.kt)("p",null,"Its value and editing possibility is linked to ",(0,o.kt)("a",{parentName:"p",href:"usergroup#supervisor"},(0,o.kt)("strong",{parentName:"a"},"ACCESS RIGHTS"))," setting. Not used in current version."),(0,o.kt)("h3",{id:"start-date"},"Start Date"),(0,o.kt)("p",null,"Production starting date. Mandatory."),(0,o.kt)("h3",{id:"end-date"},"End Date"),(0,o.kt)("p",null,"Production ending date. The day, when the quantity will appear in the stock of the given warehouse."),(0,o.kt)("h3",{id:"product-no"},"Product No."),(0,o.kt)("p",null,"One of the items of ",(0,o.kt)("a",{parentName:"p",href:"product"},(0,o.kt)("strong",{parentName:"a"},"PRODUCT")),", the subject of the production, the form is related to this product. Mandatory. Value can be defined with a search field or with the barcode of the product."),(0,o.kt)("h3",{id:"reference-no"},"Reference No."),(0,o.kt)("p",null,"Other reference number. Optional, its value can be freely defined."),(0,o.kt)("h3",{id:"warehouse"},"Warehouse"),(0,o.kt)("p",null,"The produced quantity will be shown in this warehouse. A warehouse can be chosen from the search field from ",(0,o.kt)("em",{parentName:"p"},"warehouse")," type items of ",(0,o.kt)("a",{parentName:"p",href:"place#type"},(0,o.kt)("strong",{parentName:"a"},"PLACE")),". Mandatory."),(0,o.kt)("h3",{id:"batch-no"},"Batch No."),(0,o.kt)("p",null,"The quantity is put into this group. Usage is optional."),(0,o.kt)("h3",{id:"quantity"},"Quantity"),(0,o.kt)("p",null,"Quantity of produced goods."),(0,o.kt)("h3",{id:"comment"},"Comment"),(0,o.kt)("p",null,"Remarks field."),(0,o.kt)("h3",{id:"internal-notes"},"Internal notes"),(0,o.kt)("p",null,"Internal comments. Text defined in this field will not appear on the document."),(0,o.kt)("h2",{id:"related-data"},"Related data"),(0,o.kt)("h3",{id:"metadata"},(0,o.kt)("a",{parentName:"h3",href:"metadata"},(0,o.kt)("strong",{parentName:"a"},"METADATA"))),(0,o.kt)("p",null,"Unlimited number of supplementary data can be added."),(0,o.kt)("h3",{id:"report-notes"},(0,o.kt)("a",{parentName:"h3",href:"notes"},(0,o.kt)("strong",{parentName:"a"},"REPORT NOTES"))),(0,o.kt)("p",null,"Editable remarks, data for reports."),(0,o.kt)("h3",{id:"document-item"},"DOCUMENT ITEM"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},"Shipping Date"),": The usage date of raw materials. The program will deduct the\nquantity from the given warehouse stock on this day."),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},"Warehouse"),": The raw material quantity will be deducted from this warehouse by the program. A warehouse can be chosen from the search field from ",(0,o.kt)("em",{parentName:"li"},"warehouse")," type items of ",(0,o.kt)("a",{parentName:"li",href:"place#type"},(0,o.kt)("strong",{parentName:"a"},"PLACE")),". Mandatory."),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},"Product No."),": Raw material used for production. Mandatory."),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},"Batch No."),": The quantity is taken off from this group. Usage is optional."),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},"Quantity"),": The used quantity. Should not contain any directional sign, will\nautomatically interpreted as stock decrease by the program. If negative value is given,\nthen it will increase the stock level (eg. by-product formation).")),(0,o.kt)("h2",{id:"operations"},"Operations"),(0,o.kt)("h3",{id:"copy-from"},"COPY FROM"),(0,o.kt)("p",null,"Create a new, ",(0,o.kt)("em",{parentName:"p"},"same transaction type")," document on the basis of current document's data. "),(0,o.kt)("p",null,"The dates and information related to creation gets updated and the references from the original document will not be transferred either."),(0,o.kt)("h3",{id:"load-formula"},"LOAD FORMULA"),(0,o.kt)("p",null,"The program will display the ",(0,o.kt)("a",{parentName:"p",href:"formula"},(0,o.kt)("strong",{parentName:"a"},"FORMULA"))," templates available for the produced product in a drop down list. The list of raw materials will be loaded according to the chosen template. ",(0,o.kt)("strong",{parentName:"p"},"If there were lines already, those will be deleted!")," Raw material quantity is taken proportionally to the product quantity to be produced."),(0,o.kt)("h3",{id:"report"},"REPORT"),(0,o.kt)("p",null,(0,o.kt)("a",{parentName:"p",href:"export"},(0,o.kt)("strong",{parentName:"a"},"DATA EXPORT"))),(0,o.kt)("h3",{id:"bookmark"},"BOOKMARK"),(0,o.kt)("p",null,"Set a bookmark for the record. Later can be loaded from bookmarks at any time."))}p.isMDXComponent=!0}}]);