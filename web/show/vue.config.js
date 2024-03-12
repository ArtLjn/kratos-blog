const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave:false,
  assetsDir:'static',
  productionSourceMap:false,
  devServer: {
    port:8033,
    proxy: {
      '/api': {
        target: 'http://localhost:8080/api',
        // ws: true,
        changeOrigin: true, //是否跨域
        pathRewrite: {
          '^/api': ''
        }
      },
      '/util':{
        target: 'http://localhost:8080/util',
        // ws: true,
        changeOrigin: true, //是否跨域
        pathRewrite: {
          '^/util': ''
        }
      }
    },
    
    historyApiFallback: true //增加这个选项
  },
})
