import React from 'react';
import {Button} from 'react-bootstrap';

export default class Layout extends React.Component {
    constructor(){
        super();
        const button = (
            <Button bsStyle='primary'></Button>
        )
    }

    render(){
        return (
            <div>
                {this.button}
            </div>
        )
    }
}