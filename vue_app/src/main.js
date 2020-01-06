import Vue from 'vue'
import App from './App.vue'
import  Router  from  'vue-router'
import SearchDomain from './views/SearchDomain'
import History from './views/History'
import Instructions from './views/Instructions'

Vue.config.productionTip = false
Vue.use(Router)

const  router  =  new  Router({
  routes: [
    {
      path:  '/',
      name:  'instructions',
      component:  Instructions
    },
    {
      path:  '/search',
      name:  'search',
      component:  SearchDomain
    },
    {
      path:  '/history',
      name:  'history',
      component:  History
    }
  ]
})

new  Vue({router,
render:  h  =>  h(App)
}).$mount('#app')
