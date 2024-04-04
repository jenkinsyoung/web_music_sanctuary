import axios from "axios";

const URL = ''

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
        const response = await axios.get(`${URL}`)
        return response.data
    } catch (err: any){
        console.error(err.toJSON())
    }
}

export const getAddById = async(ID:number) =>{
    try{
        const response = await axios.get(`${URL}`)
        return response.data
    } catch (err: any){
        console.error(err.toJSON())
    }
}