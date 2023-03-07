import React from 'react';

const Navbar = () => {
    return (
        <>
            <div className="navbar">
                <nav>
                    <div className="nav-wrapper black"> 
                        <a href="/home" className="brand-logo" id="brand">
                            <i className="material-icons">cloud</i>Shop
                        </a>
                        <ul className="right hide-on-med-and-down">
                            <li><a href="/categories">Categories</a></li>
                            <li><a href="/home">Home</a></li>
                        </ul>
                    </div>
                </nav>
            </div>
        </>
    );
};

export default Navbar;