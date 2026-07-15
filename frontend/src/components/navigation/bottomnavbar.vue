<script setup lang="ts">
import { navigationItems } from '@/router/navigation';
import Bottomnavigationbutton from './bottomnavigationbutton.vue';
</script>

<template>
  <nav
    class="bottom-navigation">
    <Bottomnavigationbutton
      v-for="item in navigationItems"
      :key="item.to"
      :to="item.to"
      :navigation-item="item">
      {{ item.label }}
    </Bottomnavigationbutton>
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
</style>
