import { createStore } from "vuex";
import UserService from "@/services/UserServices";
import { jwtDecode } from "jwt-decode";

const store = createStore({
  state: {
    token: localStorage.getItem("token") || "",
    user: JSON.parse(localStorage.getItem("user")) || null,
  },
  mutations: {
    setToken(state, token) {
      state.token = token;
    },
    setUser(state, user) {
      state.user = user;
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
    async logInUser({ commit }, body) {
      try {
        const response = await UserService.logInUser(body);
        const userToken = jwtDecode(response.token);
        commit("setToken", response.token);
        commit("setUser", userToken);
      } catch (error) {
        throw error;
      }
    },
  },
  getters: {
    getToken(state) {
      return state.token;
    },
    getUser(state) {
      return state.user;
    },
  },
});

export default store;
