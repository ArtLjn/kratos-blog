<template>
  <div>
    <el-select v-model="selectedItem" placeholder="请选择">
      <el-option
          v-for="item in selectOptions"
          :key="item.value"
          :label="item.label"
          :value="item.value"
      ></el-option>
    </el-select>

    <el-button @click="updateOptions">更新选项</el-button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      selectedItem: '',
      selectOptions: [
        { value: '1', label: '选项1' },
        { value: '2', label: '选项2' },
        { value: '3', label: '选项3' }
      ]
    };
  },
  methods: {
    updateOptions() {
      // 模拟异步获取数据
      // setTimeout(() => {
      //
      // }, 1000);
      this.selectOptions = [
        { value: '4', label: '选项4' },
        { value: '5', label: '选项5' },
        { value: '6', label: '选项6' }
      ];
    }
  }
};
</script>