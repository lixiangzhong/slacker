import Cookies from 'js-cookie'

const TokenKey = '{{.ProjectName}}-token'

export function getToken() {
  return Cookies.get(TokenKey)
}

export function setToken(token) {
  return Cookies.set(TokenKey, token)
}

export function removeToken() {
  return Cookies.remove(TokenKey)
}

export function Get(key) {
  return Cookies.get(key)
}

export function Set(key, value) {
  return Cookies.set(key, value)
}

export function Del(key) {
  return Cookies.remove(key)
}
