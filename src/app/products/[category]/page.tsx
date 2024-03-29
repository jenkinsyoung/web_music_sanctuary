import Product from '@/ui/Product'
import styles from './ProductByCategoryPage.module.scss'

const ProductsByCategoryPage = ({params} : any) =>{
    return (
        <div className={styles.container}>
            <div className={styles.filters}>
                Здесь нужны фильтры
            </div>

            <div className={styles.wrapper}>
                <Product />
                <Product />
                <Product />
                <Product />
                <Product />
                <Product />
                <Product />
                <Product />
                <Product />
            </div>
        </div>
    )
}

export default ProductsByCategoryPage