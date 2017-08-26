import React from 'react';
import {Button, Col, Row, Container} from 'reactstrap';
import {BootstrapTable, TableHeaderColumn} from 'react-bootstrap-table';
import caret from '../styles/fa-caret-right.png';

export default class PrintQuote extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            comments: '',
            orderNum: ''
        };
    }

    render() {
        return (
            <div>
                <Container>
                    <Row>
                        <Col sm='6' className='text-left'>
                            <Row><h3>YOUR ULTIMATE DRIVING EXPERIENCE</h3></Row>
                            <Row><img style={{width: 5, height: 5}} src={caret} alt=''/> NEW {'&'} USED TIRES ALL BRANDS</Row>
                            <Row><img style={{width: 5, height: 5}} src={caret} alt=''/> COMPUTER BALANCING</Row>
                            <Row><img style={{width: 5, height: 5}} src={caret} alt=''/> WHEELS {'&'} TIRES</Row>
                            <Row><img style={{width: 5, height: 5}} src={caret} alt=''/> LAY-AWAY AVAILABLE</Row>
                            <Row><img style={{width: 5, height: 5}} src={caret} alt=''/> CAR ACCESSORIES</Row>
                        </Col>
                        <Col sm='6' className='text-right'>
                            <br/>
                            <h4>6901 SHELDON ROAD | TAMPA, FL 33615</h4>
                            <h4>(813) 886-8072 | FAX: (813) 866-8092</h4>
                            <h4>cruzinsolutions@juno.com</h4>
                        </Col>
                    </Row>
                    <Row><h3>QUOTE</h3></Row>
                    <br/><br/>
                    <Row>Date of Quote: {this.props.date}</Row>
                    <Row>(PRICE VALID ONLY FOR THIS DATE)</Row>
                    <br/><br/>
                    <BootstrapTable 
                        data={this.props.items} 
                        condensed
                        containerStyle={{
                            background: '#2F2F2F'
                        }}>
                        <TableHeaderColumn dataField="itemnum" width='60' isKey>Item #</TableHeaderColumn>
                        <TableHeaderColumn dataField="description" width='auto'>Description</TableHeaderColumn>
                        <TableHeaderColumn dataField="price" width='100'>Price</TableHeaderColumn>
                        <TableHeaderColumn dataField="qty" width='60'>Qty</TableHeaderColumn>
                        <TableHeaderColumn dataField="amount" width='100'>Amount</TableHeaderColumn>
                    </BootstrapTable>
                    <br/><br/>
                    <Row>
                        <p>Subtotal: {this.props.subtotal}</p>
                        <p>Tax: {this.props.tax}</p>
                        <p>Total: {this.props.total}</p>
                    </Row>
                    <br/><br/>
                    <Button color='info' className='hidden-sm' onClick={this.props.onPrint ? () => this.props.onPrint() : () => this.onConfirm(window)}>{this.props.printTitle}</Button>
                    <br/><br/><br/>
                    <Row className='text-left'>
                        <Col sm='6'>
                            <p>- Always inflate, rotate and balance your tires at intervals of 7,000 miles. Improper care can affect the warranty on new tires.</p>
                            <p>- No warranty on low profile tires: 25, 30, 35, 40 series</p>
                            <p>- Customer must re-tighten lug nuts after 25 miles</p>
                            <p>- All wheel sales are final when mounted.</p>
                            <p>- No refund on deposit</p>
                            <p>- No refund on special orders</p>
                        </Col>
                        <Col sm='6'>
                            <p>- No warranty on used tires</p>
                            <p>- 25% Re-Stock fee on all returned goods</p>
                            <p>- No warranty on spacers, adapters or any vehicle alterations</p>
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
                            <p>BALANCING: FREE BALANCING AT TIME OF PURCHASE ONLY UNLESS OTHER ARRANGEMENTS MADE</p>
                        </Col>
                        <Col sm='6'>
                            <p>SALE PRICE: NO WARRANTY ON ANY SPECIAL MARK DOWN PRODUCTS - ALL SPECIAL MARK DOWN PRODUCTS ARE SOLD AS NON-AVAILABILITY AND ARE NOT COVERED UNDER WARRANTY - SEE STORE FOR DETAILS</p>
                        </Col>
                    </Row>
                </Container>
            </div>
        );
    }
}