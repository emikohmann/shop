import React, {useEffect, useState} from "react";

import M from 'materialize-css';

import {useParams} from "react-router-dom";

import Loading from '../components/loading';

import ImageGallery from 'react-image-gallery';

import 'react-image-gallery/styles/css/image-gallery.css';

const Item = () => {
  const [itemData, setItemData] = useState([{}]);
  const [itemImages, setItemImages] = useState([{}]);

  let { id } = useParams();

  useEffect(() => {
    setItemData(undefined);
    fetch(`http://localhost:5001/api/items/${id}`).then(
      response => response.json()
    ).then(
      data => {
        setItemData(data);
        var itemImages = [];
        data['images'].map(img => {
          itemImages.push({
            original: img,
            thumbnail: img
          })
        });
        setItemImages(itemImages);
      }
    );
  }, []);

  return (
    <>
      {typeof itemData === 'undefined' ? (
        <>
          <br />
          <Loading />
        </>
      ) : (
        <div>
          <div>
            <h1>
              {itemData['name']}
            </h1>
            <div>
              {itemData['description']}
            </div>
            <br />
            <div className="left">
              <ImageGallery 
                items={itemImages}
                showFullscreenButton={true}
                useBrowserFullscreen={true}
                showPlayButton={false}
                showBullets={false}
                disableThumbnailScroll={false}
                slideDuration={100}
                thumbnailPosition="right"
              />
            </div>
          </div>
          <table>
            <tbody>
              <tr>
                <td>ID</td>
                <td>{itemData['id']}</td>
              </tr>
              <tr>
                <td>Thumbnail</td>
                <td><img src={itemData['thumbnail']} /></td>
              </tr>
              <tr>
                <td>Name</td>
                <td>{itemData['name']}</td>
              </tr>
              <tr>
                <td>Description</td>
                <td>{itemData['description']}</td>
              </tr>
              <tr>
                <td>Price</td>
                <td>{itemData['price']}</td>
              </tr>
            </tbody>
          </table>
        </div>
      )}
    </>
  );
};
  
export default Item;