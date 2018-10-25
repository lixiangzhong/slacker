import Vue from 'vue'
import Vuex from "vuex";
import * as cookie from "@/utils/cookie";
import * as loginapi from "@/api/login";

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    token: cookie.getToken(),
  },
  mutations: {
    SetToken(state, token) {
      state.token = token
    },
  },
  actions: {
    // 登录
    Login({
      commit
    }, userInfo) {
      return new Promise((resolve, reject) => {
        loginapi.getToken(userInfo).then(response => {
          const data = response.data
          cookie.setToken(data.token)
          commit('SetToken', data.token)
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 前端 登出
    LogOut({
      commit
    }) {
      return new Promise(resolve => {
        commit('SetToken', '')
        cookie.removeToken()
        resolve()
      })
    }
    //
  }
})

export default store
