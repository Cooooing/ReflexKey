import type { NavItem } from "@/types/common";

export const NAV_ITEMS: readonly NavItem[] = [
  { name: "home", icon: "icon-home", text: "主页" },
  { name: "account", icon: "icon-user", text: "账号" },
  { name: "tools", icon: "icon-code", text: "工具" },
] as const;

export const SETTINGS_NAV: NavItem = {
  name: "settings",
  icon: "icon-shezhi",
  text: "设置",
} as const;

// 导出类型
export type NavItemType = (typeof NAV_ITEMS)[number];
export type NavItems = typeof NAV_ITEMS;
