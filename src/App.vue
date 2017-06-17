<template>
<div id="app">
    <div class="block">
        <div class="name">
            <span class="lato-hairline">George </span><span class="lato-bold">Plukov</span>
        </div>
        <div class="menu raleway">
            <a v-on:mouseover="mouseOver(0)" v-on:mouseleave="mouseLeave(0)" href="#">Gallery</a>
            <a v-on:mouseover="mouseOver(1)" v-on:mouseleave="mouseLeave(1)" href="#">Blog</a>
            <a v-on:mouseover="mouseOver(2)" v-on:mouseleave="mouseLeave(2)" href="#">
              Github
              <!-- <img class="github" src="./assets/GitHub-Mark-120px.png" alt="Link to George Plukov's github page. www.github.com/GeorgePlukov"> -->
            </a>
            <a v-on:mouseover="mouseOver(3)" v-on:mouseleave="mouseLeave(3)" href="#">Pinterest</a>
            <a v-on:mouseover="mouseOver(4)" v-on:mouseleave="mouseLeaveButton(4)" href="#">Insta
<!-- <img class="github" src="./assets/Insta-Mark-120px.png" alt="Link to George Plukov's github page. www.github.com/GeorgePlukov"> -->
            </a>


        </div>
        <transition name="slide-fade">
            <div v-if="current_active != -1" class="preview">
                <div class="lineblock">
                    <!-- Line that fits the gallery item -->
                    <svg v-show="current_active == 0" id="line-gallery" class="line" height="10" width="100%">
                    <line x1="0px" y1="10px" x2="2em" y2="10px"  />
                    <line class="lighter" x1="2em" y1="10px" x2="2.5em" y2="0px" />
                    <line class="lighter" x1="2.5em" y1="0px" x2="3em" y2="10px" />
                    <line x1="3em" y1="10px" x2="83%" y2="10px" />
                  </svg>
                  <div v-show="current_active == 0" v-on:mouseleave="mouseLeave(0)" class="todo1">

                      <!-- <img class="preview-image" v-bind:src="" alt=""> -->

                      <!-- <img class="preview-image" v-bind:src="" alt=""> -->

                      <img class="preview-image" v-bind:src="insta_data['items'][2]['images']['standard_resolution']['url']" alt="">
                  </div>


                  <svg v-show="current_active == 1" id="line-blog" class="line" height="10" width="100%">
                    <line x1="0px" y1="10px" x2="6.25em" y2="10px"  />
                    <line class="lighter" x1="6.25em" y1="10px" x2="6.75em" y2="0px" />
                    <line class="lighter" x1="6.75em" y1="0px" x2="7.25em" y2="10px" />
                    <line x1="7.25em" y1="10px" x2="83%" y2="10px" />
                  </svg>

                  <!-- Github -->
                  <svg v-show="current_active == 2" id="line-github" class="line" height="10" width="100%">
                    <line x1="0px" y1="10px" x2="10em" y2="10px"  />
                    <line class="lighter" x1="10em" y1="10px" x2="10.5em" y2="0px" />
                    <line class="lighter" x1="10.5em" y1="0px" x2="11em" y2="10px" />
                    <line x1="11em" y1="10px" x2="83%" y2="10px" />
                  </svg>
                  <div v-show="current_active == 2" v-on:mouseleave="mouseLeave(2)" class="todo1">

                      <!-- <img class="preview-image" v-bind:src="" alt=""> -->

                      <!-- <img class="preview-image" v-bind:src="" alt=""> -->

                      <img class="preview-image" v-bind:src="insta_data['items'][2]['images']['standard_resolution']['url']" alt="">
                  </div>



                    <svg v-show="current_active == 3" id="line-pinterest" class="line" height="10" width="100%">
                    <line x1="0px" y1="10px" x2="15em" y2="10px"  />
                    <line class="lighter" x1="15em" y1="10px" x2="15.5em" y2="0px" />
                    <line class="lighter" x1="15.5em" y1="0px" x2="16em" y2="10px" />
                    <line x1="16em" y1="10px" x2="83%" y2="10px" />
                  </svg>
                  <!-- <div v-show="current_active == 3" v-on:mouseleave="mouseLeave(3)" class="todo1">

                    <a data-pin-do="embedBoard" data-pin-board-width="400" data-pin-scale-height="240" data-pin-scale-width="80" href="https://www.pinterest.com/pinterest/official-news/"></a>

                      - <img class="preview-image" v-bind:src="insta_data['items'][1]['images']['standard_resolution']['url']" alt="">

                      <img class="preview-image" v-bind:src="insta_data['items'][2]['images']['standard_resolution']['url']" alt=""> ->
                  </div> -->

                    <!-- instagram -->
                    <svg v-show="current_active == 4" id="line-instagram" class="line" height="10" width="100%">
                    <line x1="0px" y1="10px" x2="19.5em" y2="10px"  />
                    <line class="lighter" x1="19.5em" y1="10px" x2="20em" y2="0px" />
                    <line class="lighter" x1="20em" y1="0px" x2="20.5em" y2="10px" />
                    <line x1="20.5em" y1="10px" x2="83%" y2="10px" />
                  </svg>
                    <div v-show="current_active == 4" v-on:mouseleave="mouseLeave(4)" class="todo1">

                        <img class="preview-image" v-bind:src="insta_data['items'][0]['images']['standard_resolution']['url']" alt="">

                        <img class="preview-image" v-bind:src="insta_data['items'][1]['images']['standard_resolution']['url']" alt="">

                        <img class="preview-image" v-bind:src="insta_data['items'][2]['images']['standard_resolution']['url']" alt="">
                    </div>
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
            current_active: -1,
            continue: false,
            insta_data: null
        }
    },
    created: function() {
        this.axios.get('https://www.instagram.com/georgeplukov/media/').then((response) => {
            this.insta_data = response.data

        })
    },
    methods: {
        mouseOver: function(num) {
            this.current_active = num;

        },
        mouseLeaveButton: function(num) {

          setTimeout(function() {
            // if (this.current_active != )
            this.current_active = -1;

          }, 50);

        },
        mouseLeave: function(num) {


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
.slide-fade-enter {
    transform: translateY(40%);
    opacity: 0;
}
.slide-fade-leave-to {
    opacity: 0;

}
.item-box {
    display: inline-block;
    // max-height: 10vh;
    height: auto;
    width: 25%;

    // max-height: 150px;
    margin-left: 10px;
    margin-right: 5px;
    border: 1px solid lightgrey;
    border-radius: 10px;
}
.preview-image {
    height: auto;
    width: 25%;

    display: inline-block;

    // max-height: 150px;
    margin-left: 10px;
    margin-right: 5px;
    border: 1px solid lightgrey;
    border-radius: 10px;
}
</style>
