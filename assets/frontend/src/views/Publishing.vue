<template>
        <body>
        <form class="decor" enctype="multipart/form-data">
            <div class="form-left-decoration"></div>
            <div class="form-right-decoration"></div>
            <div class="circle"></div>
            <div class="form-inner">
                <h3>Sumbit your game!</h3>
                <h5>About game...</h5>
                <input type="text" v-model.trim="game.name" placeholder="Game title">
                <div class="example-2">
                    <input type="file" name="file" id="file" class="input-file">
                    <label for="file" class="btn btn-tertiary js-labelFile">
                        <span class="js-fileName">Upload gameplay video</span>
                    </label>
                </div>
                <div class="example-2">
                    <input type="file" @change="onFilesSelected">
                    <label for="file" class="btn btn-tertiary js-labelFile">
                        <span class="js-fileName">Upload icon</span>
                    </label>
                </div>

                <textarea placeholder="Tell us more about your game" v-model.trim="game.description" rows="3"></textarea>

                <h5>About team...</h5>
                <input type="text" v-model.trim="sender.name" placeholder="Studio name">
                <input type="email" v-model.trim="sender.email" placeholder="Email">
                <input type="text" v-model.trim="sender.skype_id" placeholder="Skype ID">
                <button @click="onUpload">Upload</button>
            </div>
        </form>
        </body>
</template>

<script>
    import HTTP from "../components/http";
    export default {

        name: "Publishing.vue",
        data: () => ({
            game: {
                name: '',
                description: '',
                icon: 'null',
                video: '',
                appstorelink: ''
            },
                sender:{
                skype_id: '',
                    name:'',
                    email:''
                }
        }),
        methods: {
            onFilesSelected(event){
                this.game.icon = event.target.files[0]
            },
            onUpload(){
                HTTP.post('/games', {
                    game: this.game,
                    sender: this.sender
                }).then(console.log(this.game.icon))
            },
            getBase64(file) {
                var reader = new FileReader();
                reader.readAsDataURL(file);
                reader.onload = function () {
                    console.log(reader.result);
                };
                reader.onerror = function (error) {
                console.log('Error: ', error);
                return reader.result;
            };
    }
        }
    }
</script>

<style scoped>
    body {background: #856396}
    .decor {
        position: relative;
        max-width: 400px;
        margin: 50px auto 0;
        background: white;
        border-radius: 30px;
    }
    .form-left-decoration,
    .form-right-decoration {
        content: "";
        position: absolute;
        width: 50px;
        height: 20px;
        background: #856396;
        border-radius: 20px;
    }
    .form-left-decoration {
        bottom: 60px;
        left: -30px;
    }
    .form-right-decoration {
        top: 60px;
        right: -30px;
    }
    .form-left-decoration:before,
    .form-left-decoration:after,
    .form-right-decoration:before,
    .form-right-decoration:after {
        content: "";
        position: absolute;
        width: 50px;
        height: 20px;
        border-radius: 30px;
        background: white;
    }
    .form-left-decoration:before {top: -20px;}
    .form-left-decoration:after {
        top: 20px;
        left: 10px;
    }
    .form-right-decoration:before {
        top: -20px;
        right: 0;
    }
    .form-right-decoration:after {
        top: 20px;
        right: 10px;
    }
    .circle {
        position: absolute;
        bottom: 80px;
        left: -55px;
        width: 20px;
        height: 20px;
        border-radius: 50%;
        background: white;
    }
    .form-inner {padding: 50px;}
    .form-inner input,
    .form-inner textarea {
        display: block;
        width: 100%;
        padding: 0 20px;
        margin-bottom: 10px;
        background: #E9EFF6;
        line-height: 40px;
        border-width: 0;
        border-radius: 20px;
        font-family: 'Roboto', sans-serif;
    }
    .form-inner input[type="submit"] {
        margin-top: 30px;
        background: #856396;
        border-bottom: 4px solid #59118c;
        color: white;
        font-size: 14px;
    }
    .form-inner textarea {resize: none;}
    .form-inner h3 {
        margin-top: 0;
        font-family: 'Roboto', sans-serif;
        font-weight: 500;
        font-size: 24px;
        color: #707981;
    }
    @media(max-width:480px){
        .form-left-decoration,
        .form-right-decoration,
        .circle {display: none;}
        .form-inner{padding: 30px 20px;}
        .form-inner h3{text-align:center;}
    }
    .example-2 .btn-tertiary{color:#555;padding:0;line-height:40px;width:300px;
        display:block;border:2px solid #555;
        margin: 10px auto 10px;
    }
    .example-2 .btn-tertiary:hover,.example-2 .btn-tertiary:focus{color:#888;border-color:#888}
    .example-2 .input-file{width:.1px;height:.1px;opacity:0;overflow:hidden;position:absolute;z-index:-1}
    .example-2 .input-file + .js-labelFile{overflow:hidden;text-overflow:ellipsis;white-space:nowrap;padding:0 10px;cursor:pointer}
    .example-2 .input-file + .js-labelFile .icon:before{content:"\f093"}
    .example-2 .input-file + .js-labelFile.has-file .icon:before{content:"\f00c";color:#5AAC7B}
</style>