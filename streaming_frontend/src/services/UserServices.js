import userClient from "../clients/UserClient.js";

const UserService = {
  createUser: async (body) => {
    const { data } = await userClient.post("/users", body);
    return data;
  },

  logInUser: async (body) => {
    const { data } = await userClient.post("/logIn", body);
    return data;
  },
};

export default UserService;
