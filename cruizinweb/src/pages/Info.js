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
            break;
            case 'vehicles':
                return (
                    <VehicleGrid />
                );
            break;
            default:
                return (
                        <div>
                            <h1>none selected</h1>
                        </div>
                    );
        }
    }
}