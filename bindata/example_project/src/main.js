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

Vue.filter("unix2datetime", function(i) {
  return utils.parseTime(i);
});

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
