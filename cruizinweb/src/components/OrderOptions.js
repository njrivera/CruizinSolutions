import React from 'react';
import TireGrid from './TireGrid';
import RimGrid from './RimGrid';
import PartGrid from './PartGrid';
import ServiceGrid from './ServiceGrid';

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
            case 'rims':
            return (
                <div>
                    <RimGrid 
                        extra={true}
                        extraTitle={'Add Item'}
                        extraFunction={this.props.extraFunction}/>
                </div>
            );
            case 'parts':
            return (
                <div>
                    <PartGrid 
                        extra={true}
                        extraTitle={'Add Item'}
                        extraFunction={this.props.extraFunction}/>
                </div>
            );
            case 'services':
            return (
                <div>
                    <ServiceGrid 
                        extra={true}
                        extraTitle={'Add Item'}
                        extraFunction={this.props.extraFunction}/>
                </div>
            );
            default: return (<div></div>);
        }
    }
}