import"./el-button.wSr2vfOs.js";import{E as e,a,v as t}from"./el-table-column.CRjwOlSm.js";import{E as l}from"./el-progress.DD9jSVW7.js";import{E as r}from"./el-dialog.QrPPJRCD.js";import{e as i,s as o}from"./message.bOUkP4cz.js";import"./el-tooltip.l0sNRNKZ.js";import"./el-popper.B5lc-jMO.js";/* empty css               */import{E as n}from"./el-card.qxUhNo9k.js";import{E as s,a as d,b as u}from"./el-dropdown-item.BwzSe7Kc.js";import{_ as c}from"./index.vFwLIVie.js";import{h as f,i as m,j as p,k as h,l as v}from"./index.BqDuBUYv.js";import{d as g,y as _,r as y,o as j,k as C,f as b,m as w,w as k,u as x,Y as A,F as V,a as z,A as E,e as O,b as U,i as T,t as Y,S as I,U as $,_ as D}from"./index.C-XteKAQ.js";import{c as M}from"./convert.B8xZ5nZn.js";import{g as R,d as q}from"./dayjs.min.BglOJP2w.js";import{E as F,a as H}from"./index.BusDn2qi.js";import{E as J}from"./index.W1JDYdlb.js";import"./index.HJW_Xvjo.js";import"./isEqual.Bn-GIEBR.js";import"./_Uint8Array.DhThSVct.js";import"./use-form-item.BHwM85F6.js";import"./_initCloneObject.A5AD6tfW.js";import"./_baseClone.BXZGK5jP.js";import"./use-dialog.0QHuDEHA.js";import"./vnode.DQET2EQQ.js";import"./scroll.COMsDKDo.js";import"./focus-trap.BwkBla3c.js";import"./refs.D6GIAKzc.js";import"./castArray.DN6WFAt4.js";import"./index.Bs0cFoC_.js";function L(e){if("string"!=typeof e)throw new TypeError("Path must be a string. Received "+JSON.stringify(e))}function N(e,a){for(var t,l="",r=0,i=-1,o=0,n=0;n<=e.length;++n){if(n<e.length)t=e.charCodeAt(n);else{if(47===t)break;t=47}if(47===t){if(i===n-1||1===o);else if(i!==n-1&&2===o){if(l.length<2||2!==r||46!==l.charCodeAt(l.length-1)||46!==l.charCodeAt(l.length-2))if(l.length>2){var s=l.lastIndexOf("/");if(s!==l.length-1){-1===s?(l="",r=0):r=(l=l.slice(0,s)).length-1-l.lastIndexOf("/"),i=n,o=0;continue}}else if(2===l.length||1===l.length){l="",r=0,i=n,o=0;continue}a&&(l.length>0?l+="/..":l="..",r=2)}else l.length>0?l+="/"+e.slice(i+1,n):l=e.slice(i+1,n),r=n-i-1;i=n,o=0}else 46===t&&-1!==o?++o:o=-1}return l}var P={resolve:function(){for(var e,a="",t=!1,l=arguments.length-1;l>=-1&&!t;l--){var r;l>=0?r=arguments[l]:(void 0===e&&(e=process.cwd()),r=e),L(r),0!==r.length&&(a=r+"/"+a,t=47===r.charCodeAt(0))}return a=N(a,!t),t?a.length>0?"/"+a:"/":a.length>0?a:"."},normalize:function(e){if(L(e),0===e.length)return".";var a=47===e.charCodeAt(0),t=47===e.charCodeAt(e.length-1);return 0!==(e=N(e,!a)).length||a||(e="."),e.length>0&&t&&(e+="/"),a?"/"+e:e},isAbsolute:function(e){return L(e),e.length>0&&47===e.charCodeAt(0)},join:function(){if(0===arguments.length)return".";for(var e,a=0;a<arguments.length;++a){var t=arguments[a];L(t),t.length>0&&(void 0===e?e=t:e+="/"+t)}return void 0===e?".":P.normalize(e)},relative:function(e,a){if(L(e),L(a),e===a)return"";if((e=P.resolve(e))===(a=P.resolve(a)))return"";for(var t=1;t<e.length&&47===e.charCodeAt(t);++t);for(var l=e.length,r=l-t,i=1;i<a.length&&47===a.charCodeAt(i);++i);for(var o=a.length-i,n=r<o?r:o,s=-1,d=0;d<=n;++d){if(d===n){if(o>n){if(47===a.charCodeAt(i+d))return a.slice(i+d+1);if(0===d)return a.slice(i+d)}else r>n&&(47===e.charCodeAt(t+d)?s=d:0===d&&(s=0));break}var u=e.charCodeAt(t+d);if(u!==a.charCodeAt(i+d))break;47===u&&(s=d)}var c="";for(d=t+s+1;d<=l;++d)d!==l&&47!==e.charCodeAt(d)||(0===c.length?c+="..":c+="/..");return c.length>0?c+a.slice(i+s):(i+=s,47===a.charCodeAt(i)&&++i,a.slice(i))},_makeLong:function(e){return e},dirname:function(e){if(L(e),0===e.length)return".";for(var a=e.charCodeAt(0),t=47===a,l=-1,r=!0,i=e.length-1;i>=1;--i)if(47===(a=e.charCodeAt(i))){if(!r){l=i;break}}else r=!1;return-1===l?t?"/":".":t&&1===l?"//":e.slice(0,l)},basename:function(e,a){if(void 0!==a&&"string"!=typeof a)throw new TypeError('"ext" argument must be a string');L(e);var t,l=0,r=-1,i=!0;if(void 0!==a&&a.length>0&&a.length<=e.length){if(a.length===e.length&&a===e)return"";var o=a.length-1,n=-1;for(t=e.length-1;t>=0;--t){var s=e.charCodeAt(t);if(47===s){if(!i){l=t+1;break}}else-1===n&&(i=!1,n=t+1),o>=0&&(s===a.charCodeAt(o)?-1==--o&&(r=t):(o=-1,r=n))}return l===r?r=n:-1===r&&(r=e.length),e.slice(l,r)}for(t=e.length-1;t>=0;--t)if(47===e.charCodeAt(t)){if(!i){l=t+1;break}}else-1===r&&(i=!1,r=t+1);return-1===r?"":e.slice(l,r)},extname:function(e){L(e);for(var a=-1,t=0,l=-1,r=!0,i=0,o=e.length-1;o>=0;--o){var n=e.charCodeAt(o);if(47!==n)-1===l&&(r=!1,l=o+1),46===n?-1===a?a=o:1!==i&&(i=1):-1!==a&&(i=-1);else if(!r){t=o+1;break}}return-1===a||-1===l||0===i||1===i&&a===l-1&&a===t+1?"":e.slice(a,l)},format:function(e){if(null===e||"object"!=typeof e)throw new TypeError('The "pathObject" argument must be of type Object. Received type '+typeof e);return function(e,a){var t=a.dir||a.root,l=a.base||(a.name||"")+(a.ext||"");return t?t===a.root?t+l:t+e+l:l}("/",e)},parse:function(e){L(e);var a={root:"",dir:"",base:"",ext:"",name:""};if(0===e.length)return a;var t,l=e.charCodeAt(0),r=47===l;r?(a.root="/",t=1):t=0;for(var i=-1,o=0,n=-1,s=!0,d=e.length-1,u=0;d>=t;--d)if(47!==(l=e.charCodeAt(d)))-1===n&&(s=!1,n=d+1),46===l?-1===i?i=d:1!==u&&(u=1):-1!==i&&(u=-1);else if(!s){o=d+1;break}return-1===i||-1===n||0===u||1===u&&i===n-1&&i===o+1?-1!==n&&(a.base=a.name=0===o&&r?e.slice(1,n):e.slice(o,n)):(0===o&&r?(a.name=e.slice(1,i),a.base=e.slice(1,n)):(a.name=e.slice(o,i),a.base=e.slice(o,n)),a.ext=e.slice(i,n)),o>0?a.dir=e.slice(0,o-1):r&&(a.dir="/"),a},sep:"/",delimiter:":",win32:null,posix:null};P.posix=P;const S=R(P),B={class:"am-host-header"},K={class:"am-host-operator"},G={class:"am-table"},Q={style:{display:"flex","align-items":"center"}},W=["onClick"],X={class:"am-host-file-create"},Z={class:"am-host-file-create"},ee={class:"am-host-file-upload"},ae=(e=>(I("data-v-70569ddc"),e=e(),$(),e))((()=>b("div",{class:"el-upload__text"},[E("Drop file here or "),b("em",null,"click to upload")],-1))),te=D(g({__name:"index",setup(g){const I=_(),$=y(!1),D=y(""),R=y(""),L=y([]);j((()=>{te(D.value,!0)}));const N=()=>{if("/"==D.value)return;R.value=D.value;let e=D.value.split("/");e.pop(),te(e.join("/"),!0)},P=()=>{"/"!=D.value&&te(R.value,!0)},te=async(e,a)=>{if(!a)return void i("该路径不是文件夹");""==e&&(e="/"),$.value=!0,D.value=e;let t={path:e};const{data:l}=await f(t);L.value=l.files,L.value.sort(((e,a)=>e.is_dir&&!a.is_dir?-1:!e.is_dir&&a.is_dir?1:0)),$.value=!1},le=y(0),re=y(!1),ie=()=>{o("上传成功"),re.value=!1,te(D.value,!0)},oe=y(!1),ne=y(!1),se=y(""),de=async()=>{const e={path:D.value,folder_name:se.value},{data:a}=await h(e);o("创建成功"),ne.value=!1,te(D.value,!0)},ue=y(!1),ce=y(""),fe=async()=>{const e={path:D.value,file_name:ce.value},{data:a}=await v(e);o("创建成功"),ue.value=!1,te(D.value,!0)};return(i,f)=>{const h=c,v=F,g=H,_=s,y=d,j=u,R=n,me=e,pe=a,he=J,ve=r,ge=l,_e=t;return z(),C(V,null,[b("div",B,[b("span",{onClick:f[0]||(f[0]=e=>i.$router.push("/host/monitor"))},"监控"),b("span",{onClick:f[1]||(f[1]=e=>i.$router.push("/host/file"))},"文件"),b("span",{onClick:f[2]||(f[2]=e=>i.$router.push("/host/settings"))},"设置")]),b("div",K,[w(R,{shadow:"never"},{default:k((()=>[w(g,{class:"ml-4",style:{"margin-right":"16px"}},{default:k((()=>[w(v,{type:"primary",onClick:N},{default:k((()=>[w(h,{"icon-class":"back"})])),_:1}),w(v,{type:"primary",onClick:P},{default:k((()=>[w(h,{"icon-class":"next"})])),_:1})])),_:1}),w(j,null,{dropdown:k((()=>[w(y,null,{default:k((()=>[w(_,null,{default:k((()=>[w(v,{type:"primary",size:"small",text:"",onClick:f[3]||(f[3]=e=>ne.value=!0)},{default:k((()=>[w(h,{"icon-class":"folder",style:{color:"#105eeb","margin-right":"4px"}}),E(" 文件夹 ")])),_:1})])),_:1}),w(_,null,{default:k((()=>[w(v,{type:"primary",size:"small",text:"",onClick:f[4]||(f[4]=e=>ue.value=!0)},{default:k((()=>[E(" 文件 ")])),_:1})])),_:1})])),_:1})])),default:k((()=>[w(v,{type:"primary",plain:""},{default:k((()=>[E(" 新建 "),w(h,{"icon-class":"down"})])),_:1})])),_:1}),w(v,{type:"primary",plain:"",onClick:f[5]||(f[5]=e=>re.value=!0)},{default:k((()=>[E(" 上传 ")])),_:1})])),_:1})]),w(R,{shadow:"never"},{default:k((()=>[b("div",G,[O((z(),U(pe,{data:x(L),key:x(le),stripe:"",ref:"multipleTable"},{default:k((()=>[w(me,{type:"selection",width:"55"}),w(me,{prop:"name",label:"名称","min-width":"150",align:"center",fixed:""},{default:k((e=>[b("div",Q,[1==e.row.is_dir?(z(),U(h,{key:0,"icon-class":"folder",style:{color:"#105eeb","margin-right":"4px"}})):T("",!0),b("span",{onClick:a=>te(x(S).join(x(D),e.row.name),e.row.is_dir)},Y(e.row.name),9,W)])])),_:1}),w(me,{prop:"size",label:"大小",align:"center","min-width":"100",sortable:""},{default:k((e=>[b("span",null,Y(x(M)(e.row.size)),1)])),_:1}),w(me,{prop:"mode",label:"权限",align:"center","min-width":"100"}),w(me,{prop:"mod_time",label:"修改时间",align:"center","min-width":"100"},{default:k((e=>[b("span",null,Y(x(q)(1e3*e.row.mod_time).format("YYYY-MM-DD HH:mm:ss")),1)])),_:1}),w(me,{label:"操作",width:"200",fixed:"right",align:"center"},{default:k((e=>[w(v,{type:"danger",size:"small",text:"",onClick:a=>(async e=>{const a={filepath:S.join(D.value,e)},{data:t}=await m(a);o("删除成功"),te(D.value,!0)})(e.row.name)},{default:k((()=>[E(" 删除 ")])),_:2},1032,["onClick"]),w(v,{type:"primary",size:"small",text:"",onClick:a=>(async e=>{const a={filepath:S.join(D.value,e)},{data:t}=await p(a);o("下载成功")})(e.row.name)},{default:k((()=>[E(" 下载 ")])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"])),[[_e,x($)]])])])),_:1}),b("div",X,[w(ve,{modelValue:x(ue),"onUpdate:modelValue":f[7]||(f[7]=e=>A(ue)?ue.value=e:null),title:"新建文件",width:"50%"},{default:k((()=>[w(he,{modelValue:x(ce),"onUpdate:modelValue":f[6]||(f[6]=e=>A(ce)?ce.value=e:null),placeholder:"请输入文件名"},null,8,["modelValue"]),O((z(),U(v,{size:"default",type:"info",plain:"",onClick:fe},{default:k((()=>[E(" 确定 ")])),_:1})),[[_e,x(oe)]])])),_:1},8,["modelValue"])]),b("div",Z,[w(ve,{modelValue:x(ne),"onUpdate:modelValue":f[9]||(f[9]=e=>A(ne)?ne.value=e:null),title:"新建文件夹",width:"50%"},{default:k((()=>[w(he,{modelValue:x(se),"onUpdate:modelValue":f[8]||(f[8]=e=>A(se)?se.value=e:null),placeholder:"请输入文件夹名称"},null,8,["modelValue"]),O((z(),U(v,{size:"default",type:"info",plain:"",onClick:de},{default:k((()=>[E(" 确定 ")])),_:1})),[[_e,x(oe)]])])),_:1},8,["modelValue"])]),b("div",ee,[w(ve,{modelValue:x(re),"onUpdate:modelValue":f[10]||(f[10]=e=>A(re)?re.value=e:null),title:"上传文件",width:"50%"},{default:k((()=>[w(ge,{ref:"uploadRef",drag:"",action:"/app/api/v1/host/file_upload",limit:1,data:{prefix:x(D)},multiple:"",headers:{Authorization:`Bearer ${x(I).user.token}`},"on-success":ie},{default:k((()=>[w(h,{"icon-class":"upload"}),ae])),_:1},8,["data","headers"])])),_:1},8,["modelValue"])])],64)}}}),[["__scopeId","data-v-70569ddc"]]);export{te as default};