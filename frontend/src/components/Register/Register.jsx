import React, { useState } from 'react';
import axios from 'axios'
import './register.css'
import { useNavigate } from 'react-router';

const Register = () => {
  const navigate = useNavigate()
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');

  //called on button click
  const handleRegister = (e) => {
    e.preventDefault();
    const headers = {
        'Content-Type': 'application/json',
    };

    if (password != confirmPassword) {
      alert("Passowrds are not matching")
      return
    }

    const data = {
        id: 1,
        email: email,
        password: password,
        authLevel: 1,
    };

    //make post request to create user
    axios.post(process.env.API_BASE_URL + "/create-user", data, {headers})
    .then((response) => { 
      navigate('/login')
   })
   .catch((error) => {
      alert(error.response.data.error)
   })
  }

  return (
    <div className="register-container">
      <h1>Register</h1>
      <form onSubmit={handleRegister} className="registerReact">
        <div className="form-group registerReact">
          <input
            className="registerReact"
            type="email" 
            id="email-input" 
            value={email} 
            placeholder='email'
            onChange={(e) => setEmail(e.target.value)} 
          />
        </div>
        <div className="form-group registerReact">
          <input 
            className="registerReact"
            type="password" 
            id="password-input" 
            value={password} 
            placeholder='password'
            onChange={(e) => setPassword(e.target.value)} 
          />
        </div>
        <div className="form-group registerReact">
          <input 
            className="registerReact"
            type="password" 
            id="confirm-password-input"
            placeholder='confirm password'
            value={confirmPassword} 
            onChange={(e) => setConfirmPassword(e.target.value)} 
          />
        </div>
        <button type="submit" className="registerReact">Register</button>
      </form>
    </div>
  );
}

export default Register;
