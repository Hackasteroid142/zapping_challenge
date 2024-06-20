<script>
import { ref } from "vue";
import VideoPlayer from "./VideoPlayer.vue";
import "video.js/dist/video-js.css";
import axios from "axios";

export default {
  components: {
    VideoPlayer,
  },
  data() {
    return {
      videoOptions: {
        autoplay: true,
        controls: false,
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

const count = ref(0);
</script>

<template>
  <div>
    <video-player :options="videoOptions" />
  </div>

  <div class="card">
    <p>
      Edit
      <code>components/HelloWorld.vue</code> to test HMR
    </p>
  </div>

  <p>
    Check out
    <a href="https://vuejs.org/guide/quick-start.html#local" target="_blank"
      >create-vue</a
    >, the official Vue + Vite starter
  </p>
  <p>
    Install
    <a href="https://github.com/vuejs/language-tools" target="_blank">Volar</a>
    in your IDE for a better DX
  </p>
  <p class="read-the-docs">Click on the Vite and Vue logos to learn more</p>
</template>

<style scoped>
.read-the-docs {
  color: #888;
}
</style>
