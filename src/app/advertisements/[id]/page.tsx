import styles from './AdvertisementsByIDPage.module.scss'
import Slider from '@/components/Slider'
import Product from '@/ui/Product'

const AdvertisementsByIDPage = ({params} : any) =>{
    return (
        <div className={styles.container}>
            <form className={styles.info}>
                <div className={styles.main}>
                    <input className={styles.h1} placeholder={`Гитара 2239402-21ш349`}/>
                    <Slider />
                </div>
                <div className={styles.desc}>
                    <h2>Описание:</h2>
                    <div><textarea placeholder={`Lorem ipsum dolor sit amet consectetur adipisicing elit. Tempore quos iste ipsam reiciendis soluta nemo iusto vero quis unde! Beatae id quidem rerum eligendi culpa possimus dolorum nam facere? Culpa.`} /></div>
                    <div className={styles.cost}>
                        <h2>Цена:</h2>
                        <span><input type='number' placeholder={`12000`}/>₽</span>
                    </div>

                    <button type='submit'>Сохранить</button>
                </div>
            </form>
            <div className={styles.contacts}>
                <h2>Контакты:</h2>
            </div>
            <div className={styles.wrapper}>
                <Product />
                <Product />
                <Product />
                <Product />
            </div>
        {/* <UploadImage /> */}
        </div>
    )
}

export default AdvertisementsByIDPage

