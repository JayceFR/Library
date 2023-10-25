import { useState } from "react";
import Input from "../components/input";
import { email_regex_value, password_regex_value } from "../constants/regexConstants";

function Login() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [can_submit, setCan_submit] = useState(false)
    const email_regex = new RegExp(email_regex_value);
    const password_regex = new RegExp(password_regex_value);
    function submit(e) {
        e.preventDefault();
        console.log(email);
        console.log(password);
        login();
    }
    const login = async () => {
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
        const json_result = await result.json();
        console.log(json_result);
    }
    if (email_regex.test(email) && password_regex.test(password)) {
        if (!can_submit) {
            setCan_submit(true)
        }
    }
    else {
        if (can_submit) {
            setCan_submit(false)
        }
    }
    return (
        <>
            <p>Login Page</p>
            <br></br>
            <form onSubmit={submit}>
                <label>Email</label>
                <br></br>
                <Input password={false} value={email} change_method={setEmail} />
                <br></br>
                <label>Password</label>
                <br></br>
                <Input password={true} value={password} change_method={setPassword} />
                <br></br>
                <br></br>
                {
                    can_submit && <input type='submit' />
                }
            </form>
        </>
    )
}
export default Login;