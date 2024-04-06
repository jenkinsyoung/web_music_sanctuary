"use client"
import styles from './HopePage.module.scss';
import Image from 'next/image';
import Link from 'next/link';
const HomePage = () => {
  return (
    <div className={styles.container}>
      <div className={styles.first_line}>
        <div className={styles.main_add}>
          <h1>Sell unnecessary music equipment</h1>
          <h4>music sanctuary</h4>
          <div className={styles.desc}>
            <p>It has now become easier to sell music equipment. You only need to take photos, and we will make an announcement for you, select the optimal price, and offer it to other customers.</p>
            <p>All you have to do is specify your contacts and wait for the message.</p>
          </div>
          <button><Link href ='/creation'>create an advertisement</Link></button>
          <div className={styles.image}><div className={styles.img} style={{backgroundImage: 'url(/music_equipment.svg)'}}/></div>
        </div>
        <div className={styles.hotline}>
          <div className={styles.hotline_card}>
            <div className={styles.image}>
              <Image 
              src='/phone.svg'
              alt=''
              width={266.97}
              height={189.83}/>
            </div>
            <h2>The Hotline</h2>
            <div className={styles.contacts}>
              <h3>phone</h3>
              <p>8 999 999 99 99</p>
              <h3>email</h3>
              <button>Click here</button>
            </div>
          </div>
        </div>
      </div>

      <div className={styles.second_line}>
        <div className={styles.story_add}>
          <h1>music sanctuary STORY</h1>
          <Image priority
          src='/guitar.svg'
          alt=''
          width={427.1}
          height={390.88}/>
          <p>Music Sanctuary was developed with intent to help people sell and exchange their music gear faster. Our goal is to harvest the marvel of modern AI technologies to provide the best experience possible.</p>
        </div>
        <div className={styles.advantages_add}>
          <h1>Why you should choose us</h1>
          <h4>Advantages</h4>
          <li>Quick listing generation based on the latest AI models</li>
          <li>95% accuracy guitar model recognition</li>
          <li>Trusted sellers and customers</li>
          <li>United community from all over the country</li>
          <li>No additional fees</li>
          <li>Free use for everyone</li>
          <Image priority
            src='/music_girl.svg'
            alt=''
            width={559}
            height={503}
          />
        </div>
        <div className={styles.ai_add}>
        <h1>ai & music</h1>
        <Image priority
          src='/ai-music.svg'
          alt=''
          width={475.28}
          height={393.71}/>
        <p>Use the latest advancements in the field of artificial intelligence to create the best gear shopping experience</p>
        </div>
        <h1 onClick={() => {window.scrollTo({
        top: 0,
        left: 0,
        behavior: 'smooth'
      });}}>Music Sanctuary</h1>
      </div>
      
    </div>
  );
}
export default HomePage