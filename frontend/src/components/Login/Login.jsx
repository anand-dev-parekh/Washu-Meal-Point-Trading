import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios'

import './login.css'


const Login = () => {
  const navigate = useNavigate();
  
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleLogin = (e) => {
    e.preventDefault();
    const headers = {
        'Content-Type': 'application/json',
    };

    const data = {
        email: email,
        password: password,
    };
    
    //get jwt token
    axios.post("http://localhost:8080/generate-token", data, {headers})
     .then((response) => { 
        localStorage.setItem("token", response.data.token) //maybe change later?
        localStorage.setItem("userID", response.data.userID) 
        navigate('/offers')
     })
     .catch((error) => {
        alert(error.response.data.error)
     })
  }

  return (
    <div className="login-container">
      <h1>Login</h1>
      <form onSubmit={handleLogin} className="loginReact">
        <div className="form-group loginReact">
          <input
            className="loginReact"
            type="email" 
            id="email-input" 
            value={email} 
            placeholder='email'
            onChange={(e) => setEmail(e.target.value)} 
          />
        </div>
        <div className="form-group loginReact">
          <input 
            className="loginReact"
            type="password" 
            id="password-input" 
            value={password} 
            placeholder='password'
            onChange={(e) => setPassword(e.target.value)} 
          />
        </div>
        <button type="submit" className="loginReact">Login</button>
      </form>
    </div>
  );
}

export default Login;
