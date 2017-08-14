import React from 'react';
import CustomerGrid from '../components/CustomerGrid';
import CustVehiclesGrid from '../components/CustVehiclesGrid';
import VehicleGrid from '../components/VehicleGrid';
import {DropdownButton, MenuItem, FormGroup, FormControl, ControlLabel} from 'react-bootstrap';
import OrderOptions from '../components/OrderOptions';
import ItemList from '../components/ItemList';
import Invoice from '../components/Invoice';
import {Container, Row, Col, Button} from 'reactstrap';


export default class Order extends React.Component {
    constructor() {
        super();
        this.state = {
            grid: 'customer',
            customer: null,
            vehicle: null,
            product: 'tires',
            items: [],
            finished: false,
            subtotal: "",
            tax: "",
            total: "",
            taxRate: .07,
            date: new Date(),
            invoiceNum: null,
            comments: '',
            odometer: ''
        };
        this.chooseCustomer = this.chooseCustomer.bind(this);
        this.addVehicle = this.addVehicle.bind(this);
        this.chooseVehicle = this.chooseVehicle.bind(this);
        this.addItem = this.addItem.bind(this);
        this.removeItem = this.removeItem.bind(this);
        this.changePrice = this.changePrice.bind(this);
        this.finishOrder = this.finishOrder.bind(this);
        this.addUpTotal = this.addUpTotal.bind(this);
    }

    addUpTotal() {
        var subtotal = 0.00;
        var tax = 0.00;
        for(var i = 0; i < this.state.items.length; i++) {
            subtotal += this.state.items[i].amount;
            if (this.state.items[i].tax) {
                tax += this.state.taxRate * this.state.items[i].amount;
            }
        }
        this.setState({subtotal: subtotal.toFixed(2)});
        this.setState({tax: tax.toFixed(2)});
        this.setState({total: (subtotal + tax).toFixed(2)});
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
            temp = temp.filter(function(r) {
                return r.itemnum !== record.itemnum;
            });
            this.setState({items: temp}, this.addUpTotal);
        }
    }

    addItem(record) {
        if(record) {
            var temp = JSON.parse(JSON.stringify(this.state.items));
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
                        amount: record.price,
                        price: record.price,
                        tax: record.condition === 'NEW' ? true : false
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
                        amount: record.price,
                        price: record.price,
                        tax: record.condition === 'NEW' ? true : false
                    });
                    break;
                    case 'parts':
                    temp.push({
                        itemnum: record.itemnum,
                        description: 
                            record.description + ' (' +
                            record.condition + ')',
                        qty: 1,
                        amount: record.price,
                        price: record.price,
                        tax: record.condition === 'NEW' ? true : false
                    });
                    break;
                    case 'services':
                    temp.push({
                        itemnum: record.itemnum,
                        description: record.description,
                        qty: 1,
                        amount: record.price,
                        price: record.price,
                        tax: true
                    });
                    break;
                    default: break;
                }
                this.setState({items: temp}, this.addUpTotal);
            }
        }
    }

    changePrice(itemnum, qty) {
        var temp = JSON.parse(JSON.stringify(this.state.items));
        for (var i = 0; i < temp.length; i++) {
            if (temp[i].itemnum === itemnum) {
                temp[i].amount = temp[i].price * parseFloat(qty);
                break;
            }
        }
        this.setState({items: temp}, this.addUpTotal);
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
                <Container>
                    <Row>
                        <DropdownButton title='Choose Product' id='products' onSelect={(event) => {this.setState({product: event})}}>
                            <MenuItem eventKey='tires'>Tires</MenuItem>
                            <MenuItem eventKey='rims'>Rims</MenuItem>
                            <MenuItem eventKey='parts'>Parts</MenuItem>
                            <MenuItem eventKey='services'>Services</MenuItem>
                        </DropdownButton>
                        <p></p>
                        <OrderOptions
                            product={this.state.product}
                            extraFunction={this.addItem}/>
                    </Row>
                    <p></p>
                    <Row>
                        <Col sm='6'>
                            <p className="text-left">{this.state.customer.name}</p>
                            <p className="text-left">{this.state.customer.address}</p>
                            <p className="text-left">{this.state.customer.city}, {this.state.customer.state} {this.state.customer.zipcode}</p>
                        </Col>
                        <Col sm='6'>
                            <p className="text-right">{this.state.vehicle.year} {this.state.vehicle.make} {this.state.vehicle.model}</p>
                            <Row>
                                <Col sm='9' className='text-right'><h4>Odometer:</h4></Col>
                                <Col sm='3'>
                                    <form>
                                        <FormGroup controlId='commentForm'>
                                            <FormControl type='text' value={this.state.odometer} onChange={(event) => this.setState({odometer: event.target.value})}/>
                                        </FormGroup>
                                    </form>
                                </Col>
                            </Row>
                        </Col>
                    </Row>
                    <Row>
                        <ItemList
                            items={this.state.items}
                            removeItem={this.removeItem}
                            changePrice={this.changePrice}
                            finishOrder={this.finishOrder}/>
                    </Row>
                    <Row>
                        <Col sm='8'>
                            <Row className='text-left'>
                                <form>
                                    <FormGroup controlId='commentForm'>
                                        <h4>Comments:</h4>
                                        <FormControl type='text' componentClass='textarea' value={this.state.comments} onChange={(event) => this.setState({comments: event.target.value})}/>
                                    </FormGroup>
                                </form>
                            </Row>
                        </Col>
                        <p></p>
                        <Col sm='4'>
                            <p className="text-right">Subtotal: {this.state.subtotal}</p>
                            <p className="text-right">Tax: {this.state.tax}</p>
                            <p className="text-right">Total: {this.state.total}</p>
                        </Col>
                    </Row>
                </Container>
            );
            case 'invoice':
            return (
                <Invoice 
                    finished={this.state.finished}
                    invoiceNum={this.state.invoiceNum}
                    date={this.state.date.getMonth() + 1 + '/' + this.state.date.getDate() + '/' + this.state.date.getFullYear()}
                    customer={this.state.customer}
                    vehicle={this.state.vehicle}
                    items={this.state.items}
                    subtotal={this.state.subtotal}
                    tax={this.state.tax}
                    total={this.state.total}
                    odometer={this.state.odometer}
                    comments={this.state.comments}/>
            );
            default: return (<div></div>);
        }
    }
}