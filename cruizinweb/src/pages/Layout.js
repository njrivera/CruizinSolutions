import React from 'react';
import '../styles/Layout.css';

export default class Layout extends React.Component {

    render() {
        return (
            <div className='Layout'>
                <h1>Hi there!</h1>
                <button onClick={() => window.print()}>Print</button>
            </div>
        );
    }
}