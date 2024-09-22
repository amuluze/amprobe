import{E as e}from"./index.1DI8xxX2.js";import{c as o,u as n,b as t,a as l,E as r,O as i,w as a}from"./el-popper.wQZhTUhG.js";import{_ as s,b as d,d as u,i as c,E as p,q as f,j as m,y as v,k as g,z as b,m as w,n as h}from"./el-button.Bh3a0xQF.js";import{u as y,E as I}from"./message.DctIvaHE.js";import{d as E,j as C,r as _,G as F,E as T,o as S,u as R,H as k,c as x,a0 as B,N as O,B as $,O as D,a as M,b as K,w as L,m as N,a1 as P,a2 as G,M as A,k as j,a3 as z,D as H,n as J,i as Y,J as U,f as q,h as V,l as W,F as Q,g as X}from"./index.Dr9aOSLV.js";import{u as Z}from"./use-form-item.CQAx4y3m.js";import{c as ee}from"./castArray.DwF5sPJL.js";import{c as oe}from"./refs.eEe29D1w.js";import{F as ne}from"./focus-trap.ZmI0uCgZ.js";var te=s(E({inheritAttrs:!1}),[["render",function(e,o,n,t,l,r){return C(e.$slots,"default")}],["__file","collection.vue"]]);var le=s(E({name:"ElCollectionItem",inheritAttrs:!1}),[["render",function(e,o,n,t,l,r){return C(e.$slots,"default")}],["__file","collection-item.vue"]]);const re="data-el-collection-item",ie=e=>{const o=`El${e}Collection`,n=`${o}Item`,t=Symbol(o),l=Symbol(n),r={...te,name:o,setup(){const e=_(null),o=new Map;F(t,{itemMap:o,getItems:()=>{const n=R(e);if(!n)return[];const t=Array.from(n.querySelectorAll(`[${re}]`));return[...o.values()].sort(((e,o)=>t.indexOf(e.ref)-t.indexOf(o.ref)))},collectionRef:e})}},i={...le,name:n,setup(e,{attrs:o}){const n=_(null),r=T(t,void 0);F(l,{collectionItemRef:n}),S((()=>{const e=R(n);e&&r.itemMap.set(e,{ref:e,...o})})),k((()=>{const e=R(n);r.itemMap.delete(e)}))}};return{COLLECTION_INJECTION_KEY:t,COLLECTION_ITEM_INJECTION_KEY:l,ElCollection:r,ElCollectionItem:i}},ae=d({style:{type:u([String,Array,Object])},currentTabId:{type:u(String)},defaultCurrentTabId:String,loop:Boolean,dir:{type:String,values:["ltr","rtl"],default:"ltr"},orientation:{type:u(String)},onBlur:Function,onFocus:Function,onMousedown:Function}),{ElCollection:se,ElCollectionItem:de,COLLECTION_INJECTION_KEY:ue,COLLECTION_ITEM_INJECTION_KEY:ce}=ie("RovingFocusGroup"),pe=Symbol("elRovingFocusGroup"),fe=Symbol("elRovingFocusGroupItem"),me={ArrowLeft:"prev",ArrowUp:"prev",ArrowRight:"next",ArrowDown:"next",PageUp:"first",Home:"first",PageDown:"last",End:"last"},ve=e=>{const{activeElement:o}=document;for(const n of e){if(n===o)return;if(n.focus(),o!==document.activeElement)return}},ge="currentTabIdChange",be="rovingFocusGroup.entryFocus",we={bubbles:!1,cancelable:!0},he=E({name:"ElRovingFocusGroupImpl",inheritAttrs:!1,props:ae,emits:[ge,"entryFocus"],setup(e,{emit:n}){var t;const l=_(null!=(t=e.currentTabId||e.defaultCurrentTabId)?t:null),r=_(!1),i=_(!1),a=_(null),{getItems:s}=T(ue,void 0),d=x((()=>[{outline:"none"},e.style])),u=o((o=>{var n;null==(n=e.onMousedown)||n.call(e,o)}),(()=>{i.value=!0})),c=o((o=>{var n;null==(n=e.onFocus)||n.call(e,o)}),(e=>{const o=!R(i),{target:n,currentTarget:t}=e;if(n===t&&o&&!R(r)){const e=new Event(be,we);if(null==t||t.dispatchEvent(e),!e.defaultPrevented){const e=s().filter((e=>e.focusable)),o=[e.find((e=>e.active)),e.find((e=>e.id===R(l))),...e].filter(Boolean).map((e=>e.ref));ve(o)}}i.value=!1})),p=o((o=>{var n;null==(n=e.onBlur)||n.call(e,o)}),(()=>{r.value=!1}));F(pe,{currentTabbedId:B(l),loop:O(e,"loop"),tabIndex:x((()=>R(r)?-1:0)),rovingFocusGroupRef:a,rovingFocusGroupRootStyle:d,orientation:O(e,"orientation"),dir:O(e,"dir"),onItemFocus:e=>{n(ge,e)},onItemShiftTab:()=>{r.value=!0},onBlur:p,onFocus:c,onMousedown:u}),$((()=>e.currentTabId),(e=>{l.value=null!=e?e:null})),y(a,be,((...e)=>{n("entryFocus",...e)}))}});var ye=s(E({name:"ElRovingFocusGroup",components:{ElFocusGroupCollection:se,ElRovingFocusGroupImpl:s(he,[["render",function(e,o,n,t,l,r){return C(e.$slots,"default")}],["__file","roving-focus-group-impl.vue"]])}}),[["render",function(e,o,n,t,l,r){const i=D("el-roving-focus-group-impl"),a=D("el-focus-group-collection");return M(),K(a,null,{default:L((()=>[N(i,P(G(e.$attrs)),{default:L((()=>[C(e.$slots,"default")])),_:3},16)])),_:3})}],["__file","roving-focus-group.vue"]]);var Ie=s(E({components:{ElRovingFocusCollectionItem:de},props:{focusable:{type:Boolean,default:!0},active:{type:Boolean,default:!1}},emits:["mousedown","focus","keydown"],setup(e,{emit:n}){const{currentTabbedId:t,loop:l,onItemFocus:r,onItemShiftTab:i}=T(pe,void 0),{getItems:a}=T(ue,void 0),s=Z(),d=_(null),u=o((e=>{n("mousedown",e)}),(o=>{e.focusable?r(R(s)):o.preventDefault()})),c=o((e=>{n("focus",e)}),(()=>{r(R(s))})),p=o((e=>{n("keydown",e)}),(e=>{const{key:o,shiftKey:n,target:t,currentTarget:r}=e;if(o===I.tab&&n)return void i();if(t!==r)return;const s=((e,o,n)=>{const t=((e,o)=>e)(e.key);return me[t]})(e);if(s){e.preventDefault();let o=a().filter((e=>e.focusable)).map((e=>e.ref));switch(s){case"last":o.reverse();break;case"prev":case"next":{"prev"===s&&o.reverse();const e=o.indexOf(r);o=l.value?(u=e+1,(d=o).map(((e,o)=>d[(o+u)%d.length]))):o.slice(e+1);break}}A((()=>{ve(o)}))}var d,u})),f=x((()=>t.value===R(s)));return F(fe,{rovingFocusGroupItemRef:d,tabIndex:x((()=>R(f)?0:-1)),handleMousedown:u,handleFocus:c,handleKeydown:p}),{id:s,handleKeydown:p,handleFocus:c,handleMousedown:u}}}),[["render",function(e,o,n,t,l,r){const i=D("el-roving-focus-collection-item");return M(),K(i,{id:e.id,focusable:e.focusable,active:e.active},{default:L((()=>[C(e.$slots,"default")])),_:3},8,["id","focusable","active"])}],["__file","roving-focus-item.vue"]]);const Ee=d({trigger:n.trigger,effect:{...t.effect,default:"light"},type:{type:u(String)},placement:{type:u(String),default:"bottom"},popperOptions:{type:u(Object),default:()=>({})},id:String,size:{type:String,default:""},splitButton:Boolean,hideOnClick:{type:Boolean,default:!0},loop:{type:Boolean,default:!0},showTimeout:{type:Number,default:150},hideTimeout:{type:Number,default:150},tabindex:{type:u([Number,String]),default:0},maxHeight:{type:u([Number,String]),default:""},popperClass:{type:String,default:""},disabled:{type:Boolean,default:!1},role:{type:String,default:"menu"},buttonProps:{type:u(Object)},teleported:t.teleported}),Ce=d({command:{type:[Object,String,Number],default:()=>({})},disabled:Boolean,divided:Boolean,textValue:String,icon:{type:c}}),_e=d({onKeydown:{type:u(Function)}}),Fe=[I.down,I.pageDown,I.home],Te=[I.up,I.pageUp,I.end],Se=[...Fe,...Te],{ElCollection:Re,ElCollectionItem:ke,COLLECTION_INJECTION_KEY:xe,COLLECTION_ITEM_INJECTION_KEY:Be}=ie("Dropdown"),Oe=Symbol("elDropdown"),{ButtonGroup:$e}=e;var De=s(E({name:"ElDropdown",components:{ElButton:e,ElButtonGroup:$e,ElScrollbar:l,ElDropdownCollection:Re,ElTooltip:r,ElRovingFocusGroup:ye,ElOnlyChild:i,ElIcon:p,ArrowDown:f},props:Ee,emits:["visible-change","click","command"],setup(e,{emit:o}){const n=U(),t=m("dropdown"),{t:l}=v(),r=_(),i=_(),a=_(null),s=_(null),d=_(null),u=_(null),c=_(!1),p=[I.enter,I.space,I.down],f=x((()=>({maxHeight:g(e.maxHeight)}))),w=x((()=>[t.m(T.value)])),h=x((()=>ee(e.trigger))),y=Z().value,E=x((()=>e.id||y));function C(){var e;null==(e=a.value)||e.onClose()}$([r,h],(([e,o],[n])=>{var t,l,r;(null==(t=null==n?void 0:n.$el)?void 0:t.removeEventListener)&&n.$el.removeEventListener("pointerenter",S),(null==(l=null==e?void 0:e.$el)?void 0:l.removeEventListener)&&e.$el.removeEventListener("pointerenter",S),(null==(r=null==e?void 0:e.$el)?void 0:r.addEventListener)&&o.includes("hover")&&e.$el.addEventListener("pointerenter",S)}),{immediate:!0}),k((()=>{var e,o;(null==(o=null==(e=r.value)?void 0:e.$el)?void 0:o.removeEventListener)&&r.value.$el.removeEventListener("pointerenter",S)}));const T=b();function S(){var e,o;null==(o=null==(e=r.value)?void 0:e.$el)||o.focus()}F(Oe,{contentRef:s,role:x((()=>e.role)),triggerId:E,isUsingKeyboard:c,onItemEnter:function(){},onItemLeave:function(){const e=R(s);h.value.includes("hover")&&(null==e||e.focus()),u.value=null}}),F("elDropdown",{instance:n,dropdownSize:T,handleClick:function(){C()},commandHandler:function(...e){o("command",...e)},trigger:O(e,"trigger"),hideOnClick:O(e,"hideOnClick")});return{t:l,ns:t,scrollbar:d,wrapStyle:f,dropdownTriggerKls:w,dropdownSize:T,triggerId:E,triggerKeys:p,currentTabId:u,handleCurrentTabIdChange:function(e){u.value=e},handlerMainButtonClick:e=>{o("click",e)},handleEntryFocus:function(e){c.value||(e.preventDefault(),e.stopImmediatePropagation())},handleClose:C,handleOpen:function(){var e;null==(e=a.value)||e.onOpen()},handleBeforeShowTooltip:function(){o("visible-change",!0)},handleShowTooltip:function(e){"keydown"===(null==e?void 0:e.type)&&s.value.focus()},handleBeforeHideTooltip:function(){o("visible-change",!1)},onFocusAfterTrapped:e=>{var o,n;e.preventDefault(),null==(n=null==(o=s.value)?void 0:o.focus)||n.call(o,{preventScroll:!0})},popperRef:a,contentRef:s,triggeringElementRef:r,referenceElementRef:i}}}),[["render",function(e,o,n,t,l,r){var i;const a=D("el-dropdown-collection"),s=D("el-roving-focus-group"),d=D("el-scrollbar"),u=D("el-only-child"),c=D("el-tooltip"),p=D("el-button"),f=D("arrow-down"),m=D("el-icon"),v=D("el-button-group");return M(),j("div",{class:J([e.ns.b(),e.ns.is("disabled",e.disabled)])},[N(c,{ref:"popperRef",role:e.role,effect:e.effect,"fallback-placements":["bottom","top"],"popper-options":e.popperOptions,"gpu-acceleration":!1,"hide-after":"hover"===e.trigger?e.hideTimeout:0,"manual-mode":!0,placement:e.placement,"popper-class":[e.ns.e("popper"),e.popperClass],"reference-element":null==(i=e.referenceElementRef)?void 0:i.$el,trigger:e.trigger,"trigger-keys":e.triggerKeys,"trigger-target-el":e.contentRef,"show-after":"hover"===e.trigger?e.showTimeout:0,"stop-popper-mouse-event":!1,"virtual-ref":e.triggeringElementRef,"virtual-triggering":e.splitButton,disabled:e.disabled,transition:`${e.ns.namespace.value}-zoom-in-top`,teleported:e.teleported,pure:"",persistent:"",onBeforeShow:e.handleBeforeShowTooltip,onShow:e.handleShowTooltip,onBeforeHide:e.handleBeforeHideTooltip},z({content:L((()=>[N(d,{ref:"scrollbar","wrap-style":e.wrapStyle,tag:"div","view-class":e.ns.e("list")},{default:L((()=>[N(s,{loop:e.loop,"current-tab-id":e.currentTabId,orientation:"horizontal",onCurrentTabIdChange:e.handleCurrentTabIdChange,onEntryFocus:e.handleEntryFocus},{default:L((()=>[N(a,null,{default:L((()=>[C(e.$slots,"dropdown")])),_:3})])),_:3},8,["loop","current-tab-id","onCurrentTabIdChange","onEntryFocus"])])),_:3},8,["wrap-style","view-class"])])),_:2},[e.splitButton?void 0:{name:"default",fn:L((()=>[N(u,{id:e.triggerId,ref:"triggeringElementRef",role:"button",tabindex:e.tabindex},{default:L((()=>[C(e.$slots,"default")])),_:3},8,["id","tabindex"])]))}]),1032,["role","effect","popper-options","hide-after","placement","popper-class","reference-element","trigger","trigger-keys","trigger-target-el","show-after","virtual-ref","virtual-triggering","disabled","transition","teleported","onBeforeShow","onShow","onBeforeHide"]),e.splitButton?(M(),K(v,{key:0},{default:L((()=>[N(p,H({ref:"referenceElementRef"},e.buttonProps,{size:e.dropdownSize,type:e.type,disabled:e.disabled,tabindex:e.tabindex,onClick:e.handlerMainButtonClick}),{default:L((()=>[C(e.$slots,"default")])),_:3},16,["size","type","disabled","tabindex","onClick"]),N(p,H({id:e.triggerId,ref:"triggeringElementRef"},e.buttonProps,{role:"button",size:e.dropdownSize,type:e.type,class:e.ns.e("caret-button"),disabled:e.disabled,tabindex:e.tabindex,"aria-label":e.t("el.dropdown.toggleDropdown")}),{default:L((()=>[N(m,{class:J(e.ns.e("icon"))},{default:L((()=>[N(f)])),_:1},8,["class"])])),_:1},16,["id","size","type","class","disabled","tabindex","aria-label"])])),_:3})):Y("v-if",!0)],2)}],["__file","dropdown.vue"]]);const Me=E({name:"DropdownItemImpl",components:{ElIcon:p},props:Ce,emits:["pointermove","pointerleave","click","clickimpl"],setup(e,{emit:n}){const t=m("dropdown"),{role:l}=T(Oe,void 0),{collectionItemRef:r}=T(Be,void 0),{collectionItemRef:i}=T(ce,void 0),{rovingFocusGroupItemRef:a,tabIndex:s,handleFocus:d,handleKeydown:u,handleMousedown:c}=T(fe,void 0),p=oe(r,i,a),f=x((()=>"menu"===l.value?"menuitem":"navigation"===l.value?"link":"button")),v=o((e=>{const{code:o}=e;if(o===I.enter||o===I.space)return e.preventDefault(),e.stopImmediatePropagation(),n("clickimpl",e),!0}),u);return{ns:t,itemRef:p,dataset:{[re]:""},role:f,tabIndex:s,handleFocus:d,handleKeydown:v,handleMousedown:c}}}),Ke=["aria-disabled","tabindex","role"];const Le=()=>{const e=T("elDropdown",{}),o=x((()=>null==e?void 0:e.dropdownSize));return{elDropdown:e,_elDropdownSize:o}};var Ne=s(E({name:"ElDropdownItem",components:{ElDropdownCollectionItem:ke,ElRovingFocusItem:Ie,ElDropdownItemImpl:s(Me,[["render",function(e,o,n,t,l,r){const i=D("el-icon");return M(),j(Q,null,[e.divided?(M(),j("li",H({key:0,role:"separator",class:e.ns.bem("menu","item","divided")},e.$attrs),null,16)):Y("v-if",!0),q("li",H({ref:e.itemRef},{...e.dataset,...e.$attrs},{"aria-disabled":e.disabled,class:[e.ns.be("menu","item"),e.ns.is("disabled",e.disabled)],tabindex:e.tabIndex,role:e.role,onClick:o[0]||(o[0]=o=>e.$emit("clickimpl",o)),onFocus:o[1]||(o[1]=(...o)=>e.handleFocus&&e.handleFocus(...o)),onKeydown:o[2]||(o[2]=W(((...o)=>e.handleKeydown&&e.handleKeydown(...o)),["self"])),onMousedown:o[3]||(o[3]=(...o)=>e.handleMousedown&&e.handleMousedown(...o)),onPointermove:o[4]||(o[4]=o=>e.$emit("pointermove",o)),onPointerleave:o[5]||(o[5]=o=>e.$emit("pointerleave",o))}),[e.icon?(M(),K(i,{key:0},{default:L((()=>[(M(),K(V(e.icon)))])),_:1})):Y("v-if",!0),C(e.$slots,"default")],16,Ke)],64)}],["__file","dropdown-item-impl.vue"]])},inheritAttrs:!1,props:Ce,emits:["pointermove","pointerleave","click"],setup(e,{emit:n,attrs:t}){const{elDropdown:l}=Le(),r=U(),i=_(null),s=x((()=>{var e,o;return null!=(o=null==(e=R(i))?void 0:e.textContent)?o:""})),{onItemEnter:d,onItemLeave:u}=T(Oe,void 0),c=o((e=>(n("pointermove",e),e.defaultPrevented)),a((o=>{if(e.disabled)return void u(o);const n=o.currentTarget;n===document.activeElement||n.contains(document.activeElement)||(d(o),o.defaultPrevented||null==n||n.focus())}))),p=o((e=>(n("pointerleave",e),e.defaultPrevented)),a((e=>{u(e)})));return{handleClick:o((o=>{if(!e.disabled)return n("click",o),"keydown"!==o.type&&o.defaultPrevented}),(o=>{var n,t,i;e.disabled?o.stopImmediatePropagation():((null==(n=null==l?void 0:l.hideOnClick)?void 0:n.value)&&(null==(t=l.handleClick)||t.call(l)),null==(i=l.commandHandler)||i.call(l,e.command,r,o))})),handlePointerMove:c,handlePointerLeave:p,textContent:s,propsAndAttrs:x((()=>({...e,...t})))}}}),[["render",function(e,o,n,t,l,r){var i;const a=D("el-dropdown-item-impl"),s=D("el-roving-focus-item"),d=D("el-dropdown-collection-item");return M(),K(d,{disabled:e.disabled,"text-value":null!=(i=e.textValue)?i:e.textContent},{default:L((()=>[N(s,{focusable:!e.disabled},{default:L((()=>[N(a,H(e.propsAndAttrs,{onPointerleave:e.handlePointerLeave,onPointermove:e.handlePointerMove,onClickimpl:e.handleClick}),{default:L((()=>[C(e.$slots,"default")])),_:3},16,["onPointerleave","onPointermove","onClickimpl"])])),_:3},8,["focusable"])])),_:3},8,["disabled","text-value"])}],["__file","dropdown-item.vue"]]);const Pe=E({name:"ElDropdownMenu",props:_e,setup(e){const n=m("dropdown"),{_elDropdownSize:t}=Le(),l=t.value,{focusTrapRef:r,onKeydown:i}=T(ne,void 0),{contentRef:a,role:s,triggerId:d}=T(Oe,void 0),{collectionRef:u,getItems:c}=T(xe,void 0),{rovingFocusGroupRef:p,rovingFocusGroupRootStyle:f,tabIndex:v,onBlur:g,onFocus:b,onMousedown:w}=T(pe,void 0),{collectionRef:h}=T(ue,void 0),y=x((()=>[n.b("menu"),n.bm("menu",null==l?void 0:l.value)])),E=oe(a,u,r,p,h),C=o((o=>{var n;null==(n=e.onKeydown)||n.call(e,o)}),(e=>{const{currentTarget:o,code:n,target:t}=e;if(o.contains(t),I.tab===n&&e.stopImmediatePropagation(),e.preventDefault(),t!==R(a))return;if(!Se.includes(n))return;const l=c().filter((e=>!e.disabled)).map((e=>e.ref));Te.includes(n)&&l.reverse(),ve(l)}));return{size:l,rovingFocusGroupRootStyle:f,tabIndex:v,dropdownKls:y,role:s,triggerId:d,dropdownListWrapperRef:E,handleKeydown:e=>{C(e),i(e)},onBlur:g,onFocus:b,onMousedown:w}}}),Ge=["role","aria-labelledby"];var Ae=s(Pe,[["render",function(e,o,n,t,l,r){return M(),j("ul",{ref:e.dropdownListWrapperRef,class:J(e.dropdownKls),style:X(e.rovingFocusGroupRootStyle),tabindex:-1,role:e.role,"aria-labelledby":e.triggerId,onBlur:o[0]||(o[0]=(...o)=>e.onBlur&&e.onBlur(...o)),onFocus:o[1]||(o[1]=(...o)=>e.onFocus&&e.onFocus(...o)),onKeydown:o[2]||(o[2]=W(((...o)=>e.handleKeydown&&e.handleKeydown(...o)),["self"])),onMousedown:o[3]||(o[3]=W(((...o)=>e.onMousedown&&e.onMousedown(...o)),["self"]))},[C(e.$slots,"default")],46,Ge)}],["__file","dropdown-menu.vue"]]);const je=w(De,{DropdownItem:Ne,DropdownMenu:Ae}),ze=h(Ne),He=h(Ae);export{ze as E,He as a,je as b};
