import { logInDTO, signInDTO } from "@/types/dto/AuthDtos";
import { api } from "./http";

// API service
export const authApi = {
    login: (data: logInDTO) => api.post("/auth/login", data),

    register: (data: signInDTO) => api.post("/auth/signin", data),

    logout: () => api.post("/auth/logout"),
}