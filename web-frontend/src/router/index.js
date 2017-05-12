import Vue from 'vue'
import Router from 'vue-router'
import Navbar from '@/components/Navbar'
import Dynatitle from '@/components/Dynatitle'
import vMediaQuery from 'v-media-query'

Vue.use(Router)



export default new Router({
  routes: [
    {
      path: '/',
      components: {
           navbar: Navbar,
           dynatitle: Dynatitle
       }
    }
  ]
})
