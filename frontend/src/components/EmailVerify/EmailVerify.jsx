import React from 'react';
import axios from 'axios'
import './emailVerify.css'
import { useNavigate } from 'react-router';

const EmailVerify = () => {
  const navigate = useNavigate()

  const handleSendEmail = () => {
    const headers = {
        'Content-Type': 'application/json',
        'Authorization': localStorage.getItem("token"),
    };

    //send confirmation email thru put request
    axios.put(process.env.API_BASE_URL + "/send-confirmation-email", {}, {headers})
      .then((response) => {
        alert("Sent verification email")
      })
      .catch((error) => {
        console.log(error)
        alert(error.response.data.error)
      })
  };

  return (
    <div className="email-verify-container">
      <h1>Email verification</h1>
      <p>Please click the button below to send your verification email.</p>
      <button className="email-verify-button" onClick={handleSendEmail}>Send verification email</button>
      <button className="email-verify-button" onClick={() => navigate("/login")}> Back to Login</button>
    </div>
  );
};

export default EmailVerify;
