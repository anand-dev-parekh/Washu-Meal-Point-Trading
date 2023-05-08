import React from 'react';
import { Link } from 'react-router-dom'
import './auth.css'

const Auth = () => {
  return (
    <div className="authReact" id="authContainer">
      <h1>Welcome to Washu Meal Points Trader</h1>
      <Link to="/login">
        <button className="login authReact">Login</button>
      </Link>
      <Link to="/register">
        <button className="register authReact">Register</button>
      </Link>
    </div>
  );
};

export default Auth;
