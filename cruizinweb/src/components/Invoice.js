import React from 'react';
import {Button, Col, Row, Container, ListGroup, ListGroupItem} from 'reactstrap';
import {Modal, ModalBody, ModalFooter, ModalHeader, Form, FormGroup, FormControl, ControlLabel} from 'react-bootstrap';
import tire_icon from '../styles/tire_icon.jpg';
import {BootstrapTable, TableHeaderColumn} from 'react-bootstrap-table';
import caret from '../styles/fa-caret-right.png';

export default class Invoice extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            comments: ''
        };
    }

    render() {
        var scope = this;
        return (
            <div>
                <Container>
                    <Row>
                        <Col sm='6' className='text-left'>
                            <Row><h3>YOUR ULTIMATE DRIVING EXPERIENCE</h3></Row>
                            <Row><img style={{width: 5, height: 5}} src={caret}/> NEW {'&'} USED TIRES ALL BRANDS</Row>
                            <Row><img style={{width: 5, height: 5}} src={caret}/> COMPUTER BALANCING</Row>
                            <Row><img style={{width: 5, height: 5}} src={caret}/> WHEELS {'&'} TIRES</Row>
                            <Row><img style={{width: 5, height: 5}} src={caret}/> LAY-AWAY AVAILABLE</Row>
                            <Row><img style={{width: 5, height: 5}} src={caret}/> CAR ACCESSORIES</Row>
                        </Col>
                        <Col sm='6' className='text-right'>
                            <br/>
                            <h4>6901 SHELDON ROAD | TAMPA, FL 33615</h4>
                            <h4>(813) 886-8072 | FAX: (813) 866-8092)</h4>
                            <h4>cruzinsolutions@juno.com</h4>
                        </Col>
                    </Row>
                    <br/><br/>
                    <Row>
                        <Col sm='6' className='text-left'>
                            <Row>{this.props.customer.name}</Row>
                            <Row>{this.props.customer.address}</Row>
                            <Row>{this.props.customer.city}, {this.props.customer.state} {this.props.customer.zipcode}</Row>
                            <Row>{this.props.customer.phone}</Row>
                        </Col>
                        <Col sm='6' className='text-right'>
                            <Row>Date of Order: {this.props.date}</Row>
                            <br/>
                            <Row>{this.props.vehicle.year} {this.props.vehicle.make} {this.props.vehicle.model}</Row>
                            <Row>Odometer: {this.props.odometer}</Row>
                        </Col>
                    </Row>
                    <br/><br/>
                    <BootstrapTable 
                        data={this.props.items} 
                        condensed
                        containerStyle={{
                            background: '#2F2F2F'
                        }}>
                        <TableHeaderColumn dataField="itemnum" width='60' isKey>Item #</TableHeaderColumn>
                        <TableHeaderColumn dataField="description" width='auto'>Description</TableHeaderColumn>
                        <TableHeaderColumn dataField="price" width='100'>Price Per</TableHeaderColumn>
                        <TableHeaderColumn dataField="qty" width='60'>Qty</TableHeaderColumn>
                        <TableHeaderColumn dataField="amount" width='100'>Amount</TableHeaderColumn>
                    </BootstrapTable>
                    <br/><br/>
                    <Row>
                        <Col sm='9' className='text-left'>
                            <Row><h4>Comments:</h4></Row>
                            <Row>
                                <Col sm='1'></Col>
                                <Col sm='11'><p>{this.props.comments}</p></Col>
                            </Row>
                        </Col>
                        <Col sm='3' className='text-right'>
                            <p>Subtotal: {this.props.subtotal}</p>
                            <p>Tax: {this.props.tax}</p>
                            <p>Total: {this.props.total}</p>
                        </Col>
                    </Row>
                    <br/><br/>
                    <Row className='text-left'>
                        Signature:   _______________________________________________
                    </Row>
                    <button className='hidden-sm' onClick={() => window.print()}>Print</button>
                    <br/><br/><br/>
                    <Row className='text-left'>
                        <Col sm='6'>
                            <p>- Always inflate, rotate and balance your tires at intervals of 7,000 miles. Improper care can affect the warranty on new tires.</p>
                            <p>- No warranty on low profile tires: 25, 30, 35, 40 series</p>
                            <p>- Customer must retighten lug nuts after 25 miles</p>
                            <p>- All wheel sales are final when mounted.</p>
                            <p>- No refund on deposit</p>
                            <p>- No refund on special orders</p>
                        </Col>
                        <Col sm='6'>
                            <p>- No warranty on used tires</p>
                            <p>- 25% Re-Stock fee on all returned goods</p>
                            <p>- No warrnaty on spacers, adapters or any vehicle alterations</p>
                            <p>- We are not responsible for any alterations.</p>
                            <p>- No cash refund on returned goods - only store credit given</p>
                            <p>- No refund on Lay-Away left for more than 90 days</p>
                            <p>- No refund on any Lay-Away special orders</p>
                        </Col>
                    </Row>
                    <Row><h4>TERMS {'&'} CONDITIONS</h4></Row>
                    <br/>
                    <Row className='text-left'>
                        <Col sm='6'>
                            <p>MOUNTING: FREE MOUNTING AT TIME OF PURCHASE ONLY</p>
                            <p>FLAT REPAIR: FREE FLAT REPAIR AVAILABLE ONY WITH PROOF OF PURCHASE ON NEW TIRES: DOES NOT INCLUDE SHOULDER OR SIDEWALL REPAIR, OR ANY TYPE OF DAMAGE EXCEEDING A PLUG REPAIR</p>
                            <p>LIFE-TIME ROTATION: FREE ROTATION IS AVAILABLE WITH PROOF OF PURCHASE RECEIPTS</p>    
                        </Col>
                        <Col sm='6'>
                            <p>BALANCING: FREE BALANCING AT TIME OF PURCHASE ONLY UNLESS OTHER ARRANGEMENTS MADE</p>
                            <p>SALE PRICE: NO WARRANTY ON ANY SPECIAL MARK DOWN PRODUCTS - ALL SPECIAL MARK DOWN PRODUCTS ARE SOLD AS NONAVAILABILITY AND ARE NOT COVERED UNDER WARRANTY - SEE STORE FOR DETIALS</p>
                        </Col>
                    </Row>
                </Container>
            </div>
        );
    }
}