import {Modal, ModalBody, ModalFooter, ModalHeader, Form, FormGroup, FormControl} from 'react-bootstrap';
import {Button, Col, Row} from 'reactstrap';
import React from 'react';
import axios from 'axios';

export default class GridModal extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            inputNodes: {},
            record: {},
            values: {},
            keys: []
        };
        this.assignValue = this.assignValue.bind(this);
    }

    onSave(act, nodes, id) {
        this.state.values = {};
        this.state.keys = [];
        var record = {};
        for(var key in nodes) {
            if (nodes[key].value === '')
                nodes[key].value = nodes[key].placeholder;
            record[nodes[key].id] = nodes[key].value;
        }
        if(act === 'add') {
            axios.post(this.props.url, record)
            .then(response => {
                this.props.toggle('saved');
            })
            .catch(error => {
                return(
                    <Modal>
                        <ModalHeader closeButton></ModalHeader>
                        <ModalBody><h1>Error!</h1></ModalBody>
                    </Modal>
                );
            });
        }
        if(act === 'edit') {
            this.props.onUpdate(record);
            axios.put(this.props.url + '/' + id, record)
            .then(response => {
                this.props.toggle('saved');
            })
            .catch(error => {
                return(
                    <Modal>
                        <ModalHeader closeButton></ModalHeader>
                        <ModalBody><h1>Error!</h1></ModalBody>
                    </Modal>
                );
            });
        }
    }

    onDelete(id) {
        this.state.values = {};
        this.state.keys = [];   
        axios.delete(this.props.url + '/' + id)
            .then(response => {
                this.props.toggle('deleted');
                this.state.values = {};
                this.state.keys = [];   
            })
            .catch(error => {
                return(
                    <Modal>
                        <ModalHeader closeButton></ModalHeader>
                        <ModalBody><h1>Error!</h1></ModalBody>
                    </Modal>
                );
            });
    }

    onCancel() {
        this.state.values = {};
        this.state.keys = [];
    }

    assignValue(event) {
        this.state.inputNodes[event.target.id] = event.target;
        var temp = JSON.parse(JSON.stringify(this.state.values));
        temp[event.target.id] = event.target.value;
        this.setState({values: temp});
    }

    render() {
        var scope = this;
        for(var key in scope.props.data) {
            if (!this.state.values[key]) {
                this.state.values[key] = scope.props.data[key];
            }
            if(this.state.keys.indexOf(key) <= -1) {
                this.state.keys.push(key);
            }
        }
        switch(this.props.action) {
            case 'add':
            case 'edit':          
                return (
                    <Modal show={this.props.modal} onHide={() => {this.props.toggle('done'), this.onCancel()}}>
                        <ModalBody>
                            <Form horizontal>
                                {scope.state.keys.slice(1).map(function(key, i){
                                    return <FormGroup controlId={key} key={key}>
                                        <Col sm='2'>{key.charAt(0).toUpperCase() + key.slice(1)}</Col>
                                        <Col sm='10'>
                                            <FormControl 
                                                inputRef={node => scope.state.inputNodes[key] = node} 
                                                value={scope.state.values[key]}
                                                onChange={scope.assignValue}>
                                            </FormControl>
                                        </Col>
                                    </FormGroup>
                                })}
                            </Form>
                        </ModalBody>
                        <ModalFooter>
                            <Button onClick={() => this.onSave(this.props.action, this.state.inputNodes, this.props.id)}>Save</Button>
                            <Button onClick={() => {this.props.toggle('done'), this.onCancel()}}>Cancel</Button>
                        </ModalFooter>
                    </Modal>
                );
                break;
            case 'delete':
                return (
                <Modal show={this.props.modal} onHide={() => {this.props.toggle('done'), this.onCancel()}}>
                    <ModalBody>
                        <h1>Are You Sure?</h1>
                    </ModalBody>
                    <ModalFooter>
                        <Button onClick={() => this.onDelete(this.props.id)}>Yes</Button>
                        <Button onClick={() => {this.props.toggle('done'), this.onCancel()}}>No</Button>
                    </ModalFooter>
                </Modal>
            );
            break;
            default: return (<div></div>);
        }
    }
}