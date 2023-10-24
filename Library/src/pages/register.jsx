//Front end issues -> Cors problem while calling apis

import { useState, useEffect } from 'react'
import Input from '../components/input'
import { email_regex_value, password_regex_value } from '../constants/regexConstants'
//import './App.css'


//Sign Up Page
function Register() {
  const [name, setName] = useState("")
  const [email , setEmail] = useState("")
  const [password , setPassword] = useState("")
  const [can_submit, setCan_submit] = useState(false)

  function submit(e){
    e.preventDefault()
    console.log(email)
    console.log(password)
    sign_up()
  }

  const email_regex = new RegExp(email_regex_value);
  const password_regex = new RegExp(password_regex_value);

  const sign_up = async () => {

    const base_url = "http://localhost:3000/account"
    const user_data = {
      "first_name":name,
      "email":email,
      "password":password
    }
    const result = await fetch(base_url, {
      method: 'POST',
      headers:{
        'Content-Type' : 'application/json',
        'Accept':'application/json'
      },
      body: JSON.stringify(user_data)
    })    

    const result2 = await fetch(base_url, {
      headers:{
        'Accept':'application/json'
      }
    })
    const json_result = await result2.json()
    console.log(json_result)

  }

  if(email_regex.test(email) && password_regex.test(password)){
    if(!can_submit){
      setCan_submit(true)
    }
  }
  else{
    if (can_submit){
      setCan_submit(false)
    }
  }

  console.log(can_submit)

  return (
    <>
     <p>Sign Up Page</p>
     <br></br>
     <form onSubmit={submit}>
      <label>Name</label>
      <br></br>
      <Input password={false} value={name} change_method={setName}/>
      <br></br>
      <label>Email</label>
      <br></br>
      <Input password={false} value={email} change_method={setEmail}/>
      <br></br>
      <label>Password</label>
      <br></br>
      <Input password={true} value={password} change_method={setPassword}/>
      <br></br>
      {
        can_submit && <input type='submit'/>
      }
     </form>
    </>
  )
}

export default Register
