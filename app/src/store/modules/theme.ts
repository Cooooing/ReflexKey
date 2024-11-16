import type { Theme, ThemeConfig } from "@/types/theme";
import { defaultTheme } from "@/themes/default-theme";
import { customTheme } from "@/themes/custom-theme";
import { ActionContext } from "vuex";

interface ThemeState {
  currentTheme: Theme;
  currentThemeColor: string;
}

// 定义 ActionContext 类型
type ThemeActionContext = ActionContext<ThemeState, any>;

const state: ThemeState = {
  currentTheme: "light",
  currentThemeColor: defaultTheme.colors["--color-primary"],
};

const mutations = {
  SET_THEME(state: ThemeState, theme: Theme) {
    state.currentTheme = theme;
  },
  SET_THEME_COLOR(state: ThemeState, color: string) {
    state.currentThemeColor = color;
  },
};

const actions = {
  toggleTheme({ commit, state, dispatch }: ThemeActionContext) {
    const newTheme = state.currentTheme === "light" ? "dark" : "light";
    commit("SET_THEME", newTheme);
    dispatch("applyTheme");
  },

  setThemeColor({ commit, dispatch }: ThemeActionContext, color: string) {
    commit("SET_THEME_COLOR", color);
    dispatch("applyTheme");
  },

  applyTheme({ state }: ThemeActionContext) {
    const themes: Record<Theme, ThemeConfig> = {
      light: defaultTheme,
      dark: customTheme,
    };

    // 修复类型问题
    const currentTheme = state.currentTheme as Theme;
    const currentThemeConfig = { ...themes[currentTheme] };
    currentThemeConfig.colors = {
      ...currentThemeConfig.colors,
      "--color-primary": state.currentThemeColor,
    };

    // 修复类型问题
    Object.entries(currentThemeConfig.colors).forEach(([key, value]) => {
      document.documentElement.style.setProperty(key, value as string);
    });
    document.documentElement.setAttribute("data-theme", state.currentTheme);
  },
};

const getters = {
  currentTheme: (state: ThemeState): Theme => state.currentTheme,
  currentThemeColor: (state: ThemeState): string => state.currentThemeColor,
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
