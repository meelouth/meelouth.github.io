<template>
  <div class="hello">
    <video playsinline autoplay muted loop>
      <source src="../Videos/PerspectiveAds02.mp4" type="video/mp4">
      Ваш браузер не поддерживает тег video. Пожалуйста, обновите свой браузер до более свежей версии.
    </video>

    <div class="overlay">
      <h2>UNLEASH THE TRUE POTENTIAL <br>OF YOUR GAMES</h2>
      <button>SUMBIT YOUR GAME</button>
    </div>

    <div id="content" class="site-content">
      <main id="main" class="site-main" role="main">
        <section id="service" class="section-services section-padding section-meta onepage-section">
          <div class="services">
            <div class="row">
              <div class="col-sm-12 col-lg-12 wow slideInUp">
                <div style="padding-bottom:0px" class="service-item ">

                  <div class="service-content">

                    <div class="block-media-text alignwide is-stacked-on-mobile aboutUs"
                         style="background-color:#ffffff;grid-template-columns:21% auto">
                      <figure class="block-media-text__media">
                        <img src="../../public/images/Stack_home.png" />

                      </figure>
                      <div class="block-media-text__content">
                        <p class="violet-text"><strong>ABOUT US</strong></p>


                        <p style="color:#000000;font-size:40px;text-align:left"
                           class="has-text-color big_text"><strong>A SUCCESS-STORY MAKER</strong></p>


                        <p style="color:#000000;text-align:left"
                           class="has-text-color has-medium-font-size">In 2014, driven by a passion for
                          mobile games, we became pioneers in the hyper-casual segment. <br>Today,
                          that same passion motivates our team of specialists to help you create the
                          next big hit! <br>There are more than 800K games in the App Store, we&#8217;re
                          here to make yours stand out.</p>
                      </div>
                    </div>


                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="services">
            <div class="row">
              <div class="col-sm-12 col-lg-12 wow slideInUp">
                <div class="service-item ">
                  <div class="service-content">

                    <div style="height:100px" aria-hidden="true"
                         class="block-spacer custom_spacer"></div>


                    <p style="text-align:center" class="has-text-color violet-text"><strong>OUR ADDED
                      VALUE</strong></p>


                    <p style="color:#000000;font-size:40px;text-align:center"
                       class="has-text-color big_text"><strong>WE SUPPORT YOUR GAME</strong></p>


                    <p style="color:#000000;text-align:center" class="has-text-color has-medium-font-size">
                      Our team provides you with tailor-made assistance </p>


                    <p style="color:#000000;text-align:center" class="has-text-color has-medium-font-size">
                      so you can focus on what you do best : make great games!</p>


                    <div style="height:100px" aria-hidden="true"
                         class="block-spacer custom_spacer"></div>


                    <div id="ourAddedValue" class="row" style="z-index:99;margin-left:10%;margin-right:10%">
                      <div class="col-12 col-sm-12 col-md-4 col-lg-4 col-xl-4"
                           style="z-index:99;text-align:center;">
                        <div class="valueBox">
                          <p class="imgBox" style="padding:10%">
                            <img class="img-support" width="150" height="150"
                                 src="../../public/images/imgBox/picto3.png">
                          </p>
                          <p class="text_whatWeOffer">
                            <strong class="violet-text">Get the most out of your game</strong><br>
                            Whether it’s advice on how to improve the retention or the monetization
                            of your game or how to make a big splash at launch,
                            we will guide you throughout the whole process of creating a game that
                            generates value for the players and for you.
                          </p>
                        </div>
                      </div>
                      <div class="col-12 col-sm-12 col-md-4 col-lg-4 col-xl-4"
                           style="z-index:99;text-align:center;">
                        <div class="valueBox">
                          <p class="imgBox" style="padding:10%">
                            <img class="img-support" width="150" height="150"
                                 src="../../public/images/imgBox/picto1.png">
                          </p>
                          <p class="text_whatWeOffer">
                            <strong class="violet-text">Reach the audience your game truly
                              deserve</strong><br>
                            Thanks to our pioneer cross-promotion system, first class
                            user-acquisition campaigns and largest hypercasual fans’ community on
                            social media, we’ll make sure your game gets the exposure it deserves!
                          </p>
                        </div>
                      </div>
                      <div class="col-12 col-sm-12 col-md-4 col-lg-4 col-xl-4"
                           style="z-index:99;text-align:center;">
                        <div class="valueBox"
                             style="margin:5%; border-radius: 10px;border: 1px solid #dedede; height:100%; padding:10%">
                          <p class="imgBox img-support" style="padding:10%">
                            <img class="img-support" width="140" height="140"
                                 src="../../public/images/imgBox/picto2.png">
                          </p>
                          <p class="text_whatWeOffer">
                            <strong class="violet-text">Enjoy a tight-knit relationship with our
                              team</strong><br>
                            With us, you‘re not just a line in a dashboard : you‘re a partner ! We
                            are a human-sized team of passionate game-makers. All along the way, we
                            offer you our undivided attention and ensure fair revenue-sharing.
                          </p>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>

      </main>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'HelloWorld',
    props: {
      msg: String
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  h3 {
    margin: 40px 0 0;
  }
  ul {
    list-style-type: none;
    padding: 0;
  }
  li {
    display: inline-block;
    margin: 0 10px;
  }
  a {
    color: #42b983;
  }
</style>
