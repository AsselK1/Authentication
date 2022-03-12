import React, {useState} from 'react';
import './Problems.css';
import TwoSum from './TwoSum';
import { BrowserRouter, Route, Link, Redirect } from 'react-router-dom';
import AddTwoNumbers from './AddTwoNumbers';
import Longest from './Longest Substring';
function handleErrors(response) {
  if (!response.ok) {
      throw Error(response.statusText);
  }
  return response;
}
export default function Problems(){
  const [err, setError] = useState(null);
  fetch('http://localhost:8080/problems',{
    method: 'GET',
    headers:{
      'Content-Type':'application/json'
    },
    credentials: 'include'
  })
  .then(handleErrors)
    .then(response => console.log("ok") )
    .catch(error => setError(error.message) );
    console.log(err)
  if (err){
    return <Redirect to='/login'/>
  }
  return(
    <div>
    <div className="wrapper">
      <button className="logout">Log out</button>
      <h2>Problems</h2>
      <div className = "list">
      <BrowserRouter>
      <ol>
        <li className = "problem"><Link to="/problems/1">Two Sum</Link></li>
        <li className = "problem"><Link to="/problems/2">Add two numbers</Link></li>
        <li className = "problem"><Link to="/problems/3">Longest Substring</Link></li>
        <li className = "problem"><Link to="/problems/1">Two Sum</Link></li>
        <li className = "problem"><Link to="/problems/2">Add two numbers</Link></li>
        <li className = "problem"><Link to="/problems/3">Longest Substring</Link></li>
        <li className = "problem"><Link to="/problems/1">Two Sum</Link></li>
        <li className = "problem"><Link to="/problems/2">Add two numbers</Link></li>
        <li className = "problem"><Link to="/problems/3">Longest Substring</Link></li>
        <li className = "problem"><Link to="/problems/1">Two Sum</Link></li>
        <li className = "problem"><Link to="/problems/2">Add two numbers</Link></li>
        <li className = "problem"><Link to="/problems/3">Longest Substring</Link></li>
        <li className = "problem"><Link to="/problems/1">Two Sum</Link></li>
        <li className = "problem"><Link to="/problems/2">Add two numbers</Link></li>
        <li className = "problem"><Link to="/problems/3">Longest Substring</Link></li>
        <li className = "problem"><Link to="/problems/1">Two Sum</Link></li>
        <li className = "problem"><Link to="/problems/2">Add two numbers</Link></li>
        <li className = "problem"><Link to="/problems/3">Longest Substring</Link></li>
        <li className = "problem"><Link to="/problems/1">Two Sum</Link></li>
        <li className = "problem"><Link to="/problems/2">Add two numbers</Link></li>
        <li className = "problem"><Link to="/problems/3">Longest Substring</Link></li>
        <li className = "problem"><Link to="/problems/1">Two Sum</Link></li>
        <li className = "problem"><Link to="/problems/2">Add two numbers</Link></li>
        <li className = "problem"><Link to="/problems/3">Longest Substring</Link></li>
        <li className = "problem"><Link to="/problems/1">Two Sum</Link></li>
        <li className = "problem"><Link to="/problems/2">Add two numbers</Link></li>
        <li className = "problem"><Link to="/problems/3">Longest Substring</Link></li>
        <li className = "problem"><Link to="/problems/1">Two Sum</Link></li>
        <li className = "problem"><Link to="/problems/2">Add two numbers</Link></li>
        <li className = "problem"><Link to="/problems/3">Longest Substring</Link></li>
        </ol>
        <Route path='/problems/1'><TwoSum /></Route>
        <Route path='/problems/2'><AddTwoNumbers /></Route>
        <Route path='/problems/3'><Longest /></Route>
      </BrowserRouter>
      </div>
    </div>
    </div>
  );
}