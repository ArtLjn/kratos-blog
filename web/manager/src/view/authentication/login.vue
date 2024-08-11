<template>
  <div class="background">
    <div class="login-container">
      <div class="logo-container">
        <img src="../../assets/img/logo.png" alt="Logo" class="logo">
        <h2>红豆南墙</h2>
      </div>
      <form class="login-form" @submit.prevent="logins">
        <div class="input-box">
          <input placeholder="username" type="text" id="username" name="username" v-model="loginForm.name" required>
        </div>
        <div class="input-box">
          <input placeholder="password" type="password" id="password" name="password" v-model="loginForm.password" required>
        </div>
        <button type="submit">Log In</button>
      </form>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue';
import { login } from "@/view/api/auth";

export default {
  setup() {
    const loginForm = ref({
      name: '',
      password: ''
    });

    const logins = () => {
      login(loginForm.value);
    };

    return {
      logins,
      loginForm,
    };
  }
};
</script>

<style scoped>
/* 背景样式 */
.background {
  background-image: url("../../assets/img/bg_01.jpg");
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  position: fixed;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  backdrop-filter: blur(8px); /* 添加模糊效果 */
}

/* 登录框样式 */
.login-container {
  background: rgba(255, 255, 255, 0.85); /* 半透明背景 */
  padding: 40px;
  border-radius: 15px;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2); /* 增加阴影效果 */
  width: 100%;
  max-width: 400px;
  box-sizing: border-box;
  backdrop-filter: blur(5px); /* 背景滤镜 */
  text-align: center; /* 居中对齐内容 */
}

/* Logo 样式 */
.logo-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 20px;
}

.logo {
  width: 80px;
  height: 80px;
  margin-bottom: 10px;
  border-radius: 30px;
}

.logo-container h2 {
  font-size: 24px;
  color: #333;
  font-weight: bold;
}

/* 标题样式 */
.login-form h1 {
  margin-bottom: 30px;
  font-size: 26px;
  text-align: center;
  color: #333;
  font-weight: bold;
}

/* 输入框样式 */
.input-box {
  margin-bottom: 20px;
}

label {
  font-size: 14px;
  color: #555;
  display: block;
  margin-bottom: 8px;
}

input {
  width: 100%;
  padding: 12px;
  border: none;
  border-radius: 25px; /* 圆角 */
  font-size: 14px;
  box-sizing: border-box;
  background-color: #f1f1f1;
  box-shadow: inset 2px 2px 5px rgba(0, 0, 0, 0.1); /* 内阴影 */
}

/* 按钮样式 */
button[type="submit"] {
  width: 100%;
  padding: 12px;
  border: none;
  border-radius: 25px;
  background: linear-gradient(135deg, #6b73ff 0%, #000dff 100%); /* 渐变背景 */
  color: #fff;
  font-size: 16px;
  font-weight: bold;
  cursor: pointer;
  transition: background 0.3s ease, transform 0.2s ease; /* 添加过渡效果 */
}

button[type="submit"]:hover {
  background: linear-gradient(135deg, #5a63d1 0%, #0009b3 100%);
  transform: translateY(-3px); /* 轻微向上移动 */
}

/* 响应式调整 */
@media screen and (max-width: 500px) {
  .login-container {
    padding: 30px 20px;
  }
}
</style>
