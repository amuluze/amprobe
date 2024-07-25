import{o as e,K as t,r as n,d as o,J as s,D as r,u as a,j as u,P as c,q as d}from"./index.uDfNkopd.js";import{f as i,e as l,_ as f,A as p}from"./message.BD1JyyvK.js";function v(e){return void 0===e}let m=[];const E=e=>{const t=e;t.key===l.esc&&m.forEach((e=>e(t)))},y="focus-trap.focus-after-trapped",w="focus-trap.focus-after-released",L={cancelable:!0,bubbles:!1},h={cancelable:!0,bubbles:!1},T="focusAfterTrapped",b="focusAfterReleased",g=Symbol("elFocusTrap"),F=n(),R=n(0),k=n(0);let K=0;const P=e=>{const t=[],n=document.createTreeWalker(e,NodeFilter.SHOW_ELEMENT,{acceptNode:e=>{const t="INPUT"===e.tagName&&"hidden"===e.type;return e.disabled||e.hidden||t?NodeFilter.FILTER_SKIP:e.tabIndex>=0||e===document.activeElement?NodeFilter.FILTER_ACCEPT:NodeFilter.FILTER_SKIP}});for(;n.nextNode();)t.push(n.currentNode);return t},N=(e,t)=>{for(const n of e)if(!S(n,t))return n},S=(e,t)=>{if("hidden"===getComputedStyle(e).visibility)return!0;for(;e;){if(t&&e===t)return!1;if("none"===getComputedStyle(e).display)return!0;e=e.parentElement}return!1},I=(e,t)=>{if(e&&e.focus){const n=document.activeElement;e.focus({preventScroll:!0}),k.value=window.performance.now(),e!==n&&(e=>e instanceof HTMLInputElement&&"select"in e)(e)&&t&&e.select()}};function _(e,t){const n=[...e],o=e.indexOf(t);return-1!==o&&n.splice(o,1),n}const A=(()=>{let e=[];return{push:t=>{const n=e[0];n&&t!==n&&n.pause(),e=_(e,t),e.unshift(t)},remove:t=>{var n,o;e=_(e,t),null==(o=null==(n=e[0])?void 0:n.resume)||o.call(n)}}})(),C=()=>{F.value="pointer",R.value=window.performance.now()},j=()=>{F.value="keyboard",R.value=window.performance.now()},x=e=>new CustomEvent("focus-trap.focusout-prevented",{...h,detail:e});var D=f(o({name:"ElFocusTrap",inheritAttrs:!1,props:{loop:Boolean,trapped:Boolean,focusTrapEl:Object,focusStartEl:{type:[Object,String],default:"first"}},emits:[T,b,"focusin","focusout","focusout-prevented","release-requested"],setup(o,{emit:u}){const f=n();let v,h;const{focusReason:S}=(e((()=>{0===K&&(document.addEventListener("mousedown",C),document.addEventListener("touchstart",C),document.addEventListener("keydown",j)),K++})),t((()=>{K--,K<=0&&(document.removeEventListener("mousedown",C),document.removeEventListener("touchstart",C),document.removeEventListener("keydown",j))})),{focusReason:F,lastUserFocusTimestamp:R,lastAutomatedFocusTimestamp:k});var _;_=e=>{o.trapped&&!D.paused&&u("release-requested",e)},e((()=>{0===m.length&&document.addEventListener("keydown",E),i&&m.push(_)})),t((()=>{m=m.filter((e=>e!==_)),0===m.length&&i&&document.removeEventListener("keydown",E)}));const D={paused:!1,pause(){this.paused=!0},resume(){this.paused=!1}},O=e=>{if(!o.loop&&!o.trapped)return;if(D.paused)return;const{key:t,altKey:n,ctrlKey:s,metaKey:r,currentTarget:a,shiftKey:c}=e,{loop:d}=o,i=t===l.tab&&!n&&!s&&!r,f=document.activeElement;if(i&&f){const t=a,[n,o]=(e=>{const t=P(e);return[N(t,e),N(t.reverse(),e)]})(t);if(n&&o)if(c||f!==o){if(c&&[n,t].includes(f)){const t=x({focusReason:S.value});u("focusout-prevented",t),t.defaultPrevented||(e.preventDefault(),d&&I(o,!0))}}else{const t=x({focusReason:S.value});u("focusout-prevented",t),t.defaultPrevented||(e.preventDefault(),d&&I(n,!0))}else if(f===t){const t=x({focusReason:S.value});u("focusout-prevented",t),t.defaultPrevented||e.preventDefault()}}};s(g,{focusTrapRef:f,onKeydown:O}),r((()=>o.focusTrapEl),(e=>{e&&(f.value=e)}),{immediate:!0}),r([f],(([e],[t])=>{e&&(e.addEventListener("keydown",O),e.addEventListener("focusin",H),e.addEventListener("focusout",M)),t&&(t.removeEventListener("keydown",O),t.removeEventListener("focusin",H),t.removeEventListener("focusout",M))}));const q=e=>{u(T,e)},B=e=>u(b,e),H=e=>{const t=a(f);if(!t)return;const n=e.target,s=e.relatedTarget,r=n&&t.contains(n);if(!o.trapped){s&&t.contains(s)||(v=s)}r&&u("focusin",e),D.paused||o.trapped&&(r?h=n:I(h,!0))},M=e=>{const t=a(f);if(!D.paused&&t)if(o.trapped){const n=e.relatedTarget;p(n)||t.contains(n)||setTimeout((()=>{if(!D.paused&&o.trapped){const e=x({focusReason:S.value});u("focusout-prevented",e),e.defaultPrevented||I(h,!0)}}),0)}else{const n=e.target;n&&t.contains(n)||u("focusout",e)}};async function U(){await c();const e=a(f);if(e){A.push(D);const t=e.contains(document.activeElement)?v:document.activeElement;v=t;if(!e.contains(t)){const n=new Event(y,L);e.addEventListener(y,q),e.dispatchEvent(n),n.defaultPrevented||c((()=>{let n=o.focusStartEl;d(n)||(I(n),document.activeElement!==n&&(n="first")),"first"===n&&((e,t=!1)=>{const n=document.activeElement;for(const o of e)if(I(o,t),document.activeElement!==n)return})(P(e),!0),document.activeElement!==t&&"container"!==n||I(e)}))}}}function W(){const e=a(f);if(e){e.removeEventListener(y,q);const t=new CustomEvent(w,{...L,detail:{focusReason:S.value}});e.addEventListener(w,B),e.dispatchEvent(t),t.defaultPrevented||"keyboard"!=S.value&&R.value>k.value&&!e.contains(document.activeElement)||I(null!=v?v:document.body),e.removeEventListener(w,B),A.remove(D)}}return e((()=>{o.trapped&&U(),r((()=>o.trapped),(e=>{e?U():W()}))})),t((()=>{o.trapped&&W()})),{onKeydown:O}}}),[["render",function(e,t,n,o,s,r){return u(e.$slots,"default",{handleKeydown:e.onKeydown})}],["__file","focus-trap.vue"]]);export{D as E,g as F,v as i};