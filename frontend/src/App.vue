<script setup lang="ts">
import { RouterView, useRoute } from 'vue-router';
import { isRouterReady } from './router';

const route = useRoute();
</script>

<template>
  <template
    v-if="isRouterReady">
    <main
      class="app-main">
      <RouterView />
    </main>
  </template>
  <template
    v-else>
    <div>Loading</div>
  </template>
</template>

<style lang="less" scoped>
.app-main {
  display: flex;
  flex: 1;
  min-width: 0;
}
</style>
