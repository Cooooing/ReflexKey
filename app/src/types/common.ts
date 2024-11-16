// 通用类型定义
export interface MenuItem {
  text: string;
  path?: string;
}

export interface BreadcrumbItem extends MenuItem {
  path?: string;
}

// 添加导航项类型定义
export interface NavItem {
  name: string;
  icon: string;
  text: string;
}
