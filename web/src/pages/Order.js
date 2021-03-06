import React from 'react';
import CustomerGrid from '../components/CustomerGrid';
import CustVehiclesGrid from '../components/CustVehiclesGrid';
import VehicleGrid from '../components/VehicleGrid';
import {DropdownButton, MenuItem, FormGroup, FormControl} from 'react-bootstrap';
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
            subtotal: '',
            tax: '',
            total: '',
            date: new Date(),
            invoiceNum: null,
            comments: '',
            odometer: '0',
            payment: 'DEBIT'
        };
        this.chooseCustomer = this.chooseCustomer.bind(this);
        this.addVehicle = this.addVehicle.bind(this);
        this.chooseVehicle = this.chooseVehicle.bind(this);
        this.addItem = this.addItem.bind(this);
        this.removeItem = this.removeItem.bind(this);
        this.changePrice = this.changePrice.bind(this);
        this.finishOrder = this.finishOrder.bind(this);
        this.addUpTotal = this.addUpTotal.bind(this);
        this.onBack = this.onBack.bind(this);
        this.onFinish = this.onFinish.bind(this);
        this.validateOdometer = this.validateOdometer.bind(this);
    }

    onBack() {
        switch(this.state.grid) {
            case 'custVehicles': this.setState({grid: 'customer'});
            break;
            case 'vehicles': this.setState({grid: 'custVehicles'});
            break;
            case 'items': this.setState({grid: 'custVehicles'});
            break;
            case 'invoice': this.setState({grid: 'items'});
            break;
            default: break;
        }
    }

    onFinish() {
        setTimeout(() => {
            this.setState({
                grid: 'customer',
                customer: null,
                vehicle: null,
                product: 'tires',
                items: [],
                subtotal: '',
                tax: '',
                total: '',
                date: new Date(),
                invoiceNum: null,
                comments: '',
                odometer: '0'
            });
        }, 100);
    }

    addUpTotal() {
        var subtotal = 0.00;
        var tax = 0.00;
        for(var i = 0; i < this.state.items.length; i++) {
            subtotal += parseFloat(this.state.items[i].amount);
            if (this.state.items[i].tax) {
                tax += this.props.taxRate * parseFloat(this.state.items[i].amount);
            }
        }
        this.setState({
            subtotal: subtotal.toFixed(2),
            tax: tax.toFixed(2),
            total: (subtotal + tax).toFixed(2)
        });
    }

    chooseCustomer(customer) {
        if(customer) {
            this.setState({
                customer: customer,
                grid: 'custVehicles'
            });
        }
    }

    addVehicle() {
        this.setState({grid: 'vehicles'});
    }

    chooseVehicle(vehicle) {
        if(vehicle) {
            this.setState({
                vehicle: vehicle,
                grid: 'items'
            });
        }
    }

    finishOrder() {
        if(this.state.items.length > 0) {
            this.setState({grid: 'invoice'});
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
                        tax: false
                    });
                    break;
                    case 'packages':
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
                temp[i].amount = (parseFloat(temp[i].price) * parseInt(qty, 10)).toFixed(2);
                break;
            }
        }
        this.setState({items: temp}, this.addUpTotal);
    }

    validateOdometer(event) {
        if(event.target.value.length === 0)
            event.target.value = 0;
        else if(!Number(event.target.value)) {
                event.target.value = event.target.value.slice(1);
                return;
        }
        event.target.value = parseInt(event.target.value, 10);
        this.setState({odometer: event.target.value});
    }

    render() {
        switch(this.state.grid){
            case 'customer':
            return (
                <Container>
                    <h1>Choose Customer</h1>
                    <CustomerGrid
                        extra={true}
                        extraTitle={'Choose Customer'}
                        extraFunction={this.chooseCustomer}/>
                </Container>
            );
            case 'custVehicles':
            return (
                <Container>
                    <Row>
                        <Col sm='1'><Button color='info' className='text-left' onClick={this.onBack}>Back</Button></Col>
                        <Col sm='11'></Col>
                    </Row>
                    <h1>Choose Customer Vehicle</h1>
                    <CustVehiclesGrid
                        id={this.state.customer.cid}
                        onAdd={this.addVehicle}
                        onChoose={this.chooseVehicle}/>
                </Container>
            );
            case 'vehicles':
            return (
                <Container>
                    <Row>
                        <Col sm='1'><Button color='info' className='text-left' onClick={this.onBack}>Back</Button></Col>
                        <Col sm='11'></Col>
                    </Row>
                    <h1>Add Vehicle</h1>
                    <VehicleGrid
                        extra={true}
                        extraTitle={'Choose Vehicle'}
                        extraFunction={this.chooseVehicle}/>
                </Container>
            );
            case 'items':
            return (
                <Container>
                    <Row>
                        <Col sm='1'><Button color='info' className='text-left' onClick={this.onBack}>Back</Button></Col>
                        <Col sm='11'></Col>
                    </Row>
                    <Row>
                        <DropdownButton bsStyle='info' title='Choose Product' id='products' onSelect={(event) => {this.setState({product: event})}}>
                            <MenuItem eventKey='tires'>Tires</MenuItem>
                            <MenuItem eventKey='rims'>Rims</MenuItem>
                            <MenuItem eventKey='parts'>Parts</MenuItem>
                            <MenuItem eventKey='services'>Services</MenuItem>
                            <MenuItem eventKey='packages'>Packages</MenuItem>
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
                            <p className="text-left">{this.state.customer.email}</p>
                        </Col>
                        <Col sm='6'>
                            <p className="text-right">{this.state.vehicle.year} {this.state.vehicle.make} {this.state.vehicle.model}</p>
                            <Row>
                                <Col sm='9' className='text-right'><h4>Odometer:</h4></Col>
                                <Col sm='3'>
                                    <form>
                                        <FormGroup controlId='commentForm'>
                                            <FormControl type='text' value={this.state.odometer} onChange={(event) => this.validateOdometer(event)}/>
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
                            finishOrder={this.finishOrder}
                            buttonTitle={'Invoice'}/>
                    </Row>
                    <Row>
                        <Col sm='7'>
                            <Row className='text-left'>
                                <form>
                                    <FormGroup controlId='commentForm'>
                                        <h4>Comments:</h4>
                                        <FormControl type='text' componentClass='textarea' value={this.state.comments} onChange={(event) => this.setState({comments: event.target.value})}/>
                                    </FormGroup>
                                </form>
                            </Row>
                        </Col>
                        <Col sm='1'></Col>
                        <Col sm='2'>
                            <h4>Payment Method</h4>
                            <FormControl
                                className="text-right"
                                componentClass='select' 
                                placeholder='DEBIT' 
                                onChange={(event) => this.setState({payment: event.target.options[event.target.selectedIndex].label})}>
                                    <option value="DEBIT">DEBIT</option>
                                    <option value="CREDIT">CREDIT</option>
                                    <option value="CASH">CASH</option>
                                    <option value="CHECK">CHECK</option>
                                    <option value="STORECREDIT">STORE CREDIT</option>
                            </FormControl>
                        </Col>
                        <Col sm='2'>
                            <p className="text-right">Subtotal: {this.state.subtotal}</p>
                            <p className="text-right">Tax: {this.state.tax}</p>
                            <p className="text-right">Total: {this.state.total}</p>
                        </Col>
                    </Row>
                </Container>
            );
            case 'invoice':
            return (
                <Container>
                    <Row>
                        <Col sm='1'><Button color='info' className='text-left hidden-sm' onClick={this.onBack}>Back</Button></Col>
                        <Col sm='11'></Col>
                    </Row>
                    <Invoice
                        onFinish={this.onFinish}
                        invoiceNum={this.state.invoiceNum}
                        date={this.state.date.getMonth() + 1 + '/' + this.state.date.getDate() + '/' + this.state.date.getFullYear()}
                        customer={this.state.customer}
                        vehicle={this.state.vehicle}
                        items={this.state.items}
                        subtotal={this.state.subtotal}
                        tax={this.state.tax}
                        total={this.state.total}
                        odometer={this.state.odometer}
                        comments={this.state.comments}
                        printTitle={'Confirm & Print'}
                        payment={this.state.payment}/>
                </Container>
            );
            default: return (<div></div>);
        }
    }
}