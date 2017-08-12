import React from 'react';
import {Button, Col, Row, Container} from 'reactstrap';
import {Modal, ModalBody, ModalFooter, ModalHeader, Form, FormGroup, FormControl} from 'react-bootstrap';

export default class Invoice extends React.Component {
    constructor(props) {
        super(props);
        this.state = {

        };
    }

    render() {
        return (
            <div>
                <Modal show={this.props.finished}>
                    <ModalBody>
                <Container>
                    <Row>
                        <Col>Logo</Col>
                        <Col>Other Logo</Col>
                    </Row>
                    <Row>
                        <Col>Extra Info</Col>
                        <Col>Invoice # and Date</Col>
                    </Row>
                    <Row>
                        <Col>Customer</Col>
                        <Col>Vehicle</Col>
                    </Row>
                    <Row>
                        <Col sm='1'>Qty</Col>
                        <Col sm='7'>Description</Col>
                        <Col sm='2'>Price Per</Col>
                        <Col sm='2'>Amount</Col>
                    </Row>
                    <Row>
                        <Col sm='8'>Comments</Col>
                        <Col sm='4'>Subtotal/tax/total</Col>
                    </Row>
                    <Row>
                        Signature
                        <button onClick={() => window.print()}>Print</button>
                    </Row>
                </Container>
                </ModalBody>
                </Modal>
            </div>
        );
    }
}