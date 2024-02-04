import { Nav, Navbar } from 'react-bootstrap';
import logo from './Hill..png'; // Import the image

const MyNavbar = () => {
 return (
    <Navbar bg="primary" variant="dark" expand="lg" collapseOnSelect>
      <Navbar.Brand href="#home">
        <img
          src={logo} // Use the imported image
          width="30"
          height="30"
          className="d-inline-block align-top"
          alt="Remote.Jobs Logo"
        />
      </Navbar.Brand>
      <Navbar.Toggle aria-controls="basic-navbar-nav" />
      <Navbar.Collapse id="basic-navbar-nav">
        <Nav className="ms-auto">
          <Nav.Link href="/work">Work</Nav.Link>
          <Nav.Link href="/new">New</Nav.Link>
          <Nav.Link href="/about">About</Nav.Link>
          <Nav.Link href="/career">Career</Nav.Link>
          <Nav.Link href="/contact">Contact</Nav.Link>
        </Nav>
      </Navbar.Collapse>
    </Navbar>
 );
};

export default MyNavbar;
