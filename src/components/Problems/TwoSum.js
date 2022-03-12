import React, {useState} from 'react';
import './Problems.css';
async function SubmitCode(credentials){
    console.log(JSON.stringify(credentials))
    return fetch('http://localhost:8080/problems/1',{
    method: 'POST',
    headers:{
      Accept: 'application/json',
      'Content-Type':'application/json',
    },
    credentials: "include",
    body: JSON.stringify(credentials)
  })
  .then(data => data.json())
}
export default function TwoSum(){
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
        <div className = "usl">
            <h2>Two Sum</h2>
            Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.<br></br>
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