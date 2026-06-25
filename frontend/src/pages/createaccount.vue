<script setup lang="ts">
import type { FormInstance, FormRules } from 'element-plus';
import { reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import DuckButton from '@/components/core/DuckButton/DuckButton.vue';
import { useAuthStore } from '@/stores/auth';

const form = reactive({
  email: '',
  firstName: '',
  lastName: '',
  password: '',
  confirmPassword: '',
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

    const response = await auth.createAccount({ email: form.email, firstName: form.firstName, lastName: form.lastName, password: form.password });

    if (response) {
      router.replace('/home');
    }
  }
  catch {
    // Errors are thrown if there are validation errors.
    // I don't want to log those because I don't care
  }
}

function navigateToLogin() {
  router.replace('/login');
}

function validatePassword(_rule: any, value: any, callback: (error?: string | Error) => void) {
  if (value === '') {
    callback(new Error('Please input the password'));
  }
  else if (value.length < 7) {
    callback(new Error('Password must be at least 7 characters'));
  }
  else {
    callback();
  }
}

function validateConfirmPassword(_rule: any, value: any, callback: (error?: string | Error) => void) {
  if (value === '') {
    callback(new Error('Please input the password again'));
  }
  else if (value !== form.password) {
    callback(new Error('The passwords don\'t match!'));
  }
  else {
    callback();
  }
}

const rules = reactive<FormRules<typeof form>>({
  email: [{ required: true, message: 'Please input an email', trigger: 'blur' }, { type: 'email', message: 'Please input a valid email', trigger: ['blur', 'change'] }],
  firstName: [{ required: true, message: 'Please input your first name', trigger: 'blur' }],
  lastName: [{ required: true, message: 'Please input your last name', trigger: 'blur' }],
  password: [{ required: true, validator: validatePassword, trigger: 'blur' }],
  confirmPassword: [{ required: true, validator: validateConfirmPassword, trigger: 'blur' }],
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
        :model="form"
        @submit.prevent="onSubmit">
        <el-form-item
          prop="email">
          <el-input
            v-model="form.email"
            size="large"
            type="email"
            placeholder="you@example.com" />
        </el-form-item>
        <div
          class="one-row">
          <el-form-item
            prop="firstName"
            class="row-input">
            <el-input
              v-model="form.firstName"
              size="large"
              type="text"
              placeholder="First name" />
          </el-form-item>
          <el-form-item
            prop="lastName"
            class="row-input">
            <el-input
              v-model="form.lastName"
              size="large"
              type="text"
              placeholder="Last name" />
          </el-form-item>
        </div>

        <el-form-item
          prop="password"
          class="row-input">
          <el-input
            v-model="form.password"
            size="large"
            type="password"
            placeholder="Password"
            :show-password="true" />
        </el-form-item>

        <el-form-item
          prop="confirmPassword"
          class="row-input">
          <el-input
            v-model="form.confirmPassword"
            size="large"
            type="password"
            placeholder="Confirm password"
            :show-password="true" />
        </el-form-item>

        <div
          class="submit-buttons">
          <DuckButton
            class="login-button"
            variant="primary"
            html-type="submit"
            @click="onSubmit">
            Create account
          </DuckButton>

          <div
            class="or-divider">
            OR
          </div>

          <DuckButton
            variant="secondary"
            @click="navigateToLogin">
            Login
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

.one-row {
  display: flex;
  flex-direction: row;
  gap: 8px;
}

.row-input {
  display: flex;
  flex-grow: 1;
}
</style>
