<script setup lang="ts">
import type { ElAlert, FormInstance, FormRules } from 'element-plus';
import { reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import DuckButton from '@/components/core/DuckButton/DuckButton.vue';
import { useAuthStore } from '@/stores/auth';

const form = reactive({
  email: '',
  password: '',
});

const router = useRouter();
const formRef = ref<FormInstance>();
const auth = useAuthStore();

async function onSubmit() {
  if (!formRef.value)
    return;

  try {
    const isValid = await formRef.value.validate();
    if (!isValid)
      return;

    const response = await auth.login(form.email, form.password);

    if (response) {
      router.replace('/home');
    }
  }
  catch {
    // Errors are thrown if there are validation errors.
    // I don't want to log those because I don't care
  }
}

function navigateToCreateAccount() {
  router.push('/createAccount');
}

const rules = reactive<FormRules<typeof form>>({
  email: [{ required: true, message: 'Please input an email', trigger: 'blur' }, { type: 'email', message: 'Please input a valid email', trigger: ['blur', 'change'] }],
  password: [{ required: true, message: 'Password is required', trigger: 'blur' }],
});
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
        ref="formRef"
        class="login-form"
        :rules="rules"
        status-icon
        :model="form">
        <el-form-item
          prop="email">
          <el-input
            v-model="form.email"
            size="large"
            type="email"
            placeholder="you@example.com" />
        </el-form-item>

        <el-form-item
          prop="password">
          <el-input
            v-model="form.password"
            size="large"
            type="password"
            placeholder="Password"
            :show-password="true" />
        </el-form-item>

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
    <ElAlert
      v-if="auth.error"
      :title="auth.error"
      type="error"
      show-icon
      :closable="false" />
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
  gap: 16px;
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
