<template>
  <div ref="chart" class="chart"></div>
</template>

<script>
import * as echarts from 'echarts';

export default {
  name: 'ChartDetectionStats',
  mounted() {
    this.initChart();
  },
  methods: {
    initChart() {
      const chart = echarts.init(this.$refs.chart);
      const option = {
        title: { text: '检测数量统计' },
        tooltip: {},
        legend: { data: ['检测数量', '存在漏洞'] },
        xAxis: { data: ['一月', '二月', '三月'] },
        yAxis: {},
        series: [
          {
            name: '检测数量',
            type: 'bar',
            data: [50, 200, 360],
          },
          {
            name: '存在漏洞',
            type: 'bar',
            data: [20, 80, 150],
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
