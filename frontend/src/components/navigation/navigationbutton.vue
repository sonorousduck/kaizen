<script setup lang="ts">
import type { NavigationItem } from '@/router/navigation';
import { useRouter } from 'vue-router';
import DuckButton from '@/components/core/DuckButton/DuckButton.vue';

interface Props {
  navigationItem: NavigationItem
}

const props = defineProps<Props>();
const router = useRouter();

function onNavigationClicked() {
  router.push(props.navigationItem.to);
}
</script>

<template>
  <DuckButton
    variant="secondary"
    @click="onNavigationClicked">
    {{ props.navigationItem.label }}
  </DuckButton>
</template>

<style lang="less" scoped>
</style>
