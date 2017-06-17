import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
import ghPinnedRepos from 'gh-pinned-repos'

Vue.use(VueAxios, axios)
ghPinnedRepos('georgeplukov')
  .then(console.log) // [...]
new Vue({
  el: '#app',
  render: h => h(App)
})
