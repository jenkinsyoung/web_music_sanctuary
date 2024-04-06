import axios from "axios";

const URL = ''
//get запросы
export const getAllData = async() =>{
    try{
        const response = await axios.get(`${URL}`)
        return response.data
    } catch (err: any){
        console.error(err.toJSON())
    }
}

export const getProductById = async(ID:number) =>{
    try{
        const response = await axios.get(`${URL}?id=${ID}`)
        return response.data
    } catch (err: any){
        console.error(err.toJSON())
    }
}

export const getAllAdds = async() =>{
    try{
        const response = await axios.get(`${URL}`)
        return response.data
    } catch (err: any){
        console.error(err.toJSON())
    }
}

export const getAddById = async(ID:number) =>{
    try{
        const response = await axios.get(`${URL}?id=${ID}`)
        return response.data
    } catch (err: any){
        console.error(err.toJSON())
    }
}

export const getUserById = async(ID:number) =>{
    try{
        const response = await axios.get(`${URL}?id=${ID}`)
        return response.data
    } catch (err: any){
        console.error(err.toJSON())
    }
}

//post запросы

export const postNewAdd = (data: any) => {
    axios.post(`${URL}`, data)
}

//put запрос

export const putUpdateAdd = (data: any)=>{
    axios.put(`${URL}`, data)
}

//delete запрос

export const deleteAddById = (ID: number)=>{
    axios.delete(`${URL}`)
}