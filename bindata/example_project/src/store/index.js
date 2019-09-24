import Vue from "vue";
import Vuex from "vuex";
import * as cookie from "@/utils/cookie";
import * as loginapi from "@/api/login";

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    token: cookie.getToken(),
    username: cookie.Get("username")
  },
  mutations: {
    SetToken(state, token) {
      state.token = token;
    },
    SetUsername(state, name) {
      state.username = name;
    }
  },
  actions: {
    // 登录
    Login({ commit }, userInfo) {
      return new Promise((resolve, reject) => {
        loginapi
          .getToken(userInfo)
          .then(response => {
            const data = response.data;
            cookie.setToken(data.token);
            cookie.Set("username", userInfo.username);
            commit("SetToken", data.token);
            commit("SetUsername", userInfo.username);
            resolve();
          })
          .catch(error => {
            reject(error);
          });
      });
    },
    // 前端 登出
    LogOut({ commit }) {
      return new Promise(resolve => {
        commit("SetToken", "");
        commit("SetUsername", "");
        cookie.removeToken();
        cookie.Del("username");
        resolve();
      });
    }
    //
  }
});

export default store;
