import { CreateUserDTO, UpdateUserDTO, UserDTO } from "@/types/dto/UserDtos";
import { api } from "./http";

// API service
export const userApi = {
    getUser: (id: number): Promise<UserDTO> => api.get(`/user/${id}`).then(r => r.data),

    createUser: (data: CreateUserDTO) => api.post("/user", data).then(r => r.data),

    updateUser: (id: number, data: UpdateUserDTO) => api.put(`/user/${id}`, data),

    deleteUser: (id: number): Promise<void> => api.delete(`/user/${id}`),
}