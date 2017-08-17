import React from 'react';
import CustomerGrid from '../components/CustomerGrid';
import CustVehiclesGrid from '../components/CustVehiclesGrid';
import VehicleGrid from '../components/VehicleGrid';
import {DropdownButton, MenuItem, FormGroup, FormControl, ControlLabel} from 'react-bootstrap';
import OrderOptions from '../components/OrderOptions';
import ItemList from '../components/ItemList';
import Invoice from '../components/Invoice';
import {Container, Row, Col, Button} from 'reactstrap';
import {BootstrapTable, TableHeaderColumn} from 'react-bootstrap-table';
import axios from 'axios';


export default class Order extends React.Component {
    constructor() {
        super();
        this.state = {
            grid: 'customer',
            orders: [],
            gridOrders: [],
            customer: null,
            selected: null,
            items: []
        };
        this.chooseCustomer = this.chooseCustomer.bind(this);
        this.chooseOrder = this.chooseOrder.bind(this);
        this.onBack = this.onBack.bind(this);
        this.onSelectRecord = this.onSelectRecord.bind(this);
    }

    onBack() {
        switch(this.state.grid) {
            case 'orders': 
                this.setState({
                    gridOrders: [],
                    orders: [],
                    grid: 'customer',
                    customer: null,
                    selected: null
                });
            break;
            case 'invoice': 
                this.setState({
                    grid: 'orders',
                    items: [],
                    selected: null
                });
            break;
            default: break;
        }
    }

    chooseCustomer(customer) {
        if(customer) {
            axios.get('/api/orders/' + customer.cid)
            .then(response => {
                var temp = this.state.gridOrders;
                for(var i = 0; i < response.data.length; i++) {
                    var order = response.data[i];
                    temp.push({
                        ordernum: order.ordernum,
                        vid: order.vid,
                        vehicle: order.year + ' ' + order.make + ' ' + order.model,
                        odometer: order.odometer,
                        date: order.date,
                        comments: order.comments,
                        subtotal: order.subtotal,
                        tax: order.tax,
                        total: order.total,
                    })
                }
                this.setState({
                    gridOrders: temp,
                    orders: response.data,
                    customer: customer,
                    grid: 'orders'});
            })
            .catch(error => {
                console.log(error);
            });  
        }
    }

    onSelectRecord(row, isSelected) {
        if (isSelected) {
            this.setState({selected: row});
        }
        else
            this.setState({selected: null});
    }

    chooseOrder() {
        if(this.state.selected) {
            axios.get('/api/orders/items/' + this.state.selected.ordernum)
            .then(response => {
                var vehicle;
                for(var i = 0; i < this.state.orders.length; i++) {
                    if (this.state.selected.ordernum === this.state.orders[i].ordernum) {
                        this.setState({vehicle: {
                            vid: this.state.orders[i].vid,
                            year: this.state.orders[i].year,
                            make: this.state.orders[i].make,
                            model: this.state.orders[i].model
                        }});
                        break;
                    }
                }
                this.setState({
                    items: response.data,
                    grid: 'invoice'});
            })
            .catch(error => {
                console.log(error);
            });  
        }
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
            case 'orders':
                return (
                    <Container>
                        <Row>
                            <Col sm='1'><Button className='text-left' onClick={this.onBack}>Back</Button></Col>
                            <Col sm='11'></Col>
                        </Row>
                        <h1>Choose Invoice</h1>
                        <Row>
                            <BootstrapTable 
                                data={this.state.gridOrders} 
                                maxHeight='500px'
                                scrollTop={'Bottom'} 
                                hover
                                condensed
                                selectRow={{
                                    mode: 'radio', 
                                    clickToSelect: true, 
                                    bgColor: 'black',
                                    hideSelectColumn: true,
                                    onSelect: this.onSelectRecord
                                }} 
                                containerStyle={{
                                    background: '#2F2F2F'
                                }}>
                                <TableHeaderColumn dataField="ordernum" width='auto' isKey>Order #</TableHeaderColumn>
                                <TableHeaderColumn dataField="vehicle" width='auto' columnTitle dataSort filter={{type: 'TextFilter'}}>Vehicle</TableHeaderColumn>
                                <TableHeaderColumn dataField="odometer" width='auto' columnTitle dataSort filter={{type: 'TextFilter'}}>Odometer</TableHeaderColumn>
                                <TableHeaderColumn dataField="date" width='auto' columnTitle dataSort filter={{type: 'TextFilter'}}>Date</TableHeaderColumn>
                                <TableHeaderColumn dataField="total" width='auto' dataSort filter={{type: 'TextFilter'}}>Total Paid</TableHeaderColumn>
                            </BootstrapTable>
                        </Row>
                        <p></p>
                        <Button onClick={() => this.chooseOrder(this.state.selected)}>Show Invoice</Button>
                    </Container>
                );
            case 'invoice':
                return (
                    <Container>
                        <Row>
                            <Col sm='1'><Button className='text-left' onClick={this.onBack}>Back</Button></Col>
                            <Col sm='11'></Col>
                        </Row>
                        <Invoice
                            invoiceNum={this.state.selected.ordernum}
                            date={this.state.selected.date}
                            customer={this.state.customer}
                            vehicle={this.state.vehicle}
                            items={this.state.items}
                            subtotal={this.state.selected.subtotal}
                            tax={this.state.selected.tax}
                            total={this.state.selected.total}
                            odometer={this.state.selected.odometer}
                            comments={this.state.selected.comments}
                            onPrint={() => window.print()}
                            printTitle={'Print'}/>
                    </Container>
                );
        }
    }
}