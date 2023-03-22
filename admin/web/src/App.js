import './App.css';

import '98.css';

import {Config} from './Config';

import {API} from './API';

function handleBuild() {
  API.build();
}

function handleStart() {
  API.start();
}

function handleStop() {
  API.stop();
}

function App() {
  return (
    <>
      <h4><img className="adminIcon" alt="Icon" src="https://www.hsenidmobile.com/wp-content/uploads/2015/09/tools-icon-white.png" /> Admin v0.0.1</h4>
      <div className="window">
        <div className="title-bar">
            <div className="title-bar-text">
              Services
            </div>
        </div>
        <div className="window-body">
          <p>Hello, world!</p>
          <menu role="tablist">
            <li role="tab" aria-selected="true"><a href="#tabs">Services</a></li>
            <li role="tab"><a href="#tabs">Service detail</a></li>
            <li role="tab"><a href="#tabs">Docker images</a></li>
          </menu>
          <div className="window" role="tabpanel">
            <div className="window-body">
            <table>
                <thead>
                  <tr>
                    <td>Type</td>
                    <td>Name</td>
                    <td>ID</td>
                    <td>Technology</td>
                    <td>Current version</td>
                    <td colSpan={3}>Actions</td>
                    <td colSpan={2}>Links</td>
                  </tr>
                </thead>
                <tbody>
                  {Config.services.map(service => (
                    <tr key={service.id}>
                      <td>{service.type}</td>
                      <td>{service.display_name}</td>
                      <td>{service.id}</td>
                      <td className="serviceIcon">
                        <img src={Config.icons[service.technology]} alt={service.technology}/>
                      </td>
                      <td>
                        0.0.1
                      </td>
                      <td className="actions">
                        Version: <input type="text" placeholder="0.0.1" />
                        <button onClick={handleBuild} >Build</button>
                      </td>
                      <td className="actions">
                        <button onClick={handleStart} >Start</button>
                      </td>
                      <td className="actions">
                        <button onClick={handleStop} >Stop</button>
                      </td>
                      <td className="links">
                        <a href={service.repository}>Github</a>
                      </td>
                      <td className="links">
                        <a href={service.docs}>Swagger Docs</a>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>  
            </div>
          </div>
        </div>
      </div>
      <div className="outer">
      </div>
    </>
  );
}

export default App;
