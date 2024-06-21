import { createStore } from "vuex";
import UserService from "@/services/UserServices";

const store = createStore({
  state: {
    token: localStorage.getItem("token") || "",
    user: JSON.parse(localStorage.getItem("user")) || null,
  },
  mutations: {
    setToken(state, token) {
      state.token = token;
    },
  },
  actions: {
    async createUser({ commit }, body) {
      try {
        const response = await UserService.createUser(body);
        return response;
      } catch (error) {
        throw error;
      }
    },
  },
  getters: {
    token(state) {
      return state.token;
    },
    user(state) {
      return state.user;
    },
  },
});

export default store;
