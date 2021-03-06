import React from 'react';
import {BootstrapTable, TableHeaderColumn} from 'react-bootstrap-table';
import GridModal from './GridModal';
import {Container, Row, Col, Button} from 'reactstrap';
import {Modal, ModalBody, ModalFooter} from 'react-bootstrap';
import axios from 'axios';

export default class CustomerGrid extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            customers: [],
            selected: null,
            modal: false,
            flag: true,
            action: null,
            error: false,
            errorMessage: ''
        };
        this.loadCustomers();
        this.loadCustomers = this.loadCustomers.bind(this);
        this.setModal = this.setModal.bind(this);
        this.onSelectCustomer = this.onSelectCustomer.bind(this);
        this.setFlag = this.setFlag.bind(this);
        this.editSelected = this.editSelected.bind(this);
        this.sortByDate = this.sortByDate.bind(this);
    }

    onSelectCustomer(row, isSelected) {
        if (isSelected) {
            this.setState({selected: row});
        }
        else
            this.setState({selected: null});
    }

    setFlag() {
        this.setState({flag: false});
    }

    editSelected(record) {
        this.setState({selected: record});
    }

    setModal() {
        this.setState({modal: false});
    }

    checkSelected() {
        if (this.state.selected) {
            this.setState({
                modal: true,
                flag: true
            });
        }
    }

    loadCustomers() {
        axios.get('/api/customers')
        .then(response => {
            this.setState({customers: response.data});
        })
        .catch(error => {
            var err = error.response.data;
            this.setState({
                error: true,
                errorMessage: err
            });
        });
    }

    sortByDate() {
        axios.get('/api/customers/sortbydate/0')
        .then(response => {
            this.setState({customers: response.data});
        })
        .catch(error => {
            var err = error.response.data;
            this.setState({
                error: true,
                errorMessage: err
            });
        });
    }

    render() {
        return (
            <div>
                <Container>
                    <Row>
                    <Button color='success' onClick={() => {
                        this.sortByDate();
                    }}>Sort By Last Order</Button>
                    </Row>
                    <p></p>
                    <Row>
                        <BootstrapTable 
                            data={this.state.customers} 
                            maxHeight='500px'
                            scrollTop={'Bottom'} 
                            hover
                            condensed
                            selectRow={{
                                mode: 'radio', 
                                clickToSelect: true, 
                                bgColor: 'black',
                                hideSelectColumn: true,
                                onSelect: this.onSelectCustomer
                            }} 
                            containerStyle={{
                                background: '#2F2F2F'
                            }}>
                            <TableHeaderColumn dataField="cid" width='auto' isKey hidden>ID</TableHeaderColumn>
                            <TableHeaderColumn dataField="name" width='auto' columnTitle dataSort filter={{type: 'TextFilter'}}>Name</TableHeaderColumn>
                            <TableHeaderColumn dataField="address" width='auto' columnTitle dataSort filter={{type: 'TextFilter'}}>Address</TableHeaderColumn>
                            <TableHeaderColumn dataField="city" width='auto' columnTitle dataSort filter={{type: 'TextFilter'}}>City</TableHeaderColumn>
                            <TableHeaderColumn dataField="state" width='auto' columnTitle dataSort filter={{type: 'TextFilter'}}>State</TableHeaderColumn>
                            <TableHeaderColumn dataField="zipcode" width='auto' columnTitle dataSort filter={{type: 'TextFilter'}}>Zipcode</TableHeaderColumn>
                            <TableHeaderColumn dataField="phone" width='auto' columnTitle dataSort filter={{type: 'TextFilter'}}>Phone</TableHeaderColumn>
                            <TableHeaderColumn dataField="email" width='auto' columnTitle dataSort filter={{type: 'TextFilter'}}>Email</TableHeaderColumn>
                        </BootstrapTable>
                    </Row>
                    <p></p>
                    <Row>
                        <Col>
                            <Button color='success' onClick={() => {
                                this.setState({
                                    action: 'add',
                                    modal: true,
                                    flag: true
                                });
                            }}>Add</Button>
                            {' '}<Button color='warning' onClick={() => {
                                this.checkSelected(); 
                                this.setState({action: 'edit'});
                            }}>Edit</Button>
                        </Col>
                    </Row>
                    <p></p>
                    <div className={!this.props.extra ? 'hidden' : ''}>
                        <Button color='info' onClick={() => this.props.extraFunction(this.state.selected)}>{this.props.extraTitle}</Button>
                    </div>
                </Container>
                <GridModal 
                    url='/api/customers'
                    record={this.state.action === 'add' ? {
                            name: '',
                            address: '',
                            city: '',
                            state: '',
                            zipcode: '',
                            phone: '',
                            email: ''
                        } : this.state.selected ? {
                                name: this.state.selected.name,
                                address: this.state.selected.address,
                                city: this.state.selected.city,
                                state: this.state.selected.state,
                                zipcode: this.state.selected.zipcode,
                                phone: this.state.selected.phone,
                                email: this.state.selected.email
                            } : {}
                    }
                    id={this.state.selected ? JSON.parse(JSON.stringify(this.state.selected)).cid : null}
                    modal={this.state.modal}
                    setModal={this.setModal}
                    loadRecords={this.loadCustomers}
                    action={this.state.action}
                    setFlag={this.setFlag}
                    flag={this.state.flag}
                    editSelected={this.editSelected}
                    editable={{
                        name: true,
                        address: true,
                        city: true,
                        state: true,
                        zipcode: true,
                        phone: true,
                        email: true
                    }}
                    validateInput={
                            (scope, event) => {
                                var temp = JSON.parse(JSON.stringify(scope.state.record));
                                temp[event.target.id] = event.target.value;
                                scope.setState({record: temp});
                            }
                    }
                    onSave={
                        (scope) => {
                            var temp = JSON.parse(JSON.stringify(scope.state.record));
                            if(scope.props.action === 'add'){
                                axios.post(scope.props.url, temp)
                                .then(response => {
                                    scope.props.setModal();
                                    scope.props.loadRecords();
                                })
                                .catch(error => {
                                    scope.props.setModal();
                                    var err = error.response.data;
                                    this.setState({error: true, errorMessage: err});
                                });
                            }
                            else{
                                axios.put(scope.props.url + '/' + scope.props.id, temp)
                                .then(response => {
                                    scope.props.editSelected(response.data);
                                    scope.props.setModal();
                                    scope.props.loadRecords();
                                })
                                .catch(error => {
                                    scope.props.setModal();
                                    var err = error.response.data;
                                    this.setState({error: true, errorMessage: err});
                                });
                            }
                        }
                    }/>
                    <Modal show={this.state.error} onHide={() => this.setState({error: false, errorMessage: ''})}>
                        <ModalBody>
                            <h1>{this.state.errorMessage}</h1>
                        </ModalBody>
                        <ModalFooter>
                            <Button onClick={() => this.setState({error: false, errorMessage: ''})}>OK</Button>
                        </ModalFooter>
                    </Modal>
            </div>
        );
    }
}