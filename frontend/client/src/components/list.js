import React, {useEffect, useState} from "react";
  
const List = () => {
  const [itemList, setItemList] = useState([{}]);

  useEffect(() => {
    fetch("http://localhost:5001/api/items?limit=10&offset=0").then(
      response => response.json()
    ).then(
      data => {
        console.log(`API layer list response: ${data}`);
        setItemList(data)
      }
    )
  }, []);

  return (
    <>
        <div>
        {(typeof itemList === 'undefined') ? (
          <p>Loading item list...</p>
        ) : (
          <ul className="collection">
            {itemList !== undefined && itemList['items'] !== undefined && itemList['items'].map(item => (
              <li className="collection-item avatar">
                <img className="circle" src={item['thumbnail']} alt=""/>
                <span className="title">{item['name']}</span>
                <p>ID: {item['id']}</p>
                <p>Description{item['description']}</p>
                <p>Price: {item['price']}</p>
                <a href="#!" className="secondary-content">
                  <i class="material-icons">grade</i>
                </a>
                <div className="center">
                  <a className="waves-effect waves-light btn blue-grey">
                    <i class="material-icons left">remove_red_eye</i> More info
                  </a>
                </div>
              </li>
            ))}
          </ul>
        )}
        </div>
    </>
  );
};
  
export default List;
