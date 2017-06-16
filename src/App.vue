<template>
<div id="app">
    <div class="block">
        <div class="name">
            <span class="lato-hairline">George </span><span class="lato-bold">Plukov</span>
        </div>
        <div class="menu raleway">
            <a v-on:mouseover="mouseOver(0)"  v-on:mouseleave="mouseLeave(0)"href="http://www.google.com">Gallery</a>
            <a v-on:mouseover="mouseOver(1)"  v-on:mouseleave="mouseLeave(1)" href="#">Blog</a>
            <a v-on:mouseover="mouseOver(2)"  v-on:mouseleave="mouseLeave(2)" href="#">
              Github
              <!-- <img class="github" src="./assets/GitHub-Mark-120px.png" alt="Link to George Plukov's github page. www.github.com/GeorgePlukov"> -->
            </a>
            <a v-on:mouseover="mouseOver(3)"  v-on:mouseleave="mouseLeave(3)" href="#">Pinterest</a>
            <a v-on:mouseover="mouseOver(4)"  v-on:mouseleave="mouseLeave(4)" href="#">Insta
<!-- <img class="github" src="./assets/Insta-Mark-120px.png" alt="Link to George Plukov's github page. www.github.com/GeorgePlukov"> -->
            </a>


        </div>
        <transition name="slide-fade">
            <div v-if="current_active != -1"  class="preview">
                <div class="lineblock">
                    <!-- Line that fits the gallery item -->
                  <svg v-show="current_active == 0" id="line-gallery" class="line" height="10" width="100%">
                    <line x1="0px" y1="10px" x2="2em" y2="10px"  />
                    <line class="lighter" x1="2em" y1="10px" x2="2.5em" y2="0px" />
                    <line class="lighter" x1="2.5em" y1="0px" x2="3em" y2="10px" />
                    <line x1="3em" y1="10px" x2="83%" y2="10px" />
                  </svg>
                  <div v-show="current_active == 0" class="todo1">
                      <div class="item-box">
                      </div>
                      <div class="item-box">
                      </div>
                      <div class="item-box">
                      </div>
                  </div>


                  <svg v-show="current_active == 1" id="line-blog" class="line" height="10" width="100%">
                    <line x1="0px" y1="10px" x2="6.25em" y2="10px"  />
                    <line class="lighter" x1="6.25em" y1="10px" x2="6.75em" y2="0px" />
                    <line class="lighter" x1="6.75em" y1="0px" x2="7.25em" y2="10px" />
                    <line x1="7.25em" y1="10px" x2="83%" y2="10px" />
                  </svg>
                  <svg v-show="current_active == 2" id="line-github" class="line" height="10" width="100%">
                    <line x1="0px" y1="10px" x2="10em" y2="10px"  />
                    <line class="lighter" x1="10em" y1="10px" x2="10.5em" y2="0px" />
                    <line class="lighter" x1="10.5em" y1="0px" x2="11em" y2="10px" />
                    <line x1="11em" y1="10px" x2="83%" y2="10px" />
                  </svg>
                  <svg v-show="current_active == 3" id="line-pinterest" class="line" height="10" width="100%">
                    <line x1="0px" y1="10px" x2="15em" y2="10px"  />
                    <line class="lighter" x1="15em" y1="10px" x2="15.5em" y2="0px" />
                    <line class="lighter" x1="15.5em" y1="0px" x2="16em" y2="10px" />
                    <line x1="16em" y1="10px" x2="83%" y2="10px" />
                  </svg>
                  <svg v-show="current_active == 4" id="line-instagram" class="line" height="10" width="100%">
                    <line x1="0px" y1="10px" x2="19.5em" y2="10px"  />
                    <line class="lighter" x1="19.5em" y1="10px" x2="20em" y2="0px" />
                    <line class="lighter" x1="20em" y1="0px" x2="20.5em" y2="10px" />
                    <line x1="20.5em" y1="10px" x2="83%" y2="10px" />
                  </svg>
                </div>

            </div>
        </transition>
    </div>
</div>
</template>

<script>
export default {
    name: 'app',
    data() {
        return {
            name: 'george',
            current_active: -1
        }
    },
    methods: {
        mouseOver: function(num) {
          this.current_active = num
        },
        mouseLeave: function(num){
          this.current_active = -1;
        }
    }
}
</script>

<style lang="scss">
#app {
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
}
.block {
    margin-top: 28vh;
    margin-left: 17%;
}
.name {
    font-size: 5em;
    user-select: none;
}
.menu {
    margin-left: 10px;
}
a {
    text-decoration: none;
    color: #333;
    padding-left: 4px;
    padding-right: 4px;
    font-size: 1.2em;
}
.github {
    height: 1.2em;
    width: 1.2em;
    vertical-align: middle;
}
.lato-hairline {
    font-family: 'Lato', sans-serif;
    font-weight: 100;
}
.lato-bold {
    font-family: 'Lato', sans-serif;
    font-weight: 300;
}
.raleway {
    font-family: 'Raleway', sans-serif;
    font-weight: 400;
}
.preview {
    height: 30vh;

}
line {
    // color:grey;
    stroke: grey;
    stroke-width: 1;
}
.lighter {
    stroke-width: 0.5;
}

.todo1 {
    height: 100%;
    // border: 1px solid #333;
    // background-color: red;
}
/* Enter and leave animations can use different */
/* durations and timing functions.              */
.slide-fade-enter-active {
    transition: all 0.4s;
}
.slide-fade-leave-active {
    transition: all 0.2s;
}
/* .slide-fade-leave-active for <2.1.8 */
.slide-fade-enter{
    transform: translateY(40%);
    opacity:0;
}
.slide-fade-leave-to {
    opacity: 0;

}
.item-box{
  display: inline-block;
  height: 100%;
  min-height: 150px;
  margin-left: 10px;
  margin-right: 5px;
  width: 25%;
  border: 1px solid lightgrey;
  border-radius: 10px;
}
</style>
