<template>
  <div class="background-image">
    <div class="box">
      <form  class="form" @submit.prevent="logins">
        <h1>登录</h1>
        <div class="input-box">
          <label for="username">Username:</label>
          <input type="text" id="username" name="username" v-model="loginForm.name" required>
  
          <label for="password">Password:</label>
          <input type="password" id="password" name="password" v-model="loginForm.password" required>
        </div>
        <button type="submit">Log In</button>
      </form>
    </div>
  </div>
  </template>
  
  <script>
  import { onMounted,ref } from 'vue';
  import {login} from "@/components/api/auth";

  export default {
    // eslint-disable-next-line vue/multi-word-component-names
    name: 'Login',
    setup() {
      const loginForm = ref({
        name: '',
        password: ''
      });
      onMounted(() => {
        login(loginForm)
      });

      const logins = () => {
        login(loginForm.value)
      };

      return {
        logins,
        loginForm,
      };
    }
  };
  </script>
  
  <style scoped>
  .background-image {
    background-image: url("../../assets/img/bg_01.jpg");
    position: fixed;
    top: 0;
    left: 0;
    height: 100%;
    width: 100%;
    z-index: 1;
  }
  
  .box {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100%;
    animation: slide-in-blurred 0.5s both;
  }
  
  @keyframes slide-in-blurred {
    0% {
      transform: translateY(-1000px) scaleY(2.5) scaleX(0.2);
      transform-origin: 50% 0%;
      filter: blur(40px);
      opacity: 0;
    }
    100% {
      transform: translateY(0px) scaleY(1) scaleX(1);
      transform-origin: 50% 50%;
      filter: blur(0px);
      opacity: 1;
    }
  }
  
  .form {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    width: 400px;
    height: auto;
    background-color: rgb(247, 243, 243,0.7);
    padding: 100px 50px;
    border-radius: 30px;
  }
  
  form h1 {
    margin: 10px;
    font-size: 28px;
    text-align: center;
    color: #283789;
  }
  
  .input-box {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    width: 100%;
  }
  
  input {
    margin: 10px;
    padding: 12px 20px;
    width: 100%;
    border: none;
    border-radius: 25px;
    box-sizing: border-box;
    background-color: #f7f7f7;
    font-size: 16px;
    color: #333333;
  }
  button[type="submit"] {
    margin: 10px;
    margin-top: 20px;
    padding: 12px 20px;
    width: 100%;
    border: none;
    border-radius: 25px;
    box-sizing: border-box;
    background-color: #283789;
    font-size: 16px;
    color: #ffffff;
    cursor: pointer;
    transition: background-color 0.3s ease-out;
    }
    
    button[type="submit"]:hover {
    background-color: #4c5aa9;
    }
    
    button[type="button"] {
    margin: 10px;
    margin-top: 0px;
    padding: 12px 20px;
    width: 100%;
    border: none;
    border-radius: 25px;
    box-sizing: border-box;
    background-color: #f7f7f7;
    font-size: 16px;
    color: #333333;
    cursor: pointer;
    transition: background-color 0.3s ease-out;
    }
    
    button[type="button"]:hover {
    background-color: #e4e4e4;
    }
    
    a {
    margin-top: 10px;
    text-align: center;
    font-size: 16px;
    color: #283789;
    text-decoration: none;
    }
  
    .verification-code-input {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
    margin: 10px;
    }
    @media screen and (max-width:500) {
      .form{
        width: 300px;
      }
    }
  </style>