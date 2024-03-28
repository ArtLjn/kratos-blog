<template>
    <div>
      <audio ref="audioElement" controls autoplay name="media" style="float:right;" class="audio">
        <source :src="currentTrack" type="audio/mpeg">
      </audio>
    </div>
  </template>
  
  <script>
  import { ref, onMounted } from 'vue';
  export default {
    setup() {
      const currentTrack = ref('');
  
      const getRandomMusic = () => {
        fetch("https://api.uomg.com/api/rand.music?sort=热歌榜&format=json")
          .then((response) => response.json())
          .then((data) => {
            const url = data.data.url;
            playMusic(url);
          });
      };
  
      const playMusic = (url) => {
        const audio = audioElement.value; // 使用ref获取音频元素
        audio.src = url;
        audio.play();
        currentTrack.value = url;
      };
  
      const nextTrack = () => {
        getRandomMusic();
      };
  
      const audioElement = ref(null); // 创建ref
  
      onMounted(() => {
        // const audioPlayer = audioElement.value; // 在组件渲染后再获取音频元素
        // audioPlayer.addEventListener('ended', nextTrack);
        // getRandomMusic();
      });
  
      return {
        currentTrack,
        audioElement // 将ref返回，以便在模板中访问DOM元素
      };
    }
  }
  </script>
  