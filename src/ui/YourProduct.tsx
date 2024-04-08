import styles from './YourProduct.module.scss'
import Link from 'next/link'

const YourProduct = ({product} : any) => {
    return (
        <div className={styles.container}>
            <h1>{product.name}</h1>
            <div className={styles.img} style={{backgroundImage: `url(/${product.image[0].src})`}} />
            <div className={styles.tags}>#{product.pickup_configuration} #{product.form}</div>
            <div className={styles.cost}><p>Цена:</p><span>{product.cost} ₽</span></div>
            <div style={{display: 'flex'}}>
            <Link href={`/advertisements/${product.id}`}><button className={styles.edit}>edit</button></Link>
            <button className={styles.delete}>delete</button>
            </div>
            
        </div>
    )
  
}

export default YourProduct