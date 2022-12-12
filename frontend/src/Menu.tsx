import type { Component } from 'solid-js'
import { Link, useNavigate } from 'solid-app-router'
import { Container, Nav, Navbar } from "solid-bootstrap"
import styles from './Menu.module.scss'

const Menu: Component = () => {
  const navigate = useNavigate()
  return (
    
    <div class={styles.Menu}>
      <Container class="pb-5">test</Container>
      <a href="#">TEST</a>
      <Navbar collapseOnSelect expand="lg" bg="dark" variant="dark">
            <Container>
            <Navbar.Brand>
            <Link href='/' class="no-underline color-white">
                Faith
            </Link>            
            </Navbar.Brand>
            <Navbar.Toggle aria-controls="responsive-navbar-nav" />
            <Navbar.Collapse id="responsive-navbar-nav">                
                <Nav>
                <Nav.Link onclick={()=>{navigate(`/test`)}}>Test</Nav.Link>
                </Nav>
            </Navbar.Collapse>
            </Container>
        </Navbar>
      
    </div>
  );
};

export default Menu;
