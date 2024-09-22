import{b as e,i as a,d as l,P as t,h as i,z as s,j as o,H as r,k as n,O as u,E as d,Q as c,_ as m,m as p}from"./el-button.Bh3a0xQF.js";import{v}from"./el-loading.D6_9u1EB.js";import{E as f}from"./el-drawer.is6L0oK4.js";import"./message.DctIvaHE.js";import{a as b,E as _}from"./el-form-item.BDwFQHfq.js";import"./el-tooltip.l0sNRNKZ.js";import{E as y}from"./el-popper.wQZhTUhG.js";/* empty css               */import{E as k,a as h}from"./el-select.CX85AUQb.js";import{E as V,a as g}from"./el-table-column.412yn3sf.js";import{E as w}from"./el-card.Dwv9l0RQ.js";import{q as x,a as j,c as I,d as C,u as U}from"./index.D4j8F93o.js";import{E as S}from"./index.1DI8xxX2.js";import{E}from"./index.D9jSBXYG.js";import{E as T}from"./index.Bjg9Z5Vi.js";import{q as z,d as A,c as B,r as q,B as N,o as P,a as $,k as K,f as O,u as R,n as F,a6 as H,b as L,w as X,h as G,i as M,t as Q,m as Y,j as Z,g as D,l as J,M as W,ac as ee,x as ae,X as le,F as te,A as ie,e as se,P as oe,_ as re}from"./index.Dr9aOSLV.js";import{i as ne}from"./validator.nxW8Vyz9.js";import{u as ue,d as de,t as ce}from"./index.BRGrvrSC.js";import{U as me,C as pe,I as ve}from"./event.BwRzfsZt.js";import{a as fe,b as be}from"./use-form-item.CQAx4y3m.js";import"./use-dialog.BrP1_dbF.js";import"./vnode._yyMxGkh.js";import"./scroll.BZMJFPW3.js";import"./focus-trap.ZmI0uCgZ.js";import"./castArray.DwF5sPJL.js";import"./_baseClone.CG0H6u9F.js";import"./_Uint8Array.DSxpQmA-.js";import"./_initCloneObject.u30XBX3X.js";import"./token.DWNpOE8r.js";import"./isEqual.qlKf3nEY.js";import"./index.nuX2YVwj.js";const _e=e({modelValue:{type:[Boolean,String,Number],default:!1},disabled:{type:Boolean,default:!1},loading:{type:Boolean,default:!1},size:{type:String,validator:ne},width:{type:[String,Number],default:""},inlinePrompt:{type:Boolean,default:!1},inactiveActionIcon:{type:a},activeActionIcon:{type:a},activeIcon:{type:a},inactiveIcon:{type:a},activeText:{type:String,default:""},inactiveText:{type:String,default:""},activeValue:{type:[Boolean,String,Number],default:!0},inactiveValue:{type:[Boolean,String,Number],default:!1},name:{type:String,default:""},validateEvent:{type:Boolean,default:!0},beforeChange:{type:l(Function)},id:String,tabindex:{type:[String,Number]},label:{type:String,default:void 0},...ue(["ariaLabel"])}),ye={[me]:e=>t(e)||z(e)||i(e),[pe]:e=>t(e)||z(e)||i(e),[ve]:e=>t(e)||z(e)||i(e)},ke=["onClick"],he=["id","aria-checked","aria-disabled","aria-label","name","true-value","false-value","disabled","tabindex","onKeydown"],Ve=["aria-hidden"],ge=["aria-hidden"],we=["aria-hidden"],xe="ElSwitch",je=A({name:xe});const Ie=p(m(A({...je,props:_e,emits:ye,setup(e,{expose:a,emit:l}){const i=e,{formItem:m}=fe(),p=s(),v=o("switch"),{inputId:f}=be(i,{formItemContext:m}),b=r(B((()=>i.loading))),_=q(!1!==i.modelValue),y=q(),k=q(),h=B((()=>[v.b(),v.m(p.value),v.is("disabled",b.value),v.is("checked",j.value)])),V=B((()=>[v.e("label"),v.em("label","left"),v.is("active",!j.value)])),g=B((()=>[v.e("label"),v.em("label","right"),v.is("active",j.value)])),w=B((()=>({width:n(i.width)})));N((()=>i.modelValue),(()=>{_.value=!0}));const x=B((()=>!!_.value&&i.modelValue)),j=B((()=>x.value===i.activeValue));[i.activeValue,i.inactiveValue].includes(x.value)||(l(me,i.inactiveValue),l(pe,i.inactiveValue),l(ve,i.inactiveValue)),N(j,(e=>{var a;y.value.checked=e,i.validateEvent&&(null==(a=null==m?void 0:m.validate)||a.call(m,"change").catch((e=>de())))}));const I=()=>{const e=j.value?i.inactiveValue:i.activeValue;l(me,e),l(pe,e),l(ve,e),W((()=>{y.value.checked=j.value}))},C=()=>{if(b.value)return;const{beforeChange:e}=i;if(!e)return void I();const a=e();[ee(a),t(a)].includes(!0)||ce(xe,"beforeChange must return type `Promise<boolean>` or `boolean`"),ee(a)?a.then((e=>{e&&I()})).catch((e=>{})):a&&I()};return P((()=>{y.value.checked=j.value})),u({from:"label",replacement:"aria-label",version:"2.8.0",scope:"el-switch",ref:"https://element-plus.org/en-US/component/switch.html"},B((()=>!!i.label))),a({focus:()=>{var e,a;null==(a=null==(e=y.value)?void 0:e.focus)||a.call(e)},checked:j}),(e,a)=>($(),K("div",{class:F(R(h)),onClick:J(C,["prevent"])},[O("input",{id:R(f),ref_key:"input",ref:y,class:F(R(v).e("input")),type:"checkbox",role:"switch","aria-checked":R(j),"aria-disabled":R(b),"aria-label":e.label||e.ariaLabel,name:e.name,"true-value":e.activeValue,"false-value":e.inactiveValue,disabled:R(b),tabindex:e.tabindex,onChange:I,onKeydown:H(C,["enter"])},null,42,he),e.inlinePrompt||!e.inactiveIcon&&!e.inactiveText?M("v-if",!0):($(),K("span",{key:0,class:F(R(V))},[e.inactiveIcon?($(),L(R(d),{key:0},{default:X((()=>[($(),L(G(e.inactiveIcon)))])),_:1})):M("v-if",!0),!e.inactiveIcon&&e.inactiveText?($(),K("span",{key:1,"aria-hidden":R(j)},Q(e.inactiveText),9,Ve)):M("v-if",!0)],2)),O("span",{ref_key:"core",ref:k,class:F(R(v).e("core")),style:D(R(w))},[e.inlinePrompt?($(),K("div",{key:0,class:F(R(v).e("inner"))},[e.activeIcon||e.inactiveIcon?($(),L(R(d),{key:0,class:F(R(v).is("icon"))},{default:X((()=>[($(),L(G(R(j)?e.activeIcon:e.inactiveIcon)))])),_:1},8,["class"])):e.activeText||e.inactiveText?($(),K("span",{key:1,class:F(R(v).is("text")),"aria-hidden":!R(j)},Q(R(j)?e.activeText:e.inactiveText),11,ge)):M("v-if",!0)],2)):M("v-if",!0),O("div",{class:F(R(v).e("action"))},[e.loading?($(),L(R(d),{key:0,class:F(R(v).is("loading"))},{default:X((()=>[Y(R(c))])),_:1},8,["class"])):R(j)?Z(e.$slots,"active-action",{key:1},(()=>[e.activeActionIcon?($(),L(R(d),{key:0},{default:X((()=>[($(),L(G(e.activeActionIcon)))])),_:1})):M("v-if",!0)])):R(j)?M("v-if",!0):Z(e.$slots,"inactive-action",{key:2},(()=>[e.inactiveActionIcon?($(),L(R(d),{key:0},{default:X((()=>[($(),L(G(e.inactiveActionIcon)))])),_:1})):M("v-if",!0)]))],2)],6),e.inlinePrompt||!e.activeIcon&&!e.activeText?M("v-if",!0):($(),K("span",{key:1,class:F(R(g))},[e.activeIcon?($(),L(R(d),{key:0},{default:X((()=>[($(),L(G(e.activeIcon)))])),_:1})):M("v-if",!0),!e.activeIcon&&e.activeText?($(),K("span",{key:1,"aria-hidden":!R(j)},Q(e.activeText),9,we)):M("v-if",!0)],2))],10,ke))}}),[["__file","switch.vue"]])),Ce={class:"am-account-header"},Ue={class:"am-user-operator"},Se={class:"am-table"},Ee={class:"am-user-create"},Te={class:"am-user-create__operator"},ze={class:"am-user-create"},Ae={class:"am-user-create__operator"},Be=re(A({__name:"index",setup(e){P((()=>{i(),r()}));const a=q(0),l=q([]),t=q(!1),i=async()=>{t.value=!0;const{data:e}=await x();l.value=e.data,a.value+=1,t.value=!1},s=e=>"admin"===e.username,o=q([]),r=async()=>{const{data:e}=await j();o.value=e.data},n=q(!1),u=q(),d=q({username:"",password:"",remark:"",role_ids:[],status:"1"}),c=ae({username:[{required:!0,message:"请输入用户名",trigger:"blur"}],password:[{required:!0,message:"请输入密码",trigger:"blur"}],role_ids:[{required:!0,message:"请选择角色",trigger:"blur"}]}),m=q(!1),p=q(!1),z=q(),A=q(!1),B=q({id:"",username:"",remark:"",role_ids:[],status:"1"}),N=q(!1);return(e,r)=>{const x=S,j=w,q=V,P=E,F=g,H=T,G=b,M=h,Z=k,D=Ie,J=y,W=_,ee=f,ae=v;return $(),K(te,null,[O("div",Ce,[O("span",{onClick:r[0]||(r[0]=a=>e.$router.push("/account/user"))},"用户"),O("span",{onClick:r[1]||(r[1]=a=>e.$router.push("/account/role"))},"角色"),O("span",{onClick:r[2]||(r[2]=a=>e.$router.push("/account/api"))},"接口")]),O("div",Ue,[Y(j,{shadow:"never"},{default:X((()=>[Y(x,{type:"primary",plain:"",onClick:r[3]||(r[3]=e=>n.value=!0)},{default:X((()=>[ie("新增用户")])),_:1})])),_:1})]),Y(j,{shadow:"never"},{default:X((()=>[O("div",Se,[se(($(),L(F,{data:R(l),height:"100%",key:R(a),stripe:"",ref:"multipleTable"},{default:X((()=>[Y(q,{prop:"username",label:"用户名","min-width":"120",align:"center"}),Y(q,{prop:"role",label:"角色","min-width":"160",align:"center"},{default:X((e=>[($(!0),K(te,null,oe(e.row.roles,((e,a)=>($(),L(P,{key:a},{default:X((()=>[ie(Q(e.name),1)])),_:2},1024)))),128))])),_:1}),Y(q,{prop:"status",label:"状态","min-width":"100",align:"center",sortable:""},{default:X((e=>[1===e.row.status?($(),L(P,{key:0,type:"success"},{default:X((()=>[ie("正常")])),_:1})):($(),L(P,{key:1,type:"danger"},{default:X((()=>[ie("禁用")])),_:1}))])),_:1}),Y(q,{prop:"created_at",label:"创建时间","min-width":"160",align:"center",sortable:""}),Y(q,{label:"操作",width:"200",fixed:"right",align:"center"},{default:X((e=>[Y(x,{type:"primary",size:"small",text:"",onClick:a=>{return l=e.row,A.value=!0,void(B.value={id:l.id,username:l.username,remark:l.remark,role_ids:l.roles.map((e=>e.id)),status:l.status.toString()});var l},disabled:s(e.row)},{default:X((()=>[ie(" 编辑 ")])),_:2},1032,["onClick","disabled"]),se(($(),L(x,{type:"danger",size:"small",text:"",onClick:a=>(async e=>{p.value=!0,C({ids:[e]}).finally((()=>{p.value=!1,i()}))})(e.row.id),disabled:s(e.row)},{default:X((()=>[ie(" 删除 ")])),_:2},1032,["onClick","disabled"])),[[ae,R(p)]])])),_:1})])),_:1},8,["data"])),[[ae,R(t)]])])])),_:1}),O("div",Ee,[Y(ee,{modelValue:R(n),"onUpdate:modelValue":r[11]||(r[11]=e=>le(n)?n.value=e:null),title:"创建用户",size:"50%"},{default:X((()=>[Y(W,{ref_key:"userCreateRef",ref:u,model:R(d),rules:R(c),"label-width":"120px","label-position":"left"},{default:X((()=>[Y(G,{label:"用户名",prop:"username"},{default:X((()=>[Y(H,{modelValue:R(d).username,"onUpdate:modelValue":r[4]||(r[4]=e=>R(d).username=e),placeholder:"请输入用户名"},null,8,["modelValue"])])),_:1}),Y(G,{label:"密码",prop:"password"},{default:X((()=>[Y(H,{modelValue:R(d).password,"onUpdate:modelValue":r[5]||(r[5]=e=>R(d).password=e),placeholder:"请输入密码"},null,8,["modelValue"])])),_:1}),Y(G,{label:"角色",prop:"role_ids"},{default:X((()=>[Y(Z,{modelValue:R(d).role_ids,"onUpdate:modelValue":r[6]||(r[6]=e=>R(d).role_ids=e),multiple:"",placeholder:"请选择角色"},{default:X((()=>[($(!0),K(te,null,oe(R(o),(e=>($(),L(M,{key:e.id,label:e.name,value:e.id},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),Y(G,{label:"备注",prop:"remark"},{default:X((()=>[Y(H,{modelValue:R(d).remark,"onUpdate:modelValue":r[7]||(r[7]=e=>R(d).remark=e),placeholder:"请输入备注"},null,8,["modelValue"])])),_:1}),Y(G,{label:"状态",prop:"status"},{default:X((()=>[Y(J,{content:"用户状态，1为正常，2为禁用",placement:"top"},{default:X((()=>[Y(D,{modelValue:R(d).status,"onUpdate:modelValue":r[8]||(r[8]=e=>R(d).status=e),"active-value":"1","inactive-value":"2"},null,8,["modelValue"])])),_:1})])),_:1})])),_:1},8,["model","rules"]),O("div",Te,[Y(x,{type:"default",size:"default",plain:"",onClick:r[9]||(r[9]=e=>n.value=!1)},{default:X((()=>[ie("取消")])),_:1}),se(($(),L(x,{type:"primary",size:"default",plain:"",onClick:r[10]||(r[10]=e=>(async e=>{e&&await e.validate((e=>{if(e){m.value=!0;const e={username:d.value.username,password:d.value.password,remark:d.value.remark,role_ids:d.value.role_ids,status:Number(d.value.status)};I(e).finally((()=>{m.value=!1,n.value=!1,i()}))}}))})(R(u)))},{default:X((()=>[ie(" 确定 ")])),_:1})),[[ae,R(m)]])])])),_:1},8,["modelValue"])]),O("div",ze,[Y(ee,{modelValue:R(A),"onUpdate:modelValue":r[18]||(r[18]=e=>le(A)?A.value=e:null),title:"编辑用户",size:"50%"},{default:X((()=>[Y(W,{ref_key:"userUpdateRef",ref:z,model:R(B),rules:R(c),"label-width":"120px","label-position":"left"},{default:X((()=>[Y(G,{label:"用户名",prop:"username"},{default:X((()=>[Y(H,{modelValue:R(B).username,"onUpdate:modelValue":r[12]||(r[12]=e=>R(B).username=e),placeholder:"请输入用户名"},null,8,["modelValue"])])),_:1}),Y(G,{label:"角色",prop:"role_ids"},{default:X((()=>[Y(Z,{modelValue:R(B).role_ids,"onUpdate:modelValue":r[13]||(r[13]=e=>R(B).role_ids=e),multiple:"",placeholder:"请选择角色"},{default:X((()=>[($(!0),K(te,null,oe(R(o),(e=>($(),L(M,{key:e.id,label:e.name,value:e.id},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),Y(G,{label:"备注",prop:"remark"},{default:X((()=>[Y(H,{modelValue:R(B).remark,"onUpdate:modelValue":r[14]||(r[14]=e=>R(B).remark=e),placeholder:"请输入备注"},null,8,["modelValue"])])),_:1}),Y(G,{label:"状态",prop:"status"},{default:X((()=>[Y(J,{content:"用户状态，1为正常，2为禁用",placement:"top"},{default:X((()=>[Y(D,{modelValue:R(B).status,"onUpdate:modelValue":r[15]||(r[15]=e=>R(B).status=e),"active-value":"1","inactive-value":"2"},null,8,["modelValue"])])),_:1})])),_:1})])),_:1},8,["model","rules"]),O("div",Ae,[Y(x,{type:"default",size:"default",plain:"",onClick:r[16]||(r[16]=e=>A.value=!1)},{default:X((()=>[ie("取消")])),_:1}),se(($(),L(x,{type:"primary",size:"default",plain:"",onClick:r[17]||(r[17]=e=>(async e=>{e&&await e.validate((e=>{if(e){N.value=!0;const e={id:B.value.id,username:B.value.username,remark:B.value.remark,role_ids:B.value.role_ids,status:Number(B.value.status)};U(e).finally((()=>{N.value=!1,A.value=!1,i()}))}}))})(R(z)))},{default:X((()=>[ie(" 确定 ")])),_:1})),[[ae,R(N)]])])])),_:1},8,["modelValue"])])],64)}}}),[["__scopeId","data-v-6fd47c10"]]);export{Be as default};
