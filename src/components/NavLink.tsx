"use client";

import Link from "next/link";
import styles from "./Header.module.scss";
import { usePathname } from "next/navigation";
import Image from 'next/image';
const NavLink = ({ item }: any) => {
    const pathName = usePathname();
  
    return (
    <>
    <Link
      href={item.path}
    >
      {item.title}
    </Link>
    <div className={`${styles.link_img} ${
      pathName === item.path && styles.active
    }`}>
        <Link
      href={item.path}
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