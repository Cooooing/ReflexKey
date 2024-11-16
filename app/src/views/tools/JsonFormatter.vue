<template>
  <div class="json-formatter">
    <div v-if="showCopyTip" class="copy-tip">
      <i class="iconfont icon-success" />
      <span>已复制</span>
    </div>

    <div class="editor-container">
      <div class="section-header">
        <h3>JSON 编辑器</h3>
        <div class="actions">
          <button class="action-btn" @click="clearInput">
            <i class="iconfont icon-error" />
            <span>清空</span>
          </button>
          <button class="action-btn" @click="formatJson">
            <i class="iconfont icon-code" />
            <span>格式化</span>
          </button>
          <button class="action-btn" @click="compressJson">
            <i class="iconfont icon-code" />
            <span>压缩</span>
          </button>
          <button class="action-btn" @click="copyOutput">
            <i class="iconfont icon-file" />
            <span>复制</span>
          </button>
        </div>
      </div>
      <div ref="editorRef" class="ace-editor"></div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted, onBeforeUnmount, computed } from "vue";
import { useStore } from "vuex";
import ace from "ace-builds";
import "ace-builds/src-noconflict/mode-json";
import "ace-builds/src-noconflict/theme-chrome";
import "ace-builds/src-noconflict/theme-dracula";
import "ace-builds/src-noconflict/ext-language_tools";

const store = useStore();
const showCopyTip = ref(false);
const editorRef = ref<HTMLElement | null>(null);
let editor: ace.Ace.Editor | null = null;

// 使用计算属性获取主题状态
const isDark = computed(() => store.getters["theme/isDark"]);
const currentTheme = computed(() => store.getters["theme/currentTheme"]);

// 初始化编辑器
const initEditor = () => {
  if (!editorRef.value) return;

  // 创建编辑器实例
  editor = ace.edit(editorRef.value, {
    mode: "ace/mode/json",
    theme: isDark.value ? "ace/theme/dracula" : "ace/theme/chrome",
    fontSize: 14,
    tabSize: 2,
    printMargin: false,
    showFoldWidgets: true,
    foldStyle: "markbegin",
    useWorker: false,
    highlightActiveLine: true,
    displayIndentGuides: true,
  });

  // 设置自动完成
  editor.setOptions({
    enableBasicAutocompletion: true,
    enableLiveAutocompletion: true,
    enableSnippets: true,
  });

  // 添加折叠功能
  editor.getSession().setFoldStyle("markbegin");
  editor.getSession().setUseWrapMode(false);

  // 自定义编辑器颜色
  const customizeEditorTheme = () => {
    const style = document.createElement("style");
    style.textContent = `
      .ace-chrome .ace_gutter,
      .ace-dracula .ace_gutter {
        background: ${currentTheme.value["--color-background"]} !important;
      }
      .ace-chrome .ace_gutter-active-line,
      .ace-dracula .ace_gutter-active-line {
        background: ${currentTheme.value["--color-hover"]} !important;
      }
      /* 添加更多自定义样式 */
    `;
    document.head.appendChild(style);
  };

  customizeEditorTheme();

  // 监听主题变化
  store.watch(
    (state) => state.theme.currentTheme,
    () => {
      customizeEditorTheme();
    }
  );
};

onMounted(() => {
  // 确保主题已初始化
  store.dispatch("theme/initTheme");
  initEditor();
});

onBeforeUnmount(() => {
  if (editor) {
    editor.destroy();
    editor = null;
  }
});

// 格式化 JSON
const formatJson = () => {
  if (!editor) return;
  try {
    const value = editor.getValue();
    if (!value.trim()) return;

    const parsed = JSON.parse(value);
    const formatted = JSON.stringify(parsed, null, 2);
    editor.setValue(formatted, -1);

    // 展开所有折叠
    const session = editor.getSession();
    const rows = session.getLength();
    for (let i = 0; i < rows; i++) {
      session.unfold(i);
    }
  } catch (error) {
    console.error("JSON 格式化失败:", error);
  }
};

// 压缩 JSON
const compressJson = () => {
  if (!editor) return;
  try {
    const value = editor.getValue();
    if (!value.trim()) return;

    const parsed = JSON.parse(value);
    const compressed = JSON.stringify(parsed);
    editor.setValue(compressed, -1);
  } catch (error) {
    console.error("JSON 压缩失败:", error);
  }
};

