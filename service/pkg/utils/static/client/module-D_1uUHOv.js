import{D as t}from"./module-Cq_Zev4P.js";import{r as s}from"./main-CidVXVv4.js";import{i,a as e,e as h}from"./module-CbpjHWOZ.js";const o=(t,s)=>{const i=t._$AN;if(void 0===i)return!1;for(const t of i)t._$AO?.(s,!1),o(t,s);return!0},n=t=>{let s,i;do{if(void 0===(s=t._$AM))break;i=s._$AN,i.delete(t),t=s}while(0===i?.size)},c=t=>{for(let s;s=t._$AM;t=s){let i=s._$AN;if(void 0===i)s._$AN=i=new Set;else if(i.has(t))break;i.add(t),l(s)}};function r(t){void 0!==this._$AN?(n(this),this._$AM=t,c(this)):this._$AM=t}function d(t,s=!1,i=0){const e=this._$AH,h=this._$AN;if(void 0!==h&&0!==h.size)if(s)if(Array.isArray(e))for(let t=i;t<e.length;t++)o(e[t],!1),n(e[t]);else null!=e&&(o(e,!1),n(e));else o(this,t)}const l=t=>{t.type==e.CHILD&&(t._$AP??=d,t._$AQ??=r)};class a extends i{constructor(){super(...arguments),this._$AN=void 0}_$AT(t,s,i){super._$AT(t,s,i),c(this),this.isConnected=t._$AU}_$AO(t,s=!0){t!==this.isConnected&&(this.isConnected=t,t?this.reconnected?.():this.disconnected?.()),s&&(o(this,t),n(this))}setValue(t){if(s(this.t))this.t._$AI(t,this);else{const s=[...this.t._$AH];s[this.i]=t,this.t._$AI(s,this,0)}}disconnected(){}reconnected(){}}const A=()=>new f;class f{}const $=new WeakMap,_=h(class extends a{render(s){return t}update(s,[i]){const e=i!==this.Y;return e&&void 0!==this.Y&&this.rt(void 0),(e||this.lt!==this.ct)&&(this.Y=i,this.ht=s.options?.host,this.rt(this.ct=s.element)),t}rt(t){if(this.isConnected||(t=void 0),"function"==typeof this.Y){const s=this.ht??globalThis;let i=$.get(s);void 0===i&&(i=new WeakMap,$.set(s,i)),void 0!==i.get(this.Y)&&this.Y.call(this.ht,void 0),i.set(this.Y,t),void 0!==t&&this.Y.call(this.ht,t)}else this.Y.value=t}get lt(){return"function"==typeof this.Y?$.get(this.ht??globalThis)?.get(this.Y):this.Y?.value}disconnected(){this.lt===this.ct&&this.rt(void 0)}reconnected(){this.rt(this.ct)}});export{_ as K,A as i};
