import { AuthProvider } from "react-admin";
import { config } from "./config.ts";

const authProvider: AuthProvider = {
  login: async ({ username, password }) => {
    const request = new Request(`${config.apiUrl}/auth/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: username,
        password: password,
      }),
    });

    const response = await fetch(request);

    if (response.status < 200 || response.status >= 300) {
      throw new Error(response.statusText);
    }
    // TODO: add admin checking
    // server set cookies and login is successfull
  },

  checkAuth: async () => {
    return Promise.resolve(); // TODO: add /me endpoint
  },

  checkError: async (error) => {
    if (error.status === 401 || error.status === 403) {
      return Promise.reject();
    }
    return Promise.resolve();
  },

  logout: async () => {
    await fetch(`${config.apiUrl}/logout`, {
      method: "POST",
      credentials: "include",
    });
    // server deleted cookies

    return Promise.resolve();
  },

  getIdentity: async () => {
    return { id: 1, name: "John Doe" }; // same TODO as above
  },
};

export default authProvider;
