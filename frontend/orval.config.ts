export default {
  api: {
    input: {
      target: '../backend/docs/swagger.json',
    },
    output: {
      target: './src/generated/api.ts',
      client: 'fetch',
      httpClient: 'fetch',
      // eslint-disable-next-line no-template-curly-in-string
      baseUrl: '${import.meta.env.VITE_API_BASE_URL}',
    },
  },
};
