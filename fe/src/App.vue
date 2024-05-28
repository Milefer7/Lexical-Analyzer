<template>
  <div id="app">
    <HomeView></HomeView>
    <!-- 添加一个链接到/popup的路由 -->
    <RouterLink to="/popup">Go to Popup</RouterLink>
    <RouterView></RouterView>
  
  </div>
</template>

<script>
import HomeView from '@/views/HomeView.vue'
export default {
  name: 'App',
  components: {
    HomeView
  }
}
</script>

<style>
/* 添加全局样式 */
</style>
