<template>
  <v-container justify="center" align="center" >
    <v-row justify="center" align="center" no-gutters>
      <v-col cols="4" class="text-center ">
        <h1 class="big-title"><v-icon icon="mdi-play-circle-outline"></v-icon>Kirora</h1>
        <v-btn to="/logIn" color="purple-darken-4" >Inicia sesión</v-btn>
        <p>¿No tienes una cuenta? <router-link to="/register"> Registrate </router-link></p>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
</script>

