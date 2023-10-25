import { NavLink } from "react-router-dom";
function Nav(){
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
        </>
    )
}

export default Nav;