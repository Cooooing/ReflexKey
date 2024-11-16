<template>
  <transition name="el-message-fade">
    <div
      v-show="visible"
      class="el-message"
      :class="[`el-message--${type}`]"
      @mouseenter="clearTimer"
      @mouseleave="startTimer"
    >
      <div class="el-message__content">
        <i :class="typeClass"></i>
        <span class="message-text">{{ message }}</span>
      </div>
      <div class="el-message__closeBtn" @click="close">
        <i class="iconfont icon-close-bold"></i>
      </div>
    </div>
  </transition>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, PropType } from "vue";

type MessageType = "success" | "warning" | "info" | "error";

const props = defineProps({
  message: {
    type: String,
    required: true,
  },
  type: {
    type: String as PropType<MessageType>,
    default: "info",
  },
  duration: {
    type: Number,
    default: 3000,
  },
  onClose: {
    type: Function as PropType<() => void>,
    default: undefined,
  },
});

const visible = ref(false);
let timer: ReturnType<typeof setTimeout> | null = null;

const typeClass = computed(() => ({
  iconfont: true,
  "el-message__icon": true,
  "icon-success": props.type === "success",
  "icon-warning": props.type === "warning",
  "icon-info": props.type === "info",
  "icon-error": props.type === "error",
}));

const close = () => {
  visible.value = false;
  props.onClose?.();
};

const clearTimer = () => {
  if (timer) {
    clearTimeout(timer);
    timer = null;
  }
};

const startTimer = () => {
  if (props.duration > 0) {
    timer = setTimeout(() => {
      close();
    }, props.duration);
  }
};

onMounted(() => {
  visible.value = true;
  startTimer();
});
</script>

<style lang="scss" scoped>
.el-message {
  min-width: 380px;
  max-width: 800px;
  width: auto;
  box-sizing: border-box;
  border-radius: 4px;
  border: 1px solid #ebeef5;
  position: fixed;
  left: 50%;
  top: 20px;
  transform: translateX(-50%);
  background-color: #edf2fc;
  padding: 15px 30px 15px 15px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  transition: opacity 0.3s, transform 0.4s;
  z-index: 2000;

  .el-message__content {
    flex: 1;
    display: flex;
    align-items: flex-start;
    justify-content: center;
    gap: 8px;
    max-width: calc(100% - 40px);

    .message-text {
      font-size: 14px;
      line-height: 1.6;
      word-break: break-word;
      white-space: pre-wrap;
      text-align: left;
      flex: 1;
      padding-right: 10px;
    }

    .el-message__icon {
      font-size: 16px;
      margin-top: 2px;
    }
  }

  .el-message__closeBtn {
    position: absolute;
    top: 15px;
    right: 15px;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 4px;
    border-radius: 4px;
    transition: background-color 0.2s;

    &:hover {
      background-color: rgba(0, 0, 0, 0.06);
    }

    .iconfont {
      font-size: 12px;
      color: #909399;
    }
  }

  &--success {
    background-color: #f0f9eb;
    border-color: #e1f3d8;

    .el-message__content span,
    .el-message__icon {
      color: #67c23a;
    }
  }

  &--warning {
    background-color: #fdf6ec;
    border-color: #faecd8;

    .el-message__content span,
    .el-message__icon {
      color: #e6a23c;
    }
  }

  &--info {
    background-color: #edf2fc;
    border-color: #ebeef5;

    .el-message__content span,
    .el-message__icon {
      color: #909399;
    }
  }

  &--error {
    background-color: #fef0f0;
    border-color: #fde2e2;

    .el-message__content span,
    .el-message__icon {
      color: #f56c6c;
    }
  }
}

.el-message-fade-enter-from,
.el-message-fade-leave-to {
  opacity: 0;
  transform: translate(-50%, -100%);
}
</style>
