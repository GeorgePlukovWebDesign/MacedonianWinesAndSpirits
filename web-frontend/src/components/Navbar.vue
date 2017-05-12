<template id="navbar">

<div class="navbar">
  <div class="flex-container">
    <div id="products" class="flex-item start">
      <div class="center-text">
        Products
      </div>
    </div>
    <div id="review" class="flex-item start">
      <div class="center-text">
        Reviews
      </div>
    </div>

    <div class="flex-item filler"></div>

    <div id="about-us" class="flex-item end">
      <div class="center-text">
        About Us
      </div>
    </div>
  </div>
  <div class="faded-line"></div>
</div>
</template>

<script>
export default {
  name: 'navbar',
  data() {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.navbar {
  background-color: white;
  position: fixed;
  width: 100%;
  height: 50px;
}

.flex-container {
  display: -webkit-flex;
  display: flex;
  flex-flow: row;
  background-color: white;
  height: 100%;
  width: 100%;
}

.flex-item {
  flex: 1 3 0;
}

.flex-item:hover {
  background-color: lightgrey;
  cursor: pointer;
  text-shadow: 0 0 5px #333;
}

.filler {
  order: 2;
  flex: 2 2 30% ;
  pointer-events: none;
}

.start {
  order: 1;
}

.end {
  order: 3;
}

.center-text {
  vertical-align: middle;
  text-align: center;
  line-height: 50px;
  height: 50px;
}


/*
    Fancy faded line
    TODO: add media queries for phone design
*/
.faded-line {
  position: fixed;
  top: 50px;
  width: 100%;
  height: 2px;
  background-repeat: no-repeat;
  background-image: radial-gradient(farthest-side at 100% 0px, #000 0%, rgba(0, 0, 0, 0) 55%),
  radial-gradient(farthest-side at 0px 0px, #000 0%, rgba(0, 0, 0, 0) 55%);
}
</style>
