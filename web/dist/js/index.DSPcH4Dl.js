import{b as e,g as t,h as a,i as l,d as s,j as n,k as o,E as r,_ as i,m as u,n as c,o as d,r as p,p as m,q as v,s as f,a as h,t as b,v as x}from"./el-button.Bh3a0xQF.js";import{d as g,r as y,c as M,q as I,B as k,a as C,k as _,g as S,u as E,b as w,w as j,h as O,j as T,n as $,C as A,D as N,T as z,E as P,x as B,G as L,o as F,H as q,I as W,e as H,v as R,F as V,J as D,K as G,L as J,M as K,N as Q,O as U,f as Z,A as X,t as Y,_ as ee,m as te,y as ae,z as le,l as se,i as ne,P as oe,Q as re,R as ie,S as ue,U as ce}from"./index.Dr9aOSLV.js";import{E as de,C as pe,a as me}from"./el-popper.wQZhTUhG.js";import"./el-tooltip.l0sNRNKZ.js";import{_ as ve}from"./amprobe.BVrJrCGO.js";import{E as fe,a as he,b as be}from"./el-dropdown-item.CPFkE0B0.js";import{a as xe}from"./index.BqPg4TRo.js";import{_ as ge}from"./index.DRsxk879.js";import{E as ye,m as Me,b as Ie}from"./message.DctIvaHE.js";import{t as ke}from"./aria.C-hsWcn7.js";import{_ as Ce}from"./index.CyV4eoda.js";import{T as _e}from"./index.1DI8xxX2.js";import{t as Se}from"./index.BRGrvrSC.js";import{f as Ee}from"./vnode._yyMxGkh.js";import"./use-form-item.CQAx4y3m.js";import"./focus-trap.ZmI0uCgZ.js";import"./castArray.DwF5sPJL.js";import"./refs.eEe29D1w.js";import"./index.nuX2YVwj.js";const we=e({size:{type:[Number,String],values:t,default:"",validator:e=>a(e)},shape:{type:String,values:["circle","square"],default:"circle"},icon:{type:l},src:{type:String,default:""},alt:String,srcSet:String,fit:{type:s(String),default:"cover"}}),je={error:e=>e instanceof Event},Oe=["src","alt","srcset"],Te=g({name:"ElAvatar"});const $e=u(i(g({...Te,props:we,emits:je,setup(e,{emit:t}){const l=e,s=n("avatar"),i=y(!1),u=M((()=>{const{size:e,icon:t,shape:a}=l,n=[s.b()];return I(e)&&n.push(s.m(e)),t&&n.push(s.m("icon")),a&&n.push(s.m(a)),n})),c=M((()=>{const{size:e}=l;return a(e)?s.cssVarBlock({size:o(e)||""}):void 0})),d=M((()=>({objectFit:l.fit})));function p(e){i.value=!0,t("error",e)}return k((()=>l.src),(()=>i.value=!1)),(e,t)=>(C(),_("span",{class:$(E(u)),style:S(E(c))},[!e.src&&!e.srcSet||i.value?e.icon?(C(),w(E(r),{key:1},{default:j((()=>[(C(),w(O(e.icon)))])),_:1})):T(e.$slots,"default",{key:2}):(C(),_("img",{key:0,src:e.src,alt:e.alt,srcset:e.srcSet,style:S(E(d)),onError:p},null,44,Oe))],6))}}),[["__file","avatar.vue"]])),Ae=g({name:"ElContainer"});var Ne=i(g({...Ae,props:{direction:{type:String}},setup(e){const t=e,a=A(),l=n("container"),s=M((()=>{if("vertical"===t.direction)return!0;if("horizontal"===t.direction)return!1;if(a&&a.default){return a.default().some((e=>{const t=e.type.name;return"ElHeader"===t||"ElFooter"===t}))}return!1}));return(e,t)=>(C(),_("section",{class:$([E(l).b(),E(l).is("vertical",E(s))])},[T(e.$slots,"default")],2))}}),[["__file","container.vue"]]);const ze=g({name:"ElAside"});var Pe=i(g({...ze,props:{width:{type:String,default:null}},setup(e){const t=e,a=n("aside"),l=M((()=>t.width?a.cssVarBlock({width:t.width}):{}));return(e,t)=>(C(),_("aside",{class:$(E(a).b()),style:S(E(l))},[T(e.$slots,"default")],6))}}),[["__file","aside.vue"]]);const Be=g({name:"ElFooter"});var Le=i(g({...Be,props:{height:{type:String,default:null}},setup(e){const t=e,a=n("footer"),l=M((()=>t.height?a.cssVarBlock({height:t.height}):{}));return(e,t)=>(C(),_("footer",{class:$(E(a).b()),style:S(E(l))},[T(e.$slots,"default")],6))}}),[["__file","footer.vue"]]);const Fe=g({name:"ElHeader"});var qe=i(g({...Fe,props:{height:{type:String,default:null}},setup(e){const t=e,a=n("header"),l=M((()=>t.height?a.cssVarBlock({height:t.height}):{}));return(e,t)=>(C(),_("header",{class:$(E(a).b()),style:S(E(l))},[T(e.$slots,"default")],6))}}),[["__file","header.vue"]]);const We=g({name:"ElMain"});var He=i(g({...We,setup(e){const t=n("main");return(e,a)=>(C(),_("main",{class:$(E(t).b())},[T(e.$slots,"default")],2))}}),[["__file","main.vue"]]);const Re=u(Ne,{Aside:Pe,Footer:Le,Header:qe,Main:He}),Ve=c(Pe);c(Le),c(qe),c(He);let De=class{constructor(e,t){this.parent=e,this.domNode=t,this.subIndex=0,this.subIndex=0,this.init()}init(){this.subMenuItems=this.domNode.querySelectorAll("li"),this.addListeners()}gotoSubIndex(e){e===this.subMenuItems.length?e=0:e<0&&(e=this.subMenuItems.length-1),this.subMenuItems[e].focus(),this.subIndex=e}addListeners(){const e=this.parent.domNode;Array.prototype.forEach.call(this.subMenuItems,(t=>{t.addEventListener("keydown",(t=>{let a=!1;switch(t.code){case ye.down:this.gotoSubIndex(this.subIndex+1),a=!0;break;case ye.up:this.gotoSubIndex(this.subIndex-1),a=!0;break;case ye.tab:ke(e,"mouseleave");break;case ye.enter:case ye.space:a=!0,t.currentTarget.click()}return a&&(t.preventDefault(),t.stopPropagation()),!1}))}))}},Ge=class{constructor(e,t){this.domNode=e,this.submenu=null,this.submenu=null,this.init(t)}init(e){this.domNode.setAttribute("tabindex","0");const t=this.domNode.querySelector(`.${e}-menu`);t&&(this.submenu=new De(this,t)),this.addListeners()}addListeners(){this.domNode.addEventListener("keydown",(e=>{let t=!1;switch(e.code){case ye.down:ke(e.currentTarget,"mouseenter"),this.submenu&&this.submenu.gotoSubIndex(0),t=!0;break;case ye.up:ke(e.currentTarget,"mouseenter"),this.submenu&&this.submenu.gotoSubIndex(this.submenu.subMenuItems.length-1),t=!0;break;case ye.tab:ke(e.currentTarget,"mouseleave");break;case ye.enter:case ye.space:t=!0,e.currentTarget.click()}t&&e.preventDefault()}))}},Je=class{constructor(e,t){this.domNode=e,this.init(t)}init(e){const t=this.domNode.childNodes;Array.from(t).forEach((t=>{1===t.nodeType&&new Ge(t,e)}))}};var Ke=i(g({name:"ElMenuCollapseTransition",setup(){const e=n("menu");return{listeners:{onBeforeEnter:e=>e.style.opacity="0.2",onEnter(t,a){d(t,`${e.namespace.value}-opacity-transition`),t.style.opacity="1",a()},onAfterEnter(t){p(t,`${e.namespace.value}-opacity-transition`),t.style.opacity=""},onBeforeLeave(t){t.dataset||(t.dataset={}),m(t,e.m("collapse"))?(p(t,e.m("collapse")),t.dataset.oldOverflow=t.style.overflow,t.dataset.scrollWidth=t.clientWidth.toString(),d(t,e.m("collapse"))):(d(t,e.m("collapse")),t.dataset.oldOverflow=t.style.overflow,t.dataset.scrollWidth=t.clientWidth.toString(),p(t,e.m("collapse"))),t.style.width=`${t.scrollWidth}px`,t.style.overflow="hidden"},onLeave(e){d(e,"horizontal-collapse-transition"),e.style.width=`${e.dataset.scrollWidth}px`}}}}}),[["render",function(e,t,a,l,s,n){return C(),w(z,N({mode:"out-in"},e.listeners),{default:j((()=>[T(e.$slots,"default")])),_:3},16)}],["__file","menu-collapse-transition.vue"]]);function Qe(e,t){const a=M((()=>{let a=e.parent;const l=[t.value];for(;"ElMenu"!==a.type.name;)a.props.index&&l.unshift(a.props.index),a=a.parent;return l}));return{parentMenu:M((()=>{let t=e.parent;for(;t&&!["ElMenu","ElSubMenu"].includes(t.type.name);)t=t.parent;return t})),indexPath:a}}function Ue(e){return M((()=>{const t=e.backgroundColor;return t?new _e(t).shade(20).toString():""}))}const Ze=(e,t)=>{const a=n("menu");return M((()=>a.cssVarBlock({"text-color":e.textColor||"","hover-text-color":e.textColor||"","bg-color":e.backgroundColor||"","hover-bg-color":Ue(e).value||"","active-color":e.activeTextColor||"",level:`${t}`})))},Xe=e({index:{type:String,required:!0},showTimeout:Number,hideTimeout:Number,popperClass:String,disabled:Boolean,teleported:{type:Boolean,default:void 0},popperOffset:Number,expandCloseIcon:{type:l},expandOpenIcon:{type:l},collapseCloseIcon:{type:l},collapseOpenIcon:{type:l}}),Ye="ElSubMenu";var et=g({name:Ye,props:Xe,setup(e,{slots:t,expose:a}){const l=D(),{indexPath:s,parentMenu:o}=Qe(l,M((()=>e.index))),i=n("menu"),u=n("sub-menu"),c=P("rootMenu");c||Se(Ye,"can not inject root menu");const d=P(`subMenu:${o.value.uid}`);d||Se(Ye,"can not inject sub menu");const p=y({}),m=y({});let b;const x=y(!1),g=y(),C=y(null),_=M((()=>"horizontal"===A.value&&E.value?"bottom-start":"right-start")),S=M((()=>"horizontal"===A.value&&E.value||"vertical"===A.value&&!c.props.collapse?e.expandCloseIcon&&e.expandOpenIcon?T.value?e.expandOpenIcon:e.expandCloseIcon:v:e.collapseCloseIcon&&e.collapseOpenIcon?T.value?e.collapseOpenIcon:e.collapseCloseIcon:f)),E=M((()=>0===d.level)),w=M((()=>{const t=e.teleported;return void 0===t?E.value:t})),j=M((()=>c.props.collapse?`${i.namespace.value}-zoom-in-left`:`${i.namespace.value}-zoom-in-top`)),O=M((()=>"horizontal"===A.value&&E.value?["bottom-start","bottom-end","top-start","top-end","right-start","left-start"]:["right-start","right","right-end","left-start","bottom-start","bottom-end","top-start","top-end"])),T=M((()=>c.openedMenus.includes(e.index))),$=M((()=>{let e=!1;return Object.values(p.value).forEach((t=>{t.active&&(e=!0)})),Object.values(m.value).forEach((t=>{t.active&&(e=!0)})),e})),A=M((()=>c.props.mode)),N=B({index:e.index,indexPath:s,active:$}),z=Ze(c.props,d.level+1),G=M((()=>{var t;return null!=(t=e.popperOffset)?t:c.props.popperOffset})),J=M((()=>{var t;return null!=(t=e.popperClass)?t:c.props.popperClass})),K=M((()=>{var t;return null!=(t=e.showTimeout)?t:c.props.showTimeout})),Q=M((()=>{var t;return null!=(t=e.hideTimeout)?t:c.props.hideTimeout})),U=e=>{var t,a,l;e||null==(l=null==(a=null==(t=C.value)?void 0:t.popperRef)?void 0:a.popperInstanceRef)||l.destroy()},Z=()=>{"hover"===c.props.menuTrigger&&"horizontal"===c.props.mode||c.props.collapse&&"vertical"===c.props.mode||e.disabled||c.handleSubMenuClick({index:e.index,indexPath:s.value,active:$.value})},X=(t,a=K.value)=>{var l;"focus"!==t.type&&("click"===c.props.menuTrigger&&"horizontal"===c.props.mode||!c.props.collapse&&"vertical"===c.props.mode||e.disabled?d.mouseInChild.value=!0:(d.mouseInChild.value=!0,null==b||b(),({stop:b}=h((()=>{c.openMenu(e.index,s.value)}),a)),w.value&&(null==(l=o.value.vnode.el)||l.dispatchEvent(new MouseEvent("mouseenter")))))},Y=(t=!1)=>{var a;"click"===c.props.menuTrigger&&"horizontal"===c.props.mode||!c.props.collapse&&"vertical"===c.props.mode?d.mouseInChild.value=!1:(null==b||b(),d.mouseInChild.value=!1,({stop:b}=h((()=>!x.value&&c.closeMenu(e.index,s.value)),Q.value)),w.value&&t&&(null==(a=d.handleMouseleave)||a.call(d,!0)))};k((()=>c.props.collapse),(e=>U(Boolean(e))));{const e=e=>{m.value[e.index]=e},t=e=>{delete m.value[e.index]};L(`subMenu:${l.uid}`,{addSubMenu:e,removeSubMenu:t,handleMouseleave:Y,mouseInChild:x,level:d.level+1})}return a({opened:T}),F((()=>{c.addSubMenu(N),d.addSubMenu(N)})),q((()=>{d.removeSubMenu(N),c.removeSubMenu(N)})),()=>{var a;const s=[null==(a=t.title)?void 0:a.call(t),W(r,{class:u.e("icon-arrow"),style:{transform:T.value?e.expandCloseIcon&&e.expandOpenIcon||e.collapseCloseIcon&&e.collapseOpenIcon&&c.props.collapse?"none":"rotateZ(180deg)":"none"}},{default:()=>I(S.value)?W(l.appContext.components[S.value]):W(S.value)})],n=c.isMenuPopup?W(de,{ref:C,visible:T.value,effect:"light",pure:!0,offset:G.value,showArrow:!1,persistent:!0,popperClass:J.value,placement:_.value,teleported:w.value,fallbackPlacements:O.value,transition:j.value,gpuAcceleration:!1},{content:()=>{var e;return W("div",{class:[i.m(A.value),i.m("popup-container"),J.value],onMouseenter:e=>X(e,100),onMouseleave:()=>Y(!0),onFocus:e=>X(e,100)},[W("ul",{class:[i.b(),i.m("popup"),i.m(`popup-${_.value}`)],style:z.value},[null==(e=t.default)?void 0:e.call(t)])])},default:()=>W("div",{class:u.e("title"),onClick:Z},s)}):W(V,{},[W("div",{class:u.e("title"),ref:g,onClick:Z},s),W(Ce,{},{default:()=>{var e;return H(W("ul",{role:"menu",class:[i.b(),i.m("inline")],style:z.value},[null==(e=t.default)?void 0:e.call(t)]),[[R,T.value]])}})]);return W("li",{class:[u.b(),u.is("active",$.value),u.is("opened",T.value),u.is("disabled",e.disabled)],role:"menuitem",ariaHaspopup:!0,ariaExpanded:T.value,onMouseenter:X,onMouseleave:()=>Y(),onFocus:X},[n])}}});const tt=e({mode:{type:String,values:["horizontal","vertical"],default:"vertical"},defaultActive:{type:String,default:""},defaultOpeneds:{type:s(Array),default:()=>Me([])},uniqueOpened:Boolean,router:Boolean,menuTrigger:{type:String,values:["hover","click"],default:"hover"},collapse:Boolean,backgroundColor:String,textColor:String,activeTextColor:String,closeOnClickOutside:Boolean,collapseTransition:{type:Boolean,default:!0},ellipsis:{type:Boolean,default:!0},popperOffset:{type:Number,default:6},ellipsisIcon:{type:l,default:()=>b},popperEffect:{type:String,values:["dark","light"],default:"dark"},popperClass:String,showTimeout:{type:Number,default:300},hideTimeout:{type:Number,default:300}}),at=e=>Array.isArray(e)&&e.every((e=>I(e)));var lt=g({name:"ElMenu",props:tt,emits:{close:(e,t)=>I(e)&&at(t),open:(e,t)=>I(e)&&at(t),select:(e,t,a,l)=>I(e)&&at(t)&&J(a)&&(void 0===l||l instanceof Promise)},setup(e,{emit:t,slots:a,expose:l}){const s=D(),o=s.appContext.config.globalProperties.$router,i=y(),u=n("menu"),c=n("sub-menu"),d=y(-1),p=y(e.defaultOpeneds&&!e.collapse?e.defaultOpeneds.slice(0):[]),m=y(e.defaultActive),v=y({}),f=y({}),h=M((()=>"horizontal"===e.mode||"vertical"===e.mode&&e.collapse)),b=(a,l)=>{p.value.includes(a)||(e.uniqueOpened&&(p.value=p.value.filter((e=>l.includes(e)))),p.value.push(a),t("open",a,l))},g=e=>{const t=p.value.indexOf(e);-1!==t&&p.value.splice(t,1)},I=(e,a)=>{g(e),t("close",e,a)},C=({index:e,indexPath:t})=>{p.value.includes(e)?I(e,t):b(e,t)},_=a=>{("horizontal"===e.mode||e.collapse)&&(p.value=[]);const{index:l,indexPath:s}=a;if(!x(l)&&!x(s))if(e.router&&o){const e=a.route||l,n=o.push(e).then((e=>(e||(m.value=l),e)));t("select",l,s,{index:l,indexPath:s,route:e},n)}else m.value=l,t("select",l,s,{index:l,indexPath:s})},S=()=>{var e,t;if(!i.value)return-1;const a=Array.from(null!=(t=null==(e=i.value)?void 0:e.childNodes)?t:[]).filter((e=>"#comment"!==e.nodeName&&("#text"!==e.nodeName||e.nodeValue))),l=getComputedStyle(i.value),s=Number.parseInt(l.paddingLeft,10),n=Number.parseInt(l.paddingRight,10),o=i.value.clientWidth-s-n;let r=0,u=0;return a.forEach(((e,t)=>{r+=(e=>{const t=getComputedStyle(e),a=Number.parseInt(t.marginLeft,10),l=Number.parseInt(t.marginRight,10);return e.offsetWidth+a+l||0})(e),r<=o-64&&(u=t+1)})),u===a.length?-1:u};let E=!0;const w=()=>{if(d.value===S())return;const e=()=>{d.value=-1,K((()=>{d.value=S()}))};E?e():((e,t=33.34)=>{let a;return()=>{a&&clearTimeout(a),a=setTimeout((()=>{e()}),t)}})(e)(),E=!1};let j;k((()=>e.defaultActive),(t=>{v.value[t]||(m.value=""),(t=>{const a=v.value,l=a[t]||m.value&&a[m.value]||a[e.defaultActive];m.value=l?l.index:t})(t)})),k((()=>e.collapse),(e=>{e&&(p.value=[])})),k(v.value,(()=>{const t=m.value&&v.value[m.value];if(!t||"horizontal"===e.mode||e.collapse)return;t.indexPath.forEach((e=>{const t=f.value[e];t&&b(e,t.indexPath)}))})),G((()=>{"horizontal"===e.mode&&e.ellipsis?j=Ie(i,w).stop:null==j||j()}));const O=y(!1);{const t=e=>{f.value[e.index]=e},a=e=>{delete f.value[e.index]},l=e=>{v.value[e.index]=e},n=e=>{delete v.value[e.index]};L("rootMenu",B({props:e,openedMenus:p,items:v,subMenus:f,activeIndex:m,isMenuPopup:h,addMenuItem:l,removeMenuItem:n,addSubMenu:t,removeSubMenu:a,openMenu:b,closeMenu:I,handleMenuItemClick:_,handleSubMenuClick:C})),L(`subMenu:${s.uid}`,{addSubMenu:t,removeSubMenu:a,mouseInChild:O,level:0})}F((()=>{"horizontal"===e.mode&&new Je(s.vnode.el,u.namespace.value)}));l({open:e=>{const{indexPath:t}=f.value[e];t.forEach((e=>b(e,t)))},close:g,handleResize:w});return()=>{var l,s;let n=null!=(s=null==(l=a.default)?void 0:l.call(a))?s:[];const o=[];if("horizontal"===e.mode&&i.value){const t=Ee(n),a=-1===d.value?t:t.slice(0,d.value),l=-1===d.value?[]:t.slice(d.value);(null==l?void 0:l.length)&&e.ellipsis&&(n=a,o.push(W(et,{index:"sub-menu-more",class:c.e("hide-arrow"),popperOffset:e.popperOffset},{title:()=>W(r,{class:c.e("icon-more")},{default:()=>W(e.ellipsisIcon)}),default:()=>l})))}const m=Ze(e,0),v=e.closeOnClickOutside?[[pe,()=>{p.value.length&&(O.value||(p.value.forEach((e=>{return t("close",e,(a=e,f.value[a].indexPath));var a})),p.value=[]))}]]:[],h=H(W("ul",{key:String(e.collapse),role:"menubar",ref:i,style:m.value,class:{[u.b()]:!0,[u.m(e.mode)]:!0,[u.m("collapse")]:e.collapse}},[...n,...o]),v);return e.collapseTransition&&"vertical"===e.mode?W(Ke,(()=>h)):h}}});const st=e({index:{type:s([String,null]),default:null},route:{type:s([String,Object])},disabled:Boolean}),nt="ElMenuItem";var ot=i(g({name:nt,components:{ElTooltip:de},props:st,emits:{click:e=>I(e.index)&&Array.isArray(e.indexPath)},setup(e,{emit:t}){const a=D(),l=P("rootMenu"),s=n("menu"),o=n("menu-item");l||Se(nt,"can not inject root menu");const{parentMenu:r,indexPath:i}=Qe(a,Q(e,"index")),u=P(`subMenu:${r.value.uid}`);u||Se(nt,"can not inject sub menu");const c=M((()=>e.index===l.activeIndex)),d=B({index:e.index,indexPath:i,active:c});return F((()=>{u.addSubMenu(d),l.addMenuItem(d)})),q((()=>{u.removeSubMenu(d),l.removeMenuItem(d)})),{parentMenu:r,rootMenu:l,active:c,nsMenu:s,nsMenuItem:o,handleClick:()=>{e.disabled||(l.handleMenuItemClick({index:e.index,indexPath:i.value,route:e.route}),t("click",d))}}}}),[["render",function(e,t,a,l,s,n){const o=U("el-tooltip");return C(),_("li",{class:$([e.nsMenuItem.b(),e.nsMenuItem.is("active",e.active),e.nsMenuItem.is("disabled",e.disabled)]),role:"menuitem",tabindex:"-1",onClick:t[0]||(t[0]=(...t)=>e.handleClick&&e.handleClick(...t))},["ElMenu"===e.parentMenu.type.name&&e.rootMenu.props.collapse&&e.$slots.title?(C(),w(o,{key:0,effect:e.rootMenu.props.popperEffect,placement:"right","fallback-placements":["left"],persistent:""},{content:j((()=>[T(e.$slots,"title")])),default:j((()=>[Z("div",{class:$(e.nsMenu.be("tooltip","trigger"))},[T(e.$slots,"default")],2)])),_:3},8,["effect"])):(C(),_(V,{key:1},[T(e.$slots,"default"),T(e.$slots,"title")],64))],2)}],["__file","menu-item.vue"]]);var rt=i(g({name:"ElMenuItemGroup",props:{title:String},setup:()=>({ns:n("menu-item-group")})}),[["render",function(e,t,a,l,s,n){return C(),_("li",{class:$(e.ns.b())},[Z("div",{class:$(e.ns.e("title"))},[e.$slots.title?T(e.$slots,"title",{key:1}):(C(),_(V,{key:0},[X(Y(e.title),1)],64))],2),Z("ul",null,[T(e.$slots,"default")])],2)}],["__file","menu-item-group.vue"]]);const it=u(lt,{MenuItem:ot,MenuItemGroup:rt,SubMenu:et}),ut=c(ot);c(rt);const ct=c(et),dt={class:"am-content-container"};const pt=ee({},[["render",function(e,t){const a=U("router-view");return C(),_("section",dt,[te(a)])}],["__scopeId","data-v-5690a20d"]]),mt={viewBox:"0 0 1024 1024",width:"1.2em",height:"1.2em"},vt=[Z("path",{fill:"currentColor",d:"M512 512a192 192 0 1 0 0-384a192 192 0 0 0 0 384m0 64a256 256 0 1 1 0-512a256 256 0 0 1 0 512m320 320v-96a96 96 0 0 0-96-96H288a96 96 0 0 0-96 96v96a32 32 0 1 1-64 0v-96a160 160 0 0 1 160-160h448a160 160 0 0 1 160 160v96a32 32 0 1 1-64 0"},null,-1)];const ft={name:"ep-user",render:function(e,t){return C(),_("svg",mt,[...vt])}},ht=ee(g({__name:"Avatar",setup(e){const t=ae(),a=le(),l=async()=>{await xe(),t.user.setToken("",""),t.app.isCollapse=!1,await a.replace("/login")},s=()=>{a.push("/profile")};return(e,t)=>{const a=ft,n=$e,o=fe,r=he,i=be;return C(),w(i,null,{dropdown:j((()=>[te(r,null,{default:j((()=>[te(o,{onClick:se(l,["prevent"])},{default:j((()=>[X("退出登录")])),_:1}),te(o,{onClick:se(s,["prevent"])},{default:j((()=>[X("个人中心")])),_:1})])),_:1})])),default:j((()=>[te(n,{size:"small"},{default:j((()=>[te(a)])),_:1})])),_:1})}}}),[["__scopeId","data-v-80ff9785"]]),bt=ee(g({__name:"CollapseIcon",setup(e){const t=ae(),a=()=>t.app.setCollapse(!t.app.isCollapse);return(e,l)=>{const s=ge;return C(),_("div",{class:"collapse-ico",onClick:a},[E(t).app.isCollapse?(C(),w(s,{key:0,"icon-class":"right"})):(C(),w(s,{key:1,"icon-class":"left"}))])}}}),[["__scopeId","data-v-c47edaaa"]]),xt=ee(g({__name:"Menuitem",props:{item:{}},setup(e){const t=ae(),a=M((()=>t.app.isCollapse)),l=le();return(e,s)=>{var n;const o=ge,r=U("menuitem",!0),i=ct,u=ut;return C(),_("div",{class:$(["am-menuitem-container",{"am-hide-sidebar":E(a)}])},[(null==(n=e.item.children)?void 0:n.length)?(C(),w(i,{key:0,class:"grid",index:e.item.path},{title:j((()=>{var a;return[e.item.meta&&e.item.meta.icon?(C(),w(o,{key:0,"icon-class":e.item.meta.icon},null,8,["icon-class"])):ne("",!0),H(Z("span",null,Y(null==(a=e.item.meta)?void 0:a.title),513),[[R,!E(t).app.isCollapse]])]})),default:j((()=>[(C(!0),_(V,null,oe(e.item.children,(e=>(C(),w(r,{key:e.name,item:e},null,8,["item"])))),128))])),_:1},8,["index"])):(C(),w(u,{key:1,index:e.item.path,onClick:s[0]||(s[0]=t=>{return a=e.item,void l.push(a.path);var a})},{title:j((()=>{var t;return[X(Y(null==(t=e.item.meta)?void 0:t.title),1)]})),default:j((()=>[e.item.meta&&e.item.meta.icon?(C(),w(o,{key:0,"icon-class":e.item.meta.icon},null,8,["icon-class"])):ne("",!0)])),_:1},8,["index"]))],2)}}}),[["__scopeId","data-v-fef109de"]]),gt={class:"am-logo"},yt=(e=>(ie("data-v-cf8d4319"),e=e(),ue(),e))((()=>Z("img",{class:"am-logo__img",src:ve,alt:""},null,-1))),Mt={class:"am-logo__text"},It={class:"am-menu"},kt={class:"am-user"},Ct={class:"am-collapse"},_t=ee(g({__name:"index",setup(e){const t=re(),a=ae(),l=M((()=>ce.filter((e=>{var t;return null==(t=e.meta)?void 0:t.show})))),s=M((()=>t.meta.activeMenu?t.meta.activeMenu:t.path));return(e,t)=>{const n=it,o=me,r=Ve;return C(),w(r,{width:E(a).app.isCollapse?"54px":"160px"},{default:j((()=>[Z("div",gt,[yt,H(Z("span",Mt,"Amprobe",512),[[R,!E(a).app.isCollapse]])]),Z("div",It,[te(o,null,{default:j((()=>[te(n,{"default-active":E(s),collapse:E(a).app.isCollapse,"unique-opened":!0,"collapse-transition":!1,"background-color":"#E9EFFD","text-color":"#000","active-text-color":"#105EEB",mode:"vertical"},{default:j((()=>[(C(!0),_(V,null,oe(E(l),((e,t)=>(C(),w(xt,{key:t,item:e},null,8,["item"])))),128))])),_:1},8,["default-active","collapse"])])),_:1})]),Z("div",kt,[te(ht)]),Z("div",Ct,[te(bt)])])),_:1},8,["width"])}}}),[["__scopeId","data-v-cf8d4319"]]),St={class:"am-layout-container"},Et={class:"am-layout-sidebar"},wt={class:"am-layout-content"},jt=ee(g({__name:"index",setup:e=>(e,t)=>{const a=Re;return C(),_("div",St,[te(a,null,{default:j((()=>[Z("div",Et,[te(_t)]),Z("div",wt,[te(pt)])])),_:1})])}}),[["__scopeId","data-v-ce0339c7"]]);export{jt as default};
