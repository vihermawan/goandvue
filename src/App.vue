<template>
  <v-app>
    <v-app-bar app color="blue" dark>
      <v-toolbar-title>
        Vue And Go
      </v-toolbar-title>
      <v-spacer></v-spacer>
    </v-app-bar>
    <v-main>
       
    </v-main>
    <router-view></router-view>
  </v-app>
  
</template>

<script>
import HelloWorld from './components/HelloWorld';

export default {
  name: 'App',

  components: {
    HelloWorld,
  },

  data() {
    return {
      
    }
  },
};
</script>
