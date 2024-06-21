<template>
  <v-app-bar :elevation="0" color="grey-darken-4">
    <v-app-bar-title class="app-title"
      ><v-icon icon="mdi-play-circle-outline"></v-icon>Kirora</v-app-bar-title
    >
    <v-spacer></v-spacer>
    <p v-if="getToken">Hola {{ this.getUser.name }}!</p>
    <v-btn icon @click="exitApp" v-if="getToken">
      <v-icon>mdi-logout</v-icon>
    </v-btn>
    <v-btn icon to="/" v-if="!getToken">
      <v-icon>mdi-home</v-icon>
    </v-btn>
  </v-app-bar>
</template>
<script>
import { mapGetters, mapActions } from "vuex";
import router from "@/router";

export default {
  computed: {
    ...mapGetters(["getToken", "getUser"])
  },

  methods: {
    ...mapActions(["logout"]),
    exitApp() {
      this.logout();
      router.push("/");
    },
  },
};
</script>
