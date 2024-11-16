import type { Theme, ThemeConfig } from "@/types/theme";
import { defaultTheme } from "@/themes/default-theme";
import { customTheme } from "@/themes/custom-theme";
import type { ActionContext } from "vuex";

interface ThemeState {
  currentTheme: Theme;
  currentThemeColor: string;
}

// 定义 ActionContext 类型，使用 unknown 替代 any
type ThemeActionContext = ActionContext<ThemeState, unknown>;

const state: ThemeState = {
  currentTheme: "light",
  currentThemeColor: defaultTheme.colors["--color-primary"],
};

const mutations = {
  SET_THEME(state: ThemeState, theme: Theme): void {
    state.currentTheme = theme;
  },
  SET_THEME_COLOR(state: ThemeState, color: string): void {
    state.currentThemeColor = color;
  },
};

const actions = {
  toggleTheme({ commit, state, dispatch }: ThemeActionContext): void {
    const newTheme = state.currentTheme === "light" ? "dark" : "light";
    commit("SET_THEME", newTheme);
    dispatch("applyTheme");
  },

  setThemeColor({ commit, dispatch }: ThemeActionContext, color: string): void {
    commit("SET_THEME_COLOR", color);
    dispatch("applyTheme");
  },

  applyTheme({ state }: ThemeActionContext): void {
    const themes: Record<Theme, ThemeConfig> = {
      light: defaultTheme,
      dark: customTheme,
    };

    const currentTheme = state.currentTheme as Theme;
    const currentThemeConfig = { ...themes[currentTheme] };
    currentThemeConfig.colors = {
      ...currentThemeConfig.colors,
      "--color-primary": state.currentThemeColor,
    };

    Object.entries(currentThemeConfig.colors).forEach(([key, value]) => {
      document.documentElement.style.setProperty(key, value as string);
    });
    document.documentElement.setAttribute("data-theme", state.currentTheme);
    saveTheme(state);
  },
};

const getters = {
  currentTheme: (state: ThemeState): Theme => state.currentTheme,
  currentThemeColor: (state: ThemeState): string => state.currentThemeColor,
};

// 添加持久化功能
const THEME_STORAGE_KEY = "app-theme";

const saveTheme = (theme: ThemeState): void => {
  localStorage.setItem(THEME_STORAGE_KEY, JSON.stringify(theme));
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
