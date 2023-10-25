import React from 'react'
import ReactDOM from 'react-dom/client'
import Register from './pages/register.jsx'
import {BrowserRouter} from "react-router-dom"
import { Route, Routes } from 'react-router-dom'
import Home from './pages/home.jsx'
import Nav from './components/nav.jsx'
import Login from './pages/login.jsx'
import "./style.css"

function App(){
  return(
    <>
    <Nav/>
    <Routes>
      <Route path = "/" element={<Home/>}/>
      <Route path = "/register" element={<Register/>}/>
      <Route path = "/login" element={<Login/>}/>
    </Routes>
    </>
  )
}

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <BrowserRouter>
      <App />
    </BrowserRouter>
  </React.StrictMode>,
)
