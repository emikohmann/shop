import React from "react";
  
const Search = () => {
  return (
    <>
      <div className="center">
            <form className="">
                <div className="row">
                    <div className="input-field col s12 blue-grey-text">
                    <i className="material-icons prefix blue-grey-text">search</i>
                    <input className="materialize-textarea blue-grey-text" type="text" placeholder="Search products" id="icon_prefix2"/>
                    </div>
                </div>
            </form>
        </div>
    </>
  );
};
  
export default Search;