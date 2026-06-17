import * as Sentry from '@sentry/vue';

type LogLevel = 'debug' | 'info' | 'warning' | 'error' | 'fatal';

interface LogOptions {
  context: Record<string, unknown>

  fingerprint?: string[]
}

class Logger {
  private isDev = import.meta.env.DEV;

  log(level: LogLevel, message: string, options?: LogOptions) {
    const { context, fingerprint } = options || {};

    if (this.isDev) {
      // eslint-disable-next-line no-console
      const logFunction = level === 'error' ? console.error : level === 'warning' ? console.warn : console.log;

      logFunction(`[${level.toUpperCase()}] ${message}`, { ...context });
    }

    Sentry.withScope((scope) => {
      scope.setLevel(level);
      if (context)
        scope.setExtras(context);
      if (fingerprint)
        scope.setFingerprint(fingerprint);
      Sentry.captureMessage(message);
    });
  }
}

export const logger = new Logger();
