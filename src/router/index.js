import Vue from 'vue'
import VueRouter from 'vue-router'

const Notepad = () => import('../views/Notepad')
const Headed = () => import('../views/Headed')
const Bottom = () => import('../views/Bottom')

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'notepad',
    component: Notepad
  },
  {
    path: '/notepad',
    name: 'notepad',
    component: Notepad
  },
  {
    path: '/headed',
    name: 'headed',
    component: Headed
  },
  {
    path: '/bottom',
    name: 'bottom',
    component: Bottom
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
