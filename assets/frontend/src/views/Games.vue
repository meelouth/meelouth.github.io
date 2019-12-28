<template>
    <section class=" games">
    <div class="position-relative background-image overflow-hidden p-3 p-md-5 m-md-3 text-center bg-light">
        <div class="col-md-5 p-lg-5 mx-auto my-5" >
            <h1 style="display: inline" class="display-4 font-weight-normal text-white"> We make</h1> <h1 style="color:blueviolet; display: inline; font-family: Pacifico" class=" display-4 font-weight-normal"> &nbsp;hit</h1> <h1 class="display-4 font-weight-normal text-white"> games</h1>
            <p class="lead font-weight-normal text-white">We build awesome games for our players.</p>
        </div>
        <div class="product-device shadow-sm d-none d-md-block"> <img class="image-in-phone" src="/images/screenshot5.png" width="276" height="500" alt="Perspective"></div>
    </div>
        <!--

            <div class="d-md-flex flex-md-equal w-100 my-md-3 pl-md-3">
                <div class="bg-dark mr-md-3 pt-3 px-3 pt-md-5 px-md-5 text-center text-white overflow-hidden">
                    <div onclick="location.href='http://localhost:63342/Mantiya-website/game.html';" style="cursor: pointer;">
                        <div class="my-3 py-3">
                            <h2 class="display-5">Perspective</h2>
                            <p class="lead">Rotate is fun!</p>
                        </div>
                        <img class="circle-border" src="assets/frontend/src/images/icon.jpg" width="320" height="320" alt="Perspective">
                    </div>
                </div>
                <div class="bg-light mr-md-3 pt-3 px-3 pt-md-5 px-md-5 text-center overflow-hidden">
                    <div onclick="location.href='https://apps.apple.com/us/app/perspective/id1469038654';" style="cursor: pointer;">
                        <div class="my-3 py-3">
                            <h2 class="display-5">Dash Ballz</h2>
                            <p class="lead">Swipe up!</p>
                        </div>
                        <img class="circle-border" src="assets/frontend/src/images/icon_2.png" width="320" height="320" alt="Perspective">
                    </div>
                </div>
            </div>

            <div class="d-md-flex flex-md-equal w-100 my-md-3 pl-md-3">
                <div class="bg-light mr-md-3 pt-3 px-3 pt-md-5 px-md-5 text-center overflow-hidden">
                    <div onclick="location.href='https://apps.apple.com/us/app/perspective/id1469038654';" style="cursor: pointer;">
                        <div class="my-3 py-3">
                            <h2 class="display-5">Jelly Escape</h2>
                            <p class="lead">Jelly want touch the sky!</p>
                        </div>
                        <img class="circle-border" src="assets/frontend/src/images/icon_3.png" width="320" height="320" alt="Perspective">
                    </div>
                </div>
                <div class="bg-yellow mr-md-3 pt-3 px-3 pt-md-5 px-md-5 text-center text-white overflow-hidden">
                    <div onclick="location.href='https://apps.apple.com/us/app/perspective/id1469038654';" style="cursor: pointer;">
                        <div class="my-3 py-3">
                            <h2 class="display-5">{{model.name}}</h2>
                            <p class="lead">New .io game!</p>
                        </div>
                        <img class="circle-border" src="assets/frontend/src/images/icon_4.jpg" width="320" height="320" alt="Perspective">
                    </div>
                </div>
            </div> -->

        <GameComponent v-for="(model,x) in models" :key="x" v-bind:x ="x" v-bind:model="model"></GameComponent>
    </section>

</template>

<script>
    import HTTP from "../components/http";
    import GameComponent from "../components/GameComponent";

    export default {
        name: "Games",
        components: {GameComponent},
        data() {
            return {
                models: [],
                gamesPair:[
                    {gamePair:[
                            {
                                name:'1',
                            },
                            {
                                name:'2',
                            }
                        ]
                    },
                    {gamePair:[
                            {
                                name:'1',
                            },
                            {
                                name:'2',
                            }
                        ]
                    },
                    {gamePair:[
                            {
                                name:'3',
                            },
                            {
                                name:'5',
                            }
                        ]
                    }
                ],


            };
        },
        created() {
            HTTP.get('http://127.0.0.1:8081/games')
                .then(response => (this.models = response.data))

        }
    }
</script>

<style scoped>

</style>