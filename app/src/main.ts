import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import Message from "@/components/Message";
import "./main.css";

const app = createApp(App);
app.use(store);
app.use(router);
app.use(Message);
app.mount("#app");
