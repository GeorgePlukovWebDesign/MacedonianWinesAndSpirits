<template>
<div id="app">
    <div class="block">
        <div class="name">
            <span class="lato-hairline">George </span><span class="lato-bold">Plukov</span>
        </div>
        <div class="menu raleway">
            <a v-on:mouseover="mouseOver(0)" v-on:mouseleave="mouseLeave(0)" href="#">Gallery</a>
            <a v-on:mouseover="mouseOver(1)" v-on:mouseleave="mouseLeave(1)" href="#">Blog</a>
            <a v-on:mouseover="mouseOver(2)" v-on:mouseleave="mouseLeaveButton(2)" href="#">
              Github
            </a>
            <a v-on:mouseover="mouseOver(4)" v-on:mouseleave="mouseLeaveButton(4)" href="#">Instagram
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
                  <div class="github-card" v-show="current_active == 2" v-on:mouseleave="mouseLeave(2)">
                    <span class="icon">K</span>

                    <span class="underline">
                      <span class="github-title">Capstone/</span><span class="github-title font-bold">FollowThrough-</span>
                    </span>

                    <div class="github-desc">
                        An application that helps improve your basketball shot! Track the arc, release height and other stats about your basketball shot.
                    </div>
                    <div class="github-links">
                      Python/OpenCV/PHP
                    </div>
                  </div>
                  <div class="github-card" v-show="current_active == 2" v-on:mouseleave="mouseLeave(2)">
                    <span class="icon">K</span>

                    <span class="underline">
                      <span class="github-title font-bold">3D Terrain generator</span>
                    </span>

                    <div class="github-desc">
                      A terrain generator built in C++ and opengl. Used to expirement with different terrain generation algorithms. <br />
                    </div>
                    <div class="github-links">
                      C++/OpenGl
                    </div>
                  </div>
                  <div class="github-card" v-show="current_active == 2" v-on:mouseleave="mouseLeave(2)">
                    <span class="icon">K</span>

                    <span class="underline">
                      <span class="github-title">Deltahacks/</span><span class="github-title font-bold">FloodWatch-</span>
                    </span>
                    <div class="github-desc">
                        Real time natural disaster early warning system. Second place prize winner at deltahacks!
                    </div>
                    <div class="github-links">
                      Python/XBee/Arduino
                    </div>
                  </div>



                    <!-- Instagram Preview -->
                    <svg v-show="current_active == 4" id="line-instagram" class="line" height="10" width="100%">
                    <line x1="0px" y1="10px" x2="15.5em" y2="10px"  />
                    <line class="lighter" x1="15.5em" y1="10px" x2="16em" y2="0px" />
                    <line class="lighter" x1="16em" y1="0px" x2="16.5em" y2="10px" />
                    <line x1="16.5em" y1="10px" x2="83%" y2="10px" />
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
    border: 1px solid #d1d5da;
    border-radius: 10px;
}
.preview-image {
    height: auto;
    width: 25%;

    display: inline-block;

    margin-left: 10px;
    margin-right: 5px;
    border: 1px solid #d1d5da;
    border-radius: 10px;
}

.github-card {
    height: auto;
    max-width: 23%;
    min-width: 23%;

    display: inline-block;

    // max-height: 150px;
    margin-left: 10px;
    margin-right: 5px;
    padding: 15px 15px 15px 15px;
    border: 1px solid lightgrey;
    border-radius: 10px;
}
.icon{
  color:#586069;
}
.github-title{
  color: #0366d6;
  font-family: Arial;
}
.github-desc{
  margin-top: 8px;
  margin-bottom: 16px;
  // min-height: 00px;
  font-family: sans-serif;
  line-height: 1.5;
  font-size: 12px;
  color: #586069;
}
.github-links{
  color: #586069;
}
.underline:hover{
  text-decoration: underline;
  color:#0366d6;
}
.font-bold{
  font-weight:600;
}
</style>
