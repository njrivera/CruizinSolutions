import {Modal, ModalBody, ModalFooter, Form, FormGroup, FormControl} from 'react-bootstrap';
import {Button, Col} from 'reactstrap';
import React from 'react';
import axios from 'axios';

export default class GridModal extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            record: {},
            editable: {}
        };
    }

    onDelete() {
        var scope = this;
        axios.delete(this.props.url + '/' + this.props.id)
        .then(response => {
            scope.props.setModal();
            scope.props.loadRecords();
            scope.props.deleteRecord();
        })
        .catch(error => {
            scope.props.setModal();
            scope.props.setError(error.response.data);
        });
    }

    componentDidUpdate = () => {
        if (this.props.flag) {
            this.setState({record: this.props.record});
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
                                    if (scope.props.action === 'edit' && !scope.props.editable[key])
                                        return null;
                                    else
                                        return <FormGroup controlId={key} key={key}>
                                            {key === 'condition' ? (
                                                    <div>
                                                        <Col sm='2'>{key.charAt(0).toUpperCase() + key.slice(1)}</Col>
                                                        <Col sm='10'>
                                                            <FormControl 
                                                                componentClass='select' 
                                                                placeholder='USED' 
                                                                onSelect={(event) => scope.props.validateInput(scope, event)}>
                                                                    <option value="NEW">NEW</option>
                                                                    <option value="USED">USED</option>
                                                            </FormControl>
                                                        </Col>
                                                    </div>
                                                ) : (
                                                    <div>
                                                        <Col sm='2'>{key.charAt(0).toUpperCase() + key.slice(1)}</Col>
                                                        <Col sm='10'>
                                                            <FormControl 
                                                                value={scope.state.record[key] || ''}
                                                                onChange={(event) => scope.props.validateInput(scope, event)}>
                                                            </FormControl>
                                                        </Col>
                                                    </div>
                                                )
                                            }
                                        </FormGroup>
                                })}
                            </Form>
                        </ModalBody>
                        <ModalFooter>
                            <Button onClick={() => this.props.onSave(scope)}>Save</Button>
                            <Button onClick={() => this.props.setModal()}>Cancel</Button>
                        </ModalFooter>
                    </Modal>
                );
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
            default: return (<div></div>);
        }
    }
}