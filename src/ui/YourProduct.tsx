import styles from './YourProduct.module.scss'
import Link from 'next/link'
const YourProduct = () => {

  
    return (
        <div className={styles.container}>
            <h1>Гитара 2239402-21ш349</h1>
            <div className={styles.img} style={{backgroundImage: `url(/add_guitar.png)`}} />
            <div className={styles.tags}>#pppp #ppppp #pppp</div>
            <div className={styles.cost}><p>Цена:</p><span>10000 ₽</span></div>
            <div style={{display: 'flex'}}>
            <Link href={`/advertisements/id`}><button className={styles.edit}>edit</button></Link>
            <button className={styles.delete}>delete</button>
            </div>
            
        </div>
    )
  
}

export default YourProduct