<template>
    <form v-if="!showRegisterForm && !showResetForm" class="form" @submit.prevent="login">
      <h1>登录</h1>
      <div class="input-box">
        <label for="username">Username:</label>
        <input type="text" id="username" name="username" v-model="loginForm.name" required>

        <label for="password">Password:</label>
        <input type="password" id="password" name="password" v-model="loginForm.pass" required>
      </div>
      <button type="submit">Log In</button>
      <button type="button" @click="showRegisterForm = true">Register</button>
      <button type="button" @click="showResetForm = true,showRegisterForm = false,handleResetForm"
      style="background-color:rgb(239, 239, 248);">忘记密码？</button><br>
    </form>
    
    <form v-if="showRegisterForm" class="form" @submit.prevent="register">
      <h1>注册</h1>
      <div class="input-box">
        <label for="username">Username:</label>
        <input type="text" id="username" name="username"
          v-model="registerForm.name" required>

        <label for="email">Email:</label>
          <input type="email" id="email" name="email" 
          pattern="[a-zA-Z0-9_-]+@qq\.com"
          placeholder="support@qq.com"
          v-model="registerForm.email" required>

        <label for="password">Password:</label>
        <input type="password" id="password" name="password" v-model="registerForm.password" required>

        <label for="verificationCode">Verification Code:</label>
        <div class="verification-code-input">
          <input type="text" id="verificationCode" name="verificationCode" v-model="code" required>
          <el-button type="submit" @click="sendEmail"
            :plain="true">{{content}}</el-button>
        </div>
      </div>
      <button type="submit">Register</button>
      <button type="button" @click="showRegisterForm = false">Cancel</button>
    </form>


    <form v-if="showResetForm" class="form" @submit.prevent="UpdateUser"  id="resetForm">
      <h1>重置</h1>
      <div class="input-box">
        <label for="username">Username:</label>
        <input type="text" id="username" name="username" v-model="registerForm.name" required>

        <label for="email">Email:</label>
          <input type="email" id="email" name="email"
          pattern="[a-zA-Z0-9_-]+@qq\.com"
          placeholder="support@qq.com"
            v-model="registerForm.email" required>

        <label for="password">Password:</label>
        <input type="password" id="password" name="password" v-model="registerForm.password" required>

        <label for="verificationCode">Verification Code:</label>
        <div class="verification-code-input">
          <input type="text" id="verificationCode" name="verificationCode" v-model="code" required>
          <el-button type="submit" @click="sendEmail"
            :plain="true">{{content}}</el-button>
        </div>
      </div>
      <button type="submit" >Reset</button>
      <button type="button" @click="showResetForm = false">Cancel</button>
    </form>
  </template>
  
  <script>  
  import { onMounted, ref } from 'vue';
  import { ElMessage } from 'element-plus';
  import {Login, Register, SendEmail, UpdatePassword} from "@/api/auth";
   
  export default {
    name: 'LoginComponent',
    setup() {

      const content = ref('发送验证码')
      const totalTime = ref(60);
      const lockClick = ref(true);
      const showRegisterForm = ref(false);
      const showResetForm = ref(false);
      const loginForm = ref({
        name: '',
        pass: ''
      });
      const registerForm = ref({
        name: '',
        email: '',
        password: '',
      });
      const code = ref('')
      const emailSent = ref(false);
      onMounted(() => {
        if (localStorage.getItem("token")) {
          login()
        }
      });
      
      const login = () => {
          Login(loginForm.value)
      };
  
      const register = () => {
        Register(registerForm.value,code.value)
      };
      
      const sendEmail = () => {
        if (registerForm.value.email) {
          if (!lockClick.value) return
          lockClick.value = false
          content.value = totalTime.value + 's后重新发送'
          let clock = window.setInterval(() => {
            totalTime.value--;
            content.value = totalTime.value + 's后重新发送'
            if (totalTime.value < 0) {
              window.clearInterval(clock)
              content.value = '重新发送验证码'
              totalTime.value = 60
              lockClick.value = true
            }
          }, 1000)
        } else {
          ElMessage({
            message: "请输入邮箱",
            type: "warning",
          })
          return;
        }
        SendEmail(registerForm.value.email)
      };

      const UpdateUser = () => {
        UpdatePassword(registerForm.value,code.value)
      }
      const handleResetForm = () => {
      const resetFormElement = document.getElementById('resetForm');
        resetFormElement.scrollIntoView({ behavior: 'smooth' });
      };


      return {
        showRegisterForm,
        loginForm,
        registerForm,
        emailSent,
        content,
        totalTime,
        login,
        register,
        sendEmail,
        handleResetForm,
        showResetForm,
        UpdateUser,
        code
      };
    }
    
  }
  </script>
  <style scoped>
  .background-image {
    position: absolute;
    position: fixed;
    top: 0px;
    left: 0px;
    height: 100%;
    width: 100%;
    z-index: 998;
  }
  
  .box {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100%;
    background-color: rgb(161, 139, 139,0);
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
    width: 100%;
    height: auto;
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
  