// 复制内容
const copyOutput = async () => {
  if (!editor) return;
  try {
    await navigator.clipboard.writeText(editor.getValue());
    showCopyTip.value = true;
    setTimeout(() => {
      showCopyTip.value = false;
    }, 2000);
  } catch (error) {
    console.error("复制失败:", error);
  }
};

// 清空输入
const clearInput = () => {
  if (!editor) return;
  editor.setValue("", -1);
};
</script>

<style scoped>
.json-formatter {
  height: 100%;
  padding: 16px;
}

.editor-container {
  height: calc(100vh - 180px);
  background: var(--color-surface);
  border-radius: 8px;
  border: 1px solid var(--color-border);
  display: flex;
  flex-direction: column;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  border-bottom: 1px solid var(--color-border);
}

.section-header h3 {
  margin: 0;
  font-size: 14px;
  font-weight: 500;
  color: var(--color-text);
}

.actions {
  display: flex;
  gap: 8px;
}

.action-btn {
  display: flex;
  align-items: center;
  gap: 4px;
  padding: 4px 8px;
  border: 1px solid var(--color-border);
  border-radius: 4px;
  background: var(--color-surface);
  color: var(--color-text);
  cursor: pointer;
  transition: all 0.2s;
  min-width: 64px;
  justify-content: center;
}

.action-btn:hover {
  background: var(--color-hover);
  border-color: var(--color-primary);
  color: var(--color-primary);
}

.action-btn .iconfont {
  font-size: 14px;
}

.action-btn span {
  font-size: 12px;
}

.ace-editor {
  flex: 1;
  width: 100%;
  font-family: Consolas, Monaco, "Courier New", monospace;
}

/* 复制提示 */
.copy-tip {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 24px;
  background: var(--color-success);
  color: white;
  border-radius: 4px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
  z-index: 1000;
  animation: fadeInOut 2s ease-in-out;
}

@keyframes fadeInOut {
  0% {
    opacity: 0;
    transform: translate(-50%, -40%);
  }
  15% {
    opacity: 1;
    transform: translate(-50%, -50%);
  }
  85% {
    opacity: 1;
    transform: translate(-50%, -50%);
  }
  100% {
    opacity: 0;
    transform: translate(-50%, -60%);
  }
}

:deep(.ace_editor) {
  background: var(--color-surface) !important;
  color: var(--color-text) !important;
}

:deep(.ace_gutter) {
  background: var(--color-background) !important;
  color: var(--color-text-secondary) !important;
}

:deep(.ace_gutter-cell) {
  color: var(--color-text-secondary) !important;
}

:deep(.ace_gutter-active-line) {
  background-color: var(--color-hover) !important;
}

:deep(.ace_marker-layer .ace_active-line) {
  background: var(--color-hover) !important;
  opacity: 0.1;
}

:deep(.ace_cursor) {
  color: var(--color-primary) !important;
}

:deep(.ace_print-margin) {
  display: none !important;
}

/* 语法高亮颜色 - 亮色主题 */
:deep([data-theme="light"] .ace_editor) {
  .ace_string {
    color: #42b983 !important;
  }
  .ace_constant.ace_numeric {
    color: #f08d49 !important;
  }
  .ace_constant.ace_boolean {
    color: #7c4dff !important;
  }
  .ace_paren {
    color: #999 !important;
  }
}

/* 语法高亮颜色 - 暗色主题 */
:deep([data-theme="dark"] .ace_editor) {
  .ace_string {
    color: #50fa7b !important;
  }
  .ace_constant.ace_numeric {
    color: #ff79c6 !important;
  }
  .ace_constant.ace_boolean {
    color: #bd93f9 !important;
  }
  .ace_paren {
    color: #f8f8f2 !important;
  }
}

/* 选中文本的背景色 */
:deep(.ace_marker-layer .ace_selection) {
  background: var(--color-primary) !important;
  opacity: 0.2;
}

/* 匹配括号的样式 */
:deep(.ace_marker-layer .ace_bracket) {
  border: 1px solid var(--color-primary) !important;
  opacity: 0.5;
}

/* 缩进指引线 */
:deep(.ace_indent-guide) {
  background: none !important;
  border-right: 1px dashed var(--color-border) !important;
  opacity: 0.5;
}

/* 折叠按钮样式 */
:deep(.ace_fold-widget) {
  background: transparent !important;
  border: none !important;
  box-shadow: none !important;
  color: var(--color-text-secondary) !important;
}

:deep(.ace_fold-widget:hover) {
  color: var(--color-primary) !important;
}
</style>
