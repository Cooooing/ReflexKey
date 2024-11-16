export interface Tool {
  id: number;
  name: string;
  description: string;
  icon: string;
}

export interface ThemeColor {
  name: string;
  value: string;
}

export interface OverviewItem {
  title: string;
  value: string | number;
  icon: string;
}

export interface JsonLine {
  content: string;
  level: number;
  isClosing: boolean;
  isCollapsible: boolean;
  collapsedUntil?: number;
}

export interface CommonResponse<T> {
  code: number;
  data: T;
  message: string;
  success: boolean;
}

export interface UserInfo {
  id: string;
  username: string;
  email: string;
  avatar?: string;
  nickname?: string;
  role: string;
  status: number;
  createdAt: string;
  updatedAt: string;
}

export type JsonValue =
  | string
  | number
  | boolean
  | null
  | JsonValue[]
  | { [key: string]: JsonValue };
