<template>
        <body>
        <form class="decor" enctype="multipart/form-data">
            <div class="form-left-decoration"></div>
            <div class="form-right-decoration"></div>
            <div class="circle"></div>
            <div class="form-inner">
                <h3>Sumbit your game!</h3>
                <h5>About game...</h5>
                <input type="text" v-model.trim="game.name" placeholder="Game title" required>

                <input type="text" v-model.trim="game.video" placeholder="Youtube video link" required>

                <div class="example-2">
                    <input type="file" @change="onFilesSelected" name="file" id="file" class="input-file" required>
                    <label for="file" class="btn btn-tertiary js-labelFile">
                        <span class="js-fileName">Upload icon</span>
                    </label>
                </div>

                <input type="text" v-model.trim="game.appstore_link" placeholder="Link to your game" required>


                <textarea placeholder="Tell us more about your game" v-model.trim="game.description" rows="3"></textarea>

                <h5>About team...</h5>
                <input type="text" v-model.trim="sender.name" placeholder="Studio name" required>
                <input type="email" v-model.trim="sender.email" placeholder="Email" required>
                <input type="text" v-model.trim="sender.skype_id" placeholder="Skype ID" required>
                <input type="submit" @click="onUpload" class="purple-button button-round">

                <div class="check_mark" v-if="savingSuccessful">
                    <div class="sa-icon sa-success animate">
                        <span class="sa-line sa-tip animateSuccessTip"></span>
                        <span class="sa-line sa-long animateSuccessLong"></span>
                        <div class="sa-placeholder"></div>
                        <div class="sa-fix"></div>
                    </div>
                </div>

                <h5 style="text-align: center" class="fade-in" v-if="savingSuccessful">Game successfully send!</h5>



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
                appstore_link: ''
            },
            sender:{
                skype_id: '',
                    name:'',
                    email:''
            },
            reader: '',
            savingSuccessful: false
        }),
        methods: {
            onFilesSelected(event) {
                this.reader = new FileReader();
                this.reader.onloadend = function () {
                    this.file = this.result
                }
                this.reader.readAsDataURL(event.target.files[0]);


            },
            onUpload() {
                this.game.icon = this.reader.result
                HTTP.post('/games', {game: this.game, sender: this.sender}).then((response) => {(console.log(response))
                this.savingSuccessful = true})
            }
        }
    }
</script>

