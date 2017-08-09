import React from 'react';
import '../styles/Layout.css'
import {Container, Row, Col} from 'reactstrap';
import CustomerGrid from '../components/CustomerGrid'

export default class Layout extends React.Component {

    render() {
        return (
            <div className='Layout'>
                <h1>Hi there!</h1>
            </div>
        );
    }
}