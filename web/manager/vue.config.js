const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave:false,
  assetsDir:'static',
  productionSourceMap:false,
  devServer: {
    port:8036,
    proxy: {
      '/api': {
        target: 'http://localhost:8080/api',
        ws: true,
        changeOrigin: true, //是否跨域
        pathRewrite: {
          '^/api': ''
        }
      }
    },
    historyApiFallback: true //增加这个选项
  },
})
