<template>
  <div class="model-tuning">
    <h2>模型调优</h2>
    <el-card class="model-card">
      <h3>{{ model.name }}</h3>
      <p>对应工具：{{ model.tool }}</p>
      <p>版本：{{ model.version }}</p>
      <el-button type="primary" @click="tuneModel">
        调优模型
      </el-button>
      <el-upload
        action="#"
        :before-upload="beforeUpload"
        :on-success="handleUploadSuccess"
      >
        <el-button type="primary">上传新模型</el-button>
      </el-upload>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'ModelTuning',
  data() {
    return {
      model: {
        name: '深度学习模型A',
        tool: '工具A',
        version: '1.0.0',
      },
    };
  },
  methods: {
    tuneModel() {
      // TODO: 调用后端接口进行模型调优
      this.$message.success('模型调优已开始');
    },
    beforeUpload(file) {
      // 限制上传文件类型和大小
      const isModelFile = file.type === 'application/octet-stream';
      const isLt10M = file.size / 1024 / 1024 < 10;
      if (!isModelFile) {
        this.$message.error('只能上传模型文件');
      }
      if (!isLt10M) {
        this.$message.error('文件大小不能超过 10MB');
      }
      return isModelFile && isLt10M;
    },
    handleUploadSuccess(response, file) {
      // TODO: 处理上传成功后的逻辑
      this.$message.success('模型上传成功');
    },
  },
};
</script>

<style scoped>
.model-tuning {
  padding: 20px;
}
.model-card {
  max-width: 400px;
}
</style>
