<template>
    <div class="mr-md-3 pt-3 px-3 pt-md-5 px-md-5 text-center overflow-hidden" :class="{'bg-dark text-white' : flag === true,
    'bg-light text-dark' : flag === false}">
        <div onclick="location.href='http://localhost:63342/Mantiya-website/game.html';" style="cursor: pointer;">
            <div class="my-3 py-3">
                <h2 class="display-5">{{model.name}}</h2>
                <p class="lead">{{model.description}}</p>
            </div>
            <img class="circle-border" :src="model.icon" width="320" height="320" alt="Perspective">
        </div>
    </div>
</template>

<script>
    export default {
        name: "GameComponent",
        props: {
            x: {
                type: Object, required :true
            },
            model: {
                type: Object,required: true
            }
        },
        data() {
            return {
            flag: false
            }
        },
        methods :{
            isEven(){
                if(this.x % 2 === 0)
                    this.flag = true;
            }
        },
        created() {
            this.isEven();
        }
    }
</script>

<style scoped>

</style>