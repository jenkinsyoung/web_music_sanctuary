import styles from './ProductPage.module.scss'
import Image from 'next/image'

const ProductPage =()=>{
    return(
        <div className={styles.container}>
            <div className={styles.first_line}>
                <div className={styles.acoustic}>
                    <h1>acoustic guitar</h1>
                    <Image priority
                    src = '/acoustic.svg'
                    alt=''
                    width={506.29}
                    height={445.28}/> 
                </div>
                <div className={styles.electric}>
                    <h1>electric guitar</h1>
                    <Image priority
                    src = '/electric.svg'
                    alt=''
                    height={465}
                    width={750}
                    /> 
                </div>
            </div>
            <div className={styles.second_line}>
                <div className={styles.bass}>
                        <h1>bass guitar</h1>
                        <Image priority
                        src = '/bass.svg'
                        alt=''
                        width={500}
                        height={433}/> 
                </div>
                <div className={styles.classic}>
                    <h1>electric guitar</h1>
                    <Image priority
                    src = '/guitar.svg'
                    alt=''
                    height={390.88}
                    width={427.1}
                    /> 
                </div>
                <div className={styles.ukulele}>
                        <h1>ukulele</h1>
                        <Image priority
                        src = '/ukulele.svg'
                        alt=''
                        width={340}
                        height={401}/> 
                </div>    
            </div>
        </div>
    )
}
export default ProductPage