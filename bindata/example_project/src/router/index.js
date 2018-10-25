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
    redirect: '/admin/ip',
    children: [{
      path: 'ip',
      component: () =>
        import ('@/views/IP'),
    }, {
      path: 'ns',
      component: () =>
        import ('@/views/Ns'),
    }, {
      path: 'position',
      component: () =>
        import ('@/views/Position'),
    }]
  }]
})
