import { AuthProvider } from "react-admin";
import { config } from "./config.ts";

const authProvider: AuthProvider = {
  login: async ({ username, password }) => {
    const request = new Request(`${config.apiUrl}/auth/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify({
        name: username,
        password: password,
      }),
    });

    const response = await fetch(request);

    if (response.status < 200 || response.status >= 300) {
      throw new Error(response.statusText);
    }
    return Promise.resolve();
    // server set cookies and login is successfull
  },

  checkAuth: async () => {
    const res = await fetch(`${config.apiUrl}/auth/me`, {
      method: "POST",
      credentials: "include",
    });

    const data = await res.json();
    if (res.ok && data.is_admin) return Promise.resolve();
    return Promise.reject();
  },

  checkError: async (error) => {
    if (error.status === 401 || error.status === 403) {
      return Promise.reject();
    }
    return Promise.resolve();
  },

  logout: async () => {
    await fetch(`${config.apiUrl}/auth/logout`, {
      method: "POST",
      credentials: "include",
    });
    // server deleted cookies

    return Promise.resolve();
  },

  getIdentity: async () => {
    const res = await fetch(`${config.apiUrl}/auth/me`, {
      method: "POST",
      credentials: "include",
    });

    if (!res.ok) return Promise.reject();

    const credentials = await res.json();

    const { id, name } = credentials;

    return { id, name };
  },
};

export default authProvider;
