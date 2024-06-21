import axios from "axios";

const userClient = axios.create({
  baseURL: "http://localhost:4000",
});

export default userClient;