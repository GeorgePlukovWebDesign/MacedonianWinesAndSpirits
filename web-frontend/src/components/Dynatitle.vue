<template id="dynatitle">

  <h1 class="title">
    <span v-for="word in words">
      <span class="letter">{{word[0]}}</span><span class="word">{{word[1]}}</span>
    </span>

  </h1>
</template>
<script>
export default {
  name: 'dynatitle',
  data() {
    return {
      words: [["M","acedonian "], ["W","ines "], ["& "], ["S","pirits "]]
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.title{
  margin-left: 10px;
  padding-right: 5px;
  /* Positioning */
  position: sticky;
  top: 50%;
  left: 50%;
  /*width: 100%;*/
  max-width: 40%;

  /* Styling */
  text-align: left;
  font-size: 10em;
  font-weight: normal;

  /* translations and animations */
  transform: translate(-50%,-50%);
  transition: transform .4s ease;
  transition: font-size 4s ease;

  /*font-size*/


}
@media (max-width: 1200px) {
  .title{
    font-size: 5em;
    transition: transform .6s;
    transition: font-size .2s ease;

  }
}
@media (max-width: 630px) {
  .title{
    /* Positioning */
    max-width: 100%;
    width:100%;
    font-size: 2em;

    /* translations and animations */
    transition: transform .6s;
    transition: font-size .2s ease;

    transform: translateY(-50%);
  }
}

</style>
