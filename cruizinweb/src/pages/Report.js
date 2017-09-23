import React from 'react';
import {Container, Row, Col, Button} from 'reactstrap';
import {FormControl, Modal, ModalBody, ModalFooter} from 'react-bootstrap';
import axios from 'axios';

export default class Report extends React.Component {
    constructor() {
        super();
        this.state = {
            month: '1',
            year: (new Date()).getFullYear().toString(),
            totalTires: '',
            totalTax: '',
            error: false,
            errorMessage: ''
        };

        this.generateReport = this.generateReport.bind(this);
    }

    generateReport() {
        axios.get('/api/reports/newtiretax/' + this.state.month + '/' + this.state.year)
        .then(response => {
            var report = response.data;
            report.Tax = '0.00';
            this.setState({
                totalTires: report.Qty.toString(),
                totalTax: report.Tax
            })
        })
        .catch(error => {
            debugger;
            var err = error.response.data;
            this.setState({
                error: true,
                errorMessage: err
            });
        });
    }

    render() {
        var scope = this;
        return (
            <Container>
                <Row>
                <Col sm='4'></Col>
                <Col sm='2'>
                    <h4>Month</h4>
                    <FormControl
                        componentClass='select' 
                        onChange={(event) => {scope.setState({month: event.target.value})}}>
                            <option value="1">January</option>
                            <option value="2">Febuary</option>
                            <option value="3">March</option>
                            <option value="4">April</option>
                            <option value="5">May</option>
                            <option value="6">June</option>
                            <option value="7">July</option>
                            <option value="8">August</option>
                            <option value="9">September</option>
                            <option value="10">October</option>
                            <option value="11">November</option>
                            <option value="12">December</option>
                    </FormControl>
                </Col>
                <Col sm='2'>
                    <h4>Year</h4>
                    <FormControl
                        value={this.state.year}
                        onChange={(event) => {scope.setState({year: event.target.value})}}>
                    </FormControl>
                </Col>
                <Col sm='4'></Col>
                </Row>
                <br/>
                <Row>
                    <Button color='info' onClick={() => scope.generateReport()}>Generate Report</Button>
                </Row>
                <br/>
                <Row>
                    <Col sm='2'></Col>
                    <Col sm='3'>
                        <h3>TOTAL TIRES</h3>
                        <br/>
                        <h4>{this.state.totalTires}</h4>
                    </Col>
                    <Col sm='2'></Col>
                    <Col sm='3'>
                        <h3>TOTAL TAX</h3>
                        <br/>
                        <h4>{this.state.totalTax}</h4>
                    </Col>
                    <Col sm='2'></Col>
                </Row>
                <Modal show={this.state.error} onHide={() => this.setState({error: false, errorMessage: ''})}>
                    <ModalBody>
                        <h1>{this.state.errorMessage}</h1>
                    </ModalBody>
                    <ModalFooter>
                        <Button onClick={() => this.setState({error: false, errorMessage: ''})}>OK</Button>
                    </ModalFooter>
                </Modal>
            </Container>
        );
    }
}