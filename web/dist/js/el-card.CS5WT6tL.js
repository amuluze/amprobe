import{b as s,d as a,n as e,_ as t,p as o}from"./message.CS_TSLeJ.js";import{d as r,a as d,k as l,n as f,u as i,j as y,A as n,t as p,i as c,f as h,g as u}from"./index.DuNUv9gK.js";const v=s({header:{type:String,default:""},footer:{type:String,default:""},bodyStyle:{type:a([String,Object,Array]),default:""},bodyClass:String,shadow:{type:String,values:["always","hover","never"],default:"always"}}),b=r({name:"ElCard"});const g=o(t(r({...b,props:v,setup(s){const a=e("card");return(s,e)=>(d(),l("div",{class:f([i(a).b(),i(a).is(`${s.shadow}-shadow`)])},[s.$slots.header||s.header?(d(),l("div",{key:0,class:f(i(a).e("header"))},[y(s.$slots,"header",{},(()=>[n(p(s.header),1)]))],2)):c("v-if",!0),h("div",{class:f([i(a).e("body"),s.bodyClass]),style:u(s.bodyStyle)},[y(s.$slots,"default")],6),s.$slots.footer||s.footer?(d(),l("div",{key:1,class:f(i(a).e("footer"))},[y(s.$slots,"footer",{},(()=>[n(p(s.footer),1)]))],2)):c("v-if",!0)],2))}}),[["__file","card.vue"]]));export{g as E};