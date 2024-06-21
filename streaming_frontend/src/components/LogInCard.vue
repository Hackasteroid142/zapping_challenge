<template>
  <v-card class="mx-auto" width="600" color="white">
    <template v-slot:title>
      <h2 class="purple-text">Inicia Sesión</h2>
    </template>

    <v-card-text class="pt-4">
      <v-text-field v-model="email" label="Correo"></v-text-field>
      <v-text-field v-model="password" label="Contraseña"></v-text-field>
      <div class="text-center mt-4">
        <v-btn color="purple-darken-4" @click="logIn" :disabled="!formFilled">Ingresar</v-btn>
      </div>
    </v-card-text>
    <v-alert v-model="showAlert" class="error-card" text="Usuario no existe" type="error"></v-alert>
    <v-card-actions class="text-center justify-center">
      <p>
        ¿No tienes una cuenta?
        <router-link to="/register">Regístrate</router-link>
      </p>
    </v-card-actions>
  </v-card>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
import router from "@/router";

export default {
  data() {
    return {
      email: "",
      password: "",
      showAlert: false
    };
  },

  computed: {
    ...mapGetters(["getToken"]),
    formFilled() {
      return this.email !== "" && this.password !== "";
    },
  },

  methods: {
    ...mapActions(["logInUser"]),
    async logIn() {
      try {
        const body = {
          email: this.email,
          password: this.password,
        };
        await this.logInUser(body);
        if (this.getToken) {
          router.push("/live");
        }
      } catch (error) {
        console.log(error.response);
        this.showAlert = true;
      }
    },
  },
};
</script>
