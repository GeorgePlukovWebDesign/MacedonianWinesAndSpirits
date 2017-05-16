<template id="singleitem">
<div class="">
  <h2></h2>
  <div class="flex-container" >
    <div class="flex-item">

      <div class="inline image">
        <img v-if="item.Url === 'default'" src="assets/placeholder.jpg" :alt="bottle.Name">
        <img v-if="bottle.Url !== 'default'" :src="bottle.Url" :alt="bottle.Name">
      </div>
      <div class="inline text">
        <span class="name">{{bottle.Name}}</span>
        <span class="alcohol">{{bottle.AlcoholPercentage}}</span>
        <span class="price">{{bottle.Price}}</span>
        <div class="desc">{{bottle.Description}}</div>

      </div>

    </div>
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
  methods: {
  }

}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.flex-container{
  display: flex;
}
.flex-item{
  width: 50%;
  /*height: 100px;*/
  /*max-height: ;*/
}
.image{
  display:inline-block;
}
.text{
  display:inline-block;
}
.name{

}
</style>
