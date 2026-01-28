import  axios from 'axios';

// Infrastructure
export const api = axios.create({
    baseURL: process.env.NEXT_PUBLIC_API_URL,
    withCredentials: true,
});

api.interceptors.response.use(res => res, onError)

async function onError(error: any) {
    const original  = error.config;
    if (error.response?.status == 401 && !original._retry) {
        original._retry = true;
        await api.post("auth/refresh");
        return api(original);
    }

    throw error;
}