import type { Theme, ThemeConfig } from "@/types/theme";
import { defaultTheme } from "@/themes/default-theme";
import { customTheme } from "@/themes/custom-theme";
import { ref, watch } from "vue";

const THEME_STORAGE_KEY = "app-theme";

export class ThemeManager {
  private themes: Map<Theme, ThemeConfig> = new Map();
  private currentTheme = ref<Theme>("light");

  constructor() {
    this.themes = new Map([
      ["light", defaultTheme],
      ["dark", customTheme],
    ]);
    this.init();
  }

  private init(): void {
    // 从本地存储加载主题
    const savedTheme = localStorage.getItem(THEME_STORAGE_KEY) as Theme;
    if (savedTheme && this.themes.has(savedTheme)) {
      this.currentTheme.value = savedTheme;
    }

    // 监听主题变化
    watch(this.currentTheme, (newTheme) => {
      this.applyTheme(newTheme);
      localStorage.setItem(THEME_STORAGE_KEY, newTheme);
    });
  }

  setTheme(theme: Theme): void {
    if (this.themes.has(theme)) {
      this.currentTheme.value = theme;
    }
  }

  toggleTheme(): void {
    const nextTheme = this.currentTheme.value === "light" ? "dark" : "light";
    this.setTheme(nextTheme);
  }

  private applyTheme(themeName: Theme): void {
    const theme = this.themes.get(themeName);
    if (!theme) return;

    const root = document.documentElement;

    // 应用颜色变量
    Object.entries(theme.colors).forEach(([key, value]) => {
      root.style.setProperty(key, value);
    });

    // 添加调试日志
    console.log("Theme applied:", themeName, theme.colors);
  }
}

export const themeManager = new ThemeManager();
