<script setup lang="ts">
import { RouterLink } from 'vue-router';
import { navigationItems } from '@/router/navigation';
</script>

<template>
  <nav
    class="bottom-navigation">
    <RouterLink
      v-for="item in navigationItems"
      :key="item.to"
      :to="item.to"
      class="bottom-navigation-item"
      active-class="bottom-navigation-item-active">
      {{ item.label }}
    </RouterLink>
  </nav>
</template>

<style lang="less" scoped>
@import '@/styles/theme.less';

.bottom-navigation {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  display: flex;
  justify-content: space-around;
  height: 56px;
  background-color: @color-background-soft;
}

.bottom-navigation-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  flex: 1;
  height: 100%;
  color: @color-text-secondary;
}

.bottom-navigation-item-active {
  color: @color-text-primary;
  font-weight: @font-weight-bold;
}
</style>
