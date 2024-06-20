<script>
import VideoPlayer from "./VideoPlayer.vue";
import "video.js/dist/video-js.css";

export default {
  components: {
    VideoPlayer,
  },
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

<template>
  <div>
    <video-player :options="videoOptions" />
  </div>
</template>

<style scoped>
.read-the-docs {
  color: #888;
}
</style>
