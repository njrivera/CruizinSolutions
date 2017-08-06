import React from 'react';
import '../styles/Layout.css'
import {Container, Row, Col} from 'reactstrap';
import CustomerGrid from './CustomerGrid'

export default class Layout extends React.Component {

    render() {
        return (
            <div className='Layout'>
                <Container>
                    <Row>
                        <Col sm='1'>Navbar</Col>
                        <Col sm='11'>
                            <Row>
                                Main Content
                            </Row>
                            <Row>
                                <CustomerGrid />
                            </Row>
                        </Col>
                    </Row>
                </Container>
            </div>
        );
    }
}