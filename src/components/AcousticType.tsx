import React from 'react'
import styles from './Filters.module.scss'
const AcousticType = () => {
  return (
    <div className={styles.container}>
       <h5>Вид гитары</h5>
        <div className={styles.window}>
            <input type="checkbox" />
            <label>Электроакустика</label>
            <input type="checkbox" />
            <label>Не электроакустика</label>
        </div>
    </div>
  )
}

export default AcousticType