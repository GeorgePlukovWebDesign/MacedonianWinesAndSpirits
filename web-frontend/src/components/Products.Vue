<template id="products">
<div class="products ">
  <h2>Products</h2>
  <div class="flex-container" >
    <h3>Red -</h3>

    <div  class="faded-line"></div>
    <div class="" v-for="item in prod" v-if="item.Type == 'Red'">
      <single-item :bottle="item" ></single-item>
    </div>

    <h3>White -</h3>
    <div class="faded-line"></div>
    <div class="" v-for="item in prod" v-if="item.Type == 'White'">
      <single-item :bottle="item" ></single-item>
    </div>

    <h3>Rose -</h3>
    <div class="faded-line"></div>
    <div class="" v-for="item in prod" v-if="item.Type == 'Rose'">
      <single-item :bottle="item" ></single-item>
    </div>

    <h3>Rakija -</h3>
    <div class="faded-line"></div>
    <div class="" v-for="item in prod" v-if="item.Type == 'Rakija'">
      <single-item :test="item" ></single-item>
    </div>


    <br><br><br><br>
    <br><br><br><br>

  </div>
</div>
</template>

<script>
import singleItem from './SingleItem.vue';


export default {
  name: 'products',
  props:['prod'],
  components: {'single-item': singleItem}
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.products{
  padding-left:10%;
  padding-right:10%;
}
.flex-container {
  display: -webkit-flex;
  display: flex;
  flex-direction: -webkit-column;
  flex-direction: column;
  height: 100%;
  width: 100%;
}
h2{
  font-weight: normal;
}
h3{
  flex:1 1 0;
  font-weight: lighter;
  margin-top:10px;

}
.plus{
  font-weight: lighter;
  /*font-family: 'Verdana';*/
  cursor: pointer;
}
.faded-line {
  width: 100%;
  height: 2px;
  border-bottom: 1px solid grey;
}
.test {
  color:red;
  font-size: 300px;
}
</style>
