import React from 'react';
import {BootstrapTable, TableHeaderColumn} from 'react-bootstrap-table';
import {Container, Row, Col, Button} from 'reactstrap';
import {Modal, ModalBody, ModalFooter} from 'react-bootstrap';
import axios from 'axios';

export default class CustVehicleGrid extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            vehicles: [],
            selected: null,
            error: false,
            errorMessage: ''
        };
        this.loadVehicles(this.props.id);
        this.loadVehicles = this.loadVehicles.bind(this);
        this.onSelectVehicle = this.onSelectVehicle.bind(this);
    }

    onSelectVehicle(row, isSelected) {
        if (isSelected) {
            this.setState({selected: row});
        }
        else
            this.setState({selected: null});
    }

    checkSelected() {
        if (this.state.selected) {
            this.setState({modal: true});
            this.setState({flag: true});
        }
    }

    loadVehicles(id) {
        axios.get('/api/orders/vehicles/' + id)
        .then(response => {
            this.setState({vehicles: response.data});
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
                            data={this.state.vehicles} 
                            maxHeight='500px'
                            scrollTop={'Bottom'} 
                            hover
                            condensed
                            selectRow={{
                                mode: 'radio', 
                                clickToSelect: true, 
                                bgColor: 'black',
                                hideSelectColumn: true,
                                onSelect: this.onSelectVehicle
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
                            <Button color='info' onClick={() => this.props.onChoose(this.state.selected)}>Choose Vehicle</Button>
                            <p></p>
                            <Button color='info' onClick={() => this.props.onAdd()}>Add Vehicle</Button>
                        </Col>
                    </Row>
                </Container>
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