import React from 'react';
import CustomerGrid from '../components/CustomerGrid';
import CustVehiclesGrid from '../components/CustVehiclesGrid';
import VehicleGrid from '../components/VehicleGrid';
import {DropdownButton, MenuItem} from 'react-bootstrap';
import OrderOptions from '../components/OrderOptions';
import ItemList from '../components/ItemList';
import Invoice from '../components/Invoice';

export default class Order extends React.Component {
    constructor() {
        super();
        this.state = {
            grid: 'customer',
            customer: null,
            vehicle: null,
            product: 'tires',
            items: [],
            defaultPrices: [],
            finished: false
        };
        this.chooseCustomer = this.chooseCustomer.bind(this);
        this.addVehicle = this.addVehicle.bind(this);
        this.chooseVehicle = this.chooseVehicle.bind(this);
        this.addItem = this.addItem.bind(this);
        this.removeItem = this.removeItem.bind(this);
        this.changePrice = this.changePrice.bind(this);
        this.finishOrder = this.finishOrder.bind(this);
    }

    chooseCustomer(customer) {
        if(customer) {
            this.setState({customer: customer});
            this.setState({grid: 'custVehicles'});
        }
    }

    addVehicle() {
        this.setState({grid: 'vehicles'});
    }

    chooseVehicle(vehicle) {
        if(vehicle) {
            this.setState({vehicle: vehicle});
            this.setState({grid: 'items'});
        }
    }

    finishOrder() {
        if(this.state.items.length > 0) {
            this.setState({grid: 'invoice'});
            this.setState({finished: true});
        }
    }

    removeItem(record) {
        if(record) {
            var temp = JSON.parse(JSON.stringify(this.state.items));
            var tempPrices = JSON.parse(JSON.stringify(this.state.defaultPrices));
            temp = temp.filter(function(r) {
                return r.itemnum !== record.itemnum;
            });
            tempPrices = tempPrices.filter(function(r) {
                return r.itemnum !== record.itemnum;
            });
            this.setState({defaultPrices: tempPrices});
            this.setState({items: temp});
        }
    }

    addItem(record) {
        if(record) {
            var temp = JSON.parse(JSON.stringify(this.state.items));
            var tempPrices = JSON.parse(JSON.stringify(this.state.defaultPrices));
            if(temp.filter(function(r){return r.itemnum === record.itemnum;}).length === 0) {
                switch(this.state.product) {
                    case 'tires':
                    temp.push({
                        itemnum: record.itemnum,
                        description: 
                            record.brand + ' ' + 
                            record.model + ' ' + 
                            record.size + ' ' +
                            record.servicedesc + ' (' +
                            record.condition + ')',
                        qty: 1,
                        amount: record.price
                    });
                    break;
                    case 'rims':
                    temp.push({
                        itemnum: record.itemnum,
                        description: 
                            record.brand + ' ' +
                            record.model + ' ' +
                            record.size + ' ' +
                            record.boltpattern + ' ' +
                            record.finish + ' (' +
                            record.condition + ')',
                        qty: 1,
                        amount: record.price
                    });
                    break;
                    case 'parts':
                    temp.push({
                        itemnum: record.itemnum,
                        description: 
                            record.description + ' (' +
                            record.condition + ')',
                        qty: 1,
                        amount: record.price
                    });
                    break;
                    case 'services':
                    temp.push({
                        itemnum: record.itemnum,
                        description: record.description,
                        qty: 1,
                        amount: record.price
                    });
                    default: break;
                }
                tempPrices.push({itemnum: record.itemnum, price: record.price});
                this.setState({defaultPrices: tempPrices});
                this.setState({items: temp});
            }
        }
    }

    changePrice(itemnum, qty) {
        var temp = JSON.parse(JSON.stringify(this.state.items));
        for (var i = 0; i < temp.length; i++) {
            if (temp[i].itemnum === itemnum) {
                temp[i].amount = (Math.round(this.state.defaultPrices[i].price * parseFloat(qty) * 100) / 100).toFixed(2);
                break;
            }
        }
        this.setState({items: temp});
    }

    render() {
        switch(this.state.grid){
            case 'customer':
            return (
                <div>
                    <h1>Choose Customer</h1>
                    <CustomerGrid
                        extra={true}
                        extraTitle={'Choose Customer'}
                        extraFunction={this.chooseCustomer}/>
                </div>
            );
            case 'custVehicles':
            return (
                <div>
                    <h1>Choose Customer Vehicle</h1>
                    <CustVehiclesGrid
                        id={this.state.customer.cid}
                        onAdd={this.addVehicle}
                        onChoose={this.chooseVehicle}/>
                </div>
            );
            case 'vehicles':
            return (
                <div>
                    <h1>Add Vehicle</h1>
                    <VehicleGrid
                        extra={true}
                        extraTitle={'Choose Vehicle'}
                        extraFunction={this.chooseVehicle}/>
                </div>
            );
            case 'items':
            return (
                <div>
                    <DropdownButton title='Choose Product' id='products' onSelect={(event) => {this.setState({product: event})}}>
                        <MenuItem eventKey='tires'>Tires</MenuItem>
                        <MenuItem eventKey='rims'>Rims</MenuItem>
                        <MenuItem eventKey='parts'>Parts</MenuItem>
                        <MenuItem eventKey='services'>Services</MenuItem>
                    </DropdownButton>
                    <OrderOptions
                        product={this.state.product}
                        extraFunction={this.addItem}/>
                    <ItemList
                        items={this.state.items}
                        removeItem={this.removeItem}
                        changePrice={this.changePrice}
                        finishOrder={this.finishOrder}/>
                </div>
            );
            case 'invoice':
            return (
                <Invoice finished={this.state.finished}/>
            );
            default: return (<div></div>);
        }
    }
}