import React from 'react';
import './Problems.css';

export default function Longest(){
    return(
        <div className = "usl">
            <h2>Longest Substring</h2>
            Given a string s, find the length of the longest substring without repeating characters.
            <form>
                <label><h3>Solution</h3></label>
                <label>Choose a language: </label>
                <select name="lang">
                    <option value="python">Python 3</option>
                    <option value="c++">C++</option>
                    <option value="java">Java</option>
                    <option value="go">Go</option>
                </select>
                <br></br>
                <textarea name="longest" className = "solution">
                </textarea>
                <br></br>
                <input type="submit" value="Submit"></input>
            </form>
        </div>
    );
}