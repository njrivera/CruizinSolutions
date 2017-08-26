import React from 'react';
import {BootstrapTable, TableHeaderColumn} from 'react-bootstrap-table';
import GridModal from './GridModal';
import {Container, Row, Col, Button} from 'reactstrap';
import {Modal, ModalBody, ModalFooter} from 'react-bootstrap';
import axios from 'axios';

export default class PackageGrid extends React.Component {
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
        this.deleteSelected = this.deleteSelected.bind(this);
        this.setModal = this.setModal.bind(this);
        this.onSelectRecord = this.onSelectRecord.bind(this);
        this.setFlag = this.setFlag.bind(this);
        this.editSelected = this.editSelected.bind(this);
        this.setError = this.setError.bind(this);
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
            this.setState({
                modal: true,
                flag: true});
        }
    }

    setError(err) {
        this.setState({
            error: true,
            errorMessage: err
        });
    }

    loadRecords() {
        axios.get('/api/packages')
        .then(response => {
            this.setState({records: response.data});
        })
        .catch(error => {
            this.setError(error.response.data);
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
                            <TableHeaderColumn dataField="itemnum" width='auto' isKey hidden>ID</TableHeaderColumn>
                            <TableHeaderColumn dataField="description" width='auto' columnTitle dataSort filter={{type: 'TextFilter'}}>Description</TableHeaderColumn>
                            <TableHeaderColumn dataField="price" width='auto' dataSort filter={{type: 'TextFilter'}}>Price</TableHeaderColumn>
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
                            {' '}<Button color='danger' onClick={() => {
                                this.checkSelected();
                                this.setState({action: 'delete'});
                            }}>Delete</Button>
                        </Col>
                    </Row>
                    <p></p>
                    <div className={!this.props.extra ? 'hidden' : ''}>
                        <Button color='info' onClick={() => this.props.extraFunction(this.state.selected)}>{this.props.extraTitle}</Button>
                    </div>
                </Container>
                <GridModal 
                    url='/api/packages'
                    record={this.state.action === 'add' ? {
                            description: '',
                            price: ''
                        } : this.state.selected ? {
                                description: this.state.selected.description,
                                price: this.state.selected.price,
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
                    editable={{
                        price: true
                    }}
                    setError={this.setError}
                    validateInput={
                            (scope, event) => {
                                if(event.target.id === 'price') {
                                    var val = event.target.value;
                                    val = val.replace('.', '');
                                    val = parseInt(val, 10).toString();
                                    if(val === 'NaN') val = '0.00';
                                    else
                                        switch(val.length) {
                                            case 0: val = '0.00';
                                            break;
                                            case 1: val = '0.0' + val;
                                            break;
                                            case 2: val = '0.' + val;
                                            break;
                                            default: val = val.substring(0, val.length - 2) + '.' + val.substring(val.length - 2)
                                        }
                                    event.target.value = val;
                                }
                                var temp = JSON.parse(JSON.stringify(scope.state.record));
                                temp[event.target.id] = event.target.value;
                                scope.setState({record: temp});
                            }
                    }
                    onSave={
                        (scope) => {
                            var temp = JSON.parse(JSON.stringify(scope.state.record));
                            if(temp.price === '') temp.price = '0.00';
                            if(scope.props.action === 'add'){
                                axios.post(scope.props.url, temp)
                                .then(response => {
                                    scope.props.setModal();
                                    scope.props.loadRecords();
                                })
                                .catch(error => {
                                    scope.props.setModal();
                                    this.setError(error.response.data);
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
                                    this.setError(error.response.data);
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