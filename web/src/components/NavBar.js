import React from 'react';
import {Link} from 'react-router-dom';
import {Navbar, NavItem, Nav, NavDropdown, MenuItem} from 'react-bootstrap';
import {LinkContainer} from 'react-router-bootstrap';

export default class NavBar extends React.Component {

    render() {
        return (
            <Navbar inverse collapseOnSelect>
                <Navbar.Header>
                    <Navbar.Brand>
                        <Link to='/'>Home</Link>
                    </Navbar.Brand>
                    <Navbar.Toggle />
                </Navbar.Header>
                <Navbar.Collapse>
                    <Nav>
                        <NavDropdown title='Info' id='infoDropDown'>
                            <LinkContainer to='/info/customers'><MenuItem>Customers</MenuItem></LinkContainer>
                            <LinkContainer to='/info/vehicles'><MenuItem>Vehicles</MenuItem></LinkContainer>
                        </NavDropdown>
                        <NavDropdown title='Inventory' id='inventoryDropDown'>
                            <LinkContainer to='/inventory/tires'><MenuItem>Tires</MenuItem></LinkContainer>
                            <LinkContainer to='/inventory/rims'><MenuItem>Rims</MenuItem></LinkContainer>
                            <LinkContainer to='/inventory/parts'><MenuItem>Parts</MenuItem></LinkContainer>
                            <LinkContainer to='/inventory/services'><MenuItem>Services</MenuItem></LinkContainer>
                            <LinkContainer to='/inventory/packages'><MenuItem>Packages</MenuItem></LinkContainer>
                        </NavDropdown>
                        <LinkContainer to='/past'><NavItem>History</NavItem></LinkContainer>
                        <LinkContainer to='/quote'><NavItem>Quote</NavItem></LinkContainer>
                        <LinkContainer to='/order'><NavItem>Work Order</NavItem></LinkContainer>
                        <LinkContainer to='/report'><NavItem>Report</NavItem></LinkContainer>
                    </Nav>
                </Navbar.Collapse>
            </Navbar>
        );
    }
}
