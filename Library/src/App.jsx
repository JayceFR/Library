//Front end issues -> Cors problem while calling apis

import { useState, useEffect } from 'react'
//import './App.css'


//Sign Up Page
function App() {
  const [name, setName] = useState("")
  const [email , setEmail] = useState("")
  const [password , setPassword] = useState("")
  const [can_submit, setCan_submit] = useState(false)

  const email_regex = new RegExp(/^[a-z0-9A-Z]{2,}@..*\.com$/gm)
  const password_regex = new RegExp(/\w/gm)

  function submit(e){
    e.preventDefault()
    console.log(email)
    console.log(password)
    get_user_id()
  }

  /*useEffect( () => {
    const fetchdata = async () => {
      const result = await fetch("https://library.mysticjayce.repl.co/user_id")
      console.log(result)
    }
    fetchdata()
  }, [])*/

  const get_user_id = async () => {
    //const base_url = 'https://library.mysticjayce.repl.co/'
    // const base_url = 'http://127.0.0.1:5000/'
    // const result = await fetch(base_url + 'user_id')
    // const json_result = await result.json()
    // console.log(json_result['id'])
    // const put_url = base_url + "user/" + json_result['id']   
    // console.log(put_url)
    // const user_data = {
    //   "email" : email,
    //   "password" : password
    // }
    // const result2 = await fetch(put_url, {
    //   method: 'PUT', 
    //   headers: {
    //     'Content-Type' : 'application/json'
    //   },
    //   body: JSON.stringify(user_data)
    // })
    // const json_result2 = await result2.json()
    // console.log(json_result2)

    const base_url = "http://localhost:3000/account"
    const user_data = {
      "first_name":name,
      "email":email,
      "password":password
    }
    const result = await fetch(base_url, {
      method: 'POST',
      mode:'no-cors',
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
    // const json_result = await result.json()
    // console.log(json_result)

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
     <p>Login Page</p>
     <br></br>
     <form onSubmit={submit}>
      <label>Name</label>
      <br></br>
      <input value={name} onChange={e => {setName(e.target.value)}} type='text'/>
      <br></br>
      <label>Email</label>
      <br></br>
      <input value={email} onChange={e => {setEmail(e.target.value)}} type='text'/>
      <br></br>
      <label>Password</label>
      <br></br>
      <input value={password} onChange={e => {setPassword(e.target.value)}} type='password'/>
      <br></br>
      {
        can_submit && <input type='submit'/>
      }
     </form>
    </>
  )
}

export default App
