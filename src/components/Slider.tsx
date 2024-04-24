'use client'
import { useState } from 'react';
import styles from './Slider.module.scss'

// const images=[
//     {
//         'id': 1,
//         'url': '/add_guitar.png'
//     },
//     {
//         'id': 2,
//         'url': '/add_guitar2.jpg'
//     },
//     {
//         'id': 3,
//         'url': '/add_guitar3.jpg'
//     }
// ]

const Slider = ({product} : any) =>{
    const image = product.img_list
    console.log(image)
    const [currentIndex, setCurrentIndex]= useState(0)

    const slideStyles ={
        width: '774px',
        height: '615px',
        borderRadius: '10px',
        backgroundPosition: 'center',
        backgroundSize:'cover',
        backgroundImage: `url(data:image/png;base64,${image[currentIndex].image})`,
        margin: 'auto',
        transition: '0.5s'
    }
    const goToSlide = (slideIndex : number) => {
        setCurrentIndex(slideIndex);
    };

    const dotsContainerStyles = {
        display: "flex",
        gap: '10px',
        justifyContent: "center",
        transition: '0.5s',
    };
    
    const dotStyle = {
        margin: "20px 5px",
        cursor: "pointer",
        fontSize: '16px',
        fontWeight: '700',
        
    };
    return(
        <div className={styles.container}>
            <div style={{height: '100%', position: 'relative'}}>
            <div style={slideStyles}/>
            <div style={dotsContainerStyles}>
        {image.map((slide : any, slideIndex : any) => (
          <div
          className={`${styles.index} ${
            (slideIndex == currentIndex) && styles.active
         }`}
            style={dotStyle}
            key={slideIndex}
            onClick={() => goToSlide(slideIndex)}
          >
            {slideIndex + 1}
          </div>
        ))}
      </div>
            </div>    
        </div>
    )
}

export default Slider