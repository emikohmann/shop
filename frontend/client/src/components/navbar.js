import React from "react";
  
const Navbar = () => {
  return (
    <>
     <nav>
        <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet" />
        <div className="nav-wrapper blue-grey">
        <a href="/" className="brand-logo center">
          <i className="large material-icons">insert_chart</i> The shop
        </a>
        <ul id="nav-mobile" className="center hide-on-med-and-down">
            <li><a href="/checkout">Checkout</a></li>
            <li><a href="/item">Item</a></li>
        </ul>
        </div>
    </nav>
    </>
  );
};
  
export default Navbar;