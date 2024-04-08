"use client"
import styles from './AdvertisementsByIDPage.module.scss'
import Slider from '@/components/Slider'
import Product from '@/ui/Product'
import { usePathname } from 'next/navigation'
import data from '@/utils/data.json'
const AdvertisementsByIDPage = ({params} : any) =>{
    const path = usePathname();
    const id = Number(path.split('/')[2]);
    const product = data[id-1]
    return (
        <div className={styles.container}>
            <form className={styles.info}>
                <div className={styles.main}>
                    <input className={styles.h1} placeholder={`${product.name}`}/>
                    <Slider product={product}/>
                </div>
                <div className={styles.desc}>
                    <h2>Описание:</h2>
                    <div><textarea placeholder={`${product.description}`} /></div>
                    <div className={styles.cost}>
                        <h2>Цена:</h2>
                        <span><input type='number' placeholder={`${product.cost}`}/>₽</span>
                    </div>

                    <button type='submit'>Сохранить</button>
                </div>
            </form>
            <div className={styles.contacts}>
                <h2>Контакты:</h2>
            </div>
        </div>
    )
}

export default AdvertisementsByIDPage

