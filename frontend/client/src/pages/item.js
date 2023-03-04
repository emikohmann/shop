import React, {useEffect, useState} from "react";
  
const Item = () => {
  const [backendData, setBackendData] = useState([{}])

  useEffect(() => {
    fetch("http://localhost:5001/api/items/1").then(
      response => response.json()
    ).then(
      data => {
        console.log(data);
        setBackendData(data)
      }
    )
  }, []);

  return (
    <div>
      <h1>
        Item page
      </h1>

      <div>
        {(typeof backendData === 'undefined') ? (
          <p>Loading...</p>
        ) : (
          <table>
            <tbody>
              <tr><td>Thumbnail</td><td><img src={backendData['thumbnail']} /></td></tr>
              <tr><td>ID</td><td>{backendData['id']}</td></tr>
              <tr><td>Name</td><td>{backendData['name']}</td></tr>
              <tr><td>Description</td><td>{backendData['description']}</td></tr>
              <tr><td>Price</td><td>{backendData['price']}</td></tr>
              <tr><td>Images</td><td>{backendData['images'] !== undefined && backendData['images'].map((value) => (
                  <img src={value} />
                ))}</td></tr>
            </tbody>
          </table>
        )}
      </div>
    </div>
  );
};
  
export default Item;