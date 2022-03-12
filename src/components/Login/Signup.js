import React, {useState} from 'react';
import './Login.css';

async function signUser(credentials) {
    console.log(JSON.stringify(credentials))
    return fetch('http://localhost:8080/signup',{
      method: 'POST',
      headers:{
        'Content-Type':'application/json'
      },
      body: JSON.stringify(credentials)
    })
    .then(data => data.json())
  }

export default function Signup(){
    const [username, setUserName] = useState();
    const [password, setPassword] = useState();
    const [firstname, setFirstName] = useState();
    const [lastname, setLastName] = useState();
    const [email, setEmail] = useState();
    const handleSubmit = async e => {
        e.preventDefault();
        const token = await signUser({
          username,
          password,
          firstname,
          lastname,
          email
        });
        console.log(token)
      }
    return(
        <div className="login">
            <form onSubmit = {handleSubmit}>
            <label>
            <p>Username</p>
            </label>
            <input type="text" required onChange={e => setUserName(e.target.value)}/>
            <label>
            <p>Email</p>
            </label>
            <input type="text" required onChange={e => setEmail(e.target.value)}/>
            <label>
            <p>First Name</p>
            </label>
            <input type="text" required onChange={e => setFirstName(e.target.value)}/>
            <label>
            <p>Last Name</p>
            </label>
            <input type="text" onChange={e => setLastName(e.target.value)}/>
            <label>
            <p>Password</p>
            </label>
            <input type="password" required onChange={e => setPassword(e.target.value)}/>
            <br />
            <button type="submit" className="Submit">Submit</button>
            </form>
        </div>
    )
}