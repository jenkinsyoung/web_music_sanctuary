import YourProduct from '@/ui/YourProduct'
import styles from './AdvertisementsPage.module.scss'
import products from '@/utils/data.json'
const AdvertisementsPage = () =>{
    return (
        <div className={styles.container}>

            <div className={styles.wrapper}>
            {products.map((product:any) =><YourProduct key={product.id} product={product}/>)} 
            </div>
        </div>
    )
}

export default AdvertisementsPage