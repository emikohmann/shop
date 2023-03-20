import './App.css';

import '98.css';

function App() {
  return (
    <>
      <h4><img className="icon" src="https://www.hsenidmobile.com/wp-content/uploads/2015/09/tools-icon-white.png" /> Admin v0.0.1</h4>
      <div className="window">
        <div className="title-bar">
            <div className="title-bar-text">
              Services
            </div>
        </div>
        <div className="window-body">
          <p>Hello, world!</p>
          <menu role="tablist">
            <li role="tab" aria-selected="true"><a href="#tabs">Backend</a></li>
            <li role="tab"><a href="#tabs">Frontend</a></li>
            <li role="tab"><a href="#tabs">Docker images</a></li>
          </menu>
          <div className="window" role="tabpanel">
            <div className="window-body">
            <table>
                <thead>
                  <tr><td>Type</td><td>Name</td><td colSpan={3}>Actions</td><td colSpan={2}>Links</td></tr>
                </thead>
                <tbody>
                  <tr>
                    <td>Backend</td>
                    <td>Items API</td>
                    <td className="actions">
                      <button>Build</button>
                    </td>
                    <td className="actions">
                      <button>Start</button>
                    </td>
                    <td className="actions">
                      <button>Stop</button>
                    </td>
                    <td className="links">
                      <a href="https://github.com">Github</a>
                    </td>
                    <td className="links">
                      <a href="http://localhost/docs/index.html">Swagger Docs</a>
                    </td>
                  </tr>
                  <tr>
                    <td>Backend</td>
                    <td>Users API</td>
                    <td className="actions">
                      <button>Build</button>
                    </td>
                    <td className="actions">
                      <button>Start</button>
                    </td>
                    <td className="actions">
                      <button>Stop</button>
                    </td>
                    <td className="links">
                      <a href="https://github.com">Github</a>
                    </td>
                    <td className="links">
                      <a href="http://localhost/docs/index.html">Swagger Docs</a>
                    </td>
                  </tr>
                  <tr>
                    <td>Backend</td>
                    <td>Stores API</td>
                    <td className="actions">
                      <button>Build</button>
                    </td>
                    <td className="actions">
                      <button>Start</button>
                    </td>
                    <td className="actions">
                      <button>Stop</button>
                    </td>
                    <td className="links">
                      <a href="https://github.com">Github</a>
                    </td>
                    <td className="links">
                      <a href="http://localhost/docs/index.html">Swagger Docs</a>
                    </td>
                  </tr>
                  <tr>
                    <td>Backend</td>
                    <td>Orders API</td>
                    <td className="actions">
                      <button>Build</button>
                    </td>
                    <td className="actions">
                      <button>Start</button>
                    </td>
                    <td className="actions">
                      <button>Stop</button>
                    </td>
                    <td className="links">
                      <a href="https://github.com">Github</a>
                    </td>
                    <td className="links">
                      <a href="http://localhost/docs/index.html">Swagger Docs</a>
                    </td>
                  </tr>
                  <tr>
                    <td>Backend</td>
                    <td>Discounts API</td>
                    <td className="actions">
                      <button>Build</button>
                    </td>
                    <td className="actions">
                      <button>Start</button>
                    </td>
                    <td className="actions">
                      <button>Stop</button>
                    </td>
                    <td className="links">
                      <a href="https://github.com">Github</a>
                    </td>
                    <td className="links">
                      <a href="http://localhost/docs/index.html">Swagger Docs</a>
                    </td>
                  </tr>
                  <tr>
                    <td>Frontend</td>
                    <td>Server</td>
                    <td className="actions">
                      <button>Build</button>
                    </td>
                    <td className="actions">
                      <button>Start</button>
                    </td>
                    <td className="actions">
                      <button>Stop</button>
                    </td>
                    <td className="links">
                      <a href="https://github.com">Github</a>
                    </td>
                    <td className="links">
                      <a href="http://localhost/docs/index.html">Swagger Docs</a>
                    </td>
                  </tr>
                  <tr>
                    <td>Frontend</td>
                    <td>Client</td>
                    <td className="actions">
                      <button>Build</button>
                    </td>
                    <td className="actions">
                      <button>Start</button>
                    </td>
                    <td className="actions">
                      <button>Stop</button>
                    </td>
                    <td className="links">
                      <a href="https://github.com">Github</a>
                    </td>
                    <td className="links">
                      <a href="http://localhost/docs/index.html">Swagger Docs</a>
                    </td>
                  </tr>
                </tbody>
              </table>  
            </div>
          </div>
        </div>
      </div>
      <div className="outer">
            Welcome! <br />
            Here will be the app logs..
      </div>
    </>
  );
}

export default App;
