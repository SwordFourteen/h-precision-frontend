<template>
  <div id="app">
    <el-container>
      <el-header>
        <div class="logo">H-Precision 平台</div>
        <el-menu
          v-if="$route.path !== '/'"
          :default-active="$route.path"
          class="el-menu-demo"
          mode="horizontal"
          @select="handleSelect"
        >
          <el-menu-item index="/home">首页</el-menu-item>
          <el-menu-item index="/tool-strategy">策略管理</el-menu-item>
          <el-menu-item index="/model-tuning">模型调优</el-menu-item>
          <el-menu-item index="/vulnerability-detection">漏洞检测</el-menu-item>
        </el-menu>
      </el-header>
      <el-main>
        <router-view />
      </el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  name: 'App',
  methods: {
    handleSelect(key) {
      this.$router.push(key);
    },
  },
};
</script>

<style>
#app {
  height: 100%;
}
.logo {
  float: left;
  color: #fff;
  font-size: 18px;
  line-height: 60px;
  margin-right: 20px;
}
.el-menu-demo {
  float: left;
}
</style>
