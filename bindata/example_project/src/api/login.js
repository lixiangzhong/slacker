import http from "./http";


export function getToken(params) {
  return http.post('/token', params)
}
