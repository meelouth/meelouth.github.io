import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/About.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
    path: '/games',
    name: 'games',
    component: () => import('../views/Games.vue')
  },
  {
    path: '/games/:id',
    name: 'games',
    component: () => import ('../views/Game.vue')
  },
  {
    path: '/publishing',
    name: 'publishing',
    component: () => import('../views/Publishing.vue')
  },
]

const router = new VueRouter({
  routes
})

export default router
