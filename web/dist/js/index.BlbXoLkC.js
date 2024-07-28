import{b as e,i as s,d as t,u as o,T as a,a as n,E as l,_ as r,C as i,c as u,e as p,f as c,g as d,w as m,h as f,l as g,j as v}from"./message.CS_TSLeJ.js";import{E as y,a as h}from"./el-form-item.THx8gv2Q.js";import{d as b,r as _,c as x,o as C,a as w,b as k,w as j,e as E,f as T,n as I,u as S,g as z,h as A,i as H,t as L,j as U,k as q,F as B,v as M,l as V,m as D,T as N,p as O,q as $,s as F,x as P,y as W,z as Z,A as G,B as J,C as K,_ as Q}from"./index.DuNUv9gK.js";import{_ as R,l as X}from"./index.CXwM-Ot8.js";import{E as Y}from"./index.Cb0Nm553.js";import{E as ee}from"./index.Ck8DkNjZ.js";import"./castArray.CZ4MN7ta.js";import"./use-form-item.CZ2M-1KL.js";import"./_baseClone.D8XUM9zm.js";import"./_Uint8Array.Bq4iIfDz.js";import"./_initCloneObject.CpUBN0YX.js";import"./index.D8XJyr-Y.js";const se=["success","info","warning","error"],te=e({customClass:{type:String,default:""},dangerouslyUseHTMLString:{type:Boolean,default:!1},duration:{type:Number,default:4500},icon:{type:s},id:{type:String,default:""},message:{type:t([String,Object]),default:""},offset:{type:Number,default:0},onClick:{type:t(Function),default:()=>{}},onClose:{type:t(Function),required:!0},position:{type:String,values:["top-right","top-left","bottom-right","bottom-left"],default:"top-right"},showClose:{type:Boolean,default:!0},title:{type:String,default:""},type:{type:String,values:[...se,""],default:""},zIndex:Number}),oe=["id"],ae=["textContent"],ne={key:0},le=["innerHTML"],re=b({name:"ElNotification"});var ie=r(b({...re,props:te,emits:{destroy:()=>!0},setup(e,{expose:s}){const t=e,{ns:r,zIndex:c}=o("notification"),{nextZIndex:d,currentZIndex:m}=c,{Close:f}=i,g=_(!1);let v;const y=x((()=>{const e=t.type;return e&&a[t.type]?r.m(e):""})),h=x((()=>t.type&&a[t.type]||t.icon)),b=x((()=>t.position.endsWith("right")?"right":"left")),O=x((()=>t.position.startsWith("top")?"top":"bottom")),$=x((()=>{var e;return{[O.value]:`${t.offset}px`,zIndex:null!=(e=t.zIndex)?e:m.value}}));function F(){t.duration>0&&({stop:v}=u((()=>{g.value&&W()}),t.duration))}function P(){null==v||v()}function W(){g.value=!1}return C((()=>{F(),d(),g.value=!0})),n(document,"keydown",(function({code:e}){e===p.delete||e===p.backspace?P():e===p.esc?g.value&&W():F()})),s({visible:g,close:W}),(e,s)=>(w(),k(N,{name:S(r).b("fade"),onBeforeLeave:e.onClose,onAfterLeave:s[1]||(s[1]=s=>e.$emit("destroy")),persisted:""},{default:j((()=>[E(T("div",{id:e.id,class:I([S(r).b(),e.customClass,S(b)]),style:z(S($)),role:"alert",onMouseenter:P,onMouseleave:F,onClick:s[0]||(s[0]=(...s)=>e.onClick&&e.onClick(...s))},[S(h)?(w(),k(S(l),{key:0,class:I([S(r).e("icon"),S(y)])},{default:j((()=>[(w(),k(A(S(h))))])),_:1},8,["class"])):H("v-if",!0),T("div",{class:I(S(r).e("group"))},[T("h2",{class:I(S(r).e("title")),textContent:L(e.title)},null,10,ae),E(T("div",{class:I(S(r).e("content")),style:z(e.title?void 0:{margin:0})},[U(e.$slots,"default",{},(()=>[e.dangerouslyUseHTMLString?(w(),q(B,{key:1},[H(" Caution here, message could've been compromised, never use user's input as message "),T("p",{innerHTML:e.message},null,8,le)],2112)):(w(),q("p",ne,L(e.message),1))]))],6),[[M,e.message]]),e.showClose?(w(),k(S(l),{key:0,class:I(S(r).e("closeBtn")),onClick:V(W,["stop"])},{default:j((()=>[D(S(f))])),_:1},8,["class","onClick"])):H("v-if",!0)],2)],46,oe),[[M,g.value]])])),_:3},8,["name","onBeforeLeave"]))}}),[["__file","notification.vue"]]);const ue={"top-left":[],"top-right":[],"bottom-left":[],"bottom-right":[]};let pe=1;const ce=function(e={},s=null){if(!c)return{close:()=>{}};("string"==typeof e||O(e))&&(e={message:e});const t=e.position||"top-right";let o=e.offset||0;ue[t].forEach((({vm:e})=>{var s;o+=((null==(s=e.el)?void 0:s.offsetHeight)||0)+16})),o+=16;const a="notification_"+pe++,n=e.onClose,l={...e,offset:o,id:a,onClose:()=>{!function(e,s,t){const o=ue[s],a=o.findIndex((({vm:s})=>{var t;return(null==(t=s.component)?void 0:t.props.id)===e}));if(-1===a)return;const{vm:n}=o[a];if(!n)return;null==t||t(n);const l=n.el.offsetHeight,r=s.split("-")[0];o.splice(a,1);const i=o.length;if(i<1)return;for(let u=a;u<i;u++){const{el:e,component:s}=o[u].vm,t=Number.parseInt(e.style[r],10)-l-16;s.props.offset=t}}(a,t,n)}};let r=document.body;d(e.appendTo)?r=e.appendTo:$(e.appendTo)&&(r=document.querySelector(e.appendTo)),d(r)||(r=document.body);const i=document.createElement("div"),u=D(ie,l,O(l.message)?{default:()=>l.message}:null);return u.appContext=null!=s?s:ce._context,u.props.onDestroy=()=>{F(null,i)},F(u,i),ue[t].push({vm:u}),r.appendChild(i.firstElementChild),{close:()=>{u.component.exposed.visible.value=!1}}};se.forEach((e=>{ce[e]=(s={})=>(("string"==typeof s||O(s))&&(s={message:s}),ce({...s,type:e}))})),ce.closeAll=function(){for(const e of Object.values(ue))e.forEach((({vm:e})=>{e.component.exposed.visible.value=!1}))},ce._context=null;const de=m(ce,"$notify");function me(){const e=(new Date).getHours();return e>=6&&e<=10?"早上好 ⛅":e>=10&&e<=14?"中午好 🌞":e>=14&&e<=18?"下午好 🌞":e>=18&&e<=24?"晚上好 🌛":e>=0&&e<=6?"凌晨好 🌛":void 0}const fe=e=>(J("data-v-58a93363"),e=e(),K(),e),ge={class:"am-login"},ve=fe((()=>T("div",{class:"am-login-left"},[T("div",{class:"am-login-left__title"},[T("img",{class:"am-login-left__img",src:R,alt:""}),T("h1",null,"Amprobe")]),T("h3",null,"Amprobe 是一款轻量级主机及 Docker 容器监控工具，它可以轻松的帮助我们完成以下几方面的工作:"),T("div",{class:"am-login-left__item"},[T("span",null,"· 监控主机的 CPU、内存、磁盘 IO、网络 IO情况"),T("span",null,"· 监控部署于主机上 Docker 容器的运行状态、CPU、内存使用情况"),T("span",null,"· 实时查看 Docker 容器的日志，并支持日志下载")])],-1))),ye={class:"am-login-right"},he={class:"am-login-right__form"},be=fe((()=>T("div",{class:"title"},[T("span",null,"登录")],-1))),_e=Q(b({__name:"index",setup(e){const s=P({username:"amprobe",password:"123456"}),t={username:[{required:!0,trigger:"blur"}],password:[{required:!0,trigger:"blur",validator:function(e,s,t){""===s?t(new Error("password is required")):s.length<6?t(new Error("The password can not be less than 6 digits")):t()}}]};const o=W(),a=Z(),n=async()=>{try{const{data:e}=await X({...s});o.user.setToken(e.access_token,e.refresh_token),await a.replace("/"),de({title:"欢迎来到 Amprobe",message:me(),type:"success"})}catch(e){e instanceof Error&&v.error(e.message)}};return(e,o)=>{const a=Y,l=h,r=ee,i=y;return w(),q("div",ge,[ve,T("div",ye,[T("div",he,[D(i,{model:S(s),rules:t},{default:j((()=>[be,D(l,{prop:"username"},{default:j((()=>[D(a,{ref:"username",modelValue:S(s).username,"onUpdate:modelValue":o[0]||(o[0]=e=>S(s).username=e),size:"large",class:"username-input",placeholder:"请输入用户名",name:"username",autocomplete:"on","prefix-icon":S(f)},null,8,["modelValue","prefix-icon"])])),_:1}),D(l,{prop:"password"},{default:j((()=>[D(a,{modelValue:S(s).password,"onUpdate:modelValue":o[1]||(o[1]=e=>S(s).password=e),size:"large",type:"password",class:"password-input",placeholder:"请输入密码",name:"password","prefix-icon":S(g),"show-password":!0},null,8,["modelValue","prefix-icon"])])),_:1}),D(r,{class:"btn",size:"large",type:"primary",onClick:V(n,["prevent"])},{default:j((()=>[G(" 登录 ")])),_:1})])),_:1},8,["model"])])])])}}}),[["__scopeId","data-v-58a93363"]]);export{_e as default};
