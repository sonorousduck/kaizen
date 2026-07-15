<script setup lang="ts">
import { storeToRefs } from 'pinia';
import Bottomnavbar from './components/navigation/bottomnavbar.vue';
import Sidebarnav from './components/navigation/sidebarnav.vue';
import { useResponsive } from './stores/responsive';

const { isMobile } = storeToRefs(useResponsive());
</script>

<template>
  <div
    class="app-layout"
    :class="{ 'app-layout--mobile': isMobile }">
    <Sidebarnav
      v-if="!isMobile" />
    <main
      class="app-layout-content">
      <slot />
    </main>
    <Bottomnavbar
      v-if="isMobile" />
  </div>
</template>

<style lang="less" scoped>
.app-layout {
  display: flex;
  flex: 1;
  min-width: 0;

  &--mobile {
    flex-direction: column;
  }
}

.app-layout-content {
  flex: 1;
  min-width: 0;
  padding: 0 8px;
}
</style>
