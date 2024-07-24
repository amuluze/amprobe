import{E as e}from"./index.v0D0o9LE.js";import{c as o,u as n,b as t,a as l,E as r,O as a,w as i}from"./el-popper.DMOJ8Ypx.js";import{_ as s,b as d,d as u,f as c,h as p,E as f,t as v,u as m,c as g,w as b,e as w}from"./base.CwKrOXiL.js";import{h,k as y,u as I}from"./el-overlay.CI2e2OO9.js";import{g as E,d as C,f as _,r as F,p as T,h as S,j as R,u as x,k,c as B,a7 as O,W as $,w as D,X as M,o as K,x as L,v as P,m as N,a8 as G,a9 as A,q as z,e as j,a5 as H,P as J,n as Y,C as U,S as W,D as q,y as V,H as X,F as Q,z as Z}from"./index.DR6loZJs.js";import{u as ee}from"./use-form-item.KseOUpWN.js";import{c as oe}from"./castArray.BhBSoVDq.js";import{F as ne}from"./focus-trap.GHIWC1jm.js";const te=(...e)=>o=>{e.forEach((e=>{E(e)?e(o):e.value=o}))};var le=s(C({inheritAttrs:!1}),[["render",function(e,o,n,t,l,r){return _(e.$slots,"default")}],["__file","collection.vue"]]);var re=s(C({name:"ElCollectionItem",inheritAttrs:!1}),[["render",function(e,o,n,t,l,r){return _(e.$slots,"default")}],["__file","collection-item.vue"]]);const ae="data-el-collection-item",ie=e=>{const o=`El${e}Collection`,n=`${o}Item`,t=Symbol(o),l=Symbol(n),r={...le,name:o,setup(){const e=F(null),o=new Map;T(t,{itemMap:o,getItems:()=>{const n=x(e);if(!n)return[];const t=Array.from(n.querySelectorAll(`[${ae}]`));return[...o.values()].sort(((e,o)=>t.indexOf(e.ref)-t.indexOf(o.ref)))},collectionRef:e})}},a={...re,name:n,setup(e,{attrs:o}){const n=F(null),r=S(t,void 0);T(l,{collectionItemRef:n}),R((()=>{const e=x(n);e&&r.itemMap.set(e,{ref:e,...o})})),k((()=>{const e=x(n);r.itemMap.delete(e)}))}};return{COLLECTION_INJECTION_KEY:t,COLLECTION_ITEM_INJECTION_KEY:l,ElCollection:r,ElCollectionItem:a}},se=d({style:{type:u([String,Array,Object])},currentTabId:{type:u(String)},defaultCurrentTabId:String,loop:Boolean,dir:{type:String,values:["ltr","rtl"],default:"ltr"},orientation:{type:u(String)},onBlur:Function,onFocus:Function,onMousedown:Function}),{ElCollection:de,ElCollectionItem:ue,COLLECTION_INJECTION_KEY:ce,COLLECTION_ITEM_INJECTION_KEY:pe}=ie("RovingFocusGroup"),fe=Symbol("elRovingFocusGroup"),ve=Symbol("elRovingFocusGroupItem"),me={ArrowLeft:"prev",ArrowUp:"prev",ArrowRight:"next",ArrowDown:"next",PageUp:"first",Home:"first",PageDown:"last",End:"last"},ge=e=>{const{activeElement:o}=document;for(const n of e){if(n===o)return;if(n.focus(),o!==document.activeElement)return}},be="currentTabIdChange",we="rovingFocusGroup.entryFocus",he={bubbles:!1,cancelable:!0},ye=C({name:"ElRovingFocusGroupImpl",inheritAttrs:!1,props:se,emits:[be,"entryFocus"],setup(e,{emit:n}){var t;const l=F(null!=(t=e.currentTabId||e.defaultCurrentTabId)?t:null),r=F(!1),a=F(!1),i=F(null),{getItems:s}=S(ce,void 0),d=B((()=>[{outline:"none"},e.style])),u=o((o=>{var n;null==(n=e.onMousedown)||n.call(e,o)}),(()=>{a.value=!0})),p=o((o=>{var n;null==(n=e.onFocus)||n.call(e,o)}),(e=>{const o=!x(a),{target:n,currentTarget:t}=e;if(n===t&&o&&!x(r)){const e=new Event(we,he);if(null==t||t.dispatchEvent(e),!e.defaultPrevented){const e=s().filter((e=>e.focusable)),o=[e.find((e=>e.active)),e.find((e=>e.id===x(l))),...e].filter(Boolean).map((e=>e.ref));ge(o)}}a.value=!1})),f=o((o=>{var n;null==(n=e.onBlur)||n.call(e,o)}),(()=>{r.value=!1}));T(fe,{currentTabbedId:O(l),loop:$(e,"loop"),tabIndex:B((()=>x(r)?-1:0)),rovingFocusGroupRef:i,rovingFocusGroupRootStyle:d,orientation:$(e,"orientation"),dir:$(e,"dir"),onItemFocus:e=>{n(be,e)},onItemShiftTab:()=>{r.value=!0},onBlur:f,onFocus:p,onMousedown:u}),D((()=>e.currentTabId),(e=>{l.value=null!=e?e:null})),c(i,we,((...e)=>{n("entryFocus",...e)}))}});var Ie=s(C({name:"ElRovingFocusGroup",components:{ElFocusGroupCollection:de,ElRovingFocusGroupImpl:s(ye,[["render",function(e,o,n,t,l,r){return _(e.$slots,"default")}],["__file","roving-focus-group-impl.vue"]])}}),[["render",function(e,o,n,t,l,r){const a=M("el-roving-focus-group-impl"),i=M("el-focus-group-collection");return K(),L(i,null,{default:P((()=>[N(a,G(A(e.$attrs)),{default:P((()=>[_(e.$slots,"default")])),_:3},16)])),_:3})}],["__file","roving-focus-group.vue"]]);var Ee=s(C({components:{ElRovingFocusCollectionItem:ue},props:{focusable:{type:Boolean,default:!0},active:{type:Boolean,default:!1}},emits:["mousedown","focus","keydown"],setup(e,{emit:n}){const{currentTabbedId:t,loop:l,onItemFocus:r,onItemShiftTab:a}=S(fe,void 0),{getItems:i}=S(ce,void 0),s=ee(),d=F(null),u=o((e=>{n("mousedown",e)}),(o=>{e.focusable?r(x(s)):o.preventDefault()})),c=o((e=>{n("focus",e)}),(()=>{r(x(s))})),f=o((e=>{n("keydown",e)}),(e=>{const{key:o,shiftKey:n,target:t,currentTarget:r}=e;if(o===p.tab&&n)return void a();if(t!==r)return;const s=((e,o,n)=>{const t=((e,o)=>e)(e.key);return me[t]})(e);if(s){e.preventDefault();let o=i().filter((e=>e.focusable)).map((e=>e.ref));switch(s){case"last":o.reverse();break;case"prev":case"next":{"prev"===s&&o.reverse();const e=o.indexOf(r);o=l.value?(u=e+1,(d=o).map(((e,o)=>d[(o+u)%d.length]))):o.slice(e+1);break}}z((()=>{ge(o)}))}var d,u})),v=B((()=>t.value===x(s)));return T(ve,{rovingFocusGroupItemRef:d,tabIndex:B((()=>x(v)?0:-1)),handleMousedown:u,handleFocus:c,handleKeydown:f}),{id:s,handleKeydown:f,handleFocus:c,handleMousedown:u}}}),[["render",function(e,o,n,t,l,r){const a=M("el-roving-focus-collection-item");return K(),L(a,{id:e.id,focusable:e.focusable,active:e.active},{default:P((()=>[_(e.$slots,"default")])),_:3},8,["id","focusable","active"])}],["__file","roving-focus-item.vue"]]);const Ce=d({trigger:n.trigger,effect:{...t.effect,default:"light"},type:{type:u(String)},placement:{type:u(String),default:"bottom"},popperOptions:{type:u(Object),default:()=>({})},id:String,size:{type:String,default:""},splitButton:Boolean,hideOnClick:{type:Boolean,default:!0},loop:{type:Boolean,default:!0},showTimeout:{type:Number,default:150},hideTimeout:{type:Number,default:150},tabindex:{type:u([Number,String]),default:0},maxHeight:{type:u([Number,String]),default:""},popperClass:{type:String,default:""},disabled:{type:Boolean,default:!1},role:{type:String,default:"menu"},buttonProps:{type:u(Object)},teleported:t.teleported}),_e=d({command:{type:[Object,String,Number],default:()=>({})},disabled:Boolean,divided:Boolean,textValue:String,icon:{type:h}}),Fe=d({onKeydown:{type:u(Function)}}),Te=[p.down,p.pageDown,p.home],Se=[p.up,p.pageUp,p.end],Re=[...Te,...Se],{ElCollection:xe,ElCollectionItem:ke,COLLECTION_INJECTION_KEY:Be,COLLECTION_ITEM_INJECTION_KEY:Oe}=ie("Dropdown"),$e=Symbol("elDropdown"),{ButtonGroup:De}=e;var Me=s(C({name:"ElDropdown",components:{ElButton:e,ElButtonGroup:De,ElScrollbar:l,ElDropdownCollection:xe,ElTooltip:r,ElRovingFocusGroup:Ie,ElOnlyChild:a,ElIcon:f,ArrowDown:v},props:Ce,emits:["visible-change","click","command"],setup(e,{emit:o}){const n=W(),t=m("dropdown"),{t:l}=y(),r=F(),a=F(),i=F(null),s=F(null),d=F(null),u=F(null),c=F(!1),f=[p.enter,p.space,p.down],v=B((()=>({maxHeight:g(e.maxHeight)}))),b=B((()=>[t.m(_.value)])),w=B((()=>oe(e.trigger))),h=ee().value,E=B((()=>e.id||h));function C(){var e;null==(e=i.value)||e.onClose()}D([r,w],(([e,o],[n])=>{var t,l,r;(null==(t=null==n?void 0:n.$el)?void 0:t.removeEventListener)&&n.$el.removeEventListener("pointerenter",S),(null==(l=null==e?void 0:e.$el)?void 0:l.removeEventListener)&&e.$el.removeEventListener("pointerenter",S),(null==(r=null==e?void 0:e.$el)?void 0:r.addEventListener)&&o.includes("hover")&&e.$el.addEventListener("pointerenter",S)}),{immediate:!0}),k((()=>{var e,o;(null==(o=null==(e=r.value)?void 0:e.$el)?void 0:o.removeEventListener)&&r.value.$el.removeEventListener("pointerenter",S)}));const _=I();function S(){var e,o;null==(o=null==(e=r.value)?void 0:e.$el)||o.focus()}T($e,{contentRef:s,role:B((()=>e.role)),triggerId:E,isUsingKeyboard:c,onItemEnter:function(){},onItemLeave:function(){const e=x(s);w.value.includes("hover")&&(null==e||e.focus()),u.value=null}}),T("elDropdown",{instance:n,dropdownSize:_,handleClick:function(){C()},commandHandler:function(...e){o("command",...e)},trigger:$(e,"trigger"),hideOnClick:$(e,"hideOnClick")});return{t:l,ns:t,scrollbar:d,wrapStyle:v,dropdownTriggerKls:b,dropdownSize:_,triggerId:E,triggerKeys:f,currentTabId:u,handleCurrentTabIdChange:function(e){u.value=e},handlerMainButtonClick:e=>{o("click",e)},handleEntryFocus:function(e){c.value||(e.preventDefault(),e.stopImmediatePropagation())},handleClose:C,handleOpen:function(){var e;null==(e=i.value)||e.onOpen()},handleBeforeShowTooltip:function(){o("visible-change",!0)},handleShowTooltip:function(e){"keydown"===(null==e?void 0:e.type)&&s.value.focus()},handleBeforeHideTooltip:function(){o("visible-change",!1)},onFocusAfterTrapped:e=>{var o,n;e.preventDefault(),null==(n=null==(o=s.value)?void 0:o.focus)||n.call(o,{preventScroll:!0})},popperRef:i,contentRef:s,triggeringElementRef:r,referenceElementRef:a}}}),[["render",function(e,o,n,t,l,r){var a;const i=M("el-dropdown-collection"),s=M("el-roving-focus-group"),d=M("el-scrollbar"),u=M("el-only-child"),c=M("el-tooltip"),p=M("el-button"),f=M("arrow-down"),v=M("el-icon"),m=M("el-button-group");return K(),j("div",{class:Y([e.ns.b(),e.ns.is("disabled",e.disabled)])},[N(c,{ref:"popperRef",role:e.role,effect:e.effect,"fallback-placements":["bottom","top"],"popper-options":e.popperOptions,"gpu-acceleration":!1,"hide-after":"hover"===e.trigger?e.hideTimeout:0,"manual-mode":!0,placement:e.placement,"popper-class":[e.ns.e("popper"),e.popperClass],"reference-element":null==(a=e.referenceElementRef)?void 0:a.$el,trigger:e.trigger,"trigger-keys":e.triggerKeys,"trigger-target-el":e.contentRef,"show-after":"hover"===e.trigger?e.showTimeout:0,"stop-popper-mouse-event":!1,"virtual-ref":e.triggeringElementRef,"virtual-triggering":e.splitButton,disabled:e.disabled,transition:`${e.ns.namespace.value}-zoom-in-top`,teleported:e.teleported,pure:"",persistent:"",onBeforeShow:e.handleBeforeShowTooltip,onShow:e.handleShowTooltip,onBeforeHide:e.handleBeforeHideTooltip},H({content:P((()=>[N(d,{ref:"scrollbar","wrap-style":e.wrapStyle,tag:"div","view-class":e.ns.e("list")},{default:P((()=>[N(s,{loop:e.loop,"current-tab-id":e.currentTabId,orientation:"horizontal",onCurrentTabIdChange:e.handleCurrentTabIdChange,onEntryFocus:e.handleEntryFocus},{default:P((()=>[N(i,null,{default:P((()=>[_(e.$slots,"dropdown")])),_:3})])),_:3},8,["loop","current-tab-id","onCurrentTabIdChange","onEntryFocus"])])),_:3},8,["wrap-style","view-class"])])),_:2},[e.splitButton?void 0:{name:"default",fn:P((()=>[N(u,{id:e.triggerId,ref:"triggeringElementRef",role:"button",tabindex:e.tabindex},{default:P((()=>[_(e.$slots,"default")])),_:3},8,["id","tabindex"])]))}]),1032,["role","effect","popper-options","hide-after","placement","popper-class","reference-element","trigger","trigger-keys","trigger-target-el","show-after","virtual-ref","virtual-triggering","disabled","transition","teleported","onBeforeShow","onShow","onBeforeHide"]),e.splitButton?(K(),L(m,{key:0},{default:P((()=>[N(p,J({ref:"referenceElementRef"},e.buttonProps,{size:e.dropdownSize,type:e.type,disabled:e.disabled,tabindex:e.tabindex,onClick:e.handlerMainButtonClick}),{default:P((()=>[_(e.$slots,"default")])),_:3},16,["size","type","disabled","tabindex","onClick"]),N(p,J({id:e.triggerId,ref:"triggeringElementRef"},e.buttonProps,{role:"button",size:e.dropdownSize,type:e.type,class:e.ns.e("caret-button"),disabled:e.disabled,tabindex:e.tabindex,"aria-label":e.t("el.dropdown.toggleDropdown")}),{default:P((()=>[N(v,{class:Y(e.ns.e("icon"))},{default:P((()=>[N(f)])),_:1},8,["class"])])),_:1},16,["id","size","type","class","disabled","tabindex","aria-label"])])),_:3})):U("v-if",!0)],2)}],["__file","dropdown.vue"]]);const Ke=C({name:"DropdownItemImpl",components:{ElIcon:f},props:_e,emits:["pointermove","pointerleave","click","clickimpl"],setup(e,{emit:n}){const t=m("dropdown"),{role:l}=S($e,void 0),{collectionItemRef:r}=S(Oe,void 0),{collectionItemRef:a}=S(pe,void 0),{rovingFocusGroupItemRef:i,tabIndex:s,handleFocus:d,handleKeydown:u,handleMousedown:c}=S(ve,void 0),f=te(r,a,i),v=B((()=>"menu"===l.value?"menuitem":"navigation"===l.value?"link":"button")),g=o((e=>{const{code:o}=e;if(o===p.enter||o===p.space)return e.preventDefault(),e.stopImmediatePropagation(),n("clickimpl",e),!0}),u);return{ns:t,itemRef:f,dataset:{[ae]:""},role:v,tabIndex:s,handleFocus:d,handleKeydown:g,handleMousedown:c}}}),Le=["aria-disabled","tabindex","role"];const Pe=()=>{const e=S("elDropdown",{}),o=B((()=>null==e?void 0:e.dropdownSize));return{elDropdown:e,_elDropdownSize:o}};var Ne=s(C({name:"ElDropdownItem",components:{ElDropdownCollectionItem:ke,ElRovingFocusItem:Ee,ElDropdownItemImpl:s(Ke,[["render",function(e,o,n,t,l,r){const a=M("el-icon");return K(),j(Q,null,[e.divided?(K(),j("li",J({key:0,role:"separator",class:e.ns.bem("menu","item","divided")},e.$attrs),null,16)):U("v-if",!0),q("li",J({ref:e.itemRef},{...e.dataset,...e.$attrs},{"aria-disabled":e.disabled,class:[e.ns.be("menu","item"),e.ns.is("disabled",e.disabled)],tabindex:e.tabIndex,role:e.role,onClick:o[0]||(o[0]=o=>e.$emit("clickimpl",o)),onFocus:o[1]||(o[1]=(...o)=>e.handleFocus&&e.handleFocus(...o)),onKeydown:o[2]||(o[2]=X(((...o)=>e.handleKeydown&&e.handleKeydown(...o)),["self"])),onMousedown:o[3]||(o[3]=(...o)=>e.handleMousedown&&e.handleMousedown(...o)),onPointermove:o[4]||(o[4]=o=>e.$emit("pointermove",o)),onPointerleave:o[5]||(o[5]=o=>e.$emit("pointerleave",o))}),[e.icon?(K(),L(a,{key:0},{default:P((()=>[(K(),L(V(e.icon)))])),_:1})):U("v-if",!0),_(e.$slots,"default")],16,Le)],64)}],["__file","dropdown-item-impl.vue"]])},inheritAttrs:!1,props:_e,emits:["pointermove","pointerleave","click"],setup(e,{emit:n,attrs:t}){const{elDropdown:l}=Pe(),r=W(),a=F(null),s=B((()=>{var e,o;return null!=(o=null==(e=x(a))?void 0:e.textContent)?o:""})),{onItemEnter:d,onItemLeave:u}=S($e,void 0),c=o((e=>(n("pointermove",e),e.defaultPrevented)),i((o=>{if(e.disabled)return void u(o);const n=o.currentTarget;n===document.activeElement||n.contains(document.activeElement)||(d(o),o.defaultPrevented||null==n||n.focus())}))),p=o((e=>(n("pointerleave",e),e.defaultPrevented)),i((e=>{u(e)})));return{handleClick:o((o=>{if(!e.disabled)return n("click",o),"keydown"!==o.type&&o.defaultPrevented}),(o=>{var n,t,a;e.disabled?o.stopImmediatePropagation():((null==(n=null==l?void 0:l.hideOnClick)?void 0:n.value)&&(null==(t=l.handleClick)||t.call(l)),null==(a=l.commandHandler)||a.call(l,e.command,r,o))})),handlePointerMove:c,handlePointerLeave:p,textContent:s,propsAndAttrs:B((()=>({...e,...t})))}}}),[["render",function(e,o,n,t,l,r){var a;const i=M("el-dropdown-item-impl"),s=M("el-roving-focus-item"),d=M("el-dropdown-collection-item");return K(),L(d,{disabled:e.disabled,"text-value":null!=(a=e.textValue)?a:e.textContent},{default:P((()=>[N(s,{focusable:!e.disabled},{default:P((()=>[N(i,J(e.propsAndAttrs,{onPointerleave:e.handlePointerLeave,onPointermove:e.handlePointerMove,onClickimpl:e.handleClick}),{default:P((()=>[_(e.$slots,"default")])),_:3},16,["onPointerleave","onPointermove","onClickimpl"])])),_:3},8,["focusable"])])),_:3},8,["disabled","text-value"])}],["__file","dropdown-item.vue"]]);const Ge=C({name:"ElDropdownMenu",props:Fe,setup(e){const n=m("dropdown"),{_elDropdownSize:t}=Pe(),l=t.value,{focusTrapRef:r,onKeydown:a}=S(ne,void 0),{contentRef:i,role:s,triggerId:d}=S($e,void 0),{collectionRef:u,getItems:c}=S(Be,void 0),{rovingFocusGroupRef:f,rovingFocusGroupRootStyle:v,tabIndex:g,onBlur:b,onFocus:w,onMousedown:h}=S(fe,void 0),{collectionRef:y}=S(ce,void 0),I=B((()=>[n.b("menu"),n.bm("menu",null==l?void 0:l.value)])),E=te(i,u,r,f,y),C=o((o=>{var n;null==(n=e.onKeydown)||n.call(e,o)}),(e=>{const{currentTarget:o,code:n,target:t}=e;if(o.contains(t),p.tab===n&&e.stopImmediatePropagation(),e.preventDefault(),t!==x(i))return;if(!Re.includes(n))return;const l=c().filter((e=>!e.disabled)).map((e=>e.ref));Se.includes(n)&&l.reverse(),ge(l)}));return{size:l,rovingFocusGroupRootStyle:v,tabIndex:g,dropdownKls:I,role:s,triggerId:d,dropdownListWrapperRef:E,handleKeydown:e=>{C(e),a(e)},onBlur:b,onFocus:w,onMousedown:h}}}),Ae=["role","aria-labelledby"];var ze=s(Ge,[["render",function(e,o,n,t,l,r){return K(),j("ul",{ref:e.dropdownListWrapperRef,class:Y(e.dropdownKls),style:Z(e.rovingFocusGroupRootStyle),tabindex:-1,role:e.role,"aria-labelledby":e.triggerId,onBlur:o[0]||(o[0]=(...o)=>e.onBlur&&e.onBlur(...o)),onFocus:o[1]||(o[1]=(...o)=>e.onFocus&&e.onFocus(...o)),onKeydown:o[2]||(o[2]=X(((...o)=>e.handleKeydown&&e.handleKeydown(...o)),["self"])),onMousedown:o[3]||(o[3]=X(((...o)=>e.onMousedown&&e.onMousedown(...o)),["self"]))},[_(e.$slots,"default")],46,Ae)}],["__file","dropdown-menu.vue"]]);const je=b(Me,{DropdownItem:Ne,DropdownMenu:ze}),He=w(Ne),Je=w(ze);export{He as E,Je as a,je as b,te as c};
