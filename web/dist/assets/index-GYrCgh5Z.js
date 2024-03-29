import{q as H,a as R,b as U,E as F,c as J,_ as K}from"./index-DTPkIpzf.js";import{d as Q,e as V,u as X,k as Y,_ as Z,w as ee,s as j}from"./index-BCNH9aZg.js";import{E as te}from"./el-card-CFODOd__.js";import{q as ae,a as se,b as oe}from"./index-DT5Tlxla.js";import{d as $,c as ne,g as ie,o as b,e as k,n as m,u as a,f as O,A as re,B as n,C as P,D as s,z as le,j as ue,r as z,b as T,m as o,v as i,N as ce,O as pe,_ as de}from"./index-C5VkDeSH.js";const fe=Q({decimalSeparator:{type:String,default:"."},groupSeparator:{type:String,default:","},precision:{type:Number,default:0},formatter:Function,value:{type:V([Number,Object]),default:0},prefix:String,suffix:String,title:String,valueStyle:{type:V([String,Object,Array])}}),me=$({name:"ElStatistic"}),_e=$({...me,props:fe,setup(c,{expose:I}){const x=c,l=X("statistic"),A=ne(()=>{const{value:e,formatter:_,precision:r,decimalSeparator:N,groupSeparator:G}=x;if(ie(_))return _(e);if(!Y(e))return e;let[u,p=""]=String(e).split(".");return p=p.padEnd(r,"0").slice(0,r>0?r:0),u=u.replace(/\B(?=(\d{3})+(?!\d))/g,G),[u,p].join(p?N:"")});return I({displayValue:A}),(e,_)=>(b(),k("div",{class:m(a(l).b())},[e.$slots.title||e.title?(b(),k("div",{key:0,class:m(a(l).e("head"))},[O(e.$slots,"title",{},()=>[re(n(e.title),1)])],2)):P("v-if",!0),s("div",{class:m(a(l).e("content"))},[e.$slots.prefix||e.prefix?(b(),k("div",{key:0,class:m(a(l).e("prefix"))},[O(e.$slots,"prefix",{},()=>[s("span",null,n(e.prefix),1)])],2)):P("v-if",!0),s("span",{class:m(a(l).e("number")),style:le(e.valueStyle)},n(a(A)),7),e.$slots.suffix||e.suffix?(b(),k("div",{key:1,class:m(a(l).e("suffix"))},[O(e.$slots,"suffix",{},()=>[s("span",null,n(e.suffix),1)])],2)):P("v-if",!0)],2)],2))}});var ve=Z(_e,[["__file","statistic.vue"]]);const he=ee(ve),ge={series:[{type:"gauge",startAngle:90,endAngle:-270,pointer:{show:!1},progress:{show:!0,overlap:!1,roundCap:!0,clip:!1,itemStyle:{borderWidth:1,borderColor:"#464646"}},axisLine:{lineStyle:{width:8}},splitLine:{show:!1,distance:0,length:10},axisTick:{show:!1},axisLabel:{show:!1,distance:50},data:[],title:{fontSize:14},detail:{width:50,height:14,fontSize:14,color:"inherit",borderColor:"inherit",borderRadius:20,borderWidth:1,formatter:"{value}%"}}]},ye={series:[{type:"gauge",startAngle:90,endAngle:-270,pointer:{show:!1},progress:{show:!0,overlap:!1,roundCap:!0,clip:!1,itemStyle:{borderWidth:1,borderColor:"#464646"}},axisLine:{lineStyle:{width:8}},splitLine:{show:!1,distance:0,length:10},axisTick:{show:!1},axisLabel:{show:!1,distance:50},data:[],title:{fontSize:14},detail:{width:50,height:14,fontSize:14,color:"inherit",borderColor:"inherit",borderRadius:20,borderWidth:1,formatter:"{value}%"}}]},E=c=>(ce("data-v-d7835d4a"),c=c(),pe(),c),Se={class:"am-overview-container"},we=E(()=>s("h4",null,"概览",-1)),Ce=E(()=>s("h4",null,"状态",-1)),be=E(()=>s("h4",null,"主机信息",-1)),ke=E(()=>s("h4",null,"容器信息",-1)),Ie=$({__name:"index",setup(c){ue(()=>{_(),N(),p(),W(),l(),A()});const I=z(0),x=z(0),l=async()=>{const{data:t}=await ae({page:1,size:10});I.value=t.total},A=async()=>{const{data:t}=await se({page:1,size:10});x.value=t.total},e=z(),_=async()=>{const{data:t}=await H();e.value={hostname:t.hostname,uptime:t.uptime,os:t.os,platform:t.platform,platform_version:t.platform_version,kernel_version:t.kernel_version,kernel_arch:t.kernel_arch}},r=z(),N=async()=>{const{data:t}=await oe();console.log(t),r.value={docker_version:t.docker_version,api_version:t.api_version,min_api_version:t.min_api_version,git_commit:t.git_commit,go_version:t.go_version,os:t.os,arch:t.arch},console.log(r.value)},G=[{value:20,name:"CPU",title:{offsetCenter:["0%","-15%"]},detail:{valueAnimation:!0,offsetCenter:["0%","15%"]}}],u=T(ge),p=async()=>{const{data:t}=await R();G[0].value=Math.round(t.percent*100)/100,j(u,"series[0].data",G),console.log("gauge",u)},M=[{value:20,name:"Memory",title:{offsetCenter:["0%","-15%"]},detail:{valueAnimation:!0,offsetCenter:["0%","15%"]}}],q=T(ye),W=async()=>{const{data:t}=await U();M[0].value=Math.round(t.percent*100)/100,j(q,"series[0].data",M),console.log("gauge",q)};return(t,xe)=>{const L=he,d=te,f=J,D=F,B=K;return b(),k("div",Se,[o(D,{gutter:8},{default:i(()=>[o(f,{span:18},{default:i(()=>[o(d,{class:"am-overview-overview"},{default:i(()=>[we,o(D,{gutter:4},{default:i(()=>[o(f,{span:12},{default:i(()=>[o(d,null,{default:i(()=>[o(L,{title:"容器数量",value:a(I)},null,8,["value"])]),_:1})]),_:1}),o(f,{span:12},{default:i(()=>[o(d,null,{default:i(()=>[o(L,{title:"镜像数量",value:a(x)},null,8,["value"])]),_:1})]),_:1})]),_:1})]),_:1}),o(d,null,{default:i(()=>[Ce,o(D,{gutter:4},{default:i(()=>[o(f,{span:12},{default:i(()=>[o(B,{option:a(u),height:"200px"},null,8,["option"])]),_:1}),o(f,{span:12},{default:i(()=>[o(B,{option:a(q),height:"200px"},null,8,["option"])]),_:1})]),_:1})]),_:1})]),_:1}),o(f,{span:6},{default:i(()=>[o(d,{class:"am-overview-host"},{default:i(()=>{var v,h,g,y,S,w,C;return[be,s("p",null,"主机名称："+n((v=a(e))==null?void 0:v.hostname),1),s("p",null,"启动时间： "+n((h=a(e))==null?void 0:h.uptime),1),s("p",null,"os："+n((g=a(e))==null?void 0:g.os),1),s("p",null,"platform: "+n((y=a(e))==null?void 0:y.platform),1),s("p",null,"platform version: "+n((S=a(e))==null?void 0:S.platform_version),1),s("p",null,"kernel version: "+n((w=a(e))==null?void 0:w.kernel_version),1),s("p",null,"kernel arch: "+n((C=a(e))==null?void 0:C.kernel_arch),1)]}),_:1}),o(d,null,{default:i(()=>{var v,h,g,y,S,w,C;return[ke,s("p",null,"Docker 版本："+n((v=a(r))==null?void 0:v.docker_version),1),s("p",null,"API 版本： "+n((h=a(r))==null?void 0:h.api_version),1),s("p",null,"Min API 版本："+n((g=a(r))==null?void 0:g.min_api_version),1),s("p",null,"Git Commit: "+n((y=a(r))==null?void 0:y.git_commit),1),s("p",null,"Go version: "+n((S=a(r))==null?void 0:S.go_version),1),s("p",null,"OS: "+n((w=a(r))==null?void 0:w.os),1),s("p",null,"Arch: "+n((C=a(r))==null?void 0:C.arch),1)]}),_:1})]),_:1})]),_:1})])}}}),qe=de(Ie,[["__scopeId","data-v-d7835d4a"]]);export{qe as default};
