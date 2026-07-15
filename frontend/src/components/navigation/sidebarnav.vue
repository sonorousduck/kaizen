<script setup lang="ts">
import DuckButton from '@/components/core/DuckButton/DuckButton.vue';
import { bottomNavigationItems, navigationItems } from '@/router/navigation';
import Navigationbutton from './navigationbutton.vue';
</script>

<template>
  <div
    class="host">
    <DuckButton>
      + New Goal
    </DuckButton>
    <div
      class="main-navigation-button-group">
      <Navigationbutton
        v-for="item in navigationItems"
        :key="item.to"
        :navigation-item="item" />
    </div>

    <div
      class="bottom-section">
      <Navigationbutton
        v-for="item in bottomNavigationItems"
        :key="item.to"
        :navigation-item="item" />
    </div>
  </div>
</template>

<style lang="less" scoped>
@import '@/styles/theme.less';

.host {
  width: 256px;
  background-color: @color-background-soft;
  display: flex;
  flex-direction: column;
  gap: 16px;
  padding: 8px;
}

.main-navigation-button-group {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.bottom-section {
  margin-top: auto;
  display: flex;
  flex-direction: column;
  gap: 2px;
}
</style>
