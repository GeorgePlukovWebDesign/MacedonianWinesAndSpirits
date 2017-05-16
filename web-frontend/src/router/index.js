import Vue from 'vue'
import Router from 'vue-router'
import Navbar from '@/components/Navbar'
import Dynatitle from '@/components/Dynatitle'
import Products from '@/components/Products'
import Titlepage from '@/components/Titlepage'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      components: {
           navbar: Navbar,
           dynatitle: Dynatitle,
           titlepage: Titlepage,
           Products: Products
       }
    }
  ]
})
