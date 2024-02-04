import { Nav, Navbar } from 'react-bootstrap';

const MyNavbar = () => {
 return (
    <Navbar bg="primary" variant="dark" expand="lg" collapseOnSelect>
    <Navbar.Brand href="#home">
        <img
          src="https://www.canva.com/design/DAF6m_8lmE4/wL_szF4w8J9RZ5vA6_etxw/view?utm_content=DAF6m_8lmE4&utm_campaign=designshare&utm_medium=link&utm_source=editor"
          width="30"
          height="30"
          className="d-inline-block align-top"
          alt="Remote.Jobs Logo"
        />
        Remote.Jobs
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
