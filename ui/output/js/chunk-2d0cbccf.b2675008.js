(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2d0cbccf"],{"4ae1":function(t,e,n){"use strict";n.r(e);var o=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div")},s=[],c={mounted:function(){var t=this,e=this.$route.query.noBADPAccess?"noBDAPAccessUser":"noAccessUser";this.$alert(this.$t(e),this.$t("common.prompt")).then((function(){t.logout()})),setTimeout((function(){t.logout()}),1e4)},methods:{logout:function(){var t=this;this.FesApi.fetchUT("/cc/".concat(this.FesEnv.ccApiVersion,"/logout"),"get").then((function(){t.$router.push("/home")}),(function(e){e.message&&t.$message.error(e.message)}))}}},u=c,i=n("2877"),r=Object(i["a"])(u,o,s,!1,null,null,null);e["default"]=r.exports}}]);