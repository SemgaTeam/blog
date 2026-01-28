
export interface PostDTO {
    id: number,
    name: string,
    authorId: number,
    createdAt: string,
    updatedAt: string,
    contents: string
}

export interface ParamsDTO {
    ids?: number[],
    page?: number,
    perPage?: number,
    sortField?: number,
    sortOrder?: number,
    name?: number,
    authorId?: number
}

export interface CreatePostDTO {
    authorId?: number,
    name: string,
    contents: string
}

export interface UpdatePostDTO {
    name: string,
    contents: string
}