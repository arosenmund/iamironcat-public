/*
 jQuery Cookie Plugin v1.4.1
 https://github.com/carhartl/jquery-cookie

 Copyright 2013 Klaus Hartl
 Released under the MIT license
*/
(function(c){"function"===typeof define&&define.amd?define(["jquery"],c):"object"===typeof exports?c(require("jquery")):c(jQuery)})(function(c){function n(a){a=e.json?JSON.stringify(a):String(a);return e.raw?a:encodeURIComponent(a)}function m(a,d){if(e.raw)var b=a;else a:{0===a.indexOf('"')&&(a=a.slice(1,-1).replace(/\\"/g,'"').replace(/\\\\/g,"\\"));try{a=decodeURIComponent(a.replace(l," "));b=e.json?JSON.parse(a):a;break a}catch(h){}b=void 0}return c.isFunction(d)?d(b):b}var l=/\+/g,e=c.cookie=
function(a,d,b){if(void 0!==d&&!c.isFunction(d)){b=c.extend({},e.defaults,b);if("number"===typeof b.expires){var h=b.expires,g=b.expires=new Date;g.setTime(+g+864E5*h)}return document.cookie=[e.raw?a:encodeURIComponent(a),"\x3d",n(d),b.expires?"; expires\x3d"+b.expires.toUTCString():"",b.path?"; path\x3d"+b.path:"",b.domain?"; domain\x3d"+b.domain:"",b.secure?"; secure":""].join("")}b=a?void 0:{};h=document.cookie?document.cookie.split("; "):[];g=0;for(var l=h.length;g<l;g++){var f=h[g].split("\x3d");
var k=f.shift();k=e.raw?k:decodeURIComponent(k);f=f.join("\x3d");if(a&&a===k){b=m(f,d);break}a||void 0===(f=m(f))||(b[k]=f)}return b};e.defaults={};c.removeCookie=function(a,d){if(void 0===c.cookie(a))return!1;c.cookie(a,"",c.extend({},d,{expires:-1}));return!c.cookie(a)}});