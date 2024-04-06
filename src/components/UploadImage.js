"use client"
import React, {useState} from 'react'
import axios from 'axios'

const UploadImage = () => {
  const [image, setImage] = useState()

  function convertToBase64(e){
    console.log(e);
    let reader = new FileReader();
    reader.readAsDataURL(e.target.files[0]);
    reader.onload = () =>{
      console.log(reader.result);
      setImage(reader.result)
    };
    reader.onerror = error =>{
      console.log("Error: ", error);
    };
  }

  function uploadImage(){
    const data = JSON.stringify({
      base64: image
    })
    console.log(data)
    axios.post(``, data).then((res) => res).then((data)=>console.log(data))
  }
  return (
    <div>
      <input
      accept="image/*"
      type="file"
      onChange={convertToBase64}
      />
      {image=="" || image==null? '' : <img width={100} height = {100} src={image}/>}
      <button onClick={uploadImage}>Upload Image</button>
    </div>
  )
}

export default UploadImage