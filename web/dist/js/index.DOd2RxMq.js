import{j as e,G as a}from"./message.CS_TSLeJ.js";import{E as t,a as o,v as l}from"./el-table-column.CUrmI2vn.js";import{E as i}from"./el-progress.CUTpAxR6.js";import{_ as s}from"./index.ebzY_tEX.js";import{E as r}from"./el-dialog.BQ3hfU_k.js";import{u as n,z as p,E as d,a as m}from"./zh-cn.PWYwJho2.js";/* empty css               */import"./el-select.D4WZ4Qh9.js";import"./el-popper.GbrdInRL.js";import"./el-tooltip.l0sNRNKZ.js";import{E as u}from"./el-card.CS5WT6tL.js";import{a as c,g,p as f,h}from"./index.9iOzwdYR.js";import{d as j,y as v,o as _,r as w,k as C,f as k,m as b,w as y,u as x,Y as z,F as E,a as V,A as B,e as A,b as U,B as $,C as I,_ as L}from"./index.DuNUv9gK.js";import{E as O}from"./index.HvR5tX07.js";import{E as S}from"./index.Ck8DkNjZ.js";import{E as T}from"./index.Cb0Nm553.js";import"./use-form-item.CZ2M-1KL.js";import"./isEqual.C8I6FPlV.js";import"./_Uint8Array.Bq4iIfDz.js";import"./_initCloneObject.CpUBN0YX.js";import"./_baseClone.D8XUM9zm.js";import"./use-dialog.CZCXc5Xp.js";import"./vnode.BrXPxjYI.js";import"./scroll.15sbxcXA.js";import"./focus-trap.D1Nefh4a.js";import"./refs.Cp7KSzlX.js";import"./index.BusWp8wr.js";import"./castArray.CZ4MN7ta.js";import"./index.D8XJyr-Y.js";import"./aria.C-hsWcn7.js";const q=e=>($("data-v-607ed9a5"),e=e(),I(),e),D={class:"am-image-title"},F={class:"am-image-operator"},K={class:"am-table"},P={class:"am-pagination"},G={class:"am-image-pull"},J={class:"am-image-import"},M=q((()=>k("div",{class:"el-upload__text"},[B("Drop file here or "),k("em",null,"click to upload")],-1))),N=q((()=>k("div",{class:"el-upload__tip"},"tar/tar.gz files",-1))),Q=L(j({__name:"index",setup(j){const $=v();_((()=>{q()}));const I=w(0),{data:L,refresh:q,loading:Q,pagination:R}=n(c,{},{}),Y=()=>{f().then((()=>{e({type:"success",message:"清理完成"})}))},H=w(!1),W=w(!1),X=w(""),Z=async()=>{H.value=!0;let e={image_name:X.value};const{data:a}=await h(e);H.value=!1,W.value=!1,q()},ee=w(!1),ae=()=>{a("导入成功"),ee.value=!1,q()};return(a,n)=>{const c=S,f=u,h=t,j=o,v=d,_=m,w=T,te=r,oe=s,le=i,ie=l;return V(),C(E,null,[k("div",D,[k("span",{onClick:n[0]||(n[0]=e=>a.$router.push("/docker/container"))},"容器"),k("span",{onClick:n[1]||(n[1]=e=>a.$router.push("/docker/image"))},"镜像"),k("span",{onClick:n[2]||(n[2]=e=>a.$router.push("/docker/network"))},"网络"),k("span",{onClick:n[3]||(n[3]=e=>a.$router.push("/docker/settings"))},"设置")]),k("div",F,[b(f,{shadow:"never"},{default:y((()=>[b(c,{type:"primary",plain:"",onClick:n[4]||(n[4]=e=>W.value=!0)},{default:y((()=>[B("拉取镜像")])),_:1}),b(c,{type:"primary",plain:"",onClick:n[5]||(n[5]=e=>ee.value=!0)},{default:y((()=>[B("导入镜像")])),_:1}),b(c,{type:"warning",plain:"",onClick:Y},{default:y((()=>[B("清理虚悬镜像")])),_:1})])),_:1})]),b(f,{shadow:"never"},{default:y((()=>[k("div",K,[A((V(),U(j,{data:x(L),key:x(I),"highlight-current-row":"",height:"100%",stripe:""},{default:y((()=>[b(h,{prop:"id",label:"镜像 ID",align:"center",width:"120",fixed:"",sortable:""}),b(h,{prop:"name",label:"镜像名称",align:"center","min-width":"150","show-overflow-tooltip":"",fixed:"",sortable:""}),b(h,{prop:"number",label:"容器数量",align:"center","show-overflow-tooltip":"",width:"120",sortable:""}),b(h,{prop:"tag",label:"镜像 Tag",align:"center","show-overflow-tooltip":"",width:"120"}),b(h,{prop:"created",label:"创建时间",align:"center",width:"200",sortable:""}),b(h,{prop:"size",label:"镜像大小",align:"center",width:"120",sortable:""}),b(h,{label:"操作",width:"160",fixed:"right",align:"center"},{default:y((a=>[b(c,{type:"danger",plain:"",size:"small",onClick:t=>{return o=a.row.id,O.confirm("该操作会删除该镜像. 继续吗?","警告",{confirmButtonText:"OK",cancelButtonText:"Cancel",type:"warning",beforeClose:(a,t,l)=>{"confirm"===a?(t.confirmButtonLoading=!0,g({image_id:o}).then((()=>{e({type:"success",message:"删除完成"})})).finally((()=>{t.confirmButtonLoading=!1,l(),q(),I.value+=1}))):l()}}).then((()=>{})).catch((()=>{e({type:"info",message:"删除取消"})})),void q();var o}},{default:y((()=>[B(" 删除 ")])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"])),[[ie,x(Q)]])]),k("div",P,[b(_,{locale:x(p)},{default:y((()=>[b(v,{"current-page":x(R).page,"onUpdate:currentPage":n[6]||(n[6]=e=>x(R).page=e),"page-size":x(R).size,layout:"total, sizes, prev, pager, next, jumper","page-sizes":x(R).sizeOption,total:x(R).total,onSizeChange:x(R).onSizeChange,onCurrentChange:x(R).onPageChange},null,8,["current-page","page-size","page-sizes","total","onSizeChange","onCurrentChange"])])),_:1},8,["locale"])])])),_:1}),k("div",G,[b(te,{modelValue:x(W),"onUpdate:modelValue":n[8]||(n[8]=e=>z(W)?W.value=e:null),title:"拉取镜像",width:"50%"},{default:y((()=>[b(w,{modelValue:x(X),"onUpdate:modelValue":n[7]||(n[7]=e=>z(X)?X.value=e:null),placeholder:"请输入镜像名称"},null,8,["modelValue"]),A((V(),U(c,{size:"default",type:"info",plain:"",onClick:Z},{default:y((()=>[B(" 确定 ")])),_:1})),[[ie,x(H)]])])),_:1},8,["modelValue"])]),k("div",J,[b(te,{modelValue:x(ee),"onUpdate:modelValue":n[9]||(n[9]=e=>z(ee)?ee.value=e:null),title:"导入镜像",width:"50%"},{default:y((()=>[b(le,{drag:"",action:"/app/api/v1/container/image_import",headers:{Authorization:`Bearer ${x($).user.token}`},limit:1,multiple:"","on-success":ae},{tip:y((()=>[N])),default:y((()=>[b(oe,{"icon-class":"upload"}),M])),_:1},8,["headers"])])),_:1},8,["modelValue"])])],64)}}}),[["__scopeId","data-v-607ed9a5"]]);export{Q as default};
