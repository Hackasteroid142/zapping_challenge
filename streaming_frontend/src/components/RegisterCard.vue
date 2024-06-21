<template>
  <v-card class="mx-auto" width="600" color="white">
    <template v-slot:title>
      <h2 class="purple-text">Registrate</h2>
    </template>

    <v-card-text class="pt-4">
      <v-text-field v-model="name" label="Nombre"></v-text-field>
      <v-text-field v-model="email" label="Correo"></v-text-field>
      <v-text-field v-model="password" label="Contraseña"></v-text-field>
      <div class="text-center mt-4">
        <v-btn color="purple-darken-4" @click="newUser" :disabled="!formFilled"
          >Registrarme</v-btn
        >
      </div>
    </v-card-text>
    <div v-if="showAlert">
      <v-alert
        class="error-card"
        text="Usuario existente"
        type="error"
      ></v-alert>
      <div class="image-container">
        <v-img
          :width="350"
          aspect-ratio="16/9"
          cover
          src="https://http.cat/409"
        ></v-img>
      </div>
    </div>
    <v-card-actions class="text-center justify-center">
      <p>
        ¿Ya tienes una cuenta?
        <router-link to="/logIn">Inicia sesión</router-link>
      </p>
    </v-card-actions>
  </v-card>
</template>

<script>
import { mapActions } from "vuex";
import router from "@/router";

export default {
  data() {
    return {
      email: "",
      password: "",
      name: "",
      showAlert: false,
    };
  },

  computed: {
    formFilled() {
      return this.name !== "" && this.email !== "" && this.password !== "";
    },
  },

  methods: {
    ...mapActions(["createUser"]),
    async newUser() {
      try {
        const body = {
          email: this.email,
          password: this.password,
          name: this.name,
        };
        const result = await this.createUser(body);
        if (result) {
          router.push("/logIn");
          this.showAlert = false;
        }
        console.log(result);
      } catch (error) {
        console.log(error.response);
        this.showAlert = true;
      }
    },
  },
};
</script>
