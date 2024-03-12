<template>
    <div class="skill">
        <h1 class="expertise">专业技能</h1>
      <div ref="chartRef" style="height: 380px"></div>
    </div>
  </template>
  
  <script>
  import { ref, onMounted } from 'vue';
  import * as echarts from 'echarts';
  
  export default {
    setup() {
      const chartRef = ref(null);
  
      onMounted(() => {
        const myChart = echarts.init(chartRef.value);
        const option = {
          angleAxis: {
            type: 'category',
            data: ['前端三件', '数据库', '摸鱼', 'SpringBoot', 'Java', 'Python', 'Vue']
          },
          radiusAxis: {},
          polar: {},
          series: [
            {
              type: 'bar',
              data: [1, 2, 3, 4, 3, 5, 1],
              coordinateSystem: 'polar',
              name: 'A',
              stack: 'a',
              emphasis: {
                focus: 'series'
              }
            },
            {
              type: 'bar',
              data: [2, 4, 6, 1, 3, 2, 1],
              coordinateSystem: 'polar',
              name: 'B',
              stack: 'a',
              emphasis: {
                focus: 'series'
              }
            },
            {
              type: 'bar',
              data: [1, 2, 3, 4, 1, 2, 5],
              coordinateSystem: 'polar',
              name: 'C',
              stack: 'a',
              emphasis: {
                focus: 'series'
              }
            }
          ],
          legend: {
            show: true,
            data: ['A', 'B', 'C']
          }
        };
  
        option && myChart.setOption(option);
      });
  
      return {
        chartRef
      };
    }
  };
  </script>
<style scoped>
.expertise{
    text-align: center;
}
</style>
  