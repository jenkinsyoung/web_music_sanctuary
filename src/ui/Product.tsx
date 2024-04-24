import styles from './Product.module.scss'
import Link from 'next/link'

const Product = ({product} : any) => {
    return (
        <div className={styles.container}>
            <h1>{product.guitar_name}</h1>
            {product.img_list != null ? <div className={styles.img} style={{backgroundImage: `url(data:image/png;base64,${product.img_list[0].image})`}} />:<div className={styles.img} />}
            {product.category == 'Electric'?<div className={styles.tags}>#{product.pickup_config} #{product.guitar_form}</div>:
            <div className={styles.tags}>#{product.guitar_form}</div>}
            <div className={styles.cost}><p>Цена:</p><span>{product.cost} ₽</span></div>
            <Link href={`/products/${product.category}/${product.guitar_id}`}><button>see more</button></Link>
        </div>
    )
  
}

export default Product