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
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
    meta: {title:'Mantiya | About'}
  },
  {
    path: '/games',
    name: 'games',
    component: () => import('../views/Games.vue'),
    meta: {title:'Mantiya | Games made with love'}

  },
  {
    path: '/games/:id',
    name: 'game',
    component: () => import ('../views/Game.vue'),
    meta: {title:'Mantiya | Game'}
  },
  {
    path: '/publishing',
    name: 'publishing',
    component: () => import('../views/Publishing.vue'),
    meta: {title:'Mantiya | Publishing'}
  },
]

const router = new VueRouter({
  routes
})

export default router
