import YourProduct from '@/ui/YourProduct'
import styles from './AdvertisementsPage.module.scss'

const AdvertisementsPage = () =>{
    return (
        <div className={styles.container}>

            <div className={styles.wrapper}>
                <YourProduct />
                <YourProduct />
                <YourProduct />
                <YourProduct />
                <YourProduct />
                <YourProduct />
            </div>
        </div>
    )
}

export default AdvertisementsPage