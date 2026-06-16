import antfu from '@antfu/eslint-config'
import noRelativeImportPaths from 'eslint-plugin-no-relative-import-paths';


export default antfu({
formatters: {
    css: true,
    html: true,
    markdown: 'prettier',
    vue: true,
  },
  vue: true,
}, {
  plugins: {
    'no-relative-import-paths': noRelativeImportPaths,
  },
  rules: {
    'style/semi': ['error', 'always'],
    'style/quotes': ['error', 'single'],
    'style/comma-dangle': ['error', 'always-multiline'],
    'vue/html-closing-bracket-newline': [
      'error',
      {
        singleline: 'never',
        multiline: 'never',
      },
    ],
    'vue/first-attribute-linebreak': [
      'error',
      {
        singleline: 'below',
        multiline: 'below',
      },
    ],
    'vue/max-attributes-per-line': [
      'error',
      {
        singleline: 3,
        multiline: 1,
      },
    ],
    'vue/html-indent': [
      'error',
      2,
      {
        attribute: 1,
        baseIndent: 1,
        closeBracket: 0,
        alignAttributesVertically: false,
      },
    ],
    'no-relative-import-paths/no-relative-import-paths': [
      'warn',
      { allowSameFolder: true, rootDir: 'src', prefix: '@' },
    ],
  },
})
