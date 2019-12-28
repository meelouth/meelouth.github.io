<template>
    <div class="d-md-flex flex-md-equal w-100 my-md-3 pl-md-3">
        <GameComponent></GameComponent>
    </div>
</template>

<script>
    import GameComponent from "./GameComponent";
    export default {
        name: "DoubleGamesComponent",
        components: {
            GameComponent
        },
        props: {
            model:{
                type: Object, required: true
            }
        }
    }
</script>


<style scoped>

</style>