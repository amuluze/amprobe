import{X as r,ad as t,Z as e,ai as n}from"./el-button.wSr2vfOs.js";import{h as o,e as u,j as a,o as f,U as c}from"./_Uint8Array.DhThSVct.js";var s=Object.create,i=function(){function t(){}return function(e){if(!r(e))return{};if(s)return s(e);t.prototype=e;var n=new t;return t.prototype=void 0,n}}();function p(r,t){var e=-1,n=r.length;for(t||(t=Array(n));++e<n;)t[e]=r[e];return t}function v(r,n,o,u){var a=!o;o||(o={});for(var f=-1,c=n.length;++f<c;){var s=n[f],i=void 0;void 0===i&&(i=r[s]),a?t(o,s,i):e(o,s,i)}return o}var l=Object.prototype.hasOwnProperty;function y(t){if(!r(t))return function(r){var t=[];if(null!=r)for(var e in Object(r))t.push(e);return t}(t);var e=o(t),n=[];for(var u in t)("constructor"!=u||!e&&l.call(t,u))&&n.push(u);return n}function b(r){return u(r)?a(r,!0):y(r)}var d=f(Object.getPrototypeOf,Object),j="object"==typeof exports&&exports&&!exports.nodeType&&exports,h=j&&"object"==typeof module&&module&&!module.nodeType&&module,m=h&&h.exports===j?n.Buffer:void 0,O=m?m.allocUnsafe:void 0;function g(r,t){if(t)return r.slice();var e=r.length,n=O?O(e):new r.constructor(e);return r.copy(n),n}function w(r){var t=new r.constructor(r.byteLength);return new c(t).set(new c(r)),t}function x(r,t){var e=t?w(r.buffer):r.buffer;return new r.constructor(e,r.byteOffset,r.length)}function U(r){return"function"!=typeof r.constructor||o(r)?{}:i(d(r))}export{w as a,x as b,v as c,p as d,g as e,d as g,U as i,b as k};
