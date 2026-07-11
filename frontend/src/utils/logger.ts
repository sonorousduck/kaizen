import * as Sentry from '@sentry/vue';

export const LogLevel = {
  Debug: 'debug',
  Info: 'info',
  Warning: 'warning',
  Error: 'error',
  Fatal: 'fatal',
} as const;

// eslint-disable-next-line ts/no-redeclare
export type LogLevel = (typeof LogLevel)[keyof typeof LogLevel];

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
      const logFunction = level === LogLevel.Error ? console.error : level === LogLevel.Warning ? console.warn : console.log;

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
