import Link from 'next/link'
import styles from './Footer.module.scss'

const Footer = () => {

    return (
      <div className={styles.container}>
        <div className={styles.letter_c}>
        ©2024 Association of sensible programmers
        </div>
        <div className={styles.contacts}>
            <div className={styles.phone}>
                <h2>phone</h2>
                <p>8 999 999 99 99</p>
            </div>
            <div className={styles.email}>
                <h2>email</h2>
                <Link href='#'>click here</Link>
            </div>
        </div>
        <div className={styles.empty} />
      </div>
    )

}

export default Footer