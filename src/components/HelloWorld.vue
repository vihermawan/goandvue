<template>
  <v-app id="inspire" >
    <v-app-bar app clipped-left color="blue">
      <v-toolbar-title dark>Vue and Go</v-toolbar-title>
    </v-app-bar>
    <v-navigation-drawer
      app
      clipped
    >
      <v-list dense>
        <v-list-item link>
          <v-list-item-action>
            <v-icon>mdi-account</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>name</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-list-item link>
          <v-list-item-action>
            <v-icon>mdi-logout</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title @click.prevent="logOut">Log Out</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-main>
      <v-container
        class="fill-height"
        fluid
      >
        <v-row
          align="center"
          justify="center">
          <v-col class="shrink">
            <h3 style="text-align:center">HelloWorld</h3>
          </v-col>
        </v-row>
      </v-container>
    </v-main>

    <v-footer app>
      <span>&copy; {{ new Date().getFullYear() }}</span>
    </v-footer>
  </v-app>
</template>

<script>
  export default {
    props: {
      source: String,
    },
    data: () => ({
      
    }),
    methods: {
      handleLogin() {
        console.log('test', this.$store.state.user);
      },
      logOut() {
        this.$store.dispatch('auth/logout');
        this.$router.push('/');
      }
    }
  }
</script>

