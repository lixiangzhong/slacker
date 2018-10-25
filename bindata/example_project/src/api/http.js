import axios from 'axios'
import store from "@/store";
import {
  Message,
  MessageBox
} from 'element-ui'



import * as cookie from '@/utils/cookie'

// 创建axios实例
const http = axios.create({
  // baseURL: process.env.BASE_API, // api的base_url
  baseURL: process.env.BaseURL,
  timeout: 15000 // 请求超时时间
})

// request拦截器
http.interceptors.request.use(config => {
  if (store.state.token) {
    config.headers['Authorization'] = cookie.getToken() // 让每个请求携带自定义token 请根据实际情况自行修改
  }
  return config
}, error => {
  // Do something with request error
  console.log(error) // for debug
  Promise.reject(error)
})

// respone拦截器
http.interceptors.response.use(
  response => {
    const res = response.data
    if (res.code == 403) {
      store.dispatch('LogOut').then(() => {
        location.reload()
      })
    }
    if (res.code !== 0) {
      Message({
        message: res.msg,
        type: 'error',
        duration: 5 * 1000
      })
      return Promise.reject('error')
    } else {
      return response.data
    }
  },
  error => {
    console.log('err' + error) // for debug
    Message({
      message: error.message,
      type: 'error',
      duration: 5 * 1000
    })
    return Promise.reject(error)
  }
)

export default http
