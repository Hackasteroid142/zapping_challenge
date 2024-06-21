import { createStore } from "vuex";
import UserService from "@/services/UserServices";
import { jwtDecode } from "jwt-decode";

const store = createStore({
  state: {
    token: localStorage.getItem("token") || null,
    user: localStorage.getItem("user") || "Usuario",
  },
  mutations: {
    setToken(state, token) {
      state.token = token;
      localStorage.setItem("token", token);
    },
    setUser(state, user) {
      state.user = user;
      localStorage.setItem("user", user);
    },
    delData(state) {
      state.user = "";
      state.token = null;
      localStorage.removeItem("token");
      localStorage.removeItem("user");
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
        commit("setUser", userToken.name);
      } catch (error) {
        throw error;
      }
    },
    logout({ commit }) {
      commit("delData");
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
