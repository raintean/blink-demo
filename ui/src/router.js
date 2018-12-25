import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Logo from './views/Logo'

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/logo',
      name: 'logo',
      component: Logo
    }
  ]
})
