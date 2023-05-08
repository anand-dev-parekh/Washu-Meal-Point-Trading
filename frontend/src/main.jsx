import React from 'react'
import ReactDOM from 'react-dom/client'
import { createBrowserRouter, RouterProvider } from "react-router-dom";

import { Auth } from './components/Auth'
import { Register } from './components/Register'
import { Login } from './components/Login'
import { Offers } from './components/Offers'
import { EmailVerify } from './components/EmailVerify'
import { Reports } from './components/Reports'

const router = createBrowserRouter([
  {
    path: "/",
    element: <Auth/>,
  },
  {
    path: "/login",
    element: <Login/>,
  },
  {
    path: "/register",
    element: <Register/>,
  },
  {
    path: "/offers",
    element: <Offers/>,
  },
  {
    path: "/verify-email",
    element: <EmailVerify/>,
  },
  {
    path: "/reports",
    element: <Reports/>,
  }
]);

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)
