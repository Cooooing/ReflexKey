export interface EditorTheme {
  background: string;
  foreground: string;
  gutter: {
    background: string;
    foreground: string;
    activeBackground: string;
  };
  selection: {
    background: string;
  };
  cursor: {
    color: string;
  };
  // 添加更多编辑器主题相关的类型定义
}
