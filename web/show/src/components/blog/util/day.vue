<template>
  <div id="main" style="width: 100%; height: 400px;"></div>
</template>

<script>
import { onMounted, ref } from 'vue';
import * as echarts from 'echarts';

export default {
  name: 'day',
  setup() {
    const chartDom = ref(null);
    let myChart;

    onMounted(() => {
      myChart = echarts.init(chartDom.value);

      const getVirtualData = (year) => {
        const date = +echarts.time.parse(year + '-01-01');
        const end = +echarts.time.parse(+year + 1 + '-01-01');
        const dayTime = 3600 * 24 * 1000;
        const data = [];
        for (let time = date; time < end; time += dayTime) {
          data.push([
            echarts.time.format(time, '{yyyy}-{MM}-{dd}', false),
            Math.floor(Math.random() * 10000)
          ]);
        }
        return data;
      };

      const option = {
        title: {
          top: 30,
          left: 'center',
          text: 'Daily Step Count'
        },
        tooltip: {},
        visualMap: {
          min: 0,
          max: 10000,
          type: 'piecewise',
          orient: 'horizontal',
          left: 'center',
          top: 65
        },
        calendar: {
          top: 120,
          left: 30,
          right: 30,
          cellSize: ['auto', 13],
          range: '2016',
          itemStyle: {
            borderWidth: 0.5
          },
          yearLabel: { show: false }
        },
        series: {
          type: 'heatmap',
          coordinateSystem: 'calendar',
          data: getVirtualData('2016')
        }
      };

      option && myChart.setOption(option);
    });

    return {
      chartDom
    };
  }
};
</script>