<style scoped>
    .fade-in {
        margin-top: 0px;
        font-size: 21px;
        text-align: center;

        -webkit-animation: fadein 4s; /* Safari, Chrome and Opera > 12.1 */
        -moz-animation: fadein 4s; /* Firefox < 16 */
        -ms-animation: fadein 4s; /* Internet Explorer */
        -o-animation: fadein 4s; /* Opera < 12.1 */
        animation: fadein 4s;
    }

    @keyframes fadein {
        from { opacity: 0; }
        to   { opacity: 1; }
    }

    /* Firefox < 16 */
    @-moz-keyframes fadein {
        from { opacity: 0; }
        to   { opacity: 1; }
    }

    /* Safari, Chrome and Opera > 12.1 */
    @-webkit-keyframes fadein {
        from { opacity: 0; }
        to   { opacity: 1; }
    }

    /* Internet Explorer */
    @-ms-keyframes fadein {
        from { opacity: 0; }
        to   { opacity: 1; }
    }

    /* Opera < 12.1 */
    @-o-keyframes fadein {
        from { opacity: 0; }
        to   { opacity: 1; }
    }




    .check_mark {
        width: 80px;
        height: 130px;
        margin: 0 auto;
    }

    button {
        cursor: pointer;
        margin-left: 15px;
    }

    .hide{
        display:none;
    }

    .sa-icon {
        width: 80px;
        height: 80px;
        border: 4px solid gray;
        -webkit-border-radius: 40px;
        border-radius: 40px;
        border-radius: 50%;
        margin: 20px auto;
        padding: 0;
        position: relative;
        box-sizing: content-box;
    }

    .sa-icon.sa-success {
        border-color: #4CAF50;
    }

    .sa-icon.sa-success::before, .sa-icon.sa-success::after {
        content: '';
        -webkit-border-radius: 40px;
        border-radius: 40px;
        border-radius: 50%;
        position: absolute;
        width: 60px;
        height: 120px;
        background: white;
        -webkit-transform: rotate(45deg);
        transform: rotate(45deg);
    }

    .sa-icon.sa-success::before {
        -webkit-border-radius: 120px 0 0 120px;
        border-radius: 120px 0 0 120px;
        top: -7px;
        left: -33px;
        -webkit-transform: rotate(-45deg);
        transform: rotate(-45deg);
        -webkit-transform-origin: 60px 60px;
        transform-origin: 60px 60px;
    }

    .sa-icon.sa-success::after {
        -webkit-border-radius: 0 120px 120px 0;
        border-radius: 0 120px 120px 0;
        top: -11px;
        left: 30px;
        -webkit-transform: rotate(-45deg);
        transform: rotate(-45deg);
        -webkit-transform-origin: 0px 60px;
        transform-origin: 0px 60px;
    }

    .sa-icon.sa-success .sa-placeholder {
        width: 80px;
        height: 80px;
        border: 4px solid rgba(76, 175, 80, .5);
        -webkit-border-radius: 40px;
        border-radius: 40px;
        border-radius: 50%;
        box-sizing: content-box;
        position: absolute;
        left: -4px;
        top: -4px;
        z-index: 2;
    }

    .sa-icon.sa-success .sa-fix {
        width: 5px;
        height: 90px;
        background-color: white;
        position: absolute;
        left: 28px;
        top: 8px;
        z-index: 1;
        -webkit-transform: rotate(-45deg);
        transform: rotate(-45deg);
    }

    .sa-icon.sa-success.animate::after {
        -webkit-animation: rotatePlaceholder 4.25s ease-in;
        animation: rotatePlaceholder 4.25s ease-in;
    }

    .sa-icon.sa-success {
        border-color: transparent\9;
    }
    .sa-icon.sa-success .sa-line.sa-tip {
        -ms-transform: rotate(45deg) \9;
    }
    .sa-icon.sa-success .sa-line.sa-long {
        -ms-transform: rotate(-45deg) \9;
    }

    .animateSuccessTip {
        -webkit-animation: animateSuccessTip 0.75s;
        animation: animateSuccessTip 0.75s;
    }

    .animateSuccessLong {
        -webkit-animation: animateSuccessLong 0.75s;
        animation: animateSuccessLong 0.75s;
    }

    @-webkit-keyframes animateSuccessLong {
        0% {
            width: 0;
            right: 46px;
            top: 54px;
        }
        65% {
            width: 0;
            right: 46px;
            top: 54px;
        }
        84% {
            width: 55px;
            right: 0px;
            top: 35px;
        }
        100% {
            width: 47px;
            right: 8px;
            top: 38px;
        }
    }
    @-webkit-keyframes animateSuccessTip {
        0% {
            width: 0;
            left: 1px;
            top: 19px;
        }
        54% {
            width: 0;
            left: 1px;
            top: 19px;
        }
        70% {
            width: 50px;
            left: -8px;
            top: 37px;
        }
        84% {
            width: 17px;
            left: 21px;
            top: 48px;
        }
        100% {
            width: 25px;
            left: 14px;
            top: 45px;
        }
    }
    @keyframes animateSuccessTip {
        0% {
            width: 0;
            left: 1px;
            top: 19px;
        }
        54% {
            width: 0;
            left: 1px;
            top: 19px;
        }
        70% {
            width: 50px;
            left: -8px;
            top: 37px;
        }
        84% {
            width: 17px;
            left: 21px;
            top: 48px;
        }
        100% {
            width: 25px;
            left: 14px;
            top: 45px;
        }
    }

    @keyframes animateSuccessLong {
        0% {
            width: 0;
            right: 46px;
            top: 54px;
        }
        65% {
            width: 0;
            right: 46px;
            top: 54px;
        }
        84% {
            width: 55px;
            right: 0px;
            top: 35px;
        }
        100% {
            width: 47px;
            right: 8px;
            top: 38px;
        }
    }

    .sa-icon.sa-success .sa-line {
        height: 5px;
        background-color: #4CAF50;
        display: block;
        border-radius: 2px;
        position: absolute;
        z-index: 2;
    }

    .sa-icon.sa-success .sa-line.sa-tip {
        width: 25px;
        left: 14px;
        top: 46px;
        -webkit-transform: rotate(45deg);
        transform: rotate(45deg);
    }

    .sa-icon.sa-success .sa-line.sa-long {
        width: 47px;
        right: 8px;
        top: 38px;
        -webkit-transform: rotate(-45deg);
        transform: rotate(-45deg);
    }

    @-webkit-keyframes rotatePlaceholder {
        0% {
            transform: rotate(-45deg);
            -webkit-transform: rotate(-45deg);
        }
        5% {
            transform: rotate(-45deg);
            -webkit-transform: rotate(-45deg);
        }
        12% {
            transform: rotate(-405deg);
            -webkit-transform: rotate(-405deg);
        }
        100% {
            transform: rotate(-405deg);
            -webkit-transform: rotate(-405deg);
        }
    }
    @keyframes rotatePlaceholder {
        0% {
            transform: rotate(-45deg);
            -webkit-transform: rotate(-45deg);
        }
        5% {
            transform: rotate(-45deg);
            -webkit-transform: rotate(-45deg);
        }
        12% {
            transform: rotate(-405deg);
            -webkit-transform: rotate(-405deg);
        }
        100% {
            transform: rotate(-405deg);
            -webkit-transform: rotate(-405deg);
        }
    }

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
    .button-round{
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