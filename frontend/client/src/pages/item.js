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
        <>
        {typeof itemData === 'undefined' ? (
          <>
            <br />
            <Loading />
          </>
        ) : (
          <div className="row">
            <div className="col s12 center grey-text text-darken-3">
              <h3>{itemData['name']}</h3>
              <br />
            </div>
            <div className="col s12 m12 l5 xl4" id="galleryContainer">
              <ImageGallery 
                  items={itemImages}
                  showFullscreenButton={true}
                  useBrowserFullscreen={true}
                  showPlayButton={false}
                  showBullets={false}
                  disableThumbnailScroll={false}
                  slideDuration={100}
                  thumbnailPosition="bottom"
              />
            </div>
            <div className="col s12 m12 l7 xl8" id="descriptionContainer">
              <i className="small material-icons yellow-text">grade</i>
              <i className="small material-icons yellow-text">grade</i>
              <i className="small material-icons yellow-text">grade</i>
              <i className="small material-icons yellow-text">grade</i>
              <i className="small material-icons yellow-text">grade</i>
              <h4>U$D {itemData['price']}</h4>
              <p>{itemData['description']}</p>
            </div>
            <div className="col s12 m12 l7 xl8">
              <br />
              <h4>Ask for {itemData['name']}</h4>
              <div class="row">
                <form class="col s12">
                  <div class="row">
                    <div class="input-field col s6">
                      <input id="first_name" type="text" />
                      <label for="first_name">First Name</label>
                    </div>
                    <div class="input-field col s6">
                      <input id="last_name" type="text" class="validate" />
                      <label for="last_name">Last Name</label>
                    </div>
                  </div>
                  <div class="row">
                    <div class="input-field col s12">
                      <input id="email" type="email" />
                      <label for="email">Email</label>
                    </div>
                  </div>
                  <div class="row">
                    <div class="input-field col s12">
                      <textarea id="textarea1" class="materialize-textarea"></textarea>
                      <label for="textarea1">Your question</label>
                    </div>
                  </div>
                </form>
              </div>
            </div>
          </div>
        )}
        </>
      )}
    </>
  );
};
  
export default Item;