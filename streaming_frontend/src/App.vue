<template>
  <v-app class="background-app">
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<style>
/* Estilos globales */
.background-app {
  background-color: #212121 !important; /* Azul según el color de Material Design */
}
</style>
