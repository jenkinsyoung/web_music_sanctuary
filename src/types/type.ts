export interface Category{
    id: number,
    name: string
}

export interface Advertisement{
    id: number,
    user_id: number,
    description?: string,
    name?: string,
    cost?: number,
    images?: JSON,
    type_id: number
}