(self.webpackChunknervatura_docs=self.webpackChunknervatura_docs||[]).push([[5834],{3905:function(e,t,n){"use strict";n.d(t,{Zo:function(){return s},kt:function(){return m}});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},o=Object.keys(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var c=r.createContext({}),d=function(e){var t=r.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},s=function(e){var t=d(e.components);return r.createElement(c.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},u=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,o=e.originalType,c=e.parentName,s=l(e,["components","mdxType","originalType","parentName"]),u=d(n),m=a,f=u["".concat(c,".").concat(m)]||u[m]||p[m]||o;return n?r.createElement(f,i(i({ref:t},s),{},{components:n})):r.createElement(f,i({ref:t},s))}));function m(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var o=n.length,i=new Array(o);i[0]=u;var l={};for(var c in t)hasOwnProperty.call(t,c)&&(l[c]=t[c]);l.originalType=e,l.mdxType="string"==typeof e?e:a,i[1]=l;for(var d=2;d<o;d++)i[d]=n[d];return r.createElement.apply(null,i)}return r.createElement.apply(null,n)}u.displayName="MDXCreateElement"},8509:function(e,t,n){"use strict";n.r(t),n.d(t,{frontMatter:function(){return l},contentTitle:function(){return c},metadata:function(){return d},toc:function(){return s},default:function(){return u}});var r=n(2122),a=n(9756),o=(n(7294),n(3905)),i=["components"],l={id:"notes",title:"REPORT NOTES",sidebar_label:"Report notes",hide_table_of_contents:!1},c=void 0,d={unversionedId:"notes",id:"notes",isDocsHomePage:!1,title:"REPORT NOTES",description:"Overview",source:"@site/docs/notes.md",sourceDirName:".",slug:"/notes",permalink:"/nervatura/docs/notes",version:"current",frontMatter:{id:"notes",title:"REPORT NOTES",sidebar_label:"Report notes",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"Report",permalink:"/nervatura/docs/report"},next:{title:"Report queue",permalink:"/nervatura/docs/printqueue"}},s=[{value:"Overview",id:"overview",children:[]},{value:"Operations",id:"operations",children:[{value:"NEW TEMPLATE",id:"new-template",children:[]},{value:"DELETE",id:"delete",children:[]},{value:"SET DEFAULT",id:"set-default",children:[]},{value:"LOAD FROM...",id:"load-from",children:[]},{value:"SAVE TO...",id:"save-to",children:[]}]}],p={toc:s};function u(e){var t=e.components,n=(0,a.Z)(e,i);return(0,o.kt)("wrapper",(0,r.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h2",{id:"overview"},"Overview"),(0,o.kt)("p",null,"Documents created with ",(0,o.kt)("a",{parentName:"p",href:"editor"},(0,o.kt)("strong",{parentName:"a"},"REPORT TEMPLATE"))," (offer, order, invoice, etc.) can be supplemented with additional information."),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},"You may specify a ",(0,o.kt)("strong",{parentName:"li"},"longer text"),", which may vary for each transaction and can provide the customers with additional information on the delivery terms, warranty rights, opening hours, etc."),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},"Simple formatting")," can also be made to the texts: bold, italic fonts, lists, indents."),(0,o.kt)("li",{parentName:"ul"},"From text ",(0,o.kt)("strong",{parentName:"li"},"predefined text templates")," can be created. The text templates are stored as document types in the program."),(0,o.kt)("li",{parentName:"ul"},"For each document type a ",(0,o.kt)("strong",{parentName:"li"},"default text template")," can be defined. These will be loaded\nand saved by  the program for each newly created document automatically.")),(0,o.kt)("h2",{id:"operations"},"Operations"),(0,o.kt)("h3",{id:"new-template"},"NEW TEMPLATE"),(0,o.kt)("p",null,"Create a new text template. The transactions do not affect the text in the edit box. For adding a new one the program will request the name of the template. The templates are stored as document type."),(0,o.kt)("h3",{id:"delete"},"DELETE"),(0,o.kt)("p",null,"Delete the selected text template."),(0,o.kt)("h3",{id:"set-default"},"SET DEFAULT"),(0,o.kt)("p",null,"To make the selected text template default."),(0,o.kt)("h3",{id:"load-from"},"LOAD FROM..."),(0,o.kt)("p",null,"Loading the selected template text. "),(0,o.kt)("h3",{id:"save-to"},"SAVE TO..."),(0,o.kt)("p",null,"Save the edited text of the selected template."),(0,o.kt)("div",{className:"admonition admonition-caution alert alert--warning"},(0,o.kt)("div",{parentName:"div",className:"admonition-heading"},(0,o.kt)("h5",{parentName:"div"},(0,o.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,o.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"16",height:"16",viewBox:"0 0 16 16"},(0,o.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M8.893 1.5c-.183-.31-.52-.5-.887-.5s-.703.19-.886.5L.138 13.499a.98.98 0 0 0 0 1.001c.193.31.53.501.886.501h13.964c.367 0 .704-.19.877-.5a1.03 1.03 0 0 0 .01-1.002L8.893 1.5zm.133 11.497H6.987v-2.003h2.039v2.003zm0-3.004H6.987V5.987h2.039v4.006z"}))),"caution")),(0,o.kt)("div",{parentName:"div",className:"admonition-content"},(0,o.kt)("p",{parentName:"div"},"When the text template is loaded the text in the field will be overwritten, and when it is saved the text in the template will be the amended one!"))))}u.isMDXComponent=!0}}]);