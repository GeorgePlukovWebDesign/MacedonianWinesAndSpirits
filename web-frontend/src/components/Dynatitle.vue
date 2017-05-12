<template id="dynatitle">

  <h1 class="title">
    <div v-for="word in words"></div>

  </h1>
</template>
<script>
export default {
  name: 'dynatitle',
  data() {
    return {
      words: ["Macedonian", "Wines", "&","Spirits"]
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.title{
  /* Positioning */
  position: sticky;
  top: 50%;
  left: 50%;
  width: 100%;
  max-width: 450px;

  transform: translate(-50%,-50%);

  /* Styling */
  text-align: left;
  font-size: 3em;
  font-weight: normal;



}
</style>
