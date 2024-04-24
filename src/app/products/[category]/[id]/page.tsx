import styles from './ProductByIDPage.module.scss'
import Slider from '@/components/Slider'
import Product from '@/ui/Product'
import { getAllData, getProductById } from '@/utils/DataFetching'
const ProductsByIDPage = async ({params} : any) =>{
    const {id} = params
    const product = await getProductById(id)
    const products = await getAllData()
    return (
        <div className={styles.container}>
            <div className={styles.info}>
                <div className={styles.main}>
                    <h1>{product.name}</h1>
                    <Slider product = {product}/>
                </div>
                <div className={styles.desc}>
                    <h2>Описание:</h2>
                    <div>{product.description}</div>
                    <div className={styles.cost}>
                        <h2>Цена:</h2>
                        <span>{product.cost} ₽</span>
                    </div>
                </div>
            </div>
            <div className={styles.contacts}>
                <h2>Контакты:</h2>
            </div>
            <div className={styles.wrapper}>
                <Product product={products.listings[0]}/>
                <Product product={products.listings[1]} />
                <Product product={products.listings[3]}/>
                <Product product={products.listings[4]}/>
            </div>
        </div>
    )
}

export default ProductsByIDPage

