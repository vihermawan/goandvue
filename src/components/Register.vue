<template>
  <v-container class="fill-height" fluid>
  <v-dialog
        v-model="dialog"
        max-width="290"
      >
        <v-card>
          <v-card-title class="headline">Account created!</v-card-title>

          <v-card-text>
            Congratulation|, account has created successfully.
            Login!
          </v-card-text>

          <v-card-actions>
            <v-spacer></v-spacer>

            <v-btn
              color="green darken-1"
              text
              @click="goLogin"
            >
              OK
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    <v-row align="center" justify="center">
      <v-col cols="12" sm="8" md="4">
        <v-card class="elevation-12">
          <v-toolbar color="primary" dark flat>
            <v-toolbar-title>Register</v-toolbar-title>
          </v-toolbar>
          <v-card-text>
            <v-alert v-if="message" type="error">
                  {{message.error}}
            </v-alert>
            <v-form @submit.prevent="handleSubmit">
              <v-text-field 
                v-model="user.name"
                label="Nama" 
                name="name" 
                prepend-icon="mdi-account" 
                type="text"
                required>
              </v-text-field>

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
                    <span>Already have account? Login <router-link to="/">here.</router-link></span>
                  </v-col>
                  <v-col cols="12" sm="8" md="4" align="right">
                    <v-btn type="submit" color="primary">Register</v-btn>
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
    name: 'Register',
    data() {
      return {
        user: new User('', '', ''),
        submitted: false,
        successful: false,
        message: '',
        dialog: false
      }
    },
    methods : {
      handleSubmit(){
        this.message = '';
        this.submitted = true;
        console.log("test")
        this.$store.dispatch('auth/register', this.user).then(
        () => {
          this.$router.push('/');
        },
        error => {
          if (error.response === undefined) {
            this.dialog = true;
          } else {
            this.message =
              (error.response && error.response.data) ||
              error.message ||
              error.toString();
            this.successful = false;
          }
        }
      );
        
      },
      goLogin() {
        this.dialog = false;
        this.$router.push('/');
      }
    }
  }
</script>
