import React from 'react';

import './App.css';

import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';

import Navbar from './components/navbar';
import Login from './pages/login';
import Portal from './pages/portal';
import Footer from './components/footer';

function App(props) {
  return (
    <>
      { /*Generic import for Material Icons*/ }
      <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet" />
      <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0-alpha.3/css/materialize.min.css" />
      <div>
        <Navbar />
      </div>
        <Router>
          <Routes>
              <Route path="/login" element={<Login />} />
              <Route path="/portal/*" element={<Portal />} />
              <Route path="/*" element={<Portal />} />
          </Routes>
        </Router>
      <div id="footerContainer">
          <Footer  />
      </div> 
    </>
  );
}

export default App;