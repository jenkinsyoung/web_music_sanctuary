import React from 'react'
import styles from './Filters.module.scss'
const FormAcoustic = () => {
  return (
    <div className={styles.container}>
       <h5>Форма гитары</h5>
        <div className={styles.window}>
            <input type="checkbox" />
            <label>Parlor</label>
            <input type="checkbox" />
            <label>Dreadnought</label>
            <input type="checkbox" />
            <label>Concert</label>
            <input type="checkbox" />
            <label>Triple-o</label>
        </div>
    </div>
  )
}

export default FormAcoustic