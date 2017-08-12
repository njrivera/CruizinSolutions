import React from 'react';
import TireGrid from '../components/TireGrid'
import RimGrid from '../components/RimGrid'
import PartGrid from '../components/PartGrid'
import ServiceGrid from '../components/ServiceGrid'

export default class Inventory extends React.Component {

    render() {
        switch(this.props.match.params.option) {
            case 'tires':
                return (
                    <TireGrid/>
                );
            case 'rims':
                return (
                    <RimGrid/>
                );
            case 'parts':
                return (
                    <PartGrid/>
                );
            case 'services':
                return (
                    <ServiceGrid/>
                );
            default:
                return (
                    <div>Sorry!!!</div>
                );
        }
    }
}