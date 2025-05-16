import { createRouter, createWebHistory} from 'vue-router'
import Dashboard from '../views/dashboard.vue'
import Login from '../views/login.vue'

const routes = [
    {
        path: '/',
        name: 'Login',
        component: Login
    },
    {
        path: '/dashboard',
        name: 'Dashboard',
        component: Dashboard,
        meta: { requiresAuth: true }
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

router.beforeEach((to, from, next) => {
    const token = localStorage.getItem('token')

    if (to.meta.requiresAuth && !token){
        next('/')
    } else if (to.path === '/' && token) {
        next('/dashboard')
    }
      else {
         next()
      }
})

export default router
