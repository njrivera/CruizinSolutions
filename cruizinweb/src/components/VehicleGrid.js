import React from 'react';
import {BootstrapTable, TableHeaderColumn} from 'react-bootstrap-table';
import GridModal from './GridModal';
import {Container, Row, Col, Button} from 'reactstrap';
import {Modal, ModalBody, ModalFooter} from 'react-bootstrap';
import axios from 'axios';

export default class VehicleGrid extends React.Component {
    constructor() {
        super();
        this.state = {
            records: [],
            selected: null,
            modal: false,
            flag: true,
            action: null,
            error: false,
            errorMessage: ''
        };
        this.loadRecords();
        this.loadRecords = this.loadRecords.bind(this);
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

    loadRecords() {
        axios.get('/api/vehicles')
        .then(response => {
            this.setState({records: response.data});
        })
        .catch(error => {
            var err = error.response.data;
            this.setState({error: true, errorMessage: err});
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
                            <TableHeaderColumn dataField="vid" width='auto' isKey hidden>ID</TableHeaderColumn>
                            <TableHeaderColumn dataField="year" width='auto' dataSort filter={{type: 'TextFilter'}}>Year</TableHeaderColumn>
                            <TableHeaderColumn dataField="make" width='auto' columnTitle dataSort filter={{type: 'TextFilter'}}>Make</TableHeaderColumn>
                            <TableHeaderColumn dataField="model" width='auto' columnTitle dataSort filter={{type: 'TextFilter'}}>Model</TableHeaderColumn>
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
                    url='/api/vehicles'
                    record={this.state.action === 'add' ? {
                            year: '',
                            make: '',
                            model: ''
                        } : this.state.selected ? {
                                year: this.state.selected.year,
                                make: this.state.selected.make,
                                model: this.state.selected.model
                            } : {}
                    }
                    id={this.state.selected ? JSON.parse(JSON.stringify(this.state.selected)).vid : null}
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