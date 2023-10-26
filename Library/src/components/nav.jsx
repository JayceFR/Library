import { NavLink } from "react-router-dom";
import { useState } from "react";
function Nav(){
    const [mode, setMode] = useState("dark");
    function changeDark(){
        var r = document.querySelector(':root');
        if (mode == "dark"){
            r.style.setProperty('--bgcolor', '#E2E2E2')
            r.style.setProperty('--txtcolor', '#1B1B1B')
            setMode("light");
        }
        else{
            r.style.setProperty('--txtcolor', '#E2E2E2')
            r.style.setProperty('--bgcolor', '#1B1B1B')
            setMode("dark");
        }
    }
    return(
        <>
            <nav>
                <ul className="navbar">
                    <li>
                        <NavLink to="/register">Register</NavLink>
                    </li>
                    <li>
                        <NavLink to="/login">Login</NavLink>
                    </li>
                    <li>
                        <NavLink to="/">Home</NavLink>
                    </li>
                </ul>
            </nav>
            <br></br>
            {mode == "dark" && <img id="light_dark_mode" src = "../../Assets/light_on.png"/>}
            {mode == "light" && <img id="light_dark_mode" src = "../../Assets/light_off.png"/>}
            {mode=="dark" && <a onClick={changeDark} id="light_dark_txt">Turn on the lights</a>}
            {mode=="light" && <a onClick={changeDark} id="light_dark_txt">Turn off the lights</a>}
        </>
    )
}

export default Nav;