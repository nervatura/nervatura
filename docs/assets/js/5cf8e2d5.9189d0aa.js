(self.webpackChunknervatura_docs=self.webpackChunknervatura_docs||[]).push([[8053],{3905:function(e,t,n){"use strict";n.d(t,{Zo:function(){return m},kt:function(){return s}});var a=n(7294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,a,r=function(e,t){if(null==e)return{};var n,a,r={},i=Object.keys(e);for(a=0;a<i.length;a++)n=i[a],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(a=0;a<i.length;a++)n=i[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var c=a.createContext({}),d=function(e){var t=a.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},m=function(e){var t=d(e.components);return a.createElement(c.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},p=a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,i=e.originalType,c=e.parentName,m=l(e,["components","mdxType","originalType","parentName"]),p=d(n),s=r,h=p["".concat(c,".").concat(s)]||p[s]||u[s]||i;return n?a.createElement(h,o(o({ref:t},m),{},{components:n})):a.createElement(h,o({ref:t},m))}));function s(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var i=n.length,o=new Array(i);o[0]=p;var l={};for(var c in t)hasOwnProperty.call(t,c)&&(l[c]=t[c]);l.originalType=e,l.mdxType="string"==typeof e?e:r,o[1]=l;for(var d=2;d<i;d++)o[d]=n[d];return a.createElement.apply(null,o)}return a.createElement.apply(null,n)}p.displayName="MDXCreateElement"},1072:function(e,t,n){"use strict";n.r(t),n.d(t,{frontMatter:function(){return l},contentTitle:function(){return c},metadata:function(){return d},toc:function(){return m},default:function(){return p}});var a=n(2122),r=n(9756),i=(n(7294),n(3905)),o=["components"],l={id:"payment",title:"PAYMENT",sidebar_label:"Payment",hide_table_of_contents:!1},c=void 0,d={unversionedId:"payment",id:"payment",isDocsHomePage:!1,title:"PAYMENT",description:"Input fields",source:"@site/docs/payment.md",sourceDirName:".",slug:"/payment",permalink:"/nervatura/docs/payment",version:"current",frontMatter:{id:"payment",title:"PAYMENT",sidebar_label:"Payment",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"Document Item",permalink:"/nervatura/docs/item"},next:{title:"Shipping",permalink:"/nervatura/docs/shipping"}},m=[{value:"Input fields",id:"input-fields",children:[{value:"Document No.",id:"document-no",children:[]},{value:"Reference No.",id:"reference-no",children:[]},{value:"Creation",id:"creation",children:[]},{value:"Closed",id:"closed",children:[]},{value:"State",id:"state",children:[]},{value:"Comment",id:"comment",children:[]},{value:"Internal notes",id:"internal-notes",children:[]},{value:"BANK",id:"bank",children:[]},{value:"Account Date",id:"account-date",children:[]},{value:"Bank Account",id:"bank-account",children:[]},{value:"PETTY CASH",id:"petty-cash",children:[]},{value:"Direction",id:"direction",children:[]},{value:"Payment Date",id:"payment-date",children:[]},{value:"Amount",id:"amount",children:[]},{value:"Petty cash",id:"petty-cash-1",children:[]},{value:"Employee No.",id:"employee-no",children:[]}]},{value:"Related data",id:"related-data",children:[{value:"METADATA",id:"metadata",children:[]},{value:"REPORT NOTES",id:"report-notes",children:[]},{value:"INVOICE",id:"invoice",children:[]},{value:"DOCUMENT ITEM",id:"document-item",children:[]}]},{value:"Operations",id:"operations",children:[{value:"COPY FROM",id:"copy-from",children:[]},{value:"REPORT",id:"report",children:[]},{value:"BOOKMARK",id:"bookmark",children:[]}]}],u={toc:m};function p(e){var t=e.components,n=(0,r.Z)(e,o);return(0,i.kt)("wrapper",(0,a.Z)({},u,n,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h2",{id:"input-fields"},"Input fields"),(0,i.kt)("h3",{id:"document-no"},"Document No."),(0,i.kt)("p",null,"Unique ID, generated at the first data save. The format and value of the next data in row is taken from the ",(0,i.kt)("a",{parentName:"p",href:"numberdef"},(0,i.kt)("strong",{parentName:"a"},"DOCUMENT NUMBERING"))," (code = bank_transfer/cash) data series. "),(0,i.kt)("h3",{id:"reference-no"},"Reference No."),(0,i.kt)("p",null,"Other reference or bank statement number. Optional, its value can be freely defined."),(0,i.kt)("h3",{id:"creation"},"Creation"),(0,i.kt)("p",null,"Date of creation. Automatic value, cannot be changed."),(0,i.kt)("h3",{id:"closed"},"Closed"),(0,i.kt)("p",null,(0,i.kt)("em",{parentName:"p"},"Technical")," closing of the document."),(0,i.kt)("div",{className:"admonition admonition-caution alert alert--warning"},(0,i.kt)("div",{parentName:"div",className:"admonition-heading"},(0,i.kt)("h5",{parentName:"div"},(0,i.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,i.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"16",height:"16",viewBox:"0 0 16 16"},(0,i.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M8.893 1.5c-.183-.31-.52-.5-.887-.5s-.703.19-.886.5L.138 13.499a.98.98 0 0 0 0 1.001c.193.31.53.501.886.501h13.964c.367 0 .704-.19.877-.5a1.03 1.03 0 0 0 .01-1.002L8.893 1.5zm.133 11.497H6.987v-2.003h2.039v2.003zm0-3.004H6.987V5.987h2.039v4.006z"}))),"caution")),(0,i.kt)("div",{parentName:"div",className:"admonition-content"},(0,i.kt)("p",{parentName:"div"},"If set, document data become read only. ",(0,i.kt)("strong",{parentName:"p"},"Marking is not revocable on the user interface!")))),(0,i.kt)("h3",{id:"state"},"State"),(0,i.kt)("p",null,"Its value and editing possibility is linked to ",(0,i.kt)("a",{parentName:"p",href:"usergroup#supervisor"},(0,i.kt)("strong",{parentName:"a"},"ACCESS RIGHTS"))," setting. Not used in current version."),(0,i.kt)("h3",{id:"comment"},"Comment"),(0,i.kt)("p",null,"Remarks field"),(0,i.kt)("h3",{id:"internal-notes"},"Internal notes"),(0,i.kt)("p",null,"Internal comments. Text defined in this field will not appear on the document"),(0,i.kt)("hr",null),(0,i.kt)("h3",{id:"bank"},"BANK"),(0,i.kt)("h3",{id:"account-date"},"Account Date"),(0,i.kt)("p",null,"The date of the statement. Field filling is compulsory."),(0,i.kt)("h3",{id:"bank-account"},"Bank Account"),(0,i.kt)("p",null,"A bank account can be chosen from the search field from ",(0,i.kt)("em",{parentName:"p"},"bank")," type items of ",(0,i.kt)("a",{parentName:"p",href:"place#type"},(0,i.kt)("strong",{parentName:"a"},"PLACE")),"."),(0,i.kt)("hr",null),(0,i.kt)("h3",{id:"petty-cash"},"PETTY CASH"),(0,i.kt)("h3",{id:"direction"},"Direction"),(0,i.kt)("p",null,"OUT, IN. Its value cannot be changed after first save!"),(0,i.kt)("h3",{id:"payment-date"},"Payment Date"),(0,i.kt)("p",null,"The date of the transaction. Field filling is compulsory."),(0,i.kt)("h3",{id:"amount"},"Amount"),(0,i.kt)("p",null,"The amount of inpayment or outpayment. Regardless of direction always with positive sign!"),(0,i.kt)("h3",{id:"petty-cash-1"},"Petty cash"),(0,i.kt)("p",null,"A petty cash can be chosen from the search field from ",(0,i.kt)("em",{parentName:"p"},"cash")," type items of ",(0,i.kt)("a",{parentName:"p",href:"place#type"},(0,i.kt)("strong",{parentName:"a"},"PLACE")),"."),(0,i.kt)("h3",{id:"employee-no"},"Employee No."),(0,i.kt)("p",null,"One of the items of ",(0,i.kt)("a",{parentName:"p",href:"employee#employee-no."},(0,i.kt)("strong",{parentName:"a"},"EMPLOYEE")),". Optional."),(0,i.kt)("h2",{id:"related-data"},"Related data"),(0,i.kt)("h3",{id:"metadata"},(0,i.kt)("a",{parentName:"h3",href:"metadata"},(0,i.kt)("strong",{parentName:"a"},"METADATA"))),(0,i.kt)("p",null,"Unlimited number of supplementary data can be added."),(0,i.kt)("h3",{id:"report-notes"},(0,i.kt)("a",{parentName:"h3",href:"notes"},(0,i.kt)("strong",{parentName:"a"},"REPORT NOTES"))),(0,i.kt)("p",null,"Editable remarks, data for reports."),(0,i.kt)("h3",{id:"invoice"},"INVOICE"),(0,i.kt)("p",null,"Invoices which are linked to bank statement items and to cash receipts."),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("em",{parentName:"li"},"Invoice No."),": The ID of the ",(0,i.kt)("a",{parentName:"li",href:"document"},(0,i.kt)("strong",{parentName:"a"},"INVOICE")),"."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("em",{parentName:"li"},"Currency"),": The currency of related amount. Cannot be changed, set by the program based on currency of Invoice No."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("em",{parentName:"li"},"Amount"),": The payment amount. Can be lower than the total sum of the invoice (in case of payment in installments), and can be higher as well, and also can be an amount of opposite sign (net transaction)."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("em",{parentName:"li"},"Rate"),": Exchange rate in case the payment currency differs from the invoice currency. Default value: 1.")),(0,i.kt)("h3",{id:"document-item"},"DOCUMENT ITEM"),(0,i.kt)("p",null,"BANK. The line items of the statement."),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("em",{parentName:"li"},"Row ID"),": Unique ID of the row item. Automatic value, cannot be changed."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("em",{parentName:"li"},"Payment Date"),": Accounting date of the amount. Field filling is compulsory."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("em",{parentName:"li"},"Amount"),": Amount with sign. Amounts with positive sign indicate the incoming, with negative signs the outgoing items."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("em",{parentName:"li"},"Description"),": Other data of the item.")),(0,i.kt)("h2",{id:"operations"},"Operations"),(0,i.kt)("h3",{id:"copy-from"},"COPY FROM"),(0,i.kt)("p",null,"Create a new, ",(0,i.kt)("em",{parentName:"p"},"same transaction type")," document on the basis of current document's data. "),(0,i.kt)("p",null,"The dates and information related to creation gets updated and the references from the original document will not be transferred either."),(0,i.kt)("h3",{id:"report"},"REPORT"),(0,i.kt)("p",null,(0,i.kt)("a",{parentName:"p",href:"export"},(0,i.kt)("strong",{parentName:"a"},"DATA EXPORT"))),(0,i.kt)("h3",{id:"bookmark"},"BOOKMARK"),(0,i.kt)("p",null,"Set a bookmark for the record. Later can be loaded from bookmarks at any time."))}p.isMDXComponent=!0}}]);