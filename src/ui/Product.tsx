import styles from './Product.module.scss'

const Product = () => {

  
    return (
        <div className={styles.container}>
            <h1>Гитара 2239402-21ш349</h1>
            <div className={styles.img} style={{backgroundImage: `url(/add_guitar.png)`}} />
            <p>Lorem ipsum, dolor sit amet consectetur adipisicing elit. Nam maxime dignissimos optio non est? Vel, quaerat iusto?</p>
            <button>see more</button>
        </div>
    )
  
}

export default Product