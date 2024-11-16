import { Module } from "vuex";
import { RootState } from "@/store";

export interface JsonEditorState {
  content: string;
  isExpanded: boolean;
}

const jsonEditor: Module<JsonEditorState, RootState> = {
  namespaced: true,

  state: {
    content: "",
    isExpanded: true,
  },

  mutations: {
    SET_CONTENT(state, content: string) {
      state.content = content;
    },
    SET_EXPANDED(state, isExpanded: boolean) {
      state.isExpanded = isExpanded;
    },
    CLEAR_CONTENT(state) {
      state.content = "";
      state.isExpanded = true;
    },
  },

  actions: {
    updateContent({ commit }, content: string) {
      commit("SET_CONTENT", content);
    },
    updateExpanded({ commit }, isExpanded: boolean) {
      commit("SET_EXPANDED", isExpanded);
    },
    clearEditor({ commit }) {
      commit("CLEAR_CONTENT");
    },
  },

  getters: {
    getContent: (state) => state.content,
    getIsExpanded: (state) => state.isExpanded,
  },
};

export default jsonEditor;
