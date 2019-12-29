<template>
  <div id="app">
    <nav class="navbar navbar-expand-md navbar-dark fixed-top bg-dark-transparent">
      <img src="../public/images/logo.png" width="30" height="30">
      <a class="navbar-brand" style="font-family: Pacifico; padding-left: 10px" href="#"> Mantiya</a>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarCollapse" aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarCollapse">
        <ul class="navbar-nav mr-auto">
          <li class="nav-item">
            <router-link to="/about"><a class="nav-link"> About</a></router-link>
          </li>
          <li class="nav-item ">
            <router-link to="/games"><a class="nav-link"> Games</a></router-link>
          </li>
          <li class="nav-item">
            <router-link to="/publishing"><a class="nav-link"> Publishing</a></router-link>
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
        <img class="custom-logo-icon" src="../public/images/logo.png"/>

      </div>
      <div class="globalFooter-legal"><p>© 2019 Mantiya Entertainment. All Rights Reserved. Mantiya logo are registered or unregistered trademarks of Mantiya Entertainment in the Russia. and/or other countries.</p>
      </div>
    </footer>

  </div>
</template>

<script>
  export default {
    name: 'App',
    watch: {
      $route(to) {
        document.title = to.meta.title || 'Some Default Title';
      },
    }
  };
</script>

<style>
#app {

}
</style>
