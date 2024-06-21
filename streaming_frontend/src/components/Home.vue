<template>
  <v-container justify="center" align="center">
    <v-row justify="center" align="center" no-gutters>
      <v-col cols="4" class="text-center">
        <VideoCat v-if="showCat" />
        <h1 class="big-title">
          <v-icon icon="mdi-play-circle-outline"></v-icon>Kirora
        </h1>
        <v-btn to="/logIn" color="purple-darken-4">Inicia sesión</v-btn>
        <p>
          ¿No tienes una cuenta?
          <router-link to="/register"> Registrate </router-link>
        </p>
        <v-btn icon flat @click="showDialog" v-if="!showCat">
          <v-icon>mdi-gift</v-icon>
        </v-btn>
        <v-btn icon flat @click="showCat = false" v-if="showCat">
          <v-icon>mdi-close</v-icon>
        </v-btn>
      </v-col>
    </v-row>
    <v-dialog v-model="dialog" width="auto">
      <v-card
        max-width="800"
        prepend-icon="mdi-party-popper"
        title="Gracias por el interés!"
        ><template v-slot:text>
          ¡Nos alegra mucho que estés explorando nuestra aplicación! Tenemos un
          regalo especial para ti.
          <br />
          Todo lo que necesitas hacer es hacer clic en el botón 'Ver'.
          <br />
          Te recomendamos encender el sonido para disfrutar al máximo de la
          experiencia.
        </template>
        <template v-slot:actions>
          <v-btn class="ms-auto" text="Salir" @click="dialog = false"></v-btn>
          <v-btn color="purple-darken-4" text="Ver" variant="flat" @click="showVideoCat"></v-btn>
        </template>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      dialog: false,
      showCat: false,
    };
  },

  methods: {
    showDialog() {
      this.dialog = true;
    },
    showVideoCat(){
      this.dialog = false;
      this.showCat = true;
    }
  }
};
</script>
