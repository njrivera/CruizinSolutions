import React from 'react';
import {BootstrapTable, TableHeaderColumn} from 'react-bootstrap-table';
import GridModal from './GridModal';
import {Container, Row, Col, Button} from 'reactstrap';
import axios from 'axios';
import TireGrid from './TireGrid';
import ItemGrid from './ItemGrid';

export default class OrderOptions extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        switch(this.props.product) {
            case 'tires':
            return (
                <div>
                    <TireGrid 
                        extra={true}
                        extraTitle={'Add Item'}
                        extraFunction={this.props.extraFunction}/>
                </div>
            );
            break;
            case 'items':
            return (
                <div>
                    <ItemGrid 
                        extra={true}
                        extraTitle={'Add Item'}
                        extraFunction={this.props.extraFunction}/>
                </div>
            );
            break;
            default: return (<div></div>);
        }
    }
}