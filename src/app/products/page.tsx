import styles from './ProductPage.module.scss'
import Image from 'next/image'

const ProductPage =()=>{
    return(
        <div className={styles.container}>
            <div className={styles.first_line}>
                <div className={styles.acoustic}>
                    <h1>acoustic guitar</h1>

                </div>
            </div>
            <div className={styles.second_line}>

            </div>
        </div>
    )
}
export default ProductPage