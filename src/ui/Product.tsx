import styles from './Product.module.scss'
import Link from 'next/link'
const Product = () => {

  
    return (
        <div className={styles.container}>
            <h1>Гитара 2239402-21ш349</h1>
            <div className={styles.img} style={{backgroundImage: `url(/add_guitar.png)`}} />
            <div className={styles.tags}>#pppp #ppppp #pppp</div>
            <div className={styles.cost}><p>Цена:</p><span>10000 ₽</span></div>
            <Link href={`/products/category/id`}><button>see more</button></Link>
        </div>
    )
  
}

export default Product