<template>
    <div class="it">
        <h1 class="aihao">乱起八糟的爱好</h1>
      <div ref="chartRef" style="height: 380px;"></div>
    </div>
  </template>
  
  <script>
  import { ref, onMounted, onBeforeUnmount } from 'vue';
  import * as echarts from 'echarts';
import { useRouter } from 'vue-router';
  
  export default {
    name:"bingTu",
    setup() {
      const router = useRouter();
      const chartRef = ref(null);
      let myChart;
      onMounted(() => {
        myChart = echarts.init(chartRef.value);
        const option = {
          legend: {
            top: 'bottom'
          },
          toolbox: {
            show: true,
            feature: {
              mark: { show: true },
              dataView: { show: true, readOnly: false },
              restore: { show: true },
              saveAsImage: { show: true }
            }
          },
          series: [
            {
              name: 'Nightingale Chart',
              type: 'pie',
              radius: [50, 90],
              center: ['50%', '50%'],
              roseType: 'area',
              itemStyle: {
                borderRadius: 8
              },
              data: [
                { value: 40, name: '美学' },
                { value: 38, name: '摄影' },
                { value: 32, name: '剪辑' },
                { value: 30, name: '鉴赏' },
                { value: 28, name: '游戏' },
                { value: 26, name: '旅游' },
                { value: 22, name: '观影' },
                { value: 18, name: '运动' }
              ]
            }
          ]
        };
  
        option && myChart.setOption(option);
        window.addEventListener('resize', ListenSize);
      });
      onBeforeUnmount(() => {
        window.addEventListener('resize',ListenSize);
        myChart.dispose();
      })
      const ListenSize = () => {
        const screenWidth = window.innerWidth;
      }
      return {
        chartRef
      };
    }
  };
  </script>
<style scoped>
.aihao{
    text-align: center;
}
</style>