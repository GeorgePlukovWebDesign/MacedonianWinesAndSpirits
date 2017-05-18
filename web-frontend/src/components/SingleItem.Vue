<template id="singleitem">
<div class="item">
  <img class="image" src="../assets/wine_bottle.png">
  <!-- <img v-if="bottle.Url !== 'default'" :src="bottle.Url" :alt="bottle.Name"> -->
  <div class="text">

    <h3 class="name">{{bottle.Name}}</h3>
    <h3 class="alcohol">%{{bottle.AlcoholPercentage}}</h3>
    <h3 class="price">${{bottle.Price}} per <img class="bottle" src="../assets/wine_bottle.svg" alt="bottle"></h3>

    <div class="desc">{{bottle.Description}}</div>

  </div>
</div>
</template>

<script>
export default {
  name: 'products',
  props: ["bottle"],
  data() {
    return {
      item: {}
    }
  },
  methods: {}

}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.item{
  border: 1px solid grey;
  margin: 5px 5px 5px 5px;
}
.image {
  width:12%;
  display: inline-block;
}

.text {
  display: inline-block;
  width:85%;
  vertical-align: top;
  padding: 10px;
}
h3{
  font-size: 1.2em;
  display: inline-block;
  font-weight: 100;
}
.name {

}
.alcohol{
  margin-left: 5px;
  float: right;

}
.bottle{
  color: #333;
  height: 1.2em;
  vertical-align: middle;
}
.price{
  float: right;

}
.desc {

}
</style>
