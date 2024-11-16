<template>
  <div class="json-formatter">
    <div class="editor-panel">
      <div class="panel-header">
        <span>JSON 编辑器</span>
        <div class="actions">
          <button class="action-btn" @click="clearInput" title="清空">
            <i class="iconfont icon-clear" />清空
          </button>
          <button class="action-btn" @click="formatJson" title="格式化">
            <i class="iconfont icon-code" />格式化
          </button>
          <button class="action-btn" @click="toggleExpand" title="展开/收起">
            <i
              class="iconfont"
              :class="isExpanded ? 'icon-collapse' : 'icon-expand'"
            />
            {{ isExpanded ? "收起" : "展开" }}
          </button>
          <button class="action-btn" @click="copyContent" title="复制">
            <i class="iconfont icon-copy" />复制
          </button>
        </div>
      </div>
      <VAceEditor
        v-model:value="content"
        :lang="'json'"
        :theme="editorTheme"
        :options="{
          showPrintMargin: false,
          fontSize: 14,
          enableBasicAutocompletion: true,
          enableLiveAutocompletion: true,
          showLineNumbers: true,
          tabSize: 2,
          wrap: true,
        }"
        class="json-editor"
        @init="editorInit"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from "vue";
import { VAceEditor } from "vue3-ace-editor";
import { useStore } from "vuex";
import "ace-builds/src-noconflict/mode-json";
import "ace-builds/src-noconflict/theme-chrome";
import "ace-builds/src-noconflict/theme-dracula";
import { Message } from "@/components/Message";

const store = useStore();
const content = ref("");
const editor = ref();
const isExpanded = ref(true);
let isFormatting = false;
let lastContent = "";
let isToggling = false;

const editorTheme = computed(() => {
  const currentTheme = store.getters["theme/currentTheme"];
  return currentTheme === "light" ? "chrome" : "dracula";
});

const editorInit = (editorInstance: unknown) => {
  editor.value = editorInstance;
  // 监听内容变化
  editor.value.on("change", () => {
    if (isToggling) return;

    const newContent = editor.value.getValue();
    // 只在内容从空变为有内容时自动格式化
    if (lastContent === "" && newContent.trim() !== "") {
      tryFormatJson(true);
    }
    lastContent = newContent;
  });
};

// 尝试格式化 JSON
const tryFormatJson = (showMessage = true) => {
  if (!content.value.trim()) return;
  if (isFormatting) return;

  try {
    isFormatting = true;
    const parsed = JSON.parse(content.value);
    const formatted = JSON.stringify(parsed, null, 2);

    if (formatted !== content.value) {
      content.value = formatted;
      isExpanded.value = true;
      if (showMessage) {
        Message.success("JSON 格式正确");
      }
    }
  } catch (error) {
    if (showMessage) {
      Message.error("无效的 JSON 格式");
    }
  } finally {
    isFormatting = false;
  }
};

// 格式化按钮点击事件
const formatJson = () => {
  tryFormatJson(true);
};

// 展开/收起
const toggleExpand = () => {
  if (!content.value.trim()) {
    Message.warning("内容为空");
    return;
  }

  try {
    // 先验证是否为有效的 JSON
    JSON.parse(content.value);

    isToggling = true;

    if (isExpanded.value) {
      // 收起：使用编辑器的折叠功能
      const session = editor.value.getSession();
      const range = session.getFoldWidgetRange(0); // 获取第一行的折叠范围
      if (range) {
        session.addFold("", range); // 添加折叠
      }
    } else {
      // 展开：展开所有折叠
      editor.value.getSession().unfold();
    }

    isExpanded.value = !isExpanded.value;
    Message.success(isExpanded.value ? "JSON 已展开" : "JSON 已压缩");

    setTimeout(() => {
      isToggling = false;
    }, 0);
  } catch (error) {
    Message.error("无效的 JSON 格式");
    isToggling = false;
  }
};

// 清空内容
const clearInput = () => {
  content.value = "";
  lastContent = ""; // 重置 lastContent
  Message.info("内容已清空");
};

// 复制内容
const copyContent = async () => {
  try {
    await navigator.clipboard.writeText(content.value);
    Message.success("已复制到剪贴板");
  } catch (err) {
    Message.error("复制失败");
  }
};
</script>

<style lang="scss" scoped>
.json-formatter {
  padding: 20px;
  height: calc(100vh - 120px);
  background: var(--color-bg);
}

.editor-panel {
  height: 100%;
  display: flex;
  flex-direction: column;
  background: var(--color-surface);
  border-radius: 8px;
  border: 1px solid var(--color-border);
  overflow: hidden;
}

.panel-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  border-bottom: 1px solid var(--color-border);
  background: var(--color-surface);

  span {
    font-size: 14px;
    font-weight: 500;
    color: var(--color-text);
  }
}

.actions {
  display: flex;
  gap: 8px;
}

.action-btn {
  display: flex;
  align-items: center;
  gap: 4px;
  padding: 4px 12px;
  border: none;
  border-radius: 4px;
  background: var(--color-surface-variant);
  color: var(--color-text);
  cursor: pointer;
  font-size: 13px;

  &:hover {
    background: var(--color-surface-hover);
  }

  .iconfont {
    font-size: 14px;
  }
}

.json-editor {
  flex: 1;
  width: 100%;
  height: 100%;
  font-family: monospace;
}
</style>
