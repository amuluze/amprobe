import{F as e,G as a}from"./message.CS_TSLeJ.js";import{E as t,a as l,v as o}from"./el-table-column.CUrmI2vn.js";import{E as r}from"./el-drawer.Q6nFi8Na.js";import{a as n,E as s}from"./el-form-item.THx8gv2Q.js";/* empty css               */import{E as i,a as d}from"./el-select.D4WZ4Qh9.js";import"./el-popper.GbrdInRL.js";import{u as p,z as u,E as m,a as w}from"./zh-cn.PWYwJho2.js";import"./el-tooltip.l0sNRNKZ.js";import{E as c}from"./el-card.CS5WT6tL.js";import{c as k,i as g,j as f}from"./index.9iOzwdYR.js";import{E as b}from"./index.Ck8DkNjZ.js";import{E as h}from"./index.Cb0Nm553.js";import{d as v,o as _,r as j,x as y,k as C,f as V,m as z,w as x,u as E,Y as S,F as U,a as G,A as M,e as N,b as q,S as P,_ as L}from"./index.DuNUv9gK.js";import"./use-form-item.CZ2M-1KL.js";import"./isEqual.C8I6FPlV.js";import"./_Uint8Array.Bq4iIfDz.js";import"./_initCloneObject.CpUBN0YX.js";import"./use-dialog.CZCXc5Xp.js";import"./vnode.BrXPxjYI.js";import"./scroll.15sbxcXA.js";import"./focus-trap.D1Nefh4a.js";import"./castArray.CZ4MN7ta.js";import"./_baseClone.D8XUM9zm.js";import"./index.BusWp8wr.js";import"./index.D8XJyr-Y.js";const A={class:"am-network-title"},$={class:"am-network-operator"},F={class:"am-table"},O={class:"am-pagination"},I={class:"am-network-create"},J={class:"am-network-create__content"},K={class:"am-network-create__operator"},Q=L(v({__name:"index",setup(v){_((()=>{R()}));const L=j(0),{data:Q,refresh:R,loading:Y,pagination:B}=p(k,{},{}),D=y({networkName:"",networkMode:"bridge",networkSubnet:"",networkGateway:"",networkLabels:""}),H=j(!1),T=j(),W=y({networkName:[{required:!0,message:"Please input network name",trigger:"blur"}],networkMode:[{required:!0,message:"Please select network mode",trigger:"blur"}],networkSubnet:[{required:!0,message:"Please input network subnet",trigger:"blur"}],networkGateway:[{required:!0,message:"Please input network gateway",trigger:"blur"}]}),X=[{value:"bridge",label:"bridge"},{value:"host",label:"host"}],Z=j(!1);return(p,k)=>{const v=b,_=c,j=t,y=l,ee=m,ae=w,te=h,le=n,oe=d,re=i,ne=s,se=r,ie=o;return G(),C(U,null,[V("div",A,[V("span",{onClick:k[0]||(k[0]=e=>p.$router.push("/docker/container"))},"容器"),V("span",{onClick:k[1]||(k[1]=e=>p.$router.push("/docker/image"))},"镜像"),V("span",{onClick:k[2]||(k[2]=e=>p.$router.push("/docker/network"))},"网络"),V("span",{onClick:k[3]||(k[3]=e=>p.$router.push("/docker/settings"))},"设置")]),V("div",$,[z(_,{shadow:"never"},{default:x((()=>[z(v,{type:"primary",plain:"",onClick:k[4]||(k[4]=e=>H.value=!0)},{default:x((()=>[M("创建网络")])),_:1})])),_:1})]),z(_,{shadow:"never"},{default:x((()=>[V("div",F,[N((G(),q(y,{data:E(Q),key:E(L),"highlight-current-row":"",height:"100%",stripe:""},{default:x((()=>[z(j,{prop:"id",label:"名称",align:"center",width:"120",fixed:""}),z(j,{prop:"name",label:"模式",align:"center","min-width":"120"}),z(j,{prop:"subnet",label:"子网",align:"center","show-overflow-tooltip":"",width:"120"}),z(j,{prop:"gateway",label:"网关",align:"center","show-overflow-tooltip":"",width:"120"}),z(j,{prop:"created",label:"创建时间",align:"center",width:"200"}),z(j,{label:"操作",width:"160",fixed:"right",align:"center"},{default:x((e=>[z(v,{type:"danger",plain:"",size:"small",onClick:t=>(async e=>{let t={network_id:e};const{data:l}=await f(t);a("删除成功"),R()})(e.row.id)},{default:x((()=>[M(" 删除 ")])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"])),[[ie,E(Y)]])]),V("div",O,[z(ae,{locale:E(u)},{default:x((()=>[z(ee,{"current-page":E(B).page,"onUpdate:currentPage":k[5]||(k[5]=e=>E(B).page=e),"page-size":E(B).size,layout:"total, sizes, prev, pager, next, jumper","page-sizes":E(B).sizeOption,total:E(B).total,onSizeChange:E(B).onSizeChange,onCurrentChange:E(B).onPageChange},null,8,["current-page","page-size","page-sizes","total","onSizeChange","onCurrentChange"])])),_:1},8,["locale"])])])),_:1}),V("div",I,[z(se,{modelValue:E(H),"onUpdate:modelValue":k[13]||(k[13]=e=>S(H)?H.value=e:null),size:"540",title:"创建网络"},{default:x((()=>[V("div",J,[z(ne,{ref_key:"networkCreateRef",ref:T,model:E(D),rules:E(W),"label-width":"100px","label-position":"left"},{default:x((()=>[z(le,{label:"名称",prop:"networkName"},{default:x((()=>[z(te,{modelValue:E(D).networkName,"onUpdate:modelValue":k[6]||(k[6]=e=>E(D).networkName=e),placeholder:"请输入名称"},null,8,["modelValue"])])),_:1}),z(le,{label:"模式",prop:"networkMode"},{default:x((()=>[z(re,{modelValue:E(D).networkMode,"onUpdate:modelValue":k[7]||(k[7]=e=>E(D).networkMode=e),style:{width:"240px"},placeholder:"请选择模式"},{default:x((()=>[(G(),C(U,null,P(X,(e=>z(oe,{key:e.value,label:e.label,value:e.value},null,8,["label","value"]))),64))])),_:1},8,["modelValue"])])),_:1}),z(le,{label:"子网",prop:"networkSubnet"},{default:x((()=>[z(te,{modelValue:E(D).networkSubnet,"onUpdate:modelValue":k[8]||(k[8]=e=>E(D).networkSubnet=e),placeholder:"172.16.10.0/24"},null,8,["modelValue"])])),_:1}),z(le,{label:"网关",prop:"networkGateway"},{default:x((()=>[z(te,{modelValue:E(D).networkGateway,"onUpdate:modelValue":k[9]||(k[9]=e=>E(D).networkGateway=e),placeholder:"172.16.10.1"},null,8,["modelValue"])])),_:1}),z(le,{label:"标签",prop:"networkLabels"},{default:x((()=>[z(te,{modelValue:E(D).networkLabels,"onUpdate:modelValue":k[10]||(k[10]=e=>E(D).networkLabels=e),type:"textarea",rows:4,placeholder:"一行一个，例如:\nkey1=value1\nkey2=value2"},null,8,["modelValue"])])),_:1})])),_:1},8,["model","rules"])]),V("div",K,[z(v,{type:"default",size:"default",plain:"",onClick:k[11]||(k[11]=e=>H.value=!1)},{default:x((()=>[M("取消")])),_:1}),N((G(),q(v,{type:"primary",size:"default",plain:"",onClick:k[12]||(k[12]=t=>(async t=>{t&&await(null==t?void 0:t.validate(((t,l)=>{if(t){Z.value=!0;const t=new Map;D.networkLabels.split("\n").forEach((e=>{const[a,l]=e.split("=");t.set(a,l)}));let l={name:D.networkName,driver:D.networkMode,network_sebnet:D.networkSubnet,network_gateway:D.networkGateway,labels:t};g(l).then((e=>{const{data:t}=e;Z.value=!1,H.value=!1,a("创建成功"),R()})).catch((a=>{e(a),Z.value=!1}))}else e("请检查输入")})))})(E(T)))},{default:x((()=>[M(" 确定 ")])),_:1})),[[ie,E(Z)]])])])),_:1},8,["modelValue"])])],64)}}}),[["__scopeId","data-v-65b8773c"]]);export{Q as default};