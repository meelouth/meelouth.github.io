<template>
    <section class=" games">
    <div class="position-relative background-image overflow-hidden p-3 p-md-5 m-md-3 text-center bg-light">
        <div class="col-md-5 p-lg-5 mx-auto my-5" >
            <h1 style="display: inline" class="display-4 font-weight-normal text-white"> We make</h1> <h1 style="color:blueviolet; display: inline; font-family: Pacifico" class=" display-4 font-weight-normal"> &nbsp;hit</h1> <h1 class="display-4 font-weight-normal text-white"> games</h1>
            <p class="lead font-weight-normal text-white">We build awesome games for our players.</p>
        </div>
        <div class="product-device shadow-sm d-none d-md-block"> <img class="image-in-phone" src="/images/screenshot5.png" width="276" height="500" alt="Perspective"></div>
    </div>

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