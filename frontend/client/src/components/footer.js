import React from "react";
  
const Footer = () => {
  return (
    <>
        <footer className="page-footer black">
          <div className="container">
            <div className="row">
              <div className="col l6 s12">
                <h5 className="white-text">Shop Footer</h5>
                <p className="grey-text text-lighten-4">Shop footer content organized in rows and columns.</p>
              </div>
              <div className="col l4 offset-l2 s12">
                <h5 className="white-text">Links</h5>
                <ul>
                  <li><a className="grey-text text-lighten-3" href="https://github.com/emikohmann/shop">Github</a></li>
                  <li><a className="grey-text text-lighten-3" href="http://localhost:8080/docs/index.html">Docs</a></li>
                </ul>
              </div>
            </div>
          </div>
          <div className="footer-copyright">
            <div className="container">
            © 2023 Copyright Text
            <a className="grey-text text-lighten-4 right" href="mailto:emikohmann@gmail.com">Contact us</a>
            </div>
          </div>
        </footer>
    </>
  );
};
  
export default Footer;
