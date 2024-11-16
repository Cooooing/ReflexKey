import type { Theme } from "@/types/theme";

// 主题相关工具函数
export const getThemeColorAlpha = (color: string, theme: Theme): string => {
  return theme === "dark"
    ? `color-mix(in srgb, ${color} 30%, transparent)`
    : `color-mix(in srgb, ${color} 15%, transparent)`;
};
