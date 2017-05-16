<template id="products">
<div class="products">
  <div class="flex-container">

    <div>asd</div>
  </div>
</div>
</template>

<script>
import axios from 'axios';


axios.get(`http://localhost:8081/products/`)
  .then(response => {
    // JSON responses are automatically parsed.
    this.posts = response.data
    console.log(this.posts)
    // alert(this.posts["producs"])
  })
  .catch(e => {
  // this.errors.push(e)/
  })

export default {
  name: 'products',
  data() {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  }

}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.flex-container {
  display: -webkit-flex;
  display: flex;
  flex-wrap: -webkit-column;
  flex-wrap: column;
  height: 100%;
  width: 100%;
}

.flex-item {
  flex: 1 1 0;
}

.flex-item:hover {
  background-color: lightgrey;
  cursor: pointer;
  text-shadow: 0 0 5px #333;
}
</style>
