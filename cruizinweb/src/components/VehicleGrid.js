import React from 'react';
import {BootstrapTable, TableHeaderColumn} from 'react-bootstrap-table';
import GridModal from './GridModal';
import {Container, Row, Col, Button} from 'reactstrap';
import axios from 'axios';

export default class VehicleGrid extends React.Component {
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
        axios.get('/api/vehicles')
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
                            <TableHeaderColumn dataField="vid" width='auto' isKey hidden>ID</TableHeaderColumn>
                            <TableHeaderColumn dataField="year" width='auto' dataSort filter={{type: 'TextFilter'}}>Year</TableHeaderColumn>
                            <TableHeaderColumn dataField="make" width='auto' dataSort filter={{type: 'TextFilter'}}>Make</TableHeaderColumn>
                            <TableHeaderColumn dataField="model" width='auto' dataSort filter={{type: 'TextFilter'}}>Model</TableHeaderColumn>
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
                                switch(event.target.id) {
                                    case 'year':
                                        if(event.target.value.length === 0)
                                            break;
                                        if(event.target.value.length > 4 || !Number(event.target.value)) {
                                            event.target.value = event.target.value.slice(1);
                                            return;
                                        }
                                        event.target.value = parseInt(event.target.value);
                                    default:
                                }
                                var temp = JSON.parse(JSON.stringify(scope.state.record));
                                temp[event.target.id] = event.target.value;
                                scope.setState({record: temp});
                            }
                    }
                    onSave={
                        (scope) => {
                            if(scope.state.record.year.toString().length === 4){
                                var temp = JSON.parse(JSON.stringify(scope.state.record));
                                temp.year = parseInt(temp.year);
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
                        }
                    }/>
            </div>
        );
    }
}