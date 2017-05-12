<template>
<div id="app">

  <router-view></router-view>
</div>
</template>

<script>
export default {
  name: 'app'
}
</script>

<style>
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
  /*text-align: center;*/
  /*color: #2c3e50;*/
  color: #333 !important;
  /*margin-top: 60px;*/
}
</style>
