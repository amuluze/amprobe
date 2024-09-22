import{d as e,C as a,c as t,r as s,a as l,b as o,m as r,w as d,e as i,u as n,f,D as c,l as u,n as p,k as v,j as b,t as y,i as m,v as h,T as k,ab as A}from"./index.Dr9aOSLV.js";import{b as C,O as w,j as E,y as F,k as L,E as R,G as _,_ as g,m as $}from"./el-button.Bh3a0xQF.js";import{c as j,e as x,f as S,E as T}from"./use-dialog.BrP1_dbF.js";import{E as z}from"./focus-trap.ZmI0uCgZ.js";const B=C({...j,direction:{type:String,default:"rtl",values:["ltr","rtl","ttb","btt"]},size:{type:[String,Number],default:"30%"},withHeader:{type:Boolean,default:!0},modalFade:{type:Boolean,default:!0},headerAriaLevel:{type:String,default:"2"}}),D=x,I=["aria-label","aria-labelledby","aria-describedby"],P=["id","aria-level"],q=["aria-label"],H=["id"],O=e({name:"ElDrawer",inheritAttrs:!1});const G=$(g(e({...O,props:B,emits:D,setup(e,{expose:C}){const g=e,$=a();w({scope:"el-drawer",from:"the title slot",replacement:"the header slot",version:"3.0.0",ref:"https://element-plus.org/en-US/component/drawer.html#slots"},t((()=>!!$.title)));const j=s(),x=s(),B=E("drawer"),{t:D}=F(),{afterEnter:O,afterLeave:G,beforeLeave:M,visible:N,rendered:U,titleId:Y,bodyId:J,zIndex:K,onModalClick:Q,onOpenAutoFocus:V,onCloseAutoFocus:W,onFocusoutPrevented:X,onCloseRequested:Z,handleClose:ee}=S(g,j),ae=t((()=>"rtl"===g.direction||"ltr"===g.direction)),te=t((()=>L(g.size)));return C({handleClose:ee,afterEnter:O,afterLeave:G}),(e,a)=>(l(),o(A,{to:"body",disabled:!e.appendToBody},[r(k,{name:n(B).b("fade"),onAfterEnter:n(O),onAfterLeave:n(G),onBeforeLeave:n(M),persisted:""},{default:d((()=>[i(r(n(T),{mask:e.modal,"overlay-class":e.modalClass,"z-index":n(K),onClick:n(Q)},{default:d((()=>[r(n(z),{loop:"",trapped:n(N),"focus-trap-el":j.value,"focus-start-el":x.value,onFocusAfterTrapped:n(V),onFocusAfterReleased:n(W),onFocusoutPrevented:n(X),onReleaseRequested:n(Z)},{default:d((()=>[f("div",c({ref_key:"drawerRef",ref:j,"aria-modal":"true","aria-label":e.title||void 0,"aria-labelledby":e.title?void 0:n(Y),"aria-describedby":n(J)},e.$attrs,{class:[n(B).b(),e.direction,n(N)&&"open"],style:n(ae)?"width: "+n(te):"height: "+n(te),role:"dialog",onClick:a[1]||(a[1]=u((()=>{}),["stop"]))}),[f("span",{ref_key:"focusStartRef",ref:x,class:p(n(B).e("sr-focus")),tabindex:"-1"},null,2),e.withHeader?(l(),v("header",{key:0,class:p(n(B).e("header"))},[e.$slots.title?b(e.$slots,"title",{key:1},(()=>[m(" DEPRECATED SLOT ")])):b(e.$slots,"header",{key:0,close:n(ee),titleId:n(Y),titleClass:n(B).e("title")},(()=>[e.$slots.title?m("v-if",!0):(l(),v("span",{key:0,id:n(Y),role:"heading","aria-level":e.headerAriaLevel,class:p(n(B).e("title"))},y(e.title),11,P))])),e.showClose?(l(),v("button",{key:2,"aria-label":n(D)("el.drawer.close"),class:p(n(B).e("close-btn")),type:"button",onClick:a[0]||(a[0]=(...e)=>n(ee)&&n(ee)(...e))},[r(n(R),{class:p(n(B).e("close"))},{default:d((()=>[r(n(_))])),_:1},8,["class"])],10,q)):m("v-if",!0)],2)):m("v-if",!0),n(U)?(l(),v("div",{key:1,id:n(J),class:p(n(B).e("body"))},[b(e.$slots,"default")],10,H)):m("v-if",!0),e.$slots.footer?(l(),v("div",{key:2,class:p(n(B).e("footer"))},[b(e.$slots,"footer")],2)):m("v-if",!0)],16,I)])),_:3},8,["trapped","focus-trap-el","focus-start-el","onFocusAfterTrapped","onFocusAfterReleased","onFocusoutPrevented","onReleaseRequested"])])),_:3},8,["mask","overlay-class","z-index","onClick"]),[[h,n(N)]])])),_:3},8,["name","onAfterEnter","onAfterLeave","onBeforeLeave"])],8,["disabled"]))}}),[["__file","drawer.vue"]]));export{G as E};
