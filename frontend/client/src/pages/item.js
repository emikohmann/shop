import React, {useEffect, useState} from "react";

import M from 'materialize-css';

import {useParams} from "react-router-dom";

import Breadcrumb from '../components/breadcrumb';

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
        // TO DO: validate data errors
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
        <div className="row">
          <Breadcrumb />
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
            {[...Array(itemData['punctuation'])].map((_, i) => 
              <i key={`grade_${i}`} className="small material-icons yellow-text">grade</i>)
            }
            <h4>U$D {itemData['price']}</h4>
            <p>{itemData['description']}</p>
          </div>
          <div className="col s12 m12 l7 xl8">
            <br />
            <h4>Ask for {itemData['name']}</h4>
            <div className="row">
              <form className="col s12">
                <div className="row">
                  <div className="input-field col s6">
                    <input id="first_name" type="text" />
                    <label htmlFor="first_name">First Name</label>
                  </div>
                  <div className="input-field col s6">
                    <input id="last_name" type="text" className="validate" />
                    <label htmlFor="last_name">Last Name</label>
                  </div>
                </div>
                <div className="row">
                  <div className="input-field col s12">
                    <input id="email" type="email" />
                    <label htmlFor="email">Email</label>
                  </div>
                </div>
                <div className="row">
                  <div className="input-field col s12">
                    <textarea id="textarea1" className="materialize-textarea"></textarea>
                    <label htmlFor="textarea1">Your question</label>
                  </div>
                </div>
              </form>
            </div>
          </div>
        </div>
      )}
    </>
  );
};
  
export default Item;