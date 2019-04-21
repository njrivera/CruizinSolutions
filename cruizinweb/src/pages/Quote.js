import React from 'react';
import OrderOptions from '../components/OrderOptions';
import {Container, Row, Col, Button} from 'reactstrap';
import {DropdownButton, MenuItem} from 'react-bootstrap';
import ItemList from '../components/ItemList';
import PrintQuote from '../components/PrintQuote';

export default class Quote extends React.Component {
    constructor() {
        super();
        this.state = {
            grid: 'items',
            product: 'tires',
            items: [],
            subtotal: '',
            tax: '',
            total: '',
            taxRate: .085,
            date: new Date(),
        };
        this.addUpTotal = this.addUpTotal.bind(this);
        this.addItem = this.addItem.bind(this);
        this.removeItem = this.removeItem.bind(this);
        this.changePrice = this.changePrice.bind(this);
        this.finishOrder = this.finishOrder.bind(this);
    }

    addUpTotal() {
        var subtotal = 0.00;
        var tax = 0.00;
        for(var i = 0; i < this.state.items.length; i++) {
            subtotal += parseFloat(this.state.items[i].amount);
            if (this.state.items[i].tax) {
                tax += this.state.taxRate * parseFloat(this.state.items[i].amount);
            }
        }
        this.setState({
            subtotal: subtotal.toFixed(2),
            tax: tax.toFixed(2),
            total: (subtotal + tax).toFixed(2)
        });
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

    finishOrder() {
        if(this.state.items.length > 0) {
            this.setState({grid: 'quote'});
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

    render() {
        switch(this.state.grid) {
            case 'items':
                return (
                    <Container>
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
                        <br/><br/>
                        <Row>
                            <ItemList
                                items={this.state.items}
                                removeItem={this.removeItem}
                                changePrice={this.changePrice}
                                finishOrder={this.finishOrder}
                                buttonTitle={'Quote'}/>
                        </Row>
                        <br/>
                        <Row>
                            <p>Subtotal: {this.state.subtotal}</p>
                            <p>Tax: {this.state.tax}</p>
                            <p>Total: {this.state.total}</p>
                        </Row>
                    </Container>
                );
            case 'quote':
                return (
                    <Container>
                        <Row>
                            <Col sm='1'><Button color='info' className='text-left hidden-sm' onClick={() => this.setState({grid: 'items'})}>Back</Button></Col>
                            <Col sm='11'></Col>
                        </Row>
                        <PrintQuote
                            date={this.state.date.getMonth() + 1 + '/' + this.state.date.getDate() + '/' + this.state.date.getFullYear()}
                            items={this.state.items}
                            subtotal={this.state.subtotal}
                            tax={this.state.tax}
                            total={this.state.total}
                            printTitle={'Print'}
                            onPrint={() => window.print()}/>
                    </Container>
                );
            default: return (<div></div>);
        }
    }
}