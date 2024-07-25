import{o as e,N as o,K as s,d as a,I as t,c as l,a as n,k as r,f as d,j as i,n as c,u,t as f,m as v,w as m,b as p,h as g,i as b,g as h,H as y,r as C,J as w,e as E,E as k,a4 as R,v as $,T as L,ac as A}from"./index.uDfNkopd.js";import{d as _,b as M,c as x,e as F,f as I,E as j,a as T}from"./use-dialog.UpcB4Oaf.js";import{o as Y,H as X,E as B,_ as q,C as z,Y as D,n as H,p as P}from"./message.BD1JyyvK.js";import{F as S,E as K}from"./focus-trap.Y0jdQh1F.js";import{c as N}from"./refs.DcSG14C5.js";const J=(a,t,l,n)=>{let r={offsetX:0,offsetY:0};const d=e=>{const o=e.clientX,s=e.clientY,{offsetX:t,offsetY:l}=r,d=a.value.getBoundingClientRect(),i=d.left,c=d.top,u=d.width,f=d.height,v=document.documentElement.clientWidth,m=document.documentElement.clientHeight,p=-i+t,g=-c+l,b=v-i-u+t,h=m-c-f+l,y=e=>{let d=t+e.clientX-o,i=l+e.clientY-s;(null==n?void 0:n.value)||(d=Math.min(Math.max(d,p),b),i=Math.min(Math.max(i,g),h)),r={offsetX:d,offsetY:i},a.value&&(a.value.style.transform=`translate(${Y(d)}, ${Y(i)})`)},C=()=>{document.removeEventListener("mousemove",y),document.removeEventListener("mouseup",C)};document.addEventListener("mousemove",y),document.addEventListener("mouseup",C)},i=()=>{t.value&&a.value&&t.value.removeEventListener("mousedown",d)};e((()=>{o((()=>{l.value?t.value&&a.value&&t.value.addEventListener("mousedown",d):i()}))})),s((()=>{i()}))},O=Symbol("dialogInjectionKey"),U=["aria-level"],W=["aria-label"],G=["id"],Q=a({name:"ElDialogContent"});var V=q(a({...Q,props:_,emits:M,setup(e){const o=e,{t:s}=X(),{Close:a}=z,{dialogRef:y,headerRef:C,bodyId:w,ns:E,style:k}=t(O),{focusTrapRef:R}=t(S),$=l((()=>[E.b(),E.is("fullscreen",o.fullscreen),E.is("draggable",o.draggable),E.is("align-center",o.alignCenter),{[E.m("center")]:o.center}])),L=N(R,y),A=l((()=>o.draggable)),_=l((()=>o.overflow));return J(y,C,A,_),(e,o)=>(n(),r("div",{ref:u(L),class:c(u($)),style:h(u(k)),tabindex:"-1"},[d("header",{ref_key:"headerRef",ref:C,class:c([u(E).e("header"),{"show-close":e.showClose}])},[i(e.$slots,"header",{},(()=>[d("span",{role:"heading","aria-level":e.ariaLevel,class:c(u(E).e("title"))},f(e.title),11,U)])),e.showClose?(n(),r("button",{key:0,"aria-label":u(s)("el.dialog.close"),class:c(u(E).e("headerbtn")),type:"button",onClick:o[0]||(o[0]=o=>e.$emit("close"))},[v(u(B),{class:c(u(E).e("close"))},{default:m((()=>[(n(),p(g(e.closeIcon||u(a))))])),_:1},8,["class"])],10,W)):b("v-if",!0)],2),d("div",{id:u(w),class:c(u(E).e("body"))},[i(e.$slots,"default")],10,G),e.$slots.footer?(n(),r("footer",{key:0,class:c(u(E).e("footer"))},[i(e.$slots,"footer")],2)):b("v-if",!0)],6))}}),[["__file","dialog-content.vue"]]);const Z=["aria-label","aria-labelledby","aria-describedby"],ee=a({name:"ElDialog",inheritAttrs:!1});const oe=P(q(a({...ee,props:x,emits:F,setup(e,{expose:o}){const s=e,a=y();D({scope:"el-dialog",from:"the title slot",replacement:"the header slot",version:"3.0.0",ref:"https://element-plus.org/en-US/component/dialog.html#slots"},l((()=>!!a.title)));const t=H("dialog"),r=C(),f=C(),g=C(),{visible:_,titleId:M,bodyId:x,style:F,overlayDialogStyle:Y,rendered:X,zIndex:B,afterEnter:q,afterLeave:z,beforeLeave:P,handleClose:S,onModalClick:N,onOpenAutoFocus:J,onCloseAutoFocus:U,onCloseRequested:W,onFocusoutPrevented:G}=I(s,r);w(O,{dialogRef:r,headerRef:f,bodyId:x,ns:t,rendered:X,style:F});const Q=T(N),ee=l((()=>s.draggable&&!s.fullscreen));return o({visible:_,dialogContentRef:g}),(e,o)=>(n(),p(A,{to:e.appendTo,disabled:"body"===e.appendTo&&!e.appendToBody},[v(L,{name:"dialog-fade",onAfterEnter:u(q),onAfterLeave:u(z),onBeforeLeave:u(P),persisted:""},{default:m((()=>[E(v(u(j),{"custom-mask-event":"",mask:e.modal,"overlay-class":e.modalClass,"z-index":u(B)},{default:m((()=>[d("div",{role:"dialog","aria-modal":"true","aria-label":e.title||void 0,"aria-labelledby":e.title?void 0:u(M),"aria-describedby":u(x),class:c(`${u(t).namespace.value}-overlay-dialog`),style:h(u(Y)),onClick:o[0]||(o[0]=(...e)=>u(Q).onClick&&u(Q).onClick(...e)),onMousedown:o[1]||(o[1]=(...e)=>u(Q).onMousedown&&u(Q).onMousedown(...e)),onMouseup:o[2]||(o[2]=(...e)=>u(Q).onMouseup&&u(Q).onMouseup(...e))},[v(u(K),{loop:"",trapped:u(_),"focus-start-el":"container",onFocusAfterTrapped:u(J),onFocusAfterReleased:u(U),onFocusoutPrevented:u(G),onReleaseRequested:u(W)},{default:m((()=>[u(X)?(n(),p(V,k({key:0,ref_key:"dialogContentRef",ref:g},e.$attrs,{center:e.center,"align-center":e.alignCenter,"close-icon":e.closeIcon,draggable:u(ee),overflow:e.overflow,fullscreen:e.fullscreen,"show-close":e.showClose,title:e.title,"aria-level":e.headerAriaLevel,onClose:u(S)}),R({header:m((()=>[e.$slots.title?i(e.$slots,"title",{key:1}):i(e.$slots,"header",{key:0,close:u(S),titleId:u(M),titleClass:u(t).e("title")})])),default:m((()=>[i(e.$slots,"default")])),_:2},[e.$slots.footer?{name:"footer",fn:m((()=>[i(e.$slots,"footer")]))}:void 0]),1040,["center","align-center","close-icon","draggable","overflow","fullscreen","show-close","title","aria-level","onClose"])):b("v-if",!0)])),_:3},8,["trapped","onFocusAfterTrapped","onFocusAfterReleased","onFocusoutPrevented","onReleaseRequested"])],46,Z)])),_:3},8,["mask","overlay-class","z-index"]),[[$,u(_)]])])),_:3},8,["onAfterEnter","onAfterLeave","onBeforeLeave"])],8,["to","disabled"]))}}),[["__file","dialog.vue"]]));export{oe as E,J as u};