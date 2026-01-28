
export interface UserDTO {
    id: number,
    createdAt: number,
    name: string,
    isAdmin: string
}

export interface CreateUserDTO {
    name: string,
    passwrod: string
}

export interface UpdateUserDTO {
    name: string,
    passwrod: string
}