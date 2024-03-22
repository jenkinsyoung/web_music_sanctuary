import React from 'react'
import styles from './Header.module.scss'
import Link from 'next/link'
import Image from 'next/image'

const Header = () => {
  return (
    <div className={styles.container}>
        <div className={styles.header}>
            <div className={styles.logo}>
                <div className={styles.logo_img}>
                    <Image 
                    src='/electric-guitar.svg'
                    alt=''
                    width={209}
                    height={209} />
                </div>
            </div>
            <div className={styles.title_wrapper}>
                <div className={styles.today}>
                    {}
                </div>
                <h1>music sanctuary</h1>
                <h3>The best music store</h3>
            </div>
            <div className={styles.creator}>
                <button>create an advertisement</button>
            </div>
            
        </div>
        <div className={styles.navbar}>
            <Link href='#'>about</Link>
            <Link href='#'>products</Link>
            <Link href='#'>your advertisement</Link>
            <Link href='#'>your profile</Link>
        </div>
    </div>
  )
}

export default Header