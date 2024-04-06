import React from 'react'
import styles from './Filters.module.scss'

const Pickups = () => {
  return (
    <div className={styles.container}>
       <h5>Звукосниматели</h5>
        <div className={styles.window}>
          <div>
            <input type="checkbox" />
            <label>SSS</label>
          </div>
          <div>
            <input type="checkbox" />
            <label>SS</label>
          </div>
          <div>
            <input type="checkbox" />
            <label>HSH</label>
          </div>
          <div>
            <input type="checkbox" />
            <label>HH</label>
          </div>
          <div>
            <input type="checkbox" />
            <label>HSS</label>
          </div>
          <div>
            <input type="checkbox" />
            <label>H</label>
          </div>
        </div>
    </div>
  )
}

export default Pickups