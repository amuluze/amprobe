import{c as e,S as a,r as t,a2 as l,w as o,g as n,a as s,d as i,aq as r,s as u,q as p,j as d,W as c,o as f,e as v,C as m,F as x,n as y,u as g,f as h,D as b,x as w,v as S,y as k,P as z,m as C,H as F,ad as I,B as E,z as B,V as $}from"./index.DR6loZJs.js";import{j as P,P as j,f as V,o as M,b as _,d as N,x as R,u as A,U as H,V as T,z as L,a as O,E as W,M as q,_ as K,w as U}from"./base.CwKrOXiL.js";import{q as D,h as Y,u as Z,p as G,V as J,l as Q}from"./el-overlay.CI2e2OO9.js";import{e as X,b as ee,f as ae}from"./use-form-item.KseOUpWN.js";import{U as te}from"./event.CWrEFZfq.js";import{d as le}from"./error.D_Dr4eZ1.js";const oe=e=>/([\uAC00-\uD7AF\u3130-\u318F])+/gi.test(e),ne=["class","style"],se=/^on[A-Z]/;function ie(e,{afterFocus:s,beforeBlur:i,afterBlur:r}={}){const u=a(),{emit:p}=u,d=l(),c=t(!1);return o(d,(e=>{e&&e.setAttribute("tabindex","-1")})),V(d,"click",(()=>{var a;null==(a=e.value)||a.focus()})),{wrapperRef:d,isFocused:c,handleFocus:e=>{c.value||(c.value=!0,p("focus",e),null==s||s())},handleBlur:e=>{var a;!!n(i)&&i(e)||e.relatedTarget&&(null==(a=d.value)?void 0:a.contains(e.relatedTarget))||(c.value=!1,p("blur",e),null==r||r())}}}let re;const ue=`\n  height:0 !important;\n  visibility:hidden !important;\n  ${P&&/firefox/i.test(window.navigator.userAgent)?"":"overflow:hidden !important;"}\n  position:absolute !important;\n  z-index:-1000 !important;\n  top:0 !important;\n  right:0 !important;\n`,pe=["letter-spacing","line-height","padding-top","padding-bottom","font-family","font-weight","font-size","text-rendering","text-transform","width","text-indent","padding-left","padding-right","border-width","box-sizing"];function de(e,a=1,t){var l;re||(re=document.createElement("textarea"),document.body.appendChild(re));const{paddingSize:o,borderSize:n,boxSizing:s,contextStyle:i}=function(e){const a=window.getComputedStyle(e),t=a.getPropertyValue("box-sizing"),l=Number.parseFloat(a.getPropertyValue("padding-bottom"))+Number.parseFloat(a.getPropertyValue("padding-top")),o=Number.parseFloat(a.getPropertyValue("border-bottom-width"))+Number.parseFloat(a.getPropertyValue("border-top-width"));return{contextStyle:pe.map((e=>`${e}:${a.getPropertyValue(e)}`)).join(";"),paddingSize:l,borderSize:o,boxSizing:t}}(e);re.setAttribute("style",`${i};${ue}`),re.value=e.value||e.placeholder||"";let r=re.scrollHeight;const u={};"border-box"===s?r+=n:"content-box"===s&&(r-=o),re.value="";const p=re.scrollHeight-o;if(M(a)){let e=p*a;"border-box"===s&&(e=e+o+n),r=Math.max(e,r),u.minHeight=`${e}px`}if(M(t)){let e=p*t;"border-box"===s&&(e=e+o+n),r=Math.min(e,r)}return u.height=`${r}px`,null==(l=re.parentNode)||l.removeChild(re),re=void 0,u}const ce=_({id:{type:String,default:void 0},size:D,disabled:Boolean,modelValue:{type:N([String,Number,Object]),default:""},maxlength:{type:[String,Number]},minlength:{type:[String,Number]},type:{type:String,default:"text"},resize:{type:String,values:["none","both","horizontal","vertical"]},autosize:{type:N([Boolean,Object]),default:!1},autocomplete:{type:String,default:"off"},formatter:{type:Function},parser:{type:Function},placeholder:{type:String},form:{type:String},readonly:{type:Boolean,default:!1},clearable:{type:Boolean,default:!1},showPassword:{type:Boolean,default:!1},showWordLimit:{type:Boolean,default:!1},suffixIcon:{type:Y},prefixIcon:{type:Y},containerRole:{type:String,default:void 0},label:{type:String,default:void 0},tabindex:{type:[String,Number],default:0},validateEvent:{type:Boolean,default:!0},inputStyle:{type:N([Object,Array,String]),default:()=>R({})},autofocus:{type:Boolean,default:!1},...X(["ariaLabel"])}),fe={[te]:e=>s(e),input:e=>s(e),change:e=>s(e),focus:e=>e instanceof FocusEvent,blur:e=>e instanceof FocusEvent,clear:()=>!0,mouseleave:e=>e instanceof MouseEvent,mouseenter:e=>e instanceof MouseEvent,keydown:e=>e instanceof Event,compositionstart:e=>e instanceof CompositionEvent,compositionupdate:e=>e instanceof CompositionEvent,compositionend:e=>e instanceof CompositionEvent},ve=["role"],me=["id","minlength","maxlength","type","disabled","readonly","autocomplete","tabindex","aria-label","placeholder","form","autofocus"],xe=["id","minlength","maxlength","tabindex","disabled","readonly","autocomplete","aria-label","placeholder","form","autofocus"],ye=i({name:"ElInput",inheritAttrs:!1});const ge=U(K(i({...ye,props:ce,emits:fe,setup(n,{expose:s,emit:i}){const V=n,M=r(),_=u(),N=e((()=>{const e={};return"combobox"===V.containerRole&&(e["aria-haspopup"]=M["aria-haspopup"],e["aria-owns"]=M["aria-owns"],e["aria-expanded"]=M["aria-expanded"]),e})),R=e((()=>["textarea"===V.type?ce.b():pe.b(),pe.m(re.value),pe.is("disabled",ue.value),pe.is("exceed",He.value),{[pe.b("group")]:_.prepend||_.append,[pe.bm("group","append")]:_.append,[pe.bm("group","prepend")]:_.prepend,[pe.m("prefix")]:_.prefix||V.prefixIcon,[pe.m("suffix")]:_.suffix||V.suffixIcon||V.clearable||V.showPassword,[pe.bm("suffix","password-clear")]:_e.value&&Ne.value,[pe.b("hidden")]:"hidden"===V.type},M.class])),K=e((()=>[pe.e("wrapper"),pe.is("focus",Ce.value)])),U=((t={})=>{const{excludeListeners:l=!1,excludeKeys:o}=t,n=e((()=>((null==o?void 0:o.value)||[]).concat(ne))),s=a();return e(s?()=>{var e;return j(Object.entries(null==(e=s.proxy)?void 0:e.$attrs).filter((([e])=>!(n.value.includes(e)||l&&se.test(e)))))}:()=>({}))})({excludeKeys:e((()=>Object.keys(N.value)))}),{form:D,formItem:Y}=ee(),{inputId:X}=ae(V,{formItemContext:Y}),re=Z(),ue=G(),pe=A("input"),ce=A("textarea"),fe=l(),ye=l(),ge=t(!1),he=t(!1),be=t(!1),we=t(),Se=l(V.inputStyle),ke=e((()=>fe.value||ye.value)),{wrapperRef:ze,isFocused:Ce,handleFocus:Fe,handleBlur:Ie}=ie(ke,{afterBlur(){var e;V.validateEvent&&(null==(e=null==Y?void 0:Y.validate)||e.call(Y,"blur").catch((e=>le())))}}),Ee=e((()=>{var e;return null!=(e=null==D?void 0:D.statusIcon)&&e})),Be=e((()=>(null==Y?void 0:Y.validateState)||"")),$e=e((()=>Be.value&&J[Be.value])),Pe=e((()=>be.value?H:T)),je=e((()=>[M.style])),Ve=e((()=>[V.inputStyle,Se.value,{resize:V.resize}])),Me=e((()=>L(V.modelValue)?"":String(V.modelValue))),_e=e((()=>V.clearable&&!ue.value&&!V.readonly&&!!Me.value&&(Ce.value||ge.value))),Ne=e((()=>V.showPassword&&!ue.value&&!V.readonly&&!!Me.value&&(!!Me.value||Ce.value))),Re=e((()=>V.showWordLimit&&!!V.maxlength&&("text"===V.type||"textarea"===V.type)&&!ue.value&&!V.readonly&&!V.showPassword)),Ae=e((()=>Me.value.length)),He=e((()=>!!Re.value&&Ae.value>Number(V.maxlength))),Te=e((()=>!!_.suffix||!!V.suffixIcon||_e.value||V.showPassword||Re.value||!!Be.value&&Ee.value)),[Le,Oe]=function(e){const a=t();return[function(){if(null==e.value)return;const{selectionStart:t,selectionEnd:l,value:o}=e.value;if(null==t||null==l)return;const n=o.slice(0,Math.max(0,t)),s=o.slice(Math.max(0,l));a.value={selectionStart:t,selectionEnd:l,value:o,beforeTxt:n,afterTxt:s}},function(){if(null==e.value||null==a.value)return;const{value:t}=e.value,{beforeTxt:l,afterTxt:o,selectionStart:n}=a.value;if(null==l||null==o||null==n)return;let s=t.length;if(t.endsWith(o))s=t.length-o.length;else if(t.startsWith(l))s=l.length;else{const e=l[n-1],a=t.indexOf(e,n-1);-1!==a&&(s=a+1)}e.value.setSelectionRange(s,s)}]}(fe);O(ye,(e=>{if(qe(),!Re.value||"both"!==V.resize)return;const a=e[0],{width:t}=a.contentRect;we.value={right:`calc(100% - ${t+15+6}px)`}}));const We=()=>{const{type:e,autosize:a}=V;if(P&&"textarea"===e&&ye.value)if(a){const e=$(a)?a.minRows:void 0,t=$(a)?a.maxRows:void 0,l=de(ye.value,e,t);Se.value={overflowY:"hidden",...l},p((()=>{ye.value.offsetHeight,Se.value=l}))}else Se.value={minHeight:de(ye.value).minHeight}},qe=(e=>{let a=!1;return()=>{var t;if(a||!V.autosize)return;null===(null==(t=ye.value)?void 0:t.offsetParent)||(e(),a=!0)}})(We),Ke=()=>{const e=ke.value,a=V.formatter?V.formatter(Me.value):Me.value;e&&e.value!==a&&(e.value=a)},Ue=async e=>{Le();let{value:a}=e.target;V.formatter&&(a=V.parser?V.parser(a):a),he.value||(a!==Me.value?(i(te,a),i("input",a),await p(),Ke(),Oe()):Ke())},De=e=>{i("change",e.target.value)},Ye=e=>{i("compositionstart",e),he.value=!0},Ze=e=>{var a;i("compositionupdate",e);const t=null==(a=e.target)?void 0:a.value,l=t[t.length-1]||"";he.value=!oe(l)},Ge=e=>{i("compositionend",e),he.value&&(he.value=!1,Ue(e))},Je=()=>{be.value=!be.value,Qe()},Qe=async()=>{var e;await p(),null==(e=ke.value)||e.focus()},Xe=e=>{ge.value=!1,i("mouseleave",e)},ea=e=>{ge.value=!0,i("mouseenter",e)},aa=e=>{i("keydown",e)},ta=()=>{i(te,""),i("change",""),i("clear"),i("input","")};return o((()=>V.modelValue),(()=>{var e;p((()=>We())),V.validateEvent&&(null==(e=null==Y?void 0:Y.validate)||e.call(Y,"change").catch((e=>le())))})),o(Me,(()=>Ke())),o((()=>V.type),(async()=>{await p(),Ke(),We()})),d((()=>{!V.formatter&&V.parser,Ke(),p(We)})),Q({from:"label",replacement:"aria-label",version:"2.8.0",scope:"el-input",ref:"https://element-plus.org/en-US/component/input.html"},e((()=>!!V.label))),s({input:fe,textarea:ye,ref:ke,textareaStyle:Ve,autosize:c(V,"autosize"),focus:Qe,blur:()=>{var e;return null==(e=ke.value)?void 0:e.blur()},select:()=>{var e;null==(e=ke.value)||e.select()},clear:ta,resizeTextarea:We}),(e,a)=>(f(),v("div",z(g(N),{class:g(R),style:g(je),role:e.containerRole,onMouseenter:ea,onMouseleave:Xe}),[m(" input "),"textarea"!==e.type?(f(),v(x,{key:0},[m(" prepend slot "),e.$slots.prepend?(f(),v("div",{key:0,class:y(g(pe).be("group","prepend"))},[h(e.$slots,"prepend")],2)):m("v-if",!0),b("div",{ref_key:"wrapperRef",ref:ze,class:y(g(K))},[m(" prefix slot "),e.$slots.prefix||e.prefixIcon?(f(),v("span",{key:0,class:y(g(pe).e("prefix"))},[b("span",{class:y(g(pe).e("prefix-inner"))},[h(e.$slots,"prefix"),e.prefixIcon?(f(),w(g(W),{key:0,class:y(g(pe).e("icon"))},{default:S((()=>[(f(),w(k(e.prefixIcon)))])),_:1},8,["class"])):m("v-if",!0)],2)],2)):m("v-if",!0),b("input",z({id:g(X),ref_key:"input",ref:fe,class:g(pe).e("inner")},g(U),{minlength:e.minlength,maxlength:e.maxlength,type:e.showPassword?be.value?"text":"password":e.type,disabled:g(ue),readonly:e.readonly,autocomplete:e.autocomplete,tabindex:e.tabindex,"aria-label":e.label||e.ariaLabel,placeholder:e.placeholder,style:e.inputStyle,form:e.form,autofocus:e.autofocus,onCompositionstart:Ye,onCompositionupdate:Ze,onCompositionend:Ge,onInput:Ue,onFocus:a[0]||(a[0]=(...e)=>g(Fe)&&g(Fe)(...e)),onBlur:a[1]||(a[1]=(...e)=>g(Ie)&&g(Ie)(...e)),onChange:De,onKeydown:aa}),null,16,me),m(" suffix slot "),g(Te)?(f(),v("span",{key:1,class:y(g(pe).e("suffix"))},[b("span",{class:y(g(pe).e("suffix-inner"))},[g(_e)&&g(Ne)&&g(Re)?m("v-if",!0):(f(),v(x,{key:0},[h(e.$slots,"suffix"),e.suffixIcon?(f(),w(g(W),{key:0,class:y(g(pe).e("icon"))},{default:S((()=>[(f(),w(k(e.suffixIcon)))])),_:1},8,["class"])):m("v-if",!0)],64)),g(_e)?(f(),w(g(W),{key:1,class:y([g(pe).e("icon"),g(pe).e("clear")]),onMousedown:F(g(I),["prevent"]),onClick:ta},{default:S((()=>[C(g(q))])),_:1},8,["class","onMousedown"])):m("v-if",!0),g(Ne)?(f(),w(g(W),{key:2,class:y([g(pe).e("icon"),g(pe).e("password")]),onClick:Je},{default:S((()=>[(f(),w(k(g(Pe))))])),_:1},8,["class"])):m("v-if",!0),g(Re)?(f(),v("span",{key:3,class:y(g(pe).e("count"))},[b("span",{class:y(g(pe).e("count-inner"))},E(g(Ae))+" / "+E(e.maxlength),3)],2)):m("v-if",!0),g(Be)&&g($e)&&g(Ee)?(f(),w(g(W),{key:4,class:y([g(pe).e("icon"),g(pe).e("validateIcon"),g(pe).is("loading","validating"===g(Be))])},{default:S((()=>[(f(),w(k(g($e))))])),_:1},8,["class"])):m("v-if",!0)],2)],2)):m("v-if",!0)],2),m(" append slot "),e.$slots.append?(f(),v("div",{key:1,class:y(g(pe).be("group","append"))},[h(e.$slots,"append")],2)):m("v-if",!0)],64)):(f(),v(x,{key:1},[m(" textarea "),b("textarea",z({id:g(X),ref_key:"textarea",ref:ye,class:[g(ce).e("inner"),g(pe).is("focus",g(Ce))]},g(U),{minlength:e.minlength,maxlength:e.maxlength,tabindex:e.tabindex,disabled:g(ue),readonly:e.readonly,autocomplete:e.autocomplete,style:g(Ve),"aria-label":e.label||e.ariaLabel,placeholder:e.placeholder,form:e.form,autofocus:e.autofocus,onCompositionstart:Ye,onCompositionupdate:Ze,onCompositionend:Ge,onInput:Ue,onFocus:a[2]||(a[2]=(...e)=>g(Fe)&&g(Fe)(...e)),onBlur:a[3]||(a[3]=(...e)=>g(Ie)&&g(Ie)(...e)),onChange:De,onKeydown:aa}),null,16,xe),g(Re)?(f(),v("span",{key:0,style:B(we.value),class:y(g(pe).e("count"))},E(g(Ae))+" / "+E(e.maxlength),7)):m("v-if",!0)],64))],16,ve))}}),[["__file","input.vue"]]));export{ge as E,oe as i,ie as u};
