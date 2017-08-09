import React from 'react';

export default class Inventory extends React.Component {

    render() {
        switch(this.props.match.params.option) {
            case 'tires':
                return (
                    <div>
                        <h1>Tires</h1>
                    </div>
                );
            break;
            case 'rims':
                return (
                    <div>
                        <h1>Rims</h1>
                    </div>
                );
            break;
            case 'items':
                return (
                    <div>
                        <h1>Items</h1>
                    </div>
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