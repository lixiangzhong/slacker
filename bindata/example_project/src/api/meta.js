import http from "./http";

export function Get() {
  return http.get("/meta");
}
