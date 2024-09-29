<template>
  <div class="login-container">
    <el-form :model="loginForm" @submit.native.prevent="handleLogin">
      <h2>H-Precision 平台登录</h2>
      <el-form-item prop="username">
        <el-input
          v-model="loginForm.username"
          placeholder="用户名"
        ></el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input
          type="password"
          v-model="loginForm.password"
          placeholder="密码"
        ></el-input>
      </el-form-item>
      <el-button type="primary" native-type="submit">登录</el-button>
    </el-form>
  </div>
</template>

<script>
export default {
  name: 'Login',
  data() {
    return {
      loginForm: {
        username: '',
        password: '',
      },
    };
  },
  methods: {
    handleLogin() {
      // TODO: 调用后端 API 进行登录验证
      // 假设验证成功，保存用户信息并跳转
      this.$router.push('/home');
    },
  },
};
</script>

<style scoped>
.login-container {
  width: 400px;
  margin: 100px auto;
  text-align: center;
}
</style>
