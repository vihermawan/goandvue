import Vue from 'vue'
import VueRouter from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import Register from '@/components/Register'
import Login from '@/components/Login'

Vue.use(VueRouter)

const router = new VueRouter({
  mode: 'history',
  routes: [
    { path: '/', component: Login, name : 'Login' },
    { path: '/home', component: HelloWorld, name: 'HelloWorld' },
    { path: '/register', component: Register, name: 'Register' }
 ]
})
router.beforeEach((to, from, next) => {
  const publicPages = ['/', '/register'];
  const auth = !publicPages.includes(to.path);
  const loggedIn = localStorage.getItem('user');

  if (auth && !loggedIn) {
    next('/');
  } else {
    next();
  }
});
export default router
