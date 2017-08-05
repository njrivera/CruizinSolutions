import React from 'react';
import '../styles/App.css';
import Layout from './Layout';

class App extends React.Component {
  render() {
    return (
      <div className="App">
        <div className="App-header">
          <h2>Welcome to React</h2>
        </div>
        <div><Layout /></div>
      </div>
    );
  }
}

export default App;
