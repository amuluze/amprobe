import{o as e,I as t,r as n,d as o,H as s,B as r,u as a,j as u,N as c,q as d}from"./index.C-XteKAQ.js";import{c as i,_ as l,v as f}from"./el-button.wSr2vfOs.js";import{E as p}from"./message.bOUkP4cz.js";function v(e){return void 0===e}let m=[];const E=e=>{const t=e;t.key===p.esc&&m.forEach((e=>e(t)))},y="focus-trap.focus-after-trapped",w="focus-trap.focus-after-released",L={cancelable:!0,bubbles:!1},h={cancelable:!0,bubbles:!1},T="focusAfterTrapped",b="focusAfterReleased",g=Symbol("elFocusTrap"),F=n(),R=n(0),k=n(0);let N=0;const K=e=>{const t=[],n=document.createTreeWalker(e,NodeFilter.SHOW_ELEMENT,{acceptNode:e=>{const t="INPUT"===e.tagName&&"hidden"===e.type;return e.disabled||e.hidden||t?NodeFilter.FILTER_SKIP:e.tabIndex>=0||e===document.activeElement?NodeFilter.FILTER_ACCEPT:NodeFilter.FILTER_SKIP}});for(;n.nextNode();)t.push(n.currentNode);return t},P=(e,t)=>{for(const n of e)if(!S(n,t))return n},S=(e,t)=>{if("hidden"===getComputedStyle(e).visibility)return!0;for(;e;){if(t&&e===t)return!1;if("none"===getComputedStyle(e).display)return!0;e=e.parentElement}return!1},I=(e,t)=>{if(e&&e.focus){const n=document.activeElement;e.focus({preventScroll:!0}),k.value=window.performance.now(),e!==n&&(e=>e instanceof HTMLInputElement&&"select"in e)(e)&&t&&e.select()}};function _(e,t){const n=[...e],o=e.indexOf(t);return-1!==o&&n.splice(o,1),n}const j=(()=>{let e=[];return{push:t=>{const n=e[0];n&&t!==n&&n.pause(),e=_(e,t),e.unshift(t)},remove:t=>{var n,o;e=_(e,t),null==(o=null==(n=e[0])?void 0:n.resume)||o.call(n)}}})(),C=()=>{F.value="pointer",R.value=window.performance.now()},x=()=>{F.value="keyboard",R.value=window.performance.now()},A=e=>new CustomEvent("focus-trap.focusout-prevented",{...h,detail:e});var O=l(o({name:"ElFocusTrap",inheritAttrs:!1,props:{loop:Boolean,trapped:Boolean,focusTrapEl:Object,focusStartEl:{type:[Object,String],default:"first"}},emits:[T,b,"focusin","focusout","focusout-prevented","release-requested"],setup(o,{emit:u}){const l=n();let v,h;const{focusReason:S}=(e((()=>{0===N&&(document.addEventListener("mousedown",C),document.addEventListener("touchstart",C),document.addEventListener("keydown",x)),N++})),t((()=>{N--,N<=0&&(document.removeEventListener("mousedown",C),document.removeEventListener("touchstart",C),document.removeEventListener("keydown",x))})),{focusReason:F,lastUserFocusTimestamp:R,lastAutomatedFocusTimestamp:k});var _;_=e=>{o.trapped&&!O.paused&&u("release-requested",e)},e((()=>{0===m.length&&document.addEventListener("keydown",E),i&&m.push(_)})),t((()=>{m=m.filter((e=>e!==_)),0===m.length&&i&&document.removeEventListener("keydown",E)}));const O={paused:!1,pause(){this.paused=!0},resume(){this.paused=!1}},q=e=>{if(!o.loop&&!o.trapped)return;if(O.paused)return;const{key:t,altKey:n,ctrlKey:s,metaKey:r,currentTarget:a,shiftKey:c}=e,{loop:d}=o,i=t===p.tab&&!n&&!s&&!r,l=document.activeElement;if(i&&l){const t=a,[n,o]=(e=>{const t=K(e);return[P(t,e),P(t.reverse(),e)]})(t);if(n&&o)if(c||l!==o){if(c&&[n,t].includes(l)){const t=A({focusReason:S.value});u("focusout-prevented",t),t.defaultPrevented||(e.preventDefault(),d&&I(o,!0))}}else{const t=A({focusReason:S.value});u("focusout-prevented",t),t.defaultPrevented||(e.preventDefault(),d&&I(n,!0))}else if(l===t){const t=A({focusReason:S.value});u("focusout-prevented",t),t.defaultPrevented||e.preventDefault()}}};s(g,{focusTrapRef:l,onKeydown:q}),r((()=>o.focusTrapEl),(e=>{e&&(l.value=e)}),{immediate:!0}),r([l],(([e],[t])=>{e&&(e.addEventListener("keydown",q),e.addEventListener("focusin",H),e.addEventListener("focusout",M)),t&&(t.removeEventListener("keydown",q),t.removeEventListener("focusin",H),t.removeEventListener("focusout",M))}));const B=e=>{u(T,e)},D=e=>u(b,e),H=e=>{const t=a(l);if(!t)return;const n=e.target,s=e.relatedTarget,r=n&&t.contains(n);if(!o.trapped){s&&t.contains(s)||(v=s)}r&&u("focusin",e),O.paused||o.trapped&&(r?h=n:I(h,!0))},M=e=>{const t=a(l);if(!O.paused&&t)if(o.trapped){const n=e.relatedTarget;f(n)||t.contains(n)||setTimeout((()=>{if(!O.paused&&o.trapped){const e=A({focusReason:S.value});u("focusout-prevented",e),e.defaultPrevented||I(h,!0)}}),0)}else{const n=e.target;n&&t.contains(n)||u("focusout",e)}};async function U(){await c();const e=a(l);if(e){j.push(O);const t=e.contains(document.activeElement)?v:document.activeElement;v=t;if(!e.contains(t)){const n=new Event(y,L);e.addEventListener(y,B),e.dispatchEvent(n),n.defaultPrevented||c((()=>{let n=o.focusStartEl;d(n)||(I(n),document.activeElement!==n&&(n="first")),"first"===n&&((e,t=!1)=>{const n=document.activeElement;for(const o of e)if(I(o,t),document.activeElement!==n)return})(K(e),!0),document.activeElement!==t&&"container"!==n||I(e)}))}}}function W(){const e=a(l);if(e){e.removeEventListener(y,B);const t=new CustomEvent(w,{...L,detail:{focusReason:S.value}});e.addEventListener(w,D),e.dispatchEvent(t),t.defaultPrevented||"keyboard"!=S.value&&R.value>k.value&&!e.contains(document.activeElement)||I(null!=v?v:document.body),e.removeEventListener(w,D),j.remove(O)}}return e((()=>{o.trapped&&U(),r((()=>o.trapped),(e=>{e?U():W()}))})),t((()=>{o.trapped&&W()})),{onKeydown:q}}}),[["render",function(e,t,n,o,s,r){return u(e.$slots,"default",{handleKeydown:e.onKeydown})}],["__file","focus-trap.vue"]]);export{O as E,g as F,v as i};