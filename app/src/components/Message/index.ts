import { App, VNode, ComponentPublicInstance, createVNode, render } from "vue";
import MessageComponent from "./MessageComponent.vue";

export type MessageType = "success" | "warning" | "info" | "error";

export interface MessageOptions {
  message: string;
  type?: MessageType;
  duration?: number;
  center?: boolean;
  onClose?: () => void;
}

interface MessageInstance {
  id: string;
  vnode: VNode;
  vm: ComponentPublicInstance | null;
  container: HTMLElement;
  props: MessageOptions;
}

type MessageMethod = (arg: MessageOptions | string) => { close: () => void };

type MessageFunction = MessageMethod & {
  success: MessageMethod;
  warning: MessageMethod;
  info: MessageMethod;
  error: MessageMethod;
  closeAll: () => void;
};

const instances: MessageInstance[] = [];
let seed = 1;

const Message = ((options: MessageOptions | string) => {
  if (typeof options === "string") {
    options = {
      message: options,
    };
  }

  const id = `message_${seed++}`;
  const container = document.createElement("div");

  const props = {
    ...options,
    onClose: () => {
      if (typeof options !== "string" && options.onClose) {
        options.onClose();
      }
      close(id);
    },
  };

  const vnode = createVNode(MessageComponent, props);
  render(vnode, container);
  document.body.appendChild(container);

  const instance: MessageInstance = {
    id,
    vnode,
    vm: vnode.component?.proxy || null,
    container,
    props: props as MessageOptions,
  };

  instances.push(instance);

  return {
    close: () => close(id),
  };
}) as MessageFunction;

// 添加类型方法
(["success", "warning", "info", "error"] as const).forEach((type) => {
  Message[type] = ((messageOptions: MessageOptions | string) => {
    if (typeof messageOptions === "string") {
      messageOptions = {
        message: messageOptions,
      };
    }
    return Message({
      ...messageOptions,
      type,
    });
  }) as MessageMethod;
});

const close = (id: string): void => {
  const idx = instances.findIndex((instance) => instance.id === id);
  if (idx === -1) return;

  const { container } = instances[idx];
  instances.splice(idx, 1);
  render(null, container);
  document.body.removeChild(container);
};

Message.closeAll = (): void => {
  instances.forEach(({ container }) => {
    render(null, container);
    document.body.removeChild(container);
  });
  instances.length = 0;
};

export default {
  install: (app: App): void => {
    app.config.globalProperties.$message = Message;
  },
};

export { Message };
