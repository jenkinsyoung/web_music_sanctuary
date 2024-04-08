import React from 'react'
import styles from './Filters.module.scss'
const FormAcoustic = ({currentForm, onChangeCategory}:any) => {
  return (
    <div className={styles.container}>
       <h5>Форма гитары</h5>
        <div className={styles.window}>
          <div>
            <input type="checkbox" checked={currentForm.includes('Parlor')} onChange={(e) => onChangeCategory('Parlor', e.target.checked)}/>
            <label>Parlor</label>
          </div>
          <div>
            <input type="checkbox" checked={currentForm.includes('Dreadnought')} onChange={(e) => onChangeCategory('Dreadnought', e.target.checked)}/>
            <label>Dreadnought</label>
          </div>
          <div>
            <input type="checkbox" checked={currentForm.includes('Concert')} onChange={(e) => onChangeCategory('Concert', e.target.checked)}/>
            <label>Concert</label>
          </div>
          <div>
            <input type="checkbox" checked={currentForm.includes('Triple-O')} onChange={(e) => onChangeCategory('Triple-O', e.target.checked)}/>
            <label>Triple-O</label>
          </div>
        </div>
    </div>
  )
}

export default FormAcoustic