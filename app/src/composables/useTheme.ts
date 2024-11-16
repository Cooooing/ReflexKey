import { ref, computed } from "vue";
import type { Theme, ThemeManager, ThemeConfig } from "@/types/theme";
import { defaultTheme } from "@/themes/default-theme";
import { customTheme } from "@/themes/custom-theme";

export const useTheme = (): ThemeManager => {
  const currentTheme = ref<Theme>("light");

  const themes: Record<Theme, ThemeConfig> = {
    light: defaultTheme,
    dark: customTheme,
  };

  const activeTheme = computed(() => themes[currentTheme.value]);

  const applyTheme = (theme: ThemeConfig) => {
    Object.entries(theme.colors).forEach(([key, value]) => {
      document.documentElement.style.setProperty(key, value);
    });
    document.documentElement.setAttribute("data-theme", theme.name);
  };

  const setThemeColor = (color: string) => {
    const theme = { ...themes[currentTheme.value] };
    theme.colors = { ...theme.colors, "--color-primary": color };
    applyTheme(theme);
  };

  const toggleTheme = () => {
    currentTheme.value = currentTheme.value === "light" ? "dark" : "light";
    applyTheme(themes[currentTheme.value]);
  };

  // 初始化主题
  applyTheme(themes[currentTheme.value]);

  return {
    currentTheme,
    setThemeColor,
    toggleTheme,
  };
};

// 创建单例实例
export const themeManager = useTheme();
