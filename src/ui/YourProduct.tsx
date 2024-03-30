import styles from './YourProduct.module.scss'
import Link from 'next/link'
const YourProduct = () => {

  
    return (
        <div className={styles.container}>
            <h1>Гитара 2239402-21ш349</h1>
            <div className={styles.img} style={{backgroundImage: `url(/add_guitar.png)`}} />
            <p>Lorem ipsum, dolor sit amet consectetur adipisicing elit. Nam maxime dignissimos optio non est? Vel, quaerat iusto?</p>
            <div style={{display: 'flex'}}>
            <Link href={`/advertisements/id`}><button className={styles.edit}>edit</button></Link>
            <button className={styles.delete}>delete</button>
            </div>
            
        </div>
    )
  
}

export default YourProduct