import  { createContext, useContext, useEffect, useState } from "react";
import { null_uuid } from "../constants/uuidConstants";
import App from "../main";
import Nav from "../components/nav";
import { useNavigate } from "react-router-dom";

export const AuthContext = createContext();

export const AuthData = () => useContext(AuthContext);

function Authenticaiton(){
    const [user, setUser] = useState({id:null, first_name : null, email : null, password : null, is_logged_in :false});
    const [mode, setMode] = useState("dark");

    function change_dark(){
        var r = document.querySelector(':root');
        if (mode == "dark"){
            r.style.setProperty('--bgcolor', '#E2E2E2')
            r.style.setProperty('--txtcolor', '#1B1B1B')
            r.style.setProperty('--empcolor', 'black')
            setMode("light");
        }
        else{
            r.style.setProperty('--txtcolor', '#E2E2E2')
            r.style.setProperty('--bgcolor', '#1B1B1B')
            r.style.setProperty('--empcolor', 'white')
            setMode("dark");
        }
    }

    const navigate = useNavigate();

    const login = async (email, password) => {
        const url = "http://localhost:3000/login";
        const user_data = {
            "email": email,
            "password": password
        }
        const result = await fetch(url, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json'
            },
            body: JSON.stringify(user_data)
        });
        const user_obj = await result.json();
        return new Promise((resolve, reject) => {
            if (user_obj['id'] == null_uuid){
                console.log("Failed in logging in ")
                setUser({...user_data, is_logged_in: false})
                reject("Incorrect email or password")
            }
            else{
                console.log("Successfully logged in ")
                setUser({...user_data, is_logged_in: true})
                resolve("Success")
            }
        })
        
    }

    const logout = () => {
        setUser({...user_data, is_logged_in: false})
        useEffect(()=>{
            navigate("/")
        }, [])
    }

    return (
        <AuthContext.Provider value={{user, login, logout, mode, change_dark}}>
            <>
                <Nav/>
                <App/>
            </>
        </AuthContext.Provider>
    )

}

export default Authenticaiton;