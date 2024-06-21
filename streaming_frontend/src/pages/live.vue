<template>
  <AppBar />
  <v-container class="d-flex align-center justify-center video-container" fluid>
    <video-player :options="videoOptions" class="video-player" />
  </v-container>
</template>

<script>
import "video.js/dist/video-js.css";

export default {
  data() {
    return {
      videoOptions: {
        autoplay: true,
        controls: true,
        muted: true,
        sources: [
          {
            src: "http://localhost:3000/segment.m3u8",
            type: "application/x-mpegURL",
          },
        ],
      },
    };
  },
};
</script>

<route lang="json">
{
  "meta": {
    "requiresAuth": true
  }
}
</route>
