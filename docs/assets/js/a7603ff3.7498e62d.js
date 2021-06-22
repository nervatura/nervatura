(self.webpackChunknervatura_docs=self.webpackChunknervatura_docs||[]).push([[8468],{3905:function(e,t,n){"use strict";n.d(t,{Zo:function(){return u},kt:function(){return f}});var r=n(67294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},o=Object.keys(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var d=r.createContext({}),c=function(e){var t=r.useContext(d),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},u=function(e){var t=c(e.components);return r.createElement(d.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},s=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,o=e.originalType,d=e.parentName,u=l(e,["components","mdxType","originalType","parentName"]),s=c(n),f=a,m=s["".concat(d,".").concat(f)]||s[f]||p[f]||o;return n?r.createElement(m,i(i({ref:t},u),{},{components:n})):r.createElement(m,i({ref:t},u))}));function f(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var o=n.length,i=new Array(o);i[0]=s;var l={};for(var d in t)hasOwnProperty.call(t,d)&&(l[d]=t[d]);l.originalType=e,l.mdxType="string"==typeof e?e:a,i[1]=l;for(var c=2;c<o;c++)i[c]=n[c];return r.createElement.apply(null,i)}return r.createElement.apply(null,n)}s.displayName="MDXCreateElement"},84568:function(e,t,n){"use strict";n.r(t),n.d(t,{frontMatter:function(){return l},contentTitle:function(){return d},metadata:function(){return c},toc:function(){return u},default:function(){return s}});var r=n(22122),a=n(19756),o=(n(67294),n(3905)),i=["components"],l={id:"event",title:"EVENT",sidebar_label:"Event",hide_table_of_contents:!1},d=void 0,c={unversionedId:"event",id:"event",isDocsHomePage:!1,title:"EVENT",description:"Input fields",source:"@site/docs/event.md",sourceDirName:".",slug:"/event",permalink:"/nervatura/docs/event",version:"current",frontMatter:{id:"event",title:"EVENT",sidebar_label:"Event",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"Project",permalink:"/nervatura/docs/project"},next:{title:"Default settings",permalink:"/nervatura/docs/setting"}},u=[{value:"Input fields",id:"input-fields",children:[{value:"Event No.",id:"event-no",children:[]},{value:"Subject",id:"subject",children:[]},{value:"Place",id:"place",children:[]},{value:"Group",id:"group",children:[]},{value:"Start Date, End Date",id:"start-date-end-date",children:[]},{value:"Comment",id:"comment",children:[]}]},{value:"Related data",id:"related-data",children:[{value:"METADATA",id:"metadata",children:[]}]},{value:"Operations",id:"operations",children:[{value:"ICAL EXPORT",id:"ical-export",children:[]},{value:"BOOKMARK",id:"bookmark",children:[]}]}],p={toc:u};function s(e){var t=e.components,n=(0,a.Z)(e,i);return(0,o.kt)("wrapper",(0,r.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h2",{id:"input-fields"},"Input fields"),(0,o.kt)("h3",{id:"event-no"},"Event No."),(0,o.kt)("p",null,"Unique ID, generated at the first data save. The format and value of the next data in row is taken from the ",(0,o.kt)("a",{parentName:"p",href:"numberdef"},(0,o.kt)("strong",{parentName:"a"},"DOCUMENT NUMBERING"))," (code = calnumber) data series."),(0,o.kt)("h3",{id:"subject"},"Subject"),(0,o.kt)("p",null,"Brief description of the event."),(0,o.kt)("h3",{id:"place"},"Place"),(0,o.kt)("p",null,"Optional. The venue of the event."),(0,o.kt)("h3",{id:"group"},"Group"),(0,o.kt)("p",null,"Optional. In this field a valid element of ",(0,o.kt)("a",{parentName:"p",href:"groups"},(0,o.kt)("strong",{parentName:"a"},"eventgroup"))," group should be given."),(0,o.kt)("h3",{id:"start-date-end-date"},"Start Date, End Date"),(0,o.kt)("p",null,"Optional. The event's start and end dates."),(0,o.kt)("h3",{id:"comment"},"Comment"),(0,o.kt)("p",null,"Remarks field."),(0,o.kt)("h2",{id:"related-data"},"Related data"),(0,o.kt)("h3",{id:"metadata"},(0,o.kt)("a",{parentName:"h3",href:"metadata"},(0,o.kt)("strong",{parentName:"a"},"METADATA"))),(0,o.kt)("p",null,"Unlimited number of supplementary data can be added."),(0,o.kt)("h2",{id:"operations"},"Operations"),(0,o.kt)("h3",{id:"ical-export"},"ICAL EXPORT"),(0,o.kt)("p",null,(0,o.kt)("a",{parentName:"p",href:"export"},(0,o.kt)("strong",{parentName:"a"},"DATA EXPORT"))),(0,o.kt)("h3",{id:"bookmark"},"BOOKMARK"),(0,o.kt)("p",null,"Set a bookmark for the record. Later can be loaded from bookmarks at any time."))}s.isMDXComponent=!0}}]);