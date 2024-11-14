module.exports = {
  root: true,
  env: {
    node: true,
  },
  extends: [
    "plugin:vue/vue3-essential",
    "eslint:recommended",
    "@vue/typescript/recommended",
    "plugin:prettier/recommended",
    "prettier", // 确保 ESLint 不会破坏 Prettier 的格式化
  ],
  parserOptions: {
    ecmaVersion: 2020,
    parser: "@typescript-eslint/parser", // 使用 TypeScript 解析器
  },
  settings: {
    "import/resolver": {
      node: {
        extensions: [".js", ".ts", ".vue"],
      },
    },
  },
  rules: {
    "no-console": process.env.NODE_ENV === "production" ? "warn" : "off",
    "no-debugger": process.env.NODE_ENV === "production" ? "warn" : "off",
    "prefer-const": "warn", // 建议使用 const 而不是 let
    "no-unused-vars": "warn", // 禁止未使用的变量
    "@typescript-eslint/no-explicit-any": "warn", // 禁止使用 any 类型
    "@typescript-eslint/explicit-module-boundary-types": "warn", // 建议在模块边界上明确指定类型
    "prettier/prettier": "error", // 确保代码符合 Prettier 规则
  },
};
