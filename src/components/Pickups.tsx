import React from 'react'
import styles from './Filters.module.scss'

const Pickups = ({currentForm, onChangeCategory}:any) => {
  return (
    <div className={styles.container}>
       <h5>Звукосниматели</h5>
        <div className={styles.window}>
          <div>
            <input type="checkbox" checked={currentForm.includes('SSS')} onChange={(e) => onChangeCategory('SSS', e.target.checked)}/>
            <label>SSS</label>
          </div>
          <div>
            <input type="checkbox" checked={currentForm.includes('SS')} onChange={(e) => onChangeCategory('SS', e.target.checked)}/>
            <label>SS</label>
          </div>
          <div>
            <input type="checkbox" checked={currentForm.includes('HSH')} onChange={(e) => onChangeCategory('HSH', e.target.checked)}/>
            <label>HSH</label>
          </div>
          <div>
            <input type="checkbox" checked={currentForm.includes('HH')} onChange={(e) => onChangeCategory('HH', e.target.checked)}/>
            <label>HH</label>
          </div>
          <div>
            <input type="checkbox" checked={currentForm.includes('HSS')} onChange={(e) => onChangeCategory('HSS', e.target.checked)}/>
            <label>HSS</label>
          </div>
          <div>
            <input type="checkbox" checked={currentForm.includes('H')} onChange={(e) => onChangeCategory('H', e.target.checked)}/>
            <label>H</label>
          </div>
        </div>
    </div>
  )
}

export default Pickups