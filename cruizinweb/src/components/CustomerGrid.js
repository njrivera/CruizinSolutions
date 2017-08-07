import React from 'react';
import {BootstrapTable, TableHeaderColumn} from 'react-bootstrap-table';
import GridModal from './GridModal';
import {Container, Row, Col, Button} from 'reactstrap';
import axios from 'axios';

export default class CustomerGrid extends React.Component {
    constructor() {
        super();
        this.state = {
            customers: [],
            modal: false,
            id: null
        };
        this.loadCustomers();
        this.onSelectRec = this.onSelectRec.bind(this);
        this.toggle = this.toggle.bind(this);
    }

    onSelectRec(row, isSelected) {
        if (isSelected) {
            this.selected = row;
            this.customer = row;
        }
        else
            this.selected = null;
    }

    loadCustomers() {
        axios.get('/customers')
        .then(response => {
            this.setState({customers: response.data});
        })
        .catch(error => {
            console.log(error);
        });
    }

    toggle(act) {
        this.loadCustomers();
        if(act === 'deleted') {
            this.selected = null;
            this.customer = null;
        }
        if(act === 'edit' || act === 'delete') {
            if (!this.selected)
                return;
            this.customer = this.selected;
        }
        if(act === 'add')
            this.customer = {
                cid: '',
                name: '',
                address: '',
                city: '',
                state: '',
                zipcode: '',
                phone: ''
            };
        this.setState({modal: !this.state.modal});
        if(act === 'done') {
            this.customer = null;
            return;
        }
        this.setState({option: act});
        this.setState({id: this.customer.cid})
    }

    render() {
        return (
            <div>
                <Container>
                    <Row>
                        <BootstrapTable 
                            data={this.state.customers} 
                            maxHeight='500px'
                            scrollTop={'Bottom'} 
                            hover
                            selectRow={{
                                mode: 'radio', 
                                clickToSelect: true, 
                                bgColor: 'black',
                                hideSelectColumn: true,
                                onSelect: this.onSelectRec
                            }} 
                            containerStyle={{
                                background: '#2F2F2F'
                            }}>
                            <TableHeaderColumn dataField="cid" width='auto' isKey hidden>ID</TableHeaderColumn>
                            <TableHeaderColumn dataField="name" width='auto' dataSort filter={{type: 'TextFilter'}}>Name</TableHeaderColumn>
                            <TableHeaderColumn dataField="address" width='auto' dataSort filter={{type: 'TextFilter'}}>Address</TableHeaderColumn>
                            <TableHeaderColumn dataField="city" width='auto' dataSort filter={{type: 'TextFilter'}}>City</TableHeaderColumn>
                            <TableHeaderColumn dataField="state" width='auto' dataSort filter={{type: 'TextFilter'}}>State</TableHeaderColumn>
                            <TableHeaderColumn dataField="zipcode" width='auto' dataSort filter={{type: 'TextFilter'}}>Zipcode</TableHeaderColumn>
                            <TableHeaderColumn dataField="phone" width='auto' dataSort filter={{type: 'TextFilter'}}>Phone</TableHeaderColumn>
                        </BootstrapTable>
                    </Row>
                    <Row>
                        <Col>
                            <Button color='success' onClick={() => this.toggle('add')}>Add</Button>
                            <Button color='info' onClick={() => this.toggle('edit')}>Edit</Button>
                            <Button color='danger' onClick={() => this.toggle('delete')}>Delete</Button>
                        </Col>
                    </Row>
                </Container>
                <GridModal 
                    toggle={this.toggle}
                    data={this.customer}
                    id={this.state.id}
                    action={this.state.option} 
                    modal={this.state.modal} 
                    url='/customers'/>
            </div>
        );
    }
}