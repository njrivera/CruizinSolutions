import React from 'react';
import '../styles/App.css';
import Layout from './Layout';
import "../../node_modules/react-bootstrap-table/dist/react-bootstrap-table.min.css";

class App extends React.Component {
  render() {
    return (
      <div className="App">
        <div className="App-header">
          <h1>Cruizin Manager</h1>
        </div>
        <div>
          <Layout />
        </div>
      </div>
    );
  }
}

export default App;
