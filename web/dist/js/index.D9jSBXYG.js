import{d as s,c as a,a as e,k as l,f as o,j as n,n as t,u as c,b as i,w as r,m as u,l as p,i as f,g as m,T as d}from"./index.Dr9aOSLV.js";import{b as k,g,z as y,j as b,G as v,E as C,_,m as h}from"./el-button.Bh3a0xQF.js";const j=k({type:{type:String,values:["primary","success","info","warning","danger"],default:"primary"},closable:Boolean,disableTransitions:Boolean,hit:Boolean,color:String,size:{type:String,values:g},effect:{type:String,values:["dark","light","plain"],default:"light"},round:Boolean}),E={close:s=>s instanceof MouseEvent,click:s=>s instanceof MouseEvent},B=s({name:"ElTag"});const S=h(_(s({...B,props:j,emits:E,setup(s,{emit:k}){const g=s,_=y(),h=b("tag"),j=a((()=>{const{type:s,hit:a,effect:e,closable:l,round:o}=g;return[h.b(),h.is("closable",l),h.m(s||"primary"),h.m(_.value),h.m(e),h.is("hit",a),h.is("round",o)]})),E=s=>{k("close",s)},B=s=>{k("click",s)};return(s,a)=>s.disableTransitions?(e(),l("span",{key:0,class:t(c(j)),style:m({backgroundColor:s.color}),onClick:B},[o("span",{class:t(c(h).e("content"))},[n(s.$slots,"default")],2),s.closable?(e(),i(c(C),{key:0,class:t(c(h).e("close")),onClick:p(E,["stop"])},{default:r((()=>[u(c(v))])),_:1},8,["class","onClick"])):f("v-if",!0)],6)):(e(),i(d,{key:1,name:`${c(h).namespace.value}-zoom-in-center`,appear:""},{default:r((()=>[o("span",{class:t(c(j)),style:m({backgroundColor:s.color}),onClick:B},[o("span",{class:t(c(h).e("content"))},[n(s.$slots,"default")],2),s.closable?(e(),i(c(C),{key:0,class:t(c(h).e("close")),onClick:p(E,["stop"])},{default:r((()=>[u(c(v))])),_:1},8,["class","onClick"])):f("v-if",!0)],6)])),_:3},8,["name"]))}}),[["__file","tag.vue"]]));export{S as E,j as t};