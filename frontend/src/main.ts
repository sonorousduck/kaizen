import * as Sentry from '@sentry/vue';
import { createPinia } from 'pinia';
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate';
import { createApp } from 'vue';
import App from './App.vue';
import router from './index.ts';
import '@/styles/theme.less';

const app = createApp(App);

Sentry.init({
  dsn: 'https://28916b703ef147d4863e1184d765b264@logging.sonorousduck.com/3',
  integrations: [
    Sentry.browserTracingIntegration({ router }),
    Sentry.vueIntegration({ app }),
  ],
  environment: import.meta.env.MODE,
});

const pinia = createPinia();
pinia.use(piniaPluginPersistedstate);

app.use(pinia);
app.use(router);

app.mount('#app');
