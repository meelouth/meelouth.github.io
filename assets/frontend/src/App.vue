<template>
  <div id="app">
    <nav class="navbar navbar-expand-md navbar-dark fixed-top bg-dark-transparent">
      <img src="./images/logo.png" width="30" height="30">
      <a class="navbar-brand" style="font-family: Pacifico; padding-left: 10px" href="#"> Mantiya</a>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarCollapse" aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarCollapse">
        <ul class="navbar-nav mr-auto">
          <li class="nav-item">
            <a class="nav-link" href="index.html">About</a>
          </li>
          <li class="nav-item ">
            <a class="nav-link" href="games.html">Games</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#">Publishing</a>
          </li>
        </ul>
      </div>
    </nav>

    <main>
      <router-view/>
    </main>

    <footer id="colophon" class="site-footer globalFooter" role="contentinfo">

      <div class="container">
      </div>
      <div class="globalFooter-logos">
        <img class="custom-logo-icon" src="./images/logo.png"/>

      </div>
      <div class="globalFooter-legal"><p>© 2019 Mantiya Entertainment. All Rights Reserved. Mantiya logo are registered or unregistered trademarks of Mantiya Entertainment in the Russia. and/or other countries.</p><div class="globalFooter-legal--links"><a href="#" target="_blank" rel="noopener
noreferrer" data-aa-location="footer" data-aa-category="action" data-aa-action="legal - contacts">Support</a><span>  |  </span><a
              href="#" target="_blank" rel="noopener noreferrer" data-aa-location="footer"
              data-aa-category="action" data-aa-action="legal - privacy">Privacy</a><span>  |  </span><a href="#"
                                                                                                         target="_blank" rel="noopener noreferrer" data-aa-location="footer" data-aa-category="action" data-aa-action="legal - legal
info">Legal Info</a><span>  |  </span><a href="#" target="_blank" rel="noopener noreferrer"
                                         data-aa-location="footer" data-aa-category="action" data-aa-action="legal - terms">Terms</a><span>  |  </span><a href="https://legal.ubi.com/"
                                                                                                                                                          target="_blank" rel="noopener noreferrer" data-aa-location="footer" data-aa-category="action" data-aa-action="legal -
documents ">Legal Documents</a><span>  |  </span><a
              href="#" target="_blank"
              rel="noopener noreferrer" data-aa-location="footer" data-aa-category="action" data-aa-action="legal - interest based
Ads">Interest-Based Ads</a><span>  |  </span><a href="#" target="_blank"
                                                rel="noopener noreferrer" data-aa-location="footer" data-aa-category="action" data-aa-action="legal - code of conduct">Code
        of Conduct</a></div>
      </div>
    </footer>

  </div>
</template>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
