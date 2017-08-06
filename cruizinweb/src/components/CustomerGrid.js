import React from 'react';
import {BootstrapTable, TableHeaderColumn} from 'react-bootstrap-table';
import {Container, Row, Col, Button, Modal, ModalFooter, ModalBody, ModalHeader} from 'reactstrap';
import axios from 'axios';

export default class CustomerGrid extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            customer: [],
            modal: false
        };
        var selected;
        const scope = this;
        axios.get('/customers')
        .then(response => {
            scope.setState({customers: response.data});
        })
        .catch(error => {
            console.log(error);
        });

        this.toggle = this.toggle.bind(this)
    }

    onSelectRec(row, isSelected) {
        if (isSelected)
            this.selected = row;
    }

    toggle() {
        this.setState({modal: !this.state.modal});
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
                            <Button color='success' onClick={this.toggle}>Add</Button>
                            <Button color='info'>Edit</Button>
                            <Button color='danger'>Delete</Button>
                        </Col>
                    </Row>
                </Container>

                <Modal isOpen={this.state.modal} toggle={this.toggle}>
                    <ModalHeader>sdf</ModalHeader>
                    <ModalBody>asdfasdf</ModalBody>
                    <ModalFooter><Button>sdfsdf</Button></ModalFooter>
                </Modal>
            </div>
        );
    }
}