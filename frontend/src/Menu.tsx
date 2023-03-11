import type { Component } from 'solid-js'
import { Link, useNavigate } from 'solid-app-router'
import { Container, Nav, Navbar } from "solid-bootstrap"
import './Menu.scss'

const Menu: Component = () => {
  const navigate = useNavigate()
  return (
    
    <div id="menu">
      <Navbar collapseOnSelect expand="lg">
            <Container>
            <Navbar.Brand>
            <Link href='/' class="no-underline color-white">
                Blog
            </Link>            
            </Navbar.Brand>
            <Navbar.Toggle aria-controls="responsive-navbar-nav" />
            <Navbar.Collapse id="responsive-navbar-nav">                
                <Nav>
                <Nav.Link onclick={()=>{navigate(`/about`)}}>About</Nav.Link>
                </Nav>
            </Navbar.Collapse>
            </Container>
        </Navbar>
      
    </div>
  );
};

export default Menu;
