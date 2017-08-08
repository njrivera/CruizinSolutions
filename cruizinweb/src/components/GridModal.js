import {Modal, ModalBody, ModalFooter, ModalHeader, Form, FormGroup, FormControl} from 'react-bootstrap';
import {Button, Col, Row} from 'reactstrap';
import React from 'react';
import axios from 'axios';

export default class GridModal extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            record : {}
        };
        this.assignValue = this.assignValue.bind(this);
        this.onSave = this.onSave.bind(this);
    }

    onSave() {
        var scope = this;
        if(this.props.action === 'add'){
            axios.post('/customers', this.state.record)
            .then(response => {
                scope.props.setModal();
                scope.props.loadRecords();
            })
            .catch(error => {
                console.log(error);
            });
        }
        else{
            axios.put('/customers/' + this.props.id, this.state.record)
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

    onDelete() {
        var scope = this;
        axios.delete('/customers/' + this.props.id)
        .then(response => {
            scope.props.setModal();
            scope.props.loadRecords();
            scope.props.deleteRecord();
        })
        .catch(error => {
            console.log(error);
        });
    }

    assignValue(event) {
        var temp = JSON.parse(JSON.stringify(this.state.record));
        temp[event.target.id] = event.target.value;
        this.setState({record: temp});
    }

    componentDidUpdate = () => {
        if (this.props.flag) {
            this.state.record = this.props.record;
            this.props.setFlag();
        }
    }

    render() {
        var scope = this;
        switch(this.props.action) {
            case 'add':
            case 'edit':          
                return (
                    <Modal show={this.props.modal} onHide={() => this.props.setModal()}>
                        <ModalBody>
                            <Form horizontal>
                                {Object.keys(scope.props.record).map(function(key){
                                    return <FormGroup controlId={key} key={key}>
                                        <Col sm='2'>{key.charAt(0).toUpperCase() + key.slice(1)}</Col>
                                        <Col sm='10'>
                                            <FormControl 
                                                value={scope.state.record[key] || ''}
                                                onChange={scope.assignValue}>
                                            </FormControl>
                                        </Col>
                                    </FormGroup>
                                })}
                            </Form>
                        </ModalBody>
                        <ModalFooter>
                            <Button onClick={() => this.onSave()}>Save</Button>
                            <Button onClick={() => this.props.setModal()}>Cancel</Button>
                        </ModalFooter>
                    </Modal>
                );
                break;
            case 'delete':
                return (
                    <Modal show={this.props.modal} onHide={() => this.props.setModal()}>
                        <ModalBody>
                            <h1>Are You Sure?</h1>
                        </ModalBody>
                        <ModalFooter>
                            <Button onClick={() => this.onDelete()}>Yes</Button>
                            <Button onClick={() => this.props.setModal()}>No</Button>
                        </ModalFooter>
                    </Modal>
                );
                break;
            default: return (<div></div>);
        }
    }
}