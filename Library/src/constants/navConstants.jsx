import Dashboard from "../pages/dashboard";
import Home from "../pages/home";
import Login from "../pages/login";
import Register from "../pages/register";

export const nav = [
    {path: "/", name:"Home", element: <Home/>, isMenu: true, isPrivate: false},
    {path: "/login", name:"Login", element: <Login/>, isMenu: false, isPrivate: false},
    {path: "/register", name:"Register", element: <Register/>, isMenu: true, isPrivate: false},
    {path: "/dashboard", name:"Dashboard", element: <Dashboard/>, isMenu: true, isPrivate: true},
]