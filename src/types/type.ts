export interface Product{
    id: number,
    user_id: number,
    guitar_id: number,
    name: string,
    cost: number,
    description: string
}

export interface Guitar{
    id: number,
    form: string,
    pickup_config: string,
    category: string
}

export interface User{
    id: number,
    name: string,
    surname: string,
    patronymic: string,
    email: string,
    password: string,
    phone: string
}