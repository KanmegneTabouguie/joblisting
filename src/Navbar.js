import { Nav, Navbar } from 'react-bootstrap';
import logo from './Hill..png'; // Import the image

const MyNavbar = () => {
 return (
    <Navbar bg="primary" variant="dark" expand="lg" collapseOnSelect>
      <Navbar.Brand href="#home">
        <img
          src={logo} // Use the imported image
          width="50"
          height="50"
          className="d-inline-block align-top"
          alt="Remote.Jobs Logo"
        />
      </Navbar.Brand>
      <Navbar.Toggle aria-controls="basic-navbar-nav" />
      <Navbar.Collapse id="basic-navbar-nav">
        <Nav className="ms-auto">
        <Nav.Link href="/work" style={{fontSize: '1.2em', color: '#fff'}}>Work</Nav.Link>
          <Nav.Link href="/new" style={{fontSize: '1.2em', color: '#fff'}}>New</Nav.Link>
          <Nav.Link href="/about" style={{fontSize: '1.2em', color: '#fff'}}>About</Nav.Link>
          <Nav.Link href="/career" style={{fontSize: '1.2em', color: '#fff'}}>Career</Nav.Link>
          <Nav.Link href="/contact" style={{fontSize: '1.2em', color: '#fff'}}>Contact</Nav.Link>
        </Nav>
      </Navbar.Collapse>
    </Navbar>
 );
};

export default MyNavbar;
