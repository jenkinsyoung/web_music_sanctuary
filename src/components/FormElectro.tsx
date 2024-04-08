import React from 'react'
import styles from './Filters.module.scss'
const FormElectro = ({currentForm, onChangeCategory}: any) => {
  return (
    <div className={styles.container}>
       <h5>Форма гитары</h5>
        <div className={styles.window}>
          <div>
            <input type="checkbox" checked={currentForm.includes('Stratocaster')} onChange={(e) => onChangeCategory('Stratocaster', e.target.checked)}/>
            <label>Stratocaster</label>
          </div>
          <div>
            <input type="checkbox" checked={currentForm.includes('Telecaster')} onChange={(e) => onChangeCategory('Telecaster', e.target.checked)}/>
            <label>Telecaster</label>            
          </div>
          <div>
            <input type="checkbox" checked={currentForm.includes('Les paul')} onChange={(e) => onChangeCategory('Les paul', e.target.checked)}/>
            <label>Les paul</label>           
          </div>
          <div>
            <input type="checkbox" checked={currentForm.includes('Explorer')} onChange={(e) => onChangeCategory('Explorer', e.target.checked)}/>
            <label>Explorer</label>           
          </div>
          <div>
            <input type="checkbox" checked={currentForm.includes('Flying V')} onChange={(e) => onChangeCategory('Flying V', e.target.checked)}/>
            <label>Flying V</label>            
          </div>
          <div>
            <input type="checkbox" checked={currentForm.includes('SG')} onChange={(e) => onChangeCategory('SG', e.target.checked)}/>
            <label>SG</label>            
          </div>
          <div>
            <input type="checkbox" checked={currentForm.includes('ES')} onChange={(e) => onChangeCategory('ES', e.target.checked)}/>
            <label>ES</label>            
          </div>
        </div>
    </div>
  )
}

export default FormElectro