"use client"
import React, { useState } from 'react';
import axios from 'axios';

const App = () => {
  const [image, setImage] = useState('');
  const [loading, setLoading] = useState(false);

  const handleUpload = async (e) => {
    setImage(e.target.files[0])
    // e.preventDefault();

    const file = e.target.files[0];
    const reader = new FileReader();

    reader.onload = async () => {
      const base64Image = reader.result.split(',')[1];

      const data = {
        image: base64Image,
      };

      try {
        const response = await axios.post('https://localhost:3000/api/upload-image', data);
        console.log(response.data);
        console.log(data);
      } catch (error) {
        console.error(error);
      } finally {
      }
    };

    reader.readAsDataURL(file);
  };

  return (
    <div>
      <h1>Загрузка изображения</h1>
      <form onSubmit={handleUpload}>
        <input type="file" name="image" />
        <button type="submit" >
           Загрузить
        </button>
      </form>
    </div>
  );
};

export default App;
