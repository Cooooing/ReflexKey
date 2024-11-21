<template>
  <div class="aes-formatter">
    <div class="editor-panel">
      <div class="panel-header">
        <span>AES 加密/解密</span>
      </div>

      <!-- 配置区域 -->
      <div class="config-section">
        <div class="config-row">
          <div class="config-item">
            <label>运算模式：</label>
            <select v-model="formState.operationMode">
              <option
                v-for="mode in operationModes"
                :key="mode.value"
                :value="mode.value"
              >
                {{ mode.label }}
              </option>
            </select>
          </div>
          <div class="config-item">
            <label>填充方式：</label>
            <select v-model="formState.fill">
              <option
                v-for="fill in fillModes"
                :key="fill.value"
                :value="fill.value"
              >
                {{ fill.label }}
              </option>
            </select>
          </div>
          <div class="config-item">
            <label>密钥长度：</label>
            <select v-model="formState.keyLength">
              <option
                v-for="len in keyLengths"
                :key="len.value"
                :value="len.value"
              >
                {{ len.label }}
              </option>
            </select>
          </div>
        </div>

        <div class="config-row">
          <div class="config-item flex-grow">
            <label>密钥：</label>
            <div class="input-group">
              <input
                v-model="formState.key"
                type="text"
                :placeholder="`请输入${formState.keyLength}位密钥`"
              />
              <select v-model="formState.keyFormat">
                <option
                  v-for="format in keyFormats"
                  :key="format.value"
                  :value="format.value"
                >
                  {{ format.label }}
                </option>
              </select>
            </div>
          </div>
        </div>

        <div class="config-row" v-if="formState.operationMode !== 'ECB'">
          <div class="config-item flex-grow">
            <label>偏移量：</label>
            <div class="input-group">
              <input
                v-model="formState.deviation"
                type="text"
                :placeholder="getDeviationPlaceholder()"
              />
              <select v-model="formState.deviationFormat">
                <option
                  v-for="format in deviationFormats"
                  :key="format.value"
                  :value="format.value"
                >
                  {{ format.label }}
                </option>
              </select>
            </div>
          </div>
        </div>
      </div>

      <!-- 输入输出区域 -->
      <div class="editor-section">
        <!-- 输入区域 -->
        <div class="editor-container">
          <div class="editor-header">
            <div class="header-left">
              <label>字符编码：</label>
              <select v-model="formState.inputFormat">
                <option
                  v-for="format in textFormats"
                  :key="format.value"
                  :value="format.value"
                >
                  {{ format.label }}
                </option>
              </select>
            </div>
            <div class="header-right">
              <button class="copy-button" @click="handleCopy" title="复制">
                <i class="iconfont icon-copy" />
              </button>
            </div>
          </div>
          <div class="editor-content">
            <textarea
              v-model="formState.inputText"
              :placeholder="
                formState.mode === 'encrypt'
                  ? '请输入要加密的内容'
                  : '请输入要解密的内容'
              "
              class="editor-textarea"
            ></textarea>
          </div>
        </div>

        <!-- 操作按钮 -->
        <div class="action-bar">
          <button class="action-btn" @click="handleClear" title="清空">
            <i class="iconfont icon-ashbin" />清空
          </button>
          <button
            class="action-btn"
            @click="handleEncrypt"
            title="加密"
            :disabled="loading"
          >
            <span
              v-if="loading && formState.mode === 'encrypt'"
              class="loading-spinner"
            ></span>
            <i v-else class="iconfont icon-lock" />加密
          </button>
          <button
            class="action-btn"
            @click="handleDecrypt"
            title="解密"
            :disabled="loading"
          >
            <span
              v-if="loading && formState.mode === 'decrypt'"
              class="loading-spinner"
            ></span>
            <i v-else class="iconfont icon-unlock" />解密
          </button>
          <button class="action-btn" @click="handleExchange" title="交换">
            <i class="iconfont icon-refresh" />交换
          </button>
        </div>

        <!-- 输出区域 -->
        <div class="editor-container">
          <div class="editor-header">
            <div class="header-left">
              <label>输出格式：</label>
              <select v-model="formState.cipherTextFormat">
                <option value="Hex">Hex</option>
                <option value="Base64">Base64</option>
              </select>
              <span class="format-hint">
                (格式加密后显示输出，解密后显示输入)
              </span>
            </div>
            <div class="header-right">
              <button
                v-if="formState.outputText"
                class="copy-button"
                @click="handleCopy"
                title="复制"
              >
                <i class="iconfont icon-copy" />
              </button>
            </div>
          </div>
          <div class="editor-content">
            <textarea
              v-model="formState.outputText"
              class="editor-textarea"
              :placeholder="
                formState.mode === 'encrypt' ? '加密结果' : '解密结果'
              "
            ></textarea>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.aes-formatter {
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

.config-section {
  padding: 20px;
  border-bottom: 1px solid var(--color-border);
  background: var(--color-surface);
}

.config-row {
  display: flex;
  gap: 16px;
  margin-bottom: 16px;

  &:last-child {
    margin-bottom: 0;
  }
}

.config-item {
  display: flex;
  align-items: center;
  gap: 8px;

  &.flex-grow {
    flex: 1;
  }

  label {
    color: var(--color-text);
    font-size: 14px;
    white-space: nowrap;
    min-width: 70px;
  }
}

.input-group {
  display: flex;
  flex: 1;
  gap: 8px;

  input {
    flex: 1;
  }

  select {
    width: 120px;
  }
}

select {
  height: 32px;
  padding: 0 12px;
  border: 1px solid var(--color-border);
  border-radius: 4px;
  background: var(--color-surface);
  color: var(--color-text);
  cursor: pointer;
  transition: all 0.2s;

  &:hover {
    border-color: var(--color-primary);
  }

  &:focus {
    border-color: var(--color-primary);
    box-shadow: 0 0 0 2px var(--color-primary-light);
    outline: none;
  }
}

input {
  height: 32px;
  padding: 0 12px;
  border: 1px solid var(--color-border);
  border-radius: 4px;
  background: var(--color-surface);
  color: var(--color-text);
  transition: all 0.2s;

  &:hover {
    border-color: var(--color-primary);
  }

  &:focus {
    border-color: var(--color-primary);
    box-shadow: 0 0 0 2px var(--color-primary-light);
    outline: none;
  }

  &::placeholder {
    color: var(--color-text-light);
  }
}

.editor-section {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 16px;
  padding: 20px;
  overflow: auto;
  background: var(--color-surface);
}

.editor-container {
  background: var(--color-surface);
  border-radius: 8px;
  overflow: hidden;
}

.editor-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background: var(--color-surface);
  border-bottom: 1px solid var(--color-border);

  .header-left {
    display: flex;
    align-items: center;
    gap: 12px;

    label {
      color: var(--color-text);
      font-size: 14px;
    }
  }

  .header-right {
    display: flex;
    align-items: center;
  }
}

