<script setup lang="ts">
import { RouterView } from 'vue-router';
import Applayout from './applayout.vue';
import { isRouterReady } from './router';
import { useAuthStore } from './stores/auth';

const auth = useAuthStore();
</script>

<template>
  <template
    v-if="isRouterReady">
    <Applayout
      v-if="auth.isLoggedIn">
      <RouterView />
    </Applayout>
    <main
      v-else
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
