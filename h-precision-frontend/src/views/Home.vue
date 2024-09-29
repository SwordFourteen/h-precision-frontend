<template>
  <div class="home-container">
    <el-row :gutter="20">
      <el-col :span="8">
        <chart-strategy-runs />
      </el-col>
      <el-col :span="8">
        <chart-vulnerabilities />
      </el-col>
      <el-col :span="8">
        <chart-detection-stats />
      </el-col>
    </el-row>
  </div>
</template>

<script>
import ChartStrategyRuns from '../components/charts/ChartStrategyRuns.vue';
import ChartVulnerabilities from '../components/charts/ChartVulnerabilities.vue';
import ChartDetectionStats from '../components/charts/ChartDetectionStats.vue';

export default {
  name: 'Home',
  components: {
    ChartStrategyRuns,
    ChartVulnerabilities,
    ChartDetectionStats,
  },
};
</script>

<style scoped>
.home-container {
  padding: 20px;
}
</style>
