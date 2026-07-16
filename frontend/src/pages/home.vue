<script setup lang="ts">
import type { ModelsGoal } from '@/generated/api';
import { computedAsync } from '@vueuse/core';
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import DuckButton from '@/components/core/DuckButton/DuckButton.vue';
import Modal from '@/components/core/modal/modal.vue';
import { getGoalsByUser } from '@/generated/api';
import { useAuthStore } from '@/stores/auth';
import { logger, LogLevel } from '@/utils/logger';
import { getResponseResult, isSuccess } from '@/utils/response';

const auth = useAuthStore();
const router = useRouter();

const goals = computedAsync(async () => {
  const userGoals = await getGoalsByUser({});

  if (!isSuccess<ModelsGoal[]>(userGoals)) {
    logger.log(LogLevel.Error, 'Failed to get user goals', {
      context: {
        ...getResponseResult(userGoals),
      },
    });

    return [];
  }

  return userGoals.data;
});

const dialogVisible = ref<boolean>(false);

async function logout() {
  const success = await auth.logout();

  if (success) {
    router.replace('/login');
  }
}

function openModal() {
  dialogVisible.value = true;
}
</script>

<template>
  <div>
    <div> Welcome to logged in world!</div>

    <div
      v-for="goal in goals"
      :key="goal.id">
      {{ goal.title }}
    </div>

    <DuckButton
      @click="logout">
      Logout
    </DuckButton>

    <DuckButton
      @click="openModal">
      Open Modal
    </DuckButton>

    <Modal
      v-model:dialog-visible="dialogVisible"
      title="Test modal" />
  </div>
</template>

<style lang="less" scoped>
  @import '@/styles/theme.less';
</style>
