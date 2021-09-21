(self.webpackChunknervatura_docs=self.webpackChunknervatura_docs||[]).push([[2847],{3905:function(e,t,a){"use strict";a.d(t,{Zo:function(){return d},kt:function(){return m}});var n=a(7294);function r(e,t,a){return t in e?Object.defineProperty(e,t,{value:a,enumerable:!0,configurable:!0,writable:!0}):e[t]=a,e}function o(e,t){var a=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),a.push.apply(a,n)}return a}function l(e){for(var t=1;t<arguments.length;t++){var a=null!=arguments[t]?arguments[t]:{};t%2?o(Object(a),!0).forEach((function(t){r(e,t,a[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(a)):o(Object(a)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(a,t))}))}return e}function i(e,t){if(null==e)return{};var a,n,r=function(e,t){if(null==e)return{};var a,n,r={},o=Object.keys(e);for(n=0;n<o.length;n++)a=o[n],t.indexOf(a)>=0||(r[a]=e[a]);return r}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)a=o[n],t.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(e,a)&&(r[a]=e[a])}return r}var s=n.createContext({}),c=function(e){var t=n.useContext(s),a=t;return e&&(a="function"==typeof e?e(t):l(l({},t),e)),a},d=function(e){var t=c(e.components);return n.createElement(s.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},u=n.forwardRef((function(e,t){var a=e.components,r=e.mdxType,o=e.originalType,s=e.parentName,d=i(e,["components","mdxType","originalType","parentName"]),u=c(a),m=r,h=u["".concat(s,".").concat(m)]||u[m]||p[m]||o;return a?n.createElement(h,l(l({ref:t},d),{},{components:a})):n.createElement(h,l({ref:t},d))}));function m(e,t){var a=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var o=a.length,l=new Array(o);l[0]=u;var i={};for(var s in t)hasOwnProperty.call(t,s)&&(i[s]=t[s]);i.originalType=e,i.mdxType="string"==typeof e?e:r,l[1]=i;for(var c=2;c<o;c++)l[c]=a[c];return n.createElement.apply(null,l)}return n.createElement.apply(null,a)}u.displayName="MDXCreateElement"},6103:function(e,t,a){"use strict";a.r(t),a.d(t,{frontMatter:function(){return s},contentTitle:function(){return c},metadata:function(){return d},toc:function(){return p},default:function(){return m}});var n=a(2122),r=a(9756),o=(a(7294),a(3905)),l=a(4996),i=["components"],s={id:"model",title:"Nervatura Object Model",hide_table_of_contents:!1},c="Nervatura Object Model",d={type:"mdx",permalink:"/nervatura/model",source:"@site/src/pages/model.md"},p=[{value:"Overview",id:"overview",children:[]},{value:"Objects",id:"objects",children:[]},{value:"Base objects",id:"base-objects",children:[]},{value:"Metadata",id:"metadata",children:[]},{value:"Events",id:"events",children:[]},{value:"Transaction",id:"transaction",children:[]},{value:"Relations",id:"relations",children:[]},{value:"Group settings",id:"group-settings",children:[]},{value:"Complex data types",id:"complex-data-types",children:[]},{value:"Other objects",id:"other-objects",children:[]},{value:"User interface objects",id:"user-interface-objects",children:[]},{value:"Relations pyramid",id:"relations-pyramid",children:[]},{value:"Objects relations",id:"objects-relations",children:[]}],u={toc:p};function m(e){var t=e.components,a=(0,r.Z)(e,i);return(0,o.kt)("wrapper",(0,n.Z)({},u,a,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"nervatura-object-model"},"Nervatura Object Model"),(0,o.kt)("h2",{id:"overview"},"Overview"),(0,o.kt)("p",null,"It is a general ",(0,o.kt)("strong",{parentName:"p"},"open-data model"),", which can store all information generated in the operation of a usual corporation. This covers all manufacturer, retailer and service companies (or governmental units) where the business operation can be defined and described within a ",(0,o.kt)("strong",{parentName:"p"},"GOODS")," (items, services to be sold, provided) \u2013 ",(0,o.kt)("strong",{parentName:"p"},"CLIENT")," (the recipient of goods) - ",(0,o.kt)("strong",{parentName:"p"},"RESOURCE")," (assets used to produce the goods) triangle."),(0,o.kt)("div",{className:"admonition admonition-note alert alert--secondary"},(0,o.kt)("div",{parentName:"div",className:"admonition-heading"},(0,o.kt)("h5",{parentName:"div"},(0,o.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,o.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"},(0,o.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"}))),"note")),(0,o.kt)("div",{parentName:"div",className:"admonition-content"},(0,o.kt)("p",{parentName:"div"},"The open-data principle regulates the access to our business data. The point is the logic of the data storage. It means that the data are defined for storage so, that those are compliant with an open data-model which could be accessed and interpreted by anyone. It doesn't concern the physical storage of the data, that can be implemented according to one's needs. However it should ensure that data can be managed safely according to published description. Retrieving, new data creation, possibility to export the entire data structure should be provided."),(0,o.kt)("p",{parentName:"div"},(0,o.kt)("em",{parentName:"p"},"What are the main advantages of open-data applications?")),(0,o.kt)("ul",{parentName:"div"},(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},"safety"),": provided one's have proper usage rights and physical access to a database, then will be able to interpret and process the data correctly without any help or permission from a third party. Information in the data becomes independent of any management system, its treatment is not tied to specific programs or technologies."),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},"efficiency and cost reduction"),": the business management system can be developed and diversified in accordance with one's needs. There is no need and pressure to be tied to a solution of any vendor, the most appropriate tools and programs can be selected for all tasks. The only criteria is that selected applications should be able to communicate and exchange data with each other or with a central database according to the open-data description. The elements of the system can at any time be flexibly developed or new ones added by choosing the best offers available on the market.")))),(0,o.kt)("p",null,"It is located between the application surfaces that are using and creating the data and the real data storage layer. It defines logical objects; data is stored in these freely defined attributes and in relations between them. Its flexible structure allows defining new properties or assigning events to our objects."),(0,o.kt)("p",null,"The number of objects is minimal, their structure is simple. It has an easy to learn, clear and straightforward logic. However it is capable to store the required data in structures. It ensures the possibility to attach defined type metadata of any kind to each object and also makes the objects linkable to each other arbitrarily."),(0,o.kt)("p",null,"The data model is independent from data storage layer. The data storage can be implemented in any way or with any device but as a main requirement the user of the data model must not sense this at all."),(0,o.kt)("h2",{id:"objects"},"Objects"),(0,o.kt)("p",null,"Such ",(0,o.kt)("strong",{parentName:"p"},(0,o.kt)("em",{parentName:"strong"},"pre-defined functional roles"))," which can have any type of attributs, events can be attached to them as well as their elements can be attached to elements of other objects."),(0,o.kt)("img",{alt:"Nervatura Object Model",style:{width:325,height:259},src:(0,l.Z)("img/nom.svg")}),(0,o.kt)("h2",{id:"base-objects"},"Base objects"),(0,o.kt)("p",null,(0,o.kt)("strong",{parentName:"p"},"CUSTOMER")," - all external partners of the company, including the buyer, consumer and supplier side"),(0,o.kt)("p",null,(0,o.kt)("strong",{parentName:"p"},"PRODUCT")," - all raw materials, semi-finished and end-products that are related to our activity (as customer or vendor), produced by us as a manufacturer or offered as service"),(0,o.kt)("p",null,(0,o.kt)("strong",{parentName:"p"},"TOOL, EMPLOYEE, PLACE")," resources which are available for executing the activity and they contribute to it. These can be human resources (EMPLOYEE), material devices, tools, machines (TOOL) or financial, potentially infrastructural conditions such as warehouses, bank account, petty cash (PLACE)"),(0,o.kt)("h2",{id:"metadata"},"Metadata"),(0,o.kt)("p",null,"All data that ",(0,o.kt)("strong",{parentName:"p"},"describe a given object"),", we want to attach to it as information. Some of them are pre-defined but further ones can freely be defined for any of the objects."),(0,o.kt)("p",null,"By using the ",(0,o.kt)("strong",{parentName:"p"},"DEFFIELD")," object we can define data storage metadata for other objects. Besides the classical data types (bool, integer, float, date, time, string, notes) these can contain list of values (valuelist), url links (urlink) or references to concrete items of other objects (customer, product, tool, employee, etc.)."),(0,o.kt)("p",null,"Through the ",(0,o.kt)("strong",{parentName:"p"},"FIELDVALUE")," object every defined feature of the elements of all objects can be queried."),(0,o.kt)("h2",{id:"events"},"Events"),(0,o.kt)("p",null,(0,o.kt)("strong",{parentName:"p"},"Extended object metadata, usually connected to a time or an interval.")," With the help of events we can make the static metadata of an object into dynamic so the feature of a given component is able to take various values at different times. An event can also be valid for a period of time, so having a start and an end date. Optional number of supplementary data of a given object can be attached to it, and it can be grouped as well."),(0,o.kt)("p",null,"We can manage the events through the ",(0,o.kt)("strong",{parentName:"p"},"EVENT")," object. Beside the base object we can also assign events to projects."),(0,o.kt)("h2",{id:"transaction"},"Transaction"),(0,o.kt)("p",null,"Transaction ",(0,o.kt)("strong",{parentName:"p"},"is such a sort of event to which at least two base objects are joined"),"."),(0,o.kt)("p",null,"An event is always attached to a given object. As a further event feature another base object can be specified but it's just an optional additional data in this case.",(0,o.kt)("br",null),"\nIn the transactions the relation between the base objects is an indispensable and essential component of the given event. The transaction doesn\u2019t belong to any of the base objects but the base objects are joined to a transaction. From these some base objects might be optional components but at least two should be indispensable part of it."),(0,o.kt)("p",null,"The most common object pair is the customer and product relationship (e.g.: offer, order, invoice) but any other combination is also possible, for example product-place (stock management), customer-tool (rental), employee-customer (worksheet) etc."),(0,o.kt)("p",null,'We can link additional data to transactions just as to events, but in contrary to events, here we don\u2019t use the features of the linked base object but we can declare own metadata. Transactions can also be linked to each other or can "originate" from each other, for instance offer -> order -> inventory move -> invoice.'),(0,o.kt)("p",null,"The object of transactions is the ",(0,o.kt)("strong",{parentName:"p"},"TRANS")," which contains the main data of transactions as well as the single object relations. ",(0,o.kt)("strong",{parentName:"p"},"ITEM")," object contains PRODUCT lines linked to transactions, ",(0,o.kt)("strong",{parentName:"p"},"PAYMENT")," object contains financial settlements, ",(0,o.kt)("strong",{parentName:"p"},"MOVEMENT")," object contains warehouse and tool movements."),(0,o.kt)("h2",{id:"relations"},"Relations"),(0,o.kt)("p",null,"There are several possibilities to link single objects. Usually the object has the possibility of applying the one to one relation by default, if it is required so by its type. In case of need the additional data pointing to the proper type of object can also be generated at any time.",(0,o.kt)("br",null),"\nFor example we can set a customer type feature for CUSTOMER object wherewith we can link a given customer to another customer. With the same method one to many relations can also be set, so in this case we can also link our customer to some other customers."),(0,o.kt)("p",null,"If any linked customer is also linked to an other customer it results in a many to many relation.",(0,o.kt)("br",null),"\n",(0,o.kt)("strong",{parentName:"p"},"LINK")," object can also be used to set relations to objects. This way two objects can be linked without setting further object features."),(0,o.kt)("h2",{id:"group-settings"},"Group settings"),(0,o.kt)("p",null,"Several options are available for grouping the objects. Using supplementary data, further to data storage opportunities allows also grouping to a certain degree."),(0,o.kt)("p",null,"In the ",(0,o.kt)("strong",{parentName:"p"},"GROUPS")," object we can create groups by object types. If needed, further features can be defined for these groups. These can then be used for assignments of pre-defined values on a given object (for example type options), but through LINK object can also be used for creating classical one to many groups (for example customer or product groups)."),(0,o.kt)("p",null,"Actually the ",(0,o.kt)("strong",{parentName:"p"},"PROJECT")," object can be interpreted as the extension of GROUPS object. Surely it is possible to set metadata here as well but at PROJECT object time related extension is also possible, just like it is in case of events vs. metadata. Optionally it can have start or end date, we can also link it to customers or places. Projects can also have their own events as well as any transaction can be linked to them."),(0,o.kt)("h2",{id:"complex-data-types"},"Complex data types"),(0,o.kt)("p",null,"When adding features to objects in some cases complex data feature setting is needed. Essentially ",(0,o.kt)("strong",{parentName:"p"},"these are such sub objects which possess own features"),". For example if we want to add address data to a customer then by setting the address we can give the city, the zip code or the name of the street as well.",(0,o.kt)("br",null),"\nSome complex data types can be linked not only to a single object but the same element can also be attached to many others. In some of them it is possible to define further metadata."),(0,o.kt)("p",null,"One to many linked sub objects:",(0,o.kt)("br",null),"\n",(0,o.kt)("strong",{parentName:"p"},"ADDRESS, BARCODE, CONTACT, PRICE")),(0,o.kt)("p",null,"One to one linked sub objects:",(0,o.kt)("br",null),"\n",(0,o.kt)("strong",{parentName:"p"},"CURRENCY, PATTERN, RATE, TAX")),(0,o.kt)("h2",{id:"other-objects"},"Other objects"),(0,o.kt)("p",null,"The objects of rights management, logging and other options:",(0,o.kt)("br",null),"\n",(0,o.kt)("strong",{parentName:"p"},"LOG, NUMBERDEF")),(0,o.kt)("h2",{id:"user-interface-objects"},"User interface objects"),(0,o.kt)("p",null,"These objects ",(0,o.kt)("strong",{parentName:"p"},"are not part of the object model"),", they are not needed for recording the workflow data. However certain applications to ensure their own operation might require data storage possibilities."),(0,o.kt)("p",null,"Storage of data of Reports:",(0,o.kt)("br",null),"\n",(0,o.kt)("strong",{parentName:"p"},"UI_REPORT")),(0,o.kt)("p",null,"Settings of user interfaces:",(0,o.kt)("br",null),"\n",(0,o.kt)("strong",{parentName:"p"},"UI_MENU, UI_MENUFIELDS")),(0,o.kt)("p",null,"User rights, settings:",(0,o.kt)("br",null),"\n",(0,o.kt)("strong",{parentName:"p"},"UI_AUDIT, UI_USERCONFIG")),(0,o.kt)("p",null,"Regional settings:",(0,o.kt)("br",null),"\n",(0,o.kt)("strong",{parentName:"p"},"UI_MESSAGE")),(0,o.kt)("p",null,"Printing:",(0,o.kt)("br",null),"\n",(0,o.kt)("strong",{parentName:"p"},"UI_PRINTQUEUE")),(0,o.kt)("h2",{id:"relations-pyramid"},"Relations pyramid"),(0,o.kt)("p",null,(0,o.kt)("em",{parentName:"p"},"For safe data export and import go from top to the bottom.")),(0,o.kt)("table",null,(0,o.kt)("thead",{parentName:"table"},(0,o.kt)("tr",{parentName:"thead"},(0,o.kt)("th",{parentName:"tr",align:null},"Level"),(0,o.kt)("th",{parentName:"tr",align:null}),(0,o.kt)("th",{parentName:"tr",align:null},"Metadata"),(0,o.kt)("th",{parentName:"tr",align:null},"Objects"))),(0,o.kt)("tbody",{parentName:"table"},(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:null},"level 1a"),(0,o.kt)("td",{parentName:"tr",align:null},"no external link"),(0,o.kt)("td",{parentName:"tr",align:null},"no"),(0,o.kt)("td",{parentName:"tr",align:null},"GROUPS, NUMBERDEF")),(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:null},"level 1b"),(0,o.kt)("td",{parentName:"tr",align:null},"no external link"),(0,o.kt)("td",{parentName:"tr",align:null},"yes*"),(0,o.kt)("td",{parentName:"tr",align:null},"CURRENCY, TAX")),(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:null},"level 2a"),(0,o.kt)("td",{parentName:"tr",align:null},"link to <= level 1"),(0,o.kt)("td",{parentName:"tr",align:null},"no"),(0,o.kt)("td",{parentName:"tr",align:null},"DEFFIELD, PATTERN")),(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:null},"level 2b"),(0,o.kt)("td",{parentName:"tr",align:null},"link to <= level 1"),(0,o.kt)("td",{parentName:"tr",align:null},"yes*"),(0,o.kt)("td",{parentName:"tr",align:null},"CUSTOMER, EMPLOYEE, PLACE, PRODUCT")),(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:null},"level 3"),(0,o.kt)("td",{parentName:"tr",align:null},"link to <= level 2"),(0,o.kt)("td",{parentName:"tr",align:null},"yes*"),(0,o.kt)("td",{parentName:"tr",align:null},"BARCODE, PRICE, PROJECT, RATE, TOOL")),(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:null},"level 4"),(0,o.kt)("td",{parentName:"tr",align:null},"link to <= level 3"),(0,o.kt)("td",{parentName:"tr",align:null},"yes*"),(0,o.kt)("td",{parentName:"tr",align:null},"TRANS")),(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:null},"level 5"),(0,o.kt)("td",{parentName:"tr",align:null},"link to <= level 4"),(0,o.kt)("td",{parentName:"tr",align:null},"yes*"),(0,o.kt)("td",{parentName:"tr",align:null},"EVENT, ITEM, MOVEMENT, PAYMENT")),(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:null},"level 6"),(0,o.kt)("td",{parentName:"tr",align:null},"link to <= level 5"),(0,o.kt)("td",{parentName:"tr",align:null},"yes*"),(0,o.kt)("td",{parentName:"tr",align:null},"ADDRESS, CONTACT")),(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:null},"level 7"),(0,o.kt)("td",{parentName:"tr",align:null},"link to <= level 6"),(0,o.kt)("td",{parentName:"tr",align:null},"yes*"),(0,o.kt)("td",{parentName:"tr",align:null},"LINK, LOG")),(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:null},"level 8"),(0,o.kt)("td",{parentName:"tr",align:null},"link to <= level 7"),(0,o.kt)("td",{parentName:"tr",align:null},"no"),(0,o.kt)("td",{parentName:"tr",align:null},"FIELDVALUE")))),(0,o.kt)("p",null,"*",(0,o.kt)("em",{parentName:"p"},"Export with the FIELDVALUE (cross-references fields)")),(0,o.kt)("table",null,(0,o.kt)("thead",{parentName:"table"},(0,o.kt)("tr",{parentName:"thead"},(0,o.kt)("th",{parentName:"tr",align:null},"Level"),(0,o.kt)("th",{parentName:"tr",align:null}),(0,o.kt)("th",{parentName:"tr",align:null},"Objects"))),(0,o.kt)("tbody",{parentName:"table"},(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:null},"level 1"),(0,o.kt)("td",{parentName:"tr",align:null},"no external link"),(0,o.kt)("td",{parentName:"tr",align:null},"UI_MENU")),(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:null},"level 2a"),(0,o.kt)("td",{parentName:"tr",align:null},"link to <= level 1"),(0,o.kt)("td",{parentName:"tr",align:null},"UI_MESSAGE")),(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:null},"level 2b"),(0,o.kt)("td",{parentName:"tr",align:null},"link to <= NOM level 1"),(0,o.kt)("td",{parentName:"tr",align:null},"UI_REPORT, UI_AUDIT")),(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:null},"level 3"),(0,o.kt)("td",{parentName:"tr",align:null},"link to <= level 2b"),(0,o.kt)("td",{parentName:"tr",align:null},"UI_MENUFIELDS")),(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:null},"level 4"),(0,o.kt)("td",{parentName:"tr",align:null},"link to <= NOM level 2"),(0,o.kt)("td",{parentName:"tr",align:null},"UI_USERCONFIG, UI_PRINTQUEUE")))),(0,o.kt)("h2",{id:"objects-relations"},"Objects relations"),(0,o.kt)("p",null,(0,o.kt)("em",{parentName:"p"},"1. picture: Document type (transtype) relations.")),(0,o.kt)("img",{alt:"Document type relations",style:{maxWidth:600},src:(0,l.Z)("img/trans.svg")}),(0,o.kt)("p",null,(0,o.kt)("em",{parentName:"p"},"2. picture: A possible relational database plan of NOM objects.")),(0,o.kt)("img",{alt:"Database plan of NOM objects",style:{maxWidth:800},src:(0,l.Z)("img/nom_rel.svg")}),(0,o.kt)("p",null,(0,o.kt)("em",{parentName:"p"},"3. picture: A possible relational database plan of user interface objects.")),(0,o.kt)("img",{alt:"Database plan of user interface objects",style:{maxWidth:560},src:(0,l.Z)("img/nom_uio.svg")}))}m.isMDXComponent=!0},3919:function(e,t,a){"use strict";function n(e){return!0===/^(\w*:|\/\/)/.test(e)}function r(e){return void 0!==e&&!n(e)}a.d(t,{b:function(){return n},Z:function(){return r}})},4996:function(e,t,a){"use strict";a.d(t,{C:function(){return o},Z:function(){return l}});var n=a(2263),r=a(3919);function o(){var e=(0,n.Z)().siteConfig,t=(e=void 0===e?{}:e).baseUrl,a=void 0===t?"/":t,o=e.url;return{withBaseUrl:function(e,t){return function(e,t,a,n){var o=void 0===n?{}:n,l=o.forcePrependBaseUrl,i=void 0!==l&&l,s=o.absolute,c=void 0!==s&&s;if(!a)return a;if(a.startsWith("#"))return a;if((0,r.b)(a))return a;if(i)return t+a;var d=a.startsWith(t)?a:t+a.replace(/^\//,"");return c?e+d:d}(o,a,e,t)}}}function l(e,t){return void 0===t&&(t={}),(0,o().withBaseUrl)(e,t)}}}]);