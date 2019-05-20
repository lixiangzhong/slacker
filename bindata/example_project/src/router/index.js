import Vue from 'vue'
import Router from 'vue-router'
import Layout from "@/layout/Layout";
import Login from '@/views/Login'

Vue.use(Router)

export default new Router({
  routes: [{
    path: '/login',
    name: 'Login',
    component: Login
  }, {
    path: '/logout',
  }, {
    path: '/',
    redirect: '/admin',
  }, {
    path: "/admin",
    name: "Admin",
    component: Layout,
    {{range $i,$v:=.Tables}}{{if eq $i 0}}redirect: '/admin/{{$v.LowerName}}',{{end}}{{end}}
    children: [{{range $i,$v:=.Tables}}{{if gt $i 0}},{{end}}{
      path: '{{$v.LowerName}}',
      component: () =>
        import ('@/views/{{$v.CamelCaseNameWithDBName}}'),
    }{{end}}]
  }]
})
