<template>
<div id="app">

    <navbar></navbar>
    <!-- <dynatitle></dynatitle> -->
    <titlepage></titlepage>

    <products v-bind:prod="prod"></products>
</div>
</template>

<script>
import navbar from './components/Navbar.vue'
import dynatitle from './components/Dynatitle.vue'
import titlepage from './components/Titlepage.vue'
import products from './components/Products.vue'
import axios from 'axios'

export default {
  name: 'app',
  components:{navbar, dynatitle, titlepage, products},
  data () {
    return {
      prod:[]
    }
  },
  mounted: function(){
    axios.get(`http://localhost:8081/products/`)
    .then(response => {
        // JSON responses are automatically parsed.
        this.prod = response.data.products
        console.log(this.prod)
    })
    .catch(e => {
      // this.errors.push(e)/
    })
  }
}
</script>

<style>
html,body{
  overflow: hidden;
}
html,body,div,p,span,ul,li,ol,textarea{
  font-family: 'Open Sans', sans-serif;

  color: #333 !important;
}
img,form,input,h1,h2,h3,h4,h5,h6,pre {
  color:#333 !important;
  margin:0;
  padding:0;
  font-family: 'Roboto', sans-serif;

}

#app {
  font-family: 'Open Sans', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  overflow-y: scroll;
  width: 100%;
  height: 100vh;

  color: #333 !important;

}
</style>
