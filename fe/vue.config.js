const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    host: '0.0.0.0',
    port: 3000,
    proxy: {
      '/': {
        target: 'http://localhost:8080',
        changeOrigin: true
      }
    }
  }
})