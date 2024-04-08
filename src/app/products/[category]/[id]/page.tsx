"use client"
import styles from './ProductByIDPage.module.scss'
import Slider from '@/components/Slider'
import Product from '@/ui/Product'
import { usePathname } from 'next/navigation'
import data from '@/utils/data.json'
const ProductsByIDPage = ({params} : any) =>{
    const path = usePathname();
    const id = Number(path.split('/')[3]);
    const product = data[id]
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
                <Product product={data[0]}/>
                <Product product={data[1]} />
                <Product product={data[3]}/>
                <Product product={data[4]}/>
            </div>
        </div>
    )
}

export default ProductsByIDPage

