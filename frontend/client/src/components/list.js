import React, {useEffect, useState} from "react";
import Loading from "./loading";

const List = () => {

  const [itemList, setItemList] = useState([{}]);

  useEffect(() => {
    setItemList(undefined);
    fetch("http://localhost:5001/api/items?limit=100&offset=0").then(
      response => response.json()
    ).then(
      data => {
        // TO DO: validate data errors
        setItemList(data)
      }
    )
  }, []);

  const round = (num) => {
    return (Math.round(num * 100) / 100).toFixed(2);
  }

  return (
    <>
      <div className="row card-container">
        {(typeof itemList === 'undefined') ? (
          <Loading />
        ) : (
          <div>
            {itemList !== undefined && itemList['items'] !== undefined && itemList['items'].map(item => (
                  <div key={item['id']} className="col s12 m6 l3 xl2">
                    <div className="card-panel lighten-4">
                      <div className="card">
                        <div className="card-image" href={`/items/${item['id']}`}>
                        <a href={`/portal/items/${item['id']}`}><img alt={item['thumbnail']} src={item['thumbnail']} /></a>
                          <a href="/#" className="btn-floating halfway-fab waves-effect waves-light blue">
                            <i className="material-icons">favorite_border</i>
                          </a>
                        </div>
                        <div className="card-content">
                          <p className="card-title activator">{item['name']}</p>
                          <p className="flow-text activator">U$D {round(item['price'])}</p>
                        </div>
                        <div className="card-reveal">
                          <span className="card-title grey-text text-darken-4">{item['name']}
                            <i className="material-icons right">close</i>
                          </span>
                          <p>{item['description']}</p>
                        </div>
                    </div>
                  </div>
                </div>
            ))}
          </div>
        )}
      </div>
      <ul className="pagination center lighten-2">
        <li className="disabled"><a href="/#"><i className="material-icons">chevron_left</i></a></li>
        <li className="active black"><a href="/#">1</a></li>
        <li className="waves-effect"><a href="/#">2</a></li>
        <li className="waves-effect"><a href="/#">3</a></li>
        <li className="waves-effect"><a href="/#">4</a></li>
        <li className="waves-effect"><a href="/#">5</a></li>
        <li className="waves-effect"><a href="/#"><i className="material-icons">chevron_right</i></a></li>
      </ul>
    </>
  );
}
  
export default List;