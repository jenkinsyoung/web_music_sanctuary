"use client"
import Product from '@/ui/Product'
import { useState } from 'react'
import styles from './ProductByCategoryPage.module.scss'
import PriceFilter from '@/components/PriceFilter'
import { getAllData } from '@/utils/DataFetching'
import { usePathname } from 'next/navigation'
import FormAcoustic from '@/components/FormAcoustic'
import FormElectro from '@/components/FormElectro'
import Pickups from '@/components/Pickups'

const products = getAllData()

const ProductsByCategoryPage = ({params} : any) =>{
    const [minPrice, setMinPrice] = useState(0);
    const [maxPrice, setMaxPrice] = useState(1000);
    const [currentProducts, setProducts] = useState(products);
    
    function handlePriceChange(newMinPrice: number, newMaxPrice: number) {
      setMinPrice(newMinPrice);
      setMaxPrice(newMaxPrice);
    }
    
    const path = usePathname();
    const filter = path.split('/')[2];

    return (
        <div className={styles.container}>
            <div className={styles.search}>
                <input type="text" placeholder='Поиск по названию' />
                <button type="submit">Поиск</button>
            </div>
            <div className={styles.filters}>
                <PriceFilter min={minPrice} max ={maxPrice}/>
                {filter == 'acoustic_guitar' ? <FormAcoustic /> : ''}
                {filter == 'electric_guitar' ? <FormElectro /> : ''}
                {filter == 'electric_guitar' ? <Pickups /> : ''}
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