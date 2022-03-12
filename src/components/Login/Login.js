import React, {useState} from 'react';
import { useHistory } from "react-router-dom";
import './Login.css';
import '../Problems/Problems';
import { BrowserRouter, Route, Link, Redirect } from 'react-router-dom';
import Cookies from 'js-cookie';
import PropTypes from 'prop-types';

async function loginUser(credentials) {
  console.log(JSON.stringify(credentials))
  return fetch('http://localhost:8080/login',{
    method: 'POST',
    headers:{
      'Content-Type':'application/json',
    },
    credentials: 'include',
    body: JSON.stringify(credentials)
  })
  .then(data => data.json())
}

async function CheckCookie(cookie){
  const [err, setError] = useState(false);
  fetch('http://localhost:8080/cookie',{
    method: 'POST',
    headers:{
      'Content-Type':'application/json'
    },
    body: JSON.stringify(cookie)
  })
  .then(data => data.json())
  .catch(error => setError(error.message))
}

export default function Login() {
  const history = useHistory();
  const [username, setUserName] = useState();
  const [password, setPassword] = useState();
  const handleSubmit = async e => {
    e.preventDefault();
    history.push('/problems')
    const token = await loginUser({
      username,
      password
    })
  }
  let cookie = Cookies.get("session_token");
  if (cookie){
    if (CheckCookie(cookie)){
      console.log(CheckCookie(cookie))
      //return <Redirect to='/problems'/>
    }
  }
  return(
    <div className="login">
      <form onSubmit={handleSubmit}>
        <label>
          <p>Username</p>
          <input type="text" onChange={e => setUserName(e.target.value)} />
        </label>
        <label>
          <p>Password</p>
          <input type="password" onChange={e => setPassword(e.target.value)}/>
        </label>
        <div>
          <button type="submit" className="Submit">Submit</button>
        </div>
      </form>
    </div>
  )
}
// Login.propTypes = {
//   setToken: PropTypes.func.isRequired
// }