import styles from './Product.module.scss'
import Link from 'next/link'

const Product = ({product} : any) => {
    return (
        <div className={styles.container}>
            <h1>{product.name}</h1>
            <div className={styles.img} style={{backgroundImage: `url(/${product.image[0].src})`}} />
            <div className={styles.tags}>#{product.pickup_configuration} #{product.form}</div>
            <div className={styles.cost}><p>Цена:</p><span>{product.cost} ₽</span></div>
            <Link href={`/products/${product.category}/${product.id}`}><button>see more</button></Link>
        </div>
    )
  
}

export default Product