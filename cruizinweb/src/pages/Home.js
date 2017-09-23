import React from 'react';
import {Container, Row, Col} from 'reactstrap';

export default class Home extends React.Component {
    render() {
        return (
            <Container>
                <h1>WELCOME</h1>
                <h2 className='text-left'>THINGS TO CONSIDER:</h2>
                <br/>
                <Row>
                    <Col sm='1'></Col>
                    <Col sm='11'>
                        <Row>
                            <Col sm='1'><h3 className='text-right'>---</h3></Col>
                            <Col sm='11'>
                                <h3 className='text-left'>DELETING an item (tire, rim, part, service, or package) WILL NOT affect past invoices.</h3>
                                <br/><br/>
                            </Col>
                            <Col sm='1'><h3 className='text-right'>---</h3></Col>
                            <Col sm='11'>
                                <h3 className='text-left'>EDITING a customer or vehicle WILL affect past invoices.</h3>
                                <br/><br/>
                            </Col>
                            <Col sm='1'><h3 className='text-right'>---</h3></Col>
                            <Col sm='11'>
                                <h3 className='text-left'>TAX:</h3>
                            </Col>
                            <Col sm='6'>     
                                <h3>TAXABLE</h3>
                                <br/>
                                <h4>Anything NEW</h4>
                                <h4>Packages</h4>
                            </Col>
                            <Col sm='6'>
                                <h3>NON-TAXABLE</h3>
                                <br/>
                                <h4>Anything USED</h4>
                                <h4>Services</h4>
                            </Col>
                        </Row>
                    </Col>
                </Row>
            </Container>
        );
    }
}