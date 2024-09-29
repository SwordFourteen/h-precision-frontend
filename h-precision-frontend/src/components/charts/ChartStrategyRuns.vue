<template>
  <div ref="chart" class="chart"></div>
</template>

<script>
import * as echarts from 'echarts';

export default {
  name: 'ChartStrategyRuns',
  mounted() {
    this.initChart();
  },
  methods: {
    initChart() {
      const chart = echarts.init(this.$refs.chart);
      const option = {
        title: { text: '策略运行次数' },
        tooltip: {},
        xAxis: {
          data: ['策略A', '策略B', '策略C'],
        },
        yAxis: {},
        series: [
          {
            name: '运行次数',
            type: 'bar',
            data: [5, 20, 36],
          },
        ],
      };
      chart.setOption(option);
    },
  },
};
</script>

<style scoped>
.chart {
  width: 100%;
  height: 300px;
}
</style>
