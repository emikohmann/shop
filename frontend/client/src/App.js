import React from 'react';

import './App.css';
import "./materialize/css/materialize.css";

import Navbar from './components/navbar';
import Search from './components/search';
import Footer from './components/footer';
import { BrowserRouter as Router, Routes, Route} from 'react-router-dom';

import Home from './pages/home';
import Item from './pages/item';
import Checkout from './pages/checkout';

function App(props) {
  return (

    <Router>
      <Navbar />
      <Search />

      <Routes>
        <Route exact path='/' element={<Home />} />
        <Route path='/item' element={<Item/>} />
        <Route path='/checkout' element={<Checkout/>} />
      </Routes>

      <Footer />
    </Router>
  );
}

export default App;