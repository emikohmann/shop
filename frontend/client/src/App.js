import React from 'react';

import './App.css';

import Navbar from './components/navbar';
import Sidebar from './components/sidebar';
import Footer from './components/footer';
import { BrowserRouter as Router, Routes, Route} from 'react-router-dom';

import Home from './pages/home';
import Item from './pages/item';
import Categories from './pages/categories';

function App(props) {
  return (
    <>
      { /*Generic import for Material Icons*/ }
      <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet" />
      <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0-alpha.3/css/materialize.min.css" />
      <div>
        <Navbar />
      </div>
      <div className="row">
        <div className="col s2" >
          <Sidebar />
        </div>
        <div className="col s10">
          <Router>
            <Routes>
              <Route exact path='/' element={<Home />} />
              <Route path='/home' element={<Home/>} />
              <Route path='/item' element={<Item/>} />
              <Route path='/categories' element={<Categories/>} />
            </Routes>
          </Router>
        </div>
      </div>
      <div id="footerContainer">
        <Footer  />
      </div> 
    </>
  );
}

export default App;