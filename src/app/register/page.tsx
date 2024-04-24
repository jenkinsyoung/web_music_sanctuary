"use client"
import { useState } from 'react';
import {useRouter} from 'next/navigation';

const LoginPage = () => {
  const [email, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const router = useRouter();
  const handleLogin = async () => {
    try {
      // Отправка запроса на сервер для аутентификации
      const response = await fetch('http://127.0.0.1/api/sing-in', {
        method: 'POST',
        body: JSON.stringify({ email, password }),
        headers: {
          'Content-Type': 'application/json'
        }
      });
      
      if (!response.ok) {
        throw new Error('Ошибка аутентификации');
      }

      const data = await response.json();

      // Сохранение токена в localStorage
      localStorage.setItem('token', data.access_token);

      // Перенаправление на страницу создания объявления
      router.push('/creation');
    } catch (error) {
      console.error('Ошибка входа:', error);
      // Обработка ошибки аутентификации
      // Можно вывести сообщение об ошибке пользователю
    }
  };

  return (
    <div>
      <input type="text" value={email} onChange={(e) => setUsername(e.target.value)} />
      <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
      <button onClick={handleLogin}>Login</button>
    </div>
  );
};

export default LoginPage;
