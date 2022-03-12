import React, { useState } from 'react';
import './App.css';
import Login from '../Login/Login';
import Signup from '../Login/Signup';
import Problems from '../Problems/Problems';
import { BrowserRouter, Route, Switch} from 'react-router-dom';



function App() {
  const [token, setToken] = useState();
  // if (!token){
  //   return(
  //     <Login setToken={setToken} />
  //   )
  // }
  // if (token){
  //   return(
  //     <Problems />
  //   )
  // }
  return (
    <div className="wrapper">
      <BrowserRouter>
        <Switch>
          <Route exact path="/" component={Login} />
          <Route path="/problems" component={Problems} />
          <Route path="/login" component={Login} />
          <Route path="/signup" component={Signup} />
        </Switch>
      </BrowserRouter>
    </div>
  );
}

export default App;
