import React from 'react';

const Navbar = () => {
    return (
        <>
            <div className="navbar">
                <nav>
                    <div className="nav-wrapper black"> 
                        <a href="/" className="brand-logo" id="brand">
                            <i className="material-icons">cloud</i>Shop
                        </a>
                    </div>
                </nav>
            </div>
        </>
    );
};

export default Navbar;