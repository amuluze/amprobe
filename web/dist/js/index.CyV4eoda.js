import{d as t,a as e,b as a,w as s,j as o,D as d,ae as l,u as i,T as n}from"./index.Dr9aOSLV.js";import{j as g,_ as r}from"./el-button.Bh3a0xQF.js";const p=t({name:"ElCollapseTransition"});var m=r(t({...p,setup(t){const r=g("collapse-transition"),p=t=>{t.style.maxHeight="",t.style.overflow=t.dataset.oldOverflow,t.style.paddingTop=t.dataset.oldPaddingTop,t.style.paddingBottom=t.dataset.oldPaddingBottom},m={beforeEnter(t){t.dataset||(t.dataset={}),t.dataset.oldPaddingTop=t.style.paddingTop,t.dataset.oldPaddingBottom=t.style.paddingBottom,t.style.height&&(t.dataset.elExistsHeight=t.style.height),t.style.maxHeight=0,t.style.paddingTop=0,t.style.paddingBottom=0},enter(t){requestAnimationFrame((()=>{t.dataset.oldOverflow=t.style.overflow,t.dataset.elExistsHeight?t.style.maxHeight=t.dataset.elExistsHeight:0!==t.scrollHeight?t.style.maxHeight=`${t.scrollHeight}px`:t.style.maxHeight=0,t.style.paddingTop=t.dataset.oldPaddingTop,t.style.paddingBottom=t.dataset.oldPaddingBottom,t.style.overflow="hidden"}))},afterEnter(t){t.style.maxHeight="",t.style.overflow=t.dataset.oldOverflow},enterCancelled(t){p(t)},beforeLeave(t){t.dataset||(t.dataset={}),t.dataset.oldPaddingTop=t.style.paddingTop,t.dataset.oldPaddingBottom=t.style.paddingBottom,t.dataset.oldOverflow=t.style.overflow,t.style.maxHeight=`${t.scrollHeight}px`,t.style.overflow="hidden"},leave(t){0!==t.scrollHeight&&(t.style.maxHeight=0,t.style.paddingTop=0,t.style.paddingBottom=0)},afterLeave(t){p(t)},leaveCancelled(t){p(t)}};return(t,g)=>(e(),a(n,d({name:i(r).b()},l(m)),{default:s((()=>[o(t.$slots,"default")])),_:3},16,["name"]))}}),[["__file","collapse-transition.vue"]]);m.install=t=>{t.component(m.name,m)};const y=m;export{y as _};