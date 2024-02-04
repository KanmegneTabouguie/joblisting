import React from 'react';
import { Container, Row, Col } from 'react-bootstrap';

const Footer = () => {
 return (
    <footer className="mt-auto py-3 bg-danger">
      <Container>
        <Row className="justify-content-around">
          <Col xs={12} sm={4} className="text-center">
            <span>Copyright &copy; 2024 Remote Jobs</span>
          </Col>
          <Col xs={12} sm={4} className="text-center">
            <span>Terms of Service</span>
          </Col>
          <Col xs={12} sm={4} className="text-center">
            <span>Privacy Policy</span>
          </Col>
        </Row>
      </Container>
    </footer>
 );
};

export default Footer;
