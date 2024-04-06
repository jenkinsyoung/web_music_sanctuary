import React from 'react'
import styles from './Filters.module.scss'
const FormElectro = () => {
  return (
    <div className={styles.container}>
       <h5>Форма гитары</h5>
        <div className={styles.window}>
          <div>
            <input type="checkbox" />
            <label>Stratocaster</label>
          </div>
          <div>
            <input type="checkbox" />
            <label>Telecaster</label>            
          </div>
          <div>
            <input type="checkbox" />
            <label>Les paul</label>           
          </div>
          <div>
            <input type="checkbox" />
            <label>Explorer</label>           
          </div>
          <div>
            <input type="checkbox" />
            <label>Flying V</label>            
          </div>
          <div>
            <input type="checkbox" />
            <label>SG</label>            
          </div>
          <div>
            <input type="checkbox" />
            <label>ES</label>            
          </div>
        </div>
    </div>
  )
}

export default FormElectro