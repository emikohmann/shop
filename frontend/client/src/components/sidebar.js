import React, { Component } from "react";

import M from 'materialize-css';

class Sidebar extends Component {
    render() {
        return (
            <>
                <ul id="slide-out" className="side-nav fixed grey-text">
                    <li>
                        <div className="input-field" id="searchContainer">
                            <i className="material-icons prefix" id="searchIcon">search</i>
                            <input id="search"  type="text"/>
                                <label htmlFor="search">Search!</label>
                        </div>
                    </li>
                    
                    <br />
                    <li><a href="/" className="subheader">Home</a></li>
                    <br />
                    <li><a href="/login" className="subheader">Login</a></li>
                    <br />
                    <li><a href="/categories" className="waves-effect">Categories</a></li>
                </ul>
            </>
        );
    }
}

export default Sidebar;