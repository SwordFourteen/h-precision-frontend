<template>
  <div ref="chart" class="chart"></div>
</template>

<script>
import * as echarts from 'echarts';

export default {
  name: 'ChartVulnerabilities',
  mounted() {
    this.initChart();
  },
  methods: {
    initChart() {
      const chart = echarts.init(this.$refs.chart);
      const option = {
        title: { text: '漏洞检测能力' },
        tooltip: {},
        legend: { data: ['漏洞数量'] },
        xAxis: { data: ['漏洞1', '漏洞2', '漏洞3'] },
        yAxis: {},
        series: [
          {
            name: '漏洞数量',
            type: 'line',
            data: [10, 15, 9],
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
