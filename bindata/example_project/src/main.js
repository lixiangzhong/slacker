// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from "vue";
import App from "./App";
import router from "./router";
import ElementUI from "element-ui";
import "element-ui/lib/theme-chalk/index.css";
import store from "./store";
import * as utils from "@/utils/utils";
Vue.config.productionTip = false;
Vue.use(ElementUI);

import mycomponents from "@/components/mycomponents";
Vue.use(mycomponents);

router.beforeEach((to, from, next) => {
  if (to.path === "/logout") {
    store.dispatch("LogOut");
    next("/login");
  }
  if (store.state.token) {
    if (to.path === "/login") {
      next("/");
    } else {
      next();
    }
  } else {
    if (to.path === "/login") {
      next();
    } else {
      next("/login");
    }
  }
});

/* eslint-disable no-new */

Vue.filter("int2ip", function(i) {
  return utils.int2ip(i);
});

Vue.filter("int2ip_reverse", function(i) {
  return utils.int2ip_reverse(i);
});

Vue.filter("unix2datetime", function(i) {
  return utils.parseTime(i);
});

Vue.filter("Kbyte", function(i) {
  if (i == NaN || i == undefined) {
    return "0";
  }
  return utils.formatKbps(i * 8);
});

Vue.filter("duration", function(i) {
  if (i == NaN || i == undefined) {
    return "0秒";
  }
  var day = Math.floor(i / (24 * 3600)); // Math.floor()向下取整
  var hour = Math.floor((i - day * 24 * 3600) / 3600);
  var minute = Math.floor((i - day * 24 * 3600 - hour * 3600) / 60);
  var second = i - day * 24 * 3600 - hour * 3600 - minute * 60;
  if (day) {
    return day + "天" + hour + "时" + minute + "分" + second + "秒";
  }
  return hour + "时" + minute + "分" + second + "秒";
});

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
