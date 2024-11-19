/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",  // 注意这里添加了 .vue
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}