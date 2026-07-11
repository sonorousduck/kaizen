<script setup lang="ts">
import type { ModelsGoal } from '@/generated/api';
import { computedAsync } from '@vueuse/core';
import { useRouter } from 'vue-router';
import DuckButton from '@/components/core/DuckButton/DuckButton.vue';
import { getGoalsByUser } from '@/generated/api';
import { useAuthStore } from '@/stores/auth';
import { logger } from '@/utils/logger';
import { isSuccess } from '@/utils/response';

const auth = useAuthStore();
const router = useRouter();

const goals = computedAsync(async () => {
  const userGoals = await getGoalsByUser({});

  if (!isSuccess<ReadonlyArray<ModelsGoal>>) {
    logger.log('info');
  }
});

async function logout() {
  const success = await auth.logout();

  if (success) {
    router.replace('/login');
  }
}
</script>

<template>
  <div>
    <div> Welcome to logged in world!</div>

    <DuckButton
      @click="logout">
      Logout
    </DuckButton>
  </div>
</template>

<style lang="less" scoped>
</style>
