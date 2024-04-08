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
import products  from '@/utils/data.json'
import { getVisibleProduct, getVisibleProductByForm, getVisibleProductByPick } from '@/services/product-queries'

const ProductsByCategoryPage = ({params} : any) =>{
    const path = usePathname();
    const filter = path.split('/')[2];
    const current = products.filter((product:any)=>(
        product.category === filter
    ))
    
    const [minPrice, setMinPrice] = useState(0);
    const [maxPrice, setMaxPrice] = useState(160000);
    const [currentProducts, setProducts] = useState(current);
    const [currentForm, setForms] = useState(['']);
    function handlePriceChange(newMinPrice: number, newMaxPrice: number) {
      setMinPrice(newMinPrice);
      setMaxPrice(newMaxPrice);
    }

    const onChangeCategoryHandler = (form : any, isChecked : any) =>{
        isChecked
        ? setForms((prevForms) => [...prevForms, form])
        : setForms(
            currentForm.filter(
                card => card !== form
            )
            
        )
        
    }
    
    const items1 = getVisibleProductByPick(currentForm, currentProducts) 
    const items2 = getVisibleProductByForm(currentForm, currentProducts)
    const items = () =>{
        if (items1.length > items2.length) return items1.filter((value : any) => items2.includes(value))

        if (items1.length <= items2.length)return items2.filter((value : any) => items1.includes(value))
    }

    return (
        <div className={styles.container}>
            <div className={styles.search}>
                <input type="text" placeholder='Поиск по названию' />
                <button type="submit">Поиск</button>
            </div>
            <div className={styles.filters}>
                <PriceFilter min={minPrice} max ={maxPrice}/>
                {filter == 'acoustic_guitar' ? <FormAcoustic currentForm = {currentForm} onChangeCategory={onChangeCategoryHandler}/> : ''}
                {filter == 'electric_guitar' ? <FormElectro currentForm = {currentForm} onChangeCategory={onChangeCategoryHandler}/> : ''}
                {filter == 'electric_guitar' ? <Pickups currentForm = {currentForm} onChangeCategory={onChangeCategoryHandler}/> : ''}
            </div>

            <div className={styles.wrapper}>
                {items().map((product:any) =><Product key={product.id} product={product}/>)} 
            </div>
        </div>
    )
}

export default ProductsByCategoryPage