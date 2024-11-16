declare module "@fluentui/web-components" {
  import type { DefineComponent } from "vue";

  export const FluentCard: DefineComponent<{}>;
  export const FluentButton: DefineComponent<{
    appearance?: "accent" | "lightweight" | "neutral" | "outline" | "stealth";
  }>;

  export const allComponents: any[];

  export function provideFluentDesignSystem(): {
    register: (components: any[]) => void;
  };
}
