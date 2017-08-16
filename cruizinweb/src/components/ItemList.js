import React from 'react';
import {BootstrapTable, TableHeaderColumn} from 'react-bootstrap-table';
import GridModal from './GridModal';
import {Container, Row, Col, Button} from 'reactstrap';
import axios from 'axios';

export default class ItemList extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            selected: null
        };
        this.onSelectItems = this.onSelectItems.bind(this);
        this.onAfterSaveCell = this.onAfterSaveCell.bind(this);
    }

    onSelectItems(row, isSelected) {
        if (isSelected) {
            this.setState({selected: row});
        }
        else
            this.setState({selected: null});
    }

    qtyValidator(value) {
        const nan = isNaN(parseInt(value, 10));
        if (nan) {
            return 'Must Be Integer';
        }
        return true;
    }

    onAfterSaveCell(row, cellName, cellValue) {
        this.props.changePrice(row.itemnum, cellValue);
    }

    render() {
        const cellEditProp = {
            mode: 'click',
            blurToSave: true,
            afterSaveCell: this.onAfterSaveCell
        };
        return (
            <div>
                <Container>
                    <Row>
                        <BootstrapTable 
                            data={this.props.items} 
                            maxHeight='500px'
                            scrollTop={'Bottom'} 
                            hover
                            condensed
                            selectRow={{
                                mode: 'radio', 
                                clickToSelect: true, 
                                bgColor: 'black',
                                onSelect: this.onSelectItems
                            }} 
                            containerStyle={{
                                background: '#2F2F2F'
                            }}
                            cellEdit={cellEditProp}>
                            <TableHeaderColumn dataField="itemnum" width='auto' isKey>Item Number</TableHeaderColumn>
                            <TableHeaderColumn dataField="description" width='auto' columnTitle editable={false}>Description</TableHeaderColumn>
                            <TableHeaderColumn dataField="qty" editable={{validator: this.qtyValidator}} width='auto'>Qty</TableHeaderColumn>
                            <TableHeaderColumn dataField="amount" editable={false} width='auto'>Amount</TableHeaderColumn>
                            <TableHeaderColumn dataField="taxable" hidden width='auto'>Taxable</TableHeaderColumn>
                        </BootstrapTable>
                    </Row>
                    <p></p>
                    <Row>
                        <Col>
                            <Button onClick={() => this.props.removeItem(this.state.selected)}>Remove Item</Button>
                            {' '}<Button onClick={() => this.props.finishOrder()}>Invoice</Button>
                        </Col>
                    </Row>
                </Container>
            </div>
        );
    }
}