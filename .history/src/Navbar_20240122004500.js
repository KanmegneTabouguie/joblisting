import { Navbar, Nav } from 'react-bootstrap';

const MyNavbar = () => {
 return (
    <Navbar bg="primary" variant="dark" expand="lg">
      <Navbar.Brand href="#home">Remote Jobs</Navbar.Brand>
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
