<template>
  <div class="strategy-management">
    <el-button type="primary" @click="openCreateStrategyDialog">
      创建策略
    </el-button>
    <el-table :data="strategies" style="width: 100%; margin-top: 20px">
      <el-table-column prop="name" label="策略名称" />
      <el-table-column prop="tool" label="检测工具" />
      <el-table-column label="操作">
        <template #default="scope">
          <el-button
            size="mini"
            @click="goToScriptManagement(scope.row)"
          >
            管理脚本
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 创建策略对话框 -->
    <el-dialog
      title="创建策略"
      :visible.sync="createStrategyDialogVisible"
    >
      <el-form :model="newStrategy">
        <el-form-item label="策略名称">
          <el-input v-model="newStrategy.name" />
        </el-form-item>
        <el-form-item label="选择工具">
          <el-select v-model="newStrategy.tool" placeholder="请选择">
            <el-option
              v-for="tool in tools"
              :key="tool.id"
              :label="tool.name"
              :value="tool.name"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="createStrategyDialogVisible = false">
          取 消
        </el-button>
        <el-button type="primary" @click="createStrategy">
          确 定
        </el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'ToolStrategyManagement',
  data() {
    return {
      tools: [
        { id: 1, name: '工具A' },
        { id: 2, name: '工具B' },
      ],
      strategies: [],
      createStrategyDialogVisible: false,
      newStrategy: {
        name: '',
        tool: '',
      },
    };
  },
  methods: {
    openCreateStrategyDialog() {
      this.createStrategyDialogVisible = true;
    },
    createStrategy() {
      // 添加新策略到列表
      this.strategies.push({ ...this.newStrategy });
      this.createStrategyDialogVisible = false;
      // 重置表单
      this.newStrategy = { name: '', tool: '' };
    },
    goToScriptManagement(strategy) {
      // 跳转到脚本管理页面，传递策略信息
      this.$router.push({
        path: '/tool-script',
        query: { strategyName: strategy.name },
      });
    },
  },
};
</script>

<style scoped>
.strategy-management {
  padding: 20px;
}
</style>
