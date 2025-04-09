<!-- 
This component is used to initially sign in a user to the phēnixweb. 
It requires a valid email (user ID) and password.
 -->

<template>
  <div id="signin">
    <b-modal :active.sync="signUpModal" has-modal-card>
      <div class="modal-card" style="width: auto">
        <header class="modal-card-head">
          <p class="modal-card-title">Create a New Account</p>
        </header>
        <section class="modal-card-body">
          <b-field label="User Name" 
            :type="{ 'is-danger' : userExists }" 
            :message="{ 'User already exists' : userExists }">
            <b-input type="text" v-model="new_username" minlength="4" maxlength="32" autofocus></b-input>
          </b-field>
          <b-field label="First Name">
            <b-input type="text" v-model="first_name"></b-input>
          </b-field>
          <b-field label="Last Name">
            <b-input type="text" v-model="last_name"></b-input>
          </b-field>
          <b-field label="Password"
            :message="createModalErrors.passwordErrorMessage"
            :type="createModalErrors.passwordErrorLevel">
            <b-input type="password" minlength="8" maxlength="32" v-model="new_password"></b-input>
          </b-field>
          <b-field label="Confirm Password">
            <b-input type="password" minlength="8" maxlength="32" v-model="confirmPassword" @keyup.native.enter="create"></b-input>
          </b-field>
        </section>
        <footer class="modal-card-foot buttons is-right">
          <button class="button is-light"
            @click="create"
            :disabled="!check_password_validity()">Create User</button>
        </footer>
      </div>
    </b-modal>
    <div class="signin-form">
      <b-field label="Username">
        <b-input type="text" v-model="username" autofocus></b-input>
      </b-field>
      <b-field label="Password">
        <b-input type="password" v-model="password" @keyup.native.enter="onSubmit"></b-input>
      </b-field>
      <b-checkbox type="is-light" size="is-small" v-model="rememberMe" native-value=false>Remember me</b-checkbox>
      <br>
      <button class="button is-light" @click="onSubmit">Submit</button>
      <button class="button is-pulled-right is-small is-text" @click="signUpModal = true">Create Account</button>
    </div>
  </div>
</template>

