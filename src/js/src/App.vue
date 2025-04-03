<!-- 
This is the main file for the Vue app. It sets the header, footer,
and body to the overall Vue window. Header and Footer are separate
Vue components. There is a dispatch that is used to check for auto
login and returns a user to Experiments component if successful.
-->

<template>
  <div @click="resetTimer" @keydown="resetTimer">
    <app-header></app-header>
    <div class="row container is-fullhd px-4">
      <div class="col-xs-12">
        <router-view></router-view>
      </div>
    </div>
    <app-footer></app-footer>
  </div>
</template>

<script>
  import Header from './components/Header.vue'
  import Footer from './components/Footer.vue'
  
  export default {
    components: {
      appHeader: Header,
      appFooter: Footer
    },
    mounted(){
      window.addEventListener('click', this.resetTimer);
      window.addEventListener('keydown', this.resetTimer);
    },
    
    beforeDestroy () {
      this.wsDisconnect();
      if ( this.unwatch ) {
        this.unwatch();
      }
      window.removeEventListener('click', this.resetTimer);
      window.removeEventListener('keydown', this.resetTimer);
    },
    
    async created () {
      try {
        let resp     = await fetch(this.$router.resolve({ name: 'features'}).href);
        let features = await resp.json();

        this.$store.commit( 'FEATURES', features['features'] );
      } catch (err) {
        console.log(`ERROR getting features: ${err}`);
      }

      try {
        let resp    = await fetch(this.$router.resolve({ name: 'options'}).href);
        let options = await resp.json();

        this.$store.commit( 'OPTIONS', options );
      } catch (err) {
        console.log(`ERROR getting options: ${err}`);
      }

      this.wsConnect();

      this.unwatch = this.$store.watch(
        ( _, getters ) => getters.token,
        () => {
          // Disconnect the websocket clients no matter what on token updates.
          this.wsDisconnect();
          this.wsConnect();
        }
      )
    },
    data () {
      return {
        socket: null,
        timeout: {
          enabled: false,
          timeout_min: 30,
          warning_min: 3,
        },
        warnToast: null,
      }
    },

    beforeUpdate(){
      this.getSettings();
    },

    methods: {
      wsConnect () {
        let path = `${process.env.BASE_URL}api/v1/ws`;

        if (this.$route.path === "/signin" || this.$route.path === "/login") {
          console.log("skipping websocket connect until login")
          return
        }

        if (this.$store.getters.token) {
          path += `?token=${this.$store.getters.token}`;
        }

        let proto = location.protocol == "https:" ? "wss://" : "ws://";
        let url   = proto + location.host + path;

        console.log("connect websocket")
        this.$connect(url);

        // Separate, stand-alone websocket connection to handle app-wide
        // notifications (e.g. new scorch terminal notifications).
        this.socket = new WebSocket(url);
        this.socket.onmessage = this.globalWsHandler;
      },

      wsDisconnect () {
        this.$disconnect();

        if ( this.socket ) {
          this.socket.close();
          this.socket = null;
        }
      },

      globalWsHandler (event) {
        event.data.split(/\r?\n/).forEach(data => {
          if (data) {
            let msg = JSON.parse(data);

            if (msg.resource.type === 'apps/scorch' && msg.resource.action === 'terminal-create') {
              this.$buefy.toast.open({
                message: `Scorch terminal created for experiment ${msg.resource.name}`,
                type:    'is-success',
                duration: 5000
              });
            }
          }
        });
      },
      getSettings(){ // TODO: this should be pushed over ws instead of only on page change
        this.$http.get('settings/timeout').then(
          response => {
            response.json().then( state => {
              this.timeout = state
            })
          }
        )
      },
      startLogoutTimer(){
        if (!this.timeout.enabled || !this.$store.getters.auth){
          return
        }
        // timeout_min: 30,
        // warning_min: 3,

        var timeout = this.timeout.timeout_min
        var warning = this.timeout.warning_min

        if (timeout <= 0) {
          timeout = 30
        }

        var diff = timeout - warning
        if (warning > 0) {
          this.logoutTimer = setTimeout(this.warnUser, 1000 * 60 * diff, warning)
        } else {
          //no warning, send straight to log out
          this.logoutTimer = setTimeout(this.logoutUser, 1000 * 60 * timeout)
        }
      },
      warnUser(timeLeft){

        var message = `Still there? Inactive auto log out in ${timeLeft} minutes.`
        if (timeLeft == 1) {
          message = `Still there? Inactive auto log out in ${timeLeft} minute.`
        }
        this.warnToast = this.$buefy.toast.open({
          message: message,
          type: 'is-warning',
          indefinite: true
        });
        this.logoutTimer = setTimeout(this.logoutUser, 1000 * 60 * timeLeft);
      },
      logoutUser(){
        if (this.warnToast) {
          this.warnToast.close();
          this.warnToast = null;
        }
        this.$http.get( 'logout' ).then(
          response => {
            if ( response.status == 204 ) {
              this.$store.commit( 'LOGOUT' );
            }
          }
        );
      },
      resetTimer() {
        if (!this.timeout.enabled) {
          return
        }
        if (this.warnToast) {
          this.warnToast.close();
          this.warnToast = null;
        }
        clearTimeout(this.logoutTimer);
        this.startLogoutTimer();
      },
    },
    watch: {
      '$route': function(to, _) {
        if (!this.socket && !(to.path === "/signin" || to.path === "/login"))
          this.wsConnect();
      }
    }
  }
