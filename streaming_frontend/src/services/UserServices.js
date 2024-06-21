import userClient from "../clients/UserClient.js";

const UserService = {
  createUser: async (body) => {
    const { data } = await userClient.post("/users", body);
    return data;
  },
};

export default UserService;
