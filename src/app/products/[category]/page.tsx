import { ByCategoryPage } from '@/components/ByCategoryPage'
import styles from './ProductByCategoryPage.module.scss'
import { getAllData } from '@/utils/DataFetching'


const ProductsByCategoryPage = async ({params} : any) =>{
    const products = await getAllData()
    console.log(products)
    return (
        <ByCategoryPage params={products} />
    )
}

export default ProductsByCategoryPage