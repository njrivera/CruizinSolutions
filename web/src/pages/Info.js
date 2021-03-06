import React from 'react';
import CustomerGrid from '../components/CustomerGrid';
import VehicleGrid from '../components/VehicleGrid';

export default class Info extends React.Component {

    render() {
        switch(this.props.match.params.option) {
            case 'customers':
                return (
                    <CustomerGrid />
                );
            case 'vehicles':
                return (
                    <VehicleGrid />
                );
            default:
                return (
                        <div>
                            <h1>none selected</h1>
                        </div>
                    );
        }
    }
}