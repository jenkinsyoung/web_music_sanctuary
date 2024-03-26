import React from 'react'
import styles from './Header.module.scss'
import Link from 'next/link'
import Image from 'next/image'
import {getCurrentDateAndYear, getDayName, getMonthName} from '@/utils/Datetime'
import NavLink from './NavLink'

const Links=[
    {
      src: "/brush-1.svg",
      title: "products",
      path: "/products",
      w: 349,
      h: 49,
    },
    {
      src: "/brush-2.svg",
      title: "your advertisements",
      path: "/advertisements",
      w: 315,
      h: 55,
    },
    {
        src: "/brush-3.svg",
        title: "your profile",
        path: "/your_profile",
        w: 361,
        h: 41,
    }
  ];

const Header = () => {
    const date = getCurrentDateAndYear()
    const month = getMonthName()
    const day = getDayName()
  return (
    <div className={styles.container}>
        <div className={styles.header}>
            <div className={styles.logo}>
                <div className={styles.logo_img}>
                    <Link href='/'><Image 
                    src='/electric-guitar.svg'
                    alt=''
                    width={209}
                    height={209} /></Link>
                </div>
            </div>
            <div className={styles.title_wrapper}>
                <div className={styles.today}>
                    {day}, {month} {date.day}, {date.year}
                </div>
                <h1>music sanctuary</h1>
                <h3>The best music store</h3>
            </div>
            <div className={styles.creator}>
                <button>create an advertisement</button>
            </div>
            
        </div>
        <div className={styles.navbar}>
            <div className={styles.link}>
            <Link href='/'>about</Link>
            
            </div>
            <div className={styles.link}>
            {/* <Link href='/products'>products</Link> */}
            <NavLink item={Links[0]} />
            </div>
            <div className={styles.link}>
            <NavLink item={Links[1]} />
            </div>
            <div className={styles.link}>
            <NavLink item={Links[2]} />
            </div>
        </div>
    </div>
  )
}

export default Header