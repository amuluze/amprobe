import{r as n}from"./index.DGmEUHa-.js";function t(t){return n.get("/api/v1/container/containers",t)}function r(t){return n.post("/api/v1/container/container_create",t)}function e(t){return n.post("/api/v1/container/container_start",t)}function i(t){return n.post("/api/v1/container/container_stop",t)}function a(t){return n.post("/api/v1/container/container_remove",t)}function o(t){return n.post("/api/v1/container/container_restart",t)}function c(t){return n.post("/api/v1/container/image_remove",t)}function s(){return n.post("/api/v1/container/images_prune",{})}function u(t){return n.get("/api/v1/container/images",t)}function p(t){return n.post("/api/v1/container/image_pull",t)}function v(t){return n.get("/api/v1/container/networks",t)}function f(t){return n.post("/api/v1/container/network_create",t)}function _(t){return n.post("/api/v1/container/network_delete",t)}function g(){return n.get("/api/v1/container/version",{})}function m(){return n.get("/api/v1/container/get_docker_registry_mirrors",{})}function k(t){return n.post("/api/v1/container/set_docker_registry_mirrors",t)}export{k as S,u as a,g as b,v as c,i as d,a as e,r as f,c as g,p as h,f as i,_ as j,m as k,s as p,t as q,o as r,e as s};