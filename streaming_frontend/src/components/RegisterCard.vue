<template>
  <v-card class="mx-auto" width="600" color="white">
    <template v-slot:title>
      <h2 class="purple-text">Registrate</h2>
    </template>

    <v-card-text class="pt-4">
      <v-form>
        <v-text-field v-model="name" label="Nombre"></v-text-field>
        <v-text-field v-model="email" label="Correo"></v-text-field>
        <v-text-field v-model="password" label="Contraseña"></v-text-field>
        <div class="text-center mt-4">
          <v-btn type="submit" color="purple-darken-4">Registrarme</v-btn>
        </div>
      </v-form>
    </v-card-text>
    <v-card-actions class="text-center justify-center">
      <p>
        ¿Ya tienes una cuenta?
        <router-link to="/logIn">Inicia sesión</router-link>
      </p>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  data() {
    return {
      email: "",
      password: "",
      name: ""
    };
  },
};
</script>
