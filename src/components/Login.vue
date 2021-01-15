<template>
  <v-container class="fill-height" fluid>
    <v-row align="center" justify="center">
      <v-col cols="12" sm="8" md="4">
        <v-card class="elevation-12">
          <v-toolbar color="primary" dark flat>
            <v-toolbar-title>Login</v-toolbar-title>
          </v-toolbar>
          <v-card-text>
            <v-alert v-if="message" type="error">
              {{message.error}}
            </v-alert>
            <v-form @submit.prevent="handleSubmit">
              <v-text-field 
                v-model="user.email"
                label="email" 
                name="email" 
                prepend-icon="mdi-account" 
                type="email"
                required>
              </v-text-field>

              <v-text-field 
                v-model="user.password"
                id="password" 
                label="Password" 
                name="password" 
                prepend-icon="mdi-lock" 
                type="password"
                required>
              </v-text-field>
              <v-card-actions>
                <v-container>
                  <v-row>
                    <v-col cols="12" sm="8" md="8">
                      <span>Don't have an account? Register <router-link to="/register">here.</router-link></span>
                    </v-col>
                    <v-col cols="12" sm="8" md="4" align="right">
                      <v-btn type="submit" color="primary">Login</v-btn>
                    </v-col>
                  </v-row>
                </v-container>
              </v-card-actions>
            </v-form>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import User from '../models/user';
  export default {
    name: 'Login',
    data() {
      return {
        user: new User('', ''),
        loading: false,
        message: ''
      };
    },
    computed: {
      loggedIn() {
        return this.$store.state.auth.status.loggedIn;
      }
    },
    created() {
      if (this.loggedIn) {
        this.$router.push('/home');
      }
    },
    methods: {
      handleSubmit() {
        console.log("test", this.user.email, "test")
        if (this.user.email && this.user.password) {
          this.$store.dispatch('auth/login', this.user).then(
            () => {
              this.$router.push('/home');
            },
            error => {
              this.loading = false;
              this.message =
                (error.response && error.response.data) ||
                error.message ||
                error.toString();
            }
          );
        }
      }
    }
  }
</script>