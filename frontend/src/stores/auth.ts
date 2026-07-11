import type { ModelsCreateUser, ModelsUserResponse } from '@/generated/api';
import { defineStore } from 'pinia';
import { ref } from 'vue';
import { createUser, getCurrentUser, loginUser, logoutUser } from '@/generated/api';
import { logger, LogLevel } from '@/utils/logger';
import { getResponseResult, isOk, isSuccess } from '@/utils/response';

export const useAuthStore = defineStore('auth', () => {
  const isLoggedIn = ref<boolean>(false);
  const initialized = ref<boolean>(false);
  const user = ref<ModelsUserResponse | undefined>(undefined);
  const error = ref<string | undefined>(undefined);

  async function login(email: string, password: string): Promise<ModelsUserResponse> {
    try {
      error.value = undefined;

      const response = await loginUser({ email, password });

      if (!isSuccess<ModelsUserResponse>(response)) {
        logger.log(LogLevel.Error, 'error logging in', {
          context: {
            ...getResponseResult(response),
          },
        });
        return Promise.reject(new Error('error logging in'));
      }

      user.value = response.data;
      isLoggedIn.value = true;
      return user.value;
    }
    catch (err) {
      logger.log(LogLevel.Error, 'login failed', {
        context: {
          err,
        },
      });
      error.value = 'login failed';
      isLoggedIn.value = false;
      return Promise.reject(err);
    }
  }

  async function createAccount(modelsCreateUser: ModelsCreateUser): Promise<ModelsUserResponse> {
    try {
      error.value = undefined;

      const response = await createUser(modelsCreateUser);

      if (!isSuccess<ModelsUserResponse>(response)) {
        logger.log(LogLevel.Error, 'error creating an account', {
          context: {
            ...getResponseResult(response),
          },
        });
        return Promise.reject(new Error('error creating an account in'));
      }

      user.value = response.data;
      isLoggedIn.value = true;
      return response.data;
    }
    catch (err) {
      logger.log(LogLevel.Error, 'create account failed', {
        context: {
          err,
        },
      });
      error.value = 'creating an account failed';
      isLoggedIn.value = false;
      return Promise.reject(err);
    }
  }

  async function logout() {
    try {
      const response = await logoutUser();

      if (!isOk(response)) {
        logger.log(LogLevel.Warning, 'logging out failed', {
          context: {
            ...getResponseResult(response),
          },
        });
        return Promise.reject(new Error('failed to log out'));
      }

      return true;
    }
    catch (err) {
      logger.log(LogLevel.Warning, 'log out failed for an unexpected reason', {
        context: {
          err,
          userId: user.value?.id,
        },
      });
      return Promise.reject(new Error('failed to log out'));
    }
    finally {
      isLoggedIn.value = false;
      user.value = undefined;
      error.value = undefined;
    }
  }

  async function checkAuth() {
    try {
      const response = await getCurrentUser();

      if (isSuccess<ModelsUserResponse>(response)) {
        user.value = response.data;
        isLoggedIn.value = true;
        error.value = undefined;
        return true;
      }
    }
    catch (err) {
      logger.log(LogLevel.Error, 'Failed to check auth for an unknown reason', {
        context: {
          err,
        },
      });

      isLoggedIn.value = false;
      user.value = undefined;
    }
    finally {
      initialized.value = true;
    }

    return false;
  }

  return { isLoggedIn, initialized, user, error, login, createAccount, logout, checkAuth };
});
