import React from 'react'
import styles from './Filters.module.scss'
const UkuleleType = () => {
  return (
    <div className={styles.container}>
       <h5>Вид гитары</h5>
        <div className={styles.window}>
            <input type="checkbox" />
            <label>Акустическое</label>
            <input type="checkbox" />
            <label>Электрическое</label>
        </div>
    </div>
  )
}

export default UkuleleType