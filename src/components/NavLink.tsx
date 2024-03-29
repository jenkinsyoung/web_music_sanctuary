"use client";

import Link from "next/link";
import styles from "./Header.module.scss";
import { usePathname } from "next/navigation";
import Image from 'next/image';

const NavLink = ({item}: any) => {
    const pathName = usePathname();
    const path = item.path;
    return (
    <>
    <Link
      href={path} 
    >
      <p className={`${styles.p} ${
       (!pathName.search(path)) && styles.p_active
    }`} style={{zIndex: '100', position: 'relative'}} >{item.title}</p>
    </Link>
    <div className={`${styles.link_img} ${
      (!pathName.search(path)) && styles.active
    }`} >
        <Link
      href={path}
    >

<Image 
    src={item.src}
    width={item.w}
    height={item.h}
    alt=''/>
    </Link>
    </div>
    
    </>
    );
  };

export default NavLink;