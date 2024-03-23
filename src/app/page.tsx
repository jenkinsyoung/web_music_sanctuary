import styles from './HopePage.module.scss';

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
          <button>create an advertisement</button>
          <div className={styles.image}><div className={styles.img} style={{backgroundImage: 'url(/music_equipment.svg)'}}/></div>
        </div>
        <div className={styles.hotline}>
          dd
        </div>
      </div>
    </div>
  );
}
export default HomePage