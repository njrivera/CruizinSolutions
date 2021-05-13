import React from 'react';
import '../styles/App.css';
import Home from '../pages/Home';
import Info from '../pages/Info';
import Quote from '../pages/Quote';
import Past from '../pages/Past';
import Inventory from '../pages/Inventory';
import Order from '../pages/Order';
import Report from '../pages/Report';
import NavBar from './NavBar';
import "../../node_modules/react-bootstrap-table/dist/react-bootstrap-table.min.css";
import {Switch, Route, BrowserRouter} from 'react-router-dom';
import axios from 'axios';

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      taxRate: 0
    };
  }

  componentDidMount() {
    axios.get('/api/taxRate')
    .then(response => {
      this.setState({taxRate: parseFloat(response.data)});
    })
    .catch(err => {
      console.log("couldn't get tax rate :(");
    });
  }

  render() {
    return (
      <div className="App">
        <BrowserRouter>
          <div>
            <div className="App-header">
              <h1 className='title text-left'>CRUIZIN' SOLUTIONS</h1>
            </div>
            <div>
              <NavBar />
            </div>
            <div>
              <Switch>
                <Route exact path='/' component={Home}/>
                <Route path='/info/:option' component={Info}/>
                <Route exact path='/quote' render={(props) => (<Quote taxRate={this.state.taxRate} {...props} />)}/>
                <Route exact path='/past' component={Past}/>
                <Route exact path='/inventory/:option' component={Inventory}/>
                <Route exact path='/order' component={(props) => (<Order taxRate={this.state.taxRate} {...props} />)}/>
                <Route exact path='/report' component={Report}/>
              </Switch>
            </div>
          </div>
        </BrowserRouter>
      </div>
    );
  }
}

export default App;
