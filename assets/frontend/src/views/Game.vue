<template>
    <section class="game">
        <header class="head-game perspective">
        </header>
        <div class="container tight bg-accent">
            <div class="row">
                <div class="col-md-3 icon-game"><img :src="model.icon" class="img-large"></div>
                <div class="col-md-9">
                    <h2 class="title">{{model.name}}</h2>
                    <p class="description">{{model.description}}</p>
                    <div>
                        <a href="https://apps.apple.com/us/app/perspective/id1469038654" target="_blank" class="btn-store"><img src="../../public/images/btn-appstore.png"></a>
                        <a href="#" target="_blank" class="btn-store"><img src="../../public/images/btn-googleplay.png"></a>
                    </div>
                </div>
            </div>
        </div>
        <div class="container bg-video">
            <div class="embed-responsive embed-responsive-16by9">
                <iframe v-bind:src="model.video" width="680" height="480" allowfullscreen></iframe>
            </div>
        </div>
    </section>
</template>

<script>
    import HTTP from "../components/http";

    export default {
        name: "Game",
        data() {
            return {
                model: {
                    name: 'Title',
                    description: 'Text ...',
                    video: 'link...',
                    icon: 'link...',
                    background: 'link...'
                }
            };
        },
        created() {
            var id = this.$route.params.id;
            HTTP.get('http://127.0.0.1:8081/games/' +id)
                .then(response => (this.model = response.data));
        }
    }
</script>

<style scoped>

</style>