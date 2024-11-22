import globals from 'globals'
import pluginJs from '@eslint/js'
import tseslint from 'typescript-eslint'
import pluginVue from 'eslint-plugin-vue'
import pluginPrettierRecommendedConfigs from 'eslint-plugin-prettier/recommended'

/** @type {import('eslint').Linter.Config[]} */
export default [
  { files: ['**/*.{js,mjs,cjs,ts,vue}'] },
  { languageOptions: { globals: globals.browser } },
  pluginJs.configs.recommended,
  ...tseslint.configs.recommended,
  ...pluginVue.configs['flat/essential'],
  // prettier 默认推荐规则
  pluginPrettierRecommendedConfigs,
  {
    files: ['**/*.vue'],
    languageOptions: {
      ecmaVersion: 'latest',
      parserOptions: {
        parser: tseslint.parser
      }
    }
  }
]
