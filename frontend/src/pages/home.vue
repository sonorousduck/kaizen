<script setup lang="ts">
import { useRouter } from 'vue-router';
import DuckButton from '@/components/core/DuckButton/DuckButton.vue';
import { useAuthStore } from '@/stores/auth';

const auth = useAuthStore();
const router = useRouter();

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
