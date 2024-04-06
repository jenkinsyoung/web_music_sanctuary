import React from 'react'
import styles from './Filters.module.scss'
const NumOfStringsFilter = () => {
  return (
    <div className={styles.container}>
       <h5>Количество струн</h5>
        <div className={styles.window_number}>
          <div>
            <input type="checkbox" />
            <label>4</label>
          </div>
          <div>
            <input type="checkbox" />
            <label>5</label>
          </div>
          <div>
            <input type="checkbox" />
            <label>6</label>
          </div>
          <div>
            <input type="checkbox" />
            <label>7</label>
          </div>
          <div>
            <input type="checkbox" />
            <label>8</label>
          </div>
          <div>
            <input type="checkbox" />
            <label>12</label>
          </div>
        </div>
    </div>
  )
}

export default NumOfStringsFilter