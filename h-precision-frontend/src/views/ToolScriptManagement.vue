<template>
  <div class="script-management">
    <h2>{{ strategyName }} - 脚本管理</h2>
    <el-button type="primary" @click="openCreateScriptDialog">
      创建脚本
    </el-button>
    <el-table :data="scripts" style="width: 100%; margin-top: 20px">
      <el-table-column prop="name" label="脚本名称" />
      <el-table-column prop="description" label="描述" />
      <el-table-column label="操作">
        <template #default="scope">
          <el-button size="mini" @click="editScript(scope.row)">
            编辑
          </el-button>
          <el-button
            size="mini"
            type="danger"
            @click="deleteScript(scope.row)"
          >
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 创建脚本对话框 -->
    <el-dialog
      title="创建脚本"
      :visible.sync="createScriptDialogVisible"
    >
      <el-form :model="newScript">
        <el-form-item label="脚本名称">
          <el-input v-model="newScript.name" />
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="newScript.description" />
        </el-form-item>
        <el-form-item label="内容">
          <el-input
            type="textarea"
            v-model="newScript.content"
            rows="5"
          />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="createScriptDialogVisible = false">
          取 消
        </el-button>
        <el-button type="primary" @click="createScript">
          确 定
        </el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'ToolScriptManagement',
  data() {
    return {
      strategyName: this.$route.query.strategyName || '未知策略',
      scripts: [],
      createScriptDialogVisible: false,
      newScript: {
        name: '',
        description: '',
        content: '',
      },
    };
  },
  methods: {
    openCreateScriptDialog() {
      this.createScriptDialogVisible = true;
    },
    createScript() {
      // 添加新脚本到列表
      this.scripts.push({ ...this.newScript });
      this.createScriptDialogVisible = false;
      // 重置表单
      this.newScript = { name: '', description: '', content: '' };
    },
    editScript(script) {
      // TODO: 实现编辑脚本的功能
    },
    deleteScript(script) {
      // 从列表中移除脚本
      this.scripts = this.scripts.filter((s) => s !== script);
    },
  },
};
</script>

<style scoped>
.script-management {
  padding: 20px;
}
</style>
