import{c as e,J as t,r as a,Y as l,B as o,V as n,q as s,d as i,aq as r,C as u,M as p,o as d,N as c,a as f,k as v,i as m,F as x,n as y,u as g,j as h,f as b,b as w,w as S,h as k,D as z,m as C,l as $,a5 as F,t as I,g as E,L as B}from"./index.Dr9aOSLV.js";import{u as j,m as P,b as M}from"./message.DctIvaHE.js";import{c as N,al as V,h as _,b as R,$ as A,d as H,i as L,z as O,H as T,j as K,a1 as W,av as D,aw as q,v as U,O as Y,E as J,D as Z,_ as G,m as Q}from"./el-button.Bh3a0xQF.js";import{u as X,d as ee}from"./index.BRGrvrSC.js";import{U as te}from"./event.BwRzfsZt.js";import{a as ae,b as le}from"./use-form-item.CQAx4y3m.js";const oe=e=>/([\uAC00-\uD7AF\u3130-\u318F])+/gi.test(e),ne=["class","style"],se=/^on[A-Z]/;function ie(e,{afterFocus:s,beforeBlur:i,afterBlur:r}={}){const u=t(),{emit:p}=u,d=l(),c=a(!1);return o(d,(e=>{e&&e.setAttribute("tabindex","-1")})),j(d,"click",(()=>{var t;null==(t=e.value)||t.focus()})),{wrapperRef:d,isFocused:c,handleFocus:e=>{c.value||(c.value=!0,p("focus",e),null==s||s())},handleBlur:e=>{var t;!!n(i)&&i(e)||e.relatedTarget&&(null==(t=d.value)?void 0:t.contains(e.relatedTarget))||(c.value=!1,p("blur",e),null==r||r())}}}let re;const ue=`\n  height:0 !important;\n  visibility:hidden !important;\n  ${N&&/firefox/i.test(window.navigator.userAgent)?"":"overflow:hidden !important;"}\n  position:absolute !important;\n  z-index:-1000 !important;\n  top:0 !important;\n  right:0 !important;\n`,pe=["letter-spacing","line-height","padding-top","padding-bottom","font-family","font-weight","font-size","text-rendering","text-transform","width","text-indent","padding-left","padding-right","border-width","box-sizing"];function de(e,t=1,a){var l;re||(re=document.createElement("textarea"),document.body.appendChild(re));const{paddingSize:o,borderSize:n,boxSizing:s,contextStyle:i}=function(e){const t=window.getComputedStyle(e),a=t.getPropertyValue("box-sizing"),l=Number.parseFloat(t.getPropertyValue("padding-bottom"))+Number.parseFloat(t.getPropertyValue("padding-top")),o=Number.parseFloat(t.getPropertyValue("border-bottom-width"))+Number.parseFloat(t.getPropertyValue("border-top-width"));return{contextStyle:pe.map((e=>`${e}:${t.getPropertyValue(e)}`)).join(";"),paddingSize:l,borderSize:o,boxSizing:a}}(e);re.setAttribute("style",`${i};${ue}`),re.value=e.value||e.placeholder||"";let r=re.scrollHeight;const u={};"border-box"===s?r+=n:"content-box"===s&&(r-=o),re.value="";const p=re.scrollHeight-o;if(_(t)){let e=p*t;"border-box"===s&&(e=e+o+n),r=Math.max(e,r),u.minHeight=`${e}px`}if(_(a)){let e=p*a;"border-box"===s&&(e=e+o+n),r=Math.min(e,r)}return u.height=`${r}px`,null==(l=re.parentNode)||l.removeChild(re),re=void 0,u}const ce=R({id:{type:String,default:void 0},size:A,disabled:Boolean,modelValue:{type:H([String,Number,Object]),default:""},maxlength:{type:[String,Number]},minlength:{type:[String,Number]},type:{type:String,default:"text"},resize:{type:String,values:["none","both","horizontal","vertical"]},autosize:{type:H([Boolean,Object]),default:!1},autocomplete:{type:String,default:"off"},formatter:{type:Function},parser:{type:Function},placeholder:{type:String},form:{type:String},readonly:{type:Boolean,default:!1},clearable:{type:Boolean,default:!1},showPassword:{type:Boolean,default:!1},showWordLimit:{type:Boolean,default:!1},suffixIcon:{type:L},prefixIcon:{type:L},containerRole:{type:String,default:void 0},label:{type:String,default:void 0},tabindex:{type:[String,Number],default:0},validateEvent:{type:Boolean,default:!0},inputStyle:{type:H([Object,Array,String]),default:()=>P({})},autofocus:{type:Boolean,default:!1},...X(["ariaLabel"])}),fe={[te]:e=>s(e),input:e=>s(e),change:e=>s(e),focus:e=>e instanceof FocusEvent,blur:e=>e instanceof FocusEvent,clear:()=>!0,mouseleave:e=>e instanceof MouseEvent,mouseenter:e=>e instanceof MouseEvent,keydown:e=>e instanceof Event,compositionstart:e=>e instanceof CompositionEvent,compositionupdate:e=>e instanceof CompositionEvent,compositionend:e=>e instanceof CompositionEvent},ve=["role"],me=["id","minlength","maxlength","type","disabled","readonly","autocomplete","tabindex","aria-label","placeholder","form","autofocus"],xe=["id","minlength","maxlength","tabindex","disabled","readonly","autocomplete","aria-label","placeholder","form","autofocus"],ye=i({name:"ElInput",inheritAttrs:!1});const ge=Q(G(i({...ye,props:ce,emits:fe,setup(n,{expose:s,emit:i}){const j=n,P=r(),_=u(),R=e((()=>{const e={};return"combobox"===j.containerRole&&(e["aria-haspopup"]=P["aria-haspopup"],e["aria-owns"]=P["aria-owns"],e["aria-expanded"]=P["aria-expanded"]),e})),A=e((()=>["textarea"===j.type?ce.b():pe.b(),pe.m(re.value),pe.is("disabled",ue.value),pe.is("exceed",He.value),{[pe.b("group")]:_.prepend||_.append,[pe.m("prefix")]:_.prefix||j.prefixIcon,[pe.m("suffix")]:_.suffix||j.suffixIcon||j.clearable||j.showPassword,[pe.bm("suffix","password-clear")]:Ve.value&&_e.value,[pe.b("hidden")]:"hidden"===j.type},P.class])),H=e((()=>[pe.e("wrapper"),pe.is("focus",Ce.value)])),L=((a={})=>{const{excludeListeners:l=!1,excludeKeys:o}=a,n=e((()=>((null==o?void 0:o.value)||[]).concat(ne))),s=t();return e(s?()=>{var e;return V(Object.entries(null==(e=s.proxy)?void 0:e.$attrs).filter((([e])=>!(n.value.includes(e)||l&&se.test(e)))))}:()=>({}))})({excludeKeys:e((()=>Object.keys(R.value)))}),{form:G,formItem:Q}=ae(),{inputId:X}=le(j,{formItemContext:Q}),re=O(),ue=T(),pe=K("input"),ce=K("textarea"),fe=l(),ye=l(),ge=a(!1),he=a(!1),be=a(!1),we=a(),Se=l(j.inputStyle),ke=e((()=>fe.value||ye.value)),{wrapperRef:ze,isFocused:Ce,handleFocus:$e,handleBlur:Fe}=ie(ke,{afterBlur(){var e;j.validateEvent&&(null==(e=null==Q?void 0:Q.validate)||e.call(Q,"blur").catch((e=>ee())))}}),Ie=e((()=>{var e;return null!=(e=null==G?void 0:G.statusIcon)&&e})),Ee=e((()=>(null==Q?void 0:Q.validateState)||"")),Be=e((()=>Ee.value&&W[Ee.value])),je=e((()=>be.value?D:q)),Pe=e((()=>[P.style])),Me=e((()=>[j.inputStyle,Se.value,{resize:j.resize}])),Ne=e((()=>U(j.modelValue)?"":String(j.modelValue))),Ve=e((()=>j.clearable&&!ue.value&&!j.readonly&&!!Ne.value&&(Ce.value||ge.value))),_e=e((()=>j.showPassword&&!ue.value&&!j.readonly&&!!Ne.value&&(!!Ne.value||Ce.value))),Re=e((()=>j.showWordLimit&&!!j.maxlength&&("text"===j.type||"textarea"===j.type)&&!ue.value&&!j.readonly&&!j.showPassword)),Ae=e((()=>Ne.value.length)),He=e((()=>!!Re.value&&Ae.value>Number(j.maxlength))),Le=e((()=>!!_.suffix||!!j.suffixIcon||Ve.value||j.showPassword||Re.value||!!Ee.value&&Ie.value)),[Oe,Te]=function(e){const t=a();return[function(){if(null==e.value)return;const{selectionStart:a,selectionEnd:l,value:o}=e.value;if(null==a||null==l)return;const n=o.slice(0,Math.max(0,a)),s=o.slice(Math.max(0,l));t.value={selectionStart:a,selectionEnd:l,value:o,beforeTxt:n,afterTxt:s}},function(){if(null==e.value||null==t.value)return;const{value:a}=e.value,{beforeTxt:l,afterTxt:o,selectionStart:n}=t.value;if(null==l||null==o||null==n)return;let s=a.length;if(a.endsWith(o))s=a.length-o.length;else if(a.startsWith(l))s=l.length;else{const e=l[n-1],t=a.indexOf(e,n-1);-1!==t&&(s=t+1)}e.value.setSelectionRange(s,s)}]}(fe);M(ye,(e=>{if(We(),!Re.value||"both"!==j.resize)return;const t=e[0],{width:a}=t.contentRect;we.value={right:`calc(100% - ${a+15+6}px)`}}));const Ke=()=>{const{type:e,autosize:t}=j;if(N&&"textarea"===e&&ye.value)if(t){const e=B(t)?t.minRows:void 0,a=B(t)?t.maxRows:void 0,l=de(ye.value,e,a);Se.value={overflowY:"hidden",...l},p((()=>{ye.value.offsetHeight,Se.value=l}))}else Se.value={minHeight:de(ye.value).minHeight}},We=(e=>{let t=!1;return()=>{var a;if(t||!j.autosize)return;null===(null==(a=ye.value)?void 0:a.offsetParent)||(e(),t=!0)}})(Ke),De=()=>{const e=ke.value,t=j.formatter?j.formatter(Ne.value):Ne.value;e&&e.value!==t&&(e.value=t)},qe=async e=>{Oe();let{value:t}=e.target;j.formatter&&(t=j.parser?j.parser(t):t),he.value||(t!==Ne.value?(i(te,t),i("input",t),await p(),De(),Te()):De())},Ue=e=>{i("change",e.target.value)},Ye=e=>{i("compositionstart",e),he.value=!0},Je=e=>{var t;i("compositionupdate",e);const a=null==(t=e.target)?void 0:t.value,l=a[a.length-1]||"";he.value=!oe(l)},Ze=e=>{i("compositionend",e),he.value&&(he.value=!1,qe(e))},Ge=()=>{be.value=!be.value,Qe()},Qe=async()=>{var e;await p(),null==(e=ke.value)||e.focus()},Xe=e=>{ge.value=!1,i("mouseleave",e)},et=e=>{ge.value=!0,i("mouseenter",e)},tt=e=>{i("keydown",e)},at=()=>{i(te,""),i("change",""),i("clear"),i("input","")};return o((()=>j.modelValue),(()=>{var e;p((()=>Ke())),j.validateEvent&&(null==(e=null==Q?void 0:Q.validate)||e.call(Q,"change").catch((e=>ee())))})),o(Ne,(()=>De())),o((()=>j.type),(async()=>{await p(),De(),Ke()})),d((()=>{!j.formatter&&j.parser,De(),p(Ke)})),Y({from:"label",replacement:"aria-label",version:"2.8.0",scope:"el-input",ref:"https://element-plus.org/en-US/component/input.html"},e((()=>!!j.label))),s({input:fe,textarea:ye,ref:ke,textareaStyle:Me,autosize:c(j,"autosize"),focus:Qe,blur:()=>{var e;return null==(e=ke.value)?void 0:e.blur()},select:()=>{var e;null==(e=ke.value)||e.select()},clear:at,resizeTextarea:Ke}),(e,t)=>(f(),v("div",z(g(R),{class:[g(A),{[g(pe).bm("group","append")]:e.$slots.append,[g(pe).bm("group","prepend")]:e.$slots.prepend}],style:g(Pe),role:e.containerRole,onMouseenter:et,onMouseleave:Xe}),[m(" input "),"textarea"!==e.type?(f(),v(x,{key:0},[m(" prepend slot "),e.$slots.prepend?(f(),v("div",{key:0,class:y(g(pe).be("group","prepend"))},[h(e.$slots,"prepend")],2)):m("v-if",!0),b("div",{ref_key:"wrapperRef",ref:ze,class:y(g(H))},[m(" prefix slot "),e.$slots.prefix||e.prefixIcon?(f(),v("span",{key:0,class:y(g(pe).e("prefix"))},[b("span",{class:y(g(pe).e("prefix-inner"))},[h(e.$slots,"prefix"),e.prefixIcon?(f(),w(g(J),{key:0,class:y(g(pe).e("icon"))},{default:S((()=>[(f(),w(k(e.prefixIcon)))])),_:1},8,["class"])):m("v-if",!0)],2)],2)):m("v-if",!0),b("input",z({id:g(X),ref_key:"input",ref:fe,class:g(pe).e("inner")},g(L),{minlength:e.minlength,maxlength:e.maxlength,type:e.showPassword?be.value?"text":"password":e.type,disabled:g(ue),readonly:e.readonly,autocomplete:e.autocomplete,tabindex:e.tabindex,"aria-label":e.label||e.ariaLabel,placeholder:e.placeholder,style:e.inputStyle,form:e.form,autofocus:e.autofocus,onCompositionstart:Ye,onCompositionupdate:Je,onCompositionend:Ze,onInput:qe,onFocus:t[0]||(t[0]=(...e)=>g($e)&&g($e)(...e)),onBlur:t[1]||(t[1]=(...e)=>g(Fe)&&g(Fe)(...e)),onChange:Ue,onKeydown:tt}),null,16,me),m(" suffix slot "),g(Le)?(f(),v("span",{key:1,class:y(g(pe).e("suffix"))},[b("span",{class:y(g(pe).e("suffix-inner"))},[g(Ve)&&g(_e)&&g(Re)?m("v-if",!0):(f(),v(x,{key:0},[h(e.$slots,"suffix"),e.suffixIcon?(f(),w(g(J),{key:0,class:y(g(pe).e("icon"))},{default:S((()=>[(f(),w(k(e.suffixIcon)))])),_:1},8,["class"])):m("v-if",!0)],64)),g(Ve)?(f(),w(g(J),{key:1,class:y([g(pe).e("icon"),g(pe).e("clear")]),onMousedown:$(g(F),["prevent"]),onClick:at},{default:S((()=>[C(g(Z))])),_:1},8,["class","onMousedown"])):m("v-if",!0),g(_e)?(f(),w(g(J),{key:2,class:y([g(pe).e("icon"),g(pe).e("password")]),onClick:Ge},{default:S((()=>[(f(),w(k(g(je))))])),_:1},8,["class"])):m("v-if",!0),g(Re)?(f(),v("span",{key:3,class:y(g(pe).e("count"))},[b("span",{class:y(g(pe).e("count-inner"))},I(g(Ae))+" / "+I(e.maxlength),3)],2)):m("v-if",!0),g(Ee)&&g(Be)&&g(Ie)?(f(),w(g(J),{key:4,class:y([g(pe).e("icon"),g(pe).e("validateIcon"),g(pe).is("loading","validating"===g(Ee))])},{default:S((()=>[(f(),w(k(g(Be))))])),_:1},8,["class"])):m("v-if",!0)],2)],2)):m("v-if",!0)],2),m(" append slot "),e.$slots.append?(f(),v("div",{key:1,class:y(g(pe).be("group","append"))},[h(e.$slots,"append")],2)):m("v-if",!0)],64)):(f(),v(x,{key:1},[m(" textarea "),b("textarea",z({id:g(X),ref_key:"textarea",ref:ye,class:[g(ce).e("inner"),g(pe).is("focus",g(Ce))]},g(L),{minlength:e.minlength,maxlength:e.maxlength,tabindex:e.tabindex,disabled:g(ue),readonly:e.readonly,autocomplete:e.autocomplete,style:g(Me),"aria-label":e.label||e.ariaLabel,placeholder:e.placeholder,form:e.form,autofocus:e.autofocus,onCompositionstart:Ye,onCompositionupdate:Je,onCompositionend:Ze,onInput:qe,onFocus:t[2]||(t[2]=(...e)=>g($e)&&g($e)(...e)),onBlur:t[3]||(t[3]=(...e)=>g(Fe)&&g(Fe)(...e)),onChange:Ue,onKeydown:tt}),null,16,xe),g(Re)?(f(),v("span",{key:0,style:E(we.value),class:y(g(pe).e("count"))},I(g(Ae))+" / "+I(e.maxlength),7)):m("v-if",!0)],64))],16,ve))}}),[["__file","input.vue"]]));export{ge as E,oe as i,ie as u};
