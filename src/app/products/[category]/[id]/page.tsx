import styles from './ProductByIDPage.module.scss'

const ProductsByIDPage = ({params} : any) =>{
    return (
        <div className={styles.container}>
            <div className={styles.info}>
                <div className={styles.main}>
                    <h1>Guitar</h1>
                    <div className={styles.img} />
                </div>
            </div>
        </div>
    )
}

export default ProductsByIDPage

