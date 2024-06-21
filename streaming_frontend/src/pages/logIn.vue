
<template>
  <AppBar />
  <v-container class="d-flex align-center justify-center" fluid>
    <LogInCard />
  </v-container>
</template>

<script>
import LogInCard from '@/components/LogInCard.vue'; // Ajusta la ruta según la ubicación de tu componente

export default {
  components: {
    LogInCard
  }
};
</script>
