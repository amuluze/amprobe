import{b as t,d as e,j as s,_ as i,m as r}from"./el-button.wSr2vfOs.js";import{d as a,c as o,a as l,k as n,n as d,u as c,j as u,i as f,g as v}from"./index.C-XteKAQ.js";const p=t({direction:{type:String,values:["horizontal","vertical"],default:"horizontal"},contentPosition:{type:String,values:["left","center","right"],default:"center"},borderStyle:{type:e(String),default:"solid"}}),y=a({name:"ElDivider"});const m=r(i(a({...y,props:p,setup(t){const e=t,i=s("divider"),r=o((()=>i.cssVar({"border-style":e.borderStyle})));return(t,e)=>(l(),n("div",{class:d([c(i).b(),c(i).m(t.direction)]),style:v(c(r)),role:"separator"},[t.$slots.default&&"vertical"!==t.direction?(l(),n("div",{key:0,class:d([c(i).e("text"),c(i).is(t.contentPosition)])},[u(t.$slots,"default")],2)):f("v-if",!0)],6))}}),[["__file","divider.vue"]]));export{m as E};
