import { api } from "./http";
import { PostDTO, ParamsDTO, CreatePostDTO, UpdatePostDTO } from "@/types/dto/PostDtos";

// API service
export const postApi = {
  getPost: (id: number): Promise<PostDTO> => api.get(`/post/${id}`).then(r => r.data),

  getPosts: (params?: ParamsDTO): Promise<PostDTO[]> => api.get("/post", { params }).then(r => r.data),

  createPost: (data: CreatePostDTO) => api.post("/post", data).then(r => r.data),

  updatePost: (id: number, data: UpdatePostDTO) => api.put(`/post/${id}`, data).then(r => r.data),

  deletePost: (id: number): Promise<void> => api.delete(`/post/${id}`),
}