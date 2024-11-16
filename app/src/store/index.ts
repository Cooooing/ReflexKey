import { createStore } from "vuex";
import theme from "./modules/theme";
import jsonEditor from "./modules/jsonEditor";
import { JsonEditorState } from "./modules/jsonEditor";

// 定义根状态接口
export interface RootState {
  theme: {
    currentTheme: string;
  };
  jsonEditor: JsonEditorState;
}

export default createStore<RootState>({
  modules: {
    theme,
    jsonEditor,
  },
});