<script>
  export default {
    async created() {
      this.$http.get('settings/password').then(
        response => {
          response.json().then( state => {
            console.log(state);
            this.passwordReqs = state
          });
        }
      )
    },
    methods: {
      //  this method is called when the Submit button is pressed (or 
      //  return key is) executed. It will check that an email address 
      //  is used, and/or a password. It does not check if they are valid. 
      onSubmit () {
        if ( !this.username ) {
          this.$buefy.toast.open({
            message: 'You must include an email address',
            type: 'is-warning',
            duration: 4000
          });
          return {}
        }
        
        if ( !this.password ) {
          this.$buefy.toast.open({
            message: 'You must include a password',
            type: 'is-warning',
            duration: 4000
          });
          
          return {}
        }

        this.$http.post(
          'login', {'user': this.username, 'pass': this.password}
        ).then(
          response => {
            response.json().then(
              user => {
                this.$store.commit( 'LOGIN', { "loginResponse": user, "remember": this.rememberMe } );
              }
            )
          }, response => {
            if ( response.status == 401 ) {
              this.$buefy.toast.open({
                message: 'The username and/or password is incorrect',
                type: 'is-warning',
                duration: 4000
              });
              
              this.username = null;
              this.password = null;
            } else if ( response.status == 0 ) {
              this.$buefy.toast.open({
                message: 'The data server is not available.',
                type: 'is-danger',
                duration: 6000
              });
              
              this.username = null;
              this.password = null;
            } else {
              this.$buefy.toast.open({
                message: 'Getting the user information failed.',
                type: 'is-danger',
                duration: 4000
              });
              
              this.username = null;
              this.password = null;
            }
          }
        );     
      },
      
      create () {
        if ( !this.username ) {
          this.$buefy.toast.open({
            message: 'You must include an username',
            type: 'is-warning',
            duration: 4000
          });
          
          return {}
        }
        
        if ( !this.first_name ) {
          this.$buefy.toast.open({
            message: 'You must include a first name',
            type: 'is-warning',
            duration: 4000
          });
          
          return {}
        }
        
        if ( !this.last_name ) {
          this.$buefy.toast.open({
            message: 'You must include a last name',
            type: 'is-warning',
            duration: 4000
          });
          
          return {}
        }
        
        if ( !this.password ) {
          this.$buefy.toast.open({
            message: 'You must include a password',
            type: 'is-warning',
            duration: 4000
          });
          
          return {}
        }
        
        if ( !this.confirmPassword ) {
          this.$buefy.toast.open({
            message: 'You must include a password confirmation',
            type: 'is-warning',
            duration: 4000
          });
          
          return {}
        }
        
        if ( this.password != this.confirmPassword ) {
          this.$buefy.toast.open({
            message: 'Your passwords do not match',
            type: 'is-warning',
            duration: 4000
          });
          
          return {}
        }       

        if (! this.check_password()) {
          //check_password handles the toasts
          return {}
        }

        this.$http.post(
          'signup', {
            "username": this.username,
            "password": this.password,
            "first_name": this.first_name,
            "last_name": this.last_name
          }
        ).then(
          response => { 
            return response.json().then(
              user => {
                this.$store.commit( 'SIGN_UP', user );
              }
            )
          }, err => {
            this.errorNotification(err);
            this.username = null;
            this.password = null;
          }
        );

        this.signUpModal = false;
      },
      check_password_validity() {
        let password = this.password
        if (password == undefined) {
          return true
        }
        if (password.length == 0 ){
          //don't want errors immediately on modal if they haven't typed in anything yet
          this.createModalErrors.passwordErrorMessage = null
          this.createModalErrors.passwordErrorLevel = null
          return true
        }
        if (password.length < this.passwordReqs.min_length) {
          this.createModalErrors.passwordErrorMessage = "Password must be longer than " + this.passwordReqs.min_length + " characters."
          this.createModalErrors.passwordErrorLevel = "is-danger"
          return false
        }
        if (! /[a-z]/.test(password) && this.passwordReqs.lowercase_req) {
          this.createModalErrors.passwordErrorMessage = "Password must contain a lowercase letter"
          this.createModalErrors.passwordErrorLevel = "is-danger"
          return false
        }
        if (! /[A-Z]/.test(password) && this.passwordReqs.uppercase_req) {
          this.createModalErrors.passwordErrorMessage = "Password must contain an uppercase letter"
          this.createModalErrors.passwordErrorLevel = "is-danger"
          return false
        }
        if (! /\d/.test(password) && this.passwordReqs.number_req) {
          this.createModalErrors.passwordErrorMessage = "Password must contain a number"
          this.createModalErrors.passwordErrorLevel = "is-danger"
          return false
        }
        if (! /\W/.test(password) && this.passwordReqs.symbol_req) {
          this.createModalErrors.passwordErrorMessage = "Password must contain a symbol"
          this.createModalErrors.passwordErrorLevel = "is-danger"
          return false
        }

        this.createModalErrors.passwordErrorMessage = null
        this.createModalErrors.passwordErrorLevel = null
        return true
      },
      getPasswordRequirements(){
        this.$http.get('settings/password').then(
          response => {
            response.json().then( state => {
              console.log(state);
              this.passwordReqs = state
            });
          }
        )
      },
    },

    data () {
      return {
        signUpModal: false,
        username: null,
        password: null,

        new_username: null,
        new_password: null,
        confirmPassword: null,
        first_name: null,
        last_name: null,
        rememberMe: false,
        userExists: false,
        passwordReqs: {
          number_req: false,
          symbol_req: false,
          lowercase_req: false,
          uppercase_req: false,
          min_length: 8,
        },
        createModalErrors:{
          passwordErrorMessage: null,
          passwordErrorLevel: null
        },
      }
    }
  }
</script>

<!-- This styling is used for the sign in form. -->
<style scoped>
  .signin-form {
    width: 400px;
    margin: 30px auto;
    border: 1px solid #eee;
    padding: 20px;
    box-shadow: 0 2px 3px #ccc;
  }
  
  .signin-form >>> .label {
    color: whitesmoke;
  }

  label.checkbox:hover {
    color: whitesmoke;
  }
</style>
