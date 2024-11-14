<template>
  <fluent-card>
    <h2>{{ msg }}</h2>
    <fluent-button appearance="accent" v-on:click="onClick">
      Click Me
    </fluent-button>
  </fluent-card>
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";
import { ref } from "vue";
import FluentCard from "@fluentui/web-components";
import FluentButton from "@fluentui/web-components";

@Options({
  components: {
    FluentCard,
    FluentButton,
  },
})
export default class App extends Vue {
  private msg = ref("Hello, World!");

  public onClick(): void {
    // 处理点击事件
    console.log("Button clicked!");
  }
}
</script>

<style></style>