.editor-content {
  padding: 0;
}

.editor-textarea {
  width: 100%;
  min-height: 120px;
  padding: 16px;
  border: none;
  background: var(--color-surface);
  color: var(--color-text);
  font-family: monospace;
  font-size: 14px;
  line-height: 1.6;
  resize: vertical;

  &:focus {
    outline: none;
  }

  &::placeholder {
    color: var(--color-text-light);
  }
}

.action-bar {
  display: flex;
  justify-content: center;
  gap: 12px;
  padding: 12px;
  background: var(--color-surface);
  border: 1px solid var(--color-border);
  border-radius: 8px;
}

.action-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  min-width: 90px;
  height: 32px;
  padding: 0 16px;
  border: 1px solid var(--color-border);
  border-radius: 4px;
  background: var(--color-surface);
  color: var(--color-text);
  font-size: 14px;
  cursor: pointer;
  transition: all 0.2s;

  .iconfont {
    font-size: 16px;
    color: inherit;
  }

  &:hover:not(:disabled) {
    border-color: var(--color-primary);
    color: var(--color-primary);
    background: var(--color-surface);
  }

  &:active:not(:disabled) {
    transform: translateY(1px);
  }

  &:disabled {
    opacity: 0.6;
    cursor: not-allowed;
    background: var(--color-surface);
    color: var(--color-text);
  }
}

