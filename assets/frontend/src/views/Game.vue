<template>
    <section class="game">
        <header class="head-game perspective">
        </header>
        <div class="container tight bg-accent">
            <div class="row">
                <div class="col-md-3 icon-game"><img v-bind:src="model.icon" class="img-large"></div>
                <div class="col-md-9">
                    <h2 class="title">{{model.gameTitle}}</h2>
                    <p class="description">{{model.description}}</p>
                    <div>
                        <a href="https://apps.apple.com/us/app/perspective/id1469038654" target="_blank" class="btn-store"><img src="../images/btn-appstore.png"></a>
                        <a href="#" target="_blank" class="btn-store"><img src="../images/btn-googleplay.png"></a>
                    </div>
                </div>
            </div>
        </div>
        <div class="container bg-video">
            <div class="embed-responsive embed-responsive-16by9">
                <iframe v-bind:src="model.videoLink" width="680" height="480" allowfullscreen></iframe>
            </div>
        </div>
    </section>
</template>

<script>
    export default {
        name: "Game",
        data(){
            return {
                model:{
                    gameTitle: 'Perspective',
                    icon: '../images/icon.jpg',
                    description: 'New addictive puzzle game! Rotate and connect the platforms so that the ball can roll on them. Very simple and challenging game at the same time!',
                    videoLink: 'https://www.youtube.com/embed/LlCJUH6ntIQ'
                }
            }
        }
    }
</script>

<style scoped>

</style>