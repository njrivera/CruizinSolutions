import React from 'react';
import {BootstrapTable, TableHeaderColumn} from 'react-bootstrap-table';
import GridModal from './GridModal';
import {Container, Row, Col, Button} from 'reactstrap';
import axios from 'axios';

export default class TireGrid extends React.Component {
    constructor() {
        super();
        this.state = {
            records: [],
            selected: null,
            modal: false,
            flag: true,
            action: null
        };
        this.loadRecords();
        this.loadRecords = this.loadRecords.bind(this);
        this.deleteSelected = this.deleteSelected.bind(this);
        this.setModal = this.setModal.bind(this);
        this.onSelectRecord = this.onSelectRecord.bind(this);
        this.setFlag = this.setFlag.bind(this);
        this.editSelected = this.editSelected.bind(this);
    }

    onSelectRecord(row, isSelected) {
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

    deleteSelected() {
        this.setState({selected: null});
    }

    setModal() {
        this.setState({modal: false});
    }

    checkSelected() {
        if (this.state.selected) {
            this.setState({modal: true});
            this.setState({flag: true});
        }
    }

    loadRecords() {
        axios.get('/api/tires')
        .then(response => {
            this.setState({records: response.data});
        })
        .catch(error => {
            console.log(error);
        });
    }

    render() {
        return (
            <div>
                <Container>
                    <Row>
                        <BootstrapTable 
                            data={this.state.records} 
                            maxHeight='500px'
                            scrollTop={'Bottom'} 
                            hover
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
                            <TableHeaderColumn dataField="itemnum" width='auto' isKey hidden>ID</TableHeaderColumn>
                            <TableHeaderColumn dataField="brand" width='auto' dataSort filter={{type: 'TextFilter'}}>Brand</TableHeaderColumn>
                            <TableHeaderColumn dataField="model" width='auto' dataSort filter={{type: 'TextFilter'}}>Model</TableHeaderColumn>
                            <TableHeaderColumn dataField="size" width='auto' dataSort filter={{type: 'TextFilter'}}>Size</TableHeaderColumn>
                            <TableHeaderColumn dataField="servicedesc" width='auto' dataSort filter={{type: 'TextFilter'}}>Service Description</TableHeaderColumn>
                            <TableHeaderColumn dataField="condition" width='auto' dataSort filter={{type: 'TextFilter'}}>Condition</TableHeaderColumn>
                            <TableHeaderColumn dataField="price" width='auto' dataSort filter={{type: 'TextFilter'}}>Price</TableHeaderColumn>
                            <TableHeaderColumn dataField="qty" width='auto' dataSort filter={{type: 'TextFilter'}}>Qty</TableHeaderColumn>
                        </BootstrapTable>
                    </Row>
                    <Row>
                        <Col>
                            <Button color='success' onClick={() => {this.setState({action: 'add'}), this.setState({modal: true}), this.setState({flag: true})}}>Add</Button>
                            <Button color='info' onClick={() => {this.checkSelected(), this.setState({action: 'edit'})}}>Edit</Button>
                            <Button color='danger' onClick={() => {this.checkSelected(), this.setState({action: 'delete'})}}>Delete</Button>
                        </Col>
                    </Row>
                    <div className={!this.props.extra ? 'hidden' : ''}>
                        <Button onClick={() => this.props.extraFunction(this.state.selected)}>{this.props.extraTitle}</Button>
                    </div>
                </Container>
                <GridModal 
                    url='/api/tires'
                    record={this.state.action === 'add' ? {
                            brand: '',
                            model: '',
                            size: '',
                            servicedesc: '',
                            condition: '',
                            price: '',
                            qty: ''
                        } : this.state.selected ? {
                                brand: this.state.selected.brand,
                                model: this.state.selected.model,
                                size: this.state.selected.size,
                                servicedesc: this.state.selected.servicedesc,
                                condition: this.state.selected.condition,
                                price: this.state.selected.price,
                                qty: this.state.selected.qty
                            } : {}
                    }
                    id={this.state.selected ? JSON.parse(JSON.stringify(this.state.selected)).itemnum : null}
                    deleteRecord={this.deleteSelected}
                    modal={this.state.modal}
                    setModal={this.setModal}
                    loadRecords={this.loadRecords}
                    action={this.state.action}
                    setFlag={this.setFlag}
                    flag={this.state.flag}
                    editSelected={this.editSelected}
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
                            temp.price = parseFloat(temp.price);
                            temp.qty = parseInt(temp.qty);
                            if(scope.props.action === 'add'){
                                axios.post(scope.props.url, temp)
                                .then(response => {
                                    scope.props.setModal();
                                    scope.props.loadRecords();
                                })
                                .catch(error => {
                                    console.log(error);
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
                                    console.log(error);
                                });
                            }
                        }
                    }/>
            </div>
        );
    }
}