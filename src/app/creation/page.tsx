import React from 'react'
import styles from './CreatePage.module.scss'
import UploadImage from '@/components/UploadImage'
const CreatePage = () => {
  return (
   
        <div className={styles.container}>
            <div className={styles.instruction}>На данной странице Вам необходимо загрузить фотографии, ввести название и придумать описание своего товара. Позже Вы сможете отредактировать данные в своих объявлениях.</div>
            <form className={styles.info}>
                <div className={styles.main}>
                    <input className={styles.h1} placeholder={`Введите название`}/>
                    <UploadImage />
                </div>
                <div className={styles.desc}>
                    <h2>Описание:</h2>
                    <div><textarea placeholder={`Введите описание`} /></div>

                    <button type='submit'>Создать обьявление</button>
                </div>
            </form>
        </div>
    )
}

export default CreatePage;