</script>

<!-- 
This styling was based on some Buefy examples; there is some copied
values from the pervious phēnixweb styling. The rest was guessed at
until the window looked half way presentable. Otherwise, there is no
clue what this stuff does.
 -->

<style lang="scss">
  html {
    background-repeat: no-repeat;
    background-image: url( "assets/phenix.png" );
    background-size: background;
  }

  html, body {
    margin: 0;
    height: 100%;
  }

  body {
    padding: 20px;
    color: whitesmoke !important;
  }

  h1 {
    color: whitesmoke !important;
  }
  
  h3 {
    color: whitesmoke !important;
  }
  
  tr, td {
    color: whitesmoke !important;
  }
  
  th {
    background-color: #686868;
    color: whitesmoke !important;
  }
  
  a {
    color: whitesmoke !important;
  }
  
  p {
    color: #202020;
  }
  
  ul {
    columns: 2;
    -webkit-columns: 2;
    -moz-columns: 2;
  }
  
  li {
    color: whitesmoke !important;
  }

  #app {
    display: flex;
    flex-flow: column wrap;
    margin: 0 auto;
    height: 600px;
    justify-content: flex-start;
    align-content: flex-start;
  }

  #app > * {
    border-radius: 2px;
    transition: all ease 0.3s;
  }
  
  #app>div {
    position: relative;
    width: 200px;
    padding: 8px;
    margin: 10px;
    border: 1px solid #ccc;
  }
  
  textarea {
    display: block;
    width: 200px;
    height: 50px;
    padding: 8px;
    margin: 10px;
    border: 1px solid #ccc;
  }
  
  textarea:focus {
    border-color: black;
  }
  
  .top {
    text-align: right;
    display: flex;
    flex-direction: row-reverse;
    justify-content: space-between;
    margin-bottom: 0.5em;
  }

  .close {
    text-align: right;
    height: 10px;
    width: 10px;
    position: relative;
    box-sizing: border-box;
    line-height: 10px;
    display: inline-block;
  }
  
  .close:before, .close:after {
    transform: rotate( -45deg );
    content: "";
    position: absolute;
    top: 50%;
    left: 50%;
    margin-top: -1px;
    margin-left: -5px;
    display: block;
    height: 2px;
    width: 10px;
    background-color: black;
    transition: all 0.25s ease-out;
  }
  
  .close:after {
    transform: rotate( -135deg );
  }
  
  .close:hover:before, .close:hover:after {
    transform: rotate( 0deg );
  }

  .b-table {
    .table {
      td {
        vertical-align: middle;
      }
    }
  }

  // Import Bulma's core
  @import "~bulma/sass/utilities/_all";

  $body-background-color: #333;
  $table-background-color: #484848;
  $table-row-hover-background-color: #777777;
  
  $button-text-color: whitesmoke;


  $breadcrumb-item-color: $info !important;
  $breadcrumb-item-active-color: $light-invert !important;

  $light: #686868;
  $light-invert: findColorInvert( $light );

  $progress-text-color: black;

  $fullhd: 1536px + (2 * $gap);

  $colors: (
    "light": ( $light, $light-invert ),
    "dark": ( $dark, $dark-invert ),
    "white": ( $white, $black ),
    "black": ( $black, $white ),
    "primary": ( $primary, $primary-invert ),
    "info": ( $info, $info-invert ),
    "success": ( $success, $success-invert ),
    "warning": ( $warning, $warning-invert ),
    "danger": ( $danger, $danger-invert )
  );
  
  // Import Bulma and Buefy styles
  @import "~bulma";
  @import "~buefy/src/scss/buefy";

  a.navbar-item:hover {
    background: #404040;
  }

  div.is-success {
    color: $success;
  }
</style>
