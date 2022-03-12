import React,{useState} from 'react';
import './Problems.css';
async function SubmitCode(credentials){
    console.log(JSON.stringify(credentials))
    return fetch('http://localhost:8080/problems/2',{
    method: 'POST',
    headers:{
      'Content-Type':'application/json'
    },
    body: JSON.stringify(credentials)
  })
  .then(data => data.json())
}
export default function AddTwoNumbers(){
    const [code, setCode] = useState();
    const [lang, setLang] = useState();
    const handleSubmit = async e => {
        e.preventDefault();
        const token = await SubmitCode({
          code,
          lang
        });
        console.log(token)
      }
    return(
        <div className="usl">
            <h2>Add two numbers</h2>
            You are given two non-empty linked lists representing two non-negative integers. 
            <form onSubmit={handleSubmit}>
            <label><h3>Solution</h3></label>
                <label>Choose a language: </label>
                <select name="lang" onChange={e => setLang(e.target.value)}>
                    <option disabled selected value> -- Select language -- </option>
                    <option value="python">Python 3</option>
                    <option value="c++">C++</option>
                    <option value="java">Java</option>
                    <option value="go">Go</option>
                </select>
                <br></br>
                <textarea name="twosum" className = "solution" onChange={e => setCode(e.target.value)}>
                </textarea>
                <br></br>
                <input type="submit" value="Submit"></input>
            </form>
        </div>
    );
}