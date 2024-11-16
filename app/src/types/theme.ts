import { Ref } from "vue";

export type Theme = "light" | "dark";

export interface ThemeColors {
  "--color-primary": string;
  "--color-background": string;
  "--color-surface": string;
  "--color-text": string;
  "--color-text-secondary": string;
  "--color-border": string;
  "--color-hover": string;
  "--color-success": string;
  [key: string]: string;
}

export interface ThemeConfig {
  name: Theme;
  colors: ThemeColors;
}

export interface ThemeManager {
  currentTheme: Ref<Theme>;
  setThemeColor: (color: string) => void;
  toggleTheme: () => void;
}
