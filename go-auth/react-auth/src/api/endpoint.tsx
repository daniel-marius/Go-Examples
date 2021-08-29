import axios from "axios";

const baseURL: string | undefined = "http://localhost:4000/api";
const headers: any = {
  "Content-Type": "application/json"
};

export default axios.create({
  baseURL,
  headers
});
