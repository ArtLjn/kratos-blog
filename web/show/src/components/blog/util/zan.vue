<template>
    <div>
      <button @click="toggleLike">
        {{ liked ? '取消点赞' : '点赞' }}
      </button>
      <p>点赞数: {{ likeCount }}</p>
    </div>
  </template>
  
  <script>
  import { ref } from 'vue';
  
  export default {
    setup() {
      // 使用 ref 创建响应式数据
      const likeCount = ref(0);
      const liked = ref(false);
  
      // 点赞和取消点赞的方法
      const toggleLike = () => {
        if (liked.value) {
          likeCount.value--;
        } else {
          likeCount.value++;
        }
        liked.value = !liked.value;
      };
  
      return {
        likeCount,
        liked,
        toggleLike
      };
    }
  };
  </script>