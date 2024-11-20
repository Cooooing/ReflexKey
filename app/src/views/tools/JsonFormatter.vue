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
      <div class="search-bar">
        <div class="search-input-wrapper">
          <i class="iconfont icon-search"></i>
          <input
            type="text"
            v-model="searchText"
            placeholder="搜索..."
            @keydown.enter="findNext"
            @keydown.esc="clearSearch"
          />
          <div class="search-actions">
            <button
              v-if="searchText"
              class="search-btn"
              @click="clearSearch"
              title="清除搜索"
            >
              <i class="iconfont icon-close-bold"></i>
            </button>
            <button
              class="search-btn"
              @click="findPrevious"
              title="上一个"
              :disabled="!searchText"
            >
              <i class="iconfont icon-arrow-up-bold"></i>
            </button>
            <button
              class="search-btn"
              @click="findNext"
              title="下一个"
              :disabled="!searchText"
            >
              <i class="iconfont icon-arrow-down-bold"></i>
            </button>
          </div>
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
import { computed, onMounted, ref, watch } from "vue";
import { VAceEditor } from "vue3-ace-editor";
import { useStore } from "vuex";
import "ace-builds/src-noconflict/mode-json";
import "ace-builds/src-noconflict/theme-chrome";
import "ace-builds/src-noconflict/theme-dracula";
import * as ace from "ace-builds";
import { Message } from "@/components/Message";

const store = useStore();
const editor = ref();
let isFormatting = false;
let lastContent = "";
let isToggling = false;

// 使用 computed 属性连接 Vuex 状态
const content = computed({
  get: () => store.getters["jsonEditor/getContent"],
  set: (value) => store.dispatch("jsonEditor/updateContent", value),
});

const isExpanded = computed({
  get: () => store.getters["jsonEditor/getIsExpanded"],
  set: (value) => store.dispatch("jsonEditor/updateExpanded", value),
});

const editorTheme = computed(() => {
  const currentTheme = store.getters["theme/currentTheme"];
  return currentTheme === "light" ? "chrome" : "dracula";
});

// 初始化时从 store 加载状态
onMounted(() => {
  lastContent = content.value;
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
      const range = session.getFoldWidgetRange(0);
      if (range) {
        session.addFold("", range);
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
  store.dispatch("jsonEditor/clearEditor");
  lastContent = "";
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

const searchText = ref("");
let searchRange: any = null;

// 搜索相关方法
const findNext = () => {
  if (!searchText.value || !editor.value) return;

  const editorInstance = editor.value;

  // 如果是新的搜索词，重新初始化搜索
  if (!searchRange || editorInstance.getSelectedText() !== searchText.value) {
    editorInstance.clearSelection();

    // 设置搜索选项
    const searchOptions = {
      needle: searchText.value,
      caseSensitive: false,
      wholeWord: false,
      regExp: false,
      preventScroll: false,
    };

    // 清除之前的标记
    const markers = editorInstance.session.getMarkers();
    if (markers) {
      Object.keys(markers).forEach((markerId) => {
        editorInstance.session.removeMarker(Number(markerId));
      });
    }

    // 查找所有匹配项并高亮
    const Range = ace.require("ace/range").Range;
    const doc = editorInstance.session.getDocument();
    const lines = doc.getAllLines();

    lines.forEach((line: string, row: number) => {
      let match;
      const searchRegex = new RegExp(searchText.value, "gi");

      while ((match = searchRegex.exec(line)) !== null) {
        const range = new Range(
          row,
          match.index,
          row,
          match.index + match[0].length,
        );
        editorInstance.session.addMarker(
          range,
          "ace_selected-word",
          "text",
          false,
        );
      }
    });

    // 设置当前搜索
    editorInstance.$search.set(searchOptions);
  }

  // 移动到下一个匹配项
  editorInstance.findNext({
    skipCurrent: true,
    wrap: true,
  });
};

const findPrevious = () => {
  if (!searchText.value || !editor.value) return;

  const editorInstance = editor.value;
  editorInstance.findPrevious({
    skipCurrent: true,
    wrap: true,
  });
};

const clearSearch = () => {
  if (!editor.value) return;

  const editorInstance = editor.value;
  searchText.value = "";
  searchRange = null;

  // 清除所有标记
  const markers = editorInstance.session.getMarkers();
  if (markers) {
    Object.keys(markers).forEach((markerId) => {
      editorInstance.session.removeMarker(Number(markerId));
    });
  }

  editorInstance.clearSelection();
  editorInstance.$search.set({
    needle: "",
  });
};

// 监听搜索文本变化
watch(searchText, (newValue) => {
  if (newValue) {
    findNext();
  } else {
    clearSearch();
  }
});
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

.search-bar {
  display: flex;
  padding: 8px 16px;
  border-bottom: 1px solid var(--color-border);
  background: var(--color-bg);
}

.search-input-wrapper {
  position: relative;
  display: flex;
  align-items: center;
  width: 100%;
  gap: 8px;

  .iconfont {
    position: absolute;
    color: var(--color-text-light);
    font-size: 14px;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 32px;
    height: 32px;

    &.icon-search {
      left: 0;
      pointer-events: none;
    }
  }

  input {
    flex: 1;
    height: 32px;
    padding: 0 8px 0 32px;
    border: 1px solid var(--color-border);
    border-radius: 4px;
    background: var(--color-surface);
    color: var(--color-text);
    font-size: 14px;
    outline: none;
    transition: all 0.2s;

    &:focus {
      border-color: var(--color-primary);
      box-shadow: 0 0 0 2px var(--color-primary-light);
    }

    &::placeholder {
      color: var(--color-text-light);
    }
  }
}

.search-actions {
  display: flex;
  gap: 4px;
  margin-left: auto;

  .search-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 32px;
    height: 32px;
    padding: 0;
    border: 1px solid var(--color-border);
    border-radius: 4px;
    background: var(--color-surface);
    color: var(--color-text);
    cursor: pointer;
    transition: all 0.2s;

    &:hover:not(:disabled) {
      border-color: var(--color-primary);
      color: var(--color-primary);
    }

    &:active:not(:disabled) {
      background: var(--color-primary-light);
    }

    &:disabled {
      opacity: 0.5;
      cursor: not-allowed;
    }

    .iconfont {
      font-size: 14px;
    }
  }
}

// 修改高亮样式
:deep(.ace_selected-word) {
  border: none;
  background: var(--color-primary-light);
  opacity: 1;
}

:deep(.ace_selected) {
  background: var(--color-primary) !important;
  opacity: 1;
}
</style>
