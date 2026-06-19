<script setup lang="ts">
import { reactive } from 'vue';
import { useRouter } from 'vue-router';
import DuckButton from '@/components/core/DuckButton/DuckButton.vue';

const form = reactive({
  email: '',
  password: '',
});

const router = useRouter();

async function onSubmit() {
  console.log('YEET!');
}

function navigateToCreateAccount() {
  router.push('/createAccount');
}
</script>

<template>
  <div
    class="host">
    <div
      class="login">
      <div
        class="title-subtitle">
        <span
          class="title">Hello there</span>
        <span
          class="subtitle">Welcome back to Kaizen</span>
      </div>

      <el-form
        class="login-form"
        :model="form">
        <el-input
          v-model="form.email"
          class="input"
          type="email"
          placeholder="you@example.com" />

        <el-input
          v-model="form.password"
          class="input"
          type="password"
          placeholder="Password"
          :show-password="true" />

        <div
          class="submit-buttons">
          <DuckButton
            class="login-button"
            variant="primary"
            @click="onSubmit">
            Login
          </DuckButton>

          <div
            class="or-divider">
            OR
          </div>

          <DuckButton
            variant="secondary"
            @click="navigateToCreateAccount">
            Create an account
          </DuckButton>
        </div>
      </el-form>
    </div>
  </div>
</template>

<style lang="less" scoped>
@import '@/styles/theme.less';

.host {
  display: flex;
  flex-grow: 1;
  justify-content: center;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.or-divider {
  display: flex;
  align-items: center;
  text-align: center;
  color: @color-text-secondary;
  font-size: @font-size-sm;

  &::before,
  &::after {
    content: '';
    flex: 1;
    height: 1px;
    background-color: @color-border;
  }

  &::before {
    margin-right: 8px;
  }

  &::after {
    margin-left: 8px;
  }
}

.login {
  display: flex;
  flex-direction: column;
  flex-grow: 1;
  max-width: 640px;
  align-self: center;
  border: 1px solid @color-border;
  border-radius: @border-radius-md;
  padding: 16px;
  gap: 16px;
  margin: 8px;
  background-color: @color-background;
}

.title-subtitle {
  display: flex;
  flex-direction: column;
  gap: 4px;
  align-items: center;
}

.title {
  font-size: @font-size-xl;
  font-weight: @font-weight-bold;
}

.subtitle {
  color: @color-text-secondary;
  font-size: @font-size-sm;
}

.submit-buttons {
  display: flex;
  flex-direction: column;
  gap: 16px;
  padding-top: 32px;
}
</style>
