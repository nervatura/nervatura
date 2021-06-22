(self.webpackChunknervatura_docs=self.webpackChunknervatura_docs||[]).push([[3674],{3905:function(e,t,a){"use strict";a.d(t,{Zo:function(){return d},kt:function(){return m}});var n=a(67294);function r(e,t,a){return t in e?Object.defineProperty(e,t,{value:a,enumerable:!0,configurable:!0,writable:!0}):e[t]=a,e}function i(e,t){var a=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),a.push.apply(a,n)}return a}function o(e){for(var t=1;t<arguments.length;t++){var a=null!=arguments[t]?arguments[t]:{};t%2?i(Object(a),!0).forEach((function(t){r(e,t,a[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(a)):i(Object(a)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(a,t))}))}return e}function l(e,t){if(null==e)return{};var a,n,r=function(e,t){if(null==e)return{};var a,n,r={},i=Object.keys(e);for(n=0;n<i.length;n++)a=i[n],t.indexOf(a)>=0||(r[a]=e[a]);return r}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)a=i[n],t.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(e,a)&&(r[a]=e[a])}return r}var s=n.createContext({}),c=function(e){var t=n.useContext(s),a=t;return e&&(a="function"==typeof e?e(t):o(o({},t),e)),a},d=function(e){var t=c(e.components);return n.createElement(s.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},p=n.forwardRef((function(e,t){var a=e.components,r=e.mdxType,i=e.originalType,s=e.parentName,d=l(e,["components","mdxType","originalType","parentName"]),p=c(a),m=r,h=p["".concat(s,".").concat(m)]||p[m]||u[m]||i;return a?n.createElement(h,o(o({ref:t},d),{},{components:a})):n.createElement(h,o({ref:t},d))}));function m(e,t){var a=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var i=a.length,o=new Array(i);o[0]=p;var l={};for(var s in t)hasOwnProperty.call(t,s)&&(l[s]=t[s]);l.originalType=e,l.mdxType="string"==typeof e?e:r,o[1]=l;for(var c=2;c<i;c++)o[c]=a[c];return n.createElement.apply(null,o)}return n.createElement.apply(null,a)}p.displayName="MDXCreateElement"},27945:function(e,t,a){"use strict";a.r(t),a.d(t,{frontMatter:function(){return l},contentTitle:function(){return s},metadata:function(){return c},toc:function(){return d},default:function(){return p}});var n=a(22122),r=a(19756),i=(a(67294),a(3905)),o=["components"],l={id:"metadata",title:"METADATA",sidebar_label:"Metadata",hide_table_of_contents:!1},s=void 0,c={unversionedId:"metadata",id:"metadata",isDocsHomePage:!1,title:"METADATA",description:"Overview",source:"@site/docs/metadata.md",sourceDirName:".",slug:"/metadata",permalink:"/nervatura/docs/metadata",version:"current",frontMatter:{id:"metadata",title:"METADATA",sidebar_label:"Metadata",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"Menu shortcuts",permalink:"/nervatura/docs/uimenu"},next:{title:"Groups",permalink:"/nervatura/docs/groups"}},d=[{value:"Overview",id:"overview",children:[]},{value:"Input fields",id:"input-fields",children:[{value:"Data GUID",id:"data-guid",children:[]},{value:"Description",id:"description",children:[]},{value:"Data Type",id:"data-type",children:[]},{value:"Value Type",id:"value-type",children:[]},{value:"Auto create",id:"auto-create",children:[]},{value:"Visible",id:"visible",children:[]},{value:"Readonly",id:"readonly",children:[]}]}],u={toc:d};function p(e){var t=e.components,a=(0,r.Z)(e,o);return(0,i.kt)("wrapper",(0,n.Z)({},u,a,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h2",{id:"overview"},"Overview"),(0,i.kt)("p",null,"With Nervatura it is easy to store a variety of data. If some new information is needed for which there have not been any data collected yet, the case is simple. Just create a new attribute, specify its type and connect it to the data you would like to use it with."),(0,i.kt)("h2",{id:"input-fields"},"Input fields"),(0,i.kt)("h3",{id:"data-guid"},"Data GUID"),(0,i.kt)("p",null,"Automatically generated internal ID. A unique value, can not be changed."),(0,i.kt)("h3",{id:"description"},"Description"),(0,i.kt)("p",null,"The name set here will be displayed on the user interface."),(0,i.kt)("h3",{id:"data-type"},"Data Type"),(0,i.kt)("p",null,"Valid values: customer, employee, event, place, product, project, tool, trans. Those data types which we would like the new feature to be connected to. If the trans value is selected then it can be used on all forms under DOCUMENT, PAYMENT and STOCK CONTROL menus (eg. offer, order, invoice, etc.)."),(0,i.kt)("h3",{id:"value-type"},"Value Type"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("strong",{parentName:"li"},"bool"),": Two positions, YES/NO or TRUE/FALSE values"),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("strong",{parentName:"li"},"integer, float"),": Only numbers can be entered, in case of integer only numbers without decimals"),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("strong",{parentName:"li"},"date, time"),": Only valid date or time can be set"),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("strong",{parentName:"li"},"string, notes"),": Any text type value. In case of string shorter, for notes longer editable text."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("strong",{parentName:"li"},"password"),': For data entry a "password" type field will be displayed, where the typed in characters cannot be identified. Note that the value in the database is not going to be encrypted, stored only as plain text!'),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("strong",{parentName:"li"},"valuelist"),": After saving a ",(0,i.kt)("strong",{parentName:"li"},"Value list")," field will be displayed, in which a list of items can be entered. The elements can be separated from each other by | sign. When used, only those values will be allowed to enter which have been set in this list, other value is not accepted."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("strong",{parentName:"li"},"urlink"),": Any URL link."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("strong",{parentName:"li"},"customer, employee, place, product, project, tool"),": only valid data stored in the database and chosen through the search engine can be used."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("strong",{parentName:"li"},"transitem"),": we can choose from data in DOCUMENT menu (eg offer, order, invoice, etc.) by using the search tool."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("strong",{parentName:"li"},"transpayment"),": we can choose from data in PAYMENT menu (bank, petty cash) by using the search tool."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("strong",{parentName:"li"},"transmovement"),": we can choose from data in STOCK CONTROL menu (delivery, inventory, etc.) by using the search tool.")),(0,i.kt)("h3",{id:"auto-create"},"Auto create"),(0,i.kt)("p",null,"When selected, the attribute in case of adding a new element (eg a new customer or employee is created) will automatically be created with the default value according to its type and also will be attached to the new element."),(0,i.kt)("h3",{id:"visible"},"Visible"),(0,i.kt)("p",null,"Can appear or not (hidden value) on the entry forms."),(0,i.kt)("h3",{id:"readonly"},"Readonly"),(0,i.kt)("p",null,"The value of the attribute can not be changed in the program interface."),(0,i.kt)("div",{className:"admonition admonition-tip alert alert--success"},(0,i.kt)("div",{parentName:"div",className:"admonition-heading"},(0,i.kt)("h5",{parentName:"div"},(0,i.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,i.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"12",height:"16",viewBox:"0 0 12 16"},(0,i.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M6.5 0C3.48 0 1 2.19 1 5c0 .92.55 2.25 1 3 1.34 2.25 1.78 2.78 2 4v1h5v-1c.22-1.22.66-1.75 2-4 .45-.75 1-2.08 1-3 0-2.81-2.48-5-5.5-5zm3.64 7.48c-.25.44-.47.8-.67 1.11-.86 1.41-1.25 2.06-1.45 3.23-.02.05-.02.11-.02.17H5c0-.06 0-.13-.02-.17-.2-1.17-.59-1.83-1.45-3.23-.2-.31-.42-.67-.67-1.11C2.44 6.78 2 5.65 2 5c0-2.2 2.02-4 4.5-4 1.22 0 2.36.42 3.22 1.19C10.55 2.94 11 3.94 11 5c0 .66-.44 1.78-.86 2.48zM4 14h5c-.23 1.14-1.3 2-2.5 2s-2.27-.86-2.5-2z"}))),"tip")),(0,i.kt)("div",{parentName:"div",className:"admonition-content"},(0,i.kt)("p",{parentName:"div"},"The additional data can be useful also to group our customers, products (value list data type)."))))}p.isMDXComponent=!0}}]);