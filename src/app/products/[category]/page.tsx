"use client"
import Product from '@/ui/Product'
import { useState } from 'react'
import styles from './ProductByCategoryPage.module.scss'
import PriceFilter from '@/components/PriceFilter'
import { getAllData } from '@/utils/DataFetching'

const products = getAllData()

const ProductsByCategoryPage = ({params} : any) =>{
    const [minPrice, setMinPrice] = useState(0);
    const [maxPrice, setMaxPrice] = useState(1000);
    const [currentProducts, setProducts] = useState(products);

    
    function handlePriceChange(newMinPrice: number, newMaxPrice: number) {
      setMinPrice(newMinPrice);
      setMaxPrice(newMaxPrice);
    }
    return (
        <div className={styles.container}>
            <div className={styles.filters}>
                <PriceFilter />
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