import UploadImage from '@/components/UploadImage'
import styles from './ProductPage.module.scss'
import Image from 'next/image'
import Link from 'next/link'

const ProductPage =()=>{
    return(
        <div className={styles.container}>
            <div className={styles.first_line}>
                <Link href={`/products/acoustic_guitar`}>
                <div className={styles.acoustic}>
                    <h1>acoustic guitar</h1>
                    <Image priority
                    src = '/acoustic.svg'
                    alt=''
                    width={506.29}
                    height={445.28}/> 
                </div>
                </Link>
                <Link href={`/products/electric_guitar`}>
                <div className={styles.electric}>
                    <h1>electric guitar</h1>
                    <Image priority
                    src = '/electric.svg'
                    alt=''
                    height={465}
                    width={750}
                    /> 
                </div>
                </Link>
            </div>
            <div className={styles.second_line}>
            <Link href={`/products/bass_guitar`}>
                <div className={styles.bass}>
                        <h1>bass guitar</h1>
                        <Image priority
                        src = '/bass.svg'
                        alt=''
                        width={500}
                        height={433}/> 
                </div>
            </Link>
            <Link href={`/products/classical_guitar`}>
                <div className={styles.classic}>
                    <h1>classical guitar</h1>
                    <Image priority
                    src = '/guitar.svg'
                    alt=''
                    height={390.88}
                    width={427.1}
                    /> 
                </div>
            </Link>
            <Link href={`/products/ukulele`}>
                <div className={styles.ukulele}>
                        <h1>ukulele</h1>
                        <Image priority
                        src = '/ukulele.svg'
                        alt=''
                        width={340}
                        height={401}/> 
                </div>  
            </Link>  
            </div>

            <UploadImage />
        </div>
    )
}
export default ProductPage