.loading-spinner {
  display: inline-block;
  width: 14px;
  height: 14px;
  border: 2px solid var(--color-text);
  border-top-color: transparent;
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

.text-button {
  display: flex;
  align-items: center;
  gap: 4px;
  padding: 6px 12px;
  border: none;
  border-radius: 4px;
  background: transparent;
  color: var(--color-text);
  font-size: 13px;
  cursor: pointer;
  transition: all 0.2s;

  &:hover {
    background: var(--color-surface-hover);
  }

  .iconfont {
    font-size: 14px;
  }
}

.format-hint {
  font-size: 12px;
  color: var(--color-text-secondary);
}

.copy-button {
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
  opacity: 0.6;

  &:hover {
    opacity: 1;
    border-color: var(--color-primary);
    color: var(--color-primary);
  }

  .iconfont {
    font-size: 16px;
  }
}
</style>

<script lang="ts" setup>
import { reactive, ref, watch } from "vue";
import { aesDecrypt, aesEncrypt } from "@/api/crypto";
import type { AesDecryptParams, AesEncryptParams } from "@/api/types";
import { Message } from "@/components/Message"; // 修改为正确的 Message 组件导入

// 配置选项
const operationModes = [
  { value: "CBC", label: "CBC (密码分组链接)" },
  { value: "ECB", label: "ECB (电子密码本)" },
  { value: "CFB", label: "CFB (密码反馈)" },
  { value: "OFB", label: "OFB (输出反馈)" },
  { value: "CTS", label: "CTS (密文窃取)" },
  { value: "CTR", label: "CTR (计数器)" },
  { value: "GCM", label: "GCM (伽罗瓦/计数器)" },
];

const fillModes = [
  { value: "PKCS7", label: "PKCS7" },
  { value: "Zeros", label: "Zeros" },
  { value: "No", label: "No" },
];

const keyLengths = [
  { value: "16", label: "128 bits" },
  { value: "24", label: "192 bits" },
  { value: "32", label: "256 bits" },
];

const keyFormats = [
  { value: "Hex", label: "Hex" },
  { value: "Base64", label: "Base64" },
  { value: "Text", label: "Text" },
];

const deviationFormats = [
  { value: "Hex", label: "Hex" },
  { value: "Base64", label: "Base64" },
  { value: "Text", label: "Text" },
];

const textFormats = [
  { value: "UTF-8", label: "UTF-8" },
  { value: "Hex", label: "Hex" },
  { value: "Base64", label: "Base64" },
];

// 表单状态
interface FormState {
  mode: "encrypt" | "decrypt";
  operationMode: string;
  fill: string;
  keyLength: string;
  key: string;
  keyFormat: string;
  deviation: string;
  deviationFormat: string;
  inputText: string;
  inputFormat: string;
  outputText: string;
  cipherTextFormat: string;
}

const formState = reactive<FormState>({
  mode: "encrypt",
  operationMode: "CBC",
  fill: "PKCS7",
  keyLength: "16",
  key: "",
  keyFormat: "Hex",
  deviation: "",
  deviationFormat: "Hex",
  inputText: "",
  inputFormat: "UTF-8",
  outputText: "",
  cipherTextFormat: "Base64",
});

const loading = ref(false);

// 获取偏移量提示文本
const getDeviationPlaceholder = () => {
  if (formState.operationMode === "GCM") {
    return "请输入12字节偏移量";
  }
  return "请输入16字节偏移量";
};

// 处理加密解密
const handleProcess = async () => {
  if (!formState.key || !formState.inputText) {
    Message.warning("请填写必要信息");
    return;
  }

  loading.value = true;
  try {
    if (formState.mode === "encrypt") {
      const params: AesEncryptParams = {
        plainText: formState.inputText,
        plainTextFormat: formState.inputFormat,
        operationMode: formState.operationMode,
        fill: formState.fill,
        key: formState.key,
        keyFormat: formState.keyFormat,
        deviation: formState.deviation || "",
        deviationFormat: formState.deviationFormat,
        cipherTextFormat: "Base64",
      };
      const { data } = await aesEncrypt(params);
      formState.outputText = data;
      Message.success("加密成功");
    } else {
      const params: AesDecryptParams = {
        cipherText: formState.inputText,
        cipherTextFormat: formState.inputFormat,
        operationMode: formState.operationMode,
        fill: formState.fill,
        key: formState.key,
        keyFormat: formState.keyFormat,
        deviation: formState.deviation || "",
        deviationFormat: formState.deviationFormat,
        plainTextFormat: "UTF-8",
      };
      const { data } = await aesDecrypt(params);
      formState.outputText = data;
      Message.success("解密成功");
    }
  } catch (error: any) {
    Message.error(error.message);
  } finally {
    loading.value = false;
  }
};

// 添加监听器，当操作模式改变时，重置偏移量
watch(
  () => formState.operationMode,
  (newMode) => {
    if (newMode === "ECB") {
      formState.deviation = "";
    }
  },
);

// 清空表单
const handleClear = () => {
  formState.key = "";
  formState.deviation = "";
  formState.inputText = "";
  formState.outputText = "";
  Message.info("内容已清空");
};

// 复制结果
const handleCopy = async () => {
  try {
    await navigator.clipboard.writeText(formState.outputText);
    Message.success("已复制到剪贴板");
  } catch (err) {
    Message.error("复制失败");
  }
};

// 处理加密
const handleEncrypt = async () => {
  formState.mode = "encrypt";
  await handleProcess();
};

// 处理解密
const handleDecrypt = async () => {
  formState.mode = "decrypt";
  await handleProcess();
};

// 处理交换
const handleExchange = () => {
  // 交换输入和输出内容
  const tempText = formState.inputText;
  formState.inputText = formState.outputText;
  formState.outputText = tempText;

  // 交换格式
  const tempFormat = formState.inputFormat;
  formState.inputFormat = formState.cipherTextFormat;
  formState.cipherTextFormat = tempFormat;

  // 切换模式
  formState.mode = formState.mode === "encrypt" ? "decrypt" : "encrypt";

  Message.info(`已切换为${formState.mode === "encrypt" ? "加密" : "解密"}模式`);
};
</script>
