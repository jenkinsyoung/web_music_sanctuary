import React from 'react'
import styles from './Filters.module.scss'
const FormBass = () => {
  return (
    <div className={styles.container}>
       <h5>Форма гитары</h5>
        <div className={styles.window}>
            <input type="checkbox" />
            <label>Precision</label>
            <input type="checkbox" />
            <label>Jazz</label>
            <input type="checkbox" />
            <label>Fusion</label>
        </div>
    </div>
  )
}

export default